package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 uint16; F2 int32 }
func genFn0(p0 int64) uint8 {
    var v6 bool = false
    v7 := []int{428, 584, 4294967295}
    _ = p0
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 int32, p1 uint) uint {
    v8 := GenS1{F0: -2147483648, F1: (p0 >> ((1) & 63)), F2: (p1 & p1)}
    v9 := GenS2{F0: 61}
    fmt.Println(p1, p0)
    v10 := GenS2{F0: -1}
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return (^(p1))
}

func main() {
    var v1 uint32 = 0
    var v2 float64 = (1e10 + 1e10)
    v3 := []int{-1, 32768, 0}
    var v4 float64 = ((v2 / 1.0) * (v2 - 1.5))
    v5 := []uint32{}
    var v11 uint16 = 32767
    var v12 float64 = 1e10
    for fi13 := 0; fi13 < 12; fi13++ {
        for fi14 := 0; fi14 < 15; fi14++ {
            switch "" {
            case "a":
                v2 += (v12 * ((v4 * v4) - v2))
                switch "" {
                case "\n\t":
                    var v15 string = "a"
                    _ = v15
                case "hello":
                    v16 := GenS3{F0: 0, F1: v11, F2: 0}
                    _ = v16
                default:
                    v2 = 0.1
                }
                if (!(!(916 != 1))) {
                    _ = len(v3)
                    v11 = (v11 ^ v11)
                }
            }
            _ = fi14
        }
        switch ((("a" + "hello") + "a") + ("" + "hello")) {
        case "hello":
            var v17 bool = ("foo bar" == "")
            _ = v17
        }
        _ = fi13
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        v2 -= v2
        _ = fi18
    }
    var v19 uint8 = 42
    switch "" {
    case "":
        for fi20 := 0; fi20 < 12; fi20++ {
            v21 := GenS2{F0: -9223372036854775808}
            v12 = -2.25
            fmt.Println(v19)
            _ = fi20
            _ = v21
        }
    case "\n\t":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    v22 := GenS1{F0: 987, F1: 42, F2: 4294967295}
    if true {
        v1 = 2147483647
        switch -9223372036854775808 {
        case 32767:
            var v23 int32 = -1
            _ = v23
        case 723:
            var v24 int16 = -32768
            _ = v24
        }
    } else {
        _ = len(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
