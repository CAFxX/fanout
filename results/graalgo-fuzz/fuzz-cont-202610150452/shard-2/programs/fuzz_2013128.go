package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
func genFn0(p0 uint16, p1 string, p2 int64) uint8 {
    switch "hello" {
    case "":
        for fi7 := 0; fi7 < 13; fi7++ {
            p1 = (p1 + ((p1 + p1) + (p1 + p1)))
            p1 = p1
            for fi8 := 0; fi8 < 10; fi8++ {
                p0 = p0
                var v9 float64 = (((1.5 * 0.0) - ((-2.25 / 1e10) / ((1e10 + 1e10) + 1e10))) * 1.5)
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
    case "foo bar":
        var v10 int = -1
        _ = v10
    case "\n\t":
        p1 = p1
    default:
        var v11 uint8 = 128
        _ = v11
    }
    p0 = 364
    fmt.Println(p2, p1, p2)
    _ = p0
    _ = p1
    _ = p2
    return 127
}

func main() {
    var v1 bool = false
    var v2 string = "hello"
    var v3 uint64 = 733
    v4 := []uint32{4294967295, 2147483648, 71}
    var v5 uint16 = 255
    var v6 int16 = 1000
    var v12 int = 128
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v2, v6)
    var v13 uint32 = 132
    var v14 uint32 = (v13 % (((v13 & 533)) | 1))
    v15 := GenS1{F0: (v5 / ((0) | 1)), F1: v2}
    for fi16 := 0; fi16 < 15; fi16++ {
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 uint32 = 1
            switch v2 {
            case "":
                fmt.Println(v2, fi17, fi16)
            case "a":
                fmt.Println(v12, v18, v5)
            case "hello":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            default:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            _ = fi17
            _ = v18
        }
        _ = fi16
    }
    v2 = (v2 + "foo bar")
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
