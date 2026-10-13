package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }

func main() {
    var v1 int16 = 0
    var v2 int = -9223372036854775808
    v3 := []int{(v2 / (((v2 << ((255) & 63))) | 1)), 42}
    var v4 uint8 = 2
    v5 := []uint32{2147483647, 193, 134}
    v6 := GenS1{F0: 65535, F1: 156}
    switch ("a" + "\n\t") {
    case "\n\t":
        fmt.Println(v4)
    default:
        if true {
            if ((v1 >= (v1 * (^(v1)))) || (false && (!(42 == 138)))) {
                var v7 uint32 = 481
                _ = v7
            } else {
                fmt.Println(v1, v1, v4)
            }
            v4 += (^(v4))
            var v8 uint8 = v4
            _ = v8
        }
    }
    var v9 bool = (false || true)
    if false {
        v2 -= (v2 % ((0) | 1))
    }
    switch "" {
    case "\n\t":
        for fi10 := 0; fi10 < 19; fi10++ {
            if (-2 > -128) {
                if (0 <= 1) {
                    v1 = 32767
                    if (v1 <= v1) {
                        if v9 {
                            if len(v3) > 0 {
                                v3[0] = v2
                            }
                            v11 := GenS1{F0: 18446744073709551615, F1: (v4 * (v4 >> ((255) & 63)))}
                            v1 += 1
                            _ = v11
                        }
                        _ = len(v3)
                    } else {
                        fmt.Println(v4, v1, v4)
                    }
                    var v12 string = ""
                    _ = v12
                }
            } else {
                v4 = v4
                v4 -= (v4 & (v4 >> ((0) & 63)))
            }
            _ = fi10
        }
    case "foo bar":
        v5 = append(v5, 4294967295)
    case "hello":
        var v13 int8 = -76
        _ = v13
    }
    _ = len(v3)
    v1 = -519
    if true {
        v2 -= v2
    } else {
        var v14 int8 = 0
        fmt.Println(v4, v14)
        v3 = append(v3, v2)
        _ = v14
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        var v16 int32 = -232
        _ = fi15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
