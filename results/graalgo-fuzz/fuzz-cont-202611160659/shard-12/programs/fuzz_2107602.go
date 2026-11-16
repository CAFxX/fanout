package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string }
func genFn0(p0 uint32, p1 string) uint8 {
    v4 := GenS1{F0: 255}
    if (!false) {
        p0 -= (p0 & p0)
        fmt.Println(p0, p0)
    } else {
        if ((true || ((-2.25 < ((0.0 / -2.25) * 1.5)) || false)) || ((1.5 <= 3.14159) && (("a" + ("hello" + ("foo bar" + "hello"))) == "hello"))) {
            var v5 uint = 159
            _ = v5
        } else {
            p0 += p0
            v6 := GenS2{F0: 42, F1: 2147483647}
            _ = v6
        }
        p1 = p1
    }
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := GenS2{F0: 32767, F1: 2147483647}
        p0 -= 608
        v9 := []int{9223372036854775807, 127, (-(fi7))}
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v4
    return 128
}
func genFn1() int {
    var v10 int8 = -128
    fmt.Println(v10)
    v11 := GenS3{F0: 9223372036854775807}
    _ = v10
    _ = v11
    return 678
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = -221
    v3 := make(map[uint32]bool)
    v2 -= 92
    switch -1 {
    case -171:
        v2 -= (v2 / (((^(v2))) | 1))
    case 42:
        switch "a" {
        case "a":
            switch "" {
            case "":
                switch -163 {
                case -289:
                    v12 := GenS4{F0: ((("hello" + "") + "hello") + ("hello" + "foo bar"))}
                    _ = v12
                default:
                    var v13 float64 = 1.5
                    _ = v13
                }
            }
        }
    case 32767:
        v2 = v2
    }
    v2 -= v2
    v14 := GenS1{F0: 0}
    fmt.Println(v2)
    var v15 uint64 = 42
    if false {
        v15 += (v15 << ((42) & 63))
    }
    var v16 float64 = (1.5 / 1.5)
    v16 = v16
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
