package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 uint32; F2 int64 }
type GenS4 struct { F0 int32; F1 uint64; F2 uint32 }
type GenS5 struct { F0 uint32; F1 bool; F2 uint16 }
type GenS6 struct { F0 bool; F1 string; F2 int64 }
func genFn0() int16 {
    var v7 uint32 = 490
    switch 42 {
    case 42:
        v7 += 59
    case 9223372036854775807:
        v8 := GenS1{F0: (-1 != 9223372036854775807), F1: -1, F2: ("" + ("foo bar" + ""))}
        _ = v8
    default:
        var v9 uint = 171
        _ = v9
    }
    if ((((1e10 - 3.14159) - -2.25) == 0.0) || (417 < 4294967295)) {
        v7 = (v7 | v7)
    } else {
        var v10 uint32 = (v7 - 65535)
        _ = v10
    }
    switch 128 {
    case -2147483648:
        v11 := GenS1{F0: true, F1: 501, F2: ("hello" + (("\n\t" + ("" + "\n\t")) + "a"))}
        _ = v11
    case -1:
        v12 := []int{2147483647}
        _ = v12
    case 2147483647:
        switch 2147483647 {
        case 2147483647:
            var v13 int = 127
            _ = v13
        case 32767:
            v14 := []int{-705, 42, -1, 0}
            _ = v14
        default:
            v15 := []int{42, 127, -980}
            _ = v15
        }
    }
    fmt.Println(v7)
    _ = v7
    return -32768
}
func genFn1(p0 int32) uint {
    var v16 uint64 = 2
    v16 += v16
    v17 := []int{}
    if (false && ((2 != 1) || true)) {
        if true {
            switch ("foo bar" + "hello") {
            case "":
                fmt.Println(p0)
            default:
                v18 := GenS2{F0: (v16 & 0)}
                _ = v18
            }
            v19 := GenS1{F0: (245 <= 1), F1: -9223372036854775808, F2: ""}
            _ = v19
        }
    }
    var v20 int8 = 81
    _ = p0
    _ = v16
    _ = v17
    _ = v20
    return 42
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]string)
    v3 := []int{}
    var v4 int64 = 499
    var v5 int32 = 670
    var v6 uint16 = 2
    for fi21 := 0; fi21 < 12; fi21++ {
        for fi22 := 0; fi22 < 2; fi22++ {
            fmt.Println(fi21)
            var v23 uint8 = 2
            v24 := GenS3{F0: ("\n\t" + ""), F1: 306, F2: int64(v6)}
            _ = fi22
            _ = v23
            _ = v24
        }
        if true {
            var v25 int32 = v5
            if len(v3) > 0 {
                v3[0] = (fi21 * (^(fi21)))
            }
            _ = v25
        }
        delete(v2, "hello")
        _ = fi21
    }
    fmt.Println(v5, v6, v4)
    v4 -= (v4 / ((v4) | 1))
    if (v4 <= (v4 & 1)) {
        if ((v6 / ((0) | 1)) < (v6 ^ v6)) {
            var v26 uint8 = 27
            fmt.Println(v26, v5, v5)
            v27 := GenS4{F0: (v5 | v5), F1: 18446744073709551615, F2: 35}
            _ = v26
            _ = v27
        } else {
            for fi28 := 0; fi28 < 16; fi28++ {
                _ = len(v1)
                v4 += v4
                _ = fi28
            }
        }
        var v29 uint = 4294967295
        _ = v29
    } else {
        for fi30 := 0; fi30 < 15; fi30++ {
            for fi31 := 0; fi31 < 4; fi31++ {
                v6 -= (v6 | v6)
                if false {
                    switch "" {
                    case "\n\t":
                        if (true || true) {
                            var v32 string = (("\n\t" + "a") + ("hello" + "hello"))
                            switch v32 {
                            case "a":
                                v4 -= v4
                            case "":
                                if true {
                                    switch fi31 {
                                    case 2147483647:
                                        v6 += v6
                                        switch fi31 {
                                        case 414:
                                            v6 -= 525
                                        case 229:
                                            var v33 int16 = genFn0()
                                            _ = v33
                                        default:
                                            var v34 bool = true
                                            _ = v34
                                        }
                                        fmt.Println(v4, v6, v32)
                                    default:
                                        if true {
                                            v32 = (v32 + (("foo bar" + "hello") + "a"))
                                            if (false || true) {
                                                _ = len(v2)
                                                v32 = v32
                                                fmt.Println(fi31, v5)
                                            }
                                            var v35 int16 = 1000
                                            _ = v35
                                        }
                                    }
                                }
                            case "\n\t":
                                v6 += v6
                            default:
                                var v36 uint64 = 4294967295
                                _ = v36
                            }
                            fmt.Println(v4, v4)
                            _ = v32
                        } else {
                            var v37 int = (fi30 ^ fi30)
                            if false {
                                v38 := GenS3{F0: ("foo bar" + ((((("a" + "\n\t") + "foo bar") + "") + "foo bar") + ("\n\t" + "hello"))), F1: 0, F2: v4}
                                fmt.Println(v5)
                                _ = v38
                            }
                            var v39 string = ("hello" + "foo bar")
                            v39 = v39
                            if false {
                                v5 = (v5 ^ 753)
                                switch (v37 % ((fi30) | 1)) {
                                case -1:
                                    v40 := GenS2{F0: 4294967295}
                                    _ = v40
                                case 788:
                                    if (!true) {
                                        if (288 > 255) {
                                            delete(v2, v39)
                                            fmt.Println(v39, v39)
                                        } else {
                                            fmt.Println(fi30, v5)
                                            fmt.Println(v6, v5, fi30)
                                            if (!false) {
                                                if len(v3) > 0 {
                                                    v3[0] = v37
                                                }
                                            }
                                            v41 := GenS3{F0: "", F1: 510, F2: (v4 * (v4 - v4))}
                                            _ = v41
                                        }
                                    } else {
                                        fmt.Println(v39, fi30)
                                        var v42 int8 = -14
                                        _ = v42
                                    }
                                default:
                                    var v43 int = (fi31 + fi31)
                                    _ = v43
                                }
                                var v44 int64 = v4
                                _ = v44
                            }
                            _ = v37
                            _ = v39
                        }
                    }
                    var v45 string = ("" + "hello")
                    delete(v2, v45)
                    var v46 uint32 = 2147483647
                    var v47 int16 = 0
                    var v48 int32 = v5
                    var v49 uint64 = 1
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v50 uint16 = 25
                    v6 += 1
                    _ = v45
                    _ = v46
                    _ = v47
                    _ = v48
                    _ = v49
                    _ = v50
                } else {
                    delete(v2, "foo bar")
                    var v51 uint64 = 18446744073709551615
                    v52 := GenS3{F0: "a", F1: 0, F2: (-(v4))}
                    _ = v51
                    _ = v52
                }
                _ = fi31
            }
            v4 += int64(v6)
            for fi53 := 0; fi53 < 10; fi53++ {
                v5 = 1
                v4 -= -617
                v5 += (v5 - (^(v5)))
                _ = fi53
            }
            _ = fi30
        }
        var v54 int16 = int16(v5)
        switch ("foo bar" + "foo bar") {
        case "a":
            v55 := GenS5{F0: 883, F1: false, F2: 65535}
            _ = v55
        case "hello":
            var v56 uint64 = 18446744073709551615
            _ = v56
        case "":
            var v57 int16 = (v54 * v54)
            _ = v57
        }
        _ = v54
    }
    v58 := GenS3{F0: ("foo bar" + "a"), F1: 2, F2: v4}
    var v59 int32 = v5
    v60 := GenS6{F0: ((true || ((!false) && (2 < 2))) && false), F1: "hello", F2: v4}
    var v61 uint = 735
    var v62 int = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v6
    _ = v60
    _ = v61
    _ = v62
}
