package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint16; F1 uint16; F2 int }
type GenS3 struct { F0 uint }
func genFn0(p0 int16, p1 int64, p2 int) int32 {
    var v5 uint16 = 726
    fmt.Println(p0)
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := []int{p2}
        _ = fi6
        _ = v7
    }
    var v8 int64 = p1
    var v9 string = "hello"
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v8
    _ = v9
    return -2147483648
}

func main() {
    var v1 uint64 = 4294967295
    var v2 uint = 0
    var v3 int64 = 0
    var v4 uint64 = (v1 / ((v1) | 1))
    v10 := []int{32767, 468, 127}
    switch 4294967295 {
    case 127:
        if (false && (!false)) {
            v11 := GenS1{F0: 367, F1: (true || (true || true))}
            _ = v11
        } else {
            for fi12 := 0; fi12 < 16; fi12++ {
                var v13 int = (fi12 - 32768)
                _ = fi12
                _ = v13
            }
            var v14 string = (("\n\t" + (("foo bar" + "a") + "hello")) + "hello")
            for fi15 := 0; fi15 < 4; fi15++ {
                var v16 int8 = 100
                var v17 uint16 = 32767
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = v14
        }
    }
    var v18 uint32 = 4294967295
    v19 := GenS2{F0: 65535, F1: uint16(v1), F2: 227}
    v1 += uint64(v18)
    fmt.Println(v4, v1, v2)
    v10 = append(v10, 2147483647)
    if false {
        v10 = append(v10, -466)
    }
    if (1 < 0) {
        switch 128 {
        case 128:
            var v20 uint16 = 255
            _ = v20
        default:
            var v21 int8 = 109
            _ = v21
        }
        v2 = 842
        var v22 int64 = v3
        _ = v22
    }
    switch "" {
    case "":
        if (1e10 < 0.1) {
            v3 -= int64(v18)
            for fi23 := 0; fi23 < 12; fi23++ {
                if (254 > 128) {
                    v1 = v1
                    v24 := GenS3{F0: 102}
                    v2 += (^(v2))
                    _ = v24
                }
                _ = fi23
            }
        } else {
            v10 = append(v10, -9223372036854775808)
        }
    default:
        for fi25 := 0; fi25 < 10; fi25++ {
            for fi26 := 0; fi26 < 1; fi26++ {
                v4 -= (^(v1))
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                _ = fi26
            }
            v4 -= v1
            v10 = append(v10, 353)
            _ = fi25
        }
    }
    _ = len(v10)
    fmt.Println(v18, v2, v2)
    fmt.Println(v1, v18, v4)
    fmt.Println(v4, v18)
    if len(v10) > 0 {
        v10[0] = 128
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
