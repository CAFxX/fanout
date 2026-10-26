package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int64, p1 int64) string {
    var v5 uint32 = 79
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 float64 = (float64(-2.25) + (-2.25 / 1.5))
        _ = fi6
        _ = v7
    }
    v8 := []int{-447, 4294967295}
    v9 := GenS1{F0: 65535, F1: 128}
    if false {
        p1 -= p1
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v8
    _ = v9
    return ""
}

func main() {
    var v1 uint16 = 0
    var v2 uint32 = 65535
    var v3 int16 = -32768
    var v4 float32 = ((-2.25 + 1.5) * ((1.5 * (0.0 + 0.1)) / (1.5 / (float32(0.0) - (0.0 - (1.5 * float32(3.14159)))))))
    v4 = float32(1.5)
    v2 = (v2 ^ 2147483647)
    var v10 string = "foo bar"
    if ((("foo bar" + ("hello" + v10)) <= (("hello" + v10) + "")) || false) {
        if false {
            v11 := []int{742, -9223372036854775808, 32767, 2147483647}
            if (32768 <= 1) {
                v12 := GenS2{F0: 18446744073709551615}
                v10 = (v10 + ("hello" + ((v10 + ((v10 + v10) + (("a" + (v10 + ((v10 + v10) + "foo bar"))) + (v10 + v10)))) + v10)))
                _ = v12
            }
            _ = v11
        }
        v1 -= v1
        v2 += 2147483648
    }
    var v13 uint32 = (v2 & (v2 | v2))
    if false {
        v13 = v13
        if ((v3 / (((v3 - v3)) | 1)) > v3) {
            if false {
                for fi14 := 0; fi14 < 8; fi14++ {
                    switch ("foo bar" + "\n\t") {
                    case "":
                        fmt.Println(v10, v3)
                    case "hello":
                        var v15 uint = 2
                        _ = v15
                    default:
                        switch "\n\t" {
                        case "a":
                            v13 = 4294967295
                        case "\n\t":
                            var v16 string = (v10 + "\n\t")
                            _ = v16
                        case "foo bar":
                            v2 = 884
                        default:
                            var v17 string = "foo bar"
                            _ = v17
                        }
                    }
                    fmt.Println(v4)
                    _ = fi14
                }
            } else {
                var v18 int8 = 1
                _ = v18
            }
            fmt.Println(v2)
            v2 = (v13 * 4294967295)
        }
    }
    if false {
        var v19 uint8 = 2
        var v20 uint8 = v19
        fmt.Println(v4)
        _ = v19
        _ = v20
    } else {
        if (uint8(v13) < uint8(v1)) {
            if false {
                v2 += (^(v2))
            }
        }
    }
    var v21 int32 = 245
    var v22 uint = 4294967295
    v23 := GenS1{F0: 436, F1: 184}
    v4 = ((v4 - v4) * 0.1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
