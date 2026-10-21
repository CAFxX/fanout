package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint64; F1 int32; F2 uint32 }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 int64; F1 uint; F2 uint16 }
func genFn0() int {
    if true {
        v6 := []int{2147483647}
        _ = v6
    }
    switch 523 {
    case 607:
        var v7 int8 = 127
        _ = v7
    case 32767:
        for fi8 := 0; fi8 < 12; fi8++ {
            var v9 uint = 1
            switch (("a" + "a") + "") {
            case "\n\t":
                switch (fi8 / (((-(fi8))) | 1)) {
                case -1:
                    fmt.Println(v9)
                case -1000:
                    v9 = v9
                case 0:
                    v10 := GenS1{F0: 0, F1: 42}
                    _ = v10
                }
            default:
                if true {
                    if (!(!true)) {
                        if false {
                            fmt.Println(v9)
                            var v11 uint8 = 42
                            if false {
                                v11 -= v11
                                var v12 string = "\n\t"
                                v13 := []int{}
                                _ = v12
                                _ = v13
                            }
                            _ = v11
                        }
                        if (true && (2147483647 < -1)) {
                            v14 := GenS2{F0: 443, F1: 2147483647, F2: 628}
                            var v15 uint16 = 2
                            if (true && (-536 >= 32767)) {
                                for fi16 := 0; fi16 < 5; fi16++ {
                                    var v17 int64 = -9223372036854775808
                                    switch "hello" {
                                    case "\n\t":
                                        var v18 uint8 = 42
                                        _ = v18
                                    case "a":
                                        var v19 float32 = 0.1
                                        switch "a" {
                                        case "foo bar":
                                            if (fi8 >= fi8) {
                                                v20 := []int{127}
                                                _ = v20
                                            }
                                        }
                                        v15 = (v15 & 0)
                                        _ = v19
                                    }
                                    _ = fi16
                                    _ = v17
                                }
                                if (0 >= 896) {
                                    v21 := []int{(^(fi8)), 1}
                                    _ = v21
                                } else {
                                    v9 = 65535
                                }
                            }
                            _ = v14
                            _ = v15
                        } else {
                            if (0 > 594) {
                                v9 -= 111
                            }
                        }
                        for fi22 := 0; fi22 < 18; fi22++ {
                            var v23 uint16 = 1
                            if (!(!(!false))) {
                                v9 += 4294967295
                            } else {
                                v24 := GenS1{F0: 42, F1: 18446744073709551615}
                                v9 -= 65535
                                var v25 bool = (-307 < 0)
                                v9 = (v9 >> ((v9) & 63))
                                _ = v24
                                _ = v25
                            }
                            _ = fi22
                            _ = v23
                        }
                    } else {
                        v9 = (v9 + 617)
                        var v26 int = (fi8 % ((fi8) | 1))
                        switch (-(v26)) {
                        case 128:
                            fmt.Println(v9)
                        case 2147483647:
                            fmt.Println(v9, v26)
                        default:
                            var v27 bool = (!(831 >= 266))
                            _ = v27
                        }
                        _ = v26
                    }
                    var v28 bool = false
                    _ = v28
                } else {
                    switch (^(fi8)) {
                    case 32767:
                        v9 -= 18446744073709551615
                    }
                }
            }
            var v29 int32 = 513
            _ = fi8
            _ = v29
            _ = v9
        }
    }
    var v30 int = -1
    if (true || ("a" > ("\n\t" + ""))) {
        if (!true) {
            if (!true) {
                v31 := []int{}
                v30 -= 579
                fmt.Println(v30)
                _ = v31
            }
            if true {
                v32 := []int{v30, v30, v30, v30}
                _ = v32
            } else {
                var v33 int32 = 42
                v33 -= (^(v33))
                v34 := []int{(v30 - v30), (v30 * (v30 / ((v30) | 1)))}
                _ = v33
                _ = v34
            }
        }
    } else {
        for fi35 := 0; fi35 < 11; fi35++ {
            for fi36 := 0; fi36 < 4; fi36++ {
                v37 := []int{32767, (-(v30)), 128}
                v38 := GenS2{F0: 18446744073709551615, F1: -1, F2: 2147483648}
                v39 := GenS1{F0: 1, F1: 18446744073709551615}
                _ = fi36
                _ = v37
                _ = v38
                _ = v39
            }
            switch 2147483647 {
            case 613:
                for fi40 := 0; fi40 < 11; fi40++ {
                    v30 = v30
                    switch fi35 {
                    case 814:
                        v41 := []int{fi40, fi35}
                        _ = v41
                    case 128:
                        v30 = (v30 - 32768)
                        var v42 float64 = (0.1 + (3.14159 / 1.5))
                        _ = v42
                    case -9223372036854775808:
                        v30 = (v30 / ((fi35) | 1))
                    default:
                        v30 -= 128
                    }
                    v30 += (fi40 >> ((4294967295) & 63))
                    _ = fi40
                }
            case 2147483647:
                var v43 uint64 = 18446744073709551615
                _ = v43
            }
            _ = fi35
        }
    }
    _ = v30
    return v30
}
func genFn1() int32 {
    var v44 bool = false
    v44 = v44
    v44 = true
    _ = v44
    return -456
}

