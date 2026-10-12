package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }

func main() {
    var v1 int16 = 620
    var v2 uint8 = 2
    var v3 uint32 = 4294967295
    var v4 int64 = -511
    v5 := []int{-1, 9223372036854775807, -426, 32767}
    var v6 int16 = v1
    v5 = append(v5, 128)
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 float64 = float64(3.14159)
        v5 = append(v5, -9223372036854775808)
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 18; fi9++ {
        v10 := GenS1{F0: "foo bar", F1: -128}
        v5 = append(v5, 466)
        _ = fi9
        _ = v10
    }
    switch 641 {
    case 127:
        if len(v5) > 0 {
            v5[0] = int(v6)
        }
    case 0:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    v11 := GenS1{F0: (("\n\t" + (("" + "hello") + "foo bar")) + "\n\t"), F1: 88}
    fmt.Println(v3)
    var v12 int16 = v6
    var v13 int64 = (v4 << ((585) & 63))
    var v14 int16 = (v12 ^ (-(v1)))
    for fi15 := 0; fi15 < 14; fi15++ {
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 string = "a"
            var v18 float64 = (((1e10 + 3.14159) - -2.25) - -2.25)
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi15
    }
    var v19 int64 = v4
    var v20 bool = false
    v5 = append(v5, -1)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
