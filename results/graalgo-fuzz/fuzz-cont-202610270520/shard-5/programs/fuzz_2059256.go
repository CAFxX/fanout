package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint64 }
type GenS2 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 string = "foo bar"
    var v2 int16 = -420
    var v3 uint = 1
    var v4 int32 = 42
    v4 -= 2147483647
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 uint = (v3 / (((v3 - 109)) | 1))
        for fi7 := 0; fi7 < 17; fi7++ {
            v8 := GenS1{F0: v4, F1: (fi5 >> (((v3 >> ((18446744073709551615) & 63))) & 63)), F2: 0}
            v2 += int16(fi5)
            v2 += 1
            _ = fi7
            _ = v8
        }
        fmt.Println(v6, v6)
        _ = fi5
        _ = v6
    }
    var v9 int16 = 1
    switch v1 {
    case "\n\t":
        fmt.Println(v3, v1)
    default:
        var v10 int16 = 1000
        _ = v10
    }
    var v11 int32 = v4
    v12 := []int{36, 128, 279}
    v13 := GenS2{F0: (v2 | (v9 - (v9 | v9))), F1: 2, F2: 1}
    if (65535 < 0) {
        v14 := GenS2{F0: (v2 & 42), F1: 2, F2: 255}
        _ = v14
    } else {
        v15 := GenS2{F0: v9, F1: 159, F2: 255}
        _ = v15
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        v12 = append(v12, fi16)
        for fi17 := 0; fi17 < 3; fi17++ {
            if len(v12) > 0 {
                v12[0] = fi16
            }
            v11 += (v11 + v4)
            _ = fi17
        }
        v18 := GenS3{F0: -1}
        _ = fi16
        _ = v18
    }
    switch int(v3) {
    case 127:
        v4 = (v11 + v4)
    }
    v2 += (v2 + (v9 ^ (v9 / ((960) | 1))))
    if len(v12) > 0 {
        v12[0] = -9223372036854775808
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
