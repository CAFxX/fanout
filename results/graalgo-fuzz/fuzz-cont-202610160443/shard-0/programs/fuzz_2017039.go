package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 uint32; F1 uint }
type GenS5 struct { F0 uint; F1 uint; F2 int16 }

func main() {
    var v1 uint8 = 42
    var v2 int32 = 863
    var v3 uint16 = 1
    v4 := make(map[string]int)
    var v5 string = "foo bar"
    var v6 int32 = (v2 << ((557) & 63))
    var v7 int8 = 0
    v5 = v5
    if (1 < 42) {
        if ((v7 % (((v7 << ((18446744073709551615) & 63))) | 1)) > (v7 ^ (v7 % ((v7) | 1)))) {
            for fi8 := 0; fi8 < 5; fi8++ {
                var v9 int64 = 749
                v10 := []int{fi8, (fi8 % (((fi8 / ((fi8) | 1))) | 1)), (fi8 >> ((1) & 63))}
                switch ("foo bar" + v5) {
                case "":
                    v11 := GenS1{F0: v1, F1: 18446744073709551615, F2: 2147483647}
                    _ = v11
                case "hello":
                    if (false && false) {
                        v12 := GenS1{F0: (^(v1)), F1: 1, F2: fi8}
                        _ = v12
                    }
                case "\n\t":
                    var v13 bool = (-2.25 == (0.0 / (0.0 + 3.14159)))
                    _ = v13
                default:
                    var v14 int16 = 32767
                    _ = v14
                }
                _ = fi8
                _ = v10
                _ = v9
            }
            var v15 int16 = 599
            _ = v15
        } else {
            if (!(!(!false))) {
                switch v5 {
                case "\n\t":
                    v6 += (-(v2))
                case "a":
                    if false {
                        switch -352 {
                        case 9223372036854775807:
                            if (!(872 < 2147483647)) {
                                var v16 int16 = 0
                                _ = v16
                            }
                        case -9223372036854775808:
                            v6 += -1
                        case -638:
                            var v17 int32 = v6
                            _ = v17
                        }
                    }
                case "foo bar":
                    for fi18 := 0; fi18 < 1; fi18++ {
                        v6 = (v2 >> ((uint(v3)) & 63))
                        var v19 uint8 = v1
                        for fi20 := 0; fi20 < 4; fi20++ {
                            v21 := GenS2{F0: uint8(v3)}
                            v22 := v4[v5]
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                        _ = fi18
                        _ = v19
                    }
                default:
                    var v23 string = ("" + "a")
                    _ = v23
                }
            } else {
                v24 := GenS2{F0: (v1 << ((65535) & 63))}
                switch v5 {
                case "foo bar":
                    var v25 int32 = 1
                    _ = v25
                case "":
                    if true {
                        switch v5 {
                        case "\n\t":
                            v26 := []int{}
                            _ = v26
                        case "":
                            fmt.Println(v6, v5, v6)
                        case "hello":
                            v2 = v6
                        }
                        v1 += (^(v1))
                        v27 := GenS3{F0: v5, F1: "a"}
                        _ = v27
                    } else {
                        var v28 int32 = v6
                        _ = v28
                    }
                case "a":
                    for fi29 := 0; fi29 < 6; fi29++ {
                        var v30 uint64 = uint64(v3)
                        if (true || (!(false || false))) {
                            if false {
                                _ = len(v4)
                                var v31 int32 = (v6 + v2)
                                _ = v31
                            }
                            switch "" {
                            case "hello":
                                if true {
                                    for fi32 := 0; fi32 < 1; fi32++ {
                                        v5 = ("hello" + "hello")
                                        var v33 uint32 = 2
                                        if (v7 == v7) {
                                            var v34 bool = false
                                            _ = v34
                                        }
                                        _ = fi32
                                        _ = v33
                                    }
                                }
                            case "\n\t":
                                v6 = (v6 / (((v6 / ((v6) | 1))) | 1))
                            }
                        }
                        _ = fi29
                        _ = v30
                    }
                default:
                    var v35 uint8 = v1
                    _ = v35
                }
                v36 := GenS3{F0: v5, F1: v5}
                _ = v24
                _ = v36
            }
            for fi37 := 0; fi37 < 1; fi37++ {
                v38 := []int{(fi37 + fi37), fi37, (fi37 * (-(fi37)))}
                var v39 int64 = -1
                _ = fi37
                _ = v38
                _ = v39
            }
            var v40 string = "foo bar"
            _ = v40
        }
        var v41 uint8 = 14
        v6 -= v2
        _ = v41
    }
    var v42 int64 = 42
    v43 := GenS4{F0: 42, F1: 18446744073709551615}
    v44 := GenS5{F0: uint(v1), F1: 2, F2: -32768}
    var v45 uint64 = 1
    var v46 float32 = (((0.1 - (-2.25 * 0.1)) + float32(0.0)) + float32((float64(((3.14159 + (0.0 * ((float32((1.5 * float64(-2.25))) * (0.1 + (float32(1e10) - -2.25))) - -2.25))) / (1e10 / 1e10))) * (0.1 - 0.0))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
    _ = v7
}
