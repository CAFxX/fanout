package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 int; F1 int64; F2 int8 }
type GenS3 struct { F0 int64; F1 uint32; F2 uint32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int; F1 string; F2 int8 }
type GenS6 struct { F0 uint64; F1 int }
type GenS7 struct { F0 int8 }
func genFn0(p0 uint32, p1 uint8) int64 {
    var v4 uint32 = p0
    v4 = (v4 ^ (v4 - p0))
    for fi5 := 0; fi5 < 8; fi5++ {
        var v6 float64 = float64((float32((1.5 * 3.14159)) - ((-2.25 - float32(((1.5 / 0.1) / ((1.5 - ((1e10 / 1.0) + 0.0)) + 1.5)))) / 3.14159)))
        v7 := []int{}
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 uint = 345
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v4
    return 42
}
func genFn1() bool {
    for fi10 := 0; fi10 < 15; fi10++ {
        v11 := GenS1{F0: 394, F1: 0}
        var v12 int32 = 42
        v12 -= v12
        _ = fi10
        _ = v11
        _ = v12
    }
    switch (("foo bar" + ("\n\t" + "hello")) + "\n\t") {
    case "hello":
        switch "\n\t" {
        case "hello":
            for fi13 := 0; fi13 < 16; fi13++ {
                var v14 uint64 = 0
                _ = fi13
                _ = v14
            }
        case "":
            var v15 uint16 = 0
            _ = v15
        default:
            v16 := []int{}
            _ = v16
        }
    default:
        v17 := GenS1{F0: 985, F1: 0}
        _ = v17
    }
    for fi18 := 0; fi18 < 16; fi18++ {
        var v19 bool = (-117 > -1)
        v19 = ("a" >= "foo bar")
        var v20 uint32 = 4294967295
        _ = fi18
        _ = v19
        _ = v20
    }
    v21 := []int{-569, 127, 128, 793}
    _ = v21
    return (!true)
}

func main() {
    var v1 uint16 = 32767
    var v2 int8 = -128
    var v3 float32 = float32(3.14159)
    var v22 uint8 = 0
    fmt.Println(v2)
    if true {
        v23 := GenS1{F0: 65535, F1: uint64(v22)}
        v24 := GenS1{F0: 96, F1: 0}
        _ = v23
        _ = v24
    }
    switch -1 {
    case 42:
        var v25 float32 = -2.25
        _ = v25
    case 32768:
        v26 := []int{302, -444}
        _ = v26
    case -643:
        v22 = (v22 ^ (v22 % (((v22 + v22)) | 1)))
    default:
        v27 := []int{-1, -423, -9223372036854775808, 32767}
        _ = v27
    }
    var v28 uint8 = (^(v22))
    if ((0.1 * 1.5) != float32((1e10 + float64((v3 * v3))))) {
        var v29 float64 = float64((v3 + (v3 / v3)))
        v30 := GenS2{F0: 767, F1: -83, F2: 127}
        switch -336 {
        case -1:
            v29 = -2.25
        case 42:
            for fi31 := 0; fi31 < 15; fi31++ {
                var v32 bool = (false || true)
                v33 := GenS3{F0: -1, F1: 0, F2: 4294967295}
                _ = fi31
                _ = v32
                _ = v33
            }
        case 2147483647:
            v22 -= (v22 ^ (v28 / ((v22) | 1)))
        default:
            for fi34 := 0; fi34 < 1; fi34++ {
                for fi35 := 0; fi35 < 17; fi35++ {
                    v22 -= (^(v22))
                    if (fi35 != int(v1)) {
                        var v36 int64 = 519
                        var v37 uint64 = 516
                        v38 := []int{(^(fi35)), 1, fi35}
                        _ = v36
                        _ = v37
                        _ = v38
                    }
                    _ = fi35
                }
                v39 := []int{fi34, (fi34 >> ((42) & 63))}
                _ = fi34
                _ = v39
            }
        }
        _ = v29
        _ = v30
    }
    var v40 uint16 = v1
    v41 := GenS2{F0: -940, F1: 42, F2: v2}
    switch -292 {
    case 32768:
        var v42 int16 = 564
        _ = v42
    case -218:
        v22 = v28
    case 4294967295:
        fmt.Println(v22)
    }
    var v43 int16 = 42
    fmt.Println(v2, v2)
    if true {
        for fi44 := 0; fi44 < 6; fi44++ {
            if (!false) {
                v1 += v40
            } else {
                for fi45 := 0; fi45 < 8; fi45++ {
                    v46 := GenS2{F0: (fi44 ^ 42), F1: 1, F2: v2}
                    _ = fi45
                    _ = v46
                }
            }
            if ((-2.25 + (v3 / (v3 - (v3 + (v3 - ((((v3 + -2.25) + v3) - 0.0) - ((1e10 / v3) - v3))))))) < v3) {
                switch "hello" {
                case "\n\t":
                    fmt.Println(v28, v2, v43)
                default:
                    var v47 int = (fi44 / (((-(fi44))) | 1))
                    _ = v47
                }
                for fi48 := 0; fi48 < 16; fi48++ {
                    var v49 int32 = 1
                    switch ("" + (("\n\t" + ("foo bar" + "\n\t")) + "")) {
                    case "\n\t":
                        v50 := []int{}
                        _ = v50
                    case "":
                        fmt.Println(v22, v22, fi48)
                    case "foo bar":
                        v1 += v1
                    }
                    var v51 uint16 = (v40 * v40)
                    var v52 float32 = float32(((1.5 - -2.25) / 1e10))
                    _ = fi48
                    _ = v49
                    _ = v51
                    _ = v52
                }
                v28 = v28
            }
            var v53 int16 = v43
            _ = fi44
            _ = v53
        }
        for fi54 := 0; fi54 < 6; fi54++ {
            v55 := []int{(fi54 % (((fi54 - (fi54 % (((^(fi54))) | 1)))) | 1)), (^(fi54))}
            var v56 uint64 = 18446744073709551615
            _ = fi54
            _ = v55
            _ = v56
        }
        var v57 uint32 = 266
        _ = v57
    } else {
        v58 := GenS4{F0: uint(v28)}
        _ = v58
    }
    v40 = 2
    for fi59 := 0; fi59 < 19; fi59++ {
        for fi60 := 0; fi60 < 19; fi60++ {
            fmt.Println(fi59, fi60, fi60)
            v28 = (v22 / ((v22) | 1))
            v61 := GenS5{F0: 0, F1: ("foo bar" + ("" + "\n\t")), F2: v2}
            v2 += 73
            _ = fi60
            _ = v61
        }
        fmt.Println(v22, v1, v22)
        for fi62 := 0; fi62 < 18; fi62++ {
            fmt.Println(fi62, v43)
            fmt.Println(v40)
            fmt.Println(v2, v28, v22)
            var v63 uint64 = 608
            if false {
                v64 := GenS5{F0: fi62, F1: "", F2: 67}
                v65 := GenS6{F0: (v63 ^ (v63 * (v63 & (v63 * v63)))), F1: 127}
                v66 := GenS7{F0: (-(v2))}
                _ = v64
                _ = v65
                _ = v66
            } else {
                var v67 float32 = v3
                v3 = (((v3 + v67) + -2.25) / float32(-2.25))
                _ = v67
            }
            v2 = -86
            _ = fi62
            _ = v63
        }
        _ = fi59
    }
    for fi68 := 0; fi68 < 3; fi68++ {
        var v69 int16 = (v43 - (v43 >> ((0) & 63)))
        v70 := []int{(fi68 + (fi68 + (^(fi68))))}
        if (1e10 == v3) {
            for fi71 := 0; fi71 < 3; fi71++ {
                v70 = append(v70, fi71)
                if (42 > (v1 / ((1) | 1))) {
                    _ = len(v70)
                } else {
                    _ = len(v70)
                    var v72 uint64 = 623
                    v73 := GenS7{F0: (v2 + (v2 * (v2 + -46)))}
                    _ = v72
                    _ = v73
                }
                v3 += v3
                _ = fi71
            }
        }
        _ = fi68
        _ = v69
        _ = v70
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v22
    _ = v28
    _ = v3
    _ = v40
    _ = v41
    _ = v43
}
