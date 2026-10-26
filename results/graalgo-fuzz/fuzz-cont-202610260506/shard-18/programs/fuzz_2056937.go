package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 bool; F2 uint8 }
type GenS3 struct { F0 string; F1 uint8 }
type GenS4 struct { F0 int64; F1 uint64; F2 int8 }
type GenS5 struct { F0 bool; F1 uint64; F2 int8 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int; F1 string }
type GenS9 struct { F0 uint8; F1 uint64; F2 int }
func genFn0() uint16 {
    v5 := GenS1{F0: "foo bar"}
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := GenS2{F0: ((3.14159 == 1e10) && true), F1: (681 < 4294967295), F2: 0}
        switch ("hello" + ("foo bar" + "foo bar")) {
        case "\n\t":
            var v8 string = (((("foo bar" + "\n\t") + ("\n\t" + "a")) + "foo bar") + "a")
            _ = v8
        default:
            for fi9 := 0; fi9 < 17; fi9++ {
                var v10 uint16 = 982
                var v11 int32 = -192
                _ = fi9
                _ = v10
                _ = v11
            }
        }
        var v12 int32 = 1
        _ = fi6
        _ = v12
        _ = v7
    }
    _ = v5
    return 255
}
func genFn1(p0 uint16, p1 uint) uint32 {
    var v13 uint64 = 947
    p0 += p0
    var v14 int32 = -97
    v13 += v13
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return 811
}

