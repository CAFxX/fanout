package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint8; F2 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8; F1 uint64 }
type GenS5 struct { F0 int64; F1 uint; F2 uint }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 bool; F1 string }
type GenS9 struct { F0 int16 }
func genFn0(p0 int64) int8 {
    fmt.Println(p0)
    if (0 <= 2) {
        for fi5 := 0; fi5 < 8; fi5++ {
            if ((354 <= 0) && ((18446744073709551615 != 1) || false)) {
                var v6 float32 = (0.1 * 1e10)
                _ = v6
            }
            _ = fi5
        }
    }
    fmt.Println(p0)
    fmt.Println(p0)
    p0 = p0
    _ = p0
    return -128
}

func main() {
    v1 := make(map[string]bool)
    var v2 string = "hello"
    var v3 bool = false
    var v4 uint8 = 1
    v7 := GenS1{F0: 263, F1: 2}
    for fi8 := 0; fi8 < 18; fi8++ {
        switch "\n\t" {
        case "hello":
            switch "" {
            case "a":
                v2 = (("a" + (v2 + v2)) + v2)
            case "\n\t":
                if ((0.1 < float64(0.1)) && false) {
                    fmt.Println(v4)
                    var v9 uint8 = v4
                    v10 := GenS2{F0: 747, F1: v4, F2: fi8}
                    _ = v10
                    _ = v9
                }
            default:
                var v11 int16 = -1
                _ = v11
            }
        }
        if (true || v3) {
            for fi12 := 0; fi12 < 10; fi12++ {
                fmt.Println(fi12)
                v13 := GenS3{F0: 680}
                v3 = (!v3)
                _ = fi12
                _ = v13
            }
            for fi14 := 0; fi14 < 19; fi14++ {
                if false {
                    if (3.14159 <= ((float64(0.1) / 1e10) + ((1e10 + 1.5) - -2.25))) {
                        var v15 uint64 = 0
                        if (!v3) {
                            if false {
                                v16 := GenS4{F0: -80, F1: v15}
                                _ = v16
                            } else {
                                v15 += 18446744073709551615
                                if true {
                                    if false {
                                        if v3 {
                                            v1[("" + v2)] = v3
                                            switch v2 {
                                            case "hello":
                                                switch v2 {
                                                case "":
                                                    v15 = v15
                                                }
                                            case "foo bar":
                                                if (v15 < 443) {
                                                    var v17 bool = (false || (1 >= 1))
                                                    if v17 {
                                                        v18 := GenS3{F0: v15}
                                                        var v19 uint8 = (v4 ^ 0)
                                                        _ = v18
                                                        _ = v19
                                                    } else {
                                                        v20 := []int{}
                                                        if (!(((1e10 / 1.5) * 1e10) != 3.14159)) {
                                                            var v21 int64 = 0
                                                            _ = v21
                                                        } else {
                                                            fmt.Println(fi14, v2)
                                                            var v22 int = 128
                                                            _ = v22
                                                        }
                                                        fmt.Println(v2)
                                                        _ = v20
                                                    }
                                                    var v23 uint16 = 0
                                                    _ = v17
                                                    _ = v23
                                                } else {
                                                    var v24 int16 = -32768
                                                    _ = v24
                                                }
                                            default:
                                                fmt.Println(v3, v4)
                                            }
                                        }
                                        var v25 uint16 = 2
                                        var v26 int32 = 744
                                        v25 -= 319
                                        _ = v25
                                        _ = v26
                                    }
                                    v27 := GenS3{F0: (v15 & v15)}
                                    v15 -= (v15 & v15)
                                    var v28 float64 = 1e10
                                    v29 := GenS5{F0: -1, F1: 65535, F2: 1}
                                    var v30 bool = v3
                                    v31 := []int{-2147483648, 751, 32767}
                                    _ = len(v31)
                                    v32 := GenS6{F0: (42 > 255)}
                                    var v33 int64 = 42
                                    _ = v27
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                    _ = v31
                                    _ = v32
                                    _ = v33
                                }
                                if ((!v3) || (0 >= 420)) {
                                    if (!v3) {
                                        fmt.Println(v3)
                                        switch fi14 {
                                        case 0:
                                            v34 := []int{fi14, fi8, fi14, (fi8 / ((-825) | 1))}
                                            _ = v34
                                        case 42:
                                            if (107 >= 42) {
                                                v1[(v2 + (((("foo bar" + ((((v2 + v2) + "") + (v2 + v2)) + v2)) + v2) + v2) + ""))] = true
                                                v2 = "foo bar"
                                                var v35 uint16 = 277
                                                _ = v35
                                            }
                                        }
                                    }
                                }
                            }
                            var v36 bool = (!v3)
                            v37 := GenS7{F0: 255}
                            _ = v36
                            _ = v37
                        }
                        _ = v15
                    }
                }
                _ = fi14
            }
            v4 = 1
        } else {
            var v38 string = (v2 + "\n\t")
            fmt.Println(fi8, v38, fi8)
            _ = v38
        }
        _ = fi8
    }
    v39 := GenS8{F0: v3, F1: (v2 + "foo bar")}
    if true {
        v40 := v1["hello"]
        _ = v40
    }
    if (((v2 + "hello") + (v2 + v2)) > v2) {
        for fi41 := 0; fi41 < 11; fi41++ {
            v2 = (v2 + v2)
            v3 = true
            var v42 int16 = -82
            _ = fi41
            _ = v42
        }
    } else {
        v1["\n\t"] = (3.14159 > 1e10)
    }
    switch v2 {
    case "foo bar":
        var v43 float64 = 1.5
        _ = v43
    case "":
        var v44 uint = 255
        _ = v44
    case "hello":
        v45 := []int{0}
        _ = v45
    default:
        if (42 == 0) {
            v46 := GenS9{F0: 236}
            fmt.Println(v3, v3, v2)
            v47 := []int{32, 32767}
            _ = v46
            _ = v47
        } else {
            v48 := []int{642}
            var v49 uint = 999
            _ = v48
            _ = v49
        }
    }
    v4 = (v4 | v4)
    v2 = v2
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v7
}
