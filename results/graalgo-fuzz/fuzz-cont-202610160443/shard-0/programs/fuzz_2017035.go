package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 uint16; F2 int32 }
func genFn0(p0 int64, p1 uint64, p2 uint64) int {
    v7 := []int{-322, 0}
    p1 = 18446744073709551615
    var v8 bool = (int16(p0) > 33)
    switch -699 {
    case 4294967295:
        if (true && (-1 != 643)) {
            v7 = append(v7, -2147483648)
            var v9 int8 = int8(p2)
            _ = v9
        }
    case 1:
        if v8 {
            fmt.Println(p0, v8, v8)
            switch -1 {
            case 9223372036854775807:
                fmt.Println(p1, p1, v8)
            }
            v7 = append(v7, 32768)
        } else {
            v10 := GenS1{F0: true}
            var v11 int16 = -632
            var v12 uint = 65535
            _ = v10
            _ = v11
            _ = v12
        }
    }
    var v13 int8 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 uint8, p1 int16) uint8 {
    v14 := GenS2{F0: (("a" + "") + ("\n\t" + "")), F1: 65535, F2: 1}
    v15 := GenS2{F0: ("" + "a"), F1: 65, F2: 0}
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return p0
}

func main() {
    var v1 string = "foo bar"
    var v2 uint8 = 2
    var v3 int16 = -1
    var v4 int8 = 27
    var v5 int = 42
    var v6 uint32 = 28
    v16 := GenS1{F0: (2 > v6)}
    if (4294967295 != 4294967295) {
        v17 := []int{(^(v5)), genFn0(-9223372036854775808, 1, 4294967295), (v5 - v5)}
        _ = v17
    } else {
        for fi18 := 0; fi18 < 1; fi18++ {
            if (18446744073709551615 < 255) {
                var v19 int64 = 0
                fmt.Println(v3)
                _ = v19
            }
            var v20 float64 = 1e10
            _ = fi18
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        switch v5 {
        case -2147483648:
            fmt.Println(v6, v4, v2)
        }
        _ = fi21
    }
    v4 += (v4 % (((v4 / (((v4 & (v4 / (((v4 ^ (v4 / (((v4 << ((244) & 63))) | 1)))) | 1)))) | 1))) | 1))
    if (18446744073709551615 == 255) {
        v22 := []int{(-(v5))}
        _ = v22
    }
    v23 := GenS2{F0: (((v1 + (v1 + "\n\t")) + v1) + "foo bar"), F1: 297, F2: 42}
    v24 := GenS2{F0: "", F1: 65535, F2: 1}
    var v25 float64 = 3.14159
    for fi26 := 0; fi26 < 18; fi26++ {
        v3 = -1
        fmt.Println(v25, fi26)
        _ = fi26
    }
    v3 = v3
    for fi27 := 0; fi27 < 17; fi27++ {
        var v28 string = "a"
        fmt.Println(v3, v25, v6)
        _ = fi27
        _ = v28
    }
    v29 := []int{127, v5}
    fmt.Println(v25, v4, v4)
    v1 = v1
    fmt.Println(v5, v5, v4)
    if false {
        var v30 uint16 = 0
        _ = v30
    } else {
        if len(v29) >= 2 {
            v29 = v29[1:len(v29)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
