package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 int8 }

func main() {
    var v1 uint32 = 0
    v2 := make(map[uint32]bool)
    var v3 int64 = 1
    var v4 int32 = 0
    v5 := make(map[string]uint32)
    v6 := make(map[string]uint32)
    if false {
        v7 := GenS1{F0: v3, F1: 65535}
        _ = v7
    } else {
        _ = len(v5)
        v8 := GenS1{F0: (-(v3)), F1: 776}
        var v9 int64 = v3
        _ = v8
        _ = v9
    }
    fmt.Println(v1, v3, v3)
    var v10 float32 = ((0.0 / 3.14159) + 0.0)
    if (!(false || false)) {
        v3 -= (v3 >> ((515) & 63))
        v3 = (v3 * 0)
        switch ((("foo bar" + ("a" + "foo bar")) + "foo bar") + "\n\t") {
        case "":
            v10 = v10
        case "hello":
            v1 = (^(v1))
        case "a":
            var v11 float64 = 0.0
            _ = v11
        }
    } else {
        v1 = v1
        for fi12 := 0; fi12 < 7; fi12++ {
            v13 := GenS1{F0: 42, F1: 18446744073709551615}
            var v14 int16 = -32768
            v15 := GenS2{F0: 47}
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
    }
    switch 426 {
    case 4294967295:
        if (32767 >= 42) {
            v16 := GenS1{F0: 9223372036854775807, F1: 4294967295}
            v1 = (v1 / ((65535) | 1))
            fmt.Println(v1, v1)
            _ = v16
        }
    default:
        v3 = (v3 % ((1) | 1))
    }
    v4 = (v4 >> ((1) & 63))
    fmt.Println(v1, v3, v1)
    for fi17 := 0; fi17 < 3; fi17++ {
        if false {
            v4 = (v4 ^ v4)
            v2[(v1 << ((65535) & 63))] = (!true)
        }
        var v18 int8 = 95
        _ = fi17
        _ = v18
    }
    if false {
        v19 := []int{-1, 32767}
        if ((false || (42 >= 18446744073709551615)) && (!((false || (42 == 408)) || true))) {
            var v20 int = 127
            v10 += v10
            v19 = append(v19, 1)
            _ = v20
        } else {
            if false {
                _ = len(v6)
            }
            if len(v19) > 0 {
                v19[0] = -495
            }
            switch -2147483648 {
            case 32767:
                fmt.Println(v10)
            }
        }
        _ = v19
    }
    var v21 uint64 = 1
    v22 := []int{-717, -1}
    fmt.Println(v3)
    var v23 int16 = -32768
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    var v24 float32 = v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
