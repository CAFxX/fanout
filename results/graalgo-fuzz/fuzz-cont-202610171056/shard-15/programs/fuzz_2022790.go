package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint16 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 uint; F1 int64 }

func main() {
    v1 := []uint32{1, 1, 2}
    var v2 string = "a"
    v3 := []int{-2147483648}
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]string)
    v6 := []int{}
    switch v2 {
    case "\n\t":
        switch "\n\t" {
        case "foo bar":
            v7 := GenS1{F0: 9223372036854775807, F1: -886, F2: 42}
            _ = v7
        case "hello":
            for fi8 := 0; fi8 < 15; fi8++ {
                switch (fi8 - fi8) {
                case -1:
                    for fi9 := 0; fi9 < 20; fi9++ {
                        if (1 != 0) {
                            v10 := GenS2{F0: -128, F1: -128}
                            _ = v10
                        }
                        var v11 uint8 = 128
                        var v12 int32 = 0
                        _ = fi9
                        _ = v11
                        _ = v12
                    }
                }
                _ = fi8
            }
        case "":
            v13 := GenS3{F0: 2147483648, F1: 0}
            _ = v13
        }
    case "a":
        v14 := v4[1]
        _ = v14
    case "foo bar":
        v2 = ""
    default:
        v2 = (v2 + "foo bar")
    }
    v2 = ("foo bar" + "foo bar")
    v2 = "hello"
    var v15 int16 = 1
    var v16 string = (("foo bar" + v2) + (((v2 + ((v2 + "hello") + v2)) + "\n\t") + v2))
    for fi17 := 0; fi17 < 10; fi17++ {
        if false {
            var v18 uint8 = 221
            _ = v18
        }
        _ = fi17
    }
    v19 := GenS3{F0: 0, F1: -128}
    v20 := GenS4{F0: 2, F1: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
