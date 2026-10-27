package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 string; F1 uint8; F2 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 int32; F1 uint8; F2 int16 }
type GenS6 struct { F0 bool; F1 string }
type GenS7 struct { F0 int; F1 uint8; F2 int }
type GenS8 struct { F0 uint16; F1 int64 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int; F1 string }
func genFn0(p0 uint8, p1 uint8) uint {
    var v5 int16 = 116
    v6 := GenS1{F0: 1, F1: 255}
    fmt.Println(p1)
    v7 := GenS1{F0: -1, F1: 18446744073709551615}
    for fi8 := 0; fi8 < 1; fi8++ {
        p0 = p1
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 18446744073709551615
}
func genFn1() uint8 {
    var v9 uint64 = 30
    fmt.Println(v9)
    if true {
        var v10 uint64 = v9
        if (0 < 32767) {
            for fi11 := 0; fi11 < 7; fi11++ {
                for fi12 := 0; fi12 < 7; fi12++ {
                    if (127 < 127) {
                        switch "" {
                        case "":
                            v10 = v10
                        }
                        if (3.14159 >= 0.1) {
                            if (!(0.0 >= ((1e10 / 0.1) - 1.5))) {
                                v13 := []int{}
                                var v14 int16 = 42
                                var v15 uint16 = 32767
                                fmt.Println(fi12)
                                v16 := GenS1{F0: -128, F1: 0}
                                _ = v13
                                _ = v14
                                _ = v15
                                _ = v16
                            } else {
                                if true {
                                    var v17 int32 = 522
                                    var v18 int32 = (-(v17))
                                    _ = v17
                                    _ = v18
                                } else {
                                    var v19 string = "hello"
                                    if true {
                                        v10 = (v10 ^ v10)
                                        v20 := []int{}
                                        v21 := GenS1{F0: 1, F1: genFn0(42, 222)}
                                        _ = v20
                                        _ = v21
                                    }
                                    _ = v19
                                }
                            }
                            var v22 uint32 = 312
                            _ = v22
                        }
                        v23 := []int{}
                        _ = v23
                    } else {
                        var v24 string = ""
                        switch fi12 {
                        case -9223372036854775808:
                            var v25 uint32 = 2
                            _ = v25
                        }
                        _ = v24
                    }
                    v26 := GenS2{F0: ("hello" + ("\n\t" + ("hello" + "foo bar"))), F1: 128, F2: (fi12 ^ (-(fi12)))}
                    if (1 > 620) {
                        v27 := GenS3{F0: "a"}
                        fmt.Println(fi11, v10)
                        _ = v27
                    }
                    v10 -= v9
                    v10 = (v9 % ((v10) | 1))
                    if (0 != 2) {
                        switch "a" {
                        case "\n\t":
                            v28 := GenS4{F0: 2147483647, F1: (^(v10))}
                            _ = v28
                        case "a":
                            var v29 int8 = -63
                            _ = v29
                        }
                        fmt.Println(fi12)
                    } else {
                        switch (("hello" + "") + "hello") {
                        case "hello":
                            var v30 int64 = 0
                            _ = v30
                        default:
                            v9 -= 20
                        }
                    }
                    _ = fi12
                    _ = v26
                }
                _ = fi11
            }
            v31 := GenS5{F0: -120, F1: 89, F2: 1}
            _ = v31
        } else {
            v32 := GenS6{F0: false, F1: "hello"}
            _ = v32
        }
        _ = v10
    } else {
        v33 := GenS4{F0: 42, F1: v9}
        if (9223372036854775807 > 0) {
            switch 4294967295 {
            case 42:
                var v34 uint = genFn0(255, 9)
                _ = v34
            case -9223372036854775808:
                for fi35 := 0; fi35 < 19; fi35++ {
                    var v36 uint32 = 2147483647
                    switch (("\n\t" + "foo bar") + ("\n\t" + "\n\t")) {
                    case "":
                        for fi37 := 0; fi37 < 4; fi37++ {
                            v38 := []int{fi37}
                            if (true || false) {
                                var v39 int = 4294967295
                                switch ("" + "hello") {
                                case "\n\t":
                                    if true {
                                        v39 = 128
                                        if (false && false) {
                                            var v40 string = (("" + "\n\t") + (("foo bar" + ("\n\t" + "a")) + ("hello" + "a")))
                                            if true {
                                                var v41 bool = (32767 <= 0)
                                                var v42 bool = (!v41)
                                                switch v40 {
                                                case "":
                                                    v40 = "\n\t"
                                                case "foo bar":
                                                    var v43 uint16 = 655
                                                    _ = v43
                                                default:
                                                    var v44 uint = 4294967295
                                                    _ = v44
                                                }
                                                _ = v41
                                                _ = v42
                                            }
                                            switch fi37 {
                                            case 1:
                                                _ = len(v38)
                                            case 419:
                                                var v45 string = v40
                                                _ = v45
                                            default:
                                                var v46 uint32 = 2
                                                _ = v46
                                            }
                                            _ = v40
                                        } else {
                                            if ((((0.0 * 0.1) != 0.0) && false) && (127 >= -1)) {
                                                v47 := GenS1{F0: 29, F1: 42}
                                                _ = len(v38)
                                                _ = v47
                                            }
                                            v36 += (v36 | uint32(v9))
                                        }
                                        var v48 int64 = -309
                                        _ = v48
                                    }
                                default:
                                    var v49 uint16 = 786
                                    _ = v49
                                }
                                var v50 bool = false
                                _ = v39
                                _ = v50
                            } else {
                                v9 += (v9 ^ (v9 >> ((42) & 63)))
                                v9 -= v9
                            }
                            if true {
                                fmt.Println(fi35)
                                var v51 float64 = float64((1e10 / (-2.25 - (float32(-2.25) / 0.1))))
                                _ = v51
                            }
                            if false {
                                var v52 int8 = 42
                                var v53 uint16 = 255
                                _ = v52
                                _ = v53
                            }
                            if true {
                                v9 += (v9 * v9)
                            } else {
                                var v54 string = ""
                                _ = v54
                            }
                            _ = fi37
                            _ = v38
                        }
                    }
                    _ = fi35
                    _ = v36
                }
            case 4294967295:
                for fi55 := 0; fi55 < 17; fi55++ {
                    v9 = v9
                    _ = fi55
                }
            default:
                var v56 uint8 = 2
                _ = v56
            }
        } else {
            v9 = (v9 - (v9 - v9))
            var v57 uint = 65535
            _ = v57
        }
        switch 127 {
        case -9223372036854775808:
            v9 += v9
        default:
            fmt.Println(v9)
        }
        _ = v33
    }
    fmt.Println(v9)
    _ = v9
    return 250
}

func main() {
    var v1 int = -968
    var v2 int32 = -348
    var v3 uint64 = 1
    var v4 int16 = -392
    fmt.Println(v1, v3)
    if (155 >= 255) {
        if (v4 >= -32768) {
            switch "foo bar" {
            case "a":
                for fi58 := 0; fi58 < 20; fi58++ {
                    v59 := []int{}
                    if false {
                        fmt.Println(v1)
                        v60 := GenS7{F0: (fi58 / (((v1 & v1)) | 1)), F1: genFn1(), F2: (v1 << ((0) & 63))}
                        var v61 float32 = (3.14159 / (float32((0.0 - 1.5)) + 1e10))
                        _ = v60
                        _ = v61
                    }
                    if len(v59) >= 2 {
                        v59 = v59[1:len(v59)]
                    }
                    _ = fi58
                    _ = v59
                }
            default:
                var v62 float32 = -2.25
                _ = v62
            }
            fmt.Println(v3, v2, v4)
        } else {
            var v63 uint64 = v3
            _ = v63
        }
    } else {
        fmt.Println(v4)
    }
    v64 := GenS2{F0: ((("" + "") + "hello") + "\n\t"), F1: 2, F2: v1}
    var v65 uint8 = 128
    for fi66 := 0; fi66 < 16; fi66++ {
        if (uint8(v3) >= (v65 * genFn1())) {
            var v67 uint64 = (v3 & v3)
            for fi68 := 0; fi68 < 4; fi68++ {
                switch 32768 {
                case 32767:
                    switch fi68 {
                    case 32767:
                        v67 += v3
                    case 32768:
                        if true {
                            switch (fi66 & (^(fi66))) {
                            case 0:
                                v69 := []int{0, 1}
                                _ = v69
                            case 32768:
                                v3 -= v67
                                var v70 float64 = 1e10
                                if false {
                                    var v71 int64 = 1
                                    fmt.Println(v4)
                                    _ = v71
                                } else {
                                    fmt.Println(v3)
                                }
                                _ = v70
                            }
                        } else {
                            v65 = uint8(v67)
                        }
                    default:
                        fmt.Println(v65, v4)
                    }
                case -167:
                    v3 += 4294967295
                }
                v72 := []int{(fi66 / ((v1) | 1)), (fi68 + -95), -603, 787}
                var v73 float32 = 0.1
                switch v1 {
                case 42:
                    var v74 uint16 = 714
                    _ = v74
                case 127:
                    v75 := GenS8{F0: 625, F1: 1}
                    _ = v75
                default:
                    var v76 uint = 255
                    _ = v76
                }
                var v77 int16 = v4
                _ = fi68
                _ = v72
                _ = v73
                _ = v77
            }
            _ = v67
        }
        v78 := GenS2{F0: "\n\t", F1: v65, F2: (fi66 >> ((genFn0(v65, genFn1())) & 63))}
        _ = fi66
        _ = v78
    }
    fmt.Println(v1, v2)
    v79 := GenS9{F0: true}
    for fi80 := 0; fi80 < 13; fi80++ {
        for fi81 := 0; fi81 < 20; fi81++ {
            var v82 uint32 = 542
            _ = fi81
            _ = v82
        }
        if false {
            for fi83 := 0; fi83 < 14; fi83++ {
                if (1 == 294) {
                    v84 := GenS9{F0: ((!true) || (127 < 127))}
                    if (42 < (v4 * 0)) {
                        var v85 uint32 = 42
                        v3 = (v3 | v3)
                        v86 := GenS1{F0: -128, F1: 4294967295}
                        _ = v85
                        _ = v86
                    }
                    _ = v84
                } else {
                    v2 = -38
                    var v87 uint = 65535
                    _ = v87
                }
                if true {
                    v65 = 1
                    v88 := GenS10{F0: fi83, F1: ""}
                    fmt.Println(fi80, fi80)
                    _ = v88
                } else {
                    v65 += v65
                    var v89 bool = true
                    var v90 int32 = (v2 ^ -2)
                    _ = v89
                    _ = v90
                }
                _ = fi83
            }
        }
        v1 = (fi80 + (v1 + (-(fi80))))
        _ = fi80
    }
    v91 := []int{4294967295, 4294967295, (^(v1)), (v1 + v1)}
    var v92 int32 = (v2 << ((65535) & 63))
    switch v1 {
    case 127:
        var v93 int = v1
        _ = v93
    default:
        if (false || true) {
            var v94 int16 = int16(v65)
            _ = v94
        }
    }
    v1 = (v1 % (((v1 * (-(v1)))) | 1))
    switch v1 {
    case 42:
        fmt.Println(v1)
    case 32767:
        v1 += v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v65:", v65)
    fmt.Println("slicelen v91:", len(v91), v91)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v64
    _ = v65
    _ = v79
    _ = v91
    _ = v92
}
