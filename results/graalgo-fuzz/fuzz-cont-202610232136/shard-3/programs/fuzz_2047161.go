package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 bool }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 string; F1 uint16; F2 int }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int32; F1 uint16 }

func main() {
    var v1 uint16 = 65535
    var v2 float64 = (0.1 * 0.1)
    v3 := make(map[string]bool)
    v4 := make(map[int]string)
    v5 := GenS1{F0: 255, F1: -1, F2: true}
    if false {
        v6 := GenS1{F0: 42, F1: -2147483648, F2: (false || true)}
        var v7 uint16 = 2
        _ = v6
        _ = v7
    } else {
        v1 += v1
        if ((42 != 42) && false) {
            for fi8 := 0; fi8 < 2; fi8++ {
                delete(v3, ("a" + "a"))
                if false {
                    for fi9 := 0; fi9 < 19; fi9++ {
                        v3["\n\t"] = false
                        _ = fi9
                    }
                    if ((("hello" + ("\n\t" + "\n\t")) + "a") >= "hello") {
                        for fi10 := 0; fi10 < 6; fi10++ {
                            switch "\n\t" {
                            case "\n\t":
                                var v11 uint = 1
                                _ = v11
                            case "a":
                                var v12 uint = 42
                                _ = v12
                            }
                            _ = fi10
                        }
                    }
                    for fi13 := 0; fi13 < 10; fi13++ {
                        if true {
                            switch "" {
                            case "foo bar":
                                var v14 uint = uint(v1)
                                _ = v14
                            default:
                                var v15 float64 = (float64((1e10 + (0.1 * 0.0))) + v2)
                                _ = v15
                            }
                            fmt.Println(v1, fi8, fi13)
                            switch "a" {
                            case "hello":
                                v16 := v4[(fi13 / ((fi13) | 1))]
                                _ = v16
                            case "a":
                                if (4294967295 > 4294967295) {
                                    v1 -= 951
                                    fmt.Println(v2, fi8, v2)
                                    fmt.Println(fi13, v2, fi13)
                                }
                            default:
                                switch "a" {
                                case "\n\t":
                                    var v17 uint64 = 4294967295
                                    _ = v17
                                case "foo bar":
                                    v18 := []int{}
                                    _ = v18
                                case "a":
                                    var v19 string = "foo bar"
                                    _ = v19
                                default:
                                    v20 := GenS2{F0: -1, F1: (^(fi8))}
                                    _ = v20
                                }
                            }
                        }
                        var v21 bool = (true && (1 >= -854))
                        v3[""] = v21
                        _ = fi13
                        _ = v21
                    }
                } else {
                    v2 = v2
                    v22 := GenS3{F0: 1, F1: 42}
                    if (!(true || (false && (511 != (v1 * v1))))) {
                        switch ((("\n\t" + "hello") + "foo bar") + "foo bar") {
                        case "":
                            if (true && (("" + "foo bar") <= ("" + ("hello" + ("\n\t" + ("hello" + "foo bar")))))) {
                                v3[(("\n\t" + ("hello" + "foo bar")) + "foo bar")] = (true || (((1 >= 0) || (65535 == 1)) || true))
                                var v23 uint32 = 42
                                for fi24 := 0; fi24 < 17; fi24++ {
                                    v1 -= 1
                                    _ = fi24
                                }
                                _ = v23
                            }
                        }
                        v25 := []int{127, fi8}
                        if false {
                            for fi26 := 0; fi26 < 4; fi26++ {
                                v27 := GenS4{F0: (fi8 ^ (fi26 % (((fi26 * (fi8 | fi8))) | 1))), F1: (("" + ("" + (("foo bar" + "hello") + "foo bar"))) != (("" + "hello") + "\n\t"))}
                                _ = fi26
                                _ = v27
                            }
                        }
                        _ = v25
                    } else {
                        var v28 uint64 = 42
                        if false {
                            v1 = v1
                        }
                        _ = v28
                    }
                    _ = v22
                }
                for fi29 := 0; fi29 < 1; fi29++ {
                    v30 := GenS1{F0: 42, F1: -874, F2: ((7 >= 42) && false)}
                    fmt.Println(fi29, v1, fi8)
                    v1 -= v1
                    _ = fi29
                    _ = v30
                }
                _ = fi8
            }
        } else {
            if false {
                v1 -= (v1 / (((^(v1))) | 1))
            }
        }
    }
    fmt.Println(v2, v1)
    var v31 uint32 = 65535
    var v32 int16 = 523
    fmt.Println(v2, v32, v2)
    var v33 uint8 = 42
    if true {
        if false {
            v34 := GenS5{F0: "\n\t", F1: (v1 * v1), F2: 4294967295}
            if false {
                var v35 float32 = (-2.25 / 1.0)
                v36 := []int{533, -9223372036854775808, -462, 32767}
                v37 := GenS5{F0: "\n\t", F1: v1, F2: 9223372036854775807}
                _ = v35
                _ = v36
                _ = v37
            } else {
                v38 := []int{128, 738, -9223372036854775808}
                if len(v38) > 0 {
                    v38[0] = 4294967295
                }
                delete(v3, "")
                _ = v38
            }
            _ = v34
        } else {
            fmt.Println(v1)
        }
        if true {
            delete(v4, -1)
        } else {
            v39 := GenS6{F0: v32}
            fmt.Println(v1, v1, v32)
            var v40 uint = 4294967295
            _ = v39
            _ = v40
        }
    } else {
        v41 := GenS7{F0: 2147483647, F1: (v1 + (^(v1)))}
        _ = v41
    }
    for fi42 := 0; fi42 < 14; fi42++ {
        var v43 uint32 = 183
        var v44 bool = true
        v44 = ((42 == 42) || v44)
        _ = fi42
        _ = v43
        _ = v44
    }
    v45 := []int{}
    fmt.Println(v32)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v45:", len(v45), v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v45
    _ = v5
}
