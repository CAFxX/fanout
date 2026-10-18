package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint8 }
type GenS2 struct { F0 int; F1 uint16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint8 = 128
    var v3 int64 = 1
    var v4 uint = 255
    v3 = v3
    var v5 uint = (v4 ^ (v4 + v4))
    if (42 == 42) {
        v3 = (v3 ^ (v3 ^ v3))
        var v6 float32 = (0.1 + -2.25)
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 uint32 = 2147483647
            if (!(42 < 0)) {
                v9 := v1[65535]
                var v10 float32 = 0.1
                switch 4294967295 {
                case -1:
                    _ = len(v1)
                case 128:
                    for fi11 := 0; fi11 < 13; fi11++ {
                        if false {
                            v6 = ((float32(1.5) / 1.0) * 0.0)
                            v4 -= v4
                            var v12 int8 = -128
                            v13 := []int{(fi7 & (fi11 >> ((42) & 63))), 32768, fi11, fi7}
                            _ = v12
                            _ = v13
                        } else {
                            v3 = 208
                            v14 := GenS1{F0: 509, F1: 565, F2: (v2 - 1)}
                            _ = v14
                        }
                        v8 -= v8
                        _ = fi11
                    }
                default:
                    var v15 uint64 = 42
                    _ = v15
                }
                _ = v10
                _ = v9
            }
            fmt.Println(v4, v8)
            _ = fi7
            _ = v8
        }
        _ = v6
    }
    if (255 == 409) {
        var v16 uint = (v5 & (v4 | v5))
        fmt.Println(v3)
        _ = v16
    }
    var v17 int = 0
    var v18 uint64 = 18446744073709551615
    v3 = v3
    fmt.Println(v2)
    v4 = 0
    v17 = v17
    v18 = v18
    v19 := []int{76, (-(v17)), v17, -1}
    v20 := GenS2{F0: (v17 + (^(v17))), F1: 65535}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
