package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 int64; F1 bool; F2 uint32 }
type GenS5 struct { F0 uint64; F1 int32; F2 string }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint16; F1 int32; F2 int8 }
type GenS9 struct { F0 uint16; F1 uint }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int; F1 int32 }
type GenS12 struct { F0 uint; F1 int32 }
func genFn0(p0 int32, p1 int) uint16 {
    p1 += (p1 >> ((18446744073709551615) & 63))
    for fi5 := 0; fi5 < 5; fi5++ {
        fmt.Println(p0, p1, p1)
        v6 := GenS1{F0: p0, F1: 255}
        _ = fi5
        _ = v6
    }
    v7 := []int{442, 170, (p1 ^ (p1 % (((p1 % ((127) | 1))) | 1)))}
    if ("\n\t" < "a") {
        if ((^(p0)) < (^(p0))) {
            var v8 uint8 = 255
            var v9 bool = (100 > -65)
            if ((p1 < p1) && (965 < 9223372036854775807)) {
                v8 = v8
                fmt.Println(p1)
            } else {
                v10 := GenS2{F0: 451, F1: ("hello" + "\n\t")}
                fmt.Println(v9, p0, p0)
                _ = v10
            }
            _ = v8
            _ = v9
        }
        var v11 uint64 = 2
        _ = v11
    } else {
        switch (((("" + "hello") + "") + ("" + "")) + "a") {
        case "\n\t":
            fmt.Println(p1)
        case "":
            var v12 int8 = 97
            _ = v12
        case "a":
            var v13 float64 = 0.0
            _ = v13
        }
        v14 := GenS3{F0: 42, F1: 449}
        p1 = (p1 >> ((42) & 63))
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v7
    return 2
}
func genFn1() uint32 {
    var v15 int32 = -363
    var v16 string = ("" + "foo bar")
    _ = v15
    _ = v16
    return 65535
}

func main() {
    var v1 bool = ((107 > 1) && (-1 <= -1))
    var v2 uint8 = 255
    var v3 string = ("hello" + "")
    v4 := make(map[uint32]int)
    v2 += (v2 / ((0) | 1))
    v17 := GenS4{F0: 0, F1: v1, F2: 42}
    if (v1 || true) {
        var v18 uint32 = 42
        fmt.Println(v18)
        var v19 uint = 42
        _ = v18
        _ = v19
    } else {
        v1 = false
        var v20 int64 = 9223372036854775807
        _ = v20
    }
    switch 1 {
    case -146:
        switch "\n\t" {
        case "":
            for fi21 := 0; fi21 < 8; fi21++ {
                for fi22 := 0; fi22 < 9; fi22++ {
                    v23 := GenS5{F0: 18446744073709551615, F1: -98, F2: ""}
                    _ = fi22
                    _ = v23
                }
                if ((2 > 42) && (v1 && v1)) {
                    if (2 <= 42) {
                        v3 = (v3 + v3)
                    } else {
                        v3 = ""
                        var v24 float32 = 1e10
                        _ = v24
                    }
                    var v25 uint64 = 4294967295
                    fmt.Println(v25, v2)
                    _ = v25
                } else {
                    fmt.Println(v3)
                }
                v26 := []int{(^(fi21)), fi21, fi21}
                _ = fi21
                _ = v26
            }
        case "foo bar":
            var v27 int8 = 64
            _ = v27
        case "hello":
            var v28 int = -9223372036854775808
            _ = v28
        }
    case -479:
        v3 = v3
    case 127:
        v1 = true
    default:
        v1 = false
    }
    var v29 string = ""
    switch (v29 + v3) {
    case "":
        for fi30 := 0; fi30 < 18; fi30++ {
            for fi31 := 0; fi31 < 14; fi31++ {
                v2 += (^(v2))
                if false {
                    var v32 uint16 = 2
                    if (-2147483648 > 95) {
                        v3 = (v3 + (((v29 + (v29 + v29)) + v3) + ((v3 + (v3 + v3)) + v29)))
                        v29 = "a"
                        var v33 int64 = 315
                        if false {
                            v34 := v4[42]
                            var v35 uint8 = v2
                            switch v3 {
                            case "a":
                                fmt.Println(v1, v2)
                            }
                            var v36 int32 = 1
                            v34 = fi31
                            _ = v34
                            _ = v35
                            _ = v36
                        } else {
                            fmt.Println(fi31, fi31)
                        }
                        _ = v33
                    }
                    _ = v32
                } else {
                    if (genFn1() != 639) {
                        v37 := GenS2{F0: 0, F1: ("foo bar" + v3)}
                        fmt.Println(fi30, v3)
                        _ = v37
                    } else {
                        var v38 float64 = ((3.14159 / -2.25) - 0.1)
                        _ = v38
                    }
                }
                _ = fi31
            }
            _ = fi30
        }
    default:
        for fi39 := 0; fi39 < 2; fi39++ {
            if v1 {
                v3 = v3
                v1 = false
            }
            v40 := []int{fi39, fi39}
            _ = fi39
            _ = v40
        }
    }
    for fi41 := 0; fi41 < 13; fi41++ {
        v42 := GenS4{F0: 264, F1: ((42 > 42) && (!(0 > 42))), F2: 4294967295}
        _ = fi41
        _ = v42
    }
    var v43 int64 = -9223372036854775808
    for fi44 := 0; fi44 < 17; fi44++ {
        for fi45 := 0; fi45 < 16; fi45++ {
            var v46 uint16 = 840
            var v47 bool = ((("\n\t" + v3) > v29) || (true && (1 < 2147483647)))
            _ = fi45
            _ = v46
            _ = v47
        }
        _ = fi44
    }
    v48 := GenS4{F0: (v43 >> ((679) & 63)), F1: v1, F2: 4294967295}
    v1 = (v2 == 1)
    if (65535 == 318) {
        v49 := v4[4294967295]
        v50 := GenS6{F0: (!v1)}
        for fi51 := 0; fi51 < 2; fi51++ {
            var v52 uint64 = 2
            if v1 {
                switch v29 {
                case "\n\t":
                    switch "\n\t" {
                    case "a":
                        v49 = 32768
                    case "\n\t":
                        v52 = 42
                    case "foo bar":
                        switch (fi51 ^ 0) {
                        case 127:
                            for fi53 := 0; fi53 < 19; fi53++ {
                                v54 := GenS7{F0: 543}
                                _ = fi53
                                _ = v54
                            }
                        case -1:
                            var v55 int = (fi51 >> ((42) & 63))
                            _ = v55
                        case 4294967295:
                            v29 = "foo bar"
                        }
                    default:
                        if (3.14159 < (0.1 + (((0.0 + 1.5) - ((float32((3.14159 - 0.0)) - (-2.25 + 1.5)) + (1e10 - 0.1))) + 1.5))) {
                            var v56 uint32 = 706
                            _ = v56
                        }
                    }
                case "":
                    v4[864] = fi51
                default:
                    v57 := []int{128}
                    _ = v57
                }
            }
            var v58 uint = 1
            _ = fi51
            _ = v52
            _ = v58
        }
        _ = v49
        _ = v50
    } else {
        v1 = v1
        if false {
            if false {
                var v59 int16 = -884
                _ = v59
            } else {
                for fi60 := 0; fi60 < 16; fi60++ {
                    var v61 uint = 792
                    _ = fi60
                    _ = v61
                }
                var v62 uint64 = 2
                v63 := GenS8{F0: 65535, F1: -851, F2: 1}
                _ = v62
                _ = v63
            }
        } else {
            if (-2147483648 < 1) {
                v3 = v3
                var v64 float32 = -2.25
                _ = v64
            } else {
                for fi65 := 0; fi65 < 8; fi65++ {
                    if v1 {
                        var v66 float32 = 3.14159
                        _ = v66
                    } else {
                        for fi67 := 0; fi67 < 11; fi67++ {
                            var v68 int64 = (v43 % ((-1) | 1))
                            _ = fi67
                            _ = v68
                        }
                        v69 := GenS3{F0: 1000, F1: 42}
                        for fi70 := 0; fi70 < 10; fi70++ {
                            if v1 {
                                var v71 uint32 = 283
                                _ = v71
                            }
                            _ = fi70
                        }
                        _ = v69
                    }
                    _ = fi65
                }
                v1 = v1
                if true {
                    var v72 uint16 = 0
                    _ = len(v4)
                    v73 := GenS9{F0: 1, F1: 316}
                    _ = v72
                    _ = v73
                } else {
                    v74 := GenS3{F0: 0, F1: 1}
                    if true {
                        var v75 int = 4294967295
                        _ = v75
                    } else {
                        fmt.Println(v1, v29, v29)
                        switch ("hello" + v3) {
                        case "\n\t":
                            var v76 uint8 = (v2 * 9)
                            _ = v76
                        }
                    }
                    if false {
                        var v77 int = int(v2)
                        fmt.Println(v43, v77, v3)
                        _ = v77
                    }
                    _ = v74
                }
            }
        }
    }
    if v1 {
        var v78 float64 = 1e10
        _ = v78
    } else {
        for fi79 := 0; fi79 < 16; fi79++ {
            var v80 int64 = (v43 - v43)
            var v81 float32 = (float32((-2.25 + -2.25)) + 3.14159)
            for fi82 := 0; fi82 < 13; fi82++ {
                if (!(true || false)) {
                    var v83 int16 = 42
                    if v1 {
                        switch fi79 {
                        case -9223372036854775808:
                            if v1 {
                                if v1 {
                                    v2 -= v2
                                } else {
                                    if (v1 || v1) {
                                        v83 += (v83 << ((255) & 63))
                                        v84 := []int{32767}
                                        _ = v84
                                    } else {
                                        _ = len(v4)
                                    }
                                }
                                v85 := GenS10{F0: -991}
                                v86 := GenS11{F0: fi82, F1: 0}
                                _ = v85
                                _ = v86
                            }
                        case 183:
                            switch fi82 {
                            case 9223372036854775807:
                                switch ("hello" + "") {
                                case "":
                                    var v87 bool = ("hello" < ("\n\t" + v29))
                                    _ = v87
                                default:
                                    var v88 string = "hello"
                                    _ = v88
                                }
                            case 32767:
                                if v1 {
                                    v89 := GenS3{F0: v83, F1: genFn0(-1, fi79)}
                                    _ = v89
                                }
                            }
                        case -1:
                            v80 = (^(v43))
                        }
                        var v90 uint8 = (v2 ^ 0)
                        _ = v90
                    }
                    _ = len(v4)
                    _ = v83
                }
                v91 := GenS12{F0: 4294967295, F1: 0}
                v1 = v1
                _ = fi82
                _ = v91
            }
            _ = fi79
            _ = v80
            _ = v81
        }
        v92 := GenS3{F0: 14, F1: 65535}
        _ = v92
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    _ = v1
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v43
    _ = v48
}
