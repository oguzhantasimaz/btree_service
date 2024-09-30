package main

import (
	"bytes"
	"encoding/json"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	domain "github.com/oguzhantasimaz/btree_service/models"
	"github.com/stretchr/testify/assert"
)

func TestMaxPathSum(t *testing.T) {
	// Initialize Echo
	e := echo.New()

	// Test cases
	tests := []struct {
		name           string
		input          Request
		expectedStatus int
		expectedSum    int
	}{
		{
			name: "Test Case 1 - Regular Binary domain.Tree",
			input: Request{
				Tree: domain.Tree{
					Nodes: []domain.Node{
						{ID: "1", Left: stringPointer("2"), Right: stringPointer("3"), Value: 1},
						{ID: "2", Left: stringPointer("4"), Right: stringPointer("5"), Value: 2},
						{ID: "3", Left: stringPointer("6"), Right: stringPointer("7"), Value: 3},
						{ID: "4", Left: nil, Right: nil, Value: 4},
						{ID: "5", Left: nil, Right: nil, Value: 5},
						{ID: "6", Left: nil, Right: nil, Value: 6},
						{ID: "7", Left: nil, Right: nil, Value: 7},
					},
					Root: "1",
				},
			},
			expectedStatus: http.StatusOK,
			expectedSum:    18, // 5 + 2 + 1 + 3 + 7
		},
		{
			name: "Test Case 2 - Smaller Binary domain.Tree",
			input: Request{
				Tree: domain.Tree{
					Nodes: []domain.Node{
						{ID: "1", Left: stringPointer("2"), Right: stringPointer("3"), Value: 1},
						{ID: "2", Left: nil, Right: nil, Value: 2},
						{ID: "3", Left: nil, Right: nil, Value: 3},
					},
					Root: "1",
				},
			},
			expectedStatus: http.StatusOK,
			expectedSum:    6,
		},
		{
			name: "Test Case 3 - domain.Tree with Negative Values",
			input: Request{
				Tree: domain.Tree{
					Nodes: []domain.Node{
						{ID: "1", Left: stringPointer("-10"), Right: stringPointer("-5"), Value: 1},
						{ID: "-10", Left: stringPointer("30"), Right: stringPointer("45"), Value: -10},
						{ID: "-5", Left: stringPointer("-20"), Right: stringPointer("-21"), Value: -5},
						{ID: "30", Left: nil, Right: nil, Value: 30},
						{ID: "45", Left: nil, Right: nil, Value: 45},
						{ID: "-20", Left: nil, Right: nil, Value: -20},
						{ID: "-21", Left: nil, Right: nil, Value: -21},
					},
					Root: "1",
				},
			},
			expectedStatus: http.StatusOK,
			expectedSum:    65,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Prepare the request
			reqBody, _ := json.Marshal(test.input)
			req := httptest.NewRequest(http.MethodPost, "/max-path-sum", bytes.NewReader(reqBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			// Handler function
			if assert.NoError(t, calculateMaxPathSumHandler(c)) {
				assert.Equal(t, test.expectedStatus, rec.Code)

				if rec.Code == http.StatusOK {
					var response Response
					if err := json.Unmarshal(rec.Body.Bytes(), &response); err == nil {
						assert.Equal(t, test.expectedSum, response.MaxPathSum)
					} else {
						t.Errorf("Failed to unmarshal response")
					}
				}
			}
		})
	}
}

// stringPointer is a helper function to convert string to a pointer
func stringPointer(s string) *string {
	return &s
}

// calculateMaxPathSumHandler is the handler function for the POST /max-path-sum
func calculateMaxPathSumHandler(c echo.Context) error {
	var req Request
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Input")
	}

	nodesMap := buildNodesMap(req.Tree.Nodes)
	rootNode := nodesMap[req.Tree.Root]

	maxSum := math.MinInt32
	calculateMaxPathSum(nodesMap, rootNode, &maxSum)
	return c.JSON(http.StatusOK, Response{MaxPathSum: maxSum})
}
