package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int8) uint32 {
    fmt.Println(p0)
    p0 += (p0 / (((p0 + (p0 / ((p0) | 1)))) | 1))
    if (0.1 != (0.0 - 1.5)) {
        var v6 uint8 = 1
        fmt.Println(p0)
        _ = v6
    } else {
        v7 := []int{710}
        _ = len(v7)
        _ = v7
    }
    _ = p0
    return 1
}

func main() {
    var v1 uint32 = 42
    var v2 string = "a"
    var v3 uint8 = 255
    var v4 bool = false
    var v5 uint64 = 2
    v3 = (v3 * (v3 >> ((648) & 63)))
    v1 -= (v1 + 249)
    var v8 uint64 = 1
    var v9 float64 = 1.5
    if (-32768 < -1) {
        var v10 uint16 = 1
        if true {
            for fi11 := 0; fi11 < 7; fi11++ {
                fmt.Println(v2)
                _ = fi11
            }
            fmt.Println(v5, v5)
            var v12 int32 = 427
            _ = v12
        } else {
            v5 += (v5 + v8)
            var v13 string = (v2 + v2)
            fmt.Println(v13, v10, v4)
            _ = v13
        }
        var v14 uint32 = genFn0(110)
        _ = v10
        _ = v14
    } else {
        for fi15 := 0; fi15 < 9; fi15++ {
            v16 := []int{(fi15 & fi15), (fi15 % (((fi15 * (fi15 >> ((18446744073709551615) & 63)))) | 1))}
            var v17 int32 = 42
            v18 := GenS1{F0: 18, F1: ((^(v17)) <= v17)}
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
    }
    v1 += genFn0(1)
    v5 -= (v8 >> ((42) & 63))
    switch v2 {
    case "":
        var v19 uint16 = 2
        _ = v19
    case "\n\t":
        var v20 float32 = (0.0 + float32(float64((float32(v9) * (0.1 - 0.1)))))
        _ = v20
    default:
        if (v4 && (-667 <= 1)) {
            v21 := GenS2{F0: (v8 * v8)}
            _ = v21
        }
    }
    v22 := []int{-14, 9223372036854775807, 0, -9223372036854775808}
    v4 = true
    var v23 int8 = 127
    fmt.Println(v8, v4, v1)
    switch 127 {
    case -9223372036854775808:
        for fi24 := 0; fi24 < 4; fi24++ {
            if ((fi24 / ((fi24) | 1)) > fi24) {
                var v25 bool = ((0 > 429) || true)
                if v25 {
                    v26 := GenS3{F0: 0}
                    var v27 float32 = 0.1
                    _ = v26
                    _ = v27
                }
                for fi28 := 0; fi28 < 17; fi28++ {
                    v22 = append(v22, fi24)
                    v25 = v4
                    _ = fi28
                }
                _ = v25
            }
            if len(v22) > 0 {
                v22[0] = fi24
            }
            v9 = 1e10
            _ = fi24
        }
    case 4294967295:
        var v29 float64 = 3.14159
        _ = v29
    default:
        v8 -= (v8 + (^(v8)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
