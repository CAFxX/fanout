package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 int32; F1 int16 }
type GenS6 struct { F0 int8; F1 uint8; F2 string }
type GenS7 struct { F0 int64; F1 int; F2 uint8 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint32; F1 bool; F2 uint8 }

func main() {
    var v1 int = 76
    v2 := make(map[string]uint32)
    v3 := []string{"foo bar", "a", "\n\t"}
    var v4 int8 = -1
    if false {
        v1 = v1
        if (((((("a" + "hello") + "hello") + "") + ("a" + "")) + "foo bar") <= ("" + "")) {
            switch ((("foo bar" + "a") + (((((((("" + "foo bar") + "hello") + "") + ("" + "a")) + "hello") + "") + ("a" + (("" + "hello") + ""))) + "hello")) + ("a" + "hello")) {
            case "":
                fmt.Println(v4, v4)
            case "hello":
                var v5 int32 = 0
                _ = v5
            }
        } else {
            var v6 float64 = 0.1
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if (950 > 2) {
                v7 := GenS1{F0: 2}
                var v8 float64 = (float64(0.0) * v6)
                v9 := v2["a"]
                _ = v7
                _ = v8
                _ = v9
            }
            _ = v6
        }
    }
    v1 = v1
    var v10 int = 9223372036854775807
    for fi11 := 0; fi11 < 19; fi11++ {
        for fi12 := 0; fi12 < 1; fi12++ {
            if (1 == 855) {
                if true {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                if true {
                    v13 := GenS1{F0: 0}
                    var v14 int32 = 42
                    var v15 int8 = (v4 % (((v4 / (((-(v4))) | 1))) | 1))
                    _ = v13
                    _ = v14
                    _ = v15
                }
                if (986 != 1) {
                    var v16 bool = (324 >= 717)
                    var v17 int16 = 42
                    _ = v16
                    _ = v17
                }
                delete(v2, "")
                v4 += -91
                if (0.0 == (0.1 + 1e10)) {
                    _ = len(v2)
                    if (2147483647 == -601) {
                        switch "foo bar" {
                        case "hello":
                            if true {
                                var v18 int8 = 0
                                if ((false || false) || (("foo bar" + "") < "foo bar")) {
                                    fmt.Println(fi11, fi12)
                                    if (42 < 301) {
                                        v19 := GenS1{F0: 1}
                                        var v20 int64 = -221
                                        _ = v19
                                        _ = v20
                                    } else {
                                        v1 = -9223372036854775808
                                        v4 = 100
                                    }
                                }
                                _ = len(v2)
                                _ = v18
                            }
                        default:
                            switch (fi11 / ((127) | 1)) {
                            case 9223372036854775807:
                                var v21 uint16 = 427
                                _ = v21
                            default:
                                v22 := GenS2{F0: 782, F1: -1}
                                _ = v22
                            }
                        }
                        var v23 int64 = 9223372036854775807
                        _ = v23
                    }
                    v24 := GenS1{F0: 0}
                    _ = v24
                } else {
                    var v25 uint8 = 203
                    if true {
                        v2[("" + (("\n\t" + "hello") + ""))] = 157
                        var v26 int32 = -519
                        v27 := GenS2{F0: 2, F1: (v26 | v26)}
                        if false {
                            switch ("a" + "a") {
                            case "":
                                v26 = v26
                            case "hello":
                                var v28 uint64 = 4294967295
                                _ = v28
                            case "a":
                                fmt.Println(fi12, v4)
                            default:
                                v29 := GenS1{F0: 42}
                                _ = v29
                            }
                            var v30 uint64 = 4294967295
                            switch (fi11 >> ((18446744073709551615) & 63)) {
                            case 1:
                                v31 := GenS3{F0: v25, F1: (^(v30)), F2: 0}
                                _ = v31
                            }
                            var v32 int8 = (^(v4))
                            v33 := GenS2{F0: 2, F1: 0}
                            v4 = 0
                            v1 = (-(fi12))
                            var v34 uint64 = (v30 * v30)
                            _ = v30
                            _ = v32
                            _ = v33
                            _ = v34
                        } else {
                            var v35 int32 = (v26 * -165)
                            fmt.Println(v35, v25)
                            var v36 int16 = 1000
                            _ = v35
                            _ = v36
                        }
                        if true {
                            var v37 int8 = v4
                            v4 = (v37 & (v4 / (((v37 | int8(v25))) | 1)))
                            _ = len(v2)
                            _ = v37
                        } else {
                            v38 := GenS3{F0: 12, F1: 2, F2: -32768}
                            v39 := GenS3{F0: (v25 ^ (v25 * v25)), F1: 18446744073709551615, F2: 463}
                            var v40 uint16 = 0
                            _ = v38
                            _ = v39
                            _ = v40
                        }
                        if (fi11 != 528) {
                            fmt.Println(fi11)
                            var v41 int16 = -850
                            v42 := GenS4{F0: (fi12 + v10), F1: 127}
                            _ = v41
                            _ = v42
                        }
                        _ = v26
                        _ = v27
                    }
                    switch "\n\t" {
                    case "foo bar":
                        var v43 uint16 = 32767
                        _ = v43
                    case "a":
                        _ = len(v3)
                    case "\n\t":
                        _ = len(v2)
                    default:
                        v44 := GenS5{F0: -745, F1: int16(v4)}
                        _ = v44
                    }
                    _ = v25
                }
            } else {
                if true {
                    var v45 int64 = 0
                    var v46 bool = ((^(v45)) == (v45 % ((615) | 1)))
                    delete(v2, ("foo bar" + ("hello" + "")))
                    _ = v45
                    _ = v46
                }
            }
            var v47 int32 = -785
            v10 -= (fi12 + (v10 & 32767))
            switch v10 {
            case 9223372036854775807:
                switch int(v4) {
                case 32768:
                    v48 := GenS6{F0: (v4 * (v4 + (v4 << ((18446744073709551615) & 63)))), F1: 91, F2: ""}
                    _ = v48
                default:
                    v10 -= (v1 | -9223372036854775808)
                    var v49 int = v10
                    var v50 uint64 = 0
                    _ = v49
                    _ = v50
                }
            case 32767:
                v51 := GenS7{F0: 0, F1: (fi11 + fi11), F2: 176}
                v47 += v47
                _ = v51
            }
            _ = fi12
            _ = v47
        }
        v52 := GenS3{F0: 150, F1: 4294967295, F2: -320}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi11
        _ = v52
    }
    for fi53 := 0; fi53 < 13; fi53++ {
        var v54 uint8 = 1
        var v55 uint32 = 601
        if false {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        _ = fi53
        _ = v54
        _ = v55
    }
    _ = len(v2)
    v56 := GenS8{F0: ("a" + (("\n\t" + "\n\t") + ("foo bar" + "")))}
    v57 := GenS3{F0: 42, F1: 1, F2: -1}
    v58 := GenS9{F0: 0, F1: (2 == 0), F2: 105}
    for fi59 := 0; fi59 < 16; fi59++ {
        for fi60 := 0; fi60 < 13; fi60++ {
            v4 = v4
            fmt.Println(v1, v4, fi59)
            switch (-(v1)) {
            case -851:
                v10 += 32767
            case -45:
                var v61 int8 = 10
                _ = v61
            default:
                _ = len(v3)
            }
            _ = fi60
        }
        _ = fi59
    }
    v62 := v2["foo bar"]
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v56
    _ = v57
    _ = v58
    _ = v62
}
