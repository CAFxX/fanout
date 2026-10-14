package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }

func main() {
    v1 := make(map[string]int)
    var v2 uint = 18446744073709551615
    var v3 int64 = 0
    var v4 int32 = 92
    v5 := make(map[string]int)
    var v6 bool = (!false)
    v7 := []int{4294967295, 448, 4294967295}
    v2 = (v2 - 18446744073709551615)
    v4 += (v4 + 1)
    fmt.Println(v4, v4, v3)
    fmt.Println(v3)
    delete(v1, ("hello" + "hello"))
    if ((-2.25 + 0.0) < 0.0) {
        for fi8 := 0; fi8 < 18; fi8++ {
            switch 0 {
            case -2147483648:
                fmt.Println(fi8, fi8)
            case 888:
                switch ("\n\t" + ("a" + "a")) {
                case "":
                    for fi9 := 0; fi9 < 6; fi9++ {
                        v10 := GenS1{F0: 0, F1: v2}
                        _ = fi9
                        _ = v10
                    }
                case "\n\t":
                    var v11 bool = (false && (991 == 1))
                    _ = v11
                case "a":
                    v3 = (v3 ^ v3)
                default:
                    if false {
                        var v12 uint = v2
                        _ = v12
                    }
                }
            case 127:
                v4 -= v4
            }
            _ = fi8
        }
        var v13 int = 32767
        _ = v13
    }
    v4 += v4
    var v14 uint = 1
    var v15 int64 = v3
    var v16 uint = (^(v14))
    for fi17 := 0; fi17 < 7; fi17++ {
        v3 -= (v15 * -572)
        for fi18 := 0; fi18 < 6; fi18++ {
            v19 := GenS1{F0: 0, F1: 4294967295}
            _ = fi18
            _ = v19
        }
        var v20 uint16 = 255
        _ = fi17
        _ = v20
    }
    v6 = v6
    v3 -= (v15 | v3)
    var v21 uint = v16
    switch "\n\t" {
    case "a":
        v22 := v1["\n\t"]
        _ = v22
    default:
        var v23 float64 = 1.5
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
