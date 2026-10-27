package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
func genFn0(p0 uint, p1 bool, p2 uint) int16 {
    v7 := []int{127, 458, 9223372036854775807}
    p0 -= (p0 << ((65535) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 32767
}

func main() {
    var v1 uint = 2
    v2 := make(map[int]int)
    v3 := make(map[string]bool)
    v4 := []uint32{}
    var v5 float32 = 0.1
    var v6 string = "hello"
    if true {
        if (-128 != 42) {
            if false {
                if ((42 != -1) && (42 < 923)) {
                    if len(v4) > 0 {
                        v4[0] = 0
                    }
                    var v8 int = -2147483648
                    _ = v8
                } else {
                    if false {
                        for fi9 := 0; fi9 < 6; fi9++ {
                            if true {
                                v10 := GenS1{F0: 0, F1: 42}
                                var v11 bool = true
                                _ = v10
                                _ = v11
                            }
                            v2[fi9] = (fi9 | (fi9 - 685))
                            v12 := GenS1{F0: 1, F1: 4294967295}
                            _ = fi9
                            _ = v12
                        }
                    } else {
                        if true {
                            if (0 <= 1) {
                                v5 = (v5 * ((float32(0.1) - v5) - (v5 + 1.5)))
                            } else {
                                v6 = v6
                            }
                        }
                    }
                }
                v5 = (v5 - float32(-2.25))
                var v13 uint32 = 4294967295
                _ = v13
            }
            v3[("hello" + (v6 + v6))] = true
            v5 = v5
        }
    } else {
        var v14 uint64 = 18
        fmt.Println(v14, v14, v6)
        var v15 uint32 = 0
        _ = v14
        _ = v15
    }
    var v16 uint32 = 278
    v6 = v6
    var v17 uint8 = 42
    var v18 uint = 42
    var v19 uint8 = (v17 % ((uint8(v16)) | 1))
    var v20 int16 = -784
    delete(v3, v6)
    var v21 int16 = 42
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
