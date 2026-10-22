package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 uint8; F1 bool; F2 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int32 }

func main() {
    v1 := make(map[string]bool)
    v2 := []string{}
    v3 := make(map[string]bool)
    v4 := make(map[uint32]string)
    v5 := make(map[int]uint32)
    var v6 int16 = 1
    var v7 uint8 = 1
    fmt.Println(v7, v6)
    if true {
        var v8 string = ("hello" + "foo bar")
        for fi9 := 0; fi9 < 16; fi9++ {
            switch -1 {
            case 9223372036854775807:
                for fi10 := 0; fi10 < 11; fi10++ {
                    if true {
                        var v11 int8 = -128
                        _ = v11
                    }
                    v12 := GenS1{F0: (("\n\t" + v8) + v8), F1: true, F2: (-(fi10))}
                    if (!false) {
                        v8 = ("foo bar" + "\n\t")
                        if len(v2) > 0 {
                            v2[0] = (v8 + ("" + "\n\t"))
                        }
                        v13 := GenS1{F0: (v8 + "hello"), F1: (116 < -1), F2: fi10}
                        _ = v13
                    }
                    _ = fi10
                    _ = v12
                }
            default:
                v14 := GenS2{F0: (-(fi9)), F1: fi9}
                _ = v14
            }
            var v15 uint = 2
            _ = fi9
            _ = v15
        }
        v6 -= 0
        _ = v8
    }
    var v16 int64 = -500
    switch "a" {
    case "foo bar":
        v16 -= (v16 / ((-141) | 1))
    case "a":
        if (419 > 42) {
            switch ("foo bar" + "\n\t") {
            case "":
                for fi17 := 0; fi17 < 1; fi17++ {
                    if true {
                        fmt.Println(v16, v6)
                        fmt.Println(v16, v6, fi17)
                    }
                    if (!false) {
                        v7 = (^(v7))
                        fmt.Println(v7, fi17, v16)
                    }
                    var v18 int64 = (v16 / ((v16) | 1))
                    _ = fi17
                    _ = v18
                }
            case "foo bar":
                v19 := GenS3{F0: 23, F1: false, F2: (v6 >> ((42) & 63))}
                _ = v19
            case "a":
                switch (("\n\t" + "hello") + "a") {
                case "foo bar":
                    var v20 int = 359
                    _ = v20
                case "":
                    if true {
                        v7 += 143
                    }
                default:
                    v7 += v7
                }
            }
            v16 = (v16 % ((v16) | 1))
            var v21 uint = uint(v7)
            _ = v21
        } else {
            var v22 string = "\n\t"
            fmt.Println(v16, v22)
            if true {
                if false {
                    v23 := GenS4{F0: 1}
                    var v24 int8 = 54
                    _ = v23
                    _ = v24
                }
                v7 = 137
                v25 := GenS4{F0: 346}
                _ = v25
            } else {
                v6 += -1
                v7 = (^(v7))
                for fi26 := 0; fi26 < 14; fi26++ {
                    v6 -= 652
                    var v27 int64 = -9223372036854775808
                    for fi28 := 0; fi28 < 13; fi28++ {
                        _ = len(v5)
                        _ = len(v2)
                        var v29 string = v22
                        var v30 int8 = 1
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                    _ = fi26
                    _ = v27
                }
            }
            _ = v22
        }
    case "":
        _ = len(v4)
    }
    var v31 float32 = 3.14159
    _ = len(v2)
    for fi32 := 0; fi32 < 14; fi32++ {
        fmt.Println(fi32, v16, fi32)
        _ = fi32
    }
    v33 := GenS5{F0: 0}
    var v34 uint64 = 910
    v2 = append(v2, ("a" + ("hello" + "\n\t")))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
