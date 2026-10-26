package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 string; F2 int8 }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32; F1 string; F2 int }
type GenS6 struct { F0 string }
func genFn0(p0 int16, p1 string, p2 int32) uint {
    var v4 uint16 = 2
    v5 := GenS1{F0: -100}
    v6 := []int{32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 255
}
func genFn1(p0 int8, p1 uint, p2 bool) string {
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 uint = (p1 >> ((65535) & 63))
        for fi9 := 0; fi9 < 14; fi9++ {
            v8 = (v8 | v8)
            _ = fi9
        }
        fmt.Println(fi7)
        _ = fi7
        _ = v8
    }
    v10 := []int{127, -483}
    for fi11 := 0; fi11 < 18; fi11++ {
        var v12 uint64 = 549
        var v13 uint32 = 2147483647
        switch "foo bar" {
        case "hello":
            switch (fi11 << ((35) & 63)) {
            case 2147483647:
                for fi14 := 0; fi14 < 12; fi14++ {
                    if p2 {
                        v15 := GenS1{F0: 42}
                        p2 = true
                        _ = v15
                    } else {
                        var v16 float32 = -2.25
                        p0 = (p0 * (p0 & 42))
                        _ = v16
                    }
                    if p2 {
                        if len(v10) >= 2 {
                            v10 = v10[1:len(v10)]
                        }
                        if true {
                            var v17 float64 = 3.14159
                            _ = v17
                        }
                    } else {
                        if (p1 > 1) {
                            if (uint8(p1) <= 148) {
                                var v18 uint16 = 846
                                _ = v18
                            }
                            if (-1 <= 214) {
                                fmt.Println(p0, p0, p1)
                                if (65535 == 0) {
                                    v19 := GenS1{F0: int64(p0)}
                                    p1 = p1
                                    p0 -= (p0 + p0)
                                    _ = v19
                                }
                            } else {
                                if ((!p2) || false) {
                                    var v20 int16 = 736
                                    switch "hello" {
                                    case "hello":
                                        switch fi11 {
                                        case 942:
                                            var v21 float32 = (((0.0 - 1.5) - 1.5) - float32(3.14159))
                                            _ = v21
                                        case 2147483647:
                                            if ((!((-2.25 - 1.5) <= 1e10)) && p2) {
                                                var v22 uint = (p1 + (p1 % (((p1 * 4294967295)) | 1)))
                                                var v23 uint32 = 2147483647
                                                v24 := GenS2{F0: (p0 | (p0 - (p0 << ((v22) & 63)))), F1: "\n\t", F2: 1}
                                                var v25 int8 = (^(p0))
                                                if len(v10) > 0 {
                                                    v10[0] = (^(fi11))
                                                }
                                                v26 := GenS2{F0: p0, F1: "\n\t", F2: (^(v25))}
                                                if ((1 == int64(v12)) || p2) {
                                                    v27 := GenS3{F0: -2147483648, F1: 42}
                                                    p0 = (v25 << (((v22 * (p1 / (((v22 & genFn0(v20, "foo bar", 42))) | 1)))) & 63))
                                                    _ = v27
                                                }
                                                fmt.Println(fi11, v23, fi11)
                                                v28 := GenS4{F0: 127}
                                                v29 := GenS5{F0: 42, F1: "foo bar", F2: (fi14 + fi11)}
                                                _ = len(v10)
                                                var v30 bool = false
                                                _ = v22
                                                _ = v23
                                                _ = v24
                                                _ = v25
                                                _ = v26
                                                _ = v28
                                                _ = v29
                                                _ = v30
                                            }
                                        case -2147483648:
                                            fmt.Println(fi11, p0)
                                        default:
                                            var v31 uint32 = (v13 + (v13 % ((v13) | 1)))
                                            _ = v31
                                        }
                                    case "\n\t":
                                        if false {
                                            var v32 uint = 255
                                            _ = v32
                                        }
                                    }
                                    if ((^(p1)) <= genFn0(v20, "a", -1)) {
                                        v12 = v12
                                        switch (fi11 / ((fi14) | 1)) {
                                        case 4294967295:
                                            _ = len(v10)
                                        default:
                                            v10 = append(v10, -1)
                                        }
                                        var v33 uint16 = 42
                                        _ = v33
                                    }
                                    _ = v20
                                }
                                switch 2147483647 {
                                case 42:
                                    v34 := GenS2{F0: (p0 % ((1) | 1)), F1: ("a" + "hello"), F2: -1}
                                    if len(v10) >= 2 {
                                        v10 = v10[1:len(v10)]
                                    }
                                    if len(v10) > 0 {
                                        v10[0] = 32768
                                    }
                                    _ = v34
                                default:
                                    var v35 float32 = (0.1 - float32(((((-2.25 / ((-2.25 - float64((1e10 * 0.0))) + 0.1)) / (((1e10 / 3.14159) / 0.1) * 1e10)) / float64((0.1 - -2.25))) - 1.5)))
                                    _ = v35
                                }
                                v10 = append(v10, fi11)
                                p1 += p1
                                switch "foo bar" {
                                case "":
                                    v36 := GenS5{F0: v13, F1: ("foo bar" + ""), F2: fi14}
                                    _ = v36
                                case "a":
                                    switch (-(fi14)) {
                                    case 1:
                                        var v37 uint16 = 938
                                        _ = v37
                                    default:
                                        p1 += p1
                                    }
                                case "hello":
                                    v10 = append(v10, fi11)
                                default:
                                    p0 += p0
                                }
                            }
                        }
                    }
                    fmt.Println(fi14, p2, fi14)
                    fmt.Println(v12, v12)
                    _ = fi14
                }
            case 32768:
                for fi38 := 0; fi38 < 3; fi38++ {
                    var v39 string = "\n\t"
                    v13 = (v13 ^ (v13 | v13))
                    v40 := GenS3{F0: fi11, F1: 0}
                    _ = fi38
                    _ = v39
                    _ = v40
                }
            }
        case "foo bar":
            fmt.Println(fi11)
        default:
            v10 = append(v10, (fi11 - (fi11 | (fi11 & (fi11 >> ((p1) & 63))))))
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return "hello"
}

func main() {
    var v1 uint64 = 796
    var v2 bool = true
    v3 := make(map[uint32]uint32)
    for fi41 := 0; fi41 < 12; fi41++ {
        v42 := []int{(fi41 & -648)}
        _ = fi41
        _ = v42
    }
    _ = len(v3)
    v43 := GenS1{F0: 487}
    var v44 int8 = 100
    fmt.Println(v1)
    for fi45 := 0; fi45 < 9; fi45++ {
        v2 = (980 != 246)
        _ = len(v3)
        _ = fi45
    }
    var v46 bool = (!(!(2 <= 118)))
    if (v44 < v44) {
        var v47 int64 = 1
        _ = v47
    } else {
        v1 = 18446744073709551615
        fmt.Println(v44, v44)
        v44 = -89
    }
    fmt.Println(v1, v46)
    v2 = (255 > 32767)
    for fi48 := 0; fi48 < 6; fi48++ {
        v46 = (!v2)
        fmt.Println(v2, v2)
        _ = fi48
    }
    _ = len(v3)
    fmt.Println(v46, v44)
    var v49 int64 = 1
    v50 := GenS6{F0: ""}
    var v51 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v49:", v49)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v44
    _ = v46
    _ = v49
    _ = v50
    _ = v51
}
