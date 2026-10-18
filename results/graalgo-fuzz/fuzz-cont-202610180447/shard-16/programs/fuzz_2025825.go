package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int64; F1 uint64; F2 uint8 }
func genFn0() bool {
    if false {
        var v6 int = 2147483647
        _ = v6
    }
    var v7 float32 = 1e10
    var v8 float64 = 3.14159
    var v9 string = "foo bar"
    _ = v7
    _ = v8
    _ = v9
    return true
}
func genFn1(p0 int32, p1 uint8, p2 uint32) uint {
    if ("hello" <= "") {
        v10 := []int{-358, 2147483647}
        switch 32767 {
        case 9223372036854775807:
            if (4294967295 != 4294967295) {
                p1 = (p1 >> ((74) & 63))
            } else {
                v11 := GenS1{F0: 127, F1: (p2 != (p2 >> ((2) & 63)))}
                if len(v10) > 0 {
                    v10[0] = -751
                }
                _ = v11
            }
        case 4294967295:
            for fi12 := 0; fi12 < 10; fi12++ {
                switch fi12 {
                case 127:
                    var v13 float64 = 1.5
                    _ = v13
                case -9223372036854775808:
                    switch (fi12 - -187) {
                    case -1:
                        p1 = p1
                    }
                }
                _ = fi12
            }
        default:
            p2 = p2
        }
        fmt.Println(p0)
        _ = v10
    } else {
        p0 = p0
        var v14 int8 = 100
        _ = v14
    }
    var v15 float64 = 0.1
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 60
}

func main() {
    var v1 int16 = 95
    var v2 string = "hello"
    v3 := make(map[int]bool)
    var v4 float64 = -2.25
    v5 := make(map[int]int)
    if false {
        for fi16 := 0; fi16 < 15; fi16++ {
            v1 -= -1
            _ = fi16
        }
        v2 = (v2 + ("" + ""))
    }
    fmt.Println(v1, v4, v4)
    var v17 int64 = 205
    var v18 bool = false
    var v19 int8 = 127
    fmt.Println(v1, v18)
    var v20 int16 = -1
    fmt.Println(v4)
    var v21 float64 = (v4 * v4)
    var v22 bool = (float32(0.0) <= 0.1)
    if true {
        for fi23 := 0; fi23 < 2; fi23++ {
            v24 := []int{(fi23 / ((fi23) | 1)), 4294967295, 383, -9223372036854775808}
            fmt.Println(fi23, v19, v4)
            v25 := GenS1{F0: (v19 + v19), F1: v22}
            _ = fi23
            _ = v24
            _ = v25
        }
        v4 -= (float64(float32(v4)) / ((float64(1e10) / v4) - v21))
    }
    switch v2 {
    case "":
        v21 += (float64(0.1) + v4)
    }
    v26 := []int{}
    var v27 int32 = 2147483647
    if len(v26) > 0 {
        v26[0] = 1
    }
    fmt.Println(v27, v18, v1)
    v28 := GenS1{F0: int8(v1), F1: v22}
    for fi29 := 0; fi29 < 20; fi29++ {
        fmt.Println(fi29)
        if (8 >= 127) {
            switch 127 {
            case 188:
                fmt.Println(fi29)
            case 9223372036854775807:
                var v30 int16 = 777
                _ = v30
            case -1:
                if v18 {
                    switch (fi29 / ((int(v19)) | 1)) {
                    case 4294967295:
                        v27 = (v27 | (v27 & v27))
                    case -336:
                        var v31 uint16 = 811
                        _ = v31
                    case 128:
                        var v32 int32 = v27
                        _ = v32
                    default:
                        switch 2147483647 {
                        case 127:
                            var v33 float64 = -2.25
                            _ = v33
                        case -347:
                            var v34 int = -275
                            _ = v34
                        case 128:
                            switch "foo bar" {
                            case "a":
                                for fi35 := 0; fi35 < 13; fi35++ {
                                    fmt.Println(v18, v18, v22)
                                    fmt.Println(v21, v20)
                                    _ = fi35
                                }
                            default:
                                v2 = v2
                            }
                        }
                    }
                    v36 := GenS2{F0: v17, F1: 338, F2: 79}
                    switch 127 {
                    case 2147483647:
                        v37 := GenS1{F0: (v19 / ((v19) | 1)), F1: (4294967295 >= 794)}
                        _ = v37
                    case 32768:
                        v4 = (0.1 - ((((v4 - ((((float64(1.5) * (v4 / 1e10)) * 1.5) * (v21 + (v4 + v4))) - (v21 / -2.25))) + 0.1) * ((v4 * v21) / 1e10)) + v21))
                    default:
                        var v38 uint = 31
                        _ = v38
                    }
                    _ = v36
                }
            }
        } else {
            var v39 uint16 = 478
            fmt.Println(v18)
            _ = v39
        }
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
