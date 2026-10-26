package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 string }

func main() {
    var v1 float64 = (-2.25 / -2.25)
    v2 := []uint32{65535, 2147483647}
    v3 := []uint32{384, 2147483648, 2}
    v4 := []int{424, -1}
    v5 := make(map[int]int)
    switch "foo bar" {
    case "foo bar":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "\n\t":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    _ = len(v5)
    v1 = (v1 - v1)
    var v6 float32 = (3.14159 - (-2.25 - -2.25))
    v1 -= float64(v6)
    switch -1 {
    case -2147483648:
        v1 += v1
    }
    var v7 uint16 = 384
    var v8 string = ""
    switch "hello" {
    case "hello":
        v2 = append(v2, 0)
    case "foo bar":
        fmt.Println(v1, v6, v1)
    }
    fmt.Println(v6, v7, v8)
    v7 -= v7
    v3 = append(v3, 492)
    if true {
        var v9 bool = false
        v10 := GenS1{F0: 548, F1: (!(v9 && (-1000 == 391))), F2: v8}
        _ = v10
        _ = v9
    } else {
        for fi11 := 0; fi11 < 5; fi11++ {
            var v12 uint8 = 2
            _ = fi11
            _ = v12
        }
    }
    v13 := v5[0]
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
