package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int; F1 uint; F2 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 uint32 }
func genFn0(p0 int16, p1 int, p2 int64) uint {
    var v7 float64 = (3.14159 + 1.5)
    var v8 int8 = 15
    p1 = (p1 | (p1 >> ((255) & 63)))
    v9 := GenS1{F0: (p0 + p0), F1: 127}
    var v10 float64 = float64((((1e10 - 3.14159) - 1.5) / (3.14159 - float32(float64(((-2.25 / 1.0) - 1.5))))))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 760
}

func main() {
    var v1 uint = 0
    v2 := []int{-9223372036854775808, 32767, 128, 42}
    v3 := []uint32{65535, 42}
    var v4 uint32 = 1
    var v5 float64 = 1.5
    v6 := make(map[string]bool)
    var v11 bool = (v1 > 402)
    var v12 int = 42
    switch ("hello" + "a") {
    case "\n\t":
        if v11 {
            v13 := GenS2{F0: v12, F1: (v1 - (^(v1))), F2: v12}
            _ = v13
        } else {
            v4 = uint32(v1)
        }
    case "":
        var v14 float64 = 0.0
        _ = v14
    }
    if v11 {
        var v15 uint32 = v4
        _ = v15
    } else {
        v11 = (!(0 >= 0))
    }
    if len(v2) > 0 {
        v2[0] = v12
    }
    v1 += (v1 ^ 0)
    v16 := GenS2{F0: v12, F1: v1, F2: v12}
    var v17 float32 = ((0.1 * 3.14159) / 1e10)
    if true {
        switch "foo bar" {
        case "\n\t":
            for fi18 := 0; fi18 < 14; fi18++ {
                if v11 {
                    v5 = 0.0
                } else {
                    if false {
                        var v19 float64 = ((3.14159 * v5) + (v5 / 1.5))
                        v20 := GenS1{F0: 32767, F1: 2}
                        v4 = v4
                        _ = v19
                        _ = v20
                    }
                }
                _ = fi18
            }
        }
    } else {
        fmt.Println(v12)
        if v11 {
            v17 = (-2.25 * (((v17 - float32(0.1)) - float32(v5)) + (float32(v5) * v17)))
            switch (v12 * (v12 * (v12 % ((v12) | 1)))) {
            case 221:
                v21 := GenS3{F0: (v4 | (v4 / ((65535) | 1)))}
                _ = v21
            case -20:
                if ((-459 > (v12 & (v12 << (((v1 >> ((100) & 63))) & 63)))) || v11) {
                    v6[(("foo bar" + "a") + "foo bar")] = (v1 == (^(v1)))
                    for fi22 := 0; fi22 < 14; fi22++ {
                        delete(v6, "hello")
                        for fi23 := 0; fi23 < 6; fi23++ {
                            if (v11 && v11) {
                                v24 := GenS3{F0: 42}
                                _ = v24
                            } else {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                            switch (v12 | v12) {
                            case 326:
                                _ = len(v6)
                            case 1:
                                v25 := GenS1{F0: -1, F1: 192}
                                _ = v25
                            default:
                                switch (v12 & fi22) {
                                case 42:
                                    var v26 int32 = -28
                                    _ = v26
                                }
                            }
                            fmt.Println(v11)
                            _ = fi23
                        }
                        _ = fi22
                    }
                }
            case 32768:
                v6[""] = (-1 != -83)
            default:
                var v27 int8 = -61
                _ = v27
            }
        } else {
            fmt.Println(v4, v5)
            v28 := GenS4{F0: "foo bar", F1: uint32(v1)}
            _ = v28
        }
    }
    var v29 uint16 = 410
    var v30 int8 = -70
    v31 := GenS1{F0: 42, F1: 255}
    var v32 float64 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
