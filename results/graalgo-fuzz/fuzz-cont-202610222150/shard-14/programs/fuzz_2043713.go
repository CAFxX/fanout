package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 bool; F1 int32 }
func genFn0(p0 int32, p1 int16, p2 uint32) bool {
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 int16 = (-(p1))
        var v8 int16 = (v7 ^ -1)
        fmt.Println(p1, v8, v8)
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 42, F1: (p2 & p2)}
    if ("" < "a") {
        p2 = p2
    } else {
        p2 += p2
        switch 4294967295 {
        case 42:
            p0 = 78
        case 127:
            if (1.5 < 0.0) {
                fmt.Println(p2, p2, p2)
                var v10 uint32 = (p2 % (((p2 << ((4294967295) & 63))) | 1))
                var v11 int = 547
                _ = v10
                _ = v11
            }
        case -405:
            var v12 string = ""
            _ = v12
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return false
}
func genFn1(p0 string, p1 uint) int16 {
    var v13 uint16 = 255
    var v14 int64 = 1
    v15 := GenS1{F0: (v14 / ((0) | 1)), F1: 0}
    v13 = (v13 ^ (^(v13)))
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    _ = v15
    return -668
}

func main() {
    var v1 uint = 255
    v2 := make(map[string]int)
    var v3 uint8 = 242
    var v4 uint8 = (v3 ^ v3)
    var v5 uint64 = 303
    for fi16 := 0; fi16 < 14; fi16++ {
        var v17 int16 = 1000
        switch (-(fi16)) {
        case 84:
            var v18 int64 = -1
            _ = v18
        case 289:
            v17 = v17
        }
        _ = fi16
        _ = v17
    }
    fmt.Println(v3, v4)
    v19 := GenS2{F0: genFn1("\n\t", 489), F1: (v1 % ((v1) | 1))}
    for fi20 := 0; fi20 < 19; fi20++ {
        _ = len(v2)
        _ = fi20
    }
    var v21 int = 9223372036854775807
    v22 := GenS2{F0: 42, F1: (v1 << ((18446744073709551615) & 63))}
    switch "\n\t" {
    case "\n\t":
        var v23 float64 = (((1e10 + 3.14159) / ((3.14159 - (((1e10 - ((-2.25 - 0.0) / 1e10)) - (((0.1 / float64((0.0 + (1e10 / ((((1e10 - ((3.14159 - float32((3.14159 + 0.1))) + 0.0)) + float32(1e10)) + 1e10) / ((3.14159 / 1.0) - 0.0)))))) / 1.5) - 1.5)) / 1e10)) - (1e10 / (3.14159 * 3.14159)))) + (0.0 / (((-2.25 * float64((((0.0 + ((0.1 - 3.14159) - ((1e10 + -2.25) - 0.1))) - (-2.25 * 1e10)) + 0.0))) + 0.1) / 1.0)))
        _ = v23
    case "foo bar":
        for fi24 := 0; fi24 < 20; fi24++ {
            for fi25 := 0; fi25 < 6; fi25++ {
                var v26 uint8 = 0
                v27 := []int{(fi25 * (fi24 | 0))}
                v21 -= (fi25 << (((v1 >> (((v1 | (v1 | v1))) & 63))) & 63))
                var v28 uint32 = 65535
                _ = fi25
                _ = v26
                _ = v27
                _ = v28
            }
            v4 += 128
            _ = fi24
        }
    default:
        var v29 string = ("foo bar" + ("hello" + "\n\t"))
        _ = v29
    }
    v5 = (v5 & v5)
    if false {
        v30 := []int{}
        var v31 int64 = 1
        v32 := GenS3{F0: true, F1: 685}
        _ = v30
        _ = v31
        _ = v32
    } else {
        v4 = (^(v3))
        v2["a"] = v21
        var v33 float64 = (1.5 + (3.14159 * (1.5 * 0.1)))
        _ = v33
    }
    fmt.Println(v21, v3)
    var v34 int8 = -1
    fmt.Println(v1, v3, v34)
    v35 := GenS2{F0: int16(v21), F1: v1}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
