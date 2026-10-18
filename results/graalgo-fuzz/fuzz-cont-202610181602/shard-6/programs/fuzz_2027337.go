package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint }

func main() {
    v1 := []string{"", (("a" + ("" + "hello")) + "foo bar")}
    var v2 uint8 = 42
    var v3 int64 = 339
    switch "hello" {
    case "a":
        v2 = (v2 >> ((42) & 63))
    case "foo bar":
        var v4 uint16 = 324
        _ = v4
    case "":
        _ = len(v1)
    default:
        var v5 uint = 4294967295
        _ = v5
    }
    v6 := GenS1{F0: v3, F1: 715}
    if true {
        v7 := GenS1{F0: (v3 & (^(v3))), F1: 905}
        var v8 int32 = 263
        _ = v7
        _ = v8
    } else {
        if false {
            v9 := GenS1{F0: (v3 % (((v3 / ((v3) | 1))) | 1)), F1: 398}
            var v10 uint8 = v2
            _ = v10
            _ = v9
        } else {
            var v11 float64 = ((float64((((1e10 / (-2.25 + 1.5)) - 0.1) * -2.25)) - 0.1) / 3.14159)
            _ = v11
        }
        v12 := GenS2{F0: 2}
        v3 -= 42
        _ = v12
    }
    var v13 uint = 18446744073709551615
    for fi14 := 0; fi14 < 5; fi14++ {
        var v15 uint = (v13 ^ (v13 - 40))
        _ = fi14
        _ = v15
    }
    v16 := GenS2{F0: v13}
    for fi17 := 0; fi17 < 19; fi17++ {
        for fi18 := 0; fi18 < 18; fi18++ {
            v19 := GenS2{F0: v13}
            var v20 uint64 = 42
            v3 += (v3 + int64(fi17))
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi17
    }
    switch ("\n\t" + "") {
    case "\n\t":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        var v21 uint64 = 4294967295
        _ = v21
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    v13 -= v13
    if (true || (219 >= 1)) {
        var v22 int32 = 42
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v22
    } else {
        fmt.Println(v13)
    }
    var v23 uint64 = 392
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v6
}
