package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
func genFn0() uint {
    v7 := []int{1, -954, -498, -2147483648}
    var v8 int64 = -290
    v8 -= (-(v8))
    v9 := GenS1{F0: 42, F1: 610}
    _ = v7
    _ = v8
    _ = v9
    return 2
}
func genFn1() uint {
    var v10 uint32 = 685
    for fi11 := 0; fi11 < 18; fi11++ {
        var v12 bool = false
        var v13 uint8 = 127
        _ = fi11
        _ = v12
        _ = v13
    }
    if ((!(9223372036854775807 < 1)) && false) {
        var v14 string = "hello"
        for fi15 := 0; fi15 < 5; fi15++ {
            v16 := []int{fi15, fi15}
            _ = fi15
            _ = v16
        }
        v10 = v10
        _ = v14
    } else {
        v10 = 42
        var v17 float32 = (0.0 - ((0.0 / ((1e10 * (float32((-2.25 - (0.0 - (0.0 / (1.5 - 1e10))))) - 0.1)) - (1.5 / 1e10))) * (0.1 * (3.14159 + (1.5 / 1.5)))))
        _ = v17
    }
    _ = v10
    return 255
}

func main() {
    var v1 int64 = -1
    var v2 int64 = v1
    v3 := []string{}
    v4 := []string{"", (("a" + "hello") + ""), (("\n\t" + "\n\t") + "hello")}
    v5 := []int{}
    v6 := []uint32{1, 596, 1}
    v1 = v2
    v1 = v1
    _ = len(v3)
    v1 = v2
    var v18 int16 = -664
    var v19 uint = 42
    fmt.Println(v2, v2)
    switch 9223372036854775807 {
    case 32768:
        v2 = v2
    default:
        v20 := GenS1{F0: 210, F1: 839}
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
