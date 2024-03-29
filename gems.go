package main

import (
	"net/http"
	"sync"
	"strconv"

	// Echo
	"github.com/labstack/echo/v4"
)

type Gemstone struct {
	ID 		 int 	   `json:"id"`
	Name	 string    `json:"name"`
	Hardness int       `json:"hardness"`
	Color 	 string    `json:"color"`
}

var (
	gems = map[int]*Gemstone{}
	idCount = 1
	lock = sync.Mutex{}
);

/*
 * List all gems GET
 */
func listGems(c echo.Context) error {
	lock.Lock();
	defer lock.Unlock();
	return c.JSON(http.StatusOK, gems)
}

/*
 * List particular gem GET
 */
func listGem(c echo.Context) error {
	lock.Lock();
	defer lock.Unlock();
	id, _ := strconv.Atoi(c.Param("id"));
	return c.JSON(http.StatusOK, gems[id]);
}

/*
 * Create gem POST
 */

func createGem(c echo.Context) error {
	lock.Lock();
	defer lock.Unlock();

	g := &Gemstone{
		ID: idCount,
	}

	if err := c.Bind(g); err != nil { return err }
	gems[g.ID] = g;
	idCount++;

	return c.JSON(http.StatusOK, map[string]string{ "Message": "Gemstone added"})
}

/*
 * Update gem PUT
 */
func updateGem(c echo.Context) error {
	lock.Lock();
	defer lock.Unlock();

	g := new(Gemstone);
	if err := c.Bind(g); err != nil { return err }
	id, _ := strconv.Atoi(c.Param("id"));

	gems[id].Name = g.Name;
	gems[id].Hardness = g.Hardness;
	gems[id].Color = g.Color;

	return c.JSON(http.StatusOK, map[string]string{ "Message": "Gemstone updated"});
}

/*
 * Delete gem DELETE
 */
func deleteGem(c echo.Context) error {
	lock.Lock();
	defer lock.Unlock();
	id, _ := strconv.Atoi(c.Param("id"));
	delete(gems, id);
	return c.NoContent(http.StatusNoContent);
}
