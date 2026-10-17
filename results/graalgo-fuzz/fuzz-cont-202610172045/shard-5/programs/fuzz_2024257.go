package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint64; F1 uint16 }
type GenS7 struct { F0 uint32; F1 int8 }
type GenS8 struct { F0 uint64; F1 uint16 }
type GenS9 struct { F0 uint32; F1 bool }
func genFn0(p0 uint, p1 int, p2 uint16) int {
    fmt.Println(p0, p0, p0)
    v4 := []int{p1}
    switch ("hello" + "a") {
    case "":
        var v5 string = "foo bar"
        _ = v5
    case "hello":
        var v6 int32 = 2147483647
        _ = v6
    case "\n\t":
        var v7 uint32 = 0
        _ = v7
    }
    p1 += (p1 ^ 42)
    v8 := GenS1{F0: 65535, F1: 972}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v8
    return p1
}
func genFn1(p0 int16, p1 int8) uint {
    for fi9 := 0; fi9 < 12; fi9++ {
        p0 -= (p0 & p0)
        var v10 float64 = (float64((1e10 + 1e10)) / ((-2.25 / 0.1) - 1.5))
        _ = fi9
        _ = v10
    }
    fmt.Println(p0, p1)
    var v11 float32 = ((0.0 / 1.0) + 1.5)
    fmt.Println(p1, v11)
    _ = p0
    _ = p1
    _ = v11
    return 215
}

