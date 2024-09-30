package main

import (
	"context"
	"math"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	domain "github.com/oguzhantasimaz/btree_service/models"
)

type Request struct {
	Tree domain.Tree `json:"tree"`
}

type Response struct {
	MaxPathSum int `json:"maxPathSum"`
}

func main() {
	e := echo.New()

	e.Use(middleware.CORS())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// Endpoint Handler
	e.POST("/max_path_sum", func(c echo.Context) error {
		var req Request
		if err := c.Bind(&req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid Input")
		}

		nodesMap := buildNodesMap(req.Tree.Nodes)
		rootNode := nodesMap[req.Tree.Root]

		maxSum := math.MinInt32
		calculateMaxPathSum(nodesMap, rootNode, &maxSum)
		return c.JSON(http.StatusOK, Response{MaxPathSum: maxSum})
	})

	// Graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	go func() {
		if err := e.Start(":3000"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

// Calculation Functions

func buildNodesMap(nodes []domain.Node) map[string]*domain.Node {
	nodeMap := make(map[string]*domain.Node)
	for _, node := range nodes {
		nodeCopy := node
		nodeMap[node.ID] = &nodeCopy
	}
	return nodeMap
}

func calculateMaxPathSum(nodes map[string]*domain.Node, node *domain.Node, maxSum *int) int {
	if node == nil {
		return 0
	}

	leftMax := 0
	if node.Left != nil {
		leftMax = calculateMaxPathSum(nodes, nodes[*node.Left], maxSum)
	}

	rightMax := 0
	if node.Right != nil {
		rightMax = calculateMaxPathSum(nodes, nodes[*node.Right], maxSum)
	}

	leftMax = max(leftMax, 0)
	rightMax = max(rightMax, 0)

	currentMaxPathSum := leftMax + rightMax + node.Value
	*maxSum = max(*maxSum, currentMaxPathSum)

	return node.Value + max(leftMax, rightMax)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
