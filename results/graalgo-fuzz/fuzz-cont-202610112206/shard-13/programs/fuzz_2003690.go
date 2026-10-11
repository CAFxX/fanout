package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint16; F1 string; F2 int16 }

func main() {
    var v1 uint32 = 2147483648
    var v2 float32 = 1e10
    v3 := make(map[int]bool)
    v4 := make(map[int]string)
    if false {
        v1 += 42
    }
    v1 -= v1
    fmt.Println(v1)
    for fi5 := 0; fi5 < 5; fi5++ {
        if (true && false) {
            var v6 string = "foo bar"
            _ = v6
        } else {
            v1 += v1
            switch fi5 {
            case -419:
                for fi7 := 0; fi7 < 17; fi7++ {
                    switch "a" {
                    case "\n\t":
                        fmt.Println(v2, fi7, fi5)
                    case "foo bar":
                        if (fi7 <= fi5) {
                            v8 := GenS1{F0: 4294967295}
                            fmt.Println(fi7, v2)
                            switch ("foo bar" + "\n\t") {
                            case "hello":
                                v1 = v1
                            case "foo bar":
                                var v9 uint16 = 32767
                                _ = v9
                            case "":
                                v10 := GenS1{F0: 642}
                                _ = v10
                            default:
                                v11 := GenS2{F0: 42, F1: 128, F2: 128}
                                v2 = v2
                                _ = v11
                            }
                            _ = v8
                        }
                        fmt.Println(fi5)
                        v1 = (v1 >> ((2) & 63))
                    case "":
                        var v12 uint64 = 878
                        _ = v12
                    default:
                        var v13 int = (^(fi5))
                        _ = v13
                    }
                    var v14 string = "a"
                    v14 = ""
                    if (!(!(1 != 1))) {
                        var v15 int32 = 2147483647
                        v1 -= (v1 ^ v1)
                        v14 = v14
                        v16 := GenS1{F0: 1}
                        v2 = 1.5
                        _ = v15
                        _ = v16
                    }
                    _ = fi7
                    _ = v14
                }
            }
            var v17 int = fi5
            _ = v17
        }
        _ = fi5
    }
    var v18 uint8 = 127
    var v19 int32 = -1
    switch -138 {
    case 2147483647:
        var v20 bool = true
        _ = v20
    case 127:
        v18 -= 58
    }
    if true {
        for fi21 := 0; fi21 < 2; fi21++ {
            var v22 int8 = 1
            _ = fi21
            _ = v22
        }
        delete(v3, -2147483648)
    } else {
        var v23 uint16 = 2
        var v24 string = "foo bar"
        var v25 float64 = 0.0
        _ = v23
        _ = v24
        _ = v25
    }
    for fi26 := 0; fi26 < 11; fi26++ {
        v27 := GenS3{F0: 9223372036854775807}
        _ = fi26
        _ = v27
    }
    v28 := GenS1{F0: 2}
    if false {
        v18 = v18
    }
    v19 -= v19
    v29 := GenS4{F0: 0}
    v30 := GenS5{F0: uint16(v1), F1: "\n\t", F2: -120}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
