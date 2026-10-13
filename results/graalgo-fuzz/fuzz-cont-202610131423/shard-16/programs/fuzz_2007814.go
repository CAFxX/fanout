package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int64 }
type GenS2 struct { F0 bool; F1 uint; F2 uint64 }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 string; F1 int32 }

func main() {
    var v1 float64 = 0.0
    v2 := []string{}
    v3 := []uint32{0, 4294967295}
    v1 += ((1.5 / v1) * (((1e10 / v1) + -2.25) / ((v1 + v1) * (float64(-2.25) * float64(1e10)))))
    if false {
        var v4 float64 = float64(float32(((1.5 * 1e10) * (1e10 + ((3.14159 / 1.0) * float64(0.1))))))
        _ = v4
    } else {
        var v5 int64 = 1
        _ = v5
    }
    fmt.Println(v1)
    v1 += (v1 - v1)
    fmt.Println(v1)
    fmt.Println(v1)
    v6 := GenS1{F0: 42, F1: 42, F2: -1}
    switch "hello" {
    case "foo bar":
        if ((!false) && ((!true) && false)) {
            var v7 float64 = v1
            _ = v7
        } else {
            switch "hello" {
            case "hello":
                switch "\n\t" {
                case "foo bar":
                    v1 = v1
                case "\n\t":
                    v1 = v1
                default:
                    var v8 float64 = (v1 - 3.14159)
                    _ = v8
                }
            case "":
                v1 = (v1 + v1)
            case "a":
                var v9 uint16 = 2
                _ = v9
            }
        }
    default:
        v10 := GenS2{F0: (float32(1e10) <= -2.25), F1: 65535, F2: 18446744073709551615}
        _ = v10
    }
    v2 = append(v2, (("hello" + (("a" + "a") + "hello")) + ""))
    switch 0 {
    case -517:
        if len(v2) > 0 {
            v2[0] = (((("foo bar" + "\n\t") + "hello") + "hello") + "")
        }
    case 42:
        v2 = append(v2, "foo bar")
    default:
        v3 = append(v3, 2147483647)
    }
    v11 := GenS3{F0: 1, F1: 459}
    v12 := GenS4{F0: ("a" + ("a" + ("a" + ("foo bar" + "a")))), F1: -2147483648}
    var v13 float32 = float32(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v6
}
