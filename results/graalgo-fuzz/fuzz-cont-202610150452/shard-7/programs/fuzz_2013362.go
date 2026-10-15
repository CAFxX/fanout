package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS5 struct { F0 bool; F1 int }
type GenS6 struct { F0 int16; F1 uint }
type GenS7 struct { F0 string; F1 uint }
type GenS8 struct { F0 uint; F1 int64; F2 int16 }
type GenS9 struct { F0 int16; F1 uint16 }
func genFn0(p0 uint64, p1 string, p2 int) uint64 {
    var v4 int32 = 1
    var v5 int64 = -382
    v6 := []int{(p2 + -135), p2}
    var v7 uint16 = 42
    v7 = (v7 >> ((uint(p0)) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 uint) uint16 {
    fmt.Println(p0)
    switch "" {
    case "":
        p0 = p0
    case "\n\t":
        for fi8 := 0; fi8 < 9; fi8++ {
            v9 := []int{fi8, (fi8 ^ 9223372036854775807), fi8, (fi8 | (fi8 | fi8))}
            for fi10 := 0; fi10 < 11; fi10++ {
                if false {
                    switch "a" {
                    case "foo bar":
                        p0 = p0
                    case "\n\t":
                        v11 := GenS1{F0: "hello", F1: 70}
                        _ = v11
                    }
                    v12 := GenS1{F0: ("\n\t" + (((("foo bar" + (("" + "a") + ("" + "hello"))) + ("a" + "hello")) + ("hello" + "a")) + ("" + "a"))), F1: 255}
                    v13 := GenS1{F0: "hello", F1: 255}
                    var v14 int16 = -32768
                    _ = v12
                    _ = v13
                    _ = v14
                } else {
                    p0 = p0
                    v15 := GenS2{F0: "hello"}
                    _ = v15
                }
                if len(v9) > 0 {
                    v9[0] = (fi10 << ((952) & 63))
                }
                _ = fi10
            }
            _ = fi8
            _ = v9
        }
    case "foo bar":
        if (!(-128 == 0)) {
            for fi16 := 0; fi16 < 20; fi16++ {
                if false {
                    v17 := GenS3{F0: 770, F1: 0}
                    v18 := []int{fi16, (fi16 & (fi16 - fi16)), -144}
                    if true {
                        v19 := GenS2{F0: "foo bar"}
                        for fi20 := 0; fi20 < 4; fi20++ {
                            var v21 uint8 = 128
                            p0 = (p0 - (p0 << ((2) & 63)))
                            _ = fi20
                            _ = v21
                        }
                        _ = v19
                    }
                    _ = v17
                    _ = v18
                }
                if false {
                    v22 := GenS4{F0: genFn0(1, ("hello" + ""), fi16), F1: 0, F2: 32767}
                    _ = v22
                }
                for fi23 := 0; fi23 < 17; fi23++ {
                    if false {
                        v24 := GenS5{F0: false, F1: fi23}
                        var v25 int32 = 2147483647
                        _ = v24
                        _ = v25
                    }
                    _ = fi23
                }
                _ = fi16
            }
        }
    default:
        v26 := GenS6{F0: 1, F1: (p0 + p0)}
        _ = v26
    }
    for fi27 := 0; fi27 < 20; fi27++ {
        p0 += p0
        p0 = 460
        _ = fi27
    }
    v28 := GenS3{F0: 853, F1: 151}
    for fi29 := 0; fi29 < 17; fi29++ {
        if (32767 >= 68) {
            v30 := []int{1, 483, (fi29 + (fi29 ^ 662))}
            _ = len(v30)
            _ = v30
        } else {
            v31 := []int{fi29}
            for fi32 := 0; fi32 < 4; fi32++ {
                _ = len(v31)
                if len(v31) > 0 {
                    v31[0] = (^(fi32))
                }
                if true {
                    var v33 uint32 = 277
                    var v34 uint16 = 485
                    _ = v33
                    _ = v34
                }
                _ = fi32
            }
            _ = len(v31)
            _ = v31
        }
        _ = fi29
    }
    _ = p0
    _ = v28
    return 32767
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []string{(("" + "a") + ((("" + "") + (("hello" + (("" + "a") + "\n\t")) + "\n\t")) + ("hello" + "foo bar")))}
    var v3 int32 = -1
    v3 = (v3 / ((v3) | 1))
    v3 -= (v3 & (v3 / ((344) | 1)))
    if true {
        var v35 uint32 = 4294967295
        if (901 <= 18446744073709551615) {
            v36 := GenS2{F0: "foo bar"}
            var v37 bool = false
            _ = v36
            _ = v37
        }
        if false {
            v38 := GenS3{F0: genFn1(963), F1: 566}
            delete(v1, v35)
            _ = v38
        }
        _ = v35
    }
    switch 639 {
    case -2147483648:
        for fi39 := 0; fi39 < 5; fi39++ {
            v3 = v3
            _ = fi39
        }
    }
    _ = len(v1)
    v3 = 0
    if ("" < "foo bar") {
        fmt.Println(v3)
        v40 := GenS7{F0: "", F1: 42}
        v41 := GenS5{F0: false, F1: -2147483648}
        _ = v40
        _ = v41
    }
    var v42 float32 = (1.5 / (float32(float64((3.14159 * 0.1))) + (1e10 - 1.5)))
    if (v42 > v42) {
        v42 = -2.25
        var v43 uint = 65535
        v2 = append(v2, ("a" + "foo bar"))
        _ = v43
    }
    var v44 float64 = -2.25
    v45 := GenS6{F0: -32768, F1: 65535}
    if false {
        for fi46 := 0; fi46 < 6; fi46++ {
            var v47 int = fi46
            v44 = v44
            _ = fi46
            _ = v47
        }
    } else {
        v42 -= (((3.14159 * (v42 + float32(((1e10 - v44) - (v44 * 1e10))))) * v42) + (1.5 * v42))
    }
    if false {
        v44 += v44
        if true {
            var v48 string = "hello"
            _ = v48
        }
        for fi49 := 0; fi49 < 16; fi49++ {
            v42 -= (((v42 - (v42 / v42)) * 1e10) * 3.14159)
            _ = fi49
        }
    } else {
        switch "\n\t" {
        case "":
            fmt.Println(v42)
        }
        if false {
            v50 := v1[0]
            if false {
                switch "\n\t" {
                case "hello":
                    for fi51 := 0; fi51 < 2; fi51++ {
                        v52 := GenS5{F0: (v42 != v42), F1: fi51}
                        for fi53 := 0; fi53 < 16; fi53++ {
                            var v54 uint16 = 65535
                            v44 = v44
                            if true {
                                var v55 float64 = 0.0
                                if (v54 < genFn1(952)) {
                                    delete(v1, v50)
                                } else {
                                    fmt.Println(v55, v50, v3)
                                    var v56 uint16 = (v54 * (v54 + 32767))
                                    v42 = float32((((float64(((1.5 - float32(v55)) * ((-2.25 - 1.5) - (float32(1e10) + float32(-2.25))))) * 1e10) + (v44 / v44)) - (v44 / float64((3.14159 - ((3.14159 / v42) - (1e10 / float32((v55 / v55)))))))))
                                    _ = v56
                                }
                                switch (("hello" + ("a" + "a")) + "a") {
                                case "hello":
                                    switch fi53 {
                                    case 128:
                                        fmt.Println(fi53, v55, v42)
                                    case 127:
                                        if true {
                                            var v57 uint8 = 69
                                            switch (^(fi53)) {
                                            case 42:
                                                v54 = v54
                                            }
                                            if ((float64((3.14159 + float32(v44))) - (v55 / (v44 * v55))) != (3.14159 - (((v55 + v44) + v44) / v44))) {
                                                v58 := GenS2{F0: ""}
                                                _ = v58
                                            }
                                            _ = v57
                                        } else {
                                            var v59 uint8 = 57
                                            if true {
                                                _ = len(v2)
                                                fmt.Println(v59, fi53, v42)
                                                v44 = v44
                                                var v60 uint = 224
                                                fmt.Println(v3)
                                                fmt.Println(v50, v54, v3)
                                                var v61 uint64 = 42
                                                if true {
                                                    switch fi51 {
                                                    case 2147483647:
                                                        var v62 uint64 = (v61 ^ (v61 % ((v61) | 1)))
                                                        _ = v62
                                                    }
                                                    v55 = (v44 * v44)
                                                    _ = len(v2)
                                                }
                                                _ = v60
                                                _ = v61
                                            }
                                            var v63 bool = (42 <= v3)
                                            var v64 uint64 = 42
                                            _ = v59
                                            _ = v63
                                            _ = v64
                                        }
                                    case -9223372036854775808:
                                        var v65 int8 = 127
                                        _ = v65
                                    }
                                case "":
                                    var v66 int64 = 2
                                    _ = v66
                                }
                                _ = v55
                            }
                            _ = fi53
                            _ = v54
                        }
                        var v67 int16 = 32767
                        _ = fi51
                        _ = v52
                        _ = v67
                    }
                case "\n\t":
                    fmt.Println(v44, v3, v44)
                }
            } else {
                v68 := GenS8{F0: 2, F1: 237, F2: 32767}
                v42 = (1.5 / -2.25)
                v1[(v50 - v50)] = (v50 & (v50 & (v50 & v50)))
                _ = v68
            }
            _ = v50
        }
    }
    switch 42 {
    case 32768:
        v42 += v42
    case -1:
        var v69 int = -710
        _ = v69
    case 32767:
        for fi70 := 0; fi70 < 4; fi70++ {
            v71 := GenS2{F0: "a"}
            var v72 string = ((("foo bar" + "foo bar") + "\n\t") + "hello")
            _ = fi70
            _ = v71
            _ = v72
        }
    default:
        fmt.Println(v44)
    }
    var v73 float32 = v42
    if (((("foo bar" + ("a" + "a")) + "\n\t") + "\n\t") != "a") {
        var v74 int32 = (-(v3))
        _ = v74
    } else {
        fmt.Println(v44)
        v42 += 0.0
    }
    if true {
        fmt.Println(v42, v3)
        for fi75 := 0; fi75 < 2; fi75++ {
            _ = len(v2)
            v76 := v1[2147483648]
            if (4294967295 != 2) {
                switch "foo bar" {
                case "hello":
                    if (!(true && true)) {
                        if len(v2) > 0 {
                            v2[0] = "hello"
                        }
                        var v77 float64 = 1e10
                        var v78 int16 = -612
                        _ = v77
                        _ = v78
                    }
                case "foo bar":
                    if len(v2) > 0 {
                        v2[0] = ("hello" + "\n\t")
                    }
                }
                var v79 float64 = 0.0
                _ = v79
            } else {
                v80 := GenS5{F0: false, F1: (fi75 + (fi75 - 302))}
                for fi81 := 0; fi81 < 2; fi81++ {
                    if true {
                        fmt.Println(v3, v73, fi75)
                        var v82 float32 = v73
                        _ = v82
                    }
                    if false {
                        v83 := GenS4{F0: 599, F1: 0, F2: 2}
                        _ = v83
                    } else {
                        switch (^(fi75)) {
                        case 94:
                            v84 := GenS9{F0: -677, F1: 32767}
                            _ = v84
                        case 371:
                            v76 += 423
                        case -2147483648:
                            var v85 uint16 = 2
                            _ = v85
                        default:
                            v42 = float32(1e10)
                        }
                        var v86 int8 = 42
                        _ = v86
                    }
                    var v87 int64 = -56
                    v42 -= (-2.25 + v42)
                    var v88 uint = 60
                    _ = fi81
                    _ = v87
                    _ = v88
                }
                _ = v80
            }
            _ = fi75
            _ = v76
        }
        for fi89 := 0; fi89 < 3; fi89++ {
            v42 = 1.5
            for fi90 := 0; fi90 < 9; fi90++ {
                v44 = v44
                fmt.Println(v3, fi90)
                v44 += 0.0
                _ = fi90
            }
            if false {
                v3 = -1
                fmt.Println(v44, v3)
            } else {
                var v91 uint = 2
                if (!true) {
                    if (float64(float32((v44 - (float64(v73) * v44)))) == 1e10) {
                        v92 := GenS2{F0: ""}
                        for fi93 := 0; fi93 < 2; fi93++ {
                            v44 -= (v44 + (1.5 / -2.25))
                            v94 := GenS2{F0: "a"}
                            var v95 uint32 = 1
                            _ = fi93
                            _ = v94
                            _ = v95
                        }
                        _ = v92
                    }
                }
                _ = v91
            }
            _ = fi89
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v44
    _ = v45
    _ = v73
}
