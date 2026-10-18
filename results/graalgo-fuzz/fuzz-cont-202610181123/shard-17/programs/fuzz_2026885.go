package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 string; F1 int; F2 bool }

func main() {
    v1 := []string{((("" + "a") + (("\n\t" + "a") + ("hello" + (("" + ("\n\t" + "hello")) + "\n\t")))) + ("hello" + ("foo bar" + "a")))}
    var v2 uint32 = 409
    v3 := []uint32{}
    for fi4 := 0; fi4 < 12; fi4++ {
        var v5 int32 = 2147483647
        if (!(982 <= v5)) {
            var v6 uint = 4294967295
            switch ("" + "foo bar") {
            case "":
                v2 = uint32(v6)
            case "\n\t":
                var v7 int64 = 42
                _ = v7
            default:
                if false {
                    v2 = (v2 ^ 412)
                    v5 = v5
                    for fi8 := 0; fi8 < 10; fi8++ {
                        v5 = int32(v6)
                        v9 := GenS1{F0: 42, F1: v2, F2: -856}
                        fmt.Println(fi8)
                        _ = fi8
                        _ = v9
                    }
                } else {
                    fmt.Println(v2, v5)
                    var v10 int8 = 127
                    _ = v10
                }
            }
            var v11 uint64 = 718
            _ = v11
            _ = v6
        }
        _ = fi4
        _ = v5
    }
    v2 += (v2 / (((v2 % ((4294967295) | 1))) | 1))
    v3 = append(v3, v2)
    fmt.Println(v2)
    v12 := GenS1{F0: 143, F1: v2, F2: 0}
    _ = len(v3)
    switch "\n\t" {
    case "a":
        for fi13 := 0; fi13 < 17; fi13++ {
            if len(v3) > 0 {
                v3[0] = v2
            }
            v14 := GenS2{F0: v2, F1: 2}
            _ = fi13
            _ = v14
        }
    case "":
        fmt.Println(v2)
    case "hello":
        _ = len(v1)
    default:
        switch "hello" {
        case "":
            var v15 uint64 = 948
            _ = v15
        default:
            for fi16 := 0; fi16 < 5; fi16++ {
                fmt.Println(v2, fi16)
                for fi17 := 0; fi17 < 4; fi17++ {
                    var v18 float64 = ((0.1 + 3.14159) * 3.14159)
                    _ = fi17
                    _ = v18
                }
                _ = fi16
            }
        }
    }
    switch "\n\t" {
    case "\n\t":
        v2 -= (v2 | 65535)
    default:
        v2 = (v2 >> ((664) & 63))
    }
    var v19 float64 = 0.0
    v20 := GenS3{F0: "hello", F1: 128, F2: true}
    fmt.Println(v2, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
