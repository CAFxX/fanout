package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint32 }
func genFn0() int64 {
    v7 := []int{32767}
    for fi8 := 0; fi8 < 7; fi8++ {
        switch "foo bar" {
        case "a":
            var v9 uint32 = 470
            _ = v9
        default:
            if len(v7) > 0 {
                v7[0] = fi8
            }
        }
        _ = fi8
    }
    _ = v7
    return -1
}

func main() {
    var v1 string = ""
    v2 := make(map[uint32]int)
    v3 := make(map[int]bool)
    var v4 int32 = 0
    v5 := []uint32{2147483647, 2147483648, 2147483647}
    var v6 bool = true
    for fi10 := 0; fi10 < 14; fi10++ {
        switch v1 {
        case "":
            v1 = v1
        case "foo bar":
            var v11 float32 = ((1.5 + 0.0) * (0.0 - 1e10))
            _ = v11
        }
        _ = fi10
    }
    v12 := GenS1{F0: 42, F1: 65535}
    if false {
        v13 := GenS1{F0: genFn0(), F1: 2}
        if false {
            for fi14 := 0; fi14 < 2; fi14++ {
                if v6 {
                    v1 = v1
                    v1 = v1
                    switch (fi14 ^ fi14) {
                    case 1:
                        v15 := GenS2{F0: 858, F1: 4294967295}
                        _ = v15
                    case -745:
                        switch ((v1 + (("\n\t" + "") + v1)) + (v1 + "foo bar")) {
                        case "foo bar":
                            if len(v5) > 0 {
                                v5[0] = 2
                            }
                        case "hello":
                            v4 -= (-(v4))
                        default:
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                    default:
                        fmt.Println(v4)
                    }
                } else {
                    v6 = false
                }
                if false {
                    v16 := GenS2{F0: 2, F1: 0}
                    switch "\n\t" {
                    case "a":
                        for fi17 := 0; fi17 < 10; fi17++ {
                            switch -191 {
                            case 1:
                                var v18 int32 = v4
                                _ = v18
                            case 2147483647:
                                var v19 int16 = 32767
                                _ = v19
                            default:
                                var v20 string = "foo bar"
                                _ = v20
                            }
                            var v21 int = -404
                            _ = fi17
                            _ = v21
                        }
                    default:
                        var v22 bool = (!v6)
                        _ = v22
                    }
                    for fi23 := 0; fi23 < 15; fi23++ {
                        var v24 uint32 = 517
                        fmt.Println(v4, fi14)
                        var v25 uint64 = 4294967295
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    _ = v16
                }
                switch v1 {
                case "foo bar":
                    switch ((v1 + "hello") + "a") {
                    case "\n\t":
                        var v26 bool = ((float64((-2.25 + ((0.1 / (0.0 + 1.5)) + ((1.5 - 1.5) * 3.14159)))) + float64(0.0)) != 0.0)
                        _ = v26
                    case "":
                        var v27 bool = (-1 > 1)
                        _ = v27
                    default:
                        if len(v5) > 0 {
                            v5[0] = 65535
                        }
                    }
                default:
                    var v28 uint64 = 18446744073709551615
                    _ = v28
                }
                _ = fi14
            }
            switch "hello" {
            case "":
                for fi29 := 0; fi29 < 16; fi29++ {
                    v6 = (false || true)
                    fmt.Println(v6, v4, v6)
                    switch v1 {
                    case "foo bar":
                        if (0.0 < (0.1 * 0.0)) {
                            delete(v2, 639)
                        } else {
                            v30 := GenS3{F0: 2147483648}
                            _ = v30
                        }
                    case "a":
                        var v31 int8 = 0
                        _ = v31
                    }
                    _ = fi29
                }
            default:
                switch ("foo bar" + "foo bar") {
                case "foo bar":
                    switch -1 {
                    case 2147483647:
                        for fi32 := 0; fi32 < 16; fi32++ {
                            switch (v1 + v1) {
                            case "foo bar":
                                var v33 float64 = 1.5
                                _ = v33
                            default:
                                _ = len(v5)
                            }
                            fmt.Println(fi32, v1)
                            var v34 uint64 = 2
                            _ = fi32
                            _ = v34
                        }
                    default:
                        var v35 uint8 = 255
                        _ = v35
                    }
                case "":
                    v6 = v6
                }
            }
            var v36 int16 = 42
            _ = v36
        }
        switch -2147483648 {
        case 32768:
            v6 = v6
        case 754:
            v6 = ((((((0.0 + 1e10) - (float64(float32(float64((1e10 + 0.1)))) + 1.5)) + 1.5) == 1e10) && v6) && (1 < 1))
        }
        _ = v13
    }
    v37 := GenS2{F0: 785, F1: 475}
    var v38 int = 0
    v39 := GenS2{F0: 1, F1: 48}
    for fi40 := 0; fi40 < 12; fi40++ {
        v4 = 42
        _ = fi40
    }
    v4 = (v4 | 42)
    v38 = (v38 + 32767)
    if v6 {
        v4 = v4
    } else {
        v2[910] = -1
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
