package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }

func main() {
    v1 := make(map[string]int)
    var v2 int = 128
    var v3 uint16 = 42
    v4 := make(map[string]string)
    if (2 <= (v3 << ((0) & 63))) {
        for fi5 := 0; fi5 < 4; fi5++ {
            for fi6 := 0; fi6 < 11; fi6++ {
                fmt.Println(fi6, fi6)
                var v7 int = v2
                _ = fi6
                _ = v7
            }
            if false {
                if (699 != 42) {
                    if true {
                        v8 := GenS1{F0: 255, F1: 127}
                        delete(v1, ("a" + ""))
                        _ = v8
                    }
                }
            }
            fmt.Println(v2)
            _ = fi5
        }
        var v9 float32 = float32(-2.25)
        _ = v9
    } else {
        for fi10 := 0; fi10 < 19; fi10++ {
            v2 = v2
            var v11 uint64 = 18446744073709551615
            _ = fi10
            _ = v11
        }
        fmt.Println(v2, v3)
    }
    v3 -= v3
    delete(v4, "\n\t")
    var v12 int32 = 0
    var v13 float32 = ((-2.25 + 0.1) * -2.25)
    var v14 uint64 = 1
    switch v2 {
    case 1:
        fmt.Println(v3, v2, v2)
    case 128:
        v2 -= (v2 * v2)
    default:
        v13 -= float32(((0.1 / 0.1) / (float64((v13 / ((1.5 + (v13 + v13)) + v13))) / 1.5)))
    }
    var v15 int64 = -397
    var v16 uint64 = v14
    v17 := []int{}
    var v18 uint = 0
    fmt.Println(v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
