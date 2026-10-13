package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 int16 }

func main() {
    var v1 uint64 = 4294967295
    var v2 float32 = ((1e10 * -2.25) / 0.1)
    v3 := []string{"", "hello", "\n\t", "hello"}
    v4 := []string{"foo bar"}
    fmt.Println(v1)
    var v5 uint32 = 42
    v6 := GenS1{F0: v5, F1: 2, F2: 98}
    _ = len(v4)
    if (false || (("" >= ("foo bar" + ("\n\t" + ("" + ("a" + "\n\t"))))) && (436 > -1))) {
        v2 += v2
        if ("a" != "a") {
            v2 = v2
            var v7 bool = true
            _ = v7
        } else {
            v8 := GenS1{F0: (v5 * v5), F1: 949, F2: -9223372036854775808}
            _ = v8
        }
        v5 += v5
    } else {
        v1 += (v1 >> ((255) & 63))
        var v9 int = -1
        var v10 int32 = -898
        _ = v10
        _ = v9
    }
    v11 := GenS2{F0: 443}
    v5 -= 1
    for fi12 := 0; fi12 < 1; fi12++ {
        _ = len(v3)
        v13 := GenS2{F0: int16(v5)}
        fmt.Println(v5, v1)
        _ = fi12
        _ = v13
    }
    switch (("\n\t" + "") + "a") {
    case "":
        if (!true) {
            fmt.Println(v1)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
    case "a":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "hello":
        v14 := GenS2{F0: 42}
        _ = v14
    default:
        v5 -= (v5 ^ v5)
    }
    v15 := GenS3{F0: (^(v5)), F1: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
