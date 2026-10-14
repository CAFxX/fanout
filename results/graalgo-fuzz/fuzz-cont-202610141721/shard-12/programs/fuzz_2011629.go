package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 int8; F2 uint8 }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 uint32; F1 string; F2 int8 }
type GenS6 struct { F0 int32; F1 uint8 }
type GenS7 struct { F0 int8; F1 int32; F2 int32 }
type GenS8 struct { F0 uint64; F1 int16 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int64; F1 int32; F2 string }
type GenS11 struct { F0 int32; F1 string; F2 int }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 int32; F1 int8; F2 uint64 }
func genFn0(p0 uint32, p1 string) string {
    var v7 uint8 = 2
    for fi8 := 0; fi8 < 6; fi8++ {
        p1 = (p1 + ("hello" + (((p1 + p1) + p1) + "a")))
        if true {
            var v9 float32 = 0.1
            fmt.Println(v9)
            var v10 int16 = 42
            _ = v10
            _ = v9
        } else {
            for fi11 := 0; fi11 < 10; fi11++ {
                if false {
                    p1 = p1
                } else {
                    fmt.Println(p1, p0, fi11)
                }
                v12 := GenS1{F0: (true && (!false)), F1: 636, F2: 255}
                if true {
                    var v13 uint8 = v7
                    var v14 uint16 = 255
                    if ((true || false) && ((!((32767 >= 32767) || (!(!(fi11 == (fi11 << ((4294967295) & 63))))))) && false)) {
                        v15 := GenS1{F0: true, F1: 928, F2: v7}
                        _ = v15
                    } else {
                        p1 = p1
                    }
                    _ = v13
                    _ = v14
                } else {
                    if true {
                        switch 128 {
                        case -3:
                            if ((float64((0.1 + 1.5)) + 0.0) != (0.0 + float64(3.14159))) {
                                v16 := GenS1{F0: (!(false || (!false))), F1: 657, F2: (v7 * 1)}
                                if false {
                                    fmt.Println(fi8)
                                    var v17 uint16 = 255
                                    _ = v17
                                } else {
                                    p0 += p0
                                }
                                if false {
                                    v18 := GenS2{F0: int16(p0)}
                                    fmt.Println(p1)
                                    v19 := GenS1{F0: ((!false) && (255 <= 838)), F1: 1, F2: v7}
                                    v20 := GenS3{F0: 2147483647, F1: 100, F2: v7}
                                    v21 := []int{42, fi11, (fi11 - fi11)}
                                    _ = v18
                                    _ = v19
                                    _ = v20
                                    _ = v21
                                }
                                _ = v16
                            } else {
                                v7 += (v7 & v7)
                            }
                            fmt.Println(v7, fi8)
                            v22 := GenS1{F0: (true && false), F1: 821, F2: 53}
                            var v23 string = (p1 + p1)
                            _ = v22
                            _ = v23
                        case 809:
                            var v24 int64 = 640
                            v25 := []int{-2147483648}
                            _ = v24
                            _ = v25
                        case -9223372036854775808:
                            switch p1 {
                            case "\n\t":
                                if false {
                                    fmt.Println(p0, fi11)
                                    p1 = p1
                                } else {
                                    var v26 int16 = 1000
                                    switch (fi11 << ((815) & 63)) {
                                    case 2147483647:
                                        switch fi8 {
                                        case -94:
                                            if (!true) {
                                                if false {
                                                    fmt.Println(fi8)
                                                    v27 := GenS4{F0: -105, F1: (-(fi8))}
                                                    _ = v27
                                                }
                                            } else {
                                                v28 := GenS1{F0: (true && true), F1: 18446744073709551615, F2: 38}
                                                v29 := []int{(fi8 * 4294967295)}
                                                _ = v28
                                                _ = v29
                                            }
                                        }
                                    case 42:
                                        v26 = -257
                                    case -9223372036854775808:
                                        v30 := []int{(^(fi11)), 1, fi11, 32767}
                                        _ = v30
                                    default:
                                        v31 := GenS5{F0: p0, F1: p1, F2: -1}
                                        _ = v31
                                    }
                                    _ = v26
                                }
                            case "hello":
                                switch "" {
                                case "foo bar":
                                    v32 := GenS5{F0: (p0 >> ((42) & 63)), F1: "", F2: -17}
                                    _ = v32
                                case "hello":
                                    var v33 int = (fi11 >> ((132) & 63))
                                    v34 := []int{}
                                    if len(v34) > 0 {
                                        v34[0] = (fi8 + (fi8 >> ((65535) & 63)))
                                    }
                                    _ = v33
                                    _ = v34
                                }
                            case "":
                                fmt.Println(p1, p1, fi11)
                            }
                        }
                        switch p1 {
                        case "a":
                            p0 -= (^(p0))
                        }
                        if false {
                            v35 := GenS6{F0: -284, F1: (v7 >> ((4294967295) & 63))}
                            _ = v35
                        } else {
                            v36 := []int{fi11}
                            p0 -= p0
                            fmt.Println(fi8, p0)
                            _ = v36
                        }
                    } else {
                        switch (p1 + "\n\t") {
                        case "a":
                            var v37 int8 = 0
                            _ = v37
                        case "\n\t":
                            var v38 float32 = ((-2.25 - 0.0) * 1.5)
                            _ = v38
                        case "hello":
                            p1 = (p1 + "a")
                        default:
                            fmt.Println(fi11, p1)
                        }
                        v39 := GenS7{F0: -1, F1: -1, F2: 1}
                        _ = v39
                    }
                    fmt.Println(v7, p1, v7)
                    fmt.Println(p1, p1, v7)
                }
                _ = fi11
                _ = v12
            }
            p0 -= p0
        }
        var v40 int = fi8
        _ = fi8
        _ = v40
    }
    if false {
        v41 := GenS6{F0: 1, F1: (v7 ^ v7)}
        v42 := []int{-9223372036854775808, 32768, 40}
        v43 := GenS8{F0: 0, F1: 320}
        _ = v41
        _ = v42
        _ = v43
    }
    _ = p0
    _ = p1
    _ = v7
    return p1
}

