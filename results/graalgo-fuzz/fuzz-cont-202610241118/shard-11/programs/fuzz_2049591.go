package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int16; F1 int; F2 bool }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []uint32{}
    v3 := make(map[string]string)
    var v4 int8 = 120
    var v5 bool = (!false)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v6 int = 42
    var v7 int16 = -714
    if (("hello" + "a") < "foo bar") {
        v1[42] = v5
    }
    v8 := GenS1{F0: v4, F1: 1}
    for fi9 := 0; fi9 < 18; fi9++ {
        fmt.Println(v6, fi9, v4)
        if v5 {
            v5 = ((!v5) || v5)
        }
        switch "hello" {
        case "hello":
            if true {
                for fi10 := 0; fi10 < 2; fi10++ {
                    var v11 string = ""
                    v12 := GenS1{F0: (v4 ^ v4), F1: 255}
                    v2 = append(v2, 887)
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                var v13 uint32 = 809
                _ = v13
            }
        default:
            for fi14 := 0; fi14 < 17; fi14++ {
                v5 = (int64(fi9) > 42)
                v15 := GenS2{F0: 32767, F1: (^(fi14)), F2: (!v5)}
                fmt.Println(v4, v5)
                v7 -= 767
                v2 = append(v2, 562)
                _ = fi14
                _ = v15
            }
        }
        _ = fi9
    }
    var v16 float32 = float32(1.5)
    for fi17 := 0; fi17 < 14; fi17++ {
        if false {
            v18 := GenS2{F0: v7, F1: (^(fi17)), F2: v5}
            if v5 {
                var v19 uint32 = 4294967295
                _ = v19
            }
            _ = len(v1)
            _ = v18
        } else {
            fmt.Println(v5, v5)
            var v20 uint16 = 0
            _ = v20
        }
        v6 = fi17
        v21 := GenS2{F0: (v7 ^ (v7 << ((940) & 63))), F1: (v6 | 807), F2: v5}
        _ = fi17
        _ = v21
    }
    for fi22 := 0; fi22 < 11; fi22++ {
        var v23 uint8 = 40
        var v24 int16 = (v7 / (((v7 - (^(v7)))) | 1))
        _ = fi22
        _ = v23
        _ = v24
    }
    v6 = v6
    var v25 uint32 = 961
    switch "a" {
    case "a":
        delete(v3, ("hello" + "a"))
    }
    v26 := GenS2{F0: v7, F1: v6, F2: (!(v6 <= (v6 * v6)))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
