package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 int32 = 921
    var v2 float32 = -2.25
    var v3 int64 = 42
    var v4 uint16 = 1
    var v5 int8 = 62
    var v6 bool = true
    v5 -= (v5 + 100)
    v7 := []int{9223372036854775807, -1, 2147483647}
    var v8 int64 = (v3 % ((v3) | 1))
    _ = len(v7)
    for fi9 := 0; fi9 < 19; fi9++ {
        for fi10 := 0; fi10 < 12; fi10++ {
            v11 := GenS1{F0: v6, F1: v4}
            _ = fi10
            _ = v11
        }
        if len(v7) > 0 {
            v7[0] = -2147483648
        }
        _ = fi9
    }
    _ = len(v7)
    for fi12 := 0; fi12 < 6; fi12++ {
        v13 := GenS2{F0: 0}
        switch (fi12 & fi12) {
        case 0:
            v14 := GenS2{F0: 4294967295}
            _ = v14
        default:
            v15 := GenS2{F0: 202}
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    switch ("hello" + "\n\t") {
    case "":
        if (1 < 2) {
            var v16 bool = (!(642 != 4294967295))
            _ = v16
        }
    case "foo bar":
        var v17 string = "foo bar"
        _ = v17
    default:
        var v18 uint32 = 42
        _ = v18
    }
    var v19 bool = (!(true && ("foo bar" <= "")))
    if len(v7) > 0 {
        v7[0] = 758
    }
    switch "" {
    case "foo bar":
        var v20 uint = 65535
        _ = v20
    case "a":
        for fi21 := 0; fi21 < 18; fi21++ {
            var v22 uint16 = v4
            _ = fi21
            _ = v22
        }
    case "":
        v1 -= (v1 ^ v1)
    default:
        v4 -= v4
    }
    for fi23 := 0; fi23 < 20; fi23++ {
        v24 := GenS1{F0: false, F1: (v4 - (v4 << ((91) & 63)))}
        var v25 string = ""
        fmt.Println(v6)
        _ = fi23
        _ = v24
        _ = v25
    }
    var v26 bool = (3.14159 >= (float64(1.5) - 1.5))
    var v27 float64 = 1e10
    var v28 int16 = 32767
    if (!(-1 <= 474)) {
        var v29 int64 = v3
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
