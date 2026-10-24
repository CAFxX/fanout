package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64; F1 uint; F2 uint16 }
type GenS5 struct { F0 int; F1 int8 }
type GenS6 struct { F0 string; F1 uint16 }
type GenS7 struct { F0 int8; F1 bool }
func genFn0(p0 string) int64 {
    v5 := []int{}
    var v6 int64 = 1
    fmt.Println(v6)
    var v7 uint = 65535
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return (v6 * 204)
}
func genFn1(p0 int8, p1 int8) bool {
    v8 := GenS1{F0: 688, F1: (0 != -804), F2: 2147483647}
    if (100 <= (-(p0))) {
        var v9 int16 = 1000
        v10 := GenS2{F0: 0}
        _ = v10
        _ = v9
    } else {
        p1 = (p0 << ((65535) & 63))
        p1 -= (^(p1))
    }
    _ = p0
    _ = p1
    _ = v8
    return true
}

func main() {
    var v1 int = 42
    var v2 uint64 = 42
    var v3 uint8 = 127
    v4 := make(map[int]bool)
    switch ("hello" + "hello") {
    case "\n\t":
        var v11 int8 = -128
        _ = v11
    case "hello":
        _ = len(v4)
    default:
        v3 = 0
    }
    var v12 int = int(v3)
    v13 := GenS3{F0: 127}
    v14 := GenS4{F0: v2, F1: 255, F2: 251}
    var v15 uint32 = 4294967295
    v16 := []int{v1, 128, (v1 - v1), (v12 >> ((4294967295) & 63))}
    _ = len(v4)
    v2 = v2
    var v17 int16 = 1000
    var v18 uint8 = (v3 << ((255) & 63))
    var v19 int32 = 558
    v4[(v1 % ((v12) | 1))] = false
    if false {
        switch "\n\t" {
        case "a":
            switch "hello" {
            case "a":
                v20 := GenS1{F0: 32767, F1: false, F2: v15}
                _ = v20
            case "":
                for fi21 := 0; fi21 < 4; fi21++ {
                    if false {
                        v1 -= (v12 >> ((642) & 63))
                        v12 += (v1 & v1)
                        if (v19 == v19) {
                            switch "" {
                            case "foo bar":
                                v2 = v2
                            default:
                                v19 -= v19
                            }
                        } else {
                            v17 = -1
                            v3 = 127
                            v22 := GenS5{F0: v12, F1: -128}
                            _ = v22
                        }
                    } else {
                        v23 := GenS6{F0: (("\n\t" + "\n\t") + "hello"), F1: 0}
                        v1 = fi21
                        _ = v23
                    }
                    switch "\n\t" {
                    case "hello":
                        if (!true) {
                            for fi24 := 0; fi24 < 2; fi24++ {
                                var v25 string = "hello"
                                if (!false) {
                                    if false {
                                        var v26 float64 = (-2.25 - (0.1 / 1.0))
                                        fmt.Println(fi21, v15, v25)
                                        _ = v26
                                    }
                                    v27 := GenS7{F0: 0, F1: true}
                                    _ = v27
                                } else {
                                    fmt.Println(v19)
                                    var v28 float64 = 3.14159
                                    _ = v28
                                }
                                if (278 > v15) {
                                    v1 = fi21
                                    var v29 int32 = v19
                                    switch v25 {
                                    case "":
                                        if false {
                                            v30 := GenS2{F0: 9223372036854775807}
                                            _ = v30
                                        }
                                    }
                                    _ = v29
                                } else {
                                    v25 = v25
                                    var v31 float32 = (((0.1 / 1.0) + float32(3.14159)) - (1e10 / (float32(1e10) / (3.14159 + -2.25))))
                                    var v32 float64 = (0.1 / (0.1 - 1.5))
                                    _ = v31
                                    _ = v32
                                }
                                _ = fi24
                                _ = v25
                            }
                        }
                    case "foo bar":
                        switch v12 {
                        case -340:
                            for fi33 := 0; fi33 < 19; fi33++ {
                                v34 := GenS5{F0: 32767, F1: -128}
                                v35 := GenS5{F0: 9223372036854775807, F1: -1}
                                _ = fi33
                                _ = v34
                                _ = v35
                            }
                        default:
                            v36 := v4[v1]
                            _ = v36
                        }
                    }
                    if len(v16) > 0 {
                        v16[0] = (v12 % ((int(v15)) | 1))
                    }
                    _ = fi21
                }
            }
        default:
            v18 = 42
        }
        v16 = append(v16, (v12 - (v1 << ((2) & 63))))
    }
    fmt.Println(v17, v1)
    fmt.Println(v3, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
