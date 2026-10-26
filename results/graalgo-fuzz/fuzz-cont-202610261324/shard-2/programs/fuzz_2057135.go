package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 string; F1 int8; F2 int64 }
func genFn0(p0 uint) string {
    if (("hello" != "a") && true) {
        var v5 string = ("a" + "\n\t")
        v6 := []int{128, -9223372036854775808, 127}
        var v7 bool = false
        _ = v5
        _ = v6
        _ = v7
    }
    fmt.Println(p0)
    v8 := GenS1{F0: -51, F1: -32768}
    _ = p0
    _ = v8
    return "foo bar"
}

func main() {
    v1 := make(map[int]int)
    var v2 float64 = float64((float32(float64(-2.25)) * ((float32(-2.25) * ((1.5 - (float32(0.0) - 0.1)) + ((3.14159 * 0.0) / -2.25))) + (float32(0.1) + float32(1.5)))))
    v3 := make(map[uint32]string)
    var v4 int = -493
    switch ("" + "a") {
    case "a":
        for fi9 := 0; fi9 < 5; fi9++ {
            v4 = (fi9 & fi9)
            fmt.Println(v2)
            _ = fi9
        }
    case "foo bar":
        if false {
            switch (^(v4)) {
            case 1:
                v10 := []int{(v4 << ((255) & 63)), (v4 | (v4 / (((v4 | 0)) | 1))), 0}
                _ = v10
            default:
                var v11 float64 = (v2 + float64(float32((0.1 * float64((float32(((float64(3.14159) - (0.1 * v2)) + 1e10)) + (-2.25 / float32((v2 - v2)))))))))
                _ = v11
            }
            var v12 int64 = -9223372036854775808
            _ = v12
        } else {
            var v13 float64 = ((v2 / (v2 + 1.5)) - (-2.25 + v2))
            var v14 int16 = 1
            delete(v3, 0)
            _ = v13
            _ = v14
        }
    case "\n\t":
        v4 -= v4
    default:
        v2 -= v2
    }
    fmt.Println(v4, v2)
    for fi15 := 0; fi15 < 12; fi15++ {
        switch "\n\t" {
        case "":
            var v16 bool = ((v4 & -744) < fi15)
            _ = v16
        case "a":
            v17 := []int{(fi15 / (((v4 / ((v4) | 1))) | 1)), 4294967295, v4, fi15}
            _ = v17
        default:
            for fi18 := 0; fi18 < 16; fi18++ {
                if true {
                    v19 := v3[66]
                    _ = v19
                }
                if true {
                    switch (v4 * (fi15 / ((fi15) | 1))) {
                    case 0:
                        if ((9223372036854775807 != -9223372036854775808) && ((255 <= 2) && true)) {
                            v2 = -2.25
                            var v20 bool = false
                            v2 = v2
                            _ = v20
                        }
                    case 1:
                        if false {
                            v21 := []int{4294967295, 9223372036854775807, 42}
                            v4 = 4294967295
                            if len(v21) >= 2 {
                                v21 = v21[1:len(v21)]
                            }
                            _ = v21
                        } else {
                            if (0 <= 0) {
                                v22 := GenS2{F0: (("\n\t" + "a") + "a"), F1: 87, F2: 42}
                                v4 -= (v4 & 32767)
                                switch fi15 {
                                case 464:
                                    v2 += v2
                                case -1:
                                    var v23 uint16 = 946
                                    _ = v23
                                default:
                                    fmt.Println(v4, fi18)
                                }
                                _ = v22
                            } else {
                                switch -2147483648 {
                                case 1:
                                    var v24 string = ("foo bar" + "\n\t")
                                    _ = v24
                                case 4294967295:
                                    var v25 int64 = 0
                                    _ = v25
                                case -532:
                                    v2 += ((3.14159 / ((v2 / 1.5) * v2)) - 0.1)
                                default:
                                    v2 = (3.14159 - (v2 - ((v2 * 3.14159) / 1e10)))
                                }
                                var v26 uint64 = 0
                                _ = v26
                            }
                            v2 = 1e10
                        }
                    }
                    v4 = fi18
                }
                if (2 <= 255) {
                    v27 := v3[404]
                    _ = v27
                }
                _ = fi18
            }
        }
        _ = fi15
    }
    var v28 float64 = (v2 / v2)
    v1[(v4 >> ((255) & 63))] = (-(v4))
    v3[42] = (("foo bar" + ("\n\t" + "\n\t")) + "a")
    var v29 int16 = -1
    switch "a" {
    case "":
        v28 += (v2 - 1e10)
    default:
        v28 += v2
    }
    v30 := GenS2{F0: ("foo bar" + "a"), F1: 60, F2: 0}
    v29 += 32767
    delete(v3, 65535)
    if (1 > -1) {
        if ((^(v29)) >= -32768) {
            v31 := []int{}
            if len(v31) >= 2 {
                v31 = v31[1:len(v31)]
            }
            v32 := GenS1{F0: 42, F1: v29}
            _ = v31
            _ = v32
        }
        var v33 float64 = float64(1.5)
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
