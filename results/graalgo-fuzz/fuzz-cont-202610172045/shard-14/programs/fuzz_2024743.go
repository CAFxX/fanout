package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 bool; F2 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 uint64; F2 uint }
type GenS5 struct { F0 uint8; F1 string }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 string; F1 uint16 }
type GenS8 struct { F0 int64; F1 int; F2 uint32 }
type GenS9 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 uint32; F1 bool; F2 uint8 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 0
    var v3 string = ""
    var v4 int64 = 396
    var v5 int8 = -128
    v6 := make(map[string]bool)
    v2 += 42
    v7 := GenS1{F0: -2147483648}
    v2 -= v2
    fmt.Println(v3, v4)
    var v8 int8 = v5
    switch v3 {
    case "hello":
        if false {
            if false {
                var v9 string = ("\n\t" + v3)
                var v10 bool = false
                var v11 uint32 = 4294967295
                _ = v10
                _ = v11
                _ = v9
            } else {
                switch ("a" + ("foo bar" + v3)) {
                case "foo bar":
                    v2 = v2
                case "":
                    v3 = "foo bar"
                default:
                    v12 := GenS2{F0: (v3 + ("\n\t" + v3)), F1: true, F2: (v2 + v2)}
                    _ = v12
                }
                var v13 int16 = 0
                _ = v13
            }
        }
    case "foo bar":
        v14 := GenS1{F0: -9223372036854775808}
        _ = v14
    }
    var v15 string = v3
    switch v15 {
    case "":
        switch 2147483647 {
        case 127:
            v4 = -623
        case -1:
            if (false || false) {
                if (!false) {
                    v8 -= v5
                    var v16 bool = (-2.25 != 0.0)
                    v17 := []int{-652, 127, 42}
                    _ = v16
                    _ = v17
                }
                switch 128 {
                case 9223372036854775807:
                    var v18 uint64 = 42
                    _ = v18
                case 127:
                    v19 := GenS3{F0: 615}
                    _ = v19
                case 4294967295:
                    v20 := GenS1{F0: 2147483647}
                    _ = v20
                default:
                    if true {
                        v21 := []int{128, 695}
                        _ = v21
                    }
                }
                v22 := GenS2{F0: "foo bar", F1: true, F2: 2}
                _ = v22
            } else {
                for fi23 := 0; fi23 < 14; fi23++ {
                    v24 := []int{fi23}
                    for fi25 := 0; fi25 < 14; fi25++ {
                        if ((1 >= 1) && false) {
                            var v26 uint64 = 2
                            switch (fi23 / (((-(fi23))) | 1)) {
                            case 127:
                                var v27 int8 = v5
                                _ = v27
                            default:
                                fmt.Println(v2)
                            }
                            _ = v26
                        }
                        fmt.Println(v8, fi25)
                        _ = fi25
                    }
                    var v28 uint = v2
                    _ = fi23
                    _ = v24
                    _ = v28
                }
            }
        case -64:
            if false {
                var v29 int = -839
                v2 += v2
                v30 := []int{(^(v29)), 0, (v29 | v29), v29}
                _ = v29
                _ = v30
            } else {
                if (((^(v2)) >= 0) && false) {
                    switch 127 {
                    case 42:
                        v2 = (v2 + 283)
                    case -812:
                        for fi31 := 0; fi31 < 1; fi31++ {
                            if false {
                                v32 := []int{fi31}
                                v4 += v4
                                _ = v32
                            } else {
                                var v33 bool = ((!(42 > 2)) && (v8 < 0))
                                _ = v33
                            }
                            var v34 uint64 = 350
                            _ = fi31
                            _ = v34
                        }
                    }
                    var v35 string = "hello"
                    _ = v35
                }
                v36 := []int{2147483647}
                fmt.Println(v15)
                _ = v36
            }
        }
    case "foo bar":
        if true {
            fmt.Println(v3)
            fmt.Println(v8)
            for fi37 := 0; fi37 < 3; fi37++ {
                fmt.Println(v2, v15, fi37)
                var v38 uint8 = 4
                _ = fi37
                _ = v38
            }
        } else {
            v6[v15] = (!((!true) && true))
            if false {
                switch 1 {
                case 2147483647:
                    delete(v6, ("\n\t" + ""))
                case 32767:
                    var v39 uint = (v2 >> (((v2 * (v2 ^ 255))) & 63))
                    _ = v39
                case 128:
                    var v40 int64 = (v4 ^ (-(v4)))
                    _ = v40
                default:
                    v5 += -1
                }
                v41 := GenS4{F0: v5, F1: 4294967295, F2: (v2 % ((v2) | 1))}
                _ = v41
            } else {
                v42 := GenS5{F0: 255, F1: v15}
                var v43 float32 = -2.25
                v44 := GenS4{F0: (v8 ^ 77), F1: 42, F2: 18446744073709551615}
                _ = v42
                _ = v43
                _ = v44
            }
            v45 := []int{-2147483648}
            _ = v45
        }
    case "hello":
        v46 := GenS6{F0: 258}
        _ = v46
    default:
        var v47 uint = v2
        _ = v47
    }
    for fi48 := 0; fi48 < 10; fi48++ {
        v49 := GenS5{F0: 1, F1: "a"}
        for fi50 := 0; fi50 < 9; fi50++ {
            if ((fi50 ^ 42) != 2147483647) {
                v51 := GenS7{F0: "hello", F1: 0}
                _ = v51
            }
            v15 = v3
            v3 = ""
            _ = fi50
        }
        v52 := GenS1{F0: fi48}
        _ = fi48
        _ = v49
        _ = v52
    }
    var v53 float32 = 0.1
    if true {
        v54 := []int{128}
        for fi55 := 0; fi55 < 2; fi55++ {
            var v56 uint8 = 96
            if true {
                fmt.Println(v5)
            }
            _ = fi55
            _ = v56
        }
        _ = v54
    }
    for fi57 := 0; fi57 < 1; fi57++ {
        v58 := GenS6{F0: 0}
        for fi59 := 0; fi59 < 4; fi59++ {
            if false {
                switch "a" {
                case "\n\t":
                    switch fi57 {
                    case 2147483647:
                        v8 = (v5 % ((v5) | 1))
                    case -132:
                        fmt.Println(fi57, v15, v4)
                    default:
                        switch "\n\t" {
                        case "":
                            v60 := GenS8{F0: -9223372036854775808, F1: (fi59 ^ (fi59 ^ (fi59 | (fi59 ^ 666)))), F2: 439}
                            _ = v60
                        }
                    }
                case "":
                    v61 := GenS6{F0: 845}
                    _ = v61
                case "a":
                    var v62 int = (^(fi59))
                    _ = v62
                default:
                    switch v15 {
                    case "\n\t":
                        v63 := []int{-2147483648, fi57}
                        delete(v6, v15)
                        if true {
                            v64 := GenS9{F0: true, F1: 960, F2: 4294967295}
                            v65 := GenS7{F0: "", F1: 2}
                            v3 = (v3 + v3)
                            _ = v64
                            _ = v65
                        } else {
                            var v66 uint32 = 785
                            _ = v66
                        }
                        v67 := GenS4{F0: v8, F1: 59, F2: (v2 - (v2 << (((v2 ^ v2)) & 63)))}
                        v68 := GenS5{F0: uint8(v2), F1: v3}
                        _ = v63
                        _ = v67
                        _ = v68
                    }
                }
                v69 := GenS10{F0: 973}
                v6[(v3 + v15)] = (-1 < -2147483648)
                _ = v69
            }
            _ = fi59
        }
        v70 := GenS11{F0: 265, F1: (v53 != (v53 * (0.1 / (0.1 / (1.5 / ((((0.1 * (v53 * float32((float64(((v53 / 1.0) * v53)) * 1.5)))) - v53) - v53) / float32(1.5))))))), F2: 128}
        _ = fi57
        _ = v58
        _ = v70
    }
    fmt.Println(v3, v4, v53)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v6
    _ = v7
    _ = v8
}
