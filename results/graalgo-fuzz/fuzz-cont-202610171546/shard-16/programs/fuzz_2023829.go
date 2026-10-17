package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 string }
type GenS2 struct { F0 uint; F1 int8; F2 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint }
func genFn0(p0 int8, p1 int16, p2 uint64) bool {
    v5 := []int{}
    fmt.Println(p2, p2, p0)
    for fi6 := 0; fi6 < 13; fi6++ {
        p0 = (p0 % ((-128) | 1))
        var v7 uint64 = (p2 << ((1) & 63))
        if (4294967295 < 18446744073709551615) {
            p1 = 0
        } else {
            v7 = (v7 & (v7 | (^(p2))))
            v8 := GenS1{F0: p0, F1: fi6, F2: ("hello" + "foo bar")}
            if false {
                _ = len(v5)
                p2 -= 258
            } else {
                p0 += (p0 ^ p0)
            }
            _ = v8
        }
        _ = fi6
        _ = v7
    }
    if false {
        p2 = p2
        for fi9 := 0; fi9 < 7; fi9++ {
            if (!false) {
                var v10 float32 = ((float32((0.1 * ((0.1 * 1e10) - 1.5))) - ((1e10 + -2.25) / ((-2.25 - 3.14159) * 3.14159))) + 1.5)
                _ = v10
            }
            _ = fi9
        }
        switch ("a" + "foo bar") {
        case "hello":
            p1 = -406
        default:
            fmt.Println(p0, p1, p2)
        }
    } else {
        v11 := GenS1{F0: 100, F1: 128, F2: ("foo bar" + ("foo bar" + "hello"))}
        var v12 uint64 = (p2 / ((18446744073709551615) | 1))
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (141 != -2147483648)
}
func genFn1(p0 uint8) int8 {
    var v13 uint8 = 128
    v14 := []int{2147483647, -5}
    for fi15 := 0; fi15 < 7; fi15++ {
        v16 := GenS1{F0: -128, F1: 2147483647, F2: ("hello" + "hello")}
        if false {
            for fi17 := 0; fi17 < 16; fi17++ {
                var v18 int = fi15
                p0 -= v13
                _ = fi17
                _ = v18
            }
        }
        v13 = (p0 << ((941) & 63))
        _ = fi15
        _ = v16
    }
    if true {
        v19 := GenS2{F0: 4294967295, F1: 1, F2: 1}
        if len(v14) > 0 {
            v14[0] = 358
        }
        _ = v19
    }
    _ = p0
    _ = v13
    _ = v14
    return 0
}

func main() {
    v1 := []int{1, 781, 0}
    var v2 float64 = -2.25
    var v3 int64 = 865
    var v4 float64 = v2
    v20 := GenS2{F0: 122, F1: 1, F2: 1}
    v2 = 3.14159
    for fi21 := 0; fi21 < 2; fi21++ {
        fmt.Println(fi21, fi21, fi21)
        _ = fi21
    }
    var v22 uint64 = 251
    var v23 int = 42
    for fi24 := 0; fi24 < 20; fi24++ {
        v25 := GenS1{F0: -113, F1: v23, F2: (("a" + "hello") + "foo bar")}
        _ = fi24
        _ = v25
    }
    v26 := GenS1{F0: 127, F1: (-(v23)), F2: ""}
    for fi27 := 0; fi27 < 1; fi27++ {
        v2 = v4
        _ = fi27
    }
    if (65535 <= 949) {
        for fi28 := 0; fi28 < 19; fi28++ {
            v3 = (v3 % ((-9223372036854775808) | 1))
            _ = fi28
        }
        var v29 int16 = int16(v3)
        fmt.Println(v22, v22, v2)
        _ = v29
    } else {
        var v30 uint64 = v22
        v31 := GenS3{F0: 646}
        var v32 int64 = v3
        _ = v30
        _ = v31
        _ = v32
    }
    fmt.Println(v3, v4)
    if true {
        v33 := GenS4{F0: 65535}
        v4 = (1e10 / (1.5 + v4))
        var v34 int8 = -128
        _ = v33
        _ = v34
    }
    v2 -= (1e10 / ((1.5 * float64((0.0 / 1e10))) - float64((3.14159 - -2.25))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v4
}
