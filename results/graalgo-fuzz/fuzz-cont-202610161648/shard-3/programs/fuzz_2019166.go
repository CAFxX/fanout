package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }

func main() {
    var v1 string = ("foo bar" + ("hello" + "hello"))
    var v2 int = -9223372036854775808
    v3 := []int{v2, 32768, v2, 1}
    v4 := []int{v2, (v2 * -2147483648)}
    v5 := []string{("\n\t" + v1), v1, v1, (v1 + ((v1 + "foo bar") + (v1 + "")))}
    v6 := make(map[uint32]bool)
    fmt.Println(v1, v2)
    var v7 int64 = 113
    var v8 uint8 = 165
    v5 = append(v5, ((v1 + v1) + v1))
    if true {
        v2 = (^(v2))
        for fi9 := 0; fi9 < 16; fi9++ {
            fmt.Println(v2, v7)
            v10 := GenS1{F0: (v7 << ((42) & 63)), F1: 2}
            var v11 uint64 = 0
            _ = fi9
            _ = v10
            _ = v11
        }
    } else {
        v7 = (^(v7))
    }
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = len(v6)
        v2 = v2
    }
    var v12 int64 = v7
    v12 = 470
    if (!(2147483647 != -2147483648)) {
        var v13 uint64 = 2
        for fi14 := 0; fi14 < 10; fi14++ {
            fmt.Println(v13, v7)
            v2 -= -2147483648
            v7 += v7
            _ = fi14
        }
        _ = v13
    }
    v15 := GenS1{F0: v7, F1: 9}
    v12 = (v12 | 1)
    var v16 float64 = 0.0
    v7 += (v12 % ((1) | 1))
    var v17 uint8 = (v8 % ((v8) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
