package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 bool }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 uint32; F1 bool; F2 int16 }

func main() {
    var v1 float32 = 1e10
    v2 := make(map[int]string)
    v3 := []uint32{2147483648, 521, 65535, 4294967295}
    var v4 uint64 = 869
    for fi5 := 0; fi5 < 19; fi5++ {
        for fi6 := 0; fi6 < 7; fi6++ {
            var v7 float64 = (1.5 + 3.14159)
            _ = fi6
            _ = v7
        }
        switch "" {
        case "a":
            v1 -= v1
        case "\n\t":
            switch (fi5 * fi5) {
            case 32768:
                if true {
                    v1 = v1
                } else {
                    v1 += float32((0.0 - ((1.5 * 1e10) + (0.1 + (0.0 - 1.5)))))
                    var v8 uint64 = 811
                    _ = v8
                }
            case 1:
                v3 = append(v3, 597)
            case 9223372036854775807:
                if ((v4 * v4) >= 2) {
                    v4 -= (v4 % ((423) | 1))
                }
            default:
                switch ((((("hello" + "\n\t") + "foo bar") + "") + "\n\t") + ("hello" + ((("" + (((("\n\t" + "\n\t") + "foo bar") + "") + "a")) + "a") + "\n\t"))) {
                case "hello":
                    var v9 uint = 255
                    _ = v9
                case "\n\t":
                    v10 := GenS1{F0: int64(v4), F1: 0, F2: (1 <= -1)}
                    _ = v10
                }
            }
        case "hello":
            fmt.Println(fi5, v1, v4)
        }
        _ = fi5
    }
    v1 = v1
    if true {
        var v11 float32 = ((v1 / v1) - (v1 + 0.1))
        if (!false) {
            v11 = v11
        }
        _ = v11
    }
    var v12 uint64 = (^(v4))
    v3 = append(v3, 637)
    _ = len(v3)
    v4 += (v12 / ((4294967295) | 1))
    if false {
        v13 := v2[4294967295]
        if false {
            v1 += v1
            for fi14 := 0; fi14 < 6; fi14++ {
                v1 += -2.25
                switch v13 {
                case "hello":
                    v15 := GenS2{F0: (v13 + "a"), F1: ((2147483648 == 2147483648) || ((300 == 521) && (1e10 < 1e10)))}
                    _ = v15
                }
                _ = fi14
            }
        }
        _ = v13
    }
    var v16 bool = false
    switch -1 {
    case 4294967295:
        var v17 uint32 = 42
        _ = v17
    }
    v2[-554] = "foo bar"
    v18 := GenS3{F0: 1, F1: 386}
    if v16 {
        if false {
            fmt.Println(v1)
            v1 = float32((0.0 * -2.25))
        }
        switch "" {
        case "hello":
            switch 669 {
            case 127:
                fmt.Println(v1)
            }
        case "":
            var v19 string = ("" + "a")
            _ = v19
        default:
            _ = len(v3)
        }
        for fi20 := 0; fi20 < 18; fi20++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v21 int64 = -78
            _ = fi20
            _ = v21
        }
    }
    v22 := GenS4{F0: 4294967295, F1: v16, F2: 42}
    v23 := GenS2{F0: (("hello" + ("\n\t" + "a")) + "\n\t"), F1: v16}
    v4 += v4
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v16
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
