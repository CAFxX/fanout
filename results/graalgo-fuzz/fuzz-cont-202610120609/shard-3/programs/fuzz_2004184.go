package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 uint32 = 23
    var v2 uint64 = 415
    var v3 float64 = 1e10
    var v4 uint32 = v1
    var v5 uint = 2
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := []int{(fi6 / ((fi6) | 1))}
        switch (fi6 & fi6) {
        case 2147483647:
            switch "foo bar" {
            case "hello":
                fmt.Println(v4)
            case "\n\t":
                fmt.Println(v4, v1)
            }
        case 4294967295:
            for fi8 := 0; fi8 < 4; fi8++ {
                v9 := GenS1{F0: ((v2 - 893) == (v2 & 732))}
                var v10 uint32 = 2147483647
                v7 = append(v7, fi8)
                _ = fi8
                _ = v10
                _ = v9
            }
        case -1:
            fmt.Println(fi6)
        }
        var v11 uint64 = v2
        _ = fi6
        _ = v11
        _ = v7
    }
    v4 -= (v1 & (v4 & 2))
    fmt.Println(v3, v1, v2)
    v5 -= v5
    for fi12 := 0; fi12 < 9; fi12++ {
        v5 = 65535
        var v13 uint32 = (v1 >> (((^(v5))) & 63))
        _ = fi12
        _ = v13
    }
    var v14 int8 = -43
    v15 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
