package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 uint16; F2 bool }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 uint64; F1 uint8 }
type GenS7 struct { F0 uint8; F1 uint32 }
func genFn0(p0 int64, p1 int32, p2 uint16) int {
    fmt.Println(p1)
    v4 := GenS1{F0: 42, F1: false}
    for fi5 := 0; fi5 < 2; fi5++ {
        var v6 uint16 = 211
        _ = fi5
        _ = v6
    }
    var v7 float64 = (0.0 + 1.5)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v7
    return 699
}

func main() {
    var v1 bool = false
    v2 := []uint32{1}
    v3 := []uint32{4294967295, 324, 475}
    v1 = (0.0 != ((1e10 * 0.1) - 3.14159))
    for fi8 := 0; fi8 < 1; fi8++ {
        v9 := GenS2{F0: ((("" + ("a" + ((("\n\t" + "hello") + "foo bar") + "\n\t"))) + "") + (("a" + ("hello" + "hello")) + "foo bar"))}
        _ = fi8
        _ = v9
    }
    var v10 uint = 802
    v10 += 65535
    if (0.1 == (0.0 * 0.1)) {
        v10 = 415
        var v11 int16 = 1
        _ = v11
    }
    fmt.Println(v10, v10)
    for fi12 := 0; fi12 < 12; fi12++ {
        var v13 float64 = 3.14159
        if v1 {
            var v14 int32 = 1
            v13 = v13
            _ = v14
        } else {
            v13 -= (1e10 - v13)
            if false {
                switch (fi12 >> (((v10 ^ (v10 * v10))) & 63)) {
                case 2147483647:
                    switch (-(fi12)) {
                    case 1:
                        var v15 float32 = 0.0
                        _ = v15
                    case -1:
                        if len(v2) > 0 {
                            v2[0] = 686
                        }
                    case -474:
                        var v16 uint64 = 705
                        _ = v16
                    }
                case 1:
                    var v17 int = -2147483648
                    _ = v17
                default:
                    if (667 > -1) {
                        var v18 uint8 = 60
                        _ = v18
                    }
                }
                v19 := GenS3{F0: ("hello" + ""), F1: 616, F2: v1}
                _ = v19
            }
            fmt.Println(v1)
        }
        if (0 > -622) {
            v3 = append(v3, 0)
            v20 := GenS4{F0: -2147483648, F1: 42}
            _ = v20
        } else {
            v13 = float64(float32(1.5))
            switch "hello" {
            case "\n\t":
                switch "foo bar" {
                case "hello":
                    if (593 > 697) {
                        if len(v3) > 0 {
                            v3[0] = 65535
                        }
                    } else {
                        if len(v3) > 0 {
                            v3[0] = 42
                        }
                        switch fi12 {
                        case -1:
                            for fi21 := 0; fi21 < 5; fi21++ {
                                v22 := GenS5{F0: 235, F1: 32767}
                                if v1 {
                                    switch ("hello" + "") {
                                    case "\n\t":
                                        switch ("\n\t" + "foo bar") {
                                        case "":
                                            v13 = 0.1
                                        case "\n\t":
                                            if (4294967295 != 4294967295) {
                                                v23 := GenS1{F0: 811, F1: v1}
                                                _ = v23
                                            } else {
                                                fmt.Println(fi21)
                                                v13 += v13
                                                v24 := GenS6{F0: 2, F1: 127}
                                                _ = v24
                                            }
                                        case "a":
                                            v25 := GenS6{F0: 4294967295, F1: 83}
                                            _ = v25
                                        }
                                    case "hello":
                                        if (2147483647 == 4294967295) {
                                            v13 = 1e10
                                            fmt.Println(fi12, fi21)
                                            switch (("" + ("a" + ("hello" + ("\n\t" + ("hello" + ("hello" + "foo bar")))))) + "a") {
                                            case "hello":
                                                if ((float32(((v13 * v13) * (v13 * v13))) >= -2.25) && true) {
                                                    var v26 uint64 = 4294967295
                                                    _ = v26
                                                }
                                            default:
                                                v27 := GenS2{F0: "foo bar"}
                                                _ = v27
                                            }
                                            var v28 uint16 = 946
                                            _ = v28
                                        }
                                    case "foo bar":
                                        var v29 int32 = 42
                                        _ = v29
                                    }
                                }
                                _ = fi21
                                _ = v22
                            }
                        default:
                            fmt.Println(v13)
                        }
                        var v30 int8 = 42
                        _ = v30
                    }
                case "":
                    v13 = (v13 * v13)
                }
            default:
                v31 := GenS6{F0: 18446744073709551615, F1: 142}
                _ = v31
            }
        }
        _ = fi12
        _ = v13
    }
    v10 -= (v10 >> ((v10) & 63))
    var v32 int64 = -183
    v32 = -1
    v32 -= (v32 & v32)
    v33 := GenS7{F0: 255, F1: 42}
    var v34 float32 = ((float32(0.1) + 0.0) * 0.0)
    for fi35 := 0; fi35 < 9; fi35++ {
        for fi36 := 0; fi36 < 19; fi36++ {
            var v37 uint32 = 2147483648
            v37 = 2147483647
            _ = fi36
            _ = v37
        }
        v38 := GenS1{F0: 32767, F1: false}
        _ = fi35
        _ = v38
    }
    v10 -= v10
    v34 += float32(float64(0.0))
    if true {
        if v1 {
            if v1 {
                v34 += v34
            }
        } else {
            var v39 uint32 = 109
            var v40 bool = v1
            if false {
                fmt.Println(v40)
            }
            _ = v39
            _ = v40
        }
        if v1 {
            fmt.Println(v32, v34, v1)
            v34 += v34
        }
    } else {
        for fi41 := 0; fi41 < 15; fi41++ {
            var v42 int64 = (v32 / ((v32) | 1))
            if v1 {
                v43 := GenS7{F0: 115, F1: 182}
                for fi44 := 0; fi44 < 3; fi44++ {
                    var v45 int32 = 2147483647
                    var v46 int16 = 1
                    _ = fi44
                    _ = v45
                    _ = v46
                }
                switch (("a" + "\n\t") + "foo bar") {
                case "hello":
                    v42 = 9223372036854775807
                }
                _ = v43
            }
            _ = fi41
            _ = v42
        }
        for fi47 := 0; fi47 < 17; fi47++ {
            v34 += float32(3.14159)
            _ = fi47
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
}
