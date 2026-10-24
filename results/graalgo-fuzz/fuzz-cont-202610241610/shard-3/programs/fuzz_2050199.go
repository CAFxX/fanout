package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint64; F1 uint; F2 int64 }
type GenS8 struct { F0 int64; F1 bool; F2 bool }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint32; F1 string; F2 uint }
type GenS11 struct { F0 uint64; F1 uint }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 string; F1 uint64 }
func genFn0(p0 uint8) uint {
    var v7 bool = true
    v8 := GenS1{F0: (p0 << ((65535) & 63))}
    _ = p0
    _ = v7
    _ = v8
    return 65535
}
func genFn1(p0 uint16, p1 int32) int64 {
    switch "foo bar" {
    case "\n\t":
        for fi9 := 0; fi9 < 4; fi9++ {
            if false {
                p1 += 2147483647
                if (p0 > 374) {
                    var v10 float32 = 1e10
                    _ = v10
                } else {
                    for fi11 := 0; fi11 < 11; fi11++ {
                        if true {
                            fmt.Println(fi9, p1, p0)
                            if false {
                                v12 := []int{(^(fi9))}
                                if ((("\n\t" + ("\n\t" + "hello")) + ("a" + ("\n\t" + "hello"))) >= ("a" + "\n\t")) {
                                    var v13 uint = 841
                                    _ = v13
                                } else {
                                    if false {
                                        v14 := GenS2{F0: (p0 * p0)}
                                        if ((true && (false || (true && (p1 != (p1 + 0))))) && (0.1 < 1e10)) {
                                            if (42 >= 1) {
                                                fmt.Println(fi9)
                                                switch "hello" {
                                                case "\n\t":
                                                    var v15 int16 = 32767
                                                    _ = v15
                                                default:
                                                    v16 := GenS3{F0: p1}
                                                    _ = v16
                                                }
                                                v17 := GenS4{F0: ((("" + (((("foo bar" + "") + "foo bar") + "a") + "foo bar")) + ("\n\t" + "\n\t")) + "\n\t")}
                                                _ = v17
                                            }
                                            p0 = 42
                                            v12 = append(v12, fi9)
                                        }
                                        switch "\n\t" {
                                        case "foo bar":
                                            if ((!(1 != p0)) && false) {
                                                v12 = append(v12, (fi11 / (((fi9 << ((4294967295) & 63))) | 1)))
                                            }
                                        case "hello":
                                            p1 = p1
                                        case "":
                                            p1 = -1
                                        default:
                                            if (!(((true || (2147483647 > 2147483647)) && (!((fi9 / ((fi9) | 1)) < (fi11 - 0)))) || (int8(p0) < -35))) {
                                                switch (("a" + "hello") + "") {
                                                case "":
                                                    var v18 bool = (true && false)
                                                    _ = v18
                                                }
                                                fmt.Println(fi11)
                                            }
                                        }
                                        _ = v14
                                    }
                                }
                                _ = v12
                            }
                        }
                        _ = fi11
                    }
                }
                fmt.Println(p0)
            } else {
                v19 := GenS4{F0: "a"}
                v20 := GenS3{F0: (p1 + (p1 - p1))}
                if (-537 < -32768) {
                    p0 = p0
                    switch "" {
                    case "a":
                        v21 := GenS2{F0: p0}
                        _ = v21
                    case "hello":
                        for fi22 := 0; fi22 < 17; fi22++ {
                            v23 := GenS4{F0: "foo bar"}
                            v24 := GenS5{F0: 1, F1: 18446744073709551615, F2: 42}
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                    case "foo bar":
                        v25 := GenS6{F0: ("hello" + "a")}
                        _ = v25
                    default:
                        if (!(1000 < -1)) {
                            switch -2147483648 {
                            case 42:
                                for fi26 := 0; fi26 < 8; fi26++ {
                                    var v27 uint64 = 4294967295
                                    if (2 > 42) {
                                        p1 = p1
                                        var v28 float32 = (0.0 - float32((1e10 * (float64(-2.25) * -2.25))))
                                        _ = v28
                                    }
                                    _ = fi26
                                    _ = v27
                                }
                            }
                            var v29 float32 = float32(-2.25)
                            _ = v29
                        }
                    }
                }
                _ = v19
                _ = v20
            }
            _ = fi9
        }
    case "foo bar":
        if true {
            v30 := GenS7{F0: 42, F1: genFn0(2), F2: 912}
            if false {
                p0 -= p0
                v31 := []int{-2147483648}
                _ = v31
            }
            _ = v30
        }
    default:
        v32 := GenS8{F0: -1, F1: (true && false), F2: true}
        _ = v32
    }
    v33 := GenS8{F0: 1, F1: (int8(p1) <= int8(p0)), F2: (!(11 >= 42))}
    v34 := GenS9{F0: p1}
    var v35 float64 = 0.1
    var v36 string = "foo bar"
    _ = p0
    _ = p1
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    return 1
}

func main() {
    var v1 uint64 = 507
    var v2 int64 = -144
    var v3 uint8 = 127
    var v4 float32 = 0.1
    v5 := []string{"\n\t", "hello", ("hello" + ((("hello" + "hello") + ("hello" + "hello")) + (("foo bar" + "hello") + ("\n\t" + ((("\n\t" + "\n\t") + "\n\t") + "\n\t"))))), ""}
    var v6 uint32 = 42
    switch 1 {
    case -763:
        v37 := GenS8{F0: (-(v2)), F1: (0 <= 614), F2: true}
        _ = v37
    }
    if (true || true) {
        v2 = int64(v6)
    } else {
        var v38 int16 = int16(v1)
        for fi39 := 0; fi39 < 10; fi39++ {
            var v40 uint = 859
            fmt.Println(v1)
            var v41 uint32 = v6
            _ = fi39
            _ = v40
            _ = v41
        }
        _ = v38
    }
    v42 := GenS1{F0: (v3 ^ v3)}
    fmt.Println(v3)
    fmt.Println(v2, v2)
    v43 := GenS10{F0: 2, F1: "hello", F2: 65535}
    for fi44 := 0; fi44 < 3; fi44++ {
        if (128 > v3) {
            fmt.Println(v2, v4)
            v1 = (^(v1))
            if len(v5) > 0 {
                v5[0] = "foo bar"
            }
        }
        var v45 uint64 = v1
        v46 := GenS5{F0: genFn1(1, 0), F1: v1, F2: v6}
        _ = fi44
        _ = v45
        _ = v46
    }
    var v47 uint = 65535
    var v48 uint8 = (v3 << ((v47) & 63))
    v47 = v47
    var v49 int64 = (-(v2))
    switch 127 {
    case 32768:
        for fi50 := 0; fi50 < 20; fi50++ {
            if (genFn0(v3) < v47) {
                if (127 != -117) {
                    fmt.Println(v6, v2, v48)
                } else {
                    v2 = v2
                    v51 := GenS4{F0: "foo bar"}
                    v52 := GenS10{F0: v6, F1: "a", F2: 198}
                    _ = v51
                    _ = v52
                }
                var v53 bool = true
                _ = v53
            } else {
                v54 := GenS1{F0: (^(v48))}
                if true {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v55 := GenS5{F0: (v2 << (((^(v47))) & 63)), F1: v1, F2: 1}
                    _ = v55
                } else {
                    fmt.Println(v4, v3)
                }
                v47 += (v47 - v47)
                _ = v54
            }
            var v56 int64 = (v2 ^ genFn1(710, -993))
            v57 := GenS11{F0: 594, F1: v47}
            _ = fi50
            _ = v56
            _ = v57
        }
    case -1:
        for fi58 := 0; fi58 < 13; fi58++ {
            v4 -= (v4 * v4)
            var v59 float64 = 1.5
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi58
            _ = v59
        }
    default:
        switch ("" + (("\n\t" + "foo bar") + ("a" + "hello"))) {
        case "foo bar":
            v60 := GenS8{F0: (v2 ^ (v2 + v49)), F1: false, F2: (!false)}
            _ = v60
        case "a":
            var v61 int8 = -128
            _ = v61
        case "hello":
            var v62 int8 = 100
            _ = v62
        }
    }
    var v63 uint32 = 2147483647
    fmt.Println(v47, v6, v6)
    fmt.Println(v47, v48, v1)
    v1 = 864
    if (!true) {
        var v64 int64 = v2
        _ = v64
    } else {
        v65 := GenS6{F0: ""}
        v66 := GenS12{F0: v49}
        _ = v65
        _ = v66
    }
    v67 := GenS13{F0: (("hello" + (("a" + "hello") + (("foo bar" + "foo bar") + ((("a" + "hello") + "hello") + "hello")))) + ""), F1: (^(v1))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
    _ = v63
    _ = v67
}
