package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint16 }
type GenS2 struct { F0 int8 }
func genFn0() int16 {
    var v6 int32 = 2147483647
    switch -9223372036854775808 {
    case 546:
        var v7 string = "foo bar"
        _ = v7
    default:
        v8 := GenS1{F0: 100, F1: v6, F2: 1}
        _ = v8
    }
    _ = v6
    return -1
}
func genFn1(p0 string, p1 int64) uint64 {
    p1 = p1
    for fi9 := 0; fi9 < 9; fi9++ {
        fmt.Println(p0, p1)
        var v10 uint16 = 65535
        _ = fi9
        _ = v10
    }
    v11 := []int{1}
    v11 = append(v11, -9223372036854775808)
    _ = p0
    _ = p1
    _ = v11
    return 2
}

func main() {
    var v1 uint16 = 32767
    var v2 string = "hello"
    var v3 int64 = -11
    var v4 int32 = -2147483648
    var v5 uint32 = 2147483648
    var v12 int = 301
    v3 -= (^(v3))
    var v13 float64 = 1.5
    v14 := GenS1{F0: int8(v3), F1: (v4 % ((0) | 1)), F2: v1}
    v15 := []int{(v12 ^ (v12 + (-(v12)))), 32767, (^(v12)), v12}
    var v16 uint32 = v5
    v17 := GenS2{F0: int8(v12)}
    if true {
        switch (v12 * v12) {
        case 0:
            fmt.Println(v1)
        }
        switch "a" {
        case "":
            var v18 int32 = int32(v12)
            _ = v18
        case "foo bar":
            var v19 bool = ((0.0 - (3.14159 - 1e10)) < 1e10)
            _ = v19
        default:
            var v20 uint32 = (^(v5))
            _ = v20
        }
        if true {
            for fi21 := 0; fi21 < 18; fi21++ {
                fmt.Println(v3, v12)
                for fi22 := 0; fi22 < 4; fi22++ {
                    if true {
                        v1 -= 343
                        v15 = append(v15, 42)
                        v23 := GenS1{F0: -1, F1: int32(v5), F2: 0}
                        _ = v23
                    }
                    _ = fi22
                }
                v4 = (v4 - 1)
                _ = fi21
            }
            var v24 string = v2
            if false {
                var v25 uint = 224
                _ = v25
            }
            _ = v24
        } else {
            fmt.Println(v12, v1, v12)
            if (v13 <= -2.25) {
                fmt.Println(v13, v3, v4)
                if len(v15) > 0 {
                    v15[0] = -9223372036854775808
                }
            }
        }
    } else {
        var v26 float64 = 1e10
        var v27 uint = 65535
        if (((v4 ^ (v4 % ((int32(v12)) | 1))) >= v4) && ((127 <= -105) && (v16 < v16))) {
            var v28 int16 = 1
            var v29 uint8 = 0
            _ = v28
            _ = v29
        }
        _ = v26
        _ = v27
    }
    var v30 int16 = -1
    var v31 uint32 = 139
    var v32 bool = false
    v30 = (v30 + genFn0())
    v15 = append(v15, int(v1))
    v33 := GenS2{F0: 1}
    if len(v15) > 0 {
        v15[0] = (v12 / (((v12 - (v12 | v12))) | 1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
