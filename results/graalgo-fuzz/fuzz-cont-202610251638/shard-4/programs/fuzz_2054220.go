package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 bool }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 32767
    v3 := make(map[int]int)
    v4 := make(map[int]int)
    var v5 uint16 = (v2 | (v2 ^ v2))
    var v6 uint16 = 710
    v7 := GenS1{F0: 4294967295, F1: 1000, F2: 481}
    v8 := GenS2{F0: -284}
    if (3.14159 > 1.5) {
        var v9 bool = true
        var v10 int16 = 1000
        _ = v10
        _ = v9
    }
    v5 = (^(v5))
    for fi11 := 0; fi11 < 9; fi11++ {
        if false {
            var v12 int16 = 0
            if false {
                if false {
                    for fi13 := 0; fi13 < 1; fi13++ {
                        v6 += (v2 * v6)
                        if (-1 <= v12) {
                            fmt.Println(v6)
                            v12 += 1000
                        } else {
                            var v14 uint16 = v2
                            if (false || false) {
                                var v15 int32 = 0
                                _ = v15
                            }
                            _ = v14
                        }
                        if (((115 > 1) || (9223372036854775807 > 42)) || (65535 < 961)) {
                            var v16 int = (fi13 + 4294967295)
                            var v17 uint = 1
                            v5 += (v2 >> ((0) & 63))
                            _ = v16
                            _ = v17
                        }
                        _ = fi13
                    }
                    var v18 uint16 = v2
                    v19 := []int{}
                    _ = v18
                    _ = v19
                }
                v20 := []int{9223372036854775807}
                _ = v20
            } else {
                v5 -= (v6 - 581)
                delete(v4, (fi11 << ((655) & 63)))
            }
            _ = v12
        }
        v21 := []int{fi11}
        for fi22 := 0; fi22 < 8; fi22++ {
            if (false && false) {
                v6 -= (v5 | 255)
                switch "" {
                case "a":
                    var v23 uint64 = 1
                    _ = v23
                case "foo bar":
                    v5 = 784
                case "\n\t":
                    var v24 int = fi22
                    _ = v24
                }
                var v25 float64 = (float64((0.1 - (0.1 * 1e10))) / (1e10 / 3.14159))
                v5 = v2
                _ = v25
            }
            if (1 < 205) {
                v6 += (^(v5))
            }
            _ = fi22
        }
        _ = fi11
        _ = v21
    }
    var v26 uint16 = v6
    for fi27 := 0; fi27 < 2; fi27++ {
        var v28 uint8 = 1
        _ = fi27
        _ = v28
    }
    v29 := []int{}
    var v30 uint16 = (v6 | 2)
    var v31 float32 = ((3.14159 / 0.1) - float32(1.5))
    fmt.Println(v26, v30, v2)
    var v32 bool = true
    var v33 int32 = -1
    if (800 >= 185) {
        _ = len(v29)
    } else {
        if (!v32) {
            v5 -= 973
        }
        switch 42 {
        case -1:
            var v34 int32 = (v33 + v33)
            _ = v34
        case 1:
            _ = len(v29)
        }
    }
    v35 := GenS3{F0: 9223372036854775807, F1: true}
    var v36 bool = (3.14159 != float32(float64(((3.14159 - ((v31 + v31) - v31)) - (v31 * (float32((0.1 + (0.1 + -2.25))) - float32(float64(float32((1.5 + (((float64(v31) + 0.0) + (((((float64(v31) * -2.25) / (1e10 * (0.0 + (-2.25 + 0.0)))) / 1.0) * 0.0) + 1.5)) - 0.1)))))))))))
    v33 = 857
    v37 := GenS2{F0: -300}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
