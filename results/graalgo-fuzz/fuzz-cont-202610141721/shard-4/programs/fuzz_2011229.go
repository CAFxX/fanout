package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 int32; F2 int8 }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 bool; F1 uint8; F2 uint64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int16; F1 uint }
type GenS7 struct { F0 uint; F1 int64 }
type GenS8 struct { F0 uint32; F1 uint8; F2 bool }
func genFn0(p0 int) int {
    var v4 uint16 = 575
    p0 -= (p0 + (^(p0)))
    _ = p0
    _ = v4
    return (p0 * p0)
}
func genFn1(p0 int) string {
    p0 -= (-(p0))
    for fi5 := 0; fi5 < 6; fi5++ {
        if true {
            fmt.Println(p0, p0)
            v6 := []int{(p0 * p0)}
            switch (fi5 * fi5) {
            case 0:
                for fi7 := 0; fi7 < 2; fi7++ {
                    _ = len(v6)
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    var v8 int32 = 0
                    v9 := GenS1{F0: 0}
                    _ = fi7
                    _ = v8
                    _ = v9
                }
            case 127:
                if false {
                    var v10 float32 = ((((0.1 + 3.14159) * float32(0.1)) + 1e10) + 3.14159)
                    var v11 int8 = 42
                    _ = v10
                    _ = v11
                }
            }
            _ = v6
        } else {
            switch ("foo bar" + "") {
            case "":
                v12 := GenS2{F0: (100 == 127), F1: 2147483647, F2: -1}
                _ = v12
            case "hello":
                var v13 int32 = 0
                _ = v13
            case "\n\t":
                v14 := GenS2{F0: (true || ((1e10 + 1e10) <= 0.0)), F1: 163, F2: 107}
                _ = v14
            default:
                for fi15 := 0; fi15 < 17; fi15++ {
                    if false {
                        v16 := []int{genFn0(-956), (p0 & genFn0((fi15 - (fi5 & genFn0(p0))))), 127, (fi15 << ((2) & 63))}
                        var v17 int = fi5
                        v17 -= p0
                        _ = v16
                        _ = v17
                    }
                    v18 := []int{(^(fi5)), genFn0(p0), 32768, genFn0(p0)}
                    _ = fi15
                    _ = v18
                }
            }
        }
        var v19 string = ("foo bar" + ("" + "\n\t"))
        _ = fi5
        _ = v19
    }
    p0 += 9223372036854775807
    var v20 uint32 = 42
    _ = p0
    _ = v20
    return "a"
}

