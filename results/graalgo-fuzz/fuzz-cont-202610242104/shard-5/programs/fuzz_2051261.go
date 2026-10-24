package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 int64; F1 uint32; F2 uint32 }
func genFn0(p0 uint32) int16 {
    p0 -= (^(p0))
    switch 32767 {
    case -9223372036854775808:
        p0 -= p0
    default:
        var v7 int8 = 0
        _ = v7
    }
    p0 = (^(p0))
    v8 := GenS1{F0: (^(p0))}
    var v9 int64 = 184
    _ = p0
    _ = v8
    _ = v9
    return 1000
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = -601
    var v3 float32 = 1e10
    var v4 float32 = ((-2.25 / v3) + (0.0 + v3))
    v5 := []int{-9223372036854775808}
    var v6 uint32 = 42
    var v10 int8 = -1
    v4 += v4
    fmt.Println(v3, v4)
    if (((("a" + ("a" + "foo bar")) + "a") + "") > "a") {
        v11 := GenS2{F0: 4294967295, F1: 4294967295}
        fmt.Println(v6, v3, v2)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v11
    } else {
        var v12 bool = (204 == 458)
        _ = v12
    }
    switch "a" {
    case "\n\t":
        var v13 int = 9223372036854775807
        _ = v13
    default:
        if false {
            if false {
                var v14 uint = 357
                switch "foo bar" {
                case "foo bar":
                    for fi15 := 0; fi15 < 4; fi15++ {
                        if false {
                            v5 = append(v5, (fi15 / (((fi15 ^ (fi15 + fi15))) | 1)))
                        }
                        _ = fi15
                    }
                case "a":
                    var v16 uint = (^(v14))
                    _ = v16
                case "hello":
                    v6 = 2147483648
                }
                fmt.Println(v4)
                _ = v14
            } else {
                if (false && (v2 <= v2)) {
                    if false {
                        fmt.Println(v4)
                        v10 += (v10 << ((18446744073709551615) & 63))
                    } else {
                        fmt.Println(v6, v3, v2)
                        var v17 uint = 524
                        _ = v17
                    }
                } else {
                    var v18 uint32 = (v6 % (((v6 | (^(v6)))) | 1))
                    _ = v18
                }
                fmt.Println(v6, v4)
            }
        } else {
            switch (("foo bar" + "foo bar") + "") {
            case "\n\t":
                var v19 int64 = 112
                _ = v19
            case "foo bar":
                fmt.Println(v4, v10)
            default:
                var v20 bool = ("hello" <= "")
                _ = v20
            }
        }
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        _ = len(v1)
        _ = fi21
    }
    v3 += v3
    for fi22 := 0; fi22 < 19; fi22++ {
        fmt.Println(v6, v10, v10)
        fmt.Println(v2, v10)
        _ = fi22
    }
    var v23 int = 983
    if ((v6 * (v6 / ((v6) | 1))) != 4294967295) {
        var v24 int32 = 805
        _ = v24
    } else {
        v25 := v1[v6]
        _ = v25
    }
    fmt.Println(v10, v4, v23)
    if (!true) {
        v5 = append(v5, (v23 << ((1) & 63)))
    } else {
        v26 := GenS3{F0: (^(v2)), F1: (v6 << ((667) & 63)), F2: (v6 ^ v6)}
        _ = v26
    }
    v10 = (v10 - (v10 & v10))
    v2 -= v2
    v6 -= (v6 - (^(v6)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
