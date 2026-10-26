package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string; F1 int32 }
type GenS6 struct { F0 int16; F1 uint64; F2 uint }
type GenS7 struct { F0 uint }

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = 3.14159
    var v3 int8 = -1
    var v4 int32 = 42
    if true {
        if (true || (4294967295 >= 0)) {
            if false {
                v5 := v1[1]
                for fi6 := 0; fi6 < 19; fi6++ {
                    v7 := GenS1{F0: 1, F1: false}
                    for fi8 := 0; fi8 < 16; fi8++ {
                        var v9 int8 = v3
                        var v10 int = (fi6 >> ((255) & 63))
                        v5 = "a"
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                    v5 = (v5 + "")
                    _ = fi6
                    _ = v7
                }
                for fi11 := 0; fi11 < 14; fi11++ {
                    for fi12 := 0; fi12 < 1; fi12++ {
                        var v13 uint32 = 117
                        v3 = (-(v3))
                        switch fi12 {
                        case 128:
                            var v14 int16 = -544
                            _ = v14
                        case 0:
                            if ((0 <= 2) || (!(!true))) {
                                fmt.Println(fi12, fi11, fi12)
                                var v15 int16 = int16(v4)
                                var v16 float32 = v2
                                _ = v15
                                _ = v16
                            }
                        }
                        _ = fi12
                        _ = v13
                    }
                    _ = fi11
                }
                _ = v5
            }
            v17 := GenS2{F0: 0, F1: 127}
            if ((v4 - (v4 | -2147483648)) <= (v4 >> ((0) & 63))) {
                if ((true || true) || false) {
                    if (float32(0.1) >= ((float32(-2.25) / v2) + v2)) {
                        v18 := []int{-9223372036854775808, 4294967295}
                        _ = v18
                    }
                    v19 := GenS3{F0: "", F1: v3}
                    if ((v4 >= (v4 | 0)) && false) {
                        for fi20 := 0; fi20 < 20; fi20++ {
                            fmt.Println(v2)
                            fmt.Println(v2, v2, v3)
                            if false {
                                var v21 int16 = -1
                                _ = v21
                            } else {
                                var v22 int32 = (v4 + 2147483647)
                                v4 += (v4 - -411)
                                _ = v22
                            }
                            _ = fi20
                        }
                        v3 += v3
                        v2 += float32(3.14159)
                    }
                    _ = v19
                } else {
                    v23 := GenS4{F0: -1}
                    if (((!true) && (1 <= 297)) && true) {
                        if true {
                            var v24 float64 = (-2.25 / 0.1)
                            _ = v24
                        }
                    }
                    _ = v23
                }
            } else {
                for fi25 := 0; fi25 < 5; fi25++ {
                    if (42 != 0) {
                        var v26 uint8 = 43
                        v27 := GenS3{F0: ("\n\t" + "\n\t"), F1: v3}
                        var v28 int16 = -1
                        _ = v26
                        _ = v27
                        _ = v28
                    } else {
                        for fi29 := 0; fi29 < 7; fi29++ {
                            if true {
                                v4 = v4
                                var v30 int32 = v4
                                if (!(2 > 2)) {
                                    var v31 int16 = -1
                                    delete(v1, 2147483648)
                                    fmt.Println(v4, v3)
                                    _ = v31
                                } else {
                                    if false {
                                        var v32 int16 = 0
                                        if false {
                                            v30 -= (v30 - (v30 | v4))
                                        }
                                        _ = v32
                                    } else {
                                        if ("hello" != "\n\t") {
                                            v3 = v3
                                            v33 := GenS3{F0: "", F1: 100}
                                            _ = v33
                                        } else {
                                            v4 = (v4 - -2147483648)
                                            var v34 uint8 = 88
                                            _ = v34
                                        }
                                        v35 := []int{9223372036854775807, fi29, (fi29 / ((fi29) | 1))}
                                        _ = v35
                                    }
                                }
                                _ = v30
                            }
                            v36 := []int{}
                            fmt.Println(fi29, fi29, v3)
                            _ = fi29
                            _ = v36
                        }
                    }
                    var v37 bool = ((2 <= 0) || true)
                    var v38 uint = 208
                    _ = fi25
                    _ = v37
                    _ = v38
                }
            }
            _ = v17
        }
    }
    switch ("foo bar" + "") {
    case "\n\t":
        v4 += 0
    case "foo bar":
        v3 -= (v3 ^ v3)
    }
    if true {
        fmt.Println(v4, v4, v3)
        var v39 uint8 = 128
        _ = v39
    } else {
        v40 := GenS4{F0: 0}
        _ = v40
    }
    for fi41 := 0; fi41 < 2; fi41++ {
        v42 := GenS1{F0: 42, F1: false}
        _ = fi41
        _ = v42
    }
    v2 = 1e10
    fmt.Println(v3, v2)
    fmt.Println(v3)
    var v43 uint64 = 223
    v44 := []int{151, -2147483648, -1}
    var v45 uint32 = 2147483647
    v43 = (^(v43))
    fmt.Println(v45)
    delete(v1, 4294967295)
    v46 := GenS5{F0: "foo bar", F1: v4}
    v1[(v45 >> ((317) & 63))] = "foo bar"
    if true {
        fmt.Println(v3, v2)
        if true {
            if true {
                fmt.Println(v45, v4, v2)
                v45 = v45
                v4 += v4
            }
            var v47 uint32 = (^(v45))
            if (!true) {
                fmt.Println(v43, v2)
                for fi48 := 0; fi48 < 10; fi48++ {
                    var v49 int32 = (v4 ^ v4)
                    v1[(v47 << ((18446744073709551615) & 63))] = ((("hello" + "foo bar") + "\n\t") + "")
                    for fi50 := 0; fi50 < 9; fi50++ {
                        switch "\n\t" {
                        case "a":
                            var v51 int16 = -1
                            _ = v51
                        case "":
                            v1[(v47 & v45)] = "foo bar"
                        case "foo bar":
                            if (true && true) {
                                v52 := v1[2147483647]
                                fmt.Println(v3, v3)
                                fmt.Println(v43, v52)
                                _ = v52
                            } else {
                                if false {
                                    v4 = (v4 % (((v4 >> ((42) & 63))) | 1))
                                    if true {
                                        v53 := GenS4{F0: -1}
                                        v54 := GenS4{F0: -9223372036854775808}
                                        v2 += 1.5
                                        _ = len(v44)
                                        _ = v53
                                        _ = v54
                                    }
                                    switch fi48 {
                                    case 128:
                                        v45 = (v47 * (v45 & v47))
                                    case -2147483648:
                                        if (true && (!(504 >= -749))) {
                                            var v55 int8 = int8(fi50)
                                            _ = v55
                                        } else {
                                            v56 := GenS6{F0: -765, F1: v43, F2: 2}
                                            v47 -= v47
                                            _ = v56
                                        }
                                    }
                                }
                                fmt.Println(fi48, v4)
                                v57 := GenS6{F0: -1, F1: v43, F2: 2}
                                _ = v57
                            }
                        default:
                            fmt.Println(v4, fi48)
                        }
                        _ = fi50
                    }
                    _ = fi48
                    _ = v49
                }
            }
            _ = v47
        }
    }
    v43 = (v43 & (v43 | 2))
    for fi58 := 0; fi58 < 9; fi58++ {
        fmt.Println(v43, v3, v3)
        switch "foo bar" {
        case "a":
            v4 += (v4 % ((v4) | 1))
        case "hello":
            for fi59 := 0; fi59 < 9; fi59++ {
                delete(v1, v45)
                _ = fi59
            }
        default:
            for fi60 := 0; fi60 < 4; fi60++ {
                v61 := GenS7{F0: 18446744073709551615}
                fmt.Println(v45, v3, v43)
                v45 -= (^(v45))
                _ = fi60
                _ = v61
            }
        }
        var v62 uint = 1
        _ = fi58
        _ = v62
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
}
