package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint8 }
type GenS2 struct { F0 uint16 }
func genFn0() uint32 {
    v4 := GenS1{F0: 770, F1: ("" + (("\n\t" + "foo bar") + "hello")), F2: 255}
    var v5 uint8 = 255
    _ = v4
    _ = v5
    return 843
}
func genFn1(p0 int8) uint32 {
    switch "\n\t" {
    case "":
        fmt.Println(p0)
    }
    var v6 float32 = (0.1 - 0.1)
    p0 += p0
    v6 += v6
    v7 := []int{}
    _ = p0
    _ = v6
    _ = v7
    return 2147483648
}

func main() {
    var v1 int16 = -529
    var v2 int32 = -1
    v3 := []string{"foo bar", (("foo bar" + "\n\t") + "\n\t"), (("a" + "foo bar") + ("\n\t" + "a"))}
    v2 += -576
    var v8 int32 = (v2 ^ (v2 << ((2) & 63)))
    fmt.Println(v2, v8, v1)
    v9 := GenS2{F0: 0}
    var v10 string = "a"
    v10 = "hello"
    if true {
        var v11 int64 = 42
        _ = v11
    }
    var v12 uint16 = 434
    var v13 uint32 = 840
    v3 = append(v3, v10)
    for fi14 := 0; fi14 < 5; fi14++ {
        v12 = (v12 / (((v12 << ((4294967295) & 63))) | 1))
        for fi15 := 0; fi15 < 4; fi15++ {
            v16 := GenS2{F0: (v12 ^ (v12 + (v12 * 0)))}
            v8 -= (^(v2))
            var v17 uint = uint(v13)
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
