package main

import (
    "testing"
)

func TestMakeGreeting(t *testing.T){
    want :="Hello, haru0121!"
    got := makeGreeting("haru0121")
    if got != want {
        t.Errorf("want %s, but got %s", want, got)
    }
}