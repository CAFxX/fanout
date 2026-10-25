package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int; F1 int64; F2 string }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 string }
func genFn0(p0 uint64) uint16 {
    p0 -= (^(p0))
    v6 := []int{42}
    var v7 uint = 255
    p0 = 850
    _ = p0
    _ = v6
    _ = v7
    return 2
}

func main() {
    var v1 float32 = (0.1 - 0.1)
    v2 := []string{("a" + "hello"), "", "", ""}
    var v3 uint8 = 212
    v4 := []uint32{42}
    var v5 float32 = float32(float64(v1))
    _ = len(v2)
    var v8 uint = 621
    for fi9 := 0; fi9 < 19; fi9++ {
        v1 = v5
        v10 := GenS1{F0: (fi9 >> (((v8 << (((v8 * v8)) & 63))) & 63)), F1: 0}
        _ = fi9
        _ = v10
    }
    if (1 < -32768) {
        fmt.Println(v5)
        v8 = 279
    } else {
        switch "" {
        case "\n\t":
            v5 = 0.1
        case "":
            for fi11 := 0; fi11 < 7; fi11++ {
                var v12 string = ("foo bar" + (("hello" + "foo bar") + "foo bar"))
                var v13 uint64 = 42
                _ = fi11
                _ = v12
                _ = v13
            }
        case "hello":
            fmt.Println(v1, v3)
        }
        for fi14 := 0; fi14 < 17; fi14++ {
            if false {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                fmt.Println(v8, fi14, v1)
                v3 = (v3 & (v3 ^ v3))
            }
            fmt.Println(fi14, v8, fi14)
            _ = fi14
        }
        v8 += (v8 * uint(v3))
    }
    v15 := GenS1{F0: 4294967295, F1: 79}
    v1 = (1.5 / (((v1 - 0.1) * (v1 - v1)) - v5))
    if (42 > 18446744073709551615) {
        for fi16 := 0; fi16 < 6; fi16++ {
            switch ("foo bar" + "foo bar") {
            case "\n\t":
                v17 := GenS1{F0: 155, F1: 0}
                _ = v17
            }
            fmt.Println(v5)
            v8 = (v8 | (v8 / ((v8) | 1)))
            _ = fi16
        }
        fmt.Println(v1)
        var v18 int64 = 42
        _ = v18
    } else {
        v5 += 3.14159
        v5 = (v5 * v5)
        switch int(v8) {
        case 128:
            if true {
                v19 := GenS2{F0: 2147483647, F1: -1, F2: "\n\t"}
                _ = v19
            }
        case 32768:
            if len(v2) > 0 {
                v2[0] = "\n\t"
            }
        default:
            if (2 == 0) {
                var v20 bool = (0 > 42)
                _ = v20
            } else {
                switch 4294967295 {
                case -926:
                    var v21 int64 = 1
                    _ = v21
                case 32767:
                    var v22 bool = (1 != 65535)
                    _ = v22
                }
            }
        }
    }
    fmt.Println(v5, v5, v5)
    switch "\n\t" {
    case "\n\t":
        v23 := GenS3{F0: "\n\t", F1: 42}
        _ = v23
    case "foo bar":
        var v24 uint16 = 32767
        _ = v24
    default:
        for fi25 := 0; fi25 < 10; fi25++ {
            v26 := GenS4{F0: ""}
            _ = fi25
            _ = v26
        }
    }
    v27 := GenS4{F0: ""}
    switch 0 {
    case 0:
        fmt.Println(v8, v5, v5)
    case 32768:
        v1 -= (1.5 + v5)
    }
    v28 := GenS2{F0: 32768, F1: -9223372036854775808, F2: "foo bar"}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
