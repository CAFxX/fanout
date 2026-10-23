package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int8; F1 uint16; F2 uint64 }
func genFn0(p0 int32) bool {
    if true {
        p0 = (p0 - p0)
        v6 := GenS1{F0: 42, F1: (18446744073709551615 < 690)}
        _ = v6
    }
    v7 := []int{32767, 32767}
    v7 = append(v7, -552)
    _ = p0
    _ = v7
    return ((1e10 != (-2.25 * 1e10)) || (true || false))
}
func genFn1(p0 uint16, p1 int64, p2 int16) uint64 {
    p0 = p0
    var v8 float32 = 1e10
    v9 := []int{-2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 274
}

func main() {
    var v1 int = 132
    v2 := []int{v1}
    var v3 uint8 = 71
    v4 := make(map[int]bool)
    v5 := make(map[string]bool)
    switch v1 {
    case 4294967295:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case -858:
        fmt.Println(v3)
    }
    switch v1 {
    case -283:
        switch "hello" {
        case "hello":
            var v10 int = -1
            _ = v10
        case "":
            if true {
                var v11 uint16 = 1
                fmt.Println(v1, v11)
                if true {
                    v1 = (v1 + v1)
                    v1 -= v1
                }
                _ = v11
            }
        default:
            var v12 string = "\n\t"
            _ = v12
        }
    case 1:
        for fi13 := 0; fi13 < 5; fi13++ {
            if false {
                for fi14 := 0; fi14 < 9; fi14++ {
                    var v15 uint16 = 548
                    v16 := GenS1{F0: fi14, F1: true}
                    v1 -= (fi13 & v1)
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            }
            _ = fi13
        }
    case -1:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        fmt.Println(v1)
    }
    fmt.Println(v1, v1)
    v3 -= v3
    var v17 uint8 = 255
    for fi18 := 0; fi18 < 20; fi18++ {
        var v19 uint64 = 317
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        fmt.Println(v17, v3)
        _ = fi18
        _ = v19
    }
    delete(v4, 128)
    if true {
        v17 += 42
        switch (-(v1)) {
        case 1:
            var v20 string = "foo bar"
            _ = v20
        case 0:
            v17 -= 42
        case 42:
            for fi21 := 0; fi21 < 12; fi21++ {
                v17 += v3
                v2 = append(v2, fi21)
                _ = fi21
            }
        }
    } else {
        v1 += 0
        fmt.Println(v17)
        for fi22 := 0; fi22 < 19; fi22++ {
            v23 := GenS2{F0: -31, F1: 685, F2: 692}
            _ = fi22
            _ = v23
        }
    }
    var v24 float32 = 1e10
    v24 -= v24
    fmt.Println(v17, v3)
    for fi25 := 0; fi25 < 10; fi25++ {
        for fi26 := 0; fi26 < 5; fi26++ {
            var v27 int8 = 109
            var v28 string = (("hello" + "") + "hello")
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = fi25
    }
    var v29 int8 = 0
    var v30 float32 = (((v24 / (v24 / (v24 / ((3.14159 / 1.0) / 1e10)))) - float32(0.1)) - ((3.14159 / v24) + (v24 + 0.1)))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
