package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint8 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 int8; F1 uint32; F2 uint64 }

func main() {
    v1 := []string{"hello", ("hello" + "\n\t")}
    var v2 int = -763
    var v3 string = ("hello" + ("foo bar" + ("" + ("foo bar" + ("foo bar" + "")))))
    var v4 uint8 = 1
    var v5 int16 = -32768
    var v6 float64 = float64(1.5)
    v4 += v4
    if len(v1) > 0 {
        v1[0] = v3
    }
    var v7 string = v3
    for fi8 := 0; fi8 < 11; fi8++ {
        v5 -= v5
        var v9 int32 = 0
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 uint32 = 2
            var v12 uint16 = 1
            switch 2147483647 {
            case 937:
                fmt.Println(fi10, v6)
            default:
                v13 := GenS1{F0: -32768, F1: (v2 % ((4294967295) | 1)), F2: v4}
                _ = v13
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi8
        _ = v9
    }
    var v14 int64 = -28
    if ((int32(v4) < 117) && (v14 == -848)) {
        v4 -= v4
    }
    if (100 == 100) {
        var v15 string = "foo bar"
        _ = v15
    } else {
        var v16 uint8 = (v4 * 221)
        switch 32767 {
        case -239:
            v17 := GenS2{F0: 18446744073709551615, F1: ""}
            _ = v17
        case -2147483648:
            switch (-(v2)) {
            case 4294967295:
                v18 := GenS1{F0: v5, F1: (v2 << ((2) & 63)), F2: (v4 << ((924) & 63))}
                _ = v18
            }
        default:
            switch (v2 & (v2 % ((-422) | 1))) {
            case -947:
                _ = len(v1)
            case 4294967295:
                v19 := GenS2{F0: 255, F1: v3}
                _ = v19
            case 32767:
                switch "\n\t" {
                case "hello":
                    v2 += 128
                default:
                    v14 += v14
                }
            }
        }
        _ = v16
    }
    switch "foo bar" {
    case "":
        for fi20 := 0; fi20 < 9; fi20++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v21 uint = 674
            v7 = v7
            _ = fi20
            _ = v21
        }
    case "\n\t":
        v22 := GenS1{F0: (v5 & (v5 - -32768)), F1: -647, F2: (^(v4))}
        _ = v22
    case "a":
        var v23 uint8 = 246
        _ = v23
    }
    v24 := GenS3{F0: 100, F1: 2, F2: 2}
    _ = len(v1)
    var v25 int16 = v5
    switch v2 {
    case 1:
        if len(v1) > 0 {
            v1[0] = ((v3 + v3) + "\n\t")
        }
    default:
        v26 := GenS1{F0: 32767, F1: v2, F2: (v4 + (v4 >> ((1) & 63)))}
        _ = v26
    }
    fmt.Println(v14, v4, v4)
    switch v3 {
    case "foo bar":
        var v27 float64 = 1.5
        _ = v27
    case "\n\t":
        if len(v1) > 0 {
            v1[0] = ("" + v3)
        }
    case "a":
        fmt.Println(v25, v5, v7)
    }
    v4 -= v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
