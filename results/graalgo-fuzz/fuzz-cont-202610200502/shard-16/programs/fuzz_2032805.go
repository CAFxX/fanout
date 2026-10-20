package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 uint8; F2 string }
type GenS3 struct { F0 string; F1 int8; F2 int32 }
func genFn0(p0 int64, p1 int) uint8 {
    if true {
        p0 = (p0 >> ((42) & 63))
        if false {
            for fi5 := 0; fi5 < 16; fi5++ {
                p0 = (p0 * 9223372036854775807)
                _ = fi5
            }
            v6 := GenS1{F0: -1}
            v7 := []int{(p1 << ((42) & 63))}
            _ = v6
            _ = v7
        } else {
            v8 := []int{p1, (p1 - (p1 - 2147483647)), p1}
            _ = v8
        }
    }
    var v9 uint64 = 175
    _ = p0
    _ = p1
    _ = v9
    return 255
}

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{}
    v3 := make(map[int]bool)
    var v4 uint64 = 595
    v10 := GenS1{F0: 42}
    if false {
        var v11 int = 1
        _ = v11
    } else {
        v4 += (v4 / ((2) | 1))
        var v12 int32 = -227
        _ = v12
    }
    fmt.Println(v4)
    v13 := GenS1{F0: 0}
    var v14 int64 = -9223372036854775808
    var v15 int = 32767
    v15 = (^(v15))
    v15 -= 2147483647
    var v16 float64 = float64((3.14159 - ((1.5 * (3.14159 + 0.1)) * (1e10 - (0.0 / 1.0)))))
    fmt.Println(v16, v14, v15)
    v17 := GenS2{F0: (v14 - (-(v14))), F1: genFn0((v14 / ((v14) | 1)), 396), F2: ""}
    var v18 uint32 = 453
    fmt.Println(v4)
    switch "hello" {
    case "":
        var v19 int8 = -1
        _ = v19
    case "\n\t":
        for fi20 := 0; fi20 < 5; fi20++ {
            var v21 int64 = 0
            v22 := GenS3{F0: "", F1: 127, F2: -1}
            _ = fi20
            _ = v21
            _ = v22
        }
    default:
        for fi23 := 0; fi23 < 7; fi23++ {
            if (0 >= 573) {
                var v24 int32 = 42
                if len(v2) > 0 {
                    v2[0] = (v18 + uint32(v4))
                }
                _ = v24
            }
            _ = fi23
        }
    }
    for fi25 := 0; fi25 < 10; fi25++ {
        var v26 int8 = -1
        _ = fi25
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
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
