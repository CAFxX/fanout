package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 uint32 = 1
    v2 := make(map[int]int)
    v3 := make(map[string]uint32)
    var v4 int = 42
    v1 += 2
    v4 = v4
    fmt.Println(v4, v1)
    var v5 int64 = -376
    v6 := GenS1{F0: 128}
    switch ("a" + (("" + "\n\t") + "hello")) {
    case "":
        fmt.Println(v1, v5)
    default:
        v5 = 614
    }
    if true {
        v7 := GenS1{F0: 241}
        _ = v7
    }
    for fi8 := 0; fi8 < 6; fi8++ {
        fmt.Println(fi8, fi8, v5)
        var v9 int8 = -63
        if (-1 >= 32767) {
            v10 := GenS1{F0: 255}
            v11 := GenS2{F0: -477}
            _ = v10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    if false {
        for fi12 := 0; fi12 < 7; fi12++ {
            for fi13 := 0; fi13 < 7; fi13++ {
                fmt.Println(v1, fi12, fi12)
                _ = fi13
            }
            _ = fi12
        }
        v5 -= v5
        if (false || true) {
            fmt.Println(v4, v5, v4)
        }
    } else {
        var v14 float64 = float64(((1.5 / -2.25) - 1.5))
        switch "\n\t" {
        case "hello":
            switch "\n\t" {
            case "a":
                v1 = v1
            case "hello":
                for fi15 := 0; fi15 < 10; fi15++ {
                    switch (v4 % ((fi15) | 1)) {
                    case 116:
                        v4 = (v4 + -626)
                    }
                    for fi16 := 0; fi16 < 14; fi16++ {
                        if false {
                            v5 -= 38
                            fmt.Println(v14)
                            v14 -= v14
                        }
                        var v17 uint8 = 128
                        _ = fi16
                        _ = v17
                    }
                    var v18 string = "hello"
                    _ = fi15
                    _ = v18
                }
            case "\n\t":
                var v19 uint16 = 65535
                _ = v19
            }
        case "\n\t":
            fmt.Println(v4, v5)
        default:
            v5 -= -1
        }
        for fi20 := 0; fi20 < 15; fi20++ {
            v2[v4] = v4
            v4 = (fi20 / ((v4) | 1))
            _ = fi20
        }
        _ = v14
    }
    var v21 int = v4
    fmt.Println(v1, v1, v1)
    fmt.Println(v21)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
