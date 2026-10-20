package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint; F1 string }
func genFn0(p0 uint, p1 bool, p2 uint32) uint16 {
    var v5 int8 = 58
    var v6 float32 = (0.0 + -2.25)
    var v7 int8 = v5
    var v8 uint64 = 890
    v9 := GenS1{F0: ((-732 == 127) || (!p1)), F1: 65535}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 int8 = 5
    var v2 int = -2147483648
    var v3 bool = (2147483647 != 2)
    var v4 uint16 = 255
    var v10 float64 = 0.1
    v11 := []int{}
    var v12 uint32 = 42
    _ = len(v11)
    fmt.Println(v12, v1)
    fmt.Println(v3, v10, v12)
    var v13 int32 = 1
    if len(v11) > 0 {
        v11[0] = (^(v2))
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 uint = 1
        _ = fi14
        _ = v15
    }
    v16 := GenS1{F0: false, F1: (v4 >> ((560) & 63))}
    v11 = append(v11, 127)
    var v17 int32 = (v13 + 1)
    for fi18 := 0; fi18 < 13; fi18++ {
        switch (v2 & 490) {
        case 999:
            v19 := GenS2{F0: 42, F1: false}
            _ = v19
        case 127:
            var v20 uint8 = 190
            _ = v20
        case -9223372036854775808:
            for fi21 := 0; fi21 < 19; fi21++ {
                var v22 int = (fi21 >> ((65535) & 63))
                _ = fi21
                _ = v22
            }
        default:
            fmt.Println(v1, v10)
        }
        _ = fi18
    }
    v23 := GenS3{F0: 267, F1: (("\n\t" + (("hello" + ("foo bar" + ("a" + "\n\t"))) + ("foo bar" + "a"))) + (("foo bar" + "a") + ""))}
    if false {
        v3 = true
        var v24 uint = 1
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
