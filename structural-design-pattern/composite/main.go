package main

import "fmt"

type IComposite interface {
	perform()
}

type Leaflet struct {
	name string
}

func (leaf *Leaflet) perform() {
	fmt.Println("Leaflet " + leaf.name)
}

type Branch struct {
	leafs    []Leaflet
	name     string
	branches []Branch
}

func (branch *Branch) perform() {
	fmt.Println("Branch " + branch.name)

	for _, leaf := range branch.leafs {
		leaf.perform()
	}

	fmt.Println("Run Branch :")
	for _, branch := range branch.branches {
		branch.perform()
	}
}

func (branch *Branch) add(leaf Leaflet) {
	branch.leafs = append(branch.leafs, leaf)
}

func (branch *Branch) addBranch(newBranch Branch) {
	branch.branches = append(branch.branches, newBranch)
}

func (branch *Branch) getLeaflets() []Leaflet {
	return branch.leafs
}

func main() {
	var branch = &Branch{name: "Branch 1"}
	var leaf1 = Leaflet{name: "Leaf 1"}
	var leaf2 = Leaflet{name: "Leaf 2"}
	var leaf3 = Leaflet{name: "Leaf 3"}
	var branch2 = Branch{name: "Branch 2"}

	branch.add(leaf1)
	branch.add(leaf2)
	branch2.add(leaf3)
	branch.addBranch(branch2)

	branch.perform()
}
