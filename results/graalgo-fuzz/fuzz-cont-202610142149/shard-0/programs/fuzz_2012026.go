package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64; F1 string }
func genFn0(p0 uint32, p1 int8, p2 string) uint8 {
    v7 := []int{32768, 11}
    fmt.Println(p1, p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 128
}
func genFn1(p0 uint, p1 uint16) uint8 {
    if false {
        v8 := []int{-1, 416, 0, 9223372036854775807}
        for fi9 := 0; fi9 < 19; fi9++ {
            switch fi9 {
            case -139:
                v10 := GenS1{F0: 42, F1: genFn0(42, 1, "hello")}
                _ = v10
            case -2147483648:
                for fi11 := 0; fi11 < 4; fi11++ {
                    if false {
                        if (true || ((!true) && false)) {
                            var v12 uint32 = 162
                            _ = v12
                        } else {
                            var v13 int32 = 970
                            p0 -= p0
                            if true {
                                if (((-2.25 + -2.25) >= float64(float32(1e10))) || (true || false)) {
                                    v14 := GenS2{F0: -691}
                                    var v15 float64 = -2.25
                                    var v16 uint64 = 18446744073709551615
                                    _ = v14
                                    _ = v15
                                    _ = v16
                                }
                                var v17 uint = p0
                                var v18 uint64 = 1
                                _ = v17
                                _ = v18
                            }
                            _ = v13
                        }
                        p0 = (^(p0))
                        p1 += p1
                        v19 := GenS2{F0: 1}
                        if len(v8) > 0 {
                            v8[0] = 7
                        }
                        _ = v19
                    } else {
                        fmt.Println(fi9, p1)
                        var v20 float32 = 0.0
                        var v21 float64 = 1.5
                        _ = v20
                        _ = v21
                    }
                    _ = fi11
                }
            case 73:
                var v22 bool = true
                _ = v22
            }
            _ = fi9
        }
        for fi23 := 0; fi23 < 3; fi23++ {
            switch (-(fi23)) {
            case -9223372036854775808:
                p1 = p1
            default:
                if ((false || ((p0 ^ uint(p1)) != p0)) || ((fi23 ^ 127) != (^(fi23)))) {
                    if true {
                        var v24 bool = false
                        _ = v24
                    }
                    var v25 int64 = 0
                    _ = v25
                } else {
                    var v26 uint = (^(p0))
                    _ = v26
                }
            }
            _ = fi23
        }
        _ = v8
    }
    p1 = 65535
    var v27 string = "hello"
    v28 := GenS1{F0: -1, F1: 80}
    _ = p0
    _ = p1
    _ = v27
    _ = v28
    return 255
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]uint32)
    var v3 uint64 = 658
    var v4 uint64 = (v3 | v3)
    var v5 uint16 = 680
    v6 := []uint32{2, 989}
    v29 := GenS1{F0: 1, F1: 0}
    if (0.1 != 1.5) {
        v5 -= 32767
    }
    if true {
        v3 = (^(v4))
        v30 := GenS3{F0: 9223372036854775807, F1: "\n\t"}
        v6 = append(v6, 582)
        _ = v30
    } else {
        v3 -= v4
        v31 := v1["hello"]
        delete(v2, (-(v31)))
        _ = v31
    }
    var v32 int16 = 0
    switch 9223372036854775807 {
    case -9223372036854775808:
        _ = len(v1)
    case 1:
        var v33 uint8 = 195
        _ = v33
    case 9223372036854775807:
        v5 += 547
    default:
        fmt.Println(v5, v3)
    }
    v5 = (v5 * (v5 >> ((42) & 63)))
    v34 := GenS1{F0: -128, F1: 0}
    switch 9223372036854775807 {
    case -9223372036854775808:
        v32 -= -32768
    }
    v32 += v32
    var v35 uint64 = (v3 / (((v3 & (v4 * (v3 >> ((991) & 63))))) | 1))
    for fi36 := 0; fi36 < 19; fi36++ {
        v37 := GenS2{F0: (v32 ^ (v32 * v32))}
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = len(v6)
        _ = fi36
        _ = v37
    }
    for fi38 := 0; fi38 < 1; fi38++ {
        var v39 float32 = float32(float64((0.0 / 1e10)))
        _ = fi38
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v32
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
