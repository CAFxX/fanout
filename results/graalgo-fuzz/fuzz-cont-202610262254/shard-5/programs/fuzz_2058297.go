package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint }
func genFn0() string {
    var v4 int8 = -103
    v5 := []int{848}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = v4
    _ = v5
    return (("foo bar" + "") + "foo bar")
}

func main() {
    var v1 uint32 = 424
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 42
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 uint64 = (^(v3))
        var v8 bool = true
        _ = fi6
        _ = v7
        _ = v8
    }
    v3 -= (v3 / (((^(v3))) | 1))
    for fi9 := 0; fi9 < 6; fi9++ {
        v10 := v2[v1]
        switch "foo bar" {
        case "":
            var v11 uint8 = 128
            _ = v11
        case "foo bar":
            if (true || false) {
                v10 -= 65535
            } else {
                v10 += 926
            }
        }
        _ = fi9
        _ = v10
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        v1 += 738
        v13 := GenS1{F0: -2147483648, F1: ("foo bar" + (("a" + "a") + "hello")), F2: 355}
        var v14 uint32 = 2147483647
        _ = fi12
        _ = v13
        _ = v14
    }
    v3 = (^(v3))
    var v15 int = -632
    var v16 bool = (true || true)
    delete(v2, (v1 % ((v1) | 1)))
    var v17 int64 = 831
    switch (v15 >> ((89) & 63)) {
    case 1:
        var v18 uint64 = (v3 * 0)
        _ = v18
    case 32767:
        if v16 {
            v19 := v2[1]
            _ = v19
        } else {
            v16 = false
        }
    case -329:
        var v20 int = v15
        _ = v20
    default:
        v2[42] = (^(v1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
