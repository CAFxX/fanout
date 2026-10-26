package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8; F1 int8; F2 string }
type GenS5 struct { F0 int16; F1 int32 }
type GenS6 struct { F0 string; F1 int32; F2 int8 }
type GenS7 struct { F0 uint64; F1 int32 }
type GenS8 struct { F0 int; F1 uint; F2 int16 }
func genFn0(p0 uint8, p1 int64) uint {
    if (0 >= 1) {
        v7 := GenS1{F0: 687, F1: 65535, F2: "hello"}
        _ = v7
    }
    p0 += p0
    if (((-9223372036854775808 == 275) || (486 > 0)) || (("hello" == (("" + "\n\t") + "hello")) && false)) {
        var v8 uint16 = 42
        for fi9 := 0; fi9 < 16; fi9++ {
            v8 = 839
            p0 -= p0
            v10 := GenS1{F0: 0, F1: 65535, F2: "hello"}
            _ = fi9
            _ = v10
        }
        v11 := GenS1{F0: -1, F1: 1, F2: "foo bar"}
        _ = v11
        _ = v8
    }
    _ = p0
    _ = p1
    return 2
}
func genFn1(p0 bool, p1 uint, p2 uint64) int {
    p0 = p0
    switch "foo bar" {
    case "a":
        switch 32767 {
        case 128:
            switch 32767 {
            case 32767:
                var v12 int = 32767
                _ = v12
            case 242:
                v13 := GenS2{F0: p0}
                _ = v13
            case -944:
                switch "hello" {
                case "foo bar":
                    for fi14 := 0; fi14 < 20; fi14++ {
                        v15 := GenS1{F0: 1, F1: (p1 - 18446744073709551615), F2: ("a" + ("hello" + "a"))}
                        switch (^(fi14)) {
                        case 32767:
                            var v16 float64 = 0.1
                            _ = v16
                        case 42:
                            var v17 uint64 = 102
                            _ = v17
                        }
                        p0 = (!false)
                        _ = fi14
                        _ = v15
                    }
                case "a":
                    if (42 != 65535) {
                        v18 := []int{128}
                        v18 = append(v18, 0)
                        _ = v18
                    } else {
                        p1 = (^(p1))
                        if p0 {
                            if (!(genFn0(255, 1) >= (^(p1)))) {
                                var v19 uint32 = 1
                                _ = v19
                            }
                            var v20 int32 = 938
                            _ = v20
                        } else {
                            fmt.Println(p0, p1, p2)
                            v21 := []int{-9223372036854775808, 42, 42, 1}
                            _ = v21
                        }
                        p1 += 18446744073709551615
                    }
                default:
                    var v22 uint64 = 18446744073709551615
                    _ = v22
                }
            default:
                if p0 {
                    var v23 uint = p1
                    for fi24 := 0; fi24 < 7; fi24++ {
                        var v25 uint = 0
                        for fi26 := 0; fi26 < 16; fi26++ {
                            v25 = 4294967295
                            v23 -= genFn0(128, 462)
                            v27 := []int{}
                            _ = fi26
                            _ = v27
                        }
                        _ = fi24
                        _ = v25
                    }
                    switch 9223372036854775807 {
                    case -9223372036854775808:
                        p1 = p1
                    }
                    _ = v23
                } else {
                    v28 := GenS1{F0: -1, F1: (p1 << ((p1) & 63)), F2: "foo bar"}
                    var v29 string = "a"
                    _ = v28
                    _ = v29
                }
            }
        case 1:
            for fi30 := 0; fi30 < 14; fi30++ {
                for fi31 := 0; fi31 < 6; fi31++ {
                    var v32 string = "hello"
                    _ = fi31
                    _ = v32
                }
                var v33 uint8 = 2
                _ = fi30
                _ = v33
            }
        case 9223372036854775807:
            v34 := []int{42, -9223372036854775808}
            _ = v34
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 12
}

func main() {
    var v1 int64 = 42
    var v2 bool = true
    var v3 uint16 = 42
    v4 := make(map[string]uint32)
    v5 := make(map[int]string)
    var v6 uint = 42
    v35 := GenS3{F0: v1}
    v36 := GenS4{F0: -128, F1: -20, F2: ""}
    fmt.Println(v2)
    if (v2 && v2) {
        v37 := GenS1{F0: 32767, F1: v6, F2: "a"}
        var v38 int32 = 0
        _ = v37
        _ = v38
    } else {
        switch "foo bar" {
        case "\n\t":
            switch 32767 {
            case -2147483648:
                delete(v4, "a")
            }
        case "foo bar":
            var v39 uint32 = 112
            _ = v39
        case "":
            v40 := []int{-277}
            _ = v40
        default:
            v41 := GenS2{F0: v2}
            _ = v41
        }
    }
    switch "a" {
    case "\n\t":
        switch 32768 {
        case 237:
            v42 := GenS5{F0: -32768, F1: 1}
            _ = v42
        case 1:
            delete(v5, 128)
        default:
            v43 := []int{int(v1)}
            _ = v43
        }
    case "foo bar":
        if v2 {
            if true {
                if v2 {
                    v1 += v1
                }
                fmt.Println(v6, v3)
                for fi44 := 0; fi44 < 16; fi44++ {
                    v45 := v5[(^(fi44))]
                    if (v2 && (1 != 2)) {
                        var v46 int16 = -500
                        v2 = (int64(v6) == int64(fi44))
                        v46 -= 308
                        _ = v46
                    }
                    v47 := GenS5{F0: 32767, F1: 0}
                    _ = fi44
                    _ = v45
                    _ = v47
                }
            }
            var v48 string = "hello"
            _ = v48
        }
    default:
        v49 := GenS6{F0: ("hello" + (("a" + "") + "foo bar")), F1: -639, F2: 127}
        _ = v49
    }
    v3 -= v3
    if false {
        var v50 uint = v6
        v1 -= -9223372036854775808
        switch -2147483648 {
        case 128:
            fmt.Println(v50)
        case -544:
            v1 -= v1
        }
        _ = v50
    }
    var v51 uint32 = 2
    delete(v5, genFn1(false, v6, 880))
    if (((0.0 + 1.5) * 0.0) <= -2.25) {
        v52 := []int{2147483647}
        for fi53 := 0; fi53 < 2; fi53++ {
            v54 := GenS7{F0: 4294967295, F1: 0}
            v1 -= (v1 - v1)
            _ = fi53
            _ = v54
        }
        _ = v52
    }
    if ((0.1 + (3.14159 / 1.0)) > (3.14159 + (1.5 / (1.5 - 1e10)))) {
        v55 := GenS7{F0: 86, F1: -2147483648}
        _ = v55
    } else {
        for fi56 := 0; fi56 < 11; fi56++ {
            fmt.Println(v3)
            v1 -= (-(v1))
            if v2 {
                fmt.Println(v3, fi56)
                var v57 string = ("hello" + "a")
                v6 += (v6 + v6)
                _ = v57
            }
            _ = fi56
        }
    }
    switch "foo bar" {
    case "hello":
        v3 -= 2
    }
    v51 -= v51
    v58 := GenS8{F0: 32767, F1: 987, F2: -727}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v51
    _ = v58
    _ = v6
}
