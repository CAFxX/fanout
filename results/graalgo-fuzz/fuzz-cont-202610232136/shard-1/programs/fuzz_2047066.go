package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 uint32; F2 int }
func genFn0(p0 uint64, p1 uint) bool {
    var v4 uint8 = 42
    v5 := []int{32767, 620}
    if (!false) {
        for fi6 := 0; fi6 < 12; fi6++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi6
        }
        v7 := GenS1{F0: 65535}
        var v8 int = 0
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 2147483647}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v9
    return true
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = 108
    var v3 uint8 = 181
    v3 = v3
    v10 := []int{32767}
    for fi11 := 0; fi11 < 5; fi11++ {
        v3 -= 127
        if len(v10) > 0 {
            v10[0] = fi11
        }
        v12 := GenS2{F0: (-(v2))}
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 float32 = float32(0.0)
        v14 = 0.0
        _ = fi13
        _ = v14
    }
    switch 127 {
    case -202:
        fmt.Println(v2)
    }
    v15 := GenS1{F0: 2147483648}
    if (false && (float64(3.14159) <= 0.1)) {
        if false {
            switch 4294967295 {
            case -9223372036854775808:
                v10 = append(v10, -983)
            case 32767:
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
            case 1:
                if true {
                    fmt.Println(v2)
                    switch -2147483648 {
                    case 127:
                        var v16 int16 = 0
                        _ = v16
                    }
                } else {
                    var v17 bool = (-246 != 757)
                    if len(v10) > 0 {
                        v10[0] = -2147483648
                    }
                    _ = v17
                }
            }
        } else {
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            v18 := GenS2{F0: v2}
            delete(v1, (("hello" + ("hello" + ("foo bar" + "a"))) + ""))
            _ = v18
        }
    } else {
        switch 774 {
        case 817:
            _ = len(v10)
        case -592:
            var v19 bool = true
            _ = v19
        default:
            v20 := GenS3{F0: 128, F1: 1, F2: -9223372036854775808}
            _ = v20
        }
        var v21 uint32 = 2147483648
        _ = v21
    }
    fmt.Println(v3, v2)
    var v22 int = 128
    var v23 float32 = (1e10 - 0.1)
    for fi24 := 0; fi24 < 2; fi24++ {
        var v25 uint32 = 2
        if ((-(fi24)) > (v22 - v22)) {
            v2 = -64
        }
        _ = fi24
        _ = v25
    }
    if ((0.1 / float64(v23)) >= 3.14159) {
        for fi26 := 0; fi26 < 12; fi26++ {
            var v27 uint32 = 124
            _ = fi26
            _ = v27
        }
        var v28 uint = 1
        _ = v28
    }
    v2 = -128
    v2 += 0
    v29 := GenS3{F0: v22, F1: 4294967295, F2: (v22 % ((2147483647) | 1))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v29
    _ = v3
}
