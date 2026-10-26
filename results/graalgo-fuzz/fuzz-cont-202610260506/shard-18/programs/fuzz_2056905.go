package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 uint8; F1 int8; F2 uint64 }
type GenS3 struct { F0 int; F1 int8; F2 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32 }
func genFn0(p0 uint64) uint32 {
    v6 := GenS1{F0: 2, F1: (("\n\t" + "hello") + (("foo bar" + "\n\t") + ""))}
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 int8 = 100
        for fi9 := 0; fi9 < 16; fi9++ {
            v10 := GenS1{F0: 255, F1: "a"}
            var v11 string = ("" + "")
            switch (fi7 * fi7) {
            case 42:
                v11 = ("hello" + v11)
            case 32768:
                var v12 float64 = 0.0
                _ = v12
            case -2147483648:
                fmt.Println(p0)
                v8 += (v8 ^ (v8 >> ((18446744073709551615) & 63)))
                p0 -= p0
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    if false {
        fmt.Println(p0)
        var v13 int64 = -61
        _ = v13
    } else {
        var v14 uint32 = 2
        _ = v14
    }
    p0 += p0
    _ = p0
    _ = v6
    return 2
}
func genFn1(p0 int16, p1 uint64, p2 uint64) uint32 {
    var v15 int = 439
    p1 = (p1 ^ p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 0
}

func main() {
    v1 := []string{"\n\t", "hello"}
    var v2 int = 2147483647
    v3 := make(map[string]int)
    var v4 uint8 = 219
    var v5 uint32 = 213
    var v16 int = v2
    v2 -= int(v4)
    fmt.Println(v4)
    v3[""] = v2
    if true {
        var v17 int64 = 1
        _ = v17
    } else {
        v2 = (^(v16))
        var v18 string = ""
        v19 := GenS1{F0: 65535, F1: "hello"}
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 10; fi20++ {
        var v21 int64 = -1
        for fi22 := 0; fi22 < 7; fi22++ {
            v4 += v4
            var v23 int64 = v21
            var v24 float32 = 1e10
            var v25 uint32 = (v5 / ((v5) | 1))
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
        }
        _ = fi20
        _ = v21
    }
    if true {
        if len(v1) > 0 {
            v1[0] = "a"
        }
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if true {
            v2 += 9223372036854775807
            if (true && true) {
                var v26 bool = false
                if false {
                    for fi27 := 0; fi27 < 6; fi27++ {
                        fmt.Println(v26)
                        if v26 {
                            v4 = 58
                            var v28 string = "foo bar"
                            v26 = (!(3.14159 != 1e10))
                            _ = v28
                        } else {
                            for fi29 := 0; fi29 < 4; fi29++ {
                                if v26 {
                                    switch (v16 & v2) {
                                    case 128:
                                        v5 += (v5 << ((42) & 63))
                                    case -2147483648:
                                        v16 = fi29
                                    case 1:
                                        var v30 uint = 42
                                        _ = v30
                                    }
                                    if false {
                                        var v31 uint16 = 32767
                                        v2 -= (fi27 % (((v16 >> ((396) & 63))) | 1))
                                        _ = v31
                                    } else {
                                        v3[("a" + (("" + ("\n\t" + "a")) + "foo bar"))] = (v2 / (((fi29 * int(v5))) | 1))
                                        var v32 float32 = float32((1e10 + float64(0.1)))
                                        _ = v32
                                    }
                                    v33 := v3[("\n\t" + "foo bar")]
                                    _ = v33
                                } else {
                                    if ((!v26) || (1e10 >= -2.25)) {
                                        fmt.Println(fi27)
                                        v1 = append(v1, "foo bar")
                                        var v34 uint16 = 1
                                        _ = v34
                                    }
                                    v2 = fi27
                                }
                                delete(v3, "a")
                                var v35 uint16 = 0
                                _ = fi29
                                _ = v35
                            }
                        }
                        v36 := GenS2{F0: 1, F1: -24, F2: 4294967295}
                        _ = fi27
                        _ = v36
                    }
                    switch (v16 | (v2 << ((292) & 63))) {
                    case 4294967295:
                        for fi37 := 0; fi37 < 5; fi37++ {
                            var v38 float64 = 1e10
                            var v39 float32 = 1e10
                            v2 += 32768
                            _ = fi37
                            _ = v38
                            _ = v39
                        }
                    }
                    for fi40 := 0; fi40 < 4; fi40++ {
                        if false {
                            if (v16 == (v16 - -230)) {
                                var v41 int = (v16 ^ 2147483647)
                                var v42 uint32 = v5
                                _ = v41
                                _ = v42
                            }
                            if len(v1) > 0 {
                                v1[0] = "foo bar"
                            }
                        }
                        _ = fi40
                    }
                }
                fmt.Println(v2)
                _ = v26
            }
        } else {
            if false {
                v2 -= v2
            } else {
                fmt.Println(v4, v16)
                var v43 float32 = 1.5
                var v44 bool = true
                _ = v43
                _ = v44
            }
        }
    }
    var v45 int = (v2 | v16)
    if (!false) {
        var v46 int = v16
        if (true && false) {
            if (v4 >= (v4 * v4)) {
                delete(v3, ("" + ""))
                v2 -= v46
                v47 := v3[""]
                _ = v47
            } else {
                v4 -= (v4 & 127)
                if (536 > v5) {
                    v4 += 42
                    switch v2 {
                    case 0:
                        v5 = v5
                    default:
                        if false {
                            var v48 int32 = 1
                            if (0 != 1) {
                                v45 += (-(v2))
                                v49 := GenS3{F0: v45, F1: 100, F2: 809}
                                v50 := GenS1{F0: 951, F1: "\n\t"}
                                _ = v49
                                _ = v50
                            }
                            _ = v48
                        }
                    }
                } else {
                    var v51 int64 = 0
                    var v52 bool = false
                    fmt.Println(v52)
                    _ = v51
                    _ = v52
                }
                for fi53 := 0; fi53 < 1; fi53++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    for fi54 := 0; fi54 < 8; fi54++ {
                        _ = len(v1)
                        _ = fi54
                    }
                    _ = fi53
                }
            }
        } else {
            if (v4 <= (v4 + uint8(v5))) {
                fmt.Println(v2, v2, v4)
                switch "" {
                case "":
                    switch "hello" {
                    case "\n\t":
                        if (-589 != 42) {
                            v16 -= v16
                        } else {
                            if true {
                                if true {
                                    v16 = (^(v45))
                                } else {
                                    var v55 bool = ((1 != 42) || false)
                                    if (1.5 != 3.14159) {
                                        fmt.Println(v16)
                                        v45 = (v16 << ((18446744073709551615) & 63))
                                    }
                                    v2 -= v16
                                    _ = v55
                                }
                                switch (v45 / (((v45 / ((v46) | 1))) | 1)) {
                                case 792:
                                    v56 := GenS2{F0: 128, F1: -1, F2: 779}
                                    _ = v56
                                case -986:
                                    var v57 uint16 = 32767
                                    _ = v57
                                default:
                                    v2 = v45
                                }
                            }
                        }
                    case "":
                        fmt.Println(v5, v46, v4)
                    case "hello":
                        v46 = v45
                    }
                case "\n\t":
                    if (4294967295 != v45) {
                        var v58 int8 = -1
                        var v59 uint64 = 42
                        v46 += v45
                        _ = v58
                        _ = v59
                    }
                case "foo bar":
                    var v60 int16 = 42
                    _ = v60
                default:
                    fmt.Println(v2, v46, v16)
                }
                v61 := GenS4{F0: ((!(!false)) && false)}
                _ = v61
            } else {
                v16 -= v46
                if ("foo bar" <= "") {
                    var v62 uint8 = (v4 + v4)
                    _ = v62
                } else {
                    var v63 bool = (!false)
                    v64 := GenS1{F0: 2, F1: "hello"}
                    for fi65 := 0; fi65 < 13; fi65++ {
                        v46 = v46
                        fmt.Println(v5, v2, v5)
                        _ = fi65
                    }
                    _ = v63
                    _ = v64
                }
                v66 := GenS2{F0: v4, F1: 42, F2: 1}
                _ = v66
            }
            v5 -= (v5 >> ((863) & 63))
        }
        _ = v46
    }
    switch 678 {
    case 9223372036854775807:
        v45 -= 1
    }
    v67 := v3[("\n\t" + ("\n\t" + (("a" + "") + ("" + "\n\t"))))]
    for fi68 := 0; fi68 < 8; fi68++ {
        v16 += (v45 - (-(v67)))
        switch ("" + "foo bar") {
        case "hello":
            if (2 == 255) {
                var v69 int8 = -128
                var v70 uint = 0
                v71 := GenS5{F0: (v5 / ((genFn1(1000, 216, 339)) | 1))}
                _ = v69
                _ = v70
                _ = v71
            }
        case "a":
            v67 = (v2 << ((4294967295) & 63))
        default:
            v5 += genFn0(316)
        }
        for fi72 := 0; fi72 < 15; fi72++ {
            v67 += -9223372036854775808
            _ = fi72
        }
        _ = fi68
    }
    for fi73 := 0; fi73 < 2; fi73++ {
        var v74 bool = ((0 > 65535) || ((v4 * v4) >= (v4 ^ 2)))
        _ = fi73
        _ = v74
    }
    for fi75 := 0; fi75 < 17; fi75++ {
        switch (v45 % ((v45) | 1)) {
        case 0:
            if true {
                fmt.Println(v4, v4, v67)
            }
        case 62:
            v76 := v3["foo bar"]
            _ = v76
        }
        v5 = (v5 >> ((42) & 63))
        v2 -= 9223372036854775807
        _ = fi75
    }
    if (-104 == -1) {
        var v77 int64 = -532
        fmt.Println(v16, v67)
        v4 = (v4 >> ((65535) & 63))
        _ = v77
    }
    fmt.Println(v5, v5)
    v16 = 127
    delete(v3, "")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v67
}
