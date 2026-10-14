package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8; F1 uint8; F2 bool }
type GenS6 struct { F0 uint32; F1 int8; F2 int64 }
func genFn0(p0 bool, p1 int, p2 int64) string {
    if p0 {
        fmt.Println(p1, p0, p2)
        var v6 int8 = 100
        _ = v6
    } else {
        v7 := GenS1{F0: "hello", F1: p0}
        v8 := GenS2{F0: 44, F1: ""}
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 17; fi9++ {
        v10 := []int{fi9, fi9, p1, (^(p1))}
        switch "hello" {
        case "a":
            var v11 int32 = -673
            _ = v11
        case "foo bar":
            if p0 {
                fmt.Println(fi9, fi9)
                p1 = p1
                for fi12 := 0; fi12 < 9; fi12++ {
                    _ = len(v10)
                    v13 := GenS3{F0: 42}
                    _ = fi12
                    _ = v13
                }
            } else {
                var v14 uint = 905
                p2 = (p2 + (-(p2)))
                var v15 int8 = 119
                _ = v14
                _ = v15
            }
        case "\n\t":
            var v16 bool = (-2147483648 != 42)
            _ = v16
        default:
            p1 = 849
        }
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    return "a"
}

func main() {
    v1 := []string{(("hello" + "hello") + "foo bar")}
    v2 := make(map[string]int)
    v3 := make(map[uint32]int)
    var v4 int16 = 32767
    v5 := make(map[uint32]int)
    var v17 int32 = 1
    v4 += (-(v4))
    v17 = v17
    v18 := GenS3{F0: 42}
    var v19 float64 = ((-2.25 / -2.25) / -2.25)
    fmt.Println(v19, v4)
    if (!true) {
        _ = len(v5)
        for fi20 := 0; fi20 < 14; fi20++ {
            var v21 uint = 0
            v22 := GenS4{F0: 0}
            v19 = (0.1 / float64(0.1))
            _ = fi20
            _ = v21
            _ = v22
        }
        var v23 float32 = ((0.0 / 1.0) + -2.25)
        _ = v23
    } else {
        fmt.Println(v19, v17, v19)
        fmt.Println(v19, v4)
        fmt.Println(v17, v4, v4)
    }
    var v24 uint64 = 18446744073709551615
    v17 += int32(v4)
    var v25 int32 = v17
    var v26 bool = false
    for fi27 := 0; fi27 < 16; fi27++ {
        _ = len(v1)
        v1 = append(v1, "foo bar")
        _ = fi27
    }
    v26 = (((-9223372036854775808 < -588) && (!v26)) && false)
    v28 := v5[136]
    if v26 {
        var v29 uint8 = 153
        var v30 uint64 = 0
        _ = v29
        _ = v30
    }
    v19 -= 0.0
    for fi31 := 0; fi31 < 4; fi31++ {
        switch "\n\t" {
        case "hello":
            switch fi31 {
            case -2147483648:
                var v32 int32 = v17
                _ = v32
            case 4294967295:
                v3[2147483647] = (fi31 | (fi31 % (((v28 >> ((2) & 63))) | 1)))
            default:
                if v26 {
                    switch (fi31 - (v28 << ((1) & 63))) {
                    case -1:
                        var v33 uint = 799
                        _ = v33
                    default:
                        if ((^(v28)) <= (-(fi31))) {
                            if v26 {
                                var v34 int = 127
                                _ = v34
                            }
                        }
                    }
                    if (v26 || v26) {
                        switch v28 {
                        case -235:
                            switch (fi31 ^ v28) {
                            case -1:
                                for fi35 := 0; fi35 < 9; fi35++ {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    _ = fi35
                                }
                            case 2147483647:
                                v36 := GenS5{F0: 42, F1: 1, F2: v26}
                                _ = v36
                            case 42:
                                fmt.Println(fi31, v26)
                            }
                        case -540:
                            var v37 uint8 = 1
                            _ = v37
                        case 9223372036854775807:
                            v28 += fi31
                        }
                        var v38 int32 = (v17 & v25)
                        _ = v38
                    }
                }
            }
        case "\n\t":
            var v39 uint32 = 1
            _ = v39
        default:
            v40 := GenS1{F0: "a", F1: v26}
            _ = v40
        }
        if (v26 && (v17 > (v17 >> ((179) & 63)))) {
            v41 := GenS1{F0: "hello", F1: (!v26)}
            fmt.Println(v25, v26, v19)
            _ = v41
        } else {
            v4 = (^(v4))
        }
        for fi42 := 0; fi42 < 4; fi42++ {
            switch fi31 {
            case 32767:
                switch fi31 {
                case 483:
                    switch (^(fi31)) {
                    case -1:
                        switch "" {
                        case "a":
                            if ("a" != (("\n\t" + "") + ((("hello" + (("\n\t" + "hello") + "a")) + ("" + "foo bar")) + (("hello" + "a") + "")))) {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                if len(v1) > 0 {
                                    v1[0] = ("foo bar" + ("foo bar" + "foo bar"))
                                }
                                var v43 uint16 = 804
                                _ = v43
                            }
                        case "foo bar":
                            v28 = fi31
                        }
                    case 32768:
                        var v44 int64 = -9223372036854775808
                        _ = v44
                    case -2147483648:
                        v4 = v4
                    }
                default:
                    switch ("hello" + "hello") {
                    case "foo bar":
                        v25 = -1
                    case "":
                        fmt.Println(v28, v17)
                    default:
                        v25 += v17
                    }
                    v25 = int32(fi42)
                }
            case 128:
                if false {
                    var v45 int = fi42
                    if (!(0 >= 361)) {
                        if v26 {
                            var v46 int = fi31
                            _ = len(v5)
                            _ = v46
                        } else {
                            switch "" {
                            case "":
                                v47 := v2["hello"]
                                _ = v47
                            }
                            fmt.Println(v25, v26, v25)
                        }
                        if (float32(1e10) <= 1.5) {
                            if v26 {
                                switch 4294967295 {
                                case 427:
                                    if ((v25 << ((4294967295) & 63)) != v17) {
                                        switch (((("foo bar" + "a") + (("foo bar" + "a") + "\n\t")) + ("hello" + "foo bar")) + "\n\t") {
                                        case "":
                                            v19 += -2.25
                                        default:
                                            var v48 int = (v45 * (fi42 * (v45 & -663)))
                                            _ = v48
                                        }
                                        if v26 {
                                            var v49 int16 = (v4 | 42)
                                            var v50 int64 = 304
                                            switch "\n\t" {
                                            case "foo bar":
                                                v51 := GenS6{F0: 2147483648, F1: 0, F2: v50}
                                                _ = v51
                                            case "a":
                                                var v52 float64 = v19
                                                _ = v52
                                            case "\n\t":
                                                if (fi31 >= -413) {
                                                    var v53 uint32 = 42
                                                    _ = v53
                                                }
                                            }
                                            _ = v49
                                            _ = v50
                                        }
                                        v24 = (v24 >> ((1) & 63))
                                    }
                                case 32768:
                                    if v26 {
                                        var v54 int = v28
                                        _ = v54
                                    } else {
                                        v24 = (v24 / ((0) | 1))
                                    }
                                }
                                var v55 string = ""
                                _ = v55
                            } else {
                                var v56 bool = v26
                                var v57 float32 = 0.0
                                _ = v56
                                _ = v57
                            }
                            if (v26 || true) {
                                v4 = (v4 << ((866) & 63))
                            }
                        }
                        v58 := GenS2{F0: 127, F1: "\n\t"}
                        _ = v58
                    } else {
                        var v59 int64 = -876
                        var v60 bool = v26
                        _ = v59
                        _ = v60
                    }
                    v26 = (("hello" + "hello") >= "\n\t")
                    _ = v45
                } else {
                    var v61 string = "foo bar"
                    _ = v61
                }
            case 9223372036854775807:
                v62 := GenS4{F0: 1}
                _ = v62
            }
            if (v26 && v26) {
                v26 = ((v25 / ((42) | 1)) < -2147483648)
            }
            v4 += -32768
            _ = fi42
        }
        _ = fi31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
