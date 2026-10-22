package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int }
type GenS2 struct { F0 int64 }

func main() {
    v1 := []string{"\n\t", ("\n\t" + ("foo bar" + "\n\t")), "\n\t", ""}
    var v2 bool = false
    v3 := []string{}
    for fi4 := 0; fi4 < 3; fi4++ {
        v2 = v2
        for fi5 := 0; fi5 < 16; fi5++ {
            var v6 int64 = 344
            switch 2147483647 {
            case 1:
                v2 = v2
            case 2147483647:
                var v7 int = -751
                _ = v7
            }
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    v2 = (127 >= 255)
    for fi8 := 0; fi8 < 5; fi8++ {
        fmt.Println(v2)
        v2 = (!v2)
        var v9 int32 = 2147483647
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: v2, F1: 9223372036854775807, F2: 32768}
    v11 := GenS2{F0: 42}
    v12 := GenS1{F0: v2, F1: -1, F2: 700}
    var v13 uint = 65535
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 uint64 = 0
        switch (fi14 / ((fi14) | 1)) {
        case -2147483648:
            var v16 uint = v13
            _ = v16
        }
        v13 = (v13 ^ v13)
        _ = fi14
        _ = v15
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
