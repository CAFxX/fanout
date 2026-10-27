package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 int16; F1 uint64; F2 int }
type GenS4 struct { F0 uint16 }

func main() {
    v1 := []string{}
    v2 := make(map[int]int)
    var v3 int = 127
    var v4 string = ""
    if true {
        if false {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            v4 = (v4 + v4)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        fmt.Println(v3, v4)
    }
    fmt.Println(v4, v4)
    if (((v4 == (("hello" + v4) + v4)) || (true && false)) && (255 != 421)) {
        v5 := GenS1{F0: "", F1: 766}
        v6 := GenS1{F0: (("\n\t" + "hello") + "foo bar"), F1: 4294967295}
        var v7 uint32 = 42
        _ = v5
        _ = v6
        _ = v7
    }
    v8 := GenS2{F0: v3, F1: true}
    switch (v3 * v3) {
    case -2147483648:
        switch "\n\t" {
        case "foo bar":
            var v9 uint32 = 404
            _ = v9
        case "\n\t":
            if ((100 <= 42) || false) {
                if (((v4 + v4) + "a") != v4) {
                    fmt.Println(v4)
                }
                for fi10 := 0; fi10 < 3; fi10++ {
                    switch v3 {
                    case 2147483647:
                        if (false || true) {
                            v3 += 1
                        }
                    case 32768:
                        fmt.Println(fi10, v4, v4)
                    default:
                        v3 = (-(v3))
                    }
                    switch -2147483648 {
                    case -2147483648:
                        switch "\n\t" {
                        case "foo bar":
                            v11 := v2[fi10]
                            _ = v11
                        case "":
                            if false {
                                fmt.Println(v3, fi10)
                                for fi12 := 0; fi12 < 3; fi12++ {
                                    v13 := GenS3{F0: -69, F1: 4294967295, F2: v3}
                                    var v14 uint16 = 465
                                    var v15 uint64 = 4294967295
                                    _ = fi12
                                    _ = v13
                                    _ = v14
                                    _ = v15
                                }
                                if true {
                                    if false {
                                        _ = len(v2)
                                        v4 = (v4 + v4)
                                    } else {
                                        var v16 uint = 4294967295
                                        _ = v16
                                    }
                                    switch (-(fi10)) {
                                    case 9223372036854775807:
                                        var v17 int = (fi10 % ((2147483647) | 1))
                                        _ = v17
                                    case 32768:
                                        v3 = (^(fi10))
                                    }
                                    v18 := GenS3{F0: 1, F1: 4294967295, F2: (fi10 << ((255) & 63))}
                                    _ = v18
                                }
                            }
                        case "hello":
                            switch (fi10 | fi10) {
                            case -634:
                                v3 = (v3 - (fi10 / (((-(v3))) | 1)))
                            }
                        }
                    case 726:
                        v3 = 0
                    default:
                        var v19 string = ("hello" + v4)
                        _ = v19
                    }
                    _ = fi10
                }
            }
        default:
            var v20 int32 = 42
            _ = v20
        }
    case 0:
        v21 := GenS4{F0: 0}
        _ = v21
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        v3 = (^(fi22))
        v3 = (fi22 % ((42) | 1))
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v4, fi22, v4)
        case "foo bar":
            var v23 uint = 550
            _ = v23
        case "":
            switch "\n\t" {
            case "foo bar":
                switch v3 {
                case 32768:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
            }
        }
        _ = fi22
    }
    fmt.Println(v3, v3)
    v24 := GenS3{F0: -32768, F1: 1, F2: (v3 % ((v3) | 1))}
    fmt.Println(v4, v4)
    for fi25 := 0; fi25 < 3; fi25++ {
        var v26 uint16 = 255
        _ = fi25
        _ = v26
    }
    switch v4 {
    case "hello":
        var v27 int32 = 42
        _ = v27
    default:
        var v28 int8 = 127
        _ = v28
    }
    var v29 uint8 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v4
    _ = v8
}
