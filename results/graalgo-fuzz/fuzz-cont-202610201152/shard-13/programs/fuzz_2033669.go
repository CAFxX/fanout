package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int64; F1 uint; F2 string }
type GenS3 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS4 struct { F0 uint64; F1 string; F2 int32 }
type GenS5 struct { F0 bool; F1 int16; F2 uint16 }
type GenS6 struct { F0 int16; F1 int32 }
type GenS7 struct { F0 int64; F1 uint16; F2 int32 }
type GenS8 struct { F0 uint32 }

func main() {
    var v1 int8 = 42
    var v2 int32 = 2147483647
    var v3 string = ""
    v4 := GenS1{F0: (3.14159 > float32(float64(((1e10 + 0.1) + (1e10 + (-2.25 - 0.0)))))), F1: -79}
    v5 := GenS2{F0: -19, F1: 18446744073709551615, F2: (v3 + v3)}
    var v6 int16 = 32767
    if false {
        var v7 int64 = -9223372036854775808
        _ = v7
    }
    v6 -= v6
    if false {
        for fi8 := 0; fi8 < 14; fi8++ {
            v9 := GenS1{F0: true, F1: 848}
            _ = fi8
            _ = v9
        }
        v1 = (^(v1))
    } else {
        var v10 float32 = 1e10
        v11 := GenS3{F0: 42, F1: 42, F2: (v1 + -128)}
        if ((float64(((v10 / v10) + float32(0.1))) / -2.25) > 1e10) {
            for fi12 := 0; fi12 < 12; fi12++ {
                fmt.Println(v10, v3)
                _ = fi12
            }
            switch v3 {
            case "hello":
                fmt.Println(v6, v10, v1)
            case "\n\t":
                switch -366 {
                case 127:
                    v1 = -65
                case 42:
                    var v13 bool = ((v6 % (((v6 / ((v6) | 1))) | 1)) == (v6 * v6))
                    _ = v13
                default:
                    v14 := GenS1{F0: false, F1: 9223372036854775807}
                    _ = v14
                }
            case "foo bar":
                fmt.Println(v3)
            default:
                if (1 < 2147483648) {
                    v10 = v10
                    fmt.Println(v10)
                    if true {
                        if true {
                            v1 -= v1
                        }
                        if false {
                            switch "foo bar" {
                            case "hello":
                                for fi15 := 0; fi15 < 15; fi15++ {
                                    v3 = ((v3 + v3) + ("hello" + ""))
                                    if false {
                                        fmt.Println(v10, v10)
                                        fmt.Println(v1, v1, fi15)
                                        v16 := GenS1{F0: false, F1: -9223372036854775808}
                                        _ = v16
                                    }
                                    _ = fi15
                                }
                            }
                            for fi17 := 0; fi17 < 13; fi17++ {
                                fmt.Println(v10, v6, fi17)
                                _ = fi17
                            }
                            for fi18 := 0; fi18 < 16; fi18++ {
                                var v19 int = (fi18 / ((fi18) | 1))
                                v20 := GenS4{F0: 1, F1: (("" + v3) + v3), F2: (v2 / (((v2 | 42)) | 1))}
                                _ = fi18
                                _ = v19
                                _ = v20
                            }
                        } else {
                            fmt.Println(v2)
                            v21 := []int{32768, 413, 32768}
                            _ = v21
                        }
                        if true {
                            switch 26 {
                            case 102:
                                v22 := []int{-1}
                                _ = v22
                            case 42:
                                v3 = ("" + ("\n\t" + ("\n\t" + "foo bar")))
                            case 32767:
                                var v23 int = 127
                                _ = v23
                            }
                            v10 += float32(1.5)
                            var v24 int8 = -1
                            _ = v24
                        }
                    } else {
                        if true {
                            if (float32(1.5) == float32(1.5)) {
                                v25 := []int{128, -2147483648}
                                for fi26 := 0; fi26 < 13; fi26++ {
                                    _ = len(v25)
                                    v10 += ((v10 + 1.5) * (v10 + (v10 / (float32(((3.14159 + -2.25) - (float64(0.1) + (1e10 * (1e10 - 1e10))))) - (float32(1e10) + ((-2.25 + (v10 - (1e10 + 0.0))) - 0.1))))))
                                    _ = fi26
                                }
                                _ = v25
                            } else {
                                fmt.Println(v10)
                                fmt.Println(v3, v2, v10)
                            }
                            fmt.Println(v6)
                        } else {
                            for fi27 := 0; fi27 < 6; fi27++ {
                                if false {
                                    if false {
                                        v28 := []int{}
                                        _ = v28
                                    } else {
                                        if ((158 > 4294967295) || true) {
                                            switch (fi27 - (fi27 + (^(fi27)))) {
                                            case 197:
                                                var v29 uint8 = 0
                                                _ = v29
                                            default:
                                                switch "\n\t" {
                                                case "\n\t":
                                                    v3 = v3
                                                case "a":
                                                    var v30 int = 2147483647
                                                    _ = v30
                                                case "hello":
                                                    if (2 != 65535) {
                                                        fmt.Println(v3, v10, v1)
                                                        v2 += 812
                                                        for fi31 := 0; fi31 < 9; fi31++ {
                                                            v32 := GenS5{F0: false, F1: v6, F2: 356}
                                                            _ = fi31
                                                            _ = v32
                                                        }
                                                    }
                                                default:
                                                    switch fi27 {
                                                    case 42:
                                                        v2 = (v2 | (v2 & 1))
                                                    }
                                                }
                                            }
                                            v33 := GenS3{F0: 1, F1: 65535, F2: 100}
                                            _ = v33
                                        }
                                        fmt.Println(v1, v6, v3)
                                    }
                                } else {
                                    if (-9223372036854775808 >= 9223372036854775807) {
                                        if false {
                                            v34 := []int{}
                                            _ = v34
                                        }
                                        v1 -= (v1 & (-(v1)))
                                        var v35 int64 = 0
                                        _ = v35
                                    }
                                    var v36 int32 = v2
                                    v2 += v2
                                    _ = v36
                                }
                                _ = fi27
                            }
                        }
                        fmt.Println(v10, v1, v6)
                    }
                } else {
                    var v37 uint8 = 128
                    v38 := []int{-49}
                    _ = v37
                    _ = v38
                }
            }
        }
        _ = v10
        _ = v11
    }
    var v39 uint32 = 42
    v6 += v6
    if true {
        v40 := GenS1{F0: (!((^(v39)) <= (^(v39)))), F1: -309}
        if (!false) {
            for fi41 := 0; fi41 < 14; fi41++ {
                var v42 bool = true
                var v43 string = v3
                var v44 uint16 = 65535
                _ = fi41
                _ = v42
                _ = v43
                _ = v44
            }
        } else {
            switch "" {
            case "hello":
                var v45 uint8 = uint8(v39)
                _ = v45
            case "\n\t":
                for fi46 := 0; fi46 < 8; fi46++ {
                    var v47 float64 = 0.1
                    _ = fi46
                    _ = v47
                }
            case "":
                if false {
                    v2 = -471
                    v48 := GenS3{F0: 4294967295, F1: 744, F2: (v1 / ((v1) | 1))}
                    _ = v48
                }
            default:
                if (0.1 > 1.5) {
                    var v49 string = ((v3 + (("hello" + "hello") + (("a" + (v3 + "")) + "\n\t"))) + v3)
                    v6 = v6
                    _ = v49
                }
            }
        }
        _ = v40
    } else {
        switch "foo bar" {
        case "foo bar":
            if ((float32(0.0) / 0.1) <= (0.0 * float32((0.0 / 1e10)))) {
                for fi50 := 0; fi50 < 1; fi50++ {
                    v51 := []int{}
                    _ = fi50
                    _ = v51
                }
            } else {
                switch v3 {
                case "foo bar":
                    fmt.Println(v3)
                }
                var v52 bool = true
                var v53 uint8 = 255
                _ = v52
                _ = v53
            }
        case "hello":
            var v54 int64 = -673
            _ = v54
        default:
            switch 557 {
            case -318:
                v55 := []int{9223372036854775807, 128, 32768, 9223372036854775807}
                _ = v55
            case 1:
                if (true && false) {
                    var v56 float32 = 0.1
                    fmt.Println(v1, v1, v6)
                    _ = v56
                }
            case 0:
                var v57 uint8 = 127
                _ = v57
            default:
                for fi58 := 0; fi58 < 20; fi58++ {
                    v2 -= v2
                    _ = fi58
                }
            }
        }
        fmt.Println(v1)
    }
    var v59 uint64 = 2
    v60 := []int{4294967295}
    if true {
        fmt.Println(v6)
    } else {
        v61 := GenS5{F0: (787 < 65535), F1: v6, F2: 2}
        var v62 uint32 = v39
        _ = v61
        _ = v62
    }
    v63 := GenS2{F0: -24, F1: 0, F2: ((("foo bar" + ("\n\t" + (v3 + ("a" + ("hello" + v3))))) + ((v3 + ("hello" + (("foo bar" + v3) + v3))) + v3)) + "\n\t")}
    if len(v60) >= 2 {
        v60 = v60[1:len(v60)]
    }
    if false {
        if true {
            for fi64 := 0; fi64 < 5; fi64++ {
                v1 = int8(v59)
                _ = fi64
            }
            fmt.Println(v6, v1, v3)
            for fi65 := 0; fi65 < 6; fi65++ {
                switch ((v3 + v3) + v3) {
                case "":
                    if false {
                        fmt.Println(fi65, v59, v6)
                        switch v3 {
                        case "hello":
                            for fi66 := 0; fi66 < 11; fi66++ {
                                v67 := GenS6{F0: -32768, F1: (v2 & (v2 * (v2 | v2)))}
                                var v68 string = ("a" + "foo bar")
                                switch v68 {
                                case "a":
                                    if ((v1 ^ 120) > v1) {
                                        v39 = (^(v39))
                                    } else {
                                        v69 := GenS3{F0: 18446744073709551615, F1: (v39 >> ((2) & 63)), F2: v1}
                                        _ = v69
                                    }
                                case "\n\t":
                                    v70 := GenS6{F0: v6, F1: int32(v59)}
                                    _ = v70
                                }
                                _ = fi66
                                _ = v67
                                _ = v68
                            }
                        default:
                            var v71 int = (^(fi65))
                            _ = v71
                        }
                    } else {
                        v3 = "hello"
                    }
                default:
                    switch (fi65 | (fi65 + fi65)) {
                    case 40:
                        if len(v60) > 0 {
                            v60[0] = (fi65 - 127)
                        }
                    case 127:
                        v72 := GenS1{F0: (int(v39) < fi65), F1: -404}
                        _ = v72
                    }
                }
                var v73 int32 = v2
                if len(v60) > 0 {
                    v60[0] = (fi65 % ((fi65) | 1))
                }
                _ = fi65
                _ = v73
            }
        }
    } else {
        for fi74 := 0; fi74 < 6; fi74++ {
            for fi75 := 0; fi75 < 2; fi75++ {
                v76 := GenS7{F0: 0, F1: 65535, F2: (v2 - 42)}
                if (!((true && true) || false)) {
                    var v77 float64 = (((((-2.25 / (-2.25 / 1.0)) - 1.5) + 1.5) + 1.5) + 1.5)
                    v39 = v39
                    _ = v77
                } else {
                    if false {
                        if len(v60) >= 2 {
                            v60 = v60[1:len(v60)]
                        }
                    }
                    if ((v1 & v1) > (v1 + (^(v1)))) {
                        if false {
                            v78 := GenS7{F0: 42, F1: 255, F2: v2}
                            _ = v78
                        }
                        if len(v60) > 0 {
                            v60[0] = -2147483648
                        }
                    }
                }
                _ = fi75
                _ = v76
            }
            if ((0.0 <= 1e10) || false) {
                var v79 float32 = float32(((float64(((1.5 * 0.1) * 1e10)) - (-2.25 * -2.25)) / 1e10))
                switch (v3 + "") {
                case "":
                    v80 := GenS8{F0: v39}
                    _ = v80
                }
                v3 = (v3 + (v3 + ""))
                _ = v79
            } else {
                fmt.Println(fi74)
                v81 := GenS2{F0: 42, F1: 255, F2: ("foo bar" + v3)}
                var v82 uint32 = (^(v39))
                _ = v81
                _ = v82
            }
            switch "" {
            case "a":
                _ = len(v60)
            case "foo bar":
                _ = len(v60)
            default:
                v59 = (v59 + (^(v59)))
            }
            _ = fi74
        }
    }
    var v83 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v60:", len(v60), v60)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v60
    _ = v63
    _ = v83
}
