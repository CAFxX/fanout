package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 int; F2 string }
type GenS4 struct { F0 int16; F1 int32; F2 int }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint; F1 int32 }
type GenS7 struct { F0 int8; F1 bool; F2 int }
type GenS8 struct { F0 uint64; F1 uint32 }
func genFn0(p0 int64) uint64 {
    v6 := []int{}
    for fi7 := 0; fi7 < 16; fi7++ {
        if (42 == -1) {
            v8 := GenS1{F0: (fi7 >> ((18446744073709551615) & 63)), F1: 0}
            _ = v8
        } else {
            p0 = p0
            p0 -= p0
            if (1e10 != (float64(((1e10 - 1e10) / -2.25)) * 0.0)) {
                if ("hello" != "") {
                    v9 := GenS1{F0: 9223372036854775807, F1: 577}
                    _ = v9
                }
                v10 := GenS1{F0: (fi7 << ((42) & 63)), F1: 461}
                _ = v10
            }
        }
        var v11 string = ((("" + "a") + "foo bar") + ((("" + "a") + ((((("hello" + "a") + "") + "") + (("foo bar" + "a") + ("hello" + "foo bar"))) + ("" + "foo bar"))) + ("a" + "a")))
        _ = fi7
        _ = v11
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 float64 = 1e10
        _ = fi12
        _ = v13
    }
    v14 := GenS2{F0: 0}
    for fi15 := 0; fi15 < 16; fi15++ {
        p0 = (p0 & p0)
        v16 := GenS3{F0: 222, F1: 32767, F2: "hello"}
        var v17 uint32 = 280
        _ = fi15
        _ = v16
        _ = v17
    }
    _ = p0
    _ = v14
    _ = v6
    return 1
}
func genFn1() uint64 {
    for fi18 := 0; fi18 < 3; fi18++ {
        for fi19 := 0; fi19 < 4; fi19++ {
            switch (((("foo bar" + "\n\t") + "hello") + ("foo bar" + "\n\t")) + "hello") {
            case "hello":
                var v20 string = "\n\t"
                _ = v20
            case "\n\t":
                switch (-(fi19)) {
                case -2147483648:
                    v21 := []int{-476, fi19, 128, (fi18 << ((255) & 63))}
                    _ = v21
                case 401:
                    v22 := []int{fi19, (^(fi18))}
                    _ = v22
                case 32768:
                    v23 := GenS3{F0: 42, F1: 9223372036854775807, F2: "hello"}
                    _ = v23
                default:
                    fmt.Println(fi19)
                }
            case "":
                v24 := []int{fi18}
                _ = v24
            default:
                fmt.Println(fi18, fi19)
            }
            _ = fi19
        }
        for fi25 := 0; fi25 < 20; fi25++ {
            if (0.0 == (float64(0.1) - float64(0.0))) {
                var v26 int64 = 9223372036854775807
                v26 -= v26
                var v27 uint16 = 1
                _ = v26
                _ = v27
            }
            var v28 float32 = (3.14159 / 1.0)
            _ = fi25
            _ = v28
        }
        _ = fi18
    }
    for fi29 := 0; fi29 < 12; fi29++ {
        v30 := GenS1{F0: fi29, F1: 42}
        v31 := []int{}
        if (!(-1 >= -2147483648)) {
            var v32 int = (fi29 >> ((4294967295) & 63))
            fmt.Println(v32)
            fmt.Println(fi29, v32)
            _ = v32
        } else {
            var v33 bool = (255 <= 128)
            for fi34 := 0; fi34 < 20; fi34++ {
                if v33 {
                    fmt.Println(v33)
                    fmt.Println(fi29)
                    v33 = (42 != 1)
                } else {
                    var v35 uint32 = 1
                    v35 = v35
                    if len(v31) > 0 {
                        v31[0] = fi34
                    }
                    var v36 bool = v33
                    v33 = ((("foo bar" + ("hello" + "foo bar")) + ("\n\t" + "a")) >= (("hello" + "") + (("hello" + "\n\t") + "")))
                    _ = v35
                    _ = v36
                }
                _ = fi34
            }
            _ = v33
        }
        _ = fi29
        _ = v30
        _ = v31
    }
    var v37 uint16 = 42
    _ = v37
    return 744
}

