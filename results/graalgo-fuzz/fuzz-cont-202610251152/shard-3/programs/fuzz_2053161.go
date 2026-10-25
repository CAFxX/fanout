package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 uint }

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 293
    var v3 int8 = -128
    v4 := []int{32767, int(v3)}
    var v5 float32 = 0.1
    var v6 bool = true
    v1["a"] = -96
    switch "\n\t" {
    case "\n\t":
        var v7 uint16 = 42
        _ = v7
    case "a":
        for fi8 := 0; fi8 < 11; fi8++ {
            if v6 {
                v3 -= (^(v3))
                v3 = (v3 & (v3 / ((v3) | 1)))
            } else {
                var v9 int8 = (v3 & (v3 | v3))
                var v10 uint64 = v2
                var v11 uint64 = (v2 + (^(v2)))
                _ = v10
                _ = v11
                _ = v9
            }
            _ = fi8
        }
    case "":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    default:
        if true {
            var v12 int16 = -549
            _ = v12
        } else {
            var v13 int64 = -1
            _ = v13
        }
    }
    v5 = ((float32(0.0) - ((float32(-2.25) - float32((float64(v5) - (-2.25 + 1e10)))) * 0.0)) - float32(-2.25))
    v14 := GenS1{F0: 0, F1: 1, F2: 42}
    var v15 uint16 = 255
    v3 += (v3 >> ((1) & 63))
    v5 -= -2.25
    var v16 uint64 = (v2 & (v2 ^ (v2 & v2)))
    fmt.Println(v3, v6)
    delete(v1, "")
    if ((-1 >= 702) || (v6 && ((42 < 42) && ((!true) || v6)))) {
        v17 := GenS2{F0: v15, F1: 1}
        _ = v17
    } else {
        for fi18 := 0; fi18 < 15; fi18++ {
            v19 := GenS1{F0: 2, F1: 4294967295, F2: -1}
            var v20 int = (fi18 + (^(fi18)))
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    v21 := GenS3{F0: 405}
    switch 980 {
    case 32767:
        v3 -= v3
    }
    switch ("a" + ("a" + "hello")) {
    case "hello":
        fmt.Println(v15)
    case "":
        var v22 uint32 = 1
        _ = v22
    case "foo bar":
        v4 = append(v4, -143)
    default:
        var v23 int16 = -92
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
