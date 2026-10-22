package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
func genFn0(p0 int32) int {
    p0 += (p0 + p0)
    var v7 int64 = -224
    _ = p0
    _ = v7
    return int(p0)
}

func main() {
    var v1 int16 = 42
    var v2 float32 = (1e10 / (float32(1.5) / 1e10))
    v3 := make(map[uint32]bool)
    var v4 int32 = -1
    var v5 uint64 = 4294967295
    var v6 float64 = (0.1 - (1e10 - -2.25))
    _ = len(v3)
    switch 2147483647 {
    case -935:
        var v8 float64 = v6
        _ = v8
    case 631:
        var v9 float64 = (-2.25 + v6)
        _ = v9
    }
    var v10 uint8 = 127
    var v11 uint8 = 132
    if (v1 <= v1) {
        v2 += 0.0
    }
    v5 -= v5
    for fi12 := 0; fi12 < 18; fi12++ {
        var v13 float32 = ((float32(v6) - 3.14159) / (3.14159 + v2))
        _ = fi12
        _ = v13
    }
    if ((v11 | v10) > (v11 * v11)) {
        v14 := GenS1{F0: 0, F1: 1}
        switch "hello" {
        case "":
            v1 = (v1 << ((0) & 63))
        case "foo bar":
            switch "\n\t" {
            case "\n\t":
                v15 := []int{42, 2147483647, 776, 9223372036854775807}
                _ = v15
            case "hello":
                v1 = 1000
            }
        case "hello":
            v1 = (v1 >> ((4294967295) & 63))
        }
        _ = v14
    } else {
        v10 = (v10 & 44)
    }
    if (true && ((v2 * 0.0) != 0.1)) {
        var v16 uint64 = v5
        _ = v16
    }
    for fi17 := 0; fi17 < 5; fi17++ {
        v18 := GenS1{F0: 42, F1: 100}
        v10 -= 128
        _ = fi17
        _ = v18
    }
    if false {
        v4 += v4
    }
    v11 = v11
    switch ("" + ("\n\t" + "a")) {
    case "":
        for fi19 := 0; fi19 < 5; fi19++ {
            fmt.Println(v4)
            _ = fi19
        }
    case "hello":
        v6 -= float64(v2)
    default:
        if (2 != 4294967295) {
            v1 += -32768
            v6 -= 1e10
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
