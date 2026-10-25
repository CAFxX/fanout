package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 uint; F1 int16; F2 int }

func main() {
    v1 := []uint32{0, 65535, 890, 42}
    v2 := make(map[uint32]string)
    var v3 int16 = -260
    v4 := make(map[int]bool)
    if false {
        v3 = (v3 >> ((255) & 63))
    }
    v5 := GenS1{F0: 2147483648}
    fmt.Println(v3)
    v3 = (v3 + -539)
    if (1.5 == 0.1) {
        var v6 string = ("hello" + ((("a" + "a") + ("foo bar" + ("hello" + "\n\t"))) + ""))
        v7 := v2[0]
        _ = v6
        _ = v7
    } else {
        v3 = v3
    }
    v3 += 168
    if ("foo bar" > "\n\t") {
        for fi8 := 0; fi8 < 19; fi8++ {
            switch fi8 {
            case -64:
                for fi9 := 0; fi9 < 7; fi9++ {
                    v10 := GenS1{F0: 694}
                    _ = fi9
                    _ = v10
                }
            case 511:
                v3 = (v3 ^ (v3 + (v3 + v3)))
            }
            _ = fi8
        }
        v3 -= (v3 * v3)
    } else {
        if (898 != 1) {
            for fi11 := 0; fi11 < 11; fi11++ {
                v3 += 32767
                _ = fi11
            }
            fmt.Println(v3)
        } else {
            var v12 int16 = -32768
            _ = v12
        }
    }
    var v13 int = -9223372036854775808
    fmt.Println(v3)
    if true {
        v14 := GenS2{F0: 2, F1: 4294967295}
        _ = v14
    } else {
        var v15 string = "a"
        var v16 int16 = (v3 % ((v3) | 1))
        fmt.Println(v3)
        _ = v15
        _ = v16
    }
    v17 := GenS3{F0: 0, F1: v3, F2: 4294967295}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
