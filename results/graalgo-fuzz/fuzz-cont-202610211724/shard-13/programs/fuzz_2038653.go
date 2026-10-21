package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]int)
    v3 := []string{"\n\t"}
    var v4 uint8 = 1
    var v5 float64 = 0.0
    var v6 uint8 = v4
    v6 = (v4 | v6)
    if false {
        v6 -= v6
    }
    v6 = (^(v6))
    if (!true) {
        var v7 int8 = -1
        _ = v7
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 int64 = 0
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 8; fi10++ {
        v11 := GenS1{F0: (fi10 * fi10)}
        v4 -= v4
        for fi12 := 0; fi12 < 20; fi12++ {
            if (2147483647 > -2147483648) {
                v13 := GenS1{F0: -737}
                _ = v13
            } else {
                v5 = v5
            }
            _ = fi12
        }
        _ = fi10
        _ = v11
    }
    var v14 string = "a"
    var v15 int32 = 2147483647
    delete(v2, -989)
    if true {
        var v16 float32 = float32(-2.25)
        var v17 int16 = -136
        _ = v16
        _ = v17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
