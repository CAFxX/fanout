package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int }
type GenS2 struct { F0 int; F1 uint16; F2 int16 }

func main() {
    var v1 int32 = 283
    var v2 uint16 = 0
    var v3 int64 = 9223372036854775807
    var v4 int8 = -128
    var v5 uint64 = 18446744073709551615
    v6 := []uint32{875, 65535, 684}
    v7 := GenS1{F0: -1, F1: 915, F2: -877}
    if false {
        v6 = append(v6, 2)
        v2 = (v2 >> ((872) & 63))
    }
    fmt.Println(v3)
    var v8 uint8 = 42
    if true {
        v2 = 42
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v9 float64 = 1e10
        _ = v9
    }
    if (!false) {
        v10 := GenS1{F0: 159, F1: 127, F2: -977}
        switch "a" {
        case "foo bar":
            if true {
                v1 = int32(v3)
                if ((true || false) && true) {
                    v4 = (v4 - 127)
                } else {
                    v11 := GenS1{F0: 42, F1: 0, F2: 1}
                    _ = v11
                }
            }
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = v10
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        if false {
            var v12 string = "foo bar"
            var v13 int32 = (-(v1))
            for fi14 := 0; fi14 < 14; fi14++ {
                v12 = ""
                _ = fi14
            }
            _ = v12
            _ = v13
        } else {
            fmt.Println(v8)
            if true {
                v4 = (-(v4))
                var v15 float32 = 0.1
                _ = v15
            }
        }
        v2 += v2
    }
    var v16 float32 = (0.1 - ((0.0 / float32(((-2.25 / (1.5 - ((3.14159 / 3.14159) * 1e10))) + float64((0.1 - 3.14159))))) - ((1.5 / 0.1) / 1.5)))
    for fi17 := 0; fi17 < 13; fi17++ {
        if false {
            var v18 int = (fi17 % (((^(fi17))) | 1))
            _ = v18
        }
        switch "" {
        case "a":
            var v19 bool = true
            _ = v19
        case "":
            for fi20 := 0; fi20 < 5; fi20++ {
                switch "foo bar" {
                case "":
                    v6 = append(v6, 42)
                default:
                    switch "\n\t" {
                    case "foo bar":
                        var v21 int = (fi20 | fi20)
                        _ = v21
                    case "a":
                        v6 = append(v6, uint32(v2))
                    case "\n\t":
                        v3 = v3
                    }
                    var v22 int64 = 42
                    _ = v22
                }
                var v23 int16 = 1000
                v23 = -1
                var v24 bool = (!(false && (!true)))
                switch (^(fi20)) {
                case 476:
                    var v25 uint = 255
                    _ = v25
                case 1:
                    v26 := GenS2{F0: fi20, F1: 2, F2: 1000}
                    _ = v26
                case -744:
                    if (((v3 << ((uint(v2)) & 63)) != int64(v1)) || true) {
                        v1 = v1
                    }
                }
                _ = fi20
                _ = v23
                _ = v24
            }
        case "\n\t":
            v27 := GenS1{F0: 902, F1: fi17, F2: 0}
            _ = v27
        }
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
