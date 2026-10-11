package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int32 }

func main() {
    v1 := []int{-511, 1, -2147483648}
    v2 := []int{-2147483648, 42, 570, 9223372036854775807}
    var v3 uint32 = 0
    switch "" {
    case "hello":
        _ = len(v2)
    case "a":
        v3 = (v3 ^ (v3 - (v3 << ((4294967295) & 63))))
    case "foo bar":
        v4 := GenS1{F0: 42, F1: true}
        _ = v4
    default:
        var v5 uint32 = v3
        _ = v5
    }
    if (128 >= 175) {
        if (!(2 >= 2)) {
            v6 := GenS2{F0: 427}
            v3 = 2147483647
            _ = v6
        }
        switch ("" + "a") {
        case "":
            v3 += 2147483648
        case "hello":
            for fi7 := 0; fi7 < 12; fi7++ {
                for fi8 := 0; fi8 < 5; fi8++ {
                    fmt.Println(fi7, fi8)
                    v2 = append(v2, 9223372036854775807)
                    _ = fi8
                }
                _ = fi7
            }
        case "foo bar":
            fmt.Println(v3)
        }
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        for fi10 := 0; fi10 < 1; fi10++ {
            _ = len(v2)
            if (1.5 > (1e10 - (0.1 / 1.5))) {
                if len(v2) > 0 {
                    v2[0] = fi9
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v3 += v3
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            _ = fi10
        }
        var v11 string = ""
        switch fi9 {
        case 563:
            var v12 string = "foo bar"
            _ = v12
        case -755:
            v3 += 65535
        case 33:
            switch v11 {
            case "hello":
                _ = len(v1)
            case "a":
                var v13 int64 = -1
                _ = v13
            }
        default:
            if len(v1) > 0 {
                v1[0] = 32768
            }
        }
        _ = fi9
        _ = v11
    }
    v3 += v3
    for fi14 := 0; fi14 < 10; fi14++ {
        if true {
            var v15 int = -338
            var v16 int = (fi14 | v15)
            _ = v15
            _ = v16
        }
        _ = fi14
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 += (v3 / ((2) | 1))
    if false {
        fmt.Println(v3)
        if true {
            v3 = (v3 | (^(v3)))
        }
        var v17 int16 = -1
        _ = v17
    } else {
        var v18 int8 = 75
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v18 = v18
        _ = v18
    }
    v3 -= 223
    v19 := GenS1{F0: 0, F1: (-9223372036854775808 == 32768)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
}
