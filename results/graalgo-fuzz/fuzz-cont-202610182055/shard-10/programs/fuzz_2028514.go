package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int64 }
type GenS2 struct { F0 uint64; F1 int8 }

func main() {
    v1 := []string{}
    v2 := []string{"a", "foo bar", "a"}
    var v3 int32 = 0
    if ((true && ((false || (2 >= 78)) && (-1 != 1))) || true) {
        v3 = v3
    } else {
        for fi4 := 0; fi4 < 18; fi4++ {
            for fi5 := 0; fi5 < 3; fi5++ {
                fmt.Println(fi5)
                _ = fi5
            }
            var v6 uint32 = 460
            switch "hello" {
            case "\n\t":
                v6 = 42
            }
            _ = fi4
            _ = v6
        }
    }
    var v7 float32 = 1e10
    for fi8 := 0; fi8 < 6; fi8++ {
        v7 += (v7 * v7)
        v7 += 0.1
        _ = fi8
    }
    if (!true) {
        for fi9 := 0; fi9 < 5; fi9++ {
            v3 += (^(v3))
            switch (("a" + "\n\t") + "a") {
            case "foo bar":
                if len(v1) > 0 {
                    v1[0] = "foo bar"
                }
            default:
                v3 += v3
            }
            _ = fi9
        }
        v3 = -2147483648
        var v10 float64 = -2.25
        _ = v10
    }
    v1 = append(v1, ("hello" + ("foo bar" + "a")))
    v3 += v3
    fmt.Println(v7)
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    var v11 int = 628
    v12 := GenS1{F0: "hello", F1: 205, F2: 1}
    if (0 > 255) {
        fmt.Println(v7, v11)
    } else {
        if true {
            v3 += v3
            v13 := GenS2{F0: 8, F1: -80}
            _ = v13
        }
        var v14 uint = 725
        for fi15 := 0; fi15 < 17; fi15++ {
            _ = len(v2)
            v16 := GenS1{F0: "\n\t", F1: v14, F2: 436}
            _ = fi15
            _ = v16
        }
        _ = v14
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v7
}
