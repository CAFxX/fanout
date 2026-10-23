package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int16; F1 int16; F2 uint8 }
func genFn0(p0 int) bool {
    switch p0 {
    case 491:
        var v7 uint8 = 100
        _ = v7
    case 66:
        v8 := GenS1{F0: p0, F1: -1}
        _ = v8
    default:
        var v9 uint64 = 18446744073709551615
        _ = v9
    }
    switch (p0 + p0) {
    case -1:
        v10 := GenS1{F0: -963, F1: 100}
        _ = v10
    case 127:
        for fi11 := 0; fi11 < 11; fi11++ {
            p0 += 127
            _ = fi11
        }
    case -521:
        var v12 bool = true
        _ = v12
    default:
        var v13 float32 = 3.14159
        _ = v13
    }
    _ = p0
    return (float64(float32((float64(0.1) + float64(-2.25)))) >= (0.1 * 0.0))
}
func genFn1() int64 {
    var v14 int32 = 1
    var v15 int8 = 127
    _ = v14
    _ = v15
    return -788
}

func main() {
    var v1 int16 = 32767
    v2 := []string{"a", "hello", ""}
    var v3 int64 = 1
    v4 := make(map[uint32]string)
    var v5 int64 = (v3 | (v3 << ((1) & 63)))
    v6 := make(map[int]int)
    var v16 uint32 = 2
    v1 += (v1 % ((1000) | 1))
    var v17 int = 32768
    v5 -= (v5 << ((42) & 63))
    var v18 int16 = (^(v1))
    for fi19 := 0; fi19 < 9; fi19++ {
        v20 := GenS1{F0: (fi19 << ((4294967295) & 63)), F1: -71}
        _ = fi19
        _ = v20
    }
    switch ("hello" + "a") {
    case "":
        var v21 int64 = int64(v1)
        _ = v21
    }
    fmt.Println(v3)
    var v22 uint8 = 152
    v23 := GenS2{F0: (^(v1)), F1: v1, F2: v22}
    var v24 uint16 = 32767
    v22 += (v22 ^ (v22 & 2))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
