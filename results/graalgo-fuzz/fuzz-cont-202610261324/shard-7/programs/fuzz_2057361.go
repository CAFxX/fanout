package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int64 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 uint8; F1 int8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64 }
func genFn0() int32 {
    var v6 string = "\n\t"
    switch 842 {
    case 0:
        fmt.Println(v6)
    case -264:
        var v7 string = (v6 + (((((((((("a" + v6) + (("a" + ("foo bar" + v6)) + (v6 + "foo bar"))) + v6) + "") + ((v6 + ("" + v6)) + v6)) + v6) + v6) + v6) + ((("a" + "") + ("" + "\n\t")) + v6)) + v6))
        _ = v7
    case -2147483648:
        var v8 int32 = 818
        _ = v8
    default:
        for fi9 := 0; fi9 < 9; fi9++ {
            v10 := []int{(-(fi9)), (fi9 >> ((17) & 63))}
            if (32767 > fi9) {
                if ((!false) && false) {
                    for fi11 := 0; fi11 < 7; fi11++ {
                        v12 := GenS1{F0: -32768, F1: -624, F2: 9223372036854775807}
                        fmt.Println(fi9, fi11, fi11)
                        v10 = append(v10, 4294967295)
                        _ = fi11
                        _ = v12
                    }
                    v6 = v6
                } else {
                    v13 := GenS1{F0: 32767, F1: -111, F2: 1}
                    _ = v13
                }
            } else {
                v14 := GenS2{F0: 1, F1: 117}
                _ = v14
            }
            _ = fi9
            _ = v10
        }
    }
    if (-9223372036854775808 >= -9223372036854775808) {
        v6 = v6
    }
    if (((100 >= 127) || true) && (true || (!false))) {
        if false {
            v6 = v6
        }
    }
    _ = v6
    return 42
}

