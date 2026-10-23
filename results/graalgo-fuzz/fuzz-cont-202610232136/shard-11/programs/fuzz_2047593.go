package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8; F1 int32 }

func main() {
    v1 := make(map[int]int)
    var v2 int16 = -32768
    v3 := []string{}
    var v4 string = "a"
    var v5 uint64 = 2
    var v6 uint32 = 23
    fmt.Println(v5, v6)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if true {
        var v7 int8 = -1
        _ = v7
    }
    for fi8 := 0; fi8 < 10; fi8++ {
        switch (fi8 + fi8) {
        case 998:
            v6 -= 2
        }
        if (!((3.14159 != ((0.0 - -2.25) / ((1e10 * 0.0) - 3.14159))) || (v6 <= (v6 % ((v6) | 1))))) {
            v4 = "foo bar"
        }
        for fi9 := 0; fi9 < 17; fi9++ {
            fmt.Println(v2)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi9
        }
        _ = fi8
    }
    if true {
        v10 := GenS1{F0: 0, F1: 967}
        fmt.Println(v6, v6)
        _ = v10
    } else {
        for fi11 := 0; fi11 < 3; fi11++ {
            _ = len(v3)
            var v12 string = "a"
            v3 = append(v3, "foo bar")
            _ = fi11
            _ = v12
        }
        delete(v1, 9223372036854775807)
    }
    fmt.Println(v6)
    _ = len(v3)
    _ = len(v3)
    if true {
        v1[921] = 110
        var v13 float32 = (-2.25 - 1.5)
        fmt.Println(v6, v6)
        _ = v13
    } else {
        v4 = ("" + v4)
        v14 := GenS2{F0: -593, F1: v6}
        _ = v14
    }
    v15 := GenS3{F0: -496}
    var v16 float64 = 1.5
    switch v4 {
    case "foo bar":
        var v17 uint32 = v6
        _ = v17
    case "\n\t":
        switch v4 {
        case "a":
            var v18 int64 = -9223372036854775808
            _ = v18
        case "hello":
            var v19 int16 = 1
            _ = v19
        }
    }
    v20 := GenS4{F0: 100, F1: -500}
    fmt.Println(v5)
    var v21 uint8 = 128
    fmt.Println(v16, v5)
    for fi22 := 0; fi22 < 2; fi22++ {
        switch "hello" {
        case "a":
            var v23 uint8 = (^(v21))
            _ = v23
        }
        _ = fi22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
