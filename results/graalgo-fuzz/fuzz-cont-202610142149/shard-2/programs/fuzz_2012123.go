package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64; F1 bool; F2 bool }
type GenS6 struct { F0 string; F1 uint32 }
type GenS7 struct { F0 uint8; F1 int }
type GenS8 struct { F0 int32; F1 uint16 }
func genFn0() uint32 {
    var v6 uint = 65535
    if (79 >= 140) {
        if (!((float32(0.1) == float32((-2.25 / -2.25))) && false)) {
            if false {
                if false {
                    if (32767 >= 0) {
                        if (42 < 241) {
                            v7 := GenS1{F0: 1, F1: -9223372036854775808, F2: 128}
                            v6 = 18446744073709551615
                            _ = v7
                        }
                    } else {
                        var v8 uint32 = 0
                        v9 := GenS1{F0: 0, F1: -810, F2: 1}
                        _ = v8
                        _ = v9
                    }
                    var v10 float32 = 0.0
                    v11 := GenS2{F0: 42}
                    _ = v10
                    _ = v11
                }
                v6 = v6
                switch 128 {
                case -3:
                    v12 := GenS3{F0: 1, F1: 1}
                    _ = v12
                case 42:
                    v13 := []int{-231, -9223372036854775808}
                    _ = v13
                case 964:
                    var v14 int16 = -32768
                    _ = v14
                default:
                    var v15 float64 = -2.25
                    _ = v15
                }
            } else {
                v16 := []int{-1, -1}
                v6 += v6
                v6 = (v6 / ((v6) | 1))
                _ = v16
            }
        }
        for fi17 := 0; fi17 < 9; fi17++ {
            v6 = (^(v6))
            _ = fi17
        }
    } else {
        v6 -= (v6 % ((v6) | 1))
        var v18 uint64 = 1
        _ = v18
    }
    if true {
        v19 := []int{}
        for fi20 := 0; fi20 < 6; fi20++ {
            if (2 > 0) {
                for fi21 := 0; fi21 < 1; fi21++ {
                    v6 = v6
                    v22 := GenS3{F0: 1, F1: 310}
                    _ = fi21
                    _ = v22
                }
                fmt.Println(v6)
                v23 := GenS4{F0: (("hello" + "hello") + (("a" + "foo bar") + ""))}
                _ = v23
            } else {
                fmt.Println(v6)
                if false {
                    if (!false) {
                        v19 = append(v19, 127)
                    } else {
                        for fi24 := 0; fi24 < 1; fi24++ {
                            var v25 uint = (v6 >> ((v6) & 63))
                            v26 := GenS5{F0: 81, F1: false, F2: false}
                            if (true && false) {
                                if false {
                                    v25 -= v6
                                    switch ("\n\t" + "") {
                                    case "hello":
                                        switch "a" {
                                        case "foo bar":
                                            var v27 uint16 = 78
                                            _ = v27
                                        case "hello":
                                            var v28 int64 = -1
                                            fmt.Println(v6)
                                            _ = v28
                                        case "":
                                            if (!false) {
                                                fmt.Println(fi20, fi24, v6)
                                                v6 += 4294967295
                                            }
                                        }
                                    }
                                }
                            } else {
                                if (255 != 128) {
                                    v6 = (v25 + (v6 * (v25 + (v25 << ((42) & 63)))))
                                } else {
                                    v6 -= (^(v25))
                                    if true {
                                        v6 = v6
                                        v6 = (v6 % ((1) | 1))
                                        fmt.Println(v6, fi20, fi24)
                                    } else {
                                        var v29 uint8 = 28
                                        _ = v29
                                    }
                                }
                                var v30 uint8 = 128
                                var v31 float64 = 1e10
                                _ = v30
                                _ = v31
                            }
                            if true {
                                var v32 float64 = (3.14159 / 1.0)
                                if false {
                                    v33 := GenS3{F0: -380, F1: 42}
                                    v6 = (v25 / (((v25 / (((v6 + v6)) | 1))) | 1))
                                    _ = v33
                                } else {
                                    if true {
                                        if ("a" != "hello") {
                                            var v34 int8 = 26
                                            _ = v34
                                        }
                                        v6 = 42
                                    }
                                    var v35 int16 = 1000
                                    _ = v35
                                }
                                if true {
                                    var v36 int64 = 278
                                    _ = v36
                                } else {
                                    v37 := GenS6{F0: ("foo bar" + "\n\t"), F1: 478}
                                    _ = v37
                                }
                                _ = v32
                            }
                            fmt.Println(v6, v6, fi24)
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                    }
                    fmt.Println(v6, v6)
                    v38 := GenS4{F0: "foo bar"}
                    _ = v38
                } else {
                    if len(v19) > 0 {
                        v19[0] = fi20
                    }
                    _ = len(v19)
                    v6 -= v6
                }
            }
            _ = fi20
        }
        _ = v19
    }
    _ = v6
    return 98
}
func genFn1() uint {
    for fi39 := 0; fi39 < 4; fi39++ {
        var v40 string = "\n\t"
        _ = fi39
        _ = v40
    }
    var v41 uint = 1
    _ = v41
    return v41
}

func main() {
    var v1 float64 = (1.5 - (0.0 * ((1e10 * ((1e10 + (1e10 - -2.25)) - (float64(float32(3.14159)) * (0.1 * 1e10)))) / 1.0)))
    v2 := make(map[string]int)
    var v3 int16 = -1
    v4 := []string{"", ("" + ""), (((("a" + (("\n\t" + "foo bar") + ("" + ""))) + "\n\t") + "\n\t") + "a")}
    var v5 uint64 = 973
    var v42 bool = true
    v42 = v42
    v2["\n\t"] = 42
    for fi43 := 0; fi43 < 16; fi43++ {
        var v44 uint32 = 42
        var v45 int8 = 42
        _ = fi43
        _ = v44
        _ = v45
    }
    for fi46 := 0; fi46 < 1; fi46++ {
        v47 := GenS7{F0: 42, F1: 1}
        var v48 uint = 65535
        _ = fi46
        _ = v47
        _ = v48
    }
    v49 := GenS8{F0: 2147483647, F1: 1}
    v4 = append(v4, "")
    v1 = 0.0
    v50 := GenS7{F0: 0, F1: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v49
    _ = v5
    _ = v50
}
