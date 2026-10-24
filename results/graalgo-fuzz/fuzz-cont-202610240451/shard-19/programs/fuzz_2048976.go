package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint32 }

func main() {
    v1 := []string{"\n\t", "hello", ("a" + "hello")}
    var v2 uint16 = 733
    v3 := []string{"foo bar", ("hello" + "\n\t"), ("foo bar" + "")}
    v4 := make(map[int]int)
    fmt.Println(v2)
    for fi5 := 0; fi5 < 15; fi5++ {
        if ("foo bar" == "\n\t") {
            for fi6 := 0; fi6 < 11; fi6++ {
                if (1.5 < 0.1) {
                    var v7 int64 = -553
                    var v8 uint64 = 141
                    _ = v7
                    _ = v8
                }
                v9 := GenS1{F0: 0, F1: 1, F2: 4294967295}
                _ = fi6
                _ = v9
            }
        }
        _ = fi5
    }
    v10 := GenS1{F0: 128, F1: 4294967295, F2: 2147483647}
    v2 = v2
    fmt.Println(v2)
    switch 32768 {
    case 234:
        for fi11 := 0; fi11 < 7; fi11++ {
            v3 = append(v3, "\n\t")
            _ = fi11
        }
    default:
        switch 32768 {
        case -187:
            v3 = append(v3, ("\n\t" + "hello"))
        case 127:
            v12 := v4[210]
            _ = v12
        }
    }
    v2 = 381
    v3 = append(v3, "")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
}
