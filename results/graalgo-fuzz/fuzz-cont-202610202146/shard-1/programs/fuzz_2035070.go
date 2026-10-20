package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 string }
func genFn0() uint8 {
    var v6 int = 9223372036854775807
    v6 = v6
    var v7 bool = true
    v6 = -558
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 uint64, p1 int16) uint8 {
    v8 := GenS1{F0: 42, F1: 614}
    if true {
        if (1 == 795) {
            if true {
                if (((0.0 - 0.1) / (float64(-2.25) - 1.5)) < ((-2.25 * 0.0) + 0.1)) {
                    var v9 uint16 = 717
                    if (-2.25 >= -2.25) {
                        var v10 int8 = 1
                        v9 = uint16(p0)
                        _ = v10
                    }
                    if (!(!false)) {
                        if true {
                            v11 := GenS1{F0: 2147483647, F1: 2}
                            var v12 float32 = 1e10
                            _ = v11
                            _ = v12
                        }
                        switch -535 {
                        case -1:
                            fmt.Println(p1)
                        default:
                            v13 := []int{42, 42}
                            _ = v13
                        }
                        v14 := GenS1{F0: 4294967295, F1: 65535}
                        _ = v14
                    }
                    _ = v9
                } else {
                    p0 = (p0 >> ((0) & 63))
                    v15 := GenS1{F0: 4294967295, F1: 2}
                    _ = v15
                }
                var v16 int16 = (p1 << ((65535) & 63))
                fmt.Println(v16)
                _ = v16
            }
        } else {
            p1 = p1
            var v17 uint8 = 127
            v18 := []int{854, 1}
            _ = v17
            _ = v18
        }
    }
    _ = p0
    _ = p1
    _ = v8
    return 1
}

func main() {
    var v1 int = 4294967295
    var v2 uint16 = 42
    var v3 float32 = 3.14159
    v4 := []string{"\n\t"}
    var v5 int32 = 373
    fmt.Println(v5, v5, v1)
    fmt.Println(v3, v5)
    fmt.Println(v2, v5, v2)
    fmt.Println(v5, v3)
    switch 634 {
    case -243:
        fmt.Println(v3, v3)
    }
    var v19 uint64 = 0
    for fi20 := 0; fi20 < 16; fi20++ {
        if false {
            var v21 float64 = float64((float32(0.1) - 0.0))
            v22 := GenS2{F0: "\n\t"}
            v23 := GenS2{F0: "a"}
            _ = v21
            _ = v22
            _ = v23
        }
        _ = fi20
    }
    for fi24 := 0; fi24 < 9; fi24++ {
        var v25 string = (("\n\t" + "a") + "hello")
        _ = fi24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
