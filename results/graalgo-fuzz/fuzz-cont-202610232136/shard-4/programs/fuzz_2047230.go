package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string; F1 uint64; F2 uint64 }
type GenS5 struct { F0 bool; F1 int16 }

func main() {
    var v1 uint64 = 529
    v2 := []string{}
    v3 := []string{("" + ("\n\t" + "a")), "\n\t"}
    var v4 uint = 0
    v5 := make(map[string]string)
    var v6 int64 = -777
    if (true || ((0 != 44) || false)) {
        v4 -= (^(v4))
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v6 -= v6
    var v7 int16 = 1000
    v6 = v6
    if true {
        switch "a" {
        case "foo bar":
            for fi8 := 0; fi8 < 1; fi8++ {
                fmt.Println(v1, v4, fi8)
                _ = fi8
            }
        default:
            v9 := GenS1{F0: "\n\t", F1: uint8(v4)}
            _ = v9
        }
        _ = len(v3)
    } else {
        v10 := GenS2{F0: (("foo bar" + "foo bar") + "\n\t")}
        _ = v10
    }
    if len(v2) > 0 {
        v2[0] = ("a" + ("hello" + "a"))
    }
    if true {
        v11 := GenS3{F0: v7}
        var v12 float64 = -2.25
        _ = v11
        _ = v12
    } else {
        var v13 uint32 = 697
        fmt.Println(v6)
        v14 := GenS2{F0: ""}
        _ = v13
        _ = v14
    }
    if true {
        var v15 uint64 = v1
        v5[("" + (("\n\t" + (("a" + ("foo bar" + "foo bar")) + "a")) + ("foo bar" + "")))] = "foo bar"
        v16 := GenS4{F0: "\n\t", F1: (v15 & v15), F2: (v15 / (((v15 | 18446744073709551615)) | 1))}
        _ = v15
        _ = v16
    }
    var v17 float64 = -2.25
    switch "" {
    case "a":
        fmt.Println(v6)
    case "foo bar":
        v18 := GenS5{F0: (false && true), F1: (v7 - -32768)}
        _ = v18
    default:
        v1 -= 2
    }
    v5[(("\n\t" + "\n\t") + "hello")] = ""
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
