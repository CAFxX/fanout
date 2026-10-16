package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int32 }
type GenS2 struct { F0 int32; F1 int32 }
func genFn0() int {
    var v4 uint32 = 672
    v4 += 554
    var v5 int64 = 9223372036854775807
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 uint8, p1 int) bool {
    if true {
        switch -2147483648 {
        case 1:
            v6 := GenS1{F0: 4294967295, F1: (p1 >= genFn0()), F2: -189}
            _ = v6
        case 42:
            p0 += p0
        }
    }
    p1 = genFn0()
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 int64 = -1
            v10 := []int{genFn0(), -9223372036854775808, fi7}
            fmt.Println(p0)
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    p1 -= genFn0()
    _ = p0
    _ = p1
    return false
}

func main() {
    v1 := []string{}
    v2 := []uint32{2147483647, 31, 1, 2147483648}
    var v3 uint8 = 0
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 = 42
    var v11 int = -9223372036854775808
    if ((-626 != 42) && true) {
        var v12 uint64 = 2
        v3 -= 1
        _ = v12
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        var v14 int64 = 1
        _ = fi13
        _ = v14
    }
    v11 = genFn0()
    var v15 float64 = 0.0
    var v16 uint64 = 1
    for fi17 := 0; fi17 < 14; fi17++ {
        var v18 float32 = 3.14159
        fmt.Println(v11)
        _ = fi17
        _ = v18
    }
    v15 = float64(0.1)
    var v19 uint64 = (^(v16))
    var v20 bool = (true || (false && false))
    fmt.Println(v3, v3, v15)
    v15 -= (float64(1.5) - ((1e10 * float64(0.1)) - (v15 / v15)))
    if v20 {
        var v21 uint8 = (v3 * v3)
        _ = v21
    } else {
        for fi22 := 0; fi22 < 16; fi22++ {
            if (-32768 == 32767) {
                v20 = v20
                for fi23 := 0; fi23 < 4; fi23++ {
                    var v24 uint16 = 255
                    var v25 bool = false
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                v20 = v20
            }
            _ = fi22
        }
        v26 := GenS2{F0: 1, F1: 418}
        _ = v26
    }
    if (true || (v20 || false)) {
        fmt.Println(v15, v15, v11)
        var v27 uint8 = v3
        _ = v27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
