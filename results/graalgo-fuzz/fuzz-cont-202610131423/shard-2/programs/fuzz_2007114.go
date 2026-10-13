package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 bool }
func genFn0() int64 {
    var v6 uint8 = 173
    if (!(-1 >= 4294967295)) {
        switch -651 {
        case 4294967295:
            v7 := GenS1{F0: -128, F1: int16(v6), F2: (1.5 > (-2.25 + float32(0.0)))}
            _ = v7
        }
        if (true && (true && true)) {
            for fi8 := 0; fi8 < 12; fi8++ {
                var v9 uint8 = 255
                if (615 == 763) {
                    fmt.Println(fi8, v6, v9)
                    fmt.Println(v9)
                    for fi10 := 0; fi10 < 10; fi10++ {
                        if (0 >= 594) {
                            switch fi8 {
                            case 42:
                                v11 := []int{-2147483648, int(v6), (fi10 / ((fi8) | 1)), (-(fi8))}
                                _ = v11
                            }
                            v6 = v6
                        }
                        v9 = 128
                        v12 := []int{fi10, (fi8 >> ((402) & 63)), fi8, 128}
                        _ = fi10
                        _ = v12
                    }
                }
                _ = fi8
                _ = v9
            }
        }
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 int8 = 60
        _ = fi13
        _ = v14
    }
    _ = v6
    return -925
}
func genFn1(p0 int8, p1 uint16, p2 uint32) int {
    v15 := []int{956}
    p0 = (^(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 2147483647
}

func main() {
    var v1 uint64 = 42
    var v2 uint32 = 2
    v3 := []int{4294967295, -2147483648}
    v4 := []uint32{v2, 1, (v2 * v2)}
    var v5 int64 = -1
    v16 := GenS1{F0: 100, F1: 1, F2: (1 < 2)}
    var v17 bool = ("\n\t" > "\n\t")
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v18 string = ""
    v2 += v2
    if (255 < 0) {
        v17 = (0.1 < 1.5)
        fmt.Println(v17, v2, v5)
        for fi19 := 0; fi19 < 20; fi19++ {
            v5 += (v5 | (v5 % (((v5 | (v5 - genFn0()))) | 1)))
            if len(v4) > 0 {
                v4[0] = v2
            }
            _ = fi19
        }
    }
    v20 := GenS1{F0: -128, F1: 1000, F2: (255 > 14)}
    if len(v4) > 0 {
        v4[0] = v2
    }
    var v21 uint16 = 189
    if true {
        v4 = append(v4, 0)
        v22 := GenS1{F0: 0, F1: 510, F2: v17}
        _ = v22
    } else {
        fmt.Println(v17)
        v18 = v18
    }
    var v23 uint16 = v21
    for fi24 := 0; fi24 < 6; fi24++ {
        v23 = v23
        _ = fi24
    }
    var v25 int8 = -86
    v17 = (v25 < -20)
    var v26 uint64 = (v1 ^ 2)
    var v27 bool = (42 > -851)
    v5 = v5
    var v28 int = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
