package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint64 }

func main() {
    var v1 uint = 65535
    v2 := []int{32767, 32768, -1, 4294967295}
    var v3 uint64 = 42
    var v4 bool = (!(1 > 0))
    var v5 uint8 = 128
    v6 := make(map[string]int)
    fmt.Println(v5)
    v5 += v5
    var v7 string = "a"
    for fi8 := 0; fi8 < 10; fi8++ {
        if (!(100 <= 60)) {
            v4 = ((1 <= 1) && v4)
        } else {
            var v9 uint8 = (v5 * (v5 << ((uint(v5)) & 63)))
            v10 := v6["foo bar"]
            if false {
                if len(v2) > 0 {
                    v2[0] = -862
                }
                v9 = v5
            }
            _ = v10
            _ = v9
        }
        v11 := v6["\n\t"]
        _ = fi8
        _ = v11
    }
    if (!v4) {
        var v12 uint = 904
        for fi13 := 0; fi13 < 6; fi13++ {
            fmt.Println(v4, v4)
            _ = len(v6)
            delete(v6, v7)
            _ = fi13
        }
        _ = v12
    }
    v14 := GenS1{F0: 42, F1: (v5 % (((v5 | uint8(v1))) | 1)), F2: (v3 + (v3 << (((v1 / ((255) | 1))) & 63)))}
    for fi15 := 0; fi15 < 15; fi15++ {
        switch "hello" {
        case "foo bar":
            switch v7 {
            case "foo bar":
                fmt.Println(v4, v3, v1)
            case "a":
                fmt.Println(fi15, v3, fi15)
            case "hello":
                v6[v7] = (^(fi15))
            }
        }
        _ = fi15
    }
    fmt.Println(v7, v7)
    _ = len(v6)
    switch v7 {
    case "foo bar":
        _ = len(v2)
    default:
        var v16 bool = (v3 == (v3 + v3))
        _ = v16
    }
    if v4 {
        for fi17 := 0; fi17 < 5; fi17++ {
            v3 = 333
            _ = fi17
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        var v18 uint16 = 691
        var v19 int64 = 1
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi20
    }
    var v21 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
