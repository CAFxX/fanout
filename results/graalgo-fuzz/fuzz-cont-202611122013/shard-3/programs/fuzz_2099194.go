package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 bool }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 int16 }
func genFn0() int16 {
    v7 := GenS1{F0: 65535, F1: -535, F2: true}
    v8 := GenS1{F0: 65535, F1: 42, F2: (!false)}
    var v9 int = -1
    var v10 int8 = 74
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 uint64) int8 {
    fmt.Println(p0)
    var v11 uint32 = 42
    _ = p0
    _ = v11
    return 100
}

func main() {
    v1 := make(map[string]bool)
    var v2 int = 1
    var v3 uint32 = 2147483648
    var v4 int16 = -1
    var v5 int8 = 0
    var v6 uint32 = v3
    v6 -= 1
    switch 128 {
    case 32767:
        if (0.0 <= ((1.5 + 1e10) - 1.5)) {
            var v12 int32 = -367
            fmt.Println(v5, v6)
            fmt.Println(v12, v6, v5)
            _ = v12
        } else {
            v13 := GenS1{F0: v6, F1: 898, F2: (v3 != 2147483648)}
            _ = v13
        }
    case 9223372036854775807:
        v5 -= -1
    }
    v14 := GenS2{F0: false, F1: (v3 / ((42) | 1))}
    for fi15 := 0; fi15 < 12; fi15++ {
        v16 := []int{128, (fi15 | v2), (fi15 << ((65535) & 63)), (v2 * v2)}
        v2 -= 42
        var v17 uint32 = (v6 + v6)
        _ = fi15
        _ = v16
        _ = v17
    }
    delete(v1, ("hello" + ("foo bar" + "hello")))
    var v18 uint64 = 464
    if (genFn1((v18 / ((v18) | 1))) < (-(v5))) {
        v19 := GenS3{F0: (v4 - (v4 ^ v4))}
        _ = v19
    } else {
        var v20 int32 = 42
        _ = v20
    }
    v21 := GenS2{F0: (false && false), F1: v6}
    if true {
        v22 := GenS2{F0: (((v18 >> ((739) & 63)) >= 18446744073709551615) || false), F1: v6}
        _ = v22
    }
    switch 32767 {
    case -1:
        switch ("hello" + "") {
        case "a":
            var v23 int8 = (v5 % ((genFn1(v18)) | 1))
            _ = v23
        case "foo bar":
            v5 = v5
        }
    default:
        var v24 int64 = 0
        _ = v24
    }
    fmt.Println(v6, v4)
    for fi25 := 0; fi25 < 6; fi25++ {
        v2 -= v2
        delete(v1, "foo bar")
        _ = fi25
    }
    var v26 int32 = 102
    if false {
        v4 = genFn0()
    }
    v27 := GenS2{F0: false, F1: (v6 ^ (v3 & 0))}
    var v28 int64 = 796
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v21
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
