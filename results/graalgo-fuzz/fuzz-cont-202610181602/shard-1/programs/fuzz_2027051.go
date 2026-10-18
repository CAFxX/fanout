package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 string = ("\n\t" + "foo bar")
    var v2 uint = 65535
    var v3 int32 = 42
    var v4 float64 = 3.14159
    switch -461 {
    case -9223372036854775808:
        v1 = (v1 + "\n\t")
    case -26:
        if false {
            if false {
                fmt.Println(v2)
                switch "a" {
                case "a":
                    var v5 int = 9223372036854775807
                    _ = v5
                case "hello":
                    v6 := []int{-132, -538}
                    _ = v6
                }
            }
            if true {
                switch "foo bar" {
                case "hello":
                    for fi7 := 0; fi7 < 13; fi7++ {
                        v3 = v3
                        fmt.Println(fi7, v2)
                        _ = fi7
                    }
                }
            } else {
                switch v1 {
                case "foo bar":
                    switch v1 {
                    case "foo bar":
                        var v8 string = (("hello" + "a") + v1)
                        _ = v8
                    case "hello":
                        v9 := GenS1{F0: 127, F1: int16(v3)}
                        _ = v9
                    default:
                        var v10 uint8 = 110
                        _ = v10
                    }
                case "a":
                    var v11 uint32 = 282
                    _ = v11
                default:
                    for fi12 := 0; fi12 < 12; fi12++ {
                        v13 := GenS2{F0: 42}
                        v1 = ("foo bar" + "")
                        _ = fi12
                        _ = v13
                    }
                }
            }
        }
    case 127:
        fmt.Println(v2, v4, v3)
    }
    v14 := GenS1{F0: 42, F1: 42}
    v15 := []int{-2147483648, -375}
    v4 -= (float64(((0.1 * ((0.1 - 1.5) + float32(float64(-2.25)))) - (1.5 + float32(v4)))) * v4)
    v15 = append(v15, 0)
    var v16 uint32 = 151
    v3 -= (^(v3))
    for fi17 := 0; fi17 < 16; fi17++ {
        for fi18 := 0; fi18 < 4; fi18++ {
            fmt.Println(fi18, fi18)
            _ = len(v15)
            if (!true) {
                if len(v15) > 0 {
                    v15[0] = (fi18 / ((fi18) | 1))
                }
                fmt.Println(fi17)
                v1 = (v1 + "a")
            } else {
                v19 := GenS1{F0: 1, F1: -696}
                _ = len(v15)
                _ = v19
            }
            _ = fi18
        }
        var v20 uint64 = 2
        v2 -= 42
        _ = fi17
        _ = v20
    }
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    v2 = 910
    if ((uint64(v16) > 2) && (true && true)) {
        var v21 float64 = v4
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
