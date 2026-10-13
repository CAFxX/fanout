package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int32; F1 uint8 }

func main() {
    var v1 float32 = -2.25
    v2 := []int{-9223372036854775808}
    v3 := []string{}
    var v4 float64 = (float64((v1 - 0.0)) * 0.0)
    var v5 string = "hello"
    v6 := GenS1{F0: 0, F1: 0}
    _ = len(v3)
    for fi7 := 0; fi7 < 11; fi7++ {
        if (("hello" != ("\n\t" + ("foo bar" + (v5 + (v5 + v5))))) && false) {
            v4 -= v4
            v4 += (((((3.14159 - float64(v1)) / v4) * v4) - (-2.25 / (0.0 / (v4 / v4)))) / 0.1)
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        switch ("foo bar" + v5) {
        case "foo bar":
            var v8 float32 = 1.5
            _ = v8
        case "":
            if len(v3) > 0 {
                v3[0] = v5
            }
        default:
            switch (fi7 / (((^(fi7))) | 1)) {
            case 397:
                if (false || (!true)) {
                    v4 += 3.14159
                    for fi9 := 0; fi9 < 17; fi9++ {
                        var v10 bool = (v4 < (-2.25 * (v4 - v4)))
                        _ = fi9
                        _ = v10
                    }
                }
            case -397:
                if len(v2) > 0 {
                    v2[0] = 32768
                }
            case 32767:
                switch (fi7 - (fi7 + (fi7 << ((2) & 63)))) {
                case 9223372036854775807:
                    var v11 int32 = -755
                    _ = v11
                }
            }
        }
        fmt.Println(v1, v1)
        _ = fi7
    }
    v4 += 0.0
    v5 = "hello"
    v12 := GenS2{F0: 42, F1: 255}
    fmt.Println(v4, v1)
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 uint16 = 65535
        _ = fi13
        _ = v14
    }
    fmt.Println(v1)
    for fi15 := 0; fi15 < 15; fi15++ {
        var v16 uint32 = 968
        v16 += 426
        _ = fi15
        _ = v16
    }
    var v17 uint16 = 604
    var v18 uint32 = 2147483648
    v17 = (v17 & v17)
    if len(v2) > 0 {
        v2[0] = 9223372036854775807
    }
    fmt.Println(v18, v5)
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
