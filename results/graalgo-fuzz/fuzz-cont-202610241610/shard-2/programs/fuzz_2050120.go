package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 bool }
func genFn0() int {
    var v4 uint64 = 115
    var v5 int32 = -2147483648
    var v6 int = -512
    var v7 int = (v6 | v6)
    var v8 int16 = 184
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return v7
}

func main() {
    var v1 int8 = -63
    v2 := make(map[int]int)
    var v3 uint32 = 65535
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 float64 = -2.25
        if (!((!true) && false)) {
            if (!false) {
                if (!true) {
                    var v11 float32 = -2.25
                    v12 := []int{}
                    _ = v11
                    _ = v12
                }
                var v13 uint = 18446744073709551615
                _ = v13
            } else {
                fmt.Println(v3, v10)
            }
            switch genFn0() {
            case 42:
                v10 -= -2.25
            }
        } else {
            var v14 uint = 9
            _ = v14
        }
        for fi15 := 0; fi15 < 11; fi15++ {
            v3 = v3
            var v16 int32 = 0
            fmt.Println(v1, fi15, fi9)
            _ = fi15
            _ = v16
        }
        _ = fi9
        _ = v10
    }
    v3 = (v3 * v3)
    v17 := GenS1{F0: 42, F1: 42, F2: (76 == 128)}
    v3 += 4294967295
    v18 := v2[genFn0()]
    v19 := GenS1{F0: 1, F1: 229, F2: ((v3 % (((v3 % (((v3 % ((4294967295) | 1))) | 1))) | 1)) == (v3 & 2147483648))}
    for fi20 := 0; fi20 < 18; fi20++ {
        if false {
            for fi21 := 0; fi21 < 9; fi21++ {
                delete(v2, fi20)
                var v22 int = (fi21 % ((-721) | 1))
                _ = fi21
                _ = v22
            }
        } else {
            var v23 float64 = (0.0 + -2.25)
            fmt.Println(v23, v23)
            _ = v23
        }
        var v24 float32 = (float32(float64(0.0)) - 0.0)
        _ = fi20
        _ = v24
    }
    delete(v2, v18)
    var v25 int64 = 9223372036854775807
    if true {
        var v26 uint = 415
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v3
}
