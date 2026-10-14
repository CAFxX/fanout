package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 uint; F2 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8; F1 uint32 }

func main() {
    var v1 int = -1
    var v2 uint32 = 2
    var v3 string = (("" + (("\n\t" + "a") + "")) + "foo bar")
    var v4 string = ((v3 + (v3 + "a")) + v3)
    var v5 int32 = -2147483648
    var v6 string = ("\n\t" + "foo bar")
    v7 := GenS1{F0: uint64(v2), F1: (v2 / ((v2) | 1)), F2: 138}
    for fi8 := 0; fi8 < 10; fi8++ {
        for fi9 := 0; fi9 < 5; fi9++ {
            fmt.Println(v3, v4)
            switch (^(fi8)) {
            case -2147483648:
                switch v3 {
                case "hello":
                    if ((-880 < 32767) && false) {
                        switch (fi9 << ((42) & 63)) {
                        case 587:
                            if (((-993 >= v5) && false) || (0 <= 311)) {
                                var v10 uint = uint(v2)
                                if (((^(v2)) > 65535) || false) {
                                    v11 := GenS2{F0: 32767}
                                    switch v4 {
                                    case "hello":
                                        v1 = (^(fi9))
                                    case "\n\t":
                                        var v12 int = v1
                                        _ = v12
                                    case "foo bar":
                                        v13 := []int{(fi9 << (((v10 % ((4294967295) | 1))) & 63)), v1}
                                        _ = v13
                                    }
                                    _ = v11
                                }
                                var v14 int64 = 42
                                v15 := []int{2147483647, 127}
                                _ = v10
                                _ = v14
                                _ = v15
                            } else {
                                v16 := GenS2{F0: 237}
                                v4 = (v6 + "hello")
                                v6 = (((v3 + v6) + "foo bar") + ("a" + ("hello" + v4)))
                                _ = v16
                            }
                        case 127:
                            fmt.Println(fi8)
                        }
                    } else {
                        var v17 int64 = -9223372036854775808
                        if true {
                            switch (v1 / (((fi8 * (v1 | (v1 * (-(fi8)))))) | 1)) {
                            case 127:
                                switch v6 {
                                case "hello":
                                    v18 := []int{}
                                    _ = v18
                                case "\n\t":
                                    v4 = (v4 + v4)
                                case "":
                                    var v19 uint = 255
                                    _ = v19
                                }
                            case -1:
                                switch "hello" {
                                case "foo bar":
                                    var v20 int = v1
                                    _ = v20
                                case "":
                                    var v21 int64 = 1
                                    _ = v21
                                }
                            default:
                                v4 = ""
                            }
                        }
                        var v22 float32 = -2.25
                        _ = v17
                        _ = v22
                    }
                case "foo bar":
                    if (409 == 956) {
                        var v23 int64 = -9223372036854775808
                        _ = v23
                    }
                case "":
                    var v24 float32 = float32(0.1)
                    _ = v24
                }
            case 0:
                v25 := GenS3{F0: (("foo bar" + ("a" + ("" + v6))) + v4), F1: 2, F2: 18446744073709551615}
                _ = v25
            default:
                if false {
                    switch fi8 {
                    case 128:
                        v6 = "hello"
                    case 9223372036854775807:
                        if (-9223372036854775808 != 42) {
                            v3 = ((("foo bar" + "hello") + ((("foo bar" + ((v4 + ("foo bar" + "")) + v4)) + v6) + v6)) + (v4 + v3))
                            var v26 int8 = 100
                            v27 := []int{v1, (fi9 << ((65535) & 63)), (fi8 & v1)}
                            _ = v26
                            _ = v27
                        } else {
                            switch -369 {
                            case 32768:
                                v6 = v6
                            case 4294967295:
                                fmt.Println(fi9, v5, v4)
                                switch v1 {
                                case 42:
                                    v5 = -2147483648
                                case 373:
                                    v5 -= -193
                                default:
                                    v28 := []int{42, fi8, -2147483648}
                                    _ = v28
                                }
                                v29 := []int{-2147483648, (fi9 ^ v1)}
                                _ = v29
                            case 127:
                                var v30 int8 = -49
                                _ = v30
                            }
                        }
                    }
                    v1 -= fi8
                    if (!((v5 % ((-2147483648) | 1)) == 1)) {
                        var v31 bool = (-39 != 127)
                        v32 := []int{(v1 % ((-585) | 1)), (v1 & v1)}
                        _ = v31
                        _ = v32
                    } else {
                        var v33 uint32 = 2
                        if (!true) {
                            switch v6 {
                            case "":
                                var v34 float32 = float32(1.5)
                                _ = v34
                            }
                            if true {
                                v6 = v4
                            }
                            fmt.Println(v4, v4, fi8)
                        }
                        _ = v33
                    }
                }
            }
            if true {
                fmt.Println(fi9)
                v35 := GenS4{F0: -580}
                v36 := GenS1{F0: 2, F1: v2, F2: 1000}
                _ = v35
                _ = v36
            } else {
                fmt.Println(v3)
                var v37 uint32 = 2147483648
                v38 := []int{}
                fmt.Println(v3, fi8, v2)
                v39 := GenS5{F0: 2, F1: (v2 | (v2 - 65535))}
                _ = v37
                _ = v38
                _ = v39
            }
            _ = fi9
        }
        _ = fi8
    }
    fmt.Println(v1)
    var v40 int32 = v5
    var v41 float64 = 1e10
    v41 += (float64(((3.14159 + 1.5) / float32((3.14159 / (((-2.25 * 1.5) + (float64(3.14159) + (0.0 * (((v41 - 0.1) - ((v41 + v41) + v41)) + ((v41 + 1.5) * 1e10))))) / v41))))) + float64(3.14159))
    v40 = v40
    v42 := GenS5{F0: 255, F1: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v7
}
