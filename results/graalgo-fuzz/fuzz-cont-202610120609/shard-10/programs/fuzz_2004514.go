package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 uint8; F2 uint32 }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint, p1 string, p2 bool) int16 {
    p0 -= 255
    v4 := GenS1{F0: 42, F1: p1, F2: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 398
}

func main() {
    var v1 uint64 = 0
    var v2 uint16 = 32767
    var v3 uint16 = (v2 * (v2 + v2))
    v5 := GenS1{F0: 65535, F1: ("foo bar" + ""), F2: ("hello" + "hello")}
    var v6 int8 = -128
    if (42 > -390) {
        v1 = (^(v1))
    }
    var v7 bool = true
    v8 := []int{42, 42, -9223372036854775808}
    if (-994 != 32767) {
        _ = len(v8)
        v9 := GenS2{F0: uint8(v1)}
        v3 -= v3
        _ = v9
    } else {
        var v10 int32 = 1
        v11 := GenS3{F0: 42, F1: 127, F2: 42}
        _ = v10
        _ = v11
    }
    if (!false) {
        var v12 uint8 = 255
        _ = v12
    }
    if true {
        var v13 uint32 = 827
        switch "\n\t" {
        case "hello":
            v13 -= v13
        case "\n\t":
            v14 := GenS4{F0: v6}
            _ = v14
        case "foo bar":
            var v15 uint = uint(v3)
            _ = v15
        default:
            var v16 int16 = 32767
            _ = v16
        }
        v17 := GenS1{F0: v13, F1: "hello", F2: ("a" + ("\n\t" + ("hello" + "")))}
        _ = v13
        _ = v17
    }
    v2 += 2
    switch 1 {
    case 2147483647:
        fmt.Println(v3, v1, v2)
    }
    for fi18 := 0; fi18 < 3; fi18++ {
        var v19 uint64 = v1
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        switch fi18 {
        case -2147483648:
            if len(v8) > 0 {
                v8[0] = int(v2)
            }
        case 0:
            for fi20 := 0; fi20 < 3; fi20++ {
                _ = len(v8)
                var v21 bool = true
                _ = fi20
                _ = v21
            }
        }
        _ = fi18
        _ = v19
    }
    var v22 bool = (v7 || v7)
    fmt.Println(v1, v6, v22)
    var v23 bool = v22
    v7 = true
    for fi24 := 0; fi24 < 5; fi24++ {
        var v25 float32 = (3.14159 + -2.25)
        _ = fi24
        _ = v25
    }
    _ = len(v8)
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
