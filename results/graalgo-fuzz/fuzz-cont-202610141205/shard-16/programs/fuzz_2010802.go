package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 int32; F2 int16 }
type GenS4 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS5 struct { F0 uint64; F1 int8; F2 uint64 }
type GenS6 struct { F0 uint8 }
func genFn0(p0 uint) int {
    fmt.Println(p0)
    v7 := []int{42, -161, 2147483647}
    _ = p0
    _ = v7
    return 509
}

func main() {
    var v1 int16 = -497
    v2 := []int{}
    v3 := make(map[uint32]int)
    var v4 int64 = 42
    var v5 int8 = 21
    v6 := []string{}
    if false {
        _ = len(v3)
        for fi8 := 0; fi8 < 13; fi8++ {
            v1 = (v1 / ((v1) | 1))
            fmt.Println(v4, v5, v1)
            if true {
                var v9 uint64 = 42
                if len(v6) > 0 {
                    v6[0] = ("" + ("" + "hello"))
                }
                _ = v9
            }
            _ = fi8
        }
    }
    fmt.Println(v1, v4, v4)
    switch 9223372036854775807 {
    case 2147483647:
        if (18446744073709551615 != 18446744073709551615) {
            if false {
                if true {
                    fmt.Println(v1, v1)
                }
                v10 := GenS1{F0: "foo bar", F1: false}
                for fi11 := 0; fi11 < 14; fi11++ {
                    switch "foo bar" {
                    case "foo bar":
                        _ = len(v6)
                    case "":
                        if true {
                            v4 += 706
                        } else {
                            var v12 int64 = -1
                            fmt.Println(v12, v1)
                            v13 := GenS1{F0: (("a" + ("hello" + "hello")) + "a"), F1: false}
                            _ = v12
                            _ = v13
                        }
                    case "a":
                        for fi14 := 0; fi14 < 16; fi14++ {
                            if len(v2) > 0 {
                                v2[0] = (fi14 / (((fi11 * 2147483647)) | 1))
                            }
                            fmt.Println(v5, v5)
                            _ = fi14
                        }
                    default:
                        v15 := GenS2{F0: 255}
                        _ = v15
                    }
                    _ = fi11
                }
                _ = v10
            }
            if (1 < 0) {
                v6 = append(v6, "hello")
                if (4294967295 > 86) {
                    for fi16 := 0; fi16 < 6; fi16++ {
                        switch (fi16 | (-(fi16))) {
                        case 127:
                            v17 := GenS3{F0: 236, F1: -2147483648, F2: -731}
                            _ = v17
                        default:
                            v18 := GenS1{F0: "", F1: (!(!((v1 ^ (-(v1))) >= (v1 ^ (-(v1))))))}
                            _ = v18
                        }
                        _ = fi16
                    }
                    switch 0 {
                    case -9223372036854775808:
                        for fi19 := 0; fi19 < 9; fi19++ {
                            switch genFn0(65535) {
                            case -690:
                                v1 = v1
                            case -9223372036854775808:
                                var v20 int32 = 0
                                _ = v20
                            }
                            v21 := v3[1]
                            var v22 uint32 = 4294967295
                            _ = fi19
                            _ = v21
                            _ = v22
                        }
                    case -2147483648:
                        var v23 int8 = (-(v5))
                        _ = v23
                    case 32767:
                        var v24 uint = 42
                        _ = v24
                    default:
                        var v25 uint32 = 2
                        _ = v25
                    }
                } else {
                    switch "\n\t" {
                    case "a":
                        fmt.Println(v4, v1, v4)
                    case "":
                        v26 := GenS4{F0: 190, F1: 128, F2: 415}
                        _ = v26
                    }
                    v2 = append(v2, -9223372036854775808)
                }
            }
        }
    default:
        var v27 uint64 = 934
        _ = v27
    }
    fmt.Println(v1, v4)
    if true {
        for fi28 := 0; fi28 < 6; fi28++ {
            v1 += (v1 << ((0) & 63))
            _ = fi28
        }
        for fi29 := 0; fi29 < 17; fi29++ {
            v1 = -32768
            _ = fi29
        }
        var v30 string = ("" + ("foo bar" + ""))
        _ = v30
    }
    v3[1000] = 128
    for fi31 := 0; fi31 < 7; fi31++ {
        if (0 < -1) {
            fmt.Println(fi31)
            v1 += (^(v1))
        }
        _ = fi31
    }
    v4 += -1
    if true {
        fmt.Println(v4)
    }
    if false {
        v32 := GenS5{F0: 972, F1: (-(v5)), F2: 0}
        var v33 int16 = -66
        _ = v32
        _ = v33
    } else {
        var v34 float64 = (((1e10 + 3.14159) - float64((3.14159 / -2.25))) * (-2.25 / (((-2.25 + 0.0) / 1.5) + 1e10)))
        var v35 string = ("a" + "")
        v36 := GenS6{F0: 197}
        _ = v34
        _ = v35
        _ = v36
    }
    fmt.Println(v1, v4)
    v37 := GenS5{F0: 18446744073709551615, F1: v5, F2: 742}
    for fi38 := 0; fi38 < 19; fi38++ {
        delete(v3, 0)
        var v39 bool = (false || true)
        v5 += 100
        _ = fi38
        _ = v39
    }
    if len(v2) > 0 {
        v2[0] = 1
    }
    fmt.Println(v4, v1, v5)
    fmt.Println(v1, v4, v5)
    var v40 uint32 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
