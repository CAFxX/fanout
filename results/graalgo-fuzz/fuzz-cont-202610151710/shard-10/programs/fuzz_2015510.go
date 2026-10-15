package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 bool; F1 uint }

func main() {
    v1 := []uint32{108, 360}
    var v2 int16 = -942
    var v3 uint = 4294967295
    v4 := make(map[string]string)
    v5 := make(map[string]uint32)
    v4[""] = ""
    _ = len(v1)
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 uint = v3
        var v8 uint16 = 65535
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 float64 = (1e10 * float64(1.5))
    if (732 != 9223372036854775807) {
        var v10 uint = 18446744073709551615
        _ = v10
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 float64 = v9
        var v13 float32 = ((((0.0 / 3.14159) - (float32(v12) + 1e10)) / 1.0) - 1.5)
        _ = fi11
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v15 float32 = (1.5 + -2.25)
        v3 = v3
        _ = fi14
        _ = v15
    }
    v9 -= (v9 * ((-2.25 * ((v9 - v9) - 3.14159)) - (v9 / 1e10)))
    if (!true) {
        v9 = 0.1
        v16 := GenS1{F0: v3, F1: 4294967295}
        switch -9223372036854775808 {
        case -721:
            for fi17 := 0; fi17 < 12; fi17++ {
                v3 = (v3 % ((v3) | 1))
                if true {
                    v18 := GenS1{F0: (v3 + v3), F1: 1}
                    _ = v18
                }
                var v19 uint32 = 65535
                _ = fi17
                _ = v19
            }
        default:
            for fi20 := 0; fi20 < 6; fi20++ {
                v21 := GenS1{F0: 402, F1: 1}
                var v22 uint = (v3 * (v3 ^ 1))
                if false {
                    if (-2.25 <= -2.25) {
                        var v23 int8 = 127
                        if false {
                            var v24 float64 = -2.25
                            for fi25 := 0; fi25 < 8; fi25++ {
                                switch int(v2) {
                                case -59:
                                    var v26 uint32 = 4294967295
                                    _ = v26
                                }
                                _ = fi25
                            }
                            _ = v24
                        }
                        _ = v23
                    } else {
                        v27 := GenS1{F0: v22, F1: 908}
                        v4[(("" + "hello") + "foo bar")] = "foo bar"
                        if true {
                            if ((false || false) || false) {
                                for fi28 := 0; fi28 < 16; fi28++ {
                                    v29 := GenS1{F0: (v3 << (((^(v3))) & 63)), F1: 60}
                                    var v30 uint8 = 0
                                    _ = fi28
                                    _ = v29
                                    _ = v30
                                }
                                var v31 uint64 = 4294967295
                                v9 -= float64(-2.25)
                                _ = v31
                            }
                        }
                        _ = v27
                    }
                    if (-1 < 9223372036854775807) {
                        v32 := GenS2{F0: (v3 * v3), F1: 658}
                        var v33 int32 = -683
                        v34 := GenS2{F0: v22, F1: -652}
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                } else {
                    v3 += 18446744073709551615
                    for fi35 := 0; fi35 < 7; fi35++ {
                        var v36 uint32 = 2147483647
                        delete(v5, "")
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi35
                        _ = v36
                    }
                    v37 := GenS3{F0: 1, F1: v2}
                    _ = v37
                }
                _ = fi20
                _ = v21
                _ = v22
            }
        }
        _ = v16
    } else {
        if ((!(!(!true))) || (42 > 231)) {
            v38 := v5["hello"]
            if false {
                for fi39 := 0; fi39 < 8; fi39++ {
                    v2 = v2
                    switch "a" {
                    case "":
                        for fi40 := 0; fi40 < 12; fi40++ {
                            v41 := GenS3{F0: (v3 | v3), F1: 1}
                            _ = fi40
                            _ = v41
                        }
                    case "a":
                        fmt.Println(v3, v3)
                    case "foo bar":
                        v42 := GenS4{F0: false, F1: v3}
                        _ = v42
                    }
                    _ = fi39
                }
                v43 := GenS4{F0: true, F1: v3}
                v9 += 0.0
                _ = v43
            } else {
                var v44 float32 = 1.5
                var v45 int16 = v2
                v38 -= v38
                _ = v44
                _ = v45
            }
            _ = v38
        }
    }
    delete(v5, "\n\t")
    v1 = append(v1, 264)
    if true {
        switch "hello" {
        case "a":
            v46 := GenS4{F0: ((v3 | (v3 - (v3 & v3))) >= v3), F1: (v3 - v3)}
            _ = v46
        default:
            _ = len(v1)
        }
        if true {
            var v47 float64 = v9
            v48 := GenS3{F0: v3, F1: v2}
            v3 -= (v3 / ((255) | 1))
            _ = v47
            _ = v48
        } else {
            var v49 uint = v3
            _ = v49
        }
        var v50 string = (("a" + "hello") + (((("hello" + "foo bar") + "") + "") + "foo bar"))
        _ = v50
    } else {
        var v51 uint32 = 2
        _ = v51
    }
    fmt.Println(v2, v2, v9)
    var v52 int = -199
    for fi53 := 0; fi53 < 14; fi53++ {
        for fi54 := 0; fi54 < 3; fi54++ {
            var v55 int = (v52 - (fi54 * 9223372036854775807))
            var v56 int = v52
            if ("hello" < "a") {
                v57 := GenS2{F0: v3, F1: -9223372036854775808}
                switch v55 {
                case -2147483648:
                    var v58 float32 = float32((v9 - 3.14159))
                    _ = v58
                }
                _ = v57
            } else {
                v55 += fi54
            }
            v2 = (v2 | (v2 & (v2 * v2)))
            v59 := GenS2{F0: v3, F1: 46}
            _ = fi54
            _ = v55
            _ = v56
            _ = v59
        }
        _ = fi53
    }
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    v60 := GenS1{F0: (^(v3)), F1: 750}
    v5[""] = 2147483647
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v60
    _ = v9
}
