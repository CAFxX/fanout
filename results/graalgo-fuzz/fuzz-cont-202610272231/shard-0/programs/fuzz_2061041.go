package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int16; F1 int32; F2 uint64 }
func genFn0(p0 int16, p1 bool) int {
    if false {
        p0 += 628
        var v6 int = 263
        _ = v6
    } else {
        v7 := GenS1{F0: 2, F1: "foo bar"}
        v8 := GenS1{F0: 255, F1: (("\n\t" + "a") + ("a" + "\n\t"))}
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        fmt.Println(p0)
        switch (fi9 << ((116) & 63)) {
        case 708:
            p0 = (p0 + (p0 % ((p0) | 1)))
        case 994:
            p1 = p1
        default:
            for fi10 := 0; fi10 < 4; fi10++ {
                switch (fi10 % ((fi9) | 1)) {
                case 32767:
                    p0 = 42
                case -360:
                    v11 := []int{}
                    _ = v11
                }
                v12 := []int{(fi9 << ((581) & 63)), (fi9 >> ((540) & 63)), (fi9 + (fi10 | fi9))}
                _ = fi10
                _ = v12
            }
        }
        for fi13 := 0; fi13 < 15; fi13++ {
            var v14 int8 = -128
            if p1 {
                if p1 {
                    if true {
                        if (0.1 == 3.14159) {
                            switch -286 {
                            case 2147483647:
                                p0 += (p0 & 1000)
                            case 127:
                                switch fi13 {
                                case 2147483647:
                                    var v15 int32 = 1
                                    _ = v15
                                case -2147483648:
                                    var v16 uint32 = 828
                                    _ = v16
                                default:
                                    switch "\n\t" {
                                    case "a":
                                        var v17 uint64 = 42
                                        _ = v17
                                    case "":
                                        v18 := []int{(fi13 >> ((18446744073709551615) & 63)), 0, 9223372036854775807}
                                        _ = v18
                                    case "foo bar":
                                        var v19 int8 = -1
                                        _ = v19
                                    default:
                                        v20 := GenS1{F0: 128, F1: ("\n\t" + "a")}
                                        _ = v20
                                    }
                                }
                            }
                            if false {
                                if (true || p1) {
                                    var v21 int = (fi9 - -159)
                                    _ = v21
                                } else {
                                    v14 = (-(v14))
                                    p0 += (p0 % ((1) | 1))
                                    v22 := GenS1{F0: 2, F1: "hello"}
                                    v23 := []int{fi13, fi13, (fi13 >> ((4294967295) & 63)), fi13}
                                    p1 = ("" == "")
                                    if len(v23) >= 2 {
                                        v23 = v23[1:len(v23)]
                                    }
                                    var v24 int = (fi9 >> ((970) & 63))
                                    _ = v22
                                    _ = v23
                                    _ = v24
                                }
                                var v25 uint = 855
                                _ = v25
                            }
                            if p1 {
                                var v26 int16 = p0
                                _ = v26
                            }
                        }
                        switch "foo bar" {
                        case "foo bar":
                            v27 := GenS1{F0: 64, F1: "a"}
                            _ = v27
                        case "hello":
                            p0 = (p0 - (p0 | (^(p0))))
                        }
                    }
                    v28 := GenS1{F0: 1, F1: (("a" + "") + (("a" + "\n\t") + "\n\t"))}
                    _ = v28
                }
                fmt.Println(p0, fi13)
            } else {
                if p1 {
                    p1 = p1
                    v14 -= (v14 ^ v14)
                    p0 -= (p0 >> ((255) & 63))
                } else {
                    p0 -= 608
                    p1 = (!((255 < 843) && (!(!(42 > 42)))))
                    if (!p1) {
                        v29 := []int{}
                        fmt.Println(fi9, p0)
                        _ = v29
                    }
                }
            }
            _ = fi13
            _ = v14
        }
        _ = fi9
    }
    p0 = p0
    _ = p0
    _ = p1
    return 128
}

func main() {
    v1 := []string{"\n\t", "\n\t"}
    v2 := []uint32{2147483648}
    var v3 int16 = 1000
    v4 := make(map[int]bool)
    var v5 int8 = 1
    if ((("" == "hello") || true) || true) {
        var v30 uint32 = 290
        _ = v30
    } else {
        v31 := GenS1{F0: 116, F1: ""}
        _ = v31
    }
    _ = len(v2)
    if true {
        var v32 int64 = -9223372036854775808
        if false {
            if ((1e10 / (1e10 * -2.25)) != (0.0 * 0.1)) {
                v5 = (v5 ^ v5)
                for fi33 := 0; fi33 < 2; fi33++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    fmt.Println(v5, v5)
                    _ = fi33
                }
            } else {
                var v34 float32 = (1.5 + 0.1)
                if len(v1) > 0 {
                    v1[0] = ("a" + "a")
                }
                var v35 uint32 = 350
                _ = v34
                _ = v35
            }
        }
        v36 := GenS2{F0: v3, F1: 0, F2: 2}
        _ = v32
        _ = v36
    } else {
        if true {
            for fi37 := 0; fi37 < 12; fi37++ {
                v4[(fi37 & fi37)] = true
                fmt.Println(v5, fi37)
                var v38 int = (fi37 | genFn0(v3, true))
                _ = fi37
                _ = v38
            }
        }
    }
    if len(v1) > 0 {
        v1[0] = ("hello" + "")
    }
    var v39 int = 42
    var v40 float32 = (1e10 * 3.14159)
    v5 = (v5 + (v5 / ((100) | 1)))
    var v41 uint8 = 1
    if false {
        var v42 bool = false
        _ = v42
    } else {
        if true {
            fmt.Println(v39, v41, v41)
        } else {
            var v43 uint32 = 577
            for fi44 := 0; fi44 < 16; fi44++ {
                var v45 uint8 = v41
                if ((v3 | -1) > -32768) {
                    delete(v4, (-(v39)))
                    v2 = append(v2, v43)
                    fmt.Println(v40, v3, fi44)
                }
                _ = fi44
                _ = v45
            }
            switch (-(v39)) {
            case 174:
                var v46 int = 4294967295
                _ = v46
            case -9223372036854775808:
                _ = len(v4)
            case 128:
                v47 := GenS1{F0: 2, F1: "\n\t"}
                _ = v47
            default:
                v3 = (v3 ^ -756)
            }
            _ = v43
        }
    }
    v40 -= v40
    switch (v39 - (v39 - (-(v39)))) {
    case -9223372036854775808:
        v41 -= 42
    case 128:
        fmt.Println(v41, v39)
    case 32768:
        var v48 int32 = 0
        _ = v48
    }
    fmt.Println(v40, v3, v41)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
