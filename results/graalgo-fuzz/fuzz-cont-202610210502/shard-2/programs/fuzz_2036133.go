package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint16 }
func genFn0(p0 uint32, p1 string, p2 uint) uint32 {
    v5 := GenS1{F0: int16(p0), F1: 2147483647, F2: 2}
    v6 := []int{}
    p0 += (p0 << (((p2 / ((p2) | 1))) & 63))
    var v7 float32 = (1.5 / 3.14159)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return (p0 - (p0 % ((p0) | 1)))
}

func main() {
    var v1 int16 = -40
    var v2 uint = 18446744073709551615
    var v3 int64 = 441
    var v4 uint16 = 65535
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 int64 = (v3 >> ((4294967295) & 63))
        _ = fi8
        _ = v9
    }
    if true {
        v10 := []int{32768}
        _ = v10
    } else {
        v2 -= v2
    }
    var v11 uint16 = (^(v4))
    v4 = (v11 / ((v4) | 1))
    if (((!(!true)) || ((42 < 187) && true)) || (((("hello" + ("foo bar" + "\n\t")) < "foo bar") || true) && false)) {
        v12 := GenS1{F0: (v1 * (v1 >> (((v2 | 0)) & 63))), F1: -593, F2: 40}
        for fi13 := 0; fi13 < 8; fi13++ {
            v14 := []int{fi13}
            _ = fi13
            _ = v14
        }
        v3 += v3
        _ = v12
    } else {
        var v15 int32 = -1
        v16 := GenS1{F0: v1, F1: (-(v15)), F2: (v4 >> ((v2) & 63))}
        _ = v15
        _ = v16
    }
    v4 -= 657
    var v17 uint = v2
    var v18 int16 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
