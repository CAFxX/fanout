package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 string; F1 int32; F2 uint8 }
func genFn0(p0 uint, p1 bool) int {
    v7 := []int{32768, 42, 127, 2147483647}
    var v8 int32 = 0
    if false {
        p0 -= 2
        v8 = v8
        switch ("foo bar" + "\n\t") {
        case "":
            for fi9 := 0; fi9 < 17; fi9++ {
                var v10 uint8 = 1
                var v11 int = fi9
                _ = fi9
                _ = v10
                _ = v11
            }
        case "foo bar":
            p1 = false
        }
    }
    p0 = (p0 ^ p0)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 678
}
func genFn1(p0 bool, p1 uint8, p2 int64) uint16 {
    fmt.Println(p2)
    v12 := GenS1{F0: 42}
    for fi13 := 0; fi13 < 8; fi13++ {
        if p0 {
            fmt.Println(p1)
            v14 := []int{fi13, genFn0(65535, p0), genFn0(2, p0), 0}
            _ = v14
        } else {
            for fi15 := 0; fi15 < 13; fi15++ {
                v16 := GenS1{F0: -657}
                _ = fi15
                _ = v16
            }
            v17 := []int{(fi13 - (fi13 % (((fi13 >> ((18446744073709551615) & 63))) | 1))), (fi13 / ((fi13) | 1)), genFn0(2, ((-2147483648 <= 0) && (fi13 >= fi13)))}
            var v18 uint8 = (p1 & p1)
            _ = v17
            _ = v18
        }
        _ = fi13
    }
    switch int(p2) {
    case 127:
        v19 := GenS2{F0: (p2 + (p2 | (p2 | p2))), F1: "foo bar"}
        _ = v19
    case -2147483648:
        var v20 uint32 = 2
        _ = v20
    }
    var v21 uint = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v21
    return 381
}

func main() {
    var v1 uint64 = 2
    v2 := []int{42, 763, -9223372036854775808, 738}
    v3 := make(map[string]int)
    var v4 int32 = 263
    var v5 uint = 18446744073709551615
    var v6 uint64 = (v1 + (v1 ^ 238))
    if (!false) {
        v22 := GenS1{F0: 0}
        v23 := GenS1{F0: -1}
        v24 := GenS2{F0: 417, F1: "hello"}
        _ = v22
        _ = v23
        _ = v24
    } else {
        v6 = (v6 % (((v1 / ((v6) | 1))) | 1))
    }
    v6 -= v1
    var v25 uint8 = 2
    v26 := GenS3{F0: 32768, F1: 1}
    var v27 uint32 = 2147483647
    var v28 uint = v5
    var v29 bool = (-128 == int8(v6))
    var v30 uint32 = 1
    if v29 {
        for fi31 := 0; fi31 < 12; fi31++ {
            var v32 int32 = (v4 ^ (v4 + (v4 % (((v4 / ((v4) | 1))) | 1))))
            for fi33 := 0; fi33 < 7; fi33++ {
                v30 -= (v27 >> ((v5) & 63))
                v27 = v27
                var v34 int64 = 0
                _ = fi33
                _ = v34
            }
            switch "a" {
            case "a":
                v1 -= (v6 + v6)
            case "":
                v35 := GenS4{F0: (v25 ^ v25)}
                _ = v35
            default:
                for fi36 := 0; fi36 < 13; fi36++ {
                    switch fi36 {
                    case 127:
                        v6 += 375
                    case -9223372036854775808:
                        v37 := GenS4{F0: uint8(v30)}
                        _ = v37
                    case -959:
                        v3["a"] = 9223372036854775807
                    default:
                        var v38 uint32 = uint32(v5)
                        _ = v38
                    }
                    _ = fi36
                }
            }
            _ = fi31
            _ = v32
        }
        v28 = v28
        for fi39 := 0; fi39 < 1; fi39++ {
            var v40 int32 = (^(v4))
            _ = fi39
            _ = v40
        }
    } else {
        fmt.Println(v5, v29)
        var v41 bool = v29
        if v41 {
            v42 := GenS5{F0: (100 <= 0)}
            var v43 float64 = -2.25
            _ = v42
            _ = v43
        }
        _ = v41
    }
    var v44 uint64 = (^(v1))
    v3["foo bar"] = 127
    if len(v2) > 0 {
        v2[0] = 809
    }
    for fi45 := 0; fi45 < 13; fi45++ {
        fmt.Println(v4, v4)
        _ = fi45
    }
    switch ("a" + "hello") {
    case "a":
        switch 181 {
        case 918:
            switch 774 {
            case 32767:
                var v46 uint64 = (v44 * v1)
                _ = v46
            case 326:
                if (v29 && (9223372036854775807 < -9223372036854775808)) {
                    v28 -= 0
                    switch "foo bar" {
                    case "hello":
                        var v47 string = ("hello" + "")
                        _ = v47
                    case "\n\t":
                        if v29 {
                            var v48 uint64 = (v1 ^ (^(v44)))
                            _ = v48
                        } else {
                            if v29 {
                                fmt.Println(v29)
                            }
                            v49 := GenS6{F0: 4294967295}
                            _ = v49
                        }
                    default:
                        var v50 uint16 = genFn1(v29, 42, 0)
                        _ = v50
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
            case -2147483648:
                switch "" {
                case "hello":
                    var v51 uint8 = (^(v25))
                    _ = v51
                case "":
                    v6 += v1
                case "a":
                    v3["foo bar"] = 1
                }
            }
        case 908:
            v6 = (v44 - (^(v1)))
        }
    default:
        v52 := GenS7{F0: "\n\t", F1: int32(v5), F2: 2}
        _ = v52
    }
    var v53 uint8 = v25
    v5 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v44
    _ = v5
    _ = v53
    _ = v6
}
