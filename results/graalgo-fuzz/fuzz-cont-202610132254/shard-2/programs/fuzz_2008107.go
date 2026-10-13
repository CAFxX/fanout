package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 uint; F2 int32 }

func main() {
    var v1 int16 = 852
    v2 := make(map[int]bool)
    v3 := make(map[int]int)
    for fi4 := 0; fi4 < 11; fi4++ {
        v1 = -32768
        _ = fi4
    }
    fmt.Println(v1)
    v5 := GenS1{F0: (v1 * (v1 >> ((42) & 63)))}
    switch "a" {
    case "a":
        switch "\n\t" {
        case "foo bar":
            if true {
                v1 = (v1 | v1)
            } else {
                if (0 == 923) {
                    var v6 float32 = 1.5
                    for fi7 := 0; fi7 < 5; fi7++ {
                        fmt.Println(fi7)
                        _ = fi7
                    }
                    _ = v6
                }
                v8 := []int{}
                _ = v8
            }
        case "\n\t":
            v1 -= v1
        case "hello":
            v1 = (v1 + v1)
        }
    case "":
        if (859 == 1) {
            var v9 int16 = v1
            _ = v9
        }
    default:
        if false {
            v10 := []int{}
            var v11 uint16 = 255
            fmt.Println(v1)
            _ = v10
            _ = v11
        }
    }
    fmt.Println(v1)
    v12 := GenS2{F0: "\n\t", F1: 300, F2: -367}
    for fi13 := 0; fi13 < 13; fi13++ {
        for fi14 := 0; fi14 < 16; fi14++ {
            if (!((42 > 690) || (100 == -1))) {
                fmt.Println(v1, fi14, v1)
                v15 := GenS1{F0: (-(v1))}
                _ = v15
            } else {
                var v16 uint16 = 2
                switch (fi14 << ((1) & 63)) {
                case 0:
                    fmt.Println(fi13, fi13)
                default:
                    delete(v3, fi13)
                }
                _ = v16
            }
            switch (fi13 / ((fi14) | 1)) {
            case 366:
                v1 += (v1 % (((v1 * (v1 << ((4294967295) & 63)))) | 1))
            case -1:
                v1 += (v1 % ((v1) | 1))
            default:
                v3[(^(fi13))] = fi13
            }
            var v17 uint32 = 2147483647
            _ = fi14
            _ = v17
        }
        _ = fi13
    }
    for fi18 := 0; fi18 < 19; fi18++ {
        v19 := []int{}
        _ = fi18
        _ = v19
    }
    v1 -= v1
    v20 := GenS1{F0: v1}
    var v21 int64 = -9223372036854775808
    v21 += v21
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v5
}
