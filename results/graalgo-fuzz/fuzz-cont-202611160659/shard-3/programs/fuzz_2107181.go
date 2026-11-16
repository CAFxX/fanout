package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int16; F1 string; F2 uint }
type GenS3 struct { F0 uint; F1 int32; F2 uint }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint8; F1 string; F2 uint }
type GenS8 struct { F0 int8 }
func genFn0(p0 string, p1 uint8, p2 uint32) uint16 {
    var v5 int8 = 42
    p0 = ((p0 + "hello") + p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 2
}

func main() {
    v1 := []int{42, 829, 1}
    v2 := []uint32{0, 65535}
    var v3 int64 = 1
    var v4 int8 = 1
    v6 := GenS1{F0: ("hello" + "foo bar"), F1: 0}
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 19; fi8++ {
            fmt.Println(fi7)
            v4 = 127
            _ = fi8
        }
        v4 -= v4
        _ = fi7
    }
    if (v3 < (^(v3))) {
        v9 := GenS2{F0: 247, F1: ("a" + "a"), F2: 466}
        fmt.Println(v4)
        _ = v9
    } else {
        fmt.Println(v4, v4)
        v1 = append(v1, -2147483648)
    }
    var v10 uint8 = 0
    for fi11 := 0; fi11 < 3; fi11++ {
        for fi12 := 0; fi12 < 7; fi12++ {
            switch 253 {
            case 42:
                fmt.Println(v4, v4)
            }
            if true {
                var v13 uint32 = 198
                var v14 float64 = (1.5 - (1.5 - ((0.0 + 0.1) + (0.1 - 0.1))))
                _ = v13
                _ = v14
            } else {
                v15 := GenS2{F0: -32768, F1: "\n\t", F2: 18446744073709551615}
                _ = v15
            }
            v16 := GenS2{F0: 486, F1: "hello", F2: 42}
            _ = fi12
            _ = v16
        }
        for fi17 := 0; fi17 < 17; fi17++ {
            v10 = v10
            switch -2147483648 {
            case 4294967295:
                fmt.Println(fi17, v10, v3)
            case -634:
                switch (-(fi17)) {
                case 127:
                    var v18 int8 = -1
                    _ = v18
                case 32768:
                    if (100 <= (v4 - 100)) {
                        fmt.Println(v10)
                    }
                case 1:
                    switch "foo bar" {
                    case "":
                        if (int64(v10) >= (v3 ^ v3)) {
                            switch "\n\t" {
                            case "foo bar":
                                if (false || true) {
                                    switch "\n\t" {
                                    case "foo bar":
                                        fmt.Println(v3, v4, v3)
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    case "":
                                        if len(v2) > 0 {
                                            v2[0] = 1
                                        }
                                    }
                                    if ((v10 < v10) && true) {
                                        fmt.Println(fi17, v4)
                                    }
                                    v19 := GenS2{F0: 1, F1: "\n\t", F2: 1}
                                    _ = v19
                                }
                                if false {
                                    var v20 string = "a"
                                    var v21 string = (v20 + v20)
                                    switch fi11 {
                                    case 0:
                                        v22 := GenS3{F0: 1, F1: 2147483647, F2: 3}
                                        _ = v22
                                    default:
                                        fmt.Println(fi11, fi17, fi17)
                                    }
                                    var v23 uint8 = 42
                                    var v24 uint = 18446744073709551615
                                    _ = v20
                                    _ = v21
                                    _ = v23
                                    _ = v24
                                }
                            case "":
                                if false {
                                    v4 = v4
                                } else {
                                    if false {
                                        if len(v2) > 0 {
                                            v2[0] = 2147483647
                                        }
                                        _ = len(v2)
                                        fmt.Println(fi17, v3, fi11)
                                    } else {
                                        var v25 int64 = -1
                                        var v26 int8 = (^(v4))
                                        _ = len(v2)
                                        _ = v25
                                        _ = v26
                                    }
                                    fmt.Println(v10, v4)
                                    var v27 float64 = 1e10
                                    var v28 uint8 = v10
                                    _ = len(v2)
                                    _ = v27
                                    _ = v28
                                }
                            }
                            if true {
                                if (!(float64((float32(0.0) / (3.14159 * (0.1 - 0.0)))) > ((((0.1 + 1.5) / 0.1) * 1e10) * 1.5))) {
                                    v3 = -9223372036854775808
                                }
                                v4 = (-(v4))
                            }
                        }
                    case "foo bar":
                        v2 = append(v2, 2147483648)
                    }
                default:
                    fmt.Println(fi17, fi11, v3)
                }
            case -860:
                switch ((("a" + "hello") + "a") + "hello") {
                case "hello":
                    if (42 != 65535) {
                        v29 := GenS1{F0: "foo bar", F1: v10}
                        if (!(true || false)) {
                            fmt.Println(v3, v10, fi11)
                        }
                        _ = len(v1)
                        if true {
                            v30 := GenS3{F0: 42, F1: int32(v4), F2: 426}
                            _ = v30
                        } else {
                            v31 := GenS2{F0: 342, F1: "foo bar", F2: 255}
                            var v32 string = "\n\t"
                            var v33 int = fi11
                            _ = v31
                            _ = v32
                            _ = v33
                        }
                        switch fi17 {
                        case 127:
                            fmt.Println(fi17, fi17)
                        default:
                            switch "hello" {
                            case "":
                                var v34 bool = (0.0 == (1.5 + ((0.1 + 0.1) / -2.25)))
                                _ = v34
                            case "foo bar":
                                v35 := GenS3{F0: 18446744073709551615, F1: -1, F2: 108}
                                _ = v35
                            default:
                                v10 = (v10 % ((12) | 1))
                            }
                            v36 := GenS4{F0: 42, F1: 0}
                            if len(v1) > 0 {
                                v1[0] = fi11
                            }
                            _ = v36
                        }
                        _ = v29
                    }
                default:
                    var v37 int8 = v4
                    var v38 bool = ((2147483647 <= 672) && true)
                    v38 = (!((!((!(-1 > 0)) || (!(!(788 != (fi17 * fi17)))))) || (!false)))
                    _ = v37
                    _ = v38
                }
            default:
                v39 := GenS4{F0: 186, F1: v10}
                _ = v39
            }
            v10 = v10
            v1 = append(v1, fi11)
            _ = fi17
        }
        _ = fi11
    }
    v10 -= 128
    switch 127 {
    case 409:
        switch 9223372036854775807 {
        case 0:
            if false {
                var v40 uint16 = 682
                for fi41 := 0; fi41 < 7; fi41++ {
                    var v42 uint = 569
                    _ = fi41
                    _ = v42
                }
                var v43 string = "\n\t"
                _ = v40
                _ = v43
            } else {
                v44 := GenS1{F0: (("\n\t" + "") + "a"), F1: v10}
                _ = v44
            }
        }
    case 0:
        for fi45 := 0; fi45 < 1; fi45++ {
            fmt.Println(v3, v4, v4)
            if (923 > 42) {
                var v46 uint16 = 65535
                _ = v46
            } else {
                for fi47 := 0; fi47 < 10; fi47++ {
                    v10 = (v10 << ((18446744073709551615) & 63))
                    if (false || (255 > (v10 % ((127) | 1)))) {
                        v4 += v4
                        v3 = (v3 & v3)
                    } else {
                        var v48 int32 = 42
                        v10 += v10
                        var v49 uint64 = 464
                        _ = v48
                        _ = v49
                    }
                    var v50 int16 = -1
                    var v51 int64 = v3
                    var v52 bool = false
                    v10 += 127
                    _ = fi47
                    _ = v50
                    _ = v51
                    _ = v52
                }
            }
            _ = fi45
        }
    default:
        if false {
            var v53 uint64 = 90
            if len(v2) > 0 {
                v2[0] = 2147483648
            }
            _ = v53
        }
    }
    for fi54 := 0; fi54 < 11; fi54++ {
        v55 := GenS5{F0: 128, F1: 2, F2: 0}
        var v56 string = ("" + "\n\t")
        for fi57 := 0; fi57 < 10; fi57++ {
            if len(v2) > 0 {
                v2[0] = 548
            }
            _ = fi57
        }
        _ = fi54
        _ = v55
        _ = v56
    }
    for fi58 := 0; fi58 < 2; fi58++ {
        if true {
            if ((^(v10)) != 127) {
                if ((1.5 > (1.5 + -2.25)) && (v3 >= v3)) {
                    if true {
                        v4 = (v4 % ((v4) | 1))
                        switch "\n\t" {
                        case "\n\t":
                            for fi59 := 0; fi59 < 10; fi59++ {
                                if len(v2) > 0 {
                                    v2[0] = 507
                                }
                                _ = fi59
                            }
                        case "a":
                            v60 := GenS2{F0: -344, F1: (("hello" + "") + "\n\t"), F2: 18446744073709551615}
                            _ = v60
                        }
                        for fi61 := 0; fi61 < 9; fi61++ {
                            var v62 float32 = (-2.25 / -2.25)
                            _ = len(v1)
                            _ = fi61
                            _ = v62
                        }
                    } else {
                        var v63 float64 = float64((float32(0.0) / 1.0))
                        if len(v2) > 0 {
                            v2[0] = 42
                        }
                        _ = v63
                    }
                    v64 := GenS6{F0: 2147483647}
                    _ = v64
                }
            }
            fmt.Println(fi58, v10)
        }
        var v65 int32 = 1
        for fi66 := 0; fi66 < 3; fi66++ {
            var v67 int16 = 177
            if (986 <= 923) {
                v68 := GenS7{F0: (^(v10)), F1: ("" + ("a" + ("hello" + "foo bar"))), F2: 640}
                fmt.Println(v3, v3, v65)
                switch fi58 {
                case 127:
                    if true {
                        var v69 bool = false
                        _ = v69
                    } else {
                        var v70 uint16 = 32767
                        fmt.Println(v70, fi66, v65)
                        var v71 bool = true
                        _ = v70
                        _ = v71
                    }
                case 4294967295:
                    if true {
                        v72 := GenS8{F0: (v4 ^ (^(v4)))}
                        v65 = (v65 >> ((0) & 63))
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = v72
                    } else {
                        v65 += v65
                    }
                }
                _ = v68
            }
            _ = fi66
            _ = v67
        }
        _ = fi58
        _ = v65
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
