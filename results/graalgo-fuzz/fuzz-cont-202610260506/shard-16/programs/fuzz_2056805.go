package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 int; F2 bool }
func genFn0() uint16 {
    var v4 int32 = -1
    if (0.0 == 3.14159) {
        var v5 uint32 = 0
        v6 := GenS1{F0: 251}
        _ = v5
        _ = v6
    }
    v7 := []int{}
    v4 += v4
    _ = v4
    _ = v7
    return 255
}

func main() {
    v1 := []int{-1, -9223372036854775808}
    var v2 int16 = 950
    v3 := make(map[uint32]string)
    var v8 float32 = float32(float64(0.1))
    if ("hello" == ("hello" + "\n\t")) {
        v2 = (v2 % ((v2) | 1))
        switch "hello" {
        case "hello":
            var v9 int16 = (v2 << ((4294967295) & 63))
            _ = v9
        case "\n\t":
            delete(v3, 2147483647)
        default:
            delete(v3, 65535)
        }
    } else {
        v10 := GenS1{F0: 1}
        v1 = append(v1, 128)
        _ = v10
    }
    if false {
        for fi11 := 0; fi11 < 2; fi11++ {
            fmt.Println(v8, fi11)
            var v12 int8 = 42
            _ = fi11
            _ = v12
        }
    }
    var v13 uint16 = 2
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 uint16 = 32767
        var v16 int32 = 413
        _ = fi14
        _ = v15
        _ = v16
    }
    v13 = 0
    v17 := GenS2{F0: ("foo bar" + "a"), F1: -2147483648, F2: true}
    switch ("a" + "a") {
    case "":
        if (-128 >= 127) {
            v2 = (v2 & v2)
        } else {
            var v18 uint64 = 42
            _ = v18
        }
    case "foo bar":
        v8 -= 0.0
    case "a":
        switch 100 {
        case 61:
            var v19 uint16 = 32767
            _ = v19
        default:
            if true {
                v8 -= v8
            }
        }
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        switch fi20 {
        case -2147483648:
            fmt.Println(v2, fi20, fi20)
        default:
            var v21 uint8 = 216
            _ = v21
        }
        _ = fi20
    }
    delete(v3, 4294967295)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v8
}
