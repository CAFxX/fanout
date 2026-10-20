package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 string; F1 string }
type GenS7 struct { F0 uint16; F1 uint; F2 string }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 bool; F1 int16 }
type GenS10 struct { F0 uint8; F1 int8 }
type GenS11 struct { F0 uint; F1 int16 }
type GenS12 struct { F0 int64; F1 int8 }
func genFn0(p0 bool, p1 uint) uint32 {
    v6 := GenS1{F0: 42, F1: true, F2: 255}
    p1 = (^(p1))
    _ = p0
    _ = p1
    _ = v6
    return 694
}
func genFn1() uint32 {
    switch 0 {
    case 0:
        v7 := []int{2147483647, 4294967295, 32768}
        _ = v7
    case 42:
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 uint64 = 1
            if false {
                var v10 string = ("\n\t" + "")
                v11 := GenS2{F0: -1}
                if true {
                    var v12 int8 = 1
                    _ = v12
                } else {
                    var v13 int64 = 9223372036854775807
                    _ = v13
                }
                _ = v10
                _ = v11
            } else {
                var v14 uint8 = 79
                if false {
                    switch "" {
                    case "hello":
                        for fi15 := 0; fi15 < 17; fi15++ {
                            v16 := []int{}
                            switch "\n\t" {
                            case "\n\t":
                                if false {
                                    v17 := GenS3{F0: 2147483647, F1: 647}
                                    _ = v17
                                }
                            default:
                                fmt.Println(fi8)
                            }
                            var v18 uint32 = 4294967295
                            _ = fi15
                            _ = v16
                            _ = v18
                        }
                    case "a":
                        if false {
                            switch "foo bar" {
                            case "\n\t":
                                v9 = 18446744073709551615
                            case "a":
                                v14 = 2
                            }
                        }
                    case "foo bar":
                        v19 := []int{}
                        _ = v19
                    default:
                        v20 := []int{(fi8 << ((0) & 63)), (fi8 >> ((893) & 63))}
                        _ = v20
                    }
                    v21 := GenS1{F0: genFn0((124 < -1), 65535), F1: true, F2: v14}
                    var v22 uint = 1
                    _ = v21
                    _ = v22
                }
                _ = v14
            }
            var v23 uint64 = (v9 | (^(v9)))
            _ = fi8
            _ = v23
            _ = v9
        }
    default:
        switch 2147483647 {
        case 9223372036854775807:
            var v24 int16 = 1000
            _ = v24
        default:
            var v25 int64 = 42
            _ = v25
        }
    }
    var v26 float32 = 1.5
    for fi27 := 0; fi27 < 4; fi27++ {
        for fi28 := 0; fi28 < 8; fi28++ {
            var v29 int16 = 366
            _ = fi28
            _ = v29
        }
        _ = fi27
    }
    for fi30 := 0; fi30 < 5; fi30++ {
        if true {
            var v31 uint32 = 2147483648
            fmt.Println(v31, v26, fi30)
            _ = v31
        } else {
            v26 -= 0.1
            var v32 int8 = -35
            if false {
                v33 := GenS1{F0: genFn0((127 == 128), 826), F1: ("foo bar" == (("foo bar" + "foo bar") + "\n\t")), F2: 0}
                v34 := []int{fi30, (fi30 & (fi30 << ((1) & 63))), fi30, 32768}
                v35 := GenS4{F0: -32768}
                _ = v33
                _ = v34
                _ = v35
            }
            _ = v32
        }
        for fi36 := 0; fi36 < 17; fi36++ {
            v26 += 0.0
            if (!true) {
                var v37 int = (fi30 / ((fi30) | 1))
                _ = v37
            }
            _ = fi36
        }
        _ = fi30
    }
    v38 := GenS2{F0: 127}
    _ = v26
    _ = v38
    return 647
}

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 96
    v3 := make(map[int]uint32)
    v4 := []uint32{1, (v2 >> ((0) & 63)), v2}
    v5 := []string{}
    v2 = (v2 * 0)
    v3[2147483647] = (^(v2))
    for fi39 := 0; fi39 < 15; fi39++ {
        if true {
            v2 = v2
        } else {
            v2 -= (v2 / ((genFn0(true, 18446744073709551615)) | 1))
            v40 := GenS5{F0: 127}
            var v41 uint = 995
            _ = v40
            _ = v41
        }
        _ = fi39
    }
    v3[4294967295] = v2
    var v42 uint64 = 687
    var v43 float64 = 1e10
    var v44 uint8 = 243
    var v45 string = "hello"
    for fi46 := 0; fi46 < 17; fi46++ {
        for fi47 := 0; fi47 < 13; fi47++ {
            if false {
                if (true || (v45 == v45)) {
                    switch (^(fi47)) {
                    case 0:
                        if false {
                            fmt.Println(v44, v43)
                        }
                        if ((189 != 951) && (v43 <= (v43 / v43))) {
                            v48 := GenS6{F0: (v45 + v45), F1: v45}
                            _ = v48
                        } else {
                            v49 := GenS1{F0: v2, F1: (((v43 - ((1e10 / (v43 - v43)) - ((((1e10 + v43) + ((v43 * (3.14159 / v43)) + (v43 * float64(1.5)))) * 0.0) / float64((-2.25 / 3.14159))))) * v43) < ((0.0 + (-2.25 - -2.25)) + (v43 * ((((-2.25 * ((3.14159 + float64(float32(v43))) - v43)) + 1e10) / 3.14159) - (v43 * v43))))), F2: 203}
                            var v50 string = "foo bar"
                            _ = v49
                            _ = v50
                        }
                        var v51 int32 = 0
                        if false {
                            if false {
                                _ = len(v5)
                            } else {
                                var v52 uint16 = 220
                                _ = v52
                            }
                        } else {
                            v53 := GenS7{F0: 42, F1: 42, F2: ("" + v45)}
                            v54 := GenS8{F0: 2}
                            v44 = 0
                            _ = v53
                            _ = v54
                        }
                        switch (v45 + (("foo bar" + "a") + v45)) {
                        case "":
                            switch fi46 {
                            case 0:
                                v55 := GenS9{F0: (v44 == 118), F1: 120}
                                _ = v55
                            }
                        default:
                            v42 -= (v42 * uint64(v2))
                        }
                        _ = v51
                    default:
                        v56 := GenS10{F0: 0, F1: -1}
                        _ = v56
                    }
                    if false {
                        v45 = v45
                    }
                    fmt.Println(v42, v44, v2)
                } else {
                    v57 := GenS4{F0: 0}
                    switch fi47 {
                    case 1:
                        var v58 int8 = 0
                        _ = v58
                    default:
                        fmt.Println(v44)
                    }
                    switch (fi46 & 127) {
                    case -727:
                        if len(v4) > 0 {
                            v4[0] = v2
                        }
                    case 0:
                        switch v45 {
                        case "\n\t":
                            _ = len(v4)
                        }
                    }
                    switch "\n\t" {
                    case "hello":
                        fmt.Println(v42, v44, v44)
                    case "foo bar":
                        var v59 float64 = 1.5
                        _ = v59
                    default:
                        var v60 float64 = v43
                        _ = v60
                    }
                    _ = v57
                }
                v61 := GenS4{F0: 32767}
                var v62 bool = true
                _ = v61
                _ = v62
            }
            fmt.Println(v44, v43, fi47)
            v63 := v3[(-(fi46))]
            _ = fi47
            _ = v63
        }
        switch fi46 {
        case 1:
            switch v45 {
            case "a":
                if ((fi46 << ((2) & 63)) < (fi46 & 2147483647)) {
                    v64 := GenS3{F0: 211, F1: 255}
                    switch "hello" {
                    case "hello":
                        var v65 uint32 = genFn0(true, 478)
                        _ = v65
                    }
                    _ = v64
                } else {
                    v44 = (^(v44))
                    v44 += v44
                }
            case "\n\t":
                for fi66 := 0; fi66 < 2; fi66++ {
                    delete(v1, "")
                    _ = fi66
                }
            default:
                fmt.Println(v42)
            }
        default:
            var v67 int32 = 867
            _ = v67
        }
        for fi68 := 0; fi68 < 19; fi68++ {
            v69 := GenS8{F0: 501}
            _ = fi68
            _ = v69
        }
        _ = fi46
    }
    var v70 int = 271
    v42 = v42
    switch v45 {
    case "a":
        if false {
            var v71 int = (v70 >> ((404) & 63))
            _ = v71
        } else {
            if ((871 <= 11) || false) {
                delete(v1, "foo bar")
                if (("foo bar" <= v45) || false) {
                    if true {
                        for fi72 := 0; fi72 < 8; fi72++ {
                            v44 = v44
                            var v73 string = v45
                            _ = fi72
                            _ = v73
                        }
                    }
                    var v74 uint64 = (v42 * 1)
                    v44 = v44
                    _ = v74
                } else {
                    if len(v5) > 0 {
                        v5[0] = "\n\t"
                    }
                }
                v1[((v45 + "a") + "\n\t")] = v70
            } else {
                switch v45 {
                case "\n\t":
                    fmt.Println(v44)
                case "":
                    if len(v5) > 0 {
                        v5[0] = "a"
                    }
                }
                var v75 uint = uint(v42)
                if true {
                    switch (v70 >> (((v75 | (v75 - (v75 * (v75 % (((v75 * uint(v42))) | 1)))))) & 63)) {
                    case 2147483647:
                        v75 = v75
                    case 1:
                        if (false || (false && ((false || (-1 < 100)) && false))) {
                            v76 := v1[v45]
                            _ = v76
                        }
                    case 32767:
                        for fi77 := 0; fi77 < 17; fi77++ {
                            v78 := v1[v45]
                            _ = fi77
                            _ = v78
                        }
                    }
                    switch (v70 ^ (v70 | (v70 - v70))) {
                    case 772:
                        var v79 uint = uint(v2)
                        _ = v79
                    case -9223372036854775808:
                        v45 = "foo bar"
                    case -2147483648:
                        switch v70 {
                        case -451:
                            _ = len(v1)
                        case 0:
                            var v80 uint32 = v2
                            _ = v80
                        case -2147483648:
                            v81 := GenS11{F0: v75, F1: 608}
                            _ = v81
                        }
                    default:
                        v82 := GenS12{F0: -239, F1: -37}
                        _ = v82
                    }
                }
                _ = v75
            }
        }
    }
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v70
}
