package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 int64 = 1
    var v2 int64 = v1
    var v3 float32 = float32(-2.25)
    var v4 bool = true
    v5 := make(map[string]uint32)
    switch -1 {
    case 127:
        var v6 int = 4294967295
        _ = v6
    case -1:
        v7 := GenS1{F0: 2}
        _ = v7
    default:
        fmt.Println(v4)
    }
    v2 = v2
    var v8 uint = 255
    if v4 {
        v4 = (!v4)
        v9 := v5[(((("\n\t" + "foo bar") + "a") + "\n\t") + ("" + "hello"))]
        _ = v9
    } else {
        switch "foo bar" {
        case "foo bar":
            v1 += (^(v1))
        case "\n\t":
            var v10 float64 = ((1.5 + (1.5 + -2.25)) + 3.14159)
            _ = v10
        case "a":
            var v11 uint8 = 91
            _ = v11
        default:
            v8 += (^(v8))
        }
        switch "foo bar" {
        case "\n\t":
            switch 32768 {
            case 32768:
                var v12 float32 = v3
                _ = v12
            case 4294967295:
                v13 := []int{42, 128, -328, -546}
                _ = v13
            }
        }
        v8 = (^(v8))
    }
    var v14 int32 = -829
    switch "hello" {
    case "":
        fmt.Println(v4)
    case "\n\t":
        switch 0 {
        case 0:
            var v15 float32 = v3
            _ = v15
        default:
            v16 := v5[(("a" + "hello") + "a")]
            _ = v16
        }
    case "a":
        v2 -= -1
    default:
        v17 := GenS1{F0: 2147483648}
        _ = v17
    }
    fmt.Println(v8, v4, v14)
    fmt.Println(v14, v8)
    var v18 int16 = -641
    v19 := []int{-214, 790, 32768, 0}
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
