package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int; F1 int16; F2 int16 }
type GenS6 struct { F0 uint; F1 uint64; F2 uint16 }
type GenS7 struct { F0 uint64; F1 uint8; F2 uint }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]int)
    var v3 uint32 = 420
    var v4 uint32 = 65535
    v5 := []string{("" + ("foo bar" + ""))}
    if ((true || true) && false) {
        if (!(true && false)) {
            v6 := GenS1{F0: 916, F1: -2147483648}
            for fi7 := 0; fi7 < 2; fi7++ {
                switch "\n\t" {
                case "":
                    v3 = (^(v4))
                default:
                    var v8 int32 = -1
                    _ = v8
                }
                _ = len(v5)
                _ = fi7
            }
            switch 128 {
            case 32768:
                var v9 int = 32768
                _ = v9
            default:
                for fi10 := 0; fi10 < 4; fi10++ {
                    for fi11 := 0; fi11 < 17; fi11++ {
                        v12 := GenS2{F0: 523, F1: "hello"}
                        var v13 uint32 = v4
                        var v14 int64 = -9223372036854775808
                        v14 += 1
                        if false {
                            switch "hello" {
                            case "":
                                var v15 string = "a"
                                _ = v15
                            default:
                                if true {
                                    var v16 string = ("\n\t" + "a")
                                    _ = v16
                                } else {
                                    _ = len(v5)
                                    fmt.Println(fi11, fi10)
                                    v3 += (v4 * (v13 * 4294967295))
                                }
                            }
                        }
                        _ = fi11
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                    v17 := v2[(v4 | (^(v3)))]
                    _ = fi10
                    _ = v17
                }
            }
            _ = v6
        } else {
            v3 = (v3 & 2147483647)
            switch "hello" {
            case "\n\t":
                v4 = 2
            case "hello":
                if false {
                    if (-128 >= 1) {
                        v4 = (v3 % (((^(v4))) | 1))
                    }
                    switch "a" {
                    case "foo bar":
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                    v4 += v3
                } else {
                    v18 := GenS1{F0: 2147483647, F1: 0}
                    switch -65 {
                    case 32767:
                        if len(v5) > 0 {
                            v5[0] = ("\n\t" + "")
                        }
                    case -531:
                        fmt.Println(v4)
                    }
                    _ = v18
                }
            case "":
                if (-2.25 >= -2.25) {
                    for fi19 := 0; fi19 < 18; fi19++ {
                        v2[847] = (fi19 >> ((4294967295) & 63))
                        _ = fi19
                    }
                    v4 = (v3 >> ((924) & 63))
                    if true {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    } else {
                        v3 -= (v4 & (v3 % ((v4) | 1)))
                        switch 128 {
                        case 457:
                            v20 := v2[(v4 * 1)]
                            _ = v20
                        case 88:
                            _ = len(v5)
                        default:
                            v21 := GenS2{F0: -1, F1: "a"}
                            _ = v21
                        }
                        _ = len(v5)
                    }
                }
            default:
                if (1.5 > 0.0) {
                    v4 -= 997
                    if (1000 < 0) {
                        if false {
                            var v22 uint32 = v4
                            _ = v22
                        } else {
                            for fi23 := 0; fi23 < 3; fi23++ {
                                v5 = append(v5, ("foo bar" + ""))
                                _ = fi23
                            }
                        }
                        if true {
                            var v24 int32 = -1
                            _ = v24
                        } else {
                            var v25 uint16 = 65535
                            if ((-296 < 2147483647) && false) {
                                v3 += v4
                                if true {
                                    v25 = v25
                                }
                                v25 -= 2
                            } else {
                                v2[(v3 >> ((427) & 63))] = -457
                            }
                            switch ((("" + "hello") + "hello") + ("foo bar" + ("" + "foo bar"))) {
                            case "hello":
                                if false {
                                    v5 = append(v5, (("\n\t" + "a") + ("a" + "")))
                                    var v26 int32 = 830
                                    v4 = v4
                                    _ = v26
                                } else {
                                    var v27 int = 32767
                                    for fi28 := 0; fi28 < 11; fi28++ {
                                        _ = len(v5)
                                        v27 += (fi28 & (-(v27)))
                                        _ = fi28
                                    }
                                    _ = v27
                                }
                            case "":
                                for fi29 := 0; fi29 < 11; fi29++ {
                                    var v30 bool = true
                                    if (!v30) {
                                        if (((!((!(!((-(fi29)) < fi29))) || (!(false || false)))) && ((-(fi29)) == fi29)) || (0.0 > float32(1.5))) {
                                            var v31 string = "a"
                                            for fi32 := 0; fi32 < 14; fi32++ {
                                                var v33 uint32 = 734
                                                var v34 uint64 = 4294967295
                                                _ = fi32
                                                _ = v33
                                                _ = v34
                                            }
                                            v2[(v3 * v4)] = (fi29 >> ((947) & 63))
                                            _ = v31
                                        }
                                    }
                                    for fi35 := 0; fi35 < 15; fi35++ {
                                        fmt.Println(v30, fi29, fi29)
                                        var v36 float32 = (3.14159 / 1e10)
                                        var v37 int16 = 321
                                        v38 := GenS3{F0: 1}
                                        _ = fi35
                                        _ = v36
                                        _ = v37
                                        _ = v38
                                    }
                                    _ = fi29
                                    _ = v30
                                }
                            }
                            _ = v25
                        }
                        v4 -= (v3 / ((0) | 1))
                    } else {
                        var v39 uint64 = 2
                        if true {
                            var v40 uint16 = 148
                            v39 = (v39 ^ 2)
                            v41 := GenS4{F0: 1}
                            _ = v40
                            _ = v41
                        }
                        _ = v39
                    }
                    for fi42 := 0; fi42 < 16; fi42++ {
                        if true {
                            v3 += (v4 % ((v3) | 1))
                            _ = len(v2)
                            var v43 int = (fi42 & (fi42 % (((fi42 / (((fi42 / ((fi42) | 1))) | 1))) | 1)))
                            _ = v43
                        }
                        v44 := GenS5{F0: (fi42 ^ fi42), F1: 1, F2: -599}
                        _ = fi42
                        _ = v44
                    }
                }
            }
            var v45 uint16 = 255
            _ = v45
        }
    }
    var v46 uint16 = 65535
    v47 := GenS6{F0: 102, F1: 0, F2: (v46 - (v46 / ((v46) | 1)))}
    var v48 uint = 18446744073709551615
    var v49 int32 = -2147483648
    fmt.Println(v3, v48)
    v50 := GenS2{F0: v49, F1: "foo bar"}
    var v51 string = "a"
    v52 := GenS7{F0: 804, F1: 0, F2: (v48 << (((v48 ^ (v48 + (v48 | (^(v48)))))) & 63))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
}
