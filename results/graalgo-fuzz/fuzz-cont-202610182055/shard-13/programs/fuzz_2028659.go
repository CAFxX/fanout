package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int; F1 int64; F2 uint64 }
type GenS4 struct { F0 uint; F1 uint32; F2 string }
type GenS5 struct { F0 uint64; F1 int64; F2 int32 }

func main() {
    var v1 int = 1
    var v2 int = 42
    v3 := make(map[string]bool)
    var v4 uint32 = 42
    var v5 uint64 = 2
    var v6 uint64 = 4294967295
    v7 := []int{v2}
    var v8 uint16 = 0
    v6 = (v5 >> ((1) & 63))
    if true {
        var v9 float32 = 1.5
        v1 += 32767
        _ = v9
    } else {
        if (false || false) {
            var v10 int = v1
            v11 := GenS1{F0: 2}
            _ = v10
            _ = v11
        } else {
            for fi12 := 0; fi12 < 20; fi12++ {
                v3[("hello" + "")] = (false || (int8(v8) == 42))
                for fi13 := 0; fi13 < 2; fi13++ {
                    var v14 int64 = -9223372036854775808
                    var v15 int64 = v14
                    switch fi13 {
                    case -308:
                        var v16 float64 = (-2.25 - 0.1)
                        _ = v16
                    default:
                        if (1 == 0) {
                            var v17 float32 = 1e10
                            _ = v17
                        }
                    }
                    fmt.Println(v14, v4)
                    switch "foo bar" {
                    case "hello":
                        if false {
                            if (("\n\t" + "") <= ("" + "")) {
                                if (424 < 32767) {
                                    v18 := GenS2{F0: 109, F1: -827}
                                    _ = v18
                                } else {
                                    if (!(42 <= int32(v8))) {
                                        if true {
                                            var v19 uint = 4294967295
                                            v19 += 2
                                            var v20 int32 = 42
                                            _ = v19
                                            _ = v20
                                        } else {
                                            v2 += fi12
                                        }
                                        v8 -= v8
                                    } else {
                                        v21 := GenS3{F0: -2147483648, F1: (^(v14)), F2: (v6 - uint64(v4))}
                                        var v22 int = v2
                                        v23 := GenS1{F0: (v6 / ((v6) | 1))}
                                        v15 = -88
                                        switch fi12 {
                                        case 9223372036854775807:
                                            if (1 < 590) {
                                                v24 := GenS4{F0: 65535, F1: v4, F2: ""}
                                                _ = v24
                                            } else {
                                                fmt.Println(v5, v2, v6)
                                                v25 := GenS5{F0: 470, F1: (v14 << ((2) & 63)), F2: -1}
                                                _ = v25
                                            }
                                        default:
                                            v1 -= (v1 * (-(fi13)))
                                        }
                                        v26 := GenS4{F0: 255, F1: v4, F2: ""}
                                        v7 = append(v7, v1)
                                        var v27 int64 = (v14 >> ((2) & 63))
                                        v27 += -9223372036854775808
                                        _ = v21
                                        _ = v22
                                        _ = v23
                                        _ = v26
                                        _ = v27
                                    }
                                }
                            }
                        }
                    }
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                var v28 string = (("foo bar" + ("" + "hello")) + "a")
                _ = fi12
                _ = v28
            }
            if false {
                var v29 uint16 = (v8 & (v8 | v8))
                var v30 int64 = 1
                _ = v29
                _ = v30
            } else {
                v6 -= (v6 - v5)
                var v31 uint16 = v8
                fmt.Println(v4, v4, v1)
                _ = v31
            }
            v8 = v8
        }
        var v32 float64 = 1.5
        switch v1 {
        case -2147483648:
            var v33 int64 = 1
            _ = v33
        default:
            for fi34 := 0; fi34 < 20; fi34++ {
                v35 := GenS5{F0: v5, F1: 1, F2: 275}
                switch ("foo bar" + ("foo bar" + "")) {
                case "foo bar":
                    fmt.Println(v5, v8, v32)
                case "a":
                    fmt.Println(v4)
                case "\n\t":
                    switch -560 {
                    case 128:
                        v5 = v6
                    case 774:
                        for fi36 := 0; fi36 < 18; fi36++ {
                            var v37 float32 = -2.25
                            _ = fi36
                            _ = v37
                        }
                    case -251:
                        var v38 int16 = 1
                        _ = v38
                    }
                default:
                    if len(v7) > 0 {
                        v7[0] = (v1 / ((fi34) | 1))
                    }
                }
                var v39 float32 = 1.5
                _ = fi34
                _ = v35
                _ = v39
            }
        }
        _ = v32
    }
    var v40 bool = (!(("foo bar" + "hello") > ""))
    fmt.Println(v40, v2, v1)
    v4 -= 54
    _ = len(v7)
    v7 = append(v7, 53)
    v41 := GenS4{F0: 65535, F1: v4, F2: "foo bar"}
    var v42 float64 = 0.0
    var v43 uint64 = 118
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
