package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
func genFn0(p0 int64, p1 int32) uint64 {
    var v5 string = ("foo bar" + ((("foo bar" + ("" + (("foo bar" + ("foo bar" + "\n\t")) + ""))) + ("hello" + "a")) + ""))
    if false {
        v6 := []int{-1, 486}
        fmt.Println(p0, p0, v5)
        var v7 uint32 = 2147483647
        _ = v6
        _ = v7
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v5
    return 2
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint32 = 2147483648
    var v3 uint32 = (v2 >> ((18446744073709551615) & 63))
    var v4 uint32 = (v2 << ((65535) & 63))
    if (!(32767 <= 1000)) {
        _ = len(v1)
    }
    if (("\n\t" + (("hello" + ("hello" + "")) + "")) > "") {
        var v8 int16 = 0
        _ = v8
    }
    v4 = v3
    switch "a" {
    case "\n\t":
        var v9 int8 = -128
        _ = v9
    default:
        v3 = 69
    }
    v4 += 278
    v10 := GenS1{F0: 79, F1: false}
    switch "foo bar" {
    case "hello":
        v11 := GenS1{F0: 132, F1: (uint8(v2) > 127)}
        _ = v11
    case "a":
        var v12 uint32 = (v2 - v4)
        _ = v12
    case "foo bar":
        delete(v1, -1)
    default:
        if false {
            v2 -= v2
            for fi13 := 0; fi13 < 12; fi13++ {
                if (-128 <= 100) {
                    v3 += v4
                    var v14 int16 = 1
                    var v15 int64 = -1
                    _ = v14
                    _ = v15
                }
                fmt.Println(v2)
                v2 += (v3 * v2)
                _ = fi13
            }
        }
    }
    v16 := GenS1{F0: 127, F1: (!true)}
    for fi17 := 0; fi17 < 13; fi17++ {
        v18 := []int{}
        switch (fi17 / ((fi17) | 1)) {
        case 2147483647:
            delete(v1, fi17)
        case 42:
            for fi19 := 0; fi19 < 7; fi19++ {
                fmt.Println(fi17, fi19, v3)
                _ = fi19
            }
        case -9223372036854775808:
            v4 = 2147483648
        default:
            var v20 float32 = float32(1e10)
            _ = v20
        }
        _ = fi17
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
