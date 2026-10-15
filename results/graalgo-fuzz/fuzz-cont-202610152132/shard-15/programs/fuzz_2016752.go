package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint32 }
func genFn0() uint {
    var v4 string = "a"
    v4 = v4
    v4 = ("a" + v4)
    v4 = "a"
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 uint = 42
        v6 += v6
        _ = fi5
        _ = v6
    }
    _ = v4
    return 4294967295
}
func genFn1() uint64 {
    switch ("" + "a") {
    case "foo bar":
        var v7 int32 = 0
        _ = v7
    case "a":
        var v8 uint64 = 21
        _ = v8
    case "hello":
        var v9 int16 = 770
        _ = v9
    default:
        var v10 int32 = 2147483647
        _ = v10
    }
    v11 := []int{-2, 4294967295, 0, 32768}
    if (((!true) || ((127 == 127) || (!true))) && true) {
        switch -9223372036854775808 {
        case -2147483648:
            var v12 uint32 = 776
            _ = v12
        default:
            v13 := GenS1{F0: 2147483647, F1: 255}
            _ = v13
        }
        switch (((("hello" + (("hello" + "a") + "hello")) + "hello") + ("a" + "a")) + (("foo bar" + "") + (("" + "\n\t") + "a"))) {
        case "hello":
            var v14 int64 = 0
            _ = v14
        }
    }
    var v15 uint32 = 278
    _ = v11
    _ = v15
    return 0
}

func main() {
    var v1 int32 = -633
    var v2 uint8 = 148
    var v3 int8 = -97
    v16 := []int{981, -544}
    switch -453 {
    case 2147483647:
        v3 -= v3
    case 32767:
        var v17 uint8 = 2
        _ = v17
    default:
        v18 := GenS1{F0: 2147483647, F1: genFn0()}
        _ = v18
    }
    v16 = append(v16, 32767)
    if len(v16) > 0 {
        v16[0] = 42
    }
    v1 += v1
    var v19 uint16 = 255
    for fi20 := 0; fi20 < 14; fi20++ {
        var v21 uint8 = v2
        _ = fi20
        _ = v21
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        v1 += (v1 >> ((951) & 63))
        v1 -= v1
        _ = fi22
    }
    if (false || true) {
        v16 = append(v16, -381)
        fmt.Println(v1, v2)
        v3 = v3
    }
    if true {
        v23 := GenS1{F0: 128, F1: 65535}
        _ = len(v16)
        var v24 int = -77
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 1; fi25++ {
        for fi26 := 0; fi26 < 1; fi26++ {
            if true {
                v27 := GenS1{F0: (fi26 * int(v2)), F1: 65535}
                v28 := GenS2{F0: 2}
                fmt.Println(v19, v19)
                switch "hello" {
                case "":
                    v19 -= (v19 % ((v19) | 1))
                }
                _ = v27
                _ = v28
            }
            _ = fi26
        }
        v19 -= v19
        _ = len(v16)
        _ = fi25
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v3
}
