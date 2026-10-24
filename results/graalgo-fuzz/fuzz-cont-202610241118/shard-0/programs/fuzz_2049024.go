package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]int)
    v3 := []string{"", "\n\t"}
    var v4 int8 = 127
    v4 = 0
    for fi5 := 0; fi5 < 19; fi5++ {
        fmt.Println(fi5, v4)
        _ = fi5
    }
    var v6 int = -9223372036854775808
    v6 = v6
    for fi7 := 0; fi7 < 20; fi7++ {
        if (-9223372036854775808 > 0) {
            for fi8 := 0; fi8 < 3; fi8++ {
                v9 := GenS1{F0: 0}
                fmt.Println(fi8, v4, fi7)
                _ = fi8
                _ = v9
            }
        }
        _ = fi7
    }
    fmt.Println(v6)
    v4 = (v4 % (((^(v4))) | 1))
    var v10 uint = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
