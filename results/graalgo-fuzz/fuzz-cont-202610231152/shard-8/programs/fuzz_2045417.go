package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 int64; F1 bool; F2 uint }
func genFn0(p0 string) uint {
    v5 := []int{42, -2147483648}
    if ((!(0 == 1)) && (2 > 32767)) {
        fmt.Println(p0)
    }
    _ = len(v5)
    p0 = ("a" + p0)
    _ = p0
    _ = v5
    return 42
}

func main() {
    var v1 bool = (!true)
    var v2 string = "foo bar"
    v3 := make(map[uint32]int)
    v4 := []int{4294967295}
    v6 := GenS1{F0: 65535, F1: -695}
    v7 := GenS2{F0: 0, F1: 127}
    if (4294967295 < -9223372036854775808) {
        var v8 uint = genFn0("")
        fmt.Println(v1, v8, v8)
        var v9 string = v2
        _ = v8
        _ = v9
    }
    switch -1 {
    case 42:
        var v10 string = v2
        _ = v10
    case 128:
        var v11 int16 = 1000
        _ = v11
    case -2147483648:
        switch "\n\t" {
        case "":
            v1 = false
        case "foo bar":
            v12 := GenS2{F0: -554, F1: 0}
            _ = v12
        default:
            if len(v4) > 0 {
                v4[0] = 854
            }
        }
    }
    v3[1] = 0
    for fi13 := 0; fi13 < 14; fi13++ {
        fmt.Println(v1, v1)
        _ = fi13
    }
    if v1 {
        for fi14 := 0; fi14 < 15; fi14++ {
            v15 := GenS1{F0: 4294967295, F1: -9223372036854775808}
            _ = fi14
            _ = v15
        }
        fmt.Println(v2, v2)
        delete(v3, 0)
    }
    var v16 uint16 = 458
    v1 = v1
    switch v2 {
    case "foo bar":
        for fi17 := 0; fi17 < 5; fi17++ {
            fmt.Println(v16)
            v16 += (v16 << ((4294967295) & 63))
            _ = fi17
        }
    case "hello":
        fmt.Println(v2)
    case "a":
        for fi18 := 0; fi18 < 18; fi18++ {
            v19 := GenS3{F0: 858, F1: (!v1), F2: 665}
            _ = fi18
            _ = v19
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
