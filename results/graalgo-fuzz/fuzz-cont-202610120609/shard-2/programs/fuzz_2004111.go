package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 int32; F1 uint16; F2 uint64 }
func genFn0(p0 string, p1 bool, p2 int) string {
    v5 := []int{128, p2, 9223372036854775807}
    var v6 bool = (127 >= 100)
    fmt.Println(p2)
    v7 := GenS1{F0: p1, F1: 398}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return ("" + p0)
}

func main() {
    var v1 uint32 = 2147483648
    var v2 uint = 4294967295
    v3 := make(map[string]string)
    var v4 uint8 = 255
    switch "\n\t" {
    case "hello":
        var v8 int32 = 42
        _ = v8
    case "\n\t":
        var v9 int8 = 100
        _ = v9
    case "foo bar":
        switch "foo bar" {
        case "":
            var v10 int8 = 127
            _ = v10
        }
    default:
        v1 = (^(v1))
    }
    if false {
        v11 := []int{-419, 9223372036854775807}
        var v12 bool = true
        _ = v11
        _ = v12
    } else {
        v13 := []int{}
        switch 32768 {
        case 128:
            switch "a" {
            case "hello":
                v13 = append(v13, 32767)
            case "a":
                for fi14 := 0; fi14 < 4; fi14++ {
                    v2 -= v2
                    _ = fi14
                }
            case "foo bar":
                v15 := GenS2{F0: 1, F1: 42}
                _ = v15
            default:
                switch 42 {
                case -9223372036854775808:
                    v1 += v1
                case 0:
                    fmt.Println(v4)
                }
            }
        }
        fmt.Println(v1, v2, v1)
        _ = v13
    }
    var v16 uint = (v2 % ((42) | 1))
    for fi17 := 0; fi17 < 18; fi17++ {
        var v18 string = ""
        var v19 bool = false
        _ = fi17
        _ = v18
        _ = v19
    }
    v16 += 2
    var v20 int64 = 243
    var v21 string = ""
    v1 = v1
    var v22 int32 = -748
    var v23 uint32 = (v1 * 1)
    v4 = v4
    fmt.Println(v4)
    var v24 uint8 = v4
    v23 = 850
    for fi25 := 0; fi25 < 19; fi25++ {
        v26 := GenS3{F0: 585, F1: uint16(v24), F2: 693}
        _ = fi25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
