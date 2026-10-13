package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 bool; F1 string }
func genFn0(p0 int8, p1 uint, p2 int32) int {
    for fi5 := 0; fi5 < 12; fi5++ {
        p2 -= (-(p2))
        var v6 uint64 = 2
        var v7 bool = false
        _ = fi5
        _ = v6
        _ = v7
    }
    if (0.1 <= (-2.25 / -2.25)) {
        p2 += (p2 * (p2 * p2))
    }
    _ = p0
    _ = p1
    _ = p2
    return 32768
}
func genFn1() uint {
    var v8 int32 = 799
    v9 := GenS1{F0: -9223372036854775808}
    if ((v8 | v8) == v8) {
        var v10 uint64 = 367
        _ = v10
    }
    _ = v8
    _ = v9
    return 2
}

func main() {
    var v1 float64 = (-2.25 / 1.0)
    v2 := make(map[uint32]bool)
    var v3 uint8 = 97
    var v4 string = "\n\t"
    if true {
        fmt.Println(v3, v3)
    } else {
        switch ("hello" + "hello") {
        case "foo bar":
            var v11 string = ((v4 + v4) + ("a" + (v4 + ("" + (v4 + (v4 + ((v4 + "foo bar") + ((v4 + v4) + v4))))))))
            _ = v11
        default:
            var v12 int = 535
            _ = v12
        }
        fmt.Println(v1)
    }
    fmt.Println(v1)
    switch 9223372036854775807 {
    case 927:
        fmt.Println(v1, v4, v4)
    case 127:
        v3 = (^(v3))
    }
    var v13 uint64 = 1
    var v14 int64 = 1
    fmt.Println(v13, v1)
    v14 = -9223372036854775808
    if ((false || (v14 > -1)) || true) {
        var v15 uint64 = v13
        _ = v15
    } else {
        for fi16 := 0; fi16 < 12; fi16++ {
            v17 := GenS2{F0: 65535, F1: 42}
            var v18 int64 = int64(fi16)
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    var v19 float32 = 3.14159
    var v20 uint64 = (v13 + (^(v13)))
    v21 := GenS1{F0: 9223372036854775807}
    delete(v2, 570)
    v22 := GenS3{F0: true, F1: v4}
    for fi23 := 0; fi23 < 16; fi23++ {
        if (44 < 37) {
            v2[2147483647] = (32767 > 42)
            fmt.Println(v3, v4)
        }
        fmt.Println(v4, fi23, v4)
        _ = fi23
    }
    fmt.Println(v19)
    var v24 uint32 = 699
    if (0 <= genFn0(int8(v24), 2, int32(v13))) {
        var v25 int32 = -1
        _ = v25
    } else {
        v4 = (v4 + "")
        v3 += v3
    }
    var v26 uint = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v26
    _ = v3
    _ = v4
}
