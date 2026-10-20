package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint16 }

func main() {
    var v1 bool = (!(true && (("\n\t" + ("" + ("" + "hello"))) >= ("" + ("\n\t" + "")))))
    var v2 string = "\n\t"
    var v3 bool = (202 <= 2147483647)
    var v4 uint16 = 32767
    switch 419 {
    case 128:
        var v5 int = 220
        _ = v5
    }
    v1 = false
    var v6 uint8 = 127
    if (128 != 0) {
        v3 = v3
    }
    v6 -= (^(v6))
    v7 := GenS1{F0: 4294967295, F1: 4294967295, F2: v4}
    v8 := []int{969, -2147483648, 128, 32768}
    var v9 uint32 = 4294967295
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    v1 = v3
    v10 := GenS1{F0: 42, F1: 197, F2: 170}
    switch 127 {
    case 32767:
        var v11 uint = 247
        _ = v11
    case 4294967295:
        var v12 int32 = 0
        _ = v12
    default:
        v1 = ((0 > 42) || (!(v1 || (((((3.14159 + (0.0 - 3.14159)) - ((-2.25 + (0.0 - -2.25)) * 1e10)) * 3.14159) * (1.5 * (0.1 * ((1e10 * ((0.0 + ((1e10 - 0.1) / 1.0)) / ((3.14159 / 3.14159) / 1e10))) * 1e10)))) <= 0.0))))
    }
    v2 = v2
    var v13 int = -860
    v1 = ((v13 / ((v13) | 1)) == -327)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
