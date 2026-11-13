package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int64 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 int32; F1 string; F2 uint64 }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 int32; F1 uint8; F2 uint32 }

func main() {
    var v1 int8 = 0
    v2 := make(map[string]int)
    var v3 uint8 = 56
    var v4 string = "hello"
    var v5 int = 420
    v6 := []string{"a", v4}
    if true {
        switch v4 {
        case "\n\t":
            v7 := GenS1{F0: 9223372036854775807, F1: 2, F2: -1}
            _ = v7
        case "foo bar":
            var v8 string = ((v4 + v4) + "hello")
            _ = v8
        }
        var v9 int32 = -796
        if (65 == 32767) {
            v4 = v4
        } else {
            _ = len(v2)
        }
        _ = v9
    }
    v4 = v4
    switch 286 {
    case -9223372036854775808:
        for fi10 := 0; fi10 < 3; fi10++ {
            v1 = v1
            var v11 bool = false
            _ = fi10
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        for fi13 := 0; fi13 < 7; fi13++ {
            _ = len(v2)
            v5 = fi12
            var v14 string = "foo bar"
            v3 -= (v3 * (v3 * v3))
            _ = fi13
            _ = v14
        }
        switch v4 {
        case "a":
            if len(v6) > 0 {
                v6[0] = ""
            }
        case "\n\t":
            v15 := GenS2{F0: 2, F1: -2147483648}
            _ = v15
        case "":
            v16 := GenS3{F0: -2147483648, F1: v4, F2: 4294967295}
            _ = v16
        }
        _ = fi12
    }
    v17 := GenS3{F0: 14, F1: (v4 + v4), F2: 42}
    v4 = v4
    v5 -= (v5 - int(v1))
    switch (v5 << ((525) & 63)) {
    case -249:
        v2[v4] = v5
    case 1:
        switch (-(v5)) {
        case 42:
            var v18 uint64 = 394
            _ = v18
        case 32767:
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        v4 = v4
        v4 = v4
        if (false && (!(!false))) {
            var v20 int32 = 0
            _ = v20
        }
        _ = fi19
    }
    var v21 bool = (231 < -2147483648)
    v22 := GenS2{F0: 2, F1: -87}
    v23 := GenS4{F0: v5, F1: 1}
    var v24 uint8 = (v3 * v3)
    var v25 float32 = 1e10
    if v21 {
        switch "" {
        case "\n\t":
            v26 := GenS3{F0: 0, F1: "", F2: 4294967295}
            _ = v26
        default:
            fmt.Println(v25, v24, v21)
        }
        var v27 int32 = -2147483648
        fmt.Println(v5)
        _ = v27
    }
    if v21 {
        for fi28 := 0; fi28 < 14; fi28++ {
            var v29 uint64 = 233
            _ = fi28
            _ = v29
        }
        v30 := GenS5{F0: 42, F1: v3, F2: 42}
        _ = v30
    }
    v2["\n\t"] = (v5 << ((2) & 63))
    var v31 float32 = float32((3.14159 - ((float64(v25) + ((0.1 - 0.1) + (1.5 / 0.1))) + 0.0)))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
