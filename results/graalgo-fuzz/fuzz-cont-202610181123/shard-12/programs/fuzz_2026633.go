package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }

func main() {
    v1 := make(map[int]bool)
    var v2 bool = true
    var v3 int32 = -1
    var v4 bool = (v2 || (!(((float64(1.5) + (3.14159 / 0.1)) + 3.14159) != 1e10)))
    v3 = (v3 & v3)
    v5 := GenS1{F0: 65535, F1: 852}
    if (((v3 & v3) == v3) || v2) {
        var v6 int = 4294967295
        _ = v6
    } else {
        delete(v1, 32768)
    }
    var v7 bool = (-101 <= -236)
    var v8 bool = (!false)
    switch 0 {
    case 128:
        v9 := GenS1{F0: 4294967295, F1: 909}
        _ = v9
    case 488:
        for fi10 := 0; fi10 < 10; fi10++ {
            switch ("foo bar" + "hello") {
            case "\n\t":
                fmt.Println(v7)
            }
            _ = fi10
        }
    }
    v8 = v4
    for fi11 := 0; fi11 < 19; fi11++ {
        var v12 int64 = -9223372036854775808
        if false {
            v8 = true
            switch 161 {
            case -631:
                v4 = true
            case 42:
                switch (fi11 - (^(fi11))) {
                case 9223372036854775807:
                    v7 = (false && v7)
                case 721:
                    for fi13 := 0; fi13 < 12; fi13++ {
                        fmt.Println(v8, v4)
                        if (!true) {
                            v14 := GenS1{F0: 665, F1: 2147483647}
                            _ = v14
                        }
                        _ = fi13
                    }
                default:
                    var v15 string = ("a" + ((("foo bar" + ("" + ("foo bar" + "foo bar"))) + "\n\t") + ((("" + ("foo bar" + ("\n\t" + ("foo bar" + ("\n\t" + "a"))))) + ("hello" + "hello")) + ("foo bar" + ("a" + "hello")))))
                    _ = v15
                }
            }
            var v16 float64 = 0.0
            _ = v16
        } else {
            var v17 string = ("hello" + "")
            _ = v17
        }
        _ = fi11
        _ = v12
    }
    var v18 int8 = 100
    v19 := []int{-9223372036854775808}
    for fi20 := 0; fi20 < 10; fi20++ {
        fmt.Println(v2, v8)
        var v21 float32 = 1.5
        _ = fi20
        _ = v21
    }
    var v22 int16 = 42
    v3 -= (v3 | v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
