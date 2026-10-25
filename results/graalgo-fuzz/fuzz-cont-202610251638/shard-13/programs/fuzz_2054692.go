package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 string }
func genFn0() int {
    for fi7 := 0; fi7 < 14; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 int = (-(fi7))
            _ = fi8
            _ = v9
        }
        fmt.Println(fi7)
        if false {
            if (false && true) {
                v10 := []int{fi7}
                var v11 uint8 = 67
                for fi12 := 0; fi12 < 9; fi12++ {
                    v11 = (v11 - v11)
                    v13 := GenS1{F0: v11, F1: -128}
                    _ = fi12
                    _ = v13
                }
                _ = v10
                _ = v11
            }
        }
        _ = fi7
    }
    switch -383 {
    case 127:
        var v14 int16 = 0
        _ = v14
    case 690:
        switch "foo bar" {
        case "":
            v15 := []int{}
            _ = v15
        case "\n\t":
            var v16 uint16 = 255
            _ = v16
        }
    default:
        var v17 int = -9223372036854775808
        _ = v17
    }
    var v18 uint16 = 42
    for fi19 := 0; fi19 < 8; fi19++ {
        for fi20 := 0; fi20 < 10; fi20++ {
            if ("foo bar" <= ("\n\t" + ("" + "\n\t"))) {
                v21 := GenS1{F0: 18, F1: 42}
                if true {
                    fmt.Println(fi20, fi20, fi20)
                    v18 = 196
                    v18 -= v18
                }
                if false {
                    var v22 uint32 = 65535
                    _ = v22
                } else {
                    v23 := []int{}
                    if len(v23) >= 2 {
                        v23 = v23[1:len(v23)]
                    }
                    _ = v23
                }
                _ = v21
            } else {
                fmt.Println(fi20, v18)
                var v24 int64 = 42
                _ = v24
            }
            if false {
                switch ("\n\t" + "") {
                case "\n\t":
                    v25 := GenS2{F0: 0}
                    _ = v25
                case "foo bar":
                    var v26 bool = true
                    _ = v26
                case "":
                    v27 := []int{fi20, 32768, fi19}
                    _ = v27
                default:
                    var v28 int64 = 0
                    _ = v28
                }
                var v29 string = "hello"
                v30 := GenS2{F0: -3}
                switch v29 {
                case "a":
                    v31 := []int{fi20, 863, (fi19 ^ (-(fi19)))}
                    _ = v31
                case "":
                    v32 := GenS1{F0: 46, F1: -1}
                    _ = v32
                case "hello":
                    var v33 bool = (42 >= 293)
                    v33 = (2 > 65535)
                    v34 := []int{}
                    if v33 {
                        var v35 int = 365
                        var v36 float32 = float32((((0.1 / 3.14159) * (float64(0.1) * ((3.14159 * ((1e10 + 1e10) / -2.25)) + -2.25))) - -2.25))
                        var v37 uint16 = (^(v18))
                        _ = v35
                        _ = v36
                        _ = v37
                    } else {
                        var v38 bool = v33
                        _ = v38
                    }
                    var v39 uint64 = 298
                    v40 := GenS1{F0: 42, F1: 73}
                    v18 += v18
                    _ = v33
                    _ = v34
                    _ = v39
                    _ = v40
                }
                v18 = 228
                _ = v29
                _ = v30
            }
            _ = fi20
        }
        v18 += v18
        var v41 uint16 = (^(v18))
        _ = fi19
        _ = v41
    }
    var v42 int = 808
    _ = v18
    _ = v42
    return (v42 % ((v42) | 1))
}
func genFn1(p0 int32) string {
    var v43 bool = (0.0 == ((0.0 * 3.14159) / ((1.5 + -2.25) - 3.14159)))
    switch genFn0() {
    case 2147483647:
        fmt.Println(p0, p0)
    case -1:
        p0 = (p0 / ((292) | 1))
    default:
        var v44 int = 128
        _ = v44
    }
    var v45 int32 = p0
    _ = p0
    _ = v43
    _ = v45
    return ("a" + "")
}

