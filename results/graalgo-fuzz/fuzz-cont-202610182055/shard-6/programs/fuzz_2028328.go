package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 float32 = -2.25
    var v2 uint16 = 336
    var v3 int8 = -108
    var v4 int8 = 0
    v5 := []uint32{65535}
    v1 += -2.25
    if len(v5) > 0 {
        v5[0] = 65535
    }
    var v6 float32 = (0.0 - (v1 - float32(float64((float32((1e10 * (0.1 / 1.5))) / (((v1 / -2.25) - 0.1) - v1))))))
    switch -991 {
    case 9223372036854775807:
        v5 = append(v5, 2147483648)
    case 4294967295:
        switch -595 {
        case 590:
            v3 = (v4 >> ((0) & 63))
        case -1:
            fmt.Println(v6, v6)
        case -582:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        default:
            var v7 uint16 = 65535
            _ = v7
        }
    case 127:
        switch 9223372036854775807 {
        case 128:
            if len(v5) > 0 {
                v5[0] = 2147483648
            }
        }
    }
    switch 0 {
    case -9223372036854775808:
        v4 -= v4
    case -661:
        v2 += (v2 | v2)
    default:
        v4 -= (-(v4))
    }
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 uint32 = 0
        v10 := GenS1{F0: (v4 * (v4 >> ((2) & 63)))}
        var v11 int32 = 42
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    v6 -= v6
    var v12 uint8 = 1
    var v13 uint = 65535
    switch "" {
    case "hello":
        var v14 uint16 = v2
        _ = v14
    default:
        if (v2 < v2) {
            for fi15 := 0; fi15 < 19; fi15++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                fmt.Println(v1)
                _ = fi15
            }
        }
    }
    fmt.Println(v2)
    v16 := GenS1{F0: 82}
    if (49 >= 2) {
        var v17 int8 = v4
        v18 := GenS1{F0: (v3 % ((v17) | 1))}
        switch 128 {
        case 32768:
            v3 = v17
        default:
            switch 32768 {
            case 9223372036854775807:
                v5 = append(v5, 1)
            default:
                v6 -= (float32(-2.25) + v1)
            }
        }
        _ = v17
        _ = v18
    }
    _ = len(v5)
    for fi19 := 0; fi19 < 12; fi19++ {
        if true {
            var v20 int = fi19
            _ = v20
        } else {
            if len(v5) > 0 {
                v5[0] = 98
            }
        }
        _ = fi19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
