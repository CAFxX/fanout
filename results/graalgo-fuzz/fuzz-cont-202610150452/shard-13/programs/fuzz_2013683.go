package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 int64; F2 int64 }
func genFn0(p0 uint64) uint16 {
    p0 -= (p0 >> ((2) & 63))
    p0 = (p0 & (p0 << ((65535) & 63)))
    _ = p0
    return 427
}
func genFn1(p0 int) uint32 {
    p0 += (-(p0))
    var v4 uint32 = 42
    var v5 uint16 = 32767
    var v6 uint16 = 32767
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return v4
}

func main() {
    var v1 int8 = 104
    v2 := []uint32{42, 506, 2147483647, 1}
    var v3 uint8 = 42
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 int16 = -32768
        _ = fi7
        _ = v8
    }
    if true {
        var v9 uint8 = 1
        _ = v9
    }
    fmt.Println(v3)
    v2 = append(v2, 2)
    var v10 float64 = 0.1
    var v11 uint64 = 296
    fmt.Println(v11, v10)
    if (!((819 < 42) && (-1 <= 2147483647))) {
        v1 += v1
        v12 := GenS1{F0: -2147483648}
        _ = v12
    } else {
        var v13 float64 = v10
        _ = v13
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 bool = false
        switch (fi14 - fi14) {
        case 9223372036854775807:
            if v15 {
                v2 = append(v2, 2)
            } else {
                fmt.Println(v15)
            }
        case 2147483647:
            v16 := GenS2{F0: 1, F1: 1, F2: 9223372036854775807}
            _ = v16
        case 4294967295:
            v17 := GenS1{F0: -1}
            _ = v17
        }
        switch "" {
        case "foo bar":
            v2 = append(v2, 0)
        }
        _ = fi14
        _ = v15
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        for fi19 := 0; fi19 < 6; fi19++ {
            v2 = append(v2, 1)
            if true {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                fmt.Println(v3, v10)
                if (1.5 <= v10) {
                    if (1 < v11) {
                        _ = len(v2)
                        if false {
                            _ = len(v2)
                            fmt.Println(fi19, v3)
                            v11 = v11
                        }
                    } else {
                        var v20 float32 = (-2.25 - -2.25)
                        fmt.Println(v3)
                        _ = v20
                    }
                    var v21 float32 = float32(((((((float64(0.1) / 0.1) * v10) - v10) * (v10 - (3.14159 / 0.1))) + (float64((float32(v10) - (3.14159 * 1.5))) / (1.5 - v10))) * v10))
                    var v22 uint8 = (v3 >> ((4294967295) & 63))
                    v21 = v21
                    _ = v21
                    _ = v22
                }
            } else {
                v1 += (v1 / ((1) | 1))
                v3 += v3
                switch "\n\t" {
                case "a":
                    v1 = (v1 - (v1 * v1))
                }
            }
            var v23 int = (fi18 % ((-943) | 1))
            _ = fi19
            _ = v23
        }
        _ = fi18
    }
    fmt.Println(v10, v11, v10)
    v10 += 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
}