func main() {
    var v1 int8 = 0
    var v2 int32 = -273
    var v3 uint8 = 198
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]int)
    v6 := make(map[int]int)
    v3 = v3
    v44 := []int{833, -2147483648}
    var v45 int32 = 2147483647
    v46 := GenS5{F0: 2, F1: "a", F2: -46}
    if (1 != (v1 & 0)) {
        switch 813 {
        case -1:
            switch ("\n\t" + "a") {
            case "foo bar":
                if (9223372036854775807 > 910) {
                    v1 = v1
                    fmt.Println(v2, v1, v45)
                    for fi47 := 0; fi47 < 13; fi47++ {
                        var v48 float64 = (-2.25 + float64(float32(float64(-2.25))))
                        v3 = (v3 | v3)
                        var v49 uint64 = 0
                        _ = fi47
                        _ = v48
                        _ = v49
                    }
                }
            case "":
                if len(v44) > 0 {
                    v44[0] = 9223372036854775807
                }
            default:
                v2 = (v2 | (v2 * (v45 % ((v45) | 1))))
            }
        case 128:
            fmt.Println(v2, v3)
        case -785:
            fmt.Println(v1)
        default:
            if len(v44) >= 2 {
                v44 = v44[1:len(v44)]
            }
        }
    } else {
        if (true && false) {
            v50 := GenS9{F0: -976}
            v1 = 0
            v44 = append(v44, -232)
            _ = v50
        } else {
            v51 := GenS10{F0: 937, F1: 2147483647, F2: ""}
            _ = len(v4)
            var v52 int64 = -9223372036854775808
            _ = v51
            _ = v52
        }
    }
    _ = len(v5)
    v1 += (v1 >> ((1) & 63))
    for fi53 := 0; fi53 < 12; fi53++ {
        v3 += (v3 / ((156) | 1))
        v54 := GenS7{F0: v1, F1: int32(v3), F2: (v45 * (v45 | 937))}
        _ = fi53
        _ = v54
    }
    var v55 float64 = 0.0
    var v56 uint16 = 168
    v55 = (1.5 + (v55 + ((float64((float32(1e10) + (0.1 / (float32(v55) / 1.0)))) * 1e10) * 0.0)))
    if false {
        v5[725] = -9223372036854775808
        switch 128 {
        case -2147483648:
            switch (("\n\t" + ("\n\t" + (("a" + "hello") + ""))) + "hello") {
            case "hello":
                if true {
                    fmt.Println(v2)
                } else {
                    var v57 uint32 = 2147483647
                    var v58 int64 = -9223372036854775808
                    _ = v57
                    _ = v58
                }
            case "":
                var v59 string = ("hello" + "foo bar")
                _ = v59
            }
        case -9223372036854775808:
            if false {
                var v60 int16 = -32768
                for fi61 := 0; fi61 < 13; fi61++ {
                    switch "" {
                    case "\n\t":
                        var v62 int8 = v1
                        _ = v62
                    case "":
                        if true {
                            v3 += (v3 | v3)
                        }
                    case "foo bar":
                        var v63 uint64 = 0
                        _ = v63
                    }
                    _ = fi61
                }
                v64 := GenS11{F0: (-(v2)), F1: "a", F2: 1}
                _ = v60
                _ = v64
            }
        case -269:
            v3 += v3
        }
        for fi65 := 0; fi65 < 20; fi65++ {
            var v66 int64 = -71
            _ = fi65
            _ = v66
        }
    }
    var v67 bool = true
    if false {
        delete(v4, 2)
        var v68 int8 = -128
        v69 := GenS2{F0: 1000}
        _ = v68
        _ = v69
    } else {
        if v67 {
            var v70 float64 = float64((((1.5 + ((-2.25 / 1e10) - 0.1)) - 3.14159) + float32(float64((3.14159 / 1e10)))))
            switch 826 {
            case 694:
                switch (("\n\t" + ("" + "hello")) + "a") {
                case "\n\t":
                    switch 32767 {
                    case 32768:
                        fmt.Println(v55, v3, v1)
                    case -934:
                        var v71 float64 = (float64(0.0) * (v70 * v70))
                        _ = v71
                    default:
                        switch 2147483647 {
                        case 1:
                            for fi72 := 0; fi72 < 4; fi72++ {
                                v73 := v6[fi72]
                                v45 = int32(v1)
                                if len(v44) > 0 {
                                    v44[0] = v73
                                }
                                _ = fi72
                                _ = v73
                            }
                        default:
                            if (v67 && (!(v2 > v2))) {
                                switch "a" {
                                case "foo bar":
                                    v74 := GenS12{F0: 65535}
                                    _ = v74
                                }
                                if (128 > 32768) {
                                    v75 := v5[0]
                                    if (331 != 255) {
                                        v76 := GenS9{F0: 1}
                                        v1 = (-(v1))
                                        delete(v5, 0)
                                        _ = v76
                                    }
                                    for fi77 := 0; fi77 < 13; fi77++ {
                                        v75 -= v75
                                        v44 = append(v44, -641)
                                        var v78 int8 = (^(v1))
                                        _ = fi77
                                        _ = v78
                                    }
                                    _ = v75
                                }
                            }
                        }
                    }
                case "hello":
                    switch "a" {
                    case "foo bar":
                        v2 = (^(v2))
                    }
                }
            case 548:
                var v79 uint16 = (v56 + (v56 * (^(v56))))
                _ = v79
            case 4294967295:
                fmt.Println(v1, v1)
            default:
                for fi80 := 0; fi80 < 8; fi80++ {
                    v56 += 1
                    switch fi80 {
                    case -1:
                        switch (((("a" + "") + ("" + ("a" + (("a" + "a") + "foo bar")))) + "foo bar") + ((("hello" + "\n\t") + "hello") + "hello")) {
                        case "foo bar":
                            var v81 uint16 = (^(v56))
                            _ = v81
                        }
                    case 0:
                        v55 = 0.0
                    case -372:
                        v82 := GenS13{F0: (-(v1))}
                        _ = v82
                    default:
                        if (-32768 < -1) {
                            v55 = (3.14159 + v55)
                        }
                    }
                    _ = fi80
                }
            }
            var v83 int32 = (v45 | v2)
            _ = v70
            _ = v83
        } else {
            switch 42 {
            case 9223372036854775807:
                v2 = -2147483648
            }
        }
        v84 := GenS14{F0: (v45 * (v2 << ((345) & 63))), F1: (v1 * (^(v1))), F2: 2}
        _ = v84
    }
    if (2 <= v56) {
        v85 := GenS14{F0: v2, F1: v1, F2: 0}
        var v86 uint16 = v56
        _ = v85
        _ = v86
    }
    v2 = v2
    if len(v44) > 0 {
        v44[0] = 836
    }
    var v87 int16 = -390
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v67:", v67)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v55
    _ = v56
    _ = v6
    _ = v67
    _ = v87
}
