package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int; F1 int8; F2 uint }
func genFn0() uint16 {
    if true {
        var v4 uint32 = 4294967295
        fmt.Println(v4)
        for fi5 := 0; fi5 < 20; fi5++ {
            for fi6 := 0; fi6 < 6; fi6++ {
                var v7 int16 = 42
                var v8 int32 = 2147483647
                v7 = v7
                v8 = (v8 | v8)
                _ = fi6
                _ = v7
                _ = v8
            }
            v9 := []int{fi5, fi5}
            _ = fi5
            _ = v9
        }
        _ = v4
    } else {
        var v10 uint8 = 1
        for fi11 := 0; fi11 < 5; fi11++ {
            for fi12 := 0; fi12 < 9; fi12++ {
                var v13 uint16 = uint16(v10)
                _ = fi12
                _ = v13
            }
            v14 := []int{fi11}
            if (1e10 >= 0.0) {
                var v15 int32 = -2147483648
                _ = v15
            } else {
                switch 1 {
                case 459:
                    var v16 uint64 = 4294967295
                    _ = v16
                default:
                    switch "" {
                    case "":
                        switch "\n\t" {
                        case "foo bar":
                            v10 += 255
                        case "a":
                            v10 = (v10 ^ (v10 ^ 92))
                        }
                    case "a":
                        var v17 float64 = (((((1e10 - 3.14159) + (float64((1.5 / ((0.0 + float32(0.1)) + 1e10))) - (float64(3.14159) - float64(1.5)))) / 0.1) * -2.25) - (1e10 / -2.25))
                        _ = v17
                    case "foo bar":
                        for fi18 := 0; fi18 < 11; fi18++ {
                            if (9223372036854775807 >= (fi18 % ((fi18) | 1))) {
                                v10 += v10
                            } else {
                                v19 := GenS1{F0: ((1.5 / 0.1) == float32(0.0))}
                                if true {
                                    v10 = v10
                                    v10 -= (v10 ^ 0)
                                    var v20 bool = false
                                    if (!v20) {
                                        v10 = v10
                                        var v21 int64 = -1
                                        if ((v21 & 669) == (v21 / ((-78) | 1))) {
                                            if (591 <= 731) {
                                                if v20 {
                                                    v20 = v20
                                                    var v22 string = "a"
                                                    var v23 int16 = -32768
                                                    _ = v22
                                                    _ = v23
                                                } else {
                                                    v24 := GenS2{F0: (fi18 | fi18), F1: "foo bar"}
                                                    var v25 uint64 = 1
                                                    v21 = 1
                                                    _ = v24
                                                    _ = v25
                                                }
                                                v26 := GenS3{F0: 783}
                                                _ = v26
                                            } else {
                                                if v20 {
                                                    var v27 int16 = 1000
                                                    if len(v14) > 0 {
                                                        v14[0] = fi18
                                                    }
                                                    _ = v27
                                                } else {
                                                    switch (("foo bar" + "hello") + "foo bar") {
                                                    case "":
                                                        v20 = (65535 > 552)
                                                    case "a":
                                                        var v28 int64 = 0
                                                        _ = v28
                                                    case "\n\t":
                                                        v21 = v21
                                                    default:
                                                        v29 := GenS1{F0: (("\n\t" + "a") >= "foo bar")}
                                                        _ = v29
                                                    }
                                                }
                                                switch fi11 {
                                                case -463:
                                                    v30 := GenS2{F0: (-(fi18)), F1: ""}
                                                    _ = v30
                                                case 128:
                                                    v10 = 255
                                                case 127:
                                                    v21 -= (v21 | v21)
                                                default:
                                                    if (32767 != 32767) {
                                                        fmt.Println(fi11, v21)
                                                    }
                                                }
                                            }
                                            v21 = v21
                                        } else {
                                            v31 := GenS1{F0: v20}
                                            var v32 float32 = (float32(((-2.25 * -2.25) / (-2.25 + (0.1 + -2.25)))) / ((1e10 / 1.0) / -2.25))
                                            fmt.Println(v10, v32)
                                            _ = v31
                                            _ = v32
                                        }
                                        _ = v21
                                    }
                                    _ = v20
                                } else {
                                    if ("a" == "hello") {
                                        if true {
                                            var v33 int32 = -2147483648
                                            _ = v33
                                        }
                                        if (0 >= 1) {
                                            var v34 int16 = 684
                                            _ = v34
                                        }
                                        v10 += v10
                                    }
                                    var v35 int32 = 2147483647
                                    var v36 uint = 559
                                    _ = v35
                                    _ = v36
                                }
                                _ = v19
                            }
                            if true {
                                var v37 float64 = 0.0
                                var v38 string = ""
                                v37 += 1.5
                                _ = v37
                                _ = v38
                            }
                            var v39 uint8 = v10
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                            fmt.Println(v10)
                            _ = fi18
                            _ = v39
                        }
                    }
                }
                fmt.Println(fi11, fi11)
            }
            _ = fi11
            _ = v14
        }
        _ = v10
    }
    var v40 uint16 = 985
    var v41 float32 = (float32(0.1) - ((0.1 + (float32(-2.25) - 3.14159)) + -2.25))
    var v42 uint32 = 0
    _ = v40
    _ = v41
    _ = v42
    return v40
}
func genFn1(p0 uint16, p1 string) int {
    v43 := []int{}
    var v44 uint64 = 18446744073709551615
    _ = p0
    _ = p1
    _ = v43
    _ = v44
    return 935
}

