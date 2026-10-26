package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint; F2 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 uint16; F1 uint64; F2 uint8 }
type GenS6 struct { F0 uint; F1 int }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 uint64; F1 int16; F2 int32 }
func genFn0(p0 int64, p1 uint, p2 uint32) uint32 {
    p2 = (^(p2))
    v7 := GenS1{F0: -943}
    p2 -= (p2 | 65535)
    v8 := GenS2{F0: 1, F1: p1, F2: 42}
    var v9 bool = (!true)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return p2
}
func genFn1(p0 uint64) int16 {
    for fi10 := 0; fi10 < 10; fi10++ {
        v11 := []int{fi10, fi10, (-(fi10))}
        _ = fi10
        _ = v11
    }
    v12 := []int{42}
    p0 += (p0 ^ 315)
    if (2 < 920) {
        for fi13 := 0; fi13 < 10; fi13++ {
            if len(v12) > 0 {
                v12[0] = fi13
            }
            for fi14 := 0; fi14 < 3; fi14++ {
                v15 := GenS1{F0: 42}
                p0 = p0
                v16 := GenS2{F0: 1, F1: 691, F2: int8(p0)}
                var v17 int16 = -406
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
            }
            _ = fi13
        }
        var v18 float64 = 1.5
        p0 = 18446744073709551615
        _ = v18
    } else {
        if false {
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
            p0 += p0
            for fi19 := 0; fi19 < 12; fi19++ {
                var v20 string = "hello"
                p0 = (p0 << ((215) & 63))
                _ = fi19
                _ = v20
            }
        }
        if len(v12) > 0 {
            v12[0] = int(p0)
        }
        var v21 int32 = 42
        _ = v21
    }
    _ = p0
    _ = v12
    return -544
}

