package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 bool }

func main() {
    var v1 int64 = -17
    var v2 float32 = (0.0 / ((1.5 - 0.1) + ((3.14159 * 1.5) * 3.14159)))
    var v3 int = 0
    v4 := make(map[uint32]string)
    var v5 uint8 = 78
    var v6 uint32 = 0
    v7 := v4[v6]
    _ = len(v4)
    for fi8 := 0; fi8 < 2; fi8++ {
        v3 = v3
        v9 := GenS1{F0: (v5 << ((uint(v6)) & 63))}
        for fi10 := 0; fi10 < 9; fi10++ {
            v6 -= 2
            if false {
                var v11 uint8 = 42
                v5 -= 1
                v12 := []int{}
                _ = v11
                _ = v12
            } else {
                if true {
                    var v13 uint16 = 2
                    var v14 int32 = 174
                    var v15 uint64 = 280
                    _ = v13
                    _ = v14
                    _ = v15
                } else {
                    v1 -= int64(v6)
                    if (4294967295 > 4294967295) {
                        v2 -= v2
                        v16 := GenS2{F0: v6}
                        _ = v16
                    } else {
                        var v17 uint = 0
                        var v18 bool = (uint8(v17) >= v5)
                        if (v18 && true) {
                            fmt.Println(v6, v1)
                            var v19 uint16 = 695
                            v6 = uint32(v5)
                            _ = v19
                        } else {
                            if true {
                                v20 := []int{}
                                _ = v20
                            }
                            var v21 string = ((v7 + ("foo bar" + "")) + (v7 + v7))
                            var v22 int8 = -118
                            _ = v21
                            _ = v22
                        }
                        _ = v17
                        _ = v18
                    }
                    v3 -= fi8
                }
                v23 := GenS3{F0: 881, F1: (4294967295 > 2)}
                _ = v23
            }
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    var v24 float32 = v2
    v4[v6] = "foo bar"
    for fi25 := 0; fi25 < 18; fi25++ {
        v26 := []int{}
        _ = fi25
        _ = v26
    }
    var v27 uint = 2
    var v28 int32 = 42
    v29 := []int{(v3 % ((622) | 1)), (v3 >> ((v27) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
