package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 int8 }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 int32 }
func genFn0(p0 int) uint64 {
    fmt.Println(p0)
    v6 := GenS1{F0: false, F1: false, F2: -41}
    _ = p0
    _ = v6
    return 848
}
func genFn1(p0 int, p1 uint8, p2 uint8) int64 {
    v7 := GenS2{F0: 4294967295, F1: "hello"}
    var v8 float32 = 1.5
    v9 := []int{}
    v9 = append(v9, (p0 + p0))
    var v10 uint = 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return -9223372036854775808
}

func main() {
    var v1 uint32 = 355
    v2 := []int{2147483647, 42}
    v3 := make(map[uint32]uint32)
    v4 := []uint32{205}
    var v5 int16 = 562
    var v11 uint = 357
    var v12 float32 = float32(0.0)
    v11 = (v11 % (((^(v11))) | 1))
    v12 += (v12 - v12)
    _ = len(v3)
    fmt.Println(v1, v11)
    v12 += 1.5
    switch 127 {
    case -1:
        for fi13 := 0; fi13 < 17; fi13++ {
            v5 = v5
            if (!false) {
                v5 -= -1
                for fi14 := 0; fi14 < 20; fi14++ {
                    v1 = (^(v1))
                    if ((false || ((2 < 32767) && ((v1 + v1) < 42))) && (!true)) {
                        v15 := GenS3{F0: (!(65535 <= 302))}
                        if false {
                            var v16 uint16 = uint16(v1)
                            var v17 uint64 = 42
                            if false {
                                if ((fi13 >> ((v11) & 63)) < fi13) {
                                    v16 = (v16 | 2)
                                    var v18 uint64 = genFn0(fi13)
                                    v16 = 255
                                    _ = v18
                                } else {
                                    var v19 int8 = 15
                                    _ = v19
                                }
                                if false {
                                    switch 32768 {
                                    case 128:
                                        var v20 int8 = -128
                                        _ = v20
                                    }
                                }
                            }
                            _ = v16
                            _ = v17
                        } else {
                            if ("hello" < "\n\t") {
                                if (false && (523 <= 828)) {
                                    fmt.Println(fi13)
                                    switch "foo bar" {
                                    case "":
                                        switch int(v1) {
                                        case 128:
                                            v5 -= v5
                                        case 688:
                                            var v21 bool = true
                                            _ = v21
                                        default:
                                            fmt.Println(fi14, v5, v5)
                                        }
                                    }
                                    fmt.Println(v1, v5, v1)
                                } else {
                                    var v22 float64 = -2.25
                                    v23 := GenS3{F0: (true || (true && true))}
                                    _ = v22
                                    _ = v23
                                }
                                switch (fi13 + (fi13 >> ((v11) & 63))) {
                                case -103:
                                    v24 := GenS3{F0: false}
                                    _ = v24
                                case 0:
                                    fmt.Println(v5, v1, fi14)
                                case 1:
                                    switch 2147483647 {
                                    case 1:
                                        fmt.Println(v5)
                                    default:
                                        v25 := GenS2{F0: (v1 & v1), F1: "a"}
                                        _ = v25
                                    }
                                }
                                v5 += 42
                                v3[v1] = 540
                            }
                        }
                        if true {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v26 uint64 = genFn0((^(fi14)))
                            fmt.Println(fi13, v1, v1)
                            _ = v26
                        }
                        if true {
                            var v27 bool = false
                            _ = v27
                        }
                        v28 := GenS4{F0: 42, F1: -1}
                        _ = v15
                        _ = v28
                    }
                    var v29 uint8 = 201
                    v5 -= v5
                    switch -628 {
                    case -1:
                        var v30 uint16 = 2
                        _ = v30
                    }
                    _ = fi14
                    _ = v29
                }
            }
            var v31 uint64 = genFn0(fi13)
            _ = fi13
            _ = v31
        }
    case 32767:
        v12 = -2.25
    case 128:
        v1 -= v1
    }
    v12 = ((((v12 - ((-2.25 / v12) * 3.14159)) - -2.25) * 0.1) * v12)
    var v32 uint32 = v1
    fmt.Println(v11, v5, v11)
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
