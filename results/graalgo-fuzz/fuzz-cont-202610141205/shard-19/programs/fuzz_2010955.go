package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int64 }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16; F1 int16; F2 int8 }
type GenS5 struct { F0 int8; F1 int64; F2 uint64 }
type GenS6 struct { F0 uint16; F1 uint32 }

func main() {
    var v1 int64 = 1
    v2 := []int{-187, -9223372036854775808, 2147483647}
    v3 := make(map[uint32]int)
    var v4 string = ("hello" + "\n\t")
    v5 := make(map[string]string)
    v6 := []string{v4, v4, (("" + v4) + (v4 + v4)), v4}
    switch 1 {
    case 210:
        var v7 int = 9223372036854775807
        _ = v7
    case 32767:
        if true {
            v8 := GenS1{F0: -102, F1: 255, F2: (v1 ^ v1)}
            v4 = (v4 + v4)
            _ = v8
        }
    case -901:
        switch v4 {
        case "a":
            switch ((v4 + v4) + "hello") {
            case "a":
                switch -1 {
                case 128:
                    v9 := GenS2{F0: 2, F1: (!true)}
                    _ = v9
                default:
                    v10 := GenS3{F0: 215}
                    _ = v10
                }
            default:
                var v11 float32 = (float32(3.14159) - ((0.1 / 1e10) * float32((1e10 / 1e10))))
                _ = v11
            }
        case "hello":
            var v12 uint16 = 2
            _ = v12
        default:
            for fi13 := 0; fi13 < 5; fi13++ {
                if false {
                    v4 = ""
                }
                _ = fi13
            }
        }
    }
    switch "foo bar" {
    case "a":
        v1 = v1
    case "":
        for fi14 := 0; fi14 < 13; fi14++ {
            for fi15 := 0; fi15 < 13; fi15++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi15
            }
            _ = fi14
        }
    default:
        v4 = "hello"
    }
    if false {
        for fi16 := 0; fi16 < 11; fi16++ {
            v4 = "hello"
            v2 = append(v2, -1)
            _ = fi16
        }
        for fi17 := 0; fi17 < 4; fi17++ {
            var v18 int32 = 0
            _ = fi17
            _ = v18
        }
    }
    v19 := GenS1{F0: 0, F1: 104, F2: 42}
    switch 300 {
    case -9223372036854775808:
        v6 = append(v6, v4)
    case 309:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    case 127:
        v20 := GenS3{F0: 0}
        _ = v20
    default:
        v21 := GenS4{F0: 0, F1: 42, F2: -37}
        _ = v21
    }
    for fi22 := 0; fi22 < 12; fi22++ {
        if (-2.25 >= 0.0) {
            for fi23 := 0; fi23 < 11; fi23++ {
                v4 = v4
                _ = fi23
            }
            switch fi22 {
            case 964:
                v1 += v1
            default:
                v24 := GenS4{F0: -32768, F1: 42, F2: 127}
                _ = v24
            }
            var v25 uint = 65535
            _ = v25
        }
        for fi26 := 0; fi26 < 5; fi26++ {
            if (true || true) {
                fmt.Println(fi22, v4)
            } else {
                v1 -= v1
                var v27 string = (("foo bar" + v4) + v4)
                _ = v27
            }
            var v28 int16 = 32767
            fmt.Println(fi26)
            if (0 <= 651) {
                if true {
                    v29 := v3[1]
                    if ((v28 / ((534) | 1)) != int16(v1)) {
                        if (65535 != 301) {
                            var v30 bool = false
                            if (true || (true || v30)) {
                                var v31 uint32 = 65535
                                fmt.Println(v30, v29, v29)
                                v32 := v3[(v31 % ((v31) | 1))]
                                _ = v31
                                _ = v32
                            }
                            _ = v30
                        }
                    } else {
                        v33 := v5[((((((v4 + "foo bar") + v4) + (v4 + v4)) + v4) + ("\n\t" + ("hello" + ((v4 + v4) + (("\n\t" + v4) + v4))))) + "hello")]
                        var v34 float32 = ((1.5 * (0.0 * 0.1)) / (1e10 - -2.25))
                        v35 := GenS5{F0: 17, F1: v1, F2: 18446744073709551615}
                        _ = v33
                        _ = v34
                        _ = v35
                    }
                    _ = v29
                }
            } else {
                v36 := GenS4{F0: (^(v28)), F1: v28, F2: -1}
                switch "hello" {
                case "":
                    delete(v5, v4)
                case "a":
                    v37 := GenS4{F0: -32768, F1: (v28 >> ((1) & 63)), F2: -128}
                    _ = v37
                }
                v1 -= 42
                _ = v36
            }
            v4 = v4
            _ = fi26
            _ = v28
        }
        v1 = v1
        _ = fi22
    }
    v4 = (v4 + ("foo bar" + (((v4 + "a") + "") + v4)))
    var v38 int64 = v1
    v39 := GenS6{F0: 0, F1: 4294967295}
    v40 := v5[v4]
    var v41 uint64 = 2
    var v42 uint32 = 4294967295
    if ((v42 - v42) < (^(v42))) {
        if ((v42 & v42) < v42) {
            v40 = (v4 + (v4 + (v4 + ((((v40 + (v4 + "")) + (v4 + v40)) + v40) + v4))))
            var v43 uint16 = 255
            switch "hello" {
            case "":
                v44 := GenS3{F0: 1}
                _ = v44
            default:
                v45 := GenS2{F0: (^(v41)), F1: (v43 > 42)}
                _ = v45
            }
            _ = v43
        }
        switch v4 {
        case "a":
            if (42 != 57) {
                if (!false) {
                    for fi46 := 0; fi46 < 5; fi46++ {
                        v47 := GenS4{F0: -992, F1: 1000, F2: 0}
                        v48 := v3[(v42 % ((v42) | 1))]
                        var v49 uint32 = v42
                        _ = fi46
                        _ = v47
                        _ = v48
                        _ = v49
                    }
                } else {
                    var v50 int64 = (v1 - (v38 / ((1) | 1)))
                    v51 := v3[(v42 | 512)]
                    _ = v50
                    _ = v51
                }
                var v52 bool = ((^(v41)) >= (v41 * (v41 + (v41 & (v41 - v41)))))
                delete(v3, 4294967295)
                _ = v52
            }
        case "\n\t":
            if false {
                v53 := v3[4294967295]
                _ = v53
            }
        case "hello":
            v40 = v40
        default:
            v4 = ((v40 + ("foo bar" + ("\n\t" + "hello"))) + "")
        }
    } else {
        var v54 int = -305
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
    _ = v6
}
