package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 int16; F2 int }
type GenS3 struct { F0 int8; F1 uint16 }
func genFn0(p0 uint, p1 int32, p2 int32) uint64 {
    if (false && (float32(0.1) <= ((1.5 / 1.0) * -2.25))) {
        p0 += p0
    } else {
        var v6 int32 = 409
        p1 -= p2
        for fi7 := 0; fi7 < 16; fi7++ {
            v8 := []int{(fi7 + fi7), fi7}
            for fi9 := 0; fi9 < 18; fi9++ {
                if false {
                    switch "foo bar" {
                    case "hello":
                        var v10 uint16 = 32767
                        _ = v10
                    case "foo bar":
                        var v11 float32 = (1.5 * 0.0)
                        _ = v11
                    }
                    v12 := GenS1{F0: 0}
                    v6 += (p1 << ((97) & 63))
                    v8 = append(v8, (fi9 << ((p0) & 63)))
                    var v13 int = fi7
                    v14 := GenS1{F0: 0}
                    if true {
                        switch (("" + "\n\t") + "a") {
                        case "foo bar":
                            if (!false) {
                                var v15 float32 = (1.5 * 1e10)
                                if true {
                                    v13 = (fi9 - (fi7 >> ((42) & 63)))
                                    var v16 string = "foo bar"
                                    _ = v16
                                }
                                var v17 float64 = (((3.14159 - 0.1) - 0.1) + float64(0.1))
                                _ = v15
                                _ = v17
                            } else {
                                fmt.Println(fi9)
                                var v18 int64 = 1
                                var v19 float64 = 0.0
                                _ = v18
                                _ = v19
                            }
                        case "\n\t":
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                        default:
                            if len(v8) > 0 {
                                v8[0] = 292
                            }
                        }
                        p2 -= -2147483648
                        p1 -= (^(p1))
                    } else {
                        var v20 bool = (true && false)
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                        if v20 {
                            p0 += p0
                            var v21 int8 = 127
                            var v22 int = fi7
                            v22 += (^(fi7))
                            var v23 int = int(v21)
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                        _ = v20
                    }
                    v24 := GenS2{F0: 484, F1: 943, F2: (fi9 * fi7)}
                    _ = v12
                    _ = v13
                    _ = v14
                    _ = v24
                }
                p0 += (p0 / (((^(p0))) | 1))
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                _ = fi9
            }
            _ = fi7
            _ = v8
        }
        _ = v6
    }
    var v25 uint = 42
    v25 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    return 106
}
func genFn1(p0 bool, p1 uint8, p2 uint64) int64 {
    p0 = true
    var v26 int64 = -937
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    return -9223372036854775808
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]uint32)
    v3 := []uint32{2147483647, 4294967295}
    var v4 int64 = 42
    v5 := make(map[int]int)
    v1[("\n\t" + "hello")] = "a"
    _ = len(v3)
    for fi27 := 0; fi27 < 19; fi27++ {
        fmt.Println(fi27, v4)
        v4 -= v4
        var v28 uint16 = 255
        _ = fi27
        _ = v28
    }
    v4 = v4
    v29 := GenS1{F0: (v4 / (((v4 + (v4 % ((v4) | 1)))) | 1))}
    var v30 float64 = (float64(0.1) * (1.5 + -2.25))
    var v31 uint64 = 42
    v5[-892] = -9223372036854775808
    fmt.Println(v4)
    v1[""] = ("foo bar" + (((("hello" + "foo bar") + "hello") + ("a" + "foo bar")) + "\n\t"))
    var v32 int16 = 398
    v33 := GenS1{F0: genFn1((255 != 131), 255, v31)}
    v31 = genFn0(255, 2147483647, 42)
    v34 := GenS3{F0: -1, F1: 223}
    for fi35 := 0; fi35 < 3; fi35++ {
        if false {
            v30 = v30
            var v36 uint = 4294967295
            _ = v36
        } else {
            switch fi35 {
            case -833:
                v30 = 0.1
            }
        }
        _ = fi35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
