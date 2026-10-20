package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 uint64; F1 int; F2 int64 }

func main() {
    var v1 string = ""
    v2 := make(map[string]int)
    v3 := []uint32{}
    var v4 uint8 = 128
    if false {
        fmt.Println(v4, v1)
    } else {
        v2["a"] = 647
        for fi5 := 0; fi5 < 3; fi5++ {
            for fi6 := 0; fi6 < 13; fi6++ {
                var v7 uint64 = 945
                if true {
                    v7 += v7
                    v1 = ((v1 + "foo bar") + v1)
                    v1 = "a"
                } else {
                    if true {
                        if false {
                            if false {
                                if true {
                                    v4 -= 183
                                }
                                var v8 uint64 = v7
                                _ = v8
                            }
                        } else {
                            v3 = append(v3, 65535)
                        }
                        switch "a" {
                        case "foo bar":
                            var v9 uint = 255
                            _ = v9
                        default:
                            v10 := GenS1{F0: -1}
                            _ = v10
                        }
                        fmt.Println(v1)
                    }
                    fmt.Println(v4)
                    v11 := GenS1{F0: -128}
                    _ = v11
                }
                _ = fi6
                _ = v7
            }
            var v12 string = v1
            _ = fi5
            _ = v12
        }
    }
    fmt.Println(v1)
    for fi13 := 0; fi13 < 7; fi13++ {
        fmt.Println(v1)
        v14 := GenS1{F0: -121}
        _ = fi13
        _ = v14
    }
    fmt.Println(v1, v1)
    var v15 uint16 = 1
    if true {
        v4 += (v4 % ((v4) | 1))
        v16 := GenS2{F0: 0}
        for fi17 := 0; fi17 < 16; fi17++ {
            var v18 int16 = -32768
            if false {
                var v19 bool = false
                if v19 {
                    var v20 float32 = (0.0 * (1.5 * (0.1 - 3.14159)))
                    if (!(v19 && false)) {
                        var v21 uint16 = v15
                        _ = v21
                    } else {
                        for fi22 := 0; fi22 < 6; fi22++ {
                            if (127 > 100) {
                                v23 := GenS1{F0: -1}
                                if v19 {
                                    if (v18 != (v18 - v18)) {
                                        if len(v3) > 0 {
                                            v3[0] = 42
                                        }
                                        var v24 string = v1
                                        var v25 uint32 = 2147483647
                                        _ = v24
                                        _ = v25
                                    }
                                    var v26 string = (("" + ("foo bar" + v1)) + (((v1 + v1) + (("a" + "") + v1)) + ""))
                                    if true {
                                        v27 := GenS3{F0: v26}
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = v27
                                    } else {
                                        var v28 uint32 = 1
                                        fmt.Println(fi22, fi22)
                                        v26 = v26
                                        _ = v28
                                    }
                                    fmt.Println(fi22, v15, fi17)
                                    fmt.Println(v4, fi22, v1)
                                    fmt.Println(v18, v4)
                                    _ = v26
                                } else {
                                    var v29 int64 = 1
                                    v18 += (v18 >> ((255) & 63))
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v30 string = (((v1 + ("a" + v1)) + "a") + "a")
                                    _ = v29
                                    _ = v30
                                }
                                _ = v23
                            }
                            switch "a" {
                            case "\n\t":
                                v31 := GenS4{F0: 1, F1: v1}
                                _ = v31
                            case "foo bar":
                                if v19 {
                                    v19 = v19
                                } else {
                                    switch fi22 {
                                    case 497:
                                        var v32 uint64 = 410
                                        _ = v32
                                    case 4294967295:
                                        var v33 int64 = -9223372036854775808
                                        _ = v33
                                    }
                                }
                            default:
                                fmt.Println(v1, v18)
                            }
                            _ = fi22
                        }
                    }
                    _ = v20
                }
                _ = v19
            } else {
                if (0 <= -1) {
                    switch 42 {
                    case 32768:
                        v15 = 977
                    case -9223372036854775808:
                        if true {
                            var v34 string = v1
                            _ = v34
                        }
                    }
                    for fi35 := 0; fi35 < 2; fi35++ {
                        if true {
                            v15 -= v15
                            switch -73 {
                            case 128:
                                v18 = v18
                            case 42:
                                v1 = v1
                            case -1:
                                v3 = append(v3, 1)
                            default:
                                switch (v1 + v1) {
                                case "foo bar":
                                    v15 = (v15 | v15)
                                default:
                                    v36 := GenS2{F0: 0}
                                    _ = v36
                                }
                            }
                            v2[v1] = fi17
                        }
                        v4 = (v4 / (((v4 | (v4 - v4))) | 1))
                        v37 := GenS1{F0: -1}
                        _ = fi35
                        _ = v37
                    }
                }
                fmt.Println(v18)
            }
            for fi38 := 0; fi38 < 18; fi38++ {
                var v39 int16 = v18
                _ = fi38
                _ = v39
            }
            _ = fi17
            _ = v18
        }
        _ = v16
    } else {
        v15 += 255
        var v40 uint64 = 1
        var v41 uint64 = v40
        _ = v40
        _ = v41
    }
    if (true || true) {
        var v42 uint8 = 163
        _ = v42
    } else {
        var v43 int64 = -1
        _ = v43
    }
    var v44 uint16 = v15
    v45 := GenS3{F0: "foo bar"}
    v46 := GenS2{F0: 0}
    v47 := GenS2{F0: 970}
    if false {
        var v48 uint = 497
        _ = v48
    } else {
        v49 := GenS2{F0: 228}
        if ((float64(1e10) - (1e10 / -2.25)) == (3.14159 + float64(0.0))) {
            v4 = (v4 - v4)
        } else {
            if (int32(v4) < -1) {
                switch 128 {
                case -64:
                    v50 := GenS5{F0: 18446744073709551615, F1: 276, F2: 594}
                    _ = v50
                case 42:
                    var v51 int32 = 401
                    _ = v51
                }
                v52 := GenS4{F0: v4, F1: "a"}
                _ = v52
            }
            fmt.Println(v4, v4, v44)
        }
        var v53 float64 = ((-2.25 + 1.5) - 1e10)
        _ = v49
        _ = v53
    }
    _ = len(v2)
    var v54 int64 = 9223372036854775807
    var v55 int16 = -476
    for fi56 := 0; fi56 < 11; fi56++ {
        v57 := GenS2{F0: 1}
        v15 = 1
        if false {
            var v58 int8 = 22
            _ = v58
        }
        _ = fi56
        _ = v57
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v54
    _ = v55
}
