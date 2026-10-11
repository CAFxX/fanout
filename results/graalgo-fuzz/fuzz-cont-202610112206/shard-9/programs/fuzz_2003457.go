package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint }
type GenS2 struct { F0 bool; F1 string; F2 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 bool; F1 int64 }
type GenS8 struct { F0 int8; F1 uint8; F2 uint8 }
type GenS9 struct { F0 string; F1 int64; F2 int32 }
type GenS10 struct { F0 uint16; F1 uint64 }
type GenS11 struct { F0 int16; F1 uint16; F2 uint64 }
func genFn0() uint {
    for fi4 := 0; fi4 < 7; fi4++ {
        var v5 bool = (127 > -128)
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: 255, F1: -2147483648, F2: 0}
    v7 := GenS2{F0: (0 >= 127), F1: "a", F2: 255}
    var v8 float32 = 3.14159
    v8 += 0.1
    _ = v6
    _ = v7
    _ = v8
    return 519
}
func genFn1() uint32 {
    v9 := GenS3{F0: 1}
    var v10 int32 = 42
    v11 := GenS1{F0: 389, F1: 42, F2: 65535}
    _ = v10
    _ = v11
    _ = v9
    return 2147483647
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = 1.5
    var v3 float32 = 3.14159
    var v12 int32 = 42
    if (450 >= 232) {
        var v13 int16 = 288
        switch "\n\t" {
        case "":
            if false {
                fmt.Println(v3, v13, v2)
                for fi14 := 0; fi14 < 12; fi14++ {
                    fmt.Println(v13)
                    v2 += (((v2 * v2) - v2) - v2)
                    _ = fi14
                }
            } else {
                if false {
                    var v15 int8 = int8(v13)
                    v1[2147483648] = false
                    for fi16 := 0; fi16 < 8; fi16++ {
                        v12 += 1
                        for fi17 := 0; fi17 < 9; fi17++ {
                            switch "\n\t" {
                            case "a":
                                switch (fi16 & (fi16 ^ fi16)) {
                                case 42:
                                    v18 := GenS4{F0: 2, F1: v13}
                                    _ = v18
                                }
                            default:
                                v19 := []int{(^(fi16))}
                                _ = v19
                            }
                            v20 := []int{}
                            fmt.Println(v13, v12)
                            fmt.Println(fi16)
                            v1[1] = false
                            _ = fi17
                            _ = v20
                        }
                        for fi21 := 0; fi21 < 2; fi21++ {
                            v22 := GenS5{F0: 1}
                            _ = fi21
                            _ = v22
                        }
                        _ = fi16
                    }
                    _ = v15
                }
                v2 -= v2
            }
        default:
            v3 = v3
        }
        for fi23 := 0; fi23 < 1; fi23++ {
            for fi24 := 0; fi24 < 8; fi24++ {
                if ((100 < 127) && (!true)) {
                    v3 -= (v3 / ((v3 - -2.25) * v3))
                } else {
                    if false {
                        var v25 uint64 = 1
                        v3 = 0.1
                        if false {
                            v25 = (v25 / ((v25) | 1))
                        }
                        _ = v25
                    } else {
                        var v26 int16 = (v13 % (((v13 ^ (v13 << ((42) & 63)))) | 1))
                        switch "hello" {
                        case "foo bar":
                            v2 += 0.0
                        case "a":
                            var v27 int32 = (^(v12))
                            _ = v27
                        }
                        var v28 string = "foo bar"
                        _ = v26
                        _ = v28
                    }
                }
                v29 := GenS6{F0: (false && false)}
                _ = fi24
                _ = v29
            }
            v30 := []int{}
            _ = fi23
            _ = v30
        }
        _ = v13
    }
    fmt.Println(v12, v3)
    v31 := GenS7{F0: (-1 >= -116), F1: -486}
    v32 := GenS8{F0: -110, F1: 2, F2: 255}
    v3 += 1.5
    if ((!true) && false) {
        if (!(!(-2147483648 != 42))) {
            for fi33 := 0; fi33 < 14; fi33++ {
                v34 := GenS9{F0: "", F1: -47, F2: (^(v12))}
                _ = fi33
                _ = v34
            }
            var v35 uint8 = 42
            _ = v35
        } else {
            var v36 uint = 2
            _ = v36
        }
    } else {
        var v37 float64 = v2
        _ = v37
    }
    var v38 int64 = -177
    if true {
        v39 := []int{-2147483648, -306, 448}
        _ = v39
    } else {
        v40 := []int{}
        if (float64(3.14159) > 0.0) {
            var v41 int16 = 1000
            if (!true) {
                v38 = v38
            }
            _ = v41
        } else {
            v38 += 1
        }
        _ = v40
    }
    v42 := v1[208]
    if v42 {
        switch 32767 {
        case 32767:
            for fi43 := 0; fi43 < 2; fi43++ {
                if v42 {
                    var v44 float64 = float64(0.0)
                    v45 := GenS10{F0: 1, F1: 1}
                    var v46 uint8 = 2
                    _ = v44
                    _ = v45
                    _ = v46
                } else {
                    v12 = (-(v12))
                }
                v47 := GenS11{F0: 0, F1: 65535, F2: 42}
                switch (fi43 - fi43) {
                case 32767:
                    fmt.Println(fi43, fi43, v38)
                case 591:
                    v48 := []int{}
                    _ = v48
                case 32768:
                    var v49 string = ""
                    _ = v49
                }
                _ = fi43
                _ = v47
            }
        case -381:
            v12 += v12
        default:
            for fi50 := 0; fi50 < 16; fi50++ {
                var v51 uint16 = 333
                _ = fi50
                _ = v51
            }
        }
        var v52 uint16 = 32767
        v53 := GenS2{F0: false, F1: "a", F2: 65535}
        _ = v52
        _ = v53
    } else {
        fmt.Println(v2, v2)
        var v54 uint8 = 42
        var v55 int16 = 1
        _ = v54
        _ = v55
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v38
    _ = v42
}
