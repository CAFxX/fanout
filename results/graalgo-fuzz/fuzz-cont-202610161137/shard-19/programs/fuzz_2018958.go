package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 int32; F1 int; F2 uint16 }
type GenS3 struct { F0 int; F1 uint8 }

func main() {
    var v1 uint16 = 2
    var v2 float32 = ((-2.25 - 3.14159) * (3.14159 - 0.0))
    v3 := []string{"a", "hello", (("foo bar" + "") + "a"), "foo bar"}
    var v4 uint64 = 1
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    var v5 bool = (false && ((!true) || (79 < 58)))
    v6 := GenS1{F0: 91, F1: 18446744073709551615}
    switch "hello" {
    case "foo bar":
        if true {
            v2 = (v2 / v2)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    }
    if true {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v7 int8 = -128
        _ = v7
    } else {
        fmt.Println(v1, v1, v5)
        v8 := GenS1{F0: 100, F1: 42}
        v3 = append(v3, "hello")
        _ = v8
    }
    switch 127 {
    case -601:
        var v9 int = 127
        _ = v9
    default:
        switch -9223372036854775808 {
        case 741:
            var v10 uint16 = v1
            _ = v10
        default:
            for fi11 := 0; fi11 < 9; fi11++ {
                if (((v4 / (((^(v4))) | 1)) < (v4 >> ((42) & 63))) || ((2147483647 <= -2147483648) || true)) {
                    var v12 bool = (!(-55 <= 44))
                    v13 := GenS2{F0: int32(v1), F1: 4294967295, F2: (v1 / ((1) | 1))}
                    v4 = (v4 ^ v4)
                    _ = v12
                    _ = v13
                } else {
                    v1 = (v1 / (((v1 + v1)) | 1))
                    var v14 uint32 = 42
                    fmt.Println(v5)
                    _ = v14
                }
                _ = fi11
            }
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 -= ((float32(0.1) * v2) - ((v2 * v2) + (v2 - ((v2 - (v2 - (float32((0.0 / (float64((v2 * v2)) * 1e10))) + (float32(-2.25) - v2)))) * ((0.0 / (v2 - v2)) * v2)))))
    v5 = true
    fmt.Println(v2, v1, v2)
    for fi15 := 0; fi15 < 2; fi15++ {
        for fi16 := 0; fi16 < 20; fi16++ {
            v4 = (v4 % ((v4) | 1))
            switch "a" {
            case "\n\t":
                v4 -= v4
            case "hello":
                v17 := GenS2{F0: -635, F1: (^(fi15)), F2: (v1 >> ((42) & 63))}
                _ = v17
            }
            _ = fi16
        }
        if (!((v4 << ((2) & 63)) < (v4 * (^(v4))))) {
            v18 := GenS1{F0: -8, F1: 959}
            _ = v18
        }
        _ = fi15
    }
    fmt.Println(v2)
    v19 := GenS3{F0: 457, F1: 128}
    var v20 int64 = 672
    var v21 int16 = 42
    var v22 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
