package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 uint16 = 2
    v2 := make(map[uint32]bool)
    var v3 bool = (!true)
    v4 := make(map[int]bool)
    var v5 uint32 = 372
    var v6 uint8 = 42
    fmt.Println(v5, v1)
    switch "" {
    case "foo bar":
        if v3 {
            v7 := v2[(v5 << ((865) & 63))]
            var v8 int32 = -546
            _ = v7
            _ = v8
        }
    case "":
        v9 := GenS1{F0: 0}
        _ = v9
    case "hello":
        v1 -= (v1 * (v1 % ((120) | 1)))
    }
    v10 := []int{952, 383, -370, -1}
    v11 := v4[-1]
    var v12 float64 = ((float64(1.5) * 1e10) + float64((-2.25 * 0.1)))
    var v13 uint8 = (v6 << ((196) & 63))
    var v14 float32 = float32(1e10)
    for fi15 := 0; fi15 < 18; fi15++ {
        var v16 uint = 958
        if (!(-79 < -59)) {
            if (2147483647 >= fi15) {
                v17 := GenS2{F0: uint16(v16)}
                _ = v17
            } else {
                v18 := GenS1{F0: -9223372036854775808}
                fmt.Println(v3, fi15)
                _ = v18
            }
            fmt.Println(v5, fi15, v13)
            v10 = append(v10, fi15)
        } else {
            switch -124 {
            case 762:
                var v19 uint16 = v1
                _ = v19
            default:
                fmt.Println(v1, v14)
            }
            v20 := GenS1{F0: 9223372036854775807}
            _ = v20
        }
        _ = fi15
        _ = v16
    }
    var v21 uint = 255
    v6 = v6
    _ = len(v2)
    var v22 uint = (^(v21))
    fmt.Println(v21)
    if (334 <= 0) {
        delete(v4, 32768)
    } else {
        _ = len(v4)
        var v23 int = -9223372036854775808
        v1 += 726
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
