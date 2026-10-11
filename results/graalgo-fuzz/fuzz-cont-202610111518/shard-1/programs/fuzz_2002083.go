package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 uint16 }
type GenS2 struct { F0 string; F1 bool; F2 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64 }

func main() {
    v1 := []int{128, 1}
    var v2 int16 = 42
    v3 := []uint32{}
    v2 = (v2 * 42)
    for fi4 := 0; fi4 < 18; fi4++ {
        var v5 uint64 = 18446744073709551615
        v6 := GenS1{F0: 42, F1: 88, F2: 0}
        _ = fi4
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        fmt.Println(v2)
        _ = fi7
    }
    switch "hello" {
    case "hello":
        v2 = v2
    case "\n\t":
        fmt.Println(v2)
    case "foo bar":
        v8 := GenS2{F0: "foo bar", F1: true, F2: 1}
        _ = v8
    default:
        for fi9 := 0; fi9 < 2; fi9++ {
            v2 = (^(v2))
            _ = fi9
        }
    }
    switch ("\n\t" + ("\n\t" + "")) {
    case "":
        if (42 <= 2) {
            _ = len(v3)
            v10 := GenS3{F0: 0}
            switch "" {
            case "hello":
                var v11 uint64 = 4294967295
                _ = v11
            case "\n\t":
                switch 374 {
                case 466:
                    v2 -= (v2 ^ (v2 * v2))
                case 1:
                    for fi12 := 0; fi12 < 8; fi12++ {
                        v2 -= (v2 / (((v2 << ((583) & 63))) | 1))
                        var v13 uint8 = 0
                        _ = fi12
                        _ = v13
                    }
                }
            case "":
                v14 := GenS4{F0: 84}
                _ = v14
            }
            _ = v10
        } else {
            v2 += v2
            for fi15 := 0; fi15 < 2; fi15++ {
                v2 += 0
                fmt.Println(fi15)
                _ = fi15
            }
            if len(v3) > 0 {
                v3[0] = 2
            }
        }
    }
    if len(v3) > 0 {
        v3[0] = 4294967295
    }
    switch 2147483647 {
    case 0:
        for fi16 := 0; fi16 < 18; fi16++ {
            if (!(4294967295 <= 2)) {
                var v17 uint = 591
                var v18 int32 = 0
                _ = v17
                _ = v18
            }
            _ = fi16
        }
    }
    var v19 uint16 = 0
    if ("" <= ("foo bar" + "\n\t")) {
        var v20 bool = false
        _ = v20
    }
    fmt.Println(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
}
