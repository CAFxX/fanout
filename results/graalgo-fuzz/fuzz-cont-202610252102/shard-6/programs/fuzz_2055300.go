package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 int8; F2 uint32 }
func genFn0(p0 int8, p1 uint16, p2 int) uint16 {
    var v5 uint16 = (p1 << ((255) & 63))
    if true {
        var v6 int32 = 1
        fmt.Println(p2, p0, v5)
        var v7 uint16 = v5
        _ = v6
        _ = v7
    } else {
        p2 = p2
        fmt.Println(p1)
    }
    v8 := GenS1{F0: 42}
    v9 := GenS2{F0: 0, F1: -8, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v8
    _ = v9
    return 65535
}

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := []uint32{65535, 4294967295, 105, 0}
    var v3 int = -9223372036854775808
    v4 := []int{(-(v3))}
    v10 := GenS2{F0: 0, F1: 42, F2: 2147483648}
    switch "" {
    case "hello":
        if len(v2) > 0 {
            v2[0] = 65535
        }
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        fmt.Println(v1, fi11)
        if (!(fi11 >= int(v1))) {
            var v12 int = v3
            _ = v12
        }
        _ = fi11
    }
    var v13 uint = 1
    var v14 bool = true
    v15 := GenS2{F0: -9223372036854775808, F1: 0, F2: 905}
    v16 := GenS1{F0: 1}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi17 := 0; fi17 < 9; fi17++ {
        if (2147483647 == 390) {
            var v18 int64 = -9223372036854775808
            for fi19 := 0; fi19 < 17; fi19++ {
                switch "" {
                case "":
                    var v20 uint64 = (v1 & 18446744073709551615)
                    _ = v20
                }
                _ = fi19
            }
            _ = v18
        }
        _ = fi17
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        v2 = append(v2, 1)
        v22 := GenS2{F0: 1, F1: 127, F2: 387}
        v23 := GenS2{F0: 42, F1: 127, F2: 205}
        _ = fi21
        _ = v22
        _ = v23
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
