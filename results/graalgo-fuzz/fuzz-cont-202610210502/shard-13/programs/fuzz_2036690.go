package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 string }
type GenS2 struct { F0 int8; F1 bool; F2 int }

func main() {
    v1 := []string{("a" + ((("a" + ("a" + ("foo bar" + "a"))) + "a") + ("hello" + (("\n\t" + "") + (("a" + "foo bar") + ("" + "a")))))), ""}
    v2 := []string{}
    v3 := make(map[string]int)
    var v4 uint32 = 749
    var v5 float64 = ((((0.0 / (0.1 + 3.14159)) + 1e10) * ((1e10 / (0.0 + 3.14159)) * 0.1)) * 1.5)
    v6 := make(map[int]bool)
    fmt.Println(v5)
    switch "hello" {
    case "a":
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 float32 = 1.5
            _ = fi7
            _ = v8
        }
    }
    if false {
        v4 += 112
    }
    v4 = (v4 / ((v4) | 1))
    for fi9 := 0; fi9 < 9; fi9++ {
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 string = "a"
            if false {
                var v12 string = (((((v11 + v11) + "\n\t") + v11) + (v11 + "hello")) + (("a" + "") + ""))
                _ = v12
            }
            if true {
                var v13 bool = false
                switch (fi10 + fi10) {
                case -267:
                    fmt.Println(fi10, fi9, fi10)
                case 485:
                    v14 := v6[fi9]
                    _ = v14
                }
                switch (v11 + (v11 + ("foo bar" + v11))) {
                case "hello":
                    v15 := v6[fi10]
                    _ = v15
                case "foo bar":
                    if v13 {
                        delete(v3, v11)
                        if (false && false) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            if len(v2) > 0 {
                                v2[0] = ("hello" + v11)
                            }
                            var v16 int16 = -236
                            _ = v16
                        }
                    }
                case "a":
                    var v17 int = fi10
                    _ = v17
                default:
                    fmt.Println(v11)
                }
                _ = v13
            } else {
                fmt.Println(v4, v5)
                var v18 uint = 983
                _ = v18
            }
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    _ = len(v1)
    v5 -= 0.1
    if false {
        if len(v2) > 0 {
            v2[0] = "a"
        }
        var v19 uint64 = 42
        _ = v19
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        v5 += 1.5
        if (107 >= -1) {
            if ((v5 / -2.25) >= v5) {
                var v21 int32 = 1
                _ = v21
            } else {
                if true {
                    var v22 uint = 255
                    if (!true) {
                        v22 -= 617
                    } else {
                        if true {
                            var v23 uint16 = 129
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v24 int8 = 90
                            _ = v23
                            _ = v24
                        }
                        switch 172 {
                        case 32767:
                            var v25 int32 = 1
                            _ = v25
                        case 128:
                            for fi26 := 0; fi26 < 18; fi26++ {
                                v1 = append(v1, "\n\t")
                                fmt.Println(fi20)
                                _ = fi26
                            }
                        case 0:
                            fmt.Println(v5, fi20, v4)
                        default:
                            var v27 uint64 = 4294967295
                            _ = v27
                        }
                        var v28 uint64 = 42
                        _ = v28
                    }
                    if (!(((4294967295 == 42) || false) && true)) {
                        v5 += v5
                    } else {
                        fmt.Println(v5)
                        var v29 uint16 = 42
                        _ = v29
                    }
                    _ = v22
                }
                for fi30 := 0; fi30 < 5; fi30++ {
                    if (255 != 32767) {
                        var v31 uint32 = v4
                        fmt.Println(fi30, fi20, fi30)
                        _ = v31
                    }
                    _ = fi30
                }
                var v32 uint8 = 0
                _ = v32
            }
            var v33 int16 = 1
            if len(v1) > 0 {
                v1[0] = "a"
            }
            _ = v33
        } else {
            var v34 uint16 = 32767
            switch "a" {
            case "hello":
                v35 := GenS1{F0: 2, F1: 42, F2: ("" + "")}
                _ = v35
            case "foo bar":
                if false {
                    v36 := GenS1{F0: 18446744073709551615, F1: 42, F2: (("" + ("foo bar" + ("foo bar" + "foo bar"))) + "a")}
                    v37 := GenS1{F0: 42, F1: 55, F2: ""}
                    _ = v36
                    _ = v37
                } else {
                    var v38 int64 = -945
                    v39 := GenS1{F0: uint(v34), F1: 1, F2: "foo bar"}
                    _ = v38
                    _ = v39
                }
            case "\n\t":
                var v40 string = "\n\t"
                _ = v40
            default:
                v41 := v3["\n\t"]
                _ = v41
            }
            for fi42 := 0; fi42 < 7; fi42++ {
                v4 += v4
                if ((false && false) && false) {
                    if (false && false) {
                        var v43 int16 = 1
                        _ = len(v2)
                        v2 = append(v2, "a")
                        _ = len(v2)
                        _ = v43
                    } else {
                        v44 := GenS1{F0: 2, F1: -128, F2: "a"}
                        v34 += (^(v34))
                        _ = v44
                    }
                    var v45 int64 = 42
                    var v46 float64 = -2.25
                    var v47 float32 = 1.5
                    switch fi20 {
                    case 294:
                        v48 := GenS1{F0: 1, F1: 127, F2: ("foo bar" + "a")}
                        _ = v48
                    case 128:
                        v46 = (v5 - v46)
                    case 1:
                        fmt.Println(v47, v5, v4)
                    default:
                        if false {
                            _ = len(v3)
                            fmt.Println(fi20, v4, v45)
                            v49 := v6[459]
                            _ = v49
                        }
                    }
                    v50 := GenS1{F0: 42, F1: 0, F2: "foo bar"}
                    v51 := GenS1{F0: uint(v34), F1: 1, F2: "foo bar"}
                    _ = v45
                    _ = v46
                    _ = v47
                    _ = v50
                    _ = v51
                } else {
                    if len(v2) > 0 {
                        v2[0] = "foo bar"
                    }
                    var v52 int8 = 1
                    _ = v52
                }
                v4 -= v4
                v53 := GenS2{F0: 1, F1: false, F2: (fi20 % ((fi20) | 1))}
                v5 -= 1e10
                _ = fi42
                _ = v53
            }
            _ = v34
        }
        _ = fi20
    }
    v54 := GenS2{F0: -74, F1: (1.5 > ((0.1 / -2.25) + float32(v5))), F2: -643}
    switch -2147483648 {
    case 32767:
        v4 -= (v4 ^ (v4 | v4))
    case 1:
        for fi55 := 0; fi55 < 3; fi55++ {
            v5 = v5
            if (true || true) {
                var v56 int8 = -128
                _ = v56
            }
            if (!(!true)) {
                var v57 uint32 = 2147483648
                var v58 uint64 = 327
                if false {
                    v58 -= (v58 + v58)
                    v57 += 4294967295
                }
                _ = v57
                _ = v58
            } else {
                fmt.Println(v4)
                var v59 uint64 = 1
                v4 += (v4 * 2)
                _ = v59
            }
            _ = fi55
        }
    }
    var v60 int16 = -259
    v60 += (v60 & (v60 | (v60 - (v60 >> ((0) & 63)))))
    switch ("" + "a") {
    case "hello":
        v60 += (v60 & v60)
    case "a":
        var v61 uint = 2
        _ = v61
    case "foo bar":
        switch 32768 {
        case 9223372036854775807:
            v62 := GenS1{F0: 255, F1: 1, F2: "foo bar"}
            _ = v62
        case -842:
            var v63 float32 = 1.5
            _ = v63
        default:
            var v64 bool = (42 > 11)
            _ = v64
        }
    }
    for fi65 := 0; fi65 < 12; fi65++ {
        switch int(v60) {
        case 127:
            v4 += v4
        case 4294967295:
            v5 -= v5
        case -9223372036854775808:
            v4 = 198
        default:
            switch (("hello" + "a") + "foo bar") {
            case "":
                _ = len(v2)
            }
        }
        _ = fi65
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v6
    _ = v60
}
