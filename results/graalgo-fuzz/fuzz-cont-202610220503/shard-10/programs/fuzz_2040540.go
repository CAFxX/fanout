package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 string }
type GenS2 struct { F0 uint; F1 int; F2 uint8 }
type GenS3 struct { F0 int32; F1 uint64; F2 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16; F1 uint; F2 int }
func genFn0(p0 int8, p1 int64, p2 uint) uint64 {
    v4 := GenS1{F0: 2, F1: "hello", F2: "a"}
    for fi5 := 0; fi5 < 4; fi5++ {
        for fi6 := 0; fi6 < 20; fi6++ {
            if true {
                p2 = p2
                if (1e10 <= (float32(3.14159) + 1.5)) {
                    if true {
                        v7 := []int{9223372036854775807, 229, (fi5 * int(p2))}
                        _ = v7
                    }
                }
            } else {
                fmt.Println(p1, p1)
            }
            v8 := GenS1{F0: 18446744073709551615, F1: "hello", F2: "hello"}
            _ = fi6
            _ = v8
        }
        for fi9 := 0; fi9 < 10; fi9++ {
            if true {
                p0 += p0
                fmt.Println(p1, p1)
                var v10 string = "hello"
                _ = v10
            }
            fmt.Println(fi5)
            if false {
                var v11 float64 = (-2.25 + 3.14159)
                v12 := GenS2{F0: p2, F1: -2147483648, F2: 42}
                if true {
                    fmt.Println(fi5)
                    var v13 uint16 = 255
                    v13 = (v13 * v13)
                    v11 -= ((v11 + ((0.1 + ((float64(1.5) / ((v11 / float64(1.5)) * (v11 - 3.14159))) - v11)) * (3.14159 * (v11 + v11)))) * (v11 * (1e10 / v11)))
                    switch ("" + ("foo bar" + ("hello" + "\n\t"))) {
                    case "foo bar":
                        if (-1 < 2147483647) {
                            if (true && false) {
                                v14 := GenS3{F0: 1, F1: uint64(v13), F2: int64(v13)}
                                v11 = v11
                                _ = v14
                            } else {
                                var v15 float32 = (float32(v11) / 0.1)
                                var v16 int32 = 257
                                _ = v15
                                _ = v16
                            }
                            var v17 uint = p2
                            fmt.Println(v17)
                            _ = v17
                        } else {
                            var v18 int32 = 1
                            _ = v18
                        }
                    }
                    v13 += (^(v13))
                    _ = v13
                } else {
                    if true {
                        if ((p1 | (p1 << ((42) & 63))) >= (p1 | p1)) {
                            if (656 <= 511) {
                                p2 = p2
                                v19 := []int{}
                                fmt.Println(p1, fi5)
                                _ = v19
                            } else {
                                switch fi5 {
                                case -494:
                                    switch "hello" {
                                    case "\n\t":
                                        switch "hello" {
                                        case "hello":
                                            p2 += 255
                                        default:
                                            v20 := []int{(fi5 ^ (-(fi9)))}
                                            _ = v20
                                        }
                                    case "hello":
                                        var v21 int = (fi5 >> ((p2) & 63))
                                        _ = v21
                                    case "":
                                        v22 := GenS1{F0: 18446744073709551615, F1: "a", F2: ""}
                                        _ = v22
                                    default:
                                        if (!(true || (!false))) {
                                            v23 := GenS3{F0: int32(p0), F1: 367, F2: -9223372036854775808}
                                            _ = v23
                                        } else {
                                            v24 := []int{(fi9 ^ (fi9 - (fi5 & (fi9 + (fi5 | fi5)))))}
                                            if false {
                                                v24 = append(v24, -275)
                                                _ = len(v24)
                                                fmt.Println(p0)
                                                if len(v24) >= 2 {
                                                    v24 = v24[1:len(v24)]
                                                }
                                            }
                                            _ = v24
                                        }
                                    }
                                case 32767:
                                    var v25 int = -9223372036854775808
                                    _ = v25
                                case 0:
                                    p1 -= p1
                                default:
                                    v11 -= 0.0
                                }
                                var v26 uint32 = 4294967295
                                _ = v26
                            }
                        }
                        fmt.Println(p2)
                    } else {
                        switch 32768 {
                        case 393:
                            v27 := GenS3{F0: 1, F1: 18446744073709551615, F2: p1}
                            _ = v27
                        case 9223372036854775807:
                            fmt.Println(v11, p1)
                        case 1:
                            var v28 float32 = float32(v11)
                            _ = v28
                        default:
                            v29 := []int{}
                            _ = v29
                        }
                    }
                    v30 := []int{(fi5 * (^(fi9)))}
                    _ = v30
                }
                _ = v11
                _ = v12
            } else {
                var v31 uint = 443
                _ = v31
            }
            _ = fi9
        }
        _ = fi5
    }
    var v32 bool = (42 != p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v32
    _ = v4
    return 636
}
func genFn1() int8 {
    var v33 bool = (false || (!(!false)))
    var v34 float64 = (((float64(((0.1 - 1.5) * float32(1e10))) + -2.25) / 1e10) - (0.0 + (float64(3.14159) - float64((3.14159 / float32(0.1))))))
    v33 = v33
    v35 := GenS1{F0: 2, F1: ("foo bar" + "hello"), F2: ("\n\t" + ("\n\t" + ""))}
    _ = v33
    _ = v34
    _ = v35
    return 42
}

func main() {
    var v1 uint = 42
    var v2 float64 = (0.1 + float64((1e10 / (1e10 * (float32(0.0) - 3.14159)))))
    v3 := []string{("hello" + "foo bar"), "hello", "hello", ("" + "hello")}
    v1 += v1
    if false {
        switch -140 {
        case -2147483648:
            v36 := GenS1{F0: 18446744073709551615, F1: "\n\t", F2: "\n\t"}
            _ = v36
        default:
            v1 += (v1 * 2)
        }
    } else {
        var v37 uint32 = 937
        var v38 bool = true
        switch -9223372036854775808 {
        case 2147483647:
            for fi39 := 0; fi39 < 3; fi39++ {
                var v40 uint32 = (v37 * v37)
                var v41 float32 = 0.1
                for fi42 := 0; fi42 < 11; fi42++ {
                    v43 := GenS4{F0: 9223372036854775807}
                    v38 = (!v38)
                    if ((v40 + 2147483648) == uint32(v1)) {
                        switch (fi42 + -1) {
                        case 0:
                            if (-1 > genFn1()) {
                                var v44 int32 = 2147483647
                                fmt.Println(fi42)
                                _ = v44
                            } else {
                                var v45 int = fi42
                                var v46 int64 = 9223372036854775807
                                var v47 int16 = 1
                                _ = v45
                                _ = v46
                                _ = v47
                            }
                        case 128:
                            v2 = (((0.0 * v2) * ((0.0 - 1.5) * (0.0 - float64(v41)))) / v2)
                        case -917:
                            v38 = v38
                        }
                        v1 = 18446744073709551615
                        v2 = v2
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    v48 := GenS1{F0: 0, F1: "\n\t", F2: "foo bar"}
                    v41 -= float32(float64(3.14159))
                    fmt.Println(v40, v40, v38)
                    _ = fi42
                    _ = v43
                    _ = v48
                }
                _ = fi39
                _ = v40
                _ = v41
            }
        }
        _ = v37
        _ = v38
    }
    var v49 int8 = 127
    var v50 uint16 = 32767
    if false {
        if false {
            v51 := GenS4{F0: 42}
            _ = v51
        } else {
            if (genFn0((v49 * -1), -690, v1) != 4294967295) {
                var v52 int64 = 42
                _ = v52
            }
        }
    }
    _ = len(v3)
    v3 = append(v3, "\n\t")
    v50 = 42
    fmt.Println(v2, v49)
    var v53 uint16 = (v50 << (((^(v1))) & 63))
    v54 := GenS1{F0: 0, F1: "hello", F2: "foo bar"}
    var v55 int16 = 0
    v50 -= (v53 << ((v1) & 63))
    var v56 int = 287
    switch int(v49) {
    case 32767:
        v57 := GenS1{F0: 0, F1: (((("\n\t" + "hello") + "foo bar") + ("" + "a")) + "\n\t"), F2: "a"}
        _ = v57
    default:
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
    }
    switch (v56 & (v56 / (((v56 + -650)) | 1))) {
    case -790:
        for fi58 := 0; fi58 < 18; fi58++ {
            var v59 uint = (v1 ^ v1)
            var v60 uint8 = 155
            _ = fi58
            _ = v59
            _ = v60
        }
    case 0:
        if ((v1 ^ v1) <= 114) {
            var v61 int16 = int16(v1)
            var v62 int64 = -645
            for fi63 := 0; fi63 < 9; fi63++ {
                v3 = append(v3, ("" + "foo bar"))
                var v64 bool = (!true)
                _ = fi63
                _ = v64
            }
            _ = v61
            _ = v62
        }
    default:
        v65 := GenS3{F0: 180, F1: 1, F2: -9223372036854775808}
        _ = v65
    }
    if false {
        switch "\n\t" {
        case "hello":
            var v66 string = "\n\t"
            _ = v66
        default:
            switch v56 {
            case 1:
                v67 := GenS5{F0: (v50 | (v53 * (v53 ^ v50))), F1: (v1 | (^(v1))), F2: v56}
                _ = v67
            default:
                switch v56 {
                case -618:
                    var v68 string = ""
                    _ = v68
                default:
                    switch 32767 {
                    case -1:
                        _ = len(v3)
                    default:
                        var v69 uint = v1
                        _ = v69
                    }
                }
            }
        }
        _ = len(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    fmt.Println("v53:", v53)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v49
    _ = v50
    _ = v53
    _ = v54
    _ = v55
    _ = v56
}
