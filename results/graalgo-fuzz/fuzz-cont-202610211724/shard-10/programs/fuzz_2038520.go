package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int32; F1 uint8; F2 uint32 }
type GenS4 struct { F0 bool; F1 uint8; F2 int32 }
type GenS5 struct { F0 uint; F1 uint8 }
type GenS6 struct { F0 uint32 }
func genFn0() uint64 {
    var v5 uint16 = 32767
    var v6 uint64 = 1
    var v7 uint32 = 42
    _ = v5
    _ = v6
    _ = v7
    return 18446744073709551615
}

func main() {
    v1 := []string{(("" + "foo bar") + (("\n\t" + "") + "foo bar")), ("hello" + "hello")}
    v2 := make(map[int]bool)
    v3 := []string{("" + (((("\n\t" + ("foo bar" + ("hello" + ""))) + (("a" + "foo bar") + "a")) + ((("" + "hello") + "hello") + "")) + "a")), "\n\t", ("" + "hello"), ("hello" + ("a" + ("a" + (("foo bar" + ("foo bar" + "\n\t")) + ""))))}
    var v4 int64 = 727
    if ((3.14159 / (3.14159 + 0.0)) != 1e10) {
        for fi8 := 0; fi8 < 9; fi8++ {
            for fi9 := 0; fi9 < 18; fi9++ {
                if true {
                    if (32767 > 619) {
                        switch "hello" {
                        case "a":
                            var v10 int32 = -711
                            fmt.Println(fi8)
                            v11 := GenS1{F0: "hello"}
                            _ = v10
                            _ = v11
                        case "":
                            fmt.Println(v4)
                        default:
                            fmt.Println(fi8)
                        }
                        if (true && true) {
                            if false {
                                v12 := GenS1{F0: ("foo bar" + "a")}
                                _ = v12
                            } else {
                                v13 := GenS1{F0: (((("foo bar" + ("\n\t" + "a")) + (("a" + (("hello" + "foo bar") + "hello")) + "hello")) + ("" + ("foo bar" + "a"))) + "a")}
                                var v14 uint = 65535
                                _ = v13
                                _ = v14
                            }
                            v15 := GenS2{F0: fi9, F1: 127, F2: 18446744073709551615}
                            fmt.Println(fi8, v4, v4)
                            _ = v15
                        }
                        var v16 uint = 255
                        _ = v16
                    } else {
                        var v17 uint = 65535
                        v18 := GenS1{F0: ""}
                        v19 := GenS2{F0: (fi9 | fi8), F1: 2, F2: 706}
                        _ = v17
                        _ = v18
                        _ = v19
                    }
                }
                var v20 float32 = (0.0 / 1.5)
                _ = len(v3)
                fmt.Println(v20, fi8, fi8)
                v21 := GenS3{F0: -2147483648, F1: 42, F2: 42}
                if len(v3) > 0 {
                    v3[0] = ""
                }
                _ = fi9
                _ = v20
                _ = v21
            }
            _ = len(v3)
            _ = fi8
        }
        if true {
            for fi22 := 0; fi22 < 18; fi22++ {
                fmt.Println(v4)
                _ = fi22
            }
        }
        for fi23 := 0; fi23 < 5; fi23++ {
            var v24 int16 = 32767
            switch "\n\t" {
            case "\n\t":
                v25 := GenS3{F0: 0, F1: 42, F2: 1}
                _ = v25
            case "":
                for fi26 := 0; fi26 < 2; fi26++ {
                    var v27 uint32 = 0
                    _ = fi26
                    _ = v27
                }
            case "a":
                v4 += v4
            default:
                for fi28 := 0; fi28 < 1; fi28++ {
                    fmt.Println(fi23, v4)
                    _ = fi28
                }
            }
            _ = fi23
            _ = v24
        }
    } else {
        for fi29 := 0; fi29 < 8; fi29++ {
            if ((v4 / ((v4) | 1)) <= (^(v4))) {
                v30 := GenS4{F0: true, F1: 2, F2: 42}
                v4 = -338
                _ = v30
            } else {
                v4 += (v4 * (v4 << ((2) & 63)))
                var v31 uint8 = 255
                v32 := v2[(fi29 | 591)]
                _ = v31
                _ = v32
            }
            v4 -= v4
            _ = fi29
        }
        fmt.Println(v4)
        switch "foo bar" {
        case "":
            for fi33 := 0; fi33 < 11; fi33++ {
                v34 := GenS2{F0: fi33, F1: 127, F2: 1}
                for fi35 := 0; fi35 < 14; fi35++ {
                    switch (fi35 % ((234) | 1)) {
                    case -2147483648:
                        v4 -= (v4 / (((v4 ^ -251)) | 1))
                    case 2147483647:
                        var v36 int64 = v4
                        _ = v36
                    default:
                        var v37 float32 = 3.14159
                        _ = v37
                    }
                    v4 = -322
                    var v38 int16 = 992
                    _ = fi35
                    _ = v38
                }
                for fi39 := 0; fi39 < 2; fi39++ {
                    switch ("" + ("\n\t" + "")) {
                    case "":
                        v40 := GenS2{F0: (fi33 - fi39), F1: 127, F2: 4294967295}
                        _ = v40
                    case "foo bar":
                        v4 += (-(v4))
                    case "a":
                        v1 = append(v1, "")
                    }
                    v41 := GenS5{F0: 0, F1: 86}
                    v4 += (v4 & v4)
                    _ = fi39
                    _ = v41
                }
                _ = fi33
                _ = v34
            }
        }
    }
    fmt.Println(v4)
    if ((v4 << ((2) & 63)) > (v4 / ((-9223372036854775808) | 1))) {
        var v42 int8 = 85
        var v43 bool = false
        _ = v42
        _ = v43
    }
    switch "hello" {
    case "foo bar":
        v4 += -1
    }
    for fi44 := 0; fi44 < 9; fi44++ {
        v4 = (v4 * (^(v4)))
        _ = fi44
    }
    v45 := GenS3{F0: -1, F1: 2, F2: 65535}
    _ = len(v2)
    for fi46 := 0; fi46 < 13; fi46++ {
        fmt.Println(fi46, fi46)
        if true {
            for fi47 := 0; fi47 < 12; fi47++ {
                switch ("\n\t" + "foo bar") {
                case "\n\t":
                    if len(v3) > 0 {
                        v3[0] = ("" + "a")
                    }
                }
                _ = fi47
            }
            if true {
                for fi48 := 0; fi48 < 14; fi48++ {
                    v49 := GenS5{F0: 65535, F1: 0}
                    if false {
                        var v50 uint32 = 0
                        fmt.Println(fi48, v50)
                        _ = v50
                    } else {
                        fmt.Println(v4, v4)
                        if true {
                            v51 := GenS3{F0: 739, F1: 216, F2: 42}
                            if len(v1) > 0 {
                                v1[0] = "a"
                            }
                            v3 = append(v3, "\n\t")
                            v4 += (-(v4))
                            var v52 uint16 = 1
                            _ = v51
                            _ = v52
                        }
                        var v53 bool = (4294967295 > 272)
                        _ = v53
                    }
                    _ = fi48
                    _ = v49
                }
            }
        } else {
            var v54 int8 = -8
            v55 := GenS6{F0: 961}
            _ = v54
            _ = v55
        }
        v4 += -9223372036854775808
        _ = fi46
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
}
