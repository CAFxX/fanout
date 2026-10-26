package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 int16 }
type GenS2 struct { F0 uint; F1 int64; F2 uint }
type GenS3 struct { F0 int; F1 uint8; F2 int32 }

func main() {
    v1 := []uint32{}
    var v2 int64 = 1
    var v3 float32 = -2.25
    var v4 int8 = int8(v2)
    var v5 int64 = -9223372036854775808
    switch "" {
    case "foo bar":
        if len(v1) > 0 {
            v1[0] = 219
        }
    case "hello":
        v5 -= (v2 * (v5 & v5))
    case "":
        _ = len(v1)
    }
    v3 -= (v3 * (v3 + v3))
    if (1e10 >= float64(3.14159)) {
        var v6 uint = 4294967295
        _ = v6
    }
    v5 = 9223372036854775807
    v7 := GenS1{F0: -471, F1: 127, F2: -32768}
    var v8 int16 = 0
    v9 := GenS1{F0: 1, F1: -9223372036854775808, F2: (v8 + (v8 - (v8 >> ((2) & 63))))}
    var v10 int16 = int16(v5)
    if (32767 < 0) {
        v8 = int16(v4)
    } else {
        var v11 uint16 = 42
        _ = v11
    }
    v5 -= (v2 % ((v5) | 1))
    switch ((("" + "\n\t") + "\n\t") + ((((((("foo bar" + ("\n\t" + ("\n\t" + ("a" + (("" + ("hello" + "\n\t")) + ""))))) + "foo bar") + "a") + "\n\t") + ("foo bar" + "")) + "\n\t") + "foo bar")) {
    case "foo bar":
        var v12 bool = false
        _ = v12
    case "\n\t":
        for fi13 := 0; fi13 < 18; fi13++ {
            var v14 bool = false
            if v14 {
                v14 = (false && (((v3 + v3) + v3) <= v3))
                v15 := GenS2{F0: 723, F1: v2, F2: 897}
                _ = v15
            }
            for fi16 := 0; fi16 < 2; fi16++ {
                if (!(18446744073709551615 < 4294967295)) {
                    var v17 int32 = 1
                    if true {
                        switch "" {
                        case "foo bar":
                            v3 += 3.14159
                        case "\n\t":
                            v5 = v2
                        default:
                            var v18 bool = v14
                            _ = v18
                        }
                    }
                    _ = v17
                } else {
                    var v19 float32 = float32(-2.25)
                    _ = v19
                }
                var v20 uint16 = 42
                switch fi13 {
                case 126:
                    v3 -= (v3 / v3)
                case 32768:
                    fmt.Println(v20)
                default:
                    v5 -= (v2 | v2)
                }
                _ = fi16
                _ = v20
            }
            _ = fi13
            _ = v14
        }
    case "a":
        switch "hello" {
        case "a":
            v21 := GenS3{F0: 32768, F1: 1, F2: 72}
            _ = v21
        case "hello":
            var v22 string = "\n\t"
            _ = v22
        }
    }
    fmt.Println(v8)
    for fi23 := 0; fi23 < 14; fi23++ {
        var v24 float32 = 3.14159
        _ = fi23
        _ = v24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
