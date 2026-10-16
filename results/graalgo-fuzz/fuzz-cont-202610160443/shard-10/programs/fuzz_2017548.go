package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := []string{"foo bar", "", "foo bar"}
    var v2 int16 = -189
    var v3 int16 = (v2 & v2)
    var v4 uint32 = 140
    var v5 uint64 = 0
    v1 = append(v1, "hello")
    if true {
        v4 = (v4 ^ (v4 | v4))
    } else {
        v6 := GenS1{F0: (v4 / ((v4) | 1)), F1: 127, F2: 255}
        _ = v6
    }
    var v7 int = 128
    var v8 float32 = ((-2.25 + 0.0) * ((((1e10 / 1.5) + (3.14159 * -2.25)) / ((3.14159 / 1e10) + 3.14159)) * 0.1))
    v9 := GenS1{F0: v4, F1: -1, F2: 32767}
    v10 := GenS1{F0: 253, F1: 68, F2: 2}
    var v11 float32 = v8
    switch v7 {
    case 9223372036854775807:
        v5 -= v5
    case 128:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case 695:
        switch "foo bar" {
        case "hello":
            v5 = (^(v5))
        case "":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "foo bar":
            switch v7 {
            case 42:
                var v12 int8 = -1
                _ = v12
            }
        default:
            if true {
                fmt.Println(v2)
                v13 := GenS2{F0: 21}
                _ = v13
            } else {
                v11 += (0.0 / ((3.14159 - v11) - 3.14159))
                for fi14 := 0; fi14 < 6; fi14++ {
                    v1 = append(v1, ("\n\t" + ""))
                    v2 = 32767
                    _ = fi14
                }
                for fi15 := 0; fi15 < 7; fi15++ {
                    v11 -= (v8 / (v8 - 0.1))
                    _ = fi15
                }
            }
        }
    default:
        fmt.Println(v8)
    }
    var v16 int32 = 932
    fmt.Println(v8, v5, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
