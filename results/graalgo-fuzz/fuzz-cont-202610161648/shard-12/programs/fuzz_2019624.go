package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 int32 }
type GenS2 struct { F0 uint16; F1 string; F2 int32 }
type GenS3 struct { F0 uint8; F1 int64 }

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{272, 0, 65535, 728}
    var v3 int8 = -109
    var v4 float32 = (0.1 + 0.1)
    var v5 bool = (v3 >= (^(v3)))
    if ((true && v5) || (!false)) {
        switch "foo bar" {
        case "":
            fmt.Println(v4, v5)
        case "a":
            var v6 int32 = -2147483648
            _ = v6
        case "\n\t":
            var v7 int16 = -32768
            _ = v7
        default:
            fmt.Println(v3)
        }
        switch 9223372036854775807 {
        case -1:
            delete(v1, 515)
        case 570:
            v8 := GenS1{F0: (v3 << ((65535) & 63)), F1: true, F2: -193}
            _ = v8
        case -9223372036854775808:
            v3 = (-(v3))
        }
        if (9223372036854775807 <= -9223372036854775808) {
            v3 -= v3
        } else {
            var v9 int16 = 936
            _ = v9
        }
    } else {
        fmt.Println(v3, v5, v4)
        if (!v5) {
            if true {
                var v10 uint = 912
                v11 := GenS1{F0: 1, F1: v5, F2: 0}
                v3 += (-(v3))
                _ = v10
                _ = v11
            }
            for fi12 := 0; fi12 < 12; fi12++ {
                v13 := GenS1{F0: v3, F1: v5, F2: 0}
                _ = fi12
                _ = v13
            }
            for fi14 := 0; fi14 < 20; fi14++ {
                v4 = v4
                var v15 uint8 = 0
                v16 := GenS2{F0: 196, F1: ("a" + ""), F2: 883}
                _ = fi14
                _ = v15
                _ = v16
            }
        }
        v5 = true
    }
    switch 697 {
    case -742:
        v17 := GenS3{F0: 127, F1: 668}
        _ = v17
    }
    var v18 int32 = 527
    fmt.Println(v3, v3)
    var v19 float32 = 1.5
    v20 := GenS1{F0: v3, F1: v5, F2: (v18 >> ((375) & 63))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
