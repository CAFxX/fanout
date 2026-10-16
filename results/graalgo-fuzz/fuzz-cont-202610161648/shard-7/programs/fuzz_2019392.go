package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }

func main() {
    var v1 int16 = 32767
    var v2 bool = (!((18446744073709551615 < 1) && (!(0 == 0))))
    v3 := make(map[uint32]uint32)
    var v4 int = -764
    var v5 int16 = (v1 - int16(v4))
    var v6 int32 = 0
    if (32767 >= 1) {
        v3[2147483647] = 833
        var v7 int = 1
        _ = v7
    } else {
        var v8 int16 = 1000
        _ = v8
    }
    var v9 uint = 4294967295
    for fi10 := 0; fi10 < 17; fi10++ {
        v11 := []int{(fi10 % ((v4) | 1)), (fi10 | 9223372036854775807)}
        switch (v4 % ((v4) | 1)) {
        case 254:
            _ = len(v3)
        default:
            fmt.Println(v2, v9, fi10)
        }
        var v12 int8 = -26
        _ = fi10
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        v5 = (v1 ^ v1)
        v9 = v9
        _ = fi13
    }
    var v14 float64 = 0.0
    v15 := GenS1{F0: (v4 << (((v9 - (v9 & (v9 << ((65535) & 63))))) & 63)), F1: 100}
    fmt.Println(v9)
    v9 += v9
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
