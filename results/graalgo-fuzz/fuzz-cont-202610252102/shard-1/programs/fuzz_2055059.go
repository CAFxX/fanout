package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 int16 = 998
    var v2 int64 = -9223372036854775808
    v3 := make(map[int]string)
    var v4 int32 = 42
    v5 := make(map[uint32]string)
    var v6 uint8 = 127
    var v7 int32 = (v4 ^ 1)
    v5[65535] = "\n\t"
    for fi8 := 0; fi8 < 12; fi8++ {
        v1 -= 0
        switch (fi8 << ((4294967295) & 63)) {
        case 2147483647:
            v4 -= v4
        case -786:
            v9 := []int{fi8, 599}
            _ = v9
        case -9223372036854775808:
            v10 := []int{9223372036854775807, (fi8 & (fi8 % ((42) | 1))), fi8}
            _ = v10
        default:
            fmt.Println(v7, v6, v1)
        }
        v11 := GenS1{F0: (v6 >> ((255) & 63)), F1: (v2 ^ v2)}
        _ = fi8
        _ = v11
    }
    if false {
        for fi12 := 0; fi12 < 10; fi12++ {
            switch (fi12 | (fi12 % ((fi12) | 1))) {
            case 9223372036854775807:
                if false {
                    delete(v3, int(v2))
                    v13 := []int{fi12, fi12, (fi12 << ((1) & 63))}
                    for fi14 := 0; fi14 < 14; fi14++ {
                        if len(v13) > 0 {
                            v13[0] = (fi14 % ((fi14) | 1))
                        }
                        delete(v5, 165)
                        _ = len(v13)
                        _ = fi14
                    }
                    _ = v13
                } else {
                    var v15 uint = 65535
                    _ = v15
                }
            default:
                v4 = (v4 | (v4 >> ((255) & 63)))
            }
            var v16 uint16 = 42
            _ = fi12
            _ = v16
        }
        v1 += -32768
        fmt.Println(v6)
    }
    var v17 uint64 = 545
    for fi18 := 0; fi18 < 2; fi18++ {
        var v19 uint8 = v6
        switch "hello" {
        case "":
            var v20 uint = 4294967295
            _ = v20
        case "\n\t":
            var v21 bool = ((v7 > (^(v7))) && false)
            _ = v21
        case "a":
            v7 += (v4 & -496)
        }
        v22 := GenS2{F0: 506}
        _ = fi18
        _ = v19
        _ = v22
    }
    for fi23 := 0; fi23 < 9; fi23++ {
        var v24 uint8 = (v6 / ((0) | 1))
        _ = fi23
        _ = v24
    }
    if (!true) {
        v4 = v4
        switch "\n\t" {
        case "\n\t":
            if ((v2 / (((v2 >> ((700) & 63))) | 1)) > (-(v2))) {
                fmt.Println(v7)
            } else {
                switch 9223372036854775807 {
                case 32768:
                    var v25 int = -409
                    _ = v25
                case 978:
                    v26 := GenS1{F0: v6, F1: v2}
                    _ = v26
                case -2147483648:
                    var v27 uint = 1
                    _ = v27
                default:
                    var v28 int32 = (v7 + 1)
                    _ = v28
                }
            }
        case "a":
            if (!(-2147483648 >= -1)) {
                delete(v3, 102)
            }
        case "foo bar":
            var v29 int32 = (v7 * 2147483647)
            _ = v29
        }
    }
    v17 += (^(v17))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
