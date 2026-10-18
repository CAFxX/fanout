package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint64 }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 uint16; F1 int32; F2 int }

func main() {
    var v1 float32 = float32(3.14159)
    var v2 int8 = 0
    var v3 float32 = 0.1
    var v4 string = "foo bar"
    v5 := []int{9223372036854775807, 0, 1}
    var v6 int32 = -1
    switch -2147483648 {
    case 0:
        var v7 int32 = v6
        _ = v7
    case -547:
        var v8 float64 = 1e10
        _ = v8
    case 4294967295:
        v2 -= v2
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        switch v4 {
        case "":
            v10 := GenS1{F0: 127, F1: 2, F2: 0}
            _ = v10
        case "a":
            v2 = (v2 * (^(v2)))
        default:
            v3 += ((v3 * 1.5) * float32(0.1))
        }
        var v11 uint = 552
        _ = fi9
        _ = v11
    }
    switch "hello" {
    case "a":
        if false {
            var v12 uint = 459
            fmt.Println(v6, v12, v4)
            var v13 string = v4
            _ = v12
            _ = v13
        } else {
            v14 := GenS2{F0: v6, F1: 404}
            v4 = v4
            var v15 uint16 = 354
            _ = v14
            _ = v15
        }
    }
    v16 := GenS2{F0: v6, F1: 127}
    v17 := GenS3{F0: 18446744073709551615, F1: 2147483648}
    v18 := GenS3{F0: 255, F1: 376}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v6 = -1
    switch ("foo bar" + "hello") {
    case "":
        v5 = append(v5, 0)
    }
    v4 = v4
    if ("\n\t" <= ((v4 + v4) + (v4 + (v4 + "a")))) {
        v19 := GenS3{F0: 18446744073709551615, F1: 636}
        _ = v19
    }
    fmt.Println(v4)
    v20 := GenS4{F0: 1, F1: (^(v6)), F2: 4294967295}
    var v21 uint = 0
    v22 := GenS1{F0: 1, F1: v21, F2: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
