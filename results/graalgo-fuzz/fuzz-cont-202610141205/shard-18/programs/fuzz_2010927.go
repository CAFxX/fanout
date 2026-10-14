package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int16; F1 uint32 }
func genFn0(p0 int32, p1 uint64) uint16 {
    v6 := GenS1{F0: 52, F1: p1}
    v7 := []int{-482, 1, -424}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 42
}
func genFn1() string {
    var v8 uint16 = 255
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 int = fi9
        v10 = (^(fi9))
        v10 = v10
        _ = fi9
        _ = v10
    }
    _ = v8
    return ("hello" + "")
}

func main() {
    var v1 uint64 = 4294967295
    var v2 int32 = -1
    v3 := []int{-718, -2147483648, 1}
    var v4 uint16 = 234
    var v5 uint8 = 128
    var v11 uint16 = (v4 >> ((230) & 63))
    for fi12 := 0; fi12 < 12; fi12++ {
        if len(v3) > 0 {
            v3[0] = fi12
        }
        _ = fi12
    }
    if (!(!(0 < -1))) {
        if len(v3) > 0 {
            v3[0] = 127
        }
        v2 += -2147483648
    }
    var v13 uint = 4294967295
    for fi14 := 0; fi14 < 15; fi14++ {
        v11 = v11
        _ = fi14
    }
    fmt.Println(v2)
    switch "foo bar" {
    case "hello":
        if len(v3) > 0 {
            v3[0] = 128
        }
    case "a":
        fmt.Println(v4, v13, v4)
    case "\n\t":
        for fi15 := 0; fi15 < 12; fi15++ {
            v16 := GenS1{F0: (^(v5)), F1: uint64(v5)}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi15
            _ = v16
        }
    }
    v11 -= (^(v4))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v17 := GenS2{F0: 397, F1: 65535}
    if len(v3) > 0 {
        v3[0] = 32767
    }
    fmt.Println(v4, v4)
    fmt.Println(v11, v4)
    var v18 uint16 = 0
    v18 = (v4 & v4)
    v19 := GenS1{F0: 159, F1: (v1 ^ v1)}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
