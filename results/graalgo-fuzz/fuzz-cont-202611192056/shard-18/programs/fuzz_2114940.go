package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 string; F1 uint8; F2 uint64 }
type GenS3 struct { F0 string }
func genFn0() uint8 {
    var v7 int8 = 127
    var v8 uint8 = 42
    v8 -= 128
    var v9 uint16 = 0
    _ = v7
    _ = v8
    _ = v9
    return (v8 % (((^(v8))) | 1))
}

func main() {
    var v1 uint16 = 577
    var v2 uint32 = 631
    v3 := make(map[uint32]bool)
    var v4 int8 = 127
    v5 := make(map[int]uint32)
    v6 := make(map[uint32]string)
    fmt.Println(v2)
    for fi10 := 0; fi10 < 14; fi10++ {
        v2 -= 2147483647
        var v11 bool = (128 >= 0)
        v4 = -38
        _ = fi10
        _ = v11
    }
    v1 -= 2
    v12 := GenS1{F0: v2, F1: (v4 - v4)}
    v13 := GenS2{F0: "foo bar", F1: genFn0(), F2: 2}
    switch "foo bar" {
    case "\n\t":
        var v14 int32 = -2147483648
        _ = v14
    }
    if true {
        var v15 float64 = (1.5 + (1e10 - ((0.1 * 3.14159) * 1e10)))
        _ = v15
    }
    v16 := []int{-2147483648, 0, 42}
    fmt.Println(v1)
    for fi17 := 0; fi17 < 5; fi17++ {
        if false {
            switch (fi17 % (((^(fi17))) | 1)) {
            case -9223372036854775808:
                switch -600 {
                case 0:
                    var v18 uint32 = (v2 | (^(v2)))
                    _ = v18
                case -894:
                    v19 := GenS3{F0: "\n\t"}
                    _ = v19
                case -2147483648:
                    fmt.Println(v4, fi17, v1)
                }
            case 32767:
                if len(v16) > 0 {
                    v16[0] = (fi17 ^ -906)
                }
            }
            var v20 uint64 = 2
            _ = v20
        }
        switch "\n\t" {
        case "hello":
            var v21 int16 = 53
            _ = v21
        }
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
