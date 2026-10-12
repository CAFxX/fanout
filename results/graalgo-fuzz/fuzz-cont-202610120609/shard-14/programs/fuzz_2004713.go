package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 uint32; F1 uint32; F2 uint }
type GenS6 struct { F0 int16; F1 int32; F2 int8 }
type GenS7 struct { F0 bool; F1 string; F2 int16 }

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[int]int)
    var v3 bool = (false || (false && ((0 < 0) || true)))
    v4 := make(map[string]string)
    var v5 string = "a"
    var v6 int8 = -51
    v1 = 42
    if (!true) {
        if (127 == 2147483647) {
            var v7 int32 = (v1 ^ 1)
            _ = v7
        } else {
            var v8 uint64 = 76
            v9 := GenS1{F0: (v6 + 42), F1: v5}
            v6 = (-(v6))
            _ = v8
            _ = v9
        }
        for fi10 := 0; fi10 < 2; fi10++ {
            for fi11 := 0; fi11 < 2; fi11++ {
                var v12 float32 = ((1e10 / -2.25) - (3.14159 - float32(0.0)))
                _ = fi11
                _ = v12
            }
            v6 = v6
            switch (fi10 * fi10) {
            case 42:
                if v3 {
                    v3 = v3
                    if false {
                        v13 := GenS2{F0: 42, F1: 255}
                        v6 += -65
                        var v14 string = ""
                        _ = v13
                        _ = v14
                    } else {
                        for fi15 := 0; fi15 < 6; fi15++ {
                            v16 := v2[int(v1)]
                            switch "a" {
                            case "":
                                if false {
                                    fmt.Println(v1)
                                    var v17 float32 = float32(float64(-2.25))
                                    _ = v17
                                }
                            default:
                                var v18 uint = 0
                                _ = v18
                            }
                            v19 := []int{(fi10 ^ fi15)}
                            _ = fi15
                            _ = v16
                            _ = v19
                        }
                        var v20 int16 = 0
                        v21 := GenS1{F0: int8(fi10), F1: (v5 + "hello")}
                        _ = v20
                        _ = v21
                    }
                    v22 := []int{(fi10 % (((^(fi10))) | 1)), fi10, fi10}
                    _ = v22
                }
            case 32768:
                for fi23 := 0; fi23 < 3; fi23++ {
                    switch fi10 {
                    case 1:
                        var v24 uint = 458
                        _ = v24
                    case -549:
                        if (-836 > 9223372036854775807) {
                            fmt.Println(fi10)
                            v25 := GenS3{F0: 195}
                            var v26 float32 = -2.25
                            v27 := GenS2{F0: 0, F1: 0}
                            switch "hello" {
                            case "foo bar":
                                if v3 {
                                    switch (v5 + v5) {
                                    case "foo bar":
                                        v28 := GenS4{F0: v5, F1: 726}
                                        _ = v28
                                    case "hello":
                                        v5 = v5
                                    default:
                                        _ = len(v4)
                                    }
                                    var v29 int8 = (-(v6))
                                    v5 = v5
                                    _ = v29
                                } else {
                                    switch -643 {
                                    case -2147483648:
                                        var v30 uint16 = 677
                                        _ = v30
                                    case 1:
                                        _ = len(v4)
                                    }
                                    var v31 uint64 = 1
                                    v32 := GenS2{F0: 1, F1: 844}
                                    switch v5 {
                                    case "":
                                        delete(v4, (v5 + (v5 + "")))
                                    }
                                    _ = v31
                                    _ = v32
                                }
                            case "hello":
                                v33 := []int{(fi10 + fi23)}
                                _ = v33
                            case "":
                                if v3 {
                                    v34 := GenS5{F0: 639, F1: 65535, F2: 865}
                                    _ = v34
                                }
                            default:
                                fmt.Println(v6, v3, v1)
                            }
                            v35 := []int{2147483647, fi23, fi23}
                            v26 -= 0.1
                            _ = v25
                            _ = v26
                            _ = v27
                            _ = v35
                        }
                        v6 = 100
                    case 9223372036854775807:
                        v36 := GenS6{F0: -619, F1: (v1 * 42), F2: (-(v6))}
                        _ = v36
                    }
                    v6 = (^(v6))
                    _ = fi23
                }
            case 2147483647:
                v37 := []int{(^(fi10)), int(v6), 2147483647}
                _ = v37
            }
            _ = fi10
        }
    }
    v38 := GenS1{F0: -128, F1: v5}
    var v39 int64 = -1
    if (float64(0.1) != (float64(((1.5 * (float32(-2.25) / 1.0)) * -2.25)) * (1.5 + 3.14159))) {
        v40 := []int{-477, -2147483648, 32767}
        v41 := v2[-470]
        v3 = (2 != 0)
        _ = v40
        _ = v41
    }
    if ((597 != 2) || (-9223372036854775808 <= 2147483647)) {
        v42 := GenS1{F0: -127, F1: v5}
        var v43 uint16 = 42
        v44 := GenS7{F0: (951 < 0), F1: v5, F2: 0}
        _ = v42
        _ = v43
        _ = v44
    } else {
        var v45 int32 = v1
        var v46 bool = (0 != 66)
        _ = v45
        _ = v46
    }
    delete(v4, (v5 + ((v5 + v5) + (((v5 + "") + v5) + ((((((v5 + v5) + (v5 + v5)) + "foo bar") + ((v5 + (v5 + (v5 + "a"))) + (v5 + v5))) + ("foo bar" + "\n\t")) + (("a" + v5) + (((v5 + "") + v5) + ("foo bar" + v5))))))))
    fmt.Println(v6, v1)
    v1 += (v1 * 1)
    v47 := GenS7{F0: v3, F1: (v5 + ("foo bar" + v5)), F2: -32768}
    var v48 bool = v3
    fmt.Println(v5, v39, v48)
    v48 = true
    switch 42 {
    case -330:
        var v49 uint64 = 129
        _ = v49
    default:
        var v50 uint16 = 65535
        _ = v50
    }
    for fi51 := 0; fi51 < 11; fi51++ {
        v52 := GenS7{F0: v3, F1: v5, F2: 32767}
        _ = fi51
        _ = v52
    }
    if v48 {
        v6 = v6
        fmt.Println(v39)
        for fi53 := 0; fi53 < 2; fi53++ {
            switch 128 {
            case 23:
                fmt.Println(fi53)
            }
            var v54 int = (fi53 ^ fi53)
            v48 = (!(int64(v54) >= (v39 << ((2) & 63))))
            _ = fi53
            _ = v54
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
