package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 uint16; F1 uint64; F2 int16 }
type GenS3 struct { F0 bool; F1 uint; F2 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint; F1 uint8 }
type GenS6 struct { F0 uint32; F1 uint16 }
type GenS7 struct { F0 int64; F1 uint16 }

func main() {
    var v1 uint32 = 65535
    v2 := make(map[string]string)
    var v3 uint64 = uint64(v1)
    v4 := make(map[uint32]uint32)
    v5 := []string{"a", "foo bar", "", ("hello" + "")}
    if (42 > 644) {
        var v6 float32 = -2.25
        _ = v6
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if (!true) {
        var v7 float64 = 0.1
        if (255 == 290) {
            var v8 int32 = 42
            _ = v8
        } else {
            for fi9 := 0; fi9 < 17; fi9++ {
                v10 := v4[(v1 / (((v1 * (v1 / ((v1) | 1)))) | 1))]
                var v11 string = "foo bar"
                fmt.Println(v3)
                _ = fi9
                _ = v10
                _ = v11
            }
        }
        fmt.Println(v1, v3)
        _ = v7
    }
    v12 := GenS1{F0: 38, F1: "foo bar"}
    v13 := GenS2{F0: 65535, F1: (^(v3)), F2: 1000}
    _ = len(v5)
    if (-2.25 < 1.5) {
        fmt.Println(v3)
        v3 += v3
        switch 128 {
        case -263:
            switch ("hello" + "foo bar") {
            case "hello":
                for fi14 := 0; fi14 < 14; fi14++ {
                    switch fi14 {
                    case -1:
                        switch ((((("hello" + "\n\t") + "\n\t") + "foo bar") + "hello") + "\n\t") {
                        case "foo bar":
                            switch "a" {
                            case "a":
                                fmt.Println(fi14, v1)
                            }
                        }
                    case 32768:
                        v15 := GenS3{F0: true, F1: 4294967295, F2: "hello"}
                        _ = v15
                    case -360:
                        var v16 int64 = 42
                        _ = v16
                    }
                    if (32767 < -841) {
                        if (false && true) {
                            var v17 float64 = 1e10
                            _ = v17
                        }
                        switch fi14 {
                        case 992:
                            v3 -= (^(v3))
                        case 32768:
                            v18 := GenS4{F0: false}
                            _ = v18
                        case 0:
                            var v19 int = 4294967295
                            _ = v19
                        default:
                            if len(v5) > 0 {
                                v5[0] = (("a" + ("" + "foo bar")) + ("foo bar" + "\n\t"))
                            }
                        }
                    }
                    _ = fi14
                }
            default:
                if true {
                    switch 2147483647 {
                    case 0:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    case 107:
                        for fi20 := 0; fi20 < 14; fi20++ {
                            v3 = 2
                            _ = fi20
                        }
                    }
                    var v21 uint32 = 863
                    if true {
                        v1 = v21
                    }
                    _ = v21
                } else {
                    for fi22 := 0; fi22 < 18; fi22++ {
                        if ((-2.25 / 1e10) == (0.0 / 1.0)) {
                            _ = len(v5)
                            v23 := GenS3{F0: (195 >= 4294967295), F1: 638, F2: ("" + "foo bar")}
                            if (42 > 275) {
                                if true {
                                    v1 += (v1 % (((v1 | v1)) | 1))
                                    for fi24 := 0; fi24 < 11; fi24++ {
                                        var v25 string = "a"
                                        fmt.Println(v3, v1, fi22)
                                        if true {
                                            var v26 uint16 = uint16(v1)
                                            if len(v5) > 0 {
                                                v5[0] = ""
                                            }
                                            v25 = "hello"
                                            switch v25 {
                                            case "hello":
                                                var v27 float32 = (float32(1.5) / float32((float64(-2.25) / -2.25)))
                                                _ = v27
                                            case "\n\t":
                                                v28 := GenS3{F0: false, F1: 1, F2: (v25 + v25)}
                                                _ = v28
                                            case "a":
                                                v29 := v4[(^(v1))]
                                                _ = v29
                                            }
                                            v26 = 255
                                            _ = v26
                                        }
                                        _ = fi24
                                        _ = v25
                                    }
                                }
                                var v30 int32 = 0
                                v31 := GenS2{F0: 242, F1: (v3 & (v3 / ((42) | 1))), F2: 32767}
                                _ = v30
                                _ = v31
                            } else {
                                v32 := GenS1{F0: 168, F1: "foo bar"}
                                switch (fi22 * 337) {
                                case -650:
                                    if ((false || (((805 < 9223372036854775807) || false) || true)) || (false && true)) {
                                        v33 := GenS1{F0: 29, F1: "\n\t"}
                                        var v34 int64 = 386
                                        _ = v33
                                        _ = v34
                                    }
                                }
                                _ = v32
                            }
                            _ = v23
                        }
                        var v35 int64 = 1
                        if (v1 == v1) {
                            v3 -= (v3 >> ((429) & 63))
                            var v36 int64 = 1
                            _ = v36
                        }
                        _ = fi22
                        _ = v35
                    }
                    fmt.Println(v1, v3)
                }
            }
        default:
            var v37 int = 128
            _ = v37
        }
    }
    switch 4294967295 {
    case 0:
        if false {
            for fi38 := 0; fi38 < 11; fi38++ {
                v3 += 18446744073709551615
                var v39 uint16 = 255
                _ = fi38
                _ = v39
            }
            if true {
                var v40 int8 = 1
                _ = v40
            } else {
                v3 -= 2
                fmt.Println(v3, v3)
            }
        }
    case -2147483648:
        v1 -= (v1 * v1)
    case 2147483647:
        _ = len(v5)
    default:
        if false {
            v41 := GenS5{F0: 65535, F1: 0}
            for fi42 := 0; fi42 < 2; fi42++ {
                switch fi42 {
                case -9223372036854775808:
                    var v43 int16 = int16(fi42)
                    _ = v43
                case 32767:
                    for fi44 := 0; fi44 < 3; fi44++ {
                        var v45 int64 = 1
                        v46 := GenS6{F0: v1, F1: 0}
                        switch "" {
                        case "a":
                            v2[("foo bar" + "")] = "a"
                        case "\n\t":
                            _ = len(v5)
                            v47 := GenS5{F0: 18446744073709551615, F1: 128}
                            var v48 uint16 = 42
                            _ = v47
                            _ = v48
                        default:
                            fmt.Println(v3)
                        }
                        _ = fi44
                        _ = v45
                        _ = v46
                    }
                case 2147483647:
                    var v49 bool = (0 > 0)
                    _ = v49
                }
                fmt.Println(fi42, fi42, fi42)
                v50 := GenS2{F0: 66, F1: (v3 & 42), F2: -414}
                _ = fi42
                _ = v50
            }
            if (1 == 42) {
                _ = len(v5)
                v51 := GenS7{F0: -1, F1: 32767}
                v3 = 18446744073709551615
                _ = v51
            }
            _ = v41
        }
    }
    fmt.Println(v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
