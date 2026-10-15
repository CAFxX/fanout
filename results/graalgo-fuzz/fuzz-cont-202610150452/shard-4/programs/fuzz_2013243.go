package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 bool }
type GenS2 struct { F0 int16; F1 int32 }

func main() {
    var v1 string = "foo bar"
    v2 := []uint32{2}
    var v3 float64 = (1.5 + -2.25)
    var v4 bool = true
    var v5 int32 = 0
    v6 := []uint32{749}
    if false {
        v7 := GenS1{F0: 18446744073709551615, F1: 1, F2: v4}
        switch "hello" {
        case "foo bar":
            for fi8 := 0; fi8 < 5; fi8++ {
                var v9 float64 = (0.0 + ((((v3 / v3) - ((3.14159 + v3) - v3)) / 1.0) * v3))
                v5 = (v5 / ((v5) | 1))
                _ = fi8
                _ = v9
            }
        case "a":
            fmt.Println(v3, v4)
        case "":
            if len(v2) > 0 {
                v2[0] = 0
            }
        }
        _ = len(v6)
        _ = v7
    } else {
        for fi10 := 0; fi10 < 20; fi10++ {
            fmt.Println(v3, v5, v3)
            v3 = v3
            _ = fi10
        }
        for fi11 := 0; fi11 < 15; fi11++ {
            v1 = v1
            fmt.Println(v4, v1, v4)
            _ = fi11
        }
    }
    v12 := GenS1{F0: 2, F1: 234, F2: false}
    v3 = v3
    for fi13 := 0; fi13 < 6; fi13++ {
        var v14 uint = 0
        _ = fi13
        _ = v14
    }
    if true {
        for fi15 := 0; fi15 < 14; fi15++ {
            v5 = (v5 % (((v5 / (((^(v5))) | 1))) | 1))
            _ = fi15
        }
    }
    v16 := GenS2{F0: 652, F1: (^(v5))}
    for fi17 := 0; fi17 < 17; fi17++ {
        v4 = (42 < 2)
        var v18 int = (-(fi17))
        _ = fi17
        _ = v18
    }
    var v19 int16 = 1
    var v20 uint32 = 0
    var v21 string = ("a" + "hello")
    var v22 uint = 65535
    var v23 int32 = (v5 ^ v5)
    var v24 float32 = 1e10
    var v25 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
