package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 int64; F2 uint16 }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 uint32; F1 bool; F2 int64 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint16; F1 string }
type GenS8 struct { F0 uint16; F1 string; F2 uint32 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int16; F1 int32; F2 uint32 }
type GenS11 struct { F0 string; F1 uint32 }
type GenS12 struct { F0 uint32; F1 int32 }
func genFn0(p0 int64, p1 int32, p2 string) int16 {
    v4 := GenS1{F0: 9223372036854775807, F1: p1, F2: 2}
    v5 := []int{402}
    v6 := GenS1{F0: 308, F1: p1, F2: 0}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch p2 {
    case "hello":
        p2 = (((p2 + "foo bar") + (p2 + (p2 + p2))) + (p2 + "foo bar"))
    case "":
        _ = len(v5)
    case "a":
        if true {
            for fi7 := 0; fi7 < 13; fi7++ {
                var v8 uint64 = 18446744073709551615
                if ((-2.25 - 0.0) < 3.14159) {
                    fmt.Println(p2, p2, p0)
                    p1 += p1
                    var v9 uint64 = 179
                    _ = v9
                }
                if false {
                    v8 += (^(v8))
                    switch p2 {
                    case "a":
                        if (42 > 0) {
                            if false {
                                p2 = ("" + (p2 + p2))
                                if false {
                                    var v10 uint64 = 0
                                    _ = len(v5)
                                    _ = v10
                                } else {
                                    v5 = append(v5, (-(fi7)))
                                    var v11 int = (fi7 << ((42) & 63))
                                    _ = v11
                                }
                                fmt.Println(p1)
                            }
                        } else {
                            switch p2 {
                            case "foo bar":
                                var v12 uint8 = 176
                                _ = v12
                            }
                            if (true && (1.5 >= ((1e10 * -2.25) / 0.1))) {
                                p2 = "\n\t"
                                switch fi7 {
                                case 128:
                                    v13 := GenS2{F0: "\n\t"}
                                    _ = v13
                                default:
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                            } else {
                                for fi14 := 0; fi14 < 15; fi14++ {
                                    switch fi7 {
                                    case 32768:
                                        if false {
                                            var v15 bool = false
                                            _ = v15
                                        }
                                    case 541:
                                        var v16 int16 = 32767
                                        _ = v16
                                    default:
                                        var v17 float64 = 1.5
                                        _ = v17
                                    }
                                    fmt.Println(p0, fi7, fi7)
                                    _ = fi14
                                }
                            }
                        }
                    case "\n\t":
                        var v18 uint8 = 38
                        _ = v18
                    }
                    for fi19 := 0; fi19 < 14; fi19++ {
                        p0 = p0
                        _ = len(v5)
                        _ = fi19
                    }
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v20 float64 = (-2.25 - (3.14159 / float64(-2.25)))
                    if true {
                        p2 = "foo bar"
                        fmt.Println(fi7, v8, p1)
                    }
                    _ = v20
                }
                _ = fi7
                _ = v8
            }
            p1 += p1
        } else {
            var v21 float64 = 3.14159
            p1 += (-(p1))
            _ = v21
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return -1
}
func genFn1(p0 uint8, p1 int16) string {
    var v22 uint32 = 200
    v23 := []int{896}
    fmt.Println(p1, p0, p0)
    switch "foo bar" {
    case "\n\t":
        if len(v23) >= 2 {
            v23 = v23[1:len(v23)]
        }
    case "foo bar":
        v24 := GenS3{F0: genFn0(0, 1, "\n\t"), F1: 42, F2: 32767}
        _ = v24
    case "":
        for fi25 := 0; fi25 < 4; fi25++ {
            if len(v23) >= 2 {
                v23 = v23[1:len(v23)]
            }
            for fi26 := 0; fi26 < 5; fi26++ {
                fmt.Println(fi26, fi25)
                _ = fi26
            }
            v27 := GenS4{F0: p0, F1: 32767}
            _ = fi25
            _ = v27
        }
    }
    fmt.Println(p1, p1, p0)
    _ = p0
    _ = p1
    _ = v22
    _ = v23
    return ""
}

func main() {
    var v1 int = -2147483648
    var v2 float64 = 1e10
    v3 := []int{v1, (v1 << ((18446744073709551615) & 63))}
    v1 = (v1 + -9223372036854775808)
    v28 := GenS5{F0: 408, F1: (false || (1.5 <= 3.14159)), F2: -9223372036854775808}
    for fi29 := 0; fi29 < 16; fi29++ {
        fmt.Println(v1, fi29, fi29)
        _ = fi29
    }
    v30 := GenS6{F0: -449}
    var v31 string = "a"
    v32 := GenS7{F0: 357, F1: (("" + (v31 + (v31 + ("foo bar" + ((v31 + ((v31 + "hello") + (v31 + (v31 + (v31 + (v31 + (v31 + v31))))))) + v31))))) + "a")}
    for fi33 := 0; fi33 < 10; fi33++ {
        v34 := GenS8{F0: 255, F1: "foo bar", F2: 2}
        v35 := GenS2{F0: (((("a" + v31) + (v31 + v31)) + "") + v31)}
        _ = fi33
        _ = v34
        _ = v35
    }
    if (!true) {
        var v36 uint = 2
        for fi37 := 0; fi37 < 10; fi37++ {
            if false {
                for fi38 := 0; fi38 < 16; fi38++ {
                    if len(v3) > 0 {
                        v3[0] = (fi38 / ((-9223372036854775808) | 1))
                    }
                    _ = fi38
                }
                v31 = "foo bar"
                switch -365 {
                case -947:
                    var v39 float64 = v2
                    _ = v39
                case 9223372036854775807:
                    fmt.Println(v1, v31)
                default:
                    var v40 int = (fi37 >> ((v36) & 63))
                    _ = v40
                }
            } else {
                var v41 uint16 = 671
                v42 := GenS9{F0: 676}
                _ = v41
                _ = v42
            }
            if true {
                switch "foo bar" {
                case "\n\t":
                    v43 := GenS10{F0: 0, F1: 416, F2: 42}
                    _ = v43
                case "hello":
                    var v44 int16 = 919
                    _ = v44
                }
            } else {
                for fi45 := 0; fi45 < 12; fi45++ {
                    v46 := GenS10{F0: 1, F1: 2147483647, F2: 536}
                    var v47 float32 = 0.1
                    _ = fi45
                    _ = v46
                    _ = v47
                }
            }
            switch (-(fi37)) {
            case 1:
                v31 = (v31 + "a")
            default:
                switch "a" {
                case "":
                    fmt.Println(v36, v1)
                }
            }
            _ = fi37
        }
        if (false && false) {
            v31 = ((v31 + v31) + v31)
            v48 := GenS2{F0: v31}
            _ = v48
        }
        _ = v36
    } else {
        fmt.Println(v1)
        var v49 int = v1
        _ = v49
    }
    var v50 uint8 = 2
    v1 -= 0
    var v51 int32 = 0
    if (v50 != v50) {
        var v52 uint = 2
        var v53 int = -9223372036854775808
        _ = v52
        _ = v53
    }
    for fi54 := 0; fi54 < 20; fi54++ {
        var v55 uint8 = v50
        v56 := GenS3{F0: -1, F1: 1, F2: 255}
        _ = fi54
        _ = v55
        _ = v56
    }
    if (v50 >= (v50 * 42)) {
        switch v1 {
        case -1:
            var v57 int64 = 584
            _ = v57
        case 32768:
            _ = len(v3)
        case 0:
            v2 = 0.1
        }
        if true {
            v58 := GenS9{F0: 2}
            for fi59 := 0; fi59 < 10; fi59++ {
                fmt.Println(v31, v50, fi59)
                var v60 int32 = (v51 & (v51 | (v51 % ((v51) | 1))))
                _ = fi59
                _ = v60
            }
            _ = v58
        }
        for fi61 := 0; fi61 < 2; fi61++ {
            switch v31 {
            case "":
                var v62 string = v31
                _ = v62
            default:
                v63 := GenS11{F0: v31, F1: 520}
                _ = v63
            }
            _ = fi61
        }
    } else {
        v64 := GenS10{F0: 1000, F1: (v51 << ((42) & 63)), F2: 2147483647}
        _ = v64
    }
    v65 := GenS12{F0: 2, F1: (v51 >> ((1) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v50
    _ = v51
    _ = v65
}