func main() {
    var v1 int16 = 1000
    var v2 float32 = -2.25
    v3 := []string{("hello" + ""), "foo bar"}
    var v4 string = ("foo bar" + ("\n\t" + ("foo bar" + ("foo bar" + "hello"))))
    var v5 uint64 = 18446744073709551615
    if (true || (1 >= -988)) {
        v4 = (("hello" + v4) + (v4 + v4))
        if ("\n\t" == "a") {
            v3 = append(v3, "foo bar")
            v3 = append(v3, "hello")
            for fi38 := 0; fi38 < 11; fi38++ {
                var v39 int = (fi38 ^ fi38)
                v40 := GenS4{F0: v1, F1: -2147483648, F2: 0}
                v5 = v5
                _ = fi38
                _ = v39
                _ = v40
            }
        }
    }
    for fi41 := 0; fi41 < 13; fi41++ {
        fmt.Println(v2)
        v42 := GenS1{F0: fi41, F1: 233}
        _ = fi41
        _ = v42
    }
    fmt.Println(v4, v4, v5)
    fmt.Println(v2, v4)
    for fi43 := 0; fi43 < 5; fi43++ {
        var v44 uint16 = 32767
        _ = fi43
        _ = v44
    }
    v5 -= v5
    v45 := GenS1{F0: 4294967295, F1: 0}
    fmt.Println(v5)
    if false {
        v4 = ""
        if ((float32(0.0) - v2) < 0.0) {
            v1 = v1
            v3 = append(v3, "a")
        } else {
            switch -706 {
            case 127:
                var v46 int64 = -518
                _ = v46
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        for fi47 := 0; fi47 < 2; fi47++ {
            var v48 float64 = 1e10
            var v49 uint64 = v5
            var v50 uint64 = 4294967295
            _ = fi47
            _ = v48
            _ = v49
            _ = v50
        }
    }
    if (v2 <= float32(1.5)) {
        for fi51 := 0; fi51 < 16; fi51++ {
            for fi52 := 0; fi52 < 15; fi52++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v53 := GenS3{F0: genFn1(), F1: (fi52 + -9223372036854775808), F2: v4}
                if false {
                    v1 -= v1
                    var v54 uint64 = v5
                    var v55 int8 = 0
                    _ = v54
                    _ = v55
                } else {
                    var v56 uint64 = v5
                    _ = v56
                }
                _ = fi52
                _ = v53
            }
            var v57 int = fi51
            v2 = float32((1.5 + 0.1))
            _ = fi51
            _ = v57
        }
    }
    if (0 == 127) {
        if (v5 > (v5 % ((4294967295) | 1))) {
            switch v4 {
            case "a":
                for fi58 := 0; fi58 < 16; fi58++ {
                    if (((!(1 > 42)) || (((false || true) || (0 >= 1)) || (32767 <= 0))) && true) {
                        if (true && (42 != 1)) {
                            var v59 int64 = 1
                            fmt.Println(v1)
                            for fi60 := 0; fi60 < 12; fi60++ {
                                v61 := GenS5{F0: v1}
                                var v62 uint = 255
                                v1 = (v1 | v1)
                                v4 = v4
                                _ = len(v3)
                                if true {
                                    var v63 int16 = v1
                                    _ = v63
                                } else {
                                    v1 += (-(v1))
                                }
                                if false {
                                    v2 += float32(1.5)
                                }
                                v4 = v4
                                if (!(!(127 == 127))) {
                                    var v64 int64 = v59
                                    if true {
                                        v65 := GenS4{F0: v1, F1: -1, F2: (fi60 | 0)}
                                        _ = v65
                                    }
                                    switch "hello" {
                                    case "\n\t":
                                        v64 = (v64 / ((v59) | 1))
                                    }
                                    _ = v64
                                } else {
                                    v59 -= v59
                                }
                                _ = fi60
                                _ = v61
                                _ = v62
                            }
                            _ = v59
                        }
                        v66 := GenS1{F0: (fi58 >> ((255) & 63)), F1: 912}
                        _ = v66
                    } else {
                        switch ("foo bar" + "a") {
                        case "hello":
                            if false {
                                v2 += ((((float32(-2.25) - v2) * (1.5 * v2)) - 0.1) - -2.25)
                            } else {
                                for fi67 := 0; fi67 < 13; fi67++ {
                                    if len(v3) > 0 {
                                        v3[0] = (((v4 + v4) + "foo bar") + (("foo bar" + (v4 + v4)) + (("\n\t" + v4) + (v4 + "hello"))))
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = v4
                                    }
                                    if (!(!(0.0 > v2))) {
                                        if false {
                                            var v68 int = fi58
                                            v4 = "a"
                                            _ = v68
                                        } else {
                                            switch v4 {
                                            case "hello":
                                                if len(v3) > 0 {
                                                    v3[0] = v4
                                                }
                                            case "foo bar":
                                                v2 = v2
                                            case "":
                                                v69 := GenS6{F0: 4294967295, F1: -477}
                                                _ = v69
                                            }
                                            v70 := GenS7{F0: 1, F1: (-817 > 1), F2: (fi67 ^ fi67)}
                                            _ = v70
                                        }
                                        v5 = 0
                                        v3 = append(v3, v4)
                                    }
                                    v4 = ("a" + ((((v4 + ("hello" + "foo bar")) + ("\n\t" + ((v4 + "foo bar") + v4))) + "foo bar") + ((((("a" + v4) + v4) + ("\n\t" + (v4 + (v4 + "\n\t")))) + v4) + "foo bar")))
                                    _ = len(v3)
                                    _ = fi67
                                }
                            }
                        case "\n\t":
                            var v71 uint64 = genFn0(-9223372036854775808)
                            _ = v71
                        case "a":
                            v2 = v2
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        for fi72 := 0; fi72 < 17; fi72++ {
                            v5 = (v5 ^ (v5 ^ v5))
                            _ = fi72
                        }
                        if len(v3) > 0 {
                            v3[0] = v4
                        }
                    }
                    for fi73 := 0; fi73 < 20; fi73++ {
                        switch v4 {
                        case "foo bar":
                            _ = len(v3)
                        case "hello":
                            var v74 int16 = v1
                            _ = v74
                        case "a":
                            var v75 string = ("\n\t" + "foo bar")
                            _ = v75
                        }
                        var v76 uint16 = 699
                        if (0.0 <= -2.25) {
                            if (((0.1 * 3.14159) > (0.1 + (float64(v2) / -2.25))) || (990 == 0)) {
                                switch v4 {
                                case "":
                                    var v77 int16 = (v1 | (-(v1)))
                                    _ = v77
                                case "\n\t":
                                    fmt.Println(v76, fi73)
                                case "foo bar":
                                    fmt.Println(fi58, v1)
                                }
                            }
                        }
                        _ = fi73
                        _ = v76
                    }
                    _ = fi58
                }
            case "":
                if len(v3) > 0 {
                    v3[0] = v4
                }
            default:
                var v78 int16 = v1
                _ = v78
            }
            var v79 int32 = 1
            _ = v79
        }
        v3 = append(v3, v4)
        var v80 int16 = (v1 ^ v1)
        _ = v80
    }
    _ = len(v3)
    if false {
        v81 := GenS8{F0: v5, F1: 2}
        _ = v81
    } else {
        for fi82 := 0; fi82 < 11; fi82++ {
            var v83 uint8 = 1
            for fi84 := 0; fi84 < 13; fi84++ {
                switch fi82 {
                case -1:
                    v4 = (v4 + v4)
                case 2147483647:
                    v83 = v83
                default:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                _ = fi84
            }
            v3 = append(v3, "foo bar")
            _ = fi82
            _ = v83
        }
    }
    var v85 int16 = (v1 * (v1 >> ((4294967295) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v85
}