func main() {
    v1 := []string{"\n\t"}
    v2 := []uint32{2147483647}
    var v3 int8 = 42
    v4 := []string{"foo bar", "\n\t"}
    var v5 uint32 = 965
    v6 := []uint32{(^(v5))}
    v22 := GenS1{F0: -32768}
    fmt.Println(v3, v5)
    switch "a" {
    case "a":
        var v23 int = 9223372036854775807
        _ = v23
    case "\n\t":
        for fi24 := 0; fi24 < 2; fi24++ {
            fmt.Println(fi24, v3, fi24)
            v5 = genFn0(884, 42, (v5 | (v5 % (((v5 ^ 2147483647)) | 1))))
            _ = fi24
        }
    }
    v3 = 104
    if (-1 > 0) {
        var v25 uint16 = uint16(v5)
        var v26 int64 = 513
        fmt.Println(v25, v3)
        _ = v25
        _ = v26
    } else {
        v27 := GenS2{F0: 42, F1: 255, F2: v3}
        _ = v27
    }
    var v28 uint32 = (v5 | v5)
    if ((v3 >> ((0) & 63)) != 100) {
        v5 = (v28 << ((65535) & 63))
        fmt.Println(v5, v5)
    } else {
        switch "\n\t" {
        case "hello":
            if (1e10 < 0.0) {
                if (1 <= 953) {
                    v28 -= v28
                }
                if (128 > 0) {
                    if true {
                        if ((76 != -954) && false) {
                            var v29 int16 = 32767
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = v29
                        }
                        if (85 < 127) {
                            v30 := GenS1{F0: 1000}
                            var v31 int8 = v3
                            v31 -= (v31 ^ -1)
                            _ = v30
                            _ = v31
                        }
                        var v32 int16 = 1000
                        _ = v32
                    } else {
                        var v33 float32 = (3.14159 - 1e10)
                        for fi34 := 0; fi34 < 17; fi34++ {
                            switch -2147483648 {
                            case 9223372036854775807:
                                for fi35 := 0; fi35 < 3; fi35++ {
                                    var v36 uint32 = genFn0(0, 1, (^(v28)))
                                    if (!false) {
                                        var v37 uint64 = 7
                                        if (42 > 0) {
                                            var v38 string = "a"
                                            fmt.Println(v33, v33, fi35)
                                            var v39 bool = true
                                            _ = v38
                                            _ = v39
                                        }
                                        v37 += 18446744073709551615
                                        v33 = float32(-2.25)
                                        var v40 uint = 509
                                        var v41 bool = (false || false)
                                        _ = v37
                                        _ = v40
                                        _ = v41
                                    } else {
                                        switch (fi35 - (^(fi34))) {
                                        case -640:
                                            v3 = (-(v3))
                                        case 0:
                                            if len(v1) > 0 {
                                                v1[0] = "\n\t"
                                            }
                                        default:
                                            fmt.Println(v3, fi35)
                                        }
                                    }
                                    switch (fi34 % ((32768) | 1)) {
                                    case 489:
                                        if true {
                                            v33 = float32((((0.1 + 0.1) / -2.25) / 1.5))
                                            switch (^(fi35)) {
                                            case 0:
                                                _ = len(v4)
                                                if true {
                                                    if (false || (585 < 0)) {
                                                        var v42 uint = 255
                                                        _ = v42
                                                    }
                                                    if (18446744073709551615 < 42) {
                                                        v43 := GenS3{F0: 255}
                                                        _ = len(v4)
                                                        _ = v43
                                                    }
                                                }
                                                fmt.Println(v36)
                                            case -616:
                                                if len(v4) > 0 {
                                                    v4[0] = ""
                                                }
                                            }
                                        } else {
                                            v44 := GenS4{F0: fi35, F1: int16(v28)}
                                            if false {
                                                var v45 uint = 65535
                                                var v46 bool = (-25 <= -9223372036854775808)
                                                v47 := GenS5{F0: 583, F1: 473, F2: 79}
                                                if false {
                                                    v48 := GenS5{F0: 255, F1: 0, F2: 103}
                                                    if (!v46) {
                                                        var v49 uint16 = 551
                                                        var v50 int16 = genFn1(18446744073709551615)
                                                        _ = v49
                                                        _ = v50
                                                    }
                                                    fmt.Println(v33, v5)
                                                    if ((32767 == 32767) && (!true)) {
                                                        v3 = v3
                                                        fmt.Println(fi34, v36)
                                                        var v51 int16 = -285
                                                        _ = v51
                                                    }
                                                    switch 4294967295 {
                                                    case 127:
                                                        var v52 int = 0
                                                        _ = v52
                                                    case 481:
                                                        var v53 int8 = (v3 | v3)
                                                        _ = v53
                                                    case 32767:
                                                        v33 += v33
                                                        if true {
                                                            _ = len(v6)
                                                            fmt.Println(fi34, v3, v45)
                                                        }
                                                        var v54 float32 = (3.14159 + (v33 / float32(3.14159)))
                                                        var v55 uint8 = 127
                                                        var v56 float32 = -2.25
                                                        _ = v54
                                                        _ = v55
                                                        _ = v56
                                                    }
                                                    _ = v48
                                                }
                                                v57 := GenS6{F0: (v45 >> (((v45 << ((v45) & 63))) & 63)), F1: fi35}
                                                if (false || v46) {
                                                    fmt.Println(fi35, v5, fi35)
                                                    switch "a" {
                                                    case "foo bar":
                                                        if v46 {
                                                            if v46 {
                                                                var v58 float32 = 1e10
                                                                v36 = (v28 & 2)
                                                                _ = v58
                                                            }
                                                        }
                                                    case "\n\t":
                                                        var v59 int16 = 42
                                                        _ = v59
                                                    }
                                                } else {
                                                    v4 = append(v4, ("a" + "\n\t"))
                                                    fmt.Println(fi34, fi34)
                                                }
                                                v5 = (v28 + genFn0(0, v45, 25))
                                                if true {
                                                    v33 -= 0.1
                                                } else {
                                                    var v60 uint = 384
                                                    fmt.Println(fi34, fi35, fi35)
                                                    fmt.Println(fi35)
                                                    _ = v60
                                                }
                                                _ = v45
                                                _ = v46
                                                _ = v47
                                                _ = v57
                                            } else {
                                                var v61 uint8 = 128
                                                _ = len(v6)
                                                if (("hello" + "\n\t") == ("foo bar" + "hello")) {
                                                    v62 := GenS5{F0: uint16(v36), F1: 0, F2: (v61 / (((v61 / ((1) | 1))) | 1))}
                                                    v63 := GenS2{F0: 2147483647, F1: 65535, F2: (v3 + 106)}
                                                    var v64 uint64 = 33
                                                    v64 = uint64(v61)
                                                    _ = v62
                                                    _ = v63
                                                    _ = v64
                                                }
                                                _ = v61
                                            }
                                            if len(v6) > 0 {
                                                v6[0] = (v36 % ((1) | 1))
                                            }
                                            _ = v44
                                        }
                                    case -186:
                                        v33 += v33
                                    }
                                    _ = fi35
                                    _ = v36
                                }
                            case 2147483647:
                                var v65 uint16 = 1
                                _ = v65
                            }
                            var v66 uint32 = (v28 - (v28 - v28))
                            _ = len(v4)
                            _ = fi34
                            _ = v66
                        }
                        if (true || false) {
                            var v67 int32 = 1
                            fmt.Println(v5, v3, v67)
                            _ = len(v6)
                            _ = v67
                        } else {
                            if len(v4) > 0 {
                                v4[0] = (("\n\t" + ("hello" + "a")) + "a")
                            }
                        }
                        _ = v33
                    }
                }
                v5 -= genFn0(0, 42, v5)
            } else {
                var v68 uint = 65535
                v69 := GenS7{F0: -661}
                if false {
                    v70 := GenS8{F0: 904}
                    _ = v70
                } else {
                    fmt.Println(v3, v5, v5)
                    v71 := GenS9{F0: -1}
                    _ = v71
                }
                _ = v68
                _ = v69
            }
        case "\n\t":
            if false {
                if (!(((((1.5 + -2.25) / -2.25) * 1e10) - 1e10) == -2.25)) {
                    v4 = append(v4, "hello")
                    if true {
                        v72 := GenS10{F0: 0}
                        _ = v72
                    }
                    _ = len(v4)
                }
            } else {
                if ((1 < 1) && (!true)) {
                    v73 := GenS2{F0: -1, F1: 647, F2: (v3 - v3)}
                    _ = v73
                } else {
                    var v74 bool = (1.5 < 1.5)
                    var v75 uint = 2
                    if (2 >= 94) {
                        for fi76 := 0; fi76 < 11; fi76++ {
                            v28 += 805
                            for fi77 := 0; fi77 < 4; fi77++ {
                                fmt.Println(fi77)
                                _ = fi77
                            }
                            _ = fi76
                        }
                    } else {
                        switch 9223372036854775807 {
                        case 32767:
                            v28 = (v5 | v5)
                        default:
                            v74 = (v75 != v75)
                        }
                        for fi78 := 0; fi78 < 17; fi78++ {
                            if v74 {
                                var v79 int = fi78
                                var v80 string = (("" + "hello") + "hello")
                                _ = v79
                                _ = v80
                            } else {
                                switch ("foo bar" + "") {
                                case "foo bar":
                                    v5 += v5
                                default:
                                    for fi81 := 0; fi81 < 12; fi81++ {
                                        fmt.Println(v75)
                                        v82 := GenS3{F0: 97}
                                        v3 += (v3 / ((-128) | 1))
                                        if len(v4) > 0 {
                                            v4[0] = "foo bar"
                                        }
                                        var v83 bool = false
                                        _ = fi81
                                        _ = v82
                                        _ = v83
                                    }
                                }
                            }
                            var v84 int64 = -9223372036854775808
                            _ = fi78
                            _ = v84
                        }
                    }
                    _ = v74
                    _ = v75
                }
                v85 := GenS11{F0: 2}
                if len(v4) > 0 {
                    v4[0] = ("foo bar" + "\n\t")
                }
                _ = v85
            }
        default:
            if true {
                var v86 uint64 = 962
                _ = v86
            }
        }
        var v87 uint64 = 1
        v28 = (v5 / ((genFn0(int64(v28), 0, genFn0(0, uint(v28), genFn0(2, 65535, (v5 ^ (v5 / (((v5 << ((2) & 63))) | 1))))))) | 1))
        _ = v87
    }
    var v88 float64 = (0.0 / -2.25)
    var v89 int = -1000
    var v90 int8 = 1
    v88 -= v88
    v5 = (v5 | v28)
    for fi91 := 0; fi91 < 19; fi91++ {
        v92 := GenS5{F0: 32767, F1: 42, F2: 1}
        if true {
            fmt.Println(fi91, v90)
            if false {
                var v93 uint64 = 14
                var v94 int = 32768
                _ = v93
                _ = v94
            }
            for fi95 := 0; fi95 < 15; fi95++ {
                if ((float64(-2.25) - 0.0) < 3.14159) {
                    v5 = 799
                    v96 := GenS12{F0: 18446744073709551615, F1: genFn1(4294967295), F2: 1}
                    _ = v96
                } else {
                    v97 := GenS8{F0: 0}
                    _ = v97
                }
                _ = fi95
            }
        }
        _ = fi91
        _ = v92
    }
    var v98 uint64 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v2
    _ = v22
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v88
    _ = v89
    _ = v90
    _ = v98
}
