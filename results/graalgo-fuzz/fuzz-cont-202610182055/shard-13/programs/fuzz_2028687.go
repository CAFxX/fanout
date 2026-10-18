package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 uint8; F2 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int; F1 uint32; F2 uint }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint; F1 int32; F2 uint }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int64; F1 bool; F2 int16 }
type GenS11 struct { F0 int8; F1 int8; F2 int }
type GenS12 struct { F0 uint64; F1 uint16; F2 int16 }
type GenS13 struct { F0 uint32; F1 string; F2 int16 }
type GenS14 struct { F0 uint32; F1 int64 }
func genFn0(p0 int64, p1 uint) int64 {
    p0 = (p0 >> ((736) & 63))
    switch -560 {
    case 9223372036854775807:
        v6 := []int{9223372036854775807}
        _ = v6
    case 32767:
        var v7 uint16 = 658
        _ = v7
    default:
        if ((3.14159 + (((3.14159 / (float64(-2.25) - 1e10)) - 0.1) - -2.25)) < ((0.1 - -2.25) + (3.14159 * (0.1 - 1.5)))) {
            var v8 int64 = (p0 % (((p0 * (p0 - (p0 << (((p1 * p1)) & 63))))) | 1))
            _ = v8
        } else {
            p1 -= (p1 + p1)
            fmt.Println(p1, p1)
        }
    }
    var v9 bool = ("a" < ("\n\t" + "hello"))
    _ = p0
    _ = p1
    _ = v9
    return (p0 >> (((p1 + (p1 | 0))) & 63))
}
func genFn1(p0 uint8) int16 {
    for fi10 := 0; fi10 < 6; fi10++ {
        v11 := GenS1{F0: fi10}
        _ = fi10
        _ = v11
    }
    var v12 int16 = 1000
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 int16 = v12
        switch ("\n\t" + "") {
        case "a":
            v15 := []int{(fi13 | fi13)}
            _ = v15
        }
        for fi16 := 0; fi16 < 8; fi16++ {
            p0 = p0
            if (fi16 <= fi13) {
                v17 := []int{127, (^(fi16)), (fi13 + (fi16 / ((128) | 1)))}
                _ = v17
            }
            fmt.Println(v12, p0)
            _ = fi16
        }
        _ = fi13
        _ = v14
    }
    if true {
        for fi18 := 0; fi18 < 6; fi18++ {
            var v19 uint64 = 4294967295
            if ((v12 + 0) <= -32768) {
                v20 := []int{(fi18 & fi18), (fi18 | fi18)}
                _ = v20
            } else {
                p0 += 0
            }
            if false {
                v21 := []int{fi18, fi18, (fi18 * (^(fi18)))}
                fmt.Println(p0, fi18, fi18)
                var v22 int = fi18
                _ = v21
                _ = v22
            }
            _ = fi18
            _ = v19
        }
    } else {
        p0 = (^(p0))
    }
    v23 := GenS2{F0: "hello", F1: 2, F2: -1}
    _ = p0
    _ = v12
    _ = v23
    return v12
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{}
    var v3 uint8 = 132
    v4 := []string{("a" + "foo bar")}
    v5 := make(map[uint32]bool)
    var v24 int32 = -1
    var v25 uint8 = (^(v3))
    var v26 int8 = 42
    for fi27 := 0; fi27 < 17; fi27++ {
        var v28 uint = 700
        if false {
            v29 := GenS3{F0: (v26 & v26)}
            for fi30 := 0; fi30 < 4; fi30++ {
                if true {
                    v31 := GenS4{F0: fi27, F1: 684, F2: (v28 - (v28 >> ((4294967295) & 63)))}
                    if len(v2) > 0 {
                        v2[0] = "\n\t"
                    }
                    _ = v31
                } else {
                    if len(v4) > 0 {
                        v4[0] = "foo bar"
                    }
                    switch fi27 {
                    case 128:
                        delete(v5, 69)
                    case -923:
                        var v32 int32 = (v24 << (((v28 & uint(v3))) & 63))
                        _ = v32
                    }
                    var v33 uint = v28
                    v3 -= 42
                    if true {
                        v24 += int32(v33)
                        if (fi30 != -80) {
                            var v34 uint16 = 255
                            _ = v34
                        }
                    } else {
                        v3 = (^(v3))
                    }
                    if false {
                        var v35 bool = (float64(1.5) < (0.0 * ((float64(0.1) - 3.14159) / 1.5)))
                        v33 = 2
                        _ = len(v2)
                        _ = v35
                    }
                    _ = v33
                }
                v24 = v24
                var v36 uint64 = 2
                _ = fi30
                _ = v36
            }
            _ = v29
        }
        switch "foo bar" {
        case "a":
            v28 += (^(v28))
        case "":
            for fi37 := 0; fi37 < 11; fi37++ {
                _ = len(v2)
                switch ("hello" + "a") {
                case "":
                    if ((!true) || (42 <= 0)) {
                        if (true || true) {
                            if (2 != 0) {
                                v38 := GenS4{F0: (fi27 - (fi27 ^ fi37)), F1: 0, F2: (v28 % (((^(v28))) | 1))}
                                if ((v26 * (v26 - 1)) == int8(v24)) {
                                    var v39 float32 = -2.25
                                    fmt.Println(v26, v3)
                                    _ = v39
                                } else {
                                    fmt.Println(v3, v3)
                                    var v40 float32 = 1.5
                                    _ = v40
                                }
                                _ = v38
                            }
                            fmt.Println(v28, v28)
                            v41 := GenS4{F0: (fi27 / ((fi27) | 1)), F1: 4294967295, F2: 65535}
                            if (!false) {
                                v25 -= (v3 | 0)
                                v42 := GenS3{F0: v26}
                                _ = v42
                            } else {
                                if true {
                                    var v43 uint32 = 437
                                    var v44 uint64 = 42
                                    if false {
                                        var v45 bool = false
                                        switch (fi37 + fi37) {
                                        case 1:
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            if (true || (619 < 2147483648)) {
                                                var v46 string = ("hello" + ("a" + ""))
                                                v44 -= v44
                                                _ = v46
                                            } else {
                                                v47 := GenS3{F0: (^(v26))}
                                                _ = v47
                                            }
                                            var v48 float64 = float64((-2.25 * ((3.14159 / (0.1 - (0.0 - 1.5))) * 1.5)))
                                            _ = v48
                                        case 4294967295:
                                            v24 = v24
                                        case -257:
                                            if len(v4) > 0 {
                                                v4[0] = ("foo bar" + "\n\t")
                                            }
                                        default:
                                            var v49 string = ((("a" + "a") + "\n\t") + ("hello" + "\n\t"))
                                            _ = v49
                                        }
                                        delete(v5, 1)
                                        _ = v45
                                    } else {
                                        v50 := GenS4{F0: 414, F1: v43, F2: 255}
                                        var v51 float64 = 1.5
                                        fmt.Println(fi37, v44)
                                        _ = v50
                                        _ = v51
                                    }
                                    _ = v43
                                    _ = v44
                                } else {
                                    v52 := GenS5{F0: 4294967295}
                                    _ = v52
                                }
                                switch (("\n\t" + ("hello" + ("" + "a"))) + ("" + "")) {
                                case "\n\t":
                                    if true {
                                        if (!(65535 != 42)) {
                                            fmt.Println(fi37, v3, fi27)
                                            v53 := GenS6{F0: v28, F1: v24, F2: 255}
                                            if (true && true) {
                                                v54 := GenS5{F0: (^(v28))}
                                                v55 := GenS5{F0: v28}
                                                _ = v54
                                                _ = v55
                                            } else {
                                                v3 -= v25
                                                v24 = (-(v24))
                                            }
                                            switch "a" {
                                            case "":
                                                fmt.Println(v28, v24)
                                            case "a":
                                                var v56 float32 = 3.14159
                                                _ = v56
                                            default:
                                                if true {
                                                    v57 := GenS4{F0: (fi37 << (((v28 & 42)) & 63)), F1: 438, F2: 2}
                                                    var v58 int32 = 2147483647
                                                    var v59 int16 = 32767
                                                    _ = v57
                                                    _ = v58
                                                    _ = v59
                                                } else {
                                                    v60 := GenS7{F0: (v25 - v25)}
                                                    var v61 int64 = -1
                                                    var v62 float32 = float32(1.5)
                                                    _ = v60
                                                    _ = v61
                                                    _ = v62
                                                }
                                            }
                                            v26 = 9
                                            _ = v53
                                        }
                                        var v63 int64 = 42
                                        if (false || false) {
                                            if len(v4) > 0 {
                                                v4[0] = "hello"
                                            }
                                            v64 := GenS1{F0: fi27}
                                            _ = v64
                                        } else {
                                            var v65 int64 = v63
                                            _ = v65
                                        }
                                        _ = v63
                                    } else {
                                        v66 := GenS8{F0: 460}
                                        switch (fi37 * -667) {
                                        case 2147483647:
                                            v3 -= (v25 << (((v28 >> (((v28 - v28)) & 63))) & 63))
                                        }
                                        v5[67] = false
                                        _ = v66
                                    }
                                    v25 = v25
                                    v67 := GenS9{F0: 4294967295}
                                    if (-1 != v24) {
                                        switch ("" + "foo bar") {
                                        case "\n\t":
                                            v68 := GenS10{F0: -1, F1: false, F2: int16(v3)}
                                            _ = v68
                                        }
                                        v26 -= v26
                                        if (!false) {
                                            var v69 uint32 = 142
                                            _ = v69
                                        } else {
                                            switch (fi37 & fi37) {
                                            case 128:
                                                v3 -= (^(v25))
                                            }
                                            v70 := GenS11{F0: 85, F1: v26, F2: (^(fi37))}
                                            v3 -= v25
                                            delete(v5, 246)
                                            _ = v70
                                        }
                                    } else {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        v71 := GenS8{F0: 146}
                                        _ = v71
                                    }
                                    fmt.Println(v3)
                                    _ = v67
                                case "hello":
                                    v72 := GenS2{F0: ("" + "a"), F1: v3, F2: 0}
                                    _ = v72
                                default:
                                    switch "hello" {
                                    case "\n\t":
                                        v25 = v3
                                    case "":
                                        switch fi37 {
                                        case -88:
                                            var v73 int = fi27
                                            _ = v73
                                        }
                                    case "a":
                                        var v74 float64 = 3.14159
                                        _ = v74
                                    }
                                    v75 := v5[2147483648]
                                    _ = v75
                                }
                            }
                            _ = v41
                        }
                        fmt.Println(v24, v25, fi27)
                        if (("hello" <= ("a" + "foo bar")) || false) {
                            if (v24 < (^(v24))) {
                                v76 := GenS7{F0: 95}
                                v2 = append(v2, ("" + "hello"))
                                _ = v76
                            } else {
                                v25 -= v3
                                v26 -= (^(v26))
                            }
                        }
                        v77 := GenS2{F0: "\n\t", F1: (v25 * v25), F2: 1}
                        _ = v77
                    } else {
                        delete(v5, 65535)
                        v78 := GenS11{F0: int8(v28), F1: (^(v26)), F2: (fi27 * -202)}
                        v1[65535] = 2
                        _ = v78
                    }
                case "hello":
                    var v79 int16 = int16(v3)
                    fmt.Println(v26, v25, v79)
                    v79 += genFn1(v25)
                    _ = v79
                default:
                    fmt.Println(v24)
                }
                v24 = 957
                _ = fi37
            }
        }
        _ = fi27
        _ = v28
    }
    v80 := GenS2{F0: "hello", F1: (^(v25)), F2: genFn0(-394, 4294967295)}
    switch 127 {
    case 323:
        v81 := GenS12{F0: 4294967295, F1: 2, F2: genFn1(v25)}
        _ = v81
    }
    fmt.Println(v25, v26, v24)
    v82 := GenS9{F0: 910}
    var v83 bool = (true && (0 != 634))
    v84 := GenS3{F0: v26}
    v83 = ((v83 && v83) && v83)
    var v85 int = 255
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v86 := GenS13{F0: 2147483647, F1: ("hello" + ("" + "")), F2: -32768}
    for fi87 := 0; fi87 < 15; fi87++ {
        v26 -= 0
        if (3.14159 == float64(0.1)) {
            v88 := GenS10{F0: 42, F1: (false || true), F2: 32767}
            fmt.Println(v24, v85)
            switch "a" {
            case "":
                v89 := GenS14{F0: 65535, F1: 9223372036854775807}
                _ = v89
            default:
                v90 := GenS12{F0: 636, F1: 42, F2: -32768}
                _ = v90
            }
            _ = v88
        } else {
            var v91 bool = v83
            var v92 float32 = float32((((-2.25 - 1.5) / 1.0) - (3.14159 + (float64(-2.25) - 1.5))))
            var v93 float64 = (0.0 - (float64(v92) + 0.1))
            _ = v91
            _ = v92
            _ = v93
        }
        var v94 float64 = 1.5
        _ = fi87
        _ = v94
    }
    v25 -= (v25 * 127)
    for fi95 := 0; fi95 < 11; fi95++ {
        fmt.Println(v3, v24)
        var v96 int16 = 42
        _ = fi95
        _ = v96
    }
    var v97 int = int(v24)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v83:", v83)
    fmt.Println("v85:", v85)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v80
    _ = v82
    _ = v83
    _ = v84
    _ = v85
    _ = v86
    _ = v97
}
