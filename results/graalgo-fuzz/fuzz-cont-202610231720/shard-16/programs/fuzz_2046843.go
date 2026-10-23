package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint16 }

func main() {
    var v1 uint = 255
    v2 := []uint32{65535, 2147483647}
    v3 := []string{"\n\t"}
    var v4 uint8 = 0
    var v5 int64 = int64(v4)
    var v6 string = (("" + "\n\t") + "a")
    var v7 float64 = (-2.25 * 1.5)
    _ = len(v3)
    v8 := GenS1{F0: 18446744073709551615, F1: -2147483648, F2: 42}
    switch 360 {
    case -2147483648:
        v6 = ""
    case 128:
        switch 441 {
        case -2147483648:
            var v9 int8 = 42
            _ = v9
        case -137:
            v10 := GenS1{F0: 2, F1: int32(v4), F2: 800}
            _ = v10
        }
    case 0:
        var v11 int16 = 1000
        _ = v11
    }
    fmt.Println(v1, v1)
    v4 -= v4
    var v12 string = v6
    fmt.Println(v5, v4, v6)
    var v13 float64 = (float64((1e10 * 1.5)) * ((v7 * 1e10) * v7))
    v5 -= v5
    v3 = append(v3, "")
    v3 = append(v3, "a")
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
