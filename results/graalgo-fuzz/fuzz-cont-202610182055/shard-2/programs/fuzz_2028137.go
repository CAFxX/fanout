package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
func genFn0() uint {
    if false {
        var v6 int = -400
        _ = v6
    }
    var v7 int64 = -1
    if true {
        if ((((!false) && true) || false) || (!(true || true))) {
            v8 := []int{4294967295}
            _ = v8
        } else {
            v7 -= (v7 / (((v7 & (v7 / (((v7 << ((2) & 63))) | 1)))) | 1))
        }
        var v9 string = ("a" + ("a" + ("foo bar" + "hello")))
        _ = v9
    } else {
        v7 -= -624
        v7 = (v7 >> ((65535) & 63))
        v7 = v7
    }
    v10 := GenS1{F0: false, F1: 65535}
    switch "hello" {
    case "foo bar":
        var v11 int32 = 273
        _ = v11
    case "hello":
        v12 := []int{0, -2147483648, 667}
        _ = v12
    case "\n\t":
        if false {
            v7 -= v7
        } else {
            v13 := []int{0, 127, -343}
            var v14 int8 = 102
            v7 = 0
            _ = v13
            _ = v14
        }
    }
    _ = v10
    _ = v7
    return 831
}

func main() {
    var v1 uint64 = 42
    var v2 int16 = 420
    var v3 uint32 = 0
    v4 := []uint32{}
    var v5 uint16 = 400
    if len(v4) > 0 {
        v4[0] = (v3 / ((v3) | 1))
    }
    for fi15 := 0; fi15 < 9; fi15++ {
        if (545 == -702) {
            var v16 float32 = 1.5
            _ = v16
        }
        _ = fi15
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        _ = len(v4)
        var v18 string = ("" + "\n\t")
        _ = fi17
        _ = v18
    }
    var v19 float32 = (float32((0.0 - ((-2.25 - (float64(1e10) * (3.14159 * 1e10))) - 3.14159))) * 0.0)
    v19 -= 1e10
    v20 := GenS1{F0: (!(((131 < 128) && false) || false)), F1: (^(v5))}
    v2 -= v2
    v2 = v2
    v1 = (v1 << ((1) & 63))
    var v21 uint32 = v3
    var v22 uint = genFn0()
    var v23 bool = false
    v23 = false
    fmt.Println(v3)
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
