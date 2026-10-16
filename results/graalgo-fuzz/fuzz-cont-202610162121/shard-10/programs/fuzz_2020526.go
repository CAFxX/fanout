package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 string }

func main() {
    var v1 float64 = (((float64(1.5) * 0.0) - 0.0) / float64((1e10 * 0.1)))
    var v2 int8 = -60
    var v3 uint32 = 707
    var v4 int = -9223372036854775808
    v2 = (v2 | 42)
    var v5 uint32 = v3
    v4 = (v4 & 4294967295)
    v4 = v4
    var v6 int32 = 125
    var v7 uint8 = 0
    switch v4 {
    case 128:
        v8 := []int{}
        _ = v8
    default:
        for fi9 := 0; fi9 < 10; fi9++ {
            v10 := GenS1{F0: 1, F1: false, F2: ("hello" + "")}
            var v11 bool = true
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    var v12 int32 = (v6 << ((42) & 63))
    var v13 uint = 65535
    var v14 uint = (v13 >> (((v13 + v13)) & 63))
    if true {
        switch 1 {
        case 178:
            v6 = -2147483648
        case 32768:
            v15 := GenS1{F0: 42, F1: (false || (1 > 32767)), F2: ("\n\t" + "a")}
            _ = v15
        case 4294967295:
            v16 := []int{}
            _ = v16
        default:
            v14 -= v13
        }
        var v17 float64 = float64(float32(v1))
        _ = v17
    }
    if false {
        v18 := GenS1{F0: uint64(v7), F1: (false || (2 < 1)), F2: "a"}
        var v19 int8 = v2
        _ = v18
        _ = v19
    }
    var v20 int8 = v2
    switch v4 {
    case 1:
        v7 = v7
    case -2147483648:
        for fi21 := 0; fi21 < 1; fi21++ {
            if (float64(3.14159) >= (((v1 / 1.0) / (0.1 + 1e10)) * 0.0)) {
                v13 -= uint(v3)
            } else {
                v22 := []int{int(v2), (fi21 * (^(v4))), v4, v4}
                _ = v22
            }
            _ = fi21
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
