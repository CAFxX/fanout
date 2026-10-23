package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int; F1 uint64; F2 string }
type GenS7 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS8 struct { F0 int32; F1 int }
type GenS9 struct { F0 int; F1 uint8; F2 uint8 }
type GenS10 struct { F0 uint8; F1 uint; F2 string }
type GenS11 struct { F0 string }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 uint64; F1 uint8 }
type GenS15 struct { F0 int8; F1 uint8 }
type GenS16 struct { F0 uint32 }
type GenS17 struct { F0 int64; F1 bool }
type GenS18 struct { F0 uint8 }
func genFn0(p0 int16, p1 bool, p2 int16) uint8 {
    if (!(!p1)) {
        if p1 {
            p2 = p0
        } else {
            v5 := GenS1{F0: 2}
            var v6 int8 = int8(p2)
            _ = v5
            _ = v6
        }
    }
    switch (("foo bar" + "a") + ((("a" + "a") + "hello") + "a")) {
    case "\n\t":
        fmt.Println(p2)
    case "foo bar":
        v7 := GenS1{F0: 128}
        _ = v7
    case "a":
        if p1 {
            for fi8 := 0; fi8 < 20; fi8++ {
                for fi9 := 0; fi9 < 7; fi9++ {
                    if false {
                        if (true || (("\n\t" + "\n\t") < "hello")) {
                            v10 := []int{1, 42}
                            _ = v10
                        }
                        fmt.Println(fi9, p2)
                        p1 = (245 >= 255)
                        p2 += (p0 >> ((18446744073709551615) & 63))
                        if p1 {
                            fmt.Println(p1, fi9)
                            var v11 float32 = -2.25
                            if p1 {
                                var v12 float32 = float32(1.5)
                                if p1 {
                                    v11 -= v12
                                    var v13 uint32 = 1
                                    fmt.Println(p1)
                                    if p1 {
                                        v14 := GenS2{F0: 1, F1: v13}
                                        fmt.Println(p2)
                                        p1 = p1
                                        if true {
                                            fmt.Println(v11, fi8, v11)
                                        }
                                        _ = v14
                                    }
                                    _ = v13
                                } else {
                                    var v15 int8 = -128
                                    _ = v15
                                }
                                _ = v12
                            } else {
                                fmt.Println(p2, p2)
                                v16 := GenS2{F0: 2147483647, F1: 1}
                                var v17 float64 = 3.14159
                                if false {
                                    var v18 int32 = -2147483648
                                    v17 = 1e10
                                    _ = v18
                                }
                                _ = v16
                                _ = v17
                            }
                            _ = v11
                        }
                    } else {
                        v19 := GenS3{F0: "a", F1: 1}
                        v20 := []int{(^(fi8)), fi9, fi8}
                        _ = v19
                        _ = v20
                    }
                    p0 = p2
                    _ = fi9
                }
                fmt.Println(fi8, p0)
                _ = fi8
            }
            p0 += (p0 >> ((255) & 63))
        }
    }
    v21 := []int{268, 128, 4294967295, 1}
    v22 := GenS4{F0: 42, F1: 18446744073709551615}
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    _ = v22
    return 128
}
func genFn1(p0 uint32, p1 int, p2 uint32) int {
    v23 := GenS5{F0: 155}
    v24 := GenS6{F0: p1, F1: 4294967295, F2: ("\n\t" + "\n\t")}
    if (-117 == 127) {
        fmt.Println(p2)
        var v25 uint32 = 555
        v26 := GenS7{F0: 0, F1: 1, F2: 0}
        _ = v25
        _ = v26
    }
    p2 = (^(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    _ = v24
    return 127
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[int]int)
    var v3 uint16 = 32767
    var v4 uint16 = (^(v3))
    for fi27 := 0; fi27 < 13; fi27++ {
        switch ("" + "\n\t") {
        case "":
            _ = len(v2)
        case "a":
            fmt.Println(v1, v3, v3)
        }
        _ = fi27
    }
    switch genFn1(42, genFn1(42, 0, 2147483647), 65535) {
    case 4294967295:
        v28 := v2[2147483647]
        _ = v28
    }
    if ((v3 % (((v3 * v4)) | 1)) >= v4) {
        switch "hello" {
        case "\n\t":
            if true {
                fmt.Println(v1, v1, v3)
                v1 -= (v1 << ((624) & 63))
            }
        case "foo bar":
            v29 := GenS5{F0: 1000}
            _ = v29
        default:
            switch -1 {
            case 577:
                if false {
                    for fi30 := 0; fi30 < 3; fi30++ {
                        v31 := GenS8{F0: -92, F1: 465}
                        _ = fi30
                        _ = v31
                    }
                    var v32 uint = 0
                    fmt.Println(v1, v4)
                    _ = v32
                } else {
                    var v33 int = 42
                    _ = v33
                }
            case 127:
                if ((3.14159 == 0.1) && false) {
                    v4 -= (^(v4))
                }
            default:
                var v34 float32 = 3.14159
                _ = v34
            }
        }
    } else {
        fmt.Println(v1)
        var v35 bool = (true || false)
        if v35 {
            var v36 uint16 = (v4 + (^(v3)))
            if (!v35) {
                v37 := GenS9{F0: 128, F1: 255, F2: genFn0(467, false, 0)}
                if (!((-2.25 + -2.25) > 1e10)) {
                    v38 := GenS4{F0: -880, F1: 417}
                    for fi39 := 0; fi39 < 15; fi39++ {
                        var v40 int8 = 82
                        for fi41 := 0; fi41 < 5; fi41++ {
                            var v42 uint8 = (v1 << ((615) & 63))
                            _ = fi41
                            _ = v42
                        }
                        _ = fi39
                        _ = v40
                    }
                    v43 := GenS10{F0: (^(v1)), F1: 65535, F2: "hello"}
                    _ = v38
                    _ = v43
                }
                _ = v37
            } else {
                v44 := GenS11{F0: (("" + ((("hello" + "a") + ("hello" + (((("a" + "") + "\n\t") + ("" + "\n\t")) + "foo bar"))) + "foo bar")) + ("hello" + "a"))}
                v45 := []int{0}
                _ = v44
                _ = v45
            }
            _ = v36
        }
        _ = v35
    }
    var v46 uint8 = (v1 + v1)
    v47 := GenS12{F0: v3}
    var v48 uint = 1
    for fi49 := 0; fi49 < 13; fi49++ {
        v50 := []int{}
        _ = fi49
        _ = v50
    }
    v51 := []int{2147483647, -679, -752, -116}
    var v52 uint16 = 0
    for fi53 := 0; fi53 < 6; fi53++ {
        var v54 uint8 = (v46 | genFn0(-510, true, -1))
        if false {
            v52 += v52
            switch 956 {
            case 32768:
                var v55 uint8 = 1
                _ = v55
            }
            for fi56 := 0; fi56 < 17; fi56++ {
                var v57 int = 32768
                fmt.Println(fi56, v54)
                switch "foo bar" {
                case "a":
                    var v58 bool = (((true && true) || (false || (false && (675 < 0)))) || false)
                    _ = v58
                case "":
                    var v59 int64 = -9223372036854775808
                    _ = v59
                case "\n\t":
                    if false {
                        fmt.Println(v46, v46, v1)
                    } else {
                        var v60 float32 = 0.0
                        _ = v60
                    }
                default:
                    if len(v51) > 0 {
                        v51[0] = -1
                    }
                    if ((3.14159 - (0.0 - 0.0)) != (float32((1.5 - ((1e10 - -2.25) + float64(float32(float64(3.14159)))))) * -2.25)) {
                        _ = len(v2)
                    } else {
                        v61 := GenS2{F0: -517, F1: 1}
                        v62 := v2[(fi53 & (fi56 % ((int(v48)) | 1)))]
                        _ = v61
                        _ = v62
                    }
                    var v63 float64 = 1.5
                    _ = v63
                }
                fmt.Println(fi53, v48, v52)
                var v64 uint64 = uint64(v48)
                _ = fi56
                _ = v57
                _ = v64
            }
        } else {
            v65 := v2[fi53]
            _ = v65
        }
        if len(v51) > 0 {
            v51[0] = fi53
        }
        _ = fi53
        _ = v54
    }
    _ = len(v2)
    var v66 int16 = -32768
    for fi67 := 0; fi67 < 19; fi67++ {
        if true {
            if true {
                var v68 uint32 = 336
                v69 := GenS10{F0: v1, F1: 951, F2: "\n\t"}
                switch fi67 {
                case 32768:
                    switch "a" {
                    case "hello":
                        if true {
                            v48 = v48
                        }
                    default:
                        v2[fi67] = (fi67 << (((v48 - (v48 & v48))) & 63))
                    }
                case 32767:
                    if (2147483647 < -2147483648) {
                        var v70 uint = v48
                        v71 := GenS13{F0: (v68 % (((^(v68))) | 1))}
                        _ = v70
                        _ = v71
                    }
                case -9223372036854775808:
                    for fi72 := 0; fi72 < 18; fi72++ {
                        v52 = v3
                        _ = fi72
                    }
                default:
                    if (true || false) {
                        v73 := v2[fi67]
                        v74 := GenS14{F0: 2, F1: 42}
                        for fi75 := 0; fi75 < 12; fi75++ {
                            var v76 int64 = -331
                            if (!false) {
                                _ = len(v2)
                                if (v76 > (v76 << ((255) & 63))) {
                                    v77 := GenS9{F0: v73, F1: (v1 * 138), F2: v46}
                                    v3 = v3
                                    _ = v77
                                }
                                if true {
                                    v78 := GenS15{F0: -11, F1: v1}
                                    _ = len(v51)
                                    _ = v78
                                }
                            } else {
                                if ((!(int16(v48) != 0)) && false) {
                                    v79 := GenS16{F0: uint32(v48)}
                                    fmt.Println(v3, v76, v4)
                                    _ = v79
                                }
                            }
                            if false {
                                v66 -= v66
                                var v80 uint = v48
                                v80 = (^(v80))
                                _ = v80
                            } else {
                                var v81 float32 = 1.5
                                if true {
                                    v82 := GenS17{F0: (v76 ^ (v76 << ((uint(v4)) & 63))), F1: false}
                                    _ = v82
                                }
                                _ = v81
                            }
                            _ = fi75
                            _ = v76
                        }
                        _ = v73
                        _ = v74
                    } else {
                        v68 = uint32(v1)
                        if false {
                            v83 := GenS18{F0: genFn0(935, (!false), 1)}
                            _ = v83
                        }
                    }
                }
                _ = v68
                _ = v69
            } else {
                fmt.Println(v48, v66)
            }
            for fi84 := 0; fi84 < 6; fi84++ {
                fmt.Println(v52)
                _ = fi84
            }
        } else {
            for fi85 := 0; fi85 < 18; fi85++ {
                var v86 uint = uint(v1)
                _ = fi85
                _ = v86
            }
            for fi87 := 0; fi87 < 15; fi87++ {
                fmt.Println(v3, fi67, v52)
                _ = fi87
            }
            delete(v2, (fi67 << (((v48 | 42)) & 63)))
        }
        v88 := GenS4{F0: 682, F1: 0}
        fmt.Println(v66)
        _ = fi67
        _ = v88
    }
    v89 := GenS16{F0: 65535}
    fmt.Println(v66, v48, v48)
    if ((-973 == 1) && (42 >= -1)) {
        fmt.Println(v52, v3, v3)
        switch ("foo bar" + ("foo bar" + ("foo bar" + "\n\t"))) {
        case "\n\t":
            for fi90 := 0; fi90 < 7; fi90++ {
                delete(v2, -1)
                _ = fi90
            }
        case "hello":
            v4 = (v4 / ((v4) | 1))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("v52:", v52)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v51
    _ = v52
    _ = v66
    _ = v89
}
