package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 255
    var v3 uint64 = 867
    var v4 int64 = 0
    v5 := make(map[string]string)
    if false {
        _ = len(v1)
        var v6 uint32 = 764
        switch "\n\t" {
        case "hello":
            v4 -= v4
        default:
            switch "hello" {
            case "foo bar":
                v1[("hello" + (("a" + "hello") + "a"))] = 419
            case "hello":
                var v7 int8 = int8(v6)
                _ = v7
            }
        }
        _ = v6
    }
    if (uint8(v3) < 255) {
        v2 -= (v2 + 14)
        if ((1 >= 4) && true) {
            for fi8 := 0; fi8 < 8; fi8++ {
                fmt.Println(v4, v4)
                _ = fi8
            }
        }
        var v9 float64 = 0.0
        _ = v9
    }
    switch 32768 {
    case 32767:
        var v10 int16 = 42
        _ = v10
    }
    var v11 uint64 = 792
    fmt.Println(v3, v2)
    v11 -= v11
    v2 = v2
    var v12 uint64 = v3
    if (false && false) {
        switch 128 {
        case 4294967295:
            fmt.Println(v12, v12, v3)
        }
    } else {
        v12 -= (v11 * (v3 << ((65535) & 63)))
    }
    fmt.Println(v3, v3)
    switch -9223372036854775808 {
    case -2147483648:
        if (!false) {
            if false {
                var v13 uint8 = 128
                _ = v13
            }
        } else {
            for fi14 := 0; fi14 < 12; fi14++ {
                for fi15 := 0; fi15 < 18; fi15++ {
                    var v16 uint64 = 18446744073709551615
                    var v17 float64 = 0.0
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                v18 := []int{42, fi14, 32767, fi14}
                v19 := GenS1{F0: -2147483648, F1: -1}
                _ = fi14
                _ = v18
                _ = v19
            }
        }
    case 127:
        if (false && false) {
            switch 1 {
            case -203:
                if true {
                    fmt.Println(v2)
                    for fi20 := 0; fi20 < 7; fi20++ {
                        for fi21 := 0; fi21 < 7; fi21++ {
                            var v22 uint8 = 1
                            var v23 uint8 = 2
                            var v24 int = (fi20 ^ (fi20 / ((463) | 1)))
                            _ = fi21
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                        _ = fi20
                    }
                    v3 += v12
                } else {
                    var v25 int = 954
                    v3 = 1
                    fmt.Println(v11, v4, v12)
                    _ = v25
                }
            }
        }
    case 757:
        v3 = v3
    default:
        for fi26 := 0; fi26 < 19; fi26++ {
            v3 = (v12 % (((v3 * v3)) | 1))
            _ = fi26
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
