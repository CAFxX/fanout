package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int64 }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int64; F1 uint8; F2 int }
type GenS4 struct { F0 int; F1 string; F2 int16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS7 struct { F0 bool; F1 int }
type GenS8 struct { F0 int; F1 int32; F2 uint8 }
type GenS9 struct { F0 uint16; F1 string; F2 uint16 }
type GenS10 struct { F0 uint16; F1 int64; F2 int16 }
func genFn0() int64 {
    v4 := []int{0, -78, 0}
    if (0 != 100) {
        var v5 float32 = 3.14159
        var v6 uint8 = 2
        _ = v5
        _ = v6
    } else {
        var v7 int = -747
        if true {
            v8 := GenS1{F0: 584, F1: 1, F2: -415}
            _ = v8
        }
        if (true && (0 < 354)) {
            fmt.Println(v7)
            v7 = (v7 * v7)
        } else {
            if (2 != 1) {
                if false {
                    if len(v4) > 0 {
                        v4[0] = (v7 & (v7 << ((0) & 63)))
                    }
                    v4 = append(v4, (v7 & (v7 << ((18446744073709551615) & 63))))
                    v9 := GenS2{F0: 2147483648, F1: 160}
                    _ = v9
                } else {
                    if false {
                        var v10 int32 = -2147483648
                        fmt.Println(v10, v10)
                        _ = v10
                    }
                    for fi11 := 0; fi11 < 12; fi11++ {
                        _ = len(v4)
                        _ = fi11
                    }
                }
                switch "a" {
                case "foo bar":
                    _ = len(v4)
                default:
                    fmt.Println(v7)
                }
            }
            if ((2147483648 <= 4294967295) && true) {
                v7 -= 1
                v12 := GenS2{F0: 870, F1: 222}
                v7 += (v7 | v7)
                _ = v12
            } else {
                _ = len(v4)
                v7 = (v7 ^ (v7 * 127))
                v7 += 128
            }
        }
        _ = v7
    }
    var v13 string = "hello"
    _ = v13
    _ = v4
    return 1
}
func genFn1(p0 uint16) int {
    if (true || false) {
        v14 := GenS1{F0: p0, F1: 1, F2: 554}
        var v15 uint32 = 123
        v16 := []int{-322}
        _ = v14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        var v18 uint64 = 18446744073709551615
        v18 = 0
        fmt.Println(fi17, p0, fi17)
        _ = fi17
        _ = v18
    }
    p0 = 32767
    _ = p0
    return -756
}

func main() {
    var v1 string = ("foo bar" + ("hello" + ""))
    v2 := make(map[uint32]string)
    var v3 string = v1
    for fi19 := 0; fi19 < 8; fi19++ {
        fmt.Println(v1, fi19, v1)
        if false {
            if true {
                v20 := []int{fi19, genFn1(833), 1, genFn1(65535)}
                _ = v20
            }
            v21 := GenS3{F0: 0, F1: 255, F2: (fi19 * -22)}
            if true {
                if true {
                    fmt.Println(v1)
                    v1 = v3
                } else {
                    v22 := GenS3{F0: genFn0(), F1: 0, F2: (fi19 * (fi19 >> ((65535) & 63)))}
                    _ = v22
                }
                var v23 bool = (9223372036854775807 <= 696)
                _ = v23
            } else {
                fmt.Println(fi19)
                if true {
                    v1 = (((v3 + v1) + v1) + (("" + v1) + v1))
                    v24 := GenS1{F0: 2, F1: 1000, F2: 9223372036854775807}
                    v25 := []int{}
                    _ = v24
                    _ = v25
                }
                v26 := GenS4{F0: 0, F1: "foo bar", F2: 1}
                _ = v26
            }
            _ = v21
        } else {
            for fi27 := 0; fi27 < 20; fi27++ {
                var v28 int8 = 0
                _ = len(v2)
                v29 := []int{(-(fi27)), (fi19 + (fi27 & (fi19 * fi19))), (^(fi27)), (fi19 >> ((0) & 63))}
                _ = fi27
                _ = v28
                _ = v29
            }
        }
        _ = fi19
    }
    v30 := []int{-9223372036854775808, 32}
    switch v3 {
    case "foo bar":
        for fi31 := 0; fi31 < 3; fi31++ {
            for fi32 := 0; fi32 < 6; fi32++ {
                switch genFn1(65535) {
                case 4294967295:
                    switch v3 {
                    case "a":
                        switch v1 {
                        case "hello":
                            v33 := GenS3{F0: -1, F1: 1, F2: (fi32 << ((16) & 63))}
                            _ = v33
                        default:
                            v1 = v3
                        }
                    case "\n\t":
                        switch (("a" + (v1 + v3)) + ((v1 + (v3 + "a")) + v1)) {
                        case "":
                            fmt.Println(fi32, fi31, fi31)
                        case "a":
                            var v34 int8 = 42
                            _ = v34
                        default:
                            fmt.Println(v1)
                        }
                    default:
                        fmt.Println(fi32)
                    }
                default:
                    fmt.Println(v3, fi31, v1)
                    fmt.Println(fi32)
                    if (-1 == -128) {
                        if true {
                            v1 = v1
                            v3 = "a"
                            var v35 int = fi31
                            _ = v35
                        } else {
                            _ = len(v2)
                            v36 := GenS5{F0: -9223372036854775808}
                            _ = v36
                        }
                        switch fi31 {
                        case -1:
                            v3 = ""
                        case 744:
                            var v37 float32 = ((0.1 / 1.0) - (((1e10 / float32(0.1)) + 1e10) - 0.0))
                            _ = v37
                        default:
                            var v38 int64 = 0
                            _ = v38
                        }
                    } else {
                        if (469 != 4294967295) {
                            if false {
                                v3 = v1
                                var v39 uint8 = 255
                                _ = v39
                            }
                            switch v1 {
                            case "":
                                var v40 bool = (false || false)
                                _ = v40
                            }
                            v3 = v3
                        } else {
                            v30 = append(v30, (fi32 * (^(fi32))))
                        }
                    }
                }
                if ((float64(1e10) + (0.0 / 1.5)) < -2.25) {
                    switch v1 {
                    case "\n\t":
                        var v41 string = v1
                        _ = v41
                    default:
                        v42 := GenS3{F0: genFn0(), F1: 108, F2: fi31}
                        _ = v42
                    }
                } else {
                    v1 = v3
                    v3 = v3
                    switch v1 {
                    case "foo bar":
                        var v43 int8 = 1
                        _ = v43
                    }
                }
                v44 := GenS6{F0: 4294967295, F1: 42, F2: 532}
                _ = fi32
                _ = v44
            }
            v3 = v1
            _ = fi31
        }
    case "hello":
        var v45 int = 1
        _ = v45
    }
    if (4294967295 >= 4294967295) {
        var v46 bool = true
        for fi47 := 0; fi47 < 18; fi47++ {
            v48 := GenS2{F0: 416, F1: 1}
            var v49 uint = 2
            _ = fi47
            _ = v48
            _ = v49
        }
        for fi50 := 0; fi50 < 8; fi50++ {
            v46 = ((1e10 - -2.25) != (-2.25 + float32(-2.25)))
            _ = fi50
        }
        _ = v46
    } else {
        v1 = v3
        fmt.Println(v3)
    }
    if (!(!(true && true))) {
        var v51 int = -2147483648
        for fi52 := 0; fi52 < 6; fi52++ {
            v51 = (-(fi52))
            _ = fi52
        }
        if (false || (0 <= 566)) {
            switch "a" {
            case "a":
                var v53 int8 = 127
                _ = v53
            }
            fmt.Println(v1)
            for fi54 := 0; fi54 < 9; fi54++ {
                delete(v2, 352)
                for fi55 := 0; fi55 < 15; fi55++ {
                    fmt.Println(v51, v51, v3)
                    _ = len(v2)
                    var v56 uint = 255
                    _ = fi55
                    _ = v56
                }
                _ = fi54
            }
        } else {
            if (255 < 42) {
                delete(v2, 0)
                var v57 uint16 = 2
                v3 = v3
                _ = v57
            } else {
                if true {
                    switch "foo bar" {
                    case "foo bar":
                        v58 := GenS5{F0: -9223372036854775808}
                        _ = v58
                    }
                } else {
                    if len(v30) >= 2 {
                        v30 = v30[1:len(v30)]
                    }
                }
            }
        }
        _ = v51
    }
    fmt.Println(v3)
    v3 = v3
    var v59 string = (v1 + ("a" + "\n\t"))
    var v60 int64 = genFn0()
    v61 := GenS1{F0: 65535, F1: -605, F2: (v60 | v60)}
    var v62 int = genFn1(255)
    for fi63 := 0; fi63 < 7; fi63++ {
        var v64 string = v1
        _ = len(v30)
        v64 = (v59 + "\n\t")
        _ = fi63
        _ = v64
    }
    fmt.Println(v1)
    for fi65 := 0; fi65 < 18; fi65++ {
        switch ((v1 + ("" + "")) + "foo bar") {
        case "":
            if true {
                if true {
                    switch "hello" {
                    case "hello":
                        delete(v2, 2)
                    case "":
                        fmt.Println(v62, v1)
                    case "a":
                        fmt.Println(v1, v59, v3)
                    default:
                        if true {
                            switch fi65 {
                            case 0:
                                fmt.Println(v3, v1, v62)
                            case 1:
                                for fi66 := 0; fi66 < 5; fi66++ {
                                    var v67 int16 = -171
                                    _ = fi66
                                    _ = v67
                                }
                            case 362:
                                var v68 uint32 = 2147483648
                                _ = v68
                            }
                            var v69 string = ("foo bar" + v59)
                            if ((1 != 998) && (2 == 866)) {
                                fmt.Println(v60)
                                fmt.Println(fi65, v3)
                                for fi70 := 0; fi70 < 20; fi70++ {
                                    if (!(42 != -32768)) {
                                        var v71 bool = (((1 < 22) && false) || true)
                                        v69 = "\n\t"
                                        v1 = "a"
                                        if (1 >= 23) {
                                            if (!(!true)) {
                                                v72 := GenS2{F0: 499, F1: 65535}
                                                _ = v72
                                            }
                                            var v73 uint = 42
                                            switch v1 {
                                            case "hello":
                                                if (true && v71) {
                                                    v1 = v59
                                                } else {
                                                    var v74 int32 = -199
                                                    _ = v74
                                                }
                                            case "\n\t":
                                                var v75 uint16 = 580
                                                _ = v75
                                            case "foo bar":
                                                var v76 float64 = 1e10
                                                _ = v76
                                            }
                                            _ = v73
                                        }
                                        _ = v71
                                    } else {
                                        if (false || false) {
                                            v3 = v3
                                            v77 := GenS7{F0: (!((1e10 == -2.25) || true)), F1: fi70}
                                            v59 = "\n\t"
                                            _ = v77
                                        } else {
                                            v78 := GenS7{F0: false, F1: (fi65 * 256)}
                                            if (v60 > v60) {
                                                v69 = v59
                                                v60 -= (v60 - (v60 * -9223372036854775808))
                                                if (0 == 1) {
                                                    var v79 uint64 = 4294967295
                                                    fmt.Println(v3, v79, v1)
                                                    _ = v79
                                                }
                                            }
                                            v69 = v1
                                            _ = len(v30)
                                            v60 += 0
                                            _ = v78
                                        }
                                    }
                                    v59 = "\n\t"
                                    v3 = "a"
                                    _ = fi70
                                }
                            } else {
                                for fi80 := 0; fi80 < 6; fi80++ {
                                    var v81 bool = false
                                    fmt.Println(v1)
                                    _ = fi80
                                    _ = v81
                                }
                                v2[812] = "foo bar"
                            }
                            _ = v69
                        }
                    }
                }
                fmt.Println(v3)
                v82 := GenS7{F0: false, F1: fi65}
                _ = v82
            }
        case "hello":
            if true {
                var v83 uint = 126
                v83 = v83
                if (((3.14159 / 1.0) * 1e10) <= 1.5) {
                    switch v59 {
                    case "hello":
                        delete(v2, 2147483648)
                    case "":
                        v62 += (v62 + (-(v62)))
                    default:
                        fmt.Println(v83, v60)
                    }
                }
                _ = v83
            } else {
                var v84 bool = false
                fmt.Println(fi65, v3, v59)
                var v85 int = genFn1(255)
                _ = v84
                _ = v85
            }
        case "a":
            var v86 int32 = 0
            _ = v86
        default:
            if false {
                v87 := GenS8{F0: fi65, F1: -1, F2: 127}
                v88 := GenS8{F0: (^(fi65)), F1: 0, F2: 255}
                var v89 uint32 = 1
                _ = v87
                _ = v88
                _ = v89
            } else {
                fmt.Println(v60)
            }
        }
        for fi90 := 0; fi90 < 7; fi90++ {
            v59 = v59
            _ = fi90
        }
        _ = fi65
    }
    if false {
        for fi91 := 0; fi91 < 7; fi91++ {
            if ((((v3 + v3) + v1) <= "") && true) {
                v1 = v1
                var v92 bool = ((0.1 - 0.1) != -2.25)
                _ = v92
            } else {
                if len(v30) > 0 {
                    v30[0] = (v62 << ((1) & 63))
                }
                switch (v62 % ((148) | 1)) {
                case 127:
                    v2[864] = (v3 + v1)
                case 9223372036854775807:
                    v3 = (v1 + v3)
                case 4294967295:
                    v59 = v3
                }
                v1 = v1
            }
            v93 := GenS7{F0: true, F1: genFn1(2)}
            _ = fi91
            _ = v93
        }
    } else {
        switch v62 {
        case -2147483648:
            var v94 float64 = (0.0 - 3.14159)
            _ = v94
        case 9223372036854775807:
            if (1 <= 2) {
                fmt.Println(v59)
                _ = len(v30)
            }
        case -885:
            fmt.Println(v59, v3)
        }
    }
    v95 := GenS5{F0: (v60 % (((v60 % (((v60 >> ((769) & 63))) | 1))) | 1))}
    v96 := GenS3{F0: (v60 | (v60 % (((v60 % ((v60) | 1))) | 1))), F1: 1, F2: 127}
    if true {
        switch v62 {
        case 901:
            if false {
                v97 := GenS2{F0: 1, F1: 0}
                _ = len(v2)
                var v98 int8 = 1
                _ = v97
                _ = v98
            }
        case 690:
            v99 := GenS9{F0: 65535, F1: "", F2: 1}
            _ = v99
        case 32768:
            switch v3 {
            case "hello":
                v100 := GenS10{F0: 32767, F1: v60, F2: -32768}
                _ = v100
            }
        default:
            switch v59 {
            case "hello":
                _ = len(v2)
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v59
    _ = v60
    _ = v61
    _ = v62
    _ = v95
    _ = v96
}
