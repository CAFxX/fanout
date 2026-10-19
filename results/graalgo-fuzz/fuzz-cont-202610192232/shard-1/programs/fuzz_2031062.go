package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 string }
type GenS2 struct { F0 uint64; F1 int; F2 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 int16; F1 uint64; F2 uint }
func genFn0(p0 uint8, p1 uint64) int16 {
    p1 = (p1 >> ((975) & 63))
    var v4 int16 = 1000
    _ = p0
    _ = p1
    _ = v4
    return v4
}

func main() {
    var v1 uint64 = 981
    v2 := make(map[uint32]int)
    v3 := make(map[string]int)
    for fi5 := 0; fi5 < 6; fi5++ {
        v1 = (v1 - v1)
        if true {
            v6 := []int{(fi5 * fi5), fi5, 32768, 383}
            _ = v6
        } else {
            _ = len(v3)
            if (9223372036854775807 > (fi5 * (fi5 ^ fi5))) {
                fmt.Println(v1)
                _ = len(v3)
                v1 += 42
            }
            var v7 uint = 42
            _ = v7
        }
        fmt.Println(v1, v1)
        _ = fi5
    }
    var v8 int = 127
    var v9 uint64 = v1
    switch "a" {
    case "hello":
        v10 := GenS1{F0: 4294967295, F1: "\n\t", F2: "\n\t"}
        _ = v10
    case "foo bar":
        if (!false) {
            v8 -= v8
        }
    case "\n\t":
        v11 := GenS1{F0: 65535, F1: ("a" + ("hello" + ("\n\t" + "a"))), F2: "\n\t"}
        _ = v11
    default:
        if false {
            switch (("hello" + "\n\t") + "a") {
            case "":
                switch v8 {
                case 1:
                    var v12 uint16 = 0
                    _ = v12
                case 357:
                    for fi13 := 0; fi13 < 20; fi13++ {
                        v1 = 697
                        _ = fi13
                    }
                case -836:
                    v8 = v8
                }
            case "a":
                v14 := GenS2{F0: v9, F1: 127, F2: -1}
                _ = v14
            case "foo bar":
                var v15 uint32 = 1
                _ = v15
            default:
                delete(v3, "foo bar")
            }
        }
    }
    fmt.Println(v1, v9)
    v16 := GenS3{F0: 9223372036854775807}
    for fi17 := 0; fi17 < 3; fi17++ {
        v18 := GenS4{F0: 79, F1: 172}
        _ = fi17
        _ = v18
    }
    v19 := GenS5{F0: 1000, F1: 92, F2: 0}
    v20 := GenS3{F0: 523}
    switch "" {
    case "foo bar":
        var v21 string = "foo bar"
        _ = v21
    case "hello":
        var v22 string = "a"
        _ = v22
    case "\n\t":
        v23 := GenS3{F0: 1}
        _ = v23
    default:
        v9 += 450
    }
    for fi24 := 0; fi24 < 19; fi24++ {
        v8 += (fi24 % ((559) | 1))
        if (0 <= 214) {
            v8 -= 2147483647
            switch "a" {
            case "foo bar":
                fmt.Println(v1, v1)
            case "":
                v1 += v1
            case "a":
                v25 := v2[1]
                _ = v25
            default:
                v1 = 267
            }
            v9 -= 42
        } else {
            var v26 string = "a"
            v9 = 792
            _ = v26
        }
        _ = len(v2)
        _ = fi24
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v8
    _ = v9
}
