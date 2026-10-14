package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 bool }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 int8; F1 int8 }

func main() {
    var v1 int = 32768
    v2 := make(map[int]bool)
    var v3 int64 = 101
    v4 := []uint32{2147483647, 4294967295}
    v5 := make(map[uint32]bool)
    var v6 uint = 42
    v7 := v5[430]
    if v7 {
        fmt.Println(v3, v1, v1)
        v8 := GenS1{F0: 0, F1: v6, F2: false}
        _ = v8
    } else {
        v7 = ("\n\t" <= "")
    }
    v7 = v7
    for fi9 := 0; fi9 < 12; fi9++ {
        for fi10 := 0; fi10 < 5; fi10++ {
            var v11 float32 = float32((((1.5 * (3.14159 + ((float64(0.0) / 0.1) * 0.1))) - 0.0) - float64(float32(float64(0.1)))))
            v1 = v1
            switch 32768 {
            case -9223372036854775808:
                v1 += fi10
            case 1:
                _ = len(v4)
            default:
                var v12 int = -2147483648
                v11 = (v11 - 3.14159)
                fmt.Println(fi10, v11)
                _ = v12
            }
            fmt.Println(fi9, v1, v3)
            _ = fi10
            _ = v11
        }
        var v13 bool = (!v7)
        _ = fi9
        _ = v13
    }
    v5[0] = (!true)
    var v14 int16 = -732
    if (("foo bar" + ((("foo bar" + "") + (("a" + ("" + ("\n\t" + "hello"))) + "")) + "")) >= "foo bar") {
        var v15 float32 = 3.14159
        _ = v15
    } else {
        if (-2.25 >= 1e10) {
            var v16 int16 = (v14 & v14)
            v1 += (v1 * 127)
            _ = v16
        }
        var v17 uint16 = 1
        _ = v17
    }
    _ = len(v4)
    delete(v2, 0)
    v18 := GenS1{F0: 239, F1: 4294967295, F2: v7}
    v19 := GenS2{F0: v6, F1: (v3 % ((v3) | 1))}
    for fi20 := 0; fi20 < 3; fi20++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi20
    }
    fmt.Println(v1)
    v14 -= v14
    if v7 {
        for fi21 := 0; fi21 < 16; fi21++ {
            fmt.Println(fi21, fi21, v6)
            if v7 {
                v1 = (-(v1))
                v1 -= 128
            } else {
                v22 := GenS3{F0: -20, F1: 127}
                _ = v22
            }
            _ = fi21
        }
    } else {
        var v23 uint64 = 239
        _ = v23
    }
    v24 := GenS3{F0: 0, F1: -86}
    for fi25 := 0; fi25 < 17; fi25++ {
        var v26 float32 = 3.14159
        _ = fi25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
