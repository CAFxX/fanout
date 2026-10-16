package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 uint8; F2 uint8 }

func main() {
    v1 := []int{128, 64, 709}
    var v2 uint32 = 933
    var v3 uint32 = v2
    v3 -= (v2 << ((18446744073709551615) & 63))
    var v4 uint = 2
    var v5 string = ""
    var v6 uint16 = 903
    v4 -= 1
    var v7 int16 = int16(v6)
    var v8 bool = (false || true)
    switch 1 {
    case -9223372036854775808:
        if (true || v8) {
            v9 := GenS1{F0: 596, F1: 128, F2: 42}
            for fi10 := 0; fi10 < 18; fi10++ {
                v7 = 0
                _ = fi10
            }
            var v11 int64 = -9223372036854775808
            _ = v11
            _ = v9
        }
    }
    var v12 int16 = -185
    for fi13 := 0; fi13 < 20; fi13++ {
        if (!v8) {
            var v14 int16 = -74
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v14
        } else {
            v15 := GenS1{F0: 790, F1: 813, F2: int32(v4)}
            _ = v15
        }
        v12 += (v12 >> ((v4) & 63))
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 string = v5
            _ = fi16
            _ = v17
        }
        _ = fi13
    }
    switch -277 {
    case 535:
        switch 127 {
        case 143:
            if len(v1) > 0 {
                v1[0] = -9223372036854775808
            }
        default:
            if v8 {
                fmt.Println(v3, v7)
                switch v5 {
                case "\n\t":
                    v12 = (v7 / ((1) | 1))
                case "a":
                    v6 = (v6 ^ v6)
                case "foo bar":
                    for fi18 := 0; fi18 < 1; fi18++ {
                        var v19 uint8 = 1
                        _ = fi18
                        _ = v19
                    }
                default:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                v20 := GenS1{F0: 1, F1: -1, F2: 2147483647}
                _ = v20
            }
        }
    }
    var v21 uint16 = v6
    switch 605 {
    case 127:
        switch "a" {
        case "":
            var v22 bool = v8
            _ = v22
        case "\n\t":
            if ((v7 ^ int16(v4)) >= (v12 >> (((v4 << ((v4) & 63))) & 63))) {
                for fi23 := 0; fi23 < 11; fi23++ {
                    v6 -= v6
                    v24 := GenS2{F0: (fi23 | fi23)}
                    fmt.Println(v2, v12, v21)
                    _ = fi23
                    _ = v24
                }
                var v25 int8 = -1
                _ = v25
            }
        }
    default:
        v6 = v6
    }
    _ = len(v1)
    v26 := GenS3{F0: 980, F1: 163, F2: 182}
    var v27 bool = v8
    var v28 bool = (1 >= -70)
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
