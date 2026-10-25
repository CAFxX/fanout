package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16; F1 int8; F2 int64 }
type GenS6 struct { F0 uint; F1 bool; F2 uint }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int16; F1 int64; F2 bool }
type GenS9 struct { F0 bool; F1 uint16 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int; F1 int8 }
type GenS12 struct { F0 int64 }
func genFn0(p0 uint8, p1 string, p2 uint) uint64 {
    for fi6 := 0; fi6 < 10; fi6++ {
        p1 = p1
        p2 -= p2
        if (0.1 == 1e10) {
            if ((p0 + 255) <= (p0 ^ 104)) {
                var v7 uint8 = (p0 >> ((p2) & 63))
                var v8 int64 = 710
                _ = v7
                _ = v8
            } else {
                v9 := GenS1{F0: 25, F1: (-(fi6)), F2: 0}
                if true {
                    var v10 int32 = 791
                    fmt.Println(v10, fi6, p0)
                    _ = v10
                } else {
                    v11 := GenS1{F0: 255, F1: fi6, F2: 0}
                    _ = v11
                }
                _ = v9
            }
        } else {
            fmt.Println(p2, p0, p1)
        }
        _ = fi6
    }
    var v12 int = 2147483647
    if false {
        v13 := []int{}
        _ = len(v13)
        if (!true) {
            var v14 float32 = 0.0
            var v15 uint8 = (p0 * 2)
            v12 = (v12 >> (((p2 - (p2 & 4294967295))) & 63))
            _ = v14
            _ = v15
        } else {
            v16 := GenS2{F0: -1}
            v17 := GenS2{F0: -367}
            _ = v16
            _ = v17
        }
        _ = v13
    }
    var v18 int32 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v18
    return 1
}
func genFn1(p0 int16, p1 uint) int {
    var v19 int64 = 42
    if (p0 == p0) {
        v20 := GenS2{F0: p0}
        var v21 float64 = 1e10
        _ = v20
        _ = v21
    } else {
        for fi22 := 0; fi22 < 3; fi22++ {
            fmt.Println(v19, p0, fi22)
            for fi23 := 0; fi23 < 1; fi23++ {
                if (-435 == 0) {
                    var v24 uint = 65535
                    _ = v24
                }
                if false {
                    switch fi23 {
                    case -20:
                        var v25 bool = false
                        _ = v25
                    case -577:
                        v26 := []int{fi22, (fi23 ^ (fi23 | (fi23 >> ((255) & 63)))), -600}
                        _ = v26
                    }
                    v27 := []int{(fi23 ^ (^(fi23)))}
                    switch ((("hello" + "") + (("hello" + "\n\t") + "")) + "foo bar") {
                    case "foo bar":
                        v28 := GenS2{F0: (p0 * -1)}
                        _ = v28
                    case "\n\t":
                        var v29 int64 = v19
                        var v30 uint8 = 42
                        p0 = (p0 % ((32767) | 1))
                        v31 := GenS3{F0: 42}
                        _ = v29
                        _ = v30
                        _ = v31
                    }
                    _ = v27
                }
                if (false || (-9223372036854775808 == (fi22 % ((int(v19)) | 1)))) {
                    var v32 int64 = (v19 | -886)
                    _ = v32
                } else {
                    p0 -= (p0 % ((p0) | 1))
                    p1 = p1
                    v33 := []int{fi22, (fi23 + fi22)}
                    v34 := GenS4{F0: (^(p1))}
                    _ = v33
                    _ = v34
                }
                _ = fi23
            }
            v35 := []int{(-(fi22)), 875, fi22}
            _ = fi22
            _ = v35
        }
        v36 := []int{127, 1}
        var v37 int64 = v19
        _ = v36
        _ = v37
    }
    v38 := GenS5{F0: p0, F1: 42, F2: -9223372036854775808}
    p1 = 65535
    _ = p0
    _ = p1
    _ = v19
    _ = v38
    return 686
}

