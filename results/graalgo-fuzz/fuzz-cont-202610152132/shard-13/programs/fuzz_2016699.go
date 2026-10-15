package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 int16; F1 bool; F2 uint8 }
type GenS4 struct { F0 bool; F1 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int64; F1 int16; F2 int64 }
func genFn0(p0 uint16, p1 int32) uint {
    p1 = (p1 / ((p1) | 1))
    var v5 uint64 = 551
    p1 += (-(p1))
    v5 += 0
    _ = p0
    _ = p1
    _ = v5
    return 711
}
func genFn1() int64 {
    v6 := GenS1{F0: -128}
    var v7 float64 = (1e10 * 1.5)
    _ = v6
    _ = v7
    return -530
}

func main() {
    v1 := []string{((("foo bar" + ("hello" + "hello")) + "hello") + ("" + (((("hello" + ("hello" + "a")) + "") + "a") + "foo bar"))), "", "a"}
    var v2 bool = (true && true)
    var v3 int64 = -9223372036854775808
    var v4 uint16 = 546
    v8 := GenS2{F0: -9223372036854775808, F1: 507}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v9 := GenS3{F0: -912, F1: (((!true) || v2) && false), F2: 128}
    v1 = append(v1, (("\n\t" + (("\n\t" + (("" + "a") + "hello")) + ("foo bar" + "\n\t"))) + (("foo bar" + "a") + "")))
    var v10 uint = uint(v4)
    for fi11 := 0; fi11 < 5; fi11++ {
        for fi12 := 0; fi12 < 15; fi12++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v13 := GenS4{F0: v2, F1: 23}
            v3 -= -9
            switch ("foo bar" + "hello") {
            case "foo bar":
                switch fi11 {
                case -90:
                    fmt.Println(v10)
                default:
                    v4 += 255
                }
            }
            var v14 int = (fi11 * fi12)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi11
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v1 = append(v1, "foo bar")
    v2 = false
    switch "a" {
    case "a":
        v15 := GenS5{F0: 128}
        _ = v15
    default:
        v16 := GenS6{F0: v3, F1: 42, F2: v3}
        _ = v16
    }
    v10 += genFn0((v4 | (v4 * v4)), -2147483648)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
