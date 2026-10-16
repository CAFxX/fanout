package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
func genFn0(p0 uint64, p1 int8, p2 uint16) int64 {
    for fi4 := 0; fi4 < 10; fi4++ {
        var v5 uint64 = 852
        _ = fi4
        _ = v5
    }
    var v6 int8 = -81
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}

func main() {
    v1 := make(map[int]string)
    var v2 int64 = -9223372036854775808
    var v3 uint16 = 910
    v3 -= (v3 % (((v3 / ((v3) | 1))) | 1))
    var v7 bool = (49 < -9223372036854775808)
    var v8 uint64 = 394
    switch ("hello" + "\n\t") {
    case "foo bar":
        var v9 int = -31
        _ = v9
    case "hello":
        fmt.Println(v8)
    case "a":
        v2 -= v2
    default:
        if v7 {
            v8 = (v8 & v8)
            var v10 int16 = 1000
            var v11 int = 2147483647
            _ = v10
            _ = v11
        }
    }
    v12 := GenS1{F0: 866, F1: -679}
    switch ("foo bar" + "hello") {
    case "hello":
        v13 := []int{-1, -1, 4294967295}
        _ = v13
    }
    switch 2147483647 {
    case -339:
        if ((v8 == v8) || (v7 && (!v7))) {
            if ((^(v8)) <= 18446744073709551615) {
                v14 := []int{-2147483648, -510, 387, -2147483648}
                _ = v14
            }
        }
    case -282:
        for fi15 := 0; fi15 < 12; fi15++ {
            if false {
                v16 := []int{4294967295, fi15, (fi15 / (((fi15 | (^(fi15)))) | 1))}
                _ = v16
            } else {
                var v17 float64 = 0.0
                v17 -= (float64(-2.25) * v17)
                v7 = (false || (0 > fi15))
                _ = v17
            }
            fmt.Println(v7)
            _ = fi15
        }
    default:
        v3 += (v3 % ((v3) | 1))
    }
    switch "\n\t" {
    case "":
        v8 += 216
    case "\n\t":
        var v18 uint = 948
        _ = v18
    case "a":
        v3 = (v3 * 1)
    }
    var v19 int64 = (v2 ^ (v2 ^ -9223372036854775808))
    fmt.Println(v8, v8, v19)
    var v20 float32 = -2.25
    if v7 {
        if (2 < 1) {
            fmt.Println(v7, v3)
            var v21 float64 = (1e10 + (-2.25 * 1e10))
            fmt.Println(v7)
            _ = v21
        } else {
            switch 807 {
            case 32767:
                if (v7 && (42 != -151)) {
                    fmt.Println(v19, v7)
                }
            default:
                v22 := GenS1{F0: 42, F1: -497}
                _ = v22
            }
        }
        for fi23 := 0; fi23 < 13; fi23++ {
            v8 = v8
            for fi24 := 0; fi24 < 7; fi24++ {
                if (v7 || (!v7)) {
                    var v25 float64 = 1e10
                    var v26 uint32 = 816
                    _ = v25
                    _ = v26
                } else {
                    var v27 int64 = v19
                    var v28 int8 = 20
                    if (!v7) {
                        fmt.Println(v20)
                    } else {
                        fmt.Println(v28, v20, v7)
                    }
                    _ = v27
                    _ = v28
                }
                _ = fi24
            }
            _ = fi23
        }
    } else {
        v29 := GenS1{F0: 0, F1: 42}
        switch (("" + "a") + "foo bar") {
        case "foo bar":
            switch "\n\t" {
            case "\n\t":
                var v30 uint = 958
                _ = v30
            case "a":
                if (!true) {
                    v31 := []int{42, 9223372036854775807, 640, -2147483648}
                    _ = v31
                }
            default:
                for fi32 := 0; fi32 < 15; fi32++ {
                    v2 = (v2 + v19)
                    _ = fi32
                }
            }
        case "\n\t":
            v8 = v8
        default:
            v19 -= (v2 & int64(v3))
        }
        v33 := GenS1{F0: 0, F1: 500}
        _ = v29
        _ = v33
    }
    var v34 uint16 = 2
    v3 = v34
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v7
    _ = v8
}
