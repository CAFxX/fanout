package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int8 }

func main() {
    var v1 int16 = 1000
    var v2 int16 = -32768
    var v3 uint16 = 10
    var v4 int32 = -2147483648
    var v5 uint8 = 2
    v6 := GenS1{F0: uint(v3), F1: (0 < -322), F2: -128}
    v4 = (v4 & (v4 * v4))
    var v7 int8 = -1
    v8 := GenS1{F0: 462, F1: (65535 > 18446744073709551615), F2: v7}
    if true {
        for fi9 := 0; fi9 < 5; fi9++ {
            var v10 int = 32768
            var v11 uint32 = 809
            fmt.Println(v1)
            _ = fi9
            _ = v10
            _ = v11
        }
        v12 := GenS1{F0: 0, F1: (v7 >= -1), F2: v7}
        _ = v12
    }
    v13 := GenS1{F0: 575, F1: true, F2: (v7 << ((255) & 63))}
    if ((v1 * 236) < v1) {
        for fi14 := 0; fi14 < 18; fi14++ {
            v15 := GenS1{F0: 25, F1: false, F2: v7}
            for fi16 := 0; fi16 < 14; fi16++ {
                var v17 int8 = v7
                var v18 int = (fi16 + 9223372036854775807)
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = fi14
            _ = v15
        }
        v4 -= v4
        v5 = (v5 % (((^(v5))) | 1))
    } else {
        v1 = (v2 * (v1 ^ int16(v5)))
        for fi19 := 0; fi19 < 5; fi19++ {
            fmt.Println(v2, v2)
            v20 := []int{}
            _ = fi19
            _ = v20
        }
        var v21 uint8 = v5
        _ = v21
    }
    var v22 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
