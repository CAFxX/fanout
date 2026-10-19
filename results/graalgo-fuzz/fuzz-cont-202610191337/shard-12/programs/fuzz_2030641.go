package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int32 }
type GenS2 struct { F0 uint; F1 uint32; F2 int32 }

func main() {
    var v1 float64 = 3.14159
    var v2 uint = 42
    v3 := []string{"hello", "foo bar", "a"}
    var v4 int32 = -1
    var v5 int32 = (-(v4))
    var v6 float64 = ((float64(0.0) + v1) * float64(1e10))
    _ = len(v3)
    v6 -= v6
    v7 := GenS1{F0: 521, F1: 127, F2: (v5 * (v5 ^ 42))}
    fmt.Println(v4, v1, v1)
    var v8 uint64 = 712
    fmt.Println(v5)
    for fi9 := 0; fi9 < 4; fi9++ {
        switch (fi9 >> (((^(v2))) & 63)) {
        case 32767:
            var v10 int8 = int8(v4)
            _ = v10
        default:
            var v11 float64 = v1
            _ = v11
        }
        _ = fi9
    }
    switch 253 {
    case 815:
        switch -2147483648 {
        case 633:
            fmt.Println(v1, v6)
        }
    case 42:
        var v12 uint8 = 236
        _ = v12
    case 32767:
        if len(v3) > 0 {
            v3[0] = "hello"
        }
    default:
        if (1000 == -870) {
            if true {
                v13 := GenS1{F0: (^(v8)), F1: 0, F2: v4}
                _ = v13
            } else {
                var v14 int8 = -1
                fmt.Println(v6, v8)
                if true {
                    var v15 int32 = (^(v5))
                    switch "a" {
                    case "hello":
                        var v16 int8 = v14
                        _ = v16
                    case "a":
                        v1 -= 1.5
                    case "":
                        var v17 uint64 = (v8 >> ((65535) & 63))
                        _ = v17
                    }
                    v18 := GenS2{F0: (v2 << ((v2) & 63)), F1: 2147483647, F2: -616}
                    _ = v15
                    _ = v18
                }
                _ = v14
            }
        }
    }
    fmt.Println(v8, v5, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
