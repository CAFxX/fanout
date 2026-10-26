package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 string; F2 bool }

func main() {
    var v1 int = -217
    v2 := make(map[uint32]bool)
    var v3 uint = 217
    var v4 uint32 = 559
    v5 := make(map[int]string)
    var v6 uint8 = 185
    if (!(190 > -854)) {
        v7 := GenS1{F0: v1}
        _ = v7
    }
    v8 := []int{}
    v6 = v6
    if (1 <= (v6 - v6)) {
        for fi9 := 0; fi9 < 5; fi9++ {
            v2[(v4 + 97)] = true
            v10 := GenS1{F0: 890}
            _ = fi9
            _ = v10
        }
        v11 := v5[v1]
        if len(v8) > 0 {
            v8[0] = 662
        }
        _ = v11
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 uint64 = 112
        var v14 uint = v3
        _ = fi12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        var v16 int32 = -197
        v16 += (v16 & (^(v16)))
        v17 := GenS2{F0: 18446744073709551615}
        _ = fi15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        v19 := GenS2{F0: 0}
        var v20 uint8 = v6
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 uint = v3
    for fi22 := 0; fi22 < 5; fi22++ {
        var v23 uint32 = (v4 - (^(v4)))
        switch (("foo bar" + "hello") + "a") {
        case "\n\t":
            delete(v2, (v23 << ((255) & 63)))
        case "hello":
            for fi24 := 0; fi24 < 2; fi24++ {
                v25 := v2[v23]
                v3 = v21
                _ = fi24
                _ = v25
            }
        case "":
            v26 := GenS3{F0: v3, F1: ("\n\t" + ("\n\t" + "")), F2: (false && true)}
            _ = v26
        }
        fmt.Println(v23, v1, v6)
        _ = fi22
        _ = v23
    }
    v21 -= (v3 & v3)
    v27 := GenS3{F0: v3, F1: ("hello" + ("" + "\n\t")), F2: (v4 >= v4)}
    var v28 int16 = 1
    fmt.Println(v21)
    fmt.Println(v6, v1)
    v3 = v3
    v29 := GenS2{F0: 749}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
