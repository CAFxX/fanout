package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 string; F2 uint16 }
type GenS3 struct { F0 bool; F1 uint16; F2 uint8 }
type GenS4 struct { F0 uint16; F1 bool }
func genFn0(p0 uint16) uint16 {
    switch "" {
    case "a":
        v5 := []int{127, -414}
        _ = v5
    }
    var v6 int = -750
    v7 := GenS1{F0: 1}
    switch v6 {
    case -427:
        switch v6 {
        case -848:
            for fi8 := 0; fi8 < 13; fi8++ {
                v9 := GenS1{F0: 1}
                fmt.Println(fi8, v6)
                for fi10 := 0; fi10 < 16; fi10++ {
                    v6 += (fi8 << ((988) & 63))
                    fmt.Println(v6)
                    _ = fi10
                }
                _ = fi8
                _ = v9
            }
        }
    case -9223372036854775808:
        v11 := GenS1{F0: 0}
        _ = v11
    case -2147483648:
        var v12 uint = 562
        _ = v12
    default:
        v13 := GenS2{F0: 42, F1: ("a" + "a"), F2: p0}
        _ = v13
    }
    v14 := GenS1{F0: 39}
    _ = p0
    _ = v14
    _ = v6
    _ = v7
    return (p0 & p0)
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]int)
    v3 := make(map[int]bool)
    var v4 uint8 = 255
    var v15 int32 = 0
    var v16 int32 = 584
    var v17 uint16 = 0
    var v18 uint64 = 742
    var v19 float32 = 1.5
    switch "a" {
    case "\n\t":
        v20 := GenS2{F0: (v18 << ((603) & 63)), F1: "hello", F2: (v17 % ((712) | 1))}
        _ = v20
    }
    if true {
        v21 := GenS2{F0: 0, F1: ("hello" + "\n\t"), F2: (v17 / ((v17) | 1))}
        _ = v21
    } else {
        var v22 bool = true
        var v23 int32 = (v15 - (v16 * v15))
        _ = v22
        _ = v23
    }
    fmt.Println(v15, v4)
    if true {
        switch -9223372036854775808 {
        case 681:
            var v24 uint16 = 87
            _ = v24
        case -404:
            for fi25 := 0; fi25 < 12; fi25++ {
                var v26 int = 4294967295
                for fi27 := 0; fi27 < 9; fi27++ {
                    fmt.Println(v26, v26, v15)
                    if true {
                        switch (-(fi27)) {
                        case 32767:
                            v28 := GenS3{F0: (!(910 == 544)), F1: 42, F2: v4}
                            _ = v28
                        case -65:
                            _ = len(v3)
                        default:
                            var v29 string = "foo bar"
                            _ = v29
                        }
                    } else {
                        if true {
                            v26 -= (fi27 + fi25)
                        } else {
                            var v30 float64 = (0.0 + 0.0)
                            v31 := GenS4{F0: v17, F1: ((42 <= 297) || false)}
                            fmt.Println(fi27, v15)
                            _ = v30
                            _ = v31
                        }
                    }
                    fmt.Println(v15, fi25, v15)
                    _ = fi27
                }
                _ = fi25
                _ = v26
            }
        }
        for fi32 := 0; fi32 < 11; fi32++ {
            v33 := []int{32768, 1, fi32, (fi32 % ((fi32) | 1))}
            fmt.Println(v17)
            _ = fi32
            _ = v33
        }
        delete(v2, 2147483648)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
