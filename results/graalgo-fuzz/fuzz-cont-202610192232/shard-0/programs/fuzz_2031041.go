package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]bool)
    v3 := []int{-29}
    v4 := make(map[int]string)
    var v5 int32 = -2147483648
    v5 += v5
    if false {
        v6 := GenS1{F0: 2147483648}
        var v7 int8 = 0
        _ = v6
        _ = v7
    } else {
        v5 -= -1
        var v8 int16 = 845
        _ = v8
    }
    switch 2147483647 {
    case 9223372036854775807:
        switch ((("\n\t" + "\n\t") + ("" + ("" + ""))) + "foo bar") {
        case "foo bar":
            v5 = (v5 * (v5 ^ (v5 & (v5 & -1))))
        case "hello":
            var v9 string = ("" + ("" + ("a" + ("foo bar" + "hello"))))
            _ = v9
        }
    case 2147483647:
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 bool = (true || (!("foo bar" != ("foo bar" + "a"))))
            var v12 int32 = v5
            var v13 int8 = 0
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
    default:
        var v14 bool = true
        _ = v14
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        v2[(fi15 * (fi15 & (fi15 + (fi15 % ((32768) | 1)))))] = true
        _ = fi15
    }
    v5 = v5
    v16 := GenS1{F0: 762}
    var v17 int32 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
