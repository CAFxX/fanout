package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 42
    var v3 float64 = 0.1
    var v4 uint8 = 127
    var v5 uint8 = (v4 * (v4 >> ((4294967295) & 63)))
    var v6 int8 = 0
    var v7 uint16 = 65535
    v3 -= float64((3.14159 - float32(1e10)))
    v5 = (^(v4))
    v4 = (v4 / ((v4) | 1))
    v8 := GenS1{F0: ("hello" + "hello"), F1: ("" + "\n\t")}
    var v9 uint = 0
    v10 := v1[0]
    v11 := GenS1{F0: "foo bar", F1: ("foo bar" + "\n\t")}
    for fi12 := 0; fi12 < 8; fi12++ {
        v2 = v2
        var v13 uint32 = uint32(v7)
        _ = fi12
        _ = v13
    }
    v7 = (v7 >> (((^(v9))) & 63))
    for fi14 := 0; fi14 < 5; fi14++ {
        v5 = v5
        if (!false) {
            if (v10 || (559 == 2147483647)) {
                var v15 uint = (v9 ^ (v9 >> ((42) & 63)))
                switch "a" {
                case "":
                    var v16 float64 = -2.25
                    _ = v16
                case "a":
                    delete(v1, 2)
                }
                if (fi14 < fi14) {
                    var v17 uint32 = 1
                    v18 := GenS2{F0: v2}
                    var v19 uint32 = v17
                    _ = v17
                    _ = v18
                    _ = v19
                } else {
                    v20 := v1[42]
                    for fi21 := 0; fi21 < 5; fi21++ {
                        v5 += (v5 - (^(v4)))
                        var v22 int = (fi14 & fi21)
                        if v20 {
                            v10 = v10
                        }
                        _ = fi21
                        _ = v22
                    }
                    var v23 int = -2147483648
                    _ = v20
                    _ = v23
                }
                _ = v15
            } else {
                v24 := []int{}
                _ = v24
            }
        }
        _ = fi14
    }
    if (!v10) {
        switch "a" {
        case "\n\t":
            v10 = v10
        default:
            if (4294967295 < 127) {
                v25 := GenS2{F0: (v2 >> ((18446744073709551615) & 63))}
                v9 -= (v9 << ((v9) & 63))
                _ = v25
            }
        }
        var v26 int8 = v6
        var v27 uint = 65535
        _ = v26
        _ = v27
    } else {
        var v28 int64 = -9223372036854775808
        if (v10 || v10) {
            v3 = ((v3 / v3) - v3)
            var v29 uint64 = 637
            _ = v29
        }
        fmt.Println(v3, v5, v3)
        _ = v28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
