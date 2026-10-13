package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 string; F2 uint8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 127
    var v3 int8 = (v2 >> ((385) & 63))
    var v4 float64 = -2.25
    switch 276 {
    case 9223372036854775807:
        var v5 string = (("a" + "") + "foo bar")
        _ = v5
    case -9223372036854775808:
        if true {
            v3 = 42
            delete(v1, 27)
            v6 := []int{738, -1, -722}
            _ = v6
        }
    case 2147483647:
        v2 = (v3 - v3)
    }
    for fi7 := 0; fi7 < 15; fi7++ {
        v8 := GenS1{F0: 32767}
        var v9 int32 = 914
        switch "hello" {
        case "\n\t":
            v10 := GenS1{F0: 32767}
            _ = v10
        case "foo bar":
            v2 += (^(v2))
        case "a":
            v9 = 1
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    var v11 uint = 303
    v12 := []int{}
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    v13 := GenS2{F0: false, F1: "\n\t", F2: 1}
    var v14 uint16 = 0
    for fi15 := 0; fi15 < 12; fi15++ {
        v4 = v4
        _ = fi15
    }
    v11 -= 255
    if true {
        _ = len(v1)
    } else {
        v16 := GenS2{F0: false, F1: "foo bar", F2: 127}
        v2 += (^(v3))
        _ = v16
    }
    v17 := GenS1{F0: uint16(v11)}
    v11 = (^(v11))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
