package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type UserTable_20220713_164011 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &UserTable_20220713_164011{}
	m.Created = "20220713_164011"

	migration.Register("UserTable_20220713_164011", m)
}

// Run the migrations
func (m *UserTable_20220713_164011) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE user(id INT AUTO_INCREMENT PRIMARY KEY, username VARCHAR(255), password VARCHAR(255), email VARCHAR(255))")

}

// Reverse the migrations
func (m *UserTable_20220713_164011) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE user")

}
