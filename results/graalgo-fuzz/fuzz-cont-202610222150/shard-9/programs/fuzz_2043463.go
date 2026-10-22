package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 int64 = 809
    var v2 int8 = -63
    var v3 uint32 = 197
    var v4 uint8 = 127
    var v5 uint64 = 42
    var v6 uint32 = v3
    v7 := []int{-1, -461, -614}
    var v8 string = ""
    fmt.Println(v5, v1, v3)
    for fi9 := 0; fi9 < 9; fi9++ {
        switch "a" {
        case "":
            v7 = append(v7, 9223372036854775807)
        case "hello":
            if false {
                var v10 uint64 = (v5 ^ v5)
                if len(v7) > 0 {
                    v7[0] = (fi9 | fi9)
                }
                _ = v10
            } else {
                v11 := GenS1{F0: 0}
                if true {
                    var v12 float32 = 1.5
                    v6 = (^(v3))
                    if false {
                        for fi13 := 0; fi13 < 8; fi13++ {
                            v1 = (-(v1))
                            _ = fi13
                        }
                    } else {
                        if true {
                            v7 = append(v7, 194)
                        } else {
                            v1 = v1
                            var v14 uint32 = (^(v6))
                            _ = v14
                        }
                    }
                    _ = v12
                } else {
                    switch v8 {
                    case "":
                        switch ("a" + v8) {
                        case "a":
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                        case "\n\t":
                            fmt.Println(v3, v3)
                        }
                    }
                    v1 = v1
                }
                _ = v11
            }
        case "a":
            v2 += -1
        default:
            switch fi9 {
            case 42:
                if (!true) {
                    v8 = (((v8 + v8) + "") + v8)
                    if (!true) {
                        var v15 bool = true
                        if v15 {
                            var v16 float32 = (1e10 - float32(3.14159))
                            for fi17 := 0; fi17 < 18; fi17++ {
                                var v18 bool = (v15 && true)
                                _ = fi17
                                _ = v18
                            }
                            var v19 uint = 1
                            _ = v16
                            _ = v19
                        }
                        v20 := GenS1{F0: 1}
                        _ = v15
                        _ = v20
                    }
                }
            case 350:
                v2 -= v2
            case 109:
                v4 = (v4 & 138)
            }
        }
        _ = fi9
    }
    v3 = (v3 ^ uint32(v5))
    v4 -= v4
    var v21 int8 = v2
    v7 = append(v7, 1)
    switch v8 {
    case "foo bar":
        switch v8 {
        case "\n\t":
            var v22 string = (v8 + v8)
            _ = v22
        case "a":
            switch v8 {
            case "foo bar":
                if (!((-9 > 9223372036854775807) && true)) {
                    var v23 uint8 = 42
                    var v24 int16 = 391
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    _ = v23
                    _ = v24
                }
            case "":
                fmt.Println(v3)
            default:
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            }
        case "foo bar":
            v4 -= v4
        default:
            if false {
                v6 -= v3
                var v25 float32 = 3.14159
                _ = v25
            }
        }
    case "a":
        _ = len(v7)
    case "hello":
        for fi26 := 0; fi26 < 15; fi26++ {
            v4 += (^(v4))
            v7 = append(v7, fi26)
            var v27 uint64 = (v5 / (((v5 >> ((65535) & 63))) | 1))
            _ = fi26
            _ = v27
        }
    default:
        v3 = 739
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
