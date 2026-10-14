package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 int8; F2 int64 }
func genFn0(p0 int16, p1 uint8, p2 uint64) int64 {
    switch "hello" {
    case "\n\t":
        var v5 uint8 = 215
        _ = v5
    default:
        fmt.Println(p2, p2)
    }
    var v6 uint16 = 255
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -9223372036854775808
}

func main() {
    var v1 uint8 = 128
    v2 := make(map[int]int)
    var v3 uint8 = (v1 << ((0) & 63))
    var v4 int = 42
    delete(v2, v4)
    v1 = 255
    v7 := []int{(-(v4)), (v4 & (-(v4))), 549, (v4 / (((v4 >> ((42) & 63))) | 1))}
    v3 = v3
    v8 := v2[(v4 | int(v3))]
    if (213 == (v1 << ((65535) & 63))) {
        var v9 uint8 = v3
        switch "\n\t" {
        case "":
            if true {
                for fi10 := 0; fi10 < 20; fi10++ {
                    if false {
                        if false {
                            v11 := GenS1{F0: 65535}
                            v12 := GenS1{F0: 328}
                            _ = v11
                            _ = v12
                        }
                        switch "" {
                        case "a":
                            switch v8 {
                            case 128:
                                fmt.Println(v1)
                            case 31:
                                var v13 bool = true
                                _ = v13
                            default:
                                var v14 float32 = float32(-2.25)
                                _ = v14
                            }
                        }
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    } else {
                        v1 -= v1
                    }
                    _ = len(v2)
                    if (false || true) {
                        fmt.Println(v1, v8, v9)
                        for fi15 := 0; fi15 < 7; fi15++ {
                            v16 := GenS2{F0: 173, F1: 0, F2: 42}
                            var v17 int16 = 0
                            fmt.Println(v8)
                            fmt.Println(v3)
                            var v18 int64 = -9223372036854775808
                            _ = len(v7)
                            fmt.Println(v3)
                            _ = fi15
                            _ = v16
                            _ = v17
                            _ = v18
                        }
                    } else {
                        fmt.Println(v3)
                        v7 = append(v7, (-(fi10)))
                        var v19 uint16 = 187
                        _ = v19
                    }
                    _ = fi10
                }
            }
        default:
            if true {
                if (!(!false)) {
                    var v20 int32 = 1
                    _ = v20
                }
                if (true || false) {
                    if true {
                        for fi21 := 0; fi21 < 16; fi21++ {
                            v22 := v2[(^(fi21))]
                            v22 = 4294967295
                            _ = fi21
                            _ = v22
                        }
                        var v23 uint32 = 373
                        _ = v23
                    } else {
                        var v24 int = v4
                        v1 = v1
                        switch "\n\t" {
                        case "\n\t":
                            v8 = 42
                        case "":
                            var v25 uint32 = 4294967295
                            _ = v25
                        case "foo bar":
                            if ((int(v3) <= v4) && false) {
                                if false {
                                    fmt.Println(v9)
                                    v26 := GenS2{F0: v9, F1: -96, F2: -1}
                                    _ = v26
                                } else {
                                    fmt.Println(v4, v9)
                                    v3 = 42
                                    _ = len(v7)
                                }
                                v4 += 0
                            }
                        default:
                            v24 += (-(v8))
                        }
                        _ = v24
                    }
                    v8 -= v8
                    if (!true) {
                        if (!false) {
                            var v27 float64 = ((float64(1.5) + 0.1) + (((0.0 + 1e10) + (3.14159 * 1e10)) * float64(float32((0.1 * 1.5)))))
                            _ = v27
                        }
                    }
                }
            } else {
                fmt.Println(v9, v9)
            }
        }
        v3 = v3
        _ = v9
    }
    var v28 float32 = (((float32(((0.0 / (1.5 / ((float64((-2.25 + 1e10)) + (3.14159 + 1e10)) + -2.25))) + (1.5 / 1e10))) - float32((0.1 - (-2.25 - -2.25)))) - (3.14159 * (((3.14159 / float32(-2.25)) * ((float32((0.0 - float64(1e10))) * float32((0.1 + 0.0))) * 1e10)) / -2.25))) + ((3.14159 - 3.14159) * (0.0 + 1e10)))
    switch v8 {
    case 297:
        v28 -= v28
    default:
        v8 = (v8 * (v4 | v4))
    }
    _ = len(v2)
    if len(v7) > 0 {
        v7[0] = -675
    }
    var v29 int64 = -399
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
