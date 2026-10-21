package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int16; F1 int; F2 uint8 }
type GenS3 struct { F0 int8; F1 bool; F2 int8 }
type GenS4 struct { F0 int64; F1 uint }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int; F1 int; F2 uint8 }
type GenS7 struct { F0 int32 }
func genFn0(p0 bool, p1 uint8, p2 bool) uint32 {
    var v5 uint64 = 18446744073709551615
    for fi6 := 0; fi6 < 8; fi6++ {
        p1 -= (^(p1))
        for fi7 := 0; fi7 < 5; fi7++ {
            if (uint8(v5) <= (p1 * (p1 / ((127) | 1)))) {
                if p2 {
                    var v8 int32 = 1
                    var v9 uint64 = 1
                    var v10 int16 = 1
                    _ = v10
                    _ = v8
                    _ = v9
                } else {
                    if true {
                        v11 := GenS1{F0: v5, F1: (v5 * v5)}
                        _ = v11
                    } else {
                        p1 = 127
                        p1 = 181
                        p0 = (362 != 836)
                    }
                    fmt.Println(v5)
                    p0 = (true && (726 > 1))
                }
            }
            var v12 int = fi7
            fmt.Println(fi6, v5, p0)
            v5 = 18446744073709551615
            _ = fi7
            _ = v12
        }
        for fi13 := 0; fi13 < 19; fi13++ {
            switch fi13 {
            case -1:
                var v14 int16 = 385
                _ = v14
            case 1:
                var v15 bool = (((2 != 888) || p2) || p0)
                _ = v15
            default:
                if (214 >= 32767) {
                    v16 := GenS1{F0: 42, F1: (v5 >> ((1) & 63))}
                    v5 = v5
                    _ = v16
                }
            }
            fmt.Println(fi6)
            if (!p0) {
                var v17 uint8 = 127
                p2 = p0
                _ = v17
            }
            var v18 int64 = int64(v5)
            fmt.Println(fi6)
            _ = fi13
            _ = v18
        }
        _ = fi6
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        v20 := []int{2147483647}
        for fi21 := 0; fi21 < 18; fi21++ {
            p1 -= p1
            p1 = 53
            if len(v20) > 0 {
                v20[0] = fi19
            }
            _ = fi21
        }
        if p0 {
            var v22 string = ""
            v23 := GenS2{F0: -59, F1: 127, F2: p1}
            _ = v22
            _ = v23
        }
        _ = fi19
        _ = v20
    }
    v24 := []int{-1, 2147483647}
    for fi25 := 0; fi25 < 6; fi25++ {
        v26 := GenS3{F0: 26, F1: ((p1 * 128) >= p1), F2: -1}
        p0 = (!(-1 > -1))
        _ = fi25
        _ = v26
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v5
    return 2147483647
}
func genFn1() int8 {
    var v27 bool = true
    v28 := []int{}
    if v27 {
        if len(v28) > 0 {
            v28[0] = -2147483648
        }
        for fi29 := 0; fi29 < 4; fi29++ {
            v27 = (831 != -9223372036854775808)
            if v27 {
                if true {
                    var v30 int = (^(fi29))
                    var v31 int16 = 948
                    var v32 int16 = (v31 / (((v31 & (^(v31)))) | 1))
                    _ = v30
                    _ = v31
                    _ = v32
                }
            } else {
                v33 := GenS4{F0: -424, F1: 1}
                _ = v33
            }
            if v27 {
                _ = len(v28)
                v27 = (!v27)
            }
            _ = fi29
        }
    } else {
        if v27 {
            v34 := GenS5{F0: 42}
            _ = v34
        }
    }
    fmt.Println(v27)
    _ = v27
    _ = v28
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = 1.5
    var v3 float32 = -2.25
    var v4 int16 = 0
    v35 := GenS6{F0: -355, F1: 32768, F2: 219}
    var v36 uint = 255
    v36 += 871
    var v37 string = "a"
    v38 := []int{128, 376, 32767, -2147483648}
    v39 := GenS7{F0: -2147483648}
    for fi40 := 0; fi40 < 3; fi40++ {
        v4 = v4
        v37 = (v37 + v37)
        _ = fi40
    }
    var v41 string = (("foo bar" + ((v37 + v37) + ("" + v37))) + v37)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v41
}
