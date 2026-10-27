package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
func genFn0() uint64 {
    var v5 float64 = 1.5
    if (128 >= 91) {
        v6 := GenS1{F0: 2, F1: -963}
        v5 += float64((-2.25 - (3.14159 + 0.0)))
        _ = v6
    }
    var v7 int32 = 2147483647
    v7 = (v7 | -2147483648)
    var v8 bool = (true && ((false || true) || false))
    _ = v5
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 string = ""
    var v2 uint32 = 1
    v3 := []string{v1, v1, "", v1}
    var v4 int = 32767
    v2 = (v2 >> ((1) & 63))
    v2 = 0
    v4 = (v4 % (((v4 % ((v4) | 1))) | 1))
    var v9 uint16 = 255
    if true {
        if false {
            var v10 int8 = -128
            _ = v10
        }
    }
    var v11 float64 = 0.1
    switch v1 {
    case "":
        if (v9 <= 32767) {
            switch (^(v4)) {
            case -2147483648:
                switch (v4 >> ((42) & 63)) {
                case 127:
                    v3 = append(v3, v1)
                case -996:
                    v4 -= v4
                default:
                    v12 := GenS1{F0: (v9 & (v9 * (v9 % ((v9) | 1)))), F1: v4}
                    _ = v12
                }
            case -9223372036854775808:
                var v13 int16 = 457
                _ = v13
            case -584:
                for fi14 := 0; fi14 < 12; fi14++ {
                    v1 = v1
                    _ = fi14
                }
            default:
                _ = len(v3)
            }
        } else {
            switch "foo bar" {
            case "hello":
                v2 = v2
            case "a":
                switch v1 {
                case "\n\t":
                    v2 = 2147483647
                }
            default:
                switch (-(v4)) {
                case 0:
                    v15 := GenS1{F0: 1, F1: v4}
                    _ = v15
                case 2147483647:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                default:
                    v2 = (v2 - v2)
                }
            }
        }
    }
    fmt.Println(v1, v1, v9)
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
