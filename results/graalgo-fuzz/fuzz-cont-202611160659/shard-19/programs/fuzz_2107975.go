package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint32 }
func genFn0() string {
    v7 := []int{32767, 9223372036854775807, 2147483647, 9223372036854775807}
    switch -149 {
    case 4294967295:
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 uint32 = 42
            _ = fi8
            _ = v9
        }
    case 32768:
        v10 := GenS1{F0: 79, F1: -9223372036854775808, F2: 42}
        _ = v10
    default:
        var v11 string = "a"
        _ = v11
    }
    _ = v7
    return "\n\t"
}

func main() {
    var v1 uint16 = 42
    var v2 uint8 = 255
    var v3 bool = ((3.14159 < 3.14159) || true)
    var v4 string = ""
    v5 := make(map[int]bool)
    var v6 int16 = 42
    var v12 int = -685
    for fi13 := 0; fi13 < 7; fi13++ {
        fmt.Println(fi13, v2)
        var v14 uint32 = 4294967295
        _ = fi13
        _ = v14
    }
    var v15 string = v4
    switch (v12 | v12) {
    case 9223372036854775807:
        if ((268 != 9223372036854775807) || true) {
            v1 += v1
        }
    case 858:
        _ = len(v5)
    case 4294967295:
        v3 = (3.14159 < float32((1.5 / 3.14159)))
    }
    fmt.Println(v6)
    v3 = v3
    if v3 {
        var v16 uint = 903
        for fi17 := 0; fi17 < 8; fi17++ {
            v4 = "\n\t"
            _ = fi17
        }
        _ = v16
    }
    switch "a" {
    case "foo bar":
        switch 1 {
        case -689:
            fmt.Println(v4, v3, v6)
        case -869:
            v2 -= (v2 << ((718) & 63))
        case 813:
            v6 = (v6 ^ v6)
        default:
            switch v12 {
            case -2147483648:
                for fi18 := 0; fi18 < 12; fi18++ {
                    v19 := []int{(^(fi18)), fi18}
                    for fi20 := 0; fi20 < 19; fi20++ {
                        fmt.Println(v3)
                        v5[(v12 << ((949) & 63))] = v3
                        v15 = ("a" + v15)
                        _ = fi20
                    }
                    _ = fi18
                    _ = v19
                }
            case -1:
                fmt.Println(v4, v15)
            }
        }
    case "":
        var v21 uint8 = 55
        _ = v21
    case "hello":
        var v22 float32 = ((3.14159 - 1e10) + 1.5)
        _ = v22
    default:
        fmt.Println(v15, v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
