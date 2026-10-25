package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 string }

func main() {
    v1 := []string{("a" + ("" + "a")), "hello", "\n\t"}
    v2 := make(map[uint32]string)
    var v3 uint64 = 859
    v4 := []int{146, 9223372036854775807, 32768, 73}
    for fi5 := 0; fi5 < 18; fi5++ {
        v6 := GenS1{F0: true}
        v3 -= v3
        var v7 uint16 = 1
        _ = fi5
        _ = v6
        _ = v7
    }
    v3 -= (^(v3))
    v8 := GenS1{F0: false}
    v9 := GenS2{F0: 2147483647, F1: "a"}
    if true {
        for fi10 := 0; fi10 < 2; fi10++ {
            for fi11 := 0; fi11 < 14; fi11++ {
                var v12 uint64 = (^(v3))
                switch -99 {
                case 128:
                    var v13 uint8 = 2
                    _ = v13
                case 4294967295:
                    fmt.Println(v3)
                }
                fmt.Println(v3)
                _ = fi11
                _ = v12
            }
            var v14 float64 = 0.1
            _ = fi10
            _ = v14
        }
        var v15 bool = false
        _ = v15
    } else {
        fmt.Println(v3)
    }
    if true {
        fmt.Println(v3)
        var v16 int32 = 42
        var v17 int32 = v16
        _ = v16
        _ = v17
    } else {
        fmt.Println(v3)
        if true {
            v3 += (v3 - (v3 / (((v3 + v3)) | 1)))
            var v18 int = 9223372036854775807
            _ = len(v2)
            _ = v18
        }
    }
    var v19 float64 = (3.14159 * 0.0)
    v19 += (float64((3.14159 * (-2.25 * ((1.5 * 0.0) / 1.5)))) + v19)
    for fi20 := 0; fi20 < 12; fi20++ {
        fmt.Println(fi20, fi20, v19)
        if ((-1 == 312) && (!((!(18446744073709551615 >= 4294967295)) || ("foo bar" == "foo bar")))) {
            v21 := GenS1{F0: true}
            _ = v21
        }
        var v22 uint8 = 7
        _ = fi20
        _ = v22
    }
    if true {
        v3 = (v3 + v3)
    }
    fmt.Println(v19)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
