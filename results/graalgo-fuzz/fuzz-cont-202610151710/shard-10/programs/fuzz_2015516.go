package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 0
    v3 := make(map[string]bool)
    v4 := []int{127}
    var v5 int32 = -936
    var v6 float64 = 0.1
    for fi7 := 0; fi7 < 10; fi7++ {
        v5 += (v5 ^ (v5 - 42))
        v5 -= (v5 | v5)
        _ = fi7
    }
    v8 := GenS1{F0: 127}
    for fi9 := 0; fi9 < 13; fi9++ {
        if (0.1 == 0.1) {
            if (false || (false || false)) {
                v5 -= v5
                v6 -= 1.5
            } else {
                var v10 uint32 = v2
                var v11 bool = true
                _ = v10
                _ = v11
            }
        }
        if (!false) {
            if len(v4) > 0 {
                v4[0] = (-(fi9))
            }
            switch 314 {
            case 0:
                v2 = 65535
            case -1:
                var v12 string = "foo bar"
                _ = v12
            case 32768:
                fmt.Println(fi9, fi9, v5)
            }
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        fmt.Println(v6)
        _ = fi9
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        v14 := GenS1{F0: 100}
        for fi15 := 0; fi15 < 6; fi15++ {
            var v16 float32 = 3.14159
            _ = fi15
            _ = v16
        }
        v5 = v5
        _ = fi13
        _ = v14
    }
    v5 -= (v5 | (v5 / ((v5) | 1)))
    fmt.Println(v5, v6)
    switch -645 {
    case 127:
        var v17 float32 = 1.5
        _ = v17
    }
    fmt.Println(v2, v5, v6)
    switch "" {
    case "hello":
        if (false && false) {
            var v18 int32 = v5
            v2 = (v2 & (v2 / ((v2) | 1)))
            _ = v18
        }
    case "a":
        switch -973 {
        case 2147483647:
            for fi19 := 0; fi19 < 15; fi19++ {
                fmt.Println(fi19)
                v4 = append(v4, (fi19 - (fi19 + fi19)))
                _ = fi19
            }
        case 128:
            switch "" {
            case "foo bar":
                var v20 float64 = (float64((float32((v6 / v6)) - 0.0)) - v6)
                _ = v20
            case "\n\t":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            default:
                if ((false || (!true)) || ((v2 / ((v2) | 1)) < 33)) {
                    var v21 uint64 = 1
                    _ = v21
                }
            }
        default:
            v6 = v6
        }
    }
    var v22 int = 32768
    v22 += v22
    if true {
        switch "a" {
        case "\n\t":
            v23 := GenS1{F0: 0}
            _ = v23
        case "":
            for fi24 := 0; fi24 < 10; fi24++ {
                var v25 uint32 = v2
                fmt.Println(fi24, v5)
                switch "" {
                case "hello":
                    fmt.Println(fi24)
                case "a":
                    for fi26 := 0; fi26 < 18; fi26++ {
                        v27 := GenS1{F0: -1}
                        v22 = 156
                        _ = fi26
                        _ = v27
                    }
                }
                _ = fi24
                _ = v25
            }
        case "hello":
            for fi28 := 0; fi28 < 5; fi28++ {
                v5 -= (v5 << ((255) & 63))
                var v29 uint64 = 2
                _ = fi28
                _ = v29
            }
        default:
            v30 := GenS2{F0: v2}
            _ = v30
        }
    }
    v6 += v6
    switch "" {
    case "hello":
        switch 32768 {
        case 127:
            if false {
                var v31 int8 = int8(v22)
                _ = v31
            }
        case 1:
            v32 := GenS1{F0: 100}
            _ = v32
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
