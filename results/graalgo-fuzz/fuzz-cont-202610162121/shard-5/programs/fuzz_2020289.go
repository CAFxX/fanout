package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 string; F2 uint8 }
type GenS3 struct { F0 int32; F1 int32; F2 uint16 }
type GenS4 struct { F0 int }
func genFn0() int {
    v6 := []int{-1, 4294967295, 2147483647}
    if true {
        for fi7 := 0; fi7 < 9; fi7++ {
            for fi8 := 0; fi8 < 19; fi8++ {
                v9 := GenS1{F0: 65535}
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
        var v10 float64 = 1.5
        _ = v10
    }
    var v11 int64 = 0
    if (236 >= 255) {
        var v12 uint16 = 560
        switch 32768 {
        case 396:
            if (false || (-1 >= 0)) {
                var v13 float64 = 0.1
                _ = v13
            }
        case -9223372036854775808:
            fmt.Println(v12, v12)
        }
        _ = v12
    }
    _ = v11
    _ = v6
    return -9223372036854775808
}

func main() {
    v1 := make(map[int]int)
    var v2 float64 = float64(0.0)
    var v3 float64 = v2
    var v4 string = ""
    v5 := []uint32{}
    var v14 int32 = 1
    var v15 uint64 = 0
    v4 = v4
    if (true || (!((^(v14)) < v14))) {
        var v16 uint8 = 232
        _ = v16
    } else {
        if (!true) {
            if false {
                v17 := GenS2{F0: (v14 ^ (v14 ^ v14)), F1: v4, F2: uint8(v15)}
                v15 = v15
                fmt.Println(v15)
                _ = v17
            } else {
                switch (("foo bar" + ("" + (v4 + v4))) + ((((v4 + v4) + (v4 + (v4 + ((("foo bar" + v4) + v4) + v4)))) + "a") + "\n\t")) {
                case "foo bar":
                    v4 = (v4 + v4)
                case "hello":
                    switch 1 {
                    case 42:
                        v18 := GenS1{F0: 255}
                        _ = v18
                    default:
                        if false {
                            var v19 uint64 = 289
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            v20 := GenS3{F0: 218, F1: (-(v14)), F2: 719}
                            _ = v19
                            _ = v20
                        }
                    }
                }
                v1[-2147483648] = 32768
            }
            if false {
                v5 = append(v5, 1)
                v2 += 1.5
            }
        }
        switch -9223372036854775808 {
        case 128:
            v21 := GenS3{F0: 2147483647, F1: 343, F2: 1}
            _ = v21
        case 32767:
            v4 = "a"
        default:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        var v22 float64 = v3
        _ = v22
    }
    if false {
        var v23 int16 = -1
        var v24 int64 = 1
        _ = v23
        _ = v24
    }
    var v25 bool = (false && (!true))
    var v26 int64 = 221
    v27 := GenS2{F0: 42, F1: ("\n\t" + ""), F2: 128}
    var v28 float64 = 1.5
    var v29 uint64 = (v15 * 42)
    v15 += 2
    switch 32767 {
    case 42:
        if true {
            if v25 {
                if len(v5) > 0 {
                    v5[0] = 2
                }
                v30 := GenS1{F0: 65535}
                _ = v30
            } else {
                var v31 int8 = -128
                var v32 int16 = -870
                _ = v31
                _ = v32
            }
            if v25 {
                var v33 int = -2147483648
                _ = v33
            }
        } else {
            if (!v25) {
                switch v4 {
                case "foo bar":
                    var v34 float64 = (-2.25 + v28)
                    _ = v34
                }
                v35 := GenS4{F0: 89}
                _ = v35
            } else {
                var v36 uint8 = 42
                v26 -= v26
                _ = v36
            }
        }
    case 848:
        var v37 float32 = (0.1 - 0.1)
        _ = v37
    default:
        v38 := GenS1{F0: 65535}
        _ = v38
    }
    var v39 int8 = 0
    fmt.Println(v39)
    switch "\n\t" {
    case "hello":
        v26 = v26
    default:
        v40 := GenS2{F0: 0, F1: v4, F2: 255}
        _ = v40
    }
    v28 -= float64(3.14159)
    var v41 uint32 = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v39
    _ = v4
    _ = v41
    _ = v5
}
