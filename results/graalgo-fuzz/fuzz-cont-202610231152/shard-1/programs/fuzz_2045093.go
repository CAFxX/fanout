package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint16; F1 uint }
func genFn0() int {
    if false {
        switch "a" {
        case "foo bar":
            switch 42 {
            case 4294967295:
                var v5 uint = 65535
                _ = v5
            case 32767:
                v6 := []int{32768}
                _ = v6
            default:
                v7 := GenS1{F0: 255, F1: 158, F2: 42}
                _ = v7
            }
        case "hello":
            var v8 uint64 = 0
            _ = v8
        default:
            v9 := []int{}
            _ = v9
        }
        var v10 float64 = 0.0
        _ = v10
    } else {
        var v11 bool = true
        if false {
            v11 = true
        } else {
            switch "foo bar" {
            case "foo bar":
                switch "" {
                case "hello":
                    v11 = false
                default:
                    var v12 uint64 = 756
                    _ = v12
                }
            case "\n\t":
                var v13 uint64 = 478
                _ = v13
            case "a":
                v11 = v11
            }
            v14 := GenS2{F0: -92, F1: ("\n\t" + ("hello" + "hello"))}
            _ = v14
        }
        _ = v11
    }
    v15 := GenS3{F0: 195}
    var v16 int32 = -39
    v16 -= v16
    _ = v15
    _ = v16
    return -109
}

func main() {
    v1 := []int{-101}
    v2 := make(map[string]uint32)
    v3 := []int{42, -1}
    v4 := make(map[int]int)
    if ((-1 <= 42) || true) {
        if (1e10 >= float32(3.14159)) {
            var v17 int = 9223372036854775807
            _ = v17
        }
    }
    switch 9223372036854775807 {
    case 1:
        _ = len(v3)
    default:
        var v18 string = "a"
        _ = v18
    }
    var v19 int = 42
    v20 := GenS3{F0: 128}
    fmt.Println(v19)
    v3 = append(v3, (v19 - (v19 << ((42) & 63))))
    v19 = v19
    for fi21 := 0; fi21 < 20; fi21++ {
        if (2147483647 <= 1) {
            var v22 uint32 = 991
            if len(v1) > 0 {
                v1[0] = genFn0()
            }
            _ = v22
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi21
    }
    fmt.Println(v19)
    v19 -= 32767
    if true {
        for fi23 := 0; fi23 < 18; fi23++ {
            if (true && (true && ((874 < 1) || true))) {
                v24 := GenS2{F0: -32768, F1: "foo bar"}
                for fi25 := 0; fi25 < 6; fi25++ {
                    v19 -= genFn0()
                    _ = fi25
                }
                _ = v24
            } else {
                v26 := GenS1{F0: 2, F1: 2, F2: 784}
                _ = v26
            }
            _ = fi23
        }
        if true {
            var v27 float64 = 1.5
            _ = v27
        } else {
            v19 += v19
        }
    } else {
        fmt.Println(v19)
    }
    if (true && false) {
        if (true || ((-2.25 - (-2.25 + 0.1)) < ((0.0 * -2.25) * (3.14159 + 0.0)))) {
            switch "hello" {
            case "\n\t":
                for fi28 := 0; fi28 < 7; fi28++ {
                    if len(v3) > 0 {
                        v3[0] = v19
                    }
                    for fi29 := 0; fi29 < 6; fi29++ {
                        var v30 uint32 = 554
                        fmt.Println(v30, v30)
                        var v31 float32 = (-2.25 / 0.1)
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    _ = fi28
                }
            case "hello":
                v19 += v19
            default:
                switch ("" + "\n\t") {
                case "\n\t":
                    if (0 == (v19 << ((641) & 63))) {
                        v19 = (v19 ^ genFn0())
                        v1 = append(v1, v19)
                    } else {
                        if ((-38 <= 473) || (((false || true) && (655 < 939)) && false)) {
                            v19 = (v19 - 9223372036854775807)
                        }
                        v1 = append(v1, v19)
                    }
                case "a":
                    switch "foo bar" {
                    case "":
                        fmt.Println(v19)
                    case "foo bar":
                        v19 += (v19 | v19)
                    default:
                        v19 -= (^(v19))
                    }
                case "foo bar":
                    fmt.Println(v19)
                }
            }
            switch v19 {
            case -9223372036854775808:
                switch ("\n\t" + (("\n\t" + "foo bar") + (("hello" + "a") + ("\n\t" + "foo bar")))) {
                case "\n\t":
                    var v32 uint = 501
                    _ = v32
                case "":
                    v19 = 42
                }
            case 128:
                v19 += (v19 + (v19 - 127))
            default:
                v33 := GenS4{F0: "foo bar"}
                _ = v33
            }
            for fi34 := 0; fi34 < 1; fi34++ {
                if (1 > -9223372036854775808) {
                    fmt.Println(v19)
                } else {
                    if ((0 == 0) || false) {
                        for fi35 := 0; fi35 < 14; fi35++ {
                            v19 += fi35
                            delete(v4, (-(fi35)))
                            v2["foo bar"] = 1
                            fmt.Println(fi34, fi34, fi35)
                            _ = fi35
                        }
                    }
                }
                _ = fi34
            }
        } else {
            for fi36 := 0; fi36 < 10; fi36++ {
                var v37 uint = 4294967295
                _ = fi36
                _ = v37
            }
            var v38 int32 = 2147483647
            _ = v38
        }
        fmt.Println(v19)
    } else {
        var v39 bool = true
        switch 0 {
        case 4294967295:
            for fi40 := 0; fi40 < 10; fi40++ {
                if true {
                    v19 -= (fi40 - (v19 % ((393) | 1)))
                }
                var v41 int64 = -785
                for fi42 := 0; fi42 < 9; fi42++ {
                    if (1000 == 32767) {
                        fmt.Println(fi42)
                        switch "hello" {
                        case "\n\t":
                            v43 := GenS5{F0: 0, F1: 834}
                            v44 := GenS1{F0: 709, F1: 1, F2: 471}
                            _ = v43
                            _ = v44
                        case "":
                            _ = len(v1)
                        case "hello":
                            switch -157 {
                            case 32768:
                                v39 = v39
                            case -9223372036854775808:
                                delete(v4, genFn0())
                            case 0:
                                if (!(((((1e10 / (1.5 - float64(-2.25))) + -2.25) + float64(1e10)) - 3.14159) < 1e10)) {
                                    fmt.Println(v19, v39, v39)
                                    v41 = (v41 ^ (v41 << ((1) & 63)))
                                    v45 := GenS2{F0: -1, F1: ""}
                                    var v46 uint16 = 32767
                                    _ = v45
                                    _ = v46
                                }
                            }
                        default:
                            switch (fi42 % ((32767) | 1)) {
                            case 2147483647:
                                var v47 float64 = 0.0
                                _ = v47
                            }
                        }
                    } else {
                        v39 = (255 != 1)
                        var v48 float32 = -2.25
                        var v49 int8 = 100
                        _ = v48
                        _ = v49
                    }
                    _ = fi42
                }
                _ = fi40
                _ = v41
            }
        }
        _ = v39
    }
    fmt.Println(v19)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
