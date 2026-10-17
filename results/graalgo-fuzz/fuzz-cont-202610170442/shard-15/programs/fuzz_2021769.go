package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int64 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 int16; F1 int8 }
func genFn0(p0 uint, p1 int8, p2 int16) uint64 {
    if ((255 > 1) && true) {
        var v5 uint64 = 680
        _ = v5
    }
    p0 -= (p0 + 255)
    p2 = (p2 / (((p2 % (((p2 | p2)) | 1))) | 1))
    p2 = (^(p2))
    v6 := []int{-1, -158, -556}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 739
}

func main() {
    var v1 float32 = (0.1 / -2.25)
    var v2 int64 = 1
    var v3 uint = 160
    var v4 int32 = -2147483648
    v7 := []int{127, -710, 607, -2147483648}
    v3 -= v3
    if (65535 == 42) {
        v1 += float32(1.5)
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 uint = v3
            v9 += (v9 >> ((v3) & 63))
            var v10 int32 = v4
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = len(v7)
    }
    var v11 int16 = 1000
    v1 = v1
    var v12 float32 = (((v1 / float32(float64(1e10))) / ((v1 - (float32(1e10) / 3.14159)) * v1)) * v1)
    for fi13 := 0; fi13 < 5; fi13++ {
        v4 = v4
        _ = fi13
    }
    v7 = append(v7, 895)
    if false {
        for fi14 := 0; fi14 < 7; fi14++ {
            for fi15 := 0; fi15 < 10; fi15++ {
                var v16 uint32 = 928
                if false {
                    if true {
                        if len(v7) > 0 {
                            v7[0] = int(v16)
                        }
                        var v17 uint32 = v16
                        v3 = v3
                        _ = v17
                    } else {
                        switch (fi14 % ((fi15) | 1)) {
                        case 9223372036854775807:
                            v18 := GenS1{F0: 457, F1: v3, F2: (v2 + (v2 / ((0) | 1)))}
                            _ = v18
                        case -445:
                            if false {
                                if false {
                                    fmt.Println(v3, v16, fi15)
                                    fmt.Println(v4, fi15, v4)
                                    var v19 int16 = (v11 ^ (v11 + v11))
                                    v20 := GenS1{F0: 1, F1: 929, F2: v2}
                                    _ = v19
                                    _ = v20
                                } else {
                                    v3 = (v3 % ((v3) | 1))
                                }
                                if false {
                                    switch ("" + ("" + (("a" + "\n\t") + ("" + "foo bar")))) {
                                    case "":
                                        var v21 uint32 = (v16 % (((v16 << (((v3 + v3)) & 63))) | 1))
                                        _ = v21
                                    case "a":
                                        fmt.Println(v1, v4, v3)
                                    case "hello":
                                        fmt.Println(v2, v3)
                                        v3 = v3
                                        if false {
                                            v4 -= (v4 >> (((v3 << (((v3 ^ 255)) & 63))) & 63))
                                        } else {
                                            v22 := GenS2{F0: 128, F1: "foo bar"}
                                            _ = v22
                                        }
                                        if true {
                                            var v23 uint64 = 4294967295
                                            if len(v7) >= 2 {
                                                v7 = v7[1:len(v7)]
                                            }
                                            _ = v23
                                        } else {
                                            if false {
                                                v12 -= 1e10
                                                v16 = (v16 + v16)
                                            } else {
                                                v4 -= (v4 + (v4 - 0))
                                            }
                                            v16 += (v16 + v16)
                                            switch "\n\t" {
                                            case "":
                                                v24 := GenS2{F0: 183, F1: "a"}
                                                _ = v24
                                            }
                                        }
                                    }
                                    fmt.Println(v3)
                                    fmt.Println(v16)
                                } else {
                                    if true {
                                        var v25 uint64 = 42
                                        var v26 uint64 = (v25 - v25)
                                        _ = v25
                                        _ = v26
                                    }
                                    _ = len(v7)
                                    var v27 int32 = (v4 + v4)
                                    _ = v27
                                }
                            }
                        }
                        var v28 float32 = 1.5
                        _ = v28
                    }
                } else {
                    var v29 string = "foo bar"
                    if true {
                        fmt.Println(v4, v2, v1)
                        switch "foo bar" {
                        case "a":
                            v7 = append(v7, (fi15 & (fi14 << ((v3) & 63))))
                        case "hello":
                            var v30 int = (fi15 + (-(fi15)))
                            _ = v30
                        }
                    } else {
                        if true {
                            v31 := GenS3{F0: v11, F1: 127}
                            var v32 uint = (v3 ^ (v3 + 4294967295))
                            _ = v31
                            _ = v32
                        } else {
                            var v33 float32 = v1
                            _ = v33
                        }
                        fmt.Println(v11, v2, fi14)
                    }
                    var v34 int64 = (v2 - (-(v2)))
                    _ = v29
                    _ = v34
                }
                v35 := GenS2{F0: 231, F1: "a"}
                _ = fi15
                _ = v16
                _ = v35
            }
            var v36 uint64 = genFn0((v3 >> (((v3 >> ((v3) & 63))) & 63)), 42, (v11 / ((v11) | 1)))
            v37 := GenS1{F0: 255, F1: v3, F2: (v2 / ((v2) | 1))}
            _ = fi14
            _ = v36
            _ = v37
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
