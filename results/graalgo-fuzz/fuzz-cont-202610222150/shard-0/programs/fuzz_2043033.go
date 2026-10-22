package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint32; F1 uint32; F2 int32 }
type GenS6 struct { F0 uint8; F1 bool }
func genFn0(p0 bool) uint16 {
    var v6 uint8 = 62
    p0 = false
    v6 = 255
    fmt.Println(v6, v6)
    switch "" {
    case "\n\t":
        for fi7 := 0; fi7 < 14; fi7++ {
            p0 = (!((v6 ^ 255) == (v6 ^ v6)))
            var v8 int32 = 1
            v9 := GenS1{F0: -1, F1: (3.14159 == 0.0), F2: "\n\t"}
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    _ = p0
    _ = v6
    return 1
}
func genFn1() uint32 {
    if true {
        for fi10 := 0; fi10 < 9; fi10++ {
            var v11 uint16 = 32767
            v12 := GenS1{F0: 42, F1: (false && false), F2: (("\n\t" + "a") + "")}
            _ = fi10
            _ = v11
            _ = v12
        }
    } else {
        for fi13 := 0; fi13 < 11; fi13++ {
            v14 := GenS2{F0: (("" + "hello") + "")}
            var v15 int = 9223372036854775807
            _ = fi13
            _ = v14
            _ = v15
        }
        v16 := GenS1{F0: 162, F1: (true && false), F2: (("hello" + ("\n\t" + "hello")) + ("hello" + "\n\t"))}
        v17 := GenS2{F0: ("\n\t" + ("a" + ""))}
        _ = v16
        _ = v17
    }
    var v18 int32 = 899
    var v19 string = "foo bar"
    _ = v18
    _ = v19
    return 42
}

func main() {
    v1 := []int{573, 1, 0}
    var v2 int64 = 42
    v3 := []int{-432, -1}
    var v4 uint32 = 2
    var v5 uint16 = 255
    if (true || false) {
        if (42 >= 2) {
            fmt.Println(v2, v5)
            for fi20 := 0; fi20 < 2; fi20++ {
                var v21 uint32 = v4
                _ = fi20
                _ = v21
            }
            var v22 float64 = -2.25
            _ = v22
        } else {
            var v23 uint16 = (v5 ^ v5)
            if false {
                fmt.Println(v5, v23)
            } else {
                var v24 int32 = 1
                switch "hello" {
                case "\n\t":
                    var v25 int64 = v2
                    _ = v25
                case "hello":
                    var v26 int = 1
                    _ = v26
                case "a":
                    switch (("a" + ("hello" + (("hello" + "\n\t") + "foo bar"))) + (("" + "foo bar") + "a")) {
                    case "a":
                        switch -2147483648 {
                        case 1:
                            for fi27 := 0; fi27 < 3; fi27++ {
                                for fi28 := 0; fi28 < 7; fi28++ {
                                    fmt.Println(fi28, v24)
                                    if len(v3) > 0 {
                                        v3[0] = (fi27 >> ((4294967295) & 63))
                                    }
                                    _ = fi28
                                }
                                var v29 uint = 304
                                v23 += 42
                                _ = fi27
                                _ = v29
                            }
                        }
                    case "foo bar":
                        var v30 uint32 = 2147483648
                        _ = v30
                    default:
                        v31 := GenS3{F0: (752 != 9223372036854775807), F1: int8(v4)}
                        _ = v31
                    }
                default:
                    for fi32 := 0; fi32 < 17; fi32++ {
                        switch (-(fi32)) {
                        case 127:
                            for fi33 := 0; fi33 < 17; fi33++ {
                                if (!true) {
                                    v2 -= int64(v23)
                                    var v34 int8 = 53
                                    v35 := GenS2{F0: ("a" + "foo bar")}
                                    _ = v34
                                    _ = v35
                                }
                                v4 = v4
                                _ = fi33
                            }
                        default:
                            for fi36 := 0; fi36 < 4; fi36++ {
                                fmt.Println(v23, v23)
                                switch 4294967295 {
                                case -9223372036854775808:
                                    _ = len(v3)
                                case 0:
                                    switch "" {
                                    case "hello":
                                        v23 += (v5 + v23)
                                        if ("hello" <= ("foo bar" + ("\n\t" + ("" + (("\n\t" + ("" + "a")) + "\n\t"))))) {
                                            fmt.Println(fi32, v4)
                                            switch (fi32 ^ (fi36 ^ 292)) {
                                            case -2147483648:
                                                if false {
                                                    if false {
                                                        v4 += 698
                                                        var v37 int64 = v2
                                                        v38 := GenS4{F0: (-(v24))}
                                                        _ = v37
                                                        _ = v38
                                                    }
                                                    v39 := GenS1{F0: -1, F1: false, F2: ("" + "")}
                                                    v23 = (v5 | v23)
                                                    _ = v39
                                                } else {
                                                    var v40 uint = uint(v23)
                                                    _ = v40
                                                }
                                            case 2147483647:
                                                var v41 int32 = -1
                                                _ = v41
                                            default:
                                                v42 := GenS1{F0: -1, F1: false, F2: ("a" + "a")}
                                                _ = v42
                                            }
                                            v4 = (v4 / (((v4 ^ (v4 + v4))) | 1))
                                        }
                                    default:
                                        if false {
                                            if (!false) {
                                                fmt.Println(fi32, v24, v5)
                                            } else {
                                                fmt.Println(fi36)
                                                switch "a" {
                                                case "foo bar":
                                                    var v43 float32 = 0.1
                                                    _ = v43
                                                case "\n\t":
                                                    v44 := GenS2{F0: "a"}
                                                    _ = v44
                                                case "":
                                                    var v45 int64 = (v2 * (v2 | v2))
                                                    _ = v45
                                                default:
                                                    var v46 string = "foo bar"
                                                    _ = v46
                                                }
                                                fmt.Println(v4)
                                                var v47 string = "\n\t"
                                                v5 -= (v23 ^ 0)
                                                var v48 string = v47
                                                v24 = (v24 | (^(v24)))
                                                _ = v47
                                                _ = v48
                                            }
                                        }
                                    }
                                }
                                _ = fi36
                            }
                        }
                        if true {
                            v23 += v5
                        } else {
                            v2 = (v2 >> ((18446744073709551615) & 63))
                            var v49 uint32 = 0
                            v50 := GenS1{F0: -532, F1: true, F2: (("a" + "\n\t") + ("\n\t" + "a"))}
                            _ = v49
                            _ = v50
                        }
                        _ = fi32
                    }
                }
                _ = v24
            }
            v51 := GenS1{F0: 1000, F1: true, F2: "a"}
            _ = v23
            _ = v51
        }
    } else {
        var v52 int = 244
        _ = v52
    }
    v3 = append(v3, -956)
    v5 = (v5 - v5)
    for fi53 := 0; fi53 < 20; fi53++ {
        v1 = append(v1, (fi53 | 0))
        for fi54 := 0; fi54 < 15; fi54++ {
            v5 = v5
            v5 -= genFn0(((v2 / (((v2 << ((2) & 63))) | 1)) > v2))
            var v55 uint16 = (v5 & v5)
            _ = fi54
            _ = v55
        }
        _ = fi53
    }
    if (1 != 476) {
        if ((v5 / ((v5) | 1)) < (v5 & 255)) {
            v56 := GenS2{F0: ("\n\t" + "")}
            v57 := GenS5{F0: (v4 + (^(v4))), F1: (v4 % ((65535) | 1)), F2: 1}
            var v58 uint8 = 255
            _ = v56
            _ = v57
            _ = v58
        } else {
            v4 = 555
            v2 += -1
        }
        v59 := GenS2{F0: "hello"}
        if len(v1) > 0 {
            v1[0] = 127
        }
        _ = v59
    }
    fmt.Println(v4, v5)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v60 := GenS1{F0: -1, F1: ((!(0.1 > -2.25)) && false), F2: "\n\t"}
    v61 := GenS6{F0: 49, F1: (!((v2 % ((v2) | 1)) != 696))}
    v2 += 1
    if true {
        var v62 string = ""
        var v63 float32 = 0.0
        _ = v62
        _ = v63
    } else {
        v5 = (v5 - (v5 << ((2) & 63)))
        v64 := GenS3{F0: true, F1: 127}
        _ = v64
    }
    fmt.Println(v2, v2)
    v2 = (v2 - -9223372036854775808)
    for fi65 := 0; fi65 < 20; fi65++ {
        var v66 int16 = -551
        _ = fi65
        _ = v66
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v60
    _ = v61
}
