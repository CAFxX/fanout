package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int64 }
type GenS2 struct { F0 int16; F1 bool; F2 uint16 }
func genFn0(p0 uint32) uint16 {
    switch 2147483647 {
    case 32767:
        v6 := GenS1{F0: 42, F1: -747, F2: 471}
        _ = v6
    }
    var v7 int8 = -118
    v8 := GenS1{F0: 127, F1: 146, F2: -1}
    _ = p0
    _ = v7
    _ = v8
    return 32767
}

func main() {
    v1 := []uint32{42, 1}
    v2 := make(map[uint32]string)
    var v3 int = -9223372036854775808
    var v4 bool = (533 == 1)
    v5 := []string{"foo bar", "foo bar", "hello", ("hello" + "a")}
    v9 := GenS1{F0: 1, F1: v3, F2: 9223372036854775807}
    switch "a" {
    case "hello":
        fmt.Println(v4, v4)
    }
    fmt.Println(v3)
    fmt.Println(v4)
    if (-150 >= -365) {
        v4 = v4
        v10 := GenS1{F0: 1, F1: (v3 << ((271) & 63)), F2: -974}
        _ = v10
    }
    v4 = (0.1 == 1e10)
    for fi11 := 0; fi11 < 18; fi11++ {
        delete(v2, 65535)
        if (!v4) {
            var v12 bool = (v4 || false)
            for fi13 := 0; fi13 < 18; fi13++ {
                delete(v2, 2147483648)
                _ = fi13
            }
            v14 := GenS1{F0: 126, F1: v3, F2: -1}
            _ = v12
            _ = v14
        }
        v15 := GenS2{F0: 42, F1: true, F2: 255}
        _ = fi11
        _ = v15
    }
    v16 := GenS1{F0: 100, F1: v3, F2: -963}
    fmt.Println(v4, v4)
    var v17 uint = 42
    v4 = false
    if (0 <= -9223372036854775808) {
        switch "foo bar" {
        case "hello":
            var v18 uint64 = 2
            _ = v18
        default:
            v17 -= v17
        }
    } else {
        _ = len(v1)
        switch "" {
        case "\n\t":
            v2[1] = ("\n\t" + "a")
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
