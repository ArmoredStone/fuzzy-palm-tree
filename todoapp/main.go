/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/

// TODO: Wrap the file reading and wrting logic into functions
// TODO: a separate package for TodoEntity, import it
// TODO: implement logging into separate files
// TODO: use homedir library and root.go init to store persistent values
// TODO: add priority variable for TodoEntiies
// TODO: add tag variable for TodoEntities
// TODO: enhance list command for various output options
// TODO: check tabwriter library
// TODO: implement sort.Interface for TodoEntity
// TODO: VIPER config manager good with cobra
// TODO: add search
// TODO: https://go.dev/doc/articles/wiki/

package main

import "github.com/ArmoredStone/fuzzy-palm-tree/todoapp/cmd"

func main() {
	cmd.Execute()
}
