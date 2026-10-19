package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 int8) bool {
    var v4 string = "a"
    var v5 int = 1
    var v6 int16 = -341
    v7 := []int{v5, 0}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return (false && (!(!(!(81 >= 42)))))
}

func main() {
    var v1 int = 32767
    var v2 uint = 2
    v3 := []uint32{2147483648, 2, 1, 490}
    for fi8 := 0; fi8 < 20; fi8++ {
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 int = (v1 * fi8)
            switch v1 {
            case 800:
                v1 = (v1 + v10)
            default:
                if false {
                    v1 += 890
                }
            }
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    var v11 int8 = -46
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    _ = len(v3)
    v2 = (v2 >> ((v2) & 63))
    fmt.Println(v2, v2, v1)
    var v12 uint16 = 432
    v1 = -2147483648
    v2 += 304
    fmt.Println(v11)
    var v13 uint16 = v12
    var v14 int = (-(v1))
    v15 := GenS1{F0: 42}
    for fi16 := 0; fi16 < 11; fi16++ {
        var v17 uint = v2
        v2 -= (v2 % (((v2 * v2)) | 1))
        for fi18 := 0; fi18 < 12; fi18++ {
            var v19 float64 = (1.5 / (float64((0.1 / 1e10)) / 1.0))
            v2 = (v17 + (v17 * v17))
            var v20 float32 = 1.5
            var v21 uint8 = 42
            _ = len(v3)
            v22 := GenS1{F0: 18446744073709551615}
            var v23 int16 = -545
            if (((0 <= 42) || ((v13 + 2) < v13)) || (((v11 | v11) <= v11) || true)) {
                if len(v3) > 0 {
                    v3[0] = 65535
                }
            }
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
            _ = v22
            _ = v23
        }
        _ = fi16
        _ = v17
    }
    var v24 uint16 = 42
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (true && (v2 < v2)) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v3
}
