package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
  c "github.com/smartystreets/goconvey/convey"
)

type Point struct {
    x int
    y int
}

type IntList struct {
  val int
  tail *IntList
}

func TestSomethingWithTestify(t *testing.T) {
  assert.Equal(t, &Point{x:10, y:2}, &Point{x:1, y:3})
}

func TestIntListWithTestify(t *testing.T) {
  assert.Equal(t, &IntList{val:10}, &IntList{val: 10, tail: &IntList{val: 42}})
}

func TestSomethingWithConvey(t *testing.T) {
  c.Convey("first", t, func () {
    c.Convey("point", func () {
      c.So(&Point{x:10, y:2}, c.ShouldEqual, &Point{x:1, y:3})  
    })
    c.Convey("point2", func () {
      c.So(&IntList{val:10}, c.ShouldEqual, &IntList{val: 10, tail: &IntList{val: 42}})  
    })
    
  })
}
