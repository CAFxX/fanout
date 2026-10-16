package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 bool }
func genFn0(p0 int16, p1 int16) uint {
    fmt.Println(p1, p1)
    v6 := GenS1{F0: 65535}
    _ = p0
    _ = p1
    _ = v6
    return 42
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]int)
    v3 := []string{("hello" + "foo bar"), "hello", "\n\t"}
    var v4 int8 = -1
    v5 := make(map[uint32]int)
    switch "" {
    case "":
        if len(v3) > 0 {
            v3[0] = (("foo bar" + ("\n\t" + (("hello" + "hello") + ("a" + ("hello" + "\n\t"))))) + ("a" + "a"))
        }
    case "hello":
        v4 += v4
    case "foo bar":
        var v7 float32 = float32(float64((((-2.25 - 1e10) * 1.5) / (((0.0 - float32(3.14159)) * 1e10) / -2.25))))
        _ = v7
    default:
        v8 := GenS2{F0: 0}
        _ = v8
    }
    switch "a" {
    case "":
        switch 32767 {
        case 2147483647:
            v4 += v4
        default:
            var v9 uint64 = 1
            _ = v9
        }
    case "foo bar":
        switch "\n\t" {
        case "\n\t":
            v4 = v4
        case "foo bar":
            _ = len(v5)
        }
    default:
        var v10 int8 = 55
        _ = v10
    }
    v5[370] = -2147483648
    var v11 float64 = ((1.5 + -2.25) / 1.0)
    v4 = v4
    v12 := GenS2{F0: 32767}
    if (42 <= 112) {
        v4 -= (v4 ^ (v4 + v4))
    }
    switch 32768 {
    case -856:
        v13 := GenS2{F0: -32768}
        _ = v13
    case 32767:
        var v14 uint32 = 1
        _ = v14
    }
    v15 := GenS2{F0: 722}
    v16 := GenS3{F0: 18446744073709551615, F1: 317}
    v4 = v4
    v11 = v11
    var v17 uint16 = 65535
    if (981 <= 1) {
        if len(v3) > 0 {
            v3[0] = ("" + ("hello" + (("\n\t" + "\n\t") + "foo bar")))
        }
        v17 -= (^(v17))
        v18 := GenS4{F0: false}
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
