package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 int32; F2 string }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 uint64; F1 int8; F2 int16 }
type GenS5 struct { F0 string; F1 uint; F2 uint64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int16; F1 uint32; F2 int16 }
func genFn0() uint32 {
    var v5 uint64 = 4294967295
    for fi6 := 0; fi6 < 6; fi6++ {
        v5 += (v5 | (v5 % ((4294967295) | 1)))
        _ = fi6
    }
    _ = v5
    return 685
}

func main() {
    v1 := []int{128, 955, -1, 2147483647}
    var v2 int8 = 0
    v3 := make(map[string]string)
    var v4 bool = false
    v7 := GenS1{F0: -9223372036854775808}
    var v8 int = 151
    if ((false || v4) || (v4 && (-1 < 150))) {
        for fi9 := 0; fi9 < 15; fi9++ {
            v8 += fi9
            v2 = (v2 & (v2 >> ((18446744073709551615) & 63)))
            _ = fi9
        }
    } else {
        fmt.Println(v4, v4)
        if v4 {
            if false {
                for fi10 := 0; fi10 < 9; fi10++ {
                    for fi11 := 0; fi11 < 2; fi11++ {
                        if (854 <= 31) {
                            switch "foo bar" {
                            case "hello":
                                var v12 uint16 = 0
                                _ = v12
                            case "foo bar":
                                if false {
                                    v4 = v4
                                    _ = len(v3)
                                    if true {
                                        v13 := GenS1{F0: v8}
                                        delete(v3, "foo bar")
                                        v2 -= (-(v2))
                                        _ = v13
                                    } else {
                                        v14 := GenS1{F0: (fi11 % (((v8 * (-(fi10)))) | 1))}
                                        v4 = (float32(-2.25) >= (float32(0.1) / (1e10 / float32(((0.0 - (float64(0.0) - 0.1)) / 3.14159)))))
                                        v8 -= v8
                                        _ = v14
                                    }
                                } else {
                                    var v15 uint8 = 49
                                    _ = v15
                                }
                            case "":
                                v4 = (!(32767 == 65535))
                            default:
                                if v4 {
                                    v16 := GenS2{F0: 1, F1: 0, F2: ""}
                                    var v17 int = fi10
                                    switch (^(fi10)) {
                                    case 114:
                                        var v18 uint = 4294967295
                                        _ = v18
                                    case -198:
                                        var v19 uint = 65535
                                        _ = v19
                                    case -736:
                                        _ = len(v1)
                                    default:
                                        var v20 uint8 = 42
                                        _ = v20
                                    }
                                    _ = v16
                                    _ = v17
                                } else {
                                    if true {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        if false {
                                            fmt.Println(v2, fi11, fi11)
                                            v21 := GenS3{F0: 1, F1: 1000}
                                            _ = v21
                                        } else {
                                            switch fi10 {
                                            case 473:
                                                if true {
                                                    fmt.Println(fi10, v4)
                                                } else {
                                                    var v22 bool = (3.14159 == 3.14159)
                                                    _ = v22
                                                }
                                            case 32768:
                                                v8 = fi11
                                            case 128:
                                                _ = len(v1)
                                            }
                                            v4 = (18446744073709551615 != 0)
                                            v2 = (v2 * v2)
                                            switch "\n\t" {
                                            case "foo bar":
                                                switch "hello" {
                                                case "":
                                                    switch ("\n\t" + "") {
                                                    case "foo bar":
                                                        if v4 {
                                                            if v4 {
                                                                v23 := GenS3{F0: -561, F1: 1}
                                                                v8 = (fi10 + (fi11 & fi10))
                                                                v4 = (v4 && true)
                                                                _ = v23
                                                            } else {
                                                                if (false && v4) {
                                                                    var v24 int8 = (v2 | v2)
                                                                    _ = v24
                                                                } else {
                                                                    v3["foo bar"] = "\n\t"
                                                                    var v25 uint64 = 961
                                                                    _ = v25
                                                                }
                                                                if (!(32767 <= -1)) {
                                                                    v26 := GenS4{F0: 2, F1: -35, F2: 284}
                                                                    _ = v26
                                                                } else {
                                                                    var v27 int32 = 1
                                                                    _ = v27
                                                                }
                                                                switch fi10 {
                                                                case 128:
                                                                    var v28 uint64 = 2
                                                                    _ = v28
                                                                case -77:
                                                                    _ = len(v3)
                                                                case -155:
                                                                    var v29 uint32 = 4294967295
                                                                    _ = v29
                                                                default:
                                                                    v1 = append(v1, (fi11 << ((393) & 63)))
                                                                }
                                                            }
                                                            var v30 bool = v4
                                                            _ = v30
                                                        }
                                                    default:
                                                        v2 -= 0
                                                    }
                                                default:
                                                    v2 += v2
                                                }
                                            }
                                            var v31 uint32 = 65535
                                            var v32 int32 = 670
                                            _ = v31
                                            _ = v32
                                        }
                                    }
                                    fmt.Println(fi10, v8)
                                }
                            }
                            v33 := GenS5{F0: "\n\t", F1: 42, F2: 599}
                            v34 := GenS2{F0: 938, F1: -296, F2: ("a" + "hello")}
                            _ = v33
                            _ = v34
                        }
                        _ = fi11
                    }
                    fmt.Println(v4, v8, v4)
                    _ = fi10
                }
                if false {
                    var v35 int = v8
                    _ = v35
                }
                v8 = v8
            }
        } else {
            v2 = (v2 << ((316) & 63))
            var v36 uint8 = 255
            v8 = (-(v8))
            _ = v36
        }
        fmt.Println(v2, v2, v2)
    }
    for fi37 := 0; fi37 < 11; fi37++ {
        switch "foo bar" {
        case "\n\t":
            _ = len(v1)
        case "hello":
            if ((2 > 0) || true) {
                v38 := GenS2{F0: -342, F1: 206, F2: ""}
                for fi39 := 0; fi39 < 17; fi39++ {
                    v40 := GenS6{F0: fi37}
                    v4 = (false || v4)
                    _ = fi39
                    _ = v40
                }
                v4 = ((v8 * (fi37 * 2147483647)) <= fi37)
                _ = v38
            } else {
                if (!((!v4) || (float64(float32(-2.25)) >= ((0.1 / ((((3.14159 / 1e10) + 0.0) * float64(1e10)) + 1e10)) + ((1e10 - (float64((1e10 * 3.14159)) * 0.0)) + 1.5))))) {
                    v8 = -193
                    v2 = v2
                    var v41 int64 = 751
                    _ = v41
                }
                var v42 uint16 = 172
                _ = v42
            }
        case "foo bar":
            switch v8 {
            case 4294967295:
                if (v4 && true) {
                    v2 += (-(v2))
                }
            case -437:
                switch v8 {
                case 2147483647:
                    _ = len(v1)
                case 32768:
                    v4 = (4294967295 != 2)
                case 0:
                    var v43 uint64 = 4294967295
                    _ = v43
                }
            case 0:
                v8 -= (v8 + (^(fi37)))
            }
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if true {
            switch (^(v8)) {
            case -991:
                for fi44 := 0; fi44 < 8; fi44++ {
                    var v45 int16 = 970
                    _ = fi44
                    _ = v45
                }
            case 128:
                v46 := GenS1{F0: 4294967295}
                _ = v46
            default:
                v47 := GenS6{F0: (fi37 ^ (^(v8)))}
                _ = v47
            }
            var v48 int = 32768
            var v49 uint32 = 2147483648
            _ = v48
            _ = v49
        } else {
            fmt.Println(v8)
        }
        _ = fi37
    }
    var v50 float64 = 0.0
    var v51 int = (v8 / ((v8) | 1))
    var v52 int8 = v2
    for fi53 := 0; fi53 < 11; fi53++ {
        var v54 int8 = (v52 / (((v52 >> ((2) & 63))) | 1))
        if (("hello" + "\n\t") != ("hello" + ((("foo bar" + "") + "a") + "a"))) {
            if (!true) {
                switch "a" {
                case "foo bar":
                    fmt.Println(v52)
                case "hello":
                    v54 = v52
                }
                v55 := GenS7{F0: 0}
                _ = v55
            }
            if v4 {
                var v56 int64 = int64(fi53)
                _ = v56
            }
        } else {
            switch "a" {
            case "\n\t":
                var v57 uint32 = 2
                _ = v57
            default:
                delete(v3, ("foo bar" + ""))
            }
            switch "foo bar" {
            case "foo bar":
                switch 2147483647 {
                case -9223372036854775808:
                    var v58 uint8 = 30
                    _ = v58
                case 2147483647:
                    v54 += (v52 & (-(v54)))
                }
            }
            switch (fi53 * (v8 ^ (v51 % ((-589) | 1)))) {
            case 42:
                v59 := GenS8{F0: 966, F1: 65535, F2: 1000}
                _ = v59
            case 4294967295:
                v60 := GenS5{F0: "hello", F1: 1, F2: 671}
                _ = v60
            }
        }
        if len(v1) > 0 {
            v1[0] = fi53
        }
        _ = fi53
        _ = v54
    }
    v8 += -600
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v50
    _ = v51
    _ = v52
    _ = v7
    _ = v8
}
