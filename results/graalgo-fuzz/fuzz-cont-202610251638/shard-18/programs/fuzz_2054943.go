package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 bool }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]int)
    v3 := make(map[string]int)
    var v4 int8 = 100
    fmt.Println(v4)
    var v5 float32 = (-2.25 - ((float32(-2.25) + float32(0.0)) - -2.25))
    var v6 string = ("foo bar" + "foo bar")
    v4 -= v4
    v4 = (^(v4))
    v7 := GenS1{F0: 42, F1: 0, F2: true}
    v6 = ("a" + v6)
    v4 -= (v4 % (((v4 << ((4294967295) & 63))) | 1))
    fmt.Println(v4, v6, v6)
    for fi8 := 0; fi8 < 11; fi8++ {
        fmt.Println(fi8, v4)
        _ = fi8
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        if false {
            var v10 int64 = 0
            for fi11 := 0; fi11 < 19; fi11++ {
                v12 := GenS1{F0: 1, F1: 2, F2: true}
                _ = fi11
                _ = v12
            }
            _ = v10
        }
        if false {
            switch (v6 + (v6 + v6)) {
            case "":
                v4 -= -44
            }
            for fi13 := 0; fi13 < 6; fi13++ {
                var v14 float32 = 0.1
                var v15 int = fi13
                delete(v3, ("\n\t" + (v6 + (v6 + (v6 + (((v6 + "a") + "foo bar") + v6))))))
                var v16 int64 = 0
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
            }
        }
        var v17 uint16 = 262
        _ = fi9
        _ = v17
    }
    v18 := GenS1{F0: -153, F1: 1, F2: false}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
