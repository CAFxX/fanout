package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 uint8; F2 string }
func genFn0(p0 uint32) int8 {
    v5 := GenS1{F0: 0}
    fmt.Println(p0)
    fmt.Println(p0)
    p0 += (p0 ^ p0)
    switch 452 {
    case 1:
        fmt.Println(p0)
    case 857:
        var v6 uint8 = 255
        _ = v6
    case 88:
        switch 42 {
        case 2147483647:
            fmt.Println(p0)
        case 127:
            var v7 int16 = 1000
            _ = v7
        }
    }
    _ = p0
    _ = v5
    return -75
}
func genFn1(p0 string) string {
    p0 = p0
    var v8 int16 = 42
    p0 = ""
    var v9 uint = 65535
    _ = p0
    _ = v8
    _ = v9
    return ("foo bar" + ("\n\t" + "\n\t"))
}

func main() {
    var v1 int32 = 1
    v2 := make(map[string]bool)
    var v3 int32 = v1
    var v4 uint32 = 1
    v2[((("\n\t" + (("hello" + "foo bar") + "hello")) + "\n\t") + "")] = ((("foo bar" + "hello") + "") != ("foo bar" + "hello"))
    if false {
        v1 = 0
        v3 = 0
    } else {
        v1 += v3
    }
    var v10 uint8 = 42
    v3 += (v3 | (v3 * v3))
    fmt.Println(v10)
    if (42 < -1) {
        v11 := []int{127, 0, -67}
        v12 := GenS1{F0: 542}
        v3 = v3
        _ = v11
        _ = v12
    } else {
        var v13 uint = uint(v4)
        v3 -= (v3 | -2147483648)
        _ = v13
    }
    switch 32768 {
    case -692:
        for fi14 := 0; fi14 < 20; fi14++ {
            for fi15 := 0; fi15 < 8; fi15++ {
                fmt.Println(fi15, v10, v1)
                var v16 int16 = 32767
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
    case -195:
        v3 = v3
    default:
        if (true && ("a" <= "hello")) {
            v17 := GenS1{F0: 252}
            _ = v17
        } else {
            var v18 uint64 = uint64(v10)
            var v19 int32 = v1
            _ = v18
            _ = v19
        }
    }
    v20 := GenS2{F0: 1, F1: v10, F2: ""}
    var v21 uint = 18446744073709551615
    var v22 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
