package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 string; F1 uint32; F2 int32 }
type GenS3 struct { F0 int8; F1 int }
func genFn0() string {
    var v5 float64 = -2.25
    v6 := GenS1{F0: 42, F1: 1}
    var v7 string = ((((("" + ("" + "a")) + "") + ("a" + "a")) + "foo bar") + (("\n\t" + "hello") + "hello"))
    v5 += 3.14159
    _ = v5
    _ = v6
    _ = v7
    return v7
}
func genFn1(p0 int64, p1 int) uint32 {
    for fi8 := 0; fi8 < 11; fi8++ {
        p0 -= 0
        v9 := GenS1{F0: 42, F1: -848}
        _ = fi8
        _ = v9
    }
    var v10 bool = true
    for fi11 := 0; fi11 < 12; fi11++ {
        p1 = (p1 & (-(fi11)))
        v10 = (938 > 0)
        _ = fi11
    }
    var v12 uint = 255
    _ = p0
    _ = p1
    _ = v10
    _ = v12
    return 881
}

func main() {
    v1 := []string{}
    v2 := make(map[int]bool)
    var v3 uint = 800
    var v4 float32 = 0.1
    switch ("hello" + "") {
    case "foo bar":
        var v13 bool = false
        _ = v13
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        fmt.Println(v3)
        v15 := GenS1{F0: 2, F1: 1}
        _ = fi14
        _ = v15
    }
    switch 1 {
    case -317:
        delete(v2, 1)
    case 128:
        fmt.Println(v3, v3)
    case 4294967295:
        v4 -= (1e10 * v4)
    default:
        var v16 uint = (v3 - v3)
        _ = v16
    }
    var v17 int8 = -128
    for fi18 := 0; fi18 < 3; fi18++ {
        v19 := GenS2{F0: "", F1: 42, F2: -1}
        _ = fi18
        _ = v19
    }
    if false {
        switch 0 {
        case -938:
            v20 := GenS3{F0: (v17 * v17), F1: 1}
            _ = v20
        }
    }
    var v21 uint32 = 1
    v2[int(v21)] = (!(2 == 255))
    var v22 uint16 = 42
    v3 = 1
    switch ("\n\t" + "\n\t") {
    case "\n\t":
        v23 := GenS1{F0: 0, F1: 42}
        _ = v23
    case "foo bar":
        if (true && true) {
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
            fmt.Println(v4)
            v22 += (v22 % ((0) | 1))
        }
    default:
        v21 = (v21 % ((911) | 1))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
