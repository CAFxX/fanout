package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint; F1 int16; F2 int64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint; F1 string }

func main() {
    v1 := make(map[int]string)
    var v2 bool = (30 == 1)
    v3 := make(map[string]bool)
    var v4 int32 = 2147483647
    var v5 uint32 = 0
    var v6 float32 = float32(3.14159)
    v3[""] = (!v2)
    fmt.Println(v4, v2)
    v7 := GenS1{F0: 873, F1: 741}
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 string = "\n\t"
        _ = fi8
        _ = v9
    }
    switch "" {
    case "a":
        v4 -= (v4 | (v4 | v4))
    default:
        v4 = (v4 | (v4 ^ -1))
    }
    if ((v4 % (((v4 % ((v4) | 1))) | 1)) > (v4 & (-(v4)))) {
        fmt.Println(v6)
        for fi10 := 0; fi10 < 9; fi10++ {
            var v11 int = int(v4)
            var v12 int16 = -747
            for fi13 := 0; fi13 < 5; fi13++ {
                switch (v11 / ((32767) | 1)) {
                case 32767:
                    if ((v4 >> ((960) & 63)) == 0) {
                        if true {
                            var v14 int = 2147483647
                            var v15 bool = v2
                            var v16 string = (("\n\t" + "hello") + "")
                            _ = v14
                            _ = v15
                            _ = v16
                        } else {
                            v17 := GenS2{F0: 4294967295, F1: int16(v5)}
                            _ = v17
                        }
                        var v18 bool = v2
                        switch ("hello" + ("hello" + "\n\t")) {
                        case "hello":
                            if v2 {
                                if v2 {
                                    if (42 == 9223372036854775807) {
                                        v12 = 1
                                    } else {
                                        v19 := GenS3{F0: 65535, F1: (v12 | (v12 << ((4294967295) & 63))), F2: -9223372036854775808}
                                        _ = v19
                                    }
                                } else {
                                    v20 := GenS2{F0: 1, F1: v12}
                                    _ = v20
                                }
                                if v18 {
                                    switch (^(fi13)) {
                                    case 302:
                                        var v21 int32 = v4
                                        _ = v21
                                    case -2147483648:
                                        var v22 int64 = 1
                                        _ = v22
                                    case 4294967295:
                                        var v23 int = fi13
                                        _ = v23
                                    default:
                                        _ = len(v3)
                                    }
                                    fmt.Println(fi13, fi10, v4)
                                    fmt.Println(fi13)
                                } else {
                                    switch (("foo bar" + "foo bar") + "") {
                                    case "hello":
                                        v24 := []int{-2147483648, fi10, 4294967295}
                                        _ = v24
                                    }
                                    v25 := []int{}
                                    _ = v25
                                }
                                fmt.Println(v5, v11, v12)
                            } else {
                                if v2 {
                                    v26 := []int{}
                                    _ = v26
                                }
                                v5 = (v5 - 194)
                            }
                        }
                        v4 -= v4
                        switch ("a" + "\n\t") {
                        case "hello":
                            v4 += v4
                        case "":
                            fmt.Println(fi10, v2)
                        case "foo bar":
                            v27 := GenS4{F0: ("" + ("a" + "a"))}
                            _ = v27
                        }
                        _ = v18
                    }
                    if false {
                        switch (fi13 - fi13) {
                        case 32767:
                            fmt.Println(v4, v4)
                        case -9223372036854775808:
                            v4 = 2147483647
                        default:
                            var v28 int8 = 0
                            _ = v28
                        }
                        if v2 {
                            var v29 uint = 140
                            v30 := GenS3{F0: (v29 ^ v29), F1: 1000, F2: 9223372036854775807}
                            v29 = (v29 >> (((v29 - v29)) & 63))
                            _ = v29
                            _ = v30
                        }
                    } else {
                        switch ("" + "\n\t") {
                        case "foo bar":
                            v31 := GenS2{F0: 42, F1: v12}
                            var v32 float32 = v6
                            _ = v31
                            _ = v32
                        case "a":
                            var v33 float64 = ((((1.5 / -2.25) * (0.1 + 1.5)) + 0.1) * (float64(float32(3.14159)) + -2.25))
                            _ = v33
                        }
                        v12 = v12
                        v34 := []int{(^(v11))}
                        _ = v34
                    }
                    v35 := v1[(fi10 << ((255) & 63))]
                    fmt.Println(v4, v6)
                    delete(v3, (v35 + v35))
                    switch (fi13 >> ((345) & 63)) {
                    case 2147483647:
                        _ = len(v3)
                    case 9223372036854775807:
                        v2 = v2
                    default:
                        if false {
                            v4 = (-(v4))
                            v36 := []int{(fi13 & -883), fi13}
                            _ = v36
                        }
                    }
                    _ = v35
                case 128:
                    if (!(1 <= 2)) {
                        if v2 {
                            fmt.Println(v6, v2)
                        } else {
                            if false {
                                fmt.Println(fi13, fi13, v2)
                                var v37 string = "foo bar"
                                v11 -= 820
                                _ = v37
                            }
                        }
                    } else {
                        v38 := GenS2{F0: 0, F1: v12}
                        fmt.Println(v2, v12)
                        fmt.Println(v5, fi13)
                        _ = v38
                    }
                }
                var v39 int64 = 9223372036854775807
                _ = fi13
                _ = v39
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        v6 = (0.1 * v6)
    } else {
        v6 += ((1.5 + 1.5) * ((-2.25 + v6) - 1.5))
    }
    v40 := []int{0, 32768}
    v41 := GenS2{F0: 0, F1: -44}
    for fi42 := 0; fi42 < 20; fi42++ {
        var v43 int16 = int16(v5)
        var v44 uint32 = v5
        _ = fi42
        _ = v43
        _ = v44
    }
    switch ("a" + "foo bar") {
    case "":
        if v2 {
            fmt.Println(v4, v2)
        }
    case "\n\t":
        if v2 {
            var v45 float64 = ((0.0 - 1e10) * (1.5 / 1e10))
            _ = v45
        }
    case "hello":
        var v46 bool = (true || v2)
        _ = v46
    }
    v4 -= (v4 << ((171) & 63))
    v6 = float32(1.5)
    var v47 int64 = -9223372036854775808
    switch "foo bar" {
    case "foo bar":
        for fi48 := 0; fi48 < 3; fi48++ {
            if (v2 || (true || (1 == 4294967295))) {
                switch 955 {
                case -2147483648:
                    var v49 float64 = (1.5 - 1e10)
                    _ = v49
                case -9223372036854775808:
                    for fi50 := 0; fi50 < 2; fi50++ {
                        if len(v40) > 0 {
                            v40[0] = 1
                        }
                        fmt.Println(v4)
                        var v51 int64 = v47
                        _ = fi50
                        _ = v51
                    }
                case -1:
                    v52 := GenS5{F0: 255, F1: "foo bar"}
                    _ = v52
                default:
                    for fi53 := 0; fi53 < 8; fi53++ {
                        switch "a" {
                        case "hello":
                            var v54 int32 = v4
                            _ = v54
                        case "a":
                            v2 = v2
                        case "foo bar":
                            v2 = v2
                        default:
                            v55 := v1[fi53]
                            _ = v55
                        }
                        _ = fi53
                    }
                }
                v2 = false
            }
            if v2 {
                delete(v3, ("foo bar" + "hello"))
            }
            v47 -= 1
            _ = fi48
        }
    case "":
        var v56 uint16 = 1
        _ = v56
    default:
        var v57 int8 = -128
        _ = v57
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v47
    _ = v5
    _ = v6
    _ = v7
}