func main() {
    var v1 uint32 = 4294967295
    var v2 int64 = -1
    var v3 uint = 18446744073709551615
    var v4 int64 = (v2 + 1)
    v5 := make(map[string]string)
    v4 -= v4
    for fi15 := 0; fi15 < 7; fi15++ {
        for fi16 := 0; fi16 < 1; fi16++ {
            if true {
                v1 += (^(v1))
            } else {
                v17 := GenS1{F0: 42, F1: int32(v4), F2: (v2 + v4)}
                _ = v17
            }
            v18 := GenS3{F0: 262}
            _ = fi16
            _ = v18
        }
        _ = fi15
    }
    v19 := []int{-774}
    var v20 int32 = genFn0()
    v20 = (v20 / ((v20) | 1))
    v3 = (v3 / (((v3 / ((1000) | 1))) | 1))
    for fi21 := 0; fi21 < 14; fi21++ {
        switch fi21 {
        case -2147483648:
            v1 -= 110
        default:
            for fi22 := 0; fi22 < 13; fi22++ {
                v23 := GenS1{F0: 0, F1: genFn0(), F2: v2}
                _ = len(v19)
                _ = fi22
                _ = v23
            }
        }
        switch fi21 {
        case -905:
            for fi24 := 0; fi24 < 18; fi24++ {
                _ = len(v5)
                if (!true) {
                    v25 := GenS4{F0: fi24, F1: 1}
                    if (false || false) {
                        fmt.Println(v2, v4)
                        v4 -= v2
                    }
                    _ = v25
                } else {
                    if false {
                        fmt.Println(v2)
                        if true {
                            var v26 uint16 = 1
                            _ = v26
                        }
                    } else {
                        if false {
                            if false {
                                var v27 uint16 = 42
                                v28 := GenS5{F0: 0, F1: int8(fi24)}
                                var v29 uint16 = (v27 & 255)
                                _ = v27
                                _ = v28
                                _ = v29
                            }
                            var v30 int64 = v2
                            _ = v30
                        }
                        v4 += v2
                    }
                }
                _ = fi24
            }
        case -403:
            v5[("hello" + ("foo bar" + (("a" + ("" + "foo bar")) + "foo bar")))] = "foo bar"
        default:
            var v31 uint16 = 0
            _ = v31
        }
        v32 := GenS6{F0: genFn0()}
        _ = fi21
        _ = v32
    }
    v4 = 1
    var v33 uint32 = v1
    for fi34 := 0; fi34 < 16; fi34++ {
        switch "foo bar" {
        case "hello":
            switch "\n\t" {
            case "a":
                if true {
                    var v35 bool = (127 != 0)
                    if false {
                        var v36 uint32 = (v1 + v33)
                        _ = v36
                    } else {
                        if true {
                            _ = len(v5)
                            v1 = (v33 - v33)
                        } else {
                            v1 = (v1 ^ (v1 + (v1 / (((v1 >> ((v3) & 63))) | 1))))
                            if (!v35) {
                                v1 = v1
                                var v37 int8 = 127
                                v38 := GenS6{F0: v20}
                                _ = v37
                                _ = v38
                            } else {
                                var v39 int64 = v4
                                _ = v39
                            }
                        }
                        v20 -= -2147483648
                        var v40 float64 = 1.5
                        _ = v40
                    }
                    var v41 uint16 = 2
                    _ = v35
                    _ = v41
                }
            case "":
                v1 += v33
            case "\n\t":
                delete(v5, "")
            default:
                fmt.Println(v2)
            }
        default:
            switch ("\n\t" + "hello") {
            case "\n\t":
                fmt.Println(v1, v4, fi34)
            case "a":
                switch "hello" {
                case "\n\t":
                    var v42 string = "a"
                    _ = v42
                case "":
                    if (100 < 122) {
                        fmt.Println(v1, fi34, fi34)
                        switch (fi34 & fi34) {
                        case 2147483647:
                            v3 = v3
                        case 9223372036854775807:
                            var v43 int16 = 524
                            _ = v43
                        case 42:
                            fmt.Println(v2, v2, v2)
                        }
                        v1 -= v1
                    } else {
                        if len(v19) >= 2 {
                            v19 = v19[1:len(v19)]
                        }
                    }
                }
            }
        }
        switch fi34 {
        case -2147483648:
            switch int(v3) {
            case 0:
                v1 = 1
            case -357:
                fmt.Println(v3)
            case 42:
                v33 = (^(v33))
            }
        case -9223372036854775808:
            v33 = v33
        case -124:
            v44 := GenS6{F0: genFn0()}
            _ = v44
        default:
            if (true || (319 > 4294967295)) {
                for fi45 := 0; fi45 < 15; fi45++ {
                    v46 := GenS7{F0: 102}
                    _ = fi45
                    _ = v46
                }
                _ = len(v19)
                if (-233 > 881) {
                    var v47 float32 = 0.0
                    if false {
                        var v48 int8 = 0
                        fmt.Println(fi34)
                        _ = v48
                    }
                    _ = v47
                }
            } else {
                if false {
                    for fi49 := 0; fi49 < 20; fi49++ {
                        v50 := GenS1{F0: 42, F1: int32(v33), F2: v2}
                        var v51 uint64 = 2
                        v52 := GenS3{F0: 2}
                        _ = fi49
                        _ = v50
                        _ = v51
                        _ = v52
                    }
                    fmt.Println(v1, v3, v20)
                } else {
                    if false {
                        if (!false) {
                            v19 = append(v19, (fi34 & -623))
                        }
                        var v53 uint8 = 1
                        _ = v53
                    } else {
                        v1 = uint32(v3)
                    }
                    for fi54 := 0; fi54 < 6; fi54++ {
                        v55 := GenS3{F0: 440}
                        v33 = 1
                        var v56 bool = (false || false)
                        _ = fi54
                        _ = v55
                        _ = v56
                    }
                    var v57 uint32 = v33
                    _ = v57
                }
            }
        }
        _ = fi34
    }
    switch ("hello" + "a") {
    case "foo bar":
        var v58 uint16 = 785
        _ = v58
    default:
        fmt.Println(v3, v33)
    }
    for fi59 := 0; fi59 < 12; fi59++ {
        v5["hello"] = ("foo bar" + "a")
        v4 -= 42
        switch ("foo bar" + "hello") {
        case "hello":
            var v60 bool = true
            _ = v60
        case "":
            v33 -= (v33 ^ 1)
        case "a":
            for fi61 := 0; fi61 < 17; fi61++ {
                if true {
                    var v62 int64 = (v2 * v2)
                    v2 += (v62 - (v4 / (((v4 | (v2 * 9223372036854775807))) | 1)))
                    var v63 int8 = 127
                    _ = v62
                    _ = v63
                }
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                v5[("hello" + "")] = (("hello" + "") + "")
                _ = fi61
            }
        }
        _ = fi59
    }
    if (true && false) {
        var v64 uint8 = 232
        _ = v64
    }
    v65 := GenS2{F0: v1, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v65
}
