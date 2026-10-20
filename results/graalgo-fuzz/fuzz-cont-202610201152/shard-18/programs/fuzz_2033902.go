package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int32; F1 uint16; F2 int }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 uint64; F1 string }

func main() {
    var v1 int = -9223372036854775808
    v2 := make(map[uint32]uint32)
    var v3 int64 = 9223372036854775807
    v4 := make(map[int]uint32)
    var v5 float64 = -2.25
    var v6 uint64 = 18446744073709551615
    fmt.Println(v1, v3)
    if (false && ((("\n\t" + ("foo bar" + "foo bar")) == "") && ((!false) || true))) {
        v7 := []int{0, (v1 >> ((42) & 63))}
        _ = v7
    } else {
        v5 -= (float64(((((-2.25 + -2.25) / 3.14159) / ((0.1 * -2.25) - (1.5 / 1.0))) + 1.5)) + ((((v5 - v5) - ((v5 + 1.5) * 1.5)) + 1e10) / v5))
        fmt.Println(v6, v1, v3)
        var v8 uint64 = (^(v6))
        _ = v8
    }
    fmt.Println(v5, v3)
    fmt.Println(v1, v6)
    switch v1 {
    case 4294967295:
        v9 := GenS1{F0: 32767, F1: 65535}
        _ = v9
    case 32767:
        v10 := []int{(v1 * v1), v1, (v1 * v1)}
        _ = v10
    case 765:
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 int64 = (v3 >> ((18446744073709551615) & 63))
            _ = fi11
            _ = v12
        }
    }
    fmt.Println(v3, v5)
    if true {
        var v13 float32 = 1.5
        _ = v13
    }
    var v14 uint32 = 42
    v6 += v6
    v15 := GenS2{F0: 147, F1: 588, F2: 4294967295}
    v16 := GenS2{F0: 42, F1: 1, F2: 4294967295}
    switch v1 {
    case 0:
        var v17 uint32 = (v14 & (v14 & v14))
        _ = v17
    case 4294967295:
        if (!((!false) && true)) {
            fmt.Println(v5)
        } else {
            if false {
                if false {
                    v3 += v3
                    var v18 int8 = int8(v3)
                    _ = v18
                }
                var v19 int32 = -433
                if false {
                    if true {
                        var v20 uint32 = (v14 & (^(v14)))
                        _ = v20
                    }
                } else {
                    v3 = -80
                    _ = len(v2)
                }
                _ = v19
            } else {
                if (true || (false && false)) {
                    switch (v1 << ((255) & 63)) {
                    case -9223372036854775808:
                        v21 := GenS1{F0: 42, F1: 720}
                        _ = v21
                    case -927:
                        v22 := GenS1{F0: 32767, F1: 541}
                        _ = v22
                    case 391:
                        if false {
                            var v23 uint8 = 127
                            var v24 float64 = 1e10
                            _ = v23
                            _ = v24
                        } else {
                            for fi25 := 0; fi25 < 1; fi25++ {
                                fmt.Println(v6, v5, v14)
                                _ = fi25
                            }
                        }
                    default:
                        if ((v6 << ((2) & 63)) >= 524) {
                            switch v1 {
                            case 189:
                                var v26 uint8 = 140
                                _ = v26
                            case -2147483648:
                                for fi27 := 0; fi27 < 14; fi27++ {
                                    for fi28 := 0; fi28 < 13; fi28++ {
                                        if true {
                                            var v29 int32 = 2147483647
                                            _ = v29
                                        }
                                        v6 = v6
                                        v30 := GenS3{F0: "hello", F1: v6}
                                        _ = fi28
                                        _ = v30
                                    }
                                    _ = fi27
                                }
                            default:
                                v14 = v14
                            }
                            delete(v4, 2147483647)
                        }
                    }
                }
                for fi31 := 0; fi31 < 9; fi31++ {
                    v32 := GenS4{F0: 315, F1: "foo bar"}
                    switch -193 {
                    case 63:
                        switch "a" {
                        case "\n\t":
                            var v33 bool = (false || true)
                            _ = v33
                        case "":
                            var v34 uint64 = 53
                            _ = v34
                        case "a":
                            v14 = (v14 ^ (v14 - 906))
                        default:
                            if (!(true && false)) {
                                var v35 string = "a"
                                v1 -= (v1 >> ((0) & 63))
                                _ = v35
                            }
                        }
                    }
                    var v36 int8 = 127
                    _ = fi31
                    _ = v32
                    _ = v36
                }
            }
            if false {
                switch v1 {
                case 0:
                    var v37 int64 = (v3 << ((855) & 63))
                    _ = v37
                }
            }
        }
    }
    if (9223372036854775807 > -89) {
        v38 := []int{(v1 ^ 1), (v1 >> ((255) & 63)), v1}
        _ = v38
    } else {
        fmt.Println(v1, v14)
        v5 = 0.0
        v1 += (v1 / ((v1) | 1))
    }
    fmt.Println(v6, v14)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
