package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 int16, p1 int8, p2 int32) uint {
    var v6 int64 = 42
    if false {
        if false {
            for fi7 := 0; fi7 < 3; fi7++ {
                var v8 int8 = p1
                _ = fi7
                _ = v8
            }
            v9 := []int{}
            v6 = 9223372036854775807
            _ = v9
        }
    } else {
        var v10 float32 = float32((((-2.25 - 0.0) + 1e10) / 0.1))
        var v11 int16 = p0
        for fi12 := 0; fi12 < 12; fi12++ {
            for fi13 := 0; fi13 < 3; fi13++ {
                var v14 float32 = 0.1
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 807
}
func genFn1() uint8 {
    for fi15 := 0; fi15 < 15; fi15++ {
        switch 2147483647 {
        case -2147483648:
            for fi16 := 0; fi16 < 8; fi16++ {
                v17 := []int{4294967295, 128}
                _ = fi16
                _ = v17
            }
        case 32767:
            fmt.Println(fi15)
        }
        _ = fi15
    }
    v18 := []int{9223372036854775807, 4294967295, -1, 1}
    var v19 int = -340
    _ = v18
    _ = v19
    return 1
}

func main() {
    var v1 int = 4294967295
    var v2 int = (v1 >> ((86) & 63))
    var v3 int8 = -1
    var v4 int = v2
    v5 := []uint32{42, 1, 2147483647, 2147483647}
    switch -332 {
    case -995:
        var v20 int = (v4 << ((550) & 63))
        _ = v20
    case -360:
        var v21 int16 = -1
        _ = v21
    case 9223372036854775807:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    var v22 float32 = ((1e10 * 1e10) * -2.25)
    v22 += 0.0
    for fi23 := 0; fi23 < 19; fi23++ {
        switch ("\n\t" + ("\n\t" + "\n\t")) {
        case "hello":
            v24 := GenS1{F0: true}
            _ = v24
        case "\n\t":
            if (1 < 250) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v25 float32 = 1.5
                fmt.Println(v1)
                _ = v25
            } else {
                if len(v5) > 0 {
                    v5[0] = 2147483648
                }
                fmt.Println(v1, fi23, v3)
            }
        default:
            v5 = append(v5, 65535)
        }
        for fi26 := 0; fi26 < 9; fi26++ {
            v5 = append(v5, 2)
            var v27 uint8 = 2
            var v28 bool = ((0 >= 42) || false)
            _ = fi26
            _ = v27
            _ = v28
        }
        if true {
            var v29 int64 = -9223372036854775808
            switch fi23 {
            case 0:
                fmt.Println(v2)
            case 9223372036854775807:
                v4 = v4
            default:
                fmt.Println(v3)
            }
            _ = len(v5)
            _ = v29
        }
        _ = fi23
    }
    if (float64(1.5) != -2.25) {
        v1 += 0
        fmt.Println(v1, v1, v22)
    }
    for fi30 := 0; fi30 < 3; fi30++ {
        v22 = v22
        _ = fi30
    }
    v2 = (v1 % ((v2) | 1))
    v22 -= v22
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
