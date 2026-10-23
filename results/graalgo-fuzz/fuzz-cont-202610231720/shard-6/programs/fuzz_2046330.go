package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 string; F1 int32; F2 int16 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 uint32 }
func genFn0() int16 {
    switch "hello" {
    case "hello":
        var v7 int64 = 42
        _ = v7
    }
    var v8 int8 = 127
    v8 -= (v8 & (^(v8)))
    v8 -= (v8 ^ v8)
    v8 += -24
    _ = v8
    return 1
}

func main() {
    var v1 float32 = float32(0.0)
    v2 := make(map[string]string)
    v3 := make(map[string]bool)
    var v4 int = 1
    v5 := []int{}
    var v6 bool = (true && (4294967295 <= 0))
    if len(v5) > 0 {
        v5[0] = 128
    }
    v6 = false
    var v9 uint64 = 4294967295
    fmt.Println(v1, v9, v6)
    var v10 string = (("foo bar" + "foo bar") + "foo bar")
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 uint32 = 2
        v4 = (-(v4))
        _ = fi11
        _ = v12
    }
    v13 := GenS1{F0: 1000, F1: v10}
    if v6 {
        if v6 {
            fmt.Println(v1)
            if (false && (false || v6)) {
                _ = len(v5)
                if ((v4 & v4) >= v4) {
                    v14 := GenS2{F0: (v10 + v10), F1: 0, F2: 0}
                    v6 = v6
                    _ = v14
                }
            } else {
                if v6 {
                    if (!v6) {
                        if v6 {
                            v9 -= (v9 ^ 0)
                        } else {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                        if (1 != 0) {
                            var v15 int8 = -12
                            _ = v15
                        } else {
                            if (-128 < -128) {
                                var v16 uint = 42
                                var v17 uint16 = 483
                                var v18 string = v10
                                _ = v16
                                _ = v17
                                _ = v18
                            } else {
                                if (65535 >= 0) {
                                    v19 := GenS3{F0: 73, F1: 837}
                                    _ = v19
                                } else {
                                    switch 32767 {
                                    case -9223372036854775808:
                                        if len(v5) > 0 {
                                            v5[0] = 42
                                        }
                                    case 32768:
                                        v2[""] = "hello"
                                    case 900:
                                        for fi20 := 0; fi20 < 7; fi20++ {
                                            var v21 uint32 = 4294967295
                                            for fi22 := 0; fi22 < 3; fi22++ {
                                                var v23 uint32 = 42
                                                fmt.Println(v9, v23, v6)
                                                fmt.Println(fi20)
                                                _ = fi22
                                                _ = v23
                                            }
                                            _ = fi20
                                            _ = v21
                                        }
                                    }
                                    fmt.Println(v10, v6, v6)
                                    if true {
                                        if v6 {
                                            var v24 int64 = -9223372036854775808
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            _ = v24
                                        }
                                    }
                                }
                                v1 += 1e10
                            }
                        }
                    }
                    if true {
                        v25 := GenS2{F0: ((v10 + (v10 + v10)) + "hello"), F1: -640, F2: 178}
                        _ = v25
                    }
                    if (1e10 <= 3.14159) {
                        if v6 {
                            v26 := GenS1{F0: 1000, F1: (v10 + ("" + (((v10 + "hello") + (v10 + (v10 + ((v10 + (v10 + "hello")) + ("hello" + v10))))) + "")))}
                            for fi27 := 0; fi27 < 7; fi27++ {
                                v4 = (-(fi27))
                                fmt.Println(fi27, v4, v1)
                                _ = fi27
                            }
                            var v28 string = "a"
                            _ = v26
                            _ = v28
                        } else {
                            var v29 int = (v4 % (((v4 % ((v4) | 1))) | 1))
                            v4 += v4
                            _ = v29
                        }
                    }
                } else {
                    v5 = append(v5, v4)
                }
                var v30 int16 = 32767
                for fi31 := 0; fi31 < 9; fi31++ {
                    if ((false && false) || ((100 < 1) && v6)) {
                        var v32 uint16 = 42
                        _ = v32
                    }
                    v10 = (((((("\n\t" + "") + "foo bar") + (v10 + (v10 + v10))) + ("hello" + "a")) + v10) + "a")
                    _ = fi31
                }
                _ = v30
            }
            if v6 {
                v4 += (v4 * v4)
            }
        } else {
            v9 -= (v9 | v9)
            v6 = v6
        }
        for fi33 := 0; fi33 < 1; fi33++ {
            if false {
                for fi34 := 0; fi34 < 3; fi34++ {
                    v35 := v3[v10]
                    v9 -= 18446744073709551615
                    _ = fi34
                    _ = v35
                }
            } else {
                v36 := GenS1{F0: -1, F1: v10}
                var v37 int32 = 907
                v38 := GenS4{F0: 9223372036854775807, F1: 150}
                _ = v36
                _ = v37
                _ = v38
            }
            for fi39 := 0; fi39 < 11; fi39++ {
                v40 := GenS5{F0: 903}
                var v41 uint16 = 53
                _ = fi39
                _ = v40
                _ = v41
            }
            _ = fi33
        }
    } else {
        for fi42 := 0; fi42 < 20; fi42++ {
            if (!false) {
                if false {
                    var v43 uint32 = 161
                    _ = v43
                } else {
                    var v44 uint16 = 804
                    v6 = v6
                    _ = v44
                }
            } else {
                v4 = (v4 * 2147483647)
                v45 := v3[v10]
                var v46 uint16 = 333
                _ = v45
                _ = v46
            }
            _ = fi42
        }
        if v6 {
            fmt.Println(v9, v10, v9)
        } else {
            fmt.Println(v6, v9, v10)
            var v47 int64 = 42
            _ = v47
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
