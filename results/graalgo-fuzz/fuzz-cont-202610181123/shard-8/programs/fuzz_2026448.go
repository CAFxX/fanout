package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0() uint64 {
    v6 := GenS1{F0: 163}
    var v7 float64 = 1.5
    var v8 uint64 = 545
    v9 := GenS1{F0: -32768}
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 914
}
func genFn1(p0 int, p1 uint16, p2 uint8) uint8 {
    p2 = (^(p2))
    var v10 uint32 = 19
    fmt.Println(v10, p1)
    if (!true) {
        v11 := []int{-1}
        _ = v11
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        v10 = (v10 | v10)
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 42
}

func main() {
    var v1 uint16 = 416
    var v2 uint64 = 0
    var v3 uint32 = 0
    var v4 float32 = (((1e10 + 1.5) + (0.1 + 3.14159)) * ((3.14159 / ((0.0 + (((-2.25 - ((1.5 / 0.1) * -2.25)) - 0.0) / (-2.25 * 3.14159))) / (0.0 - (((1.5 / 1.0) + 0.1) * (0.1 * 0.1))))) * (-2.25 - ((((((0.0 * 1e10) * 3.14159) - float32(3.14159)) / 1.5) + 3.14159) - (-2.25 - ((0.1 / 1.0) * 0.1))))))
    v5 := make(map[string]bool)
    fmt.Println(v3)
    if true {
        var v13 uint16 = v1
        _ = v13
    }
    _ = len(v5)
    if (!false) {
        _ = len(v5)
        switch ("a" + ("\n\t" + "\n\t")) {
        case "\n\t":
            var v14 float64 = (float64(1.5) * 1.5)
            _ = v14
        case "":
            fmt.Println(v1, v3, v4)
        case "hello":
            v3 += v3
        }
    }
    var v15 uint8 = 2
    v16 := []int{}
    var v17 uint16 = (^(v1))
    var v18 float32 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
