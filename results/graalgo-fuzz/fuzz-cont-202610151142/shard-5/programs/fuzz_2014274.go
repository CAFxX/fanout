package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 uint16; F1 uint32 }
type GenS5 struct { F0 int32; F1 uint16 }
type GenS6 struct { F0 int16; F1 int }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]int)
    var v3 uint16 = 32767
    var v4 int = 9223372036854775807
    v5 := GenS1{F0: 1, F1: 100, F2: 1}
    var v6 uint8 = 255
    fmt.Println(v4, v3, v4)
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 float32 = 0.0
        v3 += v3
        for fi9 := 0; fi9 < 17; fi9++ {
            v3 = 318
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    v10 := GenS2{F0: (^(v6)), F1: 0}
    v6 = (^(v6))
    fmt.Println(v4, v3, v4)
    v6 = (v6 / ((v6) | 1))
    switch ("\n\t" + "hello") {
    case "\n\t":
        v3 = (v3 / ((42) | 1))
    case "hello":
        v11 := []int{-400, (v4 / ((127) | 1))}
        _ = v11
    case "a":
        v12 := []int{0, (v4 ^ v4), (v4 << ((840) & 63)), v4}
        _ = v12
    default:
        for fi13 := 0; fi13 < 12; fi13++ {
            for fi14 := 0; fi14 < 3; fi14++ {
                fmt.Println(fi14, v3, v4)
                _ = fi14
            }
            _ = fi13
        }
    }
    var v15 string = "foo bar"
    if (false && true) {
        v15 = ("foo bar" + "\n\t")
        v16 := GenS3{F0: 467, F1: ""}
        _ = v16
    } else {
        switch ("\n\t" + v15) {
        case "":
            var v17 string = (v15 + v15)
            _ = v17
        default:
            switch (v4 + 4294967295) {
            case 0:
                var v18 bool = true
                _ = v18
            case -360:
                switch (-(v4)) {
                case 32768:
                    if false {
                        v15 = v15
                        v19 := GenS4{F0: (v3 | v3), F1: 2147483648}
                        _ = v19
                    } else {
                        for fi20 := 0; fi20 < 20; fi20++ {
                            switch "" {
                            case "\n\t":
                                v21 := GenS4{F0: (v3 >> ((4294967295) & 63)), F1: 42}
                                _ = v21
                            }
                            var v22 uint16 = (v3 & 65535)
                            _ = fi20
                            _ = v22
                        }
                        v23 := []int{}
                        _ = v23
                    }
                case 128:
                    v24 := GenS5{F0: 0, F1: v3}
                    _ = v24
                case 2147483647:
                    v25 := GenS6{F0: -32768, F1: v4}
                    _ = v25
                }
            default:
                v26 := GenS2{F0: 112, F1: 42}
                _ = v26
            }
        }
        fmt.Println(v3, v4, v4)
        fmt.Println(v3)
    }
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
