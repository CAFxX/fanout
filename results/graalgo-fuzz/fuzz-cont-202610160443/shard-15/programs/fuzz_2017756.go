package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 int64; F1 int; F2 bool }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 uint }
func genFn0() int32 {
    var v4 int64 = -77
    v4 -= 1
    v4 = v4
    fmt.Println(v4)
    v5 := GenS1{F0: 0, F1: 114}
    _ = v4
    _ = v5
    return 2147483647
}

func main() {
    var v1 float64 = ((0.1 + ((-2.25 + 3.14159) * ((3.14159 * 0.1) / ((-2.25 + 1e10) * (0.0 + -2.25))))) * float64((-2.25 * 1.5)))
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 42
    switch 1 {
    case -9223372036854775808:
        v3 -= (v3 * (v3 >> ((261) & 63)))
    case -159:
        if (1 >= 42) {
            v6 := v2[1]
            _ = v6
        }
    case 832:
        switch 32767 {
        case 4294967295:
            var v7 string = ("foo bar" + "\n\t")
            _ = v7
        case 42:
            switch "" {
            case "":
                v8 := []int{42, 0}
                _ = v8
            case "\n\t":
                switch 32767 {
                case -9223372036854775808:
                    v9 := GenS2{F0: 325, F1: 4294967295, F2: false}
                    _ = v9
                case 774:
                    v10 := GenS3{F0: (v3 >> ((311) & 63)), F1: 1}
                    _ = v10
                case 2147483647:
                    var v11 float64 = (((v1 + v1) + v1) / v1)
                    _ = v11
                default:
                    fmt.Println(v1, v3)
                }
            }
        default:
            switch 1 {
            case 128:
                var v12 bool = true
                _ = v12
            case -427:
                v13 := GenS3{F0: (v3 & (^(v3))), F1: 42}
                _ = v13
            }
        }
    default:
        for fi14 := 0; fi14 < 6; fi14++ {
            v3 = 0
            v15 := []int{fi14}
            _ = fi14
            _ = v15
        }
    }
    v16 := GenS3{F0: v3, F1: -952}
    v17 := []int{-2147483648, -813, 0}
    v1 -= float64(-2.25)
    v18 := GenS4{F0: 0}
    v3 = v3
    _ = len(v2)
    v19 := GenS2{F0: 1, F1: 0, F2: true}
    switch 2147483647 {
    case 0:
        switch (("foo bar" + ("a" + "\n\t")) + "hello") {
        case "":
            for fi20 := 0; fi20 < 9; fi20++ {
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
                _ = fi20
            }
        }
    case 127:
        var v21 int16 = -1
        _ = v21
    case -670:
        fmt.Println(v1)
    }
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
