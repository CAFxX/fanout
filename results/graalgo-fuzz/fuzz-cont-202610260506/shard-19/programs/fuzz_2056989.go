package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int16, p1 int8, p2 int16) uint64 {
    p2 -= p0
    v7 := GenS1{F0: (p0 ^ p0)}
    switch -87 {
    case 32767:
        var v8 uint32 = 2147483647
        _ = v8
    case -941:
        switch "hello" {
        case "a":
            if (2 != 0) {
                fmt.Println(p0, p0, p0)
            }
        case "\n\t":
            if false {
                fmt.Println(p2, p1)
            }
        case "hello":
            p1 -= (p1 % (((p1 >> ((893) & 63))) | 1))
        }
    case -1:
        v9 := []int{1, 1}
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 2
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 66
    v3 := make(map[string]bool)
    var v4 string = ((("foo bar" + "foo bar") + "foo bar") + "\n\t")
    var v5 uint = 255
    var v6 uint8 = (v2 - (v2 >> (((v5 + (v5 / (((v5 / ((v5) | 1))) | 1)))) & 63)))
    var v10 uint64 = 4294967295
    v11 := GenS1{F0: 925}
    if ((v5 / (((v5 << ((65535) & 63))) | 1)) != (v5 | v5)) {
        v6 += 1
        for fi12 := 0; fi12 < 4; fi12++ {
            v4 = "foo bar"
            _ = fi12
        }
        var v13 uint16 = 807
        _ = v13
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        var v15 int16 = 134
        v5 = (v5 | v5)
        _ = fi14
        _ = v15
    }
    v1[42] = 1
    v10 += v10
    v16 := []int{}
    v2 += (^(v6))
    v10 = (v10 >> ((578) & 63))
    if len(v16) > 0 {
        v16[0] = 75
    }
    v17 := GenS1{F0: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
