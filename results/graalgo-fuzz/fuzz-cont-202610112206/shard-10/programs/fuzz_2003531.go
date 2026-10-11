package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 bool; F1 int32 }
func genFn0(p0 uint8, p1 uint32) uint64 {
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := GenS1{F0: 255, F1: (p1 << ((4294967295) & 63))}
        _ = fi5
        _ = v6
    }
    p1 -= (p1 % ((0) | 1))
    _ = p0
    _ = p1
    return 373
}

func main() {
    var v1 int64 = 1
    v2 := []uint32{1}
    var v3 uint8 = 128
    var v4 float64 = -2.25
    switch "foo bar" {
    case "foo bar":
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 int8 = 44
            switch 32768 {
            case 1:
                var v9 int32 = 2147483647
                _ = v9
            case 128:
                v1 += (v1 + v1)
            case 42:
                v4 += (0.0 / v4)
            default:
                var v10 int16 = -71
                _ = v10
            }
            var v11 string = "foo bar"
            _ = fi7
            _ = v11
            _ = v8
        }
    case "a":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v3 += v3
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        fmt.Println(v1, v1, v4)
        switch "\n\t" {
        case "\n\t":
            v12 := GenS1{F0: 744, F1: 4294967295}
            _ = v12
        case "foo bar":
            var v13 int64 = v1
            _ = v13
        case "":
            var v14 int64 = v1
            _ = v14
        default:
            v1 -= v1
        }
        switch "" {
        case "hello":
            switch "foo bar" {
            case "\n\t":
                switch -9223372036854775808 {
                case -1:
                    var v15 float64 = v4
                    _ = v15
                }
            case "foo bar":
                v16 := GenS1{F0: 18446744073709551615, F1: 685}
                _ = v16
            }
        case "\n\t":
            var v17 int = -2147483648
            _ = v17
        default:
            if true {
                if false {
                    var v18 int8 = -128
                    _ = v18
                }
            }
        }
    }
    fmt.Println(v4)
    switch ("" + ("" + "")) {
    case "a":
        v2 = append(v2, 0)
    case "":
        switch -130 {
        case -2147483648:
            switch ("a" + "foo bar") {
            case "a":
                var v19 uint32 = 2147483648
                _ = v19
            case "\n\t":
                v4 = v4
            default:
                v20 := GenS1{F0: 42, F1: 137}
                _ = v20
            }
        case -685:
            var v21 uint8 = v3
            _ = v21
        case -9223372036854775808:
            _ = len(v2)
        }
    case "foo bar":
        v3 -= 199
    default:
        var v22 bool = false
        _ = v22
    }
    v23 := GenS1{F0: 966, F1: 2147483648}
    var v24 string = "\n\t"
    v25 := GenS2{F0: v1, F1: 19}
    v24 = (v24 + v24)
    fmt.Println(v1, v1, v1)
    if true {
        var v26 uint = 1
        var v27 int8 = 47
        _ = v26
        _ = v27
    } else {
        switch (("a" + v24) + "a") {
        case "foo bar":
            var v28 bool = false
            _ = v28
        default:
            v2 = append(v2, 4294967295)
        }
    }
    v29 := GenS3{F0: true, F1: -2147483648}
    for fi30 := 0; fi30 < 11; fi30++ {
        fmt.Println(v3)
        _ = fi30
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v4
}
