package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint8 }
type GenS2 struct { F0 int32; F1 int64; F2 uint64 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 int8; F1 int8; F2 int32 }
type GenS5 struct { F0 uint; F1 uint8; F2 int32 }

func main() {
    v1 := make(map[string]int)
    var v2 string = "\n\t"
    var v3 float32 = (0.1 * (-2.25 - 0.1))
    var v4 int = 0
    fmt.Println(v2)
    switch "a" {
    case "\n\t":
        _ = len(v1)
    default:
        var v5 int32 = -172
        _ = v5
    }
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := GenS1{F0: 0, F1: 810, F2: 67}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 bool = ((!(-9223372036854775808 == -9223372036854775808)) && (((0.0 * float64((v3 - v3))) + 1.5) >= (float64((v3 - v3)) / (-2.25 * 1.5))))
        _ = fi8
        _ = v9
    }
    v4 = (v4 >> ((65535) & 63))
    _ = len(v1)
    var v10 int16 = 0
    if true {
        var v11 uint = 42
        _ = v11
    }
    v12 := []int{}
    v3 -= 3.14159
    for fi13 := 0; fi13 < 18; fi13++ {
        if true {
            for fi14 := 0; fi14 < 10; fi14++ {
                if (false && false) {
                    if false {
                        v15 := GenS1{F0: 187, F1: 255, F2: 42}
                        v16 := GenS2{F0: 1, F1: -1, F2: 42}
                        _ = v15
                        _ = v16
                    } else {
                        v4 = v4
                    }
                    var v17 uint16 = 0
                    v18 := GenS1{F0: (v17 - 348), F1: v17, F2: 168}
                    _ = v17
                    _ = v18
                }
                var v19 bool = false
                var v20 string = v2
                v21 := v1[v20]
                _ = fi14
                _ = v19
                _ = v20
                _ = v21
            }
            if true {
                for fi22 := 0; fi22 < 8; fi22++ {
                    v23 := GenS1{F0: 375, F1: 1, F2: 127}
                    var v24 float64 = 0.0
                    fmt.Println(v10, v2, fi22)
                    if (42 > 1) {
                        v25 := GenS2{F0: -527, F1: 42, F2: 2}
                        var v26 int = fi13
                        var v27 int64 = -988
                        v28 := GenS3{F0: 756, F1: (v10 * (v10 << ((4294967295) & 63)))}
                        v29 := GenS4{F0: -1, F1: 42, F2: 0}
                        _ = v25
                        _ = v26
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    fmt.Println(fi22, v24, fi22)
                    switch 127 {
                    case 128:
                        delete(v1, v2)
                    case 2147483647:
                        fmt.Println(v4, v2)
                    case 355:
                        var v30 float32 = float32(v24)
                        _ = v30
                    }
                    if true {
                        var v31 string = "\n\t"
                        if (865 >= -2147483648) {
                            var v32 bool = (2147483648 >= 856)
                            var v33 string = v31
                            _ = v32
                            _ = v33
                        } else {
                            v34 := v1[((v31 + "hello") + (v2 + "\n\t"))]
                            _ = v34
                        }
                        var v35 int64 = -9223372036854775808
                        _ = v31
                        _ = v35
                    } else {
                        v36 := GenS1{F0: 9, F1: 198, F2: 100}
                        _ = v36
                    }
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                if len(v12) > 0 {
                    v12[0] = (^(fi13))
                }
                v37 := GenS1{F0: 65535, F1: 65535, F2: 149}
                _ = v37
            }
        } else {
            v3 -= 1.5
            var v38 float64 = 0.0
            v3 = v3
            _ = v38
        }
        for fi39 := 0; fi39 < 19; fi39++ {
            v10 = 1000
            v3 = 1e10
            _ = fi39
        }
        _ = fi13
    }
    fmt.Println(v4)
    v2 = v2
    switch "" {
    case "\n\t":
        switch v2 {
        case "hello":
            v40 := GenS5{F0: 65535, F1: 78, F2: -37}
            _ = v40
        case "a":
            var v41 float32 = -2.25
            _ = v41
        }
    default:
        fmt.Println(v3, v2, v10)
    }
    v3 += -2.25
    v2 = ((v2 + "foo bar") + v2)
    for fi42 := 0; fi42 < 13; fi42++ {
        var v43 string = v2
        _ = fi42
        _ = v43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
