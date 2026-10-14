package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool; F1 string; F2 uint32 }
func genFn0(p0 int16, p1 int) int16 {
    if false {
        v7 := []int{(-(p1))}
        _ = v7
    }
    p0 = 1000
    var v8 uint8 = 42
    v9 := GenS1{F0: v8}
    var v10 uint64 = 2
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return p0
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]string)
    v3 := []uint32{42, 42, 4294967295}
    v4 := make(map[string]uint32)
    v5 := make(map[string]string)
    var v6 float64 = -2.25
    v11 := GenS1{F0: 98}
    switch 0 {
    case -9223372036854775808:
        v6 += v6
    default:
        for fi12 := 0; fi12 < 8; fi12++ {
            v6 = (float64(-2.25) / (0.1 / v6))
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi12
        }
    }
    v6 -= (v6 / (v6 / float64(3.14159)))
    if (4294967295 == 71) {
        v1[884] = true
        if false {
            v6 = (v6 / v6)
            v6 += 3.14159
            for fi13 := 0; fi13 < 19; fi13++ {
                v6 = float64(-2.25)
                _ = fi13
            }
        }
    }
    v6 = v6
    v6 = v6
    if (0 <= 1) {
        v14 := GenS2{F0: 1, F1: 1}
        _ = v14
    }
    v6 = v6
    v15 := GenS3{F0: 297}
    if true {
        var v16 float32 = (1e10 / ((1.5 - float32(v6)) * float32(v6)))
        delete(v1, 127)
        fmt.Println(v16, v16)
        _ = v16
    } else {
        var v17 float32 = float32(float64(0.1))
        for fi18 := 0; fi18 < 9; fi18++ {
            for fi19 := 0; fi19 < 5; fi19++ {
                switch 2147483647 {
                case 2147483647:
                    v20 := GenS4{F0: false, F1: "hello", F2: 970}
                    _ = v20
                case -571:
                    var v21 string = ("" + ("\n\t" + "a"))
                    _ = v21
                }
                if (836 >= 42) {
                    var v22 int8 = 0
                    var v23 int32 = 2147483647
                    if true {
                        v17 = float32(v6)
                        if len(v3) > 0 {
                            v3[0] = 2147483647
                        }
                    } else {
                        var v24 uint = 2
                        _ = v24
                    }
                    _ = v22
                    _ = v23
                } else {
                    fmt.Println(fi18)
                    if (!((fi19 * -32) >= -2147483648)) {
                        v4[("\n\t" + "a")] = 318
                        var v25 uint64 = 18446744073709551615
                        v26 := v1[4294967295]
                        fmt.Println(v26, v6, v26)
                        v17 += (1e10 + ((-2.25 / 3.14159) / v17))
                        var v27 uint16 = 65535
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                    fmt.Println(v6, fi19, fi18)
                }
                v3 = append(v3, 42)
                _ = fi19
            }
            v17 += (0.1 - float32(v6))
            var v28 uint64 = 379
            _ = fi18
            _ = v28
        }
        v17 -= 1e10
        _ = v17
    }
    fmt.Println(v6)
    for fi29 := 0; fi29 < 6; fi29++ {
        if (32767 != 1) {
            v3 = append(v3, 1)
            if (83 < 100) {
                v6 = ((float64(float32(v6)) / v6) / 1.0)
            }
            _ = len(v3)
        }
        if ("" < "\n\t") {
            var v30 float32 = -2.25
            _ = v30
        }
        _ = fi29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
