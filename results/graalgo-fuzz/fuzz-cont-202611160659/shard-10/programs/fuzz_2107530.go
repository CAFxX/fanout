package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]bool)
    v3 := []int{744, 2147483647, 127, -765}
    if (-1 != -884) {
        switch 1 {
        case 482:
            var v4 uint = 2
            _ = v4
        case -2147483648:
            var v5 uint = 1
            _ = v5
        }
        var v6 uint = 606
        v6 -= (v6 / ((v6) | 1))
        _ = v6
    } else {
        var v7 int8 = -128
        switch 32767 {
        case 0:
            v7 = v7
        case 557:
            if (-1 <= -111) {
                var v8 uint16 = 0
                v7 += (v7 - (v7 >> ((255) & 63)))
                _ = v8
            } else {
                var v9 int64 = 0
                v9 = v9
                _ = v9
            }
        }
        _ = v7
    }
    v10 := GenS1{F0: 2}
    if true {
        var v11 uint8 = 42
        _ = v11
    }
    switch 32768 {
    case 1:
        var v12 uint = 18446744073709551615
        _ = v12
    }
    var v13 uint64 = 4294967295
    switch "foo bar" {
    case "":
        switch "" {
        case "foo bar":
            v14 := GenS1{F0: 1}
            _ = v14
        case "\n\t":
            switch ("\n\t" + "hello") {
            case "hello":
                fmt.Println(v13)
            }
        }
    case "foo bar":
        v13 = (v13 + (v13 >> ((385) & 63)))
    case "\n\t":
        switch "foo bar" {
        case "hello":
            var v15 uint64 = (v13 & v13)
            _ = v15
        default:
            var v16 int32 = 480
            _ = v16
        }
    }
    var v17 int8 = 1
    v13 -= v13
    for fi18 := 0; fi18 < 18; fi18++ {
        if (42 >= 579) {
            v13 = v13
        } else {
            for fi19 := 0; fi19 < 14; fi19++ {
                v17 -= v17
                _ = fi19
            }
        }
        if len(v3) > 0 {
            v3[0] = fi18
        }
        _ = fi18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v13
    _ = v17
    _ = v2
    _ = v3
}
