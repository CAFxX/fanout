package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 bool }
func genFn0(p0 string) uint16 {
    if ((p0 + "a") == ("a" + p0)) {
        v7 := []int{192, 2147483647}
        var v8 float64 = (1.5 - 1.5)
        _ = v7
        _ = v8
    }
    if (!(((3.14159 == 0.0) && false) || (false || (32768 <= 1)))) {
        if (1 >= 1) {
            var v9 int8 = -61
            _ = v9
        } else {
            p0 = p0
        }
    } else {
        p0 = p0
    }
    p0 = "a"
    _ = p0
    return 65535
}

func main() {
    v1 := []uint32{65535}
    v2 := make(map[uint32]int)
    var v3 int = 903
    var v4 int64 = 332
    v5 := make(map[int]uint32)
    var v6 uint8 = 2
    fmt.Println(v4, v4)
    var v10 string = ("\n\t" + "hello")
    for fi11 := 0; fi11 < 10; fi11++ {
        var v12 int64 = v4
        var v13 bool = true
        var v14 int = (v3 + (v3 % ((1) | 1)))
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    var v15 uint8 = v6
    for fi16 := 0; fi16 < 9; fi16++ {
        v3 += (v3 & (v3 >> ((65535) & 63)))
        _ = fi16
    }
    v10 = v10
    fmt.Println(v4, v4)
    v17 := GenS1{F0: 18446744073709551615, F1: v15, F2: (!false)}
    fmt.Println(v3, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