func main() {
    v1 := []string{("foo bar" + ("a" + "foo bar"))}
    v2 := make(map[uint32]uint32)
    var v3 int8 = 127
    switch "" {
    case "foo bar":
        switch ("" + ("a" + "hello")) {
        case "":
            var v21 uint16 = 42
            _ = v21
        case "hello":
            v22 := GenS3{F0: 769, F1: 128}
            _ = v22
        }
    case "a":
        if (648 != 0) {
            fmt.Println(v3)
            if true {
                if (1 <= 0) {
                    for fi23 := 0; fi23 < 7; fi23++ {
                        v3 += 1
                        var v24 int16 = 1
                        _ = fi23
                        _ = v24
                    }
                }
            }
        }
    default:
        var v25 uint = 2
        _ = v25
    }
    var v26 string = "foo bar"
    v27 := GenS4{F0: (-1 != -32768), F1: 127, F2: 2}
    var v28 uint = 1
    if (389 <= 42) {
        if (-1 <= 32768) {
            fmt.Println(v28)
            v29 := GenS3{F0: 1, F1: -2147483648}
            if true {
                if false {
                    v26 = (v26 + (v26 + "\n\t"))
                    v28 = 4294967295
                } else {
                    var v30 uint16 = 522
                    delete(v2, 2147483647)
                    _ = v30
                }
                for fi31 := 0; fi31 < 11; fi31++ {
                    switch v26 {
                    case "hello":
                        v26 = ((("a" + v26) + (("a" + "foo bar") + "a")) + (("hello" + v26) + v26))
                    case "a":
                        v32 := GenS1{F0: 1}
                        _ = v32
                    case "foo bar":
                        var v33 uint32 = 2
                        _ = v33
                    default:
                        for fi34 := 0; fi34 < 18; fi34++ {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi34
                        }
                    }
                    _ = fi31
                }
            } else {
                v35 := GenS5{F0: -573}
                fmt.Println(v3)
                if (("a" >= v26) || false) {
                    v36 := GenS6{F0: 1, F1: v28}
                    var v37 uint = (v28 / (((^(v28))) | 1))
                    _ = v36
                    _ = v37
                }
                _ = v35
            }
            _ = v29
        } else {
            var v38 int8 = (v3 & v3)
            _ = v38
        }
        if true {
            v3 = (v3 & v3)
        }
    }
    var v39 int16 = -1
    v39 = (v39 / ((v39) | 1))
    delete(v2, 4294967295)
    for fi40 := 0; fi40 < 2; fi40++ {
        if true {
            switch ("" + v26) {
            case "\n\t":
                switch v26 {
                case "a":
                    v26 = "\n\t"
                }
            }
            var v41 bool = (!(!false))
            switch (fi40 << ((2) & 63)) {
            case 767:
                for fi42 := 0; fi42 < 3; fi42++ {
                    if (false || (!v41)) {
                        v43 := GenS7{F0: (v28 % ((v28) | 1)), F1: 1}
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v44 string = "\n\t"
                        _ = v43
                        _ = v44
                    } else {
                        fmt.Println(v41)
                        fmt.Println(fi40, fi40)
                    }
                    v3 -= (-(v3))
                    _ = fi42
                }
            }
            _ = v41
        } else {
            fmt.Println(v28, v26, v3)
        }
        _ = fi40
    }
    if (0 < 494) {
        switch v26 {
        case "a":
            fmt.Println(v26)
        case "hello":
            v26 = v26
        case "\n\t":
            for fi45 := 0; fi45 < 17; fi45++ {
                if false {
                    v46 := GenS8{F0: 675, F1: 127, F2: true}
                    var v47 int = (^(fi45))
                    switch (v47 << ((v28) & 63)) {
                    case 127:
                        var v48 uint = v28
                        _ = v48
                    case -406:
                        var v49 uint16 = 42
                        _ = v49
                    }
                    _ = v46
                    _ = v47
                } else {
                    v28 = (v28 & (v28 >> (((v28 / ((v28) | 1))) & 63)))
                    for fi50 := 0; fi50 < 13; fi50++ {
                        switch fi45 {
                        case 0:
                            v26 = ("" + v26)
                        case 976:
                            if len(v1) > 0 {
                                v1[0] = (v26 + (("a" + "hello") + (((v26 + (v26 + ((v26 + "\n\t") + "hello"))) + (v26 + "\n\t")) + "foo bar")))
                            }
                        case 223:
                            _ = len(v1)
                        }
                        if (true || true) {
                            v3 -= -76
                            v51 := GenS1{F0: 194}
                            switch (fi50 * (fi50 + (fi45 >> (((v28 - 65535)) & 63)))) {
                            case 128:
                                var v52 int32 = 2147483647
                                _ = v52
                            default:
                                v1 = append(v1, ((((v26 + v26) + (("foo bar" + ("hello" + (v26 + ""))) + ((("a" + (v26 + (v26 + v26))) + v26) + v26))) + v26) + v26))
                            }
                            _ = v51
                        } else {
                            _ = len(v1)
                            delete(v2, 2)
                        }
                        _ = fi50
                    }
                }
                var v53 float32 = 0.1
                _ = fi45
                _ = v53
            }
        }
        v39 = 1
        fmt.Println(v28)
    }
    fmt.Println(v39, v39)
    fmt.Println(v26)
    v54 := GenS2{F0: (6 == 2), F1: -702, F2: v3}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v39
    _ = v54
}
