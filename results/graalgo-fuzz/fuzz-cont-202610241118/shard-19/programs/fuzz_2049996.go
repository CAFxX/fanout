package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 int16 }
func genFn0(p0 int, p1 uint16, p2 int64) int {
    p0 = (p0 << ((4294967295) & 63))
    var v4 uint8 = 128
    for fi5 := 0; fi5 < 4; fi5++ {
        for fi6 := 0; fi6 < 3; fi6++ {
            v7 := GenS1{F0: p1}
            _ = fi6
            _ = v7
        }
        v8 := GenS1{F0: p1}
        _ = fi5
        _ = v8
    }
    if (p1 < 1) {
        switch 1 {
        case -2147483648:
            for fi9 := 0; fi9 < 7; fi9++ {
                v10 := []int{}
                var v11 int16 = 1000
                _ = fi9
                _ = v10
                _ = v11
            }
        case 209:
            switch "hello" {
            case "\n\t":
                var v12 string = "hello"
                _ = v12
            }
        default:
            switch "foo bar" {
            case "":
                var v13 uint8 = v4
                _ = v13
            default:
                var v14 bool = ((p1 % (((p1 << ((255) & 63))) | 1)) <= (p1 - (p1 * p1)))
                _ = v14
            }
        }
        v15 := GenS2{F0: (p2 + -1)}
        v16 := []int{537}
        _ = v15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (p0 >> ((18446744073709551615) & 63))
}
func genFn1(p0 int64, p1 uint16, p2 uint64) uint8 {
    v17 := GenS1{F0: (p1 / ((p1) | 1))}
    v18 := GenS2{F0: int64(p1)}
    if true {
        v19 := GenS2{F0: (p0 + p0)}
        v20 := []int{128}
        v21 := GenS3{F0: 42, F1: -32768}
        _ = v19
        _ = v20
        _ = v21
    } else {
        v22 := []int{-1}
        _ = v22
    }
    switch ("hello" + "\n\t") {
    case "hello":
        v23 := []int{0, 128}
        _ = v23
    case "":
        var v24 int = 32768
        _ = v24
    case "a":
        p0 += p0
    default:
        var v25 uint16 = p1
        _ = v25
    }
    p2 = (p2 << ((uint(p1)) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    return 255
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint32 = 42
    var v3 uint64 = 2
    for fi26 := 0; fi26 < 7; fi26++ {
        if true {
            switch "hello" {
            case "foo bar":
                switch "hello" {
                case "hello":
                    v2 = v2
                case "foo bar":
                    v3 = 42
                case "a":
                    v3 = (v3 & v3)
                }
            case "":
                if (true && (((fi26 >> ((18446744073709551615) & 63)) >= (fi26 % ((fi26) | 1))) && true)) {
                    var v27 int32 = -2147483648
                    _ = v27
                } else {
                    for fi28 := 0; fi28 < 19; fi28++ {
                        fmt.Println(v2, fi26)
                        var v29 int = fi26
                        _ = fi28
                        _ = v29
                    }
                }
            default:
                v3 = (v3 ^ (v3 & (v3 >> ((1) & 63))))
            }
        } else {
            fmt.Println(v3, fi26, v2)
        }
        v30 := []int{fi26, (fi26 & 128)}
        _ = fi26
        _ = v30
    }
    v3 = v3
    v3 = (v3 * (v3 >> ((2) & 63)))
    v3 = (^(v3))
    for fi31 := 0; fi31 < 12; fi31++ {
        var v32 int8 = 0
        var v33 int8 = (v32 ^ (v32 / (((v32 >> ((0) & 63))) | 1)))
        _ = fi31
        _ = v32
        _ = v33
    }
    fmt.Println(v2, v3)
    var v34 uint64 = 462
    v2 -= (v2 * v2)
    v3 = (^(v34))
    var v35 float32 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
}
