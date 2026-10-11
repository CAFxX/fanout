package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 bool }
func genFn0(p0 int64, p1 uint, p2 int64) int8 {
    v4 := []int{-672, 32767, 127, 302}
    var v5 float64 = (0.0 * 1e10)
    fmt.Println(p2, p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 100
}
func genFn1(p0 uint, p1 int32) string {
    fmt.Println(p0, p0)
    p0 -= p0
    var v6 float32 = ((1e10 - (0.0 / 1e10)) * (1e10 + float32(float64(1.5))))
    if ("" == (("\n\t" + "") + "a")) {
        v6 -= v6
        v7 := []int{9223372036854775807}
        fmt.Println(p0, p1)
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v6
    return "hello"
}

func main() {
    var v1 uint8 = 5
    v2 := make(map[string]string)
    var v3 string = ("hello" + "\n\t")
    v8 := GenS1{F0: 1, F1: 65535}
    for fi9 := 0; fi9 < 11; fi9++ {
        v1 += (^(v1))
        _ = fi9
    }
    var v10 bool = true
    switch -9223372036854775808 {
    case 0:
        fmt.Println(v10, v3, v3)
    case 615:
        var v11 int64 = -64
        _ = v11
    }
    v3 = v3
    var v12 int16 = 42
    v12 += 0
    switch -1 {
    case -2147483648:
        v13 := GenS2{F0: -1, F1: ((("\n\t" + v3) + "") + v3)}
        _ = v13
    default:
        for fi14 := 0; fi14 < 13; fi14++ {
            var v15 bool = ((!((42 == 278) && ((2 == 1) || false))) || (false && v10))
            for fi16 := 0; fi16 < 18; fi16++ {
                switch "hello" {
                case "hello":
                    v17 := v2[(v3 + (v3 + "foo bar"))]
                    _ = v17
                case "a":
                    v2[(v3 + "hello")] = ""
                }
                v12 += 42
                _ = fi16
            }
            _ = fi14
            _ = v15
        }
    }
    var v18 int = -100
    if true {
        v19 := GenS3{F0: v10}
        _ = v19
    }
    v3 = ((v3 + "foo bar") + v3)
    var v20 int32 = 90
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v12
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v8
}
