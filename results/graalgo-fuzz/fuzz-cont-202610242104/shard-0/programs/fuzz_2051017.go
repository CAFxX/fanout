package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 int; F2 int }
type GenS3 struct { F0 uint32; F1 int8; F2 int64 }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint8) string {
    for fi4 := 0; fi4 < 19; fi4++ {
        var v5 uint8 = (p0 << ((0) & 63))
        _ = fi4
        _ = v5
    }
    var v6 uint64 = 2
    switch "hello" {
    case "a":
        v7 := []int{}
        _ = v7
    case "hello":
        v6 -= v6
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = v6
    return "hello"
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 float32 = 1e10
    v3 := make(map[uint32]string)
    if false {
        switch (("a" + ("a" + "")) + "\n\t") {
        case "\n\t":
            v8 := GenS1{F0: -32768}
            _ = v8
        }
        v3[0] = ("a" + ("hello" + (("" + ("\n\t" + ("a" + "\n\t"))) + "foo bar")))
    }
    v9 := GenS1{F0: 0}
    for fi10 := 0; fi10 < 13; fi10++ {
        v11 := GenS2{F0: 479, F1: fi10, F2: (-(fi10))}
        _ = fi10
        _ = v11
    }
    if (1 < 4294967295) {
        v12 := GenS3{F0: 341, F1: 100, F2: (v1 / (((v1 & v1)) | 1))}
        var v13 uint8 = 171
        _ = v12
        _ = v13
    } else {
        switch "\n\t" {
        case "\n\t":
            var v14 int64 = (v1 + v1)
            _ = v14
        case "a":
            v1 += 0
        default:
            v15 := GenS4{F0: -58}
            _ = v15
        }
        var v16 float32 = v2
        v17 := []int{-1, 608, -785, 2147483647}
        _ = v16
        _ = v17
    }
    var v18 int = -648
    for fi19 := 0; fi19 < 14; fi19++ {
        fmt.Println(v2)
        _ = fi19
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        v1 -= -1
        v1 = v1
        _ = fi20
    }
    delete(v3, 2147483647)
    v2 = v2
    var v21 int32 = 0
    if false {
        v18 = 42
    } else {
        v2 = 3.14159
    }
    var v22 uint = 18446744073709551615
    v23 := v3[2147483648]
    switch 2147483647 {
    case -504:
        v23 = v23
    case -1:
        v24 := []int{-115}
        _ = v24
    case 127:
        var v25 float64 = -2.25
        _ = v25
    default:
        var v26 int16 = 42
        _ = v26
    }
    fmt.Println(v18, v1, v21)
    var v27 float64 = (((1e10 / 1e10) * (float64(v2) / ((0.1 + (1.5 + (1e10 + (0.1 * (float64((0.1 * 1.5)) + ((((0.0 + 3.14159) * 0.0) + -2.25) * -2.25)))))) - (1.5 + 1e10)))) / 1e10)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v9
}
