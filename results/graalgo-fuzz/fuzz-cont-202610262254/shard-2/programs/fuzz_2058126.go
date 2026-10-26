package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint, p1 bool) uint64 {
    p0 -= p0
    p0 = p0
    _ = p0
    _ = p1
    return 793
}

func main() {
    var v1 int = -263
    var v2 uint32 = 1
    var v3 float32 = 0.0
    v4 := []int{}
    v5 := GenS1{F0: 18446744073709551615, F1: 227, F2: 32767}
    for fi6 := 0; fi6 < 13; fi6++ {
        v2 = 4294967295
        _ = fi6
    }
    fmt.Println(v1, v2, v3)
    switch 600 {
    case 739:
        v4 = append(v4, 2147483647)
    }
    for fi7 := 0; fi7 < 14; fi7++ {
        fmt.Println(v3)
        for fi8 := 0; fi8 < 8; fi8++ {
            fmt.Println(v2, fi7)
            v9 := GenS1{F0: 4294967295, F1: -32768, F2: 0}
            v10 := GenS1{F0: 65535, F1: 32767, F2: 32767}
            if false {
                if ((180 != 65535) || ((-75 < 100) || (!((-98 < 42) || false)))) {
                    var v11 int32 = -2147483648
                    fmt.Println(v3, fi7, v1)
                    v2 -= v2
                    _ = v11
                }
                fmt.Println(fi8)
                v1 = -9223372036854775808
            } else {
                v12 := GenS1{F0: 236, F1: 1, F2: 42}
                _ = v12
            }
            if (42 >= 4294967295) {
                _ = len(v4)
            } else {
                var v13 int8 = -13
                _ = v13
            }
            v1 = -2147483648
            var v14 uint32 = (v2 << ((255) & 63))
            _ = fi8
            _ = v10
            _ = v14
            _ = v9
        }
        _ = fi7
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v1, v3, v2)
    fmt.Println(v2, v3)
    switch ("" + "\n\t") {
    case "hello":
        v3 = (float32(1e10) + v3)
    case "\n\t":
        v1 += (-(v1))
    default:
        v3 = (0.1 + v3)
    }
    if ((1 > 42) && (true || false)) {
        switch (v1 << ((421) & 63)) {
        case 128:
            if (-1 >= 1) {
                v15 := GenS2{F0: 128}
                for fi16 := 0; fi16 < 6; fi16++ {
                    v2 += 4294967295
                    v2 -= (v2 + 2147483647)
                    _ = fi16
                }
                _ = v15
            }
        case 1:
            for fi17 := 0; fi17 < 14; fi17++ {
                v3 = 0.1
                v3 += ((0.0 * (v3 / v3)) + v3)
                _ = fi17
            }
        case -156:
            fmt.Println(v1, v2)
        }
    }
    if (0 != 128) {
        if (2147483648 < v2) {
            var v18 int8 = 62
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        if len(v4) > 0 {
            v4[0] = (v1 << ((0) & 63))
        }
        _ = fi19
    }
    v20 := GenS2{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
