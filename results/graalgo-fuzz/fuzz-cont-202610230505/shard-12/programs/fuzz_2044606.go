package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int }
type GenS2 struct { F0 int; F1 string }
func genFn0(p0 uint16, p1 uint8, p2 int) uint {
    if (4294967295 == 1) {
        var v6 string = "a"
        for fi7 := 0; fi7 < 18; fi7++ {
            p2 = 1
            _ = fi7
        }
        _ = v6
    }
    if (false || (p1 > 127)) {
        var v8 uint32 = 2
        _ = v8
    } else {
        v9 := GenS1{F0: 790, F1: (^(p0)), F2: 32767}
        fmt.Println(p0)
        var v10 uint16 = p0
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return 395
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{}
    var v3 float32 = 1.5
    v4 := make(map[int]string)
    v5 := make(map[uint32]uint32)
    for fi11 := 0; fi11 < 14; fi11++ {
        for fi12 := 0; fi12 < 18; fi12++ {
            v3 += v3
            v3 = 0.0
            _ = fi12
        }
        _ = fi11
    }
    if ((3.14159 - 1.5) < (0.1 - ((1.5 / 1.5) / 0.1))) {
        if (42 == 0) {
            fmt.Println(v3)
            _ = len(v5)
            v3 = (float32(-2.25) / v3)
        }
        var v13 uint64 = 188
        _ = v13
    }
    v14 := GenS2{F0: -124, F1: "hello"}
    _ = len(v2)
    delete(v4, -825)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v15 float64 = 0.1
    var v16 uint32 = 2
    switch ("" + "\n\t") {
    case "a":
        delete(v1, 42)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
