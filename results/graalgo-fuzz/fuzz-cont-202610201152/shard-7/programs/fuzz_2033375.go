package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint8 }
func genFn0(p0 uint64) uint32 {
    v7 := []int{}
    p0 -= p0
    _ = p0
    _ = v7
    return uint32(p0)
}

func main() {
    v1 := []int{}
    var v2 int8 = 9
    var v3 uint = 18446744073709551615
    var v4 uint8 = 0
    var v5 int8 = v2
    var v6 int32 = -1
    var v8 uint32 = 136
    var v9 float64 = 0.0
    if false {
        v10 := GenS1{F0: 0}
        _ = v10
    } else {
        for fi11 := 0; fi11 < 18; fi11++ {
            if true {
                var v12 int32 = int32(fi11)
                for fi13 := 0; fi13 < 7; fi13++ {
                    if true {
                        v9 -= 1.5
                        switch (("a" + "hello") + "") {
                        case "\n\t":
                            _ = len(v1)
                        case "foo bar":
                            var v14 int = 0
                            _ = v14
                        default:
                            var v15 float64 = v9
                            switch ("foo bar" + ("" + (((("foo bar" + "hello") + ("a" + "foo bar")) + "hello") + (("" + "hello") + ("\n\t" + ("a" + ("" + "foo bar"))))))) {
                            case "foo bar":
                                if len(v1) > 0 {
                                    v1[0] = 42
                                }
                            }
                            var v16 uint8 = (v4 * v4)
                            _ = v15
                            _ = v16
                        }
                    } else {
                        switch "" {
                        case "a":
                            var v17 float64 = (float64(((3.14159 * -2.25) * 1e10)) / v9)
                            _ = v17
                        case "\n\t":
                            if true {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                var v18 int64 = 42
                                if len(v1) > 0 {
                                    v1[0] = (-(fi11))
                                }
                                _ = v18
                            } else {
                                fmt.Println(v6)
                                if ((v4 >> ((v3) & 63)) == (v4 >> (((v3 - (v3 ^ 1))) & 63))) {
                                    fmt.Println(v12)
                                    var v19 int64 = -21
                                    _ = v19
                                } else {
                                    if (v8 <= v8) {
                                        v2 += (v2 ^ 102)
                                    }
                                    v20 := GenS2{F0: ((fi13 ^ fi13) >= -913), F1: v4}
                                    v1 = append(v1, -229)
                                    if true {
                                        var v21 uint16 = 2
                                        _ = v21
                                    } else {
                                        fmt.Println(fi11)
                                        switch 386 {
                                        case -300:
                                            v22 := GenS1{F0: 272}
                                            _ = v22
                                        }
                                    }
                                    v23 := GenS3{F0: v12, F1: uint16(v4)}
                                    _ = v20
                                    _ = v23
                                }
                            }
                        default:
                            v2 -= v5
                        }
                        switch ("hello" + "foo bar") {
                        case "a":
                            if (!true) {
                                if false {
                                    v2 = -91
                                    var v24 float64 = float64(float32(((float64((-2.25 * (3.14159 * 1e10))) + (float64(0.1) + (v9 + 0.0))) + v9)))
                                    var v25 float64 = v9
                                    _ = v24
                                    _ = v25
                                } else {
                                    var v26 uint32 = v8
                                    v27 := GenS4{F0: 9223372036854775807, F1: 32767}
                                    var v28 uint32 = 391
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                }
                            }
                        default:
                            var v29 int16 = 225
                            _ = v29
                        }
                    }
                    _ = fi13
                }
                var v30 uint64 = uint64(v3)
                _ = v12
                _ = v30
            }
            var v31 int8 = v5
            if (42 >= 1) {
                fmt.Println(v4, v8)
                if false {
                    v8 += 542
                    var v32 uint16 = 2
                    _ = v32
                }
            } else {
                v5 = 1
                switch "a" {
                case "foo bar":
                    var v33 int = -1
                    _ = v33
                case "":
                    for fi34 := 0; fi34 < 16; fi34++ {
                        v35 := GenS3{F0: 42, F1: 1}
                        var v36 int32 = (v6 ^ v6)
                        var v37 uint = v3
                        if false {
                            if false {
                                v38 := GenS5{F0: "\n\t"}
                                v39 := GenS6{F0: v4}
                                fmt.Println(fi11)
                                _ = v38
                                _ = v39
                            }
                            switch ((("a" + "foo bar") + "foo bar") + "a") {
                            case "foo bar":
                                fmt.Println(v4, v36, v5)
                            default:
                                fmt.Println(v36, v8, v31)
                            }
                            if ((true && ((v5 >> (((v37 + v3)) & 63)) == v2)) && (!(42 <= int8(v3)))) {
                                fmt.Println(v6, v3)
                                fmt.Println(v3, v2, fi11)
                            } else {
                                fmt.Println(v36, v8)
                                v37 -= v37
                                var v40 float32 = 0.1
                                _ = v40
                            }
                        }
                        v1 = append(v1, (fi11 / (((fi11 << ((18446744073709551615) & 63))) | 1)))
                        v31 += int8(fi34)
                        _ = fi34
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                default:
                    var v41 int8 = (-(v2))
                    _ = v41
                }
                v42 := GenS4{F0: 0, F1: 1}
                _ = v42
            }
            _ = fi11
            _ = v31
        }
    }
    fmt.Println(v9)
    _ = len(v1)
    if len(v1) > 0 {
        v1[0] = -608
    }
    var v43 uint = (v3 & (v3 ^ v3))
    switch "a" {
    case "foo bar":
        var v44 uint = (v43 << ((864) & 63))
        _ = v44
    case "":
        v3 -= (v3 - (v43 ^ 4294967295))
    case "a":
        v4 = v4
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
