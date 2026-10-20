package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 int64; F1 int16 }
func genFn0(p0 int, p1 int8, p2 uint) int32 {
    fmt.Println(p2)
    fmt.Println(p2)
    p0 = (p0 ^ (p0 << ((18446744073709551615) & 63)))
    _ = p0
    _ = p1
    _ = p2
    return 546
}

func main() {
    var v1 int64 = 42
    var v2 int8 = 1
    var v3 uint64 = 812
    v4 := make(map[int]uint32)
    var v5 bool = false
    var v6 bool = (v3 != v3)
    if (-876 <= 238) {
        v4[-9223372036854775808] = 0
        v3 -= v3
        v7 := []int{-2147483648}
        _ = v7
    }
    if v5 {
        v6 = v5
        for fi8 := 0; fi8 < 2; fi8++ {
            fmt.Println(v1, v1, v6)
            _ = fi8
        }
    } else {
        v9 := GenS1{F0: (^(v1)), F1: 4294967295}
        _ = v9
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        v11 := []int{9223372036854775807, fi10}
        fmt.Println(v2)
        _ = fi10
        _ = v11
    }
    var v12 bool = (!v6)
    v2 = v2
    if true {
        fmt.Println(v1, v2)
    } else {
        var v13 string = "hello"
        fmt.Println(v12, v12, v12)
        _ = v13
    }
    v4[1] = 65535
    if (!v6) {
        if v12 {
            if v12 {
                if (255 <= 2) {
                    v14 := GenS1{F0: v1, F1: 998}
                    fmt.Println(v6, v3, v2)
                    if (!true) {
                        for fi15 := 0; fi15 < 16; fi15++ {
                            var v16 bool = ((v3 & 2) != (v3 - (v3 << ((4294967295) & 63))))
                            _ = fi15
                            _ = v16
                        }
                        for fi17 := 0; fi17 < 14; fi17++ {
                            v12 = false
                            switch "foo bar" {
                            case "\n\t":
                                var v18 uint16 = 255
                                _ = v18
                            }
                            _ = fi17
                        }
                    }
                    _ = v14
                }
                v19 := GenS1{F0: (v1 + (v1 >> ((65535) & 63))), F1: 4294967295}
                var v20 int8 = -1
                _ = v19
                _ = v20
            }
            v21 := []int{-961, 128, -1}
            _ = v21
        } else {
            v2 -= v2
            v5 = (false && (42 < (v3 * (v3 | 2))))
        }
        var v22 float64 = (1.5 * (((-2.25 + ((1.5 / -2.25) + 1e10)) - (0.1 - 0.0)) / float64(1e10)))
        _ = v22
    } else {
        var v23 float64 = 1.5
        switch "a" {
        case "hello":
            v23 -= (-2.25 + v23)
        }
        _ = v23
    }
    var v24 uint64 = 703
    v25 := GenS2{F0: (v1 << ((2) & 63)), F1: 32767}
    v5 = (!(-1 < 1))
    v26 := []int{0}
    v27 := GenS2{F0: -592, F1: 1}
    v2 += -1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
