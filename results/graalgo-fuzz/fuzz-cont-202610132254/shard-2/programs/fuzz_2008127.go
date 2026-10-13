package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 uint; F2 bool }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 uint16; F1 int32 }
func genFn0() uint {
    switch "hello" {
    case "foo bar":
        var v7 int8 = 1
        _ = v7
    }
    var v8 uint32 = 2147483647
    _ = v8
    return 4294967295
}
func genFn1(p0 bool, p1 uint) string {
    for fi9 := 0; fi9 < 19; fi9++ {
        for fi10 := 0; fi10 < 2; fi10++ {
            fmt.Println(p0)
            var v11 int = fi9
            _ = fi10
            _ = v11
        }
        fmt.Println(fi9, fi9)
        var v12 uint16 = 0
        _ = fi9
        _ = v12
    }
    fmt.Println(p0)
    var v13 float32 = 1.5
    var v14 bool = (878 != 32767)
    switch "foo bar" {
    case "a":
        v15 := GenS1{F0: "a"}
        _ = v15
    case "hello":
        for fi16 := 0; fi16 < 5; fi16++ {
            var v17 uint = p1
            v14 = false
            _ = fi16
            _ = v17
        }
    default:
        var v18 uint16 = 754
        _ = v18
    }
    _ = p0
    _ = p1
    _ = v13
    _ = v14
    return ("a" + ("a" + "a"))
}

func main() {
    v1 := []string{"", ("hello" + "a")}
    v2 := make(map[int]uint32)
    var v3 float64 = ((-2.25 + float64((1.5 * 0.0))) + 0.1)
    var v4 uint = 166
    v5 := make(map[string]bool)
    var v6 string = "a"
    if false {
        if true {
            v4 -= 0
        }
        if (127 > -382) {
            fmt.Println(v4, v6)
            v19 := GenS1{F0: v6}
            _ = v19
        }
        if true {
            if false {
                for fi20 := 0; fi20 < 17; fi20++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if (158 <= fi20) {
                        var v21 int32 = 42
                        switch fi20 {
                        case 2147483647:
                            if true {
                                v22 := GenS1{F0: ((((v6 + v6) + v6) + "") + (((v6 + "hello") + "foo bar") + v6))}
                                switch fi20 {
                                case 1:
                                    v23 := GenS1{F0: v6}
                                    _ = v23
                                default:
                                    v4 = v4
                                }
                                _ = v22
                            }
                        case 989:
                            var v24 uint8 = 1
                            _ = v24
                        case -2147483648:
                            var v25 int = fi20
                            _ = v25
                        default:
                            v26 := GenS2{F0: (v4 + v4)}
                            _ = v26
                        }
                        v2[fi20] = 42
                        _ = v21
                    } else {
                        v4 = v4
                    }
                    if (v4 <= genFn0()) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        fmt.Println(fi20, v6, v6)
                        v27 := GenS2{F0: (^(v4))}
                        _ = v27
                    }
                    _ = fi20
                }
            }
            var v28 string = (((v6 + ("foo bar" + (v6 + ""))) + "hello") + "hello")
            _ = v28
        }
    } else {
        var v29 float32 = (-2.25 * 0.1)
        v30 := GenS3{F0: 1, F1: (v4 - v4), F2: false}
        for fi31 := 0; fi31 < 9; fi31++ {
            var v32 float64 = (v3 / v3)
            for fi33 := 0; fi33 < 14; fi33++ {
                v3 += 0.0
                v29 -= float32(((v32 + (v3 * v32)) * v32))
                var v34 uint32 = 646
                _ = fi33
                _ = v34
            }
            _ = fi31
            _ = v32
        }
        _ = v29
        _ = v30
    }
    v4 = 1
    for fi35 := 0; fi35 < 6; fi35++ {
        if (false || (!true)) {
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
            v6 = v6
        } else {
            fmt.Println(fi35, v3)
            var v36 uint8 = 0
            _ = v36
        }
        _ = fi35
    }
    if true {
        for fi37 := 0; fi37 < 4; fi37++ {
            v6 = v6
            _ = fi37
        }
    }
    v3 -= v3
    v38 := GenS4{F0: v4, F1: genFn0()}
    for fi39 := 0; fi39 < 7; fi39++ {
        v2[fi39] = 2
        _ = len(v5)
        var v40 uint8 = 0
        _ = fi39
        _ = v40
    }
    if (1000 > 1) {
        if (true && false) {
            for fi41 := 0; fi41 < 7; fi41++ {
                for fi42 := 0; fi42 < 7; fi42++ {
                    v43 := GenS2{F0: v4}
                    var v44 float64 = v3
                    var v45 float64 = 0.0
                    var v46 bool = (9223372036854775807 == -1)
                    var v47 bool = false
                    _ = fi42
                    _ = v43
                    _ = v44
                    _ = v45
                    _ = v46
                    _ = v47
                }
                v48 := GenS1{F0: v6}
                var v49 uint64 = 18446744073709551615
                _ = fi41
                _ = v48
                _ = v49
            }
            v1 = append(v1, (v6 + v6))
            if (32767 != 42) {
                v50 := GenS5{F0: 42, F1: 0}
                if true {
                    v1 = append(v1, ("a" + ("a" + ("a" + ("hello" + (v6 + "\n\t"))))))
                    var v51 string = v6
                    switch 128 {
                    case 430:
                        _ = len(v1)
                    }
                    _ = v51
                }
                v52 := v2[127]
                _ = v50
                _ = v52
            } else {
                v6 = "hello"
                v3 = -2.25
                v3 += ((float64(0.1) + (0.0 + 1.5)) * v3)
            }
        }
        v4 -= v4
        var v53 float64 = -2.25
        _ = v53
    } else {
        if true {
            var v54 uint32 = 876
            fmt.Println(v6, v54)
            _ = len(v5)
            _ = v54
        }
        var v55 int32 = 2147483647
        _ = v55
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