func main() {
    var v1 int32 = 0
    v2 := []string{"\n\t", ("" + ""), ""}
    var v3 float32 = (1.5 - 1.5)
    var v4 uint8 = 2
    v5 := make(map[string]uint32)
    switch ("" + "\n\t") {
    case "":
        var v39 int16 = 0
        _ = v39
    case "foo bar":
        var v40 int8 = 0
        _ = v40
    }
    v41 := GenS6{F0: 255, F1: false, F2: 235}
    var v42 uint32 = 42
    v43 := GenS4{F0: 1}
    for fi44 := 0; fi44 < 1; fi44++ {
        if (true || ((true && true) && false)) {
            if (!false) {
                switch (fi44 / (((fi44 & 9223372036854775807)) | 1)) {
                case 506:
                    var v45 uint8 = v4
                    _ = v45
                }
            }
        } else {
            fmt.Println(fi44, fi44)
            for fi46 := 0; fi46 < 4; fi46++ {
                var v47 uint32 = (v42 - 192)
                v2 = append(v2, ("\n\t" + "hello"))
                v48 := GenS4{F0: 42}
                _ = fi46
                _ = v47
                _ = v48
            }
            for fi49 := 0; fi49 < 3; fi49++ {
                if false {
                    if false {
                        v3 = ((v3 + v3) / v3)
                    } else {
                        switch genFn1(-1, 1) {
                        case 9223372036854775807:
                            if (32767 != 42) {
                                v4 = (v4 >> ((542) & 63))
                                v42 = v42
                                var v50 int64 = 895
                                switch "" {
                                case "\n\t":
                                    v51 := v5["a"]
                                    _ = v51
                                case "foo bar":
                                    if len(v2) > 0 {
                                        v2[0] = "\n\t"
                                    }
                                }
                                _ = v50
                            } else {
                                var v52 int64 = 220
                                if (-61 >= 38) {
                                    if len(v2) > 0 {
                                        v2[0] = "foo bar"
                                    }
                                    if true {
                                        fmt.Println(fi49)
                                        var v53 uint32 = 0
                                        switch 0 {
                                        case 2147483647:
                                            fmt.Println(v4, v4)
                                        case 4294967295:
                                            _ = len(v2)
                                        default:
                                            v54 := GenS5{F0: 32767, F1: 127, F2: -9223372036854775808}
                                            _ = v54
                                        }
                                        _ = v53
                                    }
                                    var v55 int32 = v1
                                    v56 := GenS4{F0: 0}
                                    switch ("foo bar" + ((((("a" + "a") + "") + ("foo bar" + "\n\t")) + "\n\t") + "hello")) {
                                    case "":
                                        v42 = (v42 ^ 2)
                                    default:
                                        var v57 bool = false
                                        _ = v57
                                    }
                                    switch fi49 {
                                    case 54:
                                        var v58 string = "hello"
                                        _ = v58
                                    case 32768:
                                        if true {
                                            if false {
                                                if true {
                                                    var v59 float64 = 1.5
                                                    _ = v59
                                                } else {
                                                    v5[("a" + (("" + "\n\t") + "hello"))] = (v42 ^ 392)
                                                    fmt.Println(v3, v1, fi44)
                                                    v42 += (v42 << ((2) & 63))
                                                }
                                            }
                                            var v60 int64 = v52
                                            _ = v60
                                        } else {
                                            v61 := GenS4{F0: 1}
                                            _ = v61
                                        }
                                    case 0:
                                        if true {
                                            if len(v2) > 0 {
                                                v2[0] = "hello"
                                            }
                                            v62 := GenS1{F0: (v4 >> ((1) & 63)), F1: (^(fi49)), F2: 42}
                                            _ = v62
                                        }
                                    default:
                                        v3 = (v3 * 3.14159)
                                        v52 += (v52 - -651)
                                        var v63 int = genFn1(-32768, 18446744073709551615)
                                        _ = v63
                                    }
                                    switch "hello" {
                                    case "":
                                        v52 += v52
                                    case "hello":
                                        var v64 float32 = 0.1
                                        _ = v64
                                    case "a":
                                        _ = len(v2)
                                    }
                                    switch ("hello" + "foo bar") {
                                    case "":
                                        if (!(2 < 1)) {
                                            v65 := GenS4{F0: 1}
                                            var v66 uint64 = 1
                                            _ = v65
                                            _ = v66
                                        } else {
                                            if ((fi44 ^ (fi44 & (fi44 ^ fi49))) >= 32) {
                                                var v67 int8 = 10
                                                v68 := GenS7{F0: false}
                                                delete(v5, "")
                                                _ = v67
                                                _ = v68
                                            } else {
                                                v1 -= 813
                                            }
                                            v3 -= (-2.25 - -2.25)
                                            v2 = append(v2, "hello")
                                        }
                                    case "hello":
                                        v69 := v5["hello"]
                                        _ = v69
                                    default:
                                        var v70 uint64 = 932
                                        _ = v70
                                    }
                                    _ = v55
                                    _ = v56
                                } else {
                                    var v71 bool = true
                                    if ((!(float32(0.0) < (v3 / 0.1))) || (fi49 == (fi49 | (-(fi44))))) {
                                        v3 = v3
                                        v72 := GenS8{F0: 32767, F1: v52, F2: true}
                                        _ = v72
                                    }
                                    _ = v71
                                }
                                _ = v52
                            }
                        case 4294967295:
                            if (640 < 75) {
                                v73 := GenS9{F0: (0 < 100), F1: 83}
                                _ = v73
                            } else {
                                var v74 uint16 = 255
                                var v75 uint = 2
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v76 float32 = ((-2.25 + 0.1) * v3)
                                v76 += v3
                                _ = v74
                                _ = v75
                                _ = v76
                            }
                        case 32767:
                            v77 := GenS10{F0: -457}
                            _ = v77
                        default:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        fmt.Println(v1)
                    }
                    switch ("" + "foo bar") {
                    case "a":
                        var v78 uint32 = 695
                        _ = v78
                    default:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                } else {
                    var v79 int16 = 42
                    _ = v79
                }
                v80 := GenS11{F0: genFn1(32767, 0), F1: 0}
                v2 = append(v2, "hello")
                _ = fi49
                _ = v80
            }
        }
        _ = fi44
    }
    v81 := GenS3{F0: 1}
    for fi82 := 0; fi82 < 6; fi82++ {
        var v83 uint = 791
        var v84 int32 = 2147483647
        _ = fi82
        _ = v83
        _ = v84
    }
    v85 := GenS3{F0: 0}
    var v86 int64 = -9223372036854775808
    v87 := GenS12{F0: int64(v1)}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v81
    _ = v85
    _ = v86
    _ = v87
}
