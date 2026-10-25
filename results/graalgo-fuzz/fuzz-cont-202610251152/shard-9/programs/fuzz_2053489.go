package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 uint8 }
func genFn0(p0 bool, p1 uint16) uint8 {
    p0 = false
    p1 += 255
    p1 -= (p1 * 670)
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 uint = 65535
    v2 := make(map[string]uint32)
    v3 := make(map[string]uint32)
    v4 := []int{32767, -970, -9223372036854775808}
    v1 = (v1 % ((v1) | 1))
    v5 := GenS1{F0: 1, F1: -436, F2: 32767}
    var v6 float64 = ((3.14159 + (float64(float32(0.1)) + (1e10 * (0.1 + (float64(1e10) / (float64((0.0 + (0.0 * 1e10))) + ((((0.0 - float64(1.5)) * float64(1e10)) + 3.14159) + float64(1.5)))))))) - float64(-2.25))
    switch -1 {
    case -11:
        var v7 uint16 = 2
        _ = v7
    default:
        var v8 float32 = (0.1 / 1.5)
        _ = v8
    }
    if false {
        var v9 int32 = 0
        _ = v9
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        v11 := GenS1{F0: -1, F1: 0, F2: -553}
        for fi12 := 0; fi12 < 3; fi12++ {
            v1 += 255
            _ = fi12
        }
        _ = fi10
        _ = v11
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        if (1e10 >= ((0.0 * (3.14159 * float32(v6))) + -2.25)) {
            if (65535 >= 42) {
                fmt.Println(v6, v1, v1)
            } else {
                if false {
                    v6 = float64(0.1)
                } else {
                    v1 += (v1 & (v1 / (((^(v1))) | 1)))
                    v6 += (1.5 + v6)
                }
            }
        }
        fmt.Println(v6, fi13, v1)
        _ = fi13
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        if ((0 >= 0) || false) {
            v15 := GenS1{F0: -14, F1: -32768, F2: 1000}
            for fi16 := 0; fi16 < 7; fi16++ {
                v6 += 0.0
                if (!(!(3.14159 >= 0.0))) {
                    v1 = (v1 ^ (v1 >> (((v1 * v1)) & 63)))
                }
                _ = fi16
            }
            v17 := GenS2{F0: v1}
            _ = v15
            _ = v17
        }
        _ = fi14
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        var v19 uint16 = 0
        switch "" {
        case "":
            var v20 int64 = 206
            _ = v20
        default:
            v6 = v6
        }
        _ = fi18
        _ = v19
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v6, v6)
    for fi21 := 0; fi21 < 13; fi21++ {
        delete(v2, (("a" + "\n\t") + ("a" + ("a" + ("hello" + ("" + "foo bar"))))))
        v22 := GenS1{F0: 29, F1: -32768, F2: 32767}
        var v23 int16 = 42
        _ = fi21
        _ = v22
        _ = v23
    }
    var v24 uint16 = 82
    v6 -= 1.5
    switch 391 {
    case 128:
        v25 := GenS3{F0: -1}
        _ = v25
    case 2147483647:
        if true {
            var v26 uint64 = 12
            v27 := GenS2{F0: (v1 - (v1 + v1))}
            _ = v26
            _ = v27
        }
    }
    switch 1 {
    case -799:
        switch "" {
        case "foo bar":
            v3[("\n\t" + "foo bar")] = 2147483647
        case "\n\t":
            var v28 int = -980
            _ = v28
        }
    case -9223372036854775808:
        v29 := GenS3{F0: -1}
        _ = v29
    case 1:
        if true {
            var v30 uint = (v1 / ((v1) | 1))
            v31 := GenS4{F0: 430, F1: 158}
            _ = v30
            _ = v31
        }
    default:
        v32 := v3[(("a" + ("foo bar" + ("\n\t" + ""))) + "a")]
        _ = v32
    }
    v24 = (v24 + (v24 + 255))
    v33 := GenS4{F0: 0, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
