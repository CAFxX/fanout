package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int64 }
type GenS2 struct { F0 uint }

func main() {
    var v1 int16 = 910
    var v2 float32 = -2.25
    var v3 uint = 151
    var v4 uint8 = 65
    for fi5 := 0; fi5 < 12; fi5++ {
        fmt.Println(v3, v3)
        _ = fi5
    }
    v2 = (v2 / ((((((0.1 + (v2 / (v2 / 1e10))) + v2) * float32(-2.25)) - v2) - v2) - float32((-2.25 / 3.14159))))
    var v6 int64 = -230
    v6 = v6
    v7 := []int{}
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := GenS1{F0: 127, F1: (v3 ^ v3), F2: (v6 - 1)}
        v10 := GenS1{F0: 42, F1: (v3 >> ((v3) & 63)), F2: v6}
        _ = len(v7)
        _ = fi8
        _ = v10
        _ = v9
    }
    v1 -= v1
    switch (("a" + ("a" + "")) + "hello") {
    case "hello":
        var v11 int = 565
        _ = v11
    }
    v12 := GenS1{F0: -1, F1: (v3 | 4294967295), F2: (v6 / ((9223372036854775807) | 1))}
    switch -948 {
    case -741:
        fmt.Println(v3, v2, v6)
    case -749:
        if (true && (32767 >= v1)) {
            v1 = v1
        } else {
            var v13 uint32 = 283
            v14 := GenS2{F0: (v3 ^ uint(v4))}
            _ = v13
            _ = v14
        }
    case 4294967295:
        var v15 int = 128
        _ = v15
    }
    var v16 float32 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
