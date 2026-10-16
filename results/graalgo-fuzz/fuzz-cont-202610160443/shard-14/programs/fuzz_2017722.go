package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint32, p1 uint8, p2 int8) uint16 {
    var v5 int64 = 9223372036854775807
    var v6 float64 = (float64(1e10) + 0.1)
    for fi7 := 0; fi7 < 13; fi7++ {
        v5 = v5
        var v8 int64 = (v5 / ((v5) | 1))
        _ = fi7
        _ = v8
    }
    if ((!true) && (!false)) {
        for fi9 := 0; fi9 < 12; fi9++ {
            v10 := []int{-249}
            if len(v10) > 0 {
                v10[0] = (fi9 >> ((463) & 63))
            }
            _ = fi9
            _ = v10
        }
    } else {
        fmt.Println(p0)
    }
    p0 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 255
}
func genFn1(p0 uint16, p1 uint16, p2 uint) int8 {
    var v11 uint16 = (p0 ^ p0)
    var v12 uint8 = 2
    switch -482 {
    case 42:
        var v13 int16 = 0
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 42
}

func main() {
    var v1 float32 = 1e10
    var v2 uint32 = 96
    var v3 uint64 = 4294967295
    var v4 uint8 = 62
    var v14 bool = (0 > 2147483647)
    v15 := []int{32767, 212, 2147483647}
    var v16 bool = v14
    v17 := GenS1{F0: v3}
    v15 = append(v15, 128)
    if (!(-114 >= 100)) {
        v3 = (^(v3))
        v2 = v2
        if v16 {
            var v18 uint = 255
            switch 0 {
            case 445:
                var v19 float32 = v1
                _ = v19
            case -41:
                if (833 <= 0) {
                    fmt.Println(v1, v14)
                    for fi20 := 0; fi20 < 8; fi20++ {
                        if len(v15) > 0 {
                            v15[0] = 42
                        }
                        _ = fi20
                    }
                }
            case 32767:
                fmt.Println(v4, v14, v3)
            }
            v21 := GenS2{F0: (v2 ^ v2)}
            _ = v18
            _ = v21
        }
    }
    var v22 uint32 = uint32(v3)
    v1 += v1
    v16 = v14
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
