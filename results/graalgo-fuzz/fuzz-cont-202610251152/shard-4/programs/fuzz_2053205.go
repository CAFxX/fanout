package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint8 }
type GenS2 struct { F0 int; F1 int64; F2 int32 }
func genFn0(p0 int32, p1 uint64) int8 {
    p0 += p0
    if true {
        p0 -= (p0 * p0)
        var v6 float32 = 1.5
        if true {
            for fi7 := 0; fi7 < 19; fi7++ {
                var v8 uint16 = 39
                switch "" {
                case "foo bar":
                    v9 := GenS1{F0: 1, F1: p0, F2: uint8(v8)}
                    _ = v9
                case "\n\t":
                    var v10 string = "a"
                    _ = v10
                }
                var v11 int64 = -1
                _ = fi7
                _ = v11
                _ = v8
            }
            v12 := GenS2{F0: -372, F1: 42, F2: (^(p0))}
            if false {
                var v13 int = int(p0)
                var v14 float64 = 1.5
                _ = v13
                _ = v14
            }
            _ = v12
        }
        _ = v6
    }
    var v15 uint64 = p1
    fmt.Println(p0, p0, p0)
    v16 := GenS1{F0: -1, F1: -230, F2: 0}
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return -1
}
func genFn1(p0 bool, p1 int64) uint16 {
    var v17 uint8 = 127
    v18 := []int{1, 416}
    var v19 int8 = genFn0(0, 1)
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    _ = v19
    return 229
}

func main() {
    var v1 uint32 = 560
    v2 := make(map[int]string)
    v3 := make(map[int]uint32)
    var v4 int64 = 9223372036854775807
    var v5 int16 = -830
    var v20 int64 = (-(v4))
    var v21 float64 = 0.0
    var v22 uint8 = 0
    v4 += v4
    delete(v3, 32768)
    if false {
        _ = len(v2)
        v22 = (v22 | v22)
        if (v5 > int16(v22)) {
            var v23 uint16 = 0
            _ = v23
        }
    } else {
        var v24 uint8 = 2
        v3[-358] = (v1 ^ v1)
        _ = v24
    }
    var v25 uint = 188
    for fi26 := 0; fi26 < 7; fi26++ {
        v27 := []int{fi26, fi26}
        v28 := GenS2{F0: fi26, F1: v20, F2: 0}
        _ = fi26
        _ = v27
        _ = v28
    }
    v29 := GenS1{F0: (^(v20)), F1: -2147483648, F2: 128}
    v4 = v20
    v20 -= (v20 >> (((v25 * (v25 / (((^(v25))) | 1)))) & 63))
    v4 -= v20
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
