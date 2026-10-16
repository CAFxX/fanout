package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 uint16 }

func main() {
    var v1 string = "foo bar"
    v2 := make(map[int]uint32)
    v3 := []string{v1, "hello"}
    var v4 int8 = 127
    for fi5 := 0; fi5 < 8; fi5++ {
        delete(v2, (fi5 | fi5))
        delete(v2, (fi5 & fi5))
        _ = fi5
    }
    fmt.Println(v1, v4)
    fmt.Println(v4, v1)
    var v6 float64 = (3.14159 / 0.1)
    var v7 int8 = -128
    fmt.Println(v7, v4, v6)
    switch v1 {
    case "foo bar":
        v4 += (^(v7))
    default:
        delete(v2, -946)
    }
    v8 := GenS1{F0: 515, F1: 0}
    v1 = "foo bar"
    v1 = (v1 + (v1 + "\n\t"))
    if ((1 >= 2) || (!(0 != 483))) {
        v6 = (v6 + (float64(((1.5 / 0.1) - 1.5)) - -2.25))
        v9 := GenS1{F0: 1, F1: 2147483647}
        _ = v9
    } else {
        v10 := GenS2{F0: -1}
        _ = v10
    }
    fmt.Println(v7, v4)
    v11 := GenS3{F0: 23, F1: 11}
    var v12 uint16 = 46
    var v13 string = v1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
