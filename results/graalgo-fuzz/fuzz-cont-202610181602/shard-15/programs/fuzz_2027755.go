package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 string = "\n\t"
    v2 := []uint32{}
    v3 := []string{v1}
    fmt.Println(v1)
    v1 = ((v1 + v1) + (((v1 + "") + v1) + v1))
    v1 = v1
    fmt.Println(v1)
    v1 = v1
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    fmt.Println(v1)
    if true {
        var v4 int64 = 287
        for fi5 := 0; fi5 < 16; fi5++ {
            var v6 int = (fi5 ^ (fi5 % ((fi5) | 1)))
            _ = fi5
            _ = v6
        }
        _ = v4
    } else {
        v3 = append(v3, v1)
        for fi7 := 0; fi7 < 15; fi7++ {
            v8 := GenS1{F0: 42, F1: (v1 + v1), F2: 65535}
            if ((fi7 * fi7) < (fi7 | (-(fi7)))) {
                v9 := GenS1{F0: -1, F1: v1, F2: 65535}
                _ = v9
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                fmt.Println(v1, fi7)
                for fi10 := 0; fi10 < 18; fi10++ {
                    v11 := GenS1{F0: -740, F1: "a", F2: 787}
                    _ = fi10
                    _ = v11
                }
            }
            _ = fi7
            _ = v8
        }
    }
    v1 = v1
    for fi12 := 0; fi12 < 1; fi12++ {
        if (2 != 65535) {
            for fi13 := 0; fi13 < 7; fi13++ {
                v1 = v1
                switch (fi12 << ((18446744073709551615) & 63)) {
                case 9223372036854775807:
                    var v14 int8 = -29
                    _ = v14
                case 496:
                    v15 := GenS1{F0: 0, F1: v1, F2: 18446744073709551615}
                    _ = v15
                case 4294967295:
                    v1 = v1
                }
                switch "\n\t" {
                case "":
                    _ = len(v2)
                case "a":
                    v16 := GenS2{F0: 0}
                    _ = v16
                }
                _ = fi13
            }
            fmt.Println(v1, v1)
        }
        _ = fi12
    }
    fmt.Println(v1)
    v1 = v1
    _ = len(v2)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v17 int64 = 677
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
}
