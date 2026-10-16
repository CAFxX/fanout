package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 bool; F1 uint32 }
func genFn0(p0 int8, p1 uint16) int8 {
    p1 = p1
    switch -816 {
    case -554:
        var v7 float32 = (-2.25 * (3.14159 - (1e10 - 3.14159)))
        _ = v7
    case -1:
        switch ("hello" + "a") {
        case "hello":
            v8 := GenS1{F0: 65535, F1: -486}
            _ = v8
        case "a":
            var v9 int64 = 9223372036854775807
            _ = v9
        case "foo bar":
            for fi10 := 0; fi10 < 8; fi10++ {
                v11 := []int{575, (fi10 * (-(fi10))), (fi10 ^ (fi10 + (-(fi10))))}
                var v12 int8 = 102
                _ = fi10
                _ = v11
                _ = v12
            }
        default:
            var v13 float64 = 1.5
            _ = v13
        }
    case 42:
        for fi14 := 0; fi14 < 16; fi14++ {
            p0 = 127
            p0 -= (p0 | (p0 << ((2) & 63)))
            var v15 uint8 = 11
            _ = fi14
            _ = v15
        }
    }
    v16 := GenS1{F0: 4294967295, F1: -143}
    v17 := GenS2{F0: 162, F1: false}
    v18 := GenS3{F0: true, F1: 880}
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    _ = v18
    return (p0 * p0)
}

func main() {
    var v1 int16 = 0
    var v2 int32 = -2147483648
    var v3 int8 = 82
    var v4 int = -498
    var v5 int64 = 9223372036854775807
    var v6 string = ("foo bar" + ("\n\t" + ("\n\t" + "foo bar")))
    v5 = (v5 | 831)
    var v19 uint8 = 255
    var v20 int8 = v3
    v21 := []int{}
    for fi22 := 0; fi22 < 13; fi22++ {
        switch fi22 {
        case -2147483648:
            fmt.Println(v6)
        default:
            _ = len(v21)
        }
        if len(v21) > 0 {
            v21[0] = v4
        }
        var v23 uint64 = 42
        _ = fi22
        _ = v23
    }
    v3 = (v20 & 42)
    fmt.Println(v1, v6)
    if false {
        fmt.Println(v4)
    }
    switch v6 {
    case "\n\t":
        fmt.Println(v20)
    case "hello":
        fmt.Println(v5, v3, v4)
    default:
        switch "\n\t" {
        case "a":
            if true {
                v5 = -9223372036854775808
            } else {
                for fi24 := 0; fi24 < 13; fi24++ {
                    var v25 uint32 = 2147483647
                    v2 = v2
                    _ = fi24
                    _ = v25
                }
                if len(v21) >= 2 {
                    v21 = v21[1:len(v21)]
                }
            }
        case "":
            v5 = int64(v4)
        default:
            if (!(((255 == 32767) && false) || (1 != 522))) {
                var v26 float64 = 3.14159
                _ = v26
            } else {
                var v27 int16 = v1
                var v28 uint = 65535
                _ = v27
                _ = v28
            }
        }
    }
    v6 = ("a" + (v6 + v6))
    var v29 uint = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
