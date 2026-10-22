package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 uint8; F1 uint; F2 uint }
func genFn0(p0 uint64, p1 int) uint32 {
    var v6 float32 = 0.0
    switch p1 {
    case 128:
        if true {
            p1 -= -1
            if ((p0 ^ (p0 | p0)) < 4294967295) {
                var v7 float32 = 0.0
                v7 = ((float32((-2.25 - 0.0)) + (v7 - (((3.14159 / 0.1) - float32(0.1)) * (v7 - v6)))) / v6)
                v8 := GenS1{F0: 0, F1: 1}
                _ = v7
                _ = v8
            }
        }
    case -266:
        var v9 string = ("a" + "foo bar")
        _ = v9
    case -864:
        if ((-1 != 43) && true) {
            if true {
                v10 := []int{p1, p1}
                var v11 uint = 65535
                _ = v10
                _ = v11
            } else {
                if ((^(p1)) >= p1) {
                    var v12 int16 = -553
                    for fi13 := 0; fi13 < 5; fi13++ {
                        for fi14 := 0; fi14 < 17; fi14++ {
                            v15 := GenS1{F0: 1, F1: 42}
                            if true {
                                switch 32767 {
                                case -956:
                                    v16 := GenS2{F0: 36}
                                    _ = v16
                                }
                                if (!true) {
                                    switch "a" {
                                    case "":
                                        if (!false) {
                                            p0 = (^(p0))
                                            v17 := []int{p1, fi13, 42, (-(p1))}
                                            var v18 int = (p1 | (fi14 / ((fi14) | 1)))
                                            _ = v17
                                            _ = v18
                                        }
                                    case "hello":
                                        var v19 float32 = v6
                                        _ = v19
                                    }
                                } else {
                                    fmt.Println(v12)
                                }
                            }
                            _ = fi14
                            _ = v15
                        }
                        v20 := GenS3{F0: (v12 % ((v12) | 1))}
                        v21 := []int{fi13, fi13, fi13, 298}
                        _ = fi13
                        _ = v20
                        _ = v21
                    }
                    if true {
                        switch (p1 | p1) {
                        case 2147483647:
                            var v22 int16 = (^(v12))
                            _ = v22
                        case 830:
                            v6 = float32((0.1 + 3.14159))
                        case -2147483648:
                            p1 -= p1
                        default:
                            for fi23 := 0; fi23 < 14; fi23++ {
                                v6 += v6
                                _ = fi23
                            }
                        }
                        v6 -= v6
                        var v24 int32 = 2147483647
                        _ = v24
                    }
                    _ = v12
                } else {
                    switch p1 {
                    case -1:
                        v6 = (((float32(0.1) / float32(0.1)) + float32((((float64(float32(-2.25)) - (1.5 / 3.14159)) / (1.5 / 1e10)) + 0.0))) * (1.5 - ((float32(1e10) + (v6 - v6)) * v6)))
                    }
                    p0 += 18446744073709551615
                    var v25 uint32 = 4294967295
                    _ = v25
                }
                var v26 string = "foo bar"
                _ = v26
            }
        }
    default:
        fmt.Println(v6, v6, v6)
    }
    _ = p0
    _ = p1
    _ = v6
    return 2
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = (((1.5 * (0.1 / (0.1 / (1e10 * (3.14159 + float32(1e10)))))) / (-2.25 / 1.0)) - 1e10)
    var v3 int8 = 65
    v4 := make(map[string]uint32)
    var v5 int32 = -581
    v27 := GenS4{F0: 604, F1: 18446744073709551615}
    v3 = -1
    var v28 uint8 = 252
    var v29 int32 = (v5 >> ((960) & 63))
    v30 := GenS3{F0: 32767}
    var v31 int64 = -1
    v32 := GenS5{F0: (v28 * 214), F1: 0, F2: 2}
    var v33 uint = 56
    switch ("" + ("a" + ("" + "a"))) {
    case "":
        var v34 uint = (v33 | v33)
        _ = v34
    default:
        v31 = v31
    }
    if (!true) {
        fmt.Println(v29, v5)
        var v35 string = "\n\t"
        _ = v35
    } else {
        if false {
            v33 = v33
        }
        var v36 int = 127
        var v37 uint64 = 4294967295
        _ = v36
        _ = v37
    }
    v38 := GenS1{F0: (v28 * 115), F1: v3}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v38
    _ = v4
    _ = v5
}
