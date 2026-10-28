package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 uint; F2 string }
type GenS5 struct { F0 string; F1 int8 }
type GenS6 struct { F0 int16; F1 int; F2 string }
type GenS7 struct { F0 int8 }

func main() {
    var v1 bool = (667 < 127)
    var v2 int16 = 0
    var v3 uint32 = 2147483647
    var v4 float32 = -2.25
    v5 := make(map[uint32]string)
    var v6 int8 = 0
    _ = len(v5)
    fmt.Println(v1, v4)
    v3 += 1
    var v7 int16 = (v2 / (((v2 + (-(v2)))) | 1))
    var v8 int = -9223372036854775808
    v7 = 1
    for fi9 := 0; fi9 < 6; fi9++ {
        if (false || v1) {
            delete(v5, v3)
        }
        var v10 bool = false
        switch ("hello" + "\n\t") {
        case "a":
            v8 -= v8
        case "hello":
            if v10 {
                for fi11 := 0; fi11 < 17; fi11++ {
                    var v12 int64 = -182
                    if (false || v1) {
                        if true {
                            var v13 uint8 = 255
                            var v14 int64 = (^(v12))
                            if (v1 && v10) {
                                if false {
                                    v3 = (^(v3))
                                    var v15 float64 = (float64(3.14159) - (1.5 * (float64((v4 * float32(0.0))) * -2.25)))
                                    fmt.Println(v2, v2)
                                    var v16 uint64 = 1
                                    var v17 uint64 = 1
                                    fmt.Println(v2, v3)
                                    var v18 uint = 136
                                    var v19 uint8 = v13
                                    if false {
                                        var v20 int8 = (v6 * v6)
                                        var v21 bool = (true && (738 > 1))
                                        _ = v20
                                        _ = v21
                                    } else {
                                        if (!v10) {
                                            switch "hello" {
                                            case "a":
                                                v22 := GenS1{F0: (fi11 | (fi9 | fi9)), F1: v2, F2: (v18 + v18)}
                                                _ = v22
                                            case "\n\t":
                                                fmt.Println(v18)
                                            case "":
                                                if v1 {
                                                    var v23 int64 = v14
                                                    _ = v23
                                                }
                                            default:
                                                if v1 {
                                                    fmt.Println(v15, v12, v2)
                                                    if ((!(float64((v4 / v4)) < v15)) || ((^(v3)) > (v3 + (v3 / (((v3 & v3)) | 1))))) {
                                                        v3 -= v3
                                                    } else {
                                                        v14 = v14
                                                        v7 -= 0
                                                        var v24 float64 = (v15 * 0.1)
                                                        _ = v24
                                                    }
                                                    var v25 int64 = (v14 ^ (-(v12)))
                                                    _ = v25
                                                }
                                            }
                                            v26 := v5[(v3 % ((v3) | 1))]
                                            _ = v26
                                        } else {
                                            if ("" > "hello") {
                                                v27 := GenS2{F0: ("foo bar" + ((((("" + ("" + ("" + (("hello" + "foo bar") + ("hello" + "foo bar"))))) + "foo bar") + "hello") + "hello") + "a")), F1: (v7 << ((v18) & 63))}
                                                _ = v27
                                            }
                                            var v28 float32 = 1.5
                                            v29 := v5[(v3 / (((v3 * 1)) | 1))]
                                            _ = v28
                                            _ = v29
                                        }
                                        if (!v10) {
                                            if (!v1) {
                                                v30 := []int{(-(v8))}
                                                _ = v30
                                            } else {
                                                switch "a" {
                                                case "a":
                                                    var v31 uint = v18
                                                    _ = v31
                                                }
                                                var v32 int32 = 928
                                                v10 = true
                                                _ = v32
                                            }
                                        } else {
                                            fmt.Println(v15)
                                            var v33 uint8 = (v13 ^ (v19 >> ((0) & 63)))
                                            _ = v33
                                        }
                                        v8 -= v8
                                    }
                                    var v34 float64 = (((v15 + v15) - v15) / 1e10)
                                    var v35 int32 = 0
                                    var v36 string = "hello"
                                    _ = v15
                                    _ = v16
                                    _ = v17
                                    _ = v18
                                    _ = v19
                                    _ = v34
                                    _ = v35
                                    _ = v36
                                }
                                fmt.Println(v13, v3)
                            }
                            _ = len(v5)
                            _ = v13
                            _ = v14
                        } else {
                            v37 := GenS1{F0: 128, F1: v2, F2: 266}
                            switch ("\n\t" + ("\n\t" + "\n\t")) {
                            case "a":
                                if true {
                                    switch (fi11 << ((42) & 63)) {
                                    case 4294967295:
                                        var v38 uint8 = 126
                                        _ = v38
                                    case -477:
                                        v39 := []int{-1, -2147483648, 9223372036854775807, 128}
                                        fmt.Println(v8)
                                        _ = v39
                                    default:
                                        v40 := GenS3{F0: 222}
                                        _ = v40
                                    }
                                    v3 -= v3
                                    var v41 int8 = v6
                                    _ = v41
                                } else {
                                    v42 := []int{fi9, fi11}
                                    _ = v42
                                }
                            default:
                                v1 = true
                            }
                            if v1 {
                                switch fi9 {
                                case -9223372036854775808:
                                    var v43 float64 = (0.1 / 0.1)
                                    _ = v43
                                case 2147483647:
                                    if (!(!((false || (2 == 255)) && (-2.25 <= float64(((0.1 - (1e10 * v4)) / v4)))))) {
                                        switch (fi9 % (((^(v8))) | 1)) {
                                        case 4294967295:
                                            var v44 string = "foo bar"
                                            _ = v44
                                        case 1:
                                            v45 := GenS4{F0: "", F1: 1, F2: "foo bar"}
                                            _ = v45
                                        case 32767:
                                            fmt.Println(v7)
                                        }
                                    } else {
                                        if (!v10) {
                                            var v46 int16 = v7
                                            if v1 {
                                                v47 := GenS2{F0: ("" + "hello"), F1: v46}
                                                if ((-528 == -353) || (v4 >= v4)) {
                                                    fmt.Println(v6, v46)
                                                } else {
                                                    var v48 int64 = v12
                                                    var v49 int32 = -1
                                                    _ = v48
                                                    _ = v49
                                                }
                                                var v50 uint64 = 18446744073709551615
                                                _ = v47
                                                _ = v50
                                            }
                                            _ = v46
                                        }
                                        var v51 float32 = v4
                                        switch "hello" {
                                        case "a":
                                            v8 += (fi9 % (((fi9 * (v8 / ((9223372036854775807) | 1)))) | 1))
                                        case "":
                                            v52 := GenS5{F0: "a", F1: (v6 + (v6 & 42))}
                                            var v53 string = "hello"
                                            v7 = v2
                                            _ = v52
                                            _ = v53
                                        case "foo bar":
                                            v6 += v6
                                        }
                                        var v54 bool = false
                                        _ = v51
                                        _ = v54
                                    }
                                }
                            }
                            _ = v37
                        }
                        switch (fi9 * fi9) {
                        case -2147483648:
                            var v55 bool = false
                            _ = v55
                        default:
                            v56 := v5[(v3 * (v3 / ((v3) | 1)))]
                            _ = v56
                        }
                        switch v8 {
                        case 9223372036854775807:
                            var v57 bool = (v12 > v12)
                            _ = v57
                        }
                    } else {
                        v3 = (^(v3))
                        v58 := GenS4{F0: ("foo bar" + ("a" + ((("\n\t" + "hello") + "\n\t") + "\n\t"))), F1: 2, F2: "a"}
                        _ = v58
                    }
                    var v59 uint16 = 42
                    v7 = -987
                    _ = fi11
                    _ = v12
                    _ = v59
                }
            } else {
                v6 += (v6 - v6)
                v60 := []int{9223372036854775807, fi9, v8}
                v61 := GenS6{F0: v7, F1: -9223372036854775808, F2: "a"}
                _ = v60
                _ = v61
            }
        }
        _ = fi9
        _ = v10
    }
    var v62 uint8 = 0
    if false {
        fmt.Println(v2)
    }
    v63 := GenS7{F0: -128}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v62
    _ = v63
    _ = v7
    _ = v8
}
