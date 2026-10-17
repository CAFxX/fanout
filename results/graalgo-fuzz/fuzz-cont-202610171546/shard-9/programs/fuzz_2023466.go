package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }

func main() {
    var v1 int = -9223372036854775808
    var v2 float64 = 1e10
    v3 := make(map[int]uint32)
    for fi4 := 0; fi4 < 2; fi4++ {
        var v5 uint64 = 2
        v6 := GenS1{F0: -83, F1: 1}
        if true {
            switch "\n\t" {
            case "hello":
                for fi7 := 0; fi7 < 7; fi7++ {
                    v2 = float64(0.1)
                    fmt.Println(fi7, v2)
                    _ = fi7
                }
            case "\n\t":
                var v8 int64 = -908
                _ = v8
            default:
                for fi9 := 0; fi9 < 11; fi9++ {
                    fmt.Println(v1)
                    if (false && false) {
                        var v10 uint64 = (v5 << ((18446744073709551615) & 63))
                        fmt.Println(fi4)
                        v5 = v10
                        _ = v10
                    } else {
                        var v11 float32 = -2.25
                        if false {
                            v5 = v5
                            var v12 uint = 458
                            _ = v12
                        }
                        if false {
                            v13 := []int{fi9, (v1 & (v1 >> ((255) & 63)))}
                            _ = v13
                        } else {
                            switch "hello" {
                            case "foo bar":
                                if (fi4 > (fi9 - 42)) {
                                    var v14 int32 = 0
                                    var v15 int64 = 0
                                    _ = v14
                                    _ = v15
                                }
                            case "":
                                if false {
                                    v1 = -1
                                    v1 -= (fi9 + fi9)
                                }
                            }
                            v11 = float32(1.5)
                        }
                        _ = v11
                    }
                    _ = fi9
                }
            }
        }
        _ = fi4
        _ = v5
        _ = v6
    }
    v2 += float64((3.14159 + 0.0))
    v2 += v2
    v16 := GenS1{F0: -128, F1: 0}
    v2 = (((v2 * 1.5) + v2) + (v2 + float64(1.5)))
    var v17 int64 = 0
    var v18 float32 = 1.5
    v19 := GenS1{F0: int8(v17), F1: 127}
    var v20 bool = true
    v17 = v17
    fmt.Println(v1, v18, v1)
    var v21 uint16 = 32767
    var v22 bool = v20
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
