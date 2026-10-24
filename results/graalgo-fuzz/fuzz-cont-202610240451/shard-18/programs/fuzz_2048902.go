package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int64 }
func genFn0() uint16 {
    var v6 uint64 = 1
    v7 := GenS1{F0: 532, F1: -2147483648, F2: 9223372036854775807}
    v6 = 786
    v6 += (v6 ^ 42)
    _ = v6
    _ = v7
    return 322
}

func main() {
    var v1 bool = true
    var v2 uint8 = 220
    var v3 string = ("\n\t" + "\n\t")
    v4 := []int{}
    v5 := make(map[uint32]bool)
    v3 = ("hello" + v3)
    v3 = ("a" + v3)
    if v1 {
        fmt.Println(v1, v1)
    }
    fmt.Println(v2)
    switch 32768 {
    case 128:
        fmt.Println(v1, v2)
    case 0:
        v2 = (v2 | (^(v2)))
    case -391:
        v1 = (v1 && v1)
    default:
        v2 += v2
    }
    var v8 string = ((v3 + v3) + v3)
    if (false && true) {
        for fi9 := 0; fi9 < 11; fi9++ {
            v8 = (v8 + ((("" + v3) + "hello") + ""))
            var v10 bool = v1
            fmt.Println(v1, v1)
            _ = fi9
            _ = v10
        }
    }
    if (-9223372036854775808 <= -2147483648) {
        var v11 uint8 = (v2 ^ (v2 ^ v2))
        _ = v11
    }
    var v12 uint32 = 42
    fmt.Println(v12)
    _ = len(v5)
    fmt.Println(v3)
    v2 -= v2
    v2 = 2
    v3 = ((v3 + (v3 + "\n\t")) + (v3 + "\n\t"))
    for fi13 := 0; fi13 < 7; fi13++ {
        v1 = ((4294967295 < v12) || v1)
        var v14 int = fi13
        for fi15 := 0; fi15 < 14; fi15++ {
            var v16 uint = 255
            _ = fi15
            _ = v16
        }
        _ = fi13
        _ = v14
    }
    var v17 string = ((v3 + v8) + v8)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
