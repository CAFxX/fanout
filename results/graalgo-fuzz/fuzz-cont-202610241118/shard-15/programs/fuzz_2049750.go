package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint; F1 int8 }
type GenS6 struct { F0 int16; F1 int64 }
func genFn0(p0 uint8) uint32 {
    switch 302 {
    case 128:
        switch "hello" {
        case "\n\t":
            p0 += p0
        case "a":
            if false {
                p0 -= 188
                var v7 uint16 = 32767
                _ = v7
            } else {
                v8 := []int{-9223372036854775808, -2147483648}
                _ = v8
            }
        case "hello":
            fmt.Println(p0)
        }
    case -221:
        var v9 bool = true
        _ = v9
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        switch fi10 {
        case 9223372036854775807:
            if true {
                for fi11 := 0; fi11 < 20; fi11++ {
                    v12 := GenS1{F0: 1}
                    var v13 int = fi11
                    var v14 float64 = 3.14159
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v14
                }
                var v15 uint64 = 480
                v16 := GenS1{F0: 255}
                _ = v15
                _ = v16
            } else {
                if false {
                    var v17 int64 = -676
                    _ = v17
                }
                for fi18 := 0; fi18 < 15; fi18++ {
                    var v19 float64 = ((0.1 / 1e10) - 0.0)
                    _ = fi18
                    _ = v19
                }
                for fi20 := 0; fi20 < 13; fi20++ {
                    p0 = p0
                    v21 := []int{(-(fi10)), fi20, (fi20 * (fi20 << ((255) & 63)))}
                    p0 = 145
                    switch fi20 {
                    case 32768:
                        var v22 int = fi20
                        _ = v22
                    case 2147483647:
                        var v23 float32 = 0.0
                        _ = v23
                    }
                    if (65535 == 2) {
                        fmt.Println(p0, fi10)
                        switch "a" {
                        case "foo bar":
                            p0 = (^(p0))
                        case "hello":
                            var v24 string = "a"
                            _ = v24
                        }
                    }
                    _ = fi20
                    _ = v21
                }
            }
        default:
            for fi25 := 0; fi25 < 5; fi25++ {
                p0 -= (p0 & p0)
                _ = fi25
            }
        }
        v26 := GenS1{F0: 862}
        _ = fi10
        _ = v26
    }
    v27 := []int{-856, 396, 1, 1}
    if (((0.0 * (0.1 * (0.0 * ((((3.14159 - ((0.1 * 1e10) * -2.25)) - -2.25) / 1.0) - -2.25)))) * -2.25) <= 3.14159) {
        v28 := GenS1{F0: 0}
        if false {
            var v29 uint8 = (^(p0))
            _ = v29
        } else {
            fmt.Println(p0)
            _ = len(v27)
            var v30 uint32 = uint32(p0)
            _ = v30
        }
        _ = v28
    } else {
        v31 := GenS1{F0: 295}
        _ = v31
    }
    _ = p0
    _ = v27
    return 2
}
func genFn1() int {
    if false {
        if (1 > 4294967295) {
            if (42 > 696) {
                var v32 bool = (true && true)
                var v33 int32 = 2147483647
                _ = v32
                _ = v33
            } else {
                if (!true) {
                    var v34 bool = false
                    v35 := []int{32767, 2147483647, 1}
                    _ = v34
                    _ = v35
                } else {
                    var v36 uint16 = 0
                    v37 := GenS2{F0: true}
                    v38 := []int{1}
                    _ = v36
                    _ = v37
                    _ = v38
                }
            }
            if ((158 == 0) && false) {
                if false {
                    if (2 <= 0) {
                        var v39 int8 = 100
                        v39 = v39
                        _ = v39
                    }
                }
            } else {
                var v40 uint16 = 0
                v41 := GenS3{F0: ((("" + ("hello" + "foo bar")) + ("hello" + ("a" + ("foo bar" + ("hello" + ("a" + "\n\t")))))) + "hello"), F1: (("hello" + "foo bar") + (("" + "\n\t") + "foo bar"))}
                _ = v40
                _ = v41
            }
        } else {
            var v42 uint32 = 1
            _ = v42
        }
        v43 := []int{2147483647, 313, 0}
        if false {
            var v44 string = "foo bar"
            v44 = v44
            _ = v44
        } else {
            var v45 string = "hello"
            switch v45 {
            case "a":
                if len(v43) >= 2 {
                    v43 = v43[1:len(v43)]
                }
            }
            _ = v45
        }
        _ = v43
    }
    var v46 bool = false
    fmt.Println(v46)
    if (0 == 42) {
        var v47 int16 = -32768
        _ = v47
    }
    _ = v46
    return -899
}

