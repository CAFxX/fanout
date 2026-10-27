package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint8) uint16 {
    v4 := []int{42, 126, 946}
    var v5 string = "hello"
    if len(v4) > 0 {
        v4[0] = 197
    }
    var v6 int64 = 9223372036854775807
    if len(v4) > 0 {
        v4[0] = -62
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 int16, p1 int64) uint8 {
    fmt.Println(p0)
    var v7 int32 = 1
    var v8 float32 = 1e10
    if (-231 == 1) {
        var v9 int32 = (v7 - -2147483648)
        v7 += (-(v7))
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 32
}

func main() {
    v1 := []uint32{323, 665}
    var v2 uint8 = 235
    var v3 int64 = 42
    v2 += 0
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 float32 = 0.0
        v12 := GenS1{F0: (v3 | (v3 & -545))}
        _ = fi10
        _ = v11
        _ = v12
    }
    if len(v1) > 0 {
        v1[0] = 1
    }
    var v13 uint16 = 42
    v2 = v2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 float32 = 1.5
        var v16 uint = 1
        if (1 == 1000) {
            var v17 int32 = -1
            v18 := GenS1{F0: v3}
            _ = v17
            _ = v18
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    v13 = (v13 ^ (^(v13)))
    v2 = (v2 << ((18446744073709551615) & 63))
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    if (true && false) {
        var v19 int16 = -1
        if (true || true) {
            for fi20 := 0; fi20 < 11; fi20++ {
                _ = len(v1)
                _ = fi20
            }
            switch "\n\t" {
            case "a":
                v19 = v19
            case "foo bar":
                for fi21 := 0; fi21 < 19; fi21++ {
                    var v22 int = (fi21 | 42)
                    v3 -= v3
                    v1 = append(v1, 27)
                    _ = fi21
                    _ = v22
                }
            default:
                v13 = (v13 >> ((18446744073709551615) & 63))
            }
            v19 += 1
        }
        _ = v19
    } else {
        var v23 int16 = 32767
        _ = len(v1)
        v1 = append(v1, 4294967295)
        _ = v23
    }
    var v24 int8 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v2
    _ = v24
    _ = v3
}
