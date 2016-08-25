package main

import "errors"

type TableTop struct {
  Width, Length int
}

func NewTableTop(width, length int) (*TableTop, error) {
  if width >= 1 && length >= 1 {
    t := new(TableTop)
    
    t.Width = width
    t.Length = length
    
    return t, nil
  } else {
    return nil,  errors.New("foobar")
  }
}