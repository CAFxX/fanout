package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 uint64; F1 uint32; F2 string }
type GenS6 struct { F0 uint32 }
func genFn0(p0 int, p1 bool, p2 uint8) bool {
    if ((p2 > (^(p2))) || true) {
        switch "" {
        case "a":
            p1 = ((false || (42 == 345)) && p1)
        case "hello":
            var v7 uint64 = 4294967295
            _ = v7
        }
        v8 := GenS1{F0: -779}
        p1 = (!((p2 / (((p2 ^ (^(p2)))) | 1)) <= (p2 * p2)))
        _ = v8
    }
    fmt.Println(p1)
    var v9 float64 = 0.1
    var v10 uint32 = 4294967295
    p2 = (p2 & p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return p1
}
func genFn1(p0 uint64, p1 uint8, p2 uint16) int8 {
    switch -2147483648 {
    case 375:
        var v11 int16 = 0
        _ = v11
    case 0:
        for fi12 := 0; fi12 < 4; fi12++ {
            var v13 int32 = 42
            fmt.Println(v13)
            v14 := GenS2{F0: 65535}
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    p1 -= 127
    p2 += p2
    if ((-2.25 <= 0.1) || ((((3.14159 - 0.1) - (-2.25 - (1.5 - (3.14159 * 1e10)))) - 0.1) != (1e10 + ((0.1 * 1e10) + 3.14159)))) {
        p0 += p0
        p2 -= p2
    }
    _ = p0
    _ = p1
    _ = p2
    return 127
}

func main() {
    var v1 uint16 = 80
    var v2 string = ("foo bar" + "a")
    var v3 float32 = 0.1
    v4 := make(map[uint32]bool)
    var v5 bool = true
    var v6 string = v2
    var v15 float64 = 0.0
    v16 := []int{32767, 128, 4294967295}
    var v17 uint32 = 0
    v18 := GenS3{F0: 42, F1: (!(973 < 255))}
    switch 808 {
    case 2147483647:
        v6 = "\n\t"
    default:
        switch "\n\t" {
        case "hello":
            v17 = (^(v17))
        case "\n\t":
            var v19 uint64 = 573
            _ = v19
        case "":
            v20 := GenS4{F0: v5, F1: -42}
            _ = v20
        }
    }
    switch 32767 {
    case -1:
        v21 := GenS5{F0: 1, F1: v17, F2: ((("" + (("foo bar" + (v6 + (v6 + (v6 + v2)))) + (("a" + v6) + v2))) + "hello") + "hello")}
        _ = v21
    case 32767:
        v22 := GenS6{F0: v17}
        _ = v22
    case 128:
        _ = len(v16)
    }
    switch v6 {
    case "hello":
        if v5 {
            switch 9223372036854775807 {
            case 0:
                var v23 uint64 = 4294967295
                _ = v23
            case 566:
                fmt.Println(v5, v3, v15)
            }
        }
    }
    v24 := GenS4{F0: true, F1: 2147483647}
    fmt.Println(v2)
    var v25 int8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
