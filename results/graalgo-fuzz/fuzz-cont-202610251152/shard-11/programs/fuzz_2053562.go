package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint16; F1 uint16 }

func main() {
    v1 := make(map[int]string)
    v2 := []string{""}
    v3 := []uint32{767, 2147483648}
    v4 := []int{-904, -45, 0}
    if (!true) {
        v5 := v1[-2147483648]
        switch 32768 {
        case 9223372036854775807:
            fmt.Println(v5)
        }
        v5 = v5
        _ = v5
    }
    v6 := GenS1{F0: 42, F1: 18446744073709551615, F2: 245}
    var v7 uint = 2
    v7 -= v7
    v7 -= (v7 | (v7 >> (((v7 / (((v7 & v7)) | 1))) & 63)))
    v8 := GenS1{F0: 42, F1: 4294967295, F2: 128}
    switch "a" {
    case "\n\t":
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
    }
    var v9 int64 = 481
    v9 += -861
    if true {
        var v10 bool = (((42 > 0) || false) || ((float32(1e10) * 1.5) != 3.14159))
        v7 -= v7
        _ = v10
    } else {
        for fi11 := 0; fi11 < 14; fi11++ {
            v7 += 927
            if true {
                var v12 float32 = 0.1
                v12 += 0.1
                _ = v12
            }
            _ = fi11
        }
        var v13 uint16 = 2
        var v14 uint64 = 4294967295
        _ = v13
        _ = v14
    }
    delete(v1, -9223372036854775808)
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 int8 = 100
        _ = fi15
        _ = v16
    }
    var v17 uint32 = 2
    if true {
        v18 := v1[0]
        v19 := GenS1{F0: -1, F1: 18446744073709551615, F2: 127}
        v20 := GenS2{F0: 65535, F1: uint16(v7)}
        _ = v18
        _ = v19
        _ = v20
    }
    v21 := GenS1{F0: 872, F1: 42, F2: 255}
    var v22 bool = (!true)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
