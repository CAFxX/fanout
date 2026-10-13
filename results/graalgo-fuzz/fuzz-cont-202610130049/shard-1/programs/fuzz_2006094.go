package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 int32 }
func genFn0(p0 int, p1 uint, p2 uint) uint32 {
    p1 -= p2
    if (2 == 103) {
        v7 := GenS1{F0: 2147483648, F1: 1}
        _ = v7
    } else {
        var v8 int32 = -612
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1(p0 int) string {
    p0 = 1
    v9 := []int{p0, 93, 0, 2147483647}
    var v10 float32 = 1.5
    v10 += 3.14159
    var v11 int = p0
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return "\n\t"
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{930, 65535, 984, 2147483648}
    var v3 uint16 = 42
    var v4 int32 = 562
    var v5 float64 = (((1.5 + (float64(1e10) / (((-2.25 - (0.0 / -2.25)) - 3.14159) - 3.14159))) + 1.5) - -2.25)
    var v6 float32 = float32(float64(((1e10 + -2.25) / -2.25)))
    v4 = 2147483647
    var v12 int32 = (v4 & v4)
    v3 -= 255
    fmt.Println(v3, v3, v6)
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 float64 = ((0.1 - v5) * v5)
        if true {
            var v15 int32 = -1
            _ = v15
        }
        _ = fi13
        _ = v14
    }
    var v16 uint8 = 127
    v17 := GenS2{F0: 42, F1: -166}
    if (4294967295 > 2147483647) {
        if true {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v12 = v12
            for fi18 := 0; fi18 < 8; fi18++ {
                v6 -= (float32(v5) / float32((3.14159 * v5)))
                _ = fi18
            }
        }
        for fi19 := 0; fi19 < 17; fi19++ {
            v6 = float32(v5)
            if true {
                v6 = v6
                if (v5 > (v5 + 0.0)) {
                    if true {
                        var v20 uint32 = 0
                        for fi21 := 0; fi21 < 12; fi21++ {
                            v2 = append(v2, v20)
                            delete(v1, genFn0(int(v20), 686, 42))
                            _ = len(v1)
                            _ = fi21
                        }
                        v5 = (v5 + v5)
                        _ = v20
                    }
                    v22 := GenS1{F0: 908, F1: 77}
                    _ = v22
                } else {
                    v12 = v4
                }
            } else {
                if (3.14159 <= (v5 / (v5 - (v5 - v5)))) {
                    if true {
                        v23 := GenS2{F0: 42, F1: 935}
                        var v24 int8 = 0
                        _ = v23
                        _ = v24
                    }
                }
            }
            _ = fi19
        }
        v3 = (v3 % (((v3 & v3)) | 1))
    } else {
        for fi25 := 0; fi25 < 4; fi25++ {
            var v26 uint8 = 240
            var v27 uint32 = 2
            for fi28 := 0; fi28 < 8; fi28++ {
                if false {
                    fmt.Println(v12, v16)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v5 += v5
                }
                _ = fi28
            }
            _ = fi25
            _ = v26
            _ = v27
        }
        v29 := GenS2{F0: uint(v16), F1: 950}
        _ = v29
    }
    v5 -= ((float64(v6) / v5) - v5)
    if (257 >= uint64(v3)) {
        v3 = v3
        fmt.Println(v6)
    } else {
        v16 += 2
    }
    v30 := GenS3{F0: v12}
    v6 += v6
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
