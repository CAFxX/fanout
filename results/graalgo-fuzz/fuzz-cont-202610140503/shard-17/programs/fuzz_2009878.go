package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 string }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 uint64 }
func genFn0(p0 int64, p1 uint32) int8 {
    if (((!(-2147483648 > 32767)) || (!false)) && false) {
        fmt.Println(p0, p1)
        if (!true) {
            p0 += (p0 + (^(p0)))
        }
        switch "hello" {
        case "a":
            fmt.Println(p1, p1)
        }
    } else {
        if true {
            var v4 int64 = 811
            var v5 float64 = (0.0 - 3.14159)
            v6 := GenS1{F0: "hello", F1: true, F2: ""}
            _ = v4
            _ = v5
            _ = v6
        }
        fmt.Println(p1, p1)
    }
    v7 := GenS1{F0: "hello", F1: true, F2: "\n\t"}
    switch "foo bar" {
    case "\n\t":
        v8 := []int{525, -1}
        _ = v8
    case "a":
        v9 := []int{}
        _ = v9
    default:
        var v10 int32 = 0
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v7
    return 127
}
func genFn1(p0 int32, p1 int8) int {
    if (4294967295 < 4294967295) {
        switch 0 {
        case -9223372036854775808:
            p1 -= p1
        case 9223372036854775807:
            p1 -= 127
        default:
            var v11 int8 = p1
            _ = v11
        }
        v12 := []int{32768}
        if len(v12) > 0 {
            v12[0] = -241
        }
        _ = v12
    }
    switch "foo bar" {
    case "foo bar":
        v13 := []int{-2147483648, 34, -9223372036854775808}
        _ = v13
    case "hello":
        v14 := GenS2{F0: 594, F1: 0, F2: 45}
        _ = v14
    case "\n\t":
        if true {
            p0 += (p0 - (p0 ^ -1))
        }
    default:
        var v15 uint32 = 1
        _ = v15
    }
    p1 -= genFn0(42, 42)
    v16 := []int{4294967295, -1}
    _ = p0
    _ = p1
    _ = v16
    return 32768
}

func main() {
    var v1 uint8 = 255
    var v2 int = -836
    var v3 float64 = (0.0 + -2.25)
    v17 := []int{v2}
    v18 := GenS2{F0: 65535, F1: 367, F2: 554}
    if (649 == 1) {
        v19 := GenS3{F0: 127}
        var v20 bool = true
        _ = v19
        _ = v20
    } else {
        fmt.Println(v2)
        if true {
            var v21 uint16 = 184
            switch v2 {
            case 32767:
                if (65535 == 65535) {
                    switch (("foo bar" + "") + ("a" + "")) {
                    case "":
                        if false {
                            v2 = int(v1)
                        }
                    default:
                        var v22 float32 = 1e10
                        _ = v22
                    }
                } else {
                    switch "hello" {
                    case "hello":
                        if len(v17) > 0 {
                            v17[0] = 0
                        }
                    }
                }
            case -9223372036854775808:
                for fi23 := 0; fi23 < 6; fi23++ {
                    switch v2 {
                    case 42:
                        if len(v17) >= 2 {
                            v17 = v17[1:len(v17)]
                        }
                    case -9223372036854775808:
                        fmt.Println(v3)
                    case 32767:
                        var v24 uint64 = 0
                        _ = v24
                    }
                    for fi25 := 0; fi25 < 10; fi25++ {
                        switch "hello" {
                        case "foo bar":
                            if len(v17) > 0 {
                                v17[0] = fi25
                            }
                        }
                        var v26 uint16 = (v21 ^ (^(v21)))
                        _ = fi25
                        _ = v26
                    }
                    _ = fi23
                }
            default:
                var v27 uint32 = uint32(v21)
                _ = v27
            }
            for fi28 := 0; fi28 < 11; fi28++ {
                for fi29 := 0; fi29 < 4; fi29++ {
                    var v30 int32 = -87
                    v21 -= 255
                    if true {
                        var v31 uint16 = (v21 ^ (v21 + v21))
                        _ = len(v17)
                        _ = v31
                    }
                    _ = fi29
                    _ = v30
                }
                for fi32 := 0; fi32 < 14; fi32++ {
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                    _ = fi32
                }
                v3 = v3
                _ = fi28
            }
            _ = v21
        }
    }
    v33 := GenS2{F0: 2, F1: 4294967295, F2: 2}
    v3 = v3
    v3 = (-2.25 / float64(((-2.25 * (1e10 / (0.1 + 1e10))) + 1.5)))
    v34 := GenS2{F0: 2147483647, F1: 2, F2: 621}
    v35 := GenS1{F0: "hello", F1: (1 == 42), F2: "foo bar"}
    fmt.Println(v1)
    if false {
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
        var v36 int = (v2 + (v2 + v2))
        _ = v36
    } else {
        v37 := GenS1{F0: "\n\t", F1: (!true), F2: "foo bar"}
        for fi38 := 0; fi38 < 8; fi38++ {
            fmt.Println(v3)
            v39 := GenS4{F0: v1, F1: 52}
            v2 = fi38
            _ = fi38
            _ = v39
        }
        _ = v37
    }
    v40 := GenS4{F0: (v1 - 187), F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v40
}
