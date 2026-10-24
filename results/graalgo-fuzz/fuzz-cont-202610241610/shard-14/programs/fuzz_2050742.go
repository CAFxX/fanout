package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16 }
func genFn0() uint32 {
    var v4 uint32 = 2147483648
    var v5 uint16 = 32767
    v6 := []int{-1}
    v6 = append(v6, -2147483648)
    if false {
        fmt.Println(v5, v5)
        v4 = (v4 % (((^(v4))) | 1))
    }
    _ = v4
    _ = v5
    _ = v6
    return v4
}
func genFn1(p0 int, p1 string, p2 uint8) int8 {
    if ((582 >= 1) && true) {
        v7 := GenS1{F0: (((p1 + p1) + "") + (((p1 + "a") + "") + (((p1 + (p1 + (p1 + p1))) + "a") + p1))), F1: p0, F2: 741}
        _ = v7
    }
    switch "hello" {
    case "foo bar":
        var v8 uint32 = 2147483648
        _ = v8
    case "a":
        var v9 string = "foo bar"
        _ = v9
    case "hello":
        p2 = p2
    }
    if true {
        p1 = (((p1 + "foo bar") + (("" + p1) + "")) + p1)
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    v1 := []string{"foo bar", ("hello" + (("a" + "\n\t") + ((("" + "") + "foo bar") + "a"))), ("" + "a"), ("" + "\n\t")}
    var v2 float64 = -2.25
    v3 := []int{-1, 2147483647}
    switch ("" + "hello") {
    case "foo bar":
        _ = len(v1)
    case "":
        var v10 uint8 = 2
        _ = v10
    default:
        fmt.Println(v2)
    }
    v2 = v2
    for fi11 := 0; fi11 < 16; fi11++ {
        if false {
            for fi12 := 0; fi12 < 16; fi12++ {
                fmt.Println(v2)
                if len(v1) > 0 {
                    v1[0] = "\n\t"
                }
                fmt.Println(v2, fi11, fi12)
                _ = fi12
            }
        } else {
            var v13 int = 2147483647
            switch "hello" {
            case "\n\t":
                var v14 uint = 42
                _ = v14
            case "hello":
                fmt.Println(v2, v2)
            case "a":
                var v15 int = 127
                _ = v15
            default:
                switch "" {
                case "foo bar":
                    v16 := GenS2{F0: (1 >= -1)}
                    _ = v16
                case "":
                    switch ("" + "foo bar") {
                    case "hello":
                        v13 = (v13 - (fi11 + -67))
                    case "\n\t":
                        var v17 float32 = (0.1 - 0.1)
                        _ = v17
                    case "a":
                        var v18 int = (fi11 % ((v13) | 1))
                        _ = v18
                    default:
                        var v19 string = "hello"
                        _ = v19
                    }
                case "a":
                    if (638 != 1000) {
                        if true {
                            for fi20 := 0; fi20 < 12; fi20++ {
                                if true {
                                    v2 += (v2 * (0.0 / v2))
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                _ = fi20
                            }
                        } else {
                            var v21 int64 = 1
                            v21 -= v21
                            _ = v21
                        }
                    } else {
                        for fi22 := 0; fi22 < 16; fi22++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            fmt.Println(fi22)
                            var v23 float64 = v2
                            _ = fi22
                            _ = v23
                        }
                    }
                default:
                    var v24 int32 = 42
                    _ = v24
                }
            }
            if true {
                if (238 > 0) {
                    fmt.Println(fi11, fi11)
                }
            } else {
                for fi25 := 0; fi25 < 15; fi25++ {
                    fmt.Println(v2, v2)
                    v13 = fi25
                    v13 = fi11
                    _ = fi25
                }
                v2 = -2.25
                for fi26 := 0; fi26 < 8; fi26++ {
                    switch (v13 / ((fi11) | 1)) {
                    case 1:
                        v13 = fi26
                    case -9223372036854775808:
                        v2 = 0.1
                    }
                    var v27 uint = 447
                    _ = fi26
                    _ = v27
                }
            }
            _ = v13
        }
        switch "a" {
        case "hello":
            var v28 uint = 432
            _ = v28
        case "":
            _ = len(v1)
        default:
            v2 = v2
        }
        _ = len(v1)
        _ = fi11
    }
    _ = len(v3)
    var v29 int8 = 42
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v30 := GenS3{F0: -32768}
    var v31 uint32 = 118
    if ((42 != 1) && (2147483647 >= -1)) {
        fmt.Println(v2, v29)
        if true {
            var v32 float32 = 0.1
            v33 := GenS1{F0: "", F1: -1, F2: 18446744073709551615}
            _ = v32
            _ = v33
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
}
