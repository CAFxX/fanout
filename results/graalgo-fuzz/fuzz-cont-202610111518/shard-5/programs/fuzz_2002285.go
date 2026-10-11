package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int8 }
func genFn0(p0 bool) uint {
    p0 = true
    var v5 uint16 = 0
    if (0 <= 78) {
        for fi6 := 0; fi6 < 9; fi6++ {
            var v7 uint16 = (^(v5))
            _ = fi6
            _ = v7
        }
    }
    v8 := GenS1{F0: -128, F1: -9223372036854775808, F2: 127}
    _ = p0
    _ = v5
    _ = v8
    return 65535
}
func genFn1(p0 uint8, p1 int64, p2 int32) int {
    p2 += (p2 & (p2 + 42))
    for fi9 := 0; fi9 < 17; fi9++ {
        for fi10 := 0; fi10 < 16; fi10++ {
            if false {
                switch int(p1) {
                case 444:
                    var v11 int8 = 43
                    _ = v11
                }
                v12 := []int{fi9, 1, (fi9 * 42), 9223372036854775807}
                switch fi10 {
                case 32767:
                    var v13 uint8 = 127
                    _ = v13
                case 32768:
                    p2 += (^(p2))
                }
                _ = v12
            } else {
                fmt.Println(p1)
                v14 := []int{(fi10 + 32768), (fi9 & fi10), (fi10 % (((fi10 - (fi10 % ((42) | 1)))) | 1))}
                _ = v14
            }
            _ = fi10
        }
        _ = fi9
    }
    v15 := GenS1{F0: 42, F1: 705, F2: 40}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 127
}

func main() {
    var v1 int = 128
    v2 := make(map[int]int)
    v3 := make(map[uint32]uint32)
    var v4 int8 = 74
    v16 := []int{(-(v1)), v1, 42}
    for fi17 := 0; fi17 < 1; fi17++ {
        switch genFn1(2, 1, -602) {
        case -2147483648:
            if (255 != 65535) {
                var v18 int32 = 0
                v19 := GenS1{F0: -71, F1: 9223372036854775807, F2: v4}
                v20 := GenS1{F0: v4, F1: 42, F2: v4}
                _ = v18
                _ = v19
                _ = v20
            } else {
                for fi21 := 0; fi21 < 2; fi21++ {
                    v22 := v2[(-(v1))]
                    if (!false) {
                        if false {
                            var v23 bool = (!true)
                            _ = v23
                        } else {
                            fmt.Println(v4, fi17)
                        }
                        fmt.Println(fi21, v1, v4)
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                    } else {
                        if (120 >= 2) {
                            v24 := v2[genFn1(128, -1, 42)]
                            _ = len(v16)
                            _ = v24
                        }
                        var v25 float32 = -2.25
                        v4 += v4
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                        v3[0] = 1
                        _ = v25
                    }
                    _ = fi21
                    _ = v22
                }
                var v26 int32 = -2147483648
                _ = v26
            }
        case 252:
            delete(v3, 4294967295)
        case 32768:
            var v27 float32 = 1e10
            _ = v27
        default:
            if (128 == 255) {
                fmt.Println(v1)
            }
        }
        _ = fi17
    }
    v28 := GenS1{F0: v4, F1: -783, F2: (v4 << ((4294967295) & 63))}
    v4 = (-(v4))
    var v29 uint64 = 0
    for fi30 := 0; fi30 < 4; fi30++ {
        if ((v29 >> ((genFn0((((("a" + "") + ("foo bar" + "\n\t")) + "hello") > "a"))) & 63)) > v29) {
            if len(v16) > 0 {
                v16[0] = 32767
            }
            v31 := GenS1{F0: (v4 | 42), F1: 9223372036854775807, F2: 0}
            _ = v31
        }
        v32 := GenS1{F0: v4, F1: 0, F2: v4}
        _ = fi30
        _ = v32
    }
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    var v33 uint8 = 128
    var v34 uint64 = (v29 / ((0) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v4
}
