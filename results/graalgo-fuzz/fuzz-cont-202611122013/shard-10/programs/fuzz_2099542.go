package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 int64 = 514
    var v2 int32 = -2147483648
    v3 := []int{128, -250, -382, 434}
    v4 := []uint32{2, 260, 65535}
    var v5 float64 = 1e10
    _ = len(v4)
    if (v5 >= (0.0 - float64(0.0))) {
        v1 += v1
        if false {
            v5 = (v5 * (-2.25 + v5))
            var v6 uint32 = 2
            _ = v6
        }
    } else {
        var v7 int16 = 0
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v7
    }
    fmt.Println(v2, v1, v1)
    fmt.Println(v1)
    fmt.Println(v1, v2, v1)
    var v8 int = 128
    for fi9 := 0; fi9 < 13; fi9++ {
        v5 = ((float64((1e10 * 0.1)) - v5) / v5)
        var v10 int32 = v2
        for fi11 := 0; fi11 < 18; fi11++ {
            fmt.Println(v5, v2)
            fmt.Println(v1, fi9)
            switch 32768 {
            case -572:
                var v12 uint64 = 128
                _ = v12
            case 2147483647:
                if false {
                    fmt.Println(v2, v2, v8)
                    var v13 int32 = v10
                    _ = v13
                } else {
                    var v14 uint8 = 255
                    var v15 uint16 = 2
                    var v16 uint32 = 2147483648
                    var v17 uint8 = (v14 & (v14 + v14))
                    _ = v14
                    _ = v15
                    _ = v16
                    _ = v17
                }
            case 128:
                v1 -= (v1 + (v1 % ((v1) | 1)))
            default:
                v2 -= (v10 >> ((65535) & 63))
            }
            _ = fi11
        }
        _ = fi9
        _ = v10
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi18 := 0; fi18 < 3; fi18++ {
        fmt.Println(v2, v8)
        v1 = 87
        if (!(-37 == (v8 >> ((127) & 63)))) {
            v19 := GenS1{F0: -95}
            if false {
                for fi20 := 0; fi20 < 6; fi20++ {
                    v21 := GenS1{F0: 127}
                    _ = fi20
                    _ = v21
                }
                v4 = append(v4, 417)
                for fi22 := 0; fi22 < 10; fi22++ {
                    var v23 int32 = (v2 << ((2) & 63))
                    _ = fi22
                    _ = v23
                }
            } else {
                var v24 uint32 = 4294967295
                var v25 string = "a"
                _ = v24
                _ = v25
            }
            var v26 uint64 = 2
            _ = v19
            _ = v26
        }
        _ = fi18
    }
    switch (v8 + 1) {
    case -9223372036854775808:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