func main() {
    v1 := make(map[int]bool)
    var v2 string = "\n\t"
    var v3 int64 = -313
    var v4 uint = 4294967295
    v5 := []string{((v2 + v2) + ((v2 + v2) + v2)), (("hello" + "a") + (("\n\t" + v2) + "hello")), "", ""}
    v6 := []int{}
    fmt.Println(v2)
    _ = len(v1)
    if (true || (0 >= 42)) {
        v1[842] = (!(true || false))
    } else {
        switch 875 {
        case 2147483647:
            v1[4294967295] = (-522 == 1000)
        case -316:
            v46 := GenS3{F0: 32767, F1: "foo bar"}
            _ = v46
        case 0:
            v2 = v2
        default:
            v2 = v2
        }
    }
    if ((1 > 42) || (false && (889 <= 253))) {
        for fi47 := 0; fi47 < 17; fi47++ {
            v2 = (v2 + v2)
            v4 -= v4
            _ = fi47
        }
        v3 += (v3 + (v3 / ((v3) | 1)))
        if false {
            var v48 float64 = 1e10
            _ = len(v1)
            fmt.Println(v2)
            _ = v48
        } else {
            if (!true) {
                v2 = v2
                fmt.Println(v2)
            }
        }
    }
    if (716 == 145) {
        var v49 int = genFn0()
        if (!((164 == 42) && false)) {
            v4 = 65535
            _ = len(v5)
            for fi50 := 0; fi50 < 17; fi50++ {
                _ = len(v5)
                if ((v3 >= -223) && false) {
                    v49 -= (^(v49))
                    for fi51 := 0; fi51 < 1; fi51++ {
                        switch (v49 >> ((v4) & 63)) {
                        case 4294967295:
                            fmt.Println(fi50, v4, v49)
                            v2 = (v2 + v2)
                        case 2147483647:
                            v3 = -1
                        case -513:
                            if (float32(1.5) < 1e10) {
                                var v52 uint = 111
                                switch fi50 {
                                case -292:
                                    var v53 uint32 = 252
                                    _ = v53
                                default:
                                    if true {
                                        if (700 <= (v3 + v3)) {
                                            delete(v1, -2147483648)
                                            _ = len(v5)
                                            if (!(!((!false) && (true || ((((v2 + ("\n\t" + ((v2 + v2) + v2))) + "") + v2) < (v2 + (("hello" + "a") + "hello"))))))) {
                                                if (((^(v3)) != v3) && (917 == int32(v49))) {
                                                    var v54 float64 = (-2.25 + -2.25)
                                                    v3 = (v3 | v3)
                                                    v55 := GenS2{F0: -1}
                                                    _ = v54
                                                    _ = v55
                                                }
                                                v52 -= (v4 / ((v4) | 1))
                                                var v56 uint = 42
                                                _ = v56
                                            } else {
                                                v1[0] = ((fi51 & (fi51 & genFn0())) > 42)
                                            }
                                        } else {
                                            var v57 int32 = int32(fi50)
                                            _ = v57
                                        }
                                    }
                                }
                                v58 := GenS2{F0: -1}
                                v59 := v1[genFn0()]
                                _ = v52
                                _ = v58
                                _ = v59
                            } else {
                                var v60 int = (v49 * genFn0())
                                var v61 int64 = v3
                                _ = v60
                                _ = v61
                            }
                        default:
                            _ = len(v1)
                        }
                        switch 32767 {
                        case 0:
                            if (!(255 != 982)) {
                                var v62 int32 = 0
                                if false {
                                    _ = len(v6)
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    var v63 int = genFn0()
                                    fmt.Println(v49)
                                    if true {
                                        var v64 uint32 = uint32(v4)
                                        v2 = v2
                                        _ = v64
                                    }
                                    v62 = 2147483647
                                    v65 := GenS1{F0: 42, F1: -1}
                                    _ = v63
                                    _ = v65
                                } else {
                                    if (3.14159 >= 1e10) {
                                        _ = len(v5)
                                    }
                                }
                                v1[fi51] = true
                                _ = v62
                            } else {
                                v5 = append(v5, v2)
                            }
                        case 4294967295:
                            var v66 int8 = 2
                            _ = v66
                        }
                        _ = fi51
                    }
                } else {
                    for fi67 := 0; fi67 < 4; fi67++ {
                        v68 := v1[(-(fi67))]
                        v5 = append(v5, (("\n\t" + "\n\t") + v2))
                        var v69 int16 = 1
                        _ = fi67
                        _ = v68
                        _ = v69
                    }
                    var v70 uint64 = 1
                    switch "a" {
                    case "":
                        v49 = (fi50 >> (((v4 / (((v4 * 42)) | 1))) & 63))
                    case "hello":
                        for fi71 := 0; fi71 < 2; fi71++ {
                            v72 := GenS1{F0: 2, F1: 42}
                            _ = fi71
                            _ = v72
                        }
                    case "a":
                        var v73 int = (fi50 % ((genFn0()) | 1))
                        _ = v73
                    default:
                        v74 := v1[(v49 | v49)]
                        _ = v74
                    }
                    _ = v70
                }
                _ = len(v6)
                _ = fi50
            }
        }
        _ = v49
    } else {
        v3 -= 1
        if (-110 <= -1) {
            var v75 float64 = float64((0.1 + -2.25))
            _ = v75
        } else {
            v76 := GenS4{F0: v2}
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            for fi77 := 0; fi77 < 3; fi77++ {
                var v78 string = ("hello" + (v2 + ("a" + (("hello" + v2) + v2))))
                _ = fi77
                _ = v78
            }
            _ = v76
        }
        var v79 int = 574
        _ = v79
    }
    var v80 int32 = -2147483648
    v80 = v80
    v3 -= (v3 & v3)
    fmt.Println(v80)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v80
}
