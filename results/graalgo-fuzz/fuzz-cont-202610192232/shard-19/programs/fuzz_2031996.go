package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }

func main() {
    v1 := []string{("foo bar" + "foo bar"), "hello", (("a" + "") + ("\n\t" + "a")), "hello"}
    var v2 float32 = 0.0
    v3 := make(map[uint32]int)
    v2 = v2
    v2 = v2
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    for fi4 := 0; fi4 < 10; fi4++ {
        var v5 uint32 = 1
        _ = fi4
        _ = v5
    }
    v2 = (v2 / 1e10)
    for fi6 := 0; fi6 < 19; fi6++ {
        v2 = v2
        v2 = float32(-2.25)
        _ = fi6
    }
    v7 := GenS1{F0: (!false), F1: 239}
    v8 := GenS1{F0: (false && false), F1: 42}
    var v9 uint = 949
    v10 := GenS1{F0: true, F1: 255}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
