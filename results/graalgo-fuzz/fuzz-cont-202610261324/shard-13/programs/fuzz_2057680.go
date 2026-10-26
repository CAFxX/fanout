package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16; F1 uint; F2 string }
type GenS5 struct { F0 uint8; F1 uint8; F2 bool }

func main() {
    var v1 uint16 = 32767
    v2 := make(map[uint32]int)
    var v3 int8 = 9
    fmt.Println(v1)
    for fi4 := 0; fi4 < 17; fi4++ {
        v5 := GenS1{F0: 2, F1: 1, F2: 2}
        v1 = 32767
        _ = fi4
        _ = v5
    }
    v3 = (v3 ^ 127)
    v6 := v2[1]
    v3 -= v3
    fmt.Println(v1, v1, v3)
    v7 := GenS2{F0: 2}
    fmt.Println(v6)
    if (false && true) {
        var v8 uint64 = 4294967295
        _ = v8
    }
    v9 := GenS3{F0: v6}
    if (true || (65535 > 2147483648)) {
        if true {
            v10 := GenS4{F0: 0, F1: 1, F2: "hello"}
            fmt.Println(v1)
            var v11 int16 = 42
            _ = v10
            _ = v11
        }
    } else {
        v12 := GenS5{F0: 255, F1: 42, F2: false}
        switch "" {
        case "foo bar":
            v3 = -128
        }
        _ = v12
    }
    var v13 uint8 = 250
    v6 -= v6
    v14 := GenS4{F0: 1, F1: 18446744073709551615, F2: (("" + "hello") + (("\n\t" + ("foo bar" + "a")) + "foo bar"))}
    switch "a" {
    case "hello":
        v6 = 9223372036854775807
    case "\n\t":
        for fi15 := 0; fi15 < 18; fi15++ {
            v16 := []int{0, 128, (v6 | 32767)}
            v6 += (v6 + 42)
            _ = fi15
            _ = v16
        }
    case "a":
        switch "" {
        case "foo bar":
            var v17 uint64 = 722
            _ = v17
        case "hello":
            v6 += v6
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v9
}
