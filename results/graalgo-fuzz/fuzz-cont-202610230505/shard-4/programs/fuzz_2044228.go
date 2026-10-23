package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 string }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 uint64; F1 uint64 }
func genFn0(p0 int16, p1 bool) int8 {
    v7 := []int{128}
    p0 = p0
    _ = p0
    _ = p1
    _ = v7
    return 1
}

func main() {
    v1 := []int{-2147483648}
    var v2 int16 = 0
    v3 := make(map[uint32]uint32)
    var v4 int = 2147483647
    var v5 uint64 = 502
    v6 := make(map[uint32]string)
    if len(v1) > 0 {
        v1[0] = 9223372036854775807
    }
    delete(v3, 653)
    var v8 uint16 = 32767
    fmt.Println(v8, v4)
    v4 -= (v4 >> ((0) & 63))
    for fi9 := 0; fi9 < 14; fi9++ {
        _ = len(v3)
        v8 = (v8 + 42)
        var v10 uint = 324
        _ = fi9
        _ = v10
    }
    var v11 uint32 = 4294967295
    if ("\n\t" > "foo bar") {
        var v12 int = 42
        switch "a" {
        case "hello":
            for fi13 := 0; fi13 < 12; fi13++ {
                if (0 < (^(v11))) {
                    _ = len(v1)
                } else {
                    for fi14 := 0; fi14 < 17; fi14++ {
                        var v15 uint = 951
                        v5 -= (v5 + v5)
                        v16 := GenS1{F0: 65535, F1: 255, F2: ("foo bar" + "foo bar")}
                        _ = fi14
                        _ = v15
                        _ = v16
                    }
                    v5 -= (v5 * (v5 * (v5 & (v5 ^ v5))))
                    v2 -= -761
                }
                if false {
                    switch "" {
                    case "a":
                        v5 = 0
                    case "foo bar":
                        if false {
                            if ("" != "") {
                                for fi17 := 0; fi17 < 20; fi17++ {
                                    v2 -= v2
                                    delete(v6, 180)
                                    var v18 string = "hello"
                                    _ = fi17
                                    _ = v18
                                }
                            } else {
                                for fi19 := 0; fi19 < 15; fi19++ {
                                    if (!(15 <= -1)) {
                                        v20 := v6[(v11 + (v11 + uint32(v5)))]
                                        _ = len(v1)
                                        _ = v20
                                    }
                                    _ = fi19
                                }
                                fmt.Println(v11)
                                if false {
                                    v21 := v3[v11]
                                    v5 -= v5
                                    _ = v21
                                }
                            }
                            v2 -= v2
                        }
                    case "hello":
                        for fi22 := 0; fi22 < 16; fi22++ {
                            if false {
                                v4 += fi22
                                var v23 int8 = -126
                                v12 -= (v4 * (v4 << ((uint(v5)) & 63)))
                                fmt.Println(v11, v8)
                                if (true && (42 >= 0)) {
                                    v23 -= v23
                                    v24 := GenS1{F0: 2, F1: 1, F2: "foo bar"}
                                    _ = v24
                                } else {
                                    fmt.Println(v8, fi13)
                                }
                                _ = v23
                            } else {
                                v1 = append(v1, v4)
                                var v25 int64 = 876
                                v25 = v25
                                _ = v25
                            }
                            if false {
                                var v26 uint = 1
                                var v27 uint8 = 127
                                _ = v26
                                _ = v27
                            } else {
                                v12 = (fi13 & fi22)
                                v28 := GenS2{F0: genFn0(v2, false), F1: (v8 % (((v8 + v8)) | 1))}
                                _ = v28
                            }
                            _ = fi22
                        }
                    default:
                        for fi29 := 0; fi29 < 16; fi29++ {
                            fmt.Println(v12)
                            var v30 uint32 = 1
                            _ = fi29
                            _ = v30
                        }
                    }
                    v4 = v4
                }
                switch (fi13 % ((354) | 1)) {
                case 32768:
                    v12 -= -350
                case -462:
                    v31 := GenS2{F0: genFn0((v2 << ((0) & 63)), ((v8 ^ 920) != v8)), F1: 32767}
                    _ = v31
                default:
                    v11 -= v11
                }
                _ = fi13
            }
        case "\n\t":
            if (false && (!false)) {
                v32 := GenS3{F0: 2, F1: (^(v5))}
                _ = v32
            } else {
                var v33 uint32 = v11
                _ = v33
            }
        default:
            for fi34 := 0; fi34 < 4; fi34++ {
                var v35 uint16 = v8
                v4 = (^(v12))
                v1 = append(v1, int(v2))
                _ = fi34
                _ = v35
            }
        }
        v5 = 1
        _ = v12
    }
    v2 = (v2 / ((933) | 1))
    delete(v3, v11)
    v2 -= 1000
    var v36 float32 = ((1.5 + 1.5) * 0.1)
    v37 := GenS1{F0: (v8 / ((v8) | 1)), F1: 182, F2: ("\n\t" + (("" + "foo bar") + ""))}
    v38 := GenS1{F0: v8, F1: 128, F2: ("a" + "hello")}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
