package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
func genFn0(p0 uint) uint8 {
    fmt.Println(p0)
    if false {
        p0 -= (p0 / (((p0 % (((p0 << (((p0 | 403)) & 63))) | 1))) | 1))
        if (true || (65535 != 0)) {
            v5 := []int{0, 127, 2147483647, -9223372036854775808}
            _ = v5
        }
    }
    p0 += (p0 & 42)
    p0 = (p0 & 2)
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 int64 = 42
        var v8 bool = (1 >= 1)
        var v9 int32 = -2147483648
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    _ = p0
    return 0
}
func genFn1(p0 int8) uint64 {
    v10 := []int{984, 9223372036854775807, -9223372036854775808, 428}
    v10 = append(v10, 32768)
    var v11 int32 = 2147483647
    v11 = (v11 | v11)
    _ = p0
    _ = v10
    _ = v11
    return 853
}

func main() {
    var v1 int8 = 42
    v2 := []int{1, -240, 127}
    v3 := make(map[int]int)
    var v4 int8 = v1
    v12 := GenS1{F0: true, F1: 0}
    if (1000 != -417) {
        var v13 int16 = 1000
        if (1 == -310) {
            for fi14 := 0; fi14 < 4; fi14++ {
                v4 = v4
                for fi15 := 0; fi15 < 11; fi15++ {
                    v13 = -755
                    v4 = v4
                    _ = fi15
                }
                var v16 bool = ((true && true) || true)
                _ = fi14
                _ = v16
            }
            v17 := GenS1{F0: ((0.1 > 3.14159) && (42 < 2147483648)), F1: 9223372036854775807}
            _ = v17
        }
        switch -790 {
        case 2147483647:
            v13 = -1
        }
        _ = v13
    }
    v4 = 127
    v4 -= -128
    for fi18 := 0; fi18 < 8; fi18++ {
        var v19 int = (fi18 * (fi18 - -9223372036854775808))
        var v20 bool = true
        _ = fi18
        _ = v19
        _ = v20
    }
    if true {
        for fi21 := 0; fi21 < 19; fi21++ {
            for fi22 := 0; fi22 < 14; fi22++ {
                v3[(fi22 >> ((255) & 63))] = fi22
                _ = fi22
            }
            v23 := GenS1{F0: false, F1: (fi21 ^ (fi21 * fi21))}
            v4 -= (v4 | v1)
            _ = fi21
            _ = v23
        }
        for fi24 := 0; fi24 < 2; fi24++ {
            v25 := GenS1{F0: (0 > 2), F1: fi24}
            _ = fi24
            _ = v25
        }
    }
    fmt.Println(v4, v1)
    if ((v4 << ((0) & 63)) <= 0) {
        v1 = 127
        for fi26 := 0; fi26 < 17; fi26++ {
            v4 = (v4 * (^(v1)))
            switch (fi26 & fi26) {
            case 0:
                if true {
                    v4 -= -81
                }
            case -55:
                delete(v3, (fi26 - (fi26 | int(v4))))
            case 42:
                v4 -= v1
            }
            _ = fi26
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
