package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int64; F1 int8 }

func main() {
    var v1 uint16 = 255
    v2 := []uint32{2}
    var v3 uint64 = uint64(v1)
    v4 := make(map[string]int)
    v5 := make(map[string]string)
    v2 = append(v2, 42)
    switch 128 {
    case 2147483647:
        for fi6 := 0; fi6 < 14; fi6++ {
            if ((!true) && false) {
                v7 := GenS1{F0: -1, F1: true}
                if true {
                    v1 = (^(v1))
                }
                _ = v7
            }
            if false {
                for fi8 := 0; fi8 < 17; fi8++ {
                    _ = len(v2)
                    fmt.Println(fi8, fi6, v3)
                    _ = fi8
                }
            } else {
                if true {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    switch fi6 {
                    case 185:
                        v3 -= (v3 % (((v3 + (v3 - (v3 / ((v3) | 1))))) | 1))
                    case -9223372036854775808:
                        for fi9 := 0; fi9 < 20; fi9++ {
                            if false {
                                var v10 int32 = -2147483648
                                var v11 int32 = (v10 + v10)
                                _ = len(v2)
                                v2 = append(v2, 0)
                                _ = v10
                                _ = v11
                            }
                            v3 = (v3 ^ 918)
                            if (4294967295 > (v3 % ((v3) | 1))) {
                                v3 += v3
                                var v12 uint16 = v1
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v12
                            } else {
                                var v13 uint64 = v3
                                delete(v5, ("" + ""))
                                _ = v13
                            }
                            _ = fi9
                        }
                    default:
                        v1 = v1
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                } else {
                    var v14 int64 = 9223372036854775807
                    _ = v14
                }
            }
            delete(v5, "a")
            _ = fi6
        }
    case 127:
        fmt.Println(v3)
    }
    v15 := GenS2{F0: 365, F1: 42}
    v16 := GenS1{F0: 0, F1: (-1 > 974)}
    var v17 uint64 = 214
    v1 = (v1 & uint16(v3))
    for fi18 := 0; fi18 < 5; fi18++ {
        v17 = v17
        delete(v4, "hello")
        v3 = v3
        _ = fi18
    }
    var v19 float32 = 1.5
    var v20 int16 = -587
    var v21 float32 = 3.14159
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v5["a"] = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
