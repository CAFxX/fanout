package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 int8; F2 uint16 }
type GenS3 struct { F0 int8; F1 uint }

func main() {
    v1 := make(map[int]int)
    var v2 int64 = 534
    var v3 int16 = -32768
    var v4 bool = ((float32(((3.14159 / -2.25) - 0.0)) != 0.1) && false)
    var v5 float32 = -2.25
    var v6 uint = 514
    v7 := v1[447]
    if v4 {
        for fi8 := 0; fi8 < 11; fi8++ {
            for fi9 := 0; fi9 < 2; fi9++ {
                fmt.Println(v4, v2, v7)
                _ = fi9
            }
            _ = len(v1)
            _ = fi8
        }
        if (0 > -128) {
            var v10 float64 = 0.1
            v11 := []int{}
            _ = v10
            _ = v11
        }
        fmt.Println(v3)
    }
    v3 += v3
    v12 := GenS1{F0: v2}
    for fi13 := 0; fi13 < 9; fi13++ {
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 int64 = v2
            if v4 {
                var v16 int32 = int32(fi13)
                v15 = (-(v15))
                v3 += v3
                _ = v16
            }
            if (!v4) {
                v6 -= 2
                v6 = (^(v6))
            } else {
                v1[32767] = (-(fi13))
            }
            _ = fi14
            _ = v15
        }
        v1[fi13] = 240
        var v17 uint64 = 42
        _ = fi13
        _ = v17
    }
    v7 += v7
    var v18 uint = (v6 * (v6 ^ 299))
    for fi19 := 0; fi19 < 8; fi19++ {
        var v20 int64 = (-(v2))
        fmt.Println(v6, v6)
        var v21 uint16 = 255
        _ = fi19
        _ = v20
        _ = v21
    }
    v22 := GenS2{F0: v4, F1: 42, F2: 32767}
    if v4 {
        v3 = 47
    }
    switch (v7 / (((^(v7))) | 1)) {
    case 2147483647:
        var v23 uint8 = 97
        _ = v23
    default:
        v24 := []int{-324, int(v3), 439}
        _ = v24
    }
    var v25 bool = v4
    if v25 {
        switch 42 {
        case 32767:
            var v26 float64 = 1e10
            _ = v26
        case 4294967295:
            for fi27 := 0; fi27 < 2; fi27++ {
                for fi28 := 0; fi28 < 1; fi28++ {
                    var v29 uint8 = 127
                    if (-558 == (v7 ^ 128)) {
                        switch (fi28 & (fi27 << ((294) & 63))) {
                        case 401:
                            switch "" {
                            case "\n\t":
                                var v30 int = fi27
                                _ = v30
                            case "":
                                var v31 string = "a"
                                _ = v31
                            case "foo bar":
                                var v32 uint32 = 861
                                _ = v32
                            default:
                                if false {
                                    var v33 int32 = 41
                                    v34 := GenS2{F0: v4, F1: -9, F2: 255}
                                    _ = v33
                                    _ = v34
                                } else {
                                    if (!(v5 >= 3.14159)) {
                                        switch fi28 {
                                        case -1:
                                            var v35 uint32 = 42
                                            _ = v35
                                        case -9223372036854775808:
                                            fmt.Println(v18, v3, v29)
                                        }
                                        v36 := []int{(v7 | 212), 9223372036854775807}
                                        v5 = float32(0.1)
                                        _ = v36
                                    } else {
                                        fmt.Println(v5)
                                        fmt.Println(v29, v7, fi27)
                                        var v37 uint8 = 229
                                        _ = v37
                                    }
                                }
                            }
                        }
                        v18 += (v6 / ((v18) | 1))
                        if (1 == -1) {
                            v38 := GenS1{F0: (-(v2))}
                            var v39 int16 = int16(v2)
                            v40 := GenS1{F0: v2}
                            _ = v38
                            _ = v39
                            _ = v40
                        }
                        v41 := v1[128]
                        v42 := []int{(-(fi27))}
                        _ = v41
                        _ = v42
                    }
                    _ = fi28
                    _ = v29
                }
                switch int(v18) {
                case -9223372036854775808:
                    var v43 float64 = -2.25
                    _ = v43
                }
                _ = fi27
            }
        }
        v44 := GenS2{F0: v25, F1: -109, F2: 42}
        _ = v44
    }
    switch "a" {
    case "":
        if ((-(v7)) > (v7 * (v7 % ((9223372036854775807) | 1)))) {
            v7 += (v7 - (v7 - (v7 ^ v7)))
            if (3.14159 <= 1e10) {
                switch "a" {
                case "":
                    switch v7 {
                    case 2147483647:
                        switch (("a" + "a") + "foo bar") {
                        case "":
                            var v45 uint = 4294967295
                            _ = v45
                        case "\n\t":
                            var v46 bool = (true && (v7 == 32768))
                            _ = v46
                        case "foo bar":
                            v47 := v1[(v7 >> ((v6) & 63))]
                            _ = v47
                        default:
                            var v48 uint32 = 436
                            _ = v48
                        }
                    }
                case "hello":
                    var v49 float32 = (1e10 + v5)
                    _ = v49
                default:
                    if v4 {
                        fmt.Println(v18, v7, v6)
                    }
                }
            }
            var v50 int32 = 0
            _ = v50
        } else {
            v51 := GenS3{F0: -14, F1: 2}
            _ = v51
        }
    case "\n\t":
        var v52 int8 = 32
        _ = v52
    default:
        v5 += v5
    }
    var v53 uint8 = 128
    if (v4 && v4) {
        fmt.Println(v5, v7, v5)
        fmt.Println(v5, v7, v6)
    } else {
        v7 += 2147483647
        var v54 uint32 = 0
        _ = v54
    }
    fmt.Println(v25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v6
    _ = v7
}
