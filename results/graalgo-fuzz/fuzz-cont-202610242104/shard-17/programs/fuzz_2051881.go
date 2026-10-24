package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 uint16; F1 int8; F2 uint }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int32) int64 {
    switch 82 {
    case 127:
        for fi6 := 0; fi6 < 1; fi6++ {
            p0 -= p0
            _ = fi6
        }
    case -1:
        var v7 uint = 42
        _ = v7
    default:
        if (!(!true)) {
            v8 := GenS1{F0: 2, F1: 255}
            fmt.Println(p0)
            p0 = (p0 | (p0 << ((2) & 63)))
            _ = v8
        }
    }
    v9 := GenS2{F0: 471, F1: -71, F2: 42}
    p0 = -1
    p0 = (^(p0))
    _ = p0
    _ = v9
    return 9223372036854775807
}

func main() {
    var v1 uint8 = 90
    var v2 int32 = 833
    var v3 uint = 4294967295
    var v4 int = 0
    v5 := make(map[uint32]int)
    var v10 uint64 = 4294967295
    var v11 bool = true
    v4 = v4
    switch ("\n\t" + "a") {
    case "hello":
        var v12 uint16 = 0
        _ = v12
    case "foo bar":
        delete(v5, 938)
    case "\n\t":
        v4 -= -469
    default:
        switch -9223372036854775808 {
        case 0:
            if (!((v3 >> ((4294967295) & 63)) >= v3)) {
                for fi13 := 0; fi13 < 4; fi13++ {
                    if (4294967295 > 417) {
                        for fi14 := 0; fi14 < 3; fi14++ {
                            v1 += (v1 - v1)
                            switch fi13 {
                            case 1:
                                v15 := []int{fi13}
                                _ = v15
                            }
                            fmt.Println(v3)
                            _ = fi14
                        }
                        var v16 int32 = v2
                        _ = v16
                    } else {
                        var v17 uint64 = v10
                        var v18 uint8 = (v1 + 2)
                        _ = v17
                        _ = v18
                    }
                    var v19 float32 = 1.5
                    _ = fi13
                    _ = v19
                }
                delete(v5, 649)
                v4 -= (^(v4))
            } else {
                for fi20 := 0; fi20 < 17; fi20++ {
                    v2 = 0
                    _ = fi20
                }
                v21 := []int{}
                _ = v21
            }
        }
    }
    var v22 uint64 = v10
    v5[2] = v4
    fmt.Println(v4, v11, v22)
    v23 := GenS2{F0: 65535, F1: 127, F2: v3}
    v11 = v11
    var v24 int32 = (-(v2))
    for fi25 := 0; fi25 < 14; fi25++ {
        v26 := GenS3{F0: (v1 / ((v1) | 1))}
        v2 = v24
        _ = fi25
        _ = v26
    }
    var v27 int = v4
    var v28 float32 = float32(3.14159)
    v5[2147483648] = (v27 ^ v27)
    for fi29 := 0; fi29 < 20; fi29++ {
        if (32767 <= 1000) {
            var v30 int32 = v24
            _ = v30
        } else {
            var v31 uint8 = v1
            var v32 uint64 = (v22 >> (((^(v3))) & 63))
            v4 = (-(v27))
            _ = v31
            _ = v32
        }
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
