package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int; F1 int16 }
type GenS6 struct { F0 uint; F1 bool; F2 uint32 }
type GenS7 struct { F0 int64; F1 uint }

func main() {
    v1 := make(map[string]bool)
    v2 := []int{-9223372036854775808}
    var v3 int64 = 1
    v4 := []uint32{561, 651, 42}
    v3 = (v3 + (v3 & v3))
    v3 += (-(v3))
    if false {
        for fi5 := 0; fi5 < 18; fi5++ {
            v3 = (v3 % (((v3 & v3)) | 1))
            var v6 int32 = 42
            _ = fi5
            _ = v6
        }
    } else {
        for fi7 := 0; fi7 < 16; fi7++ {
            v8 := GenS1{F0: 42}
            var v9 int16 = 42
            _ = fi7
            _ = v8
            _ = v9
        }
        var v10 string = "hello"
        for fi11 := 0; fi11 < 8; fi11++ {
            switch fi11 {
            case 92:
                v10 = ((v10 + (v10 + "a")) + "")
            default:
                v12 := GenS2{F0: 42}
                _ = v12
            }
            if (!true) {
                for fi13 := 0; fi13 < 5; fi13++ {
                    v3 -= 9223372036854775807
                    v3 = v3
                    if (false || (!false)) {
                        var v14 uint = 255
                        v3 = -702
                        v3 -= v3
                        _ = v14
                    }
                    v3 -= 9223372036854775807
                    _ = fi13
                }
                v10 = "a"
                if (1e10 <= (0.0 + 0.0)) {
                    for fi15 := 0; fi15 < 18; fi15++ {
                        var v16 int16 = -900
                        fmt.Println(fi11, fi11)
                        if false {
                            v10 = v10
                            v1[v10] = false
                            v17 := v1[v10]
                            if true {
                                v18 := GenS3{F0: 2, F1: 2147483648}
                                _ = v18
                            }
                            _ = v17
                        }
                        v19 := GenS3{F0: 4294967295, F1: 4294967295}
                        var v20 int = (fi15 >> ((65535) & 63))
                        _ = fi15
                        _ = v16
                        _ = v19
                        _ = v20
                    }
                    var v21 uint = 800
                    _ = v21
                } else {
                    var v22 float64 = 1.5
                    if len(v2) > 0 {
                        v2[0] = fi11
                    }
                    if (!(!false)) {
                        if (!false) {
                            switch v10 {
                            case "foo bar":
                                if (477 > 38) {
                                    v10 = "hello"
                                }
                            case "\n\t":
                                v10 = "\n\t"
                            case "a":
                                for fi23 := 0; fi23 < 20; fi23++ {
                                    var v24 int32 = 1
                                    _ = fi23
                                    _ = v24
                                }
                            }
                            var v25 float64 = -2.25
                            _ = v25
                        } else {
                            v22 -= float64(0.0)
                            v3 -= (v3 << ((1) & 63))
                            v10 = v10
                        }
                    }
                    _ = v22
                }
            }
            _ = fi11
        }
        _ = v10
    }
    _ = len(v4)
    switch "" {
    case "foo bar":
        v3 -= (v3 * v3)
    case "a":
        v3 += (v3 | -463)
    default:
        var v26 uint32 = 196
        _ = v26
    }
    var v27 uint8 = 42
    for fi28 := 0; fi28 < 19; fi28++ {
        delete(v1, ("a" + (("foo bar" + "foo bar") + ("a" + "hello"))))
        _ = fi28
    }
    switch "" {
    case "a":
        v27 += v27
    default:
        if (905 < 0) {
            var v29 float32 = (0.1 * 0.0)
            if false {
                for fi30 := 0; fi30 < 12; fi30++ {
                    for fi31 := 0; fi31 < 4; fi31++ {
                        var v32 int64 = -9223372036854775808
                        v1["hello"] = (!true)
                        v27 -= v27
                        _ = fi31
                        _ = v32
                    }
                    var v33 uint64 = 42
                    for fi34 := 0; fi34 < 18; fi34++ {
                        switch "foo bar" {
                        case "\n\t":
                            v35 := GenS1{F0: v33}
                            _ = v35
                        default:
                            v36 := GenS4{F0: ((((("foo bar" + ("" + ("" + ""))) + "hello") + ("foo bar" + "a")) + "") + "\n\t")}
                            _ = v36
                        }
                        v33 = (v33 / ((uint64(v27)) | 1))
                        _ = fi34
                    }
                    _ = fi30
                    _ = v33
                }
                v1["hello"] = true
                switch -1 {
                case 128:
                    if true {
                        for fi37 := 0; fi37 < 1; fi37++ {
                            var v38 int64 = v3
                            var v39 int64 = v3
                            _ = fi37
                            _ = v38
                            _ = v39
                        }
                    } else {
                        v29 -= v29
                        var v40 float64 = 0.1
                        var v41 uint = 2
                        _ = v40
                        _ = v41
                    }
                case 127:
                    var v42 int8 = 46
                    _ = v42
                case 1:
                    v43 := GenS5{F0: -479, F1: 977}
                    _ = v43
                }
            } else {
                switch "a" {
                case "":
                    v2 = append(v2, 128)
                case "a":
                    fmt.Println(v27)
                case "\n\t":
                    switch -1 {
                    case -342:
                        switch 410 {
                        case 9223372036854775807:
                            switch 42 {
                            case 2147483647:
                                var v44 uint = 0
                                _ = v44
                            }
                        }
                    }
                default:
                    var v45 uint8 = v27
                    _ = v45
                }
                if len(v4) > 0 {
                    v4[0] = 42
                }
            }
            _ = v29
        }
    }
    v3 += v3
    var v46 float32 = float32(0.0)
    for fi47 := 0; fi47 < 4; fi47++ {
        var v48 uint = 915
        v3 = 9223372036854775807
        _ = fi47
        _ = v48
    }
    v46 = v46
    for fi49 := 0; fi49 < 19; fi49++ {
        v50 := GenS6{F0: 1, F1: true, F2: uint32(v27)}
        v2 = append(v2, -9223372036854775808)
        v46 += 1e10
        _ = fi49
        _ = v50
    }
    for fi51 := 0; fi51 < 8; fi51++ {
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
        v52 := GenS7{F0: 1, F1: 290}
        fmt.Println(v27, v46, fi51)
        _ = fi51
        _ = v52
    }
    var v53 uint8 = (v27 | v27)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v46
    _ = v53
}
