package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int16; F1 int16 }
type GenS7 struct { F0 uint64 }

func main() {
    v1 := []int{42, 9223372036854775807}
    var v2 uint16 = 42
    var v3 float64 = 0.1
    fmt.Println(v3, v2)
    var v4 int64 = int64(v2)
    v3 = (v3 / (float64((((((((((0.0 + 0.1) * 0.1) / 3.14159) * -2.25) / ((1.5 / 0.1) / 1.0)) * 1.5) - float32(0.1)) / ((((0.0 + 1.5) - 3.14159) - (0.1 * (3.14159 + (0.1 - 1e10)))) - 1.5)) - float32(float64(3.14159)))) / v3))
    for fi5 := 0; fi5 < 5; fi5++ {
        v4 = (v4 + v4)
        var v6 string = "a"
        _ = fi5
        _ = v6
    }
    switch 127 {
    case -701:
        if true {
            v4 -= v4
            v4 -= v4
            var v7 uint64 = 955
            _ = v7
        } else {
            v8 := GenS1{F0: 247, F1: -1}
            _ = v8
        }
    default:
        var v9 uint64 = 18446744073709551615
        _ = v9
    }
    var v10 float64 = float64(((1e10 + float32(0.1)) * ((1e10 + ((float32(v3) * (-2.25 + 0.1)) / 1.5)) - float32(v3))))
    v11 := GenS2{F0: -1, F1: -32768}
    v4 = v4
    v3 = v10
    v12 := GenS1{F0: 1, F1: 1}
    var v13 uint32 = 2147483647
    v14 := GenS3{F0: 32767, F1: (v2 | 440), F2: 18446744073709551615}
    for fi15 := 0; fi15 < 12; fi15++ {
        for fi16 := 0; fi16 < 9; fi16++ {
            switch ("" + "hello") {
            case "foo bar":
                v17 := GenS4{F0: false}
                _ = v17
            case "a":
                fmt.Println(fi16, v10, v10)
            }
            if ((v3 * 0.0) <= 0.1) {
                v3 = v10
                var v18 float64 = ((v3 / ((0.0 + ((((v10 / -2.25) + 0.0) + v10) + (0.0 - (((v10 * ((((0.1 / ((1.5 + 1e10) / (((v3 + float64(float32(v10))) + 1e10) / 1e10))) * float64((3.14159 * float32((1e10 - v3))))) * 0.1) / v3)) - 0.0) - 1.5)))) * v10)) - v10)
                var v19 bool = true
                _ = v18
                _ = v19
            }
            var v20 int8 = 56
            _ = fi16
            _ = v20
        }
        _ = fi15
    }
    switch "a" {
    case "foo bar":
        v2 = 277
    case "":
        if (-9223372036854775808 == -731) {
            if (v3 >= (v10 / (v10 - float64(0.0)))) {
                switch 2147483647 {
                case 32768:
                    var v21 int8 = 100
                    _ = v21
                case 2147483647:
                    var v22 int8 = -74
                    _ = v22
                case 0:
                    v10 += float64(((0.1 * ((((1.5 * (1e10 + (0.0 * 1.5))) + float32(((v10 * v10) * v3))) * 0.1) - (-2.25 / 1e10))) - -2.25))
                }
                if len(v1) > 0 {
                    v1[0] = -2147483648
                }
                v23 := GenS5{F0: v13}
                _ = v23
            } else {
                for fi24 := 0; fi24 < 5; fi24++ {
                    var v25 string = ("foo bar" + ("\n\t" + "a"))
                    v13 = 2
                    _ = fi24
                    _ = v25
                }
                v26 := GenS6{F0: -32768, F1: -1}
                _ = v26
            }
            v3 = v3
            var v27 int64 = 499
            _ = v27
        }
    case "hello":
        _ = len(v1)
    default:
        v10 -= (float64(-2.25) + -2.25)
    }
    v13 += (^(v13))
    v28 := GenS7{F0: 915}
    v13 += (v13 - (v13 << ((1) & 63)))
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
