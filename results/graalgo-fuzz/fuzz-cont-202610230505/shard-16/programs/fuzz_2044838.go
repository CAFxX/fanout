package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 uint64; F2 uint }
func genFn0(p0 int32, p1 uint8) uint {
    switch 2147483647 {
    case -379:
        var v6 float64 = (((-2.25 - (float64(1e10) * 3.14159)) - 0.0) - (-2.25 / 0.1))
        _ = v6
    }
    p1 += (p1 + (p1 * p1))
    _ = p0
    _ = p1
    return 42
}
func genFn1(p0 uint16, p1 uint) uint32 {
    if (((3.14159 - -2.25) + 0.1) >= 0.1) {
        p0 -= (p0 << (((^(p1))) & 63))
    } else {
        v7 := GenS1{F0: 1}
        _ = v7
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        if true {
            var v9 int32 = 42
            fmt.Println(v9)
            _ = v9
        }
        p1 -= p1
        for fi10 := 0; fi10 < 2; fi10++ {
            if ((true && false) && false) {
                v11 := []int{fi8}
                var v12 bool = true
                _ = v11
                _ = v12
            } else {
                v13 := GenS1{F0: 378}
                p0 = (^(p0))
                _ = v13
            }
            var v14 uint = 42
            v15 := []int{(fi8 ^ (fi10 << ((191) & 63)))}
            _ = fi10
            _ = v14
            _ = v15
        }
        _ = fi8
    }
    v16 := []int{162, 32768, 127, 32768}
    if (!(-720 == 249)) {
        v17 := GenS2{F0: p1, F1: 2, F2: (p1 & 65535)}
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v16
    return 2147483648
}

func main() {
    var v1 uint16 = 65535
    v2 := []int{-1, -341}
    var v3 int16 = 1000
    v4 := make(map[string]uint32)
    var v5 uint16 = 1
    v18 := GenS2{F0: 4294967295, F1: 18446744073709551615, F2: 4294967295}
    v5 = 0
    fmt.Println(v1, v5)
    var v19 int16 = v3
    if ((v1 % ((v5) | 1)) >= v1) {
        for fi20 := 0; fi20 < 10; fi20++ {
            v21 := GenS2{F0: genFn0(2147483647, 202), F1: 18446744073709551615, F2: 922}
            var v22 string = "foo bar"
            v3 += v3
            _ = fi20
            _ = v21
            _ = v22
        }
    } else {
        for fi23 := 0; fi23 < 5; fi23++ {
            for fi24 := 0; fi24 < 17; fi24++ {
                _ = len(v2)
                v1 = v1
                switch 128 {
                case 32768:
                    v5 -= (v5 >> ((205) & 63))
                }
                _ = fi24
            }
            var v25 float64 = float64(0.1)
            fmt.Println(v3, v19, fi23)
            _ = fi23
            _ = v25
        }
        fmt.Println(v1, v3, v3)
        v1 = 985
    }
    v5 += v1
    fmt.Println(v1, v3, v3)
    var v26 int = 127
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
