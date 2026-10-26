package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 uint8; F1 bool; F2 int32 }

func main() {
    var v1 int32 = 1
    var v2 int = 428
    v3 := make(map[string]bool)
    v4 := []int{v2, v2}
    v5 := []uint32{}
    var v6 uint64 = 338
    v7 := GenS1{F0: 969}
    if (73 <= 123) {
        if false {
            var v8 int64 = -956
            switch "\n\t" {
            case "\n\t":
                fmt.Println(v6, v2)
            }
            _ = v8
        }
    }
    if (2 <= 0) {
        v2 -= (v2 >> ((255) & 63))
        v9 := GenS1{F0: v1}
        _ = v9
    }
    if (42 != -1) {
        for fi10 := 0; fi10 < 15; fi10++ {
            fmt.Println(v1, v6, v1)
            _ = fi10
        }
        v2 += v2
        var v11 string = ("foo bar" + ("foo bar" + "a"))
        _ = v11
    } else {
        var v12 int32 = v1
        for fi13 := 0; fi13 < 5; fi13++ {
            v6 += v6
            _ = fi13
        }
        _ = v12
    }
    v14 := GenS2{F0: 18446744073709551615}
    switch (v2 ^ (v2 + -9223372036854775808)) {
    case 32768:
        for fi15 := 0; fi15 < 20; fi15++ {
            v2 += (fi15 * (v2 >> ((159) & 63)))
            v2 += 295
            _ = fi15
        }
    }
    var v16 uint8 = 102
    var v17 int16 = 42
    switch 42 {
    case -20:
        v1 = (^(v1))
    case -582:
        for fi18 := 0; fi18 < 16; fi18++ {
            var v19 int32 = (v1 - (v1 << ((255) & 63)))
            _ = len(v4)
            v20 := GenS3{F0: 0, F1: v16}
            _ = fi18
            _ = v19
            _ = v20
        }
    default:
        for fi21 := 0; fi21 < 19; fi21++ {
            var v22 uint32 = 0
            _ = fi21
            _ = v22
        }
    }
    v6 = 42
    v23 := GenS4{F0: 245, F1: (true || true), F2: -1}
    switch (("" + "a") + ("a" + ("a" + ("a" + "a")))) {
    case "\n\t":
        for fi24 := 0; fi24 < 18; fi24++ {
            v6 = (v6 + (v6 / ((v6) | 1)))
            _ = fi24
        }
    case "foo bar":
        v1 = 0
    case "a":
        fmt.Println(v2, v2, v6)
    default:
        v1 = 652
    }
    v6 -= v6
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
