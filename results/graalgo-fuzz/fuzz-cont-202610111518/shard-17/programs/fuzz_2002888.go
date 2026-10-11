package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 int; F1 int64; F2 string }
type GenS5 struct { F0 int16; F1 int64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int; F1 uint64; F2 int32 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int8; F1 uint8; F2 uint32 }
type GenS10 struct { F0 int64; F1 int16 }
func genFn0(p0 int8, p1 uint64) uint {
    v4 := []int{-9223372036854775808, 127, 32768}
    for fi5 := 0; fi5 < 10; fi5++ {
        if len(v4) > 0 {
            v4[0] = (fi5 * 184)
        }
        _ = fi5
    }
    var v6 string = ("\n\t" + "foo bar")
    if false {
        fmt.Println(p1)
        _ = len(v4)
    }
    var v7 int = 631
    _ = p0
    _ = p1
    _ = v4
    _ = v6
    _ = v7
    return 255
}
func genFn1(p0 uint8, p1 int32, p2 string) bool {
    var v8 bool = (false && false)
    v9 := GenS1{F0: 0, F1: 2}
    switch "\n\t" {
    case "hello":
        var v10 float64 = 3.14159
        _ = v10
    case "a":
        fmt.Println(p2, v8, p0)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return v8
}

func main() {
    v1 := make(map[string]string)
    var v2 int16 = 174
    v3 := []uint32{42, 368, 873}
    v11 := GenS1{F0: 42, F1: 2}
    v2 -= -32768
    for fi12 := 0; fi12 < 2; fi12++ {
        var v13 string = "foo bar"
        for fi14 := 0; fi14 < 15; fi14++ {
            var v15 int64 = 42
            var v16 int8 = 100
            fmt.Println(v2, v13)
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi12
        _ = v13
    }
    var v17 int32 = -1
    for fi18 := 0; fi18 < 10; fi18++ {
        fmt.Println(v2, fi18)
        fmt.Println(v2)
        v19 := GenS2{F0: 1, F1: 1}
        _ = fi18
        _ = v19
    }
    switch "\n\t" {
    case "foo bar":
        if (false || (0 > 128)) {
            v20 := GenS1{F0: -128, F1: 255}
            v17 = (v17 - v17)
            v21 := GenS2{F0: 42, F1: 42}
            _ = v20
            _ = v21
        } else {
            v3 = append(v3, 2147483648)
            for fi22 := 0; fi22 < 20; fi22++ {
                var v23 int8 = 100
                var v24 int16 = 1000
                _ = fi22
                _ = v23
                _ = v24
            }
        }
    case "":
        v25 := GenS2{F0: 0, F1: 76}
        _ = v25
    default:
        fmt.Println(v17, v2)
    }
    if true {
        if ((false || ((!(!(958 > 244))) || true)) || ((float64(float32(3.14159)) / float64(1e10)) == (1e10 * -2.25))) {
            if false {
                var v26 float32 = (0.1 - 0.0)
                v27 := GenS3{F0: 2147483648, F1: (v2 ^ 32767)}
                for fi28 := 0; fi28 < 6; fi28++ {
                    _ = len(v1)
                    v29 := GenS2{F0: 0, F1: 255}
                    _ = fi28
                    _ = v29
                }
                _ = v26
                _ = v27
            }
            var v30 uint32 = 65535
            _ = v30
        } else {
            if false {
                var v31 int16 = (v2 >> ((692) & 63))
                v32 := GenS4{F0: -632, F1: 137, F2: ("\n\t" + "foo bar")}
                _ = v31
                _ = v32
            } else {
                v2 = v2
                fmt.Println(v17, v17)
                var v33 int = -595
                _ = v33
            }
            fmt.Println(v2, v2)
        }
    }
    var v34 uint64 = 18446744073709551615
    v17 = (v17 | v17)
    v34 = 1
    if false {
        v35 := GenS1{F0: 127, F1: genFn0(int8(v34), v34)}
        v36 := GenS1{F0: -75, F1: 0}
        for fi37 := 0; fi37 < 16; fi37++ {
            for fi38 := 0; fi38 < 4; fi38++ {
                switch fi37 {
                case 2147483647:
                    var v39 int64 = -262
                    _ = v39
                case 1:
                    v2 = (v2 << ((4294967295) & 63))
                default:
                    if ((v2 >= v2) || (true || false)) {
                        switch fi37 {
                        case 32767:
                            v17 += 646
                        }
                        if true {
                            var v40 int8 = int8(fi38)
                            var v41 int64 = -249
                            v3 = append(v3, 13)
                            delete(v1, (("\n\t" + "") + "foo bar"))
                            _ = v40
                            _ = v41
                        } else {
                            v42 := GenS5{F0: (v2 ^ (v2 - v2)), F1: int64(fi37)}
                            var v43 string = ("a" + "\n\t")
                            var v44 uint = uint(v34)
                            _ = v42
                            _ = v43
                            _ = v44
                        }
                    } else {
                        var v45 int8 = -6
                        v2 += v2
                        _ = v45
                    }
                }
                var v46 float64 = 1e10
                _ = fi38
                _ = v46
            }
            var v47 uint8 = 42
            delete(v1, "hello")
            _ = fi37
            _ = v47
        }
        _ = v35
        _ = v36
    } else {
        v17 -= int32(v34)
        fmt.Println(v34)
    }
    v17 -= (v17 - (v17 / (((-(v17))) | 1)))
    var v48 uint8 = 1
    if ((-1 < 823) || (!true)) {
        for fi49 := 0; fi49 < 8; fi49++ {
            if ((v48 + 34) <= (v48 * v48)) {
                v34 += v34
            } else {
                v48 = (v48 / ((v48) | 1))
            }
            switch fi49 {
            case 42:
                v48 -= v48
            case -1:
                var v50 uint32 = 42
                _ = v50
            case 4294967295:
                var v51 uint32 = 2
                _ = v51
            }
            if (0 == 42) {
                if true {
                    v34 = v34
                    v48 = v48
                }
                if (26 < 1) {
                    switch 128 {
                    case 9223372036854775807:
                        for fi52 := 0; fi52 < 2; fi52++ {
                            v17 = (v17 ^ v17)
                            _ = fi52
                        }
                    case -321:
                        if true {
                            v48 = (v48 ^ v48)
                            for fi53 := 0; fi53 < 16; fi53++ {
                                switch (-(fi53)) {
                                case -427:
                                    fmt.Println(v2)
                                    var v54 uint16 = 65535
                                    _ = v54
                                case 4294967295:
                                    var v55 float32 = 0.1
                                    var v56 uint16 = 42
                                    _ = v55
                                    _ = v56
                                default:
                                    var v57 int8 = 0
                                    _ = v57
                                }
                                v17 = -380
                                var v58 int16 = v2
                                _ = fi53
                                _ = v58
                            }
                        } else {
                            _ = len(v3)
                            for fi59 := 0; fi59 < 16; fi59++ {
                                if (-2.25 <= 1.5) {
                                    if false {
                                        var v60 uint = 326
                                        if len(v3) > 0 {
                                            v3[0] = 591
                                        }
                                        if (42 <= 621) {
                                            if (false || (!false)) {
                                                var v61 uint32 = 2
                                                v62 := GenS4{F0: 4294967295, F1: 231, F2: "foo bar"}
                                                _ = v61
                                                _ = v62
                                            }
                                            if false {
                                                if true {
                                                    v34 = 2
                                                    fmt.Println(v2)
                                                    _ = len(v3)
                                                }
                                                v60 += genFn0(127, 18446744073709551615)
                                            } else {
                                                var v63 int8 = -10
                                                v64 := GenS4{F0: fi59, F1: 9223372036854775807, F2: "a"}
                                                v60 -= (v60 - 42)
                                                _ = v63
                                                _ = v64
                                            }
                                            v34 = (v34 & v34)
                                        }
                                        _ = v60
                                    }
                                    v65 := GenS2{F0: v34, F1: (v48 - (v48 + v48))}
                                    _ = v65
                                } else {
                                    fmt.Println(fi59, fi59, fi49)
                                    var v66 bool = (true && false)
                                    _ = v66
                                }
                                var v67 int16 = v2
                                v68 := GenS6{F0: ((134 > (v2 * (v67 / ((159) | 1)))) && false)}
                                _ = fi59
                                _ = v67
                                _ = v68
                            }
                        }
                    default:
                        var v69 string = "foo bar"
                        _ = v69
                    }
                }
                for fi70 := 0; fi70 < 17; fi70++ {
                    switch (((("foo bar" + ("foo bar" + "hello")) + "") + "a") + ("foo bar" + "a")) {
                    case "":
                        fmt.Println(v48, v17)
                    case "\n\t":
                        v71 := GenS3{F0: 249, F1: v2}
                        _ = v71
                    case "foo bar":
                        switch (fi49 | -1) {
                        case 128:
                            fmt.Println(v17, v34, fi70)
                        }
                    }
                    _ = fi70
                }
            } else {
                var v72 uint16 = 2
                if true {
                    v73 := GenS1{F0: 42, F1: genFn0(42, (v34 - v34))}
                    v2 += v2
                    for fi74 := 0; fi74 < 6; fi74++ {
                        fmt.Println(fi74)
                        v72 = (v72 % ((42) | 1))
                        if (fi74 == (fi49 % ((fi74) | 1))) {
                            v75 := GenS3{F0: 42, F1: (v2 & v2)}
                            _ = v75
                        }
                        _ = fi74
                    }
                    _ = v73
                } else {
                    switch "foo bar" {
                    case "\n\t":
                        if false {
                            switch fi49 {
                            case 0:
                                switch fi49 {
                                case -1:
                                    v72 += (^(v72))
                                default:
                                    if true {
                                        switch fi49 {
                                        case 9223372036854775807:
                                            v76 := GenS7{F0: fi49, F1: v34, F2: (v17 - v17)}
                                            _ = v76
                                        case 227:
                                            v34 += v34
                                        }
                                        for fi77 := 0; fi77 < 4; fi77++ {
                                            if true {
                                                v78 := GenS8{F0: (v2 ^ 155)}
                                                var v79 uint16 = (v72 & (v72 | uint16(v48)))
                                                _ = v78
                                                _ = v79
                                            }
                                            _ = fi77
                                        }
                                    }
                                }
                            case 42:
                                v80 := GenS9{F0: 1, F1: (v48 >> ((uint(v34)) & 63)), F2: uint32(v48)}
                                _ = v80
                            case -518:
                                var v81 string = (("hello" + "\n\t") + "hello")
                                _ = v81
                            default:
                                switch (fi49 % (((fi49 - 128)) | 1)) {
                                case 895:
                                    v72 = v72
                                case -1:
                                    v82 := GenS7{F0: (fi49 + int(v2)), F1: v34, F2: 786}
                                    _ = v82
                                case 9223372036854775807:
                                    if ((v34 | 0) > 42) {
                                        v83 := GenS10{F0: 714, F1: (v2 ^ v2)}
                                        _ = v83
                                    } else {
                                        v34 = (v34 * v34)
                                        if (1 <= 2147483647) {
                                            v84 := v1["\n\t"]
                                            _ = v84
                                        } else {
                                            v85 := GenS10{F0: -9223372036854775808, F1: v2}
                                            _ = v85
                                        }
                                    }
                                default:
                                    v86 := GenS8{F0: 0}
                                    _ = v86
                                }
                            }
                        } else {
                            fmt.Println(v2)
                            v17 = 42
                        }
                    case "foo bar":
                        if len(v3) > 0 {
                            v3[0] = 744
                        }
                    default:
                        for fi87 := 0; fi87 < 11; fi87++ {
                            v72 = v72
                            if false {
                                var v88 int8 = -50
                                v17 = (^(v17))
                                v89 := GenS10{F0: -9223372036854775808, F1: (-(v2))}
                                _ = v88
                                _ = v89
                            } else {
                                v48 = (^(v48))
                                v17 -= 575
                                v90 := GenS7{F0: int(v2), F1: v34, F2: 1}
                                _ = v90
                            }
                            switch ((("" + "") + ("hello" + "foo bar")) + "hello") {
                            case "\n\t":
                                v34 += 42
                            }
                            _ = fi87
                        }
                    }
                    v91 := v1["foo bar"]
                    v17 = (v17 << ((65535) & 63))
                    _ = v91
                }
                _ = v72
            }
            _ = fi49
        }
    }
    v48 = v48
    v34 += (v34 | 394)
    var v92 float64 = -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v48:", v48)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v48
    _ = v92
}
