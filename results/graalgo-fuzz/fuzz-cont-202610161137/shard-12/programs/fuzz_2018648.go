package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 int16; F1 uint16; F2 uint32 }
func genFn0() int32 {
    v5 := []int{0, -2147483648, 2147483647}
    if (!(!((1e10 != (1e10 / -2.25)) || true))) {
        var v6 uint64 = 4294967295
        _ = v6
    } else {
        var v7 string = "a"
        if (!(643 == 164)) {
            v7 = v7
        } else {
            for fi8 := 0; fi8 < 6; fi8++ {
                var v9 uint16 = 1
                var v10 int16 = 1
                var v11 float64 = (float64(-2.25) * float64(1.5))
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            var v12 bool = true
            _ = v12
        }
        var v13 float32 = 1.5
        _ = v13
        _ = v7
    }
    var v14 uint = 65535
    if true {
        v5 = append(v5, 1)
        v14 -= (v14 / ((v14) | 1))
    }
    _ = v14
    _ = v5
    return 0
}

func main() {
    var v1 int64 = -1
    v2 := make(map[uint32]uint32)
    var v3 string = "foo bar"
    var v4 uint8 = 42
    _ = len(v2)
    fmt.Println(v3, v4, v3)
    var v15 int64 = (^(v1))
    var v16 float64 = float64(float32(0.1))
    fmt.Println(v16)
    var v17 uint16 = 308
    fmt.Println(v3, v1)
    if true {
        v18 := []int{61}
        var v19 uint = 18446744073709551615
        switch v3 {
        case "":
            v20 := GenS1{F0: 100}
            _ = v20
        case "hello":
            if true {
                var v21 int8 = 127
                _ = v21
            } else {
                v22 := GenS2{F0: 30, F1: 42}
                v3 = "\n\t"
                _ = v22
            }
        case "a":
            for fi23 := 0; fi23 < 12; fi23++ {
                fmt.Println(v19)
                fmt.Println(v3, v17, v16)
                for fi24 := 0; fi24 < 5; fi24++ {
                    v1 = v15
                    _ = fi24
                }
                _ = fi23
            }
        default:
            var v25 uint64 = 527
            _ = v25
        }
        _ = v18
        _ = v19
    }
    var v26 uint16 = (v17 * v17)
    v27 := GenS3{F0: 1, F1: 32767, F2: 65535}
    v2[1] = 634
    v28 := GenS1{F0: 0}
    v29 := []int{42}
    _ = len(v2)
    var v30 int32 = 1
    if len(v29) > 0 {
        v29[0] = 2147483647
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