func main() {
    var v1 uint32 = 116
    var v2 uint32 = v1
    var v3 int16 = 592
    v1 -= 2147483647
    switch -772 {
    case 128:
        v3 += (v3 << ((255) & 63))
    case -1:
        if ("\n\t" <= "\n\t") {
            if (!(1 > -1)) {
                var v45 uint8 = 128
                _ = v45
            }
            var v46 int64 = 9223372036854775807
            _ = v46
        }
    case 4294967295:
        var v47 int32 = 689
        _ = v47
    default:
        v48 := GenS1{F0: (304 == 255)}
        _ = v48
    }
    var v49 int64 = -937
    v2 += (v1 ^ v2)
    if false {
        if false {
            v3 -= (v3 - v3)
            switch 2147483647 {
            case -240:
                if false {
                    var v50 int64 = (v49 & v49)
                    fmt.Println(v3)
                    v49 = v50
                    _ = v50
                }
            case 347:
                fmt.Println(v3, v3, v3)
            }
            if ((3.14159 * float32((1e10 - 0.1))) <= -2.25) {
                for fi51 := 0; fi51 < 15; fi51++ {
                    var v52 int = genFn1(uint16(v1), (("a" + "\n\t") + ("hello" + "")))
                    v1 = v2
                    if true {
                        for fi53 := 0; fi53 < 6; fi53++ {
                            fmt.Println(v2, v1)
                            v54 := []int{}
                            v54 = append(v54, fi51)
                            _ = fi53
                            _ = v54
                        }
                    }
                    _ = fi51
                    _ = v52
                }
                var v55 float64 = ((float64(3.14159) - ((float64((1e10 + (((-2.25 * (3.14159 - 1.5)) + (0.0 * -2.25)) * (1.5 * 0.0)))) - -2.25) * (0.1 - 3.14159))) * (((1.5 + (-2.25 * -2.25)) - ((-2.25 * (-2.25 - 1.5)) * -2.25)) + ((-2.25 / (0.1 / 3.14159)) - 1.5)))
                v56 := GenS1{F0: true}
                _ = v55
                _ = v56
            } else {
                for fi57 := 0; fi57 < 8; fi57++ {
                    v58 := []int{genFn1(genFn0(), "hello")}
                    if len(v58) >= 2 {
                        v58 = v58[1:len(v58)]
                    }
                    _ = fi57
                    _ = v58
                }
            }
        }
    } else {
        v2 = (^(v2))
        if ((v1 ^ (v1 ^ 42)) >= (v2 / ((v2) | 1))) {
            switch "\n\t" {
            case "a":
                var v59 bool = false
                _ = v59
            }
            v2 = v1
        } else {
            switch 32767 {
            case 9223372036854775807:
                v1 += 213
            case -912:
                var v60 uint64 = 42
                _ = v60
            case 346:
                switch -1 {
                case 4294967295:
                    if true {
                        var v61 int16 = 1
                        for fi62 := 0; fi62 < 3; fi62++ {
                            fmt.Println(v1, v2)
                            v63 := GenS3{F0: 2}
                            var v64 int16 = 867
                            _ = fi62
                            _ = v63
                            _ = v64
                        }
                        _ = v61
                    } else {
                        for fi65 := 0; fi65 < 4; fi65++ {
                            v66 := []int{}
                            for fi67 := 0; fi67 < 1; fi67++ {
                                v68 := GenS2{F0: (-(fi65)), F1: ((("" + ("foo bar" + ("a" + (("\n\t" + "hello") + "")))) + "") + ("\n\t" + ""))}
                                v69 := GenS4{F0: (fi67 ^ fi67), F1: 1, F2: 1}
                                v70 := GenS3{F0: genFn0()}
                                _ = fi67
                                _ = v68
                                _ = v69
                                _ = v70
                            }
                            if (42 <= 0) {
                                if (-61 > 1) {
                                    var v71 float64 = -2.25
                                    if (v49 <= 42) {
                                        fmt.Println(fi65, v71, v1)
                                    } else {
                                        for fi72 := 0; fi72 < 10; fi72++ {
                                            var v73 int32 = -77
                                            _ = fi72
                                            _ = v73
                                        }
                                        if true {
                                            v3 = (v3 * (v3 * v3))
                                            for fi74 := 0; fi74 < 4; fi74++ {
                                                if len(v66) > 0 {
                                                    v66[0] = (fi74 % (((fi65 - -497)) | 1))
                                                }
                                                _ = fi74
                                            }
                                        }
                                    }
                                    _ = v71
                                }
                                v75 := GenS4{F0: genFn1(135, "hello"), F1: 1, F2: 2}
                                _ = v75
                            }
                            _ = fi65
                            _ = v66
                        }
                        switch genFn1(833, (("foo bar" + ("hello" + ("a" + ("" + "foo bar")))) + "hello")) {
                        case 32768:
                            var v76 uint8 = 128
                            _ = v76
                        case -57:
                            switch "\n\t" {
                            case "hello":
                                v2 -= v1
                            case "":
                                fmt.Println(v2)
                            default:
                                var v77 uint = 42
                                _ = v77
                            }
                        }
                    }
                }
            }
            for fi78 := 0; fi78 < 8; fi78++ {
                var v79 float64 = 1.5
                fmt.Println(v79)
                _ = fi78
                _ = v79
            }
        }
    }
    v80 := []int{genFn1(0, ("\n\t" + "hello"))}
    var v81 uint64 = 880
    v82 := GenS4{F0: 127, F1: 127, F2: 255}
    if false {
        var v83 int = 1
        if len(v80) >= 2 {
            v80 = v80[1:len(v80)]
        }
        _ = v83
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v80:", len(v80), v80)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v49
    _ = v80
    _ = v81
    _ = v82
}
