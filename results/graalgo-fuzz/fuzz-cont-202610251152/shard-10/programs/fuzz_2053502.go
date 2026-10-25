package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int16; F1 uint; F2 int32 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 int16; F1 uint32; F2 uint64 }
type GenS5 struct { F0 uint64; F1 int16; F2 string }
type GenS6 struct { F0 int8 }

func main() {
    v1 := []string{}
    var v2 uint = 18446744073709551615
    v3 := []string{(("\n\t" + ("hello" + ("\n\t" + "hello"))) + "\n\t"), ("\n\t" + "foo bar"), "hello"}
    v4 := GenS1{F0: -1, F1: -128}
    v5 := GenS2{F0: -32768, F1: v2, F2: -1}
    for fi6 := 0; fi6 < 4; fi6++ {
        v7 := GenS3{F0: fi6, F1: (true || (1 <= 0))}
        switch fi6 {
        case 9223372036854775807:
            for fi8 := 0; fi8 < 2; fi8++ {
                if false {
                    switch (("\n\t" + (("a" + (("\n\t" + (((("\n\t" + (("\n\t" + "") + "a")) + "a") + "a") + "foo bar")) + "a")) + "\n\t")) + "\n\t") {
                    case "":
                        v2 -= (v2 >> (((v2 & v2)) & 63))
                    case "hello":
                        v2 = 354
                    }
                    v9 := GenS4{F0: 0, F1: uint32(v2), F2: 2}
                    _ = v9
                } else {
                    v2 -= (v2 % ((v2) | 1))
                }
                v2 = (v2 * (v2 - v2))
                v2 += v2
                _ = fi8
            }
        case 127:
            fmt.Println(fi6)
        case 32768:
            v10 := GenS5{F0: uint64(v2), F1: 1, F2: (("a" + "") + "")}
            _ = v10
        }
        _ = fi6
        _ = v7
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch "hello" {
    case "hello":
        switch "\n\t" {
        case "foo bar":
            if false {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                var v11 uint32 = 1
                _ = len(v1)
                v12 := GenS5{F0: 42, F1: -32768, F2: "foo bar"}
                _ = v11
                _ = v12
            }
        }
    case "a":
        v13 := GenS2{F0: 689, F1: 1, F2: 0}
        _ = v13
    default:
        var v14 float32 = ((0.0 * 0.0) + float32(0.1))
        _ = v14
    }
    v15 := GenS6{F0: 10}
    fmt.Println(v2)
    v1 = append(v1, ("" + ("" + "hello")))
    v16 := GenS6{F0: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
