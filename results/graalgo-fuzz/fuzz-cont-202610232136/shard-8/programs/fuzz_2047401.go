package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64; F1 uint }
type GenS5 struct { F0 int64; F1 bool }
type GenS6 struct { F0 string; F1 int; F2 string }
type GenS7 struct { F0 uint }
func genFn0(p0 uint32, p1 uint32, p2 string) int16 {
    p2 = (p2 + p2)
    v4 := []int{-672, 4294967295, 644, 32767}
    p0 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 1000
}

func main() {
    var v1 int64 = 42
    var v2 int32 = 0
    v3 := []string{""}
    if ((true && false) || (65535 >= 18446744073709551615)) {
        for fi5 := 0; fi5 < 17; fi5++ {
            v6 := GenS1{F0: fi5}
            _ = fi5
            _ = v6
        }
        v3 = append(v3, "hello")
    }
    switch "foo bar" {
    case "a":
        _ = len(v3)
    case "hello":
        for fi7 := 0; fi7 < 16; fi7++ {
            v2 = v2
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi7
        }
    }
    var v8 uint = 1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v9 uint = (v8 | 65535)
    v9 = v8
    if (true || true) {
        var v10 int = 9223372036854775807
        for fi11 := 0; fi11 < 10; fi11++ {
            v12 := GenS1{F0: 32768}
            _ = fi11
            _ = v12
        }
        var v13 int16 = 32767
        _ = v10
        _ = v13
    }
    switch ("\n\t" + ((("hello" + "hello") + ((("foo bar" + "a") + "\n\t") + "hello")) + ("foo bar" + "foo bar"))) {
    case "":
        v14 := GenS2{F0: 65535, F1: 255}
        _ = v14
    default:
        if (2 != 42) {
            for fi15 := 0; fi15 < 8; fi15++ {
                for fi16 := 0; fi16 < 18; fi16++ {
                    fmt.Println(v9, v1)
                    var v17 uint = (v8 | (v9 ^ (v9 * (^(v8)))))
                    _ = fi16
                    _ = v17
                }
                _ = fi15
            }
        } else {
            var v18 int64 = v1
            switch "hello" {
            case "":
                v8 += (v8 & (^(v8)))
            case "foo bar":
                switch 437 {
                case 0:
                    if (true || (v2 != -1)) {
                        var v19 uint32 = 65535
                        _ = v19
                    }
                case 32767:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                default:
                    if (false || (false && true)) {
                        for fi20 := 0; fi20 < 3; fi20++ {
                            var v21 bool = ((v18 % ((0) | 1)) > v1)
                            var v22 bool = (!(((v9 << (((v9 - 18446744073709551615)) & 63)) != (v8 + v9)) && ((v1 * v1) == (v1 - v1))))
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                        if false {
                            var v23 uint8 = 1
                            _ = v23
                        } else {
                            fmt.Println(v18, v1, v9)
                            v24 := GenS3{F0: 1}
                            var v25 uint16 = 197
                            _ = v24
                            _ = v25
                        }
                    }
                }
            default:
                if false {
                    switch 1 {
                    case 4294967295:
                        v18 = (v1 & (v1 * v18))
                    case 219:
                        for fi26 := 0; fi26 < 5; fi26++ {
                            v2 += (v2 * -827)
                            _ = fi26
                        }
                    default:
                        for fi27 := 0; fi27 < 11; fi27++ {
                            v3 = append(v3, ((("\n\t" + "a") + ("foo bar" + "")) + "foo bar"))
                            if true {
                                v8 = v8
                                if (true || false) {
                                    fmt.Println(fi27)
                                } else {
                                    if (0.0 > 0.0) {
                                        var v28 int16 = 1000
                                        if (v28 == (-(v28))) {
                                            var v29 int8 = 100
                                            _ = v29
                                        }
                                        var v30 int = fi27
                                        _ = v28
                                        _ = v30
                                    }
                                    var v31 uint = (^(v8))
                                    _ = v31
                                }
                            }
                            v3 = append(v3, "hello")
                            _ = fi27
                        }
                    }
                }
            }
            if (42 != 127) {
                var v32 int16 = 312
                if true {
                    var v33 string = ("" + "\n\t")
                    _ = v33
                }
                if (!false) {
                    v34 := GenS3{F0: 255}
                    if true {
                        v9 = v9
                        for fi35 := 0; fi35 < 7; fi35++ {
                            v36 := GenS1{F0: 42}
                            if true {
                                for fi37 := 0; fi37 < 7; fi37++ {
                                    var v38 int16 = 42
                                    _ = fi37
                                    _ = v38
                                }
                            }
                            var v39 bool = false
                            _ = fi35
                            _ = v36
                            _ = v39
                        }
                        for fi40 := 0; fi40 < 14; fi40++ {
                            v9 = (v8 >> ((0) & 63))
                            _ = fi40
                        }
                    }
                    _ = v34
                } else {
                    fmt.Println(v1, v9)
                }
                _ = v32
            } else {
                var v41 uint8 = 130
                v41 += v41
                for fi42 := 0; fi42 < 3; fi42++ {
                    var v43 uint64 = 42
                    if (("" + ("a" + ("foo bar" + ""))) > "a") {
                        var v44 uint = v8
                        v43 -= (v43 % ((v43) | 1))
                        for fi45 := 0; fi45 < 8; fi45++ {
                            switch -9223372036854775808 {
                            case 4294967295:
                                fmt.Println(fi42, v18, fi45)
                            case 32768:
                                var v46 uint64 = v43
                                _ = v46
                            default:
                                v44 = (v44 & v8)
                            }
                            var v47 bool = ((!false) || (-2.25 >= 1.5))
                            if (898 <= 644) {
                                v48 := GenS1{F0: (fi45 / ((fi45) | 1))}
                                v49 := GenS1{F0: fi45}
                                v44 = (v44 & v44)
                                _ = v48
                                _ = v49
                            } else {
                                var v50 bool = v47
                                _ = v50
                            }
                            _ = fi45
                            _ = v47
                        }
                        _ = v44
                    }
                    _ = fi42
                    _ = v43
                }
                _ = v41
            }
            _ = v18
        }
    }
    for fi51 := 0; fi51 < 11; fi51++ {
        if ((-2.25 + 0.0) > 0.0) {
            if false {
                v52 := GenS1{F0: fi51}
                if false {
                    v53 := GenS3{F0: 32767}
                    _ = v53
                } else {
                    fmt.Println(v9, v2)
                    v3 = append(v3, "a")
                }
                _ = len(v3)
                _ = v52
            }
            for fi54 := 0; fi54 < 12; fi54++ {
                if (778 <= 4294967295) {
                    if (2147483647 != 731) {
                        var v55 int = (fi54 << ((65535) & 63))
                        if false {
                            if ((0.1 >= 3.14159) || true) {
                                v1 += (v1 * v1)
                            } else {
                                v56 := GenS1{F0: fi54}
                                _ = v56
                            }
                            var v57 int = (v55 % (((fi54 & v55)) | 1))
                            _ = v57
                        } else {
                            v1 = 0
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v58 := GenS4{F0: (^(v1)), F1: v8}
                            _ = v58
                        }
                        _ = v55
                    }
                    fmt.Println(v1)
                }
                switch int(v1) {
                case -255:
                    fmt.Println(v8, fi51, fi54)
                    if false {
                        if (32767 != 42) {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v59 := GenS1{F0: (fi54 - fi54)}
                            v1 += v1
                            _ = v59
                        } else {
                            var v60 float64 = -2.25
                            var v61 uint = 311
                            _ = v60
                            _ = v61
                        }
                        var v62 uint16 = 2
                        var v63 bool = (!(false && true))
                        _ = v62
                        _ = v63
                    } else {
                        var v64 uint32 = 2
                        _ = v64
                    }
                    _ = len(v3)
                    if len(v3) > 0 {
                        v3[0] = "a"
                    }
                default:
                    v1 += 9223372036854775807
                }
                v65 := GenS5{F0: v1, F1: ((((3.14159 / 1e10) - float64((float32(3.14159) + 1e10))) - (3.14159 * -2.25)) == 1.5)}
                switch 127 {
                case -881:
                    v3 = append(v3, ("hello" + "foo bar"))
                case 32767:
                    if len(v3) > 0 {
                        v3[0] = ("hello" + "foo bar")
                    }
                case -2147483648:
                    var v66 uint = (v8 ^ (v8 - (v8 << ((v9) & 63))))
                    _ = v66
                default:
                    v9 -= (^(v8))
                }
                var v67 uint32 = 65535
                v67 = (v67 & v67)
                if false {
                    v2 = 2147483647
                } else {
                    v68 := GenS5{F0: v1, F1: (fi54 == -596)}
                    switch 833 {
                    case 680:
                        if false {
                            v69 := GenS6{F0: "a", F1: 289, F2: "foo bar"}
                            var v70 uint = (^(v8))
                            _ = v69
                            _ = v70
                        }
                    case -1:
                        if false {
                            var v71 int32 = (v2 & int32(v67))
                            var v72 uint32 = (v67 | v67)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v71
                            _ = v72
                        }
                    }
                    _ = v68
                }
                var v73 int64 = -575
                var v74 float32 = 1.5
                v3 = append(v3, ("hello" + ("\n\t" + ("" + "hello"))))
                _ = fi54
                _ = v65
                _ = v67
                _ = v73
                _ = v74
            }
            fmt.Println(fi51, v1)
        }
        v75 := GenS4{F0: v1, F1: (v9 ^ v9)}
        var v76 string = (((("a" + "hello") + ("a" + "\n\t")) + "a") + "a")
        _ = fi51
        _ = v75
        _ = v76
    }
    v77 := GenS7{F0: v9}
    fmt.Println(v8)
    var v78 uint32 = 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v78:", v78)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v77
    _ = v78
    _ = v8
    _ = v9
}
