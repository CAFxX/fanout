package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 uint8; F1 uint32; F2 uint16 }
type GenS3 struct { F0 int64; F1 uint; F2 bool }
type GenS4 struct { F0 uint16; F1 uint8; F2 uint }
type GenS5 struct { F0 int32; F1 int32; F2 string }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int64; F1 uint64; F2 int32 }
type GenS8 struct { F0 int32; F1 uint32 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int; F1 int16 }
type GenS12 struct { F0 int; F1 string }
type GenS13 struct { F0 int16; F1 bool; F2 uint32 }
func genFn0(p0 int) int32 {
    v5 := []int{}
    v6 := GenS1{F0: ("\n\t" + "\n\t"), F1: 128}
    _ = p0
    _ = v5
    _ = v6
    return 1
}
func genFn1() uint16 {
    switch "" {
    case "foo bar":
        var v7 uint32 = 161
        _ = v7
    case "hello":
        if ((65535 >= 1) && (63 > 255)) {
            if (50 != 1) {
                v8 := []int{128, 9223372036854775807, -9223372036854775808, 0}
                for fi9 := 0; fi9 < 11; fi9++ {
                    v10 := GenS2{F0: 1, F1: 2147483648, F2: 65535}
                    _ = fi9
                    _ = v10
                }
                for fi11 := 0; fi11 < 7; fi11++ {
                    for fi12 := 0; fi12 < 18; fi12++ {
                        _ = len(v8)
                        var v13 float64 = ((1e10 + float64((1e10 / -2.25))) * 3.14159)
                        _ = fi12
                        _ = v13
                    }
                    if true {
                        switch "foo bar" {
                        case "foo bar":
                            v14 := GenS3{F0: -833, F1: 1, F2: (!(float64(-2.25) <= 0.0))}
                            _ = v14
                        case "":
                            for fi15 := 0; fi15 < 18; fi15++ {
                                fmt.Println(fi11, fi15)
                                fmt.Println(fi15, fi11)
                                if (2147483648 > 4294967295) {
                                    if len(v8) > 0 {
                                        v8[0] = (fi15 / (((fi11 % (((fi15 + (-(fi11)))) | 1))) | 1))
                                    }
                                    switch -343 {
                                    case -2147483648:
                                        var v16 uint64 = 2
                                        _ = v16
                                    case 32767:
                                        v8 = append(v8, (fi11 + 32767))
                                    default:
                                        var v17 float64 = float64(0.0)
                                        _ = v17
                                    }
                                    var v18 uint64 = 612
                                    _ = v18
                                }
                                _ = fi15
                            }
                        }
                        var v19 bool = true
                        _ = v19
                    }
                    switch 4294967295 {
                    case 42:
                        var v20 bool = ((42 < 9223372036854775807) && false)
                        _ = v20
                    case 128:
                        var v21 uint16 = 42
                        _ = v21
                    case 32768:
                        var v22 int = (fi11 % (((^(fi11))) | 1))
                        _ = v22
                    }
                    _ = fi11
                }
                _ = v8
            }
            var v23 float32 = 0.1
            _ = v23
        } else {
            v24 := GenS4{F0: 42, F1: 127, F2: 255}
            _ = v24
        }
    case "\n\t":
        for fi25 := 0; fi25 < 9; fi25++ {
            var v26 uint32 = 336
            var v27 uint32 = 1
            v27 -= (^(v26))
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    var v28 string = (("\n\t" + ((("a" + "\n\t") + "") + "\n\t")) + ("\n\t" + ((("\n\t" + "hello") + "hello") + (("foo bar" + "") + "a"))))
    switch 127 {
    case 2147483647:
        var v29 float32 = 1e10
        _ = v29
    default:
        fmt.Println(v28)
    }
    v30 := GenS5{F0: -2147483648, F1: 974, F2: ("hello" + v28)}
    if (100 <= -39) {
        if false {
            fmt.Println(v28)
            v31 := GenS6{F0: -2147483648}
            for fi32 := 0; fi32 < 9; fi32++ {
                switch ((v28 + v28) + v28) {
                case "a":
                    if true {
                        v33 := GenS1{F0: (v28 + ""), F1: 2}
                        var v34 float32 = (((0.1 + -2.25) * 1e10) / 1e10)
                        _ = v33
                        _ = v34
                    } else {
                        for fi35 := 0; fi35 < 2; fi35++ {
                            fmt.Println(fi32, fi35, fi35)
                            var v36 int = fi35
                            _ = fi35
                            _ = v36
                        }
                        if true {
                            v37 := []int{fi32, (fi32 << ((18446744073709551615) & 63))}
                            if ((((!false) && (!false)) || false) || false) {
                                var v38 int16 = 1000
                                switch ((v28 + v28) + ("a" + (v28 + "hello"))) {
                                case "a":
                                    v39 := GenS7{F0: 4, F1: 1, F2: 0}
                                    _ = v39
                                case "hello":
                                    if len(v37) > 0 {
                                        v37[0] = (fi32 | fi32)
                                    }
                                default:
                                    if (20 < 2) {
                                        fmt.Println(v38, v38)
                                        if false {
                                            var v40 uint = 2
                                            for fi41 := 0; fi41 < 12; fi41++ {
                                                var v42 int = fi41
                                                v28 = ((v28 + ((v28 + v28) + v28)) + (("" + ("" + (v28 + v28))) + (("hello" + (v28 + v28)) + "hello")))
                                                fmt.Println(fi41)
                                                _ = fi41
                                                _ = v42
                                            }
                                            _ = v40
                                        } else {
                                            v43 := GenS4{F0: 42, F1: 190, F2: 4294967295}
                                            _ = v43
                                        }
                                        var v44 uint = 0
                                        _ = v44
                                    } else {
                                        switch (("hello" + v28) + v28) {
                                        case "a":
                                            var v45 int8 = -1
                                            _ = v45
                                        default:
                                            fmt.Println(fi32)
                                        }
                                        if (478 == 2147483647) {
                                            for fi46 := 0; fi46 < 5; fi46++ {
                                                if false {
                                                    fmt.Println(v38, fi32)
                                                    v37 = append(v37, 128)
                                                } else {
                                                    var v47 uint = 159
                                                    v37 = append(v37, fi46)
                                                    var v48 float32 = 0.0
                                                    var v49 uint64 = 42
                                                    v50 := GenS8{F0: genFn0(int(v38)), F1: 65535}
                                                    _ = v47
                                                    _ = v48
                                                    _ = v49
                                                    _ = v50
                                                }
                                                var v51 string = ""
                                                _ = fi46
                                                _ = v51
                                            }
                                            if false {
                                                v28 = ((v28 + (v28 + v28)) + v28)
                                                v52 := GenS4{F0: 42, F1: 1, F2: 255}
                                                _ = v52
                                            }
                                        }
                                        v53 := GenS9{F0: (fi32 - fi32)}
                                        _ = v53
                                    }
                                }
                                _ = v38
                            }
                            _ = v37
                        }
                    }
                case "\n\t":
                    if true {
                        v28 = "\n\t"
                    } else {
                        if (!(!(9223372036854775807 != 0))) {
                            v54 := []int{}
                            if len(v54) > 0 {
                                v54[0] = (fi32 / ((fi32) | 1))
                            }
                            for fi55 := 0; fi55 < 19; fi55++ {
                                if len(v54) >= 2 {
                                    v54 = v54[1:len(v54)]
                                }
                                if (false || false) {
                                    switch fi32 {
                                    case 32767:
                                        var v56 uint16 = 998
                                        _ = v56
                                    }
                                    if (false && false) {
                                        fmt.Println(v28)
                                        v54 = append(v54, -267)
                                        if (0.1 < 1e10) {
                                            v28 = ((v28 + "\n\t") + (v28 + (v28 + ((v28 + "a") + (v28 + "foo bar")))))
                                            switch 127 {
                                            case 1:
                                                var v57 float64 = 0.1
                                                _ = v57
                                            case 4294967295:
                                                var v58 bool = false
                                                _ = v58
                                            case 32768:
                                                if true {
                                                    var v59 uint32 = 731
                                                    v28 = (v28 + (v28 + "foo bar"))
                                                    _ = v59
                                                }
                                            }
                                        }
                                    } else {
                                        fmt.Println(v28)
                                        var v60 uint64 = 4294967295
                                        var v61 int8 = 127
                                        v28 = ""
                                        _ = v60
                                        _ = v61
                                    }
                                    switch 128 {
                                    case 2147483647:
                                        v62 := GenS2{F0: 202, F1: 1, F2: 0}
                                        _ = v62
                                    case 4294967295:
                                        v28 = v28
                                    case 127:
                                        var v63 int16 = -1
                                        v28 = "\n\t"
                                        var v64 uint32 = 0
                                        _ = v63
                                        _ = v64
                                    default:
                                        v28 = v28
                                    }
                                }
                                fmt.Println(fi55)
                                _ = fi55
                            }
                            _ = v54
                        } else {
                            switch (v28 + "") {
                            case "\n\t":
                                v28 = v28
                            case "a":
                                var v65 float32 = -2.25
                                _ = v65
                            }
                            for fi66 := 0; fi66 < 8; fi66++ {
                                v67 := GenS10{F0: 309}
                                switch (fi32 * fi32) {
                                case 42:
                                    var v68 int32 = 42
                                    _ = v68
                                case -9223372036854775808:
                                    v69 := GenS11{F0: fi32, F1: -642}
                                    _ = v69
                                case 303:
                                    fmt.Println(fi66, fi66)
                                    v28 = v28
                                    var v70 float32 = 1.5
                                    switch "hello" {
                                    case "\n\t":
                                        var v71 int16 = 264
                                        _ = v71
                                    case "hello":
                                        var v72 int16 = 1
                                        v70 = v70
                                        v73 := []int{}
                                        var v74 int32 = -1
                                        _ = v72
                                        _ = v73
                                        _ = v74
                                    }
                                    _ = v70
                                }
                                _ = fi66
                                _ = v67
                            }
                            if ((0.0 - float64((1.5 - 0.0))) >= (1e10 - (0.0 * (3.14159 + (float64(float32(((3.14159 * 1e10) / 1e10))) * 1e10))))) {
                                for fi75 := 0; fi75 < 20; fi75++ {
                                    v76 := []int{fi32}
                                    _ = fi75
                                    _ = v76
                                }
                            } else {
                                for fi77 := 0; fi77 < 18; fi77++ {
                                    v78 := []int{0}
                                    var v79 int32 = 2147483647
                                    var v80 int64 = 0
                                    var v81 uint32 = 0
                                    if true {
                                        v81 -= v81
                                        var v82 uint32 = 42
                                        v83 := GenS12{F0: -1, F1: v28}
                                        _ = v82
                                        _ = v83
                                    }
                                    var v84 uint32 = v81
                                    _ = fi77
                                    _ = v78
                                    _ = v79
                                    _ = v80
                                    _ = v81
                                    _ = v84
                                }
                                for fi85 := 0; fi85 < 17; fi85++ {
                                    v86 := GenS13{F0: -32768, F1: false, F2: 136}
                                    v28 = v28
                                    v28 = v28
                                    var v87 bool = false
                                    var v88 bool = (v87 && v87)
                                    fmt.Println(fi85)
                                    _ = fi85
                                    _ = v86
                                    _ = v87
                                    _ = v88
                                }
                            }
                        }
                    }
                }
                _ = fi32
            }
            _ = v31
        } else {
            var v89 uint8 = 30
            _ = v89
        }
        var v90 int8 = 42
        v91 := []int{-1, -777, 32767, -1}
        _ = v90
        _ = v91
    }
    _ = v28
    _ = v30
    return 32767
}

func main() {
    var v1 int8 = 42
    var v2 uint16 = 0
    var v3 uint32 = 4294967295
    var v4 string = "\n\t"
    fmt.Println(v3, v2, v4)
    if true {
        var v92 string = v4
        _ = v92
    } else {
        var v93 uint8 = 127
        v94 := GenS5{F0: 42, F1: 42, F2: v4}
        if false {
            var v95 int8 = (v1 >> ((330) & 63))
            v96 := GenS7{F0: -104, F1: 408, F2: 448}
            _ = v95
            _ = v96
        }
        _ = v93
        _ = v94
    }
    v97 := []int{}
    v3 = 1
    for fi98 := 0; fi98 < 13; fi98++ {
        var v99 int = 128
        _ = fi98
        _ = v99
    }
    v3 -= (^(v3))
    fmt.Println(v2, v4, v4)
    var v100 uint32 = uint32(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v97:", len(v97), v97)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v4
    _ = v97
}
