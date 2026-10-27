package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32; F1 bool; F2 uint16 }
type GenS6 struct { F0 uint64; F1 string }
type GenS7 struct { F0 uint64; F1 int32; F2 uint8 }
type GenS8 struct { F0 int64; F1 int8; F2 uint }
type GenS9 struct { F0 uint32; F1 int32; F2 uint }
type GenS10 struct { F0 uint; F1 int16; F2 int64 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 uint16 }
func genFn0() uint {
    var v5 uint32 = 2147483648
    var v6 uint16 = 1
    _ = v5
    _ = v6
    return 987
}

func main() {
    var v1 int8 = 78
    var v2 int16 = 1000
    var v3 string = (("a" + ("" + ("foo bar" + (("foo bar" + "hello") + ("hello" + "foo bar"))))) + "foo bar")
    var v4 float64 = (0.1 / (3.14159 + (3.14159 + 0.1)))
    v2 -= 76
    if false {
        var v7 int16 = -27
        _ = v7
    } else {
        v8 := []int{-2147483648, -1, 4294967295}
        for fi9 := 0; fi9 < 7; fi9++ {
            v10 := GenS1{F0: 1}
            _ = fi9
            _ = v10
        }
        v2 = (v2 ^ -909)
        _ = v8
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        for fi12 := 0; fi12 < 20; fi12++ {
            if true {
                v13 := []int{fi12, fi11, fi12, fi12}
                fmt.Println(fi11, v1)
                _ = v13
            }
            _ = fi12
        }
        var v14 bool = (!true)
        _ = fi11
        _ = v14
    }
    v15 := GenS1{F0: 42}
    v1 -= (^(v1))
    for fi16 := 0; fi16 < 19; fi16++ {
        for fi17 := 0; fi17 < 14; fi17++ {
            if false {
                switch ("foo bar" + (v3 + "")) {
                case "hello":
                    if (4294967295 == 0) {
                        fmt.Println(v1)
                        v18 := []int{42, fi17, fi17, fi17}
                        v19 := GenS1{F0: 2}
                        v20 := GenS1{F0: 1}
                        var v21 bool = false
                        _ = v18
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                case "":
                    if false {
                        fmt.Println(fi17)
                        v4 = (v4 + v4)
                        v22 := GenS1{F0: 11}
                        _ = v22
                    }
                default:
                    v3 = v3
                    fmt.Println(fi17, v4)
                    v23 := []int{fi16, fi16}
                    _ = v23
                }
                fmt.Println(fi17, v2)
                v1 -= v1
            }
            v24 := []int{fi16, (fi17 % ((42) | 1))}
            v25 := GenS1{F0: genFn0()}
            _ = fi17
            _ = v24
            _ = v25
        }
        _ = fi16
    }
    var v26 uint = 1
    v27 := GenS1{F0: v26}
    var v28 int8 = -99
    if ((((1e10 + 1e10) + ((v4 * v4) / v4)) * 0.0) <= v4) {
        if (!(1 >= 0)) {
            for fi29 := 0; fi29 < 5; fi29++ {
                switch v3 {
                case "a":
                    for fi30 := 0; fi30 < 2; fi30++ {
                        switch -774 {
                        case 698:
                            v26 += 18446744073709551615
                        case 490:
                            var v31 uint32 = 65535
                            _ = v31
                        default:
                            v32 := []int{fi29, (fi29 >> ((v26) & 63)), (fi30 << (((v26 - (v26 * v26))) & 63))}
                            _ = v32
                        }
                        var v33 int32 = -2147483648
                        var v34 uint32 = 4294967295
                        _ = fi30
                        _ = v33
                        _ = v34
                    }
                case "foo bar":
                    if ((v28 == int8(fi29)) || (false || true)) {
                        for fi35 := 0; fi35 < 11; fi35++ {
                            var v36 uint = (^(v26))
                            if false {
                                v37 := []int{fi35, (fi35 % ((32767) | 1)), (-(fi29))}
                                _ = v37
                            }
                            fmt.Println(v28)
                            if (true || (42 == -1)) {
                                if true {
                                    fmt.Println(v1, v28, v4)
                                    if false {
                                        v38 := GenS1{F0: 777}
                                        _ = v38
                                    } else {
                                        v28 = v1
                                    }
                                }
                                var v39 int64 = 0
                                if (!true) {
                                    v40 := GenS1{F0: (^(v36))}
                                    v41 := GenS2{F0: int16(v39)}
                                    v42 := []int{(fi29 & (fi35 & fi35)), (-(fi35))}
                                    _ = v40
                                    _ = v41
                                    _ = v42
                                } else {
                                    var v43 int = fi35
                                    _ = v43
                                }
                                if true {
                                    if (-561 <= 1) {
                                        if true {
                                            var v44 int16 = (^(v2))
                                            switch fi29 {
                                            case 1:
                                                switch ("" + (v3 + "hello")) {
                                                case "hello":
                                                    v45 := []int{fi35, (-(fi35))}
                                                    _ = v45
                                                case "":
                                                    var v46 uint = 65535
                                                    _ = v46
                                                case "foo bar":
                                                    if true {
                                                        v1 = v1
                                                    } else {
                                                        v26 = 705
                                                    }
                                                default:
                                                    if false {
                                                        v1 = v1
                                                        var v47 int64 = v39
                                                        var v48 uint64 = 18446744073709551615
                                                        v3 = ((("" + v3) + v3) + "a")
                                                        _ = v47
                                                        _ = v48
                                                    } else {
                                                        fmt.Println(v2, v2)
                                                        v49 := GenS2{F0: (v2 - v44)}
                                                        if ((42 >= 617) || true) {
                                                            fmt.Println(v2)
                                                        }
                                                        _ = v49
                                                    }
                                                    if false {
                                                        v50 := []int{(fi29 + 32768), -1, -1}
                                                        _ = v50
                                                    } else {
                                                        var v51 int = (fi29 * fi35)
                                                        _ = v51
                                                    }
                                                }
                                            case -148:
                                                var v52 int16 = (v2 >> (((v26 & (v36 >> ((728) & 63)))) & 63))
                                                switch int(v39) {
                                                case 119:
                                                    v53 := GenS3{F0: v39}
                                                    _ = v53
                                                case -894:
                                                    v39 = -9223372036854775808
                                                }
                                                _ = v52
                                            default:
                                                v54 := []int{}
                                                _ = v54
                                            }
                                            v55 := GenS4{F0: -67}
                                            _ = v44
                                            _ = v55
                                        } else {
                                            var v56 int32 = 0
                                            if (!(571 <= v56)) {
                                                v1 = (v28 / ((v28) | 1))
                                                switch (v3 + ("foo bar" + "\n\t")) {
                                                case "a":
                                                    switch ("a" + v3) {
                                                    case "":
                                                        v2 = (v2 | (v2 / ((1) | 1)))
                                                    case "hello":
                                                        v28 = (v28 & v28)
                                                    }
                                                case "":
                                                    var v57 int16 = (v2 % (((v2 * v2)) | 1))
                                                    _ = v57
                                                case "hello":
                                                    v58 := []int{-9223372036854775808}
                                                    v1 = (v1 % (((v28 >> ((v36) & 63))) | 1))
                                                    _ = v58
                                                }
                                                var v59 uint8 = 69
                                                var v60 int = (-(fi29))
                                                if true {
                                                    switch v3 {
                                                    case "foo bar":
                                                        v61 := GenS5{F0: v56, F1: true, F2: 42}
                                                        _ = v61
                                                    }
                                                    v26 -= (^(v36))
                                                    v62 := GenS1{F0: (v36 & (v26 | (v26 + 4294967295)))}
                                                    _ = v62
                                                } else {
                                                    v63 := GenS2{F0: 0}
                                                    if false {
                                                        v39 -= (v39 % ((v39) | 1))
                                                    }
                                                    var v64 float32 = 1e10
                                                    _ = v63
                                                    _ = v64
                                                }
                                                v65 := GenS6{F0: 2, F1: "hello"}
                                                _ = v59
                                                _ = v60
                                                _ = v65
                                            }
                                            _ = v56
                                        }
                                    } else {
                                        v1 += (^(v28))
                                        if ((!(v36 > v26)) || (!true)) {
                                            var v66 int = fi29
                                            var v67 uint16 = 0
                                            if ((false || false) || false) {
                                                v68 := []int{fi29, (^(fi35)), -237}
                                                var v69 float64 = 0.1
                                                if len(v68) > 0 {
                                                    v68[0] = (fi29 * (v66 - v66))
                                                }
                                                if ((v66 * -408) < v66) {
                                                    v70 := GenS7{F0: 0, F1: 1, F2: 156}
                                                    fmt.Println(v2, v67)
                                                    _ = v70
                                                } else {
                                                    fmt.Println(fi29, v2)
                                                    v39 -= (-(v39))
                                                    v3 = "\n\t"
                                                }
                                                _ = v68
                                                _ = v69
                                            } else {
                                                v2 = v2
                                                v28 = v1
                                                var v71 int8 = v28
                                                _ = v71
                                            }
                                            _ = v66
                                            _ = v67
                                        } else {
                                            if true {
                                                v39 = 837
                                                var v72 int8 = v28
                                                var v73 int16 = v2
                                                _ = v72
                                                _ = v73
                                            }
                                            v74 := GenS1{F0: v36}
                                            v75 := []int{32767}
                                            _ = v74
                                            _ = v75
                                        }
                                    }
                                    fmt.Println(v2, v3, fi29)
                                    v76 := []int{(fi29 >> ((genFn0()) & 63)), -1, int(v36)}
                                    _ = v76
                                } else {
                                    v28 = (v28 % ((42) | 1))
                                }
                                _ = v39
                            }
                            v26 = v36
                            v77 := GenS8{F0: 211, F1: v28, F2: v36}
                            _ = fi35
                            _ = v36
                            _ = v77
                        }
                        v78 := GenS9{F0: 722, F1: 2147483647, F2: v26}
                        v79 := GenS3{F0: 850}
                        _ = v78
                        _ = v79
                    }
                case "":
                    for fi80 := 0; fi80 < 11; fi80++ {
                        v28 -= (-(v1))
                        if true {
                            v3 = "\n\t"
                            var v81 uint16 = 745
                            _ = v81
                        } else {
                            var v82 int = fi80
                            v83 := GenS10{F0: 255, F1: 32767, F2: 9223372036854775807}
                            if false {
                                fmt.Println(fi29, v2, v1)
                            }
                            switch v3 {
                            case "a":
                                v84 := GenS7{F0: 0, F1: 0, F2: 182}
                                _ = v84
                            case "\n\t":
                                if (1 != 0) {
                                    if (!false) {
                                        if (-9223372036854775808 != 0) {
                                            v85 := GenS8{F0: -9223372036854775808, F1: int8(fi29), F2: 4294967295}
                                            v86 := GenS11{F0: (v2 >= (v2 >> ((v26) & 63)))}
                                            var v87 uint32 = 4294967295
                                            _ = v85
                                            _ = v86
                                            _ = v87
                                        }
                                    }
                                    v28 -= v1
                                    var v88 int = fi80
                                    _ = v88
                                } else {
                                    var v89 string = "\n\t"
                                    _ = v89
                                }
                            case "":
                                var v90 float64 = v4
                                _ = v90
                            }
                            _ = v82
                            _ = v83
                        }
                        switch "\n\t" {
                        case "foo bar":
                            v2 += v2
                        case "\n\t":
                            var v91 int16 = (v2 >> ((4294967295) & 63))
                            _ = v91
                        case "hello":
                            v92 := GenS4{F0: 127}
                            _ = v92
                        default:
                            var v93 int64 = 1
                            _ = v93
                        }
                        _ = fi80
                    }
                }
                fmt.Println(v3)
                _ = fi29
            }
            v3 = v3
            if false {
                switch -460 {
                case 32767:
                    var v94 string = ""
                    _ = v94
                }
            }
        }
        v28 += v1
    } else {
        var v95 uint64 = 18446744073709551615
        v4 -= -2.25
        _ = v95
    }
    switch (("a" + (v3 + "foo bar")) + v3) {
    case "a":
        for fi96 := 0; fi96 < 1; fi96++ {
            v97 := []int{fi96, (fi96 & 676), fi96, fi96}
            _ = fi96
            _ = v97
        }
    case "foo bar":
        var v98 uint32 = 2147483648
        _ = v98
    }
    for fi99 := 0; fi99 < 8; fi99++ {
        fmt.Println(v1, v26)
        v100 := []int{fi99, (^(fi99))}
        fmt.Println(v3)
        _ = fi99
        _ = v100
    }
    var v101 float32 = (float32(v4) / ((3.14159 / 1.5) - float32((v4 / v4))))
    if (!((18446744073709551615 >= 42) || (v4 != float64(v101)))) {
        v102 := GenS12{F0: 1}
        v3 = "foo bar"
        _ = v102
    }
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v101
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
