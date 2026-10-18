package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int32; F1 bool }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]bool)
    var v3 float32 = ((0.1 - 0.0) + -2.25)
    var v4 uint = 986
    var v5 uint16 = 42
    var v6 uint64 = 42
    fmt.Println(v6, v5)
    if (!(!(1000 > 32767))) {
        v7 := []int{-273, 760}
        _ = v7
    } else {
        v2[1] = true
        for fi8 := 0; fi8 < 8; fi8++ {
            if false {
                v4 -= (v4 ^ 4294967295)
                var v9 float64 = 0.1
                switch fi8 {
                case 2147483647:
                    v10 := GenS1{F0: (v6 % ((v6) | 1)), F1: -1}
                    _ = v10
                case -2147483648:
                    v4 -= 65535
                default:
                    switch fi8 {
                    case -2147483648:
                        _ = len(v1)
                    case -531:
                        fmt.Println(fi8, v5, fi8)
                    case -966:
                        v11 := []int{(-(fi8)), fi8}
                        _ = v11
                    default:
                        if (((!(!(0 < 18446744073709551615))) && (42 < 166)) || (((v6 | v6) >= uint64(v4)) && false)) {
                            v5 -= 2
                        }
                    }
                }
                _ = v9
            }
            fmt.Println(fi8)
            _ = fi8
        }
        v12 := GenS2{F0: v4, F1: 42}
        _ = v12
    }
    v13 := []int{461, 9223372036854775807, -1}
    fmt.Println(v3, v5, v6)
    v14 := GenS2{F0: v4, F1: 0}
    v4 = (v4 ^ v4)
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        var v16 int = (fi15 >> ((833) & 63))
        v17 := GenS3{F0: 233, F1: (true && (!true))}
        v18 := v1[v16]
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    var v19 uint32 = 0
    v5 = (v5 % (((v5 + (v5 / (((v5 % ((1) | 1))) | 1)))) | 1))
    switch ("hello" + "a") {
    case "foo bar":
        var v20 uint = (v4 ^ 167)
        _ = v20
    case "hello":
        _ = len(v13)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
