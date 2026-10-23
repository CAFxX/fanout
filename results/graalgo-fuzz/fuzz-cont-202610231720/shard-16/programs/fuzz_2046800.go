package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 uint8; F1 uint; F2 string }
func genFn0(p0 uint64, p1 uint8, p2 int32) string {
    var v5 float32 = float32(1.5)
    switch 4294967295 {
    case 32767:
        if (2147483648 < 121) {
            v6 := GenS1{F0: 65535, F1: (p1 % ((2) | 1))}
            var v7 string = (("\n\t" + "a") + "")
            _ = v6
            _ = v7
        } else {
            v8 := GenS1{F0: 766, F1: (p1 / (((p1 & (^(p1)))) | 1))}
            if false {
                v9 := GenS2{F0: 4294967295, F1: p0, F2: 2147483648}
                _ = v9
            } else {
                fmt.Println(p2, p2)
                if (997 == 255) {
                    v10 := []int{32768, -2147483648, 127}
                    var v11 int64 = int64(p2)
                    switch "hello" {
                    case "":
                        v12 := GenS1{F0: 42, F1: 255}
                        _ = v12
                    case "a":
                        v13 := GenS1{F0: 65535, F1: p1}
                        _ = v13
                    default:
                        var v14 uint16 = uint16(p0)
                        _ = v14
                    }
                    _ = v10
                    _ = v11
                }
            }
            _ = v8
        }
    default:
        v15 := GenS3{F0: "hello", F1: -117}
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return "hello"
}
func genFn1(p0 int16, p1 int16) uint {
    switch "\n\t" {
    case "a":
        v16 := []int{9223372036854775807, 282, 0, -1}
        _ = v16
    }
    v17 := []int{-749, 127, 32768, 130}
    var v18 int64 = 1
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return 1
}

func main() {
    v1 := []string{}
    var v2 bool = true
    var v3 uint8 = 97
    v4 := make(map[string]uint32)
    v2 = (true || v2)
    var v19 uint16 = 92
    v20 := GenS3{F0: "\n\t", F1: 42}
    var v21 float32 = (0.1 / 1e10)
    if (!(18446744073709551615 != 17)) {
        for fi22 := 0; fi22 < 18; fi22++ {
            if v2 {
                if v2 {
                    v19 -= v19
                    v3 = v3
                    for fi23 := 0; fi23 < 8; fi23++ {
                        switch fi22 {
                        case 2147483647:
                            v4["\n\t"] = 2
                        case 4294967295:
                            v24 := GenS4{F0: (v3 % ((v3) | 1)), F1: genFn1(-1, -1), F2: "\n\t"}
                            var v25 int8 = 34
                            if (!v2) {
                                var v26 int32 = -925
                                if v2 {
                                    if (18446744073709551615 > 1) {
                                        var v27 uint = 255
                                        _ = v27
                                    }
                                    if v2 {
                                        if v2 {
                                            v28 := GenS1{F0: 411, F1: v3}
                                            var v29 uint32 = 1
                                            switch ("\n\t" + ("hello" + "a")) {
                                            case "foo bar":
                                                v3 = (v3 + v3)
                                            }
                                            _ = v28
                                            _ = v29
                                        } else {
                                            var v30 int8 = (v25 >> ((65535) & 63))
                                            v31 := GenS3{F0: "hello", F1: (^(v30))}
                                            var v32 bool = ((false || v2) || true)
                                            _ = v30
                                            _ = v31
                                            _ = v32
                                        }
                                        v25 -= v25
                                        v21 -= (float32(1.5) + -2.25)
                                    }
                                    fmt.Println(v2, fi23, v2)
                                }
                                v25 -= (v25 & 1)
                                var v33 uint32 = 65535
                                _ = len(v4)
                                _ = v26
                                _ = v33
                            } else {
                                if (975 != 246) {
                                    v34 := GenS2{F0: 2147483648, F1: 18446744073709551615, F2: 0}
                                    fmt.Println(v19, fi23)
                                    var v35 uint64 = 4294967295
                                    _ = v34
                                    _ = v35
                                }
                                var v36 string = (("" + "") + "a")
                                v2 = true
                                _ = v36
                            }
                            _ = v24
                            _ = v25
                        }
                        fmt.Println(v19)
                        _ = fi23
                    }
                }
            }
            fmt.Println(fi22, v19)
            _ = fi22
        }
    }
    v19 -= (v19 * 0)
    v19 += 495
    var v37 bool = ((false && (1000 < -32768)) && (v2 || (-476 <= -1)))
    for fi38 := 0; fi38 < 3; fi38++ {
        var v39 bool = false
        _ = fi38
        _ = v39
    }
    var v40 bool = true
    v2 = (72 > 65535)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v37
    _ = v4
    _ = v40
}
