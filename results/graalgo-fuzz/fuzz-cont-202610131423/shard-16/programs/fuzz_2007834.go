package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 bool }
type GenS2 struct { F0 string; F1 int32; F2 uint16 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{65535, 0, 280, 0}
    var v3 int = 687
    v4 := []int{v3, (^(v3)), (^(v3)), -513}
    var v5 int32 = 1
    v6 := make(map[int]int)
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 int64 = -417
        _ = fi7
        _ = v8
    }
    v3 += v3
    v5 += v5
    v5 = (v5 - 0)
    v9 := GenS1{F0: ((("" + ("\n\t" + "")) + "a") + ("hello" + "hello")), F1: 0, F2: false}
    for fi10 := 0; fi10 < 18; fi10++ {
        if false {
            if len(v4) > 0 {
                v4[0] = v3
            }
            var v11 int8 = 127
            v6[(fi10 ^ (-(v3)))] = -695
            _ = v11
        } else {
            var v12 int = fi10
            _ = v12
        }
        if (true || true) {
            if ((-(v5)) < 2147483647) {
                for fi13 := 0; fi13 < 6; fi13++ {
                    if true {
                        switch fi13 {
                        case 1:
                            v14 := GenS1{F0: (("hello" + "\n\t") + ""), F1: 174, F2: false}
                            _ = v14
                        }
                    }
                    var v15 uint16 = 65535
                    v16 := GenS1{F0: ("\n\t" + ""), F1: 127, F2: (!(473 == 42))}
                    v3 = fi13
                    if (!(!(2 == 2))) {
                        if true {
                            var v17 uint16 = (v15 % (((^(v15))) | 1))
                            if (1 <= 109) {
                                var v18 int16 = 1
                                _ = v18
                            }
                            v3 = (fi13 % ((-648) | 1))
                            _ = v17
                        }
                        var v19 float64 = 0.0
                        if true {
                            v4 = append(v4, -947)
                        } else {
                            if false {
                                var v20 int = int(v5)
                                if (((1e10 - (((float32(v19) * float32(-2.25)) * 3.14159) + 0.0)) <= (-2.25 * 0.0)) && (!false)) {
                                    v3 -= 42
                                }
                                if (!true) {
                                    if true {
                                        var v21 uint8 = 255
                                        var v22 float64 = (v19 + (v19 * v19))
                                        v3 = (fi13 & (fi10 ^ (fi13 - (fi10 << ((2) & 63)))))
                                        _ = v21
                                        _ = v22
                                    }
                                    fmt.Println(v15, fi10)
                                    v3 = (v3 * (^(v20)))
                                } else {
                                    if (2147483647 != 2147483647) {
                                        v20 = (v3 - (^(v3)))
                                        fmt.Println(v15, v3, fi13)
                                        v19 = 0.0
                                        var v23 uint32 = 42
                                        var v24 bool = false
                                        v20 = (v20 % ((1) | 1))
                                        fmt.Println(v19)
                                        _ = v23
                                        _ = v24
                                    } else {
                                        fmt.Println(v5, v20)
                                    }
                                    _ = len(v2)
                                    delete(v6, (v3 - 230))
                                }
                                _ = v20
                            }
                            v5 = v5
                            if len(v4) > 0 {
                                v4[0] = (^(fi10))
                            }
                        }
                        _ = v19
                    }
                    if (false || false) {
                        var v25 int8 = 2
                        if true {
                            v25 -= (v25 - 66)
                            v25 = (v25 ^ (v25 % (((-(v25))) | 1)))
                            var v26 uint16 = v15
                            _ = v26
                        } else {
                            var v27 float32 = 0.1
                            v28 := GenS2{F0: ("\n\t" + "a"), F1: (v5 | (v5 << ((1) & 63))), F2: v15}
                            _ = v27
                            _ = v28
                        }
                        _ = v25
                    } else {
                        v5 = 1
                        if (false && false) {
                            v29 := GenS1{F0: ("" + ""), F1: 1, F2: true}
                            _ = v29
                        } else {
                            v15 = 65535
                            v1[42] = true
                        }
                    }
                    _ = fi13
                    _ = v15
                    _ = v16
                }
                var v30 string = (("hello" + "") + (("\n\t" + ("hello" + (("hello" + "\n\t") + "hello"))) + ("" + "hello")))
                var v31 float64 = ((0.0 / -2.25) * 3.14159)
                _ = v30
                _ = v31
            }
        }
        v3 += v3
        _ = fi10
    }
    delete(v1, 185)
    v5 = -1
    var v32 uint32 = 0
    v32 = (v32 & v32)
    if false {
        switch "foo bar" {
        case "foo bar":
            var v33 uint32 = 2
            _ = v33
        case "hello":
            if (255 <= uint16(v32)) {
                for fi34 := 0; fi34 < 15; fi34++ {
                    fmt.Println(v3, v32)
                    fmt.Println(v3)
                    _ = fi34
                }
                v5 += 1
                var v35 int = v3
                _ = v35
            }
        case "":
            v36 := v1[0]
            _ = v36
        }
        if (false && false) {
            v5 = (v5 ^ v5)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
