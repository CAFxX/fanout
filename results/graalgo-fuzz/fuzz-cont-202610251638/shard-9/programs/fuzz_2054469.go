package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 string; F2 uint16 }
type GenS3 struct { F0 bool }

func main() {
    var v1 float64 = 1.5
    var v2 uint = 42
    var v3 uint8 = 214
    var v4 int = -9223372036854775808
    v5 := make(map[string]string)
    v6 := make(map[string]bool)
    v7 := GenS1{F0: 1}
    v8 := GenS2{F0: 65535, F1: "hello", F2: 305}
    var v9 float32 = 0.0
    if (!(32767 != 42)) {
        var v10 uint = (v2 % ((v2) | 1))
        _ = v10
    } else {
        fmt.Println(v2, v2)
    }
    v5[("foo bar" + "\n\t")] = "hello"
    var v11 float32 = ((1.5 * float32(v1)) * ((0.0 - v9) + v9))
    switch (v4 & -9223372036854775808) {
    case 2147483647:
        for fi12 := 0; fi12 < 14; fi12++ {
            for fi13 := 0; fi13 < 2; fi13++ {
                var v14 uint16 = 0
                _ = fi13
                _ = v14
            }
            fmt.Println(v3, v4)
            v11 -= float32(float64(v11))
            _ = fi12
        }
    default:
        var v15 uint16 = 255
        _ = v15
    }
    var v16 float32 = 1e10
    switch v4 {
    case 1:
        v17 := []int{}
        _ = v17
    case 4294967295:
        var v18 uint16 = 920
        _ = v18
    case 2147483647:
        switch "a" {
        case "":
            switch "hello" {
            case "a":
                v19 := []int{32768, (v4 % ((v4) | 1)), 9223372036854775807, 832}
                _ = v19
            case "\n\t":
                var v20 uint64 = 751
                _ = v20
            case "":
                if true {
                    v4 -= (v4 | -258)
                    switch (^(v4)) {
                    case -2147483648:
                        v4 += 1
                    case 32767:
                        v2 = (v2 & (^(v2)))
                    case 511:
                        v2 += (v2 & (^(v2)))
                    default:
                        v21 := []int{(v4 + v4), (^(v4)), v4, 42}
                        _ = v21
                    }
                } else {
                    v22 := GenS3{F0: (!(2147483648 >= 0))}
                    _ = v22
                }
            }
        case "a":
            switch 4294967295 {
            case -9223372036854775808:
                for fi23 := 0; fi23 < 4; fi23++ {
                    if false {
                        var v24 string = "hello"
                        if (true && (65535 >= 2)) {
                            fmt.Println(v3)
                            fmt.Println(v1, v24)
                            v3 = (v3 / ((v3) | 1))
                        }
                        _ = v24
                    } else {
                        var v25 int8 = 127
                        switch (fi23 % ((839) | 1)) {
                        case 9223372036854775807:
                            v26 := GenS3{F0: false}
                            _ = v26
                        case -159:
                            if ((0 <= 42) || true) {
                                switch fi23 {
                                case 0:
                                    var v27 string = "a"
                                    _ = v27
                                case 210:
                                    v4 = int(v2)
                                case 577:
                                    _ = len(v6)
                                default:
                                    v2 -= 18446744073709551615
                                }
                            } else {
                                if (((0.1 - (float64(v9) + 1.5)) / v1) >= (v1 / (v1 - v1))) {
                                    if (!false) {
                                        v9 = v9
                                    }
                                    for fi28 := 0; fi28 < 1; fi28++ {
                                        var v29 float64 = (v1 - v1)
                                        if ((v4 * (fi23 | (^(fi28)))) == 917) {
                                            v30 := []int{v4, (fi28 ^ (^(fi23))), (fi23 ^ -1), (fi28 % ((0) | 1))}
                                            if (v11 > v16) {
                                                fmt.Println(v25, v3)
                                                fmt.Println(v2, v4)
                                            }
                                            _ = v30
                                        }
                                        _ = fi28
                                        _ = v29
                                    }
                                    fmt.Println(v4)
                                }
                                v31 := []int{}
                                fmt.Println(v1, v4)
                                _ = v31
                            }
                        case 42:
                            var v32 float64 = float64(v16)
                            _ = v32
                        default:
                            v33 := GenS1{F0: 2}
                            _ = v33
                        }
                        _ = v25
                    }
                    var v34 uint8 = 0
                    v9 -= (3.14159 + 0.1)
                    _ = fi23
                    _ = v34
                }
            case 32767:
                v35 := GenS3{F0: ((v4 != (v4 ^ (v4 & -145))) && false)}
                _ = v35
            case 32768:
                var v36 float64 = v1
                _ = v36
            default:
                v3 += (v3 + 128)
            }
        }
    }
    if (v1 < (0.0 + v1)) {
        var v37 int8 = 1
        switch "\n\t" {
        case "hello":
            var v38 uint8 = 42
            _ = v38
        case "a":
            v3 -= v3
        case "foo bar":
            v37 = (^(v37))
        default:
            var v39 float64 = v1
            _ = v39
        }
        _ = v37
    } else {
        for fi40 := 0; fi40 < 18; fi40++ {
            var v41 float32 = 0.1
            var v42 uint16 = 234
            _ = fi40
            _ = v41
            _ = v42
        }
    }
    fmt.Println(v4, v2)
    v5[""] = (("hello" + "\n\t") + "hello")
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
