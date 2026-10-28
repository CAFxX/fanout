package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int8 }
func genFn0(p0 string, p1 uint16, p2 int16) uint32 {
    v4 := GenS1{F0: 686, F1: (p0 + p0), F2: 65}
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 173
}

func main() {
    var v1 uint8 = 128
    var v2 uint8 = (^(v1))
    v3 := make(map[int]uint32)
    switch ("hello" + "foo bar") {
    case "a":
        for fi5 := 0; fi5 < 15; fi5++ {
            var v6 float32 = 1.5
            var v7 uint32 = 42
            _ = fi5
            _ = v6
            _ = v7
        }
    case "":
        var v8 float64 = -2.25
        _ = v8
    }
    var v9 float32 = float32(1.5)
    if ((!false) && false) {
        for fi10 := 0; fi10 < 12; fi10++ {
            v9 = v9
            v9 = v9
            fmt.Println(fi10, v1, v2)
            _ = fi10
        }
    } else {
        v9 = float32((float64((v9 * v9)) / 1e10))
        v1 += (v1 * (v1 | (^(v1))))
        v2 += (v1 - v2)
    }
    v11 := []int{42, 1}
    for fi12 := 0; fi12 < 1; fi12++ {
        fmt.Println(v9, v1, v2)
        _ = fi12
    }
    delete(v3, 819)
    v9 = v9
    var v13 int = 32767
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 float32 = 0.0
        _ = fi14
        _ = v15
    }
    var v16 float64 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v9
}
