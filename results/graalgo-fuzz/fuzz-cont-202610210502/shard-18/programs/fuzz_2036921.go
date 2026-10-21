package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 uint8 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 255
    v3 := []string{"hello"}
    var v4 uint32 = 196
    var v5 uint16 = 712
    v6 := make(map[uint32]uint32)
    for fi7 := 0; fi7 < 6; fi7++ {
        for fi8 := 0; fi8 < 5; fi8++ {
            fmt.Println(fi8, fi8, v4)
            v9 := GenS1{F0: (v2 + (^(v2)))}
            _ = fi8
            _ = v9
        }
        v10 := GenS2{F0: 42, F1: (v2 % (((v2 >> ((161) & 63))) | 1))}
        _ = fi7
        _ = v10
    }
    v4 = 0
    v5 = (v5 >> ((2) & 63))
    var v11 int16 = -688
    var v12 uint64 = 0
    v1[(v4 & 2)] = 0
    v2 = (v2 | v2)
    if false {
        var v13 uint32 = v4
        _ = v13
    } else {
        for fi14 := 0; fi14 < 9; fi14++ {
            v15 := GenS1{F0: (^(v2))}
            if (!true) {
                var v16 float32 = 1.5
                v16 = v16
                var v17 uint16 = (^(v5))
                _ = v16
                _ = v17
            }
            _ = fi14
            _ = v15
        }
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        v4 = v4
        v1[(v4 + (v4 * v4))] = int(v12)
        for fi19 := 0; fi19 < 12; fi19++ {
            v11 -= (v11 / ((int16(v12)) | 1))
            _ = fi19
        }
        _ = fi18
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        var v21 int = (fi20 >> ((0) & 63))
        switch ("foo bar" + "\n\t") {
        case "foo bar":
            v5 += (v5 - 0)
        case "hello":
            var v22 float64 = ((-2.25 / (3.14159 / -2.25)) - 1.5)
            _ = v22
        case "a":
            v23 := GenS1{F0: v2}
            _ = v23
        default:
            switch ("\n\t" + ("\n\t" + "a")) {
            case "a":
                fmt.Println(v11, v11, fi20)
            }
        }
        var v24 uint = 4294967295
        _ = fi20
        _ = v21
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
