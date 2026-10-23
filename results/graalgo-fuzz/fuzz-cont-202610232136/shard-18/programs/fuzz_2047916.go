package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0() int16 {
    v7 := GenS1{F0: -1}
    for fi8 := 0; fi8 < 14; fi8++ {
        if false {
            switch "" {
            case "\n\t":
                switch fi8 {
                case 364:
                    fmt.Println(fi8)
                default:
                    if (true || (0 >= 0)) {
                        var v9 uint32 = 4294967295
                        _ = v9
                    }
                }
            }
            var v10 int64 = 9223372036854775807
            _ = v10
        }
        var v11 int16 = -1
        _ = fi8
        _ = v11
    }
    _ = v7
    return 32767
}
func genFn1(p0 uint) uint8 {
    p0 = (p0 & 2)
    var v12 uint16 = 520
    v12 = (v12 / ((255) | 1))
    _ = p0
    _ = v12
    return 255
}

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[int]int)
    var v3 uint32 = 2147483648
    var v4 string = ("\n\t" + "")
    var v5 int64 = 1
    v6 := []uint32{v3, (v3 / ((v3) | 1)), (^(v3)), 65535}
    var v13 uint = 332
    v13 = 1
    var v14 float32 = (float32((float64(0.0) / (0.0 + (3.14159 * 1e10)))) + 1e10)
    var v15 uint8 = 140
    var v16 uint32 = (^(v3))
    v16 -= v16
    if (false && false) {
        if (42 >= 1) {
            var v17 int = 786
            _ = v17
        }
        var v18 uint64 = v1
        _ = v18
    }
    fmt.Println(v14, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
