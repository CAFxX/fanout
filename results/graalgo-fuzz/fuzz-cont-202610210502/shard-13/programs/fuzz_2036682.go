package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 bool; F1 uint; F2 bool }
type GenS3 struct { F0 int8; F1 uint32 }
func genFn0() string {
    var v6 int32 = 955
    v6 = (v6 ^ (v6 | (v6 / (((v6 - v6)) | 1))))
    for fi7 := 0; fi7 < 8; fi7++ {
        fmt.Println(v6, v6)
        fmt.Println(v6, fi7)
        _ = fi7
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        for fi9 := 0; fi9 < 9; fi9++ {
            v10 := GenS1{F0: -605, F1: 0}
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    _ = v6
    return "a"
}
func genFn1(p0 int, p1 int64, p2 uint32) uint8 {
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 float32 = float32(3.14159)
        var v13 uint8 = 1
        v14 := GenS1{F0: 32767, F1: 44}
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    var v15 int32 = 0
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 220
}

func main() {
    var v1 uint8 = 42
    var v2 int64 = 9223372036854775807
    var v3 int16 = 32767
    var v4 float32 = float32((1e10 * float64(float32(0.1))))
    v5 := []string{"\n\t"}
    var v16 float32 = (v4 * v4)
    fmt.Println(v1, v16)
    if true {
        for fi17 := 0; fi17 < 6; fi17++ {
            v3 = (v3 % (((v3 - 995)) | 1))
            if len(v5) > 0 {
                v5[0] = "a"
            }
            _ = fi17
        }
        v2 -= -1
    }
    v2 -= -159
    for fi18 := 0; fi18 < 9; fi18++ {
        if true {
            var v19 int64 = (-(v2))
            var v20 float32 = float32(0.1)
            _ = v19
            _ = v20
        }
        v5 = append(v5, "\n\t")
        for fi21 := 0; fi21 < 11; fi21++ {
            _ = len(v5)
            _ = fi21
        }
        _ = fi18
    }
    if false {
        fmt.Println(v16, v2)
        if false {
            var v22 int32 = 2147483647
            v23 := GenS1{F0: v3, F1: 0}
            _ = v22
            _ = v23
        } else {
            v24 := GenS2{F0: (1.5 >= 1.5), F1: 255, F2: true}
            fmt.Println(v2)
            v16 += ((((v4 + v4) + float32(0.0)) + v4) + 1e10)
            _ = v24
        }
    } else {
        v2 += v2
        var v25 int = 9223372036854775807
        var v26 int = v25
        _ = v25
        _ = v26
    }
    fmt.Println(v16)
    switch 128 {
    case 2147483647:
        v27 := GenS3{F0: 98, F1: 42}
        _ = v27
    case -9223372036854775808:
        for fi28 := 0; fi28 < 16; fi28++ {
            if ((((("" + "hello") + "hello") + ("" + "")) < ("foo bar" + "\n\t")) && false) {
                var v29 uint64 = 1
                _ = v29
            } else {
                var v30 int16 = v3
                v31 := GenS2{F0: false, F1: 547, F2: (((1e10 * (1e10 + ((float64((float32((-2.25 - float64((float32(float64(0.0)) - -2.25)))) - v4)) / 1e10) / 1.5))) >= 1.5) || false)}
                _ = v30
                _ = v31
            }
            v2 -= (-(v2))
            _ = fi28
        }
    case -1:
        v1 -= (^(v1))
    }
    if (42 != 42) {
        if len(v5) > 0 {
            v5[0] = "hello"
        }
    }
    v32 := GenS3{F0: 127, F1: 46}
    v4 = float32(float64(v16))
    v3 -= (v3 ^ v3)
    for fi33 := 0; fi33 < 17; fi33++ {
        var v34 int64 = (v2 - v2)
        fmt.Println(v34, v3, v3)
        _ = fi33
        _ = v34
    }
    fmt.Println(v2, v16)
    if (!(false || ((0.1 - 1.5) <= (0.0 * (0.1 - 0.1))))) {
        switch ("" + ("foo bar" + ("hello" + ""))) {
        case "a":
            v16 += v16
        default:
            if len(v5) > 0 {
                v5[0] = "foo bar"
            }
        }
    } else {
        switch 147 {
        case -1:
            v35 := GenS3{F0: 127, F1: 1}
            _ = v35
        default:
            var v36 float32 = v16
            _ = v36
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
