package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8; F1 uint32 }

func main() {
    var v1 int = -1
    var v2 string = ""
    v3 := make(map[uint32]uint32)
    v4 := []int{v1}
    var v5 int8 = 127
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 float64 = float64(3.14159)
        if true {
            for fi8 := 0; fi8 < 12; fi8++ {
                if true {
                    if true {
                        if (1 >= 42) {
                            fmt.Println(v1, v1, v7)
                            v9 := v3[2147483648]
                            _ = v9
                        } else {
                            _ = len(v4)
                            switch v2 {
                            case "foo bar":
                                v10 := GenS1{F0: 2, F1: -1, F2: v5}
                                _ = v10
                            default:
                                if (1 == 23) {
                                    v1 = fi8
                                }
                            }
                        }
                        if false {
                            var v11 string = v2
                            _ = v11
                        }
                        v5 += (v5 + 0)
                        var v12 bool = false
                        _ = v12
                    }
                    var v13 uint8 = 139
                    _ = v13
                } else {
                    if true {
                        v7 = 1.5
                    }
                    v5 = (^(v5))
                }
                v7 += (0.1 - (0.1 * (v7 / 0.1)))
                if true {
                    var v14 int32 = 894
                    v15 := v3[900]
                    if (!(!false)) {
                        if false {
                            fmt.Println(v1, fi8, v7)
                            v14 -= 572
                        }
                    } else {
                        v16 := GenS2{F0: 603}
                        if (v14 < (v14 * (v14 ^ v14))) {
                            switch ("\n\t" + v2) {
                            case "a":
                                fmt.Println(v15, v14)
                            case "foo bar":
                                v3[2147483647] = 2
                            default:
                                var v17 bool = (!true)
                                var v18 bool = (v1 > v1)
                                var v19 int64 = -9223372036854775808
                                _ = v17
                                _ = v18
                                _ = v19
                            }
                        }
                        _ = len(v4)
                        v20 := GenS1{F0: 205, F1: -1, F2: v5}
                        switch v1 {
                        case -1:
                            v5 -= v5
                            v1 += int(v14)
                            var v21 uint8 = 8
                            _ = v21
                        default:
                            v1 = (fi8 & (v1 & (-(fi6))))
                        }
                        _ = v16
                        _ = v20
                    }
                    _ = v14
                    _ = v15
                } else {
                    if (0 > 1) {
                        v2 = v2
                    }
                    v5 = v5
                    v5 -= 42
                }
                _ = fi8
            }
        }
        _ = fi6
        _ = v7
    }
    v5 -= (v5 ^ -1)
    fmt.Println(v2, v2)
    v22 := v3[584]
    if (!(true && false)) {
        if true {
            v23 := GenS2{F0: 893}
            _ = v23
        } else {
            v1 += -577
            v2 = ""
        }
        if true {
            var v24 float64 = float64(1e10)
            var v25 uint = 1
            v1 = -607
            _ = v24
            _ = v25
        } else {
            if (!(-2147483648 != 2147483647)) {
                delete(v3, v22)
            }
        }
        v26 := GenS2{F0: 42}
        _ = v26
    }
    var v27 int8 = 0
    var v28 uint8 = 0
    v3[(v22 / ((v22) | 1))] = v22
    for fi29 := 0; fi29 < 18; fi29++ {
        v30 := GenS2{F0: 36}
        if len(v4) > 0 {
            v4[0] = fi29
        }
        _ = fi29
        _ = v30
    }
    switch "a" {
    case "a":
        fmt.Println(v22, v2)
    case "hello":
        var v31 uint32 = v22
        _ = v31
    case "\n\t":
        switch "a" {
        case "\n\t":
            for fi32 := 0; fi32 < 19; fi32++ {
                _ = len(v3)
                v2 = (v2 + "a")
                _ = fi32
            }
        case "hello":
            v33 := GenS3{F0: (v28 * 42), F1: (v22 ^ v22)}
            _ = v33
        }
    default:
        for fi34 := 0; fi34 < 8; fi34++ {
            v28 -= (v28 | v28)
            switch v2 {
            case "foo bar":
                var v35 string = v2
                _ = v35
            case "":
                var v36 int64 = -822
                _ = v36
            default:
                fmt.Println(v5)
            }
            _ = fi34
        }
    }
    v1 -= v1
    var v37 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v27
    _ = v28
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
