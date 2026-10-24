package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint8 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 bool; F1 int; F2 int8 }

func main() {
    v1 := make(map[int]int)
    var v2 float64 = 0.0
    var v3 int = 42
    v4 := []string{"\n\t", ("hello" + "")}
    var v5 uint16 = 255
    var v6 uint8 = 127
    switch "hello" {
    case "a":
        fmt.Println(v6)
    case "\n\t":
        v6 = v6
    case "":
        v3 += (v3 >> ((65535) & 63))
    }
    v5 = 32767
    fmt.Println(v3, v3, v5)
    v4 = append(v4, "a")
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v7 := GenS1{F0: ((127 <= 112) || (v5 > 32767)), F1: 32767, F2: (v6 % (((v6 % (((v6 * uint8(v5))) | 1))) | 1))}
    v2 = 1.5
    switch "a" {
    case "hello":
        fmt.Println(v3, v2, v3)
    case "\n\t":
        v8 := GenS2{F0: (v2 != 1.5), F1: "\n\t"}
        _ = v8
    case "foo bar":
        switch "" {
        case "hello":
            v6 = (v6 + (v6 + (v6 % (((v6 | 59)) | 1))))
        }
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        switch "hello" {
        case "\n\t":
            for fi10 := 0; fi10 < 3; fi10++ {
                fmt.Println(v5)
                var v11 string = "\n\t"
                if ((988 <= 855) && ("" <= v11)) {
                    if (!(!true)) {
                        var v12 float64 = v2
                        fmt.Println(v12)
                        switch fi9 {
                        case 533:
                            fmt.Println(v2, v2)
                        }
                        if len(v4) > 0 {
                            v4[0] = v11
                        }
                        v13 := GenS3{F0: 0}
                        _ = v12
                        _ = v13
                    } else {
                        if false {
                            fmt.Println(v11)
                            v14 := GenS2{F0: (68 > -1), F1: ((v11 + v11) + ("a" + (v11 + v11)))}
                            v2 = -2.25
                            _ = v14
                        } else {
                            v15 := GenS1{F0: (42 != 0), F1: 851, F2: 21}
                            var v16 uint16 = (v5 | (^(v5)))
                            _ = len(v1)
                            _ = v15
                            _ = v16
                        }
                        var v17 uint8 = v6
                        _ = v17
                    }
                    fmt.Println(v11, v11)
                } else {
                    var v18 uint32 = 416
                    v18 -= (v18 / (((v18 / ((v18) | 1))) | 1))
                    _ = v18
                }
                _ = fi10
                _ = v11
            }
        default:
            v19 := GenS4{F0: v6, F1: 0}
            _ = v19
        }
        _ = fi9
    }
    var v20 uint16 = (v5 / ((65535) | 1))
    var v21 float64 = ((((v2 * v2) - 1.5) - (1.5 / (3.14159 - ((float64(3.14159) / (float64(0.1) / 1.0)) * v2)))) * float64(0.0))
    for fi22 := 0; fi22 < 15; fi22++ {
        switch 128 {
        case 968:
            for fi23 := 0; fi23 < 3; fi23++ {
                var v24 int32 = 617
                v2 = v2
                delete(v1, 32767)
                _ = fi23
                _ = v24
            }
        case -598:
            fmt.Println(fi22, v2, fi22)
        case -449:
            v25 := GenS5{F0: (2 == 42), F1: v3, F2: 1}
            _ = v25
        default:
            for fi26 := 0; fi26 < 5; fi26++ {
                _ = len(v1)
                _ = fi26
            }
        }
        fmt.Println(v3)
        var v27 int = -313
        _ = fi22
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
