package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64; F1 bool; F2 uint }
type GenS4 struct { F0 uint32 }
func genFn0() int32 {
    v7 := []int{32768, -211, 4294967295}
    var v8 uint = 1
    var v9 float32 = float32(3.14159)
    var v10 uint64 = 42
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 168
}
func genFn1(p0 bool, p1 uint16) int16 {
    v11 := GenS1{F0: 1000, F1: 346, F2: -980}
    v12 := GenS2{F0: 42}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 32767
}

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[int]int)
    var v3 uint64 = v1
    v4 := []int{-862, 32767, 2147483647, 441}
    v5 := make(map[string]uint32)
    v6 := []string{"a", "", "a"}
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 uint32 = 2147483647
        fmt.Println(v1)
        _ = fi13
        _ = v14
    }
    if (32767 >= 899) {
        v1 += 42
    }
    v15 := GenS1{F0: 1000, F1: 991, F2: 9223372036854775807}
    for fi16 := 0; fi16 < 9; fi16++ {
        v17 := v5[(("foo bar" + "a") + "\n\t")]
        switch (fi16 & fi16) {
        case 127:
            fmt.Println(fi16, v1, v17)
        case -2147483648:
            for fi18 := 0; fi18 < 6; fi18++ {
                var v19 int16 = genFn1(false, 117)
                _ = fi18
                _ = v19
            }
        default:
            fmt.Println(v17)
        }
        _ = fi16
        _ = v17
    }
    if false {
        v20 := GenS3{F0: -1, F1: (uint8(v3) != 163), F2: 0}
        v3 += (v3 % (((v1 << ((255) & 63))) | 1))
        _ = v20
    } else {
        if false {
            v3 -= v1
            if (!true) {
                v1 += v3
            } else {
                if false {
                    for fi21 := 0; fi21 < 17; fi21++ {
                        delete(v5, "")
                        _ = fi21
                    }
                } else {
                    v3 = 333
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        }
        var v22 int32 = -900
        v22 += 0
        _ = v22
    }
    v1 = (v3 ^ 217)
    for fi23 := 0; fi23 < 13; fi23++ {
        if true {
            for fi24 := 0; fi24 < 15; fi24++ {
                v3 -= (v1 | v3)
                var v25 uint = 2
                if len(v4) > 0 {
                    v4[0] = fi24
                }
                _ = fi24
                _ = v25
            }
            fmt.Println(v1)
            v26 := GenS1{F0: 0, F1: -443, F2: -1}
            _ = v26
        } else {
            if (1 < 65535) {
                fmt.Println(v3, v3, fi23)
            }
        }
        if len(v6) > 0 {
            v6[0] = ("" + "a")
        }
        v2[fi23] = fi23
        _ = fi23
    }
    var v27 float32 = (((1e10 - float32(float64((((0.0 - -2.25) * float32((1.5 - 0.1))) / 1.5)))) + (-2.25 - 0.0)) - 0.0)
    v1 = (^(v1))
    delete(v2, 127)
    if (((!false) || false) && false) {
        if true {
            var v28 int16 = 0
            for fi29 := 0; fi29 < 10; fi29++ {
                for fi30 := 0; fi30 < 7; fi30++ {
                    var v31 float64 = float64(-2.25)
                    if ((("foo bar" + ("a" + "hello")) + "") <= ("hello" + ((("foo bar" + "") + ((((("hello" + ("hello" + (("hello" + (("" + "foo bar") + "foo bar")) + (("a" + (("\n\t" + ("\n\t" + "")) + "")) + "")))) + "\n\t") + "a") + "foo bar") + "a")) + "a"))) {
                        var v32 float32 = v27
                        if false {
                            switch "hello" {
                            case "foo bar":
                                switch fi29 {
                                case -1:
                                    v27 += v32
                                }
                                _ = len(v4)
                            case "a":
                                if true {
                                    var v33 bool = (0.1 > (v32 - v27))
                                    switch 1 {
                                    case 9223372036854775807:
                                        v28 = -886
                                    default:
                                        v34 := GenS4{F0: 65535}
                                        _ = v34
                                    }
                                    var v35 string = "a"
                                    v36 := GenS1{F0: v28, F1: (v28 ^ (v28 - genFn1(v33, 255))), F2: (fi29 | -433)}
                                    _ = v33
                                    _ = v35
                                    _ = v36
                                }
                            case "\n\t":
                                var v37 uint = 18446744073709551615
                                _ = v37
                            }
                            _ = len(v6)
                            v27 += ((3.14159 - 0.0) * (v32 + 1e10))
                            var v38 uint32 = 0
                            switch (fi29 | fi30) {
                            case 127:
                                var v39 int32 = int32(v1)
                                _ = v39
                            case 675:
                                switch "foo bar" {
                                case "a":
                                    if (!true) {
                                        v40 := GenS1{F0: v28, F1: (v28 | v28), F2: (fi29 << ((146) & 63))}
                                        var v41 int8 = 100
                                        fmt.Println(fi30, v1)
                                        _ = v40
                                        _ = v41
                                    }
                                case "\n\t":
                                    if false {
                                        fmt.Println(v31, fi29)
                                        v31 -= 0.0
                                    } else {
                                        v42 := GenS3{F0: int64(v38), F1: false, F2: 624}
                                        if (false || ((!((v28 ^ (v28 / ((genFn1((-1 <= -1), 1)) | 1))) != genFn1(false, 324))) && true)) {
                                            var v43 int = (fi30 >> ((1) & 63))
                                            _ = v43
                                        }
                                        _ = v42
                                    }
                                }
                            case 42:
                                v27 = (v32 * v27)
                            default:
                                var v44 bool = (!true)
                                _ = v44
                            }
                            _ = v38
                        } else {
                            v32 += 0.0
                            v45 := GenS1{F0: v28, F1: v28, F2: (fi30 / ((fi30) | 1))}
                            _ = v45
                        }
                        fmt.Println(v32, v3, v31)
                        _ = v32
                    }
                    _ = fi30
                    _ = v31
                }
                v27 = 1.5
                _ = fi29
            }
            _ = v28
        } else {
            for fi46 := 0; fi46 < 11; fi46++ {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = fi46
            }
            var v47 uint64 = v1
            _ = v47
        }
        for fi48 := 0; fi48 < 18; fi48++ {
            v27 = (v27 / (v27 / (((1.5 - 0.0) + float32((-2.25 - 3.14159))) * (0.0 + v27))))
            _ = fi48
        }
    }
    for fi49 := 0; fi49 < 14; fi49++ {
        switch "foo bar" {
        case "a":
            fmt.Println(v27)
        case "foo bar":
            for fi50 := 0; fi50 < 17; fi50++ {
                v51 := GenS1{F0: 0, F1: 42, F2: 0}
                if (172 >= 2) {
                    var v52 uint8 = 2
                    v52 += (^(v52))
                    _ = v52
                }
                _ = fi50
                _ = v51
            }
        }
        v3 = 18446744073709551615
        _ = fi49
    }
    switch "foo bar" {
    case "":
        if (!(-742 >= 1)) {
            if true {
                for fi53 := 0; fi53 < 15; fi53++ {
                    v54 := GenS4{F0: 449}
                    fmt.Println(v3, v3, v1)
                    var v55 uint8 = 255
                    _ = fi53
                    _ = v54
                    _ = v55
                }
                fmt.Println(v27, v3, v27)
            }
        }
    default:
        var v56 uint8 = 2
        _ = v56
    }
    v5[("" + "foo bar")] = 85
    delete(v2, -2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
