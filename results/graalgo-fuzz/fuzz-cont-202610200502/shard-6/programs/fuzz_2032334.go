package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int16; F1 int32; F2 int64 }
type GenS3 struct { F0 uint; F1 uint32; F2 int }
func genFn0(p0 int) string {
    switch 2147483647 {
    case -1:
        p0 += p0
    case -887:
        for fi7 := 0; fi7 < 14; fi7++ {
            p0 = (fi7 & p0)
            _ = fi7
        }
    case 127:
        p0 = (-(p0))
    default:
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 int64 = 771
            p0 += (^(fi8))
            _ = fi8
            _ = v9
        }
    }
    fmt.Println(p0)
    _ = p0
    return (("\n\t" + "foo bar") + ("a" + ("" + "hello")))
}
func genFn1() uint32 {
    for fi10 := 0; fi10 < 19; fi10++ {
        var v11 uint16 = 2
        if ((fi10 ^ fi10) > 1) {
            var v12 string = "\n\t"
            _ = v12
        } else {
            if ((!(0 != -2147483648)) && true) {
                if (!false) {
                    if true {
                        fmt.Println(v11, fi10)
                        fmt.Println(fi10, v11)
                        fmt.Println(fi10, v11)
                    } else {
                        v13 := []int{fi10}
                        if ("hello" != "a") {
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                            if false {
                                var v14 uint = 254
                                _ = v14
                            }
                        } else {
                            if (true || true) {
                                v11 -= v11
                                v15 := GenS1{F0: 128, F1: v11}
                                v16 := GenS1{F0: 1, F1: v11}
                                _ = v15
                                _ = v16
                            } else {
                                switch (fi10 >> ((901) & 63)) {
                                case 156:
                                    fmt.Println(fi10, v11)
                                case -1:
                                    v17 := GenS1{F0: 42, F1: v11}
                                    _ = v17
                                }
                                if len(v13) > 0 {
                                    v13[0] = fi10
                                }
                            }
                            if false {
                                v18 := GenS1{F0: 98, F1: 255}
                                fmt.Println(fi10)
                                _ = v18
                            } else {
                                if len(v13) >= 2 {
                                    v13 = v13[1:len(v13)]
                                }
                                var v19 string = "foo bar"
                                _ = len(v13)
                                _ = v19
                            }
                        }
                        v13 = append(v13, -1)
                        _ = v13
                    }
                    switch (fi10 & fi10) {
                    case 0:
                        for fi20 := 0; fi20 < 2; fi20++ {
                            var v21 string = "hello"
                            _ = fi20
                            _ = v21
                        }
                    default:
                        v11 = (v11 & 42)
                    }
                } else {
                    v22 := []int{(-(fi10)), (fi10 & (^(fi10))), fi10}
                    v11 -= v11
                    _ = v22
                }
                switch fi10 {
                case 952:
                    fmt.Println(v11)
                case -9223372036854775808:
                    fmt.Println(v11, fi10)
                case 32768:
                    fmt.Println(fi10, v11)
                default:
                    if (false && true) {
                        v11 += v11
                        if (v11 == (v11 | v11)) {
                            for fi23 := 0; fi23 < 5; fi23++ {
                                if true {
                                    v24 := GenS1{F0: 127, F1: 746}
                                    fmt.Println(fi10, fi10, v11)
                                    _ = v24
                                }
                                v25 := []int{}
                                _ = fi23
                                _ = v25
                            }
                            v26 := []int{}
                            v11 += 32767
                            _ = v26
                        }
                    } else {
                        v27 := []int{}
                        _ = v27
                    }
                }
            }
            v11 += (^(v11))
        }
        var v28 uint16 = 0
        _ = fi10
        _ = v11
        _ = v28
    }
    if (4294967295 == -482) {
        v29 := []int{-623}
        var v30 int = 0
        _ = v29
        _ = v30
    }
    v31 := GenS2{F0: 1, F1: 2147483647, F2: -40}
    var v32 int = 1
    _ = v31
    _ = v32
    return 1
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{("" + ((("hello" + "a") + "") + ("" + ("a" + "hello")))), "hello"}
    v3 := make(map[uint32]bool)
    var v4 float32 = (-2.25 / 1e10)
    var v5 uint64 = 42
    var v6 uint8 = 107
    var v33 bool = false
    for fi34 := 0; fi34 < 10; fi34++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v35 float64 = ((float64((1.5 * v4)) / ((1.5 / 1.0) - 1e10)) + 3.14159)
        _ = fi34
        _ = v35
    }
    v36 := GenS3{F0: 2, F1: 2, F2: 130}
    v4 += (-2.25 / (v4 + ((v4 / (0.0 + 0.1)) - (1.5 * v4))))
    v37 := GenS1{F0: v6, F1: 0}
    v4 = v4
    v5 = (v5 << ((42) & 63))
    v33 = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
