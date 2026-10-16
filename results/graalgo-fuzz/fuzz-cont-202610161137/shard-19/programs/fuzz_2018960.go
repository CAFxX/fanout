package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint64; F1 uint64 }

func main() {
    v1 := []uint32{634, 2147483647, 2147483647, 65535}
    v2 := []string{("hello" + ("\n\t" + ""))}
    var v3 uint64 = 1
    var v4 float64 = -2.25
    fmt.Println(v4, v3)
    v3 = v3
    v3 = 392
    var v5 int64 = 1
    var v6 float32 = 1e10
    fmt.Println(v3)
    for fi7 := 0; fi7 < 20; fi7++ {
        switch "foo bar" {
        case "\n\t":
            _ = len(v1)
        case "foo bar":
            for fi8 := 0; fi8 < 10; fi8++ {
                switch "\n\t" {
                case "":
                    switch "hello" {
                    case "":
                        v9 := GenS1{F0: 480, F1: 1}
                        _ = v9
                    default:
                        var v10 float64 = v4
                        _ = v10
                    }
                case "\n\t":
                    v11 := GenS1{F0: -32768, F1: 178}
                    _ = v11
                }
                _ = fi8
            }
        case "hello":
            v6 -= v6
        }
        for fi12 := 0; fi12 < 18; fi12++ {
            fmt.Println(v6)
            if len(v1) > 0 {
                v1[0] = 2
            }
            _ = fi12
        }
        _ = fi7
    }
    v5 = -1
    if false {
        var v13 float32 = 1.5
        var v14 uint32 = 42
        for fi15 := 0; fi15 < 14; fi15++ {
            fmt.Println(v13, v14)
            _ = fi15
        }
        _ = v13
        _ = v14
    }
    if true {
        v3 += v3
        v5 = (v5 ^ -1)
        fmt.Println(v6)
    } else {
        v3 = v3
        v5 -= -363
        switch "hello" {
        case "foo bar":
            v16 := GenS1{F0: 1000, F1: 255}
            _ = v16
        case "hello":
            v2 = append(v2, "a")
        case "\n\t":
            if (!(!(true || true))) {
                for fi17 := 0; fi17 < 16; fi17++ {
                    var v18 float64 = (v4 * ((3.14159 + 1.5) * (0.1 + 0.1)))
                    _ = fi17
                    _ = v18
                }
                v3 += (v3 + v3)
            } else {
                v3 = 0
                for fi19 := 0; fi19 < 11; fi19++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v20 := GenS2{F0: (v3 - (v3 / ((v3) | 1))), F1: v3}
                    _ = fi19
                    _ = v20
                }
            }
        }
    }
    v2 = append(v2, "foo bar")
    fmt.Println(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
