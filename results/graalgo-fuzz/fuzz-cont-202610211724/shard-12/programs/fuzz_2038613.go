package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 uint16; F2 int32 }
type GenS3 struct { F0 string; F1 string; F2 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint }
func genFn0(p0 int16, p1 int64, p2 int64) int64 {
    p1 -= (p2 / (((^(p1))) | 1))
    for fi7 := 0; fi7 < 20; fi7++ {
        switch "hello" {
        case "":
            for fi8 := 0; fi8 < 10; fi8++ {
                p2 = p1
                _ = fi8
            }
        case "\n\t":
            p0 -= p0
        case "hello":
            if false {
                if true {
                    var v9 uint32 = 734
                    for fi10 := 0; fi10 < 8; fi10++ {
                        v11 := []int{fi10, fi10}
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                        var v12 uint64 = 2
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    v13 := []int{}
                    _ = v13
                    _ = v9
                } else {
                    if false {
                        for fi14 := 0; fi14 < 16; fi14++ {
                            fmt.Println(fi7, p1, p2)
                            v15 := GenS1{F0: -1}
                            if (!true) {
                                fmt.Println(p0)
                                v16 := []int{776, fi7, (fi7 + 32768), 32768}
                                v16 = append(v16, (fi7 / (((fi14 % ((fi7) | 1))) | 1)))
                                _ = v16
                            }
                            _ = fi14
                            _ = v15
                        }
                        for fi17 := 0; fi17 < 4; fi17++ {
                            var v18 uint8 = 128
                            var v19 float32 = 1e10
                            p2 += (p1 << ((222) & 63))
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                        if false {
                            p0 = p0
                            var v20 uint64 = 84
                            p0 -= (p0 >> ((2) & 63))
                            _ = v20
                        } else {
                            v21 := GenS1{F0: 127}
                            _ = v21
                        }
                    } else {
                        switch "\n\t" {
                        case "":
                            p2 += (p1 * p2)
                        case "\n\t":
                            v22 := GenS2{F0: 127, F1: 2, F2: -1}
                            _ = v22
                        }
                    }
                }
                v23 := GenS2{F0: 42, F1: 874, F2: 143}
                _ = v23
            } else {
                v24 := []int{fi7}
                var v25 int32 = -1
                if len(v24) > 0 {
                    v24[0] = fi7
                }
                _ = v24
                _ = v25
            }
        }
        _ = fi7
    }
    v26 := GenS1{F0: 44}
    var v27 float64 = 1.5
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    _ = v27
    return p2
}
func genFn1() int {
    var v28 float32 = ((3.14159 / 1e10) + (float32(float64(1e10)) - (-2.25 / 1.0)))
    var v29 uint8 = 131
    _ = v28
    _ = v29
    return 128
}

func main() {
    var v1 string = ("foo bar" + "\n\t")
    v2 := []uint32{4294967295}
    v3 := make(map[int]bool)
    v4 := []int{128, 2147483647, -592, 1}
    v5 := []int{127}
    v6 := make(map[string]int)
    v1 = "hello"
    for fi30 := 0; fi30 < 5; fi30++ {
        _ = len(v5)
        _ = fi30
    }
    v1 = v1
    var v31 int32 = -381
    v32 := GenS3{F0: ("foo bar" + v1), F1: "hello", F2: 255}
    v31 = v31
    v2 = append(v2, 65535)
    if (((false || ("hello" != ("foo bar" + v1))) && (796 <= 42)) && (!(9223372036854775807 >= genFn1()))) {
        v33 := GenS4{F0: 0}
        _ = v33
    } else {
        var v34 uint64 = 1
        for fi35 := 0; fi35 < 18; fi35++ {
            for fi36 := 0; fi36 < 8; fi36++ {
                if len(v5) > 0 {
                    v5[0] = genFn1()
                }
                v37 := GenS4{F0: 99}
                _ = fi36
                _ = v37
            }
            v38 := GenS2{F0: 1, F1: 42, F2: v31}
            _ = fi35
            _ = v38
        }
        _ = v34
    }
    v1 = v1
    var v39 string = ("hello" + v1)
    if false {
        v39 = v1
    }
    for fi40 := 0; fi40 < 13; fi40++ {
        v1 = (v39 + v1)
        if true {
            var v41 string = "foo bar"
            v42 := v6[(v1 + (v39 + "foo bar"))]
            v41 = "foo bar"
            _ = v41
            _ = v42
        } else {
            v39 = ("" + v1)
        }
        _ = len(v5)
        _ = fi40
    }
    switch ("" + (v39 + (((("a" + (v1 + v39)) + (v1 + (v39 + "hello"))) + v39) + v39))) {
    case "\n\t":
        v43 := GenS1{F0: 100}
        _ = v43
    default:
        if true {
            v39 = ("foo bar" + (v39 + v1))
            v44 := GenS5{F0: 4294967295}
            _ = v44
        } else {
            v45 := GenS1{F0: 42}
            _ = v45
        }
    }
    v1 = ((v1 + (v39 + (("foo bar" + "hello") + "foo bar"))) + ((("" + v1) + "a") + v39))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
