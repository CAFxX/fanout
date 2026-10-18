package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }

func main() {
    var v1 float64 = ((0.1 + -2.25) - float64(1.5))
    var v2 string = ("foo bar" + "a")
    var v3 uint8 = 61
    var v4 uint32 = 2
    var v5 uint64 = 18446744073709551615
    var v6 string = v2
    var v7 int16 = -1
    if true {
        var v8 uint32 = v4
        for fi9 := 0; fi9 < 18; fi9++ {
            for fi10 := 0; fi10 < 1; fi10++ {
                fmt.Println(v1, v5)
                _ = fi10
            }
            _ = fi9
        }
        if (((((42 < 100) && ((v7 - (^(v7))) < v7)) || false) || (-9223372036854775808 < 176)) || true) {
            for fi11 := 0; fi11 < 5; fi11++ {
                if false {
                    v6 = v2
                    v8 -= (v4 - 2147483647)
                }
                _ = fi11
            }
            if (0.1 != 3.14159) {
                v12 := GenS1{F0: -9223372036854775808, F1: v5}
                _ = v12
            } else {
                fmt.Println(v7, v8, v7)
                v13 := []int{-517, -9223372036854775808, 42, 810}
                _ = v13
            }
        }
        _ = v8
    }
    switch v2 {
    case "hello":
        var v14 uint = 4294967295
        _ = v14
    case "foo bar":
        var v15 int16 = (v7 >> ((4294967295) & 63))
        _ = v15
    case "a":
        v3 = (v3 + v3)
    }
    v16 := GenS1{F0: -530, F1: v5}
    v1 += -2.25
    v6 = "a"
    var v17 uint8 = uint8(v5)
    var v18 int64 = 0
    var v19 int8 = -128
    var v20 uint64 = 2
    for fi21 := 0; fi21 < 20; fi21++ {
        v2 = ("foo bar" + v2)
        if (65535 != 899) {
            var v22 uint32 = (^(v4))
            switch (v2 + v2) {
            case "a":
                var v23 uint32 = (v22 >> ((65535) & 63))
                _ = v23
            case "":
                switch v2 {
                case "foo bar":
                    for fi24 := 0; fi24 < 12; fi24++ {
                        v3 = 96
                        var v25 uint64 = v5
                        v18 = (v18 + (v18 % ((305) | 1)))
                        _ = fi24
                        _ = v25
                    }
                case "\n\t":
                    v7 += int16(v19)
                default:
                    var v26 uint16 = uint16(v17)
                    _ = v26
                }
            default:
                if false {
                    v27 := []int{32768, fi21}
                    fmt.Println(v7)
                    _ = v27
                }
            }
            _ = v22
        } else {
            v20 = 198
        }
        v28 := GenS1{F0: v18, F1: 1}
        _ = fi21
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
