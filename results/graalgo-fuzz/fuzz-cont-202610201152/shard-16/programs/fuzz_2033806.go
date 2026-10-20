package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 int16; F1 uint64; F2 int16 }
type GenS5 struct { F0 int64; F1 uint32 }
func genFn0() uint16 {
    var v4 uint64 = 18446744073709551615
    var v5 int32 = 301
    var v6 float32 = (1.5 / float32(1.5))
    _ = v4
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 uint = 1
    var v2 int32 = -422
    var v3 float64 = (0.0 - 3.14159)
    v7 := []int{127, -2147483648, 32768, 657}
    for fi8 := 0; fi8 < 12; fi8++ {
        switch fi8 {
        case 9223372036854775807:
            v3 = -2.25
        }
        v2 = (-(v2))
        _ = fi8
    }
    v9 := GenS1{F0: 21}
    for fi10 := 0; fi10 < 7; fi10++ {
        if (((v2 - 591) >= v2) && (false && true)) {
            if (float32(v3) >= (float32(0.0) * (((3.14159 / 1.0) - ((((((3.14159 * ((3.14159 - 3.14159) * 0.1)) + 3.14159) + 1.5) - 0.0) * 0.1) + 1e10)) - ((0.0 / 1.5) / (1e10 * (1.5 * (0.0 - 3.14159))))))) {
                v3 -= 1e10
            } else {
                v1 += v1
                v2 = (v2 / (((v2 >> (((v1 + (v1 - v1))) & 63))) | 1))
                v2 = v2
            }
        } else {
            switch (fi10 * -1) {
            case -916:
                var v11 uint = (v1 + v1)
                _ = v11
            default:
                switch fi10 {
                case -93:
                    switch "a" {
                    case "foo bar":
                        var v12 int = 9223372036854775807
                        _ = v12
                    case "":
                        v13 := GenS1{F0: 128}
                        _ = v13
                    }
                case -977:
                    fmt.Println(v1, v3, v3)
                case 214:
                    var v14 int8 = 42
                    _ = v14
                }
            }
            v3 -= (1e10 + (float64(3.14159) - v3))
        }
        _ = fi10
    }
    var v15 uint32 = 65535
    _ = len(v7)
    if ((false && ((("hello" + "a") + (("\n\t" + "hello") + (("foo bar" + ("" + ("" + ("hello" + "")))) + "foo bar"))) != "a")) && ((0.0 / 3.14159) < 1.5)) {
        if true {
            switch 778 {
            case -1:
                switch 2147483647 {
                case 725:
                    if true {
                        v16 := GenS2{F0: 0}
                        fmt.Println(v1, v1)
                        var v17 uint = (v1 % (((v1 >> (((v1 >> ((255) & 63))) & 63))) | 1))
                        _ = v16
                        _ = v17
                    }
                case -256:
                    v18 := GenS3{F0: "", F1: 0}
                    _ = v18
                }
            case -9223372036854775808:
                v15 -= (v15 | 699)
            case 32767:
                v2 += v2
            }
        } else {
            var v19 uint16 = 65535
            var v20 uint8 = 128
            v21 := GenS3{F0: (("\n\t" + "") + ""), F1: -9223372036854775808}
            _ = v19
            _ = v20
            _ = v21
        }
        var v22 uint32 = (^(v15))
        _ = v22
    } else {
        if (((("a" + "hello") + "hello") + "") <= "") {
            v15 += (v15 * 947)
        }
        if len(v7) > 0 {
            v7[0] = -2147483648
        }
    }
    if true {
        fmt.Println(v2, v3, v15)
        v23 := GenS2{F0: 1}
        if false {
            for fi24 := 0; fi24 < 4; fi24++ {
                var v25 uint = (^(v1))
                v1 = v25
                _ = fi24
                _ = v25
            }
            v26 := GenS4{F0: -371, F1: 2, F2: 42}
            v2 -= v2
            _ = v26
        }
        _ = v23
    }
    if (179 >= 0) {
        var v27 uint16 = 888
        var v28 int64 = -9223372036854775808
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 15; fi29++ {
        v30 := GenS1{F0: 1}
        _ = fi29
        _ = v30
    }
    for fi31 := 0; fi31 < 4; fi31++ {
        if true {
            v3 += v3
            v1 = v1
            var v32 int = fi31
            _ = v32
        } else {
            for fi33 := 0; fi33 < 4; fi33++ {
                v34 := GenS2{F0: 720}
                switch (-(fi33)) {
                case 128:
                    var v35 int64 = 148
                    _ = v35
                case -1:
                    v36 := GenS5{F0: 0, F1: (^(v15))}
                    _ = v36
                default:
                    var v37 uint8 = 128
                    _ = v37
                }
                v38 := GenS4{F0: 1, F1: 42, F2: 0}
                _ = fi33
                _ = v34
                _ = v38
            }
        }
        _ = fi31
    }
    var v39 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v39
    _ = v7
    _ = v9
}
