package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string; F1 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int32 }
func genFn0() uint8 {
    v6 := GenS1{F0: 0}
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 int = 32768
        _ = fi7
        _ = v8
    }
    var v9 uint64 = 4294967295
    v10 := GenS1{F0: 1}
    _ = v10
    _ = v6
    _ = v9
    return 60
}

func main() {
    var v1 int64 = 1
    var v2 uint16 = 158
    v3 := []uint32{42, 1, 2147483648}
    var v4 int8 = 1
    var v5 int64 = -1
    if (!(-461 != 1)) {
        var v11 int = -821
        switch (("\n\t" + "a") + ("hello" + ("" + "foo bar"))) {
        case "":
            var v12 int = 32767
            _ = v12
        case "hello":
            v13 := GenS2{F0: 544}
            _ = v13
        default:
            if (2 < 2) {
                v3 = append(v3, 2147483648)
                if ((!true) || (!false)) {
                    var v14 float64 = ((3.14159 * (float64(float32(0.1)) * -2.25)) / (0.0 + 3.14159))
                    for fi15 := 0; fi15 < 10; fi15++ {
                        var v16 uint = 240
                        if true {
                            for fi17 := 0; fi17 < 7; fi17++ {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = fi17
                            }
                            v1 = v5
                        } else {
                            v16 = 623
                        }
                        fmt.Println(fi15, v11, v5)
                        _ = fi15
                        _ = v16
                    }
                    var v18 uint32 = 33
                    _ = v14
                    _ = v18
                } else {
                    fmt.Println(v4)
                    v19 := GenS3{F0: 18446744073709551615}
                    _ = v19
                }
                switch (v11 ^ v11) {
                case 792:
                    var v20 uint64 = 4294967295
                    _ = v20
                case -211:
                    var v21 uint64 = 1
                    _ = v21
                case -1:
                    var v22 uint16 = 2
                    _ = v22
                }
            }
        }
        var v23 uint64 = 42
        _ = v11
        _ = v23
    }
    v24 := GenS4{F0: "\n\t", F1: ""}
    for fi25 := 0; fi25 < 14; fi25++ {
        v26 := GenS5{F0: 0}
        _ = fi25
        _ = v26
    }
    var v27 int64 = v5
    switch (("" + "a") + "foo bar") {
    case "\n\t":
        switch "hello" {
        case "a":
            var v28 float32 = ((0.1 / 1.0) / 3.14159)
            _ = v28
        case "":
            v27 = v27
        case "foo bar":
            switch 342 {
            case 127:
                for fi29 := 0; fi29 < 3; fi29++ {
                    var v30 uint32 = 0
                    for fi31 := 0; fi31 < 14; fi31++ {
                        fmt.Println(fi29)
                        _ = len(v3)
                        v1 = -94
                        fmt.Println(v27, v4)
                        v27 = v5
                        if true {
                            var v32 float64 = 0.1
                            switch fi31 {
                            case -181:
                                v1 = (v27 + -9223372036854775808)
                            case 2147483647:
                                v1 = (v1 / ((v1) | 1))
                            default:
                                v32 += 0.1
                            }
                            switch "foo bar" {
                            case "a":
                                var v33 float32 = 3.14159
                                _ = v33
                            case "":
                                fmt.Println(v1, v1, fi29)
                            case "foo bar":
                                v27 -= v1
                                if true {
                                    var v34 uint64 = 0
                                    v35 := GenS6{F0: fi31}
                                    v27 = 863
                                    _ = v34
                                    _ = v35
                                } else {
                                    var v36 int8 = (v4 >> ((2) & 63))
                                    fmt.Println(v2, v2, v32)
                                    _ = v36
                                }
                            default:
                                v5 = (v1 | v1)
                            }
                            _ = v32
                        } else {
                            v3 = append(v3, v30)
                        }
                        _ = fi31
                    }
                    _ = fi29
                    _ = v30
                }
            case 138:
                v1 -= (v5 & 42)
            default:
                var v37 uint32 = 1
                _ = v37
            }
        }
    case "":
        var v38 float64 = (((1.5 + 0.0) * (1.5 + -2.25)) + 1e10)
        _ = v38
    default:
        for fi39 := 0; fi39 < 18; fi39++ {
            var v40 int = -428
            v1 = v5
            v4 = v4
            _ = fi39
            _ = v40
        }
    }
    v2 -= v2
    var v41 int32 = 221
    switch 685 {
    case 4294967295:
        v27 = 441
    case 128:
        if false {
            for fi42 := 0; fi42 < 18; fi42++ {
                v5 = v1
                var v43 int32 = (v41 ^ 142)
                var v44 string = "foo bar"
                _ = fi42
                _ = v43
                _ = v44
            }
            v27 = v1
            var v45 float32 = ((1.5 - (float32(((0.1 - 1e10) + 1.5)) + float32((((float64(1.5) / 1.0) - 3.14159) * 1e10)))) * 1e10)
            _ = v45
        }
    }
    for fi46 := 0; fi46 < 10; fi46++ {
        for fi47 := 0; fi47 < 15; fi47++ {
            v48 := GenS7{F0: (v41 << ((847) & 63))}
            fmt.Println(v1)
            _ = fi47
            _ = v48
        }
        _ = fi46
    }
    v49 := GenS3{F0: 4294967295}
    v41 += 223
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v41
    _ = v49
    _ = v5
}
