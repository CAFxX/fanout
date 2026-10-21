package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool }

func main() {
    var v1 int = -367
    v2 := make(map[int]uint32)
    var v3 uint = 18
    v4 := []int{v1, (-(v1))}
    if false {
        var v5 bool = (1 < 2147483647)
        _ = v5
    }
    var v6 int32 = -2147483648
    var v7 int16 = 752
    v7 = v7
    _ = len(v4)
    var v8 float64 = 3.14159
    fmt.Println(v7, v8)
    v1 += (v1 ^ v1)
    var v9 uint32 = 1
    for fi10 := 0; fi10 < 19; fi10++ {
        for fi11 := 0; fi11 < 15; fi11++ {
            if true {
                v9 = v9
                fmt.Println(v8, v9, v9)
                v12 := GenS1{F0: 0}
                _ = v12
            }
            switch "" {
            case "foo bar":
                var v13 string = ("" + "a")
                _ = v13
            }
            _ = fi11
        }
        v7 = v7
        _ = len(v4)
        _ = fi10
    }
    v14 := GenS2{F0: (-83 <= 42)}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
