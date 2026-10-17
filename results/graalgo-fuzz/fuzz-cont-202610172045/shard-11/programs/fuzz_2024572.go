package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8; F1 int8 }
type GenS5 struct { F0 bool; F1 int16; F2 uint16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int32; F1 int64; F2 int32 }
type GenS8 struct { F0 uint; F1 bool }
type GenS9 struct { F0 int8; F1 uint8; F2 bool }
func genFn0(p0 uint64) string {
    var v5 uint8 = 0
    for fi6 := 0; fi6 < 7; fi6++ {
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 int32 = 42
            v9 := []int{-494}
            switch "foo bar" {
            case "a":
                if true {
                    fmt.Println(v8)
                } else {
                    fmt.Println(fi6, fi6)
                    switch ("" + ("foo bar" + "a")) {
                    case "\n\t":
                        v10 := GenS1{F0: 727, F1: -1}
                        _ = v10
                    case "":
                        v8 += v8
                    case "hello":
                        var v11 int16 = -32768
                        _ = v11
                    default:
                        if false {
                            v12 := GenS1{F0: 65535, F1: -417}
                            var v13 uint8 = v5
                            _ = v12
                            _ = v13
                        }
                    }
                }
            case "foo bar":
                if len(v9) > 0 {
                    v9[0] = (fi7 | 273)
                }
            default:
                p0 = 13
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        if (!true) {
            var v14 int8 = -95
            var v15 string = "a"
            _ = v14
            _ = v15
        } else {
            switch "" {
            case "a":
                switch "hello" {
                case "hello":
                    v16 := GenS2{F0: (v5 ^ 152)}
                    _ = v16
                }
            case "hello":
                if false {
                    if (((!false) || false) || ((fi6 == int(v5)) || (!(!false)))) {
                        if (("foo bar" + "") <= "") {
                            v17 := GenS3{F0: 1}
                            var v18 uint8 = v5
                            var v19 float64 = float64(1e10)
                            _ = v17
                            _ = v18
                            _ = v19
                        }
                        v20 := []int{(fi6 * 128), (fi6 * int(v5)), fi6}
                        _ = v20
                    } else {
                        switch -774 {
                        case -1:
                            if (0 != 127) {
                                if true {
                                    v5 -= (v5 % ((127) | 1))
                                    var v21 int64 = -778
                                    _ = v21
                                }
                                switch int(v5) {
                                case -9223372036854775808:
                                    fmt.Println(fi6)
                                case 32768:
                                    for fi22 := 0; fi22 < 15; fi22++ {
                                        var v23 float64 = float64(0.0)
                                        var v24 int16 = 539
                                        var v25 uint8 = (^(v5))
                                        fmt.Println(v5, v24, v23)
                                        _ = fi22
                                        _ = v23
                                        _ = v24
                                        _ = v25
                                    }
                                case 0:
                                    v5 += 128
                                }
                                for fi26 := 0; fi26 < 16; fi26++ {
                                    v27 := GenS4{F0: 112, F1: -128}
                                    v28 := []int{793, fi26, fi26}
                                    if len(v28) > 0 {
                                        v28[0] = (^(fi26))
                                    }
                                    _ = fi26
                                    _ = v27
                                    _ = v28
                                }
                            } else {
                                fmt.Println(fi6)
                                fmt.Println(fi6, p0, v5)
                            }
                        }
                        v5 = v5
                        var v29 float32 = (-2.25 - 3.14159)
                        _ = v29
                    }
                    v30 := []int{}
                    v5 -= (v5 * v5)
                    _ = v30
                } else {
                    switch "foo bar" {
                    case "\n\t":
                        v5 -= 11
                    default:
                        v31 := []int{(fi6 * (fi6 & fi6)), (fi6 - int(v5)), -725, (-(fi6))}
                        _ = v31
                    }
                    if ((-2.25 + (1.5 / 0.1)) > ((((float32((1e10 / (0.0 + ((((0.0 / (0.0 + 1e10)) - -2.25) * -2.25) * 1.5)))) - 0.1) + 0.0) - (-2.25 - (0.1 / float32(-2.25)))) + ((((((-2.25 * 0.1) + -2.25) + 1.5) - 0.0) + (1.5 - (0.1 / 1e10))) / 1.0))) {
                        if (!false) {
                            fmt.Println(v5)
                            var v32 uint32 = 289
                            _ = v32
                        } else {
                            var v33 int64 = 42
                            _ = v33
                        }
                        switch (fi6 >> ((1) & 63)) {
                        case -2147483648:
                            v5 = 127
                        default:
                            p0 += (p0 - 593)
                        }
                    }
                }
            }
            if (true || false) {
                var v34 uint16 = 537
                for fi35 := 0; fi35 < 10; fi35++ {
                    if false {
                        if ((1.5 < (0.0 - 1e10)) && (!(75 > 0))) {
                            v34 = (v34 / ((v34) | 1))
                            v36 := []int{(^(fi6)), fi35}
                            _ = v36
                        } else {
                            fmt.Println(v34)
                            fmt.Println(p0, v34)
                            var v37 uint32 = 65535
                            _ = v37
                        }
                        fmt.Println(fi6, v34, v34)
                    }
                    _ = fi35
                }
                v34 += 25
                _ = v34
            } else {
                switch (fi6 << ((255) & 63)) {
                case -1:
                    var v38 int = fi6
                    _ = v38
                }
                v39 := GenS5{F0: true, F1: -770, F2: 1}
                _ = v39
            }
            var v40 int8 = int8(p0)
            _ = v40
        }
        _ = fi6
    }
    v5 += v5
    _ = p0
    _ = v5
    return "a"
}
func genFn1() uint32 {
    switch 32768 {
    case 42:
        for fi41 := 0; fi41 < 5; fi41++ {
            fmt.Println(fi41)
            var v42 uint = 2
            _ = fi41
            _ = v42
        }
    case 0:
        for fi43 := 0; fi43 < 3; fi43++ {
            var v44 int32 = 440
            var v45 uint32 = 0
            _ = fi43
            _ = v44
            _ = v45
        }
    case 127:
        for fi46 := 0; fi46 < 20; fi46++ {
            v47 := []int{(fi46 + fi46), (fi46 >> ((42) & 63))}
            var v48 int64 = 65
            _ = fi46
            _ = v47
            _ = v48
        }
    default:
        var v49 int8 = 127
        _ = v49
    }
    v50 := []int{-744}
    switch -1 {
    case 192:
        for fi51 := 0; fi51 < 1; fi51++ {
            var v52 int16 = 0
            for fi53 := 0; fi53 < 20; fi53++ {
                fmt.Println(fi53, v52)
                switch "foo bar" {
                case "\n\t":
                    if (!((fi51 > fi53) || (557 < 962))) {
                        var v54 string = ""
                        _ = v54
                    }
                case "hello":
                    fmt.Println(v52, v52, v52)
                case "":
                    var v55 int16 = v52
                    _ = v55
                default:
                    v52 += (v52 >> ((1) & 63))
                }
                if len(v50) >= 2 {
                    v50 = v50[1:len(v50)]
                }
                if (0 >= 1) {
                    fmt.Println(v52, fi53)
                    if len(v50) >= 2 {
                        v50 = v50[1:len(v50)]
                    }
                }
                _ = fi53
            }
            _ = fi51
            _ = v52
        }
    }
    v56 := GenS5{F0: ((1.5 - (0.1 / (3.14159 - 0.1))) >= -2.25), F1: -32768, F2: 555}
    if true {
        var v57 int = 42
        _ = v57
    } else {
        v50 = append(v50, 433)
    }
    _ = v50
    _ = v56
    return 2147483648
}

func main() {
    var v1 uint8 = 128
    v2 := []int{0}
    v3 := make(map[uint32]uint32)
    v4 := []string{"", ("a" + "hello"), "\n\t"}
    v58 := GenS6{F0: 32767}
    if false {
        v59 := GenS7{F0: 258, F1: -1, F2: -405}
        _ = v59
    } else {
        for fi60 := 0; fi60 < 20; fi60++ {
            var v61 uint32 = 42
            var v62 uint = 762
            for fi63 := 0; fi63 < 10; fi63++ {
                fmt.Println(v1, v62)
                delete(v3, (v61 % ((1) | 1)))
                _ = fi63
            }
            _ = fi60
            _ = v61
            _ = v62
        }
        if true {
            v1 = 42
            for fi64 := 0; fi64 < 3; fi64++ {
                v1 = 254
                var v65 int16 = 0
                _ = fi64
                _ = v65
            }
        } else {
            fmt.Println(v1)
            fmt.Println(v1)
        }
    }
    for fi66 := 0; fi66 < 14; fi66++ {
        var v67 float64 = ((-2.25 * (1e10 + 0.1)) + 3.14159)
        var v68 uint8 = (v1 & (v1 % ((v1) | 1)))
        _ = fi66
        _ = v67
        _ = v68
    }
    var v69 uint64 = 2
    if len(v4) > 0 {
        v4[0] = ("\n\t" + "\n\t")
    }
    var v70 string = "a"
    if (127 >= -2147483648) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v3[2147483648] = 2147483648
        var v71 int32 = -2147483648
        _ = v71
    } else {
        for fi72 := 0; fi72 < 16; fi72++ {
            var v73 int64 = 0
            _ = fi72
            _ = v73
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v74 float32 = ((((0.0 + 1.5) * 0.1) / (((1.5 - ((1.5 * 1.5) + -2.25)) / -2.25) + float32(1.5))) + float32((0.1 - (0.1 * 1e10))))
        _ = v74
    }
    v75 := GenS7{F0: 1, F1: 0, F2: -2147483648}
    v76 := GenS8{F0: 1, F1: true}
    for fi77 := 0; fi77 < 7; fi77++ {
        switch fi77 {
        case 128:
            var v78 int16 = -32768
            _ = v78
        case 140:
            var v79 int32 = -2147483648
            _ = v79
        }
        _ = fi77
    }
    fmt.Println(v69, v69, v1)
    v69 -= (v69 ^ 4294967295)
    for fi80 := 0; fi80 < 16; fi80++ {
        fmt.Println(v70, v1)
        var v81 int32 = 1
        _ = fi80
        _ = v81
    }
    v82 := v3[uint32(v69)]
    v82 += 260
    if true {
        if false {
            v70 = "\n\t"
            if true {
                v83 := v3[(^(v82))]
                _ = v83
            } else {
                v84 := GenS9{F0: -85, F1: (v1 - uint8(v69)), F2: ((v82 / (((v82 % (((v82 * genFn1())) | 1))) | 1)) <= v82)}
                v82 -= uint32(v1)
                _ = len(v2)
                _ = v84
            }
            var v85 int = 9223372036854775807
            _ = v85
        } else {
            v82 = v82
            switch ((v70 + v70) + "foo bar") {
            case "a":
                var v86 int32 = -1
                _ = v86
            }
        }
        var v87 int64 = int64(v69)
        for fi88 := 0; fi88 < 19; fi88++ {
            delete(v3, 0)
            v89 := GenS9{F0: 1, F1: 128, F2: ("a" >= (v70 + (v70 + v70)))}
            if true {
                var v90 int8 = -128
                _ = v90
            }
            _ = fi88
            _ = v89
        }
        _ = v87
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v58
    _ = v69
    _ = v70
    _ = v75
    _ = v76
    _ = v82
}
