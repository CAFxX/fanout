package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 bool }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 uint32; F1 string; F2 uint8 }
type GenS4 struct { F0 string; F1 uint64; F2 int32 }

func main() {
    var v1 uint16 = 2
    var v2 uint32 = 2147483648
    var v3 int16 = 1000
    v4 := []uint32{v2, 584, (^(v2))}
    var v5 float32 = (3.14159 + (((1.5 + float32(1.5)) / -2.25) + 0.1))
    v6 := []uint32{(v2 * v2)}
    fmt.Println(v1)
    _ = len(v6)
    if true {
        for fi7 := 0; fi7 < 5; fi7++ {
            switch "a" {
            case "\n\t":
                v8 := GenS1{F0: 994, F1: (fi7 % ((fi7) | 1)), F2: ("hello" >= "\n\t")}
                _ = v8
            case "":
                switch fi7 {
                case 32768:
                    fmt.Println(fi7, fi7)
                }
            case "foo bar":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            default:
                for fi9 := 0; fi9 < 16; fi9++ {
                    v2 += v2
                    v10 := GenS2{F0: "\n\t", F1: "a"}
                    v11 := GenS2{F0: ("foo bar" + "hello"), F1: "a"}
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            }
            _ = fi7
        }
        fmt.Println(v3)
    } else {
        if len(v6) > 0 {
            v6[0] = v2
        }
        for fi12 := 0; fi12 < 15; fi12++ {
            switch ("" + (((("\n\t" + "hello") + "foo bar") + ("" + "foo bar")) + "\n\t")) {
            case "\n\t":
                v1 = v1
            }
            var v13 int = (fi12 >> ((65535) & 63))
            v1 = v1
            _ = fi12
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        switch fi14 {
        case -9223372036854775808:
            v5 = v5
        }
        for fi15 := 0; fi15 < 4; fi15++ {
            if len(v4) > 0 {
                v4[0] = 4294967295
            }
            v16 := GenS3{F0: 2147483648, F1: (("" + "hello") + (("\n\t" + "foo bar") + ("" + "\n\t"))), F2: 128}
            _ = fi15
            _ = v16
        }
        switch 903 {
        case 4294967295:
            for fi17 := 0; fi17 < 3; fi17++ {
                v18 := GenS1{F0: 382, F1: (fi14 / ((4294967295) | 1)), F2: true}
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v19 float64 = 1.5
                _ = fi17
                _ = v18
                _ = v19
            }
        case 2147483647:
            switch "" {
            case "":
                for fi20 := 0; fi20 < 17; fi20++ {
                    if false {
                        var v21 uint = 786
                        var v22 string = (((("a" + "") + "") + "\n\t") + ("hello" + "\n\t"))
                        _ = v21
                        _ = v22
                    }
                    if true {
                        v3 = (v3 | 1)
                    }
                    _ = fi20
                }
            case "a":
                if true {
                    var v23 int64 = -1
                    _ = v23
                }
            case "\n\t":
                var v24 uint64 = 908
                _ = v24
            default:
                var v25 int = fi14
                _ = v25
            }
        case -839:
            for fi26 := 0; fi26 < 7; fi26++ {
                v6 = append(v6, v2)
                fmt.Println(v5, fi26, v2)
                v27 := GenS4{F0: (("" + "hello") + "\n\t"), F1: 2, F2: -1}
                _ = fi26
                _ = v27
            }
        }
        _ = fi14
    }
    if len(v4) > 0 {
        v4[0] = v2
    }
    v3 += (v3 & (v3 % ((-32768) | 1)))
    var v28 float64 = (((3.14159 / 1e10) - 0.1) / ((0.0 + 0.1) / 3.14159))
    switch "a" {
    case "":
        var v29 uint64 = 337
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
