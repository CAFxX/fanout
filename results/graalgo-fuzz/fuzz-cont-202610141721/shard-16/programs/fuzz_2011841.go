package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint8; F1 int16; F2 string }

func main() {
    v1 := []int{-922, -9223372036854775808, 976}
    var v2 uint = 0
    var v3 string = "foo bar"
    var v4 bool = true
    for fi5 := 0; fi5 < 1; fi5++ {
        if (0.1 < -2.25) {
            fmt.Println(v3)
            var v6 int32 = 0
            _ = v6
        }
        _ = fi5
    }
    var v7 int64 = 1
    var v8 float64 = -2.25
    if true {
        if false {
            var v9 uint = (v2 ^ 18446744073709551615)
            _ = v9
        }
        switch -585 {
        case -1:
            for fi10 := 0; fi10 < 15; fi10++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                fmt.Println(v7)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi10
            }
        default:
            for fi11 := 0; fi11 < 13; fi11++ {
                _ = len(v1)
                _ = fi11
            }
        }
        v12 := GenS1{F0: false, F1: 1}
        _ = v12
    } else {
        var v13 int16 = 1
        var v14 bool = v4
        _ = v13
        _ = v14
    }
    v15 := GenS2{F0: 2, F1: 169, F2: ("a" + v3)}
    v8 = float64(0.0)
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 float64 = v8
        _ = fi16
        _ = v17
    }
    var v18 float64 = v8
    v7 -= (v7 << ((v2) & 63))
    for fi19 := 0; fi19 < 10; fi19++ {
        v18 += (1.5 * v8)
        _ = fi19
    }
    v20 := GenS2{F0: 42, F1: 0, F2: v3}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
