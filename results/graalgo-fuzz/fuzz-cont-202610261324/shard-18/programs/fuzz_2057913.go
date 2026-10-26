package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint64 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint64; F1 uint8; F2 uint }
type GenS4 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]bool)
    var v3 uint = 4294967295
    var v4 uint8 = 214
    fmt.Println(v4, v4)
    v4 = (v4 << ((18446744073709551615) & 63))
    for fi5 := 0; fi5 < 2; fi5++ {
        for fi6 := 0; fi6 < 16; fi6++ {
            fmt.Println(fi6)
            _ = fi6
        }
        _ = fi5
    }
    v4 = (v4 ^ v4)
    if true {
        for fi7 := 0; fi7 < 3; fi7++ {
            for fi8 := 0; fi8 < 8; fi8++ {
                v9 := GenS1{F0: v3, F1: fi8, F2: 42}
                v4 = 128
                switch "foo bar" {
                case "hello":
                    v3 = 0
                case "":
                    v10 := GenS1{F0: v3, F1: fi8, F2: 18446744073709551615}
                    _ = v10
                default:
                    v3 -= (^(v3))
                    v4 = v4
                }
                var v11 bool = false
                _ = fi8
                _ = v11
                _ = v9
            }
            _ = fi7
        }
        fmt.Println(v4)
    }
    if false {
        v4 += v4
    }
    v12 := GenS2{F0: 236, F1: 116}
    switch 1 {
    case 180:
        v3 = (v3 / (((v3 ^ v3)) | 1))
    case 2147483647:
        v13 := GenS2{F0: 0, F1: 328}
        _ = v13
    }
    v14 := GenS3{F0: 0, F1: (v4 >> (((v3 + (v3 >> ((65535) & 63)))) & 63)), F2: (v3 / ((v3) | 1))}
    if (true && true) {
        v15 := []int{}
        _ = v15
    } else {
        if true {
            var v16 int64 = -729
            var v17 int = -653
            v18 := GenS2{F0: 42, F1: uint64(v3)}
            _ = v16
            _ = v17
            _ = v18
        } else {
            if false {
                switch "\n\t" {
                case "hello":
                    v19 := []int{-2147483648, 0, 42, 1}
                    _ = v19
                }
            } else {
                if (-210 <= -1) {
                    for fi20 := 0; fi20 < 1; fi20++ {
                        if false {
                            v21 := []int{-331, (fi20 & fi20), (fi20 << (((^(v3))) & 63)), (fi20 % ((fi20) | 1))}
                            _ = v21
                        } else {
                            v22 := []int{fi20, fi20, 9223372036854775807, fi20}
                            _ = v22
                        }
                        var v23 int16 = 1
                        for fi24 := 0; fi24 < 12; fi24++ {
                            var v25 bool = ((1.5 * 1.5) > 0.1)
                            _ = fi24
                            _ = v25
                        }
                        _ = fi20
                        _ = v23
                    }
                    var v26 uint = v3
                    v27 := []int{9223372036854775807, -152, 0, 4294967295}
                    _ = v26
                    _ = v27
                }
                v28 := GenS4{F0: -9223372036854775808}
                _ = v28
            }
            var v29 uint32 = 2
            v4 = (v4 & 128)
            _ = v29
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
