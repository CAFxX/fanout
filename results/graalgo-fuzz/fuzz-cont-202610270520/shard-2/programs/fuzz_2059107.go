package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 int16; F1 string; F2 int }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int8; F1 uint16 }
func genFn0(p0 string, p1 int64) uint {
    v5 := GenS1{F0: 2, F1: 1}
    v6 := GenS1{F0: 65535, F1: 1000}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 255
}
func genFn1(p0 uint8, p1 uint32, p2 uint8) uint32 {
    v7 := GenS2{F0: 9223372036854775807, F1: -356}
    v8 := []int{}
    v9 := GenS2{F0: -1, F1: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return (p1 - (p1 << ((282) & 63)))
}

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]bool)
    var v3 float64 = (1.5 / (0.1 / 1e10))
    var v4 float64 = v3
    fmt.Println(v4, v4)
    v10 := GenS1{F0: 65535, F1: -754}
    fmt.Println(v3)
    v3 += v3
    fmt.Println(v4, v4)
    for fi11 := 0; fi11 < 20; fi11++ {
        v3 += float64(((3.14159 - 1e10) + 1e10))
        v1 = append(v1, 4294967295)
        _ = fi11
    }
    v3 = ((float64((0.0 / 3.14159)) / v3) - v3)
    v3 = (0.1 * 1e10)
    delete(v2, 1)
    if false {
        for fi12 := 0; fi12 < 8; fi12++ {
            switch "foo bar" {
            case "hello":
                if (1000 != 1) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v3 += 0.0
                } else {
                    if (!true) {
                        fmt.Println(v3)
                        v3 += v3
                    } else {
                        var v13 string = (("foo bar" + "hello") + "a")
                        _ = v13
                    }
                    v14 := GenS3{F0: 224, F1: 209}
                    if false {
                        v15 := GenS3{F0: 335, F1: 1}
                        var v16 float32 = 0.1
                        _ = v15
                        _ = v16
                    }
                    _ = v14
                }
            default:
                if ("a" < "foo bar") {
                    v17 := GenS2{F0: 276, F1: 42}
                    var v18 int8 = -1
                    _ = v17
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 10; fi19++ {
                        v20 := GenS2{F0: 731, F1: -1}
                        v4 += v3
                        _ = fi19
                        _ = v20
                    }
                    v1 = append(v1, 42)
                }
            }
            for fi21 := 0; fi21 < 14; fi21++ {
                _ = len(v2)
                _ = fi21
            }
            _ = fi12
        }
    }
    var v22 int64 = 1
    if (!(((v3 * v4) / (v3 * v4)) >= 3.14159)) {
        var v23 int32 = -1
        switch "foo bar" {
        case "":
            _ = len(v2)
        }
        _ = v23
    }
    var v24 int64 = v22
    for fi25 := 0; fi25 < 7; fi25++ {
        if true {
            v26 := GenS4{F0: 42, F1: "a", F2: 0}
            v1 = append(v1, 1)
            _ = v26
        } else {
            if len(v1) > 0 {
                v1[0] = 65535
            }
            if (false && true) {
                for fi27 := 0; fi27 < 9; fi27++ {
                    v1 = append(v1, 208)
                    v24 = (v24 % ((v22) | 1))
                    _ = fi27
                }
                v4 = v4
            }
            delete(v2, 0)
        }
        _ = fi25
    }
    if (false || true) {
        v3 -= v3
        v22 += (^(v22))
    }
    if (801 < -2147483648) {
        var v28 uint = genFn0("", -1)
        v3 += ((((3.14159 + v4) / v4) - v3) / (-2.25 - float64(0.1)))
        _ = v28
    } else {
        switch "foo bar" {
        case "":
            v3 -= float64(float32((((((float64((float32(0.0) - -2.25)) + float64(float32(v4))) * v4) * (v3 * v3)) * 0.1) / (v3 + (v3 - 0.0)))))
        case "foo bar":
            fmt.Println(v3, v22)
        case "hello":
            if (!(2147483647 <= -159)) {
                v29 := GenS5{F0: 100}
                var v30 uint64 = 1
                for fi31 := 0; fi31 < 13; fi31++ {
                    if (32767 == 0) {
                        v30 = (^(v30))
                        var v32 uint8 = 0
                        _ = v32
                    } else {
                        for fi33 := 0; fi33 < 15; fi33++ {
                            switch (fi31 >> ((0) & 63)) {
                            case 127:
                                var v34 uint8 = 1
                                var v35 int8 = -128
                                var v36 int8 = 0
                                _ = v34
                                _ = v35
                                _ = v36
                            case 32767:
                                if (-131 <= -1) {
                                    v37 := GenS6{F0: -1, F1: 255}
                                    _ = v37
                                }
                            case 1:
                                v38 := GenS3{F0: v30, F1: 0}
                                _ = v38
                            default:
                                if true {
                                    var v39 float64 = float64((float32(0.1) / 1.0))
                                    fmt.Println(v4)
                                    _ = v39
                                } else {
                                    if false {
                                        var v40 float32 = 1e10
                                        fmt.Println(v22, v40, v4)
                                        var v41 int32 = 1
                                        _ = v40
                                        _ = v41
                                    }
                                }
                            }
                            var v42 float32 = 3.14159
                            var v43 uint16 = 548
                            var v44 int16 = 32767
                            _ = fi33
                            _ = v42
                            _ = v43
                            _ = v44
                        }
                    }
                    v45 := GenS4{F0: -317, F1: ("foo bar" + ""), F2: fi31}
                    v30 -= (^(v30))
                    _ = fi31
                    _ = v45
                }
                _ = v29
                _ = v30
            } else {
                v22 -= v24
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v24
    _ = v3
    _ = v4
}
