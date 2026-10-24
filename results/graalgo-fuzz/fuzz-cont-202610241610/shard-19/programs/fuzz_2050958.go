package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 int32; F2 int64 }
type GenS4 struct { F0 uint; F1 int32; F2 int16 }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 uint64; F1 int16; F2 string }

func main() {
    var v1 bool = (false || (2 >= 4294967295))
    var v2 float64 = (0.1 / (3.14159 * 0.1))
    var v3 int64 = -391
    var v4 int32 = -2147483648
    var v5 uint64 = 0
    if true {
        if v1 {
            if (1000 > -791) {
                v6 := GenS1{F0: 1}
                _ = v6
            }
            for fi7 := 0; fi7 < 14; fi7++ {
                var v8 uint32 = 526
                _ = fi7
                _ = v8
            }
            v9 := []int{107, 2147483647, 9223372036854775807, 128}
            _ = v9
        }
        if (!true) {
            for fi10 := 0; fi10 < 4; fi10++ {
                v11 := GenS2{F0: 0}
                if (!v1) {
                    v12 := []int{127, (fi10 & 2147483647), (fi10 / ((-959) | 1))}
                    if v1 {
                        var v13 bool = (701 > 149)
                        v14 := GenS1{F0: v5}
                        v2 -= float64((1e10 - 0.1))
                        _ = v13
                        _ = v14
                    }
                    _ = v12
                }
                _ = fi10
                _ = v11
            }
            fmt.Println(v5, v5, v5)
            v2 += v2
        } else {
            for fi15 := 0; fi15 < 7; fi15++ {
                var v16 int16 = 750
                _ = fi15
                _ = v16
            }
        }
        if (0 <= 42) {
            v3 = (v3 ^ v3)
        }
    } else {
        var v17 uint64 = v5
        _ = v17
    }
    for fi18 := 0; fi18 < 9; fi18++ {
        v4 -= (v4 - 0)
        _ = fi18
    }
    v19 := []int{2147483647, -385, 166}
    v2 += 3.14159
    fmt.Println(v3)
    for fi20 := 0; fi20 < 10; fi20++ {
        v5 -= v5
        _ = fi20
    }
    v5 += v5
    var v21 int64 = v3
    var v22 int32 = v4
    _ = len(v19)
    switch "" {
    case "a":
        var v23 float32 = float32((v2 - float64((1.5 + 3.14159))))
        _ = v23
    case "foo bar":
        for fi24 := 0; fi24 < 13; fi24++ {
            if true {
                v19 = append(v19, fi24)
            } else {
                fmt.Println(v4, v5, fi24)
                v25 := GenS2{F0: v5}
                _ = v25
            }
            if v1 {
                v26 := GenS2{F0: 149}
                if ("\n\t" > "foo bar") {
                    if (!(((1.5 - float32(float64(1.5))) * ((-2.25 / 1.0) - (0.1 - 3.14159))) == 0.0)) {
                        for fi27 := 0; fi27 < 11; fi27++ {
                            fmt.Println(v22)
                            _ = fi27
                        }
                        var v28 float64 = (v2 + (((v2 / v2) - 1.5) * v2))
                        var v29 int = (fi24 / ((32767) | 1))
                        _ = v28
                        _ = v29
                    } else {
                        fmt.Println(v22, fi24)
                        var v30 bool = v1
                        _ = v30
                    }
                    if (2147483647 > 247) {
                        var v31 string = ""
                        v31 = v31
                        fmt.Println(v22, v31, v2)
                        _ = v31
                    }
                    var v32 uint16 = 65535
                    _ = v32
                } else {
                    _ = len(v19)
                }
                _ = v26
            } else {
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                if true {
                    v33 := GenS1{F0: (v5 % (((^(v5))) | 1))}
                    for fi34 := 0; fi34 < 13; fi34++ {
                        if v1 {
                            var v35 int16 = 687
                            if true {
                                _ = len(v19)
                            }
                            var v36 uint16 = 1
                            _ = v35
                            _ = v36
                        }
                        var v37 string = ("foo bar" + ("hello" + "hello"))
                        var v38 uint64 = v5
                        var v39 uint8 = uint8(v38)
                        _ = fi34
                        _ = v37
                        _ = v38
                        _ = v39
                    }
                    _ = v33
                }
            }
            var v40 string = "hello"
            _ = fi24
            _ = v40
        }
    case "hello":
        if true {
            for fi41 := 0; fi41 < 17; fi41++ {
                for fi42 := 0; fi42 < 2; fi42++ {
                    var v43 string = "a"
                    v44 := GenS3{F0: 255, F1: v22, F2: v21}
                    _ = fi42
                    _ = v43
                    _ = v44
                }
                switch "" {
                case "":
                    v45 := GenS4{F0: 255, F1: 42, F2: 428}
                    _ = v45
                case "\n\t":
                    fmt.Println(v4, v21, v5)
                }
                _ = fi41
            }
            v19 = append(v19, int(v5))
        }
    default:
        v5 = v5
    }
    if v1 {
        if false {
            if v1 {
                v3 = v3
                var v46 float64 = v2
                _ = v46
            }
            for fi47 := 0; fi47 < 4; fi47++ {
                var v48 bool = v1
                var v49 float64 = 0.0
                _ = fi47
                _ = v48
                _ = v49
            }
        } else {
            switch ("" + "") {
            case "hello":
                v19 = append(v19, 1)
            case "":
                v19 = append(v19, 0)
            default:
                for fi50 := 0; fi50 < 19; fi50++ {
                    if (2 != 0) {
                        if (36 <= 2) {
                            var v51 float64 = -2.25
                            fmt.Println(v21, v4)
                            if len(v19) >= 2 {
                                v19 = v19[1:len(v19)]
                            }
                            _ = v51
                        }
                        if (v3 > 649) {
                            for fi52 := 0; fi52 < 13; fi52++ {
                                if true {
                                    v22 += (v22 | (v22 ^ -89))
                                } else {
                                    fmt.Println(v3)
                                    switch (fi52 * (fi52 ^ -2147483648)) {
                                    case 32767:
                                        if true {
                                            v53 := GenS1{F0: v5}
                                            v22 += v22
                                            switch fi52 {
                                            case 32768:
                                                _ = len(v19)
                                            }
                                            _ = v53
                                        }
                                    case -1:
                                        fmt.Println(v1, v1)
                                    case -2147483648:
                                        var v54 int8 = 100
                                        var v55 int16 = int16(v21)
                                        if (("foo bar" + ("foo bar" + "hello")) < (((("hello" + "foo bar") + "a") + "\n\t") + "")) {
                                            if v1 {
                                                if (("\n\t" + ("\n\t" + (((("a" + "") + ("foo bar" + "foo bar")) + "hello") + "a"))) > "a") {
                                                    var v56 uint32 = 65535
                                                    v21 = (v3 / ((v21) | 1))
                                                    _ = v56
                                                } else {
                                                    if (v1 && (!true)) {
                                                        if (((v1 || v1) || (v1 || (v5 >= 42))) || (((fi52 % ((fi52) | 1)) >= int(v54)) || ((v1 && (float64((1.5 * 1.5)) == (float64(1e10) * (v2 * (v2 + float64(1e10)))))) || (2 <= 730)))) {
                                                            if v1 {
                                                                v1 = v1
                                                                v57 := GenS3{F0: 757, F1: v4, F2: (^(v3))}
                                                                v2 = v2
                                                                _ = v57
                                                            } else {
                                                                v58 := GenS3{F0: 1, F1: int32(v5), F2: 527}
                                                                _ = v58
                                                            }
                                                            v21 = (v21 % (((v21 | (-(v3)))) | 1))
                                                        } else {
                                                            var v59 int64 = (^(v21))
                                                            var v60 uint16 = 530
                                                            _ = v59
                                                            _ = v60
                                                        }
                                                    }
                                                }
                                                v61 := GenS5{F0: 56, F1: 2}
                                                var v62 float64 = (v2 - v2)
                                                _ = v61
                                                _ = v62
                                            }
                                        } else {
                                            _ = len(v19)
                                        }
                                        _ = v54
                                        _ = v55
                                    }
                                    v4 = (v22 / ((-2147483648) | 1))
                                }
                                _ = fi52
                            }
                            for fi63 := 0; fi63 < 13; fi63++ {
                                switch (fi63 ^ fi50) {
                                case 509:
                                    v21 += 129
                                case 9223372036854775807:
                                    if false {
                                        if (!(!(42 > 2))) {
                                            fmt.Println(v1, v1, v2)
                                        } else {
                                            v64 := GenS6{F0: v5, F1: -570, F2: (("a" + "") + "hello")}
                                            _ = len(v19)
                                            _ = v64
                                        }
                                        fmt.Println(v4)
                                        v2 += v2
                                        if len(v19) >= 2 {
                                            v19 = v19[1:len(v19)]
                                        }
                                        v5 -= v5
                                    } else {
                                        switch fi50 {
                                        case -40:
                                            var v65 int8 = -128
                                            _ = v65
                                        case -370:
                                            var v66 uint = uint(v5)
                                            _ = v66
                                        }
                                    }
                                }
                                var v67 float64 = v2
                                _ = fi63
                                _ = v67
                            }
                        } else {
                            switch (fi50 * 519) {
                            case -2147483648:
                                v21 -= (v3 % ((-253) | 1))
                            case 128:
                                var v68 string = (("\n\t" + "\n\t") + ("a" + "a"))
                                _ = v68
                            }
                            var v69 int32 = v22
                            v70 := GenS5{F0: 0, F1: 65535}
                            _ = v69
                            _ = v70
                        }
                    }
                    v71 := GenS4{F0: 7, F1: (^(v4)), F2: -32768}
                    if v1 {
                        v2 += v2
                        if len(v19) >= 2 {
                            v19 = v19[1:len(v19)]
                        }
                    }
                    _ = fi50
                    _ = v71
                }
            }
        }
        v5 = (v5 ^ v5)
    } else {
        var v72 uint16 = 571
        _ = v72
    }
    var v73 float32 = 1e10
    fmt.Println(v4, v1)
    fmt.Println(v5, v1)
    var v74 int32 = v4
    switch 127 {
    case -9223372036854775808:
        for fi75 := 0; fi75 < 9; fi75++ {
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
            v2 -= (v2 * ((v2 / 1e10) * (((((v2 - (v2 - 1.5)) * float64(((((((0.1 - 3.14159) / float32(3.14159)) * 0.0) - (1.5 / v73)) - v73) - float32(v2)))) * v2) * ((0.1 - float64(float32(v2))) + 0.1)) * (float64(1e10) - 0.1))))
            switch "\n\t" {
            case "foo bar":
                v19 = append(v19, fi75)
            case "a":
                if v1 {
                    if (590 >= 65535) {
                        v2 += (v2 - v2)
                    } else {
                        v76 := GenS3{F0: 255, F1: (v4 ^ (-(v4))), F2: (-(v3))}
                        var v77 int32 = -896
                        _ = v76
                        _ = v77
                    }
                    var v78 float32 = v73
                    _ = v78
                }
            case "":
                v5 -= (v5 >> ((255) & 63))
            }
            _ = fi75
        }
    case -596:
        if (true && v1) {
            v79 := GenS3{F0: 18446744073709551615, F1: (-(v4)), F2: v3}
            _ = v79
        }
    default:
        v1 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v73
    _ = v74
}
