package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int16; F1 bool; F2 uint8 }

func main() {
    var v1 int64 = 1
    var v2 int32 = 0
    v3 := make(map[string]bool)
    var v4 int = 884
    v5 := make(map[string]int)
    v6 := []int{}
    for fi7 := 0; fi7 < 6; fi7++ {
        switch "a" {
        case "\n\t":
            fmt.Println(fi7, v4)
        }
        var v8 float64 = 1.5
        _ = fi7
        _ = v8
    }
    if true {
        switch v4 {
        case -315:
            v9 := GenS1{F0: 128, F1: ("a" + ("" + (("foo bar" + (("\n\t" + "\n\t") + "\n\t")) + ("\n\t" + "\n\t"))))}
            _ = v9
        case 32767:
            if ((-(v1)) <= -492) {
                if len(v6) > 0 {
                    v6[0] = (v4 / ((v4) | 1))
                }
                fmt.Println(v1)
            } else {
                switch v4 {
                case 9223372036854775807:
                    var v10 int = v4
                    _ = v10
                case -579:
                    var v11 uint64 = 4294967295
                    _ = v11
                case 127:
                    var v12 float32 = 0.1
                    _ = v12
                default:
                    v2 = (v2 + v2)
                }
                fmt.Println(v1)
            }
        default:
            if (!true) {
                _ = len(v6)
            }
        }
    } else {
        v13 := GenS1{F0: 1, F1: ""}
        fmt.Println(v1, v1)
        var v14 float64 = 1.5
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        var v16 float32 = 0.1
        v1 += (v1 ^ int64(v2))
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        v2 -= (v2 & v2)
        fmt.Println(v4, v4)
        _ = fi17
    }
    v1 = -84
    switch "a" {
    case "":
        if false {
            switch ("\n\t" + "foo bar") {
            case "foo bar":
                switch (v4 ^ v4) {
                case 857:
                    switch "\n\t" {
                    case "hello":
                        fmt.Println(v1, v1, v2)
                    default:
                        v4 += (v4 << ((1) & 63))
                    }
                case 2147483647:
                    var v18 int64 = (v1 + v1)
                    _ = v18
                case 32767:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                default:
                    v4 -= v4
                }
            case "hello":
                v19 := GenS1{F0: 164, F1: ("foo bar" + "foo bar")}
                _ = v19
            case "":
                v20 := GenS1{F0: 0, F1: (("hello" + "hello") + ("" + "foo bar"))}
                _ = v20
            }
        } else {
            v2 = (v2 | (^(v2)))
        }
    case "a":
        v21 := GenS2{F0: 1000, F1: false, F2: 3}
        _ = v21
    default:
        v2 += -369
    }
    for fi22 := 0; fi22 < 12; fi22++ {
        var v23 bool = false
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        var v25 bool = (!(v2 == v2))
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        var v27 string = ""
        _ = fi26
        _ = v27
    }
    var v28 int8 = 0
    if true {
        if len(v6) > 0 {
            v6[0] = v4
        }
        var v29 uint64 = 2
        v4 = (v4 >> ((65535) & 63))
        _ = v29
    } else {
        _ = len(v3)
    }
    fmt.Println(v4)
    var v30 int8 = (v28 - v28)
    var v31 uint = 279
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
