package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 uint32 }

func main() {
    var v1 string = ((("\n\t" + ((("hello" + "") + "") + ("hello" + "\n\t"))) + (("a" + ("\n\t" + "foo bar")) + (("foo bar" + "a") + ""))) + (("hello" + "foo bar") + (((((("a" + "") + ("\n\t" + "")) + "hello") + "a") + "\n\t") + "hello")))
    v2 := make(map[int]bool)
    var v3 int32 = 1
    for fi4 := 0; fi4 < 4; fi4++ {
        var v5 uint8 = 106
        delete(v2, 859)
        var v6 float32 = (1e10 / -2.25)
        _ = fi4
        _ = v5
        _ = v6
    }
    v3 -= (^(v3))
    for fi7 := 0; fi7 < 1; fi7++ {
        for fi8 := 0; fi8 < 5; fi8++ {
            if true {
                fmt.Println(fi7, v1)
                switch v1 {
                case "foo bar":
                    v3 = -962
                case "a":
                    v9 := GenS1{F0: v3, F1: -9223372036854775808}
                    _ = v9
                }
            } else {
                v10 := GenS2{F0: 221}
                v11 := GenS3{F0: 42, F1: 2147483647}
                v3 = 1
                _ = v10
                _ = v11
            }
            if false {
                fmt.Println(fi7, fi8)
                switch "foo bar" {
                case "hello":
                    var v12 string = v1
                    _ = v12
                case "foo bar":
                    v13 := GenS1{F0: v3, F1: 494}
                    _ = v13
                case "\n\t":
                    var v14 uint8 = 33
                    _ = v14
                }
            }
            _ = fi8
        }
        _ = fi7
    }
    var v15 uint16 = 746
    var v16 string = (((v1 + ("" + ("" + "a"))) + ("" + "hello")) + (v1 + "a"))
    switch 0 {
    case 9223372036854775807:
        v17 := GenS3{F0: 2147483648, F1: 2147483648}
        _ = v17
    case -1:
        v3 += (v3 >> ((255) & 63))
    case -871:
        switch v1 {
        case "foo bar":
            var v18 int = -612
            _ = v18
        case "\n\t":
            v16 = v1
        default:
            v1 = ("" + ("foo bar" + (("" + v16) + v16)))
        }
    }
    v19 := []int{}
    fmt.Println(v1)
    v16 = (v16 + (v1 + v1))
    _ = len(v19)
    for fi20 := 0; fi20 < 15; fi20++ {
        fmt.Println(v1, v3)
        v21 := v2[2147483647]
        _ = len(v19)
        _ = fi20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
}