func main() {
    v1 := []string{"hello", "foo bar"}
    var v2 float32 = 1.5
    v3 := []string{(("a" + ("" + "\n\t")) + (("a" + (("" + "foo bar") + ("hello" + (("foo bar" + "a") + "foo bar")))) + "\n\t")), "hello", "hello"}
    v4 := make(map[int]bool)
    for fi15 := 0; fi15 < 4; fi15++ {
        v16 := GenS3{F0: "hello", F1: 236}
        if false {
            for fi17 := 0; fi17 < 5; fi17++ {
                v2 -= 0.1
                v2 = float32((1.5 * 0.0))
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi17
            }
            v18 := GenS4{F0: 9223372036854775807, F1: 4294967295, F2: -3}
            for fi19 := 0; fi19 < 3; fi19++ {
                delete(v4, (fi15 - 32768))
                fmt.Println(fi15)
                v20 := GenS5{F0: false, F1: 0, F2: 42}
                _ = fi19
                _ = v20
            }
            _ = v18
        } else {
            fmt.Println(fi15, v2)
            v2 = 0.1
        }
        v2 = (float32(1.5) + v2)
        _ = fi15
        _ = v16
    }
    var v21 bool = true
    var v22 float64 = 1.5
    var v23 string = ("\n\t" + "a")
    var v24 uint16 = 2
    v23 = v23
    switch ((v23 + v23) + v23) {
    case "foo bar":
        _ = len(v4)
    }
    fmt.Println(v21, v2)
    if (!false) {
        switch 1 {
        case 320:
            v25 := GenS3{F0: "a", F1: 128}
            _ = v25
        default:
            switch -856 {
            case -297:
                var v26 int32 = -1
                _ = v26
            case 128:
                fmt.Println(v22)
            case 127:
                var v27 uint8 = 127
                _ = v27
            }
        }
        for fi28 := 0; fi28 < 7; fi28++ {
            if false {
                v21 = (2 > 653)
                _ = len(v3)
                if (v21 || ("a" >= (((("a" + v23) + v23) + v23) + v23))) {
                    switch (v23 + v23) {
                    case "\n\t":
                        switch "\n\t" {
                        case "":
                            for fi29 := 0; fi29 < 9; fi29++ {
                                fmt.Println(fi28)
                                v2 -= (((v2 * float32((v22 + (float64(v2) + float64(v2))))) / v2) - float32(v22))
                                _ = fi29
                            }
                        }
                    case "hello":
                        var v30 bool = (!v21)
                        _ = v30
                    }
                    v31 := GenS6{F0: (fi28 / (((-(fi28))) | 1))}
                    _ = v31
                }
            } else {
                v24 -= v24
                for fi32 := 0; fi32 < 16; fi32++ {
                    fmt.Println(v21)
                    v23 = v23
                    var v33 int64 = 874
                    _ = fi32
                    _ = v33
                }
            }
            _ = fi28
        }
        delete(v4, -682)
    } else {
        v2 += v2
        if ((728 < 1) || v21) {
            for fi34 := 0; fi34 < 11; fi34++ {
                var v35 bool = true
                v4[fi34] = v21
                if v21 {
                    for fi36 := 0; fi36 < 16; fi36++ {
                        switch "foo bar" {
                        case "hello":
                            v3 = append(v3, v23)
                        case "a":
                            v4[(fi36 * (fi34 >> ((65535) & 63)))] = v35
                        case "":
                            var v37 int32 = -408
                            _ = v37
                        }
                        if v35 {
                            var v38 int16 = -1
                            v39 := GenS7{F0: v21}
                            _ = v38
                            _ = v39
                        }
                        var v40 float64 = ((0.1 / float64((v2 / (1.5 - -2.25)))) + float64(v2))
                        var v41 uint8 = 127
                        _ = fi36
                        _ = v40
                        _ = v41
                    }
                    v22 = ((((3.14159 * v22) + (v22 / 1e10)) - 1e10) * v22)
                }
                _ = fi34
                _ = v35
            }
            fmt.Println(v22, v21)
            _ = len(v3)
        } else {
            var v42 uint16 = v24
            v42 = (v24 ^ 32767)
            _ = v42
        }
        if v21 {
            v43 := v4[32768]
            for fi44 := 0; fi44 < 19; fi44++ {
                switch "hello" {
                case "":
                    v4[(fi44 % (((-(fi44))) | 1))] = (-59 != 117)
                case "foo bar":
                    v23 = ("foo bar" + v23)
                case "hello":
                    v21 = v21
                default:
                    v43 = (true || ((18446744073709551615 < 18446744073709551615) || v21))
                }
                fmt.Println(v43, fi44, v24)
                _ = fi44
            }
            v23 = v23
            _ = v43
        }
    }
    for fi45 := 0; fi45 < 4; fi45++ {
        v23 = "hello"
        v2 = v2
        _ = fi45
    }
    if (-110 == 42) {
        v2 -= 0.0
        _ = len(v3)
        v46 := GenS7{F0: (!(!(v21 && (32768 <= -973))))}
        _ = v46
    }
    var v47 int32 = 1
    var v48 int = 9223372036854775807
    v49 := GenS3{F0: v23, F1: 2}
    v24 += 32767
    switch 42 {
    case 32768:
        fmt.Println(v48, v21)
    case 128:
        for fi50 := 0; fi50 < 20; fi50++ {
            var v51 int32 = 42
            _ = fi50
            _ = v51
        }
    default:
        for fi52 := 0; fi52 < 19; fi52++ {
            var v53 int8 = 42
            switch (fi52 << ((42) & 63)) {
            case 128:
                v47 = v47
            case -609:
                for fi54 := 0; fi54 < 15; fi54++ {
                    if false {
                        switch ("" + v23) {
                        case "foo bar":
                            if len(v3) > 0 {
                                v3[0] = (v23 + "\n\t")
                            }
                        default:
                            switch v23 {
                            case "a":
                                v2 = (v2 + 0.0)
                            case "\n\t":
                                var v55 int16 = 587
                                _ = v55
                            case "hello":
                                var v56 int16 = -888
                                _ = v56
                            }
                        }
                        _ = len(v1)
                        v57 := GenS7{F0: true}
                        _ = v57
                    } else {
                        var v58 int8 = v53
                        var v59 uint16 = (^(v24))
                        v60 := GenS2{F0: v21, F1: v21, F2: 2}
                        _ = v58
                        _ = v59
                        _ = v60
                    }
                    v61 := GenS8{F0: fi54, F1: (v23 + (v23 + v23))}
                    var v62 uint32 = 2147483648
                    _ = fi54
                    _ = v61
                    _ = v62
                }
            default:
                switch v23 {
                case "a":
                    v63 := GenS7{F0: true}
                    _ = v63
                case "\n\t":
                    fmt.Println(v23)
                }
            }
            v64 := GenS9{F0: 0, F1: 4294967295, F2: int(v24)}
            _ = fi52
            _ = v53
            _ = v64
        }
    }
    var v65 bool = true
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v65
}
