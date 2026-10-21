package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool; F1 int16; F2 int8 }
func genFn0(p0 uint64, p1 int16) bool {
    v7 := GenS1{F0: 105, F1: 2}
    var v8 int16 = (p1 / (((-(p1))) | 1))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return false
}
func genFn1() uint32 {
    var v9 float32 = ((1e10 / 1.0) / 3.14159)
    var v10 bool = false
    _ = v10
    _ = v9
    return 1
}

func main() {
    v1 := []int{}
    v2 := []string{}
    var v3 uint32 = 2
    var v4 int = -1
    v5 := []uint32{v3, v3}
    v6 := make(map[uint32]string)
    switch "hello" {
    case "a":
        v11 := GenS2{F0: (583 != -9223372036854775808), F1: 0}
        _ = v11
    case "foo bar":
        if (0.0 <= -2.25) {
            if false {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v12 uint64 = 1
                v13 := GenS3{F0: 255}
                _ = v12
                _ = v13
            } else {
                v4 = -2147483648
                v14 := GenS1{F0: 2, F1: 234}
                _ = v14
            }
            _ = len(v6)
            fmt.Println(v3, v3)
        }
    case "\n\t":
        fmt.Println(v4, v4)
    default:
        var v15 uint64 = 42
        _ = v15
    }
    var v16 uint = 1
    v4 = -1
    v16 = (v16 & v16)
    fmt.Println(v16)
    if ((3.14159 - (3.14159 / 1.0)) < (-2.25 / (3.14159 + (float32(1e10) * (0.1 / 0.1))))) {
        for fi17 := 0; fi17 < 18; fi17++ {
            var v18 string = "hello"
            _ = fi17
            _ = v18
        }
        v19 := GenS2{F0: (128 != (v4 % ((v4) | 1))), F1: -32768}
        _ = v19
    } else {
        v4 -= (v4 << ((v16) & 63))
        v20 := GenS2{F0: true, F1: 416}
        var v21 int16 = 0
        _ = v20
        _ = v21
    }
    if len(v1) > 0 {
        v1[0] = v4
    }
    v4 = -500
    switch ("foo bar" + "hello") {
    case "foo bar":
        v4 = 127
    default:
        for fi22 := 0; fi22 < 19; fi22++ {
            _ = len(v6)
            v2 = append(v2, ("" + "a"))
            _ = fi22
        }
    }
    fmt.Println(v16)
    var v23 int32 = 0
    switch (v4 << (((v16 % ((v16) | 1))) & 63)) {
    case 127:
        var v24 uint = v16
        _ = v24
    case 2147483647:
        var v25 int16 = 1000
        _ = v25
    case 32768:
        v26 := GenS4{F0: false, F1: -32768, F2: 42}
        _ = v26
    }
    if false {
        if false {
            var v27 uint8 = 128
            var v28 int16 = -319
            var v29 float32 = -2.25
            _ = v27
            _ = v28
            _ = v29
        } else {
            v30 := GenS1{F0: 67, F1: (v16 * (v16 % (((v16 & (v16 - v16))) | 1)))}
            var v31 uint16 = 84
            _ = v30
            _ = v31
        }
    } else {
        v3 -= v3
    }
    v4 += (v4 - v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
