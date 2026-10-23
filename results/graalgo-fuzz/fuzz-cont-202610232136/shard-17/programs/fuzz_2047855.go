package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 uint64; F2 uint16 }
type GenS3 struct { F0 int32; F1 string; F2 int64 }
type GenS4 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int; F1 string }
type GenS8 struct { F0 uint16; F1 uint }
type GenS9 struct { F0 string }
func genFn0(p0 int16) uint64 {
    var v7 int64 = 32
    for fi8 := 0; fi8 < 20; fi8++ {
        if true {
            p0 += p0
            v7 += -88
        } else {
            var v9 uint = 65535
            fmt.Println(p0, p0, v7)
            v10 := []int{}
            _ = v10
            _ = v9
        }
        _ = fi8
    }
    _ = p0
    _ = v7
    return 321
}
func genFn1(p0 uint32) int8 {
    v11 := []int{-756, 187, 42}
    var v12 float32 = 1.5
    if false {
        var v13 int = -719
        var v14 uint64 = 18446744073709551615
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 int8 = 127
        _ = fi15
        _ = v16
    }
    if false {
        v11 = append(v11, 42)
        v17 := GenS1{F0: 255}
        switch (("hello" + "") + "foo bar") {
        case "hello":
            v12 = 3.14159
        case "\n\t":
            for fi18 := 0; fi18 < 16; fi18++ {
                if false {
                    var v19 int16 = 1
                    var v20 string = "a"
                    fmt.Println(v19, v20)
                    _ = v19
                    _ = v20
                } else {
                    fmt.Println(fi18, p0)
                    var v21 int8 = 127
                    var v22 string = ("" + "\n\t")
                    _ = v21
                    _ = v22
                }
                for fi23 := 0; fi23 < 17; fi23++ {
                    if ((-32768 == 1) && true) {
                        p0 = 741
                        var v24 bool = ((p0 + (^(p0))) >= (p0 % (((^(p0))) | 1)))
                        v24 = v24
                        _ = v24
                    } else {
                        p0 = p0
                        if ((((fi23 % (((fi23 >> ((255) & 63))) | 1)) <= (fi18 * 32768)) && true) && false) {
                            if ("" < "") {
                                if ((-1 >= 221) && false) {
                                    switch "hello" {
                                    case "hello":
                                        v25 := GenS1{F0: 535}
                                        _ = v25
                                    default:
                                        v12 -= (((v12 * float32(0.1)) - v12) * ((float32(1e10) + 1e10) - (v12 - (1.5 + 0.1))))
                                    }
                                    switch fi23 {
                                    case 2147483647:
                                        if (true || true) {
                                            if len(v11) > 0 {
                                                v11[0] = 32767
                                            }
                                            if true {
                                                switch fi18 {
                                                case 689:
                                                    v12 = v12
                                                }
                                            } else {
                                                v12 += v12
                                                v26 := GenS1{F0: 1}
                                                var v27 float32 = ((-2.25 * -2.25) * (0.0 + v12))
                                                var v28 float32 = v12
                                                v29 := GenS2{F0: "foo bar", F1: 329, F2: 358}
                                                v27 -= v27
                                                var v30 uint16 = 118
                                                _ = v26
                                                _ = v27
                                                _ = v28
                                                _ = v29
                                                _ = v30
                                            }
                                            var v31 uint64 = genFn0(0)
                                            _ = v31
                                        }
                                    }
                                    v12 += 0.1
                                }
                                p0 = (^(p0))
                            } else {
                                v32 := GenS3{F0: 0, F1: "hello", F2: -1}
                                _ = v32
                            }
                            var v33 uint = 255
                            _ = v33
                        }
                        var v34 uint = 4294967295
                        v35 := GenS2{F0: "foo bar", F1: 642, F2: 787}
                        if len(v11) > 0 {
                            v11[0] = (fi23 / ((fi23) | 1))
                        }
                        _ = v34
                        _ = v35
                    }
                    _ = fi23
                }
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                _ = fi18
            }
        case "foo bar":
            if false {
                p0 -= p0
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                var v36 bool = (-9223372036854775808 == 0)
                _ = v36
            }
        }
        _ = v17
    }
    _ = p0
    _ = v11
    _ = v12
    return 92
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 42
    var v3 uint8 = v2
    v4 := make(map[uint32]bool)
    var v5 bool = true
    v6 := []int{2147483647}
    v37 := GenS4{F0: 567, F1: true, F2: genFn0(-598)}
    var v38 int64 = 347
    if (!v5) {
        switch 128 {
        case 15:
            var v39 int64 = (v38 | -1)
            _ = v39
        case 4294967295:
            if (!true) {
                switch 670 {
                case -2147483648:
                    v38 -= 42
                }
                var v40 float32 = float32(-2.25)
                _ = v40
            } else {
                switch -1 {
                case -1:
                    var v41 float64 = (0.0 / 1.5)
                    _ = v41
                }
                var v42 int8 = genFn1(2147483648)
                if true {
                    _ = len(v1)
                    v42 += 100
                    v5 = v5
                }
                _ = v42
            }
        }
        v43 := v4[2147483647]
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = v43
    } else {
        fmt.Println(v2, v38)
    }
    v44 := GenS5{F0: 0}
    v2 = (v3 - v3)
    switch -2147483648 {
    case 603:
        switch ("hello" + "") {
        case "hello":
            v45 := GenS2{F0: "", F1: 0, F2: 2}
            _ = v45
        case "":
            if false {
                var v46 float32 = ((0.0 / 1.5) + 0.0)
                if (!((("a" + "a") + ("\n\t" + "foo bar")) >= ("foo bar" + "a"))) {
                    v2 = (v3 & 180)
                    for fi47 := 0; fi47 < 4; fi47++ {
                        for fi48 := 0; fi48 < 13; fi48++ {
                            var v49 float64 = ((((1e10 / (1e10 / 3.14159)) + 1e10) / 0.1) / 3.14159)
                            v50 := GenS6{F0: (v38 >> ((65535) & 63))}
                            _ = fi48
                            _ = v49
                            _ = v50
                        }
                        _ = fi47
                    }
                    if (("hello" + (("a" + (((("" + ("" + "hello")) + ("hello" + "")) + "\n\t") + "")) + ("hello" + "a"))) != "\n\t") {
                        if (!true) {
                            for fi51 := 0; fi51 < 13; fi51++ {
                                var v52 bool = (255 == 12)
                                _ = fi51
                                _ = v52
                            }
                            for fi53 := 0; fi53 < 12; fi53++ {
                                switch "a" {
                                case "\n\t":
                                    for fi54 := 0; fi54 < 14; fi54++ {
                                        v55 := GenS5{F0: 65535}
                                        _ = fi54
                                        _ = v55
                                    }
                                default:
                                    if (!(!(!(((!(-40 < -128)) || v5) && v5)))) {
                                        if true {
                                            v5 = v5
                                            var v56 uint8 = v3
                                            _ = v56
                                        }
                                    }
                                }
                                _ = fi53
                            }
                        } else {
                            v57 := GenS7{F0: 127, F1: ""}
                            v2 += (v3 >> ((255) & 63))
                            for fi58 := 0; fi58 < 2; fi58++ {
                                v5 = true
                                var v59 uint32 = 0
                                _ = fi58
                                _ = v59
                            }
                            _ = v57
                        }
                        var v60 uint32 = 4294967295
                        var v61 bool = true
                        _ = v60
                        _ = v61
                    }
                }
                _ = v46
            }
        case "a":
            for fi62 := 0; fi62 < 17; fi62++ {
                var v63 float64 = 3.14159
                if v5 {
                    var v64 uint32 = 2
                    fmt.Println(v5, v38, v2)
                    _ = v64
                } else {
                    v38 = (-(v38))
                    var v65 uint32 = 2147483647
                    _ = v65
                }
                _ = fi62
                _ = v63
            }
        }
    case 2147483647:
        for fi66 := 0; fi66 < 2; fi66++ {
            v38 -= v38
            _ = fi66
        }
    case 42:
        var v67 float32 = -2.25
        _ = v67
    }
    v5 = (v5 && true)
    if true {
        var v68 float32 = ((0.0 - 1e10) + (3.14159 / 1.5))
        var v69 uint8 = v3
        v70 := GenS8{F0: 255, F1: 0}
        _ = v68
        _ = v69
        _ = v70
    } else {
        for fi71 := 0; fi71 < 9; fi71++ {
            v72 := v1[(fi71 * fi71)]
            v2 -= (v2 & v3)
            _ = fi71
            _ = v72
        }
        if (!(9223372036854775807 != 9223372036854775807)) {
            var v73 uint8 = (v2 | v3)
            if true {
                v74 := GenS6{F0: (v38 + -9223372036854775808)}
                _ = v74
            } else {
                v73 -= (v73 << ((678) & 63))
                v5 = (1e10 <= 3.14159)
            }
            _ = v73
        } else {
            for fi75 := 0; fi75 < 7; fi75++ {
                switch "" {
                case "\n\t":
                    v5 = true
                case "foo bar":
                    v5 = (!v5)
                }
                for fi76 := 0; fi76 < 1; fi76++ {
                    var v77 bool = false
                    _ = fi76
                    _ = v77
                }
                fmt.Println(fi75, v2, v38)
                _ = fi75
            }
        }
    }
    var v78 int16 = -32768
    v79 := GenS9{F0: ("\n\t" + "foo bar")}
    v2 += (v2 ^ v2)
    for fi80 := 0; fi80 < 9; fi80++ {
        switch "" {
        case "":
            v81 := GenS7{F0: (fi80 - (fi80 + (-(fi80)))), F1: "hello"}
            _ = v81
        case "\n\t":
            for fi82 := 0; fi82 < 5; fi82++ {
                delete(v4, 2147483647)
                v6 = append(v6, -509)
                _ = fi82
            }
        default:
            for fi83 := 0; fi83 < 10; fi83++ {
                fmt.Println(fi80)
                var v84 uint32 = 2147483648
                _ = fi83
                _ = v84
            }
        }
        v3 = v2
        _ = fi80
    }
    fmt.Println(v38)
    v85 := v1[-9223372036854775808]
    for fi86 := 0; fi86 < 9; fi86++ {
        var v87 uint8 = (^(v3))
        var v88 uint16 = 0
        _ = fi86
        _ = v87
        _ = v88
    }
    for fi89 := 0; fi89 < 20; fi89++ {
        fmt.Println(v3)
        _ = fi89
    }
    v2 = (v2 + 42)
    v38 = v38
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v78:", v78)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v44
    _ = v5
    _ = v6
    _ = v78
    _ = v79
    _ = v85
}
