package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int16 }
type GenS2 struct { F0 string }
func genFn0() int8 {
    switch 32768 {
    case 216:
        if true {
            var v5 uint64 = 0
            for fi6 := 0; fi6 < 11; fi6++ {
                v5 += (v5 * v5)
                _ = fi6
            }
            _ = v5
        } else {
            switch "a" {
            case "hello":
                var v7 string = ("\n\t" + "a")
                _ = v7
            case "foo bar":
                var v8 int8 = 127
                _ = v8
            }
            v9 := GenS1{F0: 952, F1: 1, F2: 246}
            _ = v9
        }
    case -2147483648:
        if (!(160 == 127)) {
            v10 := []int{4294967295}
            _ = v10
        } else {
            v11 := GenS1{F0: 2, F1: 0, F2: 1000}
            _ = v11
        }
    case 1:
        v12 := []int{127}
        _ = v12
    }
    var v13 int32 = 42
    if (1 == 255) {
        var v14 int64 = 9223372036854775807
        v15 := GenS1{F0: 32767, F1: 42, F2: -986}
        _ = v14
        _ = v15
    }
    _ = v13
    return -128
}

func main() {
    var v1 float64 = 1.5
    v2 := make(map[int]int)
    var v3 uint8 = 0
    var v4 uint64 = uint64(v3)
    var v16 int8 = genFn0()
    v16 += (v16 * v16)
    if (0 < 602) {
        var v17 uint8 = v3
        fmt.Println(v16)
        v16 -= 18
        _ = v17
    } else {
        if false {
            switch -1 {
            case 9223372036854775807:
                v3 -= v3
            case 127:
                switch "hello" {
                case "a":
                    v4 = 42
                }
            case -802:
                v3 += 161
            default:
                if false {
                    v18 := []int{128}
                    _ = v18
                } else {
                    fmt.Println(v16, v16, v16)
                    switch "" {
                    case "a":
                        switch -518 {
                        case -1:
                            v4 -= v4
                        }
                    case "":
                        fmt.Println(v16, v4, v16)
                    default:
                        v1 = (float64(0.0) - float64((-2.25 + 1.5)))
                    }
                    v19 := GenS1{F0: 2, F1: v4, F2: 451}
                    _ = v19
                }
            }
            var v20 int8 = v16
            _ = v20
        }
    }
    var v21 int16 = 42
    fmt.Println(v16, v4, v4)
    v22 := []int{2147483647, 128, 42}
    v16 -= (v16 & v16)
    v21 = 1
    v4 += 42
    v23 := GenS2{F0: ("a" + "\n\t")}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
