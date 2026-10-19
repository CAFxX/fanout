package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 551
    v3 := make(map[uint32]int)
    var v4 int32 = -324
    v5 := []int{42}
    var v6 string = "foo bar"
    v4 = v4
    v6 = v6
    v4 += (v4 - (-(v4)))
    if false {
        v2 += (v2 ^ v2)
        _ = len(v5)
    }
    for fi7 := 0; fi7 < 15; fi7++ {
        v4 += (-(v4))
        if false {
            for fi8 := 0; fi8 < 3; fi8++ {
                if true {
                    var v9 int16 = -438
                    var v10 int32 = v4
                    _ = v10
                    _ = v9
                }
                if false {
                    var v11 string = ("" + v6)
                    v11 = v6
                    v12 := GenS1{F0: 65535, F1: 0, F2: 1}
                    v4 = (v4 * v4)
                    _ = v11
                    _ = v12
                } else {
                    v6 = (((v6 + v6) + v6) + v6)
                    if (2 < 2) {
                        fmt.Println(v2, v6)
                    } else {
                        var v13 int16 = -32768
                        _ = v13
                    }
                }
                _ = fi8
            }
            v6 = "hello"
        }
        _ = fi7
    }
    v6 = v6
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        switch "foo bar" {
        case "hello":
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        default:
            var v15 uint = 255
            _ = v15
        }
        v4 -= 487
        var v16 int32 = (v4 << ((65535) & 63))
        _ = fi14
        _ = v16
    }
    switch v6 {
    case "":
        v2 += (v2 | v2)
    case "hello":
        _ = len(v1)
    }
    var v17 uint8 = 255
    v18 := GenS1{F0: 729, F1: 42, F2: -459}
    v17 = v17
    v6 = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
