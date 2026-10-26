package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
func genFn0(p0 int32) int64 {
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            v9 := []int{fi7}
            fmt.Println(fi7, fi8, p0)
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    v10 := []int{128, -9223372036854775808}
    var v11 float64 = -2.25
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    var v12 int32 = (p0 & (p0 << ((65535) & 63)))
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    return -645
}

func main() {
    var v1 int8 = 1
    var v2 int16 = 1000
    v3 := make(map[int]bool)
    v4 := []int{42, 2147483647}
    var v5 uint8 = 254
    var v6 int = -1
    switch (v6 / ((v6) | 1)) {
    case 755:
        v2 = 0
    }
    if len(v4) > 0 {
        v4[0] = (v6 - 0)
    }
    var v13 float64 = 3.14159
    for fi14 := 0; fi14 < 9; fi14++ {
        delete(v3, (fi14 >> ((2) & 63)))
        v2 = v2
        _ = fi14
    }
    v2 = v2
    var v15 int = v6
    var v16 uint32 = 985
    var v17 uint = 65535
    if (v15 == (v15 << (((^(v17))) & 63))) {
        v16 += v16
        for fi18 := 0; fi18 < 19; fi18++ {
            for fi19 := 0; fi19 < 15; fi19++ {
                fmt.Println(v15, fi18, v16)
                if (448 != 1) {
                    var v20 uint = v17
                    var v21 uint64 = 2
                    _ = v20
                    _ = v21
                } else {
                    v6 += -1
                }
                _ = fi19
            }
            _ = fi18
        }
        v22 := GenS1{F0: (!(v6 < (-(v15)))), F1: 0}
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
