package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8; F1 int32; F2 uint }
type GenS7 struct { F0 string; F1 uint64 }
func genFn0(p0 int32, p1 uint64, p2 int8) int {
    var v6 uint8 = 2
    var v7 uint32 = 4294967295
    v8 := GenS1{F0: (true || (((^(v6)) == 1) || false))}
    for fi9 := 0; fi9 < 3; fi9++ {
        v10 := GenS1{F0: (((v7 ^ 1) > (v7 | v7)) || true)}
        v6 = 31
        p2 = 35
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return -464
}

func main() {
    v1 := []string{"foo bar"}
    var v2 uint = 42
    var v3 int = 1
    var v4 string = "hello"
    v5 := make(map[uint32]string)
    v3 += v3
    switch "foo bar" {
    case "\n\t":
        v11 := GenS1{F0: (32767 == 255)}
        _ = v11
    case "":
        v12 := GenS1{F0: (2147483647 > -682)}
        _ = v12
    case "hello":
        if (false || (4294967295 <= 0)) {
            var v13 float32 = (0.1 - float32(0.0))
            v13 = 1.5
            _ = v13
        }
    default:
        v14 := GenS1{F0: (-1 <= 274)}
        _ = v14
    }
    v15 := GenS2{F0: false, F1: (!false)}
    _ = len(v5)
    for fi16 := 0; fi16 < 4; fi16++ {
        v2 = v2
        var v17 string = (v4 + "a")
        _ = fi16
        _ = v17
    }
    if true {
        v3 = v3
    } else {
        switch "hello" {
        case "hello":
            for fi18 := 0; fi18 < 19; fi18++ {
                if (127 > 42) {
                    var v19 float32 = (float32(1e10) / 1.5)
                    _ = v19
                } else {
                    for fi20 := 0; fi20 < 19; fi20++ {
                        if false {
                            switch (-(fi18)) {
                            case -1:
                                _ = len(v1)
                            case 1:
                                var v21 float64 = -2.25
                                _ = v21
                            }
                            var v22 float32 = 0.1
                            _ = v22
                        }
                        v23 := GenS1{F0: true}
                        _ = fi20
                        _ = v23
                    }
                }
                _ = fi18
            }
        case "a":
            switch "\n\t" {
            case "hello":
                if true {
                    switch v3 {
                    case -107:
                        v5[897] = (v4 + (v4 + v4))
                    case 996:
                        var v24 uint64 = 1
                        _ = v24
                    case 1:
                        switch 202 {
                        case 9223372036854775807:
                            switch v4 {
                            case "a":
                                fmt.Println(v2)
                            case "hello":
                                switch (v3 / (((v3 >> ((v2) & 63))) | 1)) {
                                case 32768:
                                    v25 := GenS3{F0: 1}
                                    _ = v25
                                default:
                                    v26 := GenS4{F0: 2147483648}
                                    _ = v26
                                }
                            case "\n\t":
                                v27 := GenS5{F0: -1}
                                _ = v27
                            default:
                                for fi28 := 0; fi28 < 2; fi28++ {
                                    _ = len(v5)
                                    v3 -= genFn0(575, 18446744073709551615, 0)
                                    _ = fi28
                                }
                            }
                        case 592:
                            for fi29 := 0; fi29 < 6; fi29++ {
                                _ = len(v1)
                                v30 := GenS1{F0: false}
                                v31 := GenS3{F0: 127}
                                _ = fi29
                                _ = v30
                                _ = v31
                            }
                        case 859:
                            v32 := GenS4{F0: 42}
                            _ = v32
                        default:
                            v2 += (v2 >> (((^(v2))) & 63))
                        }
                    }
                }
            case "\n\t":
                v3 = (v3 - 32767)
            default:
                for fi33 := 0; fi33 < 16; fi33++ {
                    if true {
                        v34 := GenS3{F0: 42}
                        var v35 float64 = 1e10
                        _ = v34
                        _ = v35
                    } else {
                        if (0.0 <= (-2.25 + 0.0)) {
                            for fi36 := 0; fi36 < 5; fi36++ {
                                v37 := GenS3{F0: -63}
                                switch ("\n\t" + (v4 + ((v4 + v4) + ("hello" + (v4 + (v4 + v4)))))) {
                                case "":
                                    fmt.Println(v4)
                                case "foo bar":
                                    if false {
                                        v3 += (v3 | 747)
                                        if (false && (1000 < 32767)) {
                                            var v38 int64 = 42
                                            switch 4294967295 {
                                            case 216:
                                                v3 -= genFn0(42, 2, -128)
                                            case -9223372036854775808:
                                                v3 = 32768
                                            case 863:
                                                v39 := v5[1]
                                                _ = v39
                                            }
                                            _ = v38
                                        }
                                        switch v3 {
                                        case -1:
                                            v2 = (v2 * (v2 + (v2 - (v2 & v2))))
                                        default:
                                            v4 = ""
                                        }
                                    } else {
                                        v40 := GenS6{F0: -99, F1: -2147483648, F2: v2}
                                        v2 -= 4294967295
                                        v4 = v4
                                        _ = v40
                                    }
                                case "a":
                                    var v41 int = 926
                                    _ = v41
                                default:
                                    fmt.Println(v4)
                                }
                                _ = len(v1)
                                _ = fi36
                                _ = v37
                            }
                        } else {
                            v3 = fi33
                            v42 := v5[983]
                            _ = v42
                        }
                        switch 2147483647 {
                        case 128:
                            if (!false) {
                                for fi43 := 0; fi43 < 7; fi43++ {
                                    v3 += (-(fi43))
                                    v44 := GenS7{F0: v4, F1: 0}
                                    _ = fi43
                                    _ = v44
                                }
                            } else {
                                switch (fi33 - (v3 ^ (^(fi33)))) {
                                case 0:
                                    var v45 uint32 = 2147483647
                                    _ = v45
                                case 9223372036854775807:
                                    v3 += (fi33 | (fi33 ^ 4294967295))
                                case 1:
                                    var v46 uint32 = 2
                                    _ = v46
                                default:
                                    v5[65535] = (v4 + ((v4 + "") + v4))
                                }
                            }
                        case 42:
                            var v47 int64 = -1
                            _ = v47
                        case -848:
                            _ = len(v1)
                        default:
                            v48 := GenS3{F0: 0}
                            _ = v48
                        }
                        var v49 int64 = -9223372036854775808
                        _ = v49
                    }
                    _ = fi33
                }
            }
        default:
            fmt.Println(v2)
        }
        var v50 int16 = -1
        switch ("foo bar" + "\n\t") {
        case "\n\t":
            switch (v4 + (v4 + v4)) {
            case "a":
                var v51 int64 = -515
                _ = v51
            case "hello":
                var v52 uint = v2
                _ = v52
            }
        case "foo bar":
            v53 := GenS7{F0: "foo bar", F1: 2}
            _ = v53
        }
        _ = v50
    }
    v4 = (((v4 + v4) + ("hello" + "a")) + v4)
    switch (v3 + v3) {
    case -15:
        for fi54 := 0; fi54 < 19; fi54++ {
            v3 += fi54
            var v55 string = v4
            switch v4 {
            case "hello":
                var v56 int8 = 1
                _ = v56
            default:
                v4 = ((v4 + v4) + v4)
            }
            _ = fi54
            _ = v55
        }
    case -821:
        var v57 int = v3
        _ = v57
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
