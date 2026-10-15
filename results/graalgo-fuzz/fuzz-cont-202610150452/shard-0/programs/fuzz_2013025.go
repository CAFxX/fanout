package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int8 }
func genFn0(p0 int16) int16 {
    var v4 uint64 = 0
    switch 1 {
    case 667:
        var v5 int = -742
        _ = v5
    case -149:
        fmt.Println(p0)
    case 2147483647:
        if (-9223372036854775808 != 128) {
            v6 := []int{2147483647, 128, 2147483647, 1}
            fmt.Println(p0, p0)
            _ = v6
        } else {
            var v7 int8 = 46
            _ = v7
        }
    }
    if (9223372036854775807 >= -1) {
        var v8 float64 = 3.14159
        var v9 string = "hello"
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v4
    return (p0 << ((142) & 63))
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]bool)
    v3 := make(map[string]bool)
    _ = len(v2)
    var v10 uint32 = 0
    v10 = v10
    v10 += 447
    v10 += (v10 >> ((18446744073709551615) & 63))
    switch 1 {
    case -1:
        v11 := GenS1{F0: 127, F1: 18446744073709551615, F2: int8(v10)}
        _ = v11
    case 42:
        v12 := []int{391, 902, 32767, 32767}
        _ = v12
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        switch ("" + "\n\t") {
        case "":
            for fi14 := 0; fi14 < 19; fi14++ {
                switch "foo bar" {
                case "hello":
                    if false {
                        var v15 int32 = -1
                        v16 := GenS1{F0: 0, F1: 687, F2: 1}
                        switch "\n\t" {
                        case "a":
                            var v17 uint16 = 1
                            v17 = v17
                            _ = v17
                        case "":
                            v18 := GenS1{F0: 255, F1: 501, F2: 100}
                            _ = v18
                        default:
                            v19 := []int{(fi14 >> ((42) & 63))}
                            _ = v19
                        }
                        _ = v15
                        _ = v16
                    }
                case "a":
                    if ((!(!true)) || true) {
                        switch "a" {
                        case "":
                            v10 = v10
                        }
                    }
                case "foo bar":
                    var v20 uint32 = 42
                    _ = v20
                }
                v10 = 2
                _ = fi14
            }
        case "\n\t":
            v2[fi13] = true
        }
        switch 127 {
        case 0:
            v10 -= (v10 >> ((5) & 63))
        case -32:
            v10 += (v10 + v10)
        case 4294967295:
            v21 := []int{(^(fi13)), fi13, fi13}
            _ = v21
        default:
            v22 := GenS1{F0: 255, F1: 42, F2: -1}
            _ = v22
        }
        v23 := GenS1{F0: uint8(v10), F1: 255, F2: 127}
        _ = fi13
        _ = v23
    }
    var v24 uint16 = 42
    fmt.Println(v24, v24)
    delete(v2, int(v10))
    var v25 uint32 = 0
    v10 = (v10 / ((v10) | 1))
    v25 = uint32(v24)
    if false {
        v26 := GenS1{F0: 127, F1: 255, F2: 42}
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v25
    _ = v3
}