func main() {
    var v1 uint16 = 255
    v2 := []string{"\n\t", ("hello" + "\n\t"), "", ("foo bar" + ("a" + "\n\t"))}
    var v3 uint64 = 602
    _ = len(v2)
    var v12 float32 = 3.14159
    v13 := GenS2{F0: 127}
    v3 = v3
    if true {
        if ((826 >= 2) && ((v3 > v3) || false)) {
            if (!true) {
                switch ("hello" + "") {
                case "foo bar":
                    var v14 int64 = 1
                    _ = v14
                case "hello":
                    v1 = v1
                }
            }
            var v15 float32 = (v12 + (v12 - float32(-2.25)))
            _ = v15
        } else {
            for fi16 := 0; fi16 < 7; fi16++ {
                v12 += v12
                v1 += (^(v1))
                _ = fi16
            }
            v17 := GenS1{F0: 61, F1: (v3 * (v3 / ((1) | 1)))}
            for fi18 := 0; fi18 < 8; fi18++ {
                v19 := GenS1{F0: 263, F1: v3}
                if (100 > 0) {
                    v20 := GenS3{F0: -18}
                    v21 := GenS4{F0: -32768, F1: 1000}
                    _ = v20
                    _ = v21
                }
                var v22 uint16 = v1
                _ = fi18
                _ = v19
                _ = v22
            }
            _ = v17
        }
    } else {
        for fi23 := 0; fi23 < 10; fi23++ {
            v24 := GenS1{F0: 2, F1: v3}
            v12 -= (-2.25 - v12)
            var v25 int64 = 1
            _ = fi23
            _ = v24
            _ = v25
        }
        if false {
            for fi26 := 0; fi26 < 12; fi26++ {
                if true {
                    var v27 int32 = 1
                    v28 := GenS4{F0: 32767, F1: -1}
                    for fi29 := 0; fi29 < 11; fi29++ {
                        if true {
                            v1 -= (v1 - 65535)
                            v1 = v1
                            var v30 uint16 = (v1 - v1)
                            _ = v30
                        } else {
                            v31 := GenS5{F0: 1}
                            v32 := GenS5{F0: 0}
                            _ = v31
                            _ = v32
                        }
                        switch "foo bar" {
                        case "a":
                            v3 -= v3
                            var v33 float32 = v12
                            var v34 uint64 = (v3 / ((0) | 1))
                            _ = v33
                            _ = v34
                        case "hello":
                            if (127 == 92) {
                                v3 -= (v3 | v3)
                                v3 = 0
                                v2 = append(v2, ((("hello" + "foo bar") + ("foo bar" + "foo bar")) + "foo bar"))
                                v3 += (v3 % ((v3) | 1))
                                var v35 uint16 = (v1 << ((839) & 63))
                                _ = v35
                            }
                            v27 = -930
                            var v36 uint8 = 42
                            _ = v36
                        case "\n\t":
                            v1 = (v1 % ((v1) | 1))
                        }
                        _ = fi29
                    }
                    _ = v27
                    _ = v28
                }
                _ = fi26
            }
        }
        var v37 uint64 = v3
        _ = v37
    }
    v2 = append(v2, "a")
    v38 := GenS3{F0: 127}
    var v39 uint = genFn1(-868, 42)
    v1 -= v1
    switch ("hello" + (("" + "") + "hello")) {
    case "foo bar":
        var v40 bool = ((0.1 / 0.1) > 0.0)
        _ = v40
    case "a":
        v41 := GenS6{F0: (v3 - v3), F1: v1}
        _ = v41
    case "\n\t":
        if false {
            fmt.Println(v3)
        } else {
            if false {
                v42 := GenS7{F0: 1, F1: 127}
                var v43 uint32 = 1
                _ = v42
                _ = v43
            } else {
                v2 = append(v2, "a")
                switch ("a" + "hello") {
                case "\n\t":
                    var v44 int16 = -32768
                    _ = v44
                }
            }
        }
    }
    v45 := GenS7{F0: 2147483647, F1: 0}
    v1 += v1
    v39 = v39
    if true {
        var v46 int64 = 42
        if true {
            for fi47 := 0; fi47 < 13; fi47++ {
                v3 += 42
                var v48 int = genFn0(253, fi47, (v1 | v1))
                if false {
                    fmt.Println(v39, v48, v39)
                    if (300 != 1000) {
                        v49 := GenS1{F0: 2, F1: (v3 % ((4294967295) | 1))}
                        v48 -= fi47
                        var v50 int16 = 743
                        _ = v49
                        _ = v50
                    }
                    _ = len(v2)
                } else {
                    if (-14 >= -1) {
                        if false {
                            if (!true) {
                                var v51 int = genFn0((v39 + v39), fi47, (v1 & (v1 & 1)))
                                _ = v51
                            } else {
                                v48 += (v48 + (^(v48)))
                            }
                            if (!(!(0 <= 1))) {
                                var v52 int32 = -2147483648
                                _ = len(v2)
                                _ = v52
                            } else {
                                v1 += (v1 ^ 2)
                                _ = len(v2)
                                v53 := GenS2{F0: 115}
                                _ = v53
                            }
                            var v54 float32 = v12
                            _ = v54
                        }
                        v55 := GenS2{F0: 127}
                        if false {
                            if len(v2) > 0 {
                                v2[0] = "hello"
                            }
                            for fi56 := 0; fi56 < 5; fi56++ {
                                var v57 bool = true
                                switch ("hello" + "") {
                                case "hello":
                                    switch genFn0((^(v39)), -1, 65535) {
                                    case 32767:
                                        v2 = append(v2, ("hello" + (("\n\t" + "hello") + "hello")))
                                    }
                                case "foo bar":
                                    var v58 float64 = 0.1
                                    _ = v58
                                case "a":
                                    if false {
                                        var v59 uint32 = 383
                                        fmt.Println(v1, fi47, v3)
                                        _ = v59
                                    }
                                default:
                                    v60 := GenS8{F0: v3, F1: 560}
                                    _ = v60
                                }
                                v61 := GenS6{F0: uint64(v39), F1: (v1 + (v1 | v1))}
                                _ = fi56
                                _ = v57
                                _ = v61
                            }
                            v46 -= v46
                        }
                        _ = v55
                    } else {
                        var v62 uint64 = v3
                        v63 := GenS2{F0: 1}
                        _ = v62
                        _ = v63
                    }
                    v64 := GenS9{F0: 42, F1: (fi47 != (fi47 - -188))}
                    v65 := GenS6{F0: (v3 % ((v3) | 1)), F1: v1}
                    _ = v64
                    _ = v65
                }
                _ = fi47
                _ = v48
            }
            fmt.Println(v3, v46)
        }
        fmt.Println(v12, v12, v12)
        _ = v46
    }
    var v66 string = ("hello" + "a")
    v67 := GenS6{F0: (v3 & v3), F1: v1}
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v45
    _ = v66
    _ = v67
}