func main() {
    var v1 string = "hello"
    var v2 int8 = -107
    var v3 uint8 = 0
    var v4 int8 = v2
    var v5 int32 = -2147483648
    var v6 int64 = -736
    v2 -= 100
    fmt.Println(v1, v2)
    fmt.Println(v1, v2, v2)
    v48 := GenS2{F0: true}
    for fi49 := 0; fi49 < 18; fi49++ {
        switch "\n\t" {
        case "hello":
            var v50 uint32 = 2
            _ = v50
        case "a":
            v6 = 1
        default:
            if false {
                v4 = 27
                if false {
                    for fi51 := 0; fi51 < 9; fi51++ {
                        v52 := GenS3{F0: v1, F1: (v1 + v1)}
                        v2 -= (v2 | v2)
                        _ = fi51
                        _ = v52
                    }
                }
            }
        }
        fmt.Println(fi49, fi49)
        v5 -= (v5 + 2147483647)
        _ = fi49
    }
    fmt.Println(v1, v2)
    for fi53 := 0; fi53 < 16; fi53++ {
        switch v1 {
        case "a":
            if (2 == 4294967295) {
                v3 = (v3 % ((v3) | 1))
            }
        }
        v54 := GenS4{F0: (v3 - (v3 / (((v3 - v3)) | 1)))}
        _ = fi53
        _ = v54
    }
    var v55 int32 = -2147483648
    switch "" {
    case "hello":
        fmt.Println(v55, v6, v3)
    case "":
        switch -1 {
        case -1:
            switch 990 {
            case -9223372036854775808:
                v1 = v1
            case 849:
                for fi56 := 0; fi56 < 12; fi56++ {
                    if true {
                        var v57 float64 = (3.14159 * -2.25)
                        _ = v57
                    } else {
                        v58 := GenS4{F0: (v3 >> ((603) & 63))}
                        _ = v58
                    }
                    _ = fi56
                }
            }
        default:
            switch -823 {
            case -2147483648:
                for fi59 := 0; fi59 < 1; fi59++ {
                    v2 -= 127
                    if ((!true) || (255 > (v3 % ((v3) | 1)))) {
                        v60 := []int{genFn1(), genFn1(), (-(fi59)), (fi59 + genFn1())}
                        _ = v60
                    }
                    _ = fi59
                }
            case 1:
                for fi61 := 0; fi61 < 6; fi61++ {
                    v62 := GenS1{F0: 101}
                    if true {
                        v1 = (v1 + (v1 + ("foo bar" + "foo bar")))
                        v2 -= (v2 + (v2 % (((^(v4))) | 1)))
                        fmt.Println(v1)
                    } else {
                        for fi63 := 0; fi63 < 13; fi63++ {
                            var v64 uint32 = genFn0((v3 >> ((0) & 63)))
                            switch genFn1() {
                            case -68:
                                switch genFn1() {
                                case 2147483647:
                                    switch v1 {
                                    case "":
                                        v65 := GenS4{F0: v3}
                                        _ = v65
                                    case "hello":
                                        v55 = (v55 << ((18446744073709551615) & 63))
                                    case "foo bar":
                                        v2 -= 65
                                    default:
                                        v66 := []int{(fi63 >> ((uint(v3)) & 63)), int(v64), (-(fi63)), fi63}
                                        _ = v66
                                    }
                                    var v67 uint64 = 1
                                    _ = v67
                                case 375:
                                    var v68 int32 = (v55 & 256)
                                    _ = v68
                                }
                            }
                            var v69 string = (((v1 + (v1 + "foo bar")) + v1) + "a")
                            _ = fi63
                            _ = v64
                            _ = v69
                        }
                    }
                    _ = fi61
                    _ = v62
                }
            default:
                var v70 int64 = v6
                _ = v70
            }
        }
    default:
        fmt.Println(v3, v5)
    }
    var v71 uint16 = 65535
    v72 := []int{32767, 21, genFn1()}
    if false {
        fmt.Println(v6, v2)
    } else {
        for fi73 := 0; fi73 < 17; fi73++ {
            v4 = (v2 - (v4 << ((7) & 63)))
            if true {
                v4 = -1
                v74 := GenS5{F0: 65535, F1: (-(v4))}
                _ = v74
            } else {
                for fi75 := 0; fi75 < 7; fi75++ {
                    if (125 > 516) {
                        v2 += 42
                        fmt.Println(v2)
                        if (!(false && ((-285 >= 42) || false))) {
                            var v76 int32 = -955
                            _ = v76
                        }
                    }
                    _ = fi75
                }
                v55 = v55
                for fi77 := 0; fi77 < 6; fi77++ {
                    v6 -= (v6 + (-(v6)))
                    v78 := GenS5{F0: 957, F1: v4}
                    _ = fi77
                    _ = v78
                }
            }
            if len(v72) >= 2 {
                v72 = v72[1:len(v72)]
            }
            _ = fi73
        }
        v5 = v55
        v79 := GenS6{F0: -1, F1: (v6 % ((728) | 1))}
        _ = v79
    }
    var v80 string = v1
    var v81 int64 = v6
    var v82 uint64 = 18446744073709551615
    switch v80 {
    case "a":
        fmt.Println(v71)
    }
    var v83 bool = false
    v55 += (v5 & -295)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    fmt.Println("v71:", v71)
    fmt.Println("slicelen v72:", len(v72), v72)
    fmt.Println("v80:", v80)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v55
    _ = v6
    _ = v71
    _ = v72
    _ = v80
    _ = v81
    _ = v82
    _ = v83
}
