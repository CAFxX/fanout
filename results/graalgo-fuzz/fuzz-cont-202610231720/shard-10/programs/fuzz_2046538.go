package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int; F2 int32 }

func main() {
    var v1 int16 = 778
    var v2 uint64 = 0
    var v3 int32 = -806
    var v4 bool = ((v3 - -1) >= v3)
    v5 := GenS1{F0: 128, F1: 758, F2: v2}
    switch -9223372036854775808 {
    case 318:
        var v6 bool = v4
        _ = v6
    case -1:
        var v7 uint = 4294967295
        _ = v7
    case -832:
        for fi8 := 0; fi8 < 1; fi8++ {
            if (!v4) {
                v4 = (v4 || v4)
            } else {
                fmt.Println(fi8, fi8, v1)
                switch (fi8 | -209) {
                case 1:
                    v9 := []int{42, 9223372036854775807, (fi8 - (^(fi8))), 431}
                    _ = v9
                }
            }
            fmt.Println(fi8, fi8, fi8)
            _ = fi8
        }
    }
    var v10 int16 = v1
    fmt.Println(v3, v3, v4)
    var v11 int8 = 42
    for fi12 := 0; fi12 < 15; fi12++ {
        fmt.Println(v11, fi12)
        v3 += (v3 / (((v3 * (^(v3)))) | 1))
        _ = fi12
    }
    v11 = -125
    for fi13 := 0; fi13 < 11; fi13++ {
        if (128 > 1) {
            if (v4 && true) {
                v14 := []int{}
                switch ("foo bar" + "") {
                case "a":
                    fmt.Println(v11, v3, v4)
                case "":
                    v1 = (v1 >> ((384) & 63))
                }
                v15 := GenS1{F0: (fi13 << ((388) & 63)), F1: 255, F2: 0}
                _ = v14
                _ = v15
            }
        }
        switch fi13 {
        case -1:
            for fi16 := 0; fi16 < 12; fi16++ {
                v17 := GenS2{F0: (v2 / ((v2) | 1))}
                v10 += v10
                _ = fi16
                _ = v17
            }
        }
        v18 := GenS3{F0: false, F1: 2147483647, F2: v3}
        _ = fi13
        _ = v18
    }
    v19 := GenS1{F0: 9223372036854775807, F1: 0, F2: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
