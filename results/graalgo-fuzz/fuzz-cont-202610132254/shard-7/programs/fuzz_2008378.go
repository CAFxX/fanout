package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 int8; F1 uint16; F2 uint }

func main() {
    var v1 uint = 4294967295
    var v2 float64 = float64(float32((((1.5 * (-2.25 * -2.25)) - float64(float32((-2.25 - (1e10 - float64(1e10)))))) / 0.1)))
    var v3 uint32 = 2147483648
    v4 := make(map[string]uint32)
    for fi5 := 0; fi5 < 7; fi5++ {
        v2 = v2
        _ = fi5
    }
    v3 -= v3
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 int32 = -112
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 19; fi8++ {
        for fi9 := 0; fi9 < 2; fi9++ {
            v2 -= -2.25
            v10 := []int{619, fi9}
            _ = len(v4)
            _ = fi9
            _ = v10
        }
        var v11 float32 = 0.1
        _ = fi8
        _ = v11
    }
    fmt.Println(v3, v1, v1)
    v12 := []int{}
    if true {
        for fi13 := 0; fi13 < 11; fi13++ {
            v4[((((("foo bar" + "\n\t") + "\n\t") + "") + "foo bar") + "\n\t")] = v3
            _ = fi13
        }
        v14 := GenS1{F0: (v1 / (((v1 * (v1 * 0))) | 1)), F1: "\n\t"}
        var v15 uint16 = 431
        _ = v14
        _ = v15
    } else {
        v16 := GenS1{F0: (v1 >> ((65535) & 63)), F1: "hello"}
        _ = v16
    }
    fmt.Println(v3)
    v2 -= (1e10 / v2)
    if ("foo bar" < "hello") {
        if (((42 < 1) || false) || ((v2 + 1.5) >= 0.0)) {
            if true {
                v1 -= v1
            }
            if (v3 < (v3 >> ((v1) & 63))) {
                for fi17 := 0; fi17 < 18; fi17++ {
                    delete(v4, ((("" + "hello") + "a") + ""))
                    if (179 < 42) {
                        v1 += v1
                        var v18 int32 = 42
                        if true {
                            if true {
                                delete(v4, "foo bar")
                                v3 = 126
                                var v19 string = "\n\t"
                                _ = v19
                            } else {
                                fmt.Println(v2, fi17)
                                var v20 int = 407
                                v3 -= (^(v3))
                                _ = v20
                            }
                            v3 = 2147483648
                            var v21 uint16 = 2
                            _ = v21
                        }
                        _ = v18
                    } else {
                        v3 += v3
                        v1 = v1
                    }
                    v22 := v4[""]
                    _ = fi17
                    _ = v22
                }
            } else {
                if (float32(1.5) != 0.0) {
                    fmt.Println(v1)
                    v23 := GenS2{F0: (v3 | v3), F1: 127}
                    _ = v23
                }
                v3 += (v3 + v3)
                v12 = append(v12, 32767)
            }
        } else {
            var v24 int = 0
            var v25 int16 = 693
            _ = v24
            _ = v25
        }
        v26 := v4["\n\t"]
        _ = v26
    } else {
        switch 661 {
        case -1:
            _ = len(v4)
        case 2147483647:
            v27 := GenS3{F0: -1, F1: 255, F2: v1}
            _ = v27
        }
        v3 += (v3 ^ v3)
    }
    v2 = v2
    var v28 int32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
