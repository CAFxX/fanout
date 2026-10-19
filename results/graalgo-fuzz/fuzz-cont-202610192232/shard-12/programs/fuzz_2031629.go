package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 uint16; F1 uint16 }

func main() {
    var v1 uint8 = 128
    v2 := []string{"foo bar", "\n\t", "foo bar", ("" + "hello")}
    var v3 uint64 = 4294967295
    switch 32767 {
    case 1:
        switch "" {
        case "\n\t":
            if ("foo bar" >= ("" + "\n\t")) {
                v1 = 2
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = len(v2)
            }
        case "a":
            var v4 int8 = -43
            _ = v4
        case "foo bar":
            var v5 uint16 = 236
            _ = v5
        default:
            fmt.Println(v3, v1)
        }
    }
    var v6 int16 = 42
    if (-426 <= 0) {
        for fi7 := 0; fi7 < 12; fi7++ {
            v8 := GenS1{F0: (v6 << ((uint(v1)) & 63)), F1: v1}
            v6 -= 0
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: v6, F1: 2}
        switch "a" {
        case "a":
            if false {
                switch 32768 {
                case 128:
                    var v10 float64 = float64(-2.25)
                    _ = v10
                }
            }
        default:
            v2 = append(v2, ("foo bar" + ("a" + "foo bar")))
        }
        _ = v9
    }
    var v11 int64 = int64(v1)
    if len(v2) > 0 {
        v2[0] = ("" + "a")
    }
    v2 = append(v2, "hello")
    var v12 uint32 = 1
    v11 -= 42
    if false {
        v3 = v3
        switch "a" {
        case "":
            v13 := GenS2{F0: -22, F1: (v6 & v6)}
            _ = v13
        default:
            v12 = (v12 >> ((228) & 63))
        }
        var v14 int8 = 101
        _ = v14
    } else {
        for fi15 := 0; fi15 < 6; fi15++ {
            var v16 float64 = (float64(0.0) / 1.0)
            switch fi15 {
            case -10:
                var v17 int16 = v6
                _ = v17
            default:
                v16 += float64((float32(v16) - -2.25))
            }
            v11 -= (v11 & (v11 << ((1) & 63)))
            _ = fi15
            _ = v16
        }
    }
    var v18 string = "foo bar"
    v19 := GenS3{F0: 2, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v6
}
