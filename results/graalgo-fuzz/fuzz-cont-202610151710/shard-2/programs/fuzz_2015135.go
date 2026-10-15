package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }

func main() {
    var v1 int32 = 965
    v2 := []uint32{1, 2, 42, 4294967295}
    var v3 uint32 = 65535
    var v4 uint64 = 18446744073709551615
    var v5 uint32 = (v3 ^ 2147483648)
    v6 := GenS1{F0: -2147483648, F1: -128}
    var v7 bool = true
    switch 601 {
    case 42:
        for fi8 := 0; fi8 < 19; fi8++ {
            if ((v7 || (((!true) && ("\n\t" < "")) || (v7 && false))) && v7) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            var v9 uint8 = 127
            _ = fi8
            _ = v9
        }
    case 0:
        _ = len(v2)
    case 128:
        var v10 uint = 1
        _ = v10
    default:
        switch -2147483648 {
        case 42:
            v11 := GenS1{F0: -828, F1: 0}
            _ = v11
        default:
            switch "\n\t" {
            case "a":
                var v12 uint = 4294967295
                _ = v12
            }
        }
    }
    if (int16(v5) >= -32768) {
        switch ("hello" + ("\n\t" + ("\n\t" + ""))) {
        case "foo bar":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case "a":
            v13 := GenS1{F0: 32768, F1: 0}
            _ = v13
        case "":
            if (!v7) {
                v4 += 43
            } else {
                v1 += 0
                var v14 int8 = 42
                fmt.Println(v7)
                _ = v14
            }
        default:
            for fi15 := 0; fi15 < 20; fi15++ {
                var v16 float32 = 1.5
                var v17 int32 = (v1 ^ (v1 / (((v1 >> ((255) & 63))) | 1)))
                _ = fi15
                _ = v16
                _ = v17
            }
        }
        _ = len(v2)
        v7 = ((128 != 0) && v7)
    }
    if ((255 <= 0) && v7) {
        var v18 uint8 = 0
        var v19 string = "hello"
        var v20 int = -78
        _ = v18
        _ = v19
        _ = v20
    } else {
        var v21 float64 = -2.25
        _ = v21
    }
    switch -1 {
    case -646:
        fmt.Println(v4, v4)
    case 32767:
        fmt.Println(v4, v4)
    }
    v2 = append(v2, (v3 >> ((150) & 63)))
    var v22 uint32 = 961
    var v23 bool = (!(v7 && (!v7)))
    for fi24 := 0; fi24 < 13; fi24++ {
        v4 -= v4
        fmt.Println(v22, fi24, v1)
        _ = fi24
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
