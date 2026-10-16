package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int8 }
func genFn0() uint32 {
    v7 := GenS1{F0: 798, F1: 0, F2: -116}
    var v8 int64 = -600
    var v9 int16 = 772
    v10 := []int{4294967295}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 2147483647
}

func main() {
    var v1 uint8 = 1
    var v2 int = -2147483648
    var v3 float64 = -2.25
    var v4 int16 = 0
    var v5 uint64 = 535
    var v6 int64 = -1
    switch v2 {
    case -2147483648:
        for fi11 := 0; fi11 < 18; fi11++ {
            fmt.Println(v4)
            v12 := GenS1{F0: (v4 & int16(v1)), F1: 42, F2: 42}
            _ = fi11
            _ = v12
        }
    }
    fmt.Println(v5, v3, v3)
    for fi13 := 0; fi13 < 18; fi13++ {
        v3 += v3
        if true {
            v14 := []int{fi13, (-(fi13)), (fi13 ^ (v2 + (-(v2)))), (-(v2))}
            for fi15 := 0; fi15 < 14; fi15++ {
                v1 = (v1 / ((v1) | 1))
                var v16 string = "a"
                if false {
                    fmt.Println(fi15, fi13, v16)
                    if ((false || (((int8(fi13) != -1) && false) && false)) || (true || true)) {
                        if len(v14) > 0 {
                            v14[0] = fi15
                        }
                    }
                    v2 = fi15
                }
                _ = fi15
                _ = v16
            }
            _ = v14
        }
        _ = fi13
    }
    var v17 uint16 = 1
    var v18 uint16 = uint16(v5)
    v2 = 127
    v19 := []int{(^(v2)), v2}
    if false {
        v5 = v5
        switch v2 {
        case 964:
            for fi20 := 0; fi20 < 4; fi20++ {
                fmt.Println(fi20)
                _ = fi20
            }
        case 32767:
            v3 = (((0.1 - v3) + -2.25) / 1.0)
        case -716:
            v5 -= uint64(v1)
        }
    }
    var v21 uint8 = v1
    var v22 uint8 = v1
    var v23 bool = false
    if ((v6 ^ (v6 + int64(v4))) < v6) {
        switch (v2 | (v2 % ((v2) | 1))) {
        case -2147483648:
            if v23 {
                v6 -= (v6 / ((v6) | 1))
                var v24 int64 = v6
                fmt.Println(v24)
                _ = v24
            }
        default:
            switch (v2 ^ -9223372036854775808) {
            case -2147483648:
                v17 = (^(v18))
            default:
                v4 = (v4 / (((v4 & (v4 * v4))) | 1))
            }
        }
        for fi25 := 0; fi25 < 3; fi25++ {
            fmt.Println(v3, v22, v6)
            if (((((v3 + v3) * 0.0) * v3) + float64(float32(v3))) > 1.5) {
                var v26 int64 = (^(v6))
                _ = v26
            }
            v23 = (1 <= v2)
            _ = fi25
        }
    }
    var v27 int16 = -295
    v28 := GenS1{F0: (v27 >> ((2) & 63)), F1: v6, F2: -1}
    if v23 {
        v6 = 0
        if v23 {
            var v29 float64 = (float64(((1.5 - 1e10) * 3.14159)) * (v3 - float64(0.1)))
            _ = v29
        } else {
            v30 := GenS1{F0: -960, F1: v6, F2: 0}
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
            _ = v30
        }
    } else {
        _ = len(v19)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
