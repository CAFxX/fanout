package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 int64; F1 uint8; F2 uint32 }
func genFn0(p0 string, p1 uint16, p2 uint) uint64 {
    switch "a" {
    case "a":
        var v4 uint8 = 88
        _ = v4
    case "":
        var v5 uint = p2
        _ = v5
    case "foo bar":
        var v6 uint = (p2 & 58)
        _ = v6
    default:
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := GenS1{F0: (p1 & p1), F1: 469}
            _ = fi7
            _ = v8
        }
    }
    fmt.Println(p0)
    v9 := GenS2{F0: -9223372036854775808, F1: 127, F2: 895}
    switch -9223372036854775808 {
    case 1:
        v10 := GenS2{F0: 177, F1: 116, F2: 0}
        _ = v10
    default:
        v11 := GenS2{F0: 1, F1: 1, F2: 784}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 2
}

func main() {
    var v1 int32 = 0
    var v2 uint32 = 2147483648
    v3 := []int{32767, 128}
    var v12 uint64 = 18446744073709551615
    if len(v3) > 0 {
        v3[0] = 4294967295
    }
    var v13 int64 = 1
    switch 4294967295 {
    case 32768:
        for fi14 := 0; fi14 < 12; fi14++ {
            for fi15 := 0; fi15 < 6; fi15++ {
                v3 = append(v3, (fi15 << ((42) & 63)))
                _ = fi15
            }
            _ = fi14
        }
    case -107:
        var v16 int16 = 0
        _ = v16
    default:
        if ((v13 / ((v13) | 1)) >= -1) {
            var v17 uint16 = 171
            v18 := GenS1{F0: (v17 * (v17 ^ (v17 << ((255) & 63)))), F1: v13}
            switch 1 {
            case 1:
                if (!(!false)) {
                    v1 = v1
                }
            case 0:
                v3 = append(v3, 32768)
            }
            _ = v17
            _ = v18
        }
    }
    v12 = v12
    var v19 uint64 = 0
    var v20 float32 = float32((1e10 - (3.14159 - (0.0 * 0.1))))
    v21 := GenS2{F0: 0, F1: 255, F2: (^(v2))}
    switch 9223372036854775807 {
    case 42:
        for fi22 := 0; fi22 < 7; fi22++ {
            v19 = 18446744073709551615
            for fi23 := 0; fi23 < 2; fi23++ {
                var v24 bool = (false && ((v13 * (v13 & (v13 >> ((782) & 63)))) == (v13 & v13)))
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
    case 492:
        switch 0 {
        case 32767:
            var v25 uint64 = v12
            _ = v25
        }
    default:
        fmt.Println(v19, v12, v13)
    }
    if false {
        var v26 uint64 = (^(v19))
        for fi27 := 0; fi27 < 10; fi27++ {
            v26 = 18446744073709551615
            if len(v3) > 0 {
                v3[0] = -2147483648
            }
            _ = fi27
        }
        for fi28 := 0; fi28 < 18; fi28++ {
            var v29 bool = (3.14159 <= (0.1 * -2.25))
            for fi30 := 0; fi30 < 1; fi30++ {
                if (v29 || ((v1 | v1) == 139)) {
                    fmt.Println(v2, v12, v1)
                    v19 = (v12 & v12)
                    v31 := GenS2{F0: (v13 & (-(v13))), F1: 241, F2: v2}
                    v1 += (-(v1))
                    _ = v31
                } else {
                    if len(v3) > 0 {
                        v3[0] = (fi30 * 492)
                    }
                    v26 -= (v12 << ((4294967295) & 63))
                    switch (^(fi28)) {
                    case 2147483647:
                        var v32 float64 = 3.14159
                        _ = v32
                    case 1:
                        v13 += int64(v26)
                        v13 -= (-(v13))
                        v12 -= 42
                    case 9223372036854775807:
                        fmt.Println(v19, v12, v26)
                    default:
                        fmt.Println(v19)
                    }
                }
                v3 = append(v3, (fi30 + fi28))
                _ = fi30
            }
            _ = fi28
            _ = v29
        }
        _ = v26
    } else {
        v1 += (v1 * -1)
        var v33 int8 = 42
        v19 += v12
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
