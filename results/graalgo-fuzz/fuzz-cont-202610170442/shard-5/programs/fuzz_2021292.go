package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int }
func genFn0(p0 int16, p1 uint32, p2 int16) uint {
    v5 := []int{-2147483648, int(p2)}
    switch "a" {
    case "hello":
        p2 += (p0 | p2)
    case "\n\t":
        fmt.Println(p2)
    case "":
        switch 2147483647 {
        case 901:
            _ = len(v5)
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 39
}

func main() {
    v1 := []uint32{42, 598, 61, 0}
    var v2 uint = 4294967295
    var v3 uint8 = 1
    var v4 float32 = float32(0.1)
    var v6 int32 = -273
    v1 = append(v1, 42)
    if (!true) {
        v6 = v6
    }
    fmt.Println(v2, v3, v2)
    v7 := GenS1{F0: 0, F1: v3, F2: -626}
    var v8 int64 = int64(v6)
    v1 = append(v1, 4294967295)
    var v9 uint16 = uint16(v2)
    var v10 int32 = (v6 + (v6 ^ v6))
    if len(v1) > 0 {
        v1[0] = 0
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        for fi12 := 0; fi12 < 17; fi12++ {
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
            var v13 int64 = (-(v8))
            _ = len(v1)
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 14; fi14++ {
            switch (("foo bar" + (("foo bar" + "hello") + "foo bar")) + "foo bar") {
            case "":
                var v15 int16 = 354
                _ = v15
            case "a":
                if true {
                    if true {
                        var v16 bool = true
                        _ = v16
                    }
                } else {
                    v9 = v9
                }
            case "hello":
                v1 = append(v1, 4294967295)
            }
            _ = fi14
        }
        _ = fi11
    }
    var v17 string = "hello"
    v1 = append(v1, 2147483647)
    v8 -= -556
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
