package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 string = ("" + "")
    var v2 bool = (false || false)
    v3 := []int{-186}
    for fi4 := 0; fi4 < 6; fi4++ {
        fmt.Println(fi4)
        for fi5 := 0; fi5 < 4; fi5++ {
            var v6 uint = 255
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    v2 = v2
    var v7 int32 = -1
    if v2 {
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 float32 = ((1.5 - ((3.14159 - -2.25) / 0.1)) - -2.25)
            for fi10 := 0; fi10 < 10; fi10++ {
                if v2 {
                    v9 += v9
                    if v2 {
                        if len(v3) > 0 {
                            v3[0] = (fi8 / (((-(fi10))) | 1))
                        }
                    } else {
                        var v11 uint = 42
                        _ = v11
                    }
                }
                switch (fi8 & fi8) {
                case -885:
                    var v12 uint64 = 100
                    _ = v12
                default:
                    fmt.Println(fi10, fi8)
                }
                _ = fi10
            }
            _ = fi8
            _ = v9
        }
        fmt.Println(v2, v2, v1)
    }
    v13 := GenS1{F0: 2}
    if v2 {
        var v14 int64 = -538
        var v15 uint = 42
        fmt.Println(v1)
        _ = v14
        _ = v15
    }
    var v16 float64 = 3.14159
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 float32 = float32((0.1 - (v16 * v16)))
        v19 := GenS2{F0: -622}
        v18 = v18
        _ = fi17
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        fmt.Println(v16, fi20, fi20)
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v7
}
