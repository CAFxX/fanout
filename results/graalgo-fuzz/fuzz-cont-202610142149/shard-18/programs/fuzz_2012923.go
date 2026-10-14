package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint8 }
type GenS2 struct { F0 int64; F1 int16 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]bool)
    v3 := []int{523, -785, 661}
    v4 := []uint32{1, 2147483647, 2147483648}
    var v5 int8 = 1
    switch "foo bar" {
    case "":
        var v6 uint64 = 1
        _ = v6
    default:
        v5 = (v5 / ((127) | 1))
    }
    var v7 int16 = 42
    v7 = v7
    v5 += 42
    for fi8 := 0; fi8 < 8; fi8++ {
        fmt.Println(v7, v7)
        if true {
            v9 := GenS1{F0: ("\n\t" + ("" + "\n\t")), F1: -1, F2: 190}
            _ = v9
        } else {
            v7 += 1
            for fi10 := 0; fi10 < 13; fi10++ {
                var v11 string = (("hello" + ("foo bar" + "\n\t")) + ("foo bar" + "hello"))
                v7 -= (v7 ^ (v7 + (v7 % ((v7) | 1))))
                fmt.Println(v5, v7)
                _ = fi10
                _ = v11
            }
        }
        for fi12 := 0; fi12 < 6; fi12++ {
            if len(v4) > 0 {
                v4[0] = 0
            }
            if true {
                fmt.Println(fi12, v5, fi8)
                fmt.Println(v7, fi8, fi8)
            } else {
                v5 -= (v5 >> ((2) & 63))
                if (-2.25 == (1.5 / (1e10 + -2.25))) {
                    var v13 bool = true
                    _ = v13
                }
            }
            _ = len(v3)
            _ = fi12
        }
        _ = fi8
    }
    var v14 float32 = ((0.1 / 1e10) / (3.14159 * ((float32(1.5) + 0.1) / (1e10 / (3.14159 + (3.14159 + 0.1))))))
    switch -8 {
    case 4294967295:
        v15 := GenS1{F0: "", F1: 0, F2: 255}
        _ = v15
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        switch (fi16 - -9223372036854775808) {
        case 128:
            v17 := GenS2{F0: 9223372036854775807, F1: v7}
            _ = v17
        default:
            v14 = v14
        }
        var v18 uint16 = 163
        v14 = (v14 - (v14 * float32(float64(v14))))
        _ = fi16
        _ = v18
    }
    if (true && false) {
        for fi19 := 0; fi19 < 5; fi19++ {
            v7 -= (v7 % (((v7 - 0)) | 1))
            var v20 uint16 = 724
            var v21 int64 = -9223372036854775808
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    for fi22 := 0; fi22 < 12; fi22++ {
        delete(v2, 4294967295)
        v7 += v7
        if true {
            fmt.Println(fi22, fi22, v7)
        } else {
            if (18446744073709551615 < 0) {
                v5 = (v5 & (-(v5)))
                if ((v7 + v7) == (v7 >> ((428) & 63))) {
                    fmt.Println(v5)
                    fmt.Println(v7, v5, v5)
                } else {
                    v14 = ((float32(1e10) + v14) / (v14 - v14))
                }
                v14 += v14
            }
            switch ("a" + "foo bar") {
            case "":
                fmt.Println(fi22)
            case "\n\t":
                fmt.Println(v14, fi22)
            }
        }
        _ = fi22
    }
    v23 := GenS2{F0: -9223372036854775808, F1: v7}
    v24 := GenS2{F0: -9223372036854775808, F1: 904}
    v25 := GenS1{F0: "", F1: 42, F2: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
