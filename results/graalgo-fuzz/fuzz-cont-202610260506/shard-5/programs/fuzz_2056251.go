package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 uint; F2 int32 }
type GenS3 struct { F0 string; F1 int16 }
type GenS4 struct { F0 uint64; F1 uint32; F2 uint16 }
func genFn0() uint8 {
    var v7 int = 4294967295
    if (172 != 1) {
        var v8 uint = 42
        var v9 uint16 = 32767
        fmt.Println(v7, v9)
        _ = v8
        _ = v9
    }
    switch (-(v7)) {
    case -439:
        fmt.Println(v7)
    case 9223372036854775807:
        v7 -= v7
    default:
        v7 -= v7
    }
    _ = v7
    return 128
}
func genFn1(p0 uint16, p1 uint32) uint16 {
    var v10 uint16 = p0
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v10
    return 255
}

func main() {
    var v1 int16 = 737
    var v2 string = ""
    v3 := []string{v2, v2}
    var v4 uint64 = 820
    var v5 string = (v2 + ("hello" + v2))
    v6 := []string{("hello" + (("" + ((((("foo bar" + "hello") + v5) + v5) + "a") + "\n\t")) + v2))}
    fmt.Println(v1)
    v5 = (v5 + v5)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println(v2, v4, v2)
    if true {
        _ = len(v3)
        switch 2147483647 {
        case -1:
            var v11 uint16 = 0
            _ = v11
        case 32767:
            v12 := GenS1{F0: 1000}
            _ = v12
        }
        v13 := GenS2{F0: 42, F1: 1, F2: -503}
        _ = v13
    } else {
        fmt.Println(v4, v2)
    }
    var v14 bool = false
    v1 = (-(v1))
    v3 = append(v3, v2)
    if false {
        switch -1 {
        case 0:
            v1 = (v1 + 32767)
        case 128:
            for fi15 := 0; fi15 < 20; fi15++ {
                v16 := GenS3{F0: (v5 + v5), F1: 32767}
                v6 = append(v6, "")
                _ = fi15
                _ = v16
            }
        default:
            _ = len(v6)
        }
    } else {
        for fi17 := 0; fi17 < 9; fi17++ {
            var v18 float64 = float64(float32(3.14159))
            _ = fi17
            _ = v18
        }
        fmt.Println(v5)
    }
    v19 := GenS4{F0: (v4 << ((109) & 63)), F1: uint32(v4), F2: 0}
    var v20 uint32 = 0
    v2 = (v5 + v5)
    var v21 int8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
