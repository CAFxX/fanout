package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 int8; F1 uint; F2 bool }
type GenS6 struct { F0 int32; F1 uint64; F2 int64 }
type GenS7 struct { F0 uint32; F1 uint32 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int32; F1 bool; F2 int64 }
type GenS10 struct { F0 bool; F1 int32; F2 int64 }
type GenS11 struct { F0 uint16; F1 int; F2 uint8 }
func genFn0(p0 uint16, p1 int) int64 {
    v4 := GenS1{F0: "", F1: 32767}
    if false {
        p1 = -1
    }
    _ = p0
    _ = p1
    _ = v4
    return -9223372036854775808
}
func genFn1() uint32 {
    var v5 int = -2147483648
    for fi6 := 0; fi6 < 13; fi6++ {
        if (3.14159 == -2.25) {
            var v7 int = (fi6 | (v5 & (v5 | (v5 | (v5 >> ((500) & 63))))))
            _ = v7
        }
        var v8 int64 = -60
        if (false && true) {
            switch fi6 {
            case -506:
                fmt.Println(fi6, v8, v5)
            case -1:
                switch (v5 + v5) {
                case -215:
                    var v9 int16 = 42
                    _ = v9
                }
            case 343:
                if false {
                    var v10 int64 = genFn0(216, (fi6 << ((280) & 63)))
                    switch (v5 << ((557) & 63)) {
                    case 377:
                        for fi11 := 0; fi11 < 5; fi11++ {
                            var v12 float32 = 0.0
                            v8 = -586
                            v8 -= (v8 | v8)
                            _ = fi11
                            _ = v12
                        }
                    case -9223372036854775808:
                        if true {
                            switch v5 {
                            case -2147483648:
                                v13 := []int{}
                                _ = v13
                            case -856:
                                v5 = fi6
                            }
                            switch fi6 {
                            case 32767:
                                v14 := GenS2{F0: 78, F1: 69}
                                _ = v14
                            case -843:
                                v10 += genFn0(2, fi6)
                            case -1:
                                for fi15 := 0; fi15 < 8; fi15++ {
                                    fmt.Println(v8)
                                    var v16 uint = 65535
                                    if false {
                                        if false {
                                            switch fi15 {
                                            case 127:
                                                v17 := []int{v5, 32768, v5, (fi15 & fi15)}
                                                _ = v17
                                            default:
                                                var v18 int = fi15
                                                _ = v18
                                            }
                                            var v19 float64 = 3.14159
                                            v19 = v19
                                            _ = v19
                                        }
                                        if true {
                                            var v20 int8 = 127
                                            if true {
                                                var v21 int = (-(fi15))
                                                _ = v21
                                            }
                                            fmt.Println(v16, v20)
                                            if true {
                                                v5 = (-(v5))
                                            }
                                            _ = v20
                                        }
                                        v10 += genFn0(2, 1)
                                    }
                                    _ = fi15
                                    _ = v16
                                }
                            default:
                                for fi22 := 0; fi22 < 2; fi22++ {
                                    v23 := []int{}
                                    if false {
                                        var v24 float32 = 0.0
                                        v10 = v8
                                        fmt.Println(v10)
                                        var v25 uint8 = 2
                                        switch "" {
                                        case "foo bar":
                                            v26 := GenS3{F0: v25, F1: "a"}
                                            _ = v26
                                        case "\n\t":
                                            if (!(false || false)) {
                                                var v27 uint = 1
                                                var v28 string = "\n\t"
                                                if false {
                                                    if true {
                                                        v27 -= (v27 | v27)
                                                        var v29 int8 = 42
                                                        v30 := GenS4{F0: (fi6 & 85), F1: 0}
                                                        _ = v29
                                                        _ = v30
                                                    }
                                                }
                                                var v31 float32 = ((v24 * v24) - 1e10)
                                                _ = v27
                                                _ = v28
                                                _ = v31
                                            }
                                        default:
                                            v32 := GenS5{F0: 127, F1: 42, F2: true}
                                            _ = v32
                                        }
                                        _ = v24
                                        _ = v25
                                    } else {
                                        if true {
                                            v33 := GenS6{F0: 42, F1: 2, F2: v10}
                                            var v34 uint32 = 4294967295
                                            _ = v33
                                            _ = v34
                                        }
                                    }
                                    if false {
                                        switch (fi6 + int(v10)) {
                                        case 128:
                                            if true {
                                                var v35 int16 = 42
                                                var v36 int8 = 100
                                                _ = v35
                                                _ = v36
                                            } else {
                                                v37 := GenS3{F0: 2, F1: "a"}
                                                _ = v37
                                            }
                                        case 32767:
                                            var v38 int64 = v10
                                            _ = v38
                                        default:
                                            v39 := GenS7{F0: 4294967295, F1: 65535}
                                            v10 = genFn0(42, v5)
                                            v10 -= genFn0(1, (fi6 << ((255) & 63)))
                                            _ = v39
                                        }
                                        v40 := GenS1{F0: "\n\t", F1: -454}
                                        _ = len(v23)
                                        _ = v40
                                    } else {
                                        v41 := GenS2{F0: 2, F1: 42}
                                        v10 -= (v8 >> ((65535) & 63))
                                        var v42 string = ("a" + "")
                                        _ = v41
                                        _ = v42
                                    }
                                    _ = fi22
                                    _ = v23
                                }
                            }
                        } else {
                            v43 := GenS8{F0: 1}
                            _ = v43
                        }
                    }
                    _ = v10
                } else {
                    for fi44 := 0; fi44 < 7; fi44++ {
                        if (false && ((-(fi44)) == fi6)) {
                            v45 := []int{v5, v5, fi44, 1}
                            _ = v45
                        }
                        var v46 uint8 = 127
                        var v47 uint64 = 1
                        _ = fi44
                        _ = v46
                        _ = v47
                    }
                }
            default:
                if (!false) {
                    switch v5 {
                    case 32767:
                        v48 := []int{(fi6 >> ((255) & 63)), (fi6 & v5)}
                        _ = v48
                    case 4294967295:
                        for fi49 := 0; fi49 < 13; fi49++ {
                            v50 := GenS9{F0: -616, F1: false, F2: -1}
                            var v51 uint = 1
                            var v52 string = "foo bar"
                            _ = fi49
                            _ = v50
                            _ = v51
                            _ = v52
                        }
                    default:
                        switch (-(fi6)) {
                        case -2147483648:
                            var v53 int32 = -1
                            _ = v53
                        }
                    }
                }
            }
            v54 := []int{}
            _ = v54
        } else {
            if true {
                if ((42 <= -1) || false) {
                    v5 = (fi6 | (v5 >> ((42) & 63)))
                    if false {
                        v8 += (v8 & 9223372036854775807)
                        var v55 uint8 = 0
                        _ = v55
                    }
                    v56 := []int{128, 823, (v5 >> ((413) & 63)), (fi6 * (-(fi6)))}
                    _ = v56
                } else {
                    var v57 int16 = 0
                    v57 += v57
                    _ = v57
                }
            }
            v5 -= -9223372036854775808
        }
        _ = fi6
        _ = v8
    }
    for fi58 := 0; fi58 < 19; fi58++ {
        var v59 int16 = 42
        var v60 int8 = 0
        v61 := []int{fi58, 32768, (v5 ^ 0), (v5 * v5)}
        _ = fi58
        _ = v59
        _ = v60
        _ = v61
    }
    v5 = 826
    if ((32768 <= (v5 % ((v5) | 1))) || false) {
        var v62 float64 = (1e10 * -2.25)
        v5 += (^(v5))
        _ = v62
    }
    _ = v5
    return 0
}

