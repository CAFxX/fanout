package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint8; F1 int; F2 int64 }
func genFn0(p0 uint16, p1 uint) int8 {
    v5 := []int{}
    switch "foo bar" {
    case "hello":
        v6 := GenS1{F0: ((((("foo bar" + "") + "hello") + "hello") + "") + "hello"), F1: -543}
        _ = v6
    case "foo bar":
        var v7 int64 = -1
        _ = v7
    case "a":
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 int = (fi8 / ((626) | 1))
            var v10 int16 = -857
            var v11 bool = false
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v5
    return -7
}
func genFn1(p0 uint16) int16 {
    p0 = (p0 - 1)
    var v12 int16 = 0
    fmt.Println(p0, v12)
    _ = p0
    _ = v12
    return v12
}

func main() {
    var v1 bool = false
    var v2 int8 = 0
    var v3 uint8 = 127
    v4 := make(map[uint32]int)
    switch "\n\t" {
    case "foo bar":
        v1 = v1
    }
    v13 := GenS1{F0: ("\n\t" + "hello"), F1: -84}
    for fi14 := 0; fi14 < 16; fi14++ {
        v3 -= 128
        v15 := GenS2{F0: v3, F1: (fi14 & (fi14 * 2147483647)), F2: 1}
        fmt.Println(v1, fi14)
        _ = fi14
        _ = v15
    }
    v3 = (v3 - (v3 * (v3 + v3)))
    var v16 int64 = -9223372036854775808
    if v1 {
        if (1 >= 42) {
            for fi17 := 0; fi17 < 15; fi17++ {
                v1 = (!(-32768 != -297))
                _ = fi17
            }
            var v18 int32 = 0
            _ = v18
        }
    } else {
        fmt.Println(v16, v16)
        switch 32768 {
        case 128:
            var v19 int = 32768
            _ = v19
        case 596:
            var v20 uint64 = 901
            _ = v20
        }
    }
    switch 32767 {
    case 32768:
        var v21 string = ""
        _ = v21
    case 1:
        v1 = (false && false)
    default:
        v2 = genFn0(1, 255)
    }
    fmt.Println(v2)
    var v22 uint64 = 323
    for fi23 := 0; fi23 < 13; fi23++ {
        switch 7 {
        case 32768:
            v3 = v3
        case 0:
            v3 += (v3 * uint8(v22))
        }
        var v24 int = -20
        _ = fi23
        _ = v24
    }
    v3 += (v3 + 48)
    switch "\n\t" {
    case "a":
        var v25 uint16 = 2
        _ = v25
    case "":
        for fi26 := 0; fi26 < 11; fi26++ {
            _ = len(v4)
            _ = fi26
        }
    case "hello":
        switch "foo bar" {
        case "":
            var v27 int8 = v2
            _ = v27
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
