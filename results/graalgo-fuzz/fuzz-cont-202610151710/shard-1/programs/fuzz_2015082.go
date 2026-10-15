package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int64; F1 int64; F2 uint64 }
type GenS3 struct { F0 uint8; F1 int8; F2 uint16 }

func main() {
    var v1 uint32 = 65535
    v2 := make(map[string]uint32)
    v3 := make(map[int]string)
    v4 := make(map[uint32]uint32)
    switch (("hello" + ("\n\t" + "foo bar")) + "hello") {
    case "hello":
        fmt.Println(v1)
    case "":
        switch -893 {
        case 0:
            v1 -= v1
        case 533:
            var v5 int32 = 1
            _ = v5
        case 4294967295:
            if (1 != 71) {
                v6 := GenS1{F0: -2147483648, F1: 255}
                _ = v6
            }
        default:
            fmt.Println(v1)
        }
    }
    fmt.Println(v1)
    var v7 uint16 = 896
    switch -219 {
    case -9223372036854775808:
        switch -2147483648 {
        case 9223372036854775807:
            v1 -= (v1 % ((v1) | 1))
        }
    case 0:
        for fi8 := 0; fi8 < 3; fi8++ {
            for fi9 := 0; fi9 < 8; fi9++ {
                delete(v3, fi9)
                var v10 float32 = ((1e10 + ((3.14159 - 1.5) / 1.0)) / 0.1)
                var v11 string = ("" + "")
                v12 := GenS1{F0: 98, F1: v7}
                if false {
                    v1 = v1
                } else {
                    v13 := []int{32767, 0}
                    _ = v13
                }
                v14 := GenS2{F0: 42, F1: 1, F2: 4294967295}
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
                _ = v14
            }
            v15 := GenS1{F0: 546, F1: (v7 ^ v7)}
            _ = fi8
            _ = v15
        }
    }
    if true {
        v16 := GenS1{F0: 1, F1: (v7 * uint16(v1))}
        _ = v16
    } else {
        v17 := GenS2{F0: -764, F1: 9223372036854775807, F2: 1}
        for fi18 := 0; fi18 < 10; fi18++ {
            v19 := GenS3{F0: 217, F1: 73, F2: v7}
            _ = fi18
            _ = v19
        }
        _ = v17
    }
    var v20 int64 = -808
    var v21 uint8 = 2
    for fi22 := 0; fi22 < 5; fi22++ {
        v21 = (v21 / (((v21 >> ((1) & 63))) | 1))
        for fi23 := 0; fi23 < 12; fi23++ {
            var v24 int8 = 124
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    v1 += (^(v1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v7
}
