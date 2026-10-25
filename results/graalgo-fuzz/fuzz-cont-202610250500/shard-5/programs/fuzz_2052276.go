package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 int }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int }
func genFn0() int32 {
    v4 := GenS1{F0: 1, F1: -422, F2: 553}
    switch ("foo bar" + "foo bar") {
    case "foo bar":
        var v5 uint64 = 42
        _ = v5
    default:
        var v6 uint8 = 201
        _ = v6
    }
    _ = v4
    return 42
}
func genFn1(p0 uint16, p1 uint16) int8 {
    fmt.Println(p1, p1)
    v7 := GenS2{F0: 127, F1: 0}
    v8 := GenS1{F0: 801, F1: 32767, F2: -2147483648}
    if (!true) {
        if (92 == 367) {
            for fi9 := 0; fi9 < 1; fi9++ {
                fmt.Println(fi9, p1)
                fmt.Println(p0, p0)
                for fi10 := 0; fi10 < 12; fi10++ {
                    var v11 uint16 = 1
                    _ = fi10
                    _ = v11
                }
                _ = fi9
            }
            v12 := GenS2{F0: 2, F1: -9223372036854775808}
            v13 := GenS3{F0: 411}
            _ = v12
            _ = v13
        } else {
            p0 += p1
            for fi14 := 0; fi14 < 5; fi14++ {
                v15 := []int{(fi14 | (fi14 * fi14)), 905}
                _ = fi14
                _ = v15
            }
        }
    }
    if false {
        p1 -= (p1 + p1)
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 100
}

func main() {
    var v1 string = "foo bar"
    var v2 float32 = (0.0 / 0.1)
    v3 := []uint32{2147483648, 182, 883}
    var v16 float32 = (v2 + 1.5)
    v16 -= v2
    switch 128 {
    case 374:
        v17 := GenS3{F0: 1000}
        _ = v17
    }
    v2 += float32(-2.25)
    var v18 uint = 487
    v16 = v16
    v16 = (float32(3.14159) - float32(-2.25))
    for fi19 := 0; fi19 < 12; fi19++ {
        v20 := GenS4{F0: fi19}
        _ = fi19
        _ = v20
    }
    var v21 string = ((("hello" + v1) + v1) + ("\n\t" + v1))
    switch ((v1 + v1) + ((((v21 + (v21 + v1)) + v21) + (("\n\t" + v21) + "foo bar")) + "hello")) {
    case "hello":
        v3 = append(v3, 289)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v21
    _ = v3
}
