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
	Hardness int       `json:"mohs hardness rating"`
	Color 	 string    `json:"color"`
}

var (
	gems = map[int]*Gemstone{}
	idCount = 1
	lock = sync.Mutex{}

	// Test data will delete later
	diamond = &Gemstone{
		ID: idCount,
		Name: "Diamond",
		Hardness: 10,
		Color: "Clear",
	}
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
