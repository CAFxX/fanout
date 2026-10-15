package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 int16; F2 uint }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 int; F1 uint }

func main() {
    var v1 int = 954
    var v2 int32 = 42
    v3 := make(map[uint32]uint32)
    var v4 int32 = (v2 / ((42) | 1))
    v2 -= (v2 - v2)
    var v5 uint16 = 2
    var v6 int16 = 42
    var v7 int8 = -128
    if (!(!false)) {
        fmt.Println(v6, v5, v6)
        v8 := GenS1{F0: 18446744073709551615}
        _ = v8
    } else {
        var v9 uint64 = 99
        v6 = -1
        if false {
            if false {
                switch "\n\t" {
                case "foo bar":
                    for fi10 := 0; fi10 < 20; fi10++ {
                        if true {
                            fmt.Println(fi10, v9, v4)
                            for fi11 := 0; fi11 < 12; fi11++ {
                                v12 := GenS2{F0: 130, F1: 1000, F2: 147}
                                var v13 uint16 = (v5 ^ v5)
                                v14 := GenS3{F0: "", F1: (v4 / ((v2) | 1))}
                                _ = fi11
                                _ = v12
                                _ = v13
                                _ = v14
                            }
                        } else {
                            _ = len(v3)
                            var v15 uint8 = 0
                            _ = v15
                        }
                        _ = fi10
                    }
                case "":
                    var v16 uint64 = v9
                    _ = v16
                case "hello":
                    var v17 uint64 = (v9 * 919)
                    _ = v17
                }
            } else {
                v7 = v7
                v6 += v6
            }
        }
        _ = v9
    }
    v18 := []int{}
    v2 -= (-(v2))
    v7 = v7
    v19 := GenS4{F0: 32767, F1: 4294967295}
    v6 += (v6 | 32767)
    switch v1 {
    case 9223372036854775807:
        v20 := GenS3{F0: "", F1: (-(v4))}
        _ = v20
    case 2147483647:
        v21 := v3[5]
        _ = v21
    case -9223372036854775808:
        v22 := GenS1{F0: 42}
        _ = v22
    }
    if ((3.14159 == 1.5) && ((0 != 0) || false)) {
        var v23 int16 = v6
        var v24 uint8 = 1
        _ = v23
        _ = v24
    } else {
        fmt.Println(v4)
    }
    var v25 float32 = 3.14159
    if true {
        v7 -= -1
    }
    v25 += v25
    v1 += (v1 ^ -1)
    for fi26 := 0; fi26 < 8; fi26++ {
        fmt.Println(v4)
        _ = fi26
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
