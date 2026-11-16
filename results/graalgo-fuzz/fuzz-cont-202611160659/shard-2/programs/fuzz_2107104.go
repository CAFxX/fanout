package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
func genFn0() uint32 {
    if true {
        switch ("a" + "") {
        case "foo bar":
            var v7 float32 = -2.25
            _ = v7
        case "a":
            var v8 int16 = 42
            _ = v8
        default:
            var v9 int16 = 201
            _ = v9
        }
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 int32 = 2147483647
            switch "a" {
            case "\n\t":
                fmt.Println(v11, v11)
            case "hello":
                v11 -= (v11 - -1)
            default:
                for fi12 := 0; fi12 < 15; fi12++ {
                    var v13 uint64 = 0
                    _ = fi12
                    _ = v13
                }
            }
            for fi14 := 0; fi14 < 13; fi14++ {
                if false {
                    v11 = v11
                    v11 -= v11
                    v15 := []int{fi14, (fi14 >> ((200) & 63))}
                    var v16 int64 = 1
                    _ = v15
                    _ = v16
                }
                v11 = v11
                _ = fi14
            }
            _ = fi10
            _ = v11
        }
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        var v18 uint16 = 32767
        _ = fi17
        _ = v18
    }
    var v19 uint64 = 0
    if false {
        fmt.Println(v19)
        v20 := []int{127, 1}
        _ = v20
    } else {
        v19 = v19
    }
    _ = v19
    return 2147483648
}
func genFn1(p0 bool) uint16 {
    var v21 float64 = 3.14159
    v22 := []int{128, 42, 128}
    _ = p0
    _ = v21
    _ = v22
    return 969
}

func main() {
    var v1 int32 = -1
    v2 := make(map[int]uint32)
    var v3 uint = 0
    v4 := []uint32{65535}
    var v5 string = "a"
    var v6 bool = ((false || (0 < 141)) || (true || (0 >= -1)))
    fmt.Println(v6, v3, v3)
    var v23 float32 = 0.0
    v3 += 4294967295
    v23 += 0.0
    v6 = v6
    v3 -= 483
    _ = len(v4)
    if ((v6 || false) && v6) {
        v24 := GenS1{F0: 638, F1: 0}
        _ = v24
    }
    switch -278 {
    case 32767:
        var v25 int = 506
        _ = v25
    case 4294967295:
        fmt.Println(v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
