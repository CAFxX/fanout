package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int8; F1 int8; F2 int }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint, p1 uint32) int16 {
    p0 = (p0 + 65535)
    if true {
        v7 := GenS1{F0: uint16(p0), F1: -1}
        for fi8 := 0; fi8 < 9; fi8++ {
            if false {
                var v9 uint64 = 683
                if ("\n\t" != "a") {
                    for fi10 := 0; fi10 < 18; fi10++ {
                        v11 := GenS1{F0: 2, F1: fi8}
                        switch (-(fi10)) {
                        case -457:
                            v12 := GenS1{F0: 246, F1: (fi8 % (((fi10 | -2147483648)) | 1))}
                            _ = v12
                        }
                        _ = fi10
                        _ = v11
                    }
                    switch ("foo bar" + ("hello" + "hello")) {
                    case "a":
                        v13 := GenS2{F0: 76, F1: -1, F2: int(p0)}
                        _ = v13
                    case "\n\t":
                        v14 := GenS1{F0: 65535, F1: fi8}
                        _ = v14
                    case "":
                        var v15 uint = 42
                        _ = v15
                    default:
                        v16 := []int{}
                        _ = v16
                    }
                } else {
                    var v17 uint64 = v9
                    v18 := GenS1{F0: 255, F1: 32767}
                    _ = v17
                    _ = v18
                }
                _ = v9
            }
            fmt.Println(p0)
            var v19 uint64 = 0
            _ = fi8
            _ = v19
        }
        _ = v7
    } else {
        var v20 uint8 = 255
        _ = v20
    }
    var v21 uint = p0
    switch "a" {
    case "hello":
        v22 := []int{-397, 32768, 0, 4294967295}
        _ = v22
    default:
        var v23 int64 = -769
        _ = v23
    }
    v24 := []int{9223372036854775807, 42, -100}
    _ = p0
    _ = p1
    _ = v21
    _ = v24
    return 0
}

