package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 float32 = (((1e10 - 1.5) / (1e10 * 0.1)) * 0.0)
    var v3 float64 = float64(3.14159)
    switch "hello" {
    case "a":
        var v4 uint = 65535
        _ = v4
    case "foo bar":
        var v5 int64 = 1
        _ = v5
    case "":
        fmt.Println(v1, v3, v3)
    }
    fmt.Println(v2, v2)
    if (0 < 2) {
        var v6 int = 0
        _ = v6
    }
    fmt.Println(v2, v3, v1)
    var v7 uint64 = 2
    v8 := []int{32768}
    v7 = v1
    v7 -= v7
    var v9 int16 = -32768
    if false {
        v3 = v3
    } else {
        if (42 >= 1) {
            switch "a" {
            case "":
                switch 4294967295 {
                case 290:
                    if (65535 < 2) {
                        for fi10 := 0; fi10 < 5; fi10++ {
                            v7 += v1
                            var v11 bool = false
                            for fi12 := 0; fi12 < 6; fi12++ {
                                var v13 int64 = -1
                                v13 += v13
                                _ = fi12
                                _ = v13
                            }
                            _ = fi10
                            _ = v11
                        }
                        fmt.Println(v3)
                        if true {
                            v8 = append(v8, 127)
                            switch "\n\t" {
                            case "":
                                v3 -= 3.14159
                            default:
                                v3 = 1.5
                            }
                            v14 := GenS1{F0: false}
                            _ = v14
                        } else {
                            if true {
                                v3 -= v3
                                if len(v8) > 0 {
                                    v8[0] = 2147483647
                                }
                                var v15 bool = false
                                _ = v15
                            } else {
                                if false {
                                    fmt.Println(v1, v7, v2)
                                    switch 2147483647 {
                                    case 32767:
                                        switch -9223372036854775808 {
                                        case 0:
                                            _ = len(v8)
                                        }
                                    default:
                                        var v16 bool = ((!(-1 >= 1)) || (136 < 138))
                                        _ = v16
                                    }
                                }
                                if (189 != 372) {
                                    v7 -= 652
                                } else {
                                    if false {
                                        v7 = v1
                                        var v17 uint64 = (v1 << ((4294967295) & 63))
                                        _ = v17
                                    } else {
                                        v18 := GenS1{F0: true}
                                        _ = v18
                                    }
                                    switch 4294967295 {
                                    case -153:
                                        v19 := GenS1{F0: true}
                                        _ = v19
                                    case 565:
                                        var v20 int16 = v9
                                        _ = v20
                                    case -1:
                                        v21 := GenS1{F0: (65535 > 0)}
                                        _ = v21
                                    default:
                                        v22 := GenS2{F0: 0}
                                        _ = v22
                                    }
                                    if len(v8) > 0 {
                                        v8[0] = -306
                                    }
                                }
                                for fi23 := 0; fi23 < 9; fi23++ {
                                    var v24 bool = ((460 < 32767) || true)
                                    for fi25 := 0; fi25 < 13; fi25++ {
                                        if len(v8) >= 2 {
                                            v8 = v8[1:len(v8)]
                                        }
                                        v9 -= (v9 % ((193) | 1))
                                        v24 = v24
                                        _ = fi25
                                    }
                                    fmt.Println(v3)
                                    _ = fi23
                                    _ = v24
                                }
                            }
                        }
                    } else {
                        v26 := GenS2{F0: (v1 ^ (v1 | (v1 / ((0) | 1))))}
                        _ = v26
                    }
                default:
                    v27 := GenS2{F0: v7}
                    _ = v27
                }
            case "hello":
                if true {
                    var v28 float32 = ((float32(float64(v2)) - 0.1) * 1e10)
                    if (52 < -1) {
                        _ = len(v8)
                        var v29 uint16 = 2
                        switch 408 {
                        case 32768:
                            v7 = v1
                        case 127:
                            var v30 int32 = -1
                            _ = v30
                        default:
                            v2 += ((float32(v3) + 0.0) - float32((v3 - float64(v28))))
                        }
                        _ = v29
                    }
                    _ = v28
                } else {
                    var v31 uint64 = 18446744073709551615
                    var v32 int8 = 42
                    _ = v31
                    _ = v32
                }
            default:
                var v33 int16 = -1
                _ = v33
            }
            v34 := GenS3{F0: -133, F1: 705}
            _ = v34
        } else {
            fmt.Println(v9, v7, v9)
            if (255 <= 255) {
                v35 := GenS2{F0: v1}
                v36 := GenS4{F0: v7}
                _ = v35
                _ = v36
            }
        }
    }
    switch "hello" {
    case "\n\t":
        v9 -= v9
    }
    v37 := GenS3{F0: 42, F1: 151}
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        v2 -= (float32((v3 + v3)) * v2)
    case "":
        var v38 int8 = -128
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v7
    _ = v8
    _ = v9
}
