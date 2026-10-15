package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int32; F1 int16 }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 int32 }

func main() {
    var v1 uint8 = 42
    var v2 int16 = 32767
    var v3 uint16 = 2
    var v4 int16 = 1000
    var v5 int8 = 100
    var v6 uint32 = 2147483648
    v2 = (v2 * (v4 + int16(v6)))
    var v7 bool = false
    fmt.Println(v6, v3, v2)
    v8 := []int{}
    for fi9 := 0; fi9 < 3; fi9++ {
        for fi10 := 0; fi10 < 5; fi10++ {
            if ((578 != (v4 / ((-306) | 1))) && (v6 > 2147483648)) {
                var v11 float64 = (0.1 - ((1e10 + -2.25) + float64((1e10 * 0.0))))
                _ = v11
            }
            v12 := GenS1{F0: (fi9 + int(v5))}
            if (1.5 == 1e10) {
                var v13 int64 = -106
                if ((fi9 << ((632) & 63)) != fi9) {
                    if true {
                        v7 = (fi9 != (fi9 >> ((200) & 63)))
                    } else {
                        var v14 int64 = v13
                        if len(v8) > 0 {
                            v8[0] = (fi9 ^ (-(fi10)))
                        }
                        var v15 int = (fi9 % ((fi10) | 1))
                        _ = v14
                        _ = v15
                    }
                } else {
                    v2 += v2
                    switch "" {
                    case "hello":
                        if ((fi9 | 4294967295) <= -2147483648) {
                            _ = len(v8)
                            fmt.Println(v4)
                        } else {
                            if (v7 && v7) {
                                v2 -= (v4 + (v4 % (((^(v4))) | 1)))
                                var v16 uint = 2
                                _ = v16
                            } else {
                                _ = len(v8)
                                v6 += v6
                            }
                            var v17 uint = 434
                            _ = len(v8)
                            _ = v17
                        }
                    case "a":
                        v7 = (uint(v3) > 2)
                    case "foo bar":
                        fmt.Println(v2, v7)
                    default:
                        v18 := GenS2{F0: (v6 % ((v6) | 1)), F1: 32767}
                        _ = v18
                    }
                }
                _ = v13
            } else {
                v19 := GenS1{F0: -2147483648}
                _ = v19
            }
            _ = fi10
            _ = v12
        }
        _ = fi9
    }
    if (0.0 < 0.1) {
        if (v6 != v6) {
            var v20 int8 = (v5 / ((v5) | 1))
            _ = v20
        }
        v6 = v6
    } else {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    }
    var v21 uint64 = 42
    v8 = append(v8, -1)
    switch (("hello" + ("foo bar" + "hello")) + "foo bar") {
    case "hello":
        v3 += (v3 + v3)
    default:
        var v22 uint8 = (v1 * v1)
        _ = v22
    }
    v21 -= (v21 - v21)
    if (!(!(9223372036854775807 == 32768))) {
        var v23 string = ""
        v24 := GenS3{F0: -2147483648, F1: (v2 & (v4 ^ (v4 << ((18446744073709551615) & 63))))}
        _ = v23
        _ = v24
    } else {
        fmt.Println(v4, v6)
    }
    switch ("" + "\n\t") {
    case "foo bar":
        v25 := GenS1{F0: -240}
        _ = v25
    case "a":
        if v7 {
            switch "a" {
            case "foo bar":
                v21 += v21
            case "hello":
                if len(v8) > 0 {
                    v8[0] = -1
                }
            }
            for fi26 := 0; fi26 < 1; fi26++ {
                fmt.Println(v3, v2, v2)
                fmt.Println(v21, fi26, v1)
                _ = fi26
            }
            var v27 float32 = 0.1
            _ = v27
        }
    default:
        v7 = false
    }
    v1 += v1
    for fi28 := 0; fi28 < 18; fi28++ {
        if ((v1 == v1) && (42 != 491)) {
            fmt.Println(v2, v6, v1)
            v29 := GenS4{F0: 2147483647, F1: (!true)}
            fmt.Println(v6, fi28, v2)
            _ = v29
        }
        var v30 int64 = 42
        _ = fi28
        _ = v30
    }
    v31 := GenS5{F0: 42}
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
