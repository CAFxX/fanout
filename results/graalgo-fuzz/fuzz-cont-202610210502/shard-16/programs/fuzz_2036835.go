package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS5 struct { F0 int16; F1 string; F2 uint32 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]string)
    var v3 int = 32768
    var v4 uint8 = 127
    _ = len(v1)
    var v5 uint32 = 1
    var v6 float32 = ((1e10 * ((1e10 / 3.14159) + 1.5)) + -2.25)
    fmt.Println(v5)
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 12; fi8++ {
            var v9 uint16 = 748
            v4 += 42
            _ = fi8
            _ = v9
        }
        switch fi7 {
        case 416:
            var v10 uint32 = v5
            _ = v10
        default:
            v5 -= (v5 ^ (v5 - (v5 % ((438) | 1))))
        }
        for fi11 := 0; fi11 < 6; fi11++ {
            if (326 >= 0) {
                fmt.Println(v4, v4, v3)
            } else {
                fmt.Println(fi11)
            }
            _ = fi11
        }
        _ = fi7
    }
    if false {
        for fi12 := 0; fi12 < 11; fi12++ {
            var v13 float32 = 1.5
            v13 = (v6 * float32((float64((3.14159 * 0.0)) * (0.1 + 1e10))))
            _ = fi12
            _ = v13
        }
    } else {
        v5 += v5
        var v14 uint16 = 42
        _ = v14
    }
    if (-32768 <= 42) {
        fmt.Println(v3, v6, v3)
        switch ("" + (((("" + "\n\t") + "") + "\n\t") + "foo bar")) {
        case "\n\t":
            for fi15 := 0; fi15 < 5; fi15++ {
                v16 := GenS1{F0: uint16(v4)}
                _ = fi15
                _ = v16
            }
        default:
            if (true && true) {
                if false {
                    fmt.Println(v4, v6, v6)
                    var v17 int = (v3 % ((v3) | 1))
                    var v18 int32 = 0
                    _ = v17
                    _ = v18
                } else {
                    if false {
                        v19 := GenS2{F0: 255, F1: 0}
                        _ = v19
                    }
                    v20 := GenS1{F0: 32767}
                    v21 := GenS3{F0: (v5 - (^(v5))), F1: -2147483648}
                    _ = v20
                    _ = v21
                }
                v22 := GenS2{F0: 0, F1: int16(v4)}
                _ = v22
            }
        }
    }
    switch ((("hello" + "") + ("hello" + ("" + ""))) + "hello") {
    case "":
        if (!true) {
            v23 := GenS4{F0: 128, F1: false, F2: v4}
            _ = v23
        } else {
            v24 := GenS4{F0: v4, F1: false, F2: 170}
            v4 = v4
            _ = v24
        }
    case "foo bar":
        v25 := GenS5{F0: 1000, F1: "foo bar", F2: (^(v5))}
        _ = v25
    default:
        var v26 float32 = 0.0
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
