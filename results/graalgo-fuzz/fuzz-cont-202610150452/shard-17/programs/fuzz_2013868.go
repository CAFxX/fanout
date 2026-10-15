package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 int8 }

func main() {
    var v1 string = ("" + "foo bar")
    var v2 uint16 = 32767
    v3 := make(map[uint32]bool)
    var v4 uint8 = 142
    v4 = v4
    fmt.Println(v1)
    if false {
        for fi5 := 0; fi5 < 4; fi5++ {
            var v6 int16 = -667
            _ = fi5
            _ = v6
        }
        var v7 string = (v1 + v1)
        _ = v7
    } else {
        for fi8 := 0; fi8 < 8; fi8++ {
            for fi9 := 0; fi9 < 7; fi9++ {
                fmt.Println(v4, fi8)
                v10 := GenS1{F0: (v2 ^ 32767)}
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
    }
    v2 += v2
    v11 := GenS1{F0: 32767}
    var v12 uint8 = v4
    switch ((v1 + (("hello" + v1) + v1)) + ((v1 + ((v1 + "") + "foo bar")) + "foo bar")) {
    case "foo bar":
        if false {
            v12 = 127
        }
    case "\n\t":
        v13 := []int{}
        _ = v13
    default:
        v14 := []int{674, -9223372036854775808}
        _ = v14
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        if false {
            fmt.Println(fi15)
            v16 := GenS2{F0: "hello", F1: -22}
            _ = v16
        } else {
            v17 := GenS1{F0: 255}
            _ = v17
        }
        var v18 int16 = 42
        switch 32767 {
        case -9223372036854775808:
            v19 := []int{fi15, fi15, fi15, 128}
            _ = v19
        case 32768:
            v12 = (v12 * 1)
        default:
            for fi20 := 0; fi20 < 20; fi20++ {
                v12 -= v12
                _ = fi20
            }
        }
        _ = fi15
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
