package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 string; F1 uint8; F2 bool }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 bool }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]string)
    v3 := make(map[int]string)
    v4 := []string{("\n\t" + ""), "", "\n\t", ("hello" + "foo bar")}
    v5 := make(map[int]uint32)
    v6 := []string{""}
    var v7 uint16 = 42
    fmt.Println(v7)
    _ = len(v1)
    switch "a" {
    case "":
        var v8 string = ((("foo bar" + "foo bar") + (("hello" + "a") + ("hello" + "hello"))) + "\n\t")
        _ = v8
    case "hello":
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 int64 = 1
            fmt.Println(v10, v7, fi9)
            _ = fi9
            _ = v10
        }
    case "a":
        var v11 int = 56
        _ = v11
    default:
        if (-1 <= -63) {
            if (((262 < 1) || (!(!false))) || true) {
                fmt.Println(v7)
                if true {
                    for fi12 := 0; fi12 < 1; fi12++ {
                        v7 -= (v7 ^ 1)
                        if len(v4) > 0 {
                            v4[0] = "a"
                        }
                        _ = fi12
                    }
                    _ = len(v6)
                }
            } else {
                v13 := GenS1{F0: 1, F1: 2}
                _ = v13
            }
        } else {
            var v14 float64 = (1.5 * 1.5)
            _ = v14
        }
    }
    v3[42] = ""
    var v15 int = 2147483647
    v2[2147483648] = ""
    var v16 int = v15
    if true {
        var v17 uint64 = 1
        if true {
            if (68 == 127) {
                switch v16 {
                case -507:
                    fmt.Println(v16)
                }
            } else {
                for fi18 := 0; fi18 < 11; fi18++ {
                    fmt.Println(v15, v7)
                    v17 = 42
                    v19 := GenS2{F0: "", F1: 2, F2: true}
                    _ = fi18
                    _ = v19
                }
            }
        } else {
            v20 := GenS1{F0: 9223372036854775807, F1: v17}
            var v21 uint8 = 128
            var v22 string = ("a" + ("a" + "foo bar"))
            _ = v20
            _ = v21
            _ = v22
        }
        switch (v16 * int(v7)) {
        case -2147483648:
            if true {
                v17 += 23
            } else {
                fmt.Println(v17, v16)
            }
        default:
            if (-64 == 24) {
                var v23 float64 = 0.0
                for fi24 := 0; fi24 < 4; fi24++ {
                    var v25 int64 = -9223372036854775808
                    _ = fi24
                    _ = v25
                }
                var v26 int64 = -308
                _ = v23
                _ = v26
            }
        }
        _ = v17
    }
    var v27 float64 = 0.1
    v28 := GenS2{F0: "\n\t", F1: 147, F2: false}
    if false {
        v15 -= (v16 << ((42) & 63))
    }
    if true {
        if true {
            var v29 int16 = 0
            _ = v29
        } else {
            for fi30 := 0; fi30 < 13; fi30++ {
                var v31 string = ("foo bar" + ("hello" + ("" + ((("foo bar" + "\n\t") + "\n\t") + "a"))))
                fmt.Println(v31, v31, v27)
                v32 := v5[(v16 + v15)]
                _ = fi30
                _ = v31
                _ = v32
            }
        }
        v33 := GenS1{F0: -9223372036854775808, F1: 18446744073709551615}
        _ = v33
    } else {
        if (-32768 > -1) {
            switch (^(v15)) {
            case 2147483647:
                if false {
                    for fi34 := 0; fi34 < 15; fi34++ {
                        switch fi34 {
                        case -2147483648:
                            for fi35 := 0; fi35 < 18; fi35++ {
                                switch v15 {
                                case -2147483648:
                                    v36 := GenS1{F0: 9223372036854775807, F1: 0}
                                    _ = v36
                                case -9223372036854775808:
                                    v27 -= (1e10 - v27)
                                case 0:
                                    var v37 uint32 = 985
                                    _ = v37
                                }
                                if false {
                                    if ((uint32(v7) > 4294967295) || false) {
                                        var v38 int16 = 1
                                        _ = v38
                                    }
                                    v6 = append(v6, "foo bar")
                                    switch "" {
                                    case "hello":
                                        _ = len(v4)
                                    case "\n\t":
                                        var v39 float64 = ((1e10 * 0.0) - ((-2.25 / 0.1) + v27))
                                        _ = v39
                                    case "":
                                        v7 = (v7 + (^(v7)))
                                    default:
                                        var v40 bool = true
                                        _ = v40
                                    }
                                }
                                v7 = (v7 - (v7 * v7))
                                switch (("a" + ("hello" + "a")) + "foo bar") {
                                case "a":
                                    if true {
                                        v15 = fi34
                                        var v41 int64 = -848
                                        _ = v41
                                    } else {
                                        v7 -= v7
                                    }
                                case "\n\t":
                                    var v42 uint64 = 4294967295
                                    _ = v42
                                case "foo bar":
                                    v43 := GenS3{F0: true, F1: 2147483647}
                                    _ = v43
                                }
                                _ = fi35
                            }
                        case 256:
                            fmt.Println(v15, v15, v15)
                        case 4294967295:
                            v16 -= (-(v15))
                        }
                        _ = fi34
                    }
                    v44 := GenS1{F0: 42, F1: 2}
                    _ = v44
                }
            default:
                if (100 < -1) {
                    var v45 int32 = 2147483647
                    v15 -= v15
                    _ = v45
                } else {
                    v16 = (v16 / (((v16 + 4294967295)) | 1))
                }
            }
            for fi46 := 0; fi46 < 14; fi46++ {
                v6 = append(v6, ("foo bar" + "a"))
                var v47 uint8 = 1
                _ = fi46
                _ = v47
            }
            switch (v15 & (-(v15))) {
            case -1:
                for fi48 := 0; fi48 < 4; fi48++ {
                    if false {
                        var v49 uint8 = 128
                        v50 := GenS4{F0: true}
                        v49 = v49
                        _ = v49
                        _ = v50
                    }
                    _ = fi48
                }
            }
        } else {
            v27 = (v27 / v27)
        }
        for fi51 := 0; fi51 < 10; fi51++ {
            var v52 uint8 = 1
            var v53 uint8 = 42
            _ = fi51
            _ = v52
            _ = v53
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
