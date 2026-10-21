package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
func genFn0(p0 int8, p1 string) uint32 {
    var v6 int32 = -1
    for fi7 := 0; fi7 < 3; fi7++ {
        p1 = (((p1 + "foo bar") + ((p1 + (p1 + p1)) + p1)) + ((p1 + ((p1 + p1) + ("a" + "foo bar"))) + p1))
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v6
    return 4294967295
}

func main() {
    var v1 int = -527
    var v2 int = v1
    var v3 int8 = 2
    var v4 int = v1
    var v5 float64 = 0.0
    switch (v2 & (-(v1))) {
    case -9223372036854775808:
        fmt.Println(v2, v4)
    }
    v5 -= (v5 * ((((1.5 / 0.1) - ((((v5 + (v5 + 0.1)) - 1e10) * float64(1e10)) * float64((-2.25 / float32((v5 - (float64(0.0) - ((1e10 * 0.1) - (((v5 / 3.14159) * (v5 - (((0.0 + v5) - (((0.1 - v5) + (v5 + 0.1)) + (float64(1e10) / (v5 - v5)))) / v5))) * v5))))))))) * (1e10 - v5)) + v5))
    for fi8 := 0; fi8 < 6; fi8++ {
        var v9 int16 = 1
        _ = fi8
        _ = v9
    }
    var v10 bool = false
    v11 := GenS1{F0: "", F1: ("hello" + "foo bar")}
    v12 := []int{v4}
    var v13 bool = v10
    var v14 uint16 = 747
    v13 = ((v14 >= v14) || (127 > 1))
    var v15 string = "foo bar"
    var v16 float32 = (1.5 - 0.1)
    var v17 uint64 = 4294967295
    v4 = (v4 >> ((18446744073709551615) & 63))
    if len(v12) > 0 {
        v12[0] = v4
    }
    var v18 float64 = float64(v16)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
