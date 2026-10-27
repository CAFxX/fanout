package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 int16; F1 uint16; F2 int16 }
type GenS6 struct { F0 int8; F1 int64; F2 uint64 }

func main() {
    v1 := make(map[string]string)
    v2 := []int{-2147483648}
    var v3 uint64 = 18446744073709551615
    switch -15 {
    case 128:
        var v4 int32 = -406
        _ = v4
    case -753:
        v3 = v3
    }
    var v5 int8 = 42
    delete(v1, "\n\t")
    v5 = (v5 ^ -34)
    switch ("foo bar" + ("" + "hello")) {
    case "":
        if ((!(!true)) || ((false || (42 < 1)) && (false && false))) {
            if false {
                delete(v1, "\n\t")
                var v6 uint16 = 1
                _ = v6
            } else {
                for fi7 := 0; fi7 < 1; fi7++ {
                    if (fi7 <= (^(fi7))) {
                        for fi8 := 0; fi8 < 12; fi8++ {
                            var v9 string = "\n\t"
                            delete(v1, (v9 + v9))
                            v10 := GenS1{F0: 533, F1: (fi7 | (fi8 & (fi7 << ((1) & 63))))}
                            _ = fi8
                            _ = v10
                            _ = v9
                        }
                    }
                    _ = fi7
                }
                for fi11 := 0; fi11 < 14; fi11++ {
                    if true {
                        var v12 uint = 42
                        if len(v2) > 0 {
                            v2[0] = fi11
                        }
                        _ = v12
                    } else {
                        for fi13 := 0; fi13 < 10; fi13++ {
                            v5 += (v5 / ((v5) | 1))
                            var v14 int = -9223372036854775808
                            fmt.Println(v5)
                            v15 := GenS2{F0: 429, F1: 255}
                            var v16 float64 = ((3.14159 + (0.0 + 0.1)) * 1e10)
                            _ = fi13
                            _ = v14
                            _ = v15
                            _ = v16
                        }
                        v17 := GenS3{F0: 42}
                        _ = v17
                    }
                    switch "\n\t" {
                    case "foo bar":
                        v5 -= v5
                    case "a":
                        var v18 int32 = 746
                        _ = v18
                    }
                    _ = fi11
                }
                v19 := GenS4{F0: 515, F1: 255}
                _ = v19
            }
            fmt.Println(v5, v5)
            var v20 int8 = (v5 >> ((18446744073709551615) & 63))
            _ = v20
        }
    case "a":
        _ = len(v1)
    default:
        if true {
            var v21 int32 = 858
            var v22 int = -2147483648
            v23 := GenS4{F0: (v22 / ((4294967295) | 1)), F1: 962}
            _ = v21
            _ = v22
            _ = v23
        } else {
            if (!false) {
                var v24 bool = true
                v25 := GenS3{F0: 2147483648}
                _ = v24
                _ = v25
            }
        }
    }
    var v26 float32 = 0.1
    if false {
        switch ("" + "foo bar") {
        case "hello":
            if false {
                var v27 int64 = -1
                var v28 int64 = -9223372036854775808
                _ = v27
                _ = v28
            } else {
                if (32767 >= 665) {
                    var v29 string = ("a" + ("\n\t" + ""))
                    var v30 int64 = int64(v3)
                    _ = v29
                    _ = v30
                }
                switch -985 {
                case 127:
                    if (float64(v26) == 1e10) {
                        if false {
                            v26 += 3.14159
                            var v31 int16 = 32767
                            _ = v31
                        } else {
                            for fi32 := 0; fi32 < 20; fi32++ {
                                _ = len(v2)
                                _ = fi32
                            }
                            var v33 string = ((("" + "hello") + "foo bar") + ("a" + "a"))
                            _ = v33
                        }
                        var v34 float64 = 0.0
                        v34 = v34
                        _ = v34
                    } else {
                        v3 = v3
                        for fi35 := 0; fi35 < 19; fi35++ {
                            var v36 int = -1
                            if (!false) {
                                v36 = -2147483648
                                for fi37 := 0; fi37 < 18; fi37++ {
                                    if (42 < -494) {
                                        if len(v2) > 0 {
                                            v2[0] = 42
                                        }
                                        if (!(-1 < 1)) {
                                            if len(v2) > 0 {
                                                v2[0] = (v36 % ((-2147483648) | 1))
                                            }
                                            var v38 bool = (42 > 42)
                                            v2 = append(v2, 4294967295)
                                            var v39 bool = false
                                            _ = v38
                                            _ = v39
                                        }
                                        v26 -= v26
                                        fmt.Println(v5, v5)
                                        switch "foo bar" {
                                        case "a":
                                            v40 := GenS5{F0: 0, F1: 1, F2: -303}
                                            _ = v40
                                        }
                                    }
                                    _ = fi37
                                }
                                v5 += (v5 | (v5 << ((234) & 63)))
                            }
                            _ = fi35
                            _ = v36
                        }
                        v5 += v5
                    }
                case 32767:
                    var v41 bool = (!((-1 > -432) && false))
                    _ = v41
                default:
                    v26 = ((v26 / (1.5 - float32((0.1 + (3.14159 + (-2.25 + 1.5)))))) * 1.5)
                }
                for fi42 := 0; fi42 < 1; fi42++ {
                    var v43 int16 = 386
                    _ = fi42
                    _ = v43
                }
            }
        }
        for fi44 := 0; fi44 < 18; fi44++ {
            v5 += (v5 & (v5 << ((1) & 63)))
            var v45 string = ""
            if false {
                var v46 int = (-(fi44))
                v3 -= 913
                _ = v46
            }
            _ = fi44
            _ = v45
        }
    }
    switch 2147483647 {
    case 807:
        v47 := GenS5{F0: 42, F1: 0, F2: -1}
        _ = v47
    case 128:
        fmt.Println(v5, v3, v5)
    default:
        if (("\n\t" + "") == "a") {
            if (!true) {
                v2 = append(v2, -9223372036854775808)
            }
            for fi48 := 0; fi48 < 17; fi48++ {
                var v49 uint = 1
                v2 = append(v2, fi48)
                if (0.1 != float32((float64(v26) - float64(v26)))) {
                    for fi50 := 0; fi50 < 5; fi50++ {
                        switch fi48 {
                        case 1:
                            if true {
                                if ((v5 + 1) >= v5) {
                                    var v51 int32 = -2147483648
                                    var v52 uint64 = v3
                                    _ = v51
                                    _ = v52
                                }
                                fmt.Println(v26, v3)
                                var v53 int8 = (v5 << (((v49 - (v49 << (((^(v49))) & 63)))) & 63))
                                _ = v53
                            } else {
                                var v54 bool = true
                                v1["hello"] = ("a" + "foo bar")
                                _ = v54
                            }
                        }
                        fmt.Println(v49, fi48)
                        _ = fi50
                    }
                    v55 := v1["hello"]
                    var v56 uint64 = v3
                    _ = v55
                    _ = v56
                } else {
                    var v57 uint32 = 754
                    _ = v57
                }
                _ = fi48
                _ = v49
            }
            v58 := GenS6{F0: (v5 & int8(v3)), F1: -1, F2: (v3 + v3)}
            _ = v58
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v5
}
