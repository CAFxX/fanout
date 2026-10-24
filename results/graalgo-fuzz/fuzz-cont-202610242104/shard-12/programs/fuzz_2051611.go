package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint8; F1 uint16 }

func main() {
    var v1 uint = 0
    var v2 int = 2147483647
    var v3 int8 = 95
    var v4 int32 = 0
    var v5 uint16 = 42
    v6 := []int{}
    v7 := GenS1{F0: v2, F1: 1}
    fmt.Println(v3, v2, v4)
    v8 := GenS1{F0: v2, F1: 42}
    v9 := GenS1{F0: -2147483648, F1: -2147483648}
    var v10 uint32 = 0
    switch (("foo bar" + "a") + "") {
    case "hello":
        var v11 uint = 1
        _ = v11
    case "foo bar":
        if len(v6) > 0 {
            v6[0] = 2147483647
        }
    case "a":
        v12 := GenS2{F0: 227, F1: 1}
        _ = v12
    }
    switch v2 {
    case 0:
        if len(v6) > 0 {
            v6[0] = v2
        }
    case 2147483647:
        if ((v4 & (v4 ^ v4)) < (-(v4))) {
            v2 = 2147483647
            fmt.Println(v4, v3)
        }
    }
    switch (v2 % ((v2) | 1)) {
    case 127:
        var v13 uint = (v1 & (v1 ^ v1))
        _ = v13
    case 906:
        v5 += 565
    case 0:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    default:
        fmt.Println(v10, v4)
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        v15 := GenS2{F0: 2, F1: 2}
        if (((^(v4)) == (v4 * (v4 | -2147483648))) || ((2 <= v5) && (!(!true)))) {
            var v16 uint32 = v10
            _ = v16
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi14
        _ = v15
    }
    v4 += v4
    v2 += (v2 % ((v2) | 1))
    if (64 == 1) {
        v2 = (v2 & (v2 + (^(v2))))
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        v1 += (v1 ^ v1)
        v3 = 125
        var v18 int8 = v3
        _ = fi17
        _ = v18
    }
    v4 += 0
    var v19 int8 = v3
    if true {
        fmt.Println(v10, v19)
        fmt.Println(v2, v3)
        if (!((-(v2)) <= (v2 | 581))) {
            v2 = v2
            var v20 uint = 42
            _ = v20
        }
    } else {
        var v21 bool = ("foo bar" < ("\n\t" + "hello"))
        v10 += (v10 - 42)
        _ = v21
    }
    if true {
        v22 := GenS1{F0: (v2 + (-(v2))), F1: v4}
        switch (v2 ^ v2) {
        case 32767:
            v23 := GenS1{F0: 9223372036854775807, F1: -951}
            _ = v23
        }
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
