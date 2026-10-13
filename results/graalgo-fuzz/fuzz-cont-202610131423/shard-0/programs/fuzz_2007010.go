package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }
func genFn0(p0 int8) uint32 {
    v4 := []int{-63, -9223372036854775808, 42, 2147483647}
    p0 += p0
    var v5 string = "\n\t"
    _ = p0
    _ = v4
    _ = v5
    return 4294967295
}
func genFn1(p0 uint8) int32 {
    if ((!true) && (2 > 2147483648)) {
        p0 = 42
        var v6 int64 = -1
        fmt.Println(p0, p0)
        _ = v6
    }
    v7 := GenS1{F0: 573}
    if false {
        p0 = (p0 % (((p0 & p0)) | 1))
    }
    _ = p0
    _ = v7
    return -153
}

func main() {
    var v1 uint64 = 4294967295
    var v2 uint = 18446744073709551615
    v3 := make(map[int]int)
    fmt.Println(v1, v2)
    for fi8 := 0; fi8 < 1; fi8++ {
        for fi9 := 0; fi9 < 16; fi9++ {
            v2 += v2
            if false {
                switch 232 {
                case 32767:
                    delete(v3, (fi9 * (fi9 >> ((4294967295) & 63))))
                case 0:
                    v10 := []int{32768, fi8, (fi8 % ((fi8) | 1))}
                    _ = v10
                default:
                    fmt.Println(v1)
                }
                if (!(!(((!(false && false)) || true) || (1.5 < 1e10)))) {
                    v1 = (v1 % ((1) | 1))
                }
            } else {
                v11 := GenS1{F0: v1}
                var v12 int8 = 89
                var v13 int32 = 42
                _ = v11
                _ = v12
                _ = v13
            }
            v14 := GenS1{F0: (v1 << ((v2) & 63))}
            _ = fi9
            _ = v14
        }
        if (65535 != 394) {
            v3[0] = (fi8 / ((32767) | 1))
            if true {
                var v15 int32 = 1
                v2 -= (v2 >> (((v2 >> ((v2) & 63))) & 63))
                _ = v15
            }
        }
        var v16 float64 = (1.5 + float64(float32(((1e10 / (-2.25 + 0.0)) - float64(1.5)))))
        _ = fi8
        _ = v16
    }
    var v17 float64 = 0.1
    v18 := v3[42]
    v19 := []int{}
    if (false && false) {
        if true {
            var v20 uint = v2
            _ = v20
        }
        v21 := GenS2{F0: v18}
        _ = v21
    }
    fmt.Println(v1, v1)
    var v22 float64 = 1e10
    var v23 uint32 = 0
    if true {
        var v24 int8 = 0
        _ = v24
    } else {
        for fi25 := 0; fi25 < 14; fi25++ {
            var v26 float64 = 3.14159
            var v27 int8 = 100
            switch v18 {
            case 9223372036854775807:
                switch "\n\t" {
                case "foo bar":
                    v23 = (v23 >> ((v2) & 63))
                }
            }
            _ = fi25
            _ = v26
            _ = v27
        }
        if ((1e10 + (1e10 - 1.5)) > (((((0.0 * (float32(0.1) + (0.1 + ((1e10 * 0.0) + 0.1)))) / 1.0) * 3.14159) * 0.0) - 0.0)) {
            fmt.Println(v17, v23, v18)
            v19 = append(v19, 9223372036854775807)
            var v28 float64 = ((-2.25 - v17) * 1e10)
            _ = v28
        }
        v1 += (v1 % (((v1 * (v1 | (v1 ^ v1)))) | 1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
}
