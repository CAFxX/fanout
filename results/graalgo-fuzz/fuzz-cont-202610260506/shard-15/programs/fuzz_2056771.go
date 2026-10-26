package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 string; F1 int8 }
func genFn0(p0 int64, p1 uint32, p2 uint) int64 {
    if false {
        v7 := GenS1{F0: 42, F1: 496, F2: p2}
        p2 += p2
        _ = v7
    }
    v8 := []int{9223372036854775807, -2147483648, 42}
    p2 = (p2 >> (((p2 % (((p2 | (p2 ^ (^(p2))))) | 1))) & 63))
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return p0
}
func genFn1(p0 int64, p1 int64) uint64 {
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := GenS1{F0: 0, F1: 0, F2: 18446744073709551615}
        var v11 float32 = 0.1
        _ = fi9
        _ = v10
        _ = v11
    }
    p1 = p1
    _ = p0
    _ = p1
    return 1
}

func main() {
    v1 := []uint32{707}
    var v2 float32 = 3.14159
    v3 := make(map[int]string)
    v4 := make(map[uint32]bool)
    var v5 uint32 = 895
    v6 := []string{}
    v12 := GenS2{F0: 1, F1: "hello"}
    for fi13 := 0; fi13 < 20; fi13++ {
        v5 += 128
        if (1 < -452) {
            v14 := GenS2{F0: -218, F1: "foo bar"}
            v15 := GenS3{F0: 1, F1: 1}
            for fi16 := 0; fi16 < 8; fi16++ {
                if ((3.14159 / -2.25) >= -2.25) {
                    v17 := GenS4{F0: "\n\t", F1: -1}
                    switch (("a" + (("hello" + ("hello" + ("\n\t" + "foo bar"))) + "")) + "foo bar") {
                    case "hello":
                        var v18 int64 = genFn0(int64(v5), v5, 1)
                        _ = v18
                    case "\n\t":
                        fmt.Println(v2, fi13)
                    case "":
                        v19 := v4[v5]
                        _ = v19
                    }
                    _ = v17
                } else {
                    if true {
                        v20 := GenS3{F0: 65535, F1: 695}
                        v21 := GenS4{F0: (("hello" + (("\n\t" + "hello") + "a")) + ("hello" + "\n\t")), F1: -81}
                        v22 := GenS2{F0: 966, F1: "foo bar"}
                        _ = v20
                        _ = v21
                        _ = v22
                    } else {
                        fmt.Println(v2, fi16, v2)
                        if false {
                            if true {
                                v5 = (v5 ^ v5)
                                fmt.Println(fi13)
                            }
                            v23 := GenS1{F0: 215, F1: 289, F2: 65535}
                            var v24 uint64 = 1
                            v5 = 4294967295
                            switch 2147483647 {
                            case 193:
                                v2 = ((v2 * v2) + v2)
                            case -2147483648:
                                if false {
                                    switch "hello" {
                                    case "\n\t":
                                        if len(v1) > 0 {
                                            v1[0] = (v5 * v5)
                                        }
                                    case "foo bar":
                                        var v25 int64 = int64(v5)
                                        v5 -= (^(v5))
                                        var v26 float64 = 3.14159
                                        _ = v25
                                        _ = v26
                                    case "hello":
                                        var v27 int32 = -753
                                        _ = v27
                                    }
                                } else {
                                    switch (fi13 * (fi13 ^ (fi16 ^ fi13))) {
                                    case 9223372036854775807:
                                        if false {
                                            v28 := GenS1{F0: 0, F1: 32767, F2: 799}
                                            var v29 string = "foo bar"
                                            _ = v28
                                            _ = v29
                                        } else {
                                            v24 = (v24 - v24)
                                            var v30 int32 = 0
                                            _ = v30
                                        }
                                    }
                                }
                            default:
                                if false {
                                    var v31 float32 = (float32(3.14159) - v2)
                                    _ = v31
                                } else {
                                    v2 = v2
                                }
                            }
                            fmt.Println(v5)
                            var v32 bool = true
                            _ = v23
                            _ = v24
                            _ = v32
                        } else {
                            if (153 <= 65535) {
                                fmt.Println(v5, v2)
                            } else {
                                v5 = (v5 % ((988) | 1))
                                delete(v3, fi13)
                                if len(v6) > 0 {
                                    v6[0] = ("" + ((("\n\t" + ("foo bar" + "")) + ("a" + "hello")) + ""))
                                }
                            }
                        }
                    }
                }
                fmt.Println(fi13)
                delete(v3, 4294967295)
                _ = fi16
            }
            _ = v14
            _ = v15
        } else {
            var v33 string = (("hello" + "foo bar") + "foo bar")
            v2 += v2
            var v34 uint8 = 127
            _ = v33
            _ = v34
        }
        var v35 int8 = 100
        _ = fi13
        _ = v35
    }
    v5 += v5
    var v36 uint16 = 255
    var v37 uint = 1
    for fi38 := 0; fi38 < 5; fi38++ {
        if (!false) {
            switch "" {
            case "":
                switch (fi38 << ((v37) & 63)) {
                case 127:
                    switch "" {
                    case "":
                        for fi39 := 0; fi39 < 1; fi39++ {
                            v2 = (float32(1.5) / (v2 - (0.1 - (v2 + (1.5 / float32(-2.25))))))
                            v37 = (v37 ^ uint(v5))
                            var v40 int32 = 42
                            v40 = v40
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            var v41 int8 = 124
                            v5 += (v5 - v5)
                            _ = fi39
                            _ = v40
                            _ = v41
                        }
                    case "hello":
                        fmt.Println(v37)
                    case "\n\t":
                        v42 := GenS2{F0: 767, F1: (("a" + "\n\t") + "a")}
                        _ = v42
                    }
                case 2147483647:
                    v4[172] = true
                case 1:
                    v36 = v36
                default:
                    v4[(^(v5))] = (v37 <= (^(v37)))
                }
            case "\n\t":
                v37 -= v37
            }
        }
        _ = fi38
    }
    for fi43 := 0; fi43 < 1; fi43++ {
        var v44 string = ""
        var v45 uint8 = uint8(v36)
        _ = fi43
        _ = v44
        _ = v45
    }
    fmt.Println(v5, v36)
    fmt.Println(v37)
    var v46 uint64 = 18446744073709551615
    var v47 float64 = ((3.14159 + (((0.1 * (0.1 + -2.25)) + 0.1) + (-2.25 + 0.0))) * float64(float32(0.1)))
    switch "hello" {
    case "":
        for fi48 := 0; fi48 < 5; fi48++ {
            v37 += (v37 + (v37 - 983))
            v49 := GenS3{F0: 1, F1: 876}
            _ = fi48
            _ = v49
        }
    case "\n\t":
        v6 = append(v6, "a")
    }
    v50 := GenS3{F0: uint16(v37), F1: 42}
    v51 := GenS4{F0: (("hello" + "") + "\n\t"), F1: 87}
    var v52 uint = v37
    for fi53 := 0; fi53 < 8; fi53++ {
        var v54 uint32 = (v5 + 2147483647)
        _ = fi53
        _ = v54
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v6
}
