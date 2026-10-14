package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 string }
func genFn0() int64 {
    if false {
        var v7 int8 = 124
        _ = v7
    } else {
        v8 := []int{}
        var v9 int = 0
        v9 -= 42
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 440, F1: 290, F2: 0}
    var v11 int32 = -747
    _ = v10
    _ = v11
    return -1
}
func genFn1(p0 int32, p1 int) string {
    var v12 uint8 = 2
    for fi13 := 0; fi13 < 1; fi13++ {
        if false {
            v12 += 116
        } else {
            var v14 string = "\n\t"
            var v15 bool = ((^(p0)) < 2147483647)
            p1 -= (fi13 * (fi13 % (((fi13 ^ (p1 - (-(p1))))) | 1)))
            _ = v14
            _ = v15
        }
        for fi16 := 0; fi16 < 12; fi16++ {
            var v17 int32 = p0
            _ = fi16
            _ = v17
        }
        _ = fi13
    }
    var v18 float64 = ((3.14159 / float64(1.5)) * (-2.25 - -2.25))
    _ = p0
    _ = p1
    _ = v12
    _ = v18
    return "\n\t"
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = (1 == 884)
    v3 := make(map[int]bool)
    v4 := make(map[int]int)
    var v5 int = 32767
    var v6 float64 = 1e10
    if true {
        v19 := GenS1{F0: 18446744073709551615, F1: 636, F2: -9223372036854775808}
        var v20 uint8 = 82
        _ = v19
        _ = v20
    }
    if (v2 && v2) {
        var v21 bool = (false || (1 >= 1))
        if ((v5 + 32768) >= 127) {
            var v22 float64 = v6
            var v23 int = (v5 & 0)
            if (!v2) {
                v24 := GenS1{F0: 18446744073709551615, F1: 255, F2: v5}
                var v25 bool = v2
                v23 -= v5
                _ = v24
                _ = v25
            } else {
                var v26 string = "a"
                var v27 uint8 = 1
                v26 = v26
                _ = v26
                _ = v27
            }
            _ = v22
            _ = v23
        }
        _ = len(v1)
        _ = v21
    } else {
        switch "hello" {
        case "":
            v28 := v1[158]
            _ = v28
        case "a":
            var v29 int8 = -128
            _ = v29
        default:
            fmt.Println(v5, v6)
        }
        var v30 int32 = 785
        _ = v30
    }
    v6 = 3.14159
    v31 := v3[v5]
    if v31 {
        var v32 uint32 = 65535
        switch (("\n\t" + "a") + ("\n\t" + "a")) {
        case "\n\t":
            for fi33 := 0; fi33 < 13; fi33++ {
                if ((v5 % ((-907) | 1)) > fi33) {
                    if false {
                        v34 := []int{(v5 - fi33)}
                        switch (v5 % (((fi33 * 820)) | 1)) {
                        case -1:
                            var v35 uint8 = 17
                            _ = v35
                        case 9223372036854775807:
                            var v36 bool = (!false)
                            _ = v36
                        case 2147483647:
                            var v37 int64 = -1
                            _ = v37
                        default:
                            _ = len(v4)
                        }
                        _ = v34
                    } else {
                        v38 := []int{v5}
                        fmt.Println(v5)
                        switch "foo bar" {
                        case "hello":
                            if len(v38) > 0 {
                                v38[0] = 42
                            }
                        case "":
                            v39 := GenS1{F0: 697, F1: 9, F2: -717}
                            _ = v39
                        case "a":
                            for fi40 := 0; fi40 < 4; fi40++ {
                                var v41 float32 = 3.14159
                                _ = fi40
                                _ = v41
                            }
                        default:
                            var v42 int16 = 42
                            _ = v42
                        }
                        _ = v38
                    }
                    switch "hello" {
                    case "a":
                        var v43 uint64 = 2
                        _ = v43
                    default:
                        v44 := v3[127]
                        _ = v44
                    }
                }
                for fi45 := 0; fi45 < 13; fi45++ {
                    var v46 int16 = -148
                    switch "\n\t" {
                    case "foo bar":
                        var v47 float32 = 0.1
                        _ = v47
                    }
                    _ = fi45
                    _ = v46
                }
                _ = fi33
            }
        default:
            v48 := GenS2{F0: 4294967295, F1: 0}
            _ = v48
        }
        for fi49 := 0; fi49 < 20; fi49++ {
            v31 = ((1.5 + v6) < 3.14159)
            for fi50 := 0; fi50 < 5; fi50++ {
                _ = len(v4)
                if false {
                    switch "foo bar" {
                    case "\n\t":
                        _ = len(v4)
                        v51 := GenS1{F0: 42, F1: 65535, F2: v5}
                        _ = v51
                    default:
                        fmt.Println(v2)
                    }
                    fmt.Println(v31)
                }
                _ = fi50
            }
            v52 := GenS2{F0: 988, F1: 60}
            _ = fi49
            _ = v52
        }
        _ = v32
    }
    v5 = (v5 % (((v5 / ((v5) | 1))) | 1))
    fmt.Println(v2, v31)
    var v53 int32 = -2147483648
    fmt.Println(v5, v31)
    var v54 uint32 = 2147483647
    if (v54 > (v54 & (^(v54)))) {
        v55 := GenS3{F0: v5, F1: (v5 == (v5 ^ v5))}
        var v56 int = (v5 | v5)
        _ = v55
        _ = v56
    }
    if (-460 >= -120) {
        for fi57 := 0; fi57 < 15; fi57++ {
            if (1 == 0) {
                v58 := v1[v54]
                v6 = v6
                _ = v58
            }
            for fi59 := 0; fi59 < 4; fi59++ {
                v54 += (v54 / ((966) | 1))
                v54 = (^(v54))
                var v60 uint64 = 212
                if (!v31) {
                    v61 := GenS1{F0: v60, F1: 65535, F2: v5}
                    v62 := []int{v5}
                    fmt.Println(v31, v53, fi57)
                    var v63 float64 = v6
                    _ = v61
                    _ = v62
                    _ = v63
                }
                switch "hello" {
                case "":
                    if true {
                        v3[(v5 >> ((4294967295) & 63))] = (1 != 153)
                        v53 = 0
                        v60 += v60
                        var v64 uint64 = 477
                        _ = v64
                    }
                case "foo bar":
                    v53 = (v53 / ((v53) | 1))
                }
                _ = fi59
                _ = v60
            }
            _ = fi57
        }
        var v65 string = "hello"
        _ = v65
    } else {
        switch "a" {
        case "foo bar":
            var v66 uint32 = (v54 - (v54 << ((1) & 63)))
            _ = v66
        case "":
            for fi67 := 0; fi67 < 18; fi67++ {
                v68 := GenS4{F0: "a"}
                _ = fi67
                _ = v68
            }
        }
    }
    v4[v5] = 4294967295
    var v69 uint64 = 933
    var v70 bool = v31
    for fi71 := 0; fi71 < 13; fi71++ {
        v4[fi71] = (fi71 >> ((0) & 63))
        fmt.Println(v5)
        _ = len(v3)
        _ = fi71
    }
    if (v31 && (!(v2 && (-128 >= 43)))) {
        v5 = (v5 | (v5 ^ (v5 ^ 32768)))
        var v72 int64 = -9223372036854775808
        _ = v72
    } else {
        var v73 string = "foo bar"
        _ = v73
    }
    var v74 int = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v6
    _ = v69
    _ = v70
    _ = v74
}
