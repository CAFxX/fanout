package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 bool; F1 uint; F2 uint16 }
func genFn0() int16 {
    v4 := GenS1{F0: 32767}
    var v5 string = "\n\t"
    for fi6 := 0; fi6 < 19; fi6++ {
        v5 = v5
        v7 := GenS1{F0: 32767}
        _ = fi6
        _ = v7
    }
    var v8 uint8 = 255
    v8 -= v8
    _ = v4
    _ = v5
    _ = v8
    return 32767
}
func genFn1(p0 uint32, p1 int, p2 uint) bool {
    var v9 uint64 = 4294967295
    p1 = p1
    fmt.Println(p2, p2)
    var v10 uint32 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return ((v9 & (^(v9))) > 0)
}

func main() {
    var v1 uint32 = 2
    v2 := []uint32{2}
    v3 := make(map[string]bool)
    switch 295 {
    case -2147483648:
        switch "a" {
        case "a":
            v1 = v1
        case "\n\t":
            var v11 int8 = 0
            _ = v11
        default:
            v1 -= v1
        }
    case -414:
        v12 := GenS2{F0: 42}
        _ = v12
    default:
        var v13 int32 = -2147483648
        _ = v13
    }
    var v14 float64 = 1e10
    v14 = -2.25
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v14 = (v14 + ((((v14 + float64(0.0)) / -2.25) / float64(((0.1 * ((((1.5 / (0.0 / (((-2.25 + float32(((v14 * ((v14 / (v14 + v14)) - 1e10)) + (0.0 - (((float64((3.14159 * -2.25)) / v14) - (((v14 * v14) / (-2.25 - float64(0.1))) * v14)) / -2.25))))) * 1.5) * (3.14159 / 1.0)))) / (3.14159 + 0.1)) + -2.25) + 0.0)) / (0.0 + 1e10)))) / float64(1.5)))
    for fi15 := 0; fi15 < 5; fi15++ {
        if len(v2) > 0 {
            v2[0] = (v1 - v1)
        }
        if ("" <= "hello") {
            var v16 int8 = 0
            v17 := GenS3{F0: (v16 - v16), F1: fi15}
            var v18 int16 = -728
            _ = v16
            _ = v17
            _ = v18
        } else {
            v3["hello"] = (0.1 >= 0.0)
            for fi19 := 0; fi19 < 7; fi19++ {
                _ = len(v3)
                _ = fi19
            }
        }
        _ = fi15
    }
    v14 += (v14 + (float64(0.1) - (-2.25 - v14)))
    if true {
        v14 = (((v14 / (1.5 - v14)) + (float64(1e10) - -2.25)) + 1.5)
        v20 := GenS2{F0: 127}
        _ = v20
    } else {
        var v21 uint8 = 2
        if false {
            var v22 int8 = 94
            for fi23 := 0; fi23 < 8; fi23++ {
                v22 -= v22
                if ((v21 * v21) <= v21) {
                    var v24 int16 = -835
                    if true {
                        switch (fi23 << ((uint(v21)) & 63)) {
                        case 9223372036854775807:
                            v1 += (v1 ^ v1)
                        case 127:
                            var v25 uint64 = 2
                            _ = v25
                        case -1:
                            v3[("a" + "")] = false
                        }
                    } else {
                        if false {
                            var v26 uint32 = 42
                            _ = v26
                        }
                        if true {
                            _ = len(v2)
                        } else {
                            v27 := v3["a"]
                            if len(v2) > 0 {
                                v2[0] = 60
                            }
                            _ = v27
                        }
                    }
                    _ = v24
                } else {
                    v28 := GenS1{F0: 1}
                    var v29 uint16 = 567
                    _ = v28
                    _ = v29
                }
                var v30 int8 = v22
                _ = fi23
                _ = v30
            }
            _ = v22
        }
        if (!((428 <= 1) && true)) {
            var v31 uint8 = v21
            _ = v31
        } else {
            if false {
                var v32 float64 = 0.0
                v33 := GenS4{F0: ((!false) && false), F1: 1, F2: 32767}
                _ = v32
                _ = v33
            }
            fmt.Println(v1, v1)
        }
        _ = v21
    }
    v34 := GenS2{F0: 1}
    var v35 int16 = 1000
    var v36 uint16 = 42
    fmt.Println(v14, v14)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
}
