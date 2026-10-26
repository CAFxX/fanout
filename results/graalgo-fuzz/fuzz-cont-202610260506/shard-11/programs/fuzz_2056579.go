package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool; F1 int32; F2 string }
type GenS5 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 bool; F1 int32; F2 string }
type GenS8 struct { F0 bool; F1 int16 }
type GenS9 struct { F0 int16; F1 uint32 }
type GenS10 struct { F0 uint16 }
func genFn0(p0 uint8, p1 int) int16 {
    var v4 uint16 = 42
    p0 += (^(p0))
    var v5 uint64 = 4294967295
    var v6 int = (^(p1))
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 27
}
func genFn1() bool {
    if (true || (1 != 230)) {
        for fi7 := 0; fi7 < 10; fi7++ {
            for fi8 := 0; fi8 < 4; fi8++ {
                switch "hello" {
                case "":
                    if (0.1 != -2.25) {
                        switch fi7 {
                        case 127:
                            var v9 float64 = 1.5
                            _ = v9
                        case -1:
                            var v10 int8 = 100
                            _ = v10
                        default:
                            var v11 string = "a"
                            _ = v11
                        }
                    } else {
                        if (false && true) {
                            var v12 int32 = 0
                            if (1000 == 42) {
                                v12 -= (^(v12))
                                if true {
                                    if (!true) {
                                        v13 := GenS1{F0: genFn0(127, -1)}
                                        fmt.Println(fi8, v12, v12)
                                        if true {
                                            v14 := GenS1{F0: -1}
                                            var v15 uint16 = 65535
                                            _ = v14
                                            _ = v15
                                        } else {
                                            fmt.Println(v12)
                                            var v16 uint16 = 1
                                            v17 := []int{fi8, -2147483648}
                                            var v18 float64 = float64((3.14159 / (1.5 * 0.1)))
                                            _ = v16
                                            _ = v17
                                            _ = v18
                                        }
                                        _ = v13
                                    }
                                } else {
                                    var v19 int8 = -95
                                    _ = v19
                                }
                            } else {
                                var v20 uint64 = 252
                                _ = v20
                            }
                            v21 := []int{int(v12), -672, fi8, fi7}
                            var v22 string = "hello"
                            switch fi8 {
                            case -248:
                                var v23 uint8 = 42
                                _ = v23
                            }
                            switch "foo bar" {
                            case "hello":
                                v24 := GenS2{F0: 255, F1: 42}
                                _ = v24
                            case "foo bar":
                                var v25 int16 = 1
                                _ = v25
                            default:
                                switch (v22 + v22) {
                                case "foo bar":
                                    var v26 float32 = 0.0
                                    _ = v26
                                case "a":
                                    var v27 int16 = -860
                                    _ = v27
                                case "hello":
                                    v12 += -565
                                }
                            }
                            _ = v12
                            _ = v21
                            _ = v22
                        }
                        switch ("hello" + "\n\t") {
                        case "":
                            var v28 float32 = 0.0
                            var v29 uint64 = 0
                            _ = v28
                            _ = v29
                        case "\n\t":
                            if (((!false) && (9223372036854775807 < -1)) && ("foo bar" <= ("foo bar" + "hello"))) {
                                v30 := GenS1{F0: -32768}
                                _ = v30
                            } else {
                                var v31 uint32 = 42
                                if false {
                                    switch fi8 {
                                    case -649:
                                        var v32 bool = (-2.25 > float64(-2.25))
                                        _ = v32
                                    case -1:
                                        v33 := GenS3{F0: 32767}
                                        _ = v33
                                    default:
                                        v34 := GenS2{F0: 42, F1: 255}
                                        _ = v34
                                    }
                                }
                                _ = v31
                            }
                        case "hello":
                            v35 := GenS2{F0: 785, F1: 18446744073709551615}
                            _ = v35
                        default:
                            v36 := []int{}
                            _ = v36
                        }
                    }
                }
                v37 := []int{fi8, (-(fi8)), 2147483647}
                var v38 bool = false
                _ = fi8
                _ = v37
                _ = v38
            }
            _ = fi7
        }
    } else {
        if true {
            var v39 int16 = -32768
            if true {
                v40 := GenS4{F0: (0 > 1), F1: 0, F2: ""}
                fmt.Println(v39)
                for fi41 := 0; fi41 < 10; fi41++ {
                    var v42 bool = true
                    if (727 <= 4294967295) {
                        for fi43 := 0; fi43 < 2; fi43++ {
                            v42 = (18446744073709551615 >= 42)
                            if v42 {
                                v42 = v42
                                v42 = v42
                                var v44 int = (fi41 ^ (^(fi41)))
                                _ = v44
                            } else {
                                v45 := GenS5{F0: (^(v39)), F1: 439, F2: 564}
                                _ = v45
                            }
                            _ = fi43
                        }
                    }
                    switch (("foo bar" + "hello") + "a") {
                    case "a":
                        var v46 uint32 = 641
                        _ = v46
                    default:
                        var v47 string = "hello"
                        _ = v47
                    }
                    _ = fi41
                    _ = v42
                }
                _ = v40
            } else {
                var v48 bool = false
                _ = v48
            }
            _ = v39
        } else {
            v49 := GenS3{F0: 32767}
            _ = v49
        }
        if (true && false) {
            var v50 uint64 = 2
            fmt.Println(v50)
            _ = v50
        } else {
            for fi51 := 0; fi51 < 7; fi51++ {
                var v52 float64 = 1e10
                if true {
                    v53 := GenS4{F0: true, F1: -1, F2: "a"}
                    _ = v53
                }
                _ = fi51
                _ = v52
            }
            var v54 uint32 = 2147483648
            _ = v54
        }
    }
    v55 := GenS6{F0: 441}
    _ = v55
    return true
}

