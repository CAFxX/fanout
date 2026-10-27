package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 string }
type GenS2 struct { F0 string; F1 int64; F2 uint }
type GenS3 struct { F0 uint32; F1 uint16; F2 uint }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 0
    var v3 bool = true
    var v4 string = (("\n\t" + "hello") + "hello")
    var v5 uint64 = 1
    v6 := make(map[int]int)
    if v3 {
        v7 := []int{32767}
        _ = v7
    }
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 uint16 = 2
        _ = fi8
        _ = v9
    }
    switch "a" {
    case "\n\t":
        v10 := GenS1{F0: 18446744073709551615, F1: int8(v5), F2: (v4 + "a")}
        _ = v10
    case "":
        switch "" {
        case "\n\t":
            var v11 uint16 = 442
            _ = v11
        }
    case "hello":
        for fi12 := 0; fi12 < 9; fi12++ {
            if v3 {
                fmt.Println(v4, v4)
            }
            for fi13 := 0; fi13 < 11; fi13++ {
                v5 = v5
                if v3 {
                    v3 = (-1 == 2147483647)
                } else {
                    fmt.Println(v5, fi13)
                    v14 := GenS1{F0: (v5 + (v5 << ((18446744073709551615) & 63))), F1: -1, F2: "foo bar"}
                    _ = v14
                }
                _ = fi13
            }
            _ = fi12
        }
    }
    v2 += (v2 << ((42) & 63))
    var v15 uint32 = 2147483647
    var v16 uint64 = 4294967295
    v17 := GenS2{F0: "\n\t", F1: 42, F2: 65535}
    v18 := GenS3{F0: 0, F1: 999, F2: 0}
    var v19 int8 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
