package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint64 }
func genFn0(p0 uint8, p1 int64) uint16 {
    var v5 uint = 18446744073709551615
    p1 = p1
    v5 = 18446744073709551615
    v6 := GenS1{F0: 42, F1: -460, F2: 1}
    v7 := []int{}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 32767
}

func main() {
    v1 := []string{"\n\t", "hello", "", (("hello" + "a") + "hello")}
    var v2 uint16 = 32767
    var v3 float32 = 0.1
    var v4 float64 = float64(v3)
    if (906 >= 237) {
        v3 = -2.25
    } else {
        if (4294967295 > 2) {
            v3 = 0.1
            var v8 float64 = (1.5 * v4)
            _ = v8
        } else {
            var v9 uint64 = 314
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch -1 {
            case 4294967295:
                v9 = 0
            }
            _ = v9
        }
        var v10 int = 42
        fmt.Println(v3)
        _ = v10
    }
    fmt.Println(v3, v4, v3)
    if ((((v4 + 0.0) + float64(v3)) >= float64((v3 * v3))) || (!false)) {
        fmt.Println(v4, v4, v4)
        v3 += (v3 / v3)
        _ = len(v1)
    } else {
        v4 = 0.1
    }
    var v11 uint64 = uint64(v2)
    _ = len(v1)
    switch "" {
    case "":
        for fi12 := 0; fi12 < 10; fi12++ {
            var v13 uint64 = (^(v11))
            _ = fi12
            _ = v13
        }
    default:
        var v14 float64 = float64(0.1)
        _ = v14
    }
    v1 = append(v1, "a")
    for fi15 := 0; fi15 < 10; fi15++ {
        v11 -= (^(v11))
        _ = fi15
    }
    _ = len(v1)
    if len(v1) > 0 {
        v1[0] = ("\n\t" + "\n\t")
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
}
