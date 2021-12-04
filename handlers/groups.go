package handlers

import (
	"context"
	"fiber_api_eth/models"
	"log"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gofiber/fiber/v2"
	"github.com/holiman/uint256"
)

// ResponseHTTP represents response body of this API
type ResponseHTTP struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// GetAllGroups is a function to get all groups
// @Summary Get all groups
// @Description Get all groups
// @Tags groups
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=[]models.Group}
// @Failure 503 {object} ResponseHTTP{}
// @Router /api/v1/groups [get]
func (app *App) GetAllGroups(c *fiber.Ctx) error {

	groupIds, err := app.contractInstance.GetGroupIds(nil)
	if err != nil {
		return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	groups := make([]models.Group, 0, len(groupIds))

	for _, groupId := range groupIds {
		str, e := app.contractInstance.GetGroup(nil, groupId)
		if e != nil {
			continue
		}
		u, _ := uint256.FromBig(groupId)
		log.Println("groupID:", u.Hex())
		z := models.Group{
			Name:    str.Name,
			Indexes: str.Indexes,
		}
		groups = append(groups, z)
	}

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success get all groups.",
		Data:    groups,
	})
}

// GetGroupByID is a function to get a Group by ID
// @Summary Get Group by ID
// @Description Get Group by ID
// @Tags Groups
// @Accept json
// @Produce json
// @Param groupId path string true "Group ID"
// @Success 200 {object} ResponseHTTP{data=[]models.Group}
// @Failure 404 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /api/v1/groups/{groupId} [get]
func (app *App) GetGroupByID(c *fiber.Ctx) error {
	id := c.Params("groupId")

	log.Println(id)
	v, err := uint256.FromHex(id)
	if err != nil {
		return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	z := v.ToBig()

	log.Println(v.Hex())

	group, err := app.contractInstance.GetGroup(nil, z)
	if err != nil {
		return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success get group by ID.",
		Data:    group,
	})
}

// GetIndexByID is a function to get a Index by ID
// @Summary Get Index by ID
// @Description Get Index by ID
// @Tags Indexes
// @Accept json
// @Produce json
// @Param indexId path string true "Index ID"
// @Success 200 {object} ResponseHTTP{data=[]models.Index}
// @Failure 404 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /api/v1/indexes/{indexId} [get]
func (app *App) GetIndexByID(c *fiber.Ctx) error {

	id := c.Params("indexId")

	log.Println(id)

	v, err := uint256.FromHex(id)
	if err != nil {
		return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	z := v.ToBig()

	log.Println(v.Hex())

	index, err := app.contractInstance.GetIndex(nil, z)
	if err != nil {
		return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	res := models.Index{
		Id:                z,
		Name:              index.Name,
		EthPriceInWei:     index.EthPriceInWei,
		UsdPriceInCents:   index.UsdPriceInCents,
		UsdCapitalization: index.UsdCapitalization,
		PercentageChange:  index.PercentageChange,
	}

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success get index by ID.",
		Data:    res,
	})
}

// GetIndexByID is a function to get a Index by ID
// @Summary Get Index by ID
// @Description Get Index by ID
// @Tags Indexes
// @Accept json
// @Produce json
// @Param block path string true "BlockNumber|BlockHash|'latest'"
// @Success 200 {object} ResponseHTTP{data=[]types.Header}
// @Failure 404 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /api/v1/blocks/{block} [get]
func (app *App) GetBlockInfo(c *fiber.Ctx) error {

	id := c.Params("block")

	var err error

	cx := context.Background()
	var blockInfo *types.Block
	if id == "latest" {
		blockNumber, err := app.client.BlockNumber(cx)
		if err != nil {
			return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
				Success: false,
				Message: "Unable to fetch last block number",
				Data:    nil,
			})
		}
		z := uint256.NewInt(blockNumber)
		blockBig := z.ToBig()
		blockInfo, err = app.client.BlockByNumber(cx, blockBig)
		if err != nil {
			return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
				Success: false,
				Message: "Unable to fetch last block number",
				Data:    nil,
			})
		}
	} else {
		if len(id) > 30 {
			// block hash
			hash := common.HexToHash(id)
			blockInfo, err = app.client.BlockByHash(cx, hash)
			if err != nil {
				return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
					Success: false,
					Message: "Unable to fetch block by hash",
					Data:    nil,
				})
			}
		} else {
			// block number
			z, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
					Success: false,
					Message: "Unable to parse block number",
					Data:    nil,
				})
			}
			blockBig := big.NewInt(z)
			blockInfo, err = app.client.BlockByNumber(cx, blockBig)
			if err != nil {
				return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
					Success: false,
					Message: "Unable to fetch block by number",
					Data:    nil,
				})
			}
		}
	}

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Block info",
		Data:    blockInfo.Header(),
	})
}
