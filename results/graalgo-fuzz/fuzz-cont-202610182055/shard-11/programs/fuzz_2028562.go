package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 bool }

func main() {
    var v1 int8 = -95
    var v2 int64 = 42
    v3 := []string{"hello"}
    v4 := make(map[int]bool)
    var v5 bool = false
    var v6 float64 = float64(3.14159)
    v7 := GenS1{F0: 1}
    v1 -= int8(v2)
    for fi8 := 0; fi8 < 18; fi8++ {
        v2 = (v2 | 526)
        for fi9 := 0; fi9 < 15; fi9++ {
            v5 = v5
            v10 := GenS2{F0: 32767, F1: 1}
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    v2 = v2
    fmt.Println(v6)
    var v11 bool = (1 >= -423)
    if v11 {
        v11 = ((379 == 855) && ((0.0 > 0.1) && v5))
    }
    v12 := GenS2{F0: 252, F1: -1}
    switch 127 {
    case -1:
        var v13 uint64 = 952
        _ = v13
    case 32767:
        v14 := GenS1{F0: (v2 * v2)}
        _ = v14
    case -2147483648:
        var v15 uint16 = 355
        _ = v15
    default:
        switch 2147483647 {
        case -903:
            v16 := GenS3{F0: (!false)}
            _ = v16
        }
    }
    switch ((("foo bar" + "hello") + ((("\n\t" + "hello") + ("hello" + ("" + (("foo bar" + "a") + "")))) + "a")) + "a") {
    case "\n\t":
        switch "\n\t" {
        case "foo bar":
            var v17 int32 = 2147483647
            _ = v17
        case "":
            v2 = v2
        case "hello":
            switch -130 {
            case 963:
                var v18 uint32 = 2147483648
                _ = v18
            case -52:
                var v19 uint32 = 0
                _ = v19
            case 4294967295:
                if v11 {
                    var v20 uint8 = 1
                    if len(v3) > 0 {
                        v3[0] = "foo bar"
                    }
                    _ = v20
                } else {
                    v1 += v1
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            }
        }
    case "foo bar":
        fmt.Println(v6)
    case "":
        var v21 int8 = v1
        _ = v21
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
