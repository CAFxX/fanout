package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint64 }
type GenS2 struct { F0 bool }
func genFn0(p0 uint32, p1 uint) bool {
    switch "foo bar" {
    case "":
        switch 42 {
        case 42:
            var v5 int32 = 0
            _ = v5
        case 732:
            v6 := GenS1{F0: 889, F1: false, F2: 553}
            _ = v6
        case 127:
            fmt.Println(p0, p0)
        default:
            var v7 int8 = 1
            _ = v7
        }
    case "hello":
        for fi8 := 0; fi8 < 7; fi8++ {
            if true {
                v9 := GenS2{F0: ((0.0 == (0.0 - 0.1)) && true)}
                _ = v9
            } else {
                fmt.Println(p1, p1)
                fmt.Println(p1)
            }
            var v10 int16 = 21
            var v11 uint64 = 1
            _ = fi8
            _ = v10
            _ = v11
        }
    default:
        for fi12 := 0; fi12 < 1; fi12++ {
            var v13 uint8 = 255
            switch (-(fi12)) {
            case 128:
                var v14 uint = 1
                _ = v14
            case 127:
                if true {
                    v15 := GenS2{F0: ((true || true) && true)}
                    _ = v15
                }
            case -433:
                p0 = (p0 + 1)
            }
            _ = fi12
            _ = v13
        }
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        if false {
            v17 := []int{32768}
            var v18 uint32 = 2
            _ = v17
            _ = v18
        } else {
            p1 = 733
        }
        var v19 string = "hello"
        _ = fi16
        _ = v19
    }
    v20 := GenS1{F0: 42, F1: true, F2: 226}
    v21 := []int{257}
    _ = p0
    _ = p1
    _ = v20
    _ = v21
    return false
}

func main() {
    var v1 bool = (555 < 374)
    var v2 uint16 = 2
    v3 := make(map[int]string)
    var v4 bool = v1
    v22 := []int{127}
    if false {
        v23 := GenS1{F0: 0, F1: (1 > 0), F2: 926}
        _ = v23
    }
    var v24 float64 = (1e10 / 3.14159)
    var v25 uint8 = 2
    switch 32767 {
    case 32767:
        if false {
            var v26 int32 = 2147483647
            for fi27 := 0; fi27 < 9; fi27++ {
                v24 = v24
                _ = fi27
            }
            switch "hello" {
            case "\n\t":
                v28 := GenS2{F0: (4294967295 <= 1)}
                _ = v28
            case "a":
                v4 = v1
            default:
                var v29 string = (("" + "foo bar") + ("a" + ("hello" + "a")))
                _ = v29
            }
            _ = v26
        }
    case 127:
        switch "a" {
        case "hello":
            var v30 uint16 = (v2 ^ (v2 >> ((310) & 63)))
            _ = v30
        }
    }
    var v31 int16 = 1000
    v31 = (v31 * 1000)
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v4
}
