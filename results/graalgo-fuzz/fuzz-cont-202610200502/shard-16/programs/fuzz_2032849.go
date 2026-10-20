package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int8; F1 int; F2 int32 }
type GenS3 struct { F0 uint; F1 bool }
func genFn0(p0 string) bool {
    for fi4 := 0; fi4 < 12; fi4++ {
        var v5 int32 = -237
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 17; fi6++ {
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 uint8 = 127
            var v9 uint = uint(v8)
            fmt.Println(v9)
            v10 := []int{-999, (fi6 * fi7), (fi7 % ((127) | 1))}
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    p0 = "hello"
    if false {
        var v11 uint16 = 1
        _ = v11
    } else {
        fmt.Println(p0)
    }
    v12 := []int{4294967295, -955, 32768, 128}
    _ = p0
    _ = v12
    return (1 < 253)
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = (1.5 - float32(((((((1.5 - (0.1 + 1e10)) / -2.25) / 1e10) * -2.25) - 3.14159) / -2.25)))
    v3 := []string{"", "", "", (("\n\t" + "a") + "foo bar")}
    var v13 uint16 = 0
    var v14 int64 = 420
    for fi15 := 0; fi15 < 17; fi15++ {
        switch "\n\t" {
        case "a":
            _ = len(v1)
        case "\n\t":
            for fi16 := 0; fi16 < 14; fi16++ {
                if (!true) {
                    if false {
                        fmt.Println(v13, v14, fi16)
                    }
                    fmt.Println(fi16)
                    delete(v1, 128)
                } else {
                    v17 := GenS1{F0: 2147483648, F1: 1, F2: 127}
                    _ = v17
                }
                v13 -= v13
                v18 := GenS1{F0: 65535, F1: 18446744073709551615, F2: 114}
                _ = fi16
                _ = v18
            }
        case "foo bar":
            v1[fi15] = 2147483648
        default:
            var v19 int16 = 723
            _ = v19
        }
        _ = fi15
    }
    switch "a" {
    case "":
        switch 725 {
        case 127:
            switch "" {
            case "foo bar":
                v2 += (v2 / float32(float64(v2)))
            case "\n\t":
                var v20 float32 = float32((1.5 - ((0.1 - (1.5 - 0.1)) * (1e10 / 1.5))))
                _ = v20
            default:
                var v21 int64 = -1
                _ = v21
            }
        case 128:
            v22 := GenS2{F0: 42, F1: -515, F2: 0}
            _ = v22
        case -777:
            if true {
                if true {
                    var v23 int64 = 0
                    _ = v23
                } else {
                    v24 := GenS3{F0: 385, F1: true}
                    switch ("foo bar" + "") {
                    case "hello":
                        var v25 int8 = 42
                        _ = v25
                    case "\n\t":
                        v13 = (v13 + (v13 ^ v13))
                    }
                    v26 := GenS3{F0: 18446744073709551615, F1: (3.14159 > v2)}
                    _ = v24
                    _ = v26
                }
                v14 += (v14 ^ v14)
            } else {
                v27 := GenS3{F0: 272, F1: false}
                switch ("" + "a") {
                case "a":
                    var v28 uint = 94
                    _ = v28
                default:
                    var v29 bool = true
                    _ = v29
                }
                _ = v27
            }
        default:
            for fi30 := 0; fi30 < 18; fi30++ {
                if true {
                    v3 = append(v3, "\n\t")
                    var v31 uint64 = 42
                    _ = v31
                }
                var v32 int = (fi30 << ((4294967295) & 63))
                _ = fi30
                _ = v32
            }
        }
    case "a":
        if ((true && (!true)) && (!(-635 <= -9223372036854775808))) {
            _ = len(v3)
            _ = len(v1)
            var v33 int8 = 0
            _ = v33
        } else {
            if (false && true) {
                var v34 int32 = 2147483647
                _ = v34
            }
            v35 := v1[-2147483648]
            v36 := GenS2{F0: 54, F1: 32768, F2: 0}
            _ = v35
            _ = v36
        }
    case "\n\t":
        fmt.Println(v2)
    default:
        v2 = v2
    }
    var v37 uint8 = 1
    var v38 uint = 42
    v14 = (-(v14))
    var v39 int8 = 42
    v40 := GenS2{F0: (v39 % ((v39) | 1)), F1: 184, F2: -1}
    v39 = 0
    var v41 uint64 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v40
    _ = v41
}
