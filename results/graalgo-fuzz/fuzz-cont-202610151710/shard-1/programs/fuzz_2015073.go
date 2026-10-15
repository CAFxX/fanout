package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 int64 }
func genFn0(p0 bool, p1 int16, p2 string) int8 {
    fmt.Println(p2, p0, p0)
    var v5 int8 = -37
    var v6 uint64 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (v5 % (((v5 - v5)) | 1))
}

func main() {
    var v1 int = -695
    v2 := []uint32{42, 2}
    var v3 int64 = 42
    var v4 float32 = 0.0
    if true {
        fmt.Println(v1)
    }
    var v7 uint64 = 77
    fmt.Println(v4, v3)
    v8 := GenS1{F0: 255, F1: (!(!true))}
    fmt.Println(v7)
    v9 := GenS1{F0: 1, F1: true}
    v1 = (-(v1))
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    switch (v1 << ((4294967295) & 63)) {
    case 4294967295:
        var v10 string = ((("foo bar" + ("foo bar" + ("foo bar" + "\n\t"))) + "a") + (("\n\t" + ("" + ("\n\t" + "\n\t"))) + "foo bar"))
        _ = v10
    }
    var v11 int8 = 0
    var v12 int = v1
    var v13 float32 = 3.14159
    v11 = (v11 % ((int8(v3)) | 1))
    for fi14 := 0; fi14 < 9; fi14++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if true {
            if false {
                switch v1 {
                case -608:
                    fmt.Println(v7, v12, v1)
                default:
                    if len(v2) > 0 {
                        v2[0] = 694
                    }
                }
                v15 := GenS2{F0: (v3 << ((0) & 63))}
                _ = v15
            }
            if true {
                for fi16 := 0; fi16 < 12; fi16++ {
                    v2 = append(v2, 65535)
                    fmt.Println(fi14, v13, v3)
                    var v17 uint = 42
                    _ = fi16
                    _ = v17
                }
            }
            var v18 uint64 = (^(v7))
            _ = v18
        }
        v19 := GenS2{F0: (-(v3))}
        _ = fi14
        _ = v19
    }
    for fi20 := 0; fi20 < 19; fi20++ {
        var v21 int32 = 1
        v1 = (v12 / ((fi20) | 1))
        _ = fi20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
