package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 uint16; F1 uint64 }
func genFn0(p0 int64) int {
    switch 673 {
    case 755:
        p0 = (p0 & (p0 + (p0 / ((p0) | 1))))
    case -472:
        fmt.Println(p0)
    case -2147483648:
        p0 = (p0 * (p0 ^ -642))
    }
    v5 := GenS1{F0: 4294967295, F1: 127}
    var v6 uint16 = 42
    fmt.Println(v6, v6)
    _ = p0
    _ = v5
    _ = v6
    return 4294967295
}
func genFn1(p0 int) uint8 {
    var v7 uint = 670
    var v8 float32 = 0.1
    v9 := []int{(p0 - (^(p0))), p0}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 bool = false
    v2 := make(map[string]int)
    var v3 int = -2147483648
    var v4 uint32 = 470
    for fi10 := 0; fi10 < 8; fi10++ {
        v3 -= genFn0(-1)
        if v1 {
            if ("\n\t" >= (("foo bar" + "foo bar") + "")) {
                v11 := []int{32767, v3, (v3 * v3)}
                var v12 float64 = 3.14159
                var v13 int16 = 138
                _ = v11
                _ = v12
                _ = v13
            } else {
                var v14 int64 = 9223372036854775807
                _ = v14
            }
        }
        _ = fi10
    }
    v1 = v1
    for fi15 := 0; fi15 < 12; fi15++ {
        fmt.Println(v1, v1, fi15)
        _ = fi15
    }
    fmt.Println(v4)
    switch "a" {
    case "\n\t":
        v16 := GenS2{F0: 382, F1: 2}
        _ = v16
    default:
        v2[""] = v3
    }
    fmt.Println(v3, v3)
    v1 = true
    for fi17 := 0; fi17 < 9; fi17++ {
        for fi18 := 0; fi18 < 18; fi18++ {
            switch "\n\t" {
            case "foo bar":
                switch "foo bar" {
                case "":
                    fmt.Println(v4, v4, v4)
                default:
                    var v19 int8 = 0
                    _ = v19
                }
                if true {
                    var v20 int = fi17
                    fmt.Println(fi17)
                    _ = v20
                }
            case "\n\t":
                v21 := GenS1{F0: 65535, F1: (^(v3))}
                _ = v21
            case "hello":
                v1 = ((v4 & 2) == v4)
            default:
                v22 := GenS2{F0: 9223372036854775807, F1: 0}
                _ = v22
            }
            var v23 int8 = 42
            _ = fi18
            _ = v23
        }
        _ = fi17
    }
    v4 = (v4 << ((934) & 63))
    v3 -= v3
    var v24 int = v3
    fmt.Println(v3, v1)
    switch ("a" + "hello") {
    case "a":
        var v25 int8 = 1
        _ = v25
    case "\n\t":
        delete(v2, ("hello" + "foo bar"))
    case "hello":
        _ = len(v2)
    default:
        for fi26 := 0; fi26 < 6; fi26++ {
            v1 = v1
            v27 := GenS3{F0: 2, F1: 563}
            _ = fi26
            _ = v27
        }
    }
    fmt.Println(v1, v3, v1)
    var v28 bool = (((v24 - genFn0(841)) >= (v3 | v24)) || v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v4
}
