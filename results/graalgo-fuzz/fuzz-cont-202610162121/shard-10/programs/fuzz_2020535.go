package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool }

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{1}
    var v3 int16 = 1000
    var v4 uint = 445
    var v5 int = 42
    v6 := GenS1{F0: 971}
    var v7 int32 = 42
    if (!false) {
        v4 -= v4
    } else {
        var v8 uint8 = 1
        _ = v8
    }
    switch "a" {
    case "hello":
        v3 = (v3 / ((-404) | 1))
    case "foo bar":
        fmt.Println(v7, v3, v7)
    case "a":
        if false {
            v5 = -2147483648
            v4 += v4
            _ = len(v1)
        }
    }
    var v9 uint16 = 108
    delete(v1, (v5 % (((-(v5))) | 1)))
    if false {
        var v10 int16 = v3
        _ = v10
    }
    v11 := v1[v5]
    v12 := GenS2{F0: (!(!((!(!false)) || (v3 != (v3 + v3)))))}
    var v13 float64 = (((3.14159 / 0.1) - ((((-2.25 * 1.5) + 1e10) / 1.5) / 1.5)) + 0.1)
    _ = len(v2)
    var v14 uint16 = (^(v9))
    v7 = v7
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
