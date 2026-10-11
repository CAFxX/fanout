package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 uint; F1 uint; F2 int16 }
type GenS4 struct { F0 string; F1 uint64; F2 uint }
type GenS5 struct { F0 int64; F1 int8 }
type GenS6 struct { F0 int; F1 int }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int; F1 int32; F2 int16 }
type GenS9 struct { F0 string; F1 uint; F2 uint32 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int64; F1 uint64; F2 int8 }
type GenS12 struct { F0 int16; F1 int8; F2 int8 }
type GenS13 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS14 struct { F0 int64 }
func genFn0() uint {
    for fi4 := 0; fi4 < 1; fi4++ {
        var v5 uint = 0
        switch 127 {
        case 32767:
            v6 := []int{(^(fi4))}
            _ = v6
        default:
            if (764 == 42) {
                if false {
                    if false {
                        v7 := GenS1{F0: 2}
                        _ = v7
                    }
                    v8 := GenS2{F0: 42, F1: 32767}
                    _ = v8
                }
                var v9 float64 = 0.0
                _ = v9
            } else {
                var v10 int64 = 1
                _ = v10
            }
        }
        v5 -= 1
        _ = fi4
        _ = v5
    }
    var v11 int8 = 48
    v12 := GenS1{F0: 891}
    v13 := []int{-9223372036854775808, 127}
    var v14 uint64 = 817
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return 18446744073709551615
}
func genFn1() string {
    var v15 uint8 = 1
    var v16 uint64 = 1
    if false {
        v17 := []int{32768, 127, 127, 128}
        for fi18 := 0; fi18 < 2; fi18++ {
            for fi19 := 0; fi19 < 5; fi19++ {
                fmt.Println(v15, fi18, fi19)
                switch "hello" {
                case "hello":
                    if false {
                        var v20 int64 = -462
                        fmt.Println(v16, fi19)
                        _ = v20
                    } else {
                        var v21 uint64 = v16
                        _ = v21
                    }
                case "foo bar":
                    v16 = 42
                }
                v15 -= 2
                _ = fi19
            }
            v22 := GenS3{F0: 65535, F1: 753, F2: 42}
            v15 -= v15
            _ = fi18
            _ = v22
        }
        _ = v17
    } else {
        fmt.Println(v15, v15)
    }
    v23 := []int{1, -1, -1, 18}
    if len(v23) >= 2 {
        v23 = v23[1:len(v23)]
    }
    _ = v15
    _ = v16
    _ = v23
    return (("hello" + "foo bar") + "a")
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint = 255
    var v3 bool = (4294967295 > 65535)
    var v24 int64 = -498
    delete(v1, 1)
    v25 := v1[65535]
    v3 = (!(0.0 >= 0.0))
    var v26 int16 = 327
    var v27 float32 = 0.1
    v2 = 0
    for fi28 := 0; fi28 < 11; fi28++ {
        var v29 uint16 = 255
        fmt.Println(v29)
        v30 := GenS4{F0: "foo bar", F1: 4294967295, F2: genFn0()}
        _ = fi28
        _ = v29
        _ = v30
    }
    if (false || (!(v3 && (0 > -1)))) {
        v31 := GenS5{F0: int64(v25), F1: 42}
        fmt.Println(v2, v2, v3)
        for fi32 := 0; fi32 < 17; fi32++ {
            v3 = true
            if v3 {
                v33 := GenS6{F0: fi32, F1: -21}
                _ = v33
            } else {
                var v34 int64 = -1
                var v35 string = "hello"
                v1[(v25 + (v25 / ((v25) | 1)))] = v25
                _ = v34
                _ = v35
            }
            v36 := GenS5{F0: v24, F1: -119}
            _ = fi32
            _ = v36
        }
        _ = v31
    } else {
        for fi37 := 0; fi37 < 15; fi37++ {
            v38 := GenS3{F0: (v2 - (v2 * v2)), F1: (v2 / ((255) | 1)), F2: v26}
            if false {
                v25 += v25
                v24 = v24
                v39 := GenS1{F0: 2}
                _ = v39
            } else {
                for fi40 := 0; fi40 < 18; fi40++ {
                    fmt.Println(v2, v2, v25)
                    v26 = v26
                    _ = len(v1)
                    if ((v24 << ((18446744073709551615) & 63)) != v24) {
                        var v41 int = 2147483647
                        _ = v41
                    } else {
                        if (128 <= 255) {
                            v42 := GenS7{F0: 570}
                            var v43 int16 = v26
                            var v44 int32 = 2147483647
                            if true {
                                v45 := GenS3{F0: (v2 & (v2 - (v2 | v2))), F1: v2, F2: (v43 % ((v43) | 1))}
                                v27 = -2.25
                                if v3 {
                                    v24 = v24
                                    v46 := []int{4294967295, fi40, (fi40 / ((fi40) | 1)), (^(fi37))}
                                    _ = v46
                                } else {
                                    if false {
                                        v47 := GenS4{F0: "", F1: 18446744073709551615, F2: (v2 % ((18446744073709551615) | 1))}
                                        v2 -= 2
                                        _ = v47
                                    } else {
                                        var v48 float32 = 0.0
                                        var v49 int32 = v44
                                        var v50 int32 = -1
                                        _ = v48
                                        _ = v49
                                        _ = v50
                                    }
                                    if false {
                                        var v51 int8 = -120
                                        _ = v51
                                    } else {
                                        var v52 string = "a"
                                        if (!(64 < 128)) {
                                            v53 := GenS7{F0: 42}
                                            var v54 uint32 = v25
                                            v55 := GenS7{F0: 1}
                                            _ = v53
                                            _ = v54
                                            _ = v55
                                        }
                                        v2 -= v2
                                        _ = v52
                                    }
                                    if (v27 <= (0.0 * 1.5)) {
                                        v24 -= (v24 * -1)
                                        fmt.Println(v26)
                                        v56 := GenS8{F0: (fi37 - (fi40 % (((fi37 + fi40)) | 1))), F1: v44, F2: v43}
                                        _ = v56
                                    }
                                }
                                _ = len(v1)
                                v57 := []int{(^(fi40))}
                                _ = v45
                                _ = v57
                            }
                            _ = v42
                            _ = v43
                            _ = v44
                        }
                        _ = len(v1)
                    }
                    _ = fi40
                }
            }
            _ = fi37
            _ = v38
        }
        var v58 uint32 = v25
        for fi59 := 0; fi59 < 19; fi59++ {
            v60 := GenS9{F0: ("" + "\n\t"), F1: v2, F2: v58}
            if v3 {
                for fi61 := 0; fi61 < 8; fi61++ {
                    switch -1 {
                    case -141:
                        v62 := v1[v58]
                        _ = v62
                    case -819:
                        switch (fi61 * fi59) {
                        case 797:
                            fmt.Println(v25, v3)
                            if v3 {
                                switch "hello" {
                                case "hello":
                                    v63 := GenS7{F0: 42}
                                    _ = v63
                                case "foo bar":
                                    v26 += int16(v2)
                                case "a":
                                    switch "\n\t" {
                                    case "":
                                        var v64 int64 = (v24 + 1)
                                        v65 := []int{fi59, -1, fi59, (-(fi59))}
                                        _ = v64
                                        _ = v65
                                    case "foo bar":
                                        var v66 uint = v2
                                        _ = v66
                                    case "hello":
                                        if (3.14159 >= -2.25) {
                                            var v67 int = (fi61 | (-(fi61)))
                                            _ = v67
                                        }
                                    }
                                default:
                                    var v68 string = ""
                                    _ = v68
                                }
                            }
                            fmt.Println(v2, v58, v58)
                        case -959:
                            if ((v24 ^ v24) == v24) {
                                fmt.Println(v24, v25)
                                v69 := []int{(fi59 & (fi61 * (fi61 + (fi61 ^ (-(fi61)))))), (-(fi59)), (fi61 / ((fi61) | 1)), fi59}
                                _ = v69
                            }
                        }
                    }
                    v70 := GenS10{F0: -128}
                    _ = fi61
                    _ = v70
                }
            }
            _ = fi59
            _ = v60
        }
        _ = v58
    }
    if (!(128 != 12)) {
        for fi71 := 0; fi71 < 4; fi71++ {
            v72 := GenS11{F0: (v24 * (v24 - v24)), F1: 4294967295, F2: 0}
            v26 -= -1
            _ = fi71
            _ = v72
        }
        if false {
            v73 := []int{}
            _ = v73
        }
        var v74 bool = (-144 < 9223372036854775807)
        _ = v74
    } else {
        if v3 {
            var v75 int = 32768
            fmt.Println(v25)
            v76 := []int{v75, int(v25), (-(v75)), (v75 & (v75 * (v75 - -423)))}
            _ = v75
            _ = v76
        }
        v77 := GenS10{F0: 42}
        _ = v77
    }
    if (42 != -128) {
        var v78 uint32 = v25
        _ = v78
    }
    fmt.Println(v2, v26)
    if (267 < 54) {
        switch 196 {
        case 168:
            for fi79 := 0; fi79 < 3; fi79++ {
                if (32767 != 65535) {
                    switch "a" {
                    case "":
                        v80 := []int{}
                        _ = v80
                    default:
                        var v81 uint8 = 31
                        _ = v81
                    }
                    fmt.Println(v3)
                    switch "\n\t" {
                    case "foo bar":
                        v82 := GenS12{F0: v26, F1: -90, F2: -106}
                        _ = v82
                    case "a":
                        if v3 {
                            for fi83 := 0; fi83 < 1; fi83++ {
                                if v3 {
                                    var v84 uint = (v2 / ((255) | 1))
                                    v85 := GenS9{F0: "\n\t", F1: v2, F2: 965}
                                    _ = len(v1)
                                    fmt.Println(v24, fi79)
                                    _ = v84
                                    _ = v85
                                }
                                v25 -= (v25 * 267)
                                _ = fi83
                            }
                            var v86 int64 = -1
                            var v87 bool = (true || false)
                            _ = v86
                            _ = v87
                        } else {
                            if (int64(fi79) < (v24 & (v24 % ((v24) | 1)))) {
                                for fi88 := 0; fi88 < 17; fi88++ {
                                    v89 := GenS13{F0: 111, F1: 127, F2: (v25 % ((v25) | 1))}
                                    _ = fi88
                                    _ = v89
                                }
                            } else {
                                if (255 > 0) {
                                    v27 = v27
                                    var v90 bool = v3
                                    var v91 int64 = -1
                                    _ = v90
                                    _ = v91
                                }
                                v92 := GenS14{F0: (v24 * 9223372036854775807)}
                                _ = v92
                            }
                        }
                    default:
                        v3 = ((v3 && v3) && false)
                    }
                }
                _ = fi79
            }
        default:
            var v93 int16 = (v26 % ((v26) | 1))
            _ = v93
        }
    } else {
        switch "foo bar" {
        case "":
            fmt.Println(v26, v24)
        case "a":
            v26 = v26
        }
        var v94 string = ("" + "hello")
        _ = v94
    }
    v2 += 2
    v25 -= v25
    v3 = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
