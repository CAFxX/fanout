package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS6 struct { F0 int32; F1 int64; F2 int8 }
type GenS7 struct { F0 int8; F1 string; F2 int32 }
func genFn0(p0 bool, p1 uint32) uint16 {
    v6 := GenS1{F0: "\n\t", F1: 42}
    v7 := GenS2{F0: 0}
    var v8 uint64 = 559
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 1
}
func genFn1() int8 {
    for fi9 := 0; fi9 < 16; fi9++ {
        if false {
            var v10 uint16 = 42
            fmt.Println(fi9, fi9)
            _ = v10
        } else {
            switch (fi9 * (fi9 % (((fi9 / ((fi9) | 1))) | 1))) {
            case 32767:
                if (((1.5 - 3.14159) * 0.0) < 1e10) {
                    switch ("a" + ("a" + ("hello" + "hello"))) {
                    case "foo bar":
                        switch "" {
                        case "\n\t":
                            switch ((((("" + "") + "foo bar") + ("hello" + "foo bar")) + ("foo bar" + "a")) + "a") {
                            case "\n\t":
                                v11 := GenS3{F0: 42}
                                _ = v11
                            }
                        case "hello":
                            var v12 uint32 = 1
                            _ = v12
                        default:
                            var v13 uint64 = 488
                            _ = v13
                        }
                    case "a":
                        switch "foo bar" {
                        case "foo bar":
                            v14 := GenS4{F0: "foo bar"}
                            _ = v14
                        default:
                            var v15 string = ("hello" + "")
                            _ = v15
                        }
                    case "\n\t":
                        var v16 float32 = 0.0
                        _ = v16
                    }
                } else {
                    var v17 string = "hello"
                    switch v17 {
                    case "a":
                        v17 = v17
                    case "\n\t":
                        if true {
                            fmt.Println(fi9)
                            var v18 bool = (2 != 4294967295)
                            fmt.Println(v18, v18, v18)
                            _ = v18
                        } else {
                            v19 := []int{fi9, (^(fi9))}
                            v20 := GenS1{F0: "hello", F1: 128}
                            _ = v19
                            _ = v20
                        }
                    case "":
                        if true {
                            v21 := GenS4{F0: v17}
                            for fi22 := 0; fi22 < 14; fi22++ {
                                v17 = v17
                                v23 := []int{(fi22 - fi22)}
                                v17 = v17
                                _ = fi22
                                _ = v23
                            }
                            for fi24 := 0; fi24 < 19; fi24++ {
                                v25 := GenS1{F0: ("foo bar" + "a"), F1: 0}
                                _ = fi24
                                _ = v25
                            }
                            _ = v21
                        } else {
                            var v26 uint32 = 62
                            _ = v26
                        }
                    }
                    fmt.Println(v17, fi9)
                    _ = v17
                }
            case -570:
                var v27 int = -718
                _ = v27
            case 2147483647:
                if (127 <= -128) {
                    v28 := GenS5{F0: 127, F1: 42, F2: 488}
                    var v29 uint16 = 180
                    v29 -= v29
                    _ = v28
                    _ = v29
                }
            }
            var v30 uint16 = 255
            switch "" {
            case "":
                v30 = v30
            case "hello":
                v30 += v30
            case "foo bar":
                fmt.Println(v30, v30)
            default:
                var v31 int32 = -1
                _ = v31
            }
            _ = v30
        }
        var v32 int16 = 1000
        var v33 uint64 = 18446744073709551615
        _ = fi9
        _ = v32
        _ = v33
    }
    if true {
        var v34 uint32 = 381
        _ = v34
    }
    var v35 uint8 = 128
    _ = v35
    return 1
}

func main() {
    var v1 uint64 = 1
    v2 := []string{}
    v3 := make(map[uint32]string)
    var v4 int8 = -25
    var v5 uint32 = uint32(v1)
    v36 := GenS3{F0: genFn1()}
    fmt.Println(v5, v5)
    v4 = v4
    v2 = append(v2, "\n\t")
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    for fi37 := 0; fi37 < 13; fi37++ {
        if (18446744073709551615 != uint(v5)) {
            fmt.Println(v5)
            if false {
                v2 = append(v2, "foo bar")
                v1 += (v1 | v1)
                fmt.Println(v5)
            } else {
                var v38 uint8 = 2
                _ = v38
            }
        } else {
            v39 := GenS6{F0: 210, F1: 9223372036854775807, F2: (-(v4))}
            _ = v39
        }
        _ = fi37
    }
    v40 := GenS4{F0: (((((("foo bar" + "\n\t") + "foo bar") + "a") + "a") + ("\n\t" + "a")) + "hello")}
    if (!false) {
        _ = len(v3)
    } else {
        fmt.Println(v4)
    }
    switch "\n\t" {
    case "foo bar":
        v41 := GenS7{F0: 1, F1: (("foo bar" + ("foo bar" + ("\n\t" + ("" + ("\n\t" + "hello"))))) + ""), F2: 0}
        _ = v41
    case "hello":
        v1 -= (^(v1))
    case "":
        for fi42 := 0; fi42 < 18; fi42++ {
            if true {
                v4 += (v4 - (v4 ^ v4))
            } else {
                var v43 bool = ((!(483 == 505)) || true)
                _ = v43
            }
            v5 -= v5
            v3[(v5 % ((39) | 1))] = ("\n\t" + "a")
            _ = fi42
        }
    default:
        fmt.Println(v4)
    }
    v4 = -128
    for fi44 := 0; fi44 < 20; fi44++ {
        switch (fi44 / (((fi44 & (fi44 + (fi44 % (((fi44 & int(v1))) | 1))))) | 1)) {
        case 32768:
            delete(v3, v5)
        case 536:
            for fi45 := 0; fi45 < 17; fi45++ {
                _ = len(v2)
                fmt.Println(fi44, v5, v4)
                _ = fi45
            }
        case -639:
            v46 := v3[v5]
            _ = v46
        }
        v5 -= (v5 + 2)
        _ = fi44
    }
    switch "\n\t" {
    case "foo bar":
        if (!true) {
            var v47 uint64 = 1
            var v48 uint8 = 1
            _ = v47
            _ = v48
        } else {
            switch -668 {
            case -395:
                v49 := GenS5{F0: 255, F1: 660, F2: 965}
                _ = v49
            case -2147483648:
                v5 += v5
            case 42:
                v4 -= genFn1()
            default:
                v50 := GenS3{F0: (v4 & -128)}
                _ = v50
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v40
    _ = v5
}
