package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
func genFn0(p0 int16) uint16 {
    var v7 int64 = -1
    v7 -= -9223372036854775808
    _ = p0
    _ = v7
    return 155
}
func genFn1() int16 {
    switch 140 {
    case 2147483647:
        var v8 uint8 = 0
        _ = v8
    case 9223372036854775807:
        v9 := []int{-1}
        _ = v9
    case -126:
        if (-67 <= -128) {
            var v10 uint8 = 210
            v10 -= 128
            switch 2147483647 {
            case -220:
                for fi11 := 0; fi11 < 19; fi11++ {
                    fmt.Println(v10, fi11)
                    v12 := []int{fi11, (fi11 << ((uint(v10)) & 63)), (fi11 >> ((800) & 63))}
                    switch ("foo bar" + ("foo bar" + ("hello" + "a"))) {
                    case "foo bar":
                        var v13 float32 = 3.14159
                        _ = v13
                    case "\n\t":
                        v14 := GenS1{F0: 42, F1: 1}
                        _ = v14
                    }
                    _ = fi11
                    _ = v12
                }
            case 42:
                fmt.Println(v10)
            case 32768:
                v10 = v10
            }
            _ = v10
        } else {
            var v15 uint16 = 0
            _ = v15
        }
    }
    v16 := []int{127, 4294967295, 0}
    switch "\n\t" {
    case "hello":
        switch -46 {
        case 32767:
            v17 := GenS1{F0: 127, F1: 0}
            _ = v17
        case -9223372036854775808:
            var v18 uint16 = 255
            _ = v18
        case 9223372036854775807:
            var v19 int32 = -2147483648
            _ = v19
        default:
            switch "" {
            case "a":
                var v20 int = -897
                _ = v20
            default:
                if false {
                    var v21 int8 = -34
                    if ((32767 == 128) && (115 == 128)) {
                        fmt.Println(v21)
                        fmt.Println(v21)
                        var v22 uint = 4294967295
                        _ = v22
                    } else {
                        fmt.Println(v21)
                        _ = len(v16)
                    }
                    _ = v21
                } else {
                    _ = len(v16)
                }
            }
        }
    case "a":
        v16 = append(v16, 42)
    default:
        var v23 int8 = 100
        _ = v23
    }
    var v24 bool = true
    _ = v16
    _ = v24
    return -975
}

func main() {
    v1 := make(map[string]string)
    var v2 int8 = -1
    var v3 int64 = 42
    v4 := make(map[uint32]string)
    v5 := make(map[uint32]uint32)
    var v6 int = -367
    var v25 uint16 = 871
    fmt.Println(v25)
    var v26 int8 = 127
    for fi27 := 0; fi27 < 6; fi27++ {
        var v28 uint16 = genFn0(int16(v26))
        _ = fi27
        _ = v28
    }
    var v29 string = ""
    var v30 int16 = -32768
    var v31 int = 127
    for fi32 := 0; fi32 < 3; fi32++ {
        fmt.Println(v3, v2, v26)
        var v33 uint32 = 42
        _ = fi32
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
