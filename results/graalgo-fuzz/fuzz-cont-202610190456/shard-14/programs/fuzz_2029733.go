package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 bool; F2 uint16 }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 int8; F1 uint8; F2 int8 }
type GenS6 struct { F0 int64; F1 int8 }
type GenS7 struct { F0 string }
func genFn0(p0 int64, p1 int, p2 int) int32 {
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 float64 = 1e10
        _ = fi5
        _ = v6
    }
    var v7 uint32 = 983
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -207
}
func genFn1(p0 uint8, p1 uint8, p2 bool) uint64 {
    fmt.Println(p1)
    v8 := GenS1{F0: 4294967295}
    p0 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 18446744073709551615
}

func main() {
    var v1 uint8 = 89
    v2 := make(map[string]int)
    var v3 float32 = 0.1
    var v4 int32 = 42
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 uint32 = 2147483648
        v2["foo bar"] = fi9
        _ = fi9
        _ = v10
    }
    switch 127 {
    case 1:
        v4 -= genFn0(-1, 540, 9223372036854775807)
    case -9223372036854775808:
        v11 := GenS1{F0: -9223372036854775808}
        _ = v11
    case 42:
        for fi12 := 0; fi12 < 8; fi12++ {
            v1 = 215
            _ = fi12
        }
    default:
        var v13 int = -1
        _ = v13
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        v4 -= (v4 ^ -1)
        var v15 uint8 = v1
        v16 := GenS2{F0: fi14}
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 uint16 = 255
    var v18 int64 = 733
    v19 := []int{-1, -755, -9223372036854775808, 0}
    v4 -= (v4 | (-(v4)))
    for fi20 := 0; fi20 < 15; fi20++ {
        if false {
            for fi21 := 0; fi21 < 4; fi21++ {
                _ = len(v2)
                _ = fi21
            }
            v3 += v3
            for fi22 := 0; fi22 < 17; fi22++ {
                var v23 bool = true
                _ = len(v2)
                v24 := v2["\n\t"]
                _ = fi22
                _ = v23
                _ = v24
            }
        } else {
            v3 -= v3
        }
        switch 1 {
        case 1:
            if false {
                v17 -= v17
            }
        case -9223372036854775808:
            var v25 uint16 = v17
            _ = v25
        case -564:
            switch 2147483647 {
            case 42:
                if (v18 == 42) {
                    v18 = (v18 ^ (-(v18)))
                    if true {
                        var v26 int32 = 42
                        _ = v26
                    } else {
                        var v27 uint8 = (v1 + (v1 ^ 128))
                        _ = v27
                    }
                } else {
                    v19 = append(v19, (-(fi20)))
                    for fi28 := 0; fi28 < 12; fi28++ {
                        var v29 int64 = v18
                        switch (fi28 + fi28) {
                        case 481:
                            switch (-(fi20)) {
                            case -1:
                                if (true && false) {
                                    if true {
                                        if (false || false) {
                                            if false {
                                                if ("\n\t" == "a") {
                                                    v30 := GenS1{F0: fi20}
                                                    v3 = float32(((((1e10 + -2.25) + -2.25) * 0.0) / 1e10))
                                                    _ = v30
                                                }
                                            }
                                        }
                                        if (("\n\t" > "") || false) {
                                            if (0 >= 4294967295) {
                                                v18 = (v29 * (v29 * (v18 >> ((42) & 63))))
                                                v18 += v18
                                            }
                                            var v31 int64 = v18
                                            _ = v31
                                        } else {
                                            v32 := GenS3{F0: fi28, F1: false, F2: 2}
                                            _ = v32
                                        }
                                    }
                                }
                            default:
                                var v33 uint16 = v17
                                _ = v33
                            }
                        case 936:
                            var v34 string = "foo bar"
                            _ = v34
                        default:
                            if (0 < 4294967295) {
                                if false {
                                    fmt.Println(v4)
                                    v35 := GenS4{F0: (^(v4)), F1: (^(v4))}
                                    _ = len(v19)
                                    v36 := GenS5{F0: 42, F1: v1, F2: -112}
                                    v4 = 2147483647
                                    _ = v35
                                    _ = v36
                                }
                                var v37 uint8 = v1
                                if (0.1 == (float64(1e10) - 1e10)) {
                                    if ((v37 - (v1 + 133)) < v1) {
                                        switch "\n\t" {
                                        case "":
                                            var v38 int16 = 26
                                            _ = v38
                                        case "foo bar":
                                            v39 := GenS6{F0: -693, F1: -13}
                                            _ = v39
                                        case "\n\t":
                                            v40 := GenS7{F0: ("a" + "\n\t")}
                                            _ = v40
                                        }
                                        v1 -= v37
                                    }
                                    v41 := GenS3{F0: (fi20 % ((42) | 1)), F1: (!(!(0.0 < (float32(1e10) * v3)))), F2: v17}
                                    _ = len(v19)
                                    _ = v41
                                }
                                _ = v37
                            } else {
                                var v42 uint8 = v1
                                v29 = (v18 & (v18 * 42))
                                v17 += v17
                                _ = v42
                            }
                        }
                        var v43 float64 = float64(v3)
                        _ = fi28
                        _ = v29
                        _ = v43
                    }
                    var v44 uint8 = v1
                    _ = v44
                }
            default:
                var v45 int32 = (v4 >> ((173) & 63))
                _ = v45
            }
        }
        _ = fi20
    }
    var v46 string = "hello"
    for fi47 := 0; fi47 < 18; fi47++ {
        delete(v2, ("hello" + "\n\t"))
        fmt.Println(v3)
        _ = fi47
    }
    switch "" {
    case "\n\t":
        var v48 uint8 = v1
        _ = v48
    case "hello":
        var v49 float64 = (-2.25 * 0.1)
        _ = v49
    default:
        v4 = v4
    }
    var v50 bool = true
    switch "hello" {
    case "a":
        fmt.Println(v46, v18, v3)
    case "foo bar":
        fmt.Println(v50, v17, v46)
    }
    var v51 int8 = 42
    for fi52 := 0; fi52 < 2; fi52++ {
        for fi53 := 0; fi53 < 3; fi53++ {
            fmt.Println(v1, v46, v50)
            switch (fi53 | (fi53 << ((18446744073709551615) & 63))) {
            case 440:
                v54 := GenS1{F0: int(v4)}
                _ = v54
            case 128:
                var v55 int32 = (v4 % (((v4 ^ (v4 + (v4 + (v4 / ((v4) | 1)))))) | 1))
                _ = v55
            case 2147483647:
                if len(v19) > 0 {
                    v19[0] = 871
                }
            default:
                v17 = (^(v17))
            }
            _ = fi53
        }
        _ = fi52
    }
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v50
    _ = v51
}
