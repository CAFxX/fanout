package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 string; F2 int }
type GenS4 struct { F0 int32; F1 int32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 588
    var v3 uint16 = 132
    var v4 uint = 489
    v5 := []string{"", "foo bar", "a"}
    if (!((v2 * -52) < v2)) {
        var v6 int = 32768
        v7 := GenS1{F0: -164}
        _ = v6
        _ = v7
    }
    var v8 int = (v2 ^ (v2 & (v2 - -9223372036854775808)))
    if false {
        var v9 int64 = 4
        _ = v9
    }
    switch -904 {
    case 128:
        v4 += (v4 - v4)
    case 0:
        if true {
            v10 := GenS2{F0: "a"}
            v8 -= v2
            v5 = append(v5, ("hello" + ""))
            _ = v10
        } else {
            var v11 int8 = -25
            switch 128 {
            case 397:
                for fi12 := 0; fi12 < 6; fi12++ {
                    v2 -= v8
                    _ = fi12
                }
            default:
                var v13 int64 = 1
                _ = v13
            }
            _ = v11
        }
    case 4294967295:
        if true {
            var v14 uint64 = 811
            _ = v14
        } else {
            fmt.Println(v3, v8, v8)
        }
    default:
        var v15 int = (v8 / (((^(v2))) | 1))
        _ = v15
    }
    v1[1] = ((((("hello" + "hello") + "") + "a") + "foo bar") + "hello")
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        for fi17 := 0; fi17 < 9; fi17++ {
            fmt.Println(v3, v3)
            v18 := GenS3{F0: (18446744073709551615 != 4294967295), F1: "", F2: (fi16 + (^(v8)))}
            switch ("a" + "foo bar") {
            case "foo bar":
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case "\n\t":
                v19 := GenS1{F0: (fi16 >> (((^(v4))) & 63))}
                _ = v19
            case "a":
                v20 := GenS4{F0: 0, F1: 2147483647}
                _ = v20
            }
            _ = fi17
            _ = v18
        }
        var v21 uint64 = 523
        _ = fi16
        _ = v21
    }
    fmt.Println(v8, v4)
    v22 := GenS3{F0: false, F1: "hello", F2: v2}
    var v23 uint8 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
