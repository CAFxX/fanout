package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int }
func genFn0(p0 uint8, p1 uint32) int {
    var v4 int64 = 42
    v5 := GenS1{F0: -10}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -24
}

func main() {
    v1 := []uint32{2}
    var v2 float64 = 1.5
    var v3 uint64 = 128
    var v6 uint8 = 2
    var v7 uint = 414
    if true {
        var v8 int32 = 1
        v3 -= 2
        _ = v8
    } else {
        var v9 int64 = 42
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        v2 -= v2
        if false {
            fmt.Println(fi10)
            var v11 float64 = 0.0
            v12 := GenS1{F0: -37}
            _ = v11
            _ = v12
        } else {
            var v13 string = "a"
            _ = v13
        }
        for fi14 := 0; fi14 < 3; fi14++ {
            var v15 int8 = 34
            _ = fi14
            _ = v15
        }
        _ = fi10
    }
    v2 = float64((3.14159 - 1e10))
    v16 := GenS2{F0: -89}
    for fi17 := 0; fi17 < 6; fi17++ {
        var v18 int8 = 0
        _ = fi17
        _ = v18
    }
    var v19 int32 = -1
    for fi20 := 0; fi20 < 12; fi20++ {
        if false {
            v3 = (^(v3))
            switch ("hello" + "") {
            case "foo bar":
                if ("\n\t" >= "hello") {
                    fmt.Println(v3, v3)
                } else {
                    v2 -= -2.25
                    v21 := GenS1{F0: 48}
                    switch "a" {
                    case "":
                        fmt.Println(v19, v7, v3)
                    case "foo bar":
                        v6 = v6
                    }
                    _ = v21
                }
            default:
                _ = len(v1)
            }
            v3 -= v3
        } else {
            v2 = (3.14159 + 3.14159)
            var v22 uint32 = 843
            _ = v22
        }
        var v23 int8 = 0
        if len(v1) > 0 {
            v1[0] = 516
        }
        _ = fi20
        _ = v23
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v24 uint64 = v3
    var v25 int32 = 42
    _ = len(v1)
    var v26 uint = 65535
    var v27 int32 = v25
    var v28 int16 = 42
    v29 := GenS1{F0: -128}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v6
    _ = v7
}
