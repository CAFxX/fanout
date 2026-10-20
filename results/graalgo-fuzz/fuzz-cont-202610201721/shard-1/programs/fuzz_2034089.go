package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 string }

func main() {
    var v1 int32 = 2147483647
    var v2 uint16 = 0
    var v3 uint = 1
    v4 := make(map[uint32]bool)
    var v5 int64 = -749
    for fi6 := 0; fi6 < 3; fi6++ {
        v3 = 255
        _ = fi6
    }
    var v7 uint32 = 2147483647
    v2 = 65535
    v8 := []int{int(v5)}
    switch "" {
    case "":
        for fi9 := 0; fi9 < 5; fi9++ {
            if (v7 == v7) {
                switch "" {
                case "a":
                    switch "hello" {
                    case "\n\t":
                        for fi10 := 0; fi10 < 7; fi10++ {
                            v7 = (^(v7))
                            var v11 float64 = 1.5
                            var v12 uint8 = 57
                            v13 := GenS1{F0: (fi9 & -9223372036854775808), F1: (^(v7))}
                            _ = fi10
                            _ = v11
                            _ = v12
                            _ = v13
                        }
                    case "foo bar":
                        for fi14 := 0; fi14 < 16; fi14++ {
                            _ = len(v8)
                            _ = fi14
                        }
                    case "hello":
                        var v15 int32 = (v1 & (^(v1)))
                        _ = v15
                    }
                case "foo bar":
                    var v16 uint = (v3 % (((v3 & (v3 >> (((v3 << (((v3 | (v3 / ((v3) | 1)))) & 63))) & 63)))) | 1))
                    _ = v16
                }
                v5 = (v5 + -1)
            }
            for fi17 := 0; fi17 < 7; fi17++ {
                if ((("\n\t" + "foo bar") + "") > "a") {
                    switch ("\n\t" + (("hello" + ("\n\t" + "\n\t")) + (("hello" + "foo bar") + ("" + "hello")))) {
                    case "a":
                        switch ("hello" + "\n\t") {
                        case "hello":
                            var v18 int8 = -1
                            _ = v18
                        default:
                            fmt.Println(v3)
                        }
                    case "foo bar":
                        v3 = (v3 & uint(v2))
                    default:
                        v19 := GenS1{F0: (fi17 | 32768), F1: 0}
                        _ = v19
                    }
                } else {
                    fmt.Println(fi9, v7, fi17)
                }
                v20 := GenS1{F0: (fi9 >> ((927) & 63)), F1: 239}
                _ = fi17
                _ = v20
            }
            _ = fi9
        }
    case "hello":
        v21 := v4[298]
        _ = v21
    case "\n\t":
        v22 := GenS2{F0: int32(v2)}
        _ = v22
    }
    var v23 int8 = -128
    _ = len(v4)
    v1 = (v1 + v1)
    v24 := GenS3{F0: "hello"}
    var v25 uint16 = (v2 << ((v3) & 63))
    v4[(v7 << ((v3) & 63))] = ((^(v7)) > (v7 - 1))
    v5 += v5
    for fi26 := 0; fi26 < 6; fi26++ {
        v1 = (-(v1))
        var v27 int32 = (^(v1))
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
