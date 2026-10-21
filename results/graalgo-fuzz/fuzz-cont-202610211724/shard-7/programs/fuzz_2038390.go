package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 uint16; F1 int16; F2 string }
type GenS5 struct { F0 uint64; F1 uint8 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 string; F1 int64; F2 uint16 }
type GenS8 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS9 struct { F0 int64 }
func genFn0(p0 uint16, p1 int8) uint {
    var v5 int16 = -1
    v6 := GenS1{F0: int64(p0), F1: (1 != 454), F2: 127}
    v7 := GenS1{F0: -1, F1: (0 < 4294967295), F2: -2147483648}
    var v8 uint64 = 4294967295
    p0 += 32767
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 65535
}
func genFn1() int {
    var v9 uint = 0
    fmt.Println(v9)
    switch 32767 {
    case 32767:
        fmt.Println(v9)
    case 640:
        var v10 int32 = 335
        _ = v10
    }
    var v11 float32 = (0.1 - -2.25)
    for fi12 := 0; fi12 < 17; fi12++ {
        v13 := []int{}
        var v14 float64 = -2.25
        if len(v13) > 0 {
            v13[0] = (fi12 << ((v9) & 63))
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    _ = v11
    _ = v9
    return 9223372036854775807
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]bool)
    var v4 uint8 = 0
    switch 42 {
    case 32768:
        v4 -= v4
    case 4294967295:
        v15 := []int{32767, 4294967295, 298, genFn1()}
        _ = v15
    default:
        v4 += v4
    }
    if false {
        v16 := GenS1{F0: 490, F1: (true || true), F2: -952}
        _ = v16
    } else {
        for fi17 := 0; fi17 < 4; fi17++ {
            v18 := []int{fi17, genFn1(), (fi17 >> ((1) & 63)), fi17}
            var v19 int16 = 1000
            if false {
                v19 -= v19
            } else {
                switch "hello" {
                case "foo bar":
                    v20 := GenS2{F0: (("\n\t" + "foo bar") + "foo bar"), F1: 0}
                    _ = v20
                case "":
                    v21 := GenS3{F0: 32767, F1: 194}
                    _ = v21
                }
                if (true && true) {
                    var v22 int8 = 107
                    if false {
                        fmt.Println(fi17)
                        if true {
                            fmt.Println(fi17)
                            var v23 string = "hello"
                            _ = v23
                        } else {
                            v22 = -128
                            var v24 uint8 = v4
                            for fi25 := 0; fi25 < 1; fi25++ {
                                if (((4294967295 < 481) && true) && (524 <= 630)) {
                                    fmt.Println(fi25, fi17, fi25)
                                    var v26 int32 = 42
                                    if (!false) {
                                        var v27 int64 = 1
                                        _ = v27
                                    } else {
                                        if ((v26 + (v26 ^ -2147483648)) == (-(v26))) {
                                            var v28 uint32 = 2147483647
                                            var v29 int64 = 42
                                            var v30 float32 = (0.1 - -2.25)
                                            if ((v29 + (v29 << ((genFn0(2, (^(v22)))) & 63))) < v29) {
                                                v31 := GenS4{F0: 42, F1: 42, F2: ("a" + "")}
                                                _ = v31
                                            } else {
                                                var v32 bool = true
                                                v4 += 2
                                                _ = v32
                                            }
                                            _ = v28
                                            _ = v29
                                            _ = v30
                                        } else {
                                            _ = len(v1)
                                            v19 += -32768
                                        }
                                        _ = len(v18)
                                        v19 -= int16(v26)
                                    }
                                    v4 = (v4 * (v24 + v4))
                                    v26 -= -1
                                    _ = v26
                                }
                                if len(v18) >= 2 {
                                    v18 = v18[1:len(v18)]
                                }
                                var v33 uint64 = 4294967295
                                _ = fi25
                                _ = v33
                            }
                            _ = v24
                        }
                        v34 := GenS2{F0: "", F1: 4294967295}
                        _ = v34
                    } else {
                        v35 := GenS5{F0: 1, F1: v4}
                        if ((float32(3.14159) / (1.5 / 1.5)) <= -2.25) {
                            v22 = (v22 * 100)
                        }
                        _ = v35
                    }
                    _ = v22
                } else {
                    v4 = v4
                }
                var v36 uint64 = 1
                _ = v36
            }
            _ = fi17
            _ = v18
            _ = v19
        }
        v37 := GenS4{F0: 999, F1: -36, F2: "\n\t"}
        _ = v37
    }
    v38 := GenS6{F0: uint64(v4)}
    if false {
        if (18446744073709551615 < 0) {
            for fi39 := 0; fi39 < 13; fi39++ {
                for fi40 := 0; fi40 < 4; fi40++ {
                    v41 := GenS4{F0: 0, F1: -1, F2: "hello"}
                    _ = fi40
                    _ = v41
                }
                _ = fi39
            }
            if true {
                fmt.Println(v4)
                var v42 bool = false
                switch -9223372036854775808 {
                case 32767:
                    v42 = v42
                case 9223372036854775807:
                    if ((4294967295 != 0) && (346 == 18446744073709551615)) {
                        if (-1 == 42) {
                            if false {
                                v43 := []int{128, -2147483648, 9223372036854775807, 0}
                                v4 = 255
                                _ = v43
                            }
                            fmt.Println(v42)
                        } else {
                            v4 = v4
                        }
                        v44 := GenS2{F0: "", F1: 437}
                        _ = v44
                    }
                case 54:
                    if v42 {
                        v45 := GenS4{F0: 2, F1: 0, F2: ("a" + ("foo bar" + ""))}
                        fmt.Println(v42, v42)
                        switch "foo bar" {
                        case "hello":
                            if ("a" > "\n\t") {
                                if (0 != v4) {
                                    for fi46 := 0; fi46 < 15; fi46++ {
                                        var v47 uint8 = (v4 ^ v4)
                                        fmt.Println(fi46, v42)
                                        _ = fi46
                                        _ = v47
                                    }
                                }
                                if v42 {
                                    v42 = v42
                                } else {
                                    if v42 {
                                        v4 = v4
                                        v4 = 2
                                    } else {
                                        if ((false || (true && v42)) && v42) {
                                            v48 := GenS4{F0: 255, F1: -1, F2: ""}
                                            _ = v48
                                        }
                                    }
                                }
                                v1[666] = v42
                            }
                        case "a":
                            for fi49 := 0; fi49 < 3; fi49++ {
                                var v50 uint64 = 42
                                _ = fi49
                                _ = v50
                            }
                        default:
                            var v51 uint16 = 42
                            _ = v51
                        }
                        _ = v45
                    } else {
                        var v52 int32 = -2147483648
                        _ = v52
                    }
                default:
                    switch "hello" {
                    case "a":
                        v53 := GenS7{F0: "a", F1: 9223372036854775807, F2: 32767}
                        _ = v53
                    case "hello":
                        var v54 int16 = 410
                        _ = v54
                    }
                }
                _ = v42
            } else {
                fmt.Println(v4)
                v55 := []int{-536, 0, 42}
                _ = v55
            }
        }
    }
    v56 := []int{genFn1(), -78, 9223372036854775807}
    fmt.Println(v4)
    v1[586] = (!false)
    if false {
        var v57 int8 = -1
        if true {
            var v58 uint = 42
            v56 = append(v56, 127)
            if true {
                v59 := v2[283]
                _ = v59
            } else {
                switch "foo bar" {
                case "a":
                    for fi60 := 0; fi60 < 8; fi60++ {
                        fmt.Println(v57, v58, v58)
                        _ = len(v1)
                        _ = fi60
                    }
                default:
                    var v61 int8 = -128
                    _ = v61
                }
                v58 = v58
                v62 := GenS8{F0: (v4 % (((v4 / (((^(v4))) | 1))) | 1)), F1: (v4 ^ v4), F2: 137}
                _ = v62
            }
            _ = v58
        }
        _ = v57
    }
    for fi63 := 0; fi63 < 15; fi63++ {
        for fi64 := 0; fi64 < 3; fi64++ {
            var v65 int = genFn1()
            v66 := GenS9{F0: 9223372036854775807}
            fmt.Println(v65, v65)
            _ = fi64
            _ = v65
            _ = v66
        }
        for fi67 := 0; fi67 < 3; fi67++ {
            var v68 bool = true
            if (-379 != -290) {
                v68 = (v4 != (v4 - v4))
                v4 += (v4 & (v4 + v4))
                v69 := GenS5{F0: 4294967295, F1: 42}
                _ = v69
            }
            if true {
                v70 := GenS1{F0: -1, F1: false, F2: 9223372036854775807}
                var v71 int32 = -414
                v72 := GenS1{F0: -9223372036854775808, F1: ((v71 & v71) >= 42), F2: fi67}
                _ = v70
                _ = v71
                _ = v72
            } else {
                v68 = v68
                if (((-838 >= 32767) || ((fi67 & fi67) < fi67)) || false) {
                    v68 = (("foo bar" + (((("\n\t" + ("foo bar" + "\n\t")) + "a") + "a") + ("" + "a"))) == "")
                } else {
                    var v73 int32 = 2147483647
                    if (v68 || v68) {
                        v74 := GenS1{F0: 1, F1: (float32(1e10) <= -2.25), F2: 2147483647}
                        v73 += v73
                        _ = v74
                    }
                    v68 = (v68 && (-892 == -1))
                    _ = v73
                }
                switch ("hello" + "foo bar") {
                case "a":
                    switch ("\n\t" + (("" + "foo bar") + "")) {
                    case "a":
                        var v75 int32 = -1
                        _ = v75
                    case "foo bar":
                        if len(v56) >= 2 {
                            v56 = v56[1:len(v56)]
                        }
                        v68 = true
                        if false {
                            v4 = (v4 % (((v4 >> ((4294967295) & 63))) | 1))
                            fmt.Println(v68, fi67)
                        }
                    }
                case "foo bar":
                    switch "" {
                    case "foo bar":
                        var v76 int64 = 0
                        _ = v76
                    }
                case "\n\t":
                    switch (fi67 / ((genFn1()) | 1)) {
                    case 2147483647:
                        v77 := GenS4{F0: 354, F1: 1, F2: ("hello" + "")}
                        _ = v77
                    case 127:
                        v68 = false
                    }
                default:
                    switch (^(fi63)) {
                    case 127:
                        v4 = v4
                    case 0:
                        switch genFn1() {
                        case 128:
                            v4 += v4
                        case 76:
                            fmt.Println(fi63, v4)
                        case -2147483648:
                            fmt.Println(fi67)
                        }
                    case 1:
                        v56 = append(v56, fi67)
                    default:
                        var v78 int = genFn1()
                        _ = v78
                    }
                }
            }
            _ = fi67
            _ = v68
        }
        _ = fi63
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v56:", len(v56), v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v56
}