func main() {
    var v1 int8 = 1
    var v2 string = "foo bar"
    var v3 string = v2
    v4 := make(map[int]string)
    var v5 bool = (!false)
    var v45 uint8 = 128
    v46 := GenS2{F0: 42, F1: int32(v1), F2: 2}
    fmt.Println(v2)
    v5 = v5
    switch v3 {
    case "hello":
        for fi47 := 0; fi47 < 14; fi47++ {
            for fi48 := 0; fi48 < 14; fi48++ {
                var v49 uint32 = 4294967295
                var v50 uint8 = v45
                var v51 uint16 = 0
                _ = fi48
                _ = v49
                _ = v50
                _ = v51
            }
            v1 += (v1 | (v1 % (((v1 + (^(v1)))) | 1)))
            switch 127 {
            case 32767:
                _ = len(v4)
            case -1:
                v2 = "\n\t"
            }
            _ = fi47
        }
    case "\n\t":
        var v52 uint = 2
        _ = v52
    default:
        var v53 uint = 927
        _ = v53
    }
    var v54 int64 = 1
    var v55 uint16 = 87
    v2 = v3
    for fi56 := 0; fi56 < 7; fi56++ {
        for fi57 := 0; fi57 < 19; fi57++ {
            var v58 int8 = v1
            if ((fi57 % (((fi57 ^ fi57)) | 1)) <= 127) {
                v54 = (v54 - v54)
                v59 := []int{(fi56 % ((genFn0()) | 1)), (-(fi56)), (fi56 | fi57), fi57}
                delete(v4, fi57)
                var v60 float32 = 1e10
                _ = v59
                _ = v60
            } else {
                fmt.Println(v54, v54, v45)
                var v61 int32 = 548
                _ = v61
            }
            _ = fi57
            _ = v58
        }
        if true {
            for fi62 := 0; fi62 < 14; fi62++ {
                if (0.1 < 0.0) {
                    v63 := GenS3{F0: (v45 - 3), F1: v45}
                    _ = v63
                }
                v2 = v3
                _ = fi62
            }
        } else {
            for fi64 := 0; fi64 < 6; fi64++ {
                var v65 uint64 = 2
                var v66 string = ("\n\t" + v3)
                var v67 float32 = 3.14159
                _ = fi64
                _ = v65
                _ = v66
                _ = v67
            }
            var v68 uint32 = 1
            _ = v68
        }
        v69 := GenS3{F0: v45, F1: v45}
        _ = fi56
        _ = v69
    }
    var v70 string = (((v3 + "a") + "a") + "\n\t")
    v71 := GenS4{F0: 0, F1: 2, F2: v55}
    if ((-(v54)) < (v54 * (v54 & v54))) {
        if (350 != 65535) {
            var v72 int = 42
            switch v2 {
            case "hello":
                _ = len(v4)
            case "":
                for fi73 := 0; fi73 < 6; fi73++ {
                    v3 = "hello"
                    var v74 float32 = float32((float64(3.14159) + 1.5))
                    _ = fi73
                    _ = v74
                }
            case "a":
                v75 := v4[(v72 / ((v72) | 1))]
                _ = v75
            }
            v54 -= 848
            _ = v72
        } else {
            v4[0] = ("foo bar" + "foo bar")
            switch v3 {
            case "":
                if true {
                    if v5 {
                        switch -68 {
                        case -1:
                            var v76 int = 1
                            _ = v76
                        }
                    }
                } else {
                    var v77 uint64 = 4294967295
                    var v78 uint32 = 658
                    v79 := []int{-1, -2147483648, genFn0(), -1}
                    _ = v77
                    _ = v78
                    _ = v79
                }
            case "\n\t":
                v45 = 74
            default:
                v80 := GenS3{F0: (v45 | 42), F1: (v45 % (((^(v45))) | 1))}
                _ = v80
            }
        }
        delete(v4, 2147483647)
        switch "hello" {
        case "a":
            if v5 {
                switch v70 {
                case "":
                    var v81 uint32 = 2147483647
                    _ = v81
                case "hello":
                    v55 = (v55 - 42)
                case "a":
                    if v5 {
                        fmt.Println(v45, v45)
                        var v82 int32 = 42
                        _ = v82
                    }
                default:
                    var v83 float32 = 0.1
                    _ = v83
                }
            } else {
                fmt.Println(v70, v1, v5)
                v3 = "\n\t"
                v84 := GenS2{F0: 0, F1: 42, F2: 42}
                _ = v84
            }
        }
    }
    v4[-1] = (v2 + v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v54
    _ = v55
    _ = v70
    _ = v71
}
