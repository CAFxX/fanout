package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint }
type GenS2 struct { F0 uint64; F1 uint }
func genFn0() uint8 {
    var v5 uint16 = 65535
    if true {
        var v6 int = -889
        _ = v6
    } else {
        v7 := []int{}
        var v8 uint64 = 0
        _ = v7
        _ = v8
    }
    _ = v5
    return uint8(v5)
}
func genFn1(p0 uint16, p1 int8, p2 int32) uint64 {
    var v9 uint16 = p0
    fmt.Println(p0, p0, v9)
    var v10 float64 = 0.0
    if ((v10 - 0.0) < (3.14159 / float64(1.5))) {
        var v11 int8 = (p1 >> ((65535) & 63))
        _ = v11
    }
    v12 := []int{452}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v12
    _ = v9
    return 4294967295
}

func main() {
    var v1 uint64 = 576
    var v2 string = "\n\t"
    var v3 int16 = -831
    var v4 uint = 2
    var v13 uint16 = 0
    v14 := []int{}
    fmt.Println(v1, v2)
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    if len(v14) > 0 {
        v14[0] = 0
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        v13 -= v13
        _ = fi15
    }
    v16 := GenS1{F0: (v4 % (((v4 >> ((v4) & 63))) | 1)), F1: genFn1(v13, int8(v1), -2147483648), F2: v4}
    var v17 int = 320
    v18 := GenS2{F0: (v1 ^ (^(v1))), F1: (v4 ^ v4)}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
