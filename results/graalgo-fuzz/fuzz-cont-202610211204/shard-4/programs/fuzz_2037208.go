package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 uint16; F2 uint }
type GenS4 struct { F0 uint32 }

func main() {
    var v1 uint = 0
    var v2 float64 = (0.0 + 1.5)
    var v3 int = 9223372036854775807
    var v4 int = 127
    var v5 uint8 = 233
    var v6 int32 = -1
    for fi7 := 0; fi7 < 3; fi7++ {
        fmt.Println(v5, fi7, fi7)
        fmt.Println(v2, fi7, fi7)
        v3 += (v4 + v4)
        _ = fi7
    }
    var v8 int32 = (v6 / ((v6) | 1))
    v1 = v1
    v9 := GenS1{F0: 1}
    v4 -= (^(v3))
    switch v3 {
    case 32767:
        fmt.Println(v5)
    case 132:
        v3 = (v3 / ((v4) | 1))
    case 901:
        switch "" {
        case "foo bar":
            for fi10 := 0; fi10 < 20; fi10++ {
                switch v4 {
                case 9223372036854775807:
                    var v11 int = (fi10 % ((-372) | 1))
                    _ = v11
                default:
                    var v12 uint8 = (v5 << ((v1) & 63))
                    _ = v12
                }
                _ = fi10
            }
        }
    default:
        switch int(v5) {
        case 0:
            fmt.Println(v1, v8, v1)
        case 195:
            var v13 int8 = 100
            _ = v13
        case 128:
            for fi14 := 0; fi14 < 19; fi14++ {
                v15 := []int{}
                _ = fi14
                _ = v15
            }
        default:
            var v16 uint16 = 493
            _ = v16
        }
    }
    var v17 uint32 = 4294967295
    switch 32767 {
    case -116:
        var v18 uint16 = 32767
        _ = v18
    case -50:
        v19 := []int{}
        _ = v19
    }
    v3 = v3
    v20 := GenS1{F0: 0}
    switch "hello" {
    case "\n\t":
        v2 = v2
    case "foo bar":
        if true {
            if false {
                var v21 int16 = -30
                v22 := GenS1{F0: 42}
                _ = v21
                _ = v22
            } else {
                if (42 >= 0) {
                    v23 := []int{v4, int(v5), (^(v4))}
                    _ = v23
                }
                v24 := GenS2{F0: v1}
                _ = v24
            }
        }
    default:
        v25 := GenS3{F0: v8, F1: 65535, F2: v1}
        _ = v25
    }
    switch v3 {
    case -2147483648:
        v5 += (^(v5))
    case 2147483647:
        v1 -= v1
    case 4294967295:
        v26 := []int{80, int(v1)}
        _ = v26
    default:
        v27 := GenS3{F0: 54, F1: 691, F2: 904}
        _ = v27
    }
    v28 := GenS4{F0: 932}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
