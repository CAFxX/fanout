package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 int; F1 int8 }
func genFn0(p0 uint32) string {
    var v7 uint64 = 307
    var v8 string = ""
    fmt.Println(v7, p0, v7)
    if (!(2147483647 != p0)) {
        v9 := GenS1{F0: 1}
        fmt.Println(p0)
        fmt.Println(v7, v8)
        _ = v9
    } else {
        var v10 uint32 = 293
        v10 += 412
        _ = v10
    }
    _ = p0
    _ = v7
    _ = v8
    return v8
}
func genFn1() bool {
    var v11 bool = (!("hello" >= "foo bar"))
    switch 9223372036854775807 {
    case 128:
        var v12 uint64 = 932
        _ = v12
    default:
        v11 = v11
    }
    fmt.Println(v11)
    _ = v11
    return ((3.14159 - -2.25) < float64(0.1))
}

func main() {
    v1 := []uint32{42, 4294967295}
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]string)
    v4 := make(map[int]string)
    var v5 uint32 = 0
    v6 := make(map[int]int)
    fmt.Println(v5)
    if false {
        var v13 uint16 = 515
        var v14 uint16 = 810
        _ = v13
        _ = v14
    }
    v5 -= (^(v5))
    v5 = v5
    var v15 int16 = 32767
    v15 = (v15 << ((1) & 63))
    if (2147483647 != 42) {
        if true {
            var v16 int = 2147483647
            _ = v16
        } else {
            v15 = (v15 / ((v15) | 1))
            if ((true || true) && false) {
                if ((-18 == -54) || (!false)) {
                    for fi17 := 0; fi17 < 18; fi17++ {
                        if len(v1) > 0 {
                            v1[0] = v5
                        }
                        _ = fi17
                    }
                    var v18 float32 = (-2.25 - float32((1e10 + 1e10)))
                    var v19 uint16 = 0
                    _ = v18
                    _ = v19
                } else {
                    v1 = append(v1, v5)
                    v15 = (v15 - 160)
                    for fi20 := 0; fi20 < 15; fi20++ {
                        var v21 bool = false
                        var v22 uint = 255
                        fmt.Println(v21)
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                }
                var v23 uint16 = 356
                _ = v23
            } else {
                v24 := GenS1{F0: -1}
                for fi25 := 0; fi25 < 5; fi25++ {
                    var v26 float32 = float32(0.0)
                    _ = fi25
                    _ = v26
                }
                v15 += -32768
                _ = v24
            }
            var v27 string = ""
            _ = v27
        }
    } else {
        var v28 uint32 = (v5 + 65535)
        _ = v28
    }
    v15 += (v15 + (-(v15)))
    v29 := GenS2{F0: -128, F1: 127}
    v30 := GenS1{F0: 1}
    v5 = v5
    if true {
        if false {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        v31 := GenS3{F0: -2147483648, F1: -1}
        _ = v31
    } else {
        fmt.Println(v5, v5)
    }
    var v32 uint16 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
