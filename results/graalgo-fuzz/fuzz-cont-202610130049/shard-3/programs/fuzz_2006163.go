package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
func genFn0(p0 int16, p1 uint64) uint32 {
    for fi7 := 0; fi7 < 14; fi7++ {
        fmt.Println(p0)
        _ = fi7
    }
    switch "" {
    case "a":
        p1 += p1
    case "\n\t":
        if true {
            fmt.Println(p0, p1)
        }
    }
    _ = p0
    _ = p1
    return 389
}
func genFn1(p0 int, p1 int16) uint {
    p1 -= (p1 & (p1 ^ (^(p1))))
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return 395
}

func main() {
    var v1 int64 = 42
    var v2 uint8 = 10
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]uint32)
    var v5 int64 = v1
    var v6 bool = (-1 != 1000)
    v8 := GenS1{F0: (0 > 2147483648), F1: -784}
    fmt.Println(v5, v1)
    for fi9 := 0; fi9 < 13; fi9++ {
        fmt.Println(v2, v2)
        switch ((("foo bar" + "hello") + "foo bar") + "a") {
        case "a":
            fmt.Println(v6, v5, v2)
        case "\n\t":
            var v10 int32 = 0
            _ = v10
        }
        _ = fi9
    }
    var v11 int64 = -9223372036854775808
    v12 := []int{127, 128, 9223372036854775807, -528}
    v13 := GenS1{F0: v6, F1: 1000}
    v2 = (^(v2))
    for fi14 := 0; fi14 < 20; fi14++ {
        for fi15 := 0; fi15 < 17; fi15++ {
            fmt.Println(v5, v1, fi14)
            _ = fi15
        }
        _ = fi14
    }
    if v6 {
        v16 := GenS1{F0: (2 <= 255), F1: -32768}
        v1 = -1
        fmt.Println(v5)
        _ = v16
    }
    if (!(3.14159 >= 1.5)) {
        switch "\n\t" {
        case "hello":
            var v17 uint8 = (v2 << ((65535) & 63))
            _ = v17
        case "a":
            if v6 {
                if (!true) {
                    fmt.Println(v2, v11)
                    var v18 bool = v6
                    var v19 uint8 = v2
                    _ = v18
                    _ = v19
                }
                for fi20 := 0; fi20 < 1; fi20++ {
                    fmt.Println(v5)
                    _ = fi20
                }
            }
        }
    }
    v1 -= -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
