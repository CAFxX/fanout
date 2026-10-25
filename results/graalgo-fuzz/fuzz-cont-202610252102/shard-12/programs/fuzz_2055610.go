package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 uint16; F1 uint16 }
func genFn0() uint64 {
    var v5 float64 = ((-2.25 / -2.25) / 0.1)
    var v6 float64 = 1.5
    _ = v5
    _ = v6
    return 4294967295
}
func genFn1(p0 uint, p1 string, p2 uint32) int16 {
    p1 = p1
    p1 = p1
    var v7 uint16 = 32767
    v8 := GenS1{F0: 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 1000
}

func main() {
    var v1 uint64 = 759
    v2 := make(map[int]string)
    var v3 float64 = float64(1e10)
    v4 := make(map[uint32]uint32)
    if (2 <= 42) {
        if false {
            if (127 > 31) {
                if (true || (((851 > 65535) || false) || false)) {
                    v1 += v1
                } else {
                    v3 = ((float64((0.1 * 0.0)) + (v3 / -2.25)) - v3)
                    var v9 int16 = genFn1(447, "a", 2147483647)
                    fmt.Println(v3, v3, v9)
                    _ = v9
                }
            } else {
                v10 := GenS2{F0: 128, F1: 42, F2: 4294967295}
                _ = v10
            }
        } else {
            var v11 float32 = 1.5
            switch -251 {
            case -2147483648:
                v12 := []int{127, 9223372036854775807, 32768, 32768}
                _ = v12
            case 42:
                if false {
                    var v13 uint16 = 32767
                    _ = v13
                }
            case 128:
                if true {
                    if false {
                        v14 := GenS3{F0: 710, F1: 42}
                        var v15 float64 = (v3 - (v3 + v3))
                        v16 := GenS2{F0: 128, F1: 255, F2: 959}
                        _ = v14
                        _ = v15
                        _ = v16
                    } else {
                        var v17 string = "foo bar"
                        _ = v17
                    }
                    if (!(2147483647 < 570)) {
                        for fi18 := 0; fi18 < 10; fi18++ {
                            v11 = v11
                            for fi19 := 0; fi19 < 7; fi19++ {
                                v20 := []int{(fi19 | (fi18 ^ fi19)), fi18, fi18}
                                _ = fi19
                                _ = v20
                            }
                            var v21 uint16 = 709
                            _ = fi18
                            _ = v21
                        }
                        v3 = v3
                        fmt.Println(v3)
                    }
                    var v22 uint8 = 1
                    _ = v22
                } else {
                    v23 := GenS1{F0: 42}
                    v24 := []int{545, 32767, 813}
                    _ = v23
                    _ = v24
                }
            }
            _ = v11
        }
        v25 := GenS4{F0: 0, F1: 255}
        var v26 bool = (false || true)
        _ = v25
        _ = v26
    }
    if false {
        v27 := GenS1{F0: 0}
        if (true || true) {
            _ = len(v4)
            fmt.Println(v1, v3)
            v3 += v3
        }
        _ = v27
    } else {
        if true {
            v28 := []int{32768}
            _ = v28
        }
    }
    v29 := GenS2{F0: 72, F1: 0, F2: 2}
    var v30 float32 = -2.25
    var v31 string = "a"
    var v32 bool = false
    var v33 int16 = 1000
    v1 = 2
    v32 = (!v32)
    var v34 float32 = v30
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
