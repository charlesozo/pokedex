package main

import (
	"fmt"
	"log"


)

func commandMap(cfg *config) error {
	
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Location areas:")
   for _, area := range resp.Results{
	fmt.Printf(" -%s\n", area.Name)
   }
   cfg.nextLocationAreaURL = resp.Next
   cfg.previousLocationAreaURL = resp.Previous
   return nil	
}

func commandMapb(cfg *config) error {
	
	resp, err := cfg.pokeapiClient.ListLocationAreas( cfg.previousLocationAreaURL)
	if err != nil{
		return err
	}
	fmt.Println("Location areas:")
   for _, area := range resp.Results{
	fmt.Printf(" -%s\n", area.Name)
   }
   cfg.nextLocationAreaURL = resp.Next
   cfg.previousLocationAreaURL = resp.Previous
   return nil	
}
