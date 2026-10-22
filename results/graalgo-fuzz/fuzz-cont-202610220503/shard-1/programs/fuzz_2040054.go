package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 int; F1 string; F2 uint }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 uint64; F1 uint64; F2 bool }
type GenS6 struct { F0 int16 }

func main() {
    v1 := []string{}
    v2 := []uint32{}
    v3 := make(map[int]uint32)
    var v4 int16 = 348
    v4 += -1
    var v5 float64 = 0.1
    if (158 > 0) {
        var v6 int8 = 27
        _ = v6
    } else {
        if len(v1) > 0 {
            v1[0] = ""
        }
        var v7 int = 2147483647
        _ = v7
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        for fi9 := 0; fi9 < 7; fi9++ {
            if (!false) {
                v4 = v4
                v4 += (v4 % ((v4) | 1))
                v10 := GenS1{F0: 127, F1: 128}
                _ = v10
            } else {
                switch "foo bar" {
                case "a":
                    var v11 uint8 = 127
                    _ = v11
                case "foo bar":
                    delete(v3, (fi8 - fi9))
                default:
                    if (700 > 526) {
                        if true {
                            switch "foo bar" {
                            case "hello":
                                fmt.Println(fi8)
                            case "\n\t":
                                var v12 int16 = (v4 | 1000)
                                if ((644 != 1) && true) {
                                    v5 = v5
                                } else {
                                    v13 := GenS2{F0: 460, F1: 4294967295}
                                    var v14 string = (("hello" + "hello") + "hello")
                                    _ = v13
                                    _ = v14
                                }
                                v4 -= v12
                                _ = v12
                            case "foo bar":
                                switch fi9 {
                                case 350:
                                    if false {
                                        switch fi8 {
                                        case 128:
                                            v5 = v5
                                        case -548:
                                            v4 = v4
                                        }
                                    }
                                }
                            default:
                                if false {
                                    v2 = append(v2, 42)
                                    v5 -= ((v5 + v5) - ((v5 - 1e10) + 3.14159))
                                }
                            }
                            v15 := v3[0]
                            _ = v15
                        }
                    }
                    fmt.Println(fi9)
                    var v16 bool = false
                    _ = v16
                }
                switch fi9 {
                case 339:
                    v4 += v4
                default:
                    switch "foo bar" {
                    case "a":
                        var v17 uint8 = 1
                        _ = v17
                    case "foo bar":
                        fmt.Println(v4, fi8, fi9)
                    default:
                        var v18 int16 = v4
                        _ = v18
                    }
                }
                v19 := GenS2{F0: 1, F1: 2}
                _ = v19
            }
            _ = fi9
        }
        var v20 int64 = 42
        _ = fi8
        _ = v20
    }
    v1 = append(v1, "foo bar")
    v21 := GenS3{F0: 128, F1: "hello", F2: 1}
    for fi22 := 0; fi22 < 12; fi22++ {
        if (true || true) {
            v5 = (((1e10 + (v5 * ((-2.25 * 1.5) + (1e10 + ((v5 - v5) + v5))))) * v5) / 1.5)
        }
        for fi23 := 0; fi23 < 20; fi23++ {
            v4 -= -32768
            switch fi23 {
            case 4294967295:
                v24 := GenS4{F0: 1, F1: 0}
                _ = v24
            default:
                v5 = float64(0.1)
            }
            var v25 uint = 0
            v25 += (v25 + v25)
            _ = fi23
            _ = v25
        }
        _ = fi22
    }
    if false {
        var v26 uint32 = 42
        for fi27 := 0; fi27 < 9; fi27++ {
            v28 := GenS5{F0: 0, F1: 2, F2: (!false)}
            v29 := GenS2{F0: 4294967295, F1: 1}
            _ = fi27
            _ = v28
            _ = v29
        }
        v5 += float64(3.14159)
        _ = v26
    }
    var v30 float32 = 0.0
    for fi31 := 0; fi31 < 1; fi31++ {
        var v32 uint32 = 393
        var v33 int8 = 100
        _ = fi31
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 15; fi34++ {
        if true {
            for fi35 := 0; fi35 < 6; fi35++ {
                if (!(-509 > -2147483648)) {
                    if false {
                        v30 = (v30 * 3.14159)
                    } else {
                        var v36 uint8 = 1
                        v37 := GenS6{F0: int16(fi35)}
                        fmt.Println(v4)
                        _ = v36
                        _ = v37
                    }
                    var v38 int16 = -1
                    var v39 string = ""
                    _ = v38
                    _ = v39
                }
                _ = fi35
            }
            switch (fi34 ^ 32767) {
            case 4294967295:
                var v40 float32 = v30
                _ = v40
            case 335:
                switch 127 {
                case 9223372036854775807:
                    if true {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        var v41 uint8 = 114
                        var v42 int8 = 0
                        _ = v41
                        _ = v42
                    }
                case -2147483648:
                    var v43 int16 = v4
                    _ = v43
                default:
                    var v44 float64 = 1e10
                    _ = v44
                }
            }
            v4 += 32767
        }
        fmt.Println(v5)
        _ = fi34
    }
    var v45 float32 = v30
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v4
    _ = v45
    _ = v5
}
