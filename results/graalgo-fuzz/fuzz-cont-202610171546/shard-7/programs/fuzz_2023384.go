package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint16 }
func genFn0() string {
    var v6 int32 = 42
    switch 2147483647 {
    case 930:
        v6 -= 0
    case -425:
        switch "a" {
        case "":
            var v7 int8 = -1
            _ = v7
        default:
            if (false || ((!false) && (!true))) {
                v8 := []int{32767, -2147483648}
                for fi9 := 0; fi9 < 18; fi9++ {
                    switch 1 {
                    case -11:
                        if len(v8) >= 2 {
                            v8 = v8[1:len(v8)]
                        }
                    case 127:
                        var v10 bool = (1 != 0)
                        _ = v10
                    }
                    for fi11 := 0; fi11 < 13; fi11++ {
                        if true {
                            _ = len(v8)
                        } else {
                            fmt.Println(fi9, v6, v6)
                            v6 -= (v6 & 2147483647)
                            if (73 == 100) {
                                v12 := GenS1{F0: 0, F1: 230, F2: 237}
                                v6 = (v6 * v6)
                                v13 := GenS2{F0: 32767, F1: 202, F2: 90}
                                _ = v12
                                _ = v13
                            }
                        }
                        if true {
                            fmt.Println(fi11, fi9)
                            var v14 float32 = 1e10
                            _ = v14
                        } else {
                            var v15 int64 = 42
                            var v16 int = fi9
                            var v17 uint8 = 255
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                        fmt.Println(fi9)
                        _ = fi11
                    }
                    _ = fi9
                }
                fmt.Println(v6)
                _ = v8
            }
        }
    }
    for fi18 := 0; fi18 < 11; fi18++ {
        v6 = 0
        _ = fi18
    }
    v19 := []int{}
    _ = v19
    _ = v6
    return "\n\t"
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]bool)
    var v4 int8 = 0
    var v5 float32 = -2.25
    v4 = -128
    for fi20 := 0; fi20 < 7; fi20++ {
        var v21 uint64 = 42
        fmt.Println(v5, fi20)
        v22 := []int{(^(fi20)), (^(fi20)), fi20}
        _ = fi20
        _ = v21
        _ = v22
    }
    v4 = (v4 >> ((36) & 63))
    fmt.Println(v4, v5)
    var v23 uint32 = 1
    v24 := []int{-1, 323, -890}
    _ = len(v24)
    var v25 uint64 = 18446744073709551615
    var v26 float32 = (v5 + (float32(1.5) * ((float32(0.0) * v5) + 0.1)))
    v23 = (v23 << ((1) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
