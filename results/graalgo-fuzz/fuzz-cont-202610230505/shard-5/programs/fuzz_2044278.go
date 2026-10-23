package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 int8; F2 int }

func main() {
    var v1 string = "\n\t"
    v2 := []uint32{227, 42}
    var v3 int = 527
    var v4 uint = 65535
    v3 -= (v3 << ((v4) & 63))
    v3 = v3
    if ((42 > 42) && (1e10 != float64(((float32(0.0) * -2.25) / (float32(1e10) - 1.5))))) {
        if (804 == 42) {
            var v5 string = "foo bar"
            v1 = v5
            var v6 int = 2147483647
            _ = v5
            _ = v6
        } else {
            for fi7 := 0; fi7 < 16; fi7++ {
                fmt.Println(v4, v3, fi7)
                switch -9223372036854775808 {
                case -9223372036854775808:
                    if false {
                        fmt.Println(v3, v3)
                        v3 -= (v3 + fi7)
                    }
                case 606:
                    v8 := GenS1{F0: 777}
                    _ = v8
                }
                _ = fi7
            }
            v1 = ""
            var v9 uint16 = 0
            _ = v9
        }
    }
    v3 -= (v3 + v3)
    switch 0 {
    case 667:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        var v10 string = v1
        _ = v10
    }
    fmt.Println(v4, v3)
    var v11 int8 = -128
    if (false || (0 > (v11 * (v11 / (((v11 / ((v11) | 1))) | 1))))) {
        v4 += 178
        v2 = append(v2, 2147483647)
    }
    v12 := GenS2{F0: v4, F1: 1, F2: (v3 >> (((^(v4))) & 63))}
    fmt.Println(v3, v11)
    var v13 int16 = 32767
    var v14 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
