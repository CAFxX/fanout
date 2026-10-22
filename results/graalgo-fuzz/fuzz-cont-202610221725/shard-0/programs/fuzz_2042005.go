package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int64 }
type GenS2 struct { F0 bool }

func main() {
    var v1 uint8 = 1
    var v2 int = 60
    var v3 int8 = 71
    fmt.Println(v2, v2)
    v4 := GenS1{F0: (1.5 < (float32(0.0) / -2.25)), F1: 1000, F2: 345}
    v2 = v2
    switch "" {
    case "foo bar":
        switch "hello" {
        case "hello":
            v5 := []int{(v2 / ((int(v1)) | 1)), 227, v2, (-(v2))}
            _ = v5
        }
    case "\n\t":
        fmt.Println(v2)
    case "hello":
        var v6 int = v2
        _ = v6
    }
    var v7 uint64 = 1
    v1 -= 0
    var v8 float64 = 3.14159
    v9 := GenS1{F0: (("a" + ("a" + ("\n\t" + ""))) > "foo bar"), F1: -32768, F2: 42}
    switch v2 {
    case 703:
        v1 += 255
    case 9223372036854775807:
        v10 := GenS2{F0: true}
        _ = v10
    case 128:
        v1 = v1
    default:
        v11 := []int{v2, (v2 - v2), v2, (-(v2))}
        _ = v11
    }
    v12 := []int{-222, 0}
    for fi13 := 0; fi13 < 11; fi13++ {
        v14 := GenS1{F0: true, F1: -32768, F2: -803}
        _ = fi13
        _ = v14
    }
    v15 := GenS1{F0: true, F1: 0, F2: 9223372036854775807}
    if false {
        var v16 float32 = float32(1.5)
        _ = v16
    }
    switch "\n\t" {
    case "foo bar":
        switch v2 {
        case 250:
            for fi17 := 0; fi17 < 17; fi17++ {
                var v18 uint32 = 2
                v19 := GenS2{F0: (!(989 < 255))}
                _ = fi17
                _ = v18
                _ = v19
            }
        }
    }
    for fi20 := 0; fi20 < 10; fi20++ {
        switch v2 {
        case -996:
            switch "a" {
            case "hello":
                if ("\n\t" <= "hello") {
                    if false {
                        v21 := GenS2{F0: (!false)}
                        for fi22 := 0; fi22 < 14; fi22++ {
                            var v23 uint32 = 1
                            var v24 uint16 = 1
                            v25 := GenS1{F0: (260 > uint(v1)), F1: -780, F2: 9223372036854775807}
                            _ = fi22
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                        v12 = append(v12, 528)
                        _ = v21
                    } else {
                        var v26 int16 = -505
                        v1 = v1
                        _ = v26
                    }
                    v27 := GenS2{F0: ((0 == 1) || false)}
                    var v28 int64 = 1
                    _ = v27
                    _ = v28
                }
            case "foo bar":
                switch ("" + ("\n\t" + "foo bar")) {
                case "hello":
                    var v29 uint32 = 4294967295
                    _ = v29
                default:
                    switch v2 {
                    case 438:
                        var v30 uint16 = 2
                        _ = v30
                    case -2147483648:
                        v7 = (v7 - (v7 + 42))
                    case 594:
                        _ = len(v12)
                    default:
                        fmt.Println(fi20, v3, v2)
                    }
                }
            case "a":
                fmt.Println(v7)
            }
        default:
            var v31 string = ("a" + ("\n\t" + "\n\t"))
            _ = v31
        }
        _ = fi20
    }
    v8 += v8
    var v32 int16 = -32768
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
