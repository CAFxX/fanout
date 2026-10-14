package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint64, p1 bool, p2 int64) int8 {
    var v5 uint8 = 255
    for fi6 := 0; fi6 < 13; fi6++ {
        p1 = false
        _ = fi6
    }
    fmt.Println(p1)
    v7 := []int{-1, -956}
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v7
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]int)
    var v3 float64 = 0.1
    v4 := make(map[int]uint32)
    v3 = (v3 * 0.0)
    v8 := GenS1{F0: 32767}
    v9 := GenS2{F0: -32768}
    var v10 uint16 = 42
    if (!(((true || (2 != 0)) && (!true)) && (((1 <= 0) && true) && false))) {
        fmt.Println(v3, v10)
    }
    fmt.Println(v10)
    if true {
        for fi11 := 0; fi11 < 18; fi11++ {
            if true {
                for fi12 := 0; fi12 < 16; fi12++ {
                    v13 := GenS1{F0: 65535}
                    _ = fi12
                    _ = v13
                }
                var v14 int64 = 0
                _ = v14
            }
            _ = fi11
        }
    } else {
        switch "" {
        case "\n\t":
            v15 := GenS2{F0: -687}
            _ = v15
        default:
            if (false && true) {
                v3 = float64(0.0)
                v10 += v10
            } else {
                if ((0 > 228) && true) {
                    var v16 uint8 = 127
                    v3 -= v3
                    _ = v16
                } else {
                    v17 := v2[4294967295]
                    v3 = (v3 - v3)
                    _ = v17
                }
                delete(v4, 32767)
            }
        }
    }
    if true {
        v10 = 2
    } else {
        var v18 int16 = 0
        _ = v18
    }
    v3 += v3
    switch 128 {
    case 32767:
        var v19 float64 = float64(0.1)
        _ = v19
    case 127:
        fmt.Println(v3, v3)
    case -1:
        switch 32767 {
        case 4294967295:
            v3 -= v3
        case 9223372036854775807:
            var v20 uint = 1
            _ = v20
        case 32767:
            fmt.Println(v10, v10)
        default:
            if ((1 > 752) || true) {
                if true {
                    v21 := GenS1{F0: v10}
                    var v22 int16 = -97
                    switch 457 {
                    case 128:
                        v3 = v3
                    case -351:
                        var v23 int64 = 42
                        _ = v23
                    }
                    _ = v21
                    _ = v22
                }
            } else {
                var v24 uint32 = 56
                _ = v24
            }
        }
    }
    fmt.Println(v10, v3)
    fmt.Println(v10, v10)
    for fi25 := 0; fi25 < 18; fi25++ {
        fmt.Println(v3, v10)
        _ = fi25
    }
    var v26 uint32 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
