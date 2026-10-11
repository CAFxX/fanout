package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 uint16; F1 uint16; F2 uint64 }

func main() {
    var v1 float64 = (0.1 * 0.0)
    var v2 float32 = (1e10 - 0.1)
    var v3 int32 = -96
    var v4 float64 = (v1 * (v1 * (v1 + ((((v1 / ((1e10 - -2.25) - v1)) / 1e10) + -2.25) - v1))))
    switch ((("" + "a") + "") + ("" + (("hello" + "\n\t") + "hello"))) {
    case "a":
        var v5 uint16 = 32767
        _ = v5
    case "":
        fmt.Println(v4)
    }
    var v6 int = 1
    switch v6 {
    case 127:
        v7 := GenS1{F0: 42}
        _ = v7
    case 32767:
        v8 := GenS1{F0: 378}
        _ = v8
    default:
        if (!(!false)) {
            var v9 bool = (!("hello" > "\n\t"))
            fmt.Println(v1, v1)
            _ = v9
        } else {
            for fi10 := 0; fi10 < 10; fi10++ {
                v4 -= -2.25
                if (9223372036854775807 != -9223372036854775808) {
                    if true {
                        switch "\n\t" {
                        case "\n\t":
                            for fi11 := 0; fi11 < 9; fi11++ {
                                var v12 uint16 = 2
                                v13 := GenS1{F0: 654}
                                fmt.Println(v2, fi11, v2)
                                _ = fi11
                                _ = v12
                                _ = v13
                            }
                        case "":
                            for fi14 := 0; fi14 < 16; fi14++ {
                                v15 := GenS2{F0: ("hello" + "a"), F1: (^(v6))}
                                v1 = v4
                                var v16 uint = 2
                                fmt.Println(v6, v6)
                                fmt.Println(fi14, v1, v6)
                                _ = fi14
                                _ = v15
                                _ = v16
                            }
                        case "hello":
                            fmt.Println(v4, v6)
                        default:
                            var v17 uint = 42
                            _ = v17
                        }
                        v18 := []int{fi10}
                        v1 = v1
                        _ = v18
                    } else {
                        var v19 int = (fi10 ^ fi10)
                        _ = v19
                    }
                } else {
                    v6 -= 271
                }
                _ = fi10
            }
            var v20 float32 = 1.5
            _ = v20
        }
    }
    if (32767 >= 687) {
        v21 := GenS1{F0: 501}
        var v22 int = (^(v6))
        _ = v21
        _ = v22
    }
    v4 = v4
    v3 = (v3 >> ((183) & 63))
    v23 := GenS3{F0: 1, F1: 244, F2: 304}
    fmt.Println(v4)
    switch 32767 {
    case 968:
        v24 := []int{v6, (v6 + 4294967295), v6, v6}
        _ = v24
    default:
        var v25 int32 = 0
        _ = v25
    }
    var v26 uint8 = 150
    v27 := GenS3{F0: 336, F1: 0, F2: 18446744073709551615}
    fmt.Println(v26)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v6
}
