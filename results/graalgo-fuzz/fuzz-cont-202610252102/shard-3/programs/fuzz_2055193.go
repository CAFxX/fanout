package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 uint32; F1 int16 }
func genFn0(p0 string, p1 int16, p2 uint8) int64 {
    v6 := GenS1{F0: 42, F1: (4294967295 != 18446744073709551615)}
    var v7 uint32 = 9
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -73
}
func genFn1(p0 uint8) int64 {
    var v8 bool = (!true)
    var v9 uint32 = 2
    _ = p0
    _ = v8
    _ = v9
    return -745
}

func main() {
    var v1 string = (((("\n\t" + "") + "a") + "foo bar") + "a")
    v2 := make(map[string]uint32)
    var v3 int64 = 1
    v4 := make(map[int]bool)
    v5 := []uint32{4294967295, 2, 2147483648}
    var v10 uint32 = 2147483648
    if (0.1 != float64(0.1)) {
        switch "\n\t" {
        case "a":
            switch 9223372036854775807 {
            case 42:
                var v11 bool = false
                _ = v11
            case 0:
                for fi12 := 0; fi12 < 6; fi12++ {
                    fmt.Println(v3, fi12)
                    fmt.Println(fi12)
                    _ = fi12
                }
            case 128:
                var v13 string = v1
                _ = v13
            default:
                if false {
                    var v14 bool = false
                    var v15 uint = 643
                    _ = v14
                    _ = v15
                }
            }
        case "hello":
            fmt.Println(v10)
        default:
            v1 = v1
        }
    } else {
        v1 = "\n\t"
        var v16 bool = false
        _ = v16
    }
    var v17 int = 127
    v3 += (^(v3))
    for fi18 := 0; fi18 < 13; fi18++ {
        v3 = (v3 ^ 1)
        _ = fi18
    }
    var v19 bool = (!false)
    var v20 float64 = float64(3.14159)
    v21 := GenS2{F0: v10, F1: -504}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
