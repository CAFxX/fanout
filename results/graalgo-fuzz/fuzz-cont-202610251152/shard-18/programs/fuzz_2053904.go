package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 int32; F1 uint; F2 int }
type GenS6 struct { F0 uint64; F1 int32 }
type GenS7 struct { F0 uint32; F1 int8; F2 uint }
type GenS8 struct { F0 uint16 }
func genFn0(p0 int8, p1 uint) int8 {
    fmt.Println(p0, p1)
    v6 := GenS1{F0: -2147483648, F1: (p1 - (p1 / ((131) | 1)))}
    _ = p0
    _ = p1
    _ = v6
    return (p0 + (p0 & (p0 ^ -6)))
}

func main() {
    v1 := []string{}
    var v2 uint = 42
    var v3 uint64 = 725
    var v4 int64 = 986
    v5 := make(map[uint32]int)
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 int32 = 288
        if ((!false) || (3.14159 >= 3.14159)) {
            var v9 uint16 = 65535
            v10 := v5[4294967295]
            _ = v10
            _ = v9
        } else {
            switch 42 {
            case 5:
                fmt.Println(v3)
            case 42:
                var v11 uint32 = 65535
                _ = v11
            default:
                fmt.Println(v2)
            }
            for fi12 := 0; fi12 < 15; fi12++ {
                v8 = v8
                v8 = v8
                _ = fi12
            }
            v4 = (v4 / ((927) | 1))
        }
        for fi13 := 0; fi13 < 17; fi13++ {
            switch "\n\t" {
            case "":
                v14 := GenS1{F0: (^(v8)), F1: 42}
                _ = v14
            case "foo bar":
                v15 := GenS1{F0: v8, F1: v2}
                if ("hello" != "\n\t") {
                    if len(v1) > 0 {
                        v1[0] = "\n\t"
                    }
                    fmt.Println(fi7, v3, fi13)
                } else {
                    var v16 int16 = 105
                    _ = v16
                }
                _ = v15
            case "hello":
                if (!(!false)) {
                    v1 = append(v1, "a")
                }
            }
            var v17 int8 = 1
            if true {
                var v18 string = ("hello" + ("foo bar" + "a"))
                _ = v18
            }
            _ = fi13
            _ = v17
        }
        _ = fi7
        _ = v8
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        for fi20 := 0; fi20 < 2; fi20++ {
            v21 := GenS2{F0: (3.14159 > ((0.0 - float32(0.1)) - 1.5)), F1: 1, F2: 42}
            var v22 bool = false
            switch (^(fi20)) {
            case -2147483648:
                var v23 int64 = v4
                _ = v23
            case -1:
                if (2147483647 < 1) {
                    switch "" {
                    case "foo bar":
                        var v24 uint8 = 1
                        var v25 uint16 = 1
                        _ = len(v1)
                        _ = v24
                        _ = v25
                    case "hello":
                        switch fi20 {
                        case 247:
                            _ = len(v5)
                        }
                    case "a":
                        if len(v1) > 0 {
                            v1[0] = "a"
                        }
                    default:
                        v26 := GenS1{F0: 0, F1: (v2 * v2)}
                        _ = v26
                    }
                    var v27 int = (fi20 >> (((v2 | (v2 >> (((^(v2))) & 63)))) & 63))
                    _ = v27
                }
            default:
                if v22 {
                    if true {
                        if v22 {
                            v28 := GenS3{F0: (fi19 % ((fi20) | 1)), F1: 1}
                            if false {
                                fmt.Println(fi19)
                                if v22 {
                                    fmt.Println(v22, v2, v4)
                                }
                                if (v22 && true) {
                                    v4 = v4
                                    v29 := GenS1{F0: 690, F1: v2}
                                    _ = v29
                                }
                            }
                            var v30 uint = v2
                            if false {
                                v31 := v5[4294967295]
                                var v32 int64 = (v4 - (v4 % ((v4) | 1)))
                                _ = v31
                                _ = v32
                            } else {
                                var v33 int = (fi19 | int(v4))
                                var v34 int = (-(fi20))
                                _ = v33
                                _ = v34
                            }
                            _ = v28
                            _ = v30
                        } else {
                            if v22 {
                                switch "\n\t" {
                                case "hello":
                                    var v35 bool = v22
                                    _ = v35
                                default:
                                    v3 += v3
                                }
                                v36 := GenS2{F0: (1 >= 42), F1: 42, F2: 656}
                                _ = v36
                            }
                            v1 = append(v1, ("foo bar" + ("foo bar" + "\n\t")))
                        }
                    }
                    switch "hello" {
                    case "a":
                        var v37 bool = false
                        v4 = v4
                        v38 := GenS4{F0: 2, F1: 42}
                        _ = v37
                        _ = v38
                    case "":
                        v39 := GenS2{F0: ((131 == 68) && v22), F1: 2147483648, F2: 1}
                        _ = v39
                    case "hello":
                        if false {
                            if v22 {
                                var v40 uint16 = 65535
                                _ = v40
                            } else {
                                switch "\n\t" {
                                case "\n\t":
                                    switch ("foo bar" + "a") {
                                    case "a":
                                        if (52 > 42) {
                                            switch "a" {
                                            case "foo bar":
                                                fmt.Println(v22)
                                            case "hello":
                                                if len(v1) > 0 {
                                                    v1[0] = "foo bar"
                                                }
                                            }
                                            if false {
                                                v41 := v5[4294967295]
                                                _ = v41
                                            } else {
                                                fmt.Println(v4)
                                                v3 = v3
                                            }
                                            v42 := GenS5{F0: -696, F1: v2, F2: fi19}
                                            _ = v42
                                        }
                                    case "foo bar":
                                        var v43 int64 = (v4 / ((v4) | 1))
                                        _ = v43
                                    case "":
                                        v22 = ((!v22) && ((float32(float64(1e10)) != float32((3.14159 - 1.5))) && v22))
                                    }
                                default:
                                    if false {
                                        if ((0 > 2147483648) && true) {
                                            v1 = append(v1, "foo bar")
                                        }
                                    }
                                }
                                switch "" {
                                case "a":
                                    switch (fi19 ^ (fi19 & fi19)) {
                                    case 0:
                                        v44 := GenS1{F0: 0, F1: (v2 / (((v2 ^ 255)) | 1))}
                                        _ = v44
                                    case -693:
                                        var v45 uint32 = 2147483647
                                        _ = v45
                                    default:
                                        v4 -= (v4 / ((0) | 1))
                                    }
                                case "hello":
                                    v2 -= 18446744073709551615
                                case "\n\t":
                                    v46 := GenS6{F0: (v3 * v3), F1: -1}
                                    _ = v46
                                default:
                                    v3 = (v3 / ((v3) | 1))
                                }
                            }
                            if (367 == 82) {
                                var v47 int64 = v4
                                var v48 float32 = (0.0 - 1e10)
                                if false {
                                    if (1 <= -68) {
                                        var v49 int32 = 0
                                        var v50 int = (fi19 * (fi20 >> ((v2) & 63)))
                                        _ = len(v5)
                                        v48 += v48
                                        _ = v49
                                        _ = v50
                                    } else {
                                        var v51 uint16 = 483
                                        v22 = (!v22)
                                        if (!(-1 <= -128)) {
                                            if ((fi20 - 32767) > 594) {
                                                v48 = v48
                                            }
                                            var v52 int64 = v47
                                            _ = v52
                                        } else {
                                            switch 127 {
                                            case 128:
                                                v48 += 3.14159
                                            default:
                                                v3 += (v3 ^ (v3 + 1))
                                            }
                                            switch (("" + "foo bar") + "a") {
                                            case "foo bar":
                                                var v53 float64 = (3.14159 * -2.25)
                                                _ = v53
                                            case "a":
                                                if len(v1) > 0 {
                                                    v1[0] = ("a" + "hello")
                                                }
                                            default:
                                                fmt.Println(v22, v2)
                                            }
                                        }
                                        v51 = 1
                                        if v22 {
                                            var v54 uint8 = 202
                                            v55 := GenS3{F0: fi19, F1: v51}
                                            switch -1 {
                                            case 4294967295:
                                                var v56 string = ((("" + "") + "") + "\n\t")
                                                _ = v56
                                            case 32767:
                                                v22 = true
                                            case 1:
                                                fmt.Println(v54)
                                            default:
                                                if v22 {
                                                    v47 -= v4
                                                    var v57 int = fi19
                                                    _ = v57
                                                } else {
                                                    if v22 {
                                                        if ((v51 & (v51 / (((^(v51))) | 1))) == v51) {
                                                            v58 := GenS1{F0: 1, F1: v2}
                                                            _ = v58
                                                        } else {
                                                            v4 += (v47 * (-(v4)))
                                                        }
                                                    }
                                                    switch (("a" + "") + (("hello" + ("hello" + ((("foo bar" + "") + "") + "foo bar"))) + "a")) {
                                                    case "\n\t":
                                                        v4 -= 9223372036854775807
                                                    default:
                                                        fmt.Println(v48, v3)
                                                        var v59 bool = ((40 != -63) || v22)
                                                        _ = v59
                                                    }
                                                    v60 := GenS5{F0: 2147483647, F1: (v2 & 641), F2: (fi19 - 128)}
                                                    _ = v60
                                                }
                                            }
                                            _ = v54
                                            _ = v55
                                        } else {
                                            v61 := GenS7{F0: 4294967295, F1: 127, F2: 275}
                                            if ((true || v22) && v22) {
                                                v62 := GenS7{F0: 4294967295, F1: 127, F2: v2}
                                                v47 += v47
                                                _ = v62
                                            } else {
                                                _ = len(v5)
                                            }
                                            if ((v51 << (((^(v2))) & 63)) == 32767) {
                                                v63 := v5[42]
                                                _ = v63
                                            } else {
                                                v4 = (v4 - v47)
                                                var v64 uint = (^(v2))
                                                _ = v64
                                            }
                                            v4 = v47
                                            fmt.Println(v2)
                                            _ = v61
                                        }
                                        v5[830] = 128
                                        var v65 int16 = -1
                                        _ = v51
                                        _ = v65
                                    }
                                    _ = len(v1)
                                    delete(v5, 42)
                                } else {
                                    v66 := GenS8{F0: 32767}
                                    v67 := GenS7{F0: 4294967295, F1: genFn0(genFn0(115, 0), v2), F2: v2}
                                    _ = v66
                                    _ = v67
                                }
                                var v68 int16 = int16(v3)
                                _ = v47
                                _ = v48
                                _ = v68
                            }
                            var v69 uint64 = v3
                            _ = v69
                        }
                    }
                    fmt.Println(fi19)
                }
            }
            _ = len(v1)
            _ = fi20
            _ = v21
            _ = v22
        }
        delete(v5, 0)
        _ = fi19
    }
    _ = len(v1)
    var v70 int8 = -1
    v71 := GenS7{F0: 315, F1: genFn0(19, 42), F2: (v2 << (((v2 % ((v2) | 1))) & 63))}
    if false {
        for fi72 := 0; fi72 < 10; fi72++ {
            var v73 uint64 = v3
            if true {
                var v74 int64 = (v4 / ((0) | 1))
                _ = v74
            }
            _ = fi72
            _ = v73
        }
        for fi75 := 0; fi75 < 4; fi75++ {
            delete(v5, 65535)
            _ = fi75
        }
        for fi76 := 0; fi76 < 15; fi76++ {
            for fi77 := 0; fi77 < 15; fi77++ {
                var v78 uint64 = v3
                v2 += 758
                _ = fi77
                _ = v78
            }
            var v79 int8 = (v70 << ((v2) & 63))
            _ = fi76
            _ = v79
        }
    }
    var v80 string = ("\n\t" + "foo bar")
    v5[42] = 362
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v70:", v70)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v70
    _ = v71
    _ = v80
}
