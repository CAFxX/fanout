package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := []uint32{1, 610}
    var v2 bool = true
    v3 := make(map[uint32]int)
    var v4 uint = 953
    var v5 uint32 = uint32(v4)
    if v2 {
        for fi6 := 0; fi6 < 19; fi6++ {
            for fi7 := 0; fi7 < 11; fi7++ {
                fmt.Println(v2, v4)
                var v8 int16 = -32768
                _ = fi7
                _ = v8
            }
            var v9 uint32 = (v5 ^ v5)
            _ = fi6
            _ = v9
        }
        fmt.Println(v2, v5)
    }
    var v10 float64 = ((0.1 + 1.5) + float64((float32(0.1) * float32((-2.25 - 1e10)))))
    switch ("" + "\n\t") {
    case "\n\t":
        var v11 uint8 = 42
        _ = v11
    case "":
        if (v10 == v10) {
            fmt.Println(v10)
            fmt.Println(v4, v5, v5)
        }
    case "hello":
        v12 := GenS1{F0: -1}
        _ = v12
    default:
        for fi13 := 0; fi13 < 17; fi13++ {
            delete(v3, v5)
            v10 += v10
            _ = fi13
        }
    }
    var v14 int64 = 0
    if (v2 || false) {
        v4 += (v4 * v4)
    } else {
        var v15 uint32 = (v5 >> ((v4) & 63))
        _ = v15
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        var v17 bool = (220 > uint8(v4))
        v4 -= (v4 + 4294967295)
        v5 += (^(v5))
        _ = fi16
        _ = v17
    }
    v1 = append(v1, (v5 / ((v5) | 1)))
    v4 += v4
    var v18 int8 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
