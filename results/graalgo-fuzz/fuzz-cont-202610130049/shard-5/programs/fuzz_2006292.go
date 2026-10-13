package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint }
type GenS2 struct { F0 int16; F1 int64; F2 int }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 uint32 = 2
    var v2 int32 = 42
    v3 := make(map[string]string)
    var v4 bool = false
    var v5 bool = (2 == 897)
    v6 := make(map[int]bool)
    if true {
        if true {
            switch "foo bar" {
            case "a":
                var v7 int32 = 1
                _ = v7
            }
            if (9223372036854775807 > 9223372036854775807) {
                if (v2 > v2) {
                    for fi8 := 0; fi8 < 12; fi8++ {
                        v4 = (v1 <= v1)
                        _ = fi8
                    }
                    fmt.Println(v5, v2, v2)
                    switch -1 {
                    case 32767:
                        for fi9 := 0; fi9 < 6; fi9++ {
                            v1 = (v1 % ((v1) | 1))
                            var v10 int32 = v2
                            _ = fi9
                            _ = v10
                        }
                    case 2147483647:
                        switch 1 {
                        case 33:
                            v4 = true
                        case 127:
                            var v11 int16 = -700
                            _ = v11
                        }
                    default:
                        v12 := GenS1{F0: v1, F1: (v1 - 854), F2: 18446744073709551615}
                        _ = v12
                    }
                }
                switch 32768 {
                case 127:
                    v1 = (v1 << ((42) & 63))
                case 128:
                    var v13 float64 = (1.5 * 0.1)
                    _ = v13
                case 42:
                    var v14 float64 = 1e10
                    _ = v14
                default:
                    var v15 uint32 = 42
                    _ = v15
                }
            } else {
                var v16 uint16 = 0
                _ = v16
            }
            v17 := []int{9223372036854775807, 2147483647, 9223372036854775807, 128}
            _ = v17
        } else {
            if (1 != (v2 * v2)) {
                var v18 int32 = 42
                v19 := GenS1{F0: v1, F1: (v1 << ((255) & 63)), F2: 444}
                _ = v18
                _ = v19
            } else {
                for fi20 := 0; fi20 < 1; fi20++ {
                    for fi21 := 0; fi21 < 19; fi21++ {
                        if true {
                            fmt.Println(fi20, v4, fi20)
                            if (103 > 2) {
                                if v5 {
                                    v22 := GenS1{F0: (v1 << ((4294967295) & 63)), F1: (v1 >> ((18446744073709551615) & 63)), F2: 720}
                                    v4 = v5
                                    delete(v3, "a")
                                    if false {
                                        v2 += (v2 - (v2 ^ (v2 - (v2 ^ v2))))
                                        switch fi21 {
                                        case 737:
                                            v2 = v2
                                            if v5 {
                                                v2 -= (-(v2))
                                                v1 = (^(v1))
                                                v23 := GenS2{F0: 412, F1: 1, F2: (fi21 & (^(fi20)))}
                                                _ = v23
                                            }
                                        case 195:
                                            v4 = (50 >= 127)
                                        }
                                        v24 := GenS1{F0: (v1 >> ((28) & 63)), F1: v1, F2: 18446744073709551615}
                                        _ = v24
                                    }
                                    _ = v22
                                }
                                fmt.Println(v1)
                                v25 := []int{(fi20 * 4294967295)}
                                _ = v25
                            }
                            v5 = (128 != 127)
                            fmt.Println(v5, v1)
                        } else {
                            var v26 uint64 = 827
                            switch fi21 {
                            case 185:
                                var v27 float64 = ((((3.14159 / -2.25) + (0.0 - 0.0)) - -2.25) - 0.1)
                                _ = v27
                            case 1:
                                v6[(-(fi21))] = v5
                            default:
                                v3["a"] = "a"
                            }
                            v28 := GenS3{F0: 128}
                            _ = v26
                            _ = v28
                        }
                        var v29 uint8 = 255
                        switch 895 {
                        case 32768:
                            var v30 uint16 = 2
                            _ = v30
                        case 32767:
                            v1 = v1
                            _ = len(v3)
                            fmt.Println(v29, v2, v4)
                        case 128:
                            fmt.Println(v1, v29)
                            v1 -= v1
                        }
                        _ = fi21
                        _ = v29
                    }
                    _ = fi20
                }
            }
            delete(v6, 520)
            fmt.Println(v2, v5)
        }
    } else {
        var v31 uint16 = 42
        _ = v31
    }
    v5 = (0 == 42)
    switch 32768 {
    case -1:
        fmt.Println(v1, v4, v4)
    case -9223372036854775808:
        _ = len(v6)
    case 311:
        for fi32 := 0; fi32 < 7; fi32++ {
            var v33 int8 = 100
            var v34 int64 = 0
            _ = fi32
            _ = v33
            _ = v34
        }
    default:
        v35 := GenS1{F0: (v1 - v1), F1: v1, F2: 752}
        _ = v35
    }
    v5 = (-669 == 9223372036854775807)
    v4 = false
    var v36 uint32 = 440
    v37 := v3["hello"]
    v38 := []int{-2147483648, 128, 0, -451}
    if len(v38) > 0 {
        v38[0] = -148
    }
    v39 := v3[v37]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
