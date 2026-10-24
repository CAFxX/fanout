package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 uint8; F1 int64; F2 string }
func genFn0() uint16 {
    var v6 uint64 = 2
    var v7 int8 = -128
    var v8 float64 = -2.25
    if (-525 >= 9223372036854775807) {
        v6 += (v6 | v6)
        v7 += -81
    } else {
        v9 := GenS1{F0: 1}
        v7 += (v7 + v7)
        _ = v9
    }
    _ = v6
    _ = v7
    _ = v8
    return 1
}
func genFn1() int16 {
    v10 := GenS1{F0: 368}
    var v11 uint16 = 320
    v11 = v11
    v12 := GenS2{F0: -335}
    _ = v10
    _ = v11
    _ = v12
    return -1
}

func main() {
    var v1 int8 = -95
    var v2 int32 = 687
    v3 := []uint32{2147483648, 822, 153}
    var v4 int64 = 9223372036854775807
    var v5 float32 = (((1.5 - (0.1 * ((1e10 * (1e10 + (1.5 + (0.0 * 3.14159)))) + 0.0))) - 3.14159) - float32((((1e10 / (3.14159 / (1.5 - 3.14159))) * 0.0) * (0.1 - ((float64(0.0) * (1.5 + 3.14159)) - (1.5 + float64(float32((0.0 + -2.25)))))))))
    var v13 uint8 = 217
    v14 := GenS2{F0: 9223372036854775807}
    switch 110 {
    case -105:
        var v15 int = -2147483648
        _ = v15
    }
    if (!(1e10 > 3.14159)) {
        switch ("foo bar" + ((("a" + ("\n\t" + "")) + ((((("foo bar" + "a") + (("\n\t" + "") + "hello")) + ((((("" + "a") + ("foo bar" + ("a" + "\n\t"))) + (("foo bar" + "foo bar") + "\n\t")) + "\n\t") + "")) + ("hello" + ("a" + ""))) + (("" + "\n\t") + "hello"))) + "")) {
        case "foo bar":
            for fi16 := 0; fi16 < 17; fi16++ {
                for fi17 := 0; fi17 < 2; fi17++ {
                    fmt.Println(v2, v2)
                    var v18 uint16 = genFn0()
                    _ = fi17
                    _ = v18
                }
                var v19 float64 = (((0.0 + 0.1) + 0.0) + 1e10)
                _ = fi16
                _ = v19
            }
        case "\n\t":
            var v20 uint64 = 18446744073709551615
            _ = v20
        case "a":
            v21 := GenS3{F0: 133, F1: 127}
            _ = v21
        default:
            if ((v2 % (((v2 >> ((4294967295) & 63))) | 1)) > v2) {
                v13 = (v13 + 0)
                switch "a" {
                case "hello":
                    if true {
                        fmt.Println(v2, v13)
                    } else {
                        var v22 int64 = 9223372036854775807
                        _ = v22
                    }
                case "":
                    fmt.Println(v5, v5, v1)
                default:
                    v23 := GenS3{F0: 0, F1: 42}
                    _ = v23
                }
            }
        }
    }
    v24 := GenS4{F0: v13, F1: -368, F2: "foo bar"}
    v13 -= (v13 ^ (v13 - v13))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v5 += v5
    if (147 > 42) {
        v25 := GenS3{F0: 42, F1: 1}
        var v26 uint32 = 65535
        _ = v25
        _ = v26
    }
    if true {
        switch 9223372036854775807 {
        case -1:
            v2 -= (v2 + v2)
        default:
            if false {
                v27 := GenS2{F0: -1}
                _ = v27
            }
        }
        var v28 int64 = 0
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
