package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 int32; F2 string }
type GenS4 struct { F0 int64; F1 uint16; F2 uint }
type GenS5 struct { F0 bool; F1 bool; F2 uint }
func genFn0() int8 {
    for fi4 := 0; fi4 < 17; fi4++ {
        var v5 int = (fi4 | 32767)
        _ = fi4
        _ = v5
    }
    var v6 int16 = 0
    v7 := GenS1{F0: ("foo bar" + "a"), F1: false, F2: 335}
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 string = "hello"
        var v10 string = v9
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = v6
    _ = v7
    return -105
}

func main() {
    var v1 uint8 = 2
    v2 := make(map[int]bool)
    var v3 int64 = 42
    v11 := GenS2{F0: 424}
    var v12 int32 = int32(v3)
    for fi13 := 0; fi13 < 11; fi13++ {
        switch (fi13 >> ((0) & 63)) {
        case 128:
            switch int(v12) {
            case -9223372036854775808:
                if false {
                    for fi14 := 0; fi14 < 5; fi14++ {
                        v1 -= 255
                        var v15 string = "foo bar"
                        _ = fi14
                        _ = v15
                    }
                    var v16 string = "foo bar"
                    if (797 > 42) {
                        v12 = -396
                    } else {
                        fmt.Println(v12, fi13, v3)
                        var v17 uint = 18446744073709551615
                        switch ("" + (v16 + (v16 + v16))) {
                        case "\n\t":
                            var v18 string = (v16 + "a")
                            _ = v18
                        case "":
                            fmt.Println(v12, v17)
                        case "hello":
                            var v19 uint16 = 255
                            _ = v19
                        default:
                            if false {
                                v1 -= 255
                                var v20 string = ((v16 + "a") + "\n\t")
                                _ = v20
                            }
                        }
                        _ = v17
                    }
                    _ = v16
                } else {
                    for fi21 := 0; fi21 < 10; fi21++ {
                        var v22 int64 = v3
                        switch "a" {
                        case "hello":
                            fmt.Println(v3, v12)
                        case "\n\t":
                            switch "" {
                            case "\n\t":
                                if (0.0 != -2.25) {
                                    var v23 int32 = v12
                                    _ = v23
                                }
                            case "":
                                var v24 int = (fi21 ^ 32767)
                                _ = v24
                            case "a":
                                if (true || false) {
                                    v25 := []int{}
                                    v26 := GenS1{F0: "", F1: ("\n\t" == ""), F2: 826}
                                    var v27 int16 = 1
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                } else {
                                    v12 -= (v12 - (-(v12)))
                                }
                            default:
                                v28 := GenS3{F0: "a", F1: 1, F2: ("foo bar" + "hello")}
                                _ = v28
                            }
                        }
                        _ = fi21
                        _ = v22
                    }
                }
            case 1:
                if (1 > 65535) {
                    v3 += -1
                    fmt.Println(v12)
                    switch int(v3) {
                    case -2147483648:
                        v12 += 42
                    case 710:
                        switch "foo bar" {
                        case "a":
                            var v29 float64 = float64(float32(0.0))
                            _ = v29
                        case "\n\t":
                            for fi30 := 0; fi30 < 3; fi30++ {
                                var v31 int16 = 42
                                if false {
                                    var v32 int64 = -1
                                    switch -1 {
                                    case -711:
                                        v1 = 2
                                    }
                                    var v33 float64 = 0.1
                                    _ = v32
                                    _ = v33
                                }
                                _ = fi30
                                _ = v31
                            }
                        case "foo bar":
                            if (176 < 0) {
                                var v34 float32 = 1.5
                                _ = v34
                            } else {
                                v1 = 1
                            }
                        }
                    }
                }
            }
        default:
            var v35 uint16 = 91
            _ = v35
        }
        _ = fi13
    }
    v36 := GenS4{F0: (v3 * -644), F1: 2, F2: 42}
    var v37 float64 = 0.0
    fmt.Println(v37, v1)
    var v38 int64 = (v3 - (v3 + (-(v3))))
    for fi39 := 0; fi39 < 1; fi39++ {
        v1 = (v1 & v1)
        _ = fi39
    }
    v40 := []int{32767, 1}
    var v41 bool = true
    v40 = append(v40, 222)
    v42 := GenS1{F0: "\n\t", F1: (32767 >= 918), F2: 132}
    fmt.Println(v41)
    var v43 uint32 = 296
    switch ("hello" + "") {
    case "a":
        switch "foo bar" {
        case "foo bar":
            v37 -= v37
        case "\n\t":
            if (1.5 != 0.1) {
                v44 := GenS1{F0: "a", F1: false, F2: 275}
                for fi45 := 0; fi45 < 18; fi45++ {
                    fmt.Println(v37)
                    _ = fi45
                }
                var v46 int64 = v38
                _ = v44
                _ = v46
            } else {
                var v47 uint32 = v43
                for fi48 := 0; fi48 < 9; fi48++ {
                    v38 -= 914
                    if (uint32(v1) != v43) {
                        v3 = (v38 ^ v3)
                    }
                    _ = fi48
                }
                _ = v47
            }
        case "":
            v1 += 2
        }
    case "":
        var v49 float32 = (-2.25 / (-2.25 - 3.14159))
        _ = v49
    case "foo bar":
        v1 += (v1 >> ((4294967295) & 63))
    }
    if ((425 >= (v43 >> ((42) & 63))) && (1 < 2)) {
        v50 := GenS3{F0: "\n\t", F1: (v12 - (v12 % (((v12 % ((v12) | 1))) | 1))), F2: ("hello" + ((("" + "a") + "foo bar") + "\n\t"))}
        v51 := GenS5{F0: v41, F1: (v41 || v41), F2: 1}
        if ((1e10 >= (1.5 + (v37 * 3.14159))) || false) {
            switch ("a" + "") {
            case "\n\t":
                v41 = (v12 != (v12 | (v12 * (v12 | v12))))
            default:
                fmt.Println(v1, v38)
            }
            var v52 bool = ((v37 - (v37 + 3.14159)) <= 0.0)
            _ = v52
        } else {
            for fi53 := 0; fi53 < 3; fi53++ {
                v54 := GenS5{F0: false, F1: (v41 && v41), F2: 2}
                _ = fi53
                _ = v54
            }
        }
        _ = v50
        _ = v51
    } else {
        if true {
            v38 = v3
            var v55 int64 = -790
            fmt.Println(v37, v3, v37)
            _ = v55
        } else {
            fmt.Println(v38)
            v3 -= -343
            var v56 string = "\n\t"
            _ = v56
        }
    }
    var v57 uint32 = 1
    if v41 {
        if false {
            for fi58 := 0; fi58 < 16; fi58++ {
                for fi59 := 0; fi59 < 13; fi59++ {
                    switch fi58 {
                    case 2147483647:
                        var v60 uint64 = 1
                        _ = v60
                    }
                    _ = fi59
                }
                _ = fi58
            }
        }
        v57 = v57
        v12 += 2147483647
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v57
}
