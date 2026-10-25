package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int16 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 bool; F1 int16 }

func main() {
    var v1 bool = true
    v2 := make(map[uint32]bool)
    v3 := make(map[int]string)
    v1 = true
    var v4 int64 = 738
    var v5 int16 = -502
    var v6 uint16 = 65535
    if false {
        v5 -= v5
        var v7 int64 = v4
        _ = v7
    } else {
        var v8 int64 = v4
        v4 = (-(v8))
        var v9 uint32 = 2147483647
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        v11 := GenS1{F0: 2147483647, F1: 65535, F2: v5}
        var v12 uint64 = 537
        _ = fi10
        _ = v11
        _ = v12
    }
    if false {
        switch ("foo bar" + (("a" + "a") + ((("\n\t" + "") + ("a" + (("\n\t" + "") + ("\n\t" + "foo bar")))) + "hello"))) {
        case "foo bar":
            v13 := GenS1{F0: 2147483647, F1: 42, F2: v5}
            _ = v13
        case "":
            fmt.Println(v5)
        case "a":
            fmt.Println(v5)
        default:
            switch 4294967295 {
            case -9223372036854775808:
                for fi14 := 0; fi14 < 7; fi14++ {
                    var v15 bool = false
                    _ = fi14
                    _ = v15
                }
            case 907:
                switch 4294967295 {
                case -9223372036854775808:
                    v16 := GenS1{F0: 2147483647, F1: 1, F2: (v5 / ((v5) | 1))}
                    _ = v16
                default:
                    v17 := GenS1{F0: -2147483648, F1: 4294967295, F2: 1000}
                    _ = v17
                }
            case 964:
                v1 = true
            }
        }
        v6 = v6
        if v1 {
            delete(v2, 2147483648)
            v18 := GenS2{F0: 255, F1: 18446744073709551615}
            _ = v18
        }
    } else {
        for fi19 := 0; fi19 < 11; fi19++ {
            fmt.Println(v6)
            if (((127 > 127) && (23 > 114)) || v1) {
                delete(v2, 65535)
                switch "a" {
                case "hello":
                    var v20 string = ""
                    _ = v20
                default:
                    fmt.Println(fi19)
                }
                switch "\n\t" {
                case "\n\t":
                    v6 = (v6 - 65535)
                case "a":
                    v21 := GenS3{F0: v1, F1: int16(v6)}
                    _ = v21
                case "":
                    v22 := []int{(fi19 / (((fi19 * fi19)) | 1))}
                    _ = v22
                }
            } else {
                switch ("\n\t" + "foo bar") {
                case "":
                    v23 := GenS3{F0: v1, F1: v5}
                    _ = v23
                case "a":
                    v4 = v4
                }
            }
            v24 := []int{fi19, fi19, fi19, (fi19 ^ fi19)}
            _ = fi19
            _ = v24
        }
        v5 -= int16(v4)
    }
    var v25 int8 = -42
    v26 := GenS1{F0: -2147483648, F1: 750, F2: -990}
    v6 = v6
    var v27 uint16 = (^(v6))
    var v28 int8 = v25
    var v29 float64 = float64((0.0 * -2.25))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
