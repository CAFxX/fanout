package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 int32; F2 bool }
type GenS3 struct { F0 int16; F1 bool; F2 uint64 }
type GenS4 struct { F0 int8; F1 int16; F2 int }
type GenS5 struct { F0 int64; F1 uint32 }
type GenS6 struct { F0 int8; F1 uint }
func genFn0(p0 string, p1 int16) int {
    var v5 uint16 = 136
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 float32 = 1e10
        var v8 float32 = (v7 - 0.1)
        _ = fi6
        _ = v7
        _ = v8
    }
    if (255 > 127) {
        fmt.Println(v5, p1)
        v5 += (v5 << ((170) & 63))
        var v9 uint64 = 2
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v5
    return 9223372036854775807
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[int]string)
    v3 := []string{("foo bar" + "\n\t"), "foo bar", ("foo bar" + "\n\t"), "a"}
    var v4 float64 = (((((v1 * v1) * v1) - ((v1 - float64(1e10)) * float64(1.5))) + (v1 + v1)) - v1)
    if (!(917 != -2147483648)) {
        v10 := GenS1{F0: -789}
        var v11 int32 = -2147483648
        _ = v10
        _ = v11
    } else {
        v12 := GenS1{F0: -1}
        if len(v3) > 0 {
            v3[0] = "hello"
        }
        if (false || ((!false) || true)) {
            v13 := GenS2{F0: 2, F1: 0, F2: true}
            if false {
                if true {
                    var v14 string = "foo bar"
                    for fi15 := 0; fi15 < 5; fi15++ {
                        fmt.Println(v4)
                        v14 = ""
                        _ = fi15
                    }
                    _ = v14
                }
            }
            _ = v13
        } else {
            var v16 float64 = v1
            v17 := GenS2{F0: 0, F1: -27, F2: (2 == 0)}
            _ = v16
            _ = v17
        }
        _ = v12
    }
    v18 := GenS3{F0: -32768, F1: true, F2: 543}
    if false {
        v4 -= -2.25
        var v19 float32 = 1e10
        _ = v19
    } else {
        if false {
            v4 = (v1 * (v4 + v1))
            if len(v3) > 0 {
                v3[0] = ("a" + "foo bar")
            }
            if len(v3) > 0 {
                v3[0] = ("hello" + (((("a" + "hello") + "") + ("foo bar" + ("\n\t" + "a"))) + "foo bar"))
            }
        }
        _ = len(v2)
    }
    switch 698 {
    case -1:
        v4 += v4
    case 32768:
        v20 := GenS3{F0: -103, F1: false, F2: 2}
        _ = v20
    case 0:
        v4 = ((1e10 * 3.14159) + ((v1 + v4) * v4))
    }
    if ((!true) || false) {
        if (2147483647 == 1) {
            v21 := GenS4{F0: 0, F1: 42, F2: 1}
            for fi22 := 0; fi22 < 12; fi22++ {
                v1 += 0.1
                v23 := GenS5{F0: -906, F1: 4294967295}
                _ = fi22
                _ = v23
            }
            _ = v21
        } else {
            var v24 int16 = 1000
            _ = v24
        }
        v25 := GenS6{F0: 1, F1: 259}
        v1 += v4
        _ = v25
    } else {
        v4 += ((-2.25 - v4) - (float64((0.1 / 1e10)) + v4))
        v4 = v4
    }
    v4 = v4
    var v26 float32 = 0.0
    for fi27 := 0; fi27 < 4; fi27++ {
        fmt.Println(v1)
        var v28 float64 = (v4 / ((v1 / v4) / 0.1))
        _ = fi27
        _ = v28
    }
    var v29 int16 = 32767
    for fi30 := 0; fi30 < 8; fi30++ {
        fmt.Println(v29, v29, v4)
        v31 := v2[fi30]
        _ = fi30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v4
}
