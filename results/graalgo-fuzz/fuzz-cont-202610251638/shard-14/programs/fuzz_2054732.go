package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 int16; F2 uint }
type GenS3 struct { F0 int; F1 int64; F2 uint64 }
type GenS4 struct { F0 uint8; F1 uint; F2 string }
type GenS5 struct { F0 uint32; F1 uint16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint8; F1 int8 }
type GenS8 struct { F0 uint32; F1 uint32 }
func genFn0(p0 bool) bool {
    for fi4 := 0; fi4 < 7; fi4++ {
        p0 = (p0 && false)
        var v5 float64 = (-2.25 / ((1e10 - ((float64(((((1.5 * float32(1.5)) * 1.5) - 1.5) - 3.14159)) - 0.1) * -2.25)) / -2.25))
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: 845}
    _ = p0
    _ = v6
    return (false && (true || p0))
}

func main() {
    var v1 uint8 = 128
    var v2 float64 = 0.1
    v3 := []string{}
    v2 = (v2 * float64(1.5))
    v2 -= ((3.14159 + v2) - (v2 * 0.0))
    if (v2 != v2) {
        v7 := GenS2{F0: 18446744073709551615, F1: 229, F2: 919}
        var v8 string = ""
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 int64 = 1
            var v11 string = v8
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = v7
        _ = v8
    }
    fmt.Println(v2, v2)
    for fi12 := 0; fi12 < 3; fi12++ {
        v13 := GenS3{F0: (^(fi12)), F1: 266, F2: 1}
        if true {
            var v14 bool = false
            var v15 int8 = 100
            _ = v14
            _ = v15
        }
        for fi16 := 0; fi16 < 13; fi16++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi16
        }
        _ = fi12
        _ = v13
    }
    switch "a" {
    case "":
        for fi17 := 0; fi17 < 13; fi17++ {
            switch -169 {
            case 127:
                if ((v1 % ((v1) | 1)) <= 128) {
                    switch ("\n\t" + "a") {
                    case "foo bar":
                        v1 += v1
                    case "":
                        switch ((("\n\t" + ("" + "hello")) + "foo bar") + "") {
                        case "":
                            var v18 float32 = (3.14159 / 1e10)
                            _ = v18
                        default:
                            fmt.Println(v1)
                        }
                    }
                }
            default:
                v2 = (v2 + v2)
            }
            var v19 int8 = -128
            _ = fi17
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 7; fi20++ {
        switch ("" + (((("" + "hello") + "") + "foo bar") + ("a" + "\n\t"))) {
        case "":
            switch -1 {
            case 128:
                v2 = (v2 / -2.25)
            case -1:
                _ = len(v3)
            case 870:
                var v21 int32 = 0
                _ = v21
            default:
                if ((v1 + (v1 + v1)) <= 255) {
                    switch -274 {
                    case 4294967295:
                        switch "hello" {
                        case "a":
                            _ = len(v3)
                        case "foo bar":
                            for fi22 := 0; fi22 < 1; fi22++ {
                                v23 := GenS4{F0: v1, F1: uint(v1), F2: ("\n\t" + "a")}
                                var v24 uint32 = 73
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                        case "":
                            var v25 uint16 = 721
                            _ = v25
                        }
                    case -2147483648:
                        switch ("hello" + "foo bar") {
                        case "\n\t":
                            v26 := GenS4{F0: (^(v1)), F1: 0, F2: "foo bar"}
                            _ = v26
                        default:
                            var v27 int8 = -47
                            _ = v27
                        }
                    case 32767:
                        v28 := GenS3{F0: (fi20 / ((fi20) | 1)), F1: 9223372036854775807, F2: 662}
                        _ = v28
                    }
                }
            }
        default:
            v2 = v2
        }
        v1 -= v1
        _ = fi20
    }
    for fi29 := 0; fi29 < 4; fi29++ {
        var v30 uint8 = 128
        v1 -= 128
        v30 = 0
        _ = fi29
        _ = v30
    }
    v31 := GenS1{F0: 255}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v32 := GenS1{F0: 255}
    var v33 int = 32768
    if false {
        for fi34 := 0; fi34 < 15; fi34++ {
            v35 := GenS5{F0: 329, F1: 2}
            for fi36 := 0; fi36 < 11; fi36++ {
                v37 := GenS6{F0: fi34}
                fmt.Println(v2, v2, v33)
                if true {
                    var v38 float64 = (-2.25 * 1e10)
                    _ = v38
                } else {
                    if false {
                        _ = len(v3)
                        if (!false) {
                            fmt.Println(fi34, v33, v33)
                            switch ("hello" + ("foo bar" + "\n\t")) {
                            case "":
                                if true {
                                    if len(v3) > 0 {
                                        v3[0] = ("foo bar" + ("" + (("" + "hello") + "foo bar")))
                                    }
                                    v33 -= (v33 + (^(v33)))
                                    var v39 int = (v33 & fi34)
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = v39
                                } else {
                                    var v40 bool = ("a" <= "")
                                    if (1 <= 2147483647) {
                                        switch "foo bar" {
                                        case "":
                                            var v41 float64 = float64(0.1)
                                            _ = v41
                                        case "\n\t":
                                            var v42 uint16 = 681
                                            _ = v42
                                        default:
                                            var v43 uint8 = 255
                                            _ = v43
                                        }
                                        v33 -= v33
                                    }
                                    v44 := GenS6{F0: v33}
                                    _ = v40
                                    _ = v44
                                }
                            default:
                                v45 := GenS7{F0: v1, F1: 100}
                                _ = v45
                            }
                            v46 := GenS1{F0: 229}
                            _ = v46
                        } else {
                            var v47 bool = false
                            v47 = false
                            _ = v47
                        }
                    }
                    var v48 uint32 = 65535
                    _ = v48
                }
                v2 += ((((1.5 / 1.0) - v2) + 1e10) * (float64(1.5) + (v2 + (v2 - v2))))
                v2 -= 0.1
                v1 -= (v1 ^ 1)
                _ = fi36
                _ = v37
            }
            _ = fi34
            _ = v35
        }
        v33 += v33
    } else {
        switch v33 {
        case 127:
            var v49 uint = 2
            _ = v49
        case 476:
            fmt.Println(v33, v2)
        }
    }
    switch ((("a" + "\n\t") + ("foo bar" + "\n\t")) + "\n\t") {
    case "a":
        v50 := GenS8{F0: 708, F1: 2147483648}
        _ = v50
    case "\n\t":
        var v51 float32 = 0.0
        _ = v51
    case "":
        var v52 uint = 0
        _ = v52
    }
    for fi53 := 0; fi53 < 9; fi53++ {
        var v54 string = "\n\t"
        v55 := GenS4{F0: v1, F1: 134, F2: v54}
        var v56 uint16 = 2
        _ = fi53
        _ = v54
        _ = v55
        _ = v56
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
}
