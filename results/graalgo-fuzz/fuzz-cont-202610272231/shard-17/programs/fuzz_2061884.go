package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 string }
func genFn0() int64 {
    for fi7 := 0; fi7 < 6; fi7++ {
        for fi8 := 0; fi8 < 12; fi8++ {
            var v9 int64 = 42
            var v10 uint = 18446744073709551615
            var v11 int64 = 9223372036854775807
            switch 32768 {
            case -957:
                v11 += (^(v11))
            case -48:
                var v12 int8 = 127
                _ = v12
            }
            switch fi8 {
            case 32768:
                v13 := []int{(^(fi8)), (fi8 % (((fi7 + (fi7 / ((186) | 1)))) | 1)), (fi8 ^ (-(fi8)))}
                _ = v13
            case 1:
                if (0.0 != float64(0.0)) {
                    v14 := []int{}
                    v9 = v9
                    _ = v14
                }
            }
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        var v15 int16 = 1
        _ = fi7
        _ = v15
    }
    var v16 int16 = 157
    v17 := []int{128, 127, -732}
    for fi18 := 0; fi18 < 18; fi18++ {
        _ = len(v17)
        var v19 int32 = 2147483647
        _ = fi18
        _ = v19
    }
    _ = v16
    _ = v17
    return 1
}
func genFn1() uint64 {
    var v20 int8 = 42
    for fi21 := 0; fi21 < 16; fi21++ {
        v20 -= v20
        var v22 int = (fi21 * -9223372036854775808)
        _ = fi21
        _ = v22
    }
    _ = v20
    return 2
}

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = 1
    v3 := make(map[int]int)
    var v4 int32 = int32(v2)
    var v5 int = 242
    var v6 uint8 = 42
    v5 = v5
    fmt.Println(v5)
    for fi23 := 0; fi23 < 16; fi23++ {
        for fi24 := 0; fi24 < 4; fi24++ {
            v5 = fi24
            fmt.Println(fi23)
            var v25 int64 = genFn0()
            _ = fi24
            _ = v25
        }
        v5 = int(v2)
        _ = fi23
    }
    var v26 int32 = v4
    if true {
        if false {
            v27 := []int{(v5 * -554), (-(v5)), (v5 / (((v5 << ((0) & 63))) | 1))}
            v28 := GenS1{F0: (v6 + (^(v6))), F1: 2, F2: "\n\t"}
            _ = v27
            _ = v28
        }
        var v29 int32 = (-(v26))
        _ = v29
    } else {
        v2 = (v2 & (v2 * (v2 - 1)))
        if false {
            delete(v3, v5)
        }
    }
    fmt.Println(v2)
    var v30 int32 = (v4 + v4)
    v30 = v30
    v3[(v5 >> ((44) & 63))] = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
