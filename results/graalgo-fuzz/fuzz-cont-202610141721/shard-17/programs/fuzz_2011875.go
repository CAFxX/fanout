package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int64; F1 int32; F2 uint32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS9 struct { F0 int16; F1 int8; F2 uint64 }
func genFn0() uint64 {
    var v5 float32 = -2.25
    v5 += v5
    _ = v5
    return 4294967295
}
func genFn1(p0 int64, p1 uint64) uint8 {
    p1 = 0
    v6 := []int{}
    _ = p0
    _ = p1
    _ = v6
    return 199
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = -128
    v3 := make(map[string]int)
    v4 := make(map[uint32]int)
    fmt.Println(v2)
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            v9 := GenS1{F0: ((fi7 - 1) >= (fi8 << ((18446744073709551615) & 63)))}
            v2 = (v2 / ((v2) | 1))
            v2 += v2
            _ = fi8
            _ = v9
        }
        switch "foo bar" {
        case "a":
            for fi10 := 0; fi10 < 15; fi10++ {
                var v11 uint8 = 128
                _ = fi10
                _ = v11
            }
        case "hello":
            var v12 int = fi7
            _ = v12
        case "":
            v13 := []int{fi7, fi7}
            _ = v13
        default:
            fmt.Println(v2, v2)
        }
        v14 := GenS1{F0: true}
        _ = fi7
        _ = v14
    }
    if true {
        if true {
            v15 := GenS1{F0: (990 == 0)}
            v2 -= 42
            _ = v15
        }
        fmt.Println(v2)
        var v16 int64 = 1
        _ = v16
    } else {
        if true {
            var v17 float32 = float32((0.0 + -2.25))
            switch -164 {
            case -1:
                var v18 int = 128
                _ = v18
            case 128:
                if true {
                    var v19 uint = 18446744073709551615
                    fmt.Println(v2, v2)
                    v20 := []int{int(v2), 9223372036854775807, 50, -55}
                    _ = v19
                    _ = v20
                }
            case -389:
                v21 := GenS1{F0: ((1 <= 32767) && (!((((("" + "\n\t") + (("" + "\n\t") + "foo bar")) + ("a" + (("" + ("foo bar" + "\n\t")) + (("foo bar" + "a") + "foo bar")))) + "\n\t") < (("hello" + "") + "foo bar"))))}
                _ = v21
            }
            _ = v17
        } else {
            v2 = v2
            fmt.Println(v2)
        }
        v22 := GenS1{F0: ((0.1 < float64(-2.25)) || false)}
        fmt.Println(v2)
        _ = v22
    }
    v23 := []int{}
    v24 := GenS2{F0: (v2 * v2), F1: ""}
    for fi25 := 0; fi25 < 9; fi25++ {
        switch ("foo bar" + "") {
        case "hello":
            v23 = append(v23, fi25)
        case "foo bar":
            var v26 uint8 = 82
            _ = v26
        case "\n\t":
            if (1e10 != 0.0) {
                v27 := GenS2{F0: (v2 * v2), F1: "hello"}
                switch (fi25 * int(v2)) {
                case -695:
                    fmt.Println(fi25, fi25)
                case 32768:
                    var v28 uint16 = 0
                    _ = v28
                case -622:
                    for fi29 := 0; fi29 < 1; fi29++ {
                        var v30 string = (((("hello" + ((("\n\t" + "\n\t") + ("" + ("" + "\n\t"))) + "foo bar")) + "hello") + "") + "hello")
                        _ = fi29
                        _ = v30
                    }
                default:
                    v31 := GenS2{F0: 0, F1: ""}
                    _ = v31
                }
                if len(v23) > 0 {
                    v23[0] = (fi25 - (fi25 - (fi25 << ((42) & 63))))
                }
                _ = v27
            } else {
                var v32 uint8 = 2
                v33 := GenS1{F0: false}
                _ = v32
                _ = v33
            }
        default:
            var v34 uint16 = 218
            _ = v34
        }
        var v35 uint32 = 2147483647
        for fi36 := 0; fi36 < 3; fi36++ {
            switch fi36 {
            case 127:
                v35 += v35
            default:
                v23 = append(v23, (fi36 / ((fi36) | 1)))
            }
            v3[("a" + "\n\t")] = fi25
            v37 := GenS3{F0: (("" + ("" + "foo bar")) + "a")}
            if (969 > 42) {
                if false {
                    var v38 float64 = 3.14159
                    var v39 uint = 0
                    v40 := GenS3{F0: "hello"}
                    v41 := GenS3{F0: "a"}
                    _ = v38
                    _ = v39
                    _ = v40
                    _ = v41
                }
                var v42 string = ""
                _ = v42
            } else {
                v2 -= (-(v2))
            }
            v43 := v3["hello"]
            _ = fi36
            _ = v37
            _ = v43
        }
        _ = fi25
        _ = v35
    }
    fmt.Println(v2)
    for fi44 := 0; fi44 < 10; fi44++ {
        if (float32((1.5 - float64(float32((0.1 - ((float64(0.1) * float64((((1.5 + 1.5) / (-2.25 + 0.1)) * (((((0.1 - float32(float64((((((1.5 - 0.0) - 1.5) / 1e10) * -2.25) + 1.5)))) - (-2.25 * 1.5)) + 0.1) / 1.5) / (-2.25 + (0.1 * float32((3.14159 / -2.25)))))))) / -2.25)))))) > ((1.5 - -2.25) + 1.5)) {
            if (!(!false)) {
                var v45 float32 = 0.1
                for fi46 := 0; fi46 < 14; fi46++ {
                    v47 := GenS4{F0: 9223372036854775807, F1: 2147483647, F2: 42}
                    if len(v23) > 0 {
                        v23[0] = (fi46 << ((624) & 63))
                    }
                    if (!false) {
                        fmt.Println(fi44, fi46)
                        v23 = append(v23, -100)
                        var v48 uint8 = 0
                        _ = v48
                    } else {
                        v49 := GenS1{F0: true}
                        if (fi46 >= (fi44 + (fi44 & (^(fi46))))) {
                            fmt.Println(fi46)
                            var v50 uint32 = 65535
                            _ = v50
                        } else {
                            v2 -= 77
                        }
                        fmt.Println(fi46, v45)
                        _ = v49
                    }
                    _ = fi46
                    _ = v47
                }
                var v51 uint64 = 2
                _ = v45
                _ = v51
            } else {
                var v52 uint = 431
                _ = v52
            }
            if (3.14159 != -2.25) {
                var v53 float64 = (0.0 / 1.5)
                if len(v23) > 0 {
                    v23[0] = fi44
                }
                var v54 float32 = (3.14159 + -2.25)
                _ = v53
                _ = v54
            } else {
                if (!false) {
                    if (0.1 == (0.1 + 1e10)) {
                        switch ("" + ("\n\t" + "\n\t")) {
                        case "hello":
                            if false {
                                for fi55 := 0; fi55 < 15; fi55++ {
                                    switch (fi55 ^ (fi55 >> ((846) & 63))) {
                                    case 0:
                                        var v56 float32 = 1e10
                                        _ = v56
                                    case 128:
                                        if true {
                                            switch fi44 {
                                            case -9223372036854775808:
                                                var v57 bool = (false || (!false))
                                                _ = v57
                                            }
                                            var v58 int32 = 42
                                            v23 = append(v23, (fi44 * 2147483647))
                                            _ = v58
                                        }
                                    case 1:
                                        fmt.Println(v2, fi55)
                                    default:
                                        fmt.Println(fi55, fi55)
                                    }
                                    _ = fi55
                                }
                            }
                        case "":
                            v2 = v2
                        case "\n\t":
                            if (true || (!true)) {
                                switch 2147483647 {
                                case -2147483648:
                                    v59 := GenS5{F0: 42}
                                    _ = v59
                                default:
                                    var v60 uint = 42
                                    _ = v60
                                }
                                var v61 uint32 = 2147483647
                                _ = v61
                            }
                        default:
                            if (182 <= 31) {
                                v62 := GenS6{F0: false}
                                _ = v62
                            } else {
                                v2 = 0
                                var v63 float32 = -2.25
                                _ = v63
                            }
                        }
                        _ = len(v23)
                        v23 = append(v23, -9223372036854775808)
                    }
                    switch 128 {
                    case 0:
                        if (42 >= 0) {
                            var v64 int64 = 330
                            var v65 uint32 = 2147483647
                            v64 = v64
                            _ = v64
                            _ = v65
                        } else {
                            fmt.Println(v2, v2)
                            v66 := GenS7{F0: false}
                            var v67 int32 = int32(fi44)
                            _ = v66
                            _ = v67
                        }
                    case 782:
                        var v68 float64 = 1.5
                        _ = v68
                    default:
                        if (true && (64 <= (v2 << ((88) & 63)))) {
                            var v69 uint64 = 701
                            if false {
                                _ = len(v23)
                            } else {
                                var v70 int = 0
                                _ = v70
                            }
                            _ = v69
                        }
                    }
                } else {
                    if true {
                        switch (fi44 & (fi44 / ((fi44) | 1))) {
                        case 0:
                            v71 := GenS8{F0: 0, F1: -2147483648, F2: 1}
                            _ = v71
                        }
                    } else {
                        if true {
                            fmt.Println(fi44, fi44)
                        } else {
                            if (2 >= 682) {
                                v2 = (v2 / ((v2) | 1))
                            } else {
                                var v72 float64 = float64(((-2.25 * (0.1 - 1.5)) + 1e10))
                                _ = v72
                            }
                        }
                        if (42 > -2147483648) {
                            if len(v23) > 0 {
                                v23[0] = -714
                            }
                        }
                        v2 -= v2
                    }
                    if ((18446744073709551615 == 194) || false) {
                        var v73 uint = 2
                        _ = v73
                    } else {
                        _ = len(v1)
                        v74 := GenS9{F0: -32768, F1: (v2 | v2), F2: 78}
                        fmt.Println(v2, v2)
                        _ = v74
                    }
                }
                if true {
                    if len(v23) >= 2 {
                        v23 = v23[1:len(v23)]
                    }
                } else {
                    switch (fi44 * 128) {
                    case -416:
                        fmt.Println(fi44)
                    case 0:
                        v2 = 108
                    case -549:
                        v2 -= (v2 + -1)
                    default:
                        v2 -= v2
                    }
                    for fi75 := 0; fi75 < 8; fi75++ {
                        v76 := GenS7{F0: false}
                        _ = fi75
                        _ = v76
                    }
                    if (-730 < (fi44 / (((fi44 ^ fi44)) | 1))) {
                        for fi77 := 0; fi77 < 9; fi77++ {
                            v78 := GenS3{F0: ("\n\t" + "\n\t")}
                            _ = fi77
                            _ = v78
                        }
                    } else {
                        v2 -= v2
                        switch fi44 {
                        case 2147483647:
                            var v79 uint8 = 15
                            _ = v79
                        case 32768:
                            var v80 uint16 = 255
                            _ = v80
                        case -1:
                            switch (fi44 ^ (fi44 % ((9223372036854775807) | 1))) {
                            case -160:
                                v23 = append(v23, 9223372036854775807)
                            default:
                                var v81 uint64 = 4294967295
                                _ = v81
                            }
                        }
                        if len(v23) >= 2 {
                            v23 = v23[1:len(v23)]
                        }
                    }
                }
                fmt.Println(v2, v2)
            }
            v2 = v2
        } else {
            v82 := GenS1{F0: false}
            var v83 int16 = 42
            _ = v82
            _ = v83
        }
        v84 := v3[("\n\t" + (((((("foo bar" + "a") + "foo bar") + ("foo bar" + ("foo bar" + "foo bar"))) + "") + "hello") + ("\n\t" + "hello")))]
        v84 -= v84
        _ = fi44
        _ = v84
    }
    v2 += (v2 % ((v2) | 1))
    switch 1 {
    case -1:
        v2 = (v2 & 0)
    case -458:
        var v85 uint64 = genFn0()
        _ = v85
    default:
        v86 := GenS9{F0: -645, F1: (v2 | 1), F2: 808}
        _ = v86
    }
    var v87 uint64 = 2
    var v88 uint = 400
    var v89 uint = 4294967295
    if false {
        var v90 uint64 = (^(v87))
        var v91 int = 4294967295
        fmt.Println(v88, v2)
        _ = v90
        _ = v91
    } else {
        fmt.Println(v2)
        var v92 int64 = 42
        if (false || true) {
            v88 -= v89
        } else {
            switch 760 {
            case -1:
                switch 128 {
                case 2147483647:
                    var v93 int32 = 0
                    _ = v93
                case -2147483648:
                    v89 += (v88 + 42)
                case 0:
                    delete(v1, 2147483647)
                }
            default:
                if len(v23) > 0 {
                    v23[0] = 587
                }
            }
            fmt.Println(v2, v88, v2)
            v89 += 18446744073709551615
        }
        _ = v92
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v87:", v87)
    fmt.Println("v88:", v88)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v87
    _ = v88
    _ = v89
}
