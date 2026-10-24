package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 string; F2 int }
type GenS4 struct { F0 int64; F1 int32; F2 uint32 }
type GenS5 struct { F0 int32; F1 int; F2 int64 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint8; F1 int8; F2 int16 }
type GenS8 struct { F0 int; F1 string }
func genFn0(p0 uint64) uint8 {
    v6 := []int{1, 128}
    fmt.Println(p0)
    if true {
        switch -171 {
        case -557:
            fmt.Println(p0)
        }
        switch "" {
        case "foo bar":
            var v7 float64 = 1e10
            _ = v7
        default:
            if len(v6) > 0 {
                v6[0] = 2147483647
            }
        }
        var v8 uint32 = 656
        _ = v8
    }
    v9 := GenS1{F0: 4294967295, F1: 2147483647}
    var v10 uint8 = 1
    _ = p0
    _ = v10
    _ = v6
    _ = v9
    return 255
}
func genFn1(p0 uint8, p1 uint16) int32 {
    switch "hello" {
    case "hello":
        var v11 float32 = (1e10 + (-2.25 + (1.5 - -2.25)))
        _ = v11
    default:
        var v12 bool = true
        _ = v12
    }
    if (-103 > int(p0)) {
        var v13 uint16 = p1
        _ = v13
    } else {
        p1 = (p1 >> ((255) & 63))
        p0 -= 127
        var v14 uint = 2
        _ = v14
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    return -2147483648
}

func main() {
    var v1 uint8 = 51
    var v2 uint8 = v1
    var v3 int8 = 127
    var v4 float32 = (-2.25 - 1.5)
    v5 := make(map[string]int)
    for fi15 := 0; fi15 < 17; fi15++ {
        if (479 > 4294967295) {
            v16 := []int{(fi15 & fi15)}
            _ = v16
        }
        v17 := GenS1{F0: 42, F1: 0}
        var v18 bool = (-1 != 1)
        _ = fi15
        _ = v17
        _ = v18
    }
    fmt.Println(v4, v3, v4)
    for fi19 := 0; fi19 < 7; fi19++ {
        var v20 bool = (9223372036854775807 != 9223372036854775807)
        v21 := GenS1{F0: 622, F1: 1}
        switch "\n\t" {
        case "":
            var v22 int8 = (v3 - (v3 << ((2) & 63)))
            _ = v22
        case "\n\t":
            for fi23 := 0; fi23 < 1; fi23++ {
                v24 := GenS2{F0: 0}
                switch -1 {
                case 426:
                    if false {
                        v2 += genFn0(2)
                    } else {
                        var v25 int64 = 0
                        _ = v25
                    }
                case 9223372036854775807:
                    v4 = float32((-2.25 + 0.1))
                }
                _ = fi23
                _ = v24
            }
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        var v27 int = fi26
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 15; fi28++ {
        if (!false) {
            v29 := GenS3{F0: (!(true || false)), F1: "foo bar", F2: -2147483648}
            switch "" {
            case "hello":
                switch 0 {
                case 2147483647:
                    v30 := GenS3{F0: true, F1: "hello", F2: (fi28 % (((-(fi28))) | 1))}
                    _ = v30
                case 1:
                    if false {
                        for fi31 := 0; fi31 < 14; fi31++ {
                            _ = len(v5)
                            _ = fi31
                        }
                        v32 := GenS2{F0: 18446744073709551615}
                        var v33 uint = 436
                        _ = v32
                        _ = v33
                    }
                default:
                    fmt.Println(v3)
                }
            case "":
                var v34 string = ("a" + "a")
                _ = v34
            case "\n\t":
                var v35 uint = 4294967295
                _ = v35
            }
            _ = v29
        }
        var v36 int = -2147483648
        var v37 int16 = 42
        _ = fi28
        _ = v36
        _ = v37
    }
    fmt.Println(v3)
    var v38 int64 = -962
    var v39 int = 42
    v40 := GenS4{F0: (-(v38)), F1: 0, F2: 2147483648}
    v2 += (v1 / ((genFn0(987)) | 1))
    var v41 uint16 = 1
    v39 += (v39 * 128)
    var v42 int32 = genFn1(genFn0(1), 255)
    v41 -= v41
    v43 := v5["hello"]
    if ((^(v38)) == int64(v41)) {
        if false {
            switch (v39 ^ 434) {
            case 127:
                fmt.Println(v2, v42)
            case 485:
                var v44 string = "hello"
                _ = v44
            default:
                switch 32767 {
                case -1:
                    fmt.Println(v43, v3)
                case 42:
                    v45 := GenS5{F0: -2147483648, F1: v43, F2: (v38 / (((v38 % ((-9223372036854775808) | 1))) | 1))}
                    _ = v45
                }
            }
        } else {
            switch v43 {
            case 2147483647:
                v46 := GenS1{F0: 2, F1: v42}
                _ = v46
            case -2147483648:
                fmt.Println(v4, v43)
            default:
                v5[(((("a" + ("foo bar" + ("hello" + "hello"))) + ("hello" + ((("\n\t" + ("hello" + "hello")) + "hello") + "hello"))) + "\n\t") + "\n\t")] = 892
            }
            switch "hello" {
            case "":
                var v47 uint = 2
                _ = v47
            case "\n\t":
                var v48 uint8 = v1
                _ = v48
            case "foo bar":
                for fi49 := 0; fi49 < 7; fi49++ {
                    for fi50 := 0; fi50 < 13; fi50++ {
                        var v51 int32 = (-(v42))
                        v52 := GenS6{F0: (^(v38))}
                        var v53 float64 = 0.0
                        v54 := GenS7{F0: (v2 * (v2 / ((v2) | 1))), F1: -57, F2: -32768}
                        v5[("a" + ("foo bar" + "a"))] = (fi49 >> ((255) & 63))
                        _ = fi50
                        _ = v51
                        _ = v52
                        _ = v53
                        _ = v54
                    }
                    switch int(v41) {
                    case 990:
                        if (("" + "foo bar") != ("hello" + "\n\t")) {
                            if true {
                                if true {
                                    v5[("a" + ("\n\t" + ""))] = v39
                                    fmt.Println(v3, v39)
                                }
                            } else {
                                v3 -= -81
                                v55 := GenS8{F0: (v43 ^ v39), F1: ("\n\t" + "a")}
                                if true {
                                    v56 := GenS1{F0: 2, F1: v42}
                                    v1 = (^(v2))
                                    _ = v56
                                }
                                _ = v55
                            }
                        } else {
                            fmt.Println(fi49, v1)
                        }
                    case -1:
                        for fi57 := 0; fi57 < 2; fi57++ {
                            switch v39 {
                            case 32768:
                                v2 = uint8(v41)
                            default:
                                v58 := []int{-9223372036854775808}
                                _ = v58
                            }
                            if (!((v3 < 27) && (genFn1(v1, v41) == v42))) {
                                v4 -= (float32(float64(0.0)) + float32(0.1))
                                var v59 uint8 = uint8(v41)
                                switch "hello" {
                                case "a":
                                    v3 = v3
                                case "":
                                    fmt.Println(v43)
                                case "\n\t":
                                    var v60 int64 = (v38 * (^(v38)))
                                    fmt.Println(v60, v1)
                                    var v61 uint = 200
                                    fmt.Println(v4, v4, v43)
                                    _ = v60
                                    _ = v61
                                default:
                                    var v62 uint8 = (v1 * (v2 / ((v1) | 1)))
                                    _ = v62
                                }
                                v4 = v4
                                _ = v59
                            }
                            _ = fi57
                        }
                    case -425:
                        v63 := GenS7{F0: v2, F1: int8(v38), F2: 1}
                        _ = v63
                    default:
                        v64 := v5["a"]
                        _ = v64
                    }
                    _ = fi49
                }
            default:
                var v65 bool = ((false && false) || (v3 <= v3))
                _ = v65
            }
            v66 := []int{9223372036854775807, 1, (^(v43)), (v39 + (v43 << ((0) & 63)))}
            _ = v66
        }
        v42 -= v42
        v67 := GenS8{F0: v43, F1: ("foo bar" + "foo bar")}
        _ = v67
    } else {
        v68 := v5[""]
        _ = v68
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
