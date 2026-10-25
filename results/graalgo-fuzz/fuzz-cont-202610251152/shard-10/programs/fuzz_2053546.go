package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint8; F1 string; F2 int32 }
type GenS3 struct { F0 uint32 }
func genFn0() int64 {
    var v5 bool = true
    if false {
        if (1e10 != (0.0 + 1.5)) {
            v6 := []int{32767, 32767, -9223372036854775808}
            v7 := GenS1{F0: (!v5), F1: 1, F2: -1}
            _ = v6
            _ = v7
        }
    } else {
        for fi8 := 0; fi8 < 12; fi8++ {
            v5 = (-2.25 < 0.0)
            switch fi8 {
            case 2147483647:
                v9 := []int{(fi8 / ((fi8) | 1)), fi8, -522}
                _ = v9
            }
            _ = fi8
        }
    }
    switch 0 {
    case 423:
        v5 = false
    default:
        for fi10 := 0; fi10 < 1; fi10++ {
            if v5 {
                if (158 > -1) {
                    var v11 float32 = (3.14159 * 0.0)
                    var v12 float64 = (0.1 + float64((v11 + 1e10)))
                    for fi13 := 0; fi13 < 14; fi13++ {
                        switch "\n\t" {
                        case "a":
                            v5 = (v5 || v5)
                        }
                        _ = fi13
                    }
                    _ = v11
                    _ = v12
                }
                v14 := []int{(fi10 - -321), 487, (fi10 << ((850) & 63)), fi10}
                _ = v14
            }
            _ = fi10
        }
    }
    v5 = (float32(0.1) != (float32((((1e10 - (0.0 - 3.14159)) / 1e10) / 1.0)) - 0.1))
    fmt.Println(v5)
    _ = v5
    return 352
}

func main() {
    var v1 int = 0
    var v2 uint16 = 2
    v3 := make(map[int]uint32)
    v4 := []int{(v1 << ((2) & 63))}
    if (-2.25 == 0.1) {
        if false {
            v1 = v1
            for fi15 := 0; fi15 < 1; fi15++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v16 := GenS2{F0: 255, F1: "hello", F2: 2147483647}
                _ = fi15
                _ = v16
            }
            for fi17 := 0; fi17 < 19; fi17++ {
                v18 := GenS1{F0: (true || (1 == 489)), F1: 223, F2: 0}
                v19 := GenS1{F0: false, F1: 18446744073709551615, F2: 390}
                _ = fi17
                _ = v18
                _ = v19
            }
        }
        fmt.Println(v1)
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        v1 = 0
        if (-1 < int16(fi20)) {
            for fi21 := 0; fi21 < 14; fi21++ {
                _ = len(v4)
                var v22 int = 491
                v23 := GenS2{F0: 161, F1: "foo bar", F2: -2147483648}
                _ = fi21
                _ = v22
                _ = v23
            }
        } else {
            if (!true) {
                fmt.Println(v1)
            }
            switch fi20 {
            case 32768:
                var v24 int16 = 1
                _ = v24
            case 1:
                var v25 int16 = 42
                _ = v25
            default:
                if len(v4) > 0 {
                    v4[0] = v1
                }
            }
            v26 := GenS2{F0: uint8(v2), F1: ("" + (((("foo bar" + "foo bar") + "\n\t") + "foo bar") + "a")), F2: -2147483648}
            _ = v26
        }
        _ = fi20
    }
    var v27 int16 = -1
    if (1 >= 18446744073709551615) {
        var v28 uint16 = v2
        switch "" {
        case "\n\t":
            switch v1 {
            case 2147483647:
                var v29 float64 = 0.1
                _ = v29
            default:
                var v30 uint64 = 89
                _ = v30
            }
        case "foo bar":
            v31 := GenS2{F0: 127, F1: ("" + "\n\t"), F2: 383}
            _ = v31
        default:
            var v32 float64 = 3.14159
            _ = v32
        }
        _ = v28
    } else {
        switch "foo bar" {
        case "hello":
            v2 -= (^(v2))
        default:
            switch ("foo bar" + (("foo bar" + ("" + "")) + "hello")) {
            case "\n\t":
                v33 := GenS3{F0: 0}
                _ = v33
            case "a":
                var v34 float32 = (float32((1e10 / (1.5 / 1e10))) + (0.1 + (1.5 / float32((1.5 - ((-2.25 + 0.1) * 1e10))))))
                _ = v34
            case "":
                var v35 uint64 = 497
                _ = v35
            }
        }
        for fi36 := 0; fi36 < 4; fi36++ {
            _ = len(v4)
            _ = fi36
        }
    }
    v27 -= v27
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v37 uint32 = 42
    v37 -= uint32(v2)
    delete(v3, 42)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v37
    _ = v4
}
