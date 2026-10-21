package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool; F1 bool }
func genFn0() uint8 {
    var v6 int32 = 307
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := GenS1{F0: 65535}
        var v9 bool = false
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{-9223372036854775808}
    _ = v10
    _ = v6
    return 36
}

func main() {
    var v1 bool = false
    var v2 uint64 = 536
    var v3 int16 = 928
    var v4 int64 = 168
    var v5 bool = (42 >= -2147483648)
    for fi11 := 0; fi11 < 9; fi11++ {
        var v12 uint32 = 885
        _ = fi11
        _ = v12
    }
    fmt.Println(v2, v2)
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 uint32 = 42
        v15 := GenS1{F0: 0}
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 int64 = 42
    switch "" {
    case "a":
        var v17 bool = (-1 >= -128)
        _ = v17
    case "hello":
        for fi18 := 0; fi18 < 8; fi18++ {
            var v19 int = (fi18 * fi18)
            var v20 int = fi18
            var v21 uint16 = 0
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
    case "foo bar":
        fmt.Println(v5, v2)
    }
    var v22 int64 = 1
    var v23 int8 = 0
    v24 := []int{int(v4), 0, 128}
    v5 = v1
    v25 := GenS2{F0: (0 <= 65535), F1: (v1 && false)}
    switch ("foo bar" + ("hello" + (("\n\t" + ("hello" + "hello")) + "foo bar"))) {
    case "a":
        v26 := GenS1{F0: 255}
        _ = v26
    }
    for fi27 := 0; fi27 < 18; fi27++ {
        if (198 < 39) {
            for fi28 := 0; fi28 < 15; fi28++ {
                switch "foo bar" {
                case "a":
                    v22 = -1
                    v4 = 839
                case "hello":
                    v22 = v22
                case "":
                    v29 := GenS2{F0: v5, F1: (v1 && (2147483648 == 2))}
                    _ = v29
                }
                _ = len(v24)
                fmt.Println(v5)
                _ = fi28
            }
            fmt.Println(v23, v22)
            v30 := GenS2{F0: (!true), F1: ((int32(v3) >= -2147483648) && ((!(false && (!(!(int16(v22) != 156))))) || v5))}
            _ = v30
        }
        v22 = v4
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
