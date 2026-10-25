package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int32 }
type GenS2 struct { F0 int }
func genFn0(p0 uint8) int32 {
    var v7 int8 = 42
    fmt.Println(v7, p0)
    fmt.Println(p0, v7)
    switch -153 {
    case 42:
        var v8 float64 = 1e10
        _ = v8
    case -1:
        switch 532 {
        case 9223372036854775807:
            p0 = 127
        case 667:
            var v9 int32 = 142
            _ = v9
        case 566:
            v10 := []int{9223372036854775807, 9223372036854775807, 804, 2147483647}
            _ = v10
        }
    }
    _ = p0
    _ = v7
    return -2147483648
}
func genFn1(p0 bool) int64 {
    p0 = p0
    v11 := GenS1{F0: 473, F1: 255, F2: -736}
    _ = p0
    _ = v11
    return -124
}

func main() {
    var v1 int16 = 920
    v2 := []uint32{42, 4294967295, 626, 370}
    var v3 bool = (2147483647 < -915)
    v4 := make(map[uint32]string)
    v5 := make(map[string]bool)
    var v6 int8 = -47
    fmt.Println(v6)
    var v12 int = -1
    v6 = -35
    v3 = v3
    switch ("a" + "foo bar") {
    case "a":
        v1 -= int16(v12)
    case "\n\t":
        v12 = v12
    default:
        if true {
            v3 = false
        }
    }
    if v3 {
        v13 := GenS2{F0: (v12 + (-(v12)))}
        _ = len(v5)
        fmt.Println(v12)
        _ = v13
    }
    v14 := GenS1{F0: v1, F1: 1, F2: 0}
    switch v12 {
    case -2147483648:
        for fi15 := 0; fi15 < 9; fi15++ {
            v1 -= 1
            switch (("\n\t" + "") + ("hello" + "foo bar")) {
            case "":
                _ = len(v4)
            case "a":
                var v16 uint = 378
                _ = v16
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi15
        }
    case -9223372036854775808:
        var v17 int = v12
        _ = v17
    default:
        for fi18 := 0; fi18 < 7; fi18++ {
            v1 -= (v1 | (v1 | 1000))
            switch ("" + ((("" + "a") + "foo bar") + "foo bar")) {
            case "hello":
                var v19 uint = 0
                _ = v19
            case "\n\t":
                for fi20 := 0; fi20 < 9; fi20++ {
                    var v21 bool = (!false)
                    fmt.Println(fi18, fi18, fi18)
                    _ = fi20
                    _ = v21
                }
            case "foo bar":
                for fi22 := 0; fi22 < 10; fi22++ {
                    v23 := GenS1{F0: v1, F1: 255, F2: -2147483648}
                    var v24 uint64 = 4294967295
                    fmt.Println(v6, v24)
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            default:
                var v25 int64 = 9223372036854775807
                _ = v25
            }
            _ = fi18
        }
    }
    v1 += (v1 / ((v1) | 1))
    v12 -= v12
    v6 -= v6
    if (34 >= 0) {
        if len(v2) > 0 {
            v2[0] = 510
        }
        if (-353 > 9223372036854775807) {
            fmt.Println(v1, v12)
            v2 = append(v2, 1)
            switch "foo bar" {
            case "hello":
                v1 += (v1 - (v1 >> ((706) & 63)))
            case "foo bar":
                var v26 bool = false
                _ = v26
            }
        }
    }
    v1 -= v1
    v12 = (v12 | (^(v12)))
    if v3 {
        var v27 int8 = (v6 | (v6 << ((18446744073709551615) & 63)))
        _ = v27
    }
    _ = len(v2)
    v3 = (51 <= 426)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
