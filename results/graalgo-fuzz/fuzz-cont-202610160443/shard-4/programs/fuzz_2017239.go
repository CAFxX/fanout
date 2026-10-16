package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int8 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 976
    var v3 int32 = -1
    var v4 int32 = v3
    fmt.Println(v3, v3)
    var v5 int8 = 100
    v4 = 675
    var v6 uint16 = 336
    v3 = (v4 & 42)
    switch -195 {
    case 0:
        v7 := GenS1{F0: int16(v5), F1: 0, F2: -66}
        _ = v7
    case -804:
        var v8 uint8 = 59
        _ = v8
    default:
        if (!true) {
            if true {
                var v9 string = "hello"
                _ = v9
            } else {
                v10 := GenS1{F0: 32767, F1: 420, F2: (v5 - (v5 >> ((1) & 63)))}
                if (9223372036854775807 < 4294967295) {
                    var v11 int8 = (v5 >> ((18446744073709551615) & 63))
                    v3 += v4
                    _ = v11
                }
                _ = v10
            }
            switch 128 {
            case 9223372036854775807:
                fmt.Println(v3, v3)
            default:
                var v12 bool = (uint32(v2) < 4294967295)
                _ = v12
            }
        } else {
            if (1 <= 179) {
                v3 += (v3 & 42)
            } else {
                var v13 uint32 = 373
                if false {
                    v3 = int32(v13)
                }
                _ = v13
            }
        }
    }
    switch 796 {
    case 4294967295:
        var v14 int16 = 42
        _ = v14
    case 0:
        v15 := []int{}
        _ = v15
    case 32768:
        if (0.1 >= -2.25) {
            v3 = v4
        }
    }
    v16 := GenS1{F0: -32768, F1: 2, F2: (^(v5))}
    fmt.Println(v5, v6)
    v1[828] = 42
    v17 := []int{0, -2147483648, 621}
    delete(v1, 99)
    switch "a" {
    case "\n\t":
        v17 = append(v17, 42)
    case "hello":
        if false {
            switch "hello" {
            case "foo bar":
                v3 -= (^(v4))
            default:
                var v18 int = -1
                _ = v18
            }
        }
    case "":
        var v19 uint16 = v2
        _ = v19
    }
    v17 = append(v17, 2147483647)
    for fi20 := 0; fi20 < 12; fi20++ {
        v5 = (v5 ^ 42)
        if true {
            switch ("foo bar" + "") {
            case "foo bar":
                v17 = append(v17, 242)
            default:
                var v21 uint64 = 0
                _ = v21
            }
        } else {
            var v22 uint8 = 42
            fmt.Println(fi20)
            _ = v22
        }
        _ = fi20
    }
    if true {
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
    } else {
        var v23 float64 = (0.0 - 1.5)
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
