package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 int16 = 32767
    var v2 int32 = 1
    var v3 float64 = (-2.25 - (3.14159 / 1.5))
    var v4 int8 = 0
    var v5 int16 = -126
    v5 = (v1 + (v5 / ((-1) | 1)))
    var v6 uint16 = 65535
    v7 := GenS1{F0: 127}
    v8 := GenS1{F0: 2}
    switch -1 {
    case 947:
        v9 := GenS1{F0: uint8(v6)}
        _ = v9
    case 9223372036854775807:
        for fi10 := 0; fi10 < 14; fi10++ {
            if ((0 <= 18446744073709551615) && true) {
                v11 := []int{}
                _ = v11
            }
            var v12 uint32 = 629
            var v13 uint = 65535
            _ = fi10
            _ = v12
            _ = v13
        }
    case 1:
        var v14 int64 = -966
        _ = v14
    default:
        switch "foo bar" {
        case "\n\t":
            var v15 float64 = float64(1e10)
            _ = v15
        }
    }
    v3 += v3
    v16 := GenS2{F0: 93}
    var v17 int32 = -1
    switch "a" {
    case "foo bar":
        if true {
            v6 -= (v6 << ((18446744073709551615) & 63))
            for fi18 := 0; fi18 < 13; fi18++ {
                v19 := GenS2{F0: 255}
                _ = fi18
                _ = v19
            }
            v3 += 3.14159
        } else {
            v3 = float64(0.0)
        }
    default:
        v5 += int16(v6)
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        var v21 bool = (1 < (v17 | (v2 * v17)))
        _ = fi20
        _ = v21
    }
    var v22 int8 = v4
    for fi23 := 0; fi23 < 20; fi23++ {
        v24 := GenS2{F0: 127}
        var v25 uint8 = uint8(v6)
        _ = fi23
        _ = v24
        _ = v25
    }
    v26 := GenS3{F0: (v4 % (((v22 - (v22 ^ (v4 + (v22 - int8(v2)))))) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
