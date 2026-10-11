package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 int8; F1 uint16 }
type GenS6 struct { F0 uint64; F1 int8; F2 int16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 string }
type GenS10 struct { F0 string; F1 bool }
type GenS11 struct { F0 uint64 }
func genFn0() bool {
    var v7 int8 = 44
    if false {
        if (0 == -13) {
            v8 := []int{900}
            switch "" {
            case "a":
                var v9 int16 = 1
                _ = v9
            case "":
                fmt.Println(v7)
            case "\n\t":
                var v10 string = "\n\t"
                _ = v10
            default:
                var v11 int64 = 0
                _ = v11
            }
            v12 := GenS1{F0: (v7 & -128)}
            _ = v12
            _ = v8
        }
        var v13 int = 128
        if false {
            if true {
                for fi14 := 0; fi14 < 17; fi14++ {
                    var v15 uint32 = 2147483647
                    _ = fi14
                    _ = v15
                }
                v7 -= v7
                v13 = v13
            } else {
                v16 := GenS1{F0: (v7 >> ((42) & 63))}
                var v17 int = v13
                _ = v16
                _ = v17
            }
            v7 = v7
            var v18 uint = 160
            _ = v18
        } else {
            v13 = v13
            v13 = (v13 % (((v13 << ((2) & 63))) | 1))
            if false {
                v13 = v13
            } else {
                switch "a" {
                case "\n\t":
                    fmt.Println(v7, v7)
                case "foo bar":
                    if false {
                        v19 := GenS2{F0: 123, F1: 99}
                        v13 = v13
                        for fi20 := 0; fi20 < 20; fi20++ {
                            if false {
                                var v21 float32 = (0.0 / 1.0)
                                var v22 uint16 = 65535
                                for fi23 := 0; fi23 < 4; fi23++ {
                                    var v24 int32 = 27
                                    v25 := []int{(fi20 - (fi23 - 42))}
                                    v7 += (v7 % ((v7) | 1))
                                    _ = fi23
                                    _ = v24
                                    _ = v25
                                }
                                _ = v21
                                _ = v22
                            }
                            v26 := []int{(fi20 + (fi20 ^ v13)), (v13 * (fi20 & v13))}
                            v7 += (v7 / (((v7 << ((2) & 63))) | 1))
                            _ = fi20
                            _ = v26
                        }
                        _ = v19
                    }
                case "a":
                    v27 := GenS2{F0: 4294967295, F1: 2}
                    _ = v27
                }
            }
        }
        _ = v13
    } else {
        v28 := []int{}
        var v29 uint32 = 493
        _ = v28
        _ = v29
    }
    if false {
        var v30 uint = 449
        _ = v30
    } else {
        v31 := []int{-230, -138}
        _ = v31
    }
    v32 := []int{209, 127, 9223372036854775807}
    _ = v32
    _ = v7
    return true
}
func genFn1(p0 uint, p1 int8, p2 uint) int32 {
    v33 := []int{32768, 2147483647}
    for fi34 := 0; fi34 < 5; fi34++ {
        p0 += 4294967295
        if len(v33) >= 2 {
            v33 = v33[1:len(v33)]
        }
        if len(v33) >= 2 {
            v33 = v33[1:len(v33)]
        }
        _ = fi34
    }
    p1 = -1
    var v35 float64 = float64(1e10)
    _ = p0
    _ = p1
    _ = p2
    _ = v33
    _ = v35
    return 2147483647
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = 42
    v3 := []uint32{2, 2147483648, 4294967295, 2147483647}
    var v4 uint64 = 42
    var v5 bool = true
    var v6 uint64 = (v4 - (v4 << ((42) & 63)))
    fmt.Println(v5)
    delete(v1, -651)
    var v36 int64 = 42
    var v37 int16 = (v2 % ((v2) | 1))
    var v38 int64 = (v36 * -9223372036854775808)
    v1[int(v36)] = 65535
    v39 := GenS1{F0: 71}
    var v40 uint = 782
    var v41 uint16 = 0
    var v42 bool = v5
    for fi43 := 0; fi43 < 1; fi43++ {
        for fi44 := 0; fi44 < 10; fi44++ {
            if true {
                _ = len(v1)
                var v45 uint64 = v6
                switch -1 {
                case 9223372036854775807:
                    v46 := GenS3{F0: (v6 | (^(v4))), F1: 1}
                    _ = v46
                case -2147483648:
                    fmt.Println(v42)
                }
                _ = v45
            } else {
                v47 := v1[int(v38)]
                _ = v47
            }
            if v5 {
                var v48 int = int(v36)
                if ((v41 / ((42) | 1)) >= (v41 ^ (v41 % (((v41 | 1)) | 1)))) {
                    if v42 {
                        if true {
                            v49 := GenS4{F0: v42, F1: v40}
                            var v50 int8 = -128
                            if false {
                                var v51 string = ((("\n\t" + "hello") + "a") + "\n\t")
                                _ = v51
                            } else {
                                v36 = (v38 >> ((v40) & 63))
                            }
                            v52 := GenS3{F0: 1, F1: v41}
                            if (327 <= 42) {
                                v53 := GenS5{F0: 127, F1: (v41 / (((v41 << ((494) & 63))) | 1))}
                                _ = v53
                            } else {
                                v54 := GenS6{F0: 42, F1: 127, F2: v37}
                                if (!v5) {
                                    v55 := GenS2{F0: 4294967295, F1: 2147483648}
                                    _ = len(v3)
                                    _ = v55
                                } else {
                                    switch ("\n\t" + "") {
                                    case "":
                                        var v56 uint32 = 1
                                        _ = v56
                                    }
                                    fmt.Println(v2)
                                    fmt.Println(v50, v4, v41)
                                }
                                v38 += v36
                                if false {
                                    var v57 float32 = -2.25
                                    if len(v3) > 0 {
                                        v3[0] = 1
                                    }
                                    if v42 {
                                        var v58 bool = (v5 || (((-2.25 / 1.5) - 1.5) > (3.14159 / (((1.5 * (1e10 / (((0.0 + (1.5 * 0.1)) / 0.1) * (0.1 + 1e10)))) / 3.14159) - (3.14159 - float64((0.0 * (0.0 - ((((((v57 + 1.5) - v57) + v57) - -2.25) - 1e10) + v57)))))))))
                                        var v59 uint = v40
                                        v60 := GenS7{F0: v58}
                                        _ = v58
                                        _ = v59
                                        _ = v60
                                    }
                                    _ = v57
                                }
                                _ = v54
                            }
                            _ = v49
                            _ = v50
                            _ = v52
                        }
                    } else {
                        var v61 string = "\n\t"
                        var v62 int64 = (-(v38))
                        _ = v61
                        _ = v62
                    }
                    if v5 {
                        if (true && (v5 || true)) {
                            if (-2.25 == 1.5) {
                                switch (^(fi44)) {
                                case 127:
                                    var v63 float64 = 1.5
                                    _ = v63
                                case -1:
                                    if true {
                                        var v64 float32 = (0.0 - (0.1 + 0.0))
                                        _ = v64
                                    } else {
                                        if ("\n\t" < "\n\t") {
                                            var v65 int8 = 90
                                            var v66 int8 = 42
                                            if (false || v42) {
                                                v65 -= v65
                                                var v67 int8 = v66
                                                var v68 int = fi44
                                                _ = v67
                                                _ = v68
                                            }
                                            _ = v65
                                            _ = v66
                                        }
                                    }
                                default:
                                    switch (-(fi43)) {
                                    case 361:
                                        v40 = v40
                                    default:
                                        fmt.Println(v38, v37, fi44)
                                        v69 := GenS1{F0: 42}
                                        _ = len(v3)
                                        _ = v69
                                    }
                                }
                                v70 := GenS8{F0: uint(v41)}
                                switch (fi43 >> (((v40 | uint(v41))) & 63)) {
                                case 1:
                                    v71 := GenS1{F0: 20}
                                    _ = v71
                                case 128:
                                    switch ("hello" + "foo bar") {
                                    case "":
                                        if true {
                                            var v72 uint = 42
                                            _ = v72
                                        } else {
                                            v73 := GenS4{F0: ("foo bar" < "hello"), F1: v40}
                                            var v74 float32 = float32(3.14159)
                                            v38 = 0
                                            var v75 int = (fi44 % (((fi44 << ((42) & 63))) | 1))
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            _ = v73
                                            _ = v74
                                            _ = v75
                                        }
                                    default:
                                        var v76 float64 = 3.14159
                                        _ = v76
                                    }
                                    var v77 int32 = 1
                                    _ = len(v1)
                                    _ = v77
                                }
                                _ = v70
                            } else {
                                v48 += (fi44 & fi44)
                                fmt.Println(v41)
                                v48 = (fi43 + 9223372036854775807)
                            }
                        } else {
                            v78 := GenS9{F0: "\n\t"}
                            _ = v78
                        }
                        v79 := v1[407]
                        _ = v79
                    } else {
                        if (genFn1((v40 % ((uint(v41)) | 1)), -128, 18446744073709551615) <= -1) {
                            if (!v42) {
                                var v80 uint = (^(v40))
                                v36 += v38
                                _ = v80
                            } else {
                                v38 = 42
                                fmt.Println(v48, v42, fi44)
                            }
                        } else {
                            var v81 string = "hello"
                            fmt.Println(fi43, v6, v37)
                            _ = v81
                        }
                        var v82 uint8 = 0
                        _ = len(v1)
                        _ = v82
                    }
                } else {
                    var v83 int = v48
                    _ = v83
                }
                _ = v48
            }
            switch 798 {
            case 2147483647:
                v4 -= 2
            }
            _ = fi44
        }
        _ = fi43
    }
    v84 := GenS10{F0: "", F1: ((0.0 + 1e10) == ((3.14159 + 1.5) - 1e10))}
    var v85 int8 = -1
    var v86 string = "hello"
    v87 := GenS11{F0: v6}
    _ = len(v1)
    for fi88 := 0; fi88 < 13; fi88++ {
        var v89 uint16 = (v41 << (((v40 / ((v40) | 1))) & 63))
        _ = fi88
        _ = v89
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v85:", v85)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v84
    _ = v85
    _ = v86
    _ = v87
}
