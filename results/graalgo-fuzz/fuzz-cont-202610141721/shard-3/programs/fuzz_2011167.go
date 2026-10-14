package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 int64; F1 int16; F2 uint64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint64 }
func genFn0(p0 uint) uint8 {
    v5 := []int{2147483647, 128}
    var v6 int64 = 42
    if false {
        var v7 uint = p0
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 int8 = 127
            _ = fi8
            _ = v9
        }
        var v10 uint64 = 0
        _ = v10
        _ = v7
    }
    _ = p0
    _ = v5
    _ = v6
    return uint8(p0)
}
func genFn1() int64 {
    if true {
        v11 := GenS1{F0: 371}
        var v12 uint = 2
        _ = v11
        _ = v12
    } else {
        var v13 int16 = -1
        v13 -= 1
        if false {
            switch 791 {
            case 830:
                fmt.Println(v13)
            case 4294967295:
                v13 = 472
            }
        } else {
            var v14 int = 42
            fmt.Println(v13, v14)
            _ = v14
        }
        _ = v13
    }
    switch 9223372036854775807 {
    case 128:
        v15 := []int{516, 2147483647}
        _ = v15
    case 32767:
        var v16 int64 = -9223372036854775808
        _ = v16
    default:
        var v17 int8 = 42
        _ = v17
    }
    var v18 uint64 = 467
    v18 += 867
    switch 128 {
    case 127:
        v19 := GenS1{F0: -966}
        _ = v19
    case 1:
        fmt.Println(v18)
    }
    _ = v18
    return 9223372036854775807
}

