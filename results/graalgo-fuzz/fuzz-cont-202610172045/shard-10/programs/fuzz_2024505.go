package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint16 }
func genFn0(p0 uint32, p1 uint16, p2 uint) int8 {
    var v5 uint32 = 4294967295
    v6 := []int{}
    switch "\n\t" {
    case "hello":
        p2 = (p2 << ((p2) & 63))
    case "foo bar":
        fmt.Println(v5, p1, p0)
    default:
        p1 -= 65535
    }
    _ = len(v6)
    p1 -= p1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -7
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 42
    var v3 float64 = (((1e10 - 1e10) * 0.0) / 0.1)
    var v4 uint16 = 2
    var v7 int = 412
    fmt.Println(v2, v4, v2)
    v8 := []int{v7}
    var v9 float32 = 3.14159
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    var v10 uint = 962
    _ = len(v1)
    var v11 uint32 = 445
    if (true || (true && (419 < 2))) {
        switch (^(v7)) {
        case 128:
            for fi12 := 0; fi12 < 13; fi12++ {
                for fi13 := 0; fi13 < 17; fi13++ {
                    if false {
                        v14 := GenS1{F0: (v4 % ((v4) | 1)), F1: true, F2: 124}
                        var v15 uint64 = 0
                        _ = v14
                        _ = v15
                    }
                    _ = fi13
                }
                _ = fi12
            }
        }
    } else {
        var v16 int64 = int64(v11)
        var v17 bool = ((v11 - v11) <= 635)
        _ = v16
        _ = v17
    }
    if (v9 == v9) {
        var v18 uint8 = 127
        _ = v18
    }
    var v19 string = "\n\t"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
