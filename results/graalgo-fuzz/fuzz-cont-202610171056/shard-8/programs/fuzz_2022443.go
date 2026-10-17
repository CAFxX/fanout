package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16; F1 int32; F2 uint32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 string; F1 string; F2 bool }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]string)
    v3 := make(map[string]string)
    var v4 float32 = 0.0
    v5 := make(map[int]string)
    var v6 int32 = 42
    v7 := GenS1{F0: 1, F1: v6}
    switch "\n\t" {
    case "hello":
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 int32 = v6
            var v10 int16 = int16(fi8)
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        v12 := GenS1{F0: -1, F1: v6}
        _ = fi11
        _ = v12
    }
    v4 = 1e10
    v13 := []int{-864, 9223372036854775807, -365, int(v6)}
    var v14 float64 = (0.0 - -2.25)
    v3["hello"] = "foo bar"
    if false {
        for fi15 := 0; fi15 < 9; fi15++ {
            if (("a" >= "hello") || ((-32768 == -354) || (740 > fi15))) {
                fmt.Println(fi15, v6, fi15)
                if false {
                    for fi16 := 0; fi16 < 19; fi16++ {
                        v4 = float32((v14 * (0.0 - v14)))
                        v4 = 0.0
                        v17 := GenS2{F0: 1, F1: (^(fi15))}
                        _ = fi16
                        _ = v17
                    }
                    if (1e10 > float32(float64(float32(float64((0.0 * v4)))))) {
                        v18 := GenS3{F0: -32768}
                        _ = v18
                    }
                }
            }
            if true {
                fmt.Println(v4, v4, fi15)
                if (1e10 >= v14) {
                    if true {
                        var v19 float32 = 1e10
                        var v20 int8 = 100
                        _ = v19
                        _ = v20
                    }
                    if false {
                        if true {
                            for fi21 := 0; fi21 < 17; fi21++ {
                                var v22 uint32 = 28
                                _ = fi21
                                _ = v22
                            }
                            if len(v13) > 0 {
                                v13[0] = 127
                            }
                        } else {
                            v4 = v4
                            fmt.Println(fi15, v4)
                        }
                        v14 = (-2.25 * 0.0)
                        var v23 uint16 = 2
                        _ = v23
                    } else {
                        v24 := GenS2{F0: 0, F1: 4294967295}
                        _ = v24
                    }
                } else {
                    for fi25 := 0; fi25 < 6; fi25++ {
                        fmt.Println(fi15, v4, v14)
                        var v26 uint8 = 90
                        switch "foo bar" {
                        case "a":
                            var v27 bool = true
                            _ = v27
                        case "hello":
                            var v28 bool = false
                            _ = v28
                        }
                        _ = fi25
                        _ = v26
                    }
                }
                if ((true && (v6 <= 2147483647)) && (1 > 1)) {
                    v14 = 1e10
                    var v29 float32 = v4
                    fmt.Println(v4, fi15)
                    _ = v29
                }
            } else {
                var v30 uint = 0
                switch fi15 {
                case 854:
                    if (true && true) {
                        fmt.Println(v30, v6, v6)
                        fmt.Println(fi15, v30)
                        v31 := GenS3{F0: 1000}
                        _ = v31
                    }
                default:
                    var v32 float64 = 3.14159
                    _ = v32
                }
                if true {
                    v33 := GenS4{F0: -1, F1: v6, F2: 2147483647}
                    var v34 uint16 = 32767
                    _ = v33
                    _ = v34
                } else {
                    for fi35 := 0; fi35 < 12; fi35++ {
                        fmt.Println(fi35)
                        var v36 string = "foo bar"
                        v30 = (^(v30))
                        _ = fi35
                        _ = v36
                    }
                    if false {
                        if false {
                            v14 = 1.5
                            _ = len(v3)
                        }
                        v37 := GenS3{F0: 238}
                        _ = v37
                    }
                    if false {
                        if (!false) {
                            for fi38 := 0; fi38 < 20; fi38++ {
                                if (-1 <= -128) {
                                    _ = len(v13)
                                }
                                v4 += v4
                                _ = fi38
                            }
                            var v39 bool = false
                            _ = v39
                        }
                        var v40 int16 = -1
                        v13 = append(v13, fi15)
                        _ = v40
                    } else {
                        for fi41 := 0; fi41 < 2; fi41++ {
                            var v42 int32 = v6
                            if true {
                                if true {
                                    if (0 > 735) {
                                        var v43 string = ""
                                        switch int(v42) {
                                        case -1:
                                            v43 = v43
                                        case 128:
                                            switch (fi15 + fi15) {
                                            case -9223372036854775808:
                                                var v44 int64 = -880
                                                _ = v44
                                            case 128:
                                                if (true || (false || true)) {
                                                    var v45 uint8 = 255
                                                    v46 := GenS4{F0: 743, F1: v42, F2: 65535}
                                                    _ = v45
                                                    _ = v46
                                                } else {
                                                    delete(v1, 0)
                                                    switch ("a" + (v43 + v43)) {
                                                    case "foo bar":
                                                        v4 = v4
                                                    case "hello":
                                                        v42 = (v42 >> ((v30) & 63))
                                                    default:
                                                        var v47 uint = (^(v30))
                                                        _ = v47
                                                    }
                                                }
                                            case 4294967295:
                                                switch "\n\t" {
                                                case "":
                                                    v42 = (^(v42))
                                                case "\n\t":
                                                    fmt.Println(fi41)
                                                case "hello":
                                                    v6 -= (v42 / ((376) | 1))
                                                default:
                                                    fmt.Println(fi15, v6)
                                                }
                                            }
                                        case -2147483648:
                                            var v48 uint32 = 1
                                            v4 -= (v4 - (0.0 + v4))
                                            _ = v48
                                        }
                                        _ = v43
                                    }
                                    v49 := GenS4{F0: -32768, F1: v6, F2: 2147483647}
                                    _ = v49
                                }
                                v13 = append(v13, fi41)
                            }
                            switch "a" {
                            case "foo bar":
                                if len(v13) >= 2 {
                                    v13 = v13[1:len(v13)]
                                }
                            default:
                                var v50 int = (^(fi41))
                                _ = v50
                            }
                            _ = fi41
                            _ = v42
                        }
                    }
                }
                _ = v30
            }
            if true {
                v51 := GenS5{F0: 0}
                var v52 int8 = -128
                _ = v51
                _ = v52
            }
            _ = fi15
        }
        v4 = 1e10
    }
    if len(v13) > 0 {
        v13[0] = 128
    }
    for fi53 := 0; fi53 < 1; fi53++ {
        fmt.Println(v4)
        _ = fi53
    }
    v6 += v6
    v6 += v6
    var v54 float32 = ((v4 - v4) - v4)
    v55 := GenS6{F0: ("\n\t" + "hello"), F1: (("hello" + ((("" + ("hello" + "foo bar")) + "") + "\n\t")) + ("foo bar" + "a")), F2: false}
    v6 = (^(v6))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v6
    _ = v7
}
