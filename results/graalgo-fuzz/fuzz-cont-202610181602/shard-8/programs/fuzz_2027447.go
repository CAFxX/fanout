package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int; F1 int32; F2 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int64; F1 uint64; F2 int16 }
type GenS5 struct { F0 uint64; F1 int64; F2 uint }
type GenS6 struct { F0 int8 }
func genFn0(p0 uint8, p1 int64, p2 int64) uint64 {
    v7 := []int{2147483647, -168, 128}
    p2 = -9223372036854775808
    p0 = (p0 | 0)
    for fi8 := 0; fi8 < 9; fi8++ {
        if (false || true) {
            if (137 == 65535) {
                v9 := GenS1{F0: true, F1: p0}
                if true {
                    switch (-(fi8)) {
                    case 0:
                        if (-272 <= 259) {
                            if false {
                                var v10 int8 = -128
                                p0 += p0
                                v11 := GenS1{F0: true, F1: p0}
                                _ = v10
                                _ = v11
                            } else {
                                _ = len(v7)
                                if true {
                                    if ((1e10 / (0.1 + float64(1.5))) >= -2.25) {
                                        p1 = -9223372036854775808
                                        if (!true) {
                                            v12 := GenS1{F0: false, F1: 99}
                                            _ = v12
                                        } else {
                                            var v13 uint16 = 5
                                            if false {
                                                v14 := GenS1{F0: (576 <= 707), F1: (^(p0))}
                                                _ = v14
                                            }
                                            if false {
                                                v15 := GenS1{F0: (true && false), F1: (^(p0))}
                                                var v16 int = int(v13)
                                                v16 -= v16
                                                _ = v15
                                                _ = v16
                                            }
                                            _ = v13
                                        }
                                    } else {
                                        var v17 uint64 = 838
                                        _ = v17
                                    }
                                } else {
                                    p2 -= (^(p2))
                                    p0 += 127
                                }
                                v18 := GenS2{F0: (fi8 % ((128) | 1)), F1: 1, F2: 948}
                                _ = v18
                            }
                            var v19 string = ""
                            var v20 uint16 = 255
                            _ = v19
                            _ = v20
                        }
                    case 9223372036854775807:
                        if len(v7) > 0 {
                            v7[0] = -743
                        }
                    }
                    var v21 uint8 = p0
                    _ = v21
                } else {
                    v7 = append(v7, fi8)
                }
                var v22 int = fi8
                _ = v22
                _ = v9
            }
        } else {
            var v23 uint = 816
            var v24 float32 = (0.1 * (0.1 / 1e10))
            v23 = 0
            _ = v23
            _ = v24
        }
        var v25 int8 = 100
        _ = fi8
        _ = v25
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 457
}

func main() {
    var v1 float64 = (-2.25 * 1.5)
    var v2 uint = 4294967295
    v3 := make(map[uint32]bool)
    var v4 int64 = 1
    var v5 uint32 = 114
    var v6 int16 = int16(v5)
    var v26 int16 = v6
    var v27 float32 = (1.5 * 0.0)
    var v28 bool = true
    switch "\n\t" {
    case "foo bar":
        v2 = (v2 * v2)
    }
    v29 := GenS3{F0: 159}
    if (((!v28) || (!v28)) && v28) {
        var v30 uint16 = 1
        v30 += v30
        if (v6 >= v26) {
            if false {
                if false {
                    switch "hello" {
                    case "hello":
                        var v31 float32 = float32(v1)
                        _ = v31
                    case "a":
                        v32 := GenS4{F0: (v4 ^ v4), F1: 4294967295, F2: (v6 >> ((v2) & 63))}
                        _ = v32
                    default:
                        var v33 uint32 = (v5 * v5)
                        _ = v33
                    }
                }
                if true {
                    v26 -= 1000
                    if v28 {
                        v34 := GenS5{F0: 18446744073709551615, F1: v4, F2: 822}
                        if v28 {
                            for fi35 := 0; fi35 < 9; fi35++ {
                                for fi36 := 0; fi36 < 19; fi36++ {
                                    var v37 int16 = (v26 | v6)
                                    fmt.Println(v1, v30, v5)
                                    fmt.Println(v5)
                                    _ = fi36
                                    _ = v37
                                }
                                v4 = v4
                                _ = fi35
                            }
                            v30 = (v30 & 42)
                        }
                        _ = v34
                    }
                }
            }
            v38 := GenS2{F0: 128, F1: 556, F2: 2147483648}
            var v39 uint = 18446744073709551615
            _ = v38
            _ = v39
        } else {
            v30 -= (v30 % ((v30) | 1))
            v40 := GenS6{F0: -1}
            _ = v40
        }
        _ = v30
    } else {
        fmt.Println(v1)
    }
    for fi41 := 0; fi41 < 13; fi41++ {
        fmt.Println(v1, v26, v2)
        _ = fi41
    }
    v4 = -9223372036854775808
    var v42 bool = ("\n\t" < ("foo bar" + (("a" + "foo bar") + ("hello" + "hello"))))
    v43 := []int{9223372036854775807, 42}
    for fi44 := 0; fi44 < 3; fi44++ {
        fmt.Println(v28, v6)
        _ = fi44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
