package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 int16 }
func genFn0(p0 bool, p1 uint32) uint64 {
    p1 += p1
    v4 := []int{759, -725}
    v5 := GenS1{F0: -32768}
    if p0 {
        v6 := GenS1{F0: -32768}
        p1 -= p1
        v7 := GenS1{F0: 0}
        _ = v6
        _ = v7
    } else {
        p0 = (-1 < 42)
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 42
}

func main() {
    v1 := []string{"foo bar", "a", "foo bar", "\n\t"}
    v2 := []string{"foo bar", "\n\t"}
    var v3 uint8 = 12
    var v8 int = -2147483648
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 int64 = 264
        _ = fi9
        _ = v10
    }
    var v11 uint64 = 18446744073709551615
    v3 = v3
    _ = len(v1)
    var v12 int8 = -128
    if (v12 == v12) {
        switch (("\n\t" + ("a" + "")) + "") {
        case "foo bar":
            switch ("\n\t" + "foo bar") {
            case "a":
                for fi13 := 0; fi13 < 4; fi13++ {
                    var v14 int64 = 1
                    v2 = append(v2, ("\n\t" + "hello"))
                    _ = fi13
                    _ = v14
                }
            default:
                var v15 int16 = 871
                _ = v15
            }
        case "hello":
            fmt.Println(v12)
        }
        v8 = (v8 % (((v8 << ((42) & 63))) | 1))
        v8 += v8
    } else {
        var v16 int64 = -31
        if true {
            var v17 int16 = 487
            var v18 float64 = -2.25
            _ = v17
            _ = v18
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = len(v2)
            v2 = append(v2, ("" + "a"))
        }
        var v19 uint16 = 42
        _ = v16
        _ = v19
    }
    v8 -= (v8 ^ 32768)
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    v20 := GenS2{F0: 18, F1: -412}
    switch "" {
    case "hello":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        _ = len(v1)
    }
    _ = len(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v8
}