func main() {
    var v1 uint32 = 2147483647
    var v2 int8 = 1
    var v3 int16 = -553
    v63 := GenS5{F0: (v2 * (v2 ^ 127)), F1: 65535, F2: (!false)}
    v3 = int16(v2)
    var v64 bool = (false && (!true))
    fmt.Println(v2, v1)
    var v65 int = 0
    var v66 int = (v65 << ((42) & 63))
    var v67 int64 = 1
    switch v66 {
    case 268:
        v65 += 128
    case -2147483648:
        for fi68 := 0; fi68 < 5; fi68++ {
            var v69 uint32 = 65535
            fmt.Println(v67, v3, v64)
            var v70 uint16 = 700
            _ = fi68
            _ = v69
            _ = v70
        }
    case 127:
        if false {
            var v71 uint64 = 4294967295
            v72 := []int{}
            _ = v71
            _ = v72
        } else {
            v3 += v3
            var v73 float32 = 3.14159
            v74 := GenS10{F0: false, F1: 2147483647, F2: -1}
            _ = v73
            _ = v74
        }
    }
    var v75 uint16 = 174
    var v76 uint64 = 18446744073709551615
    v77 := GenS1{F0: "foo bar", F1: v3}
    v64 = v64
    for fi78 := 0; fi78 < 19; fi78++ {
        v76 -= 4294967295
        v79 := GenS11{F0: 638, F1: (v65 << ((2) & 63)), F2: 2}
        v2 = (v2 & -44)
        _ = fi78
        _ = v79
    }
    fmt.Println(v65)
    if (v64 || (!(true && v64))) {
        v80 := []int{(-(v65))}
        _ = v80
    } else {
        if (!true) {
            var v81 int16 = v3
            var v82 uint8 = 0
            var v83 float32 = 1.5
            _ = v81
            _ = v82
            _ = v83
        }
        v84 := []int{v66, v65}
        _ = v84
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v2
    _ = v3
    _ = v63
    _ = v64
    _ = v65
    _ = v66
    _ = v67
    _ = v75
    _ = v76
    _ = v77
}
