package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 uint8; F1 int8; F2 uint32 }
func genFn0(p0 uint16) uint32 {
    p0 += p0
    p0 = p0
    _ = p0
    return 65535
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{1, -329, 9223372036854775807, -546}
    v3 := make(map[uint32]bool)
    var v4 uint16 = 65535
    var v5 int16 = 1000
    v6 := []uint32{41}
    for fi7 := 0; fi7 < 12; fi7++ {
        if (!(true || (((false && (!(!(!true)))) || (!(v5 == (v5 + 32767)))) && true))) {
            delete(v1, 17)
            var v8 int64 = -9223372036854775808
            _ = v8
        }
        _ = fi7
    }
    var v9 int32 = 767
    var v10 bool = false
    v11 := GenS1{F0: v9, F1: v5}
    for fi12 := 0; fi12 < 9; fi12++ {
        v13 := GenS1{F0: 0, F1: (v5 / ((v5) | 1))}
        v9 -= v9
        var v14 string = "foo bar"
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 int16 = v5
    v16 := GenS2{F0: 42, F1: -1, F2: 464}
    v9 += v9
    switch "hello" {
    case "":
        v4 = v4
    }
    if (!v10) {
        if len(v2) > 0 {
            v2[0] = 1
        }
        v9 -= (v9 >> ((18446744073709551615) & 63))
    }
    v10 = (0 == 815)
    var v17 bool = (v10 && v10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
