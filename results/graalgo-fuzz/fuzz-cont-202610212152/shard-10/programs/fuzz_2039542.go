package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint32 }
type GenS2 struct { F0 bool; F1 uint32 }

func main() {
    var v1 uint = 627
    var v2 float64 = ((float64(0.0) - -2.25) - (((0.1 - 0.0) + (3.14159 - 1.5)) + (0.0 + 0.1)))
    var v3 float64 = (v2 + v2)
    v4 := make(map[string]string)
    var v5 int8 = -1
    delete(v4, ("\n\t" + ("" + "hello")))
    var v6 uint = v1
    v7 := []int{128, 32767}
    v3 = float64(3.14159)
    if false {
        var v8 string = "hello"
        for fi9 := 0; fi9 < 10; fi9++ {
            if (!true) {
                var v10 int8 = v5
                _ = v10
            } else {
                var v11 int = (fi9 | fi9)
                fmt.Println(v11)
                if false {
                    switch (fi9 & fi9) {
                    case -776:
                        v12 := GenS1{F0: -9223372036854775808, F1: -1, F2: 2147483648}
                        _ = v12
                    }
                } else {
                    var v13 int32 = 166
                    _ = v13
                }
                _ = v11
            }
            if true {
                if true {
                    for fi14 := 0; fi14 < 15; fi14++ {
                        v7 = append(v7, -2147483648)
                        fmt.Println(v6)
                        v2 = v3
                        _ = fi14
                    }
                    v8 = (v8 + v8)
                }
            } else {
                var v15 int8 = v5
                var v16 uint16 = 173
                _ = v15
                _ = v16
            }
            v17 := v4["a"]
            _ = fi9
            _ = v17
        }
        _ = v8
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        for fi19 := 0; fi19 < 11; fi19++ {
            if (1.5 >= (float64(1e10) - ((((v3 - v3) - float64((float32(1e10) / (((-2.25 + (0.0 + ((0.1 * (0.1 * 1e10)) + 1e10))) + 0.1) + 1.5)))) - v3) * 0.0))) {
                var v20 uint64 = uint64(v1)
                var v21 uint32 = 65535
                v22 := GenS1{F0: 9223372036854775807, F1: fi19, F2: v21}
                v23 := GenS2{F0: (-319 != 42), F1: (v21 << ((v1) & 63))}
                _ = v20
                _ = v21
                _ = v22
                _ = v23
            }
            v3 = (((v3 * (v3 * 0.0)) / ((v2 / (float64(-2.25) + 1e10)) / -2.25)) * ((v2 + (v3 - ((v2 / (v2 / 0.1)) - v2))) - v3))
            _ = fi19
        }
        _ = fi18
    }
    v6 = (v6 - v6)
    switch "hello" {
    case "a":
        if false {
            v24 := GenS2{F0: false, F1: uint32(v6)}
            _ = v24
        } else {
            delete(v4, ("hello" + "\n\t"))
        }
    case "foo bar":
        v3 -= float64(0.0)
    default:
        var v25 string = "\n\t"
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
