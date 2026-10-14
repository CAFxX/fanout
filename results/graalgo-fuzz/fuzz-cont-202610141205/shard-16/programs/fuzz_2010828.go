package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 string }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int32) uint64 {
    if false {
        for fi7 := 0; fi7 < 15; fi7++ {
            for fi8 := 0; fi8 < 4; fi8++ {
                var v9 uint16 = 255
                v9 = v9
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        for fi11 := 0; fi11 < 1; fi11++ {
            v12 := []int{fi10, (^(fi10))}
            if (65535 != 2) {
                if (-32768 == 1000) {
                    var v13 int = (fi10 | (^(fi10)))
                    switch (fi10 & v13) {
                    case -158:
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                    case 42:
                        var v14 uint = 18446744073709551615
                        _ = v14
                    }
                    _ = v13
                } else {
                    var v15 float32 = 1.5
                    p0 -= p0
                    v15 = v15
                    _ = v15
                }
                p0 = p0
            }
            _ = fi11
            _ = v12
        }
        if true {
            for fi16 := 0; fi16 < 3; fi16++ {
                if false {
                    p0 = 204
                    if (((fi16 - (fi10 - fi16)) >= (-(fi10))) || false) {
                        v17 := []int{(fi16 - (fi10 ^ fi10)), fi16}
                        p0 -= (p0 & (p0 | 0))
                        if (((0.1 * (float32(-2.25) / 1e10)) + 1.5) >= (((1.5 - 1e10) + float32(1.5)) / (1e10 / float32(float64(1e10))))) {
                            if false {
                                if false {
                                    fmt.Println(p0)
                                    if len(v17) > 0 {
                                        v17[0] = 128
                                    }
                                    var v18 uint16 = 255
                                    _ = v18
                                }
                            }
                            p0 = 0
                            fmt.Println(p0, fi10)
                        } else {
                            var v19 float32 = (1e10 - (0.0 - -2.25))
                            v17 = append(v17, (fi16 | 2147483647))
                            p0 = (p0 | (p0 ^ (p0 % ((p0) | 1))))
                            _ = v19
                        }
                        _ = v17
                    } else {
                        v20 := GenS1{F0: 114, F1: 0, F2: "hello"}
                        _ = v20
                    }
                } else {
                    p0 = p0
                    p0 = (p0 % (((p0 ^ 0)) | 1))
                    var v21 float32 = 1e10
                    _ = v21
                }
                v22 := []int{-544, 552, (fi10 - fi10)}
                var v23 float64 = -2.25
                _ = fi16
                _ = v22
                _ = v23
            }
            if true {
                fmt.Println(fi10, p0)
                v24 := GenS1{F0: -143, F1: 2147483648, F2: "foo bar"}
                v25 := GenS2{F0: 382, F1: -23}
                _ = v24
                _ = v25
            }
        }
        for fi26 := 0; fi26 < 12; fi26++ {
            v27 := []int{}
            var v28 bool = ((42 < 255) && ((960 < 248) || (-8 > -654)))
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = fi10
    }
    _ = p0
    return 4294967295
}

func main() {
    var v1 int8 = -87
    var v2 uint16 = 2
    var v3 int = 9223372036854775807
    v4 := make(map[int]bool)
    var v5 uint64 = 858
    v6 := []uint32{1, 0, 1}
    v29 := GenS1{F0: -107, F1: 2, F2: ""}
    var v30 uint = 0
    for fi31 := 0; fi31 < 1; fi31++ {
        v32 := GenS3{F0: v5}
        fmt.Println(v3)
        _ = fi31
        _ = v32
    }
    v5 += v5
    switch "" {
    case "":
        v33 := GenS1{F0: 260, F1: 0, F2: "\n\t"}
        _ = v33
    case "\n\t":
        if false {
            v5 -= (^(v5))
            fmt.Println(v2, v2)
            for fi34 := 0; fi34 < 5; fi34++ {
                v35 := GenS1{F0: 42, F1: 0, F2: (("\n\t" + "hello") + "a")}
                _ = fi34
                _ = v35
            }
        }
    case "hello":
        var v36 float64 = 1e10
        _ = v36
    default:
        v3 += (v3 + (v3 | (-(v3))))
    }
    fmt.Println(v3, v5)
    if false {
        fmt.Println(v30)
        fmt.Println(v5, v2, v30)
        if (2147483647 > 1) {
            if len(v6) > 0 {
                v6[0] = 2147483647
            }
            var v37 string = ""
            _ = v37
        }
    } else {
        var v38 float32 = ((float32(0.1) - 0.0) / 1.0)
        fmt.Println(v3)
        _ = v38
    }
    fmt.Println(v30, v1)
    var v39 float32 = 0.1
    v2 += v2
    var v40 float32 = v39
    v1 += -128
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
