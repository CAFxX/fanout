package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint64 }
type GenS2 struct { F0 uint; F1 int }

func main() {
    var v1 uint64 = 450
    v2 := []int{}
    var v3 int64 = 0
    var v4 uint32 = 2147483647
    v3 = (v3 >> ((18446744073709551615) & 63))
    if true {
        fmt.Println(v4)
        v5 := GenS1{F0: 42, F1: true, F2: (v1 - 252)}
        _ = v5
    }
    v3 = v3
    v1 = (v1 | v1)
    fmt.Println(v4, v4, v4)
    v1 -= v1
    if false {
        var v6 int8 = -128
        var v7 string = "\n\t"
        _ = v6
        _ = v7
    }
    v1 = v1
    v2 = append(v2, 4294967295)
    v3 += v3
    for fi8 := 0; fi8 < 10; fi8++ {
        if len(v2) > 0 {
            v2[0] = (fi8 >> ((613) & 63))
        }
        for fi9 := 0; fi9 < 12; fi9++ {
            v3 -= (v3 * v3)
            switch "foo bar" {
            case "foo bar":
                var v10 uint64 = v1
                _ = v10
            case "":
                switch "foo bar" {
                case "":
                    var v11 int16 = -1
                    var v12 int16 = v11
                    _ = v11
                    _ = v12
                case "foo bar":
                    switch 32767 {
                    case 1:
                        fmt.Println(v1, v1)
                    case 298:
                        if (v4 >= (v4 << ((255) & 63))) {
                            v1 = 0
                            v1 -= v1
                        }
                    }
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            default:
                if (!false) {
                    v13 := GenS1{F0: (v3 >> ((1) & 63)), F1: true, F2: (v1 << ((312) & 63))}
                    v14 := GenS2{F0: 65535, F1: 32767}
                    _ = v13
                    _ = v14
                }
            }
            _ = fi9
        }
        v15 := GenS2{F0: 4294967295, F1: int(v1)}
        _ = fi8
        _ = v15
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
