package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint8; F1 uint8 }

func main() {
    var v1 uint64 = 1
    var v2 int16 = 543
    var v3 bool = true
    var v4 int64 = 1
    var v5 bool = v3
    var v6 uint = 4294967295
    v7 := []int{-2147483648}
    _ = len(v7)
    v8 := GenS1{F0: (v1 >> (((v6 % ((v6) | 1))) & 63))}
    v2 += (^(v2))
    fmt.Println(v6, v4, v4)
    v5 = (false || (2 == v6))
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 int32 = -51
        fmt.Println(v2, v1)
        v5 = false
        _ = fi9
        _ = v10
    }
    switch "foo bar" {
    case "\n\t":
        if len(v7) > 0 {
            v7[0] = 127
        }
    }
    if false {
        switch "a" {
        case "hello":
            v4 = -1
        }
        if false {
            switch 42 {
            case 127:
                _ = len(v7)
            }
        }
        fmt.Println(v6, v3, v3)
    }
    fmt.Println(v4, v5)
    v11 := GenS2{F0: v6, F1: 4294967295}
    if (!(!(-722 != 21))) {
        switch ("\n\t" + (((("a" + (("a" + ("foo bar" + "")) + "\n\t")) + ("hello" + "\n\t")) + ("" + "a")) + "")) {
        case "a":
            _ = len(v7)
        case "\n\t":
            v6 = 2
        case "":
            fmt.Println(v2)
        default:
            var v12 int64 = v4
            _ = v12
        }
    } else {
        v7 = append(v7, 0)
        v2 += v2
    }
    fmt.Println(v2)
    if v5 {
        v13 := GenS2{F0: (v6 ^ v6), F1: uint32(v6)}
        if v3 {
            switch 127 {
            case 4294967295:
                v1 -= v1
            }
            v14 := GenS3{F0: 42, F1: 179}
            for fi15 := 0; fi15 < 3; fi15++ {
                var v16 int = -838
                for fi17 := 0; fi17 < 2; fi17++ {
                    var v18 string = ("\n\t" + "foo bar")
                    if false {
                        _ = len(v7)
                        var v19 int64 = 9223372036854775807
                        _ = v19
                    } else {
                        v1 += v1
                        v3 = ((v4 | (v4 + 0)) > v4)
                        var v20 int64 = (v4 << ((uint(v1)) & 63))
                        v21 := GenS3{F0: 127, F1: 255}
                        v22 := GenS2{F0: v6, F1: 65535}
                        v23 := GenS2{F0: uint(v1), F1: 429}
                        _ = v20
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                    _ = fi17
                    _ = v18
                }
                _ = fi15
                _ = v16
            }
            _ = v14
        } else {
            for fi24 := 0; fi24 < 20; fi24++ {
                fmt.Println(v3, v4)
                if v5 {
                    if false {
                        fmt.Println(v3, v2)
                    } else {
                        var v25 float64 = 0.0
                        v1 = 0
                        _ = v25
                    }
                    fmt.Println(v2, v3)
                    switch ("" + "hello") {
                    case "":
                        if len(v7) > 0 {
                            v7[0] = (fi24 / ((fi24) | 1))
                        }
                    case "\n\t":
                        var v26 uint64 = v1
                        _ = v26
                    default:
                        var v27 uint32 = 65535
                        _ = v27
                    }
                }
                fmt.Println(fi24, v3, v5)
                _ = fi24
            }
            fmt.Println(v1, v4, v4)
            if v5 {
                if false {
                    fmt.Println(v2, v5, v4)
                    if len(v7) > 0 {
                        v7[0] = 379
                    }
                    switch 32767 {
                    case -747:
                        var v28 uint64 = v1
                        _ = v28
                    }
                }
            }
        }
        _ = v13
    } else {
        var v29 uint32 = 2
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
