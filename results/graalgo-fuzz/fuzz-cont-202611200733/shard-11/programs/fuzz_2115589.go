package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
func genFn0() uint32 {
    var v4 uint64 = 18446744073709551615
    v5 := []int{-2147483648, -1, 127}
    fmt.Println(v4)
    fmt.Println(v4)
    _ = v4
    _ = v5
    return 42
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int8 = 42
    v3 := make(map[string]uint32)
    fmt.Println(v2)
    var v6 uint64 = 4294967295
    if (!true) {
        for fi7 := 0; fi7 < 14; fi7++ {
            switch fi7 {
            case 42:
                v8 := GenS1{F0: 9, F1: v2}
                _ = v8
            case 32767:
                if false {
                    v6 = v6
                    for fi9 := 0; fi9 < 3; fi9++ {
                        _ = len(v3)
                        var v10 int16 = 698
                        _ = fi9
                        _ = v10
                    }
                    v2 = v2
                }
            default:
                v11 := v3[""]
                _ = v11
            }
            _ = fi7
        }
        fmt.Println(v6, v2)
    }
    v12 := []int{32767, 128, 2147483647, -1}
    v6 = 18446744073709551615
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 bool = true
        for fi15 := 0; fi15 < 10; fi15++ {
            v6 += (v6 ^ 1)
            v16 := GenS1{F0: 569, F1: v2}
            _ = fi15
            _ = v16
        }
        v14 = (fi13 <= (fi13 >> ((821) & 63)))
        _ = fi13
        _ = v14
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        for fi18 := 0; fi18 < 19; fi18++ {
            v2 -= (v2 << ((2) & 63))
            v12 = append(v12, 4294967295)
            _ = fi18
        }
        v2 = (v2 + v2)
        _ = fi17
    }
    if (3.14159 != 3.14159) {
        var v19 float32 = 1e10
        _ = v19
    } else {
        var v20 uint64 = v6
        _ = v20
    }
    fmt.Println(v2, v6)
    for fi21 := 0; fi21 < 7; fi21++ {
        var v22 int16 = int16(v6)
        if len(v12) > 0 {
            v12[0] = fi21
        }
        for fi23 := 0; fi23 < 7; fi23++ {
            fmt.Println(v22, fi21, v6)
            _ = fi23
        }
        _ = fi21
        _ = v22
    }
    for fi24 := 0; fi24 < 7; fi24++ {
        var v25 uint64 = (v6 << ((18446744073709551615) & 63))
        var v26 int8 = v2
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v6
}
