package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 uint16; F1 int64; F2 uint }
type GenS4 struct { F0 int; F1 string; F2 uint }

func main() {
    var v1 int8 = 42
    var v2 int = 32768
    v3 := []string{"foo bar"}
    var v4 float64 = (1e10 - ((0.1 - (0.0 + (((((1e10 - (0.0 / 1e10)) * 1.5) - (0.1 - 0.1)) * (0.1 - (-2.25 / 0.1))) - 1.5))) - 0.1))
    v5 := []int{(-(v2)), (v2 % (((v2 & -617)) | 1))}
    var v6 uint32 = 1
    if (false && ((v2 << ((288) & 63)) < 42)) {
        _ = len(v3)
        v7 := GenS1{F0: 2147483647, F1: -2147483648, F2: 42}
        _ = v7
    } else {
        v4 -= 0.0
    }
    var v8 int16 = -499
    if (!(42 < -1)) {
        var v9 bool = true
        fmt.Println(v1, v2)
        var v10 int16 = (v8 & v8)
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 7; fi11++ {
        for fi12 := 0; fi12 < 19; fi12++ {
            if (float32(v4) != (((-2.25 - (1e10 + 1e10)) - (0.0 + 1e10)) / (0.1 * (float32(1e10) * (0.1 - 1e10))))) {
                switch 32768 {
                case 9223372036854775807:
                    switch "hello" {
                    case "":
                        if true {
                            if (!false) {
                                var v13 int8 = 1
                                v14 := GenS2{F0: 2147483648, F1: ""}
                                _ = v13
                                _ = v14
                            }
                            v1 += (v1 >> ((18446744073709551615) & 63))
                        } else {
                            v3 = append(v3, "a")
                            v15 := GenS3{F0: 64, F1: -9223372036854775808, F2: 65535}
                            var v16 uint = 18446744073709551615
                            _ = v15
                            _ = v16
                        }
                    case "a":
                        fmt.Println(v2)
                    }
                case 316:
                    if (true && (2 != 603)) {
                        switch v2 {
                        case 42:
                            var v17 string = "hello"
                            _ = v17
                        }
                    }
                    v8 = (v8 - v8)
                }
                v18 := GenS1{F0: 1, F1: fi12, F2: 4294967295}
                var v19 uint = 1
                _ = v18
                _ = v19
            } else {
                v8 = 1
                switch "\n\t" {
                case "foo bar":
                    if len(v3) > 0 {
                        v3[0] = ""
                    }
                case "\n\t":
                    if false {
                        var v20 uint16 = 65535
                        v6 += v6
                        _ = v20
                    }
                default:
                    if len(v5) > 0 {
                        v5[0] = fi12
                    }
                }
            }
            fmt.Println(v8)
            var v21 string = "foo bar"
            _ = fi12
            _ = v21
        }
        _ = fi11
    }
    switch "" {
    case "a":
        for fi22 := 0; fi22 < 7; fi22++ {
            v23 := GenS3{F0: 255, F1: 908, F2: 4294967295}
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi22
            _ = v23
        }
    case "foo bar":
        for fi24 := 0; fi24 < 19; fi24++ {
            if (false && true) {
                var v25 uint32 = (v6 & v6)
                if (true || true) {
                    if false {
                        fmt.Println(v4, v25, v2)
                    } else {
                        for fi26 := 0; fi26 < 11; fi26++ {
                            var v27 int16 = (v8 / ((42) | 1))
                            _ = fi26
                            _ = v27
                        }
                        for fi28 := 0; fi28 < 3; fi28++ {
                            var v29 int16 = v8
                            if true {
                                fmt.Println(v4)
                            }
                            _ = fi28
                            _ = v29
                        }
                        switch fi24 {
                        case 1:
                            for fi30 := 0; fi30 < 13; fi30++ {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = fi30
                            }
                        default:
                            if ((v25 / ((2) | 1)) == v25) {
                                if len(v5) > 0 {
                                    v5[0] = 42
                                }
                            }
                        }
                    }
                    var v31 float64 = v4
                    switch fi24 {
                    case 367:
                        for fi32 := 0; fi32 < 10; fi32++ {
                            if (false && true) {
                                var v33 uint64 = 18446744073709551615
                                fmt.Println(v2, v31)
                                _ = v33
                            }
                            switch "a" {
                            case "hello":
                                var v34 uint = 2
                                _ = v34
                            case "a":
                                switch "a" {
                                case "":
                                    switch "hello" {
                                    case "\n\t":
                                        v35 := GenS4{F0: int(v1), F1: ("" + (("" + "hello") + "\n\t")), F2: 515}
                                        _ = v35
                                    }
                                default:
                                    v2 = -9223372036854775808
                                }
                            }
                            fmt.Println(v1, v6)
                            _ = fi32
                        }
                    }
                    _ = v31
                } else {
                    fmt.Println(v2)
                }
                _ = len(v5)
                _ = v25
            }
            var v36 uint64 = 994
            _ = fi24
            _ = v36
        }
    default:
        var v37 string = ("" + "")
        _ = v37
    }
    v1 = v1
    v1 = (v1 + (^(v1)))
    var v38 bool = false
    var v39 uint16 = 1
    switch (v2 % (((v2 << ((18446744073709551615) & 63))) | 1)) {
    case -1:
        if (0.1 == 0.1) {
            _ = len(v3)
        }
    }
    switch "foo bar" {
    case "hello":
        var v40 int32 = 42
        _ = v40
    case "":
        if len(v3) > 0 {
            v3[0] = ((("" + "foo bar") + ("a" + ((("a" + "") + ("foo bar" + "a")) + ("\n\t" + "")))) + "hello")
        }
    }
    v39 += (v39 ^ (v39 + (v39 >> ((1) & 63))))
    if (-1 != (v1 >> ((1) & 63))) {
        v41 := GenS3{F0: 184, F1: -1, F2: uint(v39)}
        fmt.Println(v4, v39)
        _ = v41
    }
    switch ("hello" + ("a" + "foo bar")) {
    case "":
        var v42 int8 = v1
        _ = v42
    default:
        var v43 int8 = v1
        _ = v43
    }
    fmt.Println(v38)
    if v38 {
        var v44 uint64 = 1
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
