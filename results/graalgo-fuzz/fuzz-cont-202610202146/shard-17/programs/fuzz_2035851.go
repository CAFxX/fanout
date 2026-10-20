package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int32; F1 bool; F2 int64 }
type GenS3 struct { F0 uint32; F1 uint8; F2 bool }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint; F1 int64; F2 int }
type GenS7 struct { F0 int8; F1 uint16 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int32; F1 uint64; F2 bool }
type GenS10 struct { F0 int; F1 bool }
type GenS11 struct { F0 int; F1 string; F2 int64 }
type GenS12 struct { F0 int; F1 int16 }
func genFn0(p0 string) int64 {
    v4 := GenS1{F0: 789, F1: 245}
    v5 := GenS1{F0: 1, F1: 255}
    v6 := []int{}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 bool, p1 uint64, p2 int32) uint32 {
    for fi7 := 0; fi7 < 15; fi7++ {
        fmt.Println(p2)
        var v8 uint32 = 65535
        if p0 {
            switch (fi7 | fi7) {
            case 0:
                var v9 uint32 = uint32(p1)
                _ = v9
            default:
                if false {
                    fmt.Println(fi7)
                    for fi10 := 0; fi10 < 13; fi10++ {
                        v11 := []int{(fi7 / ((fi7) | 1))}
                        if p0 {
                            var v12 float32 = (1.5 / -2.25)
                            v13 := GenS1{F0: (v8 ^ 0), F1: 156}
                            v14 := GenS2{F0: p2, F1: p0, F2: 9223372036854775807}
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                        _ = fi10
                        _ = v11
                    }
                    if p0 {
                        var v15 int64 = genFn0(("a" + ("hello" + ("a" + "\n\t"))))
                        _ = v15
                    }
                }
            }
        } else {
            switch 42 {
            case -627:
                v8 -= (v8 >> ((4294967295) & 63))
            case 1:
                p0 = p0
            case 9223372036854775807:
                p0 = false
            default:
                var v16 uint = 530
                _ = v16
            }
            v17 := GenS3{F0: 205, F1: 1, F2: p0}
            for fi18 := 0; fi18 < 13; fi18++ {
                v19 := GenS1{F0: (v8 / ((v8) | 1)), F1: 1}
                v8 -= (v8 ^ v8)
                var v20 int64 = 9223372036854775807
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = v17
        }
        _ = fi7
        _ = v8
    }
    switch "\n\t" {
    case "foo bar":
        v21 := []int{-1, 2147483647}
        _ = v21
    case "":
        var v22 uint64 = 1
        _ = v22
    case "\n\t":
        v23 := []int{}
        _ = v23
    }
    _ = p0
    _ = p1
    _ = p2
    return 740
}

func main() {
    var v1 float64 = 1.5
    v2 := []string{"hello", "\n\t"}
    v3 := []string{("a" + "a"), "", "\n\t", ("a" + "hello")}
    if (42 >= 2147483647) {
        v1 -= v1
        var v24 float32 = (-2.25 * 3.14159)
        if (920 >= 1) {
            for fi25 := 0; fi25 < 18; fi25++ {
                v26 := GenS4{F0: 9223372036854775807, F1: 0}
                switch ("hello" + "a") {
                case "":
                    fmt.Println(v24, v24)
                case "hello":
                    var v27 string = "a"
                    _ = v27
                case "\n\t":
                    if ((968 != 2) || false) {
                        _ = len(v3)
                    } else {
                        var v28 int8 = -1
                        _ = v28
                    }
                }
                _ = fi25
                _ = v26
            }
            for fi29 := 0; fi29 < 15; fi29++ {
                var v30 uint32 = 0
                var v31 uint64 = 2
                _ = fi29
                _ = v30
                _ = v31
            }
            if false {
                var v32 uint8 = 0
                if (!true) {
                    for fi33 := 0; fi33 < 10; fi33++ {
                        v32 = 2
                        _ = fi33
                    }
                    switch -2147483648 {
                    case -907:
                        var v34 int8 = 0
                        _ = v34
                    case -561:
                        for fi35 := 0; fi35 < 15; fi35++ {
                            v32 = 103
                            v36 := GenS2{F0: 669, F1: false, F2: 9223372036854775807}
                            v37 := GenS5{F0: 616}
                            _ = fi35
                            _ = v36
                            _ = v37
                        }
                    default:
                        v24 -= (((v24 * (v24 * (v24 + 0.1))) / (v24 - 1e10)) / v24)
                    }
                    for fi38 := 0; fi38 < 2; fi38++ {
                        for fi39 := 0; fi39 < 16; fi39++ {
                            v32 = 0
                            if (1 >= -826) {
                                switch (fi39 + fi39) {
                                case 2147483647:
                                    v32 = (v32 % ((42) | 1))
                                }
                                v40 := GenS5{F0: 784}
                                v24 += (1.5 * 1e10)
                                _ = len(v3)
                                if len(v3) > 0 {
                                    v3[0] = ("a" + ("a" + ("hello" + (("" + "foo bar") + "\n\t"))))
                                }
                                _ = v40
                            } else {
                                v41 := GenS4{F0: genFn0("foo bar"), F1: 2}
                                var v42 uint8 = 127
                                v43 := GenS6{F0: 1, F1: genFn0(("a" + "\n\t")), F2: (fi39 * 1)}
                                if ((2 == 826) || false) {
                                    fmt.Println(v42)
                                    v44 := GenS4{F0: 664, F1: 4294967295}
                                    _ = v44
                                } else {
                                    switch "a" {
                                    case "\n\t":
                                        var v45 uint64 = 0
                                        _ = v45
                                    default:
                                        v46 := GenS6{F0: 298, F1: genFn0("hello"), F2: (fi39 - fi39)}
                                        _ = v46
                                    }
                                }
                                _ = v41
                                _ = v42
                                _ = v43
                            }
                            switch 1 {
                            case 4294967295:
                                var v47 float64 = float64(float32(v1))
                                _ = v47
                            case -1:
                                v32 += (v32 * 0)
                            case -770:
                                var v48 uint16 = 467
                                _ = v48
                            default:
                                switch fi39 {
                                case 1:
                                    var v49 int16 = 822
                                    _ = v49
                                default:
                                    v1 += v1
                                }
                            }
                            _ = fi39
                        }
                        if true {
                            var v50 string = "foo bar"
                            _ = v50
                        }
                        v51 := GenS7{F0: -77, F1: 2}
                        _ = fi38
                        _ = v51
                    }
                }
                _ = v32
            } else {
                fmt.Println(v24, v1)
                switch -9223372036854775808 {
                case -254:
                    for fi52 := 0; fi52 < 12; fi52++ {
                        for fi53 := 0; fi53 < 11; fi53++ {
                            v54 := GenS8{F0: "hello"}
                            v55 := GenS9{F0: 0, F1: 1, F2: true}
                            v56 := GenS10{F0: (^(fi53)), F1: false}
                            _ = fi53
                            _ = v54
                            _ = v55
                            _ = v56
                        }
                        _ = fi52
                    }
                case 128:
                    switch "foo bar" {
                    case "":
                        var v57 int16 = 877
                        _ = v57
                    case "a":
                        var v58 int64 = 0
                        _ = v58
                    case "hello":
                        for fi59 := 0; fi59 < 6; fi59++ {
                            var v60 int64 = 861
                            v60 += (v60 / (((v60 << ((42) & 63))) | 1))
                            _ = fi59
                            _ = v60
                        }
                    }
                }
                v1 += v1
            }
        }
        _ = v24
    }
    v61 := GenS11{F0: 2147483647, F1: "a", F2: 42}
    fmt.Println(v1)
    v1 -= 3.14159
    var v62 uint64 = 1
    for fi63 := 0; fi63 < 5; fi63++ {
        var v64 int64 = 1
        for fi65 := 0; fi65 < 6; fi65++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v64 += v64
            _ = fi65
        }
        _ = fi63
        _ = v64
    }
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    v2 = append(v2, "\n\t")
    if true {
        for fi66 := 0; fi66 < 1; fi66++ {
            v67 := GenS4{F0: int64(fi66), F1: v62}
            var v68 bool = (genFn1(true, (v62 / ((v62) | 1)), -2147483648) <= 134)
            _ = fi66
            _ = v67
            _ = v68
        }
        _ = len(v3)
    }
    v62 = (v62 | v62)
    for fi69 := 0; fi69 < 13; fi69++ {
        var v70 uint32 = 65535
        _ = fi69
        _ = v70
    }
    var v71 int16 = 0
    var v72 bool = (false || false)
    v73 := GenS12{F0: 600, F1: (v71 & (v71 % ((68) | 1)))}
    switch 128 {
    case 128:
        v71 = (v71 | v71)
    }
    v1 = (v1 * ((float64(1e10) + 1.5) / v1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v62:", v62)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v61
    _ = v62
    _ = v71
    _ = v72
    _ = v73
}
