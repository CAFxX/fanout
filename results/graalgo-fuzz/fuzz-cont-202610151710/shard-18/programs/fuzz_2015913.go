package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint32, p1 int16) uint {
    var v7 int16 = p1
    var v8 string = "foo bar"
    var v9 int32 = 0
    p1 -= p1
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 bool) int {
    v10 := GenS1{F0: -1}
    var v11 float32 = -2.25
    _ = p0
    _ = v10
    _ = v11
    return -1
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 0
    var v3 int = 128
    v4 := make(map[uint32]uint32)
    v5 := []int{(v3 | (^(v3))), v3, (v3 % (((v3 >> ((0) & 63))) | 1)), v3}
    var v6 float64 = (((3.14159 + 3.14159) - 1e10) - 0.0)
    v5 = append(v5, genFn1((v2 != (v2 % (((v2 % ((v2) | 1))) | 1)))))
    var v12 int64 = 42
    v12 += 192
    for fi13 := 0; fi13 < 6; fi13++ {
        for fi14 := 0; fi14 < 4; fi14++ {
            if (0.1 > 1.5) {
                var v15 uint64 = 4294967295
                var v16 int8 = 85
                _ = v15
                _ = v16
            }
            v2 = v2
            var v17 float64 = v6
            _ = fi14
            _ = v17
        }
        _ = fi13
    }
    v18 := GenS2{F0: 127, F1: 109, F2: -48}
    v5 = append(v5, 0)
    for fi19 := 0; fi19 < 19; fi19++ {
        switch (fi19 & (v3 % ((int(v2)) | 1))) {
        case -789:
            for fi20 := 0; fi20 < 10; fi20++ {
                if (!(v6 >= float64(-2.25))) {
                    if len(v5) > 0 {
                        v5[0] = (v3 * v3)
                    }
                    if true {
                        var v21 float32 = 0.0
                        _ = v21
                    }
                    var v22 float32 = (0.0 + 1e10)
                    _ = v22
                }
                v12 = int64(fi20)
                v23 := v4[(v2 - 4294967295)]
                _ = fi20
                _ = v23
            }
        case -1:
            if len(v5) > 0 {
                v5[0] = v3
            }
        case -917:
            if false {
                fmt.Println(v3, v12)
            } else {
                v2 += v2
                for fi24 := 0; fi24 < 10; fi24++ {
                    if (42 > 231) {
                        v12 = (v12 << ((860) & 63))
                        if false {
                            v25 := GenS3{F0: v2}
                            if true {
                                var v26 float64 = 3.14159
                                v27 := GenS1{F0: v3}
                                v2 += (v2 | v2)
                                _ = v26
                                _ = v27
                            }
                            _ = v25
                        }
                    } else {
                        v28 := GenS2{F0: 127, F1: 25, F2: 92}
                        _ = v28
                    }
                    _ = fi24
                }
                fmt.Println(v6, v6)
            }
        default:
            v3 = fi19
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v29 int32 = -2147483648
        _ = fi19
        _ = v29
    }
    _ = len(v4)
    v4[v2] = 1
    v30 := GenS1{F0: (v3 + v3)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
