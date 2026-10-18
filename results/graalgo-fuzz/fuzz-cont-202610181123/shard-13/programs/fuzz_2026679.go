package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 uint8; F2 int8 }
func genFn0(p0 int, p1 uint32, p2 string) uint32 {
    if (4294967295 == p0) {
        v4 := []int{2147483647, (^(p0)), -1, p0}
        _ = v4
    } else {
        var v5 uint32 = (p1 % ((p1) | 1))
        p1 = p1
        if (1 >= -1) {
            v6 := GenS1{F0: p1, F1: 0}
            _ = v6
        }
        _ = v5
    }
    p2 = p2
    if false {
        v7 := []int{(p0 - 32767), (p0 + p0), (-(p0))}
        p2 = p2
        _ = v7
    }
    var v8 int16 = 1000
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return p1
}
func genFn1(p0 int8, p1 uint64) uint8 {
    v9 := []int{-138}
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    v9 = append(v9, 1)
    var v10 string = "\n\t"
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 2
}

func main() {
    v1 := []uint32{2147483647}
    var v2 uint64 = 4294967295
    v3 := []string{"foo bar", "foo bar"}
    var v11 float64 = 3.14159
    fmt.Println(v2, v11)
    for fi12 := 0; fi12 < 3; fi12++ {
        for fi13 := 0; fi13 < 3; fi13++ {
            v11 -= v11
            _ = fi13
        }
        fmt.Println(fi12)
        _ = fi12
    }
    fmt.Println(v2)
    v1 = append(v1, 4294967295)
    for fi14 := 0; fi14 < 12; fi14++ {
        if false {
            for fi15 := 0; fi15 < 11; fi15++ {
                v16 := GenS1{F0: 65535, F1: 2147483648}
                var v17 int8 = 59
                switch fi14 {
                case 32767:
                    v17 = (v17 / ((int8(fi14)) | 1))
                default:
                    switch (fi15 ^ fi14) {
                    case -9223372036854775808:
                        var v18 float64 = v11
                        _ = v18
                    case 4294967295:
                        v11 = v11
                    case 42:
                        _ = len(v3)
                    }
                }
                fmt.Println(fi14, v17)
                switch "a" {
                case "foo bar":
                    if true {
                        v17 = 127
                    } else {
                        if (int32(v2) == 0) {
                            v17 = v17
                            v3 = append(v3, "foo bar")
                            v19 := GenS1{F0: 2147483647, F1: 0}
                            _ = v19
                        }
                    }
                case "":
                    v20 := GenS1{F0: 121, F1: 0}
                    _ = v20
                case "a":
                    switch "" {
                    case "a":
                        v21 := GenS2{F0: 127}
                        _ = v21
                    case "":
                        v22 := GenS1{F0: genFn0(fi14, 869, "a"), F1: 2}
                        fmt.Println(v11)
                        _ = v22
                    default:
                        if true {
                            v23 := GenS1{F0: 187, F1: 0}
                            var v24 int32 = -2147483648
                            _ = v23
                            _ = v24
                        }
                    }
                }
                var v25 uint64 = (^(v2))
                _ = fi15
                _ = v16
                _ = v17
                _ = v25
            }
            if len(v1) > 0 {
                v1[0] = 698
            }
        } else {
            var v26 float64 = 0.0
            if (false && false) {
                if false {
                    v27 := GenS2{F0: 136}
                    _ = v27
                } else {
                    var v28 int8 = 59
                    var v29 int64 = -9223372036854775808
                    if (((("hello" + (("\n\t" + "a") + ("hello" + "a"))) + (("hello" + "a") + "foo bar")) + "") == "foo bar") {
                        if false {
                            v2 += 0
                            var v30 uint32 = 0
                            v29 += v29
                            _ = v30
                        } else {
                            fmt.Println(v11)
                        }
                    }
                    _ = v28
                    _ = v29
                }
            }
            v31 := GenS3{F0: -9223372036854775808, F1: 2, F2: 37}
            _ = v26
            _ = v31
        }
        _ = fi14
    }
    v2 -= v2
    var v32 int8 = 100
    v33 := GenS2{F0: 42}
    if true {
        if true {
            for fi34 := 0; fi34 < 10; fi34++ {
                if false {
                    v32 = v32
                    fmt.Println(v2, v2, fi34)
                } else {
                    v35 := GenS3{F0: -1, F1: 0, F2: v32}
                    _ = v35
                }
                v32 -= v32
                var v36 int8 = (v32 ^ (v32 << ((986) & 63)))
                _ = fi34
                _ = v36
            }
        } else {
            v2 = v2
        }
        var v37 int16 = -711
        var v38 int = -454
        _ = v37
        _ = v38
    }
    v32 += v32
    v32 = (-(v32))
    v32 += (v32 % ((v32) | 1))
    var v39 uint8 = 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v39
}
