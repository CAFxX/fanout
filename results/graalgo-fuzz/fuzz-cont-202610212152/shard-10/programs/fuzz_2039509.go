package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint; F1 uint }
func genFn0(p0 uint16, p1 int) int64 {
    var v6 uint = 65535
    fmt.Println(p1, v6)
    v7 := []int{-429}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -1
}

func main() {
    var v1 uint64 = 572
    var v2 float32 = 0.0
    v3 := make(map[uint32]string)
    var v4 int32 = -964
    var v5 int = 32767
    var v8 int8 = -128
    var v9 uint32 = 911
    var v10 string = "a"
    var v11 float64 = ((3.14159 / (1.5 / 1e10)) - -2.25)
    if false {
        v1 = v1
        v12 := GenS1{F0: 42, F1: true}
        _ = v12
    }
    var v13 int = (v5 & 4294967295)
    v14 := GenS1{F0: 668, F1: true}
    v15 := GenS1{F0: 4294967295, F1: ((v9 ^ v9) > 1)}
    v8 = (v8 | v8)
    v4 = v4
    switch (v5 & (^(v5))) {
    case 4294967295:
        if true {
            v16 := GenS1{F0: 0, F1: (42 != 18446744073709551615)}
            for fi17 := 0; fi17 < 6; fi17++ {
                if (!true) {
                    v13 += -139
                }
                _ = fi17
            }
            _ = v16
        }
    default:
        switch (v10 + (v10 + "foo bar")) {
        case "a":
            v1 = 227
        case "foo bar":
            var v18 string = v10
            _ = v18
        default:
            v19 := GenS2{F0: 42, F1: 38}
            _ = v19
        }
    }
    v20 := []int{v13, 127}
    v1 += (v1 & uint64(v9))
    _ = len(v20)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
