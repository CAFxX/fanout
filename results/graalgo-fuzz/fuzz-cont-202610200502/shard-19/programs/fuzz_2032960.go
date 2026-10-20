package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int; F1 int; F2 uint }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 int8 }
func genFn0(p0 uint64) uint32 {
    if (("a" + "a") == (("foo bar" + "") + ((("" + "\n\t") + ("hello" + "")) + "a"))) {
        if false {
            if (249 <= 597) {
                for fi4 := 0; fi4 < 9; fi4++ {
                    var v5 float32 = float32((1.5 + -2.25))
                    var v6 uint64 = (^(p0))
                    _ = fi4
                    _ = v5
                    _ = v6
                }
                v7 := []int{0}
                for fi8 := 0; fi8 < 6; fi8++ {
                    var v9 int32 = 42
                    if len(v7) > 0 {
                        v7[0] = (fi8 % ((32768) | 1))
                    }
                    _ = fi8
                    _ = v9
                }
                _ = v7
            } else {
                v10 := GenS1{F0: 0, F1: 1000, F2: 294}
                for fi11 := 0; fi11 < 11; fi11++ {
                    p0 = p0
                    _ = fi11
                }
                var v12 uint8 = 127
                _ = v10
                _ = v12
            }
            v13 := []int{343, 6}
            _ = v13
        }
        p0 += p0
        if true {
            v14 := GenS2{F0: 719, F1: ("a" == ("foo bar" + "foo bar"))}
            if true {
                if (!false) {
                    v15 := []int{}
                    var v16 int16 = -1
                    _ = v15
                    _ = v16
                } else {
                    if (207 != 128) {
                        for fi17 := 0; fi17 < 20; fi17++ {
                            v18 := GenS2{F0: -35, F1: (!(true && false))}
                            v19 := GenS1{F0: 4294967295, F1: 32767, F2: 32767}
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                        switch 42 {
                        case 0:
                            p0 = (p0 << ((4294967295) & 63))
                        default:
                            p0 -= 4294967295
                        }
                    }
                }
                switch "" {
                case "foo bar":
                    fmt.Println(p0)
                case "":
                    var v20 uint = 818
                    _ = v20
                }
            }
            _ = v14
        } else {
            var v21 bool = (!(!false))
            var v22 int16 = 32767
            p0 = 4294967295
            _ = v21
            _ = v22
        }
    }
    p0 -= (^(p0))
    _ = p0
    return 2
}
func genFn1(p0 int64, p1 uint16) int64 {
    fmt.Println(p1, p0)
    var v23 int8 = 127
    _ = p0
    _ = p1
    _ = v23
    return -550
}

func main() {
    v1 := []string{}
    v2 := make(map[string]bool)
    var v3 float64 = ((-2.25 + 1e10) * 0.1)
    v3 = 1e10
    switch "\n\t" {
    case "foo bar":
        switch 2147483647 {
        case 0:
            v24 := GenS2{F0: 32767, F1: (42 < genFn0(18446744073709551615))}
            _ = v24
        case 1:
            v2["\n\t"] = false
        case 42:
            if true {
                if (-796 != 9223372036854775807) {
                    v25 := GenS1{F0: 715, F1: -1, F2: 42}
                    v3 = ((v3 + ((v3 + float64(float32(1e10))) / v3)) / (v3 - 1.5))
                    _ = v25
                }
                if false {
                    v26 := GenS1{F0: 2147483648, F1: 42, F2: 157}
                    _ = len(v2)
                    fmt.Println(v3)
                    _ = v26
                }
                v3 -= ((1.5 * ((v3 * v3) / v3)) + float64(3.14159))
            } else {
                for fi27 := 0; fi27 < 3; fi27++ {
                    for fi28 := 0; fi28 < 10; fi28++ {
                        fmt.Println(v3, fi27)
                        if (127 < 127) {
                            v3 += 0.1
                        }
                        _ = fi28
                    }
                    for fi29 := 0; fi29 < 4; fi29++ {
                        switch ("hello" + (("foo bar" + "foo bar") + ("" + "\n\t"))) {
                        case "":
                            var v30 uint64 = 42
                            v3 = v3
                            _ = v30
                        }
                        _ = fi29
                    }
                    var v31 bool = false
                    _ = fi27
                    _ = v31
                }
                delete(v2, (("a" + "a") + ""))
                v3 = v3
            }
        default:
            var v32 uint8 = 2
            _ = v32
        }
    case "hello":
        var v33 int32 = 0
        _ = v33
    case "a":
        if (false || false) {
            v34 := GenS3{F0: 65535}
            var v35 uint32 = 2147483648
            _ = v34
            _ = v35
        } else {
            v2[("\n\t" + "\n\t")] = false
        }
    default:
        switch 32767 {
        case 128:
            fmt.Println(v3)
        case 400:
            v2["hello"] = (42 >= 664)
        case 32767:
            if (905 > -9223372036854775808) {
                v3 = (float64(-2.25) + 0.0)
                switch "hello" {
                case "":
                    v36 := GenS4{F0: 0}
                    _ = v36
                default:
                    if (1 > -1) {
                        for fi37 := 0; fi37 < 15; fi37++ {
                            v3 -= (0.1 + (-2.25 - v3))
                            _ = fi37
                        }
                        for fi38 := 0; fi38 < 13; fi38++ {
                            v39 := GenS5{F0: 255}
                            var v40 uint = 764
                            v41 := GenS6{F0: (^(fi38))}
                            _ = fi38
                            _ = v39
                            _ = v40
                            _ = v41
                        }
                        switch 9223372036854775807 {
                        case -944:
                            switch 9223372036854775807 {
                            case 1:
                                v42 := GenS4{F0: 602}
                                _ = v42
                            }
                        default:
                            switch "foo bar" {
                            case "hello":
                                for fi43 := 0; fi43 < 9; fi43++ {
                                    for fi44 := 0; fi44 < 12; fi44++ {
                                        v45 := GenS2{F0: 42, F1: false}
                                        v2[("a" + "hello")] = true
                                        var v46 uint64 = 42
                                        _ = fi44
                                        _ = v45
                                        _ = v46
                                    }
                                    var v47 uint8 = 0
                                    _ = len(v1)
                                    _ = fi43
                                    _ = v47
                                }
                            }
                        }
                    } else {
                        if (1 != 42) {
                            if len(v1) > 0 {
                                v1[0] = ("a" + (("foo bar" + "a") + (("\n\t" + "a") + "\n\t")))
                            }
                            v48 := GenS7{F0: 256, F1: 128, F2: 18446744073709551615}
                            _ = v48
                        } else {
                            for fi49 := 0; fi49 < 11; fi49++ {
                                fmt.Println(fi49, v3)
                                if true {
                                    for fi50 := 0; fi50 < 13; fi50++ {
                                        fmt.Println(fi50, v3, fi49)
                                        v51 := GenS2{F0: 42, F1: true}
                                        if true {
                                            var v52 uint32 = genFn0(100)
                                            var v53 uint64 = 0
                                            _ = v52
                                            _ = v53
                                        }
                                        _ = fi50
                                        _ = v51
                                    }
                                } else {
                                    var v54 float64 = float64(float32(((0.0 * v3) + (v3 + v3))))
                                    _ = v54
                                }
                                for fi55 := 0; fi55 < 3; fi55++ {
                                    delete(v2, "foo bar")
                                    var v56 int16 = 0
                                    switch "foo bar" {
                                    case "hello":
                                        var v57 uint8 = 2
                                        _ = v57
                                    case "\n\t":
                                        var v58 uint32 = 894
                                        _ = v58
                                    default:
                                        if false {
                                            var v59 uint64 = 18446744073709551615
                                            v3 = -2.25
                                            _ = v59
                                        }
                                    }
                                    _ = fi55
                                    _ = v56
                                }
                                _ = fi49
                            }
                        }
                        for fi60 := 0; fi60 < 7; fi60++ {
                            fmt.Println(fi60, fi60)
                            if true {
                                for fi61 := 0; fi61 < 12; fi61++ {
                                    var v62 int16 = -387
                                    if ((-802 <= 42) && true) {
                                        fmt.Println(v3, fi61)
                                        v63 := GenS1{F0: genFn0(52), F1: -720, F2: 32767}
                                        if (fi60 < 0) {
                                            v64 := GenS8{F0: 0}
                                            v3 += ((3.14159 - v3) * 0.0)
                                            v1 = append(v1, "")
                                            v62 -= (^(v62))
                                            _ = v64
                                        }
                                        _ = v63
                                    } else {
                                        var v65 uint = 65535
                                        v3 = (((v3 + v3) - (v3 + ((v3 / (1.5 / (-2.25 - 3.14159))) / 1.0))) + 1e10)
                                        v3 += v3
                                        v66 := GenS4{F0: v65}
                                        _ = v65
                                        _ = v66
                                    }
                                    v67 := GenS8{F0: -1}
                                    _ = fi61
                                    _ = v62
                                    _ = v67
                                }
                                var v68 uint8 = 39
                                v69 := GenS5{F0: 65535}
                                _ = v68
                                _ = v69
                            }
                            for fi70 := 0; fi70 < 9; fi70++ {
                                v71 := GenS9{F0: genFn0(0)}
                                v3 = 1e10
                                _ = fi70
                                _ = v71
                            }
                            _ = fi60
                        }
                        if false {
                            _ = len(v2)
                        }
                    }
                }
            }
        default:
            var v72 bool = (185 < 32767)
            _ = v72
        }
    }
    v73 := GenS6{F0: 9223372036854775807}
    v3 = v3
    v3 -= 0.1
    switch 128 {
    case 128:
        v74 := GenS9{F0: 275}
        _ = v74
    }
    var v75 string = ("a" + "foo bar")
    for fi76 := 0; fi76 < 6; fi76++ {
        v2[(v75 + "hello")] = false
        v75 = ""
        _ = fi76
    }
    delete(v2, ((v75 + ((v75 + ("foo bar" + (v75 + v75))) + (v75 + v75))) + (v75 + v75)))
    var v77 uint16 = 886
    switch "foo bar" {
    case "\n\t":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if (!(615 >= 2147483647)) {
        if true {
            v78 := GenS9{F0: 155}
            v77 += v77
            _ = v78
        } else {
            fmt.Println(v77, v75)
            v79 := GenS5{F0: 736}
            var v80 int8 = 100
            _ = v79
            _ = v80
        }
        for fi81 := 0; fi81 < 3; fi81++ {
            v82 := GenS1{F0: 2147483647, F1: int16(v77), F2: v77}
            _ = fi81
            _ = v82
        }
    }
    for fi83 := 0; fi83 < 17; fi83++ {
        v84 := GenS2{F0: 1, F1: true}
        v85 := GenS10{F0: -128}
        v75 = "a"
        _ = fi83
        _ = v84
        _ = v85
    }
    var v86 int32 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v75:", v75)
    fmt.Println("v77:", v77)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v73
    _ = v75
    _ = v77
    _ = v86
}
