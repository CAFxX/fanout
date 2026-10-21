package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int32 }

func main() {
    var v1 uint16 = 248
    var v2 int8 = -87
    v3 := []string{}
    var v4 int8 = v2
    var v5 uint32 = 42
    var v6 int16 = int16(v5)
    v2 -= v2
    v2 += (v2 & (v2 << ((2) & 63)))
    if (!true) {
        v3 = append(v3, "foo bar")
        for fi7 := 0; fi7 < 1; fi7++ {
            for fi8 := 0; fi8 < 8; fi8++ {
                v6 = v6
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi8
            }
            _ = fi7
        }
    }
    switch -905 {
    case -9223372036854775808:
        var v9 int8 = 1
        _ = v9
    case 378:
        switch "\n\t" {
        case "foo bar":
            v10 := GenS1{F0: (v5 - v5), F1: 2147483647, F2: 42}
            _ = v10
        case "hello":
            v11 := GenS1{F0: (v5 | 2147483648), F1: v5, F2: -438}
            _ = v11
        }
    default:
        var v12 bool = true
        _ = v12
    }
    var v13 int32 = 42
    var v14 float64 = ((3.14159 * 3.14159) / 1e10)
    if len(v3) > 0 {
        v3[0] = ("a" + ("hello" + ("a" + ("foo bar" + "foo bar"))))
    }
    v15 := GenS1{F0: (v5 - (v5 & v5)), F1: 2147483647, F2: v13}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
