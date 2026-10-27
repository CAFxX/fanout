package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint16, p1 uint8, p2 uint16) bool {
    if true {
        switch 610 {
        case 32767:
            var v4 int64 = -9223372036854775808
            _ = v4
        default:
            var v5 float32 = ((1e10 + 1.5) * (((((((-2.25 * 1e10) * (3.14159 / 1e10)) - ((1e10 - 0.1) * float32(((-2.25 / -2.25) - 0.0)))) - (1e10 - (-2.25 - 3.14159))) - ((((0.0 - (1e10 - float32((1e10 / -2.25)))) - float32(-2.25)) + 0.0) + 0.1)) + 1e10) + (3.14159 * 3.14159)))
            _ = v5
        }
        v6 := GenS1{F0: 1, F1: 32767}
        var v7 int32 = 0
        _ = v6
        _ = v7
    } else {
        var v8 int16 = 1
        var v9 uint64 = 185
        p2 = (p2 & p0)
        _ = v8
        _ = v9
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    return false
}

func main() {
    var v1 int64 = 587
    var v2 int64 = (v1 << ((18446744073709551615) & 63))
    var v3 float64 = (((1.5 + 1e10) * 0.1) * 1.5)
    v10 := GenS2{F0: 127, F1: 2}
    switch 0 {
    case 1:
        var v11 uint64 = 997
        _ = v11
    case -2147483648:
        v12 := []int{0, -612}
        _ = v12
    case 127:
        v13 := GenS2{F0: 32767, F1: 127}
        _ = v13
    default:
        var v14 int8 = -41
        _ = v14
    }
    var v15 float64 = v3
    v16 := GenS1{F0: -255, F1: 825}
    for fi17 := 0; fi17 < 20; fi17++ {
        v1 += 538
        var v18 float32 = (3.14159 * -2.25)
        var v19 bool = (323 == 155)
        _ = fi17
        _ = v18
        _ = v19
    }
    v1 = (v1 % (((v2 & v2)) | 1))
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v1)
    }
    if true {
        v20 := GenS3{F0: 1}
        _ = v20
    }
    v21 := GenS1{F0: 2147483647, F1: 42}
    switch ("a" + ("a" + ("a" + (("\n\t" + ("hello" + "\n\t")) + ((("\n\t" + "a") + "\n\t") + "a"))))) {
    case "\n\t":
        if (!true) {
            for fi22 := 0; fi22 < 2; fi22++ {
                var v23 int64 = (v1 >> ((18446744073709551615) & 63))
                var v24 float32 = float32(float64((1e10 - 1.5)))
                _ = fi22
                _ = v23
                _ = v24
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
}
