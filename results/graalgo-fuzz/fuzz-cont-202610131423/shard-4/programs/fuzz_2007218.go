package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 int64; F2 uint }
type GenS3 struct { F0 uint32; F1 int16; F2 int64 }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16; F1 int64; F2 bool }
type GenS7 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS8 struct { F0 int; F1 int64; F2 uint16 }
type GenS9 struct { F0 int; F1 int }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int; F1 string; F2 int }
type GenS12 struct { F0 uint32; F1 int64; F2 bool }
type GenS13 struct { F0 int8 }
func genFn0(p0 bool, p1 int8) int64 {
    v6 := GenS1{F0: "a"}
    p1 += p1
    v7 := GenS1{F0: "foo bar"}
    if p0 {
        if true {
            for fi8 := 0; fi8 < 19; fi8++ {
                if true {
                    if (3.14159 != 3.14159) {
                        v9 := []int{(fi8 / (((fi8 % (((fi8 + (^(fi8)))) | 1))) | 1)), fi8}
                        v10 := GenS2{F0: (p1 + -128), F1: 988, F2: 1}
                        v11 := GenS3{F0: 65535, F1: 32767, F2: 0}
                        _ = v10
                        _ = v11
                        _ = v9
                    } else {
                        v12 := GenS4{F0: ("a" + "\n\t"), F1: ((p0 && p0) || (((151 != 0) && p0) || true))}
                        _ = v12
                    }
                } else {
                    v13 := GenS2{F0: (-(p1)), F1: 546, F2: 255}
                    var v14 uint16 = 2
                    switch "hello" {
                    case "foo bar":
                        v15 := []int{-1, (fi8 ^ (-(fi8)))}
                        _ = v15
                    case "":
                        v16 := []int{127}
                        _ = v16
                    }
                    _ = v13
                    _ = v14
                }
                _ = fi8
            }
            v17 := GenS5{F0: 9223372036854775807}
            v18 := GenS6{F0: 65535, F1: -794, F2: p0}
            _ = v17
            _ = v18
        } else {
            fmt.Println(p1)
        }
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        var v20 uint32 = 1
        _ = fi19
        _ = v20
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 9223372036854775807
}
func genFn1(p0 string) int32 {
    if (!false) {
        if (false && true) {
            v21 := []int{-200}
            v21 = append(v21, -1)
            if len(v21) > 0 {
                v21[0] = 2147483647
            }
            _ = v21
        } else {
            var v22 int32 = 42
            var v23 int32 = (v22 << ((705) & 63))
            if (493 <= 1) {
                v24 := GenS1{F0: "\n\t"}
                _ = v24
            } else {
                v22 = (^(v22))
                fmt.Println(v22, p0, p0)
            }
            _ = v22
            _ = v23
        }
        v25 := GenS7{F0: 215, F1: 255, F2: -128}
        _ = v25
    }
    var v26 uint8 = 1
    v27 := []int{4294967295, 128, 572}
    var v28 bool = false
    v28 = (0 > 42)
    _ = p0
    _ = v26
    _ = v27
    _ = v28
    return 1
}

func main() {
    v1 := make(map[int]string)
    var v2 string = ""
    var v3 string = "a"
    var v4 int64 = 42
    v5 := []string{v3}
    v5 = append(v5, "hello")
    for fi29 := 0; fi29 < 13; fi29++ {
        var v30 int16 = 42
        _ = fi29
        _ = v30
    }
    for fi31 := 0; fi31 < 17; fi31++ {
        var v32 string = v2
        v2 = (v3 + ("a" + v2))
        var v33 int64 = genFn0(true, 1)
        _ = fi31
        _ = v32
        _ = v33
    }
    if false {
        v34 := GenS8{F0: 32767, F1: (v4 + (v4 >> ((27) & 63))), F2: 2}
        v4 = -1
        _ = v34
    }
    var v35 uint16 = 32767
    for fi36 := 0; fi36 < 11; fi36++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi36
    }
    for fi37 := 0; fi37 < 7; fi37++ {
        var v38 uint32 = 2
        v39 := GenS9{F0: fi37, F1: -1}
        fmt.Println(v38, v4, v38)
        _ = fi37
        _ = v38
        _ = v39
    }
    var v40 float64 = 3.14159
    v2 = (v2 + ("a" + (("a" + v2) + v2)))
    if (-71 > 0) {
        var v41 int32 = 0
        for fi42 := 0; fi42 < 5; fi42++ {
            v43 := GenS6{F0: (v35 ^ 1), F1: (v4 - (v4 << ((833) & 63))), F2: (2 < 1)}
            var v44 uint = 217
            _ = fi42
            _ = v43
            _ = v44
        }
        for fi45 := 0; fi45 < 1; fi45++ {
            if true {
                v46 := GenS9{F0: fi45, F1: fi45}
                var v47 int64 = (v4 ^ v4)
                _ = v46
                _ = v47
            } else {
                var v48 uint = 4294967295
                _ = v48
            }
            fmt.Println(v40, v40)
            _ = fi45
        }
        _ = v41
    }
    for fi49 := 0; fi49 < 3; fi49++ {
        fmt.Println(v35)
        if true {
            v3 = "\n\t"
            var v50 bool = ((v4 < genFn0((!true), 0)) || true)
            for fi51 := 0; fi51 < 2; fi51++ {
                if v50 {
                    v1[(fi49 * 250)] = v2
                }
                _ = fi51
            }
            _ = v50
        } else {
            if (genFn1(v2) > -997) {
                var v52 int16 = 19
                fmt.Println(v35)
                _ = v52
            } else {
                v53 := GenS9{F0: fi49, F1: (fi49 ^ fi49)}
                v40 -= 1e10
                v35 -= v35
                _ = v53
            }
            _ = len(v5)
            v40 -= 0.0
        }
        if true {
            fmt.Println(v35)
        }
        _ = fi49
    }
    if false {
        for fi54 := 0; fi54 < 9; fi54++ {
            v55 := GenS10{F0: 1}
            var v56 int64 = v4
            _ = len(v1)
            _ = fi54
            _ = v55
            _ = v56
        }
        delete(v1, 2147483647)
    } else {
        v4 += (v4 & v4)
    }
    if (!(!false)) {
        v57 := GenS11{F0: 127, F1: v2, F2: -235}
        _ = v57
    } else {
        for fi58 := 0; fi58 < 1; fi58++ {
            var v59 float32 = 0.0
            if (!true) {
                v60 := GenS1{F0: "foo bar"}
                _ = v60
            }
            var v61 uint32 = 2
            _ = fi58
            _ = v59
            _ = v61
        }
        for fi62 := 0; fi62 < 9; fi62++ {
            var v63 int16 = 1000
            fmt.Println(v3, v4, v35)
            v64 := v1[430]
            _ = fi62
            _ = v63
            _ = v64
        }
    }
    v2 = v2
    var v65 float32 = (((1e10 * 3.14159) / (0.0 + (float32(0.1) * (0.1 / -2.25)))) / float32((float64(((float32(0.1) + 1e10) + (float32(((v40 + v40) - (0.1 - ((v40 / v40) - (1e10 - (3.14159 + 0.1)))))) + (3.14159 * (float32(((1e10 - v40) - float64(1.5))) / float32((float64((1e10 - (-2.25 * -2.25))) * 1.5))))))) / v40)))
    switch (v3 + v3) {
    case "\n\t":
        if (!(v65 != float32(v40))) {
            v1[32768] = "foo bar"
            if (!false) {
                fmt.Println(v2, v3)
            } else {
                var v66 float32 = (v65 - -2.25)
                v4 += genFn0(false, 30)
                _ = v66
            }
            if true {
                var v67 int16 = -1
                _ = v67
            }
        } else {
            if false {
                for fi68 := 0; fi68 < 12; fi68++ {
                    v3 = "\n\t"
                    v69 := GenS6{F0: v35, F1: 56, F2: true}
                    for fi70 := 0; fi70 < 11; fi70++ {
                        v71 := GenS6{F0: (v35 ^ (v35 << ((219) & 63))), F1: 9223372036854775807, F2: true}
                        if true {
                            _ = len(v5)
                            if (204 == 42) {
                                var v72 uint = 42
                                switch fi70 {
                                case -1:
                                    v73 := GenS12{F0: 42, F1: int64(v72), F2: true}
                                    _ = v73
                                case -9223372036854775808:
                                    switch "\n\t" {
                                    case "\n\t":
                                        if ((-2.25 * (v65 - v65)) > (((v65 / v65) + (((v65 / v65) + v65) - ((v65 - 1e10) + (v65 - (((1e10 + v65) + 1.5) / (v65 * v65)))))) * v65)) {
                                            fmt.Println(v40)
                                            switch "foo bar" {
                                            case "":
                                                v35 = (v35 - 1)
                                            case "foo bar":
                                                v1[(-(fi70))] = v3
                                            case "hello":
                                                v35 -= 32767
                                            default:
                                                var v74 string = v2
                                                _ = v74
                                            }
                                            if (!true) {
                                                if (true && true) {
                                                    v75 := GenS11{F0: 42, F1: ((v3 + v2) + v3), F2: fi70}
                                                    _ = v75
                                                } else {
                                                    fmt.Println(v2)
                                                    var v76 string = "foo bar"
                                                    _ = v76
                                                }
                                                if true {
                                                    var v77 uint = v72
                                                    fmt.Println(fi70, v40, v65)
                                                    _ = v77
                                                } else {
                                                    delete(v1, (fi70 / ((fi68) | 1)))
                                                }
                                                v2 = v3
                                            } else {
                                                v1[fi70] = ""
                                                v72 = v72
                                                var v78 int64 = (v4 + (^(v4)))
                                                _ = v78
                                            }
                                        }
                                    default:
                                        var v79 float32 = v65
                                        _ = v79
                                    }
                                }
                                _ = v72
                            }
                            fmt.Println(v35)
                        } else {
                            v35 = 32767
                        }
                        v1[(fi70 / ((fi70) | 1))] = v2
                        _ = fi70
                        _ = v71
                    }
                    _ = fi68
                    _ = v69
                }
                for fi80 := 0; fi80 < 3; fi80++ {
                    switch fi80 {
                    case -102:
                        v81 := GenS13{F0: 0}
                        _ = v81
                    case -224:
                        v4 = (v4 / ((v4) | 1))
                    case 1:
                        var v82 uint64 = 751
                        _ = v82
                    default:
                        v2 = v2
                    }
                    _ = fi80
                }
                var v83 bool = false
                _ = v83
            }
        }
    case "":
        var v84 int8 = 31
        _ = v84
    default:
        switch v3 {
        case "hello":
            var v85 int16 = 1
            _ = v85
        default:
            for fi86 := 0; fi86 < 15; fi86++ {
                v87 := GenS7{F0: (v35 ^ v35), F1: (v35 << ((18446744073709551615) & 63)), F2: -12}
                var v88 string = (v3 + v2)
                var v89 uint64 = 546
                _ = fi86
                _ = v87
                _ = v88
                _ = v89
            }
        }
    }
    for fi90 := 0; fi90 < 13; fi90++ {
        _ = len(v1)
        _ = fi90
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v40
    _ = v5
    _ = v65
}
