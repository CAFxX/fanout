package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint8; F1 int8; F2 uint32 }
func genFn0() int8 {
    var v6 bool = true
    v6 = v6
    v7 := []int{-142, -603, 4294967295}
    _ = v6
    _ = v7
    return 1
}

func main() {
    v1 := []string{"foo bar", "\n\t"}
    var v2 int64 = 210
    var v3 int8 = 62
    v4 := make(map[int]int)
    var v5 int32 = 246
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 uint8 = 237
        _ = fi8
        _ = v9
    }
    v3 = v3
    if true {
        delete(v4, -397)
        if false {
            v3 += v3
        }
    }
    _ = len(v1)
    v3 -= genFn0()
    var v10 uint8 = 1
    v3 += (v3 | v3)
    for fi11 := 0; fi11 < 6; fi11++ {
        v5 += v5
        if true {
            _ = len(v1)
        }
        _ = fi11
    }
    if ((true && (!true)) || (false && (4294967295 >= 4294967295))) {
        v10 = v10
        v12 := GenS1{F0: 32767, F1: -753}
        _ = v12
    }
    if false {
        for fi13 := 0; fi13 < 13; fi13++ {
            v10 = (^(v10))
            if (614 >= 221) {
                if false {
                    fmt.Println(v5, v3, v3)
                    switch 305 {
                    case 1:
                        for fi14 := 0; fi14 < 13; fi14++ {
                            v2 += (v2 % ((v2) | 1))
                            v15 := v4[(fi14 << ((1) & 63))]
                            _ = fi14
                            _ = v15
                        }
                    case 936:
                        var v16 uint8 = 255
                        _ = v16
                    case 0:
                        var v17 float32 = 1.5
                        _ = v17
                    }
                } else {
                    var v18 uint64 = 0
                    _ = v18
                }
            }
            var v19 uint16 = 32767
            _ = fi13
            _ = v19
        }
        var v20 int16 = 42
        _ = v20
    }
    for fi21 := 0; fi21 < 17; fi21++ {
        var v22 uint16 = 42
        var v23 string = ("a" + "\n\t")
        _ = fi21
        _ = v22
        _ = v23
    }
    _ = len(v1)
    v24 := GenS2{F0: v10, F1: -109, F2: 682}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
