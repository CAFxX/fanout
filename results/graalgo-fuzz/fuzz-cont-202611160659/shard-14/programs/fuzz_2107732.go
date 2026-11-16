package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }

func main() {
    var v1 int = 138
    var v2 string = ("a" + "a")
    v3 := []string{"", v2, ("a" + "a")}
    v4 := []string{((v2 + v2) + "a"), v2}
    var v5 float64 = 1.5
    v6 := make(map[string]uint32)
    var v7 int16 = -32768
    if (18446744073709551615 == 1) {
        if (!(-1 > 1)) {
            fmt.Println(v5, v5)
            v1 = v1
        }
    }
    v1 -= v1
    if len(v4) > 0 {
        v4[0] = "foo bar"
    }
    if len(v3) > 0 {
        v3[0] = (v2 + v2)
    }
    v8 := GenS1{F0: v1, F1: -2147483648}
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 uint64 = 82
        switch v2 {
        case "":
            v11 := v6[(("foo bar" + "a") + "hello")]
            _ = v11
        default:
            var v12 int16 = v7
            _ = v12
        }
        for fi13 := 0; fi13 < 18; fi13++ {
            v7 += (v7 << ((42) & 63))
            v10 = 4294967295
            _ = fi13
        }
        _ = fi9
        _ = v10
    }
    v1 -= (v1 - (v1 & (-(v1))))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch (v2 + v2) {
    case "\n\t":
        var v14 uint16 = 456
        _ = v14
    }
    _ = len(v4)
    v1 += 42
    v1 = v1
    var v15 uint64 = 4294967295
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
