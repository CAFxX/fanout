package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int; F1 string; F2 uint32 }
type GenS5 struct { F0 uint32; F1 int16 }
type GenS6 struct { F0 uint16; F1 int; F2 uint }
type GenS7 struct { F0 uint32 }
func genFn0(p0 uint16) uint8 {
    v7 := []int{804}
    v8 := GenS1{F0: 1, F1: "a"}
    _ = p0
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 uint8 = 38
    var v2 uint = 653
    var v3 float64 = -2.25
    v4 := []uint32{65535, 826, 4294967295}
    v5 := make(map[string]int)
    var v6 int64 = 0
    switch "a" {
    case "hello":
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 int = fi9
            var v11 int16 = 32767
            v12 := GenS1{F0: genFn0(32767), F1: ("hello" + "")}
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
    case "\n\t":
        v13 := GenS2{F0: 42}
        _ = v13
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        fmt.Println(v1)
        for fi15 := 0; fi15 < 20; fi15++ {
            switch "hello" {
            case "a":
                v1 = genFn0(191)
            case "hello":
                v16 := GenS3{F0: 2}
                _ = v16
            case "foo bar":
                v17 := GenS4{F0: (fi15 ^ fi15), F1: "hello", F2: 0}
                delete(v5, "")
                var v18 int64 = v6
                _ = v17
                _ = v18
            }
            _ = fi15
        }
        _ = fi14
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        var v20 float64 = (1.5 - float64((0.0 - float32(float64(float32(v3))))))
        _ = fi19
        _ = v20
    }
    fmt.Println(v2)
    var v21 int32 = -1
    v3 = 1.5
    var v22 int8 = -128
    v4 = append(v4, 2)
    for fi23 := 0; fi23 < 8; fi23++ {
        var v24 string = "hello"
        _ = fi23
        _ = v24
    }
    v25 := v5[((("a" + "") + "hello") + "\n\t")]
    for fi26 := 0; fi26 < 2; fi26++ {
        if false {
            v21 -= (v21 + (v21 / (((v21 * (v21 / (((v21 & (v21 & v21))) | 1)))) | 1)))
        } else {
            switch "a" {
            case "\n\t":
                delete(v5, ("" + "a"))
            case "hello":
                if true {
                    v27 := GenS5{F0: 530, F1: -1}
                    v6 = 300
                    _ = v27
                } else {
                    var v28 uint = 4294967295
                    fmt.Println(v25, v3)
                    _ = v28
                }
            default:
                switch "foo bar" {
                case "a":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case "":
                    switch v25 {
                    case 848:
                        var v29 int16 = 1000
                        _ = v29
                    default:
                        switch ("\n\t" + ("a" + "\n\t")) {
                        case "foo bar":
                            for fi30 := 0; fi30 < 5; fi30++ {
                                var v31 int8 = 100
                                if false {
                                    v2 += (v2 % ((v2) | 1))
                                    fmt.Println(v1, v22)
                                    var v32 float32 = -2.25
                                    _ = v32
                                } else {
                                    v22 = v31
                                    v6 -= (v6 * (v6 | (v6 ^ v6)))
                                    if ((1 >= 582) || false) {
                                        fmt.Println(fi30, v25)
                                        v33 := v5[""]
                                        if (1 <= 281) {
                                            v2 -= (v2 - (v2 + v2))
                                        } else {
                                            v21 += 1
                                            if true {
                                                v2 += v2
                                            } else {
                                                v22 -= (^(v31))
                                            }
                                        }
                                        v34 := GenS3{F0: 271}
                                        _ = v33
                                        _ = v34
                                    }
                                    v25 = v25
                                }
                                _ = fi30
                                _ = v31
                            }
                        case "hello":
                            v22 += v22
                        case "":
                            for fi35 := 0; fi35 < 7; fi35++ {
                                if false {
                                    if ((!(88 == 4294967295)) && false) {
                                        var v36 float32 = (1.5 + 0.0)
                                        var v37 float64 = float64((v36 / (3.14159 - v36)))
                                        _ = v36
                                        _ = v37
                                    }
                                    v21 += (^(v21))
                                    var v38 uint16 = 42
                                    v2 = (^(v2))
                                    _ = v38
                                }
                                var v39 string = "a"
                                _ = fi35
                                _ = v39
                            }
                        default:
                            fmt.Println(v22, v21, v22)
                        }
                    }
                }
            }
        }
        _ = fi26
    }
    fmt.Println(v3, v25, v25)
    switch "a" {
    case "hello":
        if false {
            v1 += 205
            var v40 int16 = 42
            _ = v40
        }
    case "\n\t":
        v21 += -2147483648
    case "foo bar":
        if (v2 <= v2) {
            v41 := GenS1{F0: v1, F1: ((("a" + "hello") + ("hello" + ("\n\t" + ("\n\t" + "a")))) + "")}
            for fi42 := 0; fi42 < 11; fi42++ {
                v2 -= v2
                v2 += (^(v2))
                switch "" {
                case "\n\t":
                    for fi43 := 0; fi43 < 12; fi43++ {
                        v44 := GenS1{F0: (v1 % ((42) | 1)), F1: "a"}
                        switch "a" {
                        case "a":
                            if true {
                                var v45 int64 = 42
                                if len(v4) > 0 {
                                    v4[0] = 476
                                }
                                _ = v45
                            }
                        case "hello":
                            var v46 int32 = (v21 ^ v21)
                            _ = v46
                        case "foo bar":
                            v47 := GenS4{F0: (v25 % (((fi42 % ((9223372036854775807) | 1))) | 1)), F1: "a", F2: 42}
                            _ = v47
                        }
                        var v48 string = "a"
                        _ = fi43
                        _ = v44
                        _ = v48
                    }
                case "hello":
                    v49 := GenS6{F0: 42, F1: 128, F2: 2}
                    _ = v49
                case "":
                    if true {
                        v50 := GenS7{F0: 687}
                        _ = v50
                    } else {
                        _ = len(v4)
                    }
                }
                _ = fi42
            }
            _ = v41
        } else {
            v2 -= 4294967295
            if true {
                var v51 uint8 = (v1 + v1)
                _ = len(v4)
                _ = v51
            }
        }
    }
    _ = len(v4)
    var v52 uint32 = 42
    v5[""] = v25
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
}