func main() {
    v1 := []string{"\n\t", ""}
    var v2 uint16 = 0
    var v3 float64 = 3.14159
    var v4 int32 = 2147483647
    var v5 float64 = v3
    v6 := make(map[uint32]int)
    var v25 uint32 = 81
    switch -384 {
    case 9223372036854775807:
        for fi26 := 0; fi26 < 20; fi26++ {
            switch ("" + ("\n\t" + "")) {
            case "a":
                var v27 bool = (false && true)
                _ = v27
            case "":
                if false {
                    if true {
                        v1 = append(v1, "foo bar")
                        var v28 int = 747
                        fmt.Println(v2, v2, v4)
                        _ = v28
                    } else {
                        for fi29 := 0; fi29 < 19; fi29++ {
                            var v30 int = (fi29 - (fi26 ^ (fi29 % ((fi26) | 1))))
                            if true {
                                var v31 float32 = 0.0
                                _ = v31
                            }
                            v32 := v6[(^(v25))]
                            _ = fi29
                            _ = v30
                            _ = v32
                        }
                        var v33 int32 = -1
                        _ = len(v6)
                        _ = v33
                    }
                    if len(v1) > 0 {
                        v1[0] = "\n\t"
                    }
                }
            default:
                fmt.Println(v3, v4, v3)
            }
            v3 -= 1.5
            if false {
                fmt.Println(fi26, v25)
                if ((3.14159 < v3) && false) {
                    var v34 uint8 = 2
                    if len(v1) > 0 {
                        v1[0] = "foo bar"
                    }
                    _ = v34
                }
                var v35 int64 = -9223372036854775808
                _ = v35
            } else {
                var v36 int8 = 127
                if (fi26 == 42) {
                    var v37 uint = 0
                    for fi38 := 0; fi38 < 13; fi38++ {
                        v36 = (v36 % (((-(v36))) | 1))
                        v4 = v4
                        _ = fi38
                    }
                    _ = v37
                }
                _ = v36
            }
            _ = fi26
        }
    case 4294967295:
        v25 -= v25
    case 42:
        if true {
            fmt.Println(v2)
        }
    default:
        v39 := GenS2{F0: 42, F1: -128, F2: 2147483647}
        _ = v39
    }
    _ = len(v1)
    var v40 uint64 = 42
    switch (((("hello" + ("" + "\n\t")) + "") + "") + "a") {
    case "hello":
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
    case "foo bar":
        var v41 uint8 = 169
        _ = v41
    case "a":
        fmt.Println(v25, v40, v4)
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v42 := GenS1{F0: (^(v2)), F1: -814}
    v3 = v5
    var v43 string = ("\n\t" + "")
    if (1.5 < float64(float32(0.0))) {
        _ = len(v1)
        fmt.Println(v40, v5, v25)
        if (((^(v2)) <= 951) && false) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v44 string = "\n\t"
            _ = v44
        } else {
            for fi45 := 0; fi45 < 7; fi45++ {
                var v46 string = v43
                delete(v6, (v25 << ((2) & 63)))
                _ = fi45
                _ = v46
            }
            switch 127 {
            case 32768:
                v47 := GenS1{F0: v2, F1: 2147483647}
                _ = v47
            case -265:
                if true {
                    v48 := GenS3{F0: (v25 + v25), F1: 1}
                    switch v43 {
                    case "\n\t":
                        var v49 bool = true
                        _ = v49
                    case "a":
                        for fi50 := 0; fi50 < 4; fi50++ {
                            v51 := GenS4{F0: (v4 % ((0) | 1))}
                            for fi52 := 0; fi52 < 9; fi52++ {
                                if (!true) {
                                    v40 = (v40 / (((v40 % (((v40 & uint64(v2))) | 1))) | 1))
                                    v53 := v6[2147483648]
                                    _ = v53
                                }
                                fmt.Println(v40, v2, fi50)
                                if (false && true) {
                                    var v54 uint32 = v25
                                    _ = v54
                                } else {
                                    var v55 int16 = 1000
                                    var v56 bool = (513 >= 2)
                                    var v57 int = 4294967295
                                    _ = v55
                                    _ = v56
                                    _ = v57
                                }
                                switch ((v43 + v43) + v43) {
                                case "hello":
                                    v43 = "a"
                                case "":
                                    fmt.Println(fi50, v43)
                                default:
                                    var v58 uint32 = v25
                                    _ = v58
                                }
                                var v59 float32 = 1.5
                                _ = fi52
                                _ = v59
                            }
                            _ = fi50
                            _ = v51
                        }
                    }
                    v40 = v40
                    _ = v48
                } else {
                    if (128 > 219) {
                        var v60 uint16 = (^(v2))
                        var v61 string = "a"
                        v62 := GenS4{F0: v4}
                        _ = v60
                        _ = v61
                        _ = v62
                    } else {
                        var v63 uint16 = (v2 & 494)
                        fmt.Println(v63, v40)
                        _ = v63
                    }
                    switch ((v43 + (("foo bar" + v43) + v43)) + ((v43 + v43) + ("\n\t" + (("\n\t" + v43) + "")))) {
                    case "hello":
                        v43 = (v43 + v43)
                    case "\n\t":
                        var v64 float64 = 3.14159
                        _ = v64
                    }
                    var v65 string = v43
                    _ = v65
                }
            default:
                for fi66 := 0; fi66 < 2; fi66++ {
                    switch v43 {
                    case "foo bar":
                        for fi67 := 0; fi67 < 4; fi67++ {
                            fmt.Println(v2, v4)
                            var v68 uint8 = 255
                            var v69 string = (v43 + "a")
                            if (-168 <= -639) {
                                v25 -= (v25 | uint32(v68))
                            }
                            _ = fi67
                            _ = v68
                            _ = v69
                        }
                    case "a":
                        v1 = append(v1, "a")
                    }
                    if len(v1) > 0 {
                        v1[0] = (v43 + v43)
                    }
                    fmt.Println(v4, v4, v25)
                    _ = fi66
                }
            }
            v1 = append(v1, v43)
        }
    }
    var v70 int64 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v40
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v70
}
