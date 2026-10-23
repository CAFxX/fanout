package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 int; F1 int }
func genFn0(p0 uint32, p1 string, p2 bool) uint16 {
    for fi4 := 0; fi4 < 17; fi4++ {
        var v5 float32 = (0.1 / 1e10)
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 13; fi6++ {
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 uint8 = 30
            if p2 {
                v8 = 91
                if (v8 >= 42) {
                    fmt.Println(p1)
                    v9 := GenS1{F0: p0, F1: 32767}
                    var v10 int8 = 64
                    _ = v10
                    _ = v9
                } else {
                    var v11 bool = p2
                    var v12 int32 = -1
                    var v13 float32 = 1.5
                    v12 = 0
                    _ = v11
                    _ = v12
                    _ = v13
                }
                v14 := []int{816, 32768, 127}
                _ = v14
            } else {
                var v15 uint64 = 18446744073709551615
                if (true || p2) {
                    var v16 int = (^(fi6))
                    var v17 int64 = 920
                    if p2 {
                        v8 -= (^(v8))
                        if false {
                            v17 = (v17 * v17)
                            fmt.Println(v8)
                            v18 := GenS1{F0: (p0 + p0), F1: 42}
                            _ = v18
                        } else {
                            var v19 bool = (151 != 255)
                            _ = v19
                        }
                    }
                    _ = v16
                    _ = v17
                }
                _ = v15
            }
            fmt.Println(p1)
            _ = fi7
            _ = v8
        }
        v20 := GenS1{F0: (p0 | p0), F1: 1000}
        fmt.Println(p0)
        _ = fi6
        _ = v20
    }
    var v21 int = 32767
    for fi22 := 0; fi22 < 7; fi22++ {
        for fi23 := 0; fi23 < 1; fi23++ {
            var v24 int8 = 1
            var v25 int16 = -1
            _ = fi23
            _ = v24
            _ = v25
        }
        _ = fi22
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return 255
}

func main() {
    var v1 uint8 = 127
    var v2 uint64 = 393
    var v3 int8 = -128
    var v26 string = ("" + ("foo bar" + "\n\t"))
    if false {
        for fi27 := 0; fi27 < 7; fi27++ {
            fmt.Println(v2, v2, fi27)
            v28 := GenS1{F0: 65535, F1: -58}
            v29 := []int{}
            _ = fi27
            _ = v28
            _ = v29
        }
    } else {
        v30 := GenS1{F0: 2, F1: 0}
        _ = v30
    }
    fmt.Println(v2, v3, v2)
    v31 := GenS2{F0: 42, F1: int(v3)}
    var v32 int32 = 2147483647
    var v33 bool = (v2 < (v2 & (^(v2))))
    for fi34 := 0; fi34 < 2; fi34++ {
        var v35 float64 = 0.0
        _ = fi34
        _ = v35
    }
    v36 := GenS1{F0: 0, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v36
}
