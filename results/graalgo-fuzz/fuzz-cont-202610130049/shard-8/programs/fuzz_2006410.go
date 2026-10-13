package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint8; F1 int64; F2 uint64 }
func genFn0(p0 bool, p1 int16, p2 uint16) uint32 {
    fmt.Println(p0, p2)
    fmt.Println(p2)
    v7 := []int{32767, 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}

func main() {
    var v1 bool = (117 <= 2147483647)
    var v2 float32 = 0.0
    var v3 int16 = -32768
    v4 := make(map[string]uint32)
    var v5 int16 = (v3 / (((v3 | (-(v3)))) | 1))
    var v6 float32 = v2
    for fi8 := 0; fi8 < 15; fi8++ {
        v1 = v1
        v2 = v6
        if false {
            v9 := []int{}
            _ = v9
        } else {
            v10 := GenS1{F0: 0, F1: fi8}
            var v11 uint32 = 42
            _ = v10
            _ = v11
        }
        _ = fi8
    }
    var v12 int16 = (v5 >> ((973) & 63))
    if ((-128 <= 127) || (!v1)) {
        fmt.Println(v5)
        v13 := GenS1{F0: 42, F1: 457}
        for fi14 := 0; fi14 < 8; fi14++ {
            v5 -= (v12 >> ((65535) & 63))
            var v15 int8 = 127
            _ = fi14
            _ = v15
        }
        _ = v13
    } else {
        var v16 string = (("" + "a") + ("hello" + ""))
        var v17 uint = 2
        _ = v16
        _ = v17
    }
    switch ("hello" + "") {
    case "":
        var v18 float64 = 0.0
        _ = v18
    default:
        var v19 uint16 = 202
        _ = v19
    }
    for fi20 := 0; fi20 < 10; fi20++ {
        if v1 {
            v21 := v4["\n\t"]
            _ = v21
        } else {
            v12 = v3
            var v22 int64 = -9223372036854775808
            v23 := GenS2{F0: 127, F1: (-(v22)), F2: 1}
            _ = v22
            _ = v23
        }
        _ = fi20
    }
    fmt.Println(v1, v5)
    if (32767 != 0) {
        v5 = -224
        v3 += -715
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v2)
        case "\n\t":
            for fi24 := 0; fi24 < 19; fi24++ {
                if (((fi24 << ((1) & 63)) <= (fi24 & fi24)) && false) {
                    var v25 uint8 = 243
                    var v26 uint32 = 683
                    _ = v25
                    _ = v26
                }
                _ = fi24
            }
        case "":
            var v27 int32 = 977
            _ = v27
        default:
            var v28 int32 = 2147483647
            _ = v28
        }
    }
    v29 := GenS2{F0: 255, F1: -1, F2: 305}
    var v30 int = 127
    switch "\n\t" {
    case "foo bar":
        delete(v4, ("foo bar" + ((("hello" + "hello") + ("a" + ("foo bar" + ("hello" + "hello")))) + "hello")))
    case "hello":
        v2 -= v2
    case "a":
        fmt.Println(v1)
    default:
        var v31 uint16 = 2
        _ = v31
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
