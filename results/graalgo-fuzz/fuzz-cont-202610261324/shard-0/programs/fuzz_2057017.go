package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 uint8; F1 bool }
func genFn0(p0 uint32) int32 {
    p0 = p0
    for fi6 := 0; fi6 < 10; fi6++ {
        fmt.Println(fi6)
        _ = fi6
    }
    var v7 string = ("\n\t" + "hello")
    _ = p0
    _ = v7
    return 1
}

func main() {
    var v1 int8 = 0
    var v2 int8 = v1
    var v3 uint64 = 0
    var v4 uint8 = 253
    var v5 uint16 = 426
    switch int(v3) {
    case 127:
        v8 := GenS1{F0: v3, F1: (v4 << ((458) & 63))}
        _ = v8
    }
    fmt.Println(v5)
    v9 := GenS2{F0: 86}
    if ((true || (!false)) && (2147483647 != 32768)) {
        switch ("foo bar" + (("foo bar" + "hello") + "hello")) {
        case "a":
            for fi10 := 0; fi10 < 5; fi10++ {
                v11 := []int{fi10}
                _ = fi10
                _ = v11
            }
        default:
            switch 2147483647 {
            case 32768:
                if ((v3 | 438) <= 435) {
                    v12 := []int{42, 42, 4294967295}
                    _ = v12
                } else {
                    v13 := GenS3{F0: (false || ((3.14159 < (1.5 * float32(3.14159))) && true)), F1: ("hello" + "foo bar")}
                    _ = v13
                }
            case 4294967295:
                switch "a" {
                case "\n\t":
                    var v14 float32 = -2.25
                    _ = v14
                default:
                    for fi15 := 0; fi15 < 13; fi15++ {
                        v4 = 127
                        var v16 float64 = (3.14159 * 0.1)
                        v4 = v4
                        _ = fi15
                        _ = v16
                    }
                }
            case -1:
                fmt.Println(v1)
            }
        }
        if false {
            switch 32768 {
            case 9223372036854775807:
                v5 += v5
            case 0:
                var v17 uint16 = 65535
                _ = v17
            }
            if (227 == 465) {
                var v18 uint = 887
                var v19 float32 = 0.0
                v19 = (v19 - ((float32(1.5) * v19) / (v19 * v19)))
                _ = v18
                _ = v19
            } else {
                v20 := GenS3{F0: false, F1: ""}
                switch "" {
                case "a":
                    v5 = 2
                }
                _ = v20
            }
            fmt.Println(v4, v4, v5)
        }
        fmt.Println(v2)
    }
    v4 = (v4 & v4)
    switch -9223372036854775808 {
    case -1:
        v1 += (v1 ^ v2)
    }
    v21 := []int{128, -678, 0, 599}
    switch "a" {
    case "a":
        var v22 uint64 = v3
        _ = v22
    case "foo bar":
        if true {
            fmt.Println(v4, v1)
            v23 := GenS4{F0: v4, F1: true}
            _ = v23
        } else {
            fmt.Println(v1, v4, v4)
            v4 += 255
        }
    default:
        switch "\n\t" {
        case "\n\t":
            v24 := GenS2{F0: 127}
            _ = v24
        case "foo bar":
            var v25 uint8 = (v4 % ((v4) | 1))
            _ = v25
        default:
            v21 = append(v21, 0)
        }
    }
    v1 = 127
    v2 -= v1
    switch "" {
    case "\n\t":
        if (0 < int32(v1)) {
            if false {
                var v26 string = "foo bar"
                _ = v26
            }
            fmt.Println(v5, v2)
            v27 := GenS2{F0: -15}
            _ = v27
        } else {
            for fi28 := 0; fi28 < 4; fi28++ {
                v1 += (v1 >> ((0) & 63))
                var v29 float32 = 3.14159
                v30 := GenS1{F0: (v3 % (((^(v3))) | 1)), F1: (^(v4))}
                _ = fi28
                _ = v29
                _ = v30
            }
        }
    default:
        fmt.Println(v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
