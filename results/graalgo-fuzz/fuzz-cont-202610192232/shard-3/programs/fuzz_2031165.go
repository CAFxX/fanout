package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8; F1 bool; F2 uint16 }
type GenS6 struct { F0 uint; F1 uint32 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint32; F1 uint8 }
func genFn0(p0 uint16, p1 int64) string {
    v6 := GenS1{F0: 1}
    var v7 bool = (545 < 4294967295)
    var v8 uint = 2
    p1 -= p1
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 uint64 = 994
        var v11 float64 = 0.1
        fmt.Println(p1, p0, v8)
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return "foo bar"
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = false
    v3 := make(map[int]uint32)
    v4 := make(map[string]uint32)
    var v5 float64 = 1e10
    v5 = (0.0 - (float64(-2.25) / 1.5))
    v2 = v2
    fmt.Println(v2, v2)
    for fi12 := 0; fi12 < 4; fi12++ {
        v5 -= v5
        fmt.Println(v2)
        _ = fi12
    }
    switch ("\n\t" + "") {
    case "a":
        if false {
            fmt.Println(v5, v2)
        }
    case "hello":
        var v13 uint16 = 255
        _ = v13
    default:
        var v14 string = (("hello" + "a") + "hello")
        _ = v14
    }
    v5 = v5
    var v15 uint8 = 0
    v16 := GenS1{F0: 339}
    if false {
        var v17 float64 = v5
        _ = v17
    }
    switch "foo bar" {
    case "hello":
        v15 += 40
    case "foo bar":
        v18 := []int{}
        _ = v18
    case "\n\t":
        var v19 float32 = -2.25
        _ = v19
    }
    v20 := GenS1{F0: 185}
    if (0.0 < v5) {
        delete(v1, 4294967295)
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        for fi22 := 0; fi22 < 11; fi22++ {
            var v23 float32 = -2.25
            _ = fi22
            _ = v23
        }
        v5 -= v5
        _ = fi21
    }
    if (v5 >= v5) {
        fmt.Println(v5, v2)
        v4[("a" + (("a" + "a") + ""))] = 926
        if v2 {
            var v24 int = 52
            var v25 int16 = 664
            for fi26 := 0; fi26 < 14; fi26++ {
                switch (fi26 ^ (fi26 % ((fi26) | 1))) {
                case 42:
                    if v2 {
                        if ((!(498 >= 523)) && (2 > 65535)) {
                            for fi27 := 0; fi27 < 7; fi27++ {
                                v25 -= v25
                                v25 = v25
                                v28 := GenS2{F0: -1, F1: 4294967295}
                                _ = fi27
                                _ = v28
                            }
                            if v2 {
                                v15 -= (v15 << ((506) & 63))
                            } else {
                                v15 -= 2
                                v29 := GenS3{F0: -128, F1: 42}
                                v30 := GenS3{F0: 100, F1: -1}
                                _ = v29
                                _ = v30
                            }
                        }
                    } else {
                        v2 = (2 != 1)
                        if v2 {
                            fmt.Println(v24, v5)
                        } else {
                            v24 += fi26
                        }
                        if v2 {
                            for fi31 := 0; fi31 < 2; fi31++ {
                                var v32 string = "foo bar"
                                if false {
                                    if v2 {
                                        if v2 {
                                            v25 = (v25 & v25)
                                        } else {
                                            v25 -= v25
                                            v32 = v32
                                        }
                                        if false {
                                            var v33 int = (fi31 | 127)
                                            _ = v33
                                        } else {
                                            switch "hello" {
                                            case "\n\t":
                                                v24 = fi31
                                            case "a":
                                                v34 := []int{fi31}
                                                _ = v34
                                            default:
                                                v35 := GenS4{F0: 294}
                                                _ = v35
                                            }
                                            var v36 int16 = -1
                                            if v2 {
                                                var v37 float32 = 0.0
                                                fmt.Println(v36)
                                                _ = v37
                                            }
                                            if (!((42 != uint64(v15)) || false)) {
                                                fmt.Println(fi31, v25, fi26)
                                            } else {
                                                if false {
                                                    v1[2147483647] = 128
                                                    v38 := GenS4{F0: 0}
                                                    v39 := []int{}
                                                    _ = v38
                                                    _ = v39
                                                } else {
                                                    v40 := GenS5{F0: 42, F1: v2, F2: 255}
                                                    v41 := GenS1{F0: 387}
                                                    v32 = v32
                                                    _ = v40
                                                    _ = v41
                                                }
                                            }
                                            v24 -= v24
                                            v42 := GenS5{F0: -128, F1: true, F2: 0}
                                            _ = v36
                                            _ = v42
                                        }
                                    } else {
                                        var v43 uint32 = 4294967295
                                        switch -1 {
                                        case -2147483648:
                                            v32 = ""
                                        default:
                                            if (!(v32 >= "foo bar")) {
                                                v4[v32] = v43
                                                if v2 {
                                                    switch (v32 + (v32 + v32)) {
                                                    case "hello":
                                                        var v44 int = fi31
                                                        _ = v44
                                                    case "a":
                                                        var v45 int = v24
                                                        _ = v45
                                                    }
                                                    v5 = (v5 + ((v5 + float64(0.0)) / v5))
                                                } else {
                                                    v4[(("a" + "") + (v32 + (v32 + "")))] = (v43 % ((235) | 1))
                                                    fmt.Println(v32)
                                                    v2 = v2
                                                }
                                            } else {
                                                var v46 bool = v2
                                                var v47 uint16 = 42
                                                if v2 {
                                                    fmt.Println(v25)
                                                    if v46 {
                                                        switch (v24 - (fi26 % (((v24 >> ((4294967295) & 63))) | 1))) {
                                                        case 2147483647:
                                                            if false {
                                                                var v48 int16 = v25
                                                                if (uint(v47) == 990) {
                                                                    v15 += v15
                                                                    var v49 uint64 = 0
                                                                    switch (fi31 & (v24 >> ((29) & 63))) {
                                                                    case -2147483648:
                                                                        if v2 {
                                                                            delete(v3, (-(v24)))
                                                                            v47 += uint16(v49)
                                                                        } else {
                                                                            if false {
                                                                                switch fi31 {
                                                                                case -515:
                                                                                    switch 9223372036854775807 {
                                                                                    case -9223372036854775808:
                                                                                        v50 := GenS4{F0: 1}
                                                                                        _ = v50
                                                                                    }
                                                                                }
                                                                            }
                                                                            v51 := v4[""]
                                                                            _ = v51
                                                                        }
                                                                    case 393:
                                                                        fmt.Println(v25, v15, v15)
                                                                        v2 = v46
                                                                    }
                                                                    _ = v49
                                                                }
                                                                _ = v48
                                                            } else {
                                                                var v52 uint8 = (v15 / ((v15) | 1))
                                                                _ = v52
                                                            }
                                                        case 972:
                                                            v53 := GenS6{F0: 1, F1: 42}
                                                            _ = v53
                                                        }
                                                        var v54 int64 = 0
                                                        _ = v54
                                                    }
                                                    v15 = (v15 % ((v15) | 1))
                                                }
                                                _ = v46
                                                _ = v47
                                            }
                                        }
                                        _ = v43
                                    }
                                    if v2 {
                                        if v2 {
                                            var v55 bool = v2
                                            v2 = v55
                                            _ = len(v3)
                                            _ = v55
                                        } else {
                                            if v2 {
                                                v25 = v25
                                                switch v32 {
                                                case "hello":
                                                    fmt.Println(v24, fi31)
                                                case "\n\t":
                                                    fmt.Println(v15, v5, v2)
                                                default:
                                                    if (!(0 <= 118)) {
                                                        v32 = v32
                                                        var v56 int64 = -9223372036854775808
                                                        _ = v56
                                                    } else {
                                                        v24 = v24
                                                        v25 -= v25
                                                        if false {
                                                            var v57 int32 = 1
                                                            v15 = 2
                                                            if true {
                                                                if (v2 || false) {
                                                                    switch fi31 {
                                                                    case -34:
                                                                        v15 = (v15 >> ((0) & 63))
                                                                    case -861:
                                                                        v5 += v5
                                                                    default:
                                                                        var v58 string = v32
                                                                        _ = v58
                                                                    }
                                                                    v32 = v32
                                                                }
                                                            }
                                                            v32 = "a"
                                                            if (((((3.14159 / 1e10) / (1.5 / (0.0 - (-2.25 - (v5 + (-2.25 / float64(float32(v5)))))))) * v5) - v5) >= (v5 * v5)) {
                                                                if v2 {
                                                                    switch v32 {
                                                                    case "hello":
                                                                        v3[fi26] = 0
                                                                    default:
                                                                        var v59 uint8 = v15
                                                                        _ = v59
                                                                    }
                                                                    v3[(fi31 ^ fi26)] = 731
                                                                }
                                                                if v2 {
                                                                    var v60 int32 = (v57 - (v57 >> ((42) & 63)))
                                                                    v32 = "hello"
                                                                    if (("\n\t" + "a") < v32) {
                                                                        _ = len(v3)
                                                                        var v61 int64 = 1
                                                                        v62 := GenS5{F0: -82, F1: ((v2 && false) || false), F2: 1}
                                                                        v57 = 889
                                                                        _ = v61
                                                                        _ = v62
                                                                    }
                                                                    if (v2 || v2) {
                                                                        v63 := []int{(fi26 << ((4294967295) & 63)), (-(fi26)), (fi31 >> ((628) & 63)), fi26}
                                                                        v15 += (^(v15))
                                                                        if ((-1 != 367) && ((v25 & 1000) > v25)) {
                                                                            if (false || (100 > 127)) {
                                                                                v64 := GenS6{F0: 42, F1: 1}
                                                                                var v65 uint8 = (v15 - v15)
                                                                                v65 = (v15 + (v15 ^ v65))
                                                                                v15 = 37
                                                                                fmt.Println(v32, v60)
                                                                                _ = v64
                                                                                _ = v65
                                                                            }
                                                                        } else {
                                                                            var v66 int16 = (v25 | v25)
                                                                            v67 := GenS6{F0: 2, F1: 2147483648}
                                                                            v24 += 803
                                                                            _ = v66
                                                                            _ = v67
                                                                        }
                                                                        var v68 float32 = 0.0
                                                                        _ = v63
                                                                        _ = v68
                                                                    }
                                                                    if v2 {
                                                                        v69 := []int{(fi26 << ((1) & 63)), (-(fi31)), fi26}
                                                                        _ = v69
                                                                    } else {
                                                                        v70 := GenS3{F0: 1, F1: 0}
                                                                        _ = v70
                                                                    }
                                                                    _ = v60
                                                                } else {
                                                                    v57 -= (v57 & (-(v57)))
                                                                    v71 := GenS5{F0: -1, F1: v2, F2: 678}
                                                                    v72 := GenS2{F0: -9223372036854775808, F1: 1}
                                                                    fmt.Println(v24, v57)
                                                                    _ = v71
                                                                    _ = v72
                                                                }
                                                            } else {
                                                                v2 = ((v24 * -323) >= 0)
                                                            }
                                                            v73 := GenS7{F0: (v24 | fi26)}
                                                            _ = v57
                                                            _ = v73
                                                        }
                                                        v74 := []int{}
                                                        if ((v32 == (v32 + v32)) || v2) {
                                                            fmt.Println(v24, fi31, v25)
                                                            v24 = fi26
                                                            v75 := GenS8{F0: 65535, F1: 147}
                                                            _ = v75
                                                        } else {
                                                            v15 = 158
                                                            switch ((v32 + "\n\t") + v32) {
                                                            case "":
                                                                v76 := GenS8{F0: 1, F1: 127}
                                                                _ = v76
                                                            case "a":
                                                                fmt.Println(v25, fi31, v2)
                                                            }
                                                        }
                                                        _ = v74
                                                    }
                                                }
                                            }
                                            var v77 uint = 545
                                            _ = v77
                                        }
                                        v5 += 1.5
                                        v78 := []int{128}
                                        _ = v78
                                    } else {
                                        v79 := []int{(fi26 & (-(v24))), 646, fi26, (v24 >> ((938) & 63))}
                                        var v80 uint16 = 1
                                        var v81 int32 = -123
                                        _ = v79
                                        _ = v80
                                        _ = v81
                                    }
                                    var v82 uint64 = 18446744073709551615
                                    _ = v82
                                }
                                _ = fi31
                                _ = v32
                            }
                        }
                    }
                default:
                    var v83 uint16 = 2
                    _ = v83
                }
                _ = fi26
            }
            _ = v24
            _ = v25
        }
    } else {
        var v84 int16 = 1000
        v2 = v2
        switch "" {
        case "foo bar":
            if ((-(v84)) >= 1) {
                if v2 {
                    var v85 float32 = 0.0
                    _ = v85
                }
                var v86 bool = true
                _ = v86
            }
        case "hello":
            var v87 uint32 = 65535
            _ = v87
        }
        _ = v84
    }
    for fi88 := 0; fi88 < 10; fi88++ {
        var v89 uint8 = (v15 & (v15 / (((v15 | 2)) | 1)))
        for fi90 := 0; fi90 < 3; fi90++ {
            var v91 int64 = 1
            if v2 {
                v5 -= v5
                v92 := GenS3{F0: 0, F1: 0}
                v91 = 39
                _ = v92
            } else {
                var v93 uint8 = v89
                _ = v93
            }
            var v94 uint = 2
            _ = fi90
            _ = v91
            _ = v94
        }
        if true {
            v89 = (v15 + 69)
            v2 = v2
        }
        _ = fi88
        _ = v89
    }
    v15 += (^(v15))
    fmt.Println(v15)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
