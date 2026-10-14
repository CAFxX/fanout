package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = ((("" + "\n\t") + ("a" + "a")) == "a")
    var v3 int64 = -848
    v4 := make(map[string]uint32)
    v3 = -1
    v2 = true
    if v2 {
        for fi5 := 0; fi5 < 3; fi5++ {
            switch -43 {
            case -1:
                if v2 {
                    if v2 {
                        v6 := []int{(fi5 | fi5)}
                        var v7 uint16 = 368
                        var v8 uint32 = 2147483648
                        _ = v6
                        _ = v7
                        _ = v8
                    }
                    for fi9 := 0; fi9 < 19; fi9++ {
                        var v10 uint32 = 4294967295
                        var v11 int64 = 583
                        var v12 bool = v2
                        _ = fi9
                        _ = v10
                        _ = v11
                        _ = v12
                    }
                } else {
                    var v13 string = "hello"
                    for fi14 := 0; fi14 < 9; fi14++ {
                        var v15 int16 = -1
                        var v16 int64 = (v3 * (v3 ^ -9223372036854775808))
                        var v17 int16 = int16(v16)
                        _ = fi14
                        _ = v15
                        _ = v16
                        _ = v17
                    }
                    _ = v13
                }
            case -778:
                var v18 uint16 = 255
                _ = v18
            case 979:
                for fi19 := 0; fi19 < 13; fi19++ {
                    v3 += (v3 % ((v3) | 1))
                    var v20 int = (fi5 - (fi19 & fi5))
                    _ = fi19
                    _ = v20
                }
            default:
                var v21 uint16 = 1
                _ = v21
            }
            _ = fi5
        }
        var v22 uint32 = 65535
        for fi23 := 0; fi23 < 14; fi23++ {
            v2 = v2
            _ = fi23
        }
        _ = v22
    } else {
        fmt.Println(v2, v2)
        v2 = (!(0.0 == 0.0))
        var v24 int32 = 0
        _ = v24
    }
    var v25 uint8 = 155
    fmt.Println(v3)
    v26 := []int{}
    fmt.Println(v2, v2)
    var v27 uint8 = v25
    v28 := GenS1{F0: ((!(true && v2)) && true), F1: -1}
    var v29 int8 = 1
    v25 = v27
    if len(v26) >= 2 {
        v26 = v26[1:len(v26)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
