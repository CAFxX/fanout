package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 int32; F2 int8 }
func genFn0(p0 uint64, p1 uint32) int {
    var v6 bool = (!(0 > 42))
    v7 := GenS1{F0: 255, F1: 42}
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 uint = 65535
        if (42 >= (v9 * (v9 * v9))) {
            if true {
                if (int64(p1) < -1) {
                    var v10 float64 = 3.14159
                    v11 := []int{(fi8 << ((v9) & 63)), 42}
                    _ = v10
                    _ = v11
                }
                var v12 uint8 = 42
                _ = v12
            }
            if v6 {
                v13 := GenS1{F0: uint8(p0), F1: 127}
                var v14 int64 = 0
                for fi15 := 0; fi15 < 9; fi15++ {
                    var v16 float64 = 1.5
                    var v17 int64 = (v14 & v14)
                    v18 := GenS1{F0: 136, F1: 127}
                    _ = fi15
                    _ = v16
                    _ = v17
                    _ = v18
                }
                _ = v13
                _ = v14
            } else {
                v19 := GenS2{F0: -32768}
                fmt.Println(v6, v6, p0)
                _ = v19
            }
        }
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    v1 := make(map[string]string)
    var v2 int64 = 0
    var v3 uint64 = 516
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]string)
    switch genFn0(v3, 634) {
    case 56:
        v20 := GenS3{F0: (150 <= 42), F1: -2147483648, F2: 127}
        _ = v20
    }
    v2 += (^(v2))
    for fi21 := 0; fi21 < 2; fi21++ {
        switch 1 {
        case -9223372036854775808:
            var v22 uint8 = 42
            _ = v22
        case 32767:
            v23 := GenS1{F0: 2, F1: 2}
            _ = v23
        }
        var v24 bool = ((v2 - (v2 * 287)) == v2)
        _ = fi21
        _ = v24
    }
    v3 -= v3
    var v25 uint16 = 32767
    var v26 string = "foo bar"
    fmt.Println(v3)
    var v27 int32 = 2147483647
    var v28 int16 = -109
    delete(v5, 778)
    v28 = v28
    v2 += v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
