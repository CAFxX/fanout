package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int64 }
type GenS2 struct { F0 uint16; F1 uint16 }

func main() {
    var v1 float64 = 3.14159
    var v2 int = 9223372036854775807
    var v3 float64 = (((v1 - v1) / ((v1 * v1) / 1.0)) * ((3.14159 * v1) + float64(-2.25)))
    var v4 uint64 = 1
    var v5 int = -2147483648
    v6 := []int{(v2 << ((615) & 63)), v5}
    if true {
        switch "a" {
        case "\n\t":
            var v7 bool = (32767 == 1000)
            _ = v7
        case "a":
            if (!(-2147483648 != 2147483647)) {
                if (((299 >= 909) && false) || (989 != 65535)) {
                    for fi8 := 0; fi8 < 9; fi8++ {
                        v6 = append(v6, fi8)
                        _ = fi8
                    }
                    var v9 int64 = 42
                    for fi10 := 0; fi10 < 1; fi10++ {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        _ = fi10
                    }
                    _ = v9
                } else {
                    v11 := GenS1{F0: -1, F1: 127, F2: 991}
                    var v12 string = "hello"
                    _ = v11
                    _ = v12
                }
                v6 = append(v6, v2)
            }
        case "foo bar":
            v13 := GenS1{F0: 42, F1: 0, F2: -777}
            _ = v13
        default:
            var v14 uint = 467
            _ = v14
        }
        var v15 string = ("foo bar" + "hello")
        _ = v15
    } else {
        for fi16 := 0; fi16 < 14; fi16++ {
            v3 -= v3
            if (!true) {
                v3 = -2.25
                var v17 int32 = 42
                _ = len(v6)
                _ = v17
            } else {
                fmt.Println(v3, fi16)
            }
            _ = fi16
        }
        if ((!false) || false) {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
        v18 := GenS2{F0: 32767, F1: 2}
        _ = v18
    }
    _ = len(v6)
    if true {
        v1 = v1
    } else {
        _ = len(v6)
    }
    switch 1 {
    case 0:
        v19 := GenS1{F0: 0, F1: 1, F2: 1}
        _ = v19
    case 32767:
        fmt.Println(v5)
    case -2147483648:
        fmt.Println(v1)
    default:
        _ = len(v6)
    }
    v5 += v5
    v4 += (v4 << ((900) & 63))
    var v20 bool = false
    v20 = (657 != 32767)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
