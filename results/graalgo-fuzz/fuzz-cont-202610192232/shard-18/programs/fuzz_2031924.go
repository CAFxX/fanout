package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
func genFn0(p0 uint, p1 int8) uint {
    var v4 int8 = p1
    switch 9223372036854775807 {
    case -874:
        var v5 int8 = (p1 << ((827) & 63))
        _ = v5
    default:
        switch "hello" {
        case "hello":
            fmt.Println(p0, v4)
        case "a":
            fmt.Println(v4, v4)
        default:
            v6 := GenS1{F0: (p0 % (((p0 % (((p0 << (((p0 + 18446744073709551615)) & 63))) | 1))) | 1)), F1: (2147483647 < 0)}
            _ = v6
        }
    }
    _ = p0
    _ = p1
    _ = v4
    return (^(p0))
}

func main() {
    var v1 uint16 = 32767
    var v2 uint64 = 57
    v3 := make(map[string]uint32)
    switch 0 {
    case 0:
        v7 := GenS1{F0: genFn0(4294967295, 1), F1: true}
        _ = v7
    }
    if (float32(((((((-2.25 + (1.5 * (float64(-2.25) * 0.1))) - ((3.14159 / (1e10 - (3.14159 * 0.0))) - 0.1)) + float64(3.14159)) / 1.0) * (0.1 - 1.5)) * -2.25)) < 0.0) {
        v8 := GenS1{F0: 42, F1: false}
        if (32768 != 32768) {
            for fi9 := 0; fi9 < 17; fi9++ {
                _ = len(v3)
                _ = fi9
            }
            if (true || false) {
                v10 := []int{782, 9223372036854775807, -272}
                _ = v10
            }
        }
        switch "foo bar" {
        case "\n\t":
            v11 := GenS1{F0: 781, F1: (float64((1e10 + (0.1 * (0.0 + ((0.0 / 1.0) - float32(float64(float32((((1.5 - (float64((1.5 / 0.1)) / -2.25)) * -2.25) * -2.25))))))))) >= (-2.25 / 3.14159))}
            _ = v11
        }
        _ = v8
    }
    v12 := GenS1{F0: 0, F1: false}
    v13 := []int{42, 32767, 2}
    v14 := GenS1{F0: 18446744073709551615, F1: false}
    if (42 > 890) {
        delete(v3, ("a" + ("a" + (("a" + "hello") + "hello"))))
    }
    v13 = append(v13, 1)
    v1 = 1
    if len(v13) > 0 {
        v13[0] = 32767
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        fmt.Println(v1)
        fmt.Println(v1, fi15)
        _ = fi15
    }
    fmt.Println(v1)
    var v16 string = (("hello" + "a") + "")
    var v17 int16 = -32768
    fmt.Println(v2, v17, v16)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