func main() {
    v1 := []string{(("a" + ("a" + (("\n\t" + (("\n\t" + (("foo bar" + ("\n\t" + "\n\t")) + ("a" + "a"))) + ("" + "foo bar"))) + "a"))) + "foo bar"), ("a" + ("hello" + "")), "foo bar", "\n\t"}
    var v2 uint32 = 534
    var v3 int16 = -975
    v4 := []uint32{(^(v2)), 341, (v2 * (v2 & v2)), v2}
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if true {
            for fi20 := 0; fi20 < 20; fi20++ {
                for fi21 := 0; fi21 < 10; fi21++ {
                    var v22 string = "\n\t"
                    _ = fi21
                    _ = v22
                }
                _ = len(v1)
                v2 -= v2
                _ = fi20
            }
        } else {
            _ = len(v4)
            v23 := GenS2{F0: v2, F1: (!true)}
            if true {
                for fi24 := 0; fi24 < 16; fi24++ {
                    v2 -= v2
                    fmt.Println(v3)
                    v2 = (v2 ^ 2147483647)
                    _ = fi24
                }
            }
            _ = v23
        }
    }
    var v25 uint = 946
    switch -930 {
    case -2147483648:
        v26 := GenS3{F0: 1, F1: v3, F2: 335}
        _ = v26
    default:
        v2 = v2
    }
    v2 -= (v2 % ((v2) | 1))
    v27 := GenS3{F0: 0, F1: 1, F2: 18446744073709551615}
    v2 = (v2 - v2)
    v28 := GenS4{F0: (v25 | v25)}
    if (!(!true)) {
        v3 = 794
        switch 2147483647 {
        case 209:
            switch "" {
            case "a":
                v4 = append(v4, v2)
            case "hello":
                for fi29 := 0; fi29 < 20; fi29++ {
                    if true {
                        for fi30 := 0; fi30 < 11; fi30++ {
                            fmt.Println(fi30)
                            _ = fi30
                        }
                        for fi31 := 0; fi31 < 7; fi31++ {
                            if true {
                                v32 := GenS5{F0: ""}
                                var v33 float64 = (-2.25 + 1.5)
                                fmt.Println(fi29)
                                _ = v32
                                _ = v33
                            }
                            _ = len(v1)
                            v34 := GenS4{F0: (^(v25))}
                            _ = fi31
                            _ = v34
                        }
                        for fi35 := 0; fi35 < 12; fi35++ {
                            switch "a" {
                            case "hello":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            case "":
                                v36 := GenS3{F0: int64(v2), F1: v3, F2: 652}
                                _ = v36
                            case "a":
                                v3 -= (v3 << (((v25 / (((v25 | (v25 - uint(v2)))) | 1))) & 63))
                            }
                            _ = fi35
                        }
                    } else {
                        var v37 int8 = 100
                        for fi38 := 0; fi38 < 11; fi38++ {
                            v2 = v2
                            _ = fi38
                        }
                        _ = v37
                    }
                    if (true && (false && (249 <= 255))) {
                        switch (fi29 & -2147483648) {
                        case 0:
                            v39 := GenS1{F0: 2147483647}
                            _ = v39
                        case 1:
                            v25 += 18446744073709551615
                        default:
                            v25 = v25
                        }
                        switch "" {
                        case "\n\t":
                            var v40 uint16 = 462
                            _ = v40
                        default:
                            if len(v4) > 0 {
                                v4[0] = 4294967295
                            }
                        }
                    }
                    _ = fi29
                }
            case "foo bar":
                if (false || (1e10 < float32(-2.25))) {
                    for fi41 := 0; fi41 < 3; fi41++ {
                        v3 += v3
                        _ = fi41
                    }
                } else {
                    fmt.Println(v3, v25, v25)
                }
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        case 32767:
            switch "" {
            case "hello":
                v42 := GenS5{F0: "\n\t"}
                _ = v42
            case "":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            default:
                v43 := GenS3{F0: 982, F1: (v3 << ((v25) & 63)), F2: 239}
                _ = v43
            }
        case 158:
            v25 += (v25 & (v25 / ((126) | 1)))
        }
        fmt.Println(v25, v2)
    }
    for fi44 := 0; fi44 < 16; fi44++ {
        var v45 uint32 = (v2 << ((v25) & 63))
        _ = fi44
        _ = v45
    }
    for fi46 := 0; fi46 < 12; fi46++ {
        var v47 int = fi46
        for fi48 := 0; fi48 < 5; fi48++ {
            var v49 float32 = 1.5
            v50 := GenS1{F0: -1}
            switch (fi48 >> ((v25) & 63)) {
            case 1:
                switch (v47 - (fi48 << ((v25) & 63))) {
                case 32768:
                    v4 = append(v4, v2)
                    if true {
                        var v51 uint = (v25 << ((42) & 63))
                        if false {
                            v51 = 42
                            v52 := GenS5{F0: ("foo bar" + "a")}
                            _ = v52
                        }
                        var v53 bool = false
                        _ = v51
                        _ = v53
                    } else {
                        v2 += (v2 - v2)
                    }
                    var v54 int = fi46
                    v54 += (v47 + -797)
                    _ = v54
                }
            }
            _ = fi48
            _ = v49
            _ = v50
        }
        _ = fi46
        _ = v47
    }
    var v55 bool = false
    var v56 float32 = (-2.25 - 0.1)
    if len(v4) > 0 {
        v4[0] = v2
    }
    v55 = true
    if false {
        var v57 int64 = -9223372036854775808
        fmt.Println(v55, v55)
        if (!(!((v3 < int16(v25)) && true))) {
            switch ((((("foo bar" + "a") + "foo bar") + (("\n\t" + ("" + ("\n\t" + ("hello" + "foo bar")))) + ("" + "hello"))) + "a") + "\n\t") {
            case "foo bar":
                fmt.Println(v56, v56)
            case "\n\t":
                v56 += (1e10 + (1e10 + 0.0))
            case "hello":
                for fi58 := 0; fi58 < 20; fi58++ {
                    fmt.Println(fi58, fi58)
                    switch (fi58 << (((v25 >> (((v25 | v25)) & 63))) & 63)) {
                    case 4294967295:
                        v57 -= (^(v57))
                    default:
                        var v59 int8 = 100
                        _ = v59
                    }
                    var v60 bool = true
                    _ = fi58
                    _ = v60
                }
            }
        }
        _ = v57
    } else {
        v2 = 2147483647
        for fi61 := 0; fi61 < 7; fi61++ {
            if v55 {
                v3 -= (v3 / ((141) | 1))
                var v62 int16 = v3
                _ = v62
            }
            v63 := GenS6{F0: uint64(v25)}
            var v64 float32 = v56
            _ = fi61
            _ = v63
            _ = v64
        }
        if v55 {
            v65 := GenS4{F0: 18446744073709551615}
            _ = v65
        }
    }
    v3 = v3
    var v66 uint64 = 162
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v55
    _ = v56
    _ = v66
}
