package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint32; F1 uint; F2 string }
type GenS3 struct { F0 bool; F1 int16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string; F1 int; F2 int64 }
type GenS6 struct { F0 string; F1 int16 }
type GenS7 struct { F0 int; F1 bool }
func genFn0(p0 int, p1 int8, p2 uint16) string {
    fmt.Println(p1, p0, p2)
    var v6 uint64 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return ("foo bar" + "a")
}
func genFn1() uint32 {
    for fi7 := 0; fi7 < 2; fi7++ {
        v8 := GenS1{F0: 385, F1: 42}
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        fmt.Println(fi9)
        var v10 int = (fi9 % ((fi9) | 1))
        _ = fi9
        _ = v10
    }
    var v11 string = "\n\t"
    switch "\n\t" {
    case "":
        switch "" {
        case "a":
            v11 = "foo bar"
        default:
            v12 := []int{1, -2147483648, 2147483647, -2147483648}
            _ = v12
        }
    case "\n\t":
        v13 := []int{}
        _ = v13
    }
    var v14 int8 = -1
    _ = v11
    _ = v14
    return 2147483648
}

func main() {
    var v1 int8 = -68
    var v2 uint32 = 4294967295
    var v3 int = 32768
    v4 := []string{("foo bar" + "a")}
    v5 := []uint32{}
    fmt.Println(v3, v3, v2)
    var v15 float32 = 0.1
    switch ("foo bar" + "hello") {
    case "":
        switch "foo bar" {
        case "\n\t":
            for fi16 := 0; fi16 < 9; fi16++ {
                _ = len(v5)
                var v17 bool = ((!(true && false)) && (0.1 > 3.14159))
                _ = fi16
                _ = v17
            }
        case "":
            var v18 bool = (!false)
            _ = v18
        case "a":
            v19 := GenS2{F0: genFn1(), F1: 18446744073709551615, F2: (("a" + "a") + ("a" + "hello"))}
            _ = v19
        default:
            if (!false) {
                if true {
                    var v20 int16 = -1
                    if true {
                        for fi21 := 0; fi21 < 17; fi21++ {
                            var v22 int32 = -2147483648
                            _ = fi21
                            _ = v22
                        }
                    }
                    for fi23 := 0; fi23 < 7; fi23++ {
                        v24 := GenS3{F0: (839 == 763), F1: v20}
                        v2 -= v2
                        for fi25 := 0; fi25 < 18; fi25++ {
                            if true {
                                var v26 int8 = 0
                                switch "" {
                                case "\n\t":
                                    v27 := GenS2{F0: (v2 * genFn1()), F1: 2, F2: ""}
                                    _ = v27
                                default:
                                    v28 := GenS2{F0: genFn1(), F1: 1, F2: (("hello" + "a") + ("a" + ""))}
                                    _ = v28
                                }
                                _ = v26
                            } else {
                                v4 = append(v4, "\n\t")
                                var v29 int8 = (v1 ^ v1)
                                v2 -= (v2 | genFn1())
                                _ = v29
                            }
                            v30 := GenS4{F0: (v1 & v1)}
                            _ = fi25
                            _ = v30
                        }
                        _ = fi23
                        _ = v24
                    }
                    _ = v20
                } else {
                    var v31 bool = (genFn1() == genFn1())
                    if len(v4) > 0 {
                        v4[0] = ""
                    }
                    _ = v31
                }
            }
        }
    case "foo bar":
        _ = len(v4)
    case "\n\t":
        v15 = v15
    }
    v32 := GenS5{F0: "", F1: (^(v3)), F2: 42}
    for fi33 := 0; fi33 < 1; fi33++ {
        for fi34 := 0; fi34 < 18; fi34++ {
            if true {
                var v35 int = fi34
                v15 = v15
                _ = v35
            } else {
                v1 = 115
            }
            fmt.Println(v15)
            v15 -= v15
            var v36 uint32 = (v2 & v2)
            _ = fi34
            _ = v36
        }
        for fi37 := 0; fi37 < 4; fi37++ {
            v3 += fi37
            if len(v4) > 0 {
                v4[0] = "a"
            }
            var v38 float64 = 3.14159
            _ = fi37
            _ = v38
        }
        _ = fi33
    }
    for fi39 := 0; fi39 < 1; fi39++ {
        v15 -= v15
        if true {
            var v40 uint64 = 18446744073709551615
            _ = v40
        }
        _ = fi39
    }
    switch 198 {
    case 0:
        _ = len(v4)
    case 2147483647:
        if true {
            v41 := GenS6{F0: "hello", F1: -836}
            var v42 bool = false
            _ = v41
            _ = v42
        }
    }
    v3 = 0
    var v43 bool = true
    if (-2147483648 >= -2147483648) {
        var v44 int64 = 0
        if v43 {
            for fi45 := 0; fi45 < 1; fi45++ {
                var v46 int16 = -32768
                _ = fi45
                _ = v46
            }
            var v47 int64 = (v44 & v44)
            _ = v47
        }
        fmt.Println(v44, v15)
        _ = v44
    } else {
        if (868 != 4294967295) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        } else {
            switch "" {
            case "hello":
                v48 := GenS2{F0: (v2 | (v2 >> ((0) & 63))), F1: 1, F2: ""}
                _ = v48
            case "\n\t":
                var v49 uint64 = 1
                _ = v49
            case "a":
                v2 = genFn1()
            default:
                var v50 uint64 = 0
                _ = v50
            }
        }
    }
    v51 := GenS2{F0: (v2 + genFn1()), F1: 4294967295, F2: ("" + "\n\t")}
    var v52 uint64 = 42
    var v53 int = (v3 / (((v3 ^ v3)) | 1))
    switch "foo bar" {
    case "":
        switch "foo bar" {
        case "a":
            v53 = v53
        case "hello":
            var v54 uint8 = 0
            _ = v54
        case "foo bar":
            for fi55 := 0; fi55 < 20; fi55++ {
                if (("" + "") >= "\n\t") {
                    var v56 int64 = 338
                    _ = v56
                }
                switch 32767 {
                case 0:
                    v57 := GenS7{F0: (v53 / (((v53 << ((0) & 63))) | 1)), F1: (32767 != 133)}
                    _ = v57
                case -997:
                    v58 := GenS5{F0: (("" + "hello") + ("a" + "hello")), F1: 176, F2: 9223372036854775807}
                    _ = v58
                case 32768:
                    v52 -= v52
                }
                var v59 uint8 = 128
                _ = fi55
                _ = v59
            }
        }
    case "a":
        var v60 int32 = 0
        _ = v60
    default:
        _ = len(v4)
    }
    v52 = (v52 / (((v52 | v52)) | 1))
    if v43 {
        fmt.Println(v53)
        fmt.Println(v15)
        for fi61 := 0; fi61 < 14; fi61++ {
            switch fi61 {
            case 128:
                var v62 uint32 = genFn1()
                _ = v62
            default:
                v2 -= v2
            }
            _ = fi61
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v43
    _ = v5
    _ = v51
    _ = v52
    _ = v53
}
