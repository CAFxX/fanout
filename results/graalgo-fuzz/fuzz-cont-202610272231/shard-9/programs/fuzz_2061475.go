package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int }
type GenS2 struct { F0 uint64; F1 uint16; F2 bool }
type GenS3 struct { F0 int64; F1 uint16; F2 int16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64; F1 bool }
func genFn0(p0 uint32, p1 uint64, p2 uint16) int32 {
    p1 = 0
    p1 = 1
    p1 -= p1
    v4 := GenS1{F0: 491, F1: true, F2: -549}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 515
}
func genFn1() int8 {
    if (65535 >= 18446744073709551615) {
        for fi5 := 0; fi5 < 10; fi5++ {
            switch fi5 {
            case 2147483647:
                var v6 int8 = 1
                _ = v6
            default:
                var v7 uint16 = 255
                _ = v7
            }
            _ = fi5
        }
        var v8 int = 375
        var v9 int32 = -503
        _ = v8
        _ = v9
    }
    var v10 uint64 = 2
    v11 := GenS2{F0: 311, F1: 32767, F2: (4294967295 != 2147483647)}
    if true {
        var v12 uint8 = 42
        switch "hello" {
        case "foo bar":
            switch "foo bar" {
            case "foo bar":
                for fi13 := 0; fi13 < 19; fi13++ {
                    fmt.Println(v10)
                    _ = fi13
                }
            default:
                fmt.Println(v12)
            }
        default:
            v14 := []int{}
            _ = v14
        }
        _ = v12
    }
    _ = v10
    _ = v11
    return -127
}

func main() {
    v1 := make(map[string]uint32)
    var v2 string = "a"
    var v3 string = (v2 + ("a" + "\n\t"))
    var v15 int = 32767
    v16 := GenS3{F0: 1, F1: 2, F2: 622}
    var v17 string = v2
    var v18 int = (v15 ^ v15)
    v3 = ((("\n\t" + "") + ("foo bar" + v3)) + ((v17 + "\n\t") + v2))
    if (0 >= genFn1()) {
        for fi19 := 0; fi19 < 17; fi19++ {
            v2 = v3
            fmt.Println(v2, v3, v2)
            _ = fi19
        }
        for fi20 := 0; fi20 < 19; fi20++ {
            fmt.Println(v3)
            v3 = "hello"
            v21 := GenS3{F0: -9223372036854775808, F1: 32767, F2: 0}
            _ = fi20
            _ = v21
        }
        v22 := GenS4{F0: 42}
        _ = v22
    }
    v23 := GenS5{F0: 9223372036854775807, F1: ((540 > 2) || (0.1 < (((0.0 * 1e10) - -2.25) * (0.0 - ((1.5 - (((0.0 * 1.5) + (3.14159 - 3.14159)) + 0.0)) - 3.14159)))))}
    for fi24 := 0; fi24 < 19; fi24++ {
        var v25 string = ("\n\t" + "\n\t")
        _ = fi24
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
}