func main() {
    var v1 float64 = (1e10 * 3.14159)
    v2 := make(map[uint32]uint32)
    var v3 int8 = 100
    for fi56 := 0; fi56 < 3; fi56++ {
        v3 += v3
        switch "a" {
        case "\n\t":
            switch ("foo bar" + "a") {
            case "a":
                v57 := []int{(fi56 * fi56), fi56, 765}
                _ = v57
            case "\n\t":
                v58 := GenS7{F0: (v3 >= (v3 / (((v3 - -114)) | 1))), F1: 296, F2: (((("" + (("" + "hello") + "foo bar")) + "\n\t") + "\n\t") + "foo bar")}
                _ = v58
            default:
                v3 = (v3 % ((52) | 1))
            }
        case "":
            var v59 uint32 = 48
            _ = v59
        case "a":
            for fi60 := 0; fi60 < 3; fi60++ {
                v3 -= (v3 + (v3 << ((2) & 63)))
                switch (fi60 / (((fi56 % (((^(fi56))) | 1))) | 1)) {
                case -306:
                    if (!true) {
                        v1 += 1.5
                        var v61 float32 = 1.5
                        var v62 float32 = float32(float64(float32(1e10)))
                        _ = v61
                        _ = v62
                    } else {
                        fmt.Println(fi56)
                        v3 += v3
                    }
                case 1:
                    v63 := GenS8{F0: false, F1: 32767}
                    _ = v63
                default:
                    v64 := GenS5{F0: 32767, F1: 604, F2: 1}
                    _ = v64
                }
                _ = fi60
            }
        }
        _ = fi56
    }
    var v65 uint64 = 766
    delete(v2, 0)
    var v66 int8 = (^(v3))
    if true {
        v67 := []int{9223372036854775807}
        _ = v67
    }
    switch "foo bar" {
    case "a":
        for fi68 := 0; fi68 < 3; fi68++ {
            v69 := []int{fi68}
            fmt.Println(fi68, fi68, v66)
            if false {
                var v70 uint = 630
                delete(v2, 2147483647)
                fmt.Println(v70)
                _ = v70
            }
            _ = fi68
            _ = v69
        }
    case "hello":
        v71 := GenS8{F0: (2147483647 <= 1), F1: -944}
        _ = v71
    case "\n\t":
        var v72 float64 = v1
        _ = v72
    default:
        var v73 float64 = ((v1 / (v1 / ((v1 - v1) - 0.1))) / v1)
        _ = v73
    }
    v74 := GenS9{F0: genFn0(1, 32768), F1: 4294967295}
    if (18446744073709551615 >= (^(v65))) {
        fmt.Println(v66, v1)
    } else {
        delete(v2, 65535)
        v75 := v2[0]
        _ = v75
    }
    if true {
        delete(v2, 42)
        _ = len(v2)
    }
    if true {
        var v76 int = -9223372036854775808
        var v77 uint8 = 2
        _ = v76
        _ = v77
    } else {
        for fi78 := 0; fi78 < 12; fi78++ {
            for fi79 := 0; fi79 < 20; fi79++ {
                v3 = -128
                v80 := []int{(-(fi78)), 128, -232}
                var v81 uint64 = (v65 + 169)
                _ = fi79
                _ = v80
                _ = v81
            }
            v82 := GenS10{F0: uint16(v65)}
            v83 := []int{-93, 128, -214, (fi78 >> ((18446744073709551615) & 63))}
            _ = fi78
            _ = v82
            _ = v83
        }
        if true {
            fmt.Println(v3, v1)
            v84 := GenS9{F0: 589, F1: 2147483647}
            _ = v84
        }
        var v85 string = ("foo bar" + ("hello" + "foo bar"))
        _ = v85
    }
    var v86 uint64 = v65
    for fi87 := 0; fi87 < 10; fi87++ {
        switch (fi87 | (fi87 * (fi87 % (((fi87 ^ -2147483648)) | 1)))) {
        case 0:
            _ = len(v2)
        case -402:
            var v88 string = "a"
            _ = v88
        case 2147483647:
            fmt.Println(v65, v86)
        }
        _ = fi87
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v65
    _ = v66
    _ = v74
    _ = v86
}
