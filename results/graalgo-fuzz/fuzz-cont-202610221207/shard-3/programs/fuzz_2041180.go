package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint32; F1 int; F2 bool }

func main() {
    var v1 int8 = -42
    var v2 int8 = (v1 % (((v1 << ((18446744073709551615) & 63))) | 1))
    v3 := make(map[uint32]uint32)
    v1 -= 1
    if false {
        if false {
            v4 := GenS1{F0: -2147483648, F1: 42}
            fmt.Println(v2, v2)
            _ = v4
        } else {
            v1 = (v1 + v1)
            var v5 float32 = 0.1
            var v6 uint8 = 1
            _ = v5
            _ = v6
        }
        v1 -= 19
        var v7 string = ("foo bar" + "hello")
        _ = v7
    } else {
        v1 -= (v1 + (v1 + (v2 ^ v1)))
        v3[2147483647] = 387
        fmt.Println(v2)
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 uint64 = 1
        v9 = v9
        switch "hello" {
        case "hello":
            fmt.Println(fi8, v2, v2)
        case "a":
            switch ("foo bar" + "foo bar") {
            case "foo bar":
                for fi10 := 0; fi10 < 13; fi10++ {
                    v11 := []int{fi8}
                    fmt.Println(v2, fi10)
                    var v12 int16 = -272
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            case "":
                for fi13 := 0; fi13 < 18; fi13++ {
                    v14 := []int{}
                    _ = fi13
                    _ = v14
                }
            case "\n\t":
                var v15 uint = 2
                _ = v15
            }
        }
        _ = fi8
        _ = v9
    }
    if ((v1 << ((466) & 63)) > v1) {
        for fi16 := 0; fi16 < 9; fi16++ {
            _ = len(v3)
            fmt.Println(v1)
            if false {
                for fi17 := 0; fi17 < 7; fi17++ {
                    v2 -= 127
                    _ = fi17
                }
                var v18 int8 = (^(v2))
                v19 := GenS2{F0: 65535, F1: (fi16 + (fi16 / ((-28) | 1))), F2: (2147483648 < 441)}
                _ = v18
                _ = v19
            }
            _ = fi16
        }
        v1 = v1
        var v20 int = 2147483647
        _ = v20
    }
    v2 = (v2 & v1)
    var v21 uint = 0
    v22 := GenS2{F0: 2, F1: -2147483648, F2: true}
    var v23 float32 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
