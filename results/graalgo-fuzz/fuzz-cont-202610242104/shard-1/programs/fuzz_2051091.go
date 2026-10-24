package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 int16 }

func main() {
    var v1 int = 42
    var v2 int8 = 127
    var v3 float32 = 3.14159
    v3 = (v3 * (((0.0 + v3) - float32(1e10)) - float32(0.0)))
    var v4 uint32 = 65535
    v3 += float32(1.5)
    v5 := GenS1{F0: 710}
    var v6 int32 = -2147483648
    switch 9223372036854775807 {
    case 9223372036854775807:
        fmt.Println(v1, v3, v6)
    default:
        var v7 bool = false
        _ = v7
    }
    if (4294967295 <= 4294967295) {
        var v8 int32 = v6
        v6 = v6
        _ = v8
    }
    v9 := GenS1{F0: 942}
    var v10 int32 = v6
    v11 := GenS1{F0: 32767}
    v12 := GenS2{F0: (!true), F1: 42}
    var v13 int = 32768
    v14 := []int{(v1 >> ((18446744073709551615) & 63)), (v13 | v13)}
    fmt.Println(v6, v3)
    v15 := GenS2{F0: (false && ((32767 > 1000) || (!(42 > 65535)))), F1: -1}
    v16 := GenS2{F0: true, F1: 0}
    if ((v1 & v13) > v13) {
        _ = len(v14)
        v2 -= v2
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
