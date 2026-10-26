package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 string }
type GenS2 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS3 struct { F0 int32; F1 int64; F2 string }
type GenS4 struct { F0 int64; F1 uint32; F2 int }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 uint8; F1 int }

func main() {
    v1 := make(map[int]string)
    var v2 float64 = (-2.25 - 1.5)
    var v3 uint32 = 0
    var v4 uint8 = 0
    var v5 uint64 = 982
    var v6 float32 = 1.5
    for fi7 := 0; fi7 < 10; fi7++ {
        if true {
            v8 := GenS1{F0: 42, F1: v3, F2: "\n\t"}
            v9 := GenS1{F0: -99, F1: v3, F2: "\n\t"}
            switch "" {
            case "a":
                fmt.Println(fi7, v3)
            case "\n\t":
                if true {
                    v4 = (v4 / ((v4) | 1))
                    switch "a" {
                    case "a":
                        var v10 uint16 = 32767
                        _ = v10
                    case "":
                        if false {
                            var v11 uint = 18446744073709551615
                            if true {
                                v12 := GenS1{F0: 42, F1: (v3 - (v3 >> ((v11) & 63))), F2: "a"}
                                v13 := GenS2{F0: v3, F1: -356, F2: 255}
                                _ = v12
                                _ = v13
                            } else {
                                v2 -= 1.5
                                var v14 int8 = 127
                                _ = v14
                            }
                            var v15 int16 = 1
                            _ = v11
                            _ = v15
                        } else {
                            var v16 int8 = 42
                            _ = v16
                        }
                    }
                    fmt.Println(v2, v5)
                } else {
                    delete(v1, fi7)
                    var v17 bool = false
                    _ = v17
                }
            default:
                v2 = v2
            }
            _ = v8
            _ = v9
        } else {
            var v18 int16 = 0
            v19 := GenS3{F0: -1, F1: 0, F2: (("\n\t" + "hello") + "a")}
            for fi20 := 0; fi20 < 16; fi20++ {
                v18 = v18
                fmt.Println(fi7)
                _ = fi20
            }
            _ = v18
            _ = v19
        }
        v4 = (v4 | 21)
        _ = fi7
    }
    var v21 int8 = -1
    v22 := GenS4{F0: -9223372036854775808, F1: 689, F2: 0}
    var v23 uint16 = 255
    v24 := GenS5{F0: v3, F1: (1000 > 42)}
    for fi25 := 0; fi25 < 5; fi25++ {
        var v26 int8 = v21
        switch fi25 {
        case 127:
            v6 = (0.0 / 1.0)
        case 4294967295:
            v2 = (v2 + ((v2 - v2) * v2))
        case -871:
            v27 := GenS2{F0: uint32(v4), F1: 295, F2: (v23 * (v23 + v23))}
            _ = v27
        }
        _ = fi25
        _ = v26
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        v29 := GenS6{F0: uint8(v3), F1: 0}
        _ = fi28
        _ = v29
    }
    v21 = (v21 ^ (^(v21)))
    for fi30 := 0; fi30 < 18; fi30++ {
        v31 := []int{fi30, fi30, int(v23)}
        v23 = v23
        _ = fi30
        _ = v31
    }
    var v32 uint8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
