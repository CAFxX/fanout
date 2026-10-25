package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint64; F1 int16; F2 uint8 }
func genFn0(p0 uint16, p1 int8) int {
    v6 := GenS1{F0: 255, F1: 1, F2: 4294967295}
    for fi7 := 0; fi7 < 3; fi7++ {
        if true {
            var v8 int16 = 70
            _ = v8
        }
        switch "\n\t" {
        case "":
            v9 := []int{fi7, fi7, (fi7 << ((65535) & 63)), (^(fi7))}
            _ = v9
        case "hello":
            p1 = 127
        }
        _ = fi7
    }
    if true {
        p0 = 936
    } else {
        v10 := GenS2{F0: 2147483647, F1: -1}
        var v11 uint64 = 104
        _ = v10
        _ = v11
    }
    p1 += -1
    _ = p0
    _ = p1
    _ = v6
    return 4294967295
}

func main() {
    var v1 float32 = ((-2.25 - float32(1e10)) * (-2.25 + 0.1))
    var v2 float64 = -2.25
    var v3 bool = (1 != 1)
    v4 := []uint32{810, 2147483647}
    var v5 int = 42
    v2 = ((v2 - (float64(1e10) * (v2 * (v2 * 1e10)))) * v2)
    switch -842 {
    case 969:
        v2 = (v2 - (v2 * (1e10 + float64(v1))))
    }
    if ("" != "") {
        switch -209 {
        case -999:
            for fi12 := 0; fi12 < 11; fi12++ {
                var v13 int64 = -9223372036854775808
                switch "a" {
                case "foo bar":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                _ = fi12
                _ = v13
            }
        default:
            var v14 int64 = 0
            _ = v14
        }
        if len(v4) > 0 {
            v4[0] = 739
        }
        if v3 {
            fmt.Println(v3, v5)
        } else {
            for fi15 := 0; fi15 < 13; fi15++ {
                for fi16 := 0; fi16 < 3; fi16++ {
                    var v17 int16 = 32767
                    v18 := GenS2{F0: -2147483648, F1: v17}
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                v2 = 0.1
                _ = fi15
            }
        }
    }
    var v19 bool = (0 <= 1000)
    var v20 uint32 = 2147483647
    v21 := GenS3{F0: 18446744073709551615, F1: 1, F2: 128}
    _ = len(v4)
    for fi22 := 0; fi22 < 20; fi22++ {
        switch (v5 & (^(fi22))) {
        case -9223372036854775808:
            for fi23 := 0; fi23 < 1; fi23++ {
                var v24 uint64 = 2
                v1 = v1
                _ = fi23
                _ = v24
            }
        default:
            var v25 uint16 = 600
            _ = v25
        }
        v26 := GenS2{F0: 0, F1: -32768}
        _ = fi22
        _ = v26
    }
    v2 = (3.14159 / 1.5)
    for fi27 := 0; fi27 < 10; fi27++ {
        var v28 bool = v3
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
