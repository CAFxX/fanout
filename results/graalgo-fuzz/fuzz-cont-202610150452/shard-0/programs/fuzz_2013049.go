package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 uint8; F2 int8 }
type GenS3 struct { F0 int8; F1 int; F2 uint32 }
type GenS4 struct { F0 string; F1 uint64; F2 uint }
type GenS5 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS6 struct { F0 int64; F1 uint }
func genFn0(p0 int32) int64 {
    var v5 string = "foo bar"
    var v6 uint64 = 4294967295
    _ = p0
    _ = v5
    _ = v6
    return -1
}

func main() {
    v1 := []string{"hello", "foo bar", "a", "\n\t"}
    var v2 int8 = -1
    var v3 uint64 = 107
    var v4 bool = (2147483648 == 845)
    switch 492 {
    case 4294967295:
        for fi7 := 0; fi7 < 1; fi7++ {
            switch ("hello" + ("a" + "hello")) {
            case "foo bar":
                var v8 uint32 = 4294967295
                _ = v8
            case "hello":
                for fi9 := 0; fi9 < 17; fi9++ {
                    var v10 uint16 = 176
                    _ = fi9
                    _ = v10
                }
            }
            if true {
                v2 = 1
            } else {
                switch (^(fi7)) {
                case 32768:
                    v11 := GenS1{F0: 959}
                    _ = v11
                case 629:
                    v4 = (0.1 <= (((((((1.5 - (((((-2.25 * ((0.1 / 3.14159) + (1.5 / 0.1))) - ((0.1 + 1.5) / (float32((1e10 + -2.25)) + 0.1))) - (-2.25 + (1.5 + ((float32(-2.25) / float32(((((float64(float32((0.1 + -2.25))) + -2.25) + -2.25) + 1.5) + float64(((-2.25 * (float32((0.0 * 1e10)) + 1.5)) / 1.0))))) * ((3.14159 - 1e10) - 0.1))))) - ((0.1 / 1.0) + 0.1)) + 3.14159)) * (-2.25 + (-2.25 - 0.1))) / -2.25) - float32((1.5 / 1.0))) / 0.1) + -2.25) - (3.14159 * 1.5)))
                case -1:
                    v4 = ((!v4) && true)
                }
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
            }
            v12 := GenS2{F0: 255, F1: 2, F2: 42}
            _ = fi7
            _ = v12
        }
    case -541:
        var v13 float32 = 3.14159
        _ = v13
    }
    v1 = append(v1, ("" + ("\n\t" + "a")))
    fmt.Println(v3)
    fmt.Println(v4)
    var v14 int16 = 1
    var v15 uint16 = 42
    fmt.Println(v2, v3, v3)
    v14 = (v14 << ((678) & 63))
    v1 = append(v1, "")
    if (-2147483648 > 162) {
        for fi16 := 0; fi16 < 2; fi16++ {
            v17 := GenS3{F0: v2, F1: fi16, F2: 632}
            var v18 int32 = 42
            _ = fi16
            _ = v17
            _ = v18
        }
        fmt.Println(v14, v4, v3)
    } else {
        v19 := GenS4{F0: ("\n\t" + ("a" + "hello")), F1: v3, F2: 255}
        _ = v19
    }
    for fi20 := 0; fi20 < 19; fi20++ {
        var v21 uint16 = (v15 | (^(v15)))
        v15 -= (v21 ^ v21)
        _ = fi20
        _ = v21
    }
    fmt.Println(v15)
    for fi22 := 0; fi22 < 8; fi22++ {
        if len(v1) > 0 {
            v1[0] = ("" + "")
        }
        switch 128 {
        case 9223372036854775807:
            for fi23 := 0; fi23 < 8; fi23++ {
                fmt.Println(fi23, fi23)
                _ = fi23
            }
        default:
            switch fi22 {
            case 128:
                if v4 {
                    v24 := GenS5{F0: 42, F1: -467, F2: 1}
                    fmt.Println(v3)
                    var v25 int16 = v14
                    _ = v24
                    _ = v25
                } else {
                    v15 += (v15 % ((v15) | 1))
                    if ((false || (3.14159 >= 1e10)) && (!(v3 >= (v3 / ((v3) | 1))))) {
                        if len(v1) > 0 {
                            v1[0] = "foo bar"
                        }
                        v14 = v14
                        v14 -= int16(fi22)
                    }
                }
            case 9223372036854775807:
                switch (fi22 / (((fi22 >> ((0) & 63))) | 1)) {
                case 32767:
                    fmt.Println(v14, v14, v4)
                default:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            default:
                v15 += (v15 | 2)
            }
        }
        _ = fi22
    }
    for fi26 := 0; fi26 < 4; fi26++ {
        for fi27 := 0; fi27 < 10; fi27++ {
            v28 := GenS6{F0: 1, F1: 2}
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
