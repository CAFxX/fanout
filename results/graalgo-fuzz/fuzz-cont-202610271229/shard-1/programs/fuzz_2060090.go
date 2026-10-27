package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 bool }
func genFn0(p0 int32, p1 int, p2 bool) uint16 {
    p0 = 0
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    return 432
}
func genFn1() uint32 {
    if false {
        var v5 int8 = -4
        var v6 uint = 255
        _ = v5
        _ = v6
    }
    var v7 float64 = 1e10
    if (1 >= 42) {
        var v8 int64 = -552
        v8 += (v8 + v8)
        _ = v8
    } else {
        v9 := GenS1{F0: true}
        _ = v9
    }
    _ = v7
    return 793
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = 0
    v3 := []uint32{2, 2147483647, 4294967295, 311}
    var v4 int = 2147483647
    for fi10 := 0; fi10 < 15; fi10++ {
        if (false || false) {
            switch v4 {
            case 9223372036854775807:
                for fi11 := 0; fi11 < 5; fi11++ {
                    if (!((2 == 154) && (62 != 2))) {
                        fmt.Println(v4, v2)
                        v12 := GenS1{F0: (!false)}
                        v13 := GenS2{F0: 42, F1: 446}
                        _ = v12
                        _ = v13
                    }
                    _ = fi11
                }
            case 32767:
                v2 = v2
            }
        } else {
            switch (fi10 - (v4 / ((v4) | 1))) {
            case 42:
                var v14 uint8 = 0
                _ = v14
            case 32767:
                var v15 uint64 = 698
                _ = v15
            case -134:
                for fi16 := 0; fi16 < 17; fi16++ {
                    delete(v1, 2147483647)
                    _ = fi16
                }
            }
            fmt.Println(v4)
        }
        var v17 uint8 = 126
        _ = fi10
        _ = v17
    }
    v18 := GenS1{F0: (-128 <= 42)}
    v2 += 361
    var v19 bool = (890 < 65535)
    v19 = (0.1 == (float32(3.14159) + (3.14159 / 3.14159)))
    delete(v1, 352)
    v20 := GenS1{F0: v19}
    var v21 int16 = 1000
    v19 = (-2.25 > (-2.25 - ((3.14159 * (1e10 / 0.1)) / -2.25)))
    v4 += -967
    v22 := GenS3{F0: (107 > 0)}
    v23 := GenS2{F0: 32767, F1: (^(v21))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
