package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 uint; F1 uint; F2 string }
type GenS4 struct { F0 int; F1 bool; F2 uint64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int; F1 uint }
func genFn0() uint64 {
    var v7 float64 = -2.25
    if true {
        var v8 int64 = 93
        _ = v8
    }
    var v9 float32 = 3.14159
    for fi10 := 0; fi10 < 8; fi10++ {
        switch ("foo bar" + ("\n\t" + "a")) {
        case "":
            fmt.Println(fi10, v7)
        case "hello":
            v7 = ((0.0 / v7) / 1.0)
        default:
            var v11 int8 = -20
            _ = v11
        }
        if true {
            v7 = 1.5
            var v12 int64 = -663
            _ = v12
        }
        for fi13 := 0; fi13 < 3; fi13++ {
            fmt.Println(v7)
            var v14 uint = 82
            switch (fi10 / ((fi13) | 1)) {
            case 390:
                if (-1 > -629) {
                    v9 = (1e10 + 0.1)
                    v7 += float64(0.1)
                    var v15 bool = (fi13 > fi10)
                    var v16 uint64 = 42
                    _ = v15
                    _ = v16
                } else {
                    var v17 uint16 = 1
                    switch "hello" {
                    case "a":
                        fmt.Println(v14, v7, v9)
                    case "":
                        v14 = (v14 ^ (v14 % ((889) | 1)))
                    default:
                        if false {
                            switch "foo bar" {
                            case "":
                                if true {
                                    var v18 uint8 = 42
                                    v19 := GenS1{F0: 42, F1: 4294967295}
                                    v20 := []int{9223372036854775807}
                                    var v21 uint = v14
                                    v18 = (v18 / ((v18) | 1))
                                    _ = v18
                                    _ = v19
                                    _ = v20
                                    _ = v21
                                }
                            }
                            v22 := []int{(fi10 & fi10), fi13, (-(fi13)), fi10}
                            if (!false) {
                                v22 = append(v22, fi10)
                                v9 = 0.0
                                if false {
                                    _ = len(v22)
                                    v7 += (v7 * (v7 + v7))
                                }
                                v23 := GenS1{F0: 1, F1: 1}
                                _ = v23
                            } else {
                                switch fi10 {
                                case -397:
                                    var v24 int = (fi13 & fi13)
                                    _ = v24
                                case 2147483647:
                                    fmt.Println(v7, v7, fi10)
                                case -188:
                                    if len(v22) >= 2 {
                                        v22 = v22[1:len(v22)]
                                    }
                                default:
                                    var v25 uint8 = 128
                                    _ = v25
                                }
                            }
                            v9 += v9
                            _ = v22
                        }
                    }
                    _ = v17
                }
            case 880:
                switch "foo bar" {
                case "\n\t":
                    v9 -= (v9 - -2.25)
                case "hello":
                    v9 = v9
                    v26 := GenS1{F0: 4294967295, F1: uint64(v14)}
                    _ = v26
                case "foo bar":
                    v27 := []int{32767, 0}
                    _ = v27
                }
            case 32767:
                var v28 uint = 4294967295
                _ = v28
            }
            v29 := GenS1{F0: 0, F1: 635}
            v9 = 3.14159
            _ = fi13
            _ = v14
            _ = v29
        }
        _ = fi10
    }
    _ = v7
    _ = v9
    return 42
}
func genFn1(p0 uint32, p1 string, p2 uint32) string {
    var v30 string = ((p1 + ((("\n\t" + "hello") + ("foo bar" + "\n\t")) + (p1 + "hello"))) + p1)
    p0 = (^(p2))
    v31 := GenS2{F0: 127, F1: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v30
    _ = v31
    return ((v30 + "foo bar") + v30)
}

func main() {
    var v1 int64 = -657
    v2 := []int{2147483647, -283}
    v3 := []string{"foo bar", "\n\t"}
    var v4 float32 = 3.14159
    v5 := make(map[int]uint32)
    var v6 int16 = -182
    v6 -= v6
    v32 := GenS1{F0: genFn0(), F1: 0}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi33 := 0; fi33 < 14; fi33++ {
        var v34 int = fi33
        switch "" {
        case "a":
            var v35 float64 = 1e10
            _ = v35
        case "hello":
            v36 := v5[(fi33 | -1)]
            _ = v36
        case "":
            for fi37 := 0; fi37 < 7; fi37++ {
                v38 := v5[32767]
                v39 := GenS1{F0: genFn0(), F1: 0}
                _ = fi37
                _ = v38
                _ = v39
            }
        default:
            switch "hello" {
            case "\n\t":
                for fi40 := 0; fi40 < 18; fi40++ {
                    if false {
                        v41 := GenS1{F0: 42, F1: 2}
                        _ = v41
                    } else {
                        v42 := GenS3{F0: 65535, F1: 2, F2: "a"}
                        fmt.Println(v1)
                        _ = v42
                    }
                    _ = fi40
                }
            case "hello":
                v4 = v4
            case "a":
                switch "foo bar" {
                case "hello":
                    v43 := GenS1{F0: 1, F1: 2}
                    _ = v43
                default:
                    var v44 float32 = v4
                    _ = v44
                }
            }
        }
        _ = fi33
        _ = v34
    }
    if false {
        var v45 float32 = 1.5
        var v46 uint = 65535
        _ = v45
        _ = v46
    } else {
        v6 += (v6 / ((v6) | 1))
        var v47 uint16 = 42
        var v48 uint = 1
        _ = v47
        _ = v48
    }
    var v49 uint64 = 2
    delete(v5, int(v49))
    var v50 string = "a"
    if false {
        var v51 uint64 = v49
        _ = v51
    } else {
        v1 = (v1 | (v1 * -974))
        var v52 float32 = v4
        for fi53 := 0; fi53 < 9; fi53++ {
            v54 := GenS4{F0: fi53, F1: false, F2: v49}
            _ = fi53
            _ = v54
        }
        _ = v52
    }
    if (true && false) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v55 := GenS5{F0: 54}
        if (!(!(42 <= 255))) {
            v1 -= v1
            v56 := GenS6{F0: 0, F1: 1}
            _ = v56
        } else {
            switch (v50 + (v50 + "\n\t")) {
            case "hello":
                switch "hello" {
                case "\n\t":
                    fmt.Println(v49)
                case "":
                    _ = len(v5)
                default:
                    for fi57 := 0; fi57 < 20; fi57++ {
                        v58 := GenS2{F0: 2, F1: (v49 % ((1) | 1))}
                        v59 := GenS5{F0: 127}
                        v5[(fi57 % (((fi57 + (fi57 + (fi57 + 4294967295)))) | 1))] = 1
                        _ = fi57
                        _ = v58
                        _ = v59
                    }
                }
            case "foo bar":
                v60 := GenS3{F0: 968, F1: 42, F2: v50}
                _ = v60
            case "\n\t":
                switch v50 {
                case "\n\t":
                    var v61 uint = 18446744073709551615
                    _ = v61
                case "a":
                    for fi62 := 0; fi62 < 18; fi62++ {
                        v63 := GenS2{F0: 255, F1: v49}
                        var v64 int64 = (v1 | (v1 * v1))
                        _ = fi62
                        _ = v63
                        _ = v64
                    }
                }
            }
            for fi65 := 0; fi65 < 6; fi65++ {
                if true {
                    var v66 float64 = float64(v4)
                    _ = v66
                } else {
                    var v67 int16 = v6
                    v68 := GenS5{F0: 42}
                    _ = v67
                    _ = v68
                }
                v4 = v4
                for fi69 := 0; fi69 < 19; fi69++ {
                    var v70 int32 = 967
                    switch (-(fi65)) {
                    case -337:
                        if (0.0 <= 0.0) {
                            v6 = 42
                        }
                    }
                    _ = fi69
                    _ = v70
                }
                _ = fi65
            }
            if (42 >= 1) {
                v71 := v5[0]
                _ = v71
            }
        }
        _ = v55
    }
    var v72 float64 = 0.1
    var v73 int16 = (v6 % ((-179) | 1))
    var v74 int64 = (v1 % ((v1) | 1))
    for fi75 := 0; fi75 < 1; fi75++ {
        fmt.Println(v73)
        var v76 uint8 = 42
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi75
        _ = v76
    }
    v77 := GenS4{F0: 4294967295, F1: false, F2: (v49 | (v49 % (((v49 % (((v49 / ((genFn0()) | 1))) | 1))) | 1)))}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    fmt.Println("v72:", v72)
    fmt.Println("v73:", v73)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v6
    _ = v72
    _ = v73
    _ = v74
    _ = v77
}
