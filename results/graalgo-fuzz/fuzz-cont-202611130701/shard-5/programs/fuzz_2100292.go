package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint16; F1 int32 }

func main() {
    var v1 int16 = -1
    v2 := []string{}
    v3 := make(map[uint32]int)
    var v4 int16 = (v1 ^ v1)
    _ = len(v2)
    switch ("" + "") {
    case "a":
        if len(v2) > 0 {
            v2[0] = (("foo bar" + (("a" + "\n\t") + "")) + "hello")
        }
    case "hello":
        fmt.Println(v4, v4)
    case "\n\t":
        v4 = (-(v1))
    }
    v4 += -32768
    for fi5 := 0; fi5 < 10; fi5++ {
        for fi6 := 0; fi6 < 20; fi6++ {
            switch (fi6 ^ (fi5 * (fi5 - fi6))) {
            case 0:
                v4 = v4
            case 9223372036854775807:
                var v7 bool = ((fi5 >= (fi6 & (fi5 / (((fi5 >> ((690) & 63))) | 1)))) || true)
                _ = v7
            case 2147483647:
                fmt.Println(v1)
            }
            var v8 bool = (v1 > v1)
            _ = fi6
            _ = v8
        }
        switch (-(fi5)) {
        case -950:
            fmt.Println(v4, v4, fi5)
        case 127:
            fmt.Println(v4)
        case 32768:
            var v9 int16 = v1
            _ = v9
        default:
            if false {
                if (-606 >= -2147483648) {
                    v4 = (v4 % (((^(v4))) | 1))
                    v10 := GenS1{F0: 0, F1: 1}
                    _ = v10
                } else {
                    var v11 float32 = (((1e10 - 3.14159) - 0.0) + 1e10)
                    _ = v11
                }
                var v12 uint64 = 18446744073709551615
                v4 += v1
                _ = v12
            } else {
                v3[65535] = fi5
                if true {
                    if (2 >= 127) {
                        var v13 uint = 2
                        _ = v13
                    }
                }
                switch "a" {
                case "\n\t":
                    var v14 int64 = 852
                    _ = v14
                case "":
                    for fi15 := 0; fi15 < 8; fi15++ {
                        if len(v2) > 0 {
                            v2[0] = ("" + ("foo bar" + ""))
                        }
                        _ = fi15
                    }
                }
            }
        }
        _ = fi5
    }
    fmt.Println(v1, v1)
    for fi16 := 0; fi16 < 12; fi16++ {
        v17 := GenS2{F0: 0, F1: 207}
        var v18 float64 = 1.5
        fmt.Println(v18)
        _ = fi16
        _ = v17
        _ = v18
    }
    v1 = v1
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
