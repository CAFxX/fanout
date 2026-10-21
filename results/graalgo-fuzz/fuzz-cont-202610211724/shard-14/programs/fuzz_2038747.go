package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 int; F2 int16 }
type GenS3 struct { F0 int64 }

func main() {
    v1 := []string{}
    var v2 uint64 = 2
    var v3 int = 1
    if false {
        v2 -= 2
        switch v3 {
        case -856:
            v4 := GenS1{F0: 2147483647}
            _ = v4
        }
        v2 = v2
    } else {
        v3 -= (v3 | (v3 / ((2147483647) | 1)))
        switch (v3 | (^(v3))) {
        case 32767:
            if true {
                fmt.Println(v2, v2)
            } else {
                v3 = v3
            }
        case 4294967295:
            v1 = append(v1, ("" + "\n\t"))
        case 246:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    v2 += (v2 % ((v2) | 1))
    v3 = 9223372036854775807
    v2 = v2
    switch 32768 {
    case -9223372036854775808:
        if true {
            v5 := GenS2{F0: (v3 - v3), F1: v3, F2: -830}
            var v6 uint = 613
            fmt.Println(v3, v6, v6)
            _ = v5
            _ = v6
        } else {
            for fi7 := 0; fi7 < 13; fi7++ {
                fmt.Println(v3)
                _ = fi7
            }
        }
    }
    fmt.Println(v2, v3)
    var v8 int64 = 880
    switch -112 {
    case -46:
        if ((true || (!(958 < 42))) || true) {
            if false {
                for fi9 := 0; fi9 < 2; fi9++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi9
                }
                for fi10 := 0; fi10 < 20; fi10++ {
                    _ = len(v1)
                    switch "\n\t" {
                    case "foo bar":
                        fmt.Println(fi10)
                    case "a":
                        var v11 int32 = 2147483647
                        _ = v11
                    case "":
                        var v12 uint8 = 140
                        _ = v12
                    }
                    _ = fi10
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            var v13 string = "a"
            fmt.Println(v2, v3)
            _ = v13
        }
    default:
        if len(v1) > 0 {
            v1[0] = "hello"
        }
    }
    switch "\n\t" {
    case "\n\t":
        v14 := GenS3{F0: -859}
        _ = v14
    default:
        var v15 float32 = 3.14159
        _ = v15
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
}
