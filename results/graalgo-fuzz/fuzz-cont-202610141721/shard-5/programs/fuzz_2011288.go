package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 bool; F1 int64; F2 uint32 }
type GenS3 struct { F0 int; F1 bool; F2 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 string }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]int)
    var v3 uint64 = 2
    var v4 float64 = 0.1
    fmt.Println(v3, v3)
    v5 := GenS1{F0: -128, F1: v3}
    v6 := GenS1{F0: 54, F1: 483}
    v3 += 1
    var v7 float32 = ((((((3.14159 / (-2.25 * (1e10 / -2.25))) * 3.14159) * (1.5 / 1e10)) - (0.0 + 3.14159)) - -2.25) + 1e10)
    if (2 > 65535) {
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 float32 = (0.0 + (0.1 + 3.14159))
            _ = fi8
            _ = v9
        }
    }
    fmt.Println(v3, v4, v4)
    v7 -= 1.5
    if true {
        v3 -= 0
        switch "foo bar" {
        case "a":
            for fi10 := 0; fi10 < 4; fi10++ {
                for fi11 := 0; fi11 < 1; fi11++ {
                    switch "hello" {
                    case "":
                        v3 = (^(v3))
                    case "a":
                        v12 := GenS1{F0: 1, F1: v3}
                        _ = v12
                    case "foo bar":
                        if ((1e10 < v7) && false) {
                            v13 := []int{}
                            v14 := v1[2147483647]
                            _ = v13
                            _ = v14
                        } else {
                            switch (fi11 * (fi10 + 1)) {
                            case 4294967295:
                                v15 := GenS1{F0: -112, F1: 0}
                                _ = v15
                            default:
                                if (!false) {
                                    switch "a" {
                                    case "":
                                        fmt.Println(v7, fi10, v7)
                                    case "foo bar":
                                        v3 = (v3 | (v3 / ((v3) | 1)))
                                    case "hello":
                                        if (42 >= -128) {
                                            switch "" {
                                            case "foo bar":
                                                var v16 uint64 = (v3 | 858)
                                                _ = v16
                                            }
                                            v4 += v4
                                        }
                                    default:
                                        v17 := GenS1{F0: -128, F1: v3}
                                        _ = v17
                                    }
                                    var v18 int16 = -32768
                                    if (false && (21 <= 1)) {
                                        fmt.Println(v3, fi10)
                                        v4 -= 1.5
                                        fmt.Println(v7, v3)
                                    } else {
                                        switch ("a" + "hello") {
                                        case "\n\t":
                                            var v19 int = fi10
                                            _ = v19
                                        default:
                                            var v20 int32 = 811
                                            _ = v20
                                        }
                                        v3 -= v3
                                        if false {
                                            v21 := GenS2{F0: true, F1: -1, F2: 2147483647}
                                            if (1 == -2147483648) {
                                                fmt.Println(v3, v4, v3)
                                                v22 := GenS1{F0: -71, F1: v3}
                                                v4 -= v4
                                                _ = v22
                                            } else {
                                                v23 := GenS3{F0: fi11, F1: ((42 > 998) || true), F2: (0 != 0)}
                                                v24 := GenS1{F0: -1, F1: (^(v3))}
                                                v18 = v18
                                                _ = v23
                                                _ = v24
                                            }
                                            switch -2147483648 {
                                            case -2147483648:
                                                v25 := GenS1{F0: 64, F1: v3}
                                                _ = v25
                                            case 42:
                                                v26 := GenS2{F0: (true || (((-(fi11)) != fi10) || (false && (true || (42 >= v3))))), F1: 0, F2: 0}
                                                v27 := GenS1{F0: 127, F1: v3}
                                                _ = v26
                                                _ = v27
                                            }
                                            v28 := []int{}
                                            switch "\n\t" {
                                            case "foo bar":
                                                v29 := GenS4{F0: 65535}
                                                _ = v29
                                            case "a":
                                                var v30 bool = (42 == 0)
                                                _ = v30
                                            case "\n\t":
                                                fmt.Println(fi11, v7)
                                            default:
                                                v18 = v18
                                            }
                                            _ = v21
                                            _ = v28
                                        } else {
                                            var v31 uint16 = 65535
                                            var v32 int8 = -128
                                            _ = v31
                                            _ = v32
                                        }
                                        v33 := GenS2{F0: (true || true), F1: -1, F2: 1}
                                        _ = v33
                                    }
                                    _ = v18
                                }
                            }
                            v34 := GenS5{F0: ("foo bar" + "a")}
                            _ = v34
                        }
                        switch ("" + "a") {
                        case "\n\t":
                            var v35 uint32 = 1
                            _ = v35
                        case "foo bar":
                            var v36 uint32 = 2147483647
                            _ = v36
                        case "a":
                            fmt.Println(v7, fi11, v4)
                        default:
                            v3 = 42
                        }
                    default:
                        var v37 bool = (65535 != 42)
                        _ = v37
                    }
                    v38 := GenS4{F0: 0}
                    _ = fi11
                    _ = v38
                }
                _ = fi10
            }
        case "hello":
            if true {
                for fi39 := 0; fi39 < 6; fi39++ {
                    var v40 int = (fi39 & fi39)
                    _ = fi39
                    _ = v40
                }
                for fi41 := 0; fi41 < 1; fi41++ {
                    switch "hello" {
                    case "hello":
                        _ = len(v1)
                    }
                    var v42 string = "a"
                    var v43 int8 = 127
                    _ = fi41
                    _ = v42
                    _ = v43
                }
            }
        }
    } else {
        v7 -= ((0.0 * v7) / (1.5 * (1.5 / 1e10)))
        var v44 float64 = float64((1.5 * 0.0))
        switch (("foo bar" + "hello") + "") {
        case "a":
            for fi45 := 0; fi45 < 2; fi45++ {
                v44 -= (((v4 - ((float64(v7) * 0.0) / v44)) + v4) * ((1e10 / v44) + v4))
                var v46 uint8 = 20
                _ = fi45
                _ = v46
            }
        }
        _ = v44
    }
    var v47 bool = (!false)
    v48 := []int{}
    if v47 {
        var v49 int32 = -1
        _ = v49
    } else {
        var v50 string = "a"
        var v51 string = (v50 + (v50 + v50))
        _ = v50
        _ = v51
    }
    v48 = append(v48, 782)
    if len(v48) > 0 {
        v48[0] = -969
    }
    fmt.Println(v7)
    if true {
        var v52 int32 = 963
        _ = v52
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v6
    _ = v7
}
