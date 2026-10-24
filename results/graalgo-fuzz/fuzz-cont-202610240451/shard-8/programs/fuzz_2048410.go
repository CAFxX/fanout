package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int }
type GenS2 struct { F0 uint8; F1 uint32 }

func main() {
    v1 := []string{"", "a", "foo bar"}
    var v2 uint = 444
    v3 := []uint32{2, 560, 391}
    var v4 string = ""
    for fi5 := 0; fi5 < 17; fi5++ {
        for fi6 := 0; fi6 < 5; fi6++ {
            v2 -= v2
            _ = fi6
        }
        switch "\n\t" {
        case "":
            v2 += v2
        case "foo bar":
            if true {
                var v7 int8 = -128
                v2 = v2
                _ = v7
            } else {
                _ = len(v1)
            }
        }
        var v8 int = fi5
        _ = fi5
        _ = v8
    }
    v2 = v2
    var v9 bool = ((3.14159 > 0.1) && true)
    fmt.Println(v9, v2, v4)
    var v10 uint8 = 0
    if (((((1e10 + (float64(((((3.14159 * 1.5) + 3.14159) * 3.14159) - ((-2.25 + 3.14159) - 0.1))) * ((((1.5 + ((((1e10 - 3.14159) + (1e10 - -2.25)) / float64(3.14159)) - float64((1.5 * ((((((1e10 + 1e10) * 0.1) * (1e10 / ((3.14159 * ((0.1 - 0.0) / 1e10)) + float32(float64(((0.1 + 1.5) / float32((float64(((1e10 * (float32(0.0) - (1e10 * 0.1))) + -2.25)) - (3.14159 / (0.0 + 0.1)))))))))) + float32((1e10 + 1e10))) * 0.1) - ((3.14159 + float32(float64(3.14159))) / -2.25)))))) - 1e10) * ((1.5 - -2.25) * 3.14159)) * ((0.1 + 0.1) * (((1e10 * float64(1.5)) + 1e10) + float64(((0.1 / ((-2.25 - (1e10 - -2.25)) / 3.14159)) + 1.5))))))) * (0.1 * float64((-2.25 + float32(0.1))))) * ((0.0 - float64(-2.25)) * (float64((1e10 + 3.14159)) - float64(0.0)))) + 1.5) != 1.5) {
        var v11 uint16 = 32767
        _ = v11
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 uint8 = (v10 & 158)
        fmt.Println(v4, v10)
        _ = fi12
        _ = v13
    }
    v9 = (((4294967295 < 326) || ((v10 << ((v2) & 63)) >= (v10 >> ((v2) & 63)))) && (1 >= 27))
    var v14 float64 = ((((-2.25 - 1e10) - (1e10 * 0.1)) * 0.1) / float64(0.1))
    for fi15 := 0; fi15 < 18; fi15++ {
        for fi16 := 0; fi16 < 9; fi16++ {
            switch "foo bar" {
            case "foo bar":
                v17 := GenS1{F0: 0, F1: 42, F2: fi15}
                _ = v17
            case "\n\t":
                v4 = (v4 + v4)
            case "a":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                if false {
                    var v18 int32 = 18
                    v9 = false
                    _ = v18
                } else {
                    fmt.Println(v10)
                }
                if ((v10 <= (v10 + v10)) && true) {
                    v10 = v10
                }
                if (v9 && v9) {
                    switch v4 {
                    case "":
                        v3 = append(v3, 4294967295)
                    default:
                        v19 := GenS2{F0: 128, F1: 4294967295}
                        _ = v19
                    }
                } else {
                    fmt.Println(v10, v4, v4)
                }
            }
            _ = len(v1)
            v4 = v4
            _ = fi16
        }
        _ = fi15
    }
    v20 := GenS2{F0: (v10 << (((^(v2))) & 63)), F1: 888}
    if ((v9 && (((v4 + v4) + ("a" + (v4 + ""))) >= "")) || true) {
        var v21 bool = (v10 < (v10 / ((v10) | 1)))
        _ = v21
    } else {
        if len(v1) > 0 {
            v1[0] = (v4 + (v4 + v4))
        }
    }
    var v22 uint16 = uint16(v10)
    var v23 int64 = -9223372036854775808
    v10 -= (v10 & 65)
    v10 -= (^(v10))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v9
}
