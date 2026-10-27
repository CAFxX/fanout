package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint64; F1 uint; F2 uint }
type GenS3 struct { F0 uint32; F1 uint }

func main() {
    v1 := []uint32{2147483648, 671}
    var v2 float32 = 1.5
    var v3 bool = false
    v4 := make(map[uint32]bool)
    switch "" {
    case "hello":
        v3 = (v3 && (!((!v3) && (2147483647 != 150))))
    }
    if (2147483647 != -2147483648) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v5 := GenS1{F0: 18446744073709551615, F1: 0}
    for fi6 := 0; fi6 < 13; fi6++ {
        v7 := v4[2147483648]
        switch "" {
        case "":
            if len(v1) > 0 {
                v1[0] = 0
            }
        case "\n\t":
            for fi8 := 0; fi8 < 16; fi8++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v9 int64 = 49
                _ = fi8
                _ = v9
            }
        default:
            var v10 int64 = -1
            _ = v10
        }
        v3 = v3
        _ = fi6
        _ = v7
    }
    v11 := GenS2{F0: 1, F1: 18446744073709551615, F2: 86}
    switch "hello" {
    case "\n\t":
        v2 = 3.14159
    }
    fmt.Println(v2)
    if false {
        v12 := GenS2{F0: 259, F1: 773, F2: 65535}
        var v13 uint32 = 0
        v3 = true
        _ = v12
        _ = v13
    }
    v4[67] = false
    fmt.Println(v2, v2)
    for fi14 := 0; fi14 < 2; fi14++ {
        v15 := GenS1{F0: 42, F1: -2147483648}
        v16 := GenS3{F0: 4294967295, F1: 490}
        var v17 uint = 2
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    v1 = append(v1, 513)
    delete(v4, 1)
    var v18 uint64 = 681
    if (-1 <= 0) {
        var v19 int64 = -840
        _ = v19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
