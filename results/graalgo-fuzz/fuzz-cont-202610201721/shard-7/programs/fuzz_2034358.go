package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint32 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 int; F1 int64; F2 int }

func main() {
    var v1 float32 = 1.5
    v2 := []uint32{812}
    var v3 int64 = 42
    var v4 uint8 = 127
    if ((-1 > 0) && false) {
        switch 32767 {
        case -2147483648:
            v4 = (^(v4))
        case 4294967295:
            var v5 uint32 = 386
            _ = v5
        case 127:
            fmt.Println(v3, v3, v4)
        }
        fmt.Println(v3)
    } else {
        v1 = float32(0.1)
        var v6 string = (("foo bar" + ("a" + "\n\t")) + "")
        if false {
            var v7 uint8 = 255
            _ = v7
        }
        _ = v6
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        v3 -= v3
        fmt.Println(fi8, fi8)
        _ = fi8
    }
    v1 -= v1
    v3 += (^(v3))
    if (32767 <= 370) {
        var v9 int = -9223372036854775808
        v9 = -670
        var v10 uint32 = 2
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: -555, F1: ("hello" + ""), F2: 42}
    switch "" {
    case "":
        v12 := GenS2{F0: 2, F1: 9223372036854775807}
        _ = v12
    default:
        if true {
            switch (("a" + (("a" + "foo bar") + ("" + "a"))) + "foo bar") {
            case "hello":
                var v13 string = ""
                _ = v13
            case "foo bar":
                v14 := GenS2{F0: 525, F1: -1}
                _ = v14
            case "":
                v3 = -267
            }
            v15 := GenS3{F0: -2147483648, F1: (v3 + v3), F2: 32768}
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        var v17 int = (fi16 * -992)
        _ = fi16
        _ = v17
    }
    v3 = (v3 >> ((255) & 63))
    v3 += 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
