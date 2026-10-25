package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 string; F2 int64 }
type GenS3 struct { F0 bool; F1 int8; F2 int8 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 int32, p1 int32, p2 uint) uint64 {
    var v5 string = "a"
    for fi6 := 0; fi6 < 8; fi6++ {
        p2 -= p2
        v7 := []int{(fi6 - (fi6 & 1))}
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 0
}

func main() {
    var v1 int8 = 1
    v2 := []uint32{144, 42, 2147483648, 684}
    v3 := make(map[string]uint32)
    var v4 uint64 = 554
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v8 := v3["foo bar"]
    v8 += (^(v8))
    var v9 int16 = 1000
    var v10 int16 = (-(v9))
    fmt.Println(v1, v10)
    v11 := GenS1{F0: 328}
    if (410 != -9223372036854775808) {
        if false {
            _ = len(v3)
            if true {
                switch -2147483648 {
                case 42:
                    _ = len(v3)
                case 4294967295:
                    fmt.Println(v9, v1)
                case 2147483647:
                    v10 = 0
                }
            } else {
                v12 := GenS1{F0: 65535}
                for fi13 := 0; fi13 < 1; fi13++ {
                    if true {
                        v14 := GenS2{F0: v8, F1: ("" + "hello"), F2: 9223372036854775807}
                        switch ("foo bar" + "hello") {
                        case "a":
                            v3["foo bar"] = v8
                        case "":
                            fmt.Println(v10, v8, fi13)
                        case "\n\t":
                            if true {
                                var v15 uint = 2
                                _ = v15
                            } else {
                                v16 := GenS3{F0: true, F1: v1, F2: (v1 + (v1 | v1))}
                                if (255 > 255) {
                                    v1 = (v1 | v1)
                                    if (v1 != v1) {
                                        var v17 uint = 440
                                        fmt.Println(fi13, v1, v9)
                                        _ = v17
                                    }
                                } else {
                                    for fi18 := 0; fi18 < 5; fi18++ {
                                        fmt.Println(v10, fi13, v9)
                                        var v19 float32 = (((float32(((1e10 / (((1e10 - float64(0.1)) * -2.25) - 3.14159)) / 0.1)) + (3.14159 * (-2.25 * 0.0))) - -2.25) - 0.0)
                                        _ = fi18
                                        _ = v19
                                    }
                                    var v20 int16 = v10
                                    delete(v3, "")
                                    _ = v20
                                }
                                v8 += (v8 * 338)
                                _ = v16
                            }
                        }
                        for fi21 := 0; fi21 < 4; fi21++ {
                            var v22 int8 = 127
                            switch (("a" + ("hello" + "a")) + "foo bar") {
                            case "":
                                if (18446744073709551615 >= (^(v4))) {
                                    v4 -= (v4 / ((4294967295) | 1))
                                } else {
                                    v1 -= v22
                                    v3["foo bar"] = (v8 ^ (v8 & 2))
                                    v23 := GenS4{F0: v4}
                                    _ = v23
                                }
                            case "\n\t":
                                v8 -= (v8 & 979)
                            }
                            _ = fi21
                            _ = v22
                        }
                        _ = v14
                    } else {
                        var v24 int8 = v1
                        _ = v24
                    }
                    _ = len(v2)
                    var v25 uint8 = 163
                    _ = fi13
                    _ = v25
                }
                _ = v12
            }
        }
    }
    if true {
        v26 := GenS1{F0: 1}
        _ = v26
    } else {
        if true {
            var v27 float32 = float32((1.5 * -2.25))
            switch "foo bar" {
            case "":
                switch -2147483648 {
                case 128:
                    v28 := GenS1{F0: 65535}
                    _ = v28
                case 2147483647:
                    v8 += 4294967295
                default:
                    var v29 uint = 583
                    _ = v29
                }
            case "foo bar":
                v27 = v27
            case "hello":
                var v30 uint8 = 0
                _ = v30
            }
            v3[("a" + ("" + "hello"))] = (^(v8))
            _ = v27
        }
        v4 = v4
    }
    var v31 uint32 = 0
    v31 += (v8 * v31)
    v10 = v9
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v8
    _ = v9
}
