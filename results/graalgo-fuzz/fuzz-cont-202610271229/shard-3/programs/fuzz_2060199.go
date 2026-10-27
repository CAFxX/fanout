package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 int; F1 uint8 }
type GenS6 struct { F0 int; F1 uint32 }
type GenS7 struct { F0 int64 }
func genFn0(p0 bool) int16 {
    var v4 uint = 1
    switch -2147483648 {
    case 127:
        if p0 {
            var v5 int8 = 1
            for fi6 := 0; fi6 < 16; fi6++ {
                v4 = 897
                _ = fi6
            }
            v7 := []int{127}
            _ = v5
            _ = v7
        } else {
            var v8 uint64 = 0
            if p0 {
                switch 32767 {
                case -2147483648:
                    v9 := GenS1{F0: v4}
                    _ = v9
                case -876:
                    if false {
                        var v10 int8 = 127
                        fmt.Println(v8, v10)
                        for fi11 := 0; fi11 < 4; fi11++ {
                            for fi12 := 0; fi12 < 13; fi12++ {
                                v13 := []int{}
                                fmt.Println(fi11, v4, v4)
                                _ = fi12
                                _ = v13
                            }
                            var v14 uint32 = uint32(v8)
                            v14 += (v14 * v14)
                            _ = fi11
                            _ = v14
                        }
                        _ = v10
                    }
                default:
                    var v15 uint16 = 65535
                    _ = v15
                }
            }
            _ = v8
        }
    case 2147483647:
        v16 := GenS2{F0: 2147483648, F1: 570}
        _ = v16
    default:
        p0 = (!(((2147483648 != 205) && (p0 || p0)) && p0))
    }
    v17 := GenS2{F0: 2, F1: 2147483647}
    _ = p0
    _ = v17
    _ = v4
    return -32768
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = false
    var v3 int16 = -1
    var v18 int = 2147483647
    v18 = (^(v18))
    var v19 int = (v18 * v18)
    v18 += (v19 * v18)
    for fi20 := 0; fi20 < 4; fi20++ {
        var v21 int64 = 9223372036854775807
        v18 += fi20
        _ = fi20
        _ = v21
    }
    if (42 != 42) {
        for fi22 := 0; fi22 < 19; fi22++ {
            v18 += v19
            switch "a" {
            case "foo bar":
                v19 += (v18 * fi22)
            case "\n\t":
                if (-2147483648 == 32767) {
                    v18 = 32767
                    fmt.Println(v3, v18)
                    if ((-1 <= 1) && false) {
                        switch 13 {
                        case 32767:
                            v19 = 127
                        }
                        var v23 int64 = 9223372036854775807
                        fmt.Println(v19)
                        _ = v23
                    } else {
                        for fi24 := 0; fi24 < 15; fi24++ {
                            fmt.Println(v18)
                            if false {
                                v25 := GenS3{F0: 18446744073709551615}
                                _ = v25
                            }
                            _ = fi24
                        }
                        if (!(!v2)) {
                            for fi26 := 0; fi26 < 11; fi26++ {
                                var v27 float64 = 1e10
                                if (((float32(v27) * 0.1) * ((3.14159 + (-2.25 * float32(v27))) * 1.5)) != -2.25) {
                                    var v28 float64 = 3.14159
                                    if v2 {
                                        switch (fi26 | 2147483647) {
                                        case 128:
                                            v29 := GenS1{F0: 65535}
                                            _ = v29
                                        case 4294967295:
                                            v19 = v18
                                        }
                                        v30 := GenS2{F0: 2147483648, F1: 127}
                                        _ = v30
                                    } else {
                                        var v31 float64 = 0.0
                                        v3 = 42
                                        _ = v31
                                    }
                                    _ = v28
                                } else {
                                    if (!(-9223372036854775808 == 0)) {
                                        var v32 uint = 42
                                        if v2 {
                                            switch v18 {
                                            case 487:
                                                var v33 uint = 255
                                                _ = v33
                                            case 569:
                                                v3 += -657
                                            }
                                            if v2 {
                                                var v34 uint64 = uint64(v32)
                                                fmt.Println(v19, v32, fi22)
                                                if (!v2) {
                                                    switch (v19 % (((v18 + fi26)) | 1)) {
                                                    case -2147483648:
                                                        v35 := GenS4{F0: 197, F1: 2}
                                                        _ = v35
                                                    case 127:
                                                        fmt.Println(v19, v32)
                                                    }
                                                    v36 := []int{(fi26 | 47), (fi26 << (((v32 + v32)) & 63)), (v19 & (v19 % (((fi22 * fi22)) | 1)))}
                                                    v37 := GenS4{F0: 25, F1: 32767}
                                                    _ = v36
                                                    _ = v37
                                                }
                                                _ = v34
                                            } else {
                                                v2 = (("hello" + "hello") == "foo bar")
                                                fmt.Println(v2)
                                                if true {
                                                    fmt.Println(v2)
                                                }
                                            }
                                        }
                                        var v38 uint8 = 255
                                        _ = v32
                                        _ = v38
                                    }
                                }
                                v18 -= fi22
                                _ = fi26
                                _ = v27
                            }
                            var v39 float64 = -2.25
                            _ = v39
                        } else {
                            for fi40 := 0; fi40 < 7; fi40++ {
                                var v41 uint32 = 833
                                _ = fi40
                                _ = v41
                            }
                            var v42 int = -870
                            _ = v42
                        }
                        fmt.Println(v19)
                    }
                } else {
                    v19 = fi22
                    v43 := GenS2{F0: 65535, F1: v19}
                    _ = v43
                }
            case "hello":
                var v44 int32 = 1
                _ = v44
            default:
                for fi45 := 0; fi45 < 1; fi45++ {
                    v19 = fi22
                    if false {
                        v46 := v1[65535]
                        v47 := GenS5{F0: fi22, F1: 127}
                        switch (("\n\t" + "hello") + "a") {
                        case "hello":
                            var v48 int16 = (v3 ^ v3)
                            _ = v48
                        }
                        if false {
                            var v49 int64 = -9223372036854775808
                            fmt.Println(fi45)
                            if (310 >= 104) {
                                var v50 int = v19
                                _ = v50
                            } else {
                                if (2147483647 < 2147483648) {
                                    switch v46 {
                                    case 127:
                                        v3 = (v3 & -1)
                                    default:
                                        v51 := []int{4294967295, (fi22 ^ fi22), v18}
                                        _ = v51
                                    }
                                    if (!v2) {
                                        var v52 bool = (412 > 4294967295)
                                        delete(v1, 336)
                                        v46 -= fi22
                                        fmt.Println(v19, fi22)
                                        _ = v52
                                    }
                                } else {
                                    _ = len(v1)
                                }
                                var v53 uint8 = 255
                                v49 += (v49 | (v49 & (v49 >> ((uint(v53)) & 63))))
                                _ = v53
                            }
                            _ = v49
                        }
                        _ = v46
                        _ = v47
                    } else {
                        if false {
                            var v54 bool = false
                            v55 := GenS3{F0: 2}
                            v56 := GenS6{F0: fi45, F1: 690}
                            _ = len(v1)
                            var v57 string = "a"
                            _ = v54
                            _ = v55
                            _ = v56
                            _ = v57
                        } else {
                            switch v18 {
                            case -1:
                                v3 -= (v3 % (((^(v3))) | 1))
                            case 857:
                                _ = len(v1)
                            case -302:
                                v58 := GenS5{F0: v19, F1: 250}
                                _ = v58
                            }
                            var v59 uint64 = 488
                            fmt.Println(fi45, v3)
                            _ = v59
                        }
                        switch (fi45 | fi45) {
                        case -452:
                            fmt.Println(v18, v3)
                        default:
                            if v2 {
                                switch (v18 << ((4294967295) & 63)) {
                                case -2147483648:
                                    v18 -= v18
                                    v60 := GenS7{F0: 9223372036854775807}
                                    v19 = (v19 % ((32767) | 1))
                                    var v61 float64 = 0.1
                                    v62 := GenS4{F0: 255, F1: 805}
                                    _ = v60
                                    _ = v61
                                    _ = v62
                                case 42:
                                    fmt.Println(fi22, v19)
                                case 0:
                                    if v2 {
                                        v18 = v18
                                        var v63 bool = (((("" + "hello") + "") + ("foo bar" + (("foo bar" + ("" + "a")) + (("\n\t" + "hello") + "\n\t")))) >= "foo bar")
                                        v1[65535] = 32768
                                        v19 = (v18 ^ 127)
                                        var v64 uint8 = 127
                                        v65 := GenS5{F0: (v19 * (v18 | (-(v18)))), F1: v64}
                                        _ = v63
                                        _ = v64
                                        _ = v65
                                    } else {
                                        fmt.Println(fi45, v19)
                                    }
                                }
                            } else {
                                _ = len(v1)
                                if ((-25 <= 0) || v2) {
                                    v3 -= (v3 - (v3 + (v3 / (((v3 * v3)) | 1))))
                                    v66 := GenS1{F0: 1}
                                    switch 9223372036854775807 {
                                    case 127:
                                        v2 = v2
                                    case -9223372036854775808:
                                        v1[65535] = 127
                                    default:
                                        var v67 int32 = 2147483647
                                        _ = v67
                                    }
                                    _ = v66
                                } else {
                                    v2 = v2
                                    switch "\n\t" {
                                    case "a":
                                        var v68 bool = (32767 == 1)
                                        _ = v68
                                    case "":
                                        if (-195 < 0) {
                                            if (457 != 4294967295) {
                                                v1[247] = fi22
                                                fmt.Println(fi45, v3, fi22)
                                                var v69 uint64 = 4294967295
                                                var v70 int64 = -1
                                                _ = v69
                                                _ = v70
                                            }
                                            v71 := GenS7{F0: int64(v3)}
                                            v72 := []int{888, fi22, v18, (fi45 / (((fi22 / ((fi45) | 1))) | 1))}
                                            _ = v71
                                            _ = v72
                                        } else {
                                            v1[2] = 625
                                            if false {
                                                fmt.Println(fi45)
                                            } else {
                                                fmt.Println(v2, fi45)
                                                v73 := GenS3{F0: 42}
                                                var v74 uint8 = 212
                                                _ = v73
                                                _ = v74
                                            }
                                        }
                                    case "\n\t":
                                        if (!v2) {
                                            var v75 bool = v2
                                            _ = v75
                                        }
                                    }
                                    switch "\n\t" {
                                    case "":
                                        v18 = fi45
                                    case "\n\t":
                                        v19 += v18
                                    case "a":
                                        v76 := []int{(fi22 * (fi45 - (fi45 - (v19 << ((42) & 63))))), (v19 / ((-884) | 1)), fi22}
                                        _ = v76
                                    default:
                                        if ((v3 + (v3 ^ (v3 / ((42) | 1)))) <= v3) {
                                            var v77 int32 = 42
                                            if true {
                                                v78 := []int{}
                                                _ = v78
                                            } else {
                                                var v79 int8 = 42
                                                if false {
                                                    var v80 uint16 = 768
                                                    var v81 bool = ((v80 % ((1) | 1)) != 0)
                                                    _ = v80
                                                    _ = v81
                                                }
                                                delete(v1, 2147483647)
                                                _ = v79
                                            }
                                            _ = v77
                                        } else {
                                            if false {
                                                var v82 int64 = 42
                                                _ = v82
                                            } else {
                                                var v83 uint64 = 4294967295
                                                v84 := GenS4{F0: 2, F1: 32767}
                                                switch (fi45 ^ fi45) {
                                                case 128:
                                                    fmt.Println(fi22)
                                                }
                                                _ = v83
                                                _ = v84
                                            }
                                        }
                                    }
                                }
                            }
                        }
                        var v85 uint32 = 263
                        _ = v85
                    }
                    _ = fi45
                }
            }
            v2 = v2
            _ = fi22
        }
        fmt.Println(v19, v3, v18)
    }
    fmt.Println(v18, v2, v2)
    var v86 int8 = 47
    var v87 string = "hello"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v86
    _ = v87
}
