package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 bool; F2 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool; F1 uint32 }
type GenS6 struct { F0 int8; F1 int; F2 int64 }
type GenS7 struct { F0 int; F1 int64 }
func genFn0(p0 int, p1 uint, p2 uint16) int16 {
    v4 := GenS1{F0: "hello"}
    p2 = 143
    v5 := []int{}
    v5 = append(v5, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 uint64, p1 uint64, p2 int8) string {
    p0 = (p0 - p1)
    v6 := GenS1{F0: "hello"}
    var v7 uint8 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return "hello"
}

func main() {
    var v1 float32 = 1e10
    var v2 uint32 = 222
    v3 := []int{32768, -2147483648, 32767, 127}
    for fi8 := 0; fi8 < 8; fi8++ {
        if false {
            var v9 uint8 = 255
            if false {
                if len(v3) > 0 {
                    v3[0] = (fi8 << ((1) & 63))
                }
                if true {
                    v9 -= 2
                }
                var v10 int8 = int8(v2)
                _ = v10
            }
            var v11 int64 = -1
            _ = v11
            _ = v9
        }
        switch "hello" {
        case "":
            switch 9223372036854775807 {
            case -626:
                switch fi8 {
                case 42:
                    for fi12 := 0; fi12 < 18; fi12++ {
                        fmt.Println(v2, v2, fi12)
                        _ = fi12
                    }
                case 4294967295:
                    if (!false) {
                        v13 := GenS1{F0: ""}
                        _ = v13
                    }
                default:
                    v3 = append(v3, (^(fi8)))
                }
            case 1:
                if (false || true) {
                    var v14 float64 = ((1.5 - 0.0) / 0.1)
                    _ = v14
                } else {
                    if (false || (!((fi8 - fi8) >= (-(fi8))))) {
                        if (0 == 921) {
                            var v15 int64 = -1
                            if len(v3) > 0 {
                                v3[0] = 42
                            }
                            v2 += (v2 & (v2 >> ((1) & 63)))
                            _ = v15
                        }
                        v2 = 1
                        for fi16 := 0; fi16 < 19; fi16++ {
                            fmt.Println(v1)
                            _ = fi16
                        }
                    } else {
                        var v17 int16 = genFn0(fi8, 888, 32767)
                        fmt.Println(v17, fi8)
                        _ = v17
                    }
                }
            case 2147483647:
                var v18 int16 = 0
                _ = v18
            }
        case "foo bar":
            if true {
                if len(v3) > 0 {
                    v3[0] = (fi8 >> ((0) & 63))
                }
                if (!(1e10 == 1.5)) {
                    fmt.Println(v2)
                    if len(v3) > 0 {
                        v3[0] = (fi8 >> ((0) & 63))
                    }
                    for fi19 := 0; fi19 < 19; fi19++ {
                        var v20 int8 = 127
                        _ = fi19
                        _ = v20
                    }
                } else {
                    var v21 uint8 = 128
                    var v22 int8 = 127
                    v22 += (v22 - v22)
                    _ = v21
                    _ = v22
                }
                fmt.Println(v2, v1, v2)
            } else {
                var v23 uint32 = v2
                _ = v23
            }
        default:
            v3 = append(v3, -115)
        }
        _ = fi8
    }
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi24 := 0; fi24 < 1; fi24++ {
            v25 := GenS1{F0: ""}
            var v26 uint = 4294967295
            _ = fi24
            _ = v25
            _ = v26
        }
        if true {
            switch ("hello" + "") {
            case "foo bar":
                v3 = append(v3, -9223372036854775808)
            case "hello":
                _ = len(v3)
            case "\n\t":
                if false {
                    if false {
                        switch 0 {
                        case 32767:
                            if true {
                                fmt.Println(v2)
                                v27 := GenS2{F0: "foo bar", F1: (492 == 127), F2: 829}
                                _ = v27
                            }
                        case 0:
                            var v28 uint16 = 255
                            _ = v28
                        case 42:
                            _ = len(v3)
                        }
                        v2 += v2
                    }
                } else {
                    var v29 string = "hello"
                    _ = v29
                }
            default:
                v2 = 1
            }
            var v30 uint8 = 41
            _ = v30
        } else {
            fmt.Println(v1, v1)
        }
    }
    switch 336 {
    case 625:
        v1 = 1.5
    case -9223372036854775808:
        for fi31 := 0; fi31 < 19; fi31++ {
            v32 := GenS1{F0: "\n\t"}
            _ = fi31
            _ = v32
        }
    default:
        v33 := GenS3{F0: 92}
        _ = v33
    }
    if false {
        v34 := GenS1{F0: "hello"}
        _ = v34
    }
    v35 := GenS3{F0: 92}
    if true {
        fmt.Println(v1, v1)
    } else {
        fmt.Println(v1, v2)
        for fi36 := 0; fi36 < 18; fi36++ {
            if true {
                var v37 bool = true
                if (v37 && (18446744073709551615 <= 42)) {
                    v2 -= (v2 >> ((109) & 63))
                    if (635 != 4294967295) {
                        v38 := GenS4{F0: 42}
                        _ = v38
                    } else {
                        v37 = v37
                        for fi39 := 0; fi39 < 3; fi39++ {
                            fmt.Println(v37)
                            var v40 float32 = (v1 * (float32(((1.5 + float64(v1)) - 1.5)) * (v1 / (0.1 * (((1.5 / v1) + (-2.25 + float32(float64(((float32(((((0.1 / 1.0) + float64((1.5 * 0.0))) - 1.5) * (1.5 - (0.0 - (0.0 + -2.25))))) + (v1 / v1)) * float32(-2.25)))))) + v1)))))
                            fmt.Println(v40)
                            if false {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v1 = float32(-2.25)
                                v41 := GenS3{F0: 37}
                                _ = v41
                            }
                            switch "foo bar" {
                            case "foo bar":
                                switch (fi39 >> ((1) & 63)) {
                                case 184:
                                    switch (("\n\t" + "foo bar") + "") {
                                    case "foo bar":
                                        v3 = append(v3, fi39)
                                    default:
                                        v3 = append(v3, fi39)
                                    }
                                case 128:
                                    v2 = (v2 + 2)
                                default:
                                    v42 := GenS2{F0: "foo bar", F1: (975 == 4294967295), F2: 2}
                                    v43 := GenS5{F0: (170 == uint(v2)), F1: (v2 - 578)}
                                    var v44 uint = 18446744073709551615
                                    _ = v42
                                    _ = v43
                                    _ = v44
                                }
                            }
                            _ = fi39
                            _ = v40
                        }
                    }
                    for fi45 := 0; fi45 < 1; fi45++ {
                        fmt.Println(v37, v2, fi45)
                        fmt.Println(fi36)
                        var v46 uint16 = 486
                        _ = fi45
                        _ = v46
                    }
                } else {
                    _ = len(v3)
                    var v47 string = ("" + "foo bar")
                    _ = v47
                }
                _ = v37
            } else {
                fmt.Println(fi36, v2, v2)
            }
            var v48 string = "hello"
            _ = fi36
            _ = v48
        }
        v2 = 42
    }
    fmt.Println(v2, v1)
    if (0 != 1) {
        v49 := GenS3{F0: 255}
        switch "a" {
        case "a":
            var v50 uint64 = 842
            _ = v50
        case "":
            v51 := GenS2{F0: "", F1: ((("\n\t" + "hello") + "hello") > "\n\t"), F2: 4294967295}
            _ = v51
        case "hello":
            fmt.Println(v1, v2)
        }
        switch 2147483647 {
        case 0:
            fmt.Println(v1, v1)
        }
        _ = v49
    }
    var v52 int64 = -1
    for fi53 := 0; fi53 < 15; fi53++ {
        for fi54 := 0; fi54 < 19; fi54++ {
            if true {
                if false {
                    v1 = (0.0 + (((float32(1.5) / -2.25) * v1) + v1))
                    fmt.Println(fi53)
                    fmt.Println(fi53, fi54)
                }
                var v55 string = "\n\t"
                v52 = 899
                _ = v55
            }
            var v56 float64 = (-2.25 + -2.25)
            fmt.Println(v2)
            _ = fi54
            _ = v56
        }
        _ = fi53
    }
    for fi57 := 0; fi57 < 2; fi57++ {
        if true {
            var v58 uint32 = 224
            for fi59 := 0; fi59 < 3; fi59++ {
                fmt.Println(v1, v58)
                var v60 uint = 620
                v52 += -9223372036854775808
                v61 := GenS6{F0: 100, F1: (fi59 << (((v60 % ((58) | 1))) & 63)), F2: (-(v52))}
                _ = fi59
                _ = v60
                _ = v61
            }
            _ = v58
        }
        v62 := GenS7{F0: 887, F1: v52}
        var v63 int32 = 929
        _ = fi57
        _ = v62
        _ = v63
    }
    var v64 uint16 = 0
    var v65 uint32 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v52:", v52)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v52
    _ = v64
    _ = v65
}
