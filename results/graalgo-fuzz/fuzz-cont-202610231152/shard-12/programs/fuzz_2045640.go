package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint8; F1 uint; F2 string }
type GenS3 struct { F0 int32 }
func genFn0(p0 int16, p1 uint32) int64 {
    fmt.Println(p1, p0)
    var v5 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = v5
    return v5
}
func genFn1(p0 uint64, p1 int64, p2 uint32) int32 {
    var v6 uint32 = p2
    fmt.Println(p2, p1, p0)
    if (!false) {
        v7 := []int{212}
        for fi8 := 0; fi8 < 12; fi8++ {
            for fi9 := 0; fi9 < 11; fi9++ {
                switch "a" {
                case "a":
                    switch ((("foo bar" + "a") + ("" + "hello")) + "a") {
                    case "":
                        v10 := GenS1{F0: 100, F1: (v6 + (v6 / ((p2) | 1))), F2: -2147483648}
                        _ = v10
                    default:
                        fmt.Println(p1)
                    }
                }
                if false {
                    _ = len(v7)
                } else {
                    v11 := GenS1{F0: 87, F1: (p2 << ((0) & 63)), F2: 460}
                    if false {
                        p2 -= 237
                        v12 := GenS1{F0: 1, F1: p2, F2: 42}
                        var v13 string = ""
                        _ = v12
                        _ = v13
                    } else {
                        v14 := GenS1{F0: 42, F1: p2, F2: 1}
                        p1 -= (p1 ^ 0)
                        _ = v14
                    }
                    if false {
                        p2 -= (p2 & uint32(p0))
                    }
                    _ = v11
                }
                v6 += (p2 | p2)
                _ = fi9
            }
            p0 = (p0 * (p0 % ((2) | 1)))
            v7 = append(v7, fi8)
            _ = fi8
        }
        p1 = (-(p1))
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 531
}

func main() {
    v1 := []uint32{1}
    var v2 int32 = 17
    var v3 uint64 = 42
    var v4 int16 = -1
    for fi15 := 0; fi15 < 3; fi15++ {
        v1 = append(v1, 0)
        _ = fi15
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi16 := 0; fi16 < 14; fi16++ {
        switch "\n\t" {
        case "a":
            v3 -= v3
        case "hello":
            v17 := GenS1{F0: -1, F1: uint32(v3), F2: (v2 * v2)}
            _ = v17
        case "":
            for fi18 := 0; fi18 < 3; fi18++ {
                if true {
                    if (fi18 != (fi18 >> ((42) & 63))) {
                        v3 = (v3 % ((131) | 1))
                        if true {
                            var v19 int16 = -32768
                            var v20 int8 = 41
                            _ = v19
                            _ = v20
                        } else {
                            if len(v1) > 0 {
                                v1[0] = 938
                            }
                            fmt.Println(v3)
                            if len(v1) > 0 {
                                v1[0] = 542
                            }
                        }
                        fmt.Println(v4)
                        var v21 int32 = (v2 & (v2 | v2))
                        _ = v21
                    } else {
                        var v22 uint64 = 0
                        _ = v22
                    }
                    var v23 bool = false
                    _ = v23
                }
                _ = fi18
            }
        }
        fmt.Println(fi16, v4)
        for fi24 := 0; fi24 < 11; fi24++ {
            v3 = v3
            v3 = 0
            v2 -= v2
            _ = fi24
        }
        _ = fi16
    }
    v3 = v3
    v3 += (v3 * (v3 - 0))
    v4 = (^(v4))
    fmt.Println(v3, v2, v3)
    fmt.Println(v4)
    var v25 int32 = v2
    for fi26 := 0; fi26 < 5; fi26++ {
        _ = len(v1)
        switch "foo bar" {
        case "hello":
            var v27 uint32 = 1
            _ = v27
        case "\n\t":
            for fi28 := 0; fi28 < 5; fi28++ {
                var v29 int8 = 100
                _ = fi28
                _ = v29
            }
        case "":
            _ = len(v1)
        }
        switch ("foo bar" + "a") {
        case "a":
            fmt.Println(v3, v25)
        case "":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case "hello":
            v25 = v2
        default:
            if ((1.5 * 0.1) < (1e10 * 1e10)) {
                v30 := GenS2{F0: 255, F1: 18446744073709551615, F2: ("" + ("" + "hello"))}
                v31 := GenS3{F0: (v2 ^ 792)}
                _ = v30
                _ = v31
            }
        }
        _ = fi26
    }
    if (32767 == 128) {
        fmt.Println(v2, v25, v2)
        for fi32 := 0; fi32 < 15; fi32++ {
            for fi33 := 0; fi33 < 15; fi33++ {
                var v34 uint32 = 42
                fmt.Println(fi32)
                _ = fi33
                _ = v34
            }
            v35 := GenS2{F0: 255, F1: 999, F2: ((("\n\t" + ("hello" + "foo bar")) + "foo bar") + (("\n\t" + "a") + "\n\t"))}
            _ = fi32
            _ = v35
        }
    }
    v3 = 1
    var v36 string = "a"
    if (((true || false) && false) && false) {
        fmt.Println(v2)
    }
    for fi37 := 0; fi37 < 7; fi37++ {
        v3 = (v3 * (v3 / ((v3) | 1)))
        _ = fi37
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v36
    _ = v4
}
