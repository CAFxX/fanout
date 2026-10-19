package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 string; F1 int64; F2 string }
type GenS3 struct { F0 bool; F1 int; F2 int16 }

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 522
    v3 := []string{""}
    v4 := make(map[int]bool)
    v5 := make(map[string]string)
    v6 := GenS1{F0: true, F1: 2147483647}
    var v7 uint8 = 128
    for fi8 := 0; fi8 < 4; fi8++ {
        if true {
            v2 += (^(v2))
        } else {
            v9 := v1["foo bar"]
            for fi10 := 0; fi10 < 20; fi10++ {
                v11 := GenS1{F0: false, F1: (v9 << ((108) & 63))}
                var v12 bool = (!true)
                fmt.Println(fi8, v9, v12)
                _ = fi10
                _ = v11
                _ = v12
            }
            var v13 string = "\n\t"
            _ = v13
            _ = v9
        }
        v2 = 18446744073709551615
        _ = fi8
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        for fi15 := 0; fi15 < 8; fi15++ {
            v16 := GenS2{F0: "a", F1: 42, F2: "hello"}
            _ = fi15
            _ = v16
        }
        v17 := GenS2{F0: "foo bar", F1: -600, F2: "\n\t"}
        switch 32767 {
        case 508:
            for fi18 := 0; fi18 < 14; fi18++ {
                v7 = v7
                v7 -= v7
                _ = fi18
            }
        }
        _ = fi14
        _ = v17
    }
    fmt.Println(v7, v2)
    v7 = uint8(v2)
    v19 := GenS3{F0: ((v2 | v2) <= (v2 * v2)), F1: 0, F2: 1000}
    fmt.Println(v7, v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
