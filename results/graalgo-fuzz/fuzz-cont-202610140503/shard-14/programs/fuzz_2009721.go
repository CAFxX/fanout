package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 uint; F1 int16; F2 string }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint32; F1 string; F2 int }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint32; F1 string; F2 uint16 }
func genFn0() uint64 {
    var v6 uint64 = 1
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 uint64 = (v6 / (((v6 / ((20) | 1))) | 1))
        if false {
            v9 := []int{}
            fmt.Println(v6)
            switch "a" {
            case "foo bar":
                var v10 int8 = -128
                _ = v10
            case "":
                v11 := GenS1{F0: 4294967295}
                _ = v11
            default:
                fmt.Println(v8, v6)
            }
            _ = v9
        }
        if true {
            for fi12 := 0; fi12 < 8; fi12++ {
                v6 = v6
                _ = fi12
            }
        }
        _ = fi7
        _ = v8
    }
    v13 := GenS2{F0: 2147483647}
    v14 := GenS2{F0: 995}
    _ = v13
    _ = v14
    _ = v6
    return (^(v6))
}
func genFn1(p0 uint64, p1 int16, p2 int8) bool {
    v15 := []int{2147483647, -9223372036854775808, -9223372036854775808}
    var v16 uint = uint(p0)
    for fi17 := 0; fi17 < 10; fi17++ {
        for fi18 := 0; fi18 < 5; fi18++ {
            var v19 uint16 = 2
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return (true || true)
}

func main() {
    var v1 int = -9223372036854775808
    var v2 float64 = 0.1
    v3 := make(map[int]int)
    var v4 uint = 42
    v5 := make(map[int]uint32)
    var v20 float64 = (float64(1e10) * (float64(1.5) * v2))
    fmt.Println(v2, v20, v20)
    fmt.Println(v20, v4)
    var v21 int64 = -1
    var v22 int = (v1 % ((v1) | 1))
    for fi23 := 0; fi23 < 13; fi23++ {
        v24 := []int{}
        _ = fi23
        _ = v24
    }
    switch ("" + ("\n\t" + "a")) {
    case "":
        var v25 uint8 = 255
        _ = v25
    case "foo bar":
        fmt.Println(v1)
    case "hello":
        for fi26 := 0; fi26 < 19; fi26++ {
            if true {
                if false {
                    var v27 uint64 = 42
                    var v28 uint16 = 1
                    var v29 int = (v1 ^ 4294967295)
                    _ = v27
                    _ = v28
                    _ = v29
                } else {
                    var v30 int32 = -2147483648
                    v20 -= float64(float32((((float64(((0.0 * ((1.5 / float32(1.5)) + 0.0)) / (0.1 + 0.1))) * v20) - (float64(1.5) * (float64(((0.0 - 0.0) / -2.25)) / (v2 / (1e10 - v2))))) / (float64((0.1 - 0.0)) + v2))))
                    _ = v30
                }
            } else {
                for fi31 := 0; fi31 < 12; fi31++ {
                    v32 := GenS2{F0: 1}
                    var v33 float64 = v2
                    if (!false) {
                        v34 := GenS3{F0: 18446744073709551615, F1: 0}
                        var v35 uint16 = 703
                        fmt.Println(v35)
                        v36 := GenS4{F0: (v4 & v4), F1: 42, F2: "a"}
                        if true {
                            var v37 string = "a"
                            _ = v37
                        }
                        _ = v34
                        _ = v35
                        _ = v36
                    } else {
                        switch ("a" + "\n\t") {
                        case "foo bar":
                            if (-1 < -729) {
                                if false {
                                    v22 = 1
                                    v38 := []int{(^(v22)), (v1 << ((18446744073709551615) & 63)), int(v21), fi26}
                                    var v39 uint = (v4 & v4)
                                    _ = v38
                                    _ = v39
                                } else {
                                    fmt.Println(v21)
                                    _ = len(v3)
                                }
                            } else {
                                fmt.Println(fi26, v4)
                                v40 := []int{v22}
                                _ = v40
                            }
                            if true {
                                if true {
                                    var v41 uint64 = 18446744073709551615
                                    if (!true) {
                                        v22 += (fi31 & (v22 ^ 62))
                                        v2 = v2
                                        v2 -= (v33 - 3.14159)
                                        if (false && true) {
                                            var v42 bool = ("" != ("\n\t" + ("hello" + "a")))
                                            v20 = (float64(1.5) + v33)
                                            _ = v42
                                        }
                                        var v43 string = "hello"
                                        _ = v43
                                    } else {
                                        var v44 int8 = -128
                                        v3[int(v44)] = (fi31 * 128)
                                        v4 = (^(v4))
                                        _ = v44
                                    }
                                    _ = v41
                                }
                            } else {
                                var v45 int = (v1 * (v22 >> ((v4) & 63)))
                                _ = v45
                            }
                            switch "foo bar" {
                            case "foo bar":
                                v46 := GenS3{F0: 42, F1: 1}
                                _ = v46
                            case "hello":
                                v20 = (v33 - float64(0.0))
                            default:
                                delete(v3, fi31)
                            }
                        case "\n\t":
                            v33 = 3.14159
                        case "a":
                            switch ("\n\t" + "") {
                            case "\n\t":
                                fmt.Println(v2, v21)
                                switch -369 {
                                case 127:
                                    if true {
                                        fmt.Println(v20)
                                        switch (("foo bar" + "hello") + "hello") {
                                        case "hello":
                                            var v47 uint32 = 921
                                            _ = v47
                                        case "\n\t":
                                            v33 = float64((0.1 - -2.25))
                                        default:
                                            v1 = (v22 ^ 2147483647)
                                        }
                                        fmt.Println(v4)
                                    }
                                case 567:
                                    v48 := v5[v22]
                                    _ = v48
                                case 56:
                                    if true {
                                        var v49 int32 = 431
                                        v50 := v5[80]
                                        _ = v49
                                        _ = v50
                                    } else {
                                        delete(v3, 232)
                                        switch fi26 {
                                        case -399:
                                            v51 := GenS2{F0: 42}
                                            _ = v51
                                        case 1:
                                            switch (v1 | (fi31 << ((v4) & 63))) {
                                            case 254:
                                                var v52 uint = v4
                                                _ = v52
                                            }
                                        case 32768:
                                            switch "" {
                                            case "hello":
                                                switch v22 {
                                                case 127:
                                                    fmt.Println(v1, v22, v22)
                                                default:
                                                    v53 := GenS3{F0: 4294967295, F1: 18446744073709551615}
                                                    _ = v53
                                                }
                                            case "\n\t":
                                                v54 := []int{v1}
                                                _ = v54
                                            default:
                                                if (false && (42 != 675)) {
                                                    v55 := GenS1{F0: 65535}
                                                    _ = v55
                                                } else {
                                                    if false {
                                                        _ = len(v3)
                                                        if true {
                                                            var v56 uint64 = 2
                                                            _ = v56
                                                        }
                                                        if ("\n\t" == "\n\t") {
                                                            fmt.Println(v21, fi31, v21)
                                                            v21 -= v21
                                                            v22 = fi26
                                                        }
                                                    } else {
                                                        var v57 uint = (^(v4))
                                                        _ = v57
                                                    }
                                                    v58 := []int{605, 127, fi31}
                                                    fmt.Println(v1)
                                                    _ = v58
                                                }
                                            }
                                        }
                                    }
                                default:
                                    v59 := GenS1{F0: 42}
                                    _ = v59
                                }
                            case "":
                                var v60 float64 = 0.1
                                _ = v60
                            case "foo bar":
                                switch ("\n\t" + "") {
                                case "a":
                                    switch (^(fi26)) {
                                    case 9223372036854775807:
                                        v61 := GenS5{F0: (v1 >> ((707) & 63))}
                                        _ = v61
                                    default:
                                        if false {
                                            v3[30] = 86
                                        } else {
                                            v2 -= -2.25
                                        }
                                    }
                                case "foo bar":
                                    switch (fi26 | fi31) {
                                    case 42:
                                        var v62 int16 = -987
                                        _ = v62
                                    case 132:
                                        v63 := GenS6{F0: 42, F1: "foo bar", F2: -637}
                                        if false {
                                            var v64 bool = ((("\n\t" + "\n\t") != "\n\t") || true)
                                            _ = v64
                                        } else {
                                            var v65 int8 = -40
                                            delete(v3, (fi31 - (v22 / (((v1 >> ((v4) & 63))) | 1))))
                                            _ = v65
                                        }
                                        v22 -= 777
                                        _ = v63
                                    }
                                case "hello":
                                    if true {
                                        if (-1 <= -128) {
                                            var v66 float64 = float64(1e10)
                                            _ = v66
                                        }
                                        var v67 uint = (v4 % ((993) | 1))
                                        _ = v67
                                    } else {
                                        v68 := GenS1{F0: 65535}
                                        var v69 string = "\n\t"
                                        var v70 uint16 = 32767
                                        _ = v68
                                        _ = v69
                                        _ = v70
                                    }
                                }
                            }
                        default:
                            var v71 int16 = int16(v1)
                            _ = v71
                        }
                        fmt.Println(v4, v33, fi26)
                    }
                    var v72 float32 = 3.14159
                    v73 := []int{}
                    _ = fi31
                    _ = v32
                    _ = v33
                    _ = v72
                    _ = v73
                }
                var v74 uint = 2
                v21 = int64(v74)
                _ = v74
            }
            _ = fi26
        }
    }
    if (143 != 32767) {
        var v75 int32 = 2147483647
        v76 := GenS7{F0: 0}
        var v77 float64 = (-2.25 * v2)
        _ = v75
        _ = v76
        _ = v77
    }
    fmt.Println(v22)
    v78 := GenS8{F0: 946, F1: "\n\t", F2: uint16(v4)}
    fmt.Println(v1)
    v1 = (v22 & v22)
    if (v4 <= v4) {
        for fi79 := 0; fi79 < 13; fi79++ {
            for fi80 := 0; fi80 < 3; fi80++ {
                v81 := GenS6{F0: 204, F1: "\n\t", F2: v22}
                var v82 float64 = 1e10
                _ = fi80
                _ = v81
                _ = v82
            }
            v83 := []int{fi79, 2147483647, fi79, v22}
            _ = fi79
            _ = v83
        }
    } else {
        var v84 float64 = v20
        if true {
            var v85 uint16 = 0
            _ = v85
        }
        v4 -= 42
        _ = v84
    }
    v86 := GenS6{F0: 0, F1: ("foo bar" + "a"), F2: (v22 / (((v1 - (^(v22)))) | 1))}
    if true {
        var v87 float32 = 0.1
        fmt.Println(v20, v22, v20)
        v88 := []int{}
        _ = v87
        _ = v88
    } else {
        for fi89 := 0; fi89 < 6; fi89++ {
            _ = len(v3)
            var v90 int = 1
            for fi91 := 0; fi91 < 4; fi91++ {
                fmt.Println(v20)
                v20 += v20
                _ = fi91
            }
            _ = fi89
            _ = v90
        }
    }
    fmt.Println(v20, v1, v22)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v78
    _ = v86
}
