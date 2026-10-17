package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint16 }
type GenS2 struct { F0 bool; F1 int32; F2 int32 }
func genFn0(p0 int32, p1 uint16) uint {
    for fi5 := 0; fi5 < 15; fi5++ {
        if true {
            p1 += (^(p1))
            v6 := []int{128, 128, 1, fi5}
            p0 += 0
            _ = v6
        }
        _ = fi5
    }
    fmt.Println(p0, p0)
    if true {
        var v7 float32 = (0.0 - 1.5)
        _ = v7
    } else {
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 int32 = (p0 & (p0 & -143))
            _ = fi8
            _ = v9
        }
        fmt.Println(p1, p0)
    }
    fmt.Println(p0, p1)
    switch 2147483647 {
    case 32767:
        v10 := GenS1{F0: 4294967295, F1: 294, F2: (p1 / ((p1) | 1))}
        _ = v10
    default:
        v11 := []int{}
        _ = v11
    }
    _ = p0
    _ = p1
    return 322
}

func main() {
    var v1 uint16 = 0
    var v2 uint64 = 4294967295
    var v3 uint32 = 42
    v4 := []uint32{2147483648}
    if true {
        fmt.Println(v1, v3, v3)
        if false {
            var v12 uint16 = v1
            _ = v12
        } else {
            v13 := GenS2{F0: false, F1: 1, F2: -127}
            switch 9223372036854775807 {
            case 0:
                v1 -= v1
            default:
                if false {
                    var v14 int = int(v1)
                    var v15 uint16 = (^(v1))
                    _ = v14
                    _ = v15
                } else {
                    fmt.Println(v1, v3)
                }
            }
            var v16 bool = true
            _ = v13
            _ = v16
        }
        v17 := GenS2{F0: true, F1: 42, F2: -1}
        _ = v17
    }
    switch ("" + "hello") {
    case "hello":
        v4 = append(v4, v3)
    }
    var v18 int16 = 1
    if false {
        var v19 float64 = (-2.25 / 3.14159)
        v19 = v19
        for fi20 := 0; fi20 < 13; fi20++ {
            v3 = v3
            _ = fi20
        }
        _ = v19
    } else {
        switch "a" {
        case "a":
            v21 := GenS1{F0: 877, F1: (v2 / ((v2) | 1)), F2: v1}
            _ = v21
        }
    }
    v1 -= v1
    for fi22 := 0; fi22 < 16; fi22++ {
        var v23 uint8 = 42
        _ = fi22
        _ = v23
    }
    var v24 int = 2147483647
    v2 = (v2 >> ((719) & 63))
    var v25 string = (("" + ("a" + ((("" + "hello") + "") + ((("" + ("\n\t" + "")) + ("" + "foo bar")) + ("" + ("\n\t" + "foo bar")))))) + "a")
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
