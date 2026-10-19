package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 uint64; F1 uint64; F2 string }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 bool; F1 int16 }
func genFn0() int16 {
    v4 := []int{763, -2147483648, 9223372036854775807, -991}
    switch "\n\t" {
    case "foo bar":
        var v5 uint8 = 1
        _ = v5
    case "\n\t":
        var v6 int16 = 477
        _ = v6
    }
    var v7 int = -1
    v7 = (v7 % ((v7) | 1))
    var v8 int16 = -32768
    _ = v4
    _ = v7
    _ = v8
    return v8
}

func main() {
    var v1 int32 = 2147483647
    var v2 uint64 = 750
    var v3 string = "hello"
    switch v3 {
    case "hello":
        v3 = (v3 + "\n\t")
    case "a":
        if true {
            fmt.Println(v2, v1)
            for fi9 := 0; fi9 < 9; fi9++ {
                fmt.Println(v2, fi9, v2)
                v10 := []int{(^(fi9)), 2147483647, fi9}
                var v11 uint = 255
                _ = fi9
                _ = v10
                _ = v11
            }
        }
    case "foo bar":
        v12 := GenS1{F0: 65535, F1: false}
        _ = v12
    }
    fmt.Println(v2)
    for fi13 := 0; fi13 < 13; fi13++ {
        v3 = "a"
        var v14 int64 = -904
        _ = fi13
        _ = v14
    }
    var v15 int8 = 42
    var v16 int64 = -1
    var v17 uint32 = 2
    switch "\n\t" {
    case "hello":
        v17 = v17
    case "":
        v18 := GenS1{F0: 42, F1: (42 == 1)}
        _ = v18
    default:
        var v19 uint = 2
        _ = v19
    }
    if false {
        switch v3 {
        case "a":
            for fi20 := 0; fi20 < 3; fi20++ {
                v16 = int64(fi20)
                _ = fi20
            }
        }
        var v21 uint16 = 342
        _ = v21
    } else {
        if false {
            for fi22 := 0; fi22 < 4; fi22++ {
                if false {
                    var v23 uint8 = 1
                    _ = v23
                }
                v1 -= v1
                _ = fi22
            }
            switch (v3 + ((v3 + v3) + ("a" + v3))) {
            case "":
                v1 -= int32(v17)
            default:
                if (false && true) {
                    if true {
                        v3 = v3
                        v15 += (^(v15))
                        for fi24 := 0; fi24 < 13; fi24++ {
                            var v25 uint16 = 150
                            _ = fi24
                            _ = v25
                        }
                    } else {
                        if true {
                            v26 := GenS1{F0: 42, F1: true}
                            if false {
                                fmt.Println(v3)
                                v27 := GenS2{F0: 100, F1: 4294967295}
                                _ = v27
                            }
                            for fi28 := 0; fi28 < 5; fi28++ {
                                v29 := []int{fi28, (fi28 << ((1) & 63)), fi28, (fi28 >> ((469) & 63))}
                                _ = fi28
                                _ = v29
                            }
                            _ = v26
                        } else {
                            var v30 uint8 = 200
                            var v31 int8 = 100
                            v32 := []int{-826, 42, 343, -9223372036854775808}
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                        var v33 uint8 = 127
                        _ = v33
                    }
                } else {
                    v34 := []int{4294967295, -707, 127, 32767}
                    v3 = v3
                    _ = v34
                }
            }
            var v35 uint8 = 255
            _ = v35
        }
        v36 := GenS2{F0: 0, F1: v17}
        _ = v36
    }
    for fi37 := 0; fi37 < 5; fi37++ {
        var v38 uint8 = 1
        _ = fi37
        _ = v38
    }
    for fi39 := 0; fi39 < 3; fi39++ {
        switch (fi39 << ((65535) & 63)) {
        case 4294967295:
            if (1 <= 421) {
                if true {
                    v17 = 2
                    fmt.Println(v3, v17, v15)
                    v3 = (v3 + v3)
                }
                v17 += (v17 >> ((1) & 63))
                v40 := []int{int(v17), fi39, fi39}
                _ = v40
            }
        }
        v41 := []int{int(v2), fi39, fi39}
        _ = fi39
        _ = v41
    }
    for fi42 := 0; fi42 < 7; fi42++ {
        for fi43 := 0; fi43 < 15; fi43++ {
            v3 = ((("a" + ("a" + v3)) + "\n\t") + v3)
            var v44 int8 = int8(fi43)
            v16 = (v16 >> ((42) & 63))
            _ = fi43
            _ = v44
        }
        _ = fi42
    }
    fmt.Println(v1, v15, v16)
    for fi45 := 0; fi45 < 17; fi45++ {
        v46 := GenS3{F0: (v2 & (v2 * (v2 | (v2 + (v2 / ((1) | 1)))))), F1: 42, F2: v3}
        if (!true) {
            if true {
                v47 := GenS1{F0: 42, F1: true}
                var v48 int32 = 1
                _ = v47
                _ = v48
            } else {
                v16 -= (v16 | v16)
            }
        } else {
            if (float64(1e10) < ((1e10 + 3.14159) - 0.1)) {
                if ((1 > v16) || false) {
                    if (false && true) {
                        if false {
                            v2 = 0
                        } else {
                            v49 := GenS4{F0: v17, F1: -688}
                            v50 := []int{fi45, (fi45 / ((fi45) | 1)), fi45, (fi45 ^ 0)}
                            if false {
                                switch 1 {
                                case -1:
                                    for fi51 := 0; fi51 < 2; fi51++ {
                                        switch (fi51 + fi51) {
                                        case -740:
                                            var v52 bool = true
                                            v1 -= v1
                                            _ = v52
                                        case 127:
                                            v50 = append(v50, -2147483648)
                                        case 4294967295:
                                            switch (fi51 | fi51) {
                                            case 874:
                                                switch "\n\t" {
                                                case "":
                                                    var v53 string = v3
                                                    _ = v53
                                                case "hello":
                                                    var v54 float64 = 0.1
                                                    _ = v54
                                                case "\n\t":
                                                    _ = len(v50)
                                                }
                                            case 501:
                                                var v55 uint16 = 1
                                                _ = v55
                                            case -978:
                                                _ = len(v50)
                                            default:
                                                v17 = v17
                                                if false {
                                                    var v56 bool = (!false)
                                                    v3 = (v3 + "a")
                                                    _ = v56
                                                } else {
                                                    fmt.Println(fi51)
                                                    if (!false) {
                                                        if len(v50) > 0 {
                                                            v50[0] = fi45
                                                        }
                                                    } else {
                                                        var v57 int8 = (v15 % (((v15 ^ v15)) | 1))
                                                        _ = v57
                                                    }
                                                    _ = len(v50)
                                                }
                                            }
                                            if len(v50) > 0 {
                                                v50[0] = (fi51 >> ((1) & 63))
                                            }
                                            var v58 uint16 = 65535
                                            _ = v58
                                        default:
                                            switch ("hello" + "a") {
                                            case "foo bar":
                                                v59 := GenS1{F0: 42, F1: (true || (-2.25 > 0.0))}
                                                _ = v59
                                            case "hello":
                                                fmt.Println(v2, fi45, v3)
                                            }
                                        }
                                        var v60 uint64 = 42
                                        v61 := GenS1{F0: 875, F1: ((^(v17)) <= (v17 >> ((255) & 63)))}
                                        _ = fi51
                                        _ = v60
                                        _ = v61
                                    }
                                case 1:
                                    v15 = v15
                                default:
                                    var v62 uint64 = v2
                                    _ = v62
                                }
                                v17 = v17
                            }
                            _ = v49
                            _ = v50
                        }
                    } else {
                        v63 := GenS2{F0: 100, F1: v17}
                        _ = v63
                    }
                    if (((v1 % (((v1 * int32(v15))) | 1)) != v1) || (!false)) {
                        fmt.Println(fi45, v16)
                        v64 := GenS5{F0: v2}
                        _ = v64
                    } else {
                        if (false || false) {
                            fmt.Println(v2, v16, v3)
                            var v65 int16 = -1
                            _ = v65
                        } else {
                            var v66 int8 = 1
                            var v67 uint32 = v17
                            _ = v66
                            _ = v67
                        }
                        v68 := []int{(-(fi45)), (^(fi45))}
                        var v69 bool = true
                        _ = v68
                        _ = v69
                    }
                    v70 := []int{fi45, 32767, (fi45 | fi45)}
                    _ = v70
                } else {
                    v15 -= (v15 * v15)
                    for fi71 := 0; fi71 < 18; fi71++ {
                        var v72 uint32 = v17
                        _ = fi71
                        _ = v72
                    }
                    var v73 uint16 = 0
                    _ = v73
                }
                switch -885 {
                case -337:
                    v74 := []int{fi45}
                    _ = v74
                case 4294967295:
                    v75 := GenS6{F0: (!(v3 != v3)), F1: 1}
                    _ = v75
                }
            } else {
                for fi76 := 0; fi76 < 19; fi76++ {
                    var v77 uint64 = 42
                    _ = fi76
                    _ = v77
                }
                v78 := []int{(fi45 / (((-(fi45))) | 1)), int(v15), fi45, fi45}
                v2 = (^(v2))
                _ = v78
            }
            var v79 float32 = (0.0 / 0.1)
            v80 := []int{fi45, fi45}
            _ = v79
            _ = v80
        }
        for fi81 := 0; fi81 < 18; fi81++ {
            var v82 bool = (!((v2 | (v2 << ((662) & 63))) < 417))
            var v83 uint16 = 32767
            _ = fi81
            _ = v82
            _ = v83
        }
        _ = fi45
        _ = v46
    }
    v84 := []int{}
    v15 = v15
    if ((2 < 379) || true) {
        v85 := GenS5{F0: 4294967295}
        _ = v85
    } else {
        for fi86 := 0; fi86 < 15; fi86++ {
            v3 = v3
            v87 := GenS1{F0: 1, F1: true}
            var v88 int8 = (v15 % (((v15 - v15)) | 1))
            _ = fi86
            _ = v87
            _ = v88
        }
        var v89 bool = true
        _ = v89
    }
    if true {
        v90 := GenS5{F0: 42}
        for fi91 := 0; fi91 < 1; fi91++ {
            switch v3 {
            case "\n\t":
                switch "" {
                case "hello":
                    if len(v84) >= 2 {
                        v84 = v84[1:len(v84)]
                    }
                default:
                    v2 -= (v2 + (v2 - v2))
                }
            }
            _ = fi91
        }
        var v92 bool = ((true && (v15 != 127)) || (1.5 != 1.5))
        _ = v90
        _ = v92
    } else {
        switch "\n\t" {
        case "foo bar":
            var v93 uint = 801
            _ = v93
        case "":
            v3 = "\n\t"
        case "\n\t":
            var v94 string = "\n\t"
            _ = v94
        }
        var v95 uint = 0
        _ = v95
    }
    var v96 int32 = (v1 << ((930) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v84:", len(v84), v84)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v84
    _ = v96
}
