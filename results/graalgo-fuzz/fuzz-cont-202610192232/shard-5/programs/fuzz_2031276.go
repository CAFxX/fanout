package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 uint8; F1 uint16; F2 uint }
type GenS4 struct { F0 int8; F1 bool; F2 uint }
type GenS5 struct { F0 string; F1 uint64; F2 uint8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int16 }

func main() {
    v1 := []uint32{0, 2}
    var v2 int32 = -192
    var v3 uint32 = 2147483647
    v4 := make(map[string]string)
    fmt.Println(v3, v2)
    var v5 int8 = -95
    switch "hello" {
    case "hello":
        fmt.Println(v3, v5)
    default:
        _ = len(v4)
    }
    v5 -= -1
    var v6 int16 = -109
    v2 = v2
    if (!((^(v5)) <= (-(v5)))) {
        var v7 int64 = -1
        for fi8 := 0; fi8 < 17; fi8++ {
            v2 -= v2
            var v9 uint64 = 0
            _ = fi8
            _ = v9
        }
        v5 = (-(v5))
        _ = v7
    }
    v3 += 86
    v5 = (v5 + v5)
    fmt.Println(v3)
    var v10 int8 = (-(v5))
    if true {
        v3 = (^(v3))
        switch (("a" + ("" + "foo bar")) + "\n\t") {
        case "foo bar":
            v11 := v4[("foo bar" + "a")]
            _ = v11
        case "hello":
            for fi12 := 0; fi12 < 1; fi12++ {
                for fi13 := 0; fi13 < 17; fi13++ {
                    v14 := GenS1{F0: v3, F1: v3, F2: 0}
                    v5 = 29
                    _ = fi13
                    _ = v14
                }
                switch 2147483647 {
                case 9223372036854775807:
                    if true {
                        v4["\n\t"] = "a"
                    } else {
                        switch (fi12 & -9223372036854775808) {
                        case -104:
                            for fi15 := 0; fi15 < 10; fi15++ {
                                v3 = 462
                                var v16 uint8 = 42
                                if true {
                                    if true {
                                        if ((^(v2)) <= 1) {
                                            v2 -= (v2 ^ 1)
                                        } else {
                                            if true {
                                                if len(v1) > 0 {
                                                    v1[0] = (v3 ^ 2)
                                                }
                                                v17 := GenS1{F0: 1, F1: (v3 << ((42) & 63)), F2: 42}
                                                var v18 uint16 = 1
                                                v18 += (v18 & (^(v18)))
                                                _ = v17
                                                _ = v18
                                            }
                                        }
                                        var v19 uint16 = 65535
                                        var v20 int64 = 797
                                        v5 -= (^(v10))
                                        _ = v19
                                        _ = v20
                                    }
                                    v21 := GenS1{F0: 682, F1: v3, F2: 358}
                                    _ = v21
                                }
                                v16 += (^(v16))
                                fmt.Println(v5, v2)
                                switch (("hello" + "hello") + (("foo bar" + ("\n\t" + "foo bar")) + "hello")) {
                                case "a":
                                    var v22 string = ("hello" + "a")
                                    _ = v22
                                case "foo bar":
                                    switch "\n\t" {
                                    case "a":
                                        if (true && true) {
                                            v23 := GenS2{F0: (-(v10)), F1: 4294967295}
                                            _ = len(v4)
                                            _ = v23
                                        } else {
                                            if (2 < 108) {
                                                v10 += (v10 ^ -128)
                                                if true {
                                                    v6 -= (v6 - (v6 | 1000))
                                                    v24 := GenS1{F0: (v3 * (^(v3))), F1: 2147483647, F2: 870}
                                                    var v25 int16 = (v6 + v6)
                                                    _ = v24
                                                    _ = v25
                                                }
                                                var v26 uint16 = 32767
                                                _ = v26
                                            }
                                            var v27 float64 = 0.1
                                            var v28 int64 = -511
                                            var v29 int16 = int16(v5)
                                            _ = v27
                                            _ = v28
                                            _ = v29
                                        }
                                        switch fi15 {
                                        case -1:
                                            var v30 int32 = (v2 / ((v2) | 1))
                                            _ = v30
                                        case 412:
                                            var v31 uint64 = 42
                                            _ = v31
                                        case 297:
                                            v3 = (v3 + 42)
                                        }
                                    case "hello":
                                        var v32 uint8 = (v16 / (((v16 + (v16 + v16))) | 1))
                                        v16 += 74
                                        var v33 uint64 = 42
                                        _ = len(v4)
                                        _ = v32
                                        _ = v33
                                    case "":
                                        var v34 float64 = -2.25
                                        _ = v34
                                    }
                                default:
                                    var v35 bool = ((0 != 65535) && true)
                                    _ = v35
                                }
                                var v36 float32 = 3.14159
                                _ = fi15
                                _ = v16
                                _ = v36
                            }
                        case -1:
                            v4["\n\t"] = ("a" + "hello")
                        case 9223372036854775807:
                            v5 -= (v5 * -1)
                        default:
                            for fi37 := 0; fi37 < 15; fi37++ {
                                var v38 uint16 = 1
                                _ = fi37
                                _ = v38
                            }
                        }
                        for fi39 := 0; fi39 < 8; fi39++ {
                            v3 = v3
                            _ = fi39
                        }
                    }
                case 0:
                    fmt.Println(v10, v6, v2)
                default:
                    fmt.Println(v2, v10, fi12)
                }
                v40 := GenS2{F0: v5, F1: 42}
                _ = fi12
                _ = v40
            }
        case "a":
            v41 := GenS3{F0: 127, F1: 42, F2: 443}
            _ = v41
        default:
            if false {
                fmt.Println(v3, v2)
                v3 = (v3 - (v3 - 367))
            }
        }
    } else {
        var v42 uint8 = 0
        var v43 int16 = v6
        for fi44 := 0; fi44 < 5; fi44++ {
            var v45 float32 = (float32(0.1) * (1.5 + -2.25))
            _ = len(v1)
            v10 += v5
            _ = fi44
            _ = v45
        }
        _ = v42
        _ = v43
    }
    fmt.Println(v10)
    if (int8(v2) == (v10 % (((-(v10))) | 1))) {
        switch (("\n\t" + ("" + "")) + (("" + "a") + "a")) {
        case "a":
            for fi46 := 0; fi46 < 2; fi46++ {
                if true {
                    var v47 float64 = (-2.25 * (-2.25 - 0.0))
                    _ = v47
                }
                v48 := v4["hello"]
                fmt.Println(v6, v10)
                _ = fi46
                _ = v48
            }
        case "":
            switch 0 {
            case -783:
                v5 = (v5 + v10)
            case 9223372036854775807:
                var v49 bool = ((v10 * v5) > v10)
                _ = v49
            case 1:
                v50 := GenS4{F0: 127, F1: (0 < 747), F2: 2}
                _ = v50
            }
        case "hello":
            switch 510 {
            case 0:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case 42:
                delete(v4, "foo bar")
            case 1:
                var v51 uint32 = v3
                _ = v51
            }
        default:
            if false {
                for fi52 := 0; fi52 < 11; fi52++ {
                    v10 = (v10 << ((42) & 63))
                    for fi53 := 0; fi53 < 14; fi53++ {
                        v10 = v5
                        _ = fi53
                    }
                    var v54 int8 = (v5 / (((v5 >> ((1) & 63))) | 1))
                    _ = fi52
                    _ = v54
                }
                v2 += v2
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
    } else {
        for fi55 := 0; fi55 < 19; fi55++ {
            var v56 int64 = 9223372036854775807
            fmt.Println(v2)
            _ = fi55
            _ = v56
        }
    }
    switch -1 {
    case 4294967295:
        var v57 uint8 = 205
        _ = v57
    case 0:
        var v58 uint16 = 65535
        _ = v58
    case 1:
        for fi59 := 0; fi59 < 11; fi59++ {
            v60 := GenS5{F0: "\n\t", F1: 42, F2: 128}
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi59
            _ = v60
        }
    }
    var v61 uint32 = v3
    v62 := GenS6{F0: 255}
    if (v2 < 338) {
        if true {
            v3 = 0
        }
        fmt.Println(v5)
        fmt.Println(v10, v10, v6)
    } else {
        var v63 float32 = (3.14159 / 1.5)
        v64 := GenS7{F0: (-(v6))}
        _ = v63
        _ = v64
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v62
}
