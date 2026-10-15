package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 string; F2 string }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8; F1 int64; F2 uint16 }
type GenS6 struct { F0 int16 }

func main() {
    var v1 string = "hello"
    var v2 int = 258
    var v3 float32 = 1e10
    v4 := make(map[string]bool)
    fmt.Println(v1, v1, v3)
    if true {
        switch "hello" {
        case "a":
            for fi5 := 0; fi5 < 17; fi5++ {
                for fi6 := 0; fi6 < 8; fi6++ {
                    var v7 uint16 = 32767
                    _ = fi6
                    _ = v7
                }
                _ = fi5
            }
        case "hello":
            var v8 int32 = 565
            _ = v8
        case "foo bar":
            for fi9 := 0; fi9 < 19; fi9++ {
                v3 -= (0.0 + float32(((3.14159 + 1e10) + 1e10)))
                v10 := v4["foo bar"]
                _ = len(v4)
                _ = fi9
                _ = v10
            }
        }
        v11 := GenS1{F0: 128}
        _ = v11
    } else {
        switch v2 {
        case 2147483647:
            var v12 float32 = ((((((v3 - (v3 + (v3 * (v3 * (0.1 + v3))))) - float32((1e10 + 0.1))) / v3) / 1e10) + 3.14159) + v3)
            _ = v12
        case 128:
            var v13 uint32 = 4294967295
            _ = v13
        }
        for fi14 := 0; fi14 < 18; fi14++ {
            fmt.Println(v2)
            switch fi14 {
            case 32767:
                fmt.Println(v2, fi14)
            case 2147483647:
                fmt.Println(fi14, v3)
            case -9223372036854775808:
                v1 = v1
            default:
                v15 := []int{-9223372036854775808, (v2 << ((755) & 63)), (fi14 << ((584) & 63)), 1}
                _ = v15
            }
            _ = fi14
        }
    }
    if false {
        var v16 uint32 = 655
        for fi17 := 0; fi17 < 19; fi17++ {
            fmt.Println(fi17, v16, v1)
            var v18 uint = 4294967295
            var v19 int8 = int8(v2)
            _ = fi17
            _ = v18
            _ = v19
        }
        v3 -= float32(0.0)
        _ = v16
    } else {
        v20 := GenS1{F0: 2}
        var v21 uint16 = 65535
        _ = v20
        _ = v21
    }
    v22 := []int{v2}
    if (448 > 18446744073709551615) {
        var v23 float32 = v3
        fmt.Println(v23, v23)
        _ = len(v22)
        _ = v23
    } else {
        _ = len(v4)
        for fi24 := 0; fi24 < 14; fi24++ {
            if false {
                var v25 int64 = 42
                for fi26 := 0; fi26 < 9; fi26++ {
                    v27 := GenS2{F0: fi26, F1: (v1 + "hello"), F2: ("\n\t" + "foo bar")}
                    _ = len(v4)
                    fmt.Println(v2, fi24)
                    _ = fi26
                    _ = v27
                }
                v3 = v3
                _ = v25
            } else {
                if (859 <= 0) {
                    fmt.Println(v3)
                    v22 = append(v22, fi24)
                    var v28 int64 = 42
                    _ = v28
                }
                for fi29 := 0; fi29 < 12; fi29++ {
                    var v30 uint = 18446744073709551615
                    if (!(1 <= 32767)) {
                        v31 := GenS3{F0: (0.1 > 1.5), F1: -2147483648}
                        _ = v31
                    } else {
                        switch fi24 {
                        case 32768:
                            v3 += (v3 / float32(((float64(v3) / (-2.25 * 1e10)) + 1.5)))
                        case -1:
                            switch (fi29 / (((v2 >> ((v30) & 63))) | 1)) {
                            case 216:
                                fmt.Println(v3, v30, v30)
                            default:
                                switch v1 {
                                case "hello":
                                    switch (v1 + ((v1 + v1) + v1)) {
                                    case "hello":
                                        var v32 uint = v30
                                        _ = v32
                                    case "\n\t":
                                        var v33 uint32 = 42
                                        _ = v33
                                    }
                                default:
                                    var v34 float32 = v3
                                    _ = v34
                                }
                            }
                        case 4294967295:
                            switch -528 {
                            case 500:
                                var v35 uint = v30
                                _ = v35
                            case 127:
                                var v36 uint = v30
                                switch ("a" + v1) {
                                case "\n\t":
                                    v37 := GenS4{F0: (v2 | -761)}
                                    _ = v37
                                }
                                var v38 int = v2
                                _ = v36
                                _ = v38
                            case -9223372036854775808:
                                var v39 uint16 = 999
                                _ = v39
                            }
                        default:
                            switch fi24 {
                            case 1:
                                var v40 uint32 = 2147483647
                                _ = v40
                            }
                        }
                        switch (^(fi24)) {
                        case -391:
                            v41 := GenS2{F0: (fi29 << ((v30) & 63)), F1: "foo bar", F2: "a"}
                            _ = v41
                        case 32768:
                            v42 := GenS5{F0: 218, F1: -9223372036854775808, F2: 32767}
                            _ = v42
                        case -1:
                            fmt.Println(v1, v30)
                        default:
                            var v43 bool = false
                            delete(v4, "hello")
                            _ = v43
                        }
                    }
                    var v44 uint64 = 2
                    _ = fi29
                    _ = v30
                    _ = v44
                }
                for fi45 := 0; fi45 < 19; fi45++ {
                    var v46 uint64 = 4294967295
                    var v47 uint8 = 0
                    _ = fi45
                    _ = v46
                    _ = v47
                }
            }
            for fi48 := 0; fi48 < 4; fi48++ {
                v22 = append(v22, -810)
                _ = fi48
            }
            _ = fi24
        }
        if (-1 <= 786) {
            v3 -= 3.14159
            var v49 int32 = 42
            v2 = v2
            _ = v49
        } else {
            fmt.Println(v2, v2, v2)
        }
    }
    switch "a" {
    case "\n\t":
        v50 := v4[(v1 + ("hello" + (("a" + ((((v1 + v1) + v1) + "foo bar") + ("a" + "\n\t"))) + ("foo bar" + ("foo bar" + ((v1 + v1) + v1))))))]
        _ = v50
    }
    var v51 int16 = -59
    _ = len(v4)
    if len(v22) > 0 {
        v22[0] = v2
    }
    v51 = v51
    v4[((v1 + v1) + ("foo bar" + (v1 + v1)))] = true
    v52 := GenS6{F0: (-(v51))}
    var v53 float32 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v51:", v51)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v51
    _ = v52
    _ = v53
}
