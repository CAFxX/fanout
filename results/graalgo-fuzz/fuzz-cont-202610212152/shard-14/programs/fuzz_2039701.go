package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint8; F1 int }

func main() {
    var v1 uint8 = 0
    v2 := make(map[string]int)
    v3 := []string{"\n\t"}
    v4 := []int{-720, 770}
    v5 := []int{2147483647, 196, 619}
    var v6 int64 = -9223372036854775808
    v6 -= (v6 | v6)
    fmt.Println(v1, v6)
    if (!(true || ((32767 >= 635) || true))) {
        var v7 bool = ("foo bar" > "a")
        for fi8 := 0; fi8 < 10; fi8++ {
            v7 = ((42 < 1) && (0.0 < 1e10))
            v9 := GenS1{F0: -32768, F1: ("foo bar" + "")}
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    v10 := GenS1{F0: -203, F1: ("a" + ("\n\t" + "\n\t"))}
    v1 -= v1
    v11 := GenS2{F0: 255, F1: 100, F2: v1}
    for fi12 := 0; fi12 < 5; fi12++ {
        v6 = v6
        v13 := GenS2{F0: 32767, F1: -72, F2: 122}
        switch "" {
        case "":
            var v14 int = (fi12 & 127)
            _ = v14
        }
        _ = fi12
        _ = v13
    }
    if false {
        _ = len(v2)
        v6 = v6
        v15 := GenS3{F0: -128}
        _ = v15
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        fmt.Println(v6, fi16)
        for fi17 := 0; fi17 < 20; fi17++ {
            var v18 int64 = (v6 % (((-(v6))) | 1))
            _ = fi17
            _ = v18
        }
        if false {
            _ = len(v5)
            for fi19 := 0; fi19 < 12; fi19++ {
                v1 = (v1 - v1)
                delete(v2, "")
                v20 := GenS3{F0: -109}
                _ = fi19
                _ = v20
            }
            v21 := GenS4{F0: 2, F1: 42}
            _ = v21
        } else {
            var v22 int8 = 42
            for fi23 := 0; fi23 < 3; fi23++ {
                v22 = v22
                switch "foo bar" {
                case "foo bar":
                    fmt.Println(fi16)
                default:
                    var v24 uint = 495
                    _ = v24
                }
                _ = fi23
            }
            _ = v22
        }
        _ = fi16
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
