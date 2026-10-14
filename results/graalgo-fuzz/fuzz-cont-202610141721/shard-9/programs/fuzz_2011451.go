package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint8; F1 uint }

func main() {
    v1 := []int{9223372036854775807, 127, 128}
    v2 := []string{"foo bar", "\n\t"}
    var v3 float64 = 0.0
    v4 := []string{"foo bar"}
    var v5 uint16 = 255
    for fi6 := 0; fi6 < 5; fi6++ {
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 uint8 = 255
            var v9 float64 = (((v3 + v3) - v3) / 1e10)
            var v10 uint = 18446744073709551615
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        v5 -= v5
        _ = fi6
    }
    v11 := GenS1{F0: v5, F1: v5}
    for fi12 := 0; fi12 < 16; fi12++ {
        v3 -= 0.1
        v5 = (v5 / (((v5 * v5)) | 1))
        v3 += 1.5
        _ = fi12
    }
    if (856 == 18446744073709551615) {
        for fi13 := 0; fi13 < 4; fi13++ {
            for fi14 := 0; fi14 < 17; fi14++ {
                if true {
                    var v15 int64 = -1
                    if false {
                        v16 := GenS1{F0: (^(v5)), F1: v5}
                        v4 = append(v4, "a")
                        _ = v16
                    } else {
                        v17 := GenS1{F0: (v5 << ((42) & 63)), F1: v5}
                        if (true || (v15 > (v15 % ((v15) | 1)))) {
                            var v18 uint16 = 65535
                            v15 -= -941
                            var v19 string = ((("\n\t" + "foo bar") + "a") + ("a" + "hello"))
                            _ = v18
                            _ = v19
                        } else {
                            v15 = (v15 >> ((243) & 63))
                            switch (("foo bar" + "a") + "a") {
                            case "":
                                if (!((v15 << ((1) & 63)) >= 42)) {
                                    _ = len(v2)
                                    fmt.Println(v15)
                                    v20 := GenS1{F0: 966, F1: 255}
                                    _ = v20
                                } else {
                                    if false {
                                        _ = len(v1)
                                        v15 = v15
                                        var v21 bool = false
                                        _ = v21
                                    } else {
                                        v5 = 1
                                    }
                                    var v22 uint32 = 0
                                    if len(v4) > 0 {
                                        v4[0] = "foo bar"
                                    }
                                    _ = v22
                                }
                            }
                        }
                        _ = v17
                    }
                    switch (fi14 | (fi14 / ((fi13) | 1))) {
                    case 32767:
                        var v23 string = "a"
                        _ = v23
                    }
                    _ = v15
                }
                v3 += ((v3 - (float64(3.14159) + v3)) - v3)
                v3 = v3
                if (100 < 100) {
                    v3 += 0.0
                    v24 := GenS1{F0: 32767, F1: (v5 + (^(v5)))}
                    _ = v24
                }
                switch fi14 {
                case 2147483647:
                    v25 := GenS1{F0: (v5 & v5), F1: 2}
                    _ = v25
                case 127:
                    var v26 string = "hello"
                    _ = v26
                case -934:
                    v5 -= 1
                default:
                    fmt.Println(fi14, fi13, fi13)
                }
                _ = fi14
            }
            _ = len(v2)
            _ = fi13
        }
    }
    var v27 float32 = 1e10
    var v28 uint16 = v5
    for fi29 := 0; fi29 < 1; fi29++ {
        for fi30 := 0; fi30 < 12; fi30++ {
            v28 += v5
            switch fi30 {
            case -446:
                v31 := GenS1{F0: v28, F1: 723}
                var v32 float32 = v27
                fmt.Println(v32)
                _ = v31
                _ = v32
            }
            v27 -= (float32(-2.25) * 1e10)
            _ = fi30
        }
        v33 := GenS2{F0: 128, F1: 0}
        switch (fi29 & 127) {
        case 127:
            fmt.Println(v27)
        default:
            var v34 uint8 = 62
            _ = v34
        }
        _ = fi29
        _ = v33
    }
    var v35 uint32 = 4294967295
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
