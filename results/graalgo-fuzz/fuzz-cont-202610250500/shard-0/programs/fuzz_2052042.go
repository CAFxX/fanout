package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 string }
type GenS2 struct { F0 uint; F1 bool; F2 uint }

func main() {
    var v1 int32 = 1
    var v2 float64 = 3.14159
    var v3 int64 = 1
    var v4 int32 = (v1 * -594)
    var v5 bool = (!(!(!false)))
    var v6 float32 = 1.5
    switch 32767 {
    case -961:
        var v7 float64 = v2
        _ = v7
    case 986:
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 uint8 = 255
            var v10 uint32 = 42
            _ = fi8
            _ = v10
            _ = v9
        }
    case 0:
        v11 := []int{-15, 541, -1, 32768}
        _ = v11
    default:
        v12 := GenS1{F0: 0, F1: 4294967295, F2: ("" + "foo bar")}
        _ = v12
    }
    var v13 int32 = v1
    if (v6 <= v6) {
        switch ("a" + ("\n\t" + "a")) {
        case "\n\t":
            var v14 uint16 = 167
            _ = v14
        case "a":
            for fi15 := 0; fi15 < 7; fi15++ {
                var v16 int8 = 1
                if (!false) {
                    var v17 int = (^(fi15))
                    switch "foo bar" {
                    case "a":
                        v2 += ((float64(v6) / 1e10) - v2)
                    case "":
                        v17 = fi15
                    default:
                        v18 := []int{}
                        _ = v18
                    }
                    _ = v17
                }
                _ = fi15
                _ = v16
            }
        default:
            var v19 uint8 = 42
            _ = v19
        }
        fmt.Println(v5, v4)
    }
    var v20 int32 = 1
    fmt.Println(v20, v13, v4)
    for fi21 := 0; fi21 < 10; fi21++ {
        switch 32768 {
        case 32768:
            var v22 uint64 = 4294967295
            _ = v22
        default:
            v23 := []int{(fi21 - fi21), -1}
            _ = v23
        }
        var v24 uint16 = 220
        switch fi21 {
        case -961:
            v13 = v20
        default:
            switch -946 {
            case 4294967295:
                v25 := []int{(fi21 & (fi21 * fi21)), (^(fi21)), 850, 42}
                _ = v25
            case -1:
                v3 += -681
            case 9223372036854775807:
                v26 := GenS2{F0: 65535, F1: false, F2: 2}
                _ = v26
            }
        }
        _ = fi21
        _ = v24
    }
    switch 32768 {
    case 79:
        v27 := []int{-9223372036854775808, 32767, -9223372036854775808}
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
