package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int16 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]int)
    var v3 uint8 = 2
    var v4 int16 = -887
    var v5 int64 = 0
    v6 := make(map[string]bool)
    var v7 uint64 = 1
    v7 += 4294967295
    for fi8 := 0; fi8 < 18; fi8++ {
        v4 -= (v4 * v4)
        if (("" + ("hello" + "a")) == "") {
            v9 := GenS1{F0: 0, F1: (v7 / ((42) | 1)), F2: 0}
            _ = v9
        } else {
            v4 = 0
            fmt.Println(v3)
        }
        for fi10 := 0; fi10 < 6; fi10++ {
            v3 -= (v3 | 42)
            v11 := []int{-602, 4294967295, (-(fi10)), (fi10 / (((fi10 / (((fi8 ^ int(v3))) | 1))) | 1))}
            v12 := GenS1{F0: -479, F1: 42, F2: v4}
            switch fi10 {
            case 42:
                v7 += v7
            }
            switch "\n\t" {
            case "hello":
                if true {
                    delete(v1, 2147483647)
                }
            case "\n\t":
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi8
    }
    if ((float64((1e10 - float32(0.1))) + -2.25) == float64(((0.1 / 1.0) / (float32((-2.25 / -2.25)) - 1.5)))) {
        v13 := []int{32768}
        var v14 uint16 = 631
        _ = v13
        _ = v14
    }
    fmt.Println(v7)
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v5)
    default:
        for fi15 := 0; fi15 < 13; fi15++ {
            switch (^(fi15)) {
            case -75:
                v5 = v5
            case -839:
                v16 := GenS1{F0: 2147483647, F1: 0, F2: (v4 >> ((2) & 63))}
                _ = v16
            }
            switch "a" {
            case "a":
                v7 -= v7
            }
            v17 := GenS1{F0: -1, F1: 171, F2: 42}
            _ = fi15
            _ = v17
        }
    }
    v5 -= (v5 % ((42) | 1))
    v4 = v4
    var v18 float32 = (((3.14159 / 0.1) / ((-2.25 * 0.1) - (1.5 / -2.25))) - 3.14159)
    var v19 float32 = 0.1
    for fi20 := 0; fi20 < 16; fi20++ {
        v19 = ((1.5 - float32(0.1)) * v18)
        _ = fi20
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        v3 -= 42
        _ = fi21
    }
    var v22 int64 = (v5 << ((1) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
