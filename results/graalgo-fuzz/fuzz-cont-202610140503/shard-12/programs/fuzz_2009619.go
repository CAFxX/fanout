package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int8; F1 bool; F2 string }
type GenS3 struct { F0 bool }
func genFn0(p0 int64, p1 int16) int {
    var v7 uint16 = 65535
    p1 += -316
    v8 := GenS1{F0: (("foo bar" + "hello") + ("foo bar" + "hello")), F1: 65535}
    var v9 string = ""
    v10 := []int{588, 32767, 1}
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return -2147483648
}

func main() {
    var v1 bool = false
    var v2 float64 = -2.25
    var v3 uint8 = 255
    var v4 float32 = 0.0
    var v5 int8 = -128
    var v6 int8 = 42
    if v1 {
        if v1 {
            switch -305 {
            case -148:
                fmt.Println(v4)
            }
            if v1 {
                v6 += v6
                var v11 int8 = (v5 ^ -128)
                _ = v11
            }
            switch "hello" {
            case "foo bar":
                switch 210 {
                case -336:
                    if (v1 || v1) {
                        for fi12 := 0; fi12 < 3; fi12++ {
                            fmt.Println(v4)
                            _ = fi12
                        }
                        var v13 uint64 = 811
                        for fi14 := 0; fi14 < 12; fi14++ {
                            v15 := []int{}
                            v2 = 1.5
                            if len(v15) > 0 {
                                v15[0] = -323
                            }
                            _ = fi14
                            _ = v15
                        }
                        _ = v13
                    }
                case 262:
                    v4 = ((v4 + v4) * v4)
                }
            case "a":
                var v16 int8 = (v6 >> ((42) & 63))
                _ = v16
            default:
                v17 := GenS2{F0: -115, F1: ((42 != 1) || false), F2: "foo bar"}
                _ = v17
            }
        }
    } else {
        var v18 float32 = 0.1
        fmt.Println(v4, v3)
        if false {
            v19 := GenS2{F0: 42, F1: v1, F2: (("foo bar" + "\n\t") + "foo bar")}
            _ = v19
        } else {
            v2 -= v2
            var v20 int32 = 42
            switch 0 {
            case 589:
                v21 := []int{-1, 2147483647}
                _ = v21
            default:
                v4 = 1e10
            }
            _ = v20
        }
        _ = v18
    }
    v22 := []int{589, 32767, 9223372036854775807}
    for fi23 := 0; fi23 < 17; fi23++ {
        for fi24 := 0; fi24 < 18; fi24++ {
            _ = len(v22)
            _ = fi24
        }
        for fi25 := 0; fi25 < 17; fi25++ {
            var v26 bool = false
            v5 = v5
            _ = len(v22)
            _ = fi25
            _ = v26
        }
        if v1 {
            var v27 uint = 42
            v2 = v2
            for fi28 := 0; fi28 < 17; fi28++ {
                if v1 {
                    _ = len(v22)
                    _ = len(v22)
                }
                v5 -= 1
                _ = fi28
            }
            _ = v27
        } else {
            v6 = v5
            if ((-(fi23)) < (fi23 - (fi23 ^ fi23))) {
                if v1 {
                    if v1 {
                        for fi29 := 0; fi29 < 8; fi29++ {
                            if false {
                                v22 = append(v22, (fi23 - fi23))
                                switch "\n\t" {
                                case "foo bar":
                                    var v30 uint8 = (v3 | v3)
                                    _ = v30
                                case "hello":
                                    if true {
                                        var v31 int16 = 0
                                        if (fi23 > -1) {
                                            var v32 int8 = v5
                                            if v1 {
                                                v6 = (v5 % ((v32) | 1))
                                            } else {
                                                var v33 float64 = v2
                                                _ = v33
                                            }
                                            _ = v32
                                        }
                                        v22 = append(v22, (fi29 & 2147483647))
                                        _ = v31
                                    }
                                case "a":
                                    var v34 int32 = -923
                                    _ = v34
                                }
                                var v35 int16 = 1000
                                _ = v35
                            }
                            v36 := GenS1{F0: (("\n\t" + ("foo bar" + ("hello" + "a"))) + ("\n\t" + "foo bar")), F1: 42}
                            _ = fi29
                            _ = v36
                        }
                    } else {
                        for fi37 := 0; fi37 < 1; fi37++ {
                            if (-148 > -2147483648) {
                                var v38 uint32 = 4294967295
                                switch "" {
                                case "foo bar":
                                    var v39 float64 = ((v2 - v2) - ((v2 * 1.5) * 0.1))
                                    _ = v39
                                default:
                                    var v40 float32 = 0.1
                                    _ = v40
                                }
                                v41 := GenS3{F0: (500 >= 18446744073709551615)}
                                _ = v38
                                _ = v41
                            }
                            v2 = (v2 - ((v2 * 1.5) + v2))
                            _ = fi37
                        }
                        if len(v22) > 0 {
                            v22[0] = fi23
                        }
                        fmt.Println(v6)
                    }
                }
                var v42 int8 = v5
                _ = v42
            }
        }
        _ = fi23
    }
    v5 += v6
    var v43 uint = 0
    var v44 uint32 = 814
    v45 := GenS3{F0: ("a" < ("a" + ""))}
    v44 = (v44 + uint32(v43))
    var v46 int32 = 42
    fmt.Println(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
