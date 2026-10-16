package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 string }
type GenS2 struct { F0 uint32; F1 uint64; F2 int64 }
type GenS3 struct { F0 uint; F1 int64; F2 int }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8; F1 uint8 }
type GenS6 struct { F0 uint8; F1 uint16; F2 int }
type GenS7 struct { F0 uint32; F1 int32 }
type GenS8 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{"hello", "foo bar", ("a" + (("hello" + "") + "hello"))}
    var v3 string = ("" + "foo bar")
    var v4 uint32 = 2147483648
    v5 := make(map[string]int)
    var v6 int8 = 121
    var v7 uint64 = 1
    if false {
        if (3.14159 <= 3.14159) {
            v8 := GenS1{F0: (v6 >> ((4294967295) & 63)), F1: 1, F2: (((v3 + (("" + v3) + (v3 + "hello"))) + v3) + (v3 + v3))}
            _ = v8
        }
    }
    if false {
        v9 := GenS1{F0: v6, F1: 1, F2: v3}
        _ = v9
    } else {
        for fi10 := 0; fi10 < 7; fi10++ {
            for fi11 := 0; fi11 < 1; fi11++ {
                switch v3 {
                case "":
                    _ = len(v2)
                    v12 := GenS1{F0: 109, F1: 58, F2: ("a" + v3)}
                    switch (("foo bar" + v3) + (v3 + ("\n\t" + v3))) {
                    case "hello":
                        var v13 int = 4294967295
                        _ = v13
                    case "foo bar":
                        if (0 > uint16(v7)) {
                            switch "foo bar" {
                            case "foo bar":
                                v14 := v5[("a" + "a")]
                                _ = v14
                            case "\n\t":
                                v6 -= (v6 ^ (v6 & v6))
                            case "hello":
                                v4 -= 954
                            }
                            v7 += (v7 % (((v7 << ((42) & 63))) | 1))
                        } else {
                            var v15 int8 = (v6 / ((42) | 1))
                            v4 = 358
                            _ = v15
                        }
                    case "\n\t":
                        switch v3 {
                        case "hello":
                            switch "\n\t" {
                            case "\n\t":
                                delete(v5, v3)
                            case "":
                                fmt.Println(v4, v6, v6)
                            }
                        case "":
                            if (!((!((v6 == v6) && true)) && true)) {
                                v4 += v4
                            } else {
                                var v16 uint16 = 1
                                _ = len(v2)
                                v3 = "hello"
                                v6 -= 1
                                _ = v16
                            }
                        default:
                            if false {
                                v17 := GenS1{F0: 0, F1: 2, F2: (v3 + v3)}
                                if (!(2 != 127)) {
                                    v18 := GenS1{F0: v6, F1: 42, F2: "a"}
                                    v19 := GenS2{F0: (v4 / (((v4 << ((2) & 63))) | 1)), F1: 2, F2: -9223372036854775808}
                                    _ = v18
                                    _ = v19
                                }
                                _ = v17
                            }
                        }
                    }
                    _ = v12
                case "hello":
                    v20 := GenS3{F0: 0, F1: -654, F2: fi10}
                    _ = v20
                default:
                    v7 = v7
                    switch fi10 {
                    case -631:
                        fmt.Println(v6, fi11, fi11)
                    default:
                        switch fi10 {
                        case 9223372036854775807:
                            if ((true || ((fi10 * 0) != fi11)) && true) {
                                switch (fi10 & (fi11 | fi11)) {
                                case -462:
                                    var v21 uint64 = v7
                                    _ = v21
                                case 32768:
                                    v3 = v3
                                case -991:
                                    if true {
                                        v3 = ("a" + "\n\t")
                                    }
                                default:
                                    switch fi11 {
                                    case -544:
                                        v22 := GenS4{F0: v3}
                                        _ = v22
                                    }
                                }
                                if (4294967295 <= fi10) {
                                    if ((18446744073709551615 <= 4294967295) || true) {
                                        var v23 float32 = ((1e10 / (0.1 + 0.1)) / 0.1)
                                        _ = v23
                                    } else {
                                        if false {
                                            fmt.Println(v6, v7, v7)
                                            v24 := GenS1{F0: (-(v6)), F1: 2, F2: v3}
                                            if false {
                                                v25 := GenS2{F0: v4, F1: (v7 << ((uint(v7)) & 63)), F2: 42}
                                                if false {
                                                    v6 = 89
                                                    var v26 float64 = 1.5
                                                    _ = v26
                                                } else {
                                                    v27 := GenS3{F0: 270, F1: 9223372036854775807, F2: fi10}
                                                    v28 := GenS1{F0: (v6 << ((65535) & 63)), F1: 90, F2: v3}
                                                    var v29 int32 = -138
                                                    _ = v27
                                                    _ = v28
                                                    _ = v29
                                                }
                                                _ = v25
                                            } else {
                                                var v30 uint64 = (v7 + 313)
                                                v4 = v4
                                                _ = v30
                                            }
                                            _ = v24
                                        }
                                        var v31 int32 = 2147483647
                                        _ = v31
                                    }
                                    var v32 int16 = 1000
                                    if true {
                                        if true {
                                            var v33 string = (v3 + v3)
                                            if ((true && (!true)) && (950 != -1)) {
                                                var v34 int16 = (v32 + (^(v32)))
                                                _ = v34
                                            }
                                            v2 = append(v2, "")
                                            _ = v33
                                        }
                                    }
                                    _ = v32
                                } else {
                                    v4 += (v4 & (^(v4)))
                                    var v35 float32 = 1e10
                                    _ = v35
                                }
                            } else {
                                fmt.Println(v7, v6)
                            }
                        }
                        if true {
                            var v36 int = 0
                            _ = v36
                        } else {
                            _ = len(v5)
                        }
                        var v37 string = ((v3 + (v3 + ("a" + (v3 + "foo bar")))) + ("" + (("hello" + v3) + ((v3 + v3) + "foo bar"))))
                        _ = v37
                    }
                    var v38 string = v3
                    _ = v38
                }
                _ = fi11
            }
            _ = fi10
        }
        if false {
            if true {
                v39 := GenS5{F0: 83, F1: 0}
                for fi40 := 0; fi40 < 14; fi40++ {
                    if false {
                        if true {
                            var v41 int = 9223372036854775807
                            var v42 int64 = -169
                            _ = v41
                            _ = v42
                        } else {
                            v3 = "hello"
                            fmt.Println(v7)
                            fmt.Println(fi40, v3)
                        }
                        v43 := GenS6{F0: 62, F1: 65535, F2: 0}
                        _ = v43
                    }
                    _ = fi40
                }
                _ = v39
            }
        } else {
            if len(v2) > 0 {
                v2[0] = v3
            }
            v6 -= v6
        }
        v44 := GenS4{F0: v3}
        _ = v44
    }
    fmt.Println(v3)
    v45 := GenS3{F0: uint(v7), F1: 42, F2: 42}
    if ((1 > 32767) && (32767 < 41)) {
        fmt.Println(v4)
        var v46 uint8 = 199
        if true {
            v6 = (v6 - 127)
        } else {
            for fi47 := 0; fi47 < 16; fi47++ {
                v3 = (("\n\t" + (v3 + (("\n\t" + v3) + ""))) + v3)
                v48 := GenS7{F0: 581, F1: 1}
                _ = fi47
                _ = v48
            }
            if false {
                v49 := GenS4{F0: v3}
                var v50 float64 = 1.5
                _ = v49
                _ = v50
            } else {
                v3 = v3
                var v51 uint64 = (v7 | 168)
                _ = v51
            }
            switch -761 {
            case 9223372036854775807:
                v52 := GenS8{F0: -2147483648}
                _ = v52
            }
        }
        _ = v46
    }
    for fi53 := 0; fi53 < 9; fi53++ {
        v4 += (v4 & (v4 / (((v4 % (((v4 - 4294967295)) | 1))) | 1)))
        _ = fi53
    }
    if (1 != -32768) {
        var v54 int8 = (v6 * 0)
        _ = v54
    }
    var v55 int16 = -1
    v3 = (v3 + v3)
    for fi56 := 0; fi56 < 12; fi56++ {
        v4 += (v4 / ((2147483648) | 1))
        _ = fi56
    }
    v4 += (v4 * 65535)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v55
    _ = v6
    _ = v7
}
