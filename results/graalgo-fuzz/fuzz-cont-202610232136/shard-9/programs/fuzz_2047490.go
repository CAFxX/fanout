package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint16 }
func genFn0(p0 uint16) int16 {
    p0 = (^(p0))
    p0 = p0
    v7 := GenS1{F0: 42, F1: 65535, F2: (p0 >> ((255) & 63))}
    var v8 string = "a"
    _ = p0
    _ = v7
    _ = v8
    return 1000
}

func main() {
    var v1 int32 = -2147483648
    v2 := []string{}
    v3 := make(map[int]bool)
    var v4 float64 = 1e10
    var v5 uint8 = 12
    v6 := []uint32{65535, 986}
    var v9 string = "\n\t"
    for fi10 := 0; fi10 < 11; fi10++ {
        v2 = append(v2, ((((v9 + "") + ((("hello" + (v9 + ("\n\t" + ""))) + v9) + v9)) + v9) + v9))
        _ = fi10
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        v5 = (v5 + v5)
        _ = fi11
    }
    var v12 int32 = v1
    for fi13 := 0; fi13 < 8; fi13++ {
        v14 := GenS1{F0: 42, F1: 18446744073709551615, F2: 1}
        _ = fi13
        _ = v14
    }
    var v15 float32 = 1e10
    var v16 int32 = (v12 >> ((0) & 63))
    delete(v3, 128)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
