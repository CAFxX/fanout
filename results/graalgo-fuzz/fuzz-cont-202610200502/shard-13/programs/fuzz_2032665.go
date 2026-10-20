package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint64 }
type GenS2 struct { F0 int8; F1 int8; F2 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64; F1 uint; F2 string }
func genFn0() uint32 {
    v4 := GenS1{F0: (("hello" + "\n\t") + ""), F1: -1, F2: 2}
    if (1 > 624) {
        v5 := GenS2{F0: -1, F1: 42, F2: 255}
        _ = v5
    } else {
        if true {
            switch 853 {
            case 688:
                var v6 uint64 = 4294967295
                _ = v6
            case 0:
                v7 := []int{-868, 2147483647}
                _ = v7
            default:
                if false {
                    var v8 float64 = 0.0
                    _ = v8
                }
            }
            for fi9 := 0; fi9 < 5; fi9++ {
                v10 := GenS3{F0: 1}
                v11 := GenS1{F0: "foo bar", F1: -1, F2: 2}
                var v12 uint16 = 42
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
            var v13 uint = 65535
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        var v15 int32 = 0
        v15 = v15
        _ = fi14
        _ = v15
    }
    v16 := []int{}
    _ = v16
    _ = v4
    return 42
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{245, 415, 2}
    v3 := make(map[int]uint32)
    var v17 int32 = 1
    var v18 uint64 = 760
    v19 := v1[("foo bar" + "a")]
    for fi20 := 0; fi20 < 11; fi20++ {
        switch "hello" {
        case "hello":
            var v21 int64 = 42
            _ = v21
        }
        if (2147483648 >= 2147483648) {
            if true {
                var v22 uint = 255
                for fi23 := 0; fi23 < 18; fi23++ {
                    if false {
                        if (1 >= -32768) {
                            v22 -= 414
                        } else {
                            fmt.Println(v17, v17)
                        }
                        var v24 int32 = (-(v17))
                        _ = v24
                    } else {
                        switch fi23 {
                        case 2147483647:
                            _ = len(v2)
                        case 32767:
                            v25 := GenS1{F0: "a", F1: -1, F2: 4294967295}
                            _ = v25
                        }
                        delete(v1, "")
                        if false {
                            if ((((-2.25 - 0.0) - (-2.25 / 0.1)) > 1e10) && (uint64(v22) > v18)) {
                                v26 := GenS2{F0: 63, F1: 42, F2: 2}
                                switch v19 {
                                case 1:
                                    if false {
                                        v22 += (v22 / (((v22 * (^(v22)))) | 1))
                                        v22 = 608
                                    }
                                default:
                                    v18 -= (v18 % (((v18 >> ((uint(v18)) & 63))) | 1))
                                }
                                _ = v26
                            }
                            var v27 uint64 = 334
                            v19 = (^(v19))
                            v18 += v27
                            if false {
                                if true {
                                    v18 -= (v27 - (v27 | 1))
                                    if false {
                                        var v28 int = (fi23 / ((884) | 1))
                                        _ = len(v1)
                                        _ = v28
                                    }
                                }
                                _ = len(v2)
                            }
                            var v29 int8 = 127
                            if false {
                                var v30 uint = 2
                                fmt.Println(v18)
                                _ = v30
                            } else {
                                var v31 int32 = 0
                                fmt.Println(v31, v27, v22)
                                _ = v31
                            }
                            _ = v27
                            _ = v29
                        } else {
                            var v32 uint = uint(v18)
                            var v33 float32 = float32(0.1)
                            v34 := GenS3{F0: 9223372036854775807}
                            _ = v32
                            _ = v33
                            _ = v34
                        }
                    }
                    _ = fi23
                }
                fmt.Println(v18)
                _ = v22
            } else {
                for fi35 := 0; fi35 < 10; fi35++ {
                    var v36 float64 = 3.14159
                    _ = fi35
                    _ = v36
                }
                if false {
                    if false {
                        fmt.Println(v18, fi20, v18)
                    }
                    switch (((("foo bar" + "hello") + "") + "") + ("" + (("foo bar" + ("\n\t" + ("foo bar" + "hello"))) + "foo bar"))) {
                    case "a":
                        if (float32(-2.25) >= -2.25) {
                            var v37 uint64 = v18
                            _ = v37
                        }
                    }
                } else {
                    if len(v2) > 0 {
                        v2[0] = 1
                    }
                    switch ("" + "foo bar") {
                    case "":
                        v38 := GenS2{F0: 0, F1: -106, F2: 521}
                        _ = v38
                    }
                }
            }
        } else {
            for fi39 := 0; fi39 < 11; fi39++ {
                if true {
                    v40 := GenS3{F0: 0}
                    v41 := GenS2{F0: -1, F1: -128, F2: 732}
                    fmt.Println(fi20, fi39)
                    _ = v40
                    _ = v41
                }
                v17 = (v17 / (((v17 + 42)) | 1))
                _ = fi39
            }
            if (128 >= 145) {
                var v42 uint = 2
                fmt.Println(fi20, fi20)
                var v43 float32 = 1.5
                _ = v42
                _ = v43
            } else {
                v17 = (v17 * v17)
                v19 -= (fi20 << ((0) & 63))
                v18 = v18
            }
        }
        _ = fi20
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v19 -= (v19 & -822)
    v44 := GenS1{F0: "hello", F1: 476, F2: v18}
    fmt.Println(v17)
    v17 = int32(v19)
    if false {
        switch "hello" {
        case "\n\t":
            v17 -= v17
        case "":
            var v45 int = v19
            _ = v45
        }
    } else {
        v17 += (v17 + 0)
    }
    v19 += (^(v19))
    if true {
        var v46 int16 = -638
        for fi47 := 0; fi47 < 6; fi47++ {
            for fi48 := 0; fi48 < 10; fi48++ {
                fmt.Println(fi47, v19)
                _ = fi48
            }
            v49 := GenS4{F0: (v18 >> ((4294967295) & 63)), F1: 957, F2: ("hello" + "hello")}
            _ = fi47
            _ = v49
        }
        _ = v46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v44
}
