package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 bool; F1 uint32; F2 uint64 }
type GenS3 struct { F0 uint8; F1 int; F2 uint }

func main() {
    var v1 int32 = -866
    v2 := make(map[int]bool)
    var v3 uint8 = 244
    v4 := make(map[string]string)
    switch "foo bar" {
    case "hello":
        var v5 int64 = 42
        _ = v5
    case "":
        fmt.Println(v3, v1)
    case "\n\t":
        var v6 uint64 = 87
        _ = v6
    }
    v7 := GenS1{F0: 42, F1: (v3 & v3)}
    v3 = v3
    v2[-9223372036854775808] = (!((v3 & (v3 + (v3 - (^(v3))))) >= 255))
    var v8 int8 = 1
    if (0 != 877) {
        v9 := GenS1{F0: 2, F1: v3}
        var v10 uint8 = 127
        fmt.Println(v8)
        _ = v10
        _ = v9
    } else {
        var v11 uint32 = 0
        var v12 float32 = 1.5
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 19; fi13++ {
        if (!true) {
            if true {
                var v14 uint = 4294967295
                _ = v14
            } else {
                v15 := GenS2{F0: ((1 < 65535) || (45 != v8)), F1: 0, F2: 18446744073709551615}
                for fi16 := 0; fi16 < 9; fi16++ {
                    switch "a" {
                    case "\n\t":
                        var v17 int32 = (v1 & 0)
                        var v18 int32 = 941
                        switch "a" {
                        case "hello":
                            if ((("foo bar" + "a") + "hello") <= "foo bar") {
                                v19 := []int{(fi13 / ((fi13) | 1))}
                                fmt.Println(fi13, v3, fi16)
                                fmt.Println(fi13, v18, v18)
                                _ = v19
                            }
                        default:
                            var v20 uint = 965
                            _ = v20
                        }
                        _ = v17
                        _ = v18
                    case "a":
                        v21 := GenS3{F0: v3, F1: (fi13 >> ((60) & 63)), F2: 253}
                        _ = v21
                    case "":
                        v1 -= -2147483648
                    }
                    _ = fi16
                }
                _ = v15
            }
            var v22 uint8 = (^(v3))
            fmt.Println(v22)
            _ = v22
        }
        _ = fi13
    }
    v8 = (v8 % (((v8 / ((v8) | 1))) | 1))
    fmt.Println(v8)
    v3 += v3
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
