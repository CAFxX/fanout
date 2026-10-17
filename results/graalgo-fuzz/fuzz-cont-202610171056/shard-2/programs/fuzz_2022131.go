package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 uint64 }

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 100
    v3 := []uint32{1, 830}
    var v4 uint16 = 0
    for fi5 := 0; fi5 < 1; fi5++ {
        var v6 int = (fi5 * -522)
        v1["hello"] = 9223372036854775807
        delete(v1, ("\n\t" + ""))
        _ = fi5
        _ = v6
    }
    var v7 int32 = 2147483647
    _ = len(v1)
    v8 := v1["\n\t"]
    v2 -= v2
    switch 128 {
    case 147:
        for fi9 := 0; fi9 < 17; fi9++ {
            v10 := GenS1{F0: 1}
            _ = fi9
            _ = v10
        }
    default:
        fmt.Println(v8, v4)
    }
    delete(v1, "a")
    fmt.Println(v2, v7, v4)
    v2 = 1
    switch ("a" + ("foo bar" + "hello")) {
    case "hello":
        var v11 string = "\n\t"
        _ = v11
    case "":
        for fi12 := 0; fi12 < 10; fi12++ {
            if ((!(1e10 <= 0.1)) || ("" == "hello")) {
                switch fi12 {
                case 42:
                    v13 := GenS2{F0: 153, F1: 176}
                    _ = v13
                }
                v14 := GenS2{F0: v4, F1: 914}
                _ = v14
            }
            if false {
                for fi15 := 0; fi15 < 4; fi15++ {
                    var v16 int64 = 998
                    _ = fi15
                    _ = v16
                }
            }
            _ = fi12
        }
    case "foo bar":
        _ = len(v1)
    }
    v7 = (^(v7))
    if len(v3) > 0 {
        v3[0] = 2
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
