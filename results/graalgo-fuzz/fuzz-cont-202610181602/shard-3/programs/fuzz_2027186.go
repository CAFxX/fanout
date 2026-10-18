package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 int64; F1 uint16; F2 uint32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 int32 }
type GenS6 struct { F0 int; F1 int32; F2 int16 }
type GenS7 struct { F0 string; F1 int32; F2 int }
type GenS8 struct { F0 int; F1 int32 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 bool; F1 bool }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int64 }

func main() {
    v1 := make(map[string]string)
    v2 := []int{}
    v3 := []string{}
    v4 := []string{}
    for fi5 := 0; fi5 < 9; fi5++ {
        switch "a" {
        case "":
            var v6 int32 = -1
            _ = v6
        case "foo bar":
            switch fi5 {
            case 1:
                if true {
                    if true {
                        var v7 bool = (592 < 65535)
                        _ = v7
                    } else {
                        if false {
                            fmt.Println(fi5)
                            var v8 uint64 = 4294967295
                            v8 = v8
                            _ = v8
                        } else {
                            if true {
                                v9 := GenS1{F0: -1, F1: 2, F2: -903}
                                _ = v9
                            } else {
                                var v10 uint8 = 2
                                v10 += v10
                                v10 -= v10
                                _ = v10
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                        v11 := GenS2{F0: 833, F1: "foo bar"}
                        _ = v11
                    }
                }
            default:
                var v12 uint8 = 132
                _ = v12
            }
        }
        switch 128 {
        case 9223372036854775807:
            if ((1000 != 864) && (!false)) {
                var v13 int8 = -84
                if (!(((false || true) && (false || ((true && ((239 >= 0) || true)) || true))) || false)) {
                    if (65535 >= 255) {
                        var v14 string = "a"
                        _ = v14
                    }
                }
                if (32767 > 42) {
                    _ = len(v1)
                }
                _ = v13
            } else {
                var v15 uint64 = 2
                _ = v15
            }
        case -747:
            for fi16 := 0; fi16 < 2; fi16++ {
                if (0 != 42) {
                    switch "a" {
                    case "foo bar":
                        var v17 int8 = -128
                        _ = v17
                    default:
                        if false {
                            var v18 uint64 = 893
                            if (!false) {
                                if (!true) {
                                    v19 := GenS3{F0: -52, F1: 32767, F2: 2147483648}
                                    var v20 int64 = 301
                                    _ = v19
                                    _ = v20
                                }
                                var v21 uint32 = 1
                                v21 = (^(v21))
                                v18 = (^(v18))
                                v18 = (v18 % ((v18) | 1))
                                _ = v21
                            } else {
                                v18 = (v18 ^ (v18 >> ((42) & 63)))
                                if (127 != 255) {
                                    if false {
                                        fmt.Println(fi16, fi16)
                                        v18 = (v18 ^ 1)
                                        switch (fi16 + fi5) {
                                        case -2147483648:
                                            var v22 uint16 = 65535
                                            _ = v22
                                        default:
                                            v18 -= v18
                                        }
                                    } else {
                                        if (!false) {
                                            var v23 bool = (0.1 < (float32(-2.25) - -2.25))
                                            v24 := GenS1{F0: -662, F1: 4294967295, F2: -1}
                                            var v25 uint8 = 77
                                            v23 = ((v23 || v23) || false)
                                            _ = v23
                                            _ = v24
                                            _ = v25
                                        } else {
                                            v18 = 1
                                            switch "a" {
                                            case "a":
                                                v18 += (v18 % (((v18 + 18446744073709551615)) | 1))
                                            }
                                            switch ("\n\t" + "hello") {
                                            case "a":
                                                v18 += 42
                                            default:
                                                v26 := GenS4{F0: 2}
                                                var v27 float64 = (3.14159 * 0.0)
                                                _ = v26
                                                _ = v27
                                            }
                                        }
                                    }
                                    v2 = append(v2, (-(fi5)))
                                    switch "a" {
                                    case "\n\t":
                                        var v28 float32 = -2.25
                                        _ = v28
                                    case "foo bar":
                                        v29 := GenS5{F0: (v18 % (((^(v18))) | 1)), F1: -1}
                                        _ = v29
                                    default:
                                        switch -2147483648 {
                                        case 42:
                                            v18 = v18
                                        default:
                                            fmt.Println(v18, v18, v18)
                                        }
                                        if (("\n\t" + "\n\t") > ("hello" + ("\n\t" + ""))) {
                                            switch fi5 {
                                            case -249:
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            case 909:
                                                v30 := GenS3{F0: 1, F1: 264, F2: 838}
                                                _ = v30
                                            default:
                                                fmt.Println(fi16)
                                                fmt.Println(fi16, fi5)
                                            }
                                            var v31 bool = (true || (32767 == 928))
                                            var v32 uint16 = 450
                                            _ = v31
                                            _ = v32
                                        }
                                    }
                                    v33 := GenS4{F0: 2}
                                    _ = v33
                                } else {
                                    if (!true) {
                                        var v34 string = "\n\t"
                                        if (false && (true || ((-128 == 17) && false))) {
                                            var v35 uint32 = 174
                                            _ = v35
                                        }
                                        _ = v34
                                    }
                                }
                            }
                            v36 := GenS6{F0: (fi16 << ((65535) & 63)), F1: -1, F2: 32767}
                            _ = v18
                            _ = v36
                        }
                    }
                    if true {
                        if len(v4) > 0 {
                            v4[0] = "hello"
                        }
                        v37 := GenS7{F0: (("a" + ("a" + ("" + ""))) + ("a" + "\n\t")), F1: 0, F2: (fi16 / (((fi5 - fi16)) | 1))}
                        var v38 int64 = 194
                        v38 -= -100
                        _ = len(v1)
                        fmt.Println(fi5, fi5, v38)
                        _ = v37
                        _ = v38
                    }
                } else {
                    _ = len(v2)
                }
                _ = fi16
            }
        case 1:
            var v39 int16 = -245
            _ = v39
        }
        for fi40 := 0; fi40 < 2; fi40++ {
            var v41 uint32 = 65535
            _ = fi40
            _ = v41
        }
        _ = fi5
    }
    if (true && true) {
        var v42 uint = 65535
        v43 := GenS8{F0: 931, F1: 1}
        _ = v42
        _ = v43
    }
    var v44 uint = 208
    fmt.Println(v44)
    var v45 int64 = -1
    var v46 uint = (^(v44))
    var v47 bool = (true && true)
    if true {
        if v47 {
            for fi48 := 0; fi48 < 3; fi48++ {
                v49 := GenS9{F0: 18446744073709551615}
                _ = fi48
                _ = v49
            }
        }
        if false {
            fmt.Println(v44)
        } else {
            v50 := GenS10{F0: (!true), F1: v47}
            v47 = v47
            _ = v50
        }
    }
    v44 = v46
    if v47 {
        fmt.Println(v47, v46)
        fmt.Println(v45, v47)
        var v51 int = 699
        _ = v51
    }
    if (v45 == (v45 & 0)) {
        if v47 {
            v45 -= 9223372036854775807
            var v52 int32 = 1
            v1["hello"] = ("" + "\n\t")
            _ = v52
        } else {
            if v47 {
                for fi53 := 0; fi53 < 9; fi53++ {
                    var v54 uint64 = 18446744073709551615
                    v45 -= (v45 % (((v45 + (v45 & 1))) | 1))
                    _ = fi53
                    _ = v54
                }
            }
            v55 := GenS6{F0: -2147483648, F1: -2147483648, F2: 42}
            switch 127 {
            case 127:
                v2 = append(v2, -796)
            case 1:
                var v56 float64 = 0.1
                _ = v56
            default:
                if true {
                    var v57 uint16 = 65535
                    v58 := GenS11{F0: 1}
                    v59 := GenS12{F0: v45}
                    _ = v57
                    _ = v58
                    _ = v59
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        var v60 bool = (!(!v47))
                        _ = v60
                    case "foo bar":
                        var v61 int8 = -128
                        _ = v61
                    }
                    v45 -= (-(v45))
                    v44 = (v44 % ((65535) | 1))
                }
            }
            _ = v55
        }
    }
    _ = len(v4)
    var v62 int = 0
    v63 := GenS10{F0: v47, F1: (2147483648 == 706)}
    if v47 {
        var v64 float32 = 0.0
        _ = v64
    }
    switch ("\n\t" + "") {
    case "\n\t":
        v46 += 65535
    default:
        fmt.Println(v62)
    }
    v65 := GenS6{F0: 0, F1: 1, F2: 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v62
    _ = v63
    _ = v65
}
