package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int16; F2 uint8 }
func genFn0() int32 {
    var v7 uint16 = 510
    fmt.Println(v7)
    v8 := GenS1{F0: 0}
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 bool = (!false)
        var v11 int64 = 848
        v12 := []int{(fi9 / ((fi9) | 1)), (fi9 + (-(fi9)))}
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    _ = v7
    _ = v8
    return 0
}

func main() {
    v1 := []int{640}
    var v2 float32 = 0.1
    var v3 int = 32767
    v4 := []string{"foo bar", "\n\t"}
    var v5 uint64 = 42
    var v6 int = v3
    for fi13 := 0; fi13 < 8; fi13++ {
        v6 += int(v5)
        v5 = 2
        v1 = append(v1, 32768)
        _ = fi13
    }
    v3 = -9223372036854775808
    var v14 uint32 = 65535
    var v15 uint16 = 1
    switch 32767 {
    case -198:
        for fi16 := 0; fi16 < 10; fi16++ {
            v6 -= fi16
            fmt.Println(v2, v5)
            _ = fi16
        }
    case -2147483648:
        switch 0 {
        case -2147483648:
            var v17 int = v3
            _ = v17
        case -328:
            var v18 uint = 668
            _ = v18
        default:
            for fi19 := 0; fi19 < 9; fi19++ {
                _ = len(v4)
                if true {
                    var v20 int8 = -128
                    if true {
                        for fi21 := 0; fi21 < 1; fi21++ {
                            if ((0 == (v14 << ((661) & 63))) || false) {
                                if false {
                                    v20 += 127
                                }
                            }
                            v2 -= (0.0 * v2)
                            _ = fi21
                        }
                        var v22 int32 = 1
                        _ = v22
                    } else {
                        var v23 int8 = (v20 << ((840) & 63))
                        _ = v23
                    }
                    var v24 uint64 = uint64(v15)
                    _ = v20
                    _ = v24
                }
                fmt.Println(v6, v15)
                _ = fi19
            }
        }
    }
    fmt.Println(v2, v2, v2)
    v25 := GenS2{F0: 186, F1: 643, F2: 127}
    fmt.Println(v14)
    v3 = v6
    fmt.Println(v5, v3, v2)
    v3 = (v3 | (v6 * (-(v6))))
    v26 := GenS1{F0: (v15 / ((uint16(v5)) | 1))}
    if false {
        v2 += (3.14159 + float32(-2.25))
        var v27 int8 = -1
        fmt.Println(v5, v6, v6)
        _ = v27
    }
    var v28 uint8 = 2
    v6 = v6
    for fi29 := 0; fi29 < 4; fi29++ {
        fmt.Println(v3, v5, v3)
        for fi30 := 0; fi30 < 2; fi30++ {
            v2 = v2
            _ = fi30
        }
        _ = fi29
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
