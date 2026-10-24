package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 uint32; F2 string }
type GenS5 struct { F0 uint; F1 int32 }
type GenS6 struct { F0 uint16; F1 uint64 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint16; F1 uint }
type GenS10 struct { F0 int64; F1 uint32; F2 string }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 uint16; F1 uint32 }
type GenS14 struct { F0 int32; F1 uint16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = 1
    var v3 string = ("\n\t" + ("" + ("foo bar" + "a")))
    var v4 uint32 = 926
    if false {
        v2 -= (v2 << ((18446744073709551615) & 63))
        fmt.Println(v4, v3)
        if ((418 != 255) || false) {
            fmt.Println(v4, v4)
            fmt.Println(v3, v4)
            v5 := GenS1{F0: "a"}
            _ = v5
        }
    } else {
        var v6 int = 0
        if false {
            v6 = (v6 ^ v6)
        } else {
            _ = len(v1)
            var v7 uint16 = 485
            switch v6 {
            case -845:
                for fi8 := 0; fi8 < 8; fi8++ {
                    v9 := GenS2{F0: -1, F1: v4, F2: 193}
                    _ = fi8
                    _ = v9
                }
            default:
                v10 := GenS3{F0: 255}
                _ = v10
            }
            _ = v7
        }
        if false {
            var v11 int16 = (v2 * (v2 << ((18446744073709551615) & 63)))
            _ = v11
        } else {
            v12 := []int{}
            var v13 float64 = 1.5
            if false {
                _ = len(v1)
                _ = len(v1)
                for fi14 := 0; fi14 < 5; fi14++ {
                    if (!(false || true)) {
                        if true {
                            for fi15 := 0; fi15 < 13; fi15++ {
                                v16 := GenS2{F0: 679, F1: (v4 & (^(v4))), F2: 1}
                                v2 = (v2 * v2)
                                var v17 float32 = ((((3.14159 / 1.0) - float32(3.14159)) / (1.5 + 1e10)) + (1.5 + 1e10))
                                _ = fi15
                                _ = v16
                                _ = v17
                            }
                        } else {
                            fmt.Println(v3, v6)
                            if true {
                                v18 := GenS4{F0: 32767, F1: 42, F2: v3}
                                v13 -= float64((1e10 + ((3.14159 + float32(float64(0.1))) + 1e10)))
                                _ = v18
                            } else {
                                switch ("" + v3) {
                                case "a":
                                    v19 := GenS4{F0: 0, F1: (v4 + (^(v4))), F2: (("a" + (((v3 + (v3 + v3)) + ("a" + "\n\t")) + "")) + v3)}
                                    _ = v19
                                default:
                                    fmt.Println(v2)
                                }
                            }
                        }
                        v1[(v4 - 2)] = "foo bar"
                        var v20 float64 = v13
                        _ = v20
                    }
                    var v21 bool = (true || false)
                    var v22 int8 = 124
                    _ = fi14
                    _ = v21
                    _ = v22
                }
            } else {
                v23 := GenS1{F0: v3}
                for fi24 := 0; fi24 < 5; fi24++ {
                    v1[(^(v4))] = "a"
                    _ = fi24
                }
                _ = v23
            }
            _ = v12
            _ = v13
        }
        _ = v6
    }
    v2 = (-(v2))
    if true {
        v4 -= (v4 * (^(v4)))
        if (!true) {
            v25 := GenS5{F0: 65535, F1: -602}
            _ = v25
        }
    } else {
        fmt.Println(v3, v2)
        if ((3.14159 - ((0.1 / 1.5) - float64(-2.25))) < 1e10) {
            fmt.Println(v2, v3)
        } else {
            for fi26 := 0; fi26 < 5; fi26++ {
                v27 := GenS6{F0: 1, F1: 18446744073709551615}
                v28 := GenS7{F0: 0}
                _ = fi26
                _ = v27
                _ = v28
            }
            fmt.Println(v3, v4)
        }
        switch (v3 + v3) {
        case "\n\t":
            v4 -= 2147483648
        case "hello":
            var v29 uint64 = 2
            _ = v29
        default:
            for fi30 := 0; fi30 < 18; fi30++ {
                switch 32768 {
                case -9223372036854775808:
                    var v31 int8 = 59
                    _ = v31
                case 42:
                    v32 := []int{}
                    _ = v32
                default:
                    var v33 uint64 = 0
                    _ = v33
                }
                var v34 uint16 = 42
                _ = fi30
                _ = v34
            }
        }
    }
    var v35 int64 = -9223372036854775808
    v3 = ("" + v3)
    v36 := GenS8{F0: v35}
    for fi37 := 0; fi37 < 9; fi37++ {
        if false {
            switch "hello" {
            case "hello":
                var v38 int = int(v4)
                _ = v38
            }
        }
        var v39 uint64 = 1
        var v40 int = fi37
        _ = fi37
        _ = v39
        _ = v40
    }
    var v41 int16 = (v2 | v2)
    v35 -= 9223372036854775807
    v2 = (^(v2))
    for fi42 := 0; fi42 < 12; fi42++ {
        v3 = v3
        if (v2 < (v41 & (^(v41)))) {
            if true {
                switch fi42 {
                case 9223372036854775807:
                    var v43 uint8 = 255
                    _ = v43
                case 0:
                    v44 := []int{}
                    _ = v44
                case 1:
                    delete(v1, v4)
                }
            } else {
                if ((-2147483648 <= -190) && false) {
                    v45 := []int{fi42, (-(fi42)), fi42}
                    fmt.Println(v35, fi42, v4)
                    if (!((!false) || true)) {
                        v35 = 9223372036854775807
                    } else {
                        v4 += v4
                    }
                    _ = v45
                } else {
                    if (129 <= 201) {
                        fmt.Println(v35, v2, fi42)
                    } else {
                        for fi46 := 0; fi46 < 4; fi46++ {
                            switch (fi46 ^ fi46) {
                            case 0:
                                v47 := []int{148, fi42, (^(fi42)), int(v4)}
                                _ = v47
                            case 128:
                                v48 := GenS9{F0: 32767, F1: 29}
                                _ = v48
                            case -818:
                                if (!(!((v3 + v3) <= ""))) {
                                    var v49 float32 = (((3.14159 * (3.14159 + -2.25)) * (((-2.25 * -2.25) * float32((0.0 / (0.1 - 3.14159)))) - (0.1 - 1.5))) + 0.1)
                                    switch v3 {
                                    case "a":
                                        v3 = "foo bar"
                                    default:
                                        if ((v35 + 850) == (v35 - v35)) {
                                            var v50 int16 = (-(v41))
                                            _ = v50
                                        } else {
                                            if (128 < 127) {
                                                v51 := v1[v4]
                                                if (!(42 >= int32(v35))) {
                                                    fmt.Println(v51)
                                                }
                                                _ = v51
                                            }
                                        }
                                    }
                                    v52 := []int{fi46}
                                    _ = v49
                                    _ = v52
                                }
                            default:
                                _ = len(v1)
                            }
                            v2 = v2
                            _ = fi46
                        }
                        v53 := GenS5{F0: 2, F1: 2147483647}
                        v1[(v4 - 1)] = v3
                        _ = v53
                    }
                    v35 = 42
                }
                v54 := []int{(fi42 - fi42), fi42, (fi42 - 84)}
                _ = v54
            }
            var v55 bool = (255 > 443)
            _ = v55
        }
        fmt.Println(v35)
        _ = fi42
    }
    v56 := []int{32767, -144, 42, 709}
    fmt.Println(v35, v4)
    for fi57 := 0; fi57 < 11; fi57++ {
        switch v3 {
        case "hello":
            v58 := GenS4{F0: 255, F1: 2, F2: "a"}
            _ = v58
        case "a":
            for fi59 := 0; fi59 < 5; fi59++ {
                if (!true) {
                    v3 = (v3 + ((("foo bar" + v3) + v3) + v3))
                } else {
                    v2 += v41
                    _ = len(v1)
                }
                if false {
                    v60 := GenS5{F0: 234, F1: -873}
                    v61 := GenS5{F0: 255, F1: 0}
                    _ = v60
                    _ = v61
                }
                v35 -= v35
                if false {
                    var v62 int32 = -1
                    v1[v4] = v3
                    _ = v62
                }
                _ = fi59
            }
        case "":
            for fi63 := 0; fi63 < 18; fi63++ {
                var v64 int64 = (v35 << ((1) & 63))
                switch (v3 + (v3 + "hello")) {
                case "a":
                    switch fi63 {
                    case -10:
                        v65 := GenS10{F0: (v35 * -476), F1: (v4 - 2), F2: "foo bar"}
                        _ = v65
                    case 4294967295:
                        if ((float32(1e10) - 3.14159) < (float32(0.0) - float32(((0.1 * 1e10) - float64(-2.25))))) {
                            var v66 float32 = (0.0 * 0.0)
                            _ = v66
                        } else {
                            v4 = (v4 / ((v4) | 1))
                        }
                    case -737:
                        if true {
                            var v67 float64 = 3.14159
                            var v68 int8 = -25
                            _ = v67
                            _ = v68
                        }
                    default:
                        v69 := GenS11{F0: 42}
                        _ = v69
                    }
                case "\n\t":
                    var v70 uint32 = v4
                    _ = v70
                case "":
                    switch (((("foo bar" + ("foo bar" + (v3 + "hello"))) + "") + v3) + ("" + "foo bar")) {
                    case "a":
                        fmt.Println(fi57)
                    case "hello":
                        v64 -= v64
                    case "\n\t":
                        if (-2.25 == 1e10) {
                            v2 = (-(v41))
                        } else {
                            var v71 uint = 18446744073709551615
                            _ = v71
                        }
                    default:
                        v41 += 588
                    }
                }
                _ = fi63
                _ = v64
            }
        default:
            for fi72 := 0; fi72 < 16; fi72++ {
                if ("" == v3) {
                    var v73 uint32 = v4
                    if false {
                        v73 = v73
                    } else {
                        v73 = (v73 % ((65535) | 1))
                        var v74 int64 = (v35 - v35)
                        switch "a" {
                        case "a":
                            switch int(v2) {
                            case -2147483648:
                                var v75 uint64 = 548
                                _ = v75
                            case 1:
                                if false {
                                    if true {
                                        var v76 uint16 = 32767
                                        var v77 uint16 = (v76 * (v76 / ((v76) | 1)))
                                        _ = v76
                                        _ = v77
                                    }
                                    if (1 <= 42) {
                                        delete(v1, 290)
                                        v78 := GenS9{F0: 255, F1: 65535}
                                        _ = v78
                                    } else {
                                        fmt.Println(fi57)
                                        if len(v56) > 0 {
                                            v56[0] = fi57
                                        }
                                        v41 -= (-(v2))
                                    }
                                    v56 = append(v56, 127)
                                    v4 -= v4
                                    v79 := GenS12{F0: 321}
                                    _ = v79
                                } else {
                                    v80 := GenS5{F0: 976, F1: -144}
                                    v73 = v73
                                    v1[v4] = ("hello" + (v3 + (v3 + ((("foo bar" + v3) + v3) + (("a" + (v3 + v3)) + "a")))))
                                    var v81 float64 = 1.5
                                    fmt.Println(v41)
                                    var v82 uint16 = 2
                                    var v83 int = -78
                                    _ = v80
                                    _ = v81
                                    _ = v82
                                    _ = v83
                                }
                                if len(v56) >= 2 {
                                    v56 = v56[1:len(v56)]
                                }
                                var v84 uint64 = 18446744073709551615
                                v4 = (v73 << ((828) & 63))
                                _ = v84
                            default:
                                delete(v1, 42)
                            }
                            v85 := GenS13{F0: 2, F1: 42}
                            _ = v85
                        case "":
                            fmt.Println(fi72)
                        }
                        _ = v74
                    }
                    _ = v73
                }
                if (65535 <= 950) {
                    v1[v4] = ("\n\t" + v3)
                    delete(v1, (v4 | v4))
                    var v86 uint8 = 127
                    _ = len(v1)
                    _ = v86
                }
                v2 = (v41 | 42)
                _ = fi72
            }
        }
        for fi87 := 0; fi87 < 14; fi87++ {
            var v88 uint32 = (v4 * 573)
            _ = fi87
            _ = v88
        }
        if false {
            fmt.Println(v35, v41)
        }
        _ = fi57
    }
    var v89 int = -2147483648
    for fi90 := 0; fi90 < 9; fi90++ {
        var v91 int = (fi90 >> ((255) & 63))
        _ = fi90
        _ = v91
    }
    v92 := GenS14{F0: -1, F1: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v56:", len(v56), v56)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v41
    _ = v56
    _ = v89
    _ = v92
}
