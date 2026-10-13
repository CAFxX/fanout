package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool; F1 uint32; F2 uint8 }
type GenS5 struct { F0 int32; F1 uint }
type GenS6 struct { F0 string }
type GenS7 struct { F0 bool; F1 int8 }
type GenS8 struct { F0 bool; F1 uint16; F2 uint }
type GenS9 struct { F0 uint; F1 uint8; F2 uint8 }
type GenS10 struct { F0 bool; F1 uint32; F2 uint64 }
type GenS11 struct { F0 uint32; F1 int16; F2 int32 }
type GenS12 struct { F0 int32; F1 uint32; F2 int32 }
type GenS13 struct { F0 int32; F1 int8 }
func genFn0(p0 uint, p1 int, p2 int32) uint8 {
    p0 = p0
    v4 := GenS1{F0: 4294967295, F1: 428}
    for fi5 := 0; fi5 < 9; fi5++ {
        switch "hello" {
        case "\n\t":
            v6 := GenS1{F0: 1, F1: 701}
            _ = v6
        case "hello":
            v7 := GenS1{F0: 2, F1: 373}
            _ = v7
        case "foo bar":
            fmt.Println(p2)
        default:
            switch ("hello" + ("foo bar" + (((((("" + "a") + "hello") + "foo bar") + "\n\t") + "a") + (("" + "") + ("a" + "\n\t"))))) {
            case "hello":
                v8 := []int{p1, (fi5 & 595), -9223372036854775808, (fi5 - p1)}
                _ = v8
            case "foo bar":
                var v9 int32 = (p2 << ((1) & 63))
                _ = v9
            case "\n\t":
                p2 = 0
            }
        }
        _ = fi5
    }
    p1 -= 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 0
}
func genFn1(p0 bool, p1 int16) int8 {
    var v10 int = 335
    var v11 uint64 = 0
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 1
}

