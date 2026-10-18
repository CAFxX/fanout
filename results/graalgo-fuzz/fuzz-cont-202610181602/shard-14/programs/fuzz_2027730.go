package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int8 }
type GenS2 struct { F0 uint8; F1 uint16; F2 int8 }
type GenS3 struct { F0 bool; F1 int64; F2 int }
func genFn0(p0 uint8, p1 uint32, p2 int16) int16 {
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 int32 = 233
        var v7 int = (fi5 / (((fi5 << ((65535) & 63))) | 1))
        for fi8 := 0; fi8 < 16; fi8++ {
            if true {
                v9 := GenS1{F0: p1, F1: (p2 - (^(p2))), F2: -24}
                if false {
                    v10 := []int{(^(fi8)), (fi5 - (fi8 << ((255) & 63))), int(p2)}
                    if len(v10) > 0 {
                        v10[0] = (fi8 << ((929) & 63))
                    }
                    fmt.Println(fi5)
                    p1 = (p1 & 42)
                    _ = len(v10)
                    if ((p0 > p0) && true) {
                        var v11 uint64 = 42
                        switch (("" + "") + "\n\t") {
                        case "a":
                            p0 = 14
                        case "":
                            v7 -= (-(v7))
                        case "foo bar":
                            var v12 uint8 = p0
                            _ = v12
                        }
                        fmt.Println(p2)
                        _ = v11
                    }
                    _ = v10
                }
                _ = v9
            }
            p1 += (p1 - 990)
            v13 := GenS1{F0: (p1 & (p1 - (p1 - p1))), F1: p2, F2: -1}
            _ = fi8
            _ = v13
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    v14 := []int{int(p0), -192, -1}
    if (!((((1e10 + -2.25) < -2.25) && ("hello" <= ("" + "foo bar"))) && (!false))) {
        for fi15 := 0; fi15 < 12; fi15++ {
            if false {
                switch (("" + "") + "\n\t") {
                case "foo bar":
                    v16 := GenS2{F0: (p0 ^ 0), F1: 255, F2: -82}
                    _ = v16
                case "":
                    var v17 string = "\n\t"
                    _ = v17
                case "hello":
                    v18 := GenS1{F0: (^(p1)), F1: -32768, F2: 127}
                    _ = v18
                default:
                    v19 := GenS2{F0: p0, F1: 0, F2: 127}
                    _ = v19
                }
                var v20 int16 = (p2 >> ((65535) & 63))
                _ = v20
            }
            fmt.Println(p2, p1, fi15)
            _ = fi15
        }
        if len(v14) > 0 {
            v14[0] = 42
        }
        p0 = (p0 << ((42) & 63))
    }
    p0 -= (p0 | (p0 * 128))
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return 42
}
func genFn1(p0 uint16) string {
    if (p0 > (p0 ^ p0)) {
        var v21 string = "hello"
        _ = v21
    } else {
        p0 += p0
    }
    v22 := []int{}
    var v23 float64 = -2.25
    _ = p0
    _ = v22
    _ = v23
    return "a"
}

func main() {
    var v1 int32 = 1
    v2 := make(map[int]string)
    v3 := []uint32{}
    v4 := make(map[string]bool)
    switch (((("" + "foo bar") + "hello") + "hello") + "foo bar") {
    case "foo bar":
        v1 = v1
    default:
        v1 = v1
    }
    v3 = append(v3, 2147483648)
    var v24 int64 = 42
    for fi25 := 0; fi25 < 19; fi25++ {
        if true {
            if (true && ((^(v24)) != -143)) {
                var v26 float64 = 3.14159
                if false {
                    if (42 > (v24 * 42)) {
                        v27 := GenS2{F0: 0, F1: 895, F2: int8(v1)}
                        v26 += 3.14159
                        v1 = (v1 / ((v1) | 1))
                        _ = v27
                    } else {
                        v4["a"] = (false || (!false))
                        switch fi25 {
                        case -9223372036854775808:
                            var v28 float64 = float64(0.0)
                            _ = v28
                        case 4294967295:
                            v29 := GenS2{F0: 169, F1: 32767, F2: 42}
                            _ = v29
                        default:
                            v30 := GenS2{F0: 132, F1: 972, F2: 2}
                            _ = v30
                        }
                        var v31 int8 = 0
                        _ = v31
                    }
                } else {
                    if (true || false) {
                        delete(v2, (fi25 + (fi25 % (((fi25 & (-(fi25)))) | 1))))
                        if (0 <= 284) {
                            v2[(^(fi25))] = ("foo bar" + "a")
                            var v32 uint64 = 2
                            _ = v32
                        }
                        var v33 string = "\n\t"
                        _ = v33
                    } else {
                        v34 := v4[("hello" + "hello")]
                        if (65535 < 2) {
                            if len(v3) > 0 {
                                v3[0] = 2147483647
                            }
                            var v35 int16 = 0
                            _ = v35
                        } else {
                            var v36 uint8 = 247
                            _ = v36
                        }
                        _ = v34
                    }
                }
                _ = v26
            }
            var v37 float64 = float64(1e10)
            _ = v37
        } else {
            var v38 int16 = 32767
            var v39 string = "foo bar"
            var v40 uint16 = 42
            _ = v38
            _ = v39
            _ = v40
        }
        v24 += (v24 / ((-9223372036854775808) | 1))
        var v41 uint = 0
        _ = fi25
        _ = v41
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v42 uint16 = 1
    fmt.Println(v24, v42)
    switch 615 {
    case 32768:
        fmt.Println(v42, v1)
    case -2147483648:
        if false {
            var v43 int8 = 0
            for fi44 := 0; fi44 < 20; fi44++ {
                v1 = (v1 - -555)
                var v45 int8 = v43
                _ = fi44
                _ = v45
            }
            _ = v43
        } else {
            var v46 string = ""
            if (!false) {
                v42 = v42
                switch (v46 + "foo bar") {
                case "":
                    v47 := GenS1{F0: 161, F1: 1, F2: 1}
                    _ = v47
                }
            }
            v1 = (v1 & (v1 | v1))
            _ = v46
        }
    default:
        if true {
            if false {
                for fi48 := 0; fi48 < 15; fi48++ {
                    v42 -= (v42 * 528)
                    _ = fi48
                }
                for fi49 := 0; fi49 < 4; fi49++ {
                    v42 = v42
                    v1 += v1
                    _ = fi49
                }
                if len(v3) > 0 {
                    v3[0] = 2147483647
                }
            }
            var v50 string = "a"
            v51 := GenS1{F0: 65535, F1: 32767, F2: -128}
            _ = v50
            _ = v51
        } else {
            for fi52 := 0; fi52 < 17; fi52++ {
                if (float64(3.14159) == 1e10) {
                    v42 += v42
                    var v53 float32 = (0.0 - 3.14159)
                    _ = v53
                } else {
                    v54 := GenS3{F0: (-2.25 > 0.1), F1: v24, F2: (fi52 - fi52)}
                    _ = v54
                }
                var v55 int = fi52
                _ = fi52
                _ = v55
            }
            var v56 float32 = 3.14159
            _ = v56
        }
    }
    if ((1e10 * 0.1) > 1.5) {
        if (genFn0(255, 0, 32767) < 32767) {
            if false {
                v57 := GenS3{F0: (!true), F1: (v24 & (v24 % ((v24) | 1))), F2: -9223372036854775808}
                switch "hello" {
                case "foo bar":
                    v1 += (v1 ^ (^(v1)))
                case "hello":
                    var v58 int = 2147483647
                    _ = v58
                }
                for fi59 := 0; fi59 < 3; fi59++ {
                    var v60 int64 = (-(v24))
                    _ = fi59
                    _ = v60
                }
                _ = v57
            }
            fmt.Println(v1, v24)
            switch "a" {
            case "\n\t":
                for fi61 := 0; fi61 < 7; fi61++ {
                    switch "\n\t" {
                    case "hello":
                        v42 += 255
                    case "foo bar":
                        var v62 float64 = 0.0
                        _ = v62
                    case "a":
                        v1 = v1
                    default:
                        v1 -= 1
                    }
                    var v63 int16 = 251
                    _ = fi61
                    _ = v63
                }
            case "a":
                var v64 float64 = 0.1
                _ = v64
            }
        }
    } else {
        fmt.Println(v24)
        v65 := GenS1{F0: 169, F1: genFn0(1, 296, 1), F2: -48}
        _ = v65
    }
    _ = len(v4)
    _ = len(v4)
    fmt.Println(v1, v24)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v42
}
