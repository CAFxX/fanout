package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 uint32 = 0
    var v2 uint64 = 18446744073709551615
    var v3 uint8 = 255
    var v4 float32 = -2.25
    var v5 uint16 = 1
    v6 := []int{int(v1), -2147483648, -310}
    for fi7 := 0; fi7 < 7; fi7++ {
        v1 += v1
        _ = fi7
    }
    v6 = append(v6, 32767)
    v4 = v4
    v3 = 1
    fmt.Println(v1)
    for fi8 := 0; fi8 < 18; fi8++ {
        v1 += (v1 * 302)
        for fi9 := 0; fi9 < 20; fi9++ {
            if false {
                v10 := GenS1{F0: false, F1: 0}
                _ = v10
            } else {
                switch ("\n\t" + "hello") {
                case "hello":
                    switch (fi8 % (((fi9 - fi9)) | 1)) {
                    case 359:
                        v11 := GenS2{F0: v1, F1: 255}
                        _ = v11
                    default:
                        v5 -= v5
                    }
                default:
                    if false {
                        v12 := GenS1{F0: false, F1: 211}
                        switch ("\n\t" + "foo bar") {
                        case "":
                            var v13 uint64 = uint64(v1)
                            _ = v13
                        case "foo bar":
                            fmt.Println(v2)
                            v3 = (v3 - (v3 ^ (v3 - v3)))
                        case "hello":
                            var v14 int32 = 526
                            _ = v14
                        }
                        _ = v12
                    } else {
                        v15 := GenS3{F0: 1}
                        var v16 uint32 = (v1 >> ((65535) & 63))
                        _ = v15
                        _ = v16
                    }
                }
            }
            var v17 bool = false
            _ = fi9
            _ = v17
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi8
    }
    var v18 uint32 = 42
    v19 := GenS2{F0: v1, F1: 21}
    v18 = 65535
    switch ("" + "foo bar") {
    case "\n\t":
        v20 := GenS1{F0: false, F1: 65535}
        _ = v20
    case "a":
        var v21 int16 = -32768
        _ = v21
    default:
        fmt.Println(v5)
    }
    if (!(!false)) {
        var v22 uint64 = v2
        _ = v22
    } else {
        var v23 int = 127
        for fi24 := 0; fi24 < 11; fi24++ {
            v25 := GenS1{F0: (!false), F1: 1}
            _ = fi24
            _ = v25
        }
        _ = v23
    }
    v26 := GenS4{F0: (v5 % (((v5 | v5)) | 1))}
    if false {
        v27 := GenS3{F0: (v1 * v18)}
        _ = v27
    }
    if len(v6) > 0 {
        v6[0] = -27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
