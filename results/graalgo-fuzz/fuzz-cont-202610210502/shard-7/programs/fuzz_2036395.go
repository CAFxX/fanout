package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 128
    var v3 string = ("foo bar" + ((("a" + (("\n\t" + ("" + ((("hello" + "") + "hello") + ""))) + "")) + "") + "a"))
    var v4 int8 = 95
    v5 := GenS1{F0: "hello", F1: 0}
    var v6 int32 = 2147483647
    for fi7 := 0; fi7 < 12; fi7++ {
        _ = len(v1)
        _ = fi7
    }
    v1[((v3 + (v3 + "\n\t")) + v3)] = (-9223372036854775808 > 2147483647)
    if (2 < 18446744073709551615) {
        if ((v6 * (v6 << ((18446744073709551615) & 63))) < (v6 - 0)) {
            v4 = (v4 / (((v4 ^ v4)) | 1))
        } else {
            fmt.Println(v4, v3)
            v8 := GenS1{F0: v3, F1: 42}
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        fmt.Println(v3)
        _ = fi9
    }
    var v10 int32 = (-(v6))
    var v11 string = v3
    v12 := GenS1{F0: (("\n\t" + v11) + ""), F1: 2}
    v13 := []int{4294967295}
    v3 = ("hello" + v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
