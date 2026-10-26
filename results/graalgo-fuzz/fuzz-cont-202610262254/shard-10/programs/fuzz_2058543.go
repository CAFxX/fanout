package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint; F1 uint32; F2 uint8 }
func genFn0() uint64 {
    switch -338 {
    case 0:
        if (255 == 1) {
            v7 := []int{4294967295, 42, -257}
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = v7
        }
    }
    if (false || (42 >= -9223372036854775808)) {
        for fi8 := 0; fi8 < 12; fi8++ {
            v9 := GenS1{F0: 280, F1: -128}
            _ = fi8
            _ = v9
        }
        var v10 int8 = 0
        var v11 float64 = float64((3.14159 / (1.5 / 3.14159)))
        _ = v10
        _ = v11
    } else {
        for fi12 := 0; fi12 < 16; fi12++ {
            var v13 int64 = 9223372036854775807
            _ = fi12
            _ = v13
        }
    }
    var v14 float64 = (3.14159 - 1.5)
    _ = v14
    return 4294967295
}

func main() {
    var v1 uint16 = 0
    v2 := make(map[string]bool)
    var v3 int = -9223372036854775808
    v4 := make(map[uint32]string)
    var v5 bool = true
    var v6 int64 = 492
    for fi15 := 0; fi15 < 13; fi15++ {
        v3 = fi15
        _ = fi15
    }
    v1 += 2
    v16 := GenS2{F0: 510, F1: 2147483648, F2: 128}
    var v17 uint8 = 2
    v5 = v5
    v6 = 487
    for fi18 := 0; fi18 < 3; fi18++ {
        switch -1 {
        case -9223372036854775808:
            var v19 int16 = -1
            _ = v19
        default:
            v20 := []int{fi18, (fi18 ^ v3)}
            _ = v20
        }
        _ = fi18
    }
    v17 = v17
    var v21 string = ("\n\t" + "hello")
    fmt.Println(v21, v3, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
