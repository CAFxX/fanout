package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int8 }
type GenS2 struct { F0 bool; F1 bool; F2 uint16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 2
    v3 := []uint32{931, 2147483647, 1, 2147483648}
    var v4 int = 32767
    var v5 int32 = 532
    var v6 bool = false
    if len(v3) > 0 {
        v3[0] = 42
    }
    fmt.Println(v2, v5)
    v5 -= -13
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if false {
        if true {
            fmt.Println(v6, v5, v4)
            v7 := GenS1{F0: 128, F1: (!((v6 && false) || false)), F2: 106}
            _ = v7
        } else {
            fmt.Println(v6, v6)
        }
        var v8 float32 = (1e10 / (3.14159 / 0.1))
        v9 := GenS2{F0: (v6 && v6), F1: (!((54 >= 974) && (32767 > 42))), F2: (v2 * (v2 & v2))}
        _ = v8
        _ = v9
    } else {
        for fi10 := 0; fi10 < 13; fi10++ {
            _ = len(v3)
            var v11 uint64 = 2
            _ = fi10
            _ = v11
        }
        v12 := GenS2{F0: (((1 >= -1) || (!(!(0.0 > (((3.14159 + 3.14159) / 1e10) - 1e10))))) && (v6 && v6)), F1: v6, F2: (v2 / ((255) | 1))}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v12
    }
    v13 := GenS2{F0: (42 > 1), F1: (!((v4 >> ((762) & 63)) >= 9223372036854775807)), F2: (v2 / ((1) | 1))}
    var v14 int8 = -29
    if (0 != 0) {
        if v6 {
            fmt.Println(v5, v14, v5)
        }
        fmt.Println(v2, v2, v6)
        v2 = 65535
    }
    var v15 int16 = 0
    fmt.Println(v6)
    var v16 int64 = 9223372036854775807
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
