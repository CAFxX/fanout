package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 string }
type GenS2 struct { F0 uint; F1 bool; F2 int16 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int) uint32 {
    for fi7 := 0; fi7 < 17; fi7++ {
        p0 += (p0 % ((fi7) | 1))
        v8 := GenS1{F0: false, F1: 1, F2: ("\n\t" + "foo bar")}
        v9 := []int{(p0 / ((p0) | 1)), (p0 ^ (p0 / (((p0 ^ fi7)) | 1))), fi7}
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS2{F0: 42, F1: false, F2: 0}
    if (!((false || false) || (42 < -9223372036854775808))) {
        var v11 int16 = 1000
        _ = v11
    }
    if (4294967295 < 65535) {
        for fi12 := 0; fi12 < 10; fi12++ {
            fmt.Println(p0, p0)
            _ = fi12
        }
        v13 := GenS2{F0: 255, F1: true, F2: 1}
        var v14 uint8 = 128
        _ = v13
        _ = v14
    }
    _ = p0
    _ = v10
    return 792
}

func main() {
    var v1 uint8 = 193
    var v2 uint8 = (^(v1))
    v3 := make(map[int]bool)
    var v4 uint16 = 645
    var v5 int32 = -2147483648
    var v6 bool = false
    v5 = v5
    switch 0 {
    case 0:
        var v15 int32 = 0
        _ = v15
    }
    var v16 uint32 = 1
    switch "\n\t" {
    case "\n\t":
        v6 = (4294967295 < 18446744073709551615)
    case "":
        var v17 int = 42
        _ = v17
    default:
        switch 329 {
        case 32767:
            v18 := []int{283, -2147483648, -1, -299}
            _ = v18
        }
    }
    v19 := []int{-334, 1}
    var v20 uint8 = 42
    var v21 int = 34
    var v22 int = 2147483647
    v21 -= (v21 + int(v4))
    fmt.Println(v16, v1, v20)
    v22 = v22
    v16 += v16
    fmt.Println(v20, v1, v5)
    _ = len(v19)
    if (!(((1e10 == -2.25) && v6) || true)) {
        var v23 uint8 = uint8(v4)
        if (v16 != (v16 | 700)) {
            v6 = (int64(v5) < 439)
            v24 := GenS1{F0: (false || false), F1: (v4 - (v4 - (v4 * v4))), F2: ""}
            for fi25 := 0; fi25 < 14; fi25++ {
                if len(v19) > 0 {
                    v19[0] = fi25
                }
                _ = fi25
            }
            _ = v24
        } else {
            v3[(v21 << ((4294967295) & 63))] = true
            fmt.Println(v1, v4, v22)
        }
        _ = v23
    }
    v26 := GenS3{F0: v20}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
