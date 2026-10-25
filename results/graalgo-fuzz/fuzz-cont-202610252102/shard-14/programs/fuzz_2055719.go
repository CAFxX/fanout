package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 int8; F2 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 uint64; F1 uint16 }
type GenS6 struct { F0 uint; F1 bool; F2 int }
func genFn0(p0 string, p1 uint64, p2 uint32) uint {
    if (0 == 127) {
        var v4 uint32 = p2
        _ = v4
    } else {
        switch p0 {
        case "\n\t":
            var v5 uint = 612
            _ = v5
        case "hello":
            p1 += p1
        }
        switch p0 {
        case "\n\t":
            v6 := GenS1{F0: 2}
            _ = v6
        case "foo bar":
            v7 := GenS2{F0: 2, F1: 127, F2: 239}
            _ = v7
        case "a":
            fmt.Println(p2, p2, p0)
        }
        p1 = (p1 << ((uint(p1)) & 63))
    }
    p0 = (((p0 + p0) + (p0 + (p0 + (p0 + p0)))) + p0)
    p1 += 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    v1 := []string{"\n\t", ("foo bar" + ""), ""}
    v2 := []int{-812, -698, -1}
    v3 := make(map[int]int)
    v8 := GenS2{F0: 1, F1: 42, F2: 0}
    switch "\n\t" {
    case "a":
        if ((!true) && false) {
            switch "" {
            case "hello":
                var v9 int64 = -634
                _ = v9
            }
            var v10 float64 = (0.0 - 0.0)
            if ((!false) && (1 != 988)) {
                v10 += (v10 + 3.14159)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                v10 = v10
                v11 := GenS2{F0: 65535, F1: 1, F2: 0}
                v10 = 3.14159
                _ = v11
            }
            _ = v10
        } else {
            var v12 int8 = 127
            if (2147483647 > 32767) {
                fmt.Println(v12)
            }
            _ = v12
        }
    case "foo bar":
        if (true || true) {
            if len(v2) > 0 {
                v2[0] = 0
            }
            if (false && (!false)) {
                var v13 string = "foo bar"
                _ = v13
            } else {
                var v14 bool = ((!((-2.25 / 1.5) <= 1e10)) && false)
                if (v14 || false) {
                    if v14 {
                        switch "" {
                        case "hello":
                            var v15 uint = 905
                            _ = v15
                        default:
                            if len(v2) > 0 {
                                v2[0] = 264
                            }
                        }
                        var v16 uint64 = 4294967295
                        _ = v16
                    } else {
                        v14 = false
                        if ((!v14) || (100 > -5)) {
                            var v17 bool = (!(!v14))
                            v17 = (!v14)
                            var v18 uint = 830
                            _ = v17
                            _ = v18
                        }
                    }
                    if len(v1) > 0 {
                        v1[0] = "a"
                    }
                } else {
                    if (127 != 53) {
                        v3[1] = 405
                        if v14 {
                            v14 = ((127 == 255) || false)
                            for fi19 := 0; fi19 < 2; fi19++ {
                                v14 = v14
                                var v20 uint64 = 2
                                _ = fi19
                                _ = v20
                            }
                        } else {
                            fmt.Println(v14)
                            if v14 {
                                v21 := GenS3{F0: 42}
                                _ = v21
                            }
                            if (!(!false)) {
                                v14 = v14
                                var v22 int8 = 127
                                _ = v22
                            } else {
                                switch -620 {
                                case -751:
                                    fmt.Println(v14)
                                case 128:
                                    v14 = (-9223372036854775808 >= -939)
                                case -9223372036854775808:
                                    fmt.Println(v14)
                                default:
                                    v23 := GenS3{F0: 9223372036854775807}
                                    _ = v23
                                }
                            }
                        }
                    }
                    if (v14 && (357 == -909)) {
                        _ = len(v1)
                        v24 := v3[1]
                        _ = v24
                    }
                    for fi25 := 0; fi25 < 10; fi25++ {
                        var v26 int32 = -1
                        var v27 int32 = 2147483647
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                }
                fmt.Println(v14)
                _ = v14
            }
            var v28 uint32 = 10
            _ = v28
        } else {
            for fi29 := 0; fi29 < 6; fi29++ {
                for fi30 := 0; fi30 < 3; fi30++ {
                    v1 = append(v1, "")
                    v3[fi29] = (fi30 * (fi30 & fi29))
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi30
                }
                _ = fi29
            }
            switch "a" {
            case "foo bar":
                switch "" {
                case "a":
                    v31 := GenS3{F0: -536}
                    _ = v31
                case "foo bar":
                    var v32 uint16 = 65535
                    _ = v32
                case "\n\t":
                    var v33 bool = (!(true || false))
                    _ = v33
                }
            default:
                if true {
                    switch 4294967295 {
                    case 128:
                        if true {
                            var v34 int64 = -9223372036854775808
                            v34 += -9223372036854775808
                            _ = v34
                        } else {
                            var v35 bool = (!false)
                            _ = v35
                        }
                    case 4294967295:
                        var v36 uint16 = 1
                        _ = v36
                    default:
                        for fi37 := 0; fi37 < 4; fi37++ {
                            var v38 int64 = 0
                            _ = fi37
                            _ = v38
                        }
                    }
                } else {
                    if false {
                        var v39 int16 = 32767
                        v39 -= 1000
                        v39 -= v39
                        _ = v39
                    }
                }
            }
        }
    default:
        v40 := GenS1{F0: genFn0(("" + "hello"), 18446744073709551615, 42)}
        _ = v40
    }
    var v41 bool = (!(42 <= 32767))
    fmt.Println(v41)
    v41 = (163 < -1)
    v42 := GenS4{F0: 920, F1: true}
    if len(v1) > 0 {
        v1[0] = ("foo bar" + ("\n\t" + "foo bar"))
    }
    var v43 uint16 = 567
    switch ("\n\t" + "hello") {
    case "a":
        v44 := GenS1{F0: 453}
        _ = v44
    default:
        var v45 float64 = float64((0.1 - 0.0))
        _ = v45
    }
    var v46 int = 128
    switch 127 {
    case -793:
        v47 := GenS4{F0: 2, F1: true}
        _ = v47
    case 9223372036854775807:
        for fi48 := 0; fi48 < 9; fi48++ {
            v46 = (v46 << ((42) & 63))
            _ = fi48
        }
    default:
        fmt.Println(v41, v46, v46)
    }
    for fi49 := 0; fi49 < 19; fi49++ {
        v50 := GenS5{F0: 18446744073709551615, F1: 255}
        v51 := GenS6{F0: 4294967295, F1: ((0.1 <= -2.25) || v41), F2: (fi49 - 0)}
        _ = fi49
        _ = v50
        _ = v51
    }
    switch v46 {
    case 0:
        fmt.Println(v46)
    case -1:
        fmt.Println(v41)
    default:
        var v52 uint32 = 65535
        _ = v52
    }
    v53 := GenS4{F0: 0, F1: (true && false)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
    _ = v42
    _ = v43
    _ = v46
    _ = v53
    _ = v8
}
