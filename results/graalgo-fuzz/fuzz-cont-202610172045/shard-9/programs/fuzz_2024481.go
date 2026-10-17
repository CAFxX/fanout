package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint8; F1 string; F2 bool }
type GenS3 struct { F0 bool; F1 int32; F2 uint8 }
type GenS4 struct { F0 bool }

func main() {
    var v1 uint64 = 0
    var v2 uint16 = 0
    var v3 int8 = 1
    v4 := make(map[string]string)
    v5 := GenS1{F0: 0, F1: 745}
    v1 += 0
    var v6 bool = ((false || (!(!true))) || true)
    v7 := GenS1{F0: 4294967295, F1: 0}
    var v8 int16 = -32768
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := GenS2{F0: 90, F1: "a", F2: (-268 == -197)}
        _ = fi9
        _ = v10
    }
    v4["foo bar"] = "hello"
    if v6 {
        v11 := GenS2{F0: 2, F1: "", F2: v6}
        var v12 uint = 42
        _ = v11
        _ = v12
    }
    var v13 int64 = 430
    for fi14 := 0; fi14 < 5; fi14++ {
        for fi15 := 0; fi15 < 16; fi15++ {
            var v16 bool = ("\n\t" <= (("a" + (("" + (("\n\t" + (("a" + "a") + "hello")) + ((("" + "hello") + "\n\t") + "hello"))) + ("a" + ("\n\t" + "\n\t")))) + "\n\t"))
            if ((v1 ^ v1) < v1) {
                v17 := GenS3{F0: v6, F1: 1, F2: 42}
                _ = v17
            } else {
                v2 = (v2 * (^(v2)))
                v18 := GenS4{F0: v6}
                _ = v18
            }
            _ = fi15
            _ = v16
        }
        var v19 int64 = (v13 - (v13 | v13))
        v6 = true
        _ = fi14
        _ = v19
    }
    var v20 uint64 = v1
    v13 += v13
    v8 = 967
    var v21 uint8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