func main() {
    v1 := []string{"hello", "\n\t", "foo bar"}
    var v2 string = "a"
    v3 := make(map[int]string)
    for fi12 := 0; fi12 < 14; fi12++ {
        _ = len(v1)
        _ = fi12
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 int16 = 717
            v16 := GenS1{F0: 18446744073709551615, F1: (-(v15))}
            switch v2 {
            case "hello":
                switch (fi14 / ((fi13) | 1)) {
                case 128:
                    var v17 uint16 = 65535
                    _ = v17
                }
            default:
                _ = len(v1)
            }
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi13
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        v2 = v2
        for fi19 := 0; fi19 < 15; fi19++ {
            fmt.Println(fi18, fi19, fi19)
            _ = fi19
        }
        if true {
            v3[fi18] = v2
            switch (fi18 - (fi18 - (^(fi18)))) {
            case -2147483648:
                fmt.Println(fi18, fi18)
            case 2147483647:
                v2 = ""
            case -438:
                if false {
                    v20 := GenS1{F0: 902, F1: 32767}
                    _ = v20
                } else {
                    var v21 int64 = -9223372036854775808
                    switch ((("foo bar" + "\n\t") + "a") + ((v2 + "a") + v2)) {
                    case "foo bar":
                        v22 := GenS1{F0: 107, F1: 1}
                        _ = v22
                    case "\n\t":
                        v21 -= v21
                    case "a":
                        for fi23 := 0; fi23 < 11; fi23++ {
                            if true {
                                v24 := GenS1{F0: 59, F1: 0}
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v24
                            } else {
                                if (60 == -128) {
                                    v2 = ""
                                    v25 := GenS2{F0: 0}
                                    v3[(fi23 ^ (^(fi23)))] = ("" + (v2 + v2))
                                    if (!true) {
                                        var v26 bool = false
                                        v21 -= (v21 >> ((65535) & 63))
                                        _ = v26
                                    }
                                    fmt.Println(v2, fi23, v2)
                                    _ = v25
                                } else {
                                    if (true || false) {
                                        v27 := GenS2{F0: 32767}
                                        var v28 string = v2
                                        _ = v27
                                        _ = v28
                                    } else {
                                        v29 := GenS2{F0: 42}
                                        switch v2 {
                                        case "foo bar":
                                            fmt.Println(fi18, v2)
                                            switch 42 {
                                            case 1:
                                                v21 = 42
                                            default:
                                                switch "\n\t" {
                                                case "foo bar":
                                                    v1 = append(v1, v2)
                                                default:
                                                    v21 += (v21 ^ 595)
                                                }
                                            }
                                        case "a":
                                            if false {
                                                delete(v3, fi18)
                                            } else {
                                                v30 := GenS1{F0: 2, F1: 32767}
                                                _ = v30
                                            }
                                        case "hello":
                                            if len(v1) > 0 {
                                                v1[0] = ("foo bar" + v2)
                                            }
                                        default:
                                            fmt.Println(v2, v21, v21)
                                        }
                                        _ = v29
                                    }
                                    v21 += 0
                                    if false {
                                        var v31 int64 = 42
                                        _ = v31
                                    }
                                }
                                v21 += v21
                            }
                            if (!true) {
                                var v32 uint = 645
                                var v33 uint32 = 4294967295
                                v21 = (v21 >> ((42) & 63))
                                _ = v32
                                _ = v33
                            }
                            var v34 uint64 = 121
                            _ = fi23
                            _ = v34
                        }
                    }
                    _ = v21
                }
            default:
                for fi35 := 0; fi35 < 5; fi35++ {
                    delete(v3, -649)
                    v36 := GenS3{F0: 212}
                    _ = fi35
                    _ = v36
                }
            }
        }
        _ = fi18
    }
    switch ("foo bar" + (((v2 + v2) + "") + v2)) {
    case "":
        if (!(false || (927 > 42))) {
            for fi37 := 0; fi37 < 14; fi37++ {
                v2 = ""
                _ = fi37
            }
            v2 = (v2 + v2)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            fmt.Println(v2)
            if true {
                v38 := GenS3{F0: 65535}
                var v39 int = -1
                _ = v38
                _ = v39
            } else {
                var v40 bool = true
                v41 := GenS4{F0: true, F1: 2, F2: 1}
                v2 = "foo bar"
                _ = v40
                _ = v41
            }
            v42 := GenS5{F0: 42, F1: 65535}
            _ = v42
        }
    case "foo bar":
        v2 = (v2 + ((("foo bar" + "a") + (v2 + "a")) + v2))
    }
    v43 := GenS6{F0: v2}
    if true {
        fmt.Println(v2)
        delete(v3, -437)
        var v44 uint = 42
        _ = v44
    } else {
        fmt.Println(v2)
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v45 := GenS7{F0: (false && (!true)), F1: 127}
    v46 := GenS8{F0: (false && ((false || (!((((0.1 + (-2.25 + float32(float64((0.1 * -2.25))))) * 1.5) <= 1e10) || false))) && true)), F1: 1, F2: 1}
    for fi47 := 0; fi47 < 14; fi47++ {
        for fi48 := 0; fi48 < 11; fi48++ {
            var v49 uint8 = 1
            if false {
                var v50 int16 = 257
                _ = v50
            } else {
                if true {
                    switch (fi48 - -1) {
                    case 42:
                        v51 := GenS9{F0: 18446744073709551615, F1: 128, F2: genFn0(712, fi47, 0)}
                        _ = v51
                    default:
                        if (false && (!(true && false))) {
                            v52 := v3[fi47]
                            var v53 float64 = 0.1
                            v54 := GenS10{F0: (v53 == v53), F1: 983, F2: 664}
                            _ = v52
                            _ = v53
                            _ = v54
                        }
                    }
                    if true {
                        var v55 int = 128
                        if false {
                            v49 = v49
                        } else {
                            v55 = fi48
                            v56 := GenS11{F0: 65535, F1: 596, F2: -824}
                            if (true || (false || (42 >= 437))) {
                                v2 = (v2 + ("" + (v2 + v2)))
                                var v57 uint = 65535
                                var v58 float32 = (-2.25 - -2.25)
                                _ = v57
                                _ = v58
                            } else {
                                fmt.Println(v49, v2)
                                v49 = (^(v49))
                                v3[(v55 % ((128) | 1))] = "hello"
                                var v59 bool = ((2147483647 == fi47) && ("\n\t" <= v2))
                                if (v59 || (-287 > -32768)) {
                                    var v60 int = fi47
                                    v1 = append(v1, v2)
                                    v49 = (v49 % ((127) | 1))
                                    _ = v60
                                }
                                _ = v59
                            }
                            _ = v56
                        }
                        _ = v55
                    } else {
                        if false {
                            var v61 int = (fi48 * -9223372036854775808)
                            v62 := GenS12{F0: 2147483647, F1: 42, F2: 42}
                            _ = v61
                            _ = v62
                        } else {
                            var v63 int32 = 2147483647
                            if false {
                                v64 := v3[27]
                                _ = v64
                            } else {
                                v65 := GenS13{F0: (v63 * 1), F1: 100}
                                v66 := GenS13{F0: (v63 - v63), F1: genFn1(false, 32767)}
                                _ = v65
                                _ = v66
                            }
                            _ = v63
                        }
                    }
                    v67 := GenS3{F0: 364}
                    _ = v67
                } else {
                    fmt.Println(v49)
                    if ((2 >= 481) && (888 >= 1)) {
                        v68 := GenS8{F0: (129 == 1), F1: 725, F2: 2}
                        switch v2 {
                        case "\n\t":
                            fmt.Println(fi48, v49)
                        case "":
                            var v69 int32 = -2147483648
                            _ = v69
                        default:
                            v2 = v2
                        }
                        _ = v68
                    }
                }
            }
            v2 = ((v2 + v2) + v2)
            switch -377 {
            case 4294967295:
                v2 = "foo bar"
                var v70 float32 = 1e10
                _ = v70
            case 677:
                switch (fi48 ^ (fi48 & (fi47 - fi48))) {
                case -2147483648:
                    var v71 uint32 = 2
                    _ = v71
                case 1:
                    v49 = genFn0(18446744073709551615, (fi48 >> ((608) & 63)), -2147483648)
                }
            default:
                fmt.Println(v49, v2, fi48)
            }
            _ = fi48
            _ = v49
        }
        _ = fi47
    }
    var v72 int64 = -9223372036854775808
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v45
    _ = v46
    _ = v72
}
