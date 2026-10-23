package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int8; F1 int; F2 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint; F1 int }
type GenS6 struct { F0 int8; F1 uint16 }
type GenS7 struct { F0 int32 }

func main() {
    var v1 int32 = 0
    var v2 uint = 303
    v3 := make(map[int]uint32)
    var v4 int = 2147483647
    for fi5 := 0; fi5 < 16; fi5++ {
        if (true || (1 != 0)) {
            var v6 float64 = 3.14159
            v6 -= -2.25
            v7 := []int{v4}
            _ = v6
            _ = v7
        } else {
            switch fi5 {
            case 32767:
                v8 := GenS1{F0: -588, F1: 705}
                _ = v8
            case 964:
                if (1 >= 0) {
                    fmt.Println(v4, v2, fi5)
                    v4 -= fi5
                    if (291 == (v2 | (^(v2)))) {
                        v9 := GenS1{F0: -9223372036854775808, F1: -9223372036854775808}
                        var v10 bool = (v1 <= (v1 % ((v1) | 1)))
                        switch v4 {
                        case -140:
                            switch ("hello" + "foo bar") {
                            case "":
                                v11 := GenS1{F0: -1, F1: 193}
                                _ = v11
                            case "hello":
                                var v12 uint16 = 32767
                                _ = v12
                            case "\n\t":
                                var v13 float32 = 1e10
                                _ = v13
                            default:
                                var v14 float32 = 1.5
                                _ = v14
                            }
                        case 131:
                            if ((v1 / ((v1) | 1)) > (^(v1))) {
                                switch (v4 & fi5) {
                                case -9223372036854775808:
                                    v15 := GenS2{F0: false, F1: "foo bar"}
                                    _ = v15
                                default:
                                    for fi16 := 0; fi16 < 11; fi16++ {
                                        var v17 string = "a"
                                        switch "" {
                                        case "hello":
                                            fmt.Println(v2, v2)
                                        case "\n\t":
                                            var v18 uint32 = 2
                                            _ = v18
                                        default:
                                            v2 += (v2 & (v2 * v2))
                                        }
                                        _ = fi16
                                        _ = v17
                                    }
                                }
                            } else {
                                var v19 int = fi5
                                if (v10 || false) {
                                    var v20 float32 = (1.5 / float32((3.14159 + 0.0)))
                                    v21 := []int{(v4 << (((^(v2))) & 63)), (v19 & (v4 | v4)), (^(v19)), (v19 + (v4 + (v4 - (v4 % ((v4) | 1)))))}
                                    _ = v20
                                    _ = v21
                                } else {
                                    if true {
                                        v22 := []int{(fi5 % (((fi5 & (fi5 + v4))) | 1)), v4, (fi5 % (((v19 + 32768)) | 1))}
                                        v1 = v1
                                        _ = v22
                                    } else {
                                        var v23 uint32 = 264
                                        _ = v23
                                    }
                                    for fi24 := 0; fi24 < 5; fi24++ {
                                        if ("" < "foo bar") {
                                            var v25 float32 = -2.25
                                            if (!v10) {
                                                var v26 int8 = -90
                                                switch ("a" + "\n\t") {
                                                case "foo bar":
                                                    v27 := GenS3{F0: (v26 / (((v26 >> ((v2) & 63))) | 1)), F1: (^(fi24)), F2: ""}
                                                    _ = v27
                                                case "a":
                                                    v10 = (!v10)
                                                case "":
                                                    if (-1 == 9223372036854775807) {
                                                        var v28 int = (fi24 / ((v19) | 1))
                                                        v19 = fi24
                                                        v29 := GenS4{F0: (v26 / (((v26 % ((v26) | 1))) | 1))}
                                                        v26 = (v26 ^ v26)
                                                        _ = v28
                                                        _ = v29
                                                    }
                                                default:
                                                    switch int(v1) {
                                                    case 707:
                                                        v19 = v4
                                                    case 9223372036854775807:
                                                        fmt.Println(v19)
                                                    }
                                                }
                                                if (!false) {
                                                    v30 := v3[(fi5 * v19)]
                                                    v31 := GenS1{F0: 1, F1: -961}
                                                    _ = v30
                                                    _ = v31
                                                } else {
                                                    v32 := GenS5{F0: 239, F1: -9223372036854775808}
                                                    v3[(v19 | (v19 + (-(fi5))))] = 65535
                                                    v33 := GenS6{F0: v26, F1: 956}
                                                    _ = v32
                                                    _ = v33
                                                }
                                                var v34 uint = 549
                                                switch "" {
                                                case "foo bar":
                                                    var v35 int16 = 32767
                                                    _ = v35
                                                case "a":
                                                    var v36 uint16 = 255
                                                    _ = v36
                                                case "\n\t":
                                                    if (true || (!v10)) {
                                                        v37 := GenS3{F0: v26, F1: (-(v19)), F2: "a"}
                                                        _ = v37
                                                    } else {
                                                        v19 += 694
                                                    }
                                                }
                                                _ = v26
                                                _ = v34
                                            }
                                            fmt.Println(fi24, v10)
                                            _ = v25
                                        }
                                        v19 -= -2147483648
                                        switch "\n\t" {
                                        case "foo bar":
                                            v10 = false
                                        case "a":
                                            switch "a" {
                                            case "\n\t":
                                                v38 := []int{v19}
                                                _ = v38
                                            case "hello":
                                                v19 += 197
                                            }
                                        default:
                                            if v10 {
                                                v39 := GenS2{F0: v10, F1: ("\n\t" + "a")}
                                                v40 := GenS6{F0: -15, F1: 32767}
                                                _ = v39
                                                _ = v40
                                            }
                                        }
                                        _ = fi24
                                    }
                                }
                                _ = v19
                            }
                        case -9223372036854775808:
                            v2 = (v2 | (v2 / ((v2) | 1)))
                        }
                        _ = v10
                        _ = v9
                    }
                }
            case 128:
                for fi41 := 0; fi41 < 13; fi41++ {
                    var v42 uint32 = 2147483647
                    v43 := []int{v4}
                    _ = fi41
                    _ = v42
                    _ = v43
                }
            default:
                fmt.Println(v1)
            }
            v44 := []int{497, v4, (v4 / ((v4) | 1))}
            _ = v44
        }
        fmt.Println(v1)
        v2 = 42
        _ = fi5
    }
    for fi45 := 0; fi45 < 19; fi45++ {
        var v46 float32 = 1e10
        _ = fi45
        _ = v46
    }
    v1 = (v1 % (((-(v1))) | 1))
    if (18446744073709551615 != (v2 << (((v2 & 2)) & 63))) {
        v4 = (v4 - (v4 - v4))
    } else {
        if false {
            v47 := v3[(-(v4))]
            v1 = v1
            _ = v47
        } else {
            fmt.Println(v2)
        }
        fmt.Println(v2, v2, v4)
        var v48 int64 = -1
        _ = v48
    }
    switch "\n\t" {
    case "foo bar":
        switch "hello" {
        case "":
            v2 = v2
        case "hello":
            var v49 float32 = 0.0
            _ = v49
        case "a":
            if true {
                v2 -= v2
            }
        default:
            v4 -= v4
        }
    case "":
        _ = len(v3)
    case "hello":
        var v50 bool = false
        _ = v50
    }
    v2 = (v2 % ((v2) | 1))
    var v51 int16 = 32767
    if false {
        v52 := GenS7{F0: (v1 + (-(v1)))}
        var v53 uint32 = 2147483648
        _ = v52
        _ = v53
    }
    v1 = (v1 % ((0) | 1))
    v1 += 1
    fmt.Println(v2, v1, v51)
    fmt.Println(v2)
    var v54 float32 = (float32(1e10) * (0.0 - 1.5))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v51:", v51)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v51
    _ = v54
}
