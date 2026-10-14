package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int32 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int64; F1 int32; F2 int64 }
type GenS4 struct { F0 bool; F1 int8; F2 bool }
type GenS5 struct { F0 int64; F1 uint8 }

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 2
    v3 := []int{}
    v4 := []int{}
    var v5 bool = false
    if v5 {
        var v6 bool = v5
        v4 = append(v4, 793)
        v7 := GenS1{F0: 0, F1: 32767, F2: 42}
        _ = v6
        _ = v7
    } else {
        v8 := GenS1{F0: 779, F1: 32767, F2: 576}
        var v9 uint16 = 32767
        _ = v8
        _ = v9
    }
    v2 += v2
    var v10 int16 = 100
    fmt.Println(v2)
    var v11 int = 48
    v4 = append(v4, 2147483647)
    for fi12 := 0; fi12 < 13; fi12++ {
        v5 = (965 != v11)
        v13 := GenS2{F0: 42, F1: (fi12 ^ 722)}
        _ = fi12
        _ = v13
    }
    switch ("" + ("\n\t" + (("foo bar" + (((("hello" + "\n\t") + "a") + ("" + "a")) + "\n\t")) + ""))) {
    case "":
        fmt.Println(v11, v11)
    case "\n\t":
        for fi14 := 0; fi14 < 19; fi14++ {
            switch "foo bar" {
            case "hello":
                if true {
                    v11 = 806
                    var v15 uint64 = 18446744073709551615
                    if ((!(255 > 487)) || true) {
                        v16 := GenS3{F0: 1, F1: 499, F2: int64(v2)}
                        v17 := GenS1{F0: v10, F1: v10, F2: -2147483648}
                        switch (fi14 ^ 32768) {
                        case 1:
                            v2 = v2
                        }
                        _ = v16
                        _ = v17
                    } else {
                        v3 = append(v3, (fi14 - (-(fi14))))
                        if false {
                            switch (v11 << ((42) & 63)) {
                            case 127:
                                v3 = append(v3, (v11 + 718))
                            case 470:
                                v10 = (v10 - v10)
                            case 128:
                                for fi18 := 0; fi18 < 11; fi18++ {
                                    v19 := GenS3{F0: -9223372036854775808, F1: -2147483648, F2: 1}
                                    if (9223372036854775807 < 9223372036854775807) {
                                        delete(v1, "a")
                                        var v20 uint8 = 40
                                        _ = v20
                                    } else {
                                        _ = len(v1)
                                        if false {
                                            var v21 int16 = v10
                                            var v22 float32 = 1e10
                                            _ = v21
                                            _ = v22
                                        } else {
                                            v23 := GenS1{F0: (v10 % (((v10 / (((-(v10))) | 1))) | 1)), F1: (v10 * (v10 / ((v10) | 1))), F2: -1}
                                            v5 = v5
                                            if len(v3) > 0 {
                                                v3[0] = (-(fi18))
                                            }
                                            v5 = (-2147483648 < -674)
                                            _ = v23
                                        }
                                        if ((fi18 << ((uint(v2)) & 63)) < -744) {
                                            var v24 float32 = 1.5
                                            _ = v24
                                        }
                                        if (!((true && (false || true)) && false)) {
                                            if v5 {
                                                if (false && v5) {
                                                    var v25 bool = (1 != 232)
                                                    v11 = (fi18 & v11)
                                                    v10 = 42
                                                    _ = v25
                                                }
                                                _ = len(v1)
                                                v1[""] = "foo bar"
                                                v26 := v1["\n\t"]
                                                _ = v26
                                            } else {
                                                var v27 float64 = (-2.25 + (-2.25 + 1e10))
                                                v15 -= 18446744073709551615
                                                v10 = v10
                                                _ = v27
                                            }
                                        } else {
                                            v11 += (v11 % ((-304) | 1))
                                            v10 = (v10 & (v10 << ((596) & 63)))
                                        }
                                        switch ("hello" + ("a" + "hello")) {
                                        case "\n\t":
                                            var v28 uint32 = v2
                                            _ = v28
                                        case "":
                                            var v29 int = (fi14 * fi14)
                                            _ = v29
                                        case "a":
                                            if (!v5) {
                                                v10 = v10
                                                if v5 {
                                                    v5 = ("foo bar" > ("a" + ("\n\t" + "foo bar")))
                                                    fmt.Println(fi18)
                                                } else {
                                                    var v30 bool = ((v2 - (v2 << ((2) & 63))) < (v2 - (v2 % ((v2) | 1))))
                                                    _ = v30
                                                }
                                            } else {
                                                switch ((("" + "") + ("a" + (("hello" + "hello") + "foo bar"))) + "hello") {
                                                case "foo bar":
                                                    v31 := GenS4{F0: (2147483647 <= -2147483648), F1: 98, F2: (80 > 255)}
                                                    _ = v31
                                                }
                                                var v32 uint64 = 0
                                                _ = v32
                                            }
                                        }
                                        v11 = 32767
                                    }
                                    _ = fi18
                                    _ = v19
                                }
                            default:
                                for fi33 := 0; fi33 < 12; fi33++ {
                                    if ((^(v2)) == (v2 / ((65535) | 1))) {
                                        var v34 uint64 = 42
                                        _ = v34
                                    } else {
                                        var v35 float32 = 3.14159
                                        _ = v35
                                    }
                                    _ = fi33
                                }
                            }
                            v36 := GenS3{F0: 9223372036854775807, F1: 0, F2: 9223372036854775807}
                            _ = v36
                        } else {
                            v11 -= v11
                        }
                    }
                    _ = v15
                }
            case "\n\t":
                if v5 {
                    var v37 uint64 = 0
                    _ = len(v4)
                    switch fi14 {
                    case -9223372036854775808:
                        delete(v1, ("\n\t" + (("" + "a") + "\n\t")))
                    }
                    _ = v37
                }
            default:
                switch "foo bar" {
                case "\n\t":
                    fmt.Println(v5, v11, v10)
                case "":
                    if (1e10 == 0.1) {
                        var v38 uint8 = 0
                        _ = v38
                    }
                default:
                    v1["foo bar"] = ("a" + "a")
                }
            }
            switch (("" + (("hello" + "hello") + ("hello" + ""))) + "foo bar") {
            case "":
                var v39 int = (fi14 >> ((42) & 63))
                _ = v39
            case "foo bar":
                var v40 int32 = -121
                _ = v40
            }
            if ((v10 << ((18446744073709551615) & 63)) >= 1) {
                for fi41 := 0; fi41 < 5; fi41++ {
                    fmt.Println(fi41, v11, v5)
                    if v5 {
                        fmt.Println(v11, v11)
                        v10 -= 40
                        var v42 int32 = 42
                        _ = v42
                    }
                    _ = fi41
                }
            }
            _ = fi14
        }
    default:
        switch "a" {
        case "hello":
            for fi43 := 0; fi43 < 12; fi43++ {
                var v44 uint32 = (v2 & v2)
                v10 = (^(v10))
                for fi45 := 0; fi45 < 17; fi45++ {
                    var v46 float64 = 0.0
                    fmt.Println(v46, v11)
                    v1["hello"] = "hello"
                    _ = fi45
                    _ = v46
                }
                _ = fi43
                _ = v44
            }
        case "foo bar":
            v2 -= v2
        case "\n\t":
            if (!((v10 + v10) >= v10)) {
                v10 = -32768
                var v47 int64 = 71
                switch "a" {
                case "a":
                    var v48 bool = (!false)
                    _ = v48
                case "":
                    if true {
                        v49 := v1["a"]
                        for fi50 := 0; fi50 < 4; fi50++ {
                            var v51 int = v11
                            _ = fi50
                            _ = v51
                        }
                        var v52 int = v11
                        _ = v49
                        _ = v52
                    } else {
                        v11 -= v11
                    }
                default:
                    if ("a" >= "") {
                        v53 := GenS5{F0: (v47 >> ((390) & 63)), F1: 128}
                        var v54 uint16 = 310
                        var v55 int16 = v10
                        _ = v53
                        _ = v54
                        _ = v55
                    }
                }
                _ = v47
            } else {
                for fi56 := 0; fi56 < 7; fi56++ {
                    var v57 uint8 = 255
                    var v58 int64 = 9223372036854775807
                    _ = fi56
                    _ = v57
                    _ = v58
                }
                var v59 int16 = 1
                _ = v59
            }
        default:
            v60 := GenS5{F0: 0, F1: 192}
            _ = v60
        }
    }
    v11 -= (-(v11))
    if (!v5) {
        switch "a" {
        case "\n\t":
            v61 := v1["foo bar"]
            _ = v61
        }
        var v62 uint64 = 18446744073709551615
        _ = v62
    }
    v5 = v5
    fmt.Println(v5, v10, v5)
    var v63 int = v11
    if v5 {
        switch ("a" + "a") {
        case "":
            v10 = v10
        }
    } else {
        _ = len(v3)
        fmt.Println(v11, v63)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v63
}
