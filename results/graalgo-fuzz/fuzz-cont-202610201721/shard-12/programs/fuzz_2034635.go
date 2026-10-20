package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 int32 = 0
    var v2 int32 = v1
    v3 := []int{}
    v1 = v1
    for fi4 := 0; fi4 < 16; fi4++ {
        var v5 string = (("" + "foo bar") + ("\n\t" + "foo bar"))
        _ = fi4
        _ = v5
    }
    var v6 float64 = ((1.5 * 1.5) / 1.0)
    var v7 int8 = 0
    v7 = (v7 / ((v7) | 1))
    _ = len(v3)
    var v8 int8 = 100
    var v9 uint = 399
    if false {
        for fi10 := 0; fi10 < 1; fi10++ {
            v1 = -1
            v3 = append(v3, 1)
            _ = fi10
        }
    } else {
        v11 := GenS1{F0: int64(v1)}
        _ = v11
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v12 float64 = (-2.25 - 0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
