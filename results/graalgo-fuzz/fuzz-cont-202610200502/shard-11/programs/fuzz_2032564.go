package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
func genFn0(p0 uint16, p1 bool, p2 uint8) uint8 {
    v6 := []int{32768, -719, 32768}
    switch -521 {
    case 42:
        var v7 uint = 463
        _ = v7
    }
    fmt.Println(p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p2
}
func genFn1(p0 uint16, p1 uint, p2 uint8) uint32 {
    v8 := []int{}
    if (223 != 1) {
        p0 += p0
        if (p2 == (p2 ^ uint8(p1))) {
            var v9 float64 = (3.14159 / float64(0.1))
            _ = len(v8)
            fmt.Println(p2, p2, v9)
            _ = v9
        }
        if ((-9223372036854775808 < 0) || ((p0 & (p0 >> ((p1) & 63))) >= p0)) {
            p2 = (p2 << ((1) & 63))
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 65535
}

func main() {
    var v1 uint16 = 32767
    v2 := make(map[uint32]int)
    v3 := []int{2147483647, 32768, 9223372036854775807}
    var v4 int16 = -658
    var v5 uint = 4294967295
    v5 += (v5 ^ v5)
    v10 := GenS1{F0: -9223372036854775808, F1: -1}
    v1 = v1
    var v11 uint = (v5 << ((42) & 63))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    delete(v2, genFn1((v1 + (v1 >> ((4294967295) & 63))), v11, 128))
    switch 1 {
    case 393:
        v1 += v1
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        fmt.Println(v1, v4)
        _ = fi12
    }
    if (-9223372036854775808 != 9223372036854775807) {
        if true {
            v2[0] = -75
            var v13 int32 = 1
            var v14 int32 = -1
            _ = v13
            _ = v14
        } else {
            v11 = (v5 | (v5 ^ (^(v11))))
            for fi15 := 0; fi15 < 15; fi15++ {
                if (!(!true)) {
                    var v16 int32 = -224
                    for fi17 := 0; fi17 < 12; fi17++ {
                        if ((v1 ^ 65535) < v1) {
                            _ = len(v3)
                        } else {
                            if (0.0 > -2.25) {
                                var v18 int32 = (v16 / ((v16) | 1))
                                v5 = 441
                                _ = v18
                            }
                            v4 -= -32768
                            fmt.Println(v5)
                        }
                        v3 = append(v3, fi15)
                        if false {
                            v1 += v1
                            var v19 uint8 = genFn0((v1 ^ v1), true, 42)
                            _ = v19
                        }
                        _ = fi17
                    }
                    for fi20 := 0; fi20 < 20; fi20++ {
                        switch fi20 {
                        case 32768:
                            v21 := GenS1{F0: 509, F1: 9223372036854775807}
                            _ = v21
                        case 4294967295:
                            var v22 uint16 = v1
                            _ = v22
                        default:
                            var v23 float64 = 1.5
                            _ = v23
                        }
                        _ = fi20
                    }
                    _ = v16
                } else {
                    for fi24 := 0; fi24 < 12; fi24++ {
                        v1 += v1
                        _ = fi24
                    }
                    var v25 uint64 = 711
                    _ = v25
                }
                _ = fi15
            }
        }
        v26 := GenS1{F0: 118, F1: 42}
        var v27 int64 = -499
        _ = v26
        _ = v27
    }
    v28 := GenS1{F0: 1, F1: -738}
    for fi29 := 0; fi29 < 18; fi29++ {
        if false {
            for fi30 := 0; fi30 < 19; fi30++ {
                v2[genFn1((v1 ^ 277), 1, genFn0(v1, ((0.0 + float64(0.1)) <= ((3.14159 - 1e10) + 1e10)), 2))] = fi29
                var v31 string = (("a" + "foo bar") + ("foo bar" + "\n\t"))
                _ = fi30
                _ = v31
            }
        } else {
            var v32 uint = 4294967295
            fmt.Println(fi29, v1, v1)
            _ = v32
        }
        v4 += v4
        _ = fi29
    }
    if (3.14159 == (0.1 * 1e10)) {
        var v33 float64 = 1.5
        var v34 uint32 = 2
        _ = v33
        _ = v34
    } else {
        v11 += (v11 - (v5 ^ (v5 ^ 42)))
        var v35 int16 = v4
        _ = v35
    }
    var v36 int64 = 602
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v28
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
