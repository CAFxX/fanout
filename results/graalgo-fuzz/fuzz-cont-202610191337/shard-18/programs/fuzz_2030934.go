package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 int16 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint }
func genFn0(p0 int, p1 int, p2 int) uint8 {
    var v7 uint16 = 2
    if true {
        v7 -= (v7 ^ v7)
    }
    if (454 > 9223372036854775807) {
        for fi8 := 0; fi8 < 16; fi8++ {
            for fi9 := 0; fi9 < 16; fi9++ {
                v10 := []int{p1, (fi8 & fi9), 32767, (fi9 >> ((18446744073709551615) & 63))}
                fmt.Println(fi9)
                _ = fi9
                _ = v10
            }
            p0 += p2
            _ = fi8
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 255
}
func genFn1(p0 int, p1 int32, p2 int8) int32 {
    p1 -= p1
    if true {
        var v11 uint8 = 128
        _ = v11
    } else {
        fmt.Println(p2, p0, p2)
        for fi12 := 0; fi12 < 1; fi12++ {
            v13 := GenS1{F0: fi12, F1: "\n\t", F2: 1}
            v14 := GenS1{F0: 166, F1: "a", F2: 1}
            for fi15 := 0; fi15 < 19; fi15++ {
                p1 -= p1
                var v16 string = "\n\t"
                var v17 uint8 = 128
                _ = fi15
                _ = v16
                _ = v17
            }
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    return (^(p1))
}

func main() {
    var v1 float32 = (3.14159 + 0.0)
    v2 := []string{}
    v3 := make(map[int]uint32)
    var v4 uint32 = 377
    var v5 uint64 = 42
    var v6 float32 = (v1 * (1e10 + -2.25))
    fmt.Println(v4, v4, v1)
    var v18 uint64 = (v5 << ((149) & 63))
    v18 = (v5 % ((130) | 1))
    v19 := GenS1{F0: 9223372036854775807, F1: "hello", F2: 1}
    v4 = (v4 << ((65535) & 63))
    switch ("foo bar" + (("a" + (("hello" + "foo bar") + ("hello" + "foo bar"))) + "\n\t")) {
    case "hello":
        for fi20 := 0; fi20 < 7; fi20++ {
            var v21 int8 = 42
            _ = fi20
            _ = v21
        }
    case "\n\t":
        var v22 bool = (!(!(((float64(1e10) * 0.0) * 0.0) != -2.25)))
        _ = v22
    default:
        for fi23 := 0; fi23 < 4; fi23++ {
            v2 = append(v2, ("a" + "a"))
            _ = fi23
        }
    }
    switch 1 {
    case -377:
        v18 = v18
    default:
        var v24 bool = true
        _ = v24
    }
    if (true && (65535 <= uint16(v5))) {
        for fi25 := 0; fi25 < 18; fi25++ {
            v26 := GenS2{F0: -1, F1: (false && true)}
            var v27 int = (fi25 * fi25)
            _ = fi25
            _ = v26
            _ = v27
        }
        for fi28 := 0; fi28 < 13; fi28++ {
            v29 := GenS3{F0: 18446744073709551615}
            _ = fi28
            _ = v29
        }
    } else {
        if len(v2) > 0 {
            v2[0] = ("" + ("" + (("foo bar" + "a") + "foo bar")))
        }
        fmt.Println(v6, v5, v5)
    }
    for fi30 := 0; fi30 < 5; fi30++ {
        v4 = 2147483648
        _ = fi30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
