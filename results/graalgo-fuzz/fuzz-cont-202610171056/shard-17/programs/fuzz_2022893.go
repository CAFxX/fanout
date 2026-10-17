package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 uint16; F1 uint64 }
func genFn0() uint64 {
    v4 := []int{127}
    for fi5 := 0; fi5 < 20; fi5++ {
        var v6 uint32 = 2147483647
        v7 := GenS1{F0: ((v6 / (((v6 % (((^(v6))) | 1))) | 1)) == (v6 ^ v6)), F1: "\n\t"}
        if false {
            for fi8 := 0; fi8 < 20; fi8++ {
                v6 = 863
                _ = fi8
            }
            v6 = (v6 + v6)
        } else {
            v9 := GenS1{F0: true, F1: (("" + "\n\t") + "hello")}
            v6 = v6
            _ = v9
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = v4
    return 42
}
func genFn1(p0 uint64, p1 string) int8 {
    if ((((42 == 92) || (true || (true || false))) || false) || true) {
        var v10 int32 = -1
        _ = v10
    }
    fmt.Println(p0, p1)
    p0 -= 1
    fmt.Println(p1)
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 int = 2147483647
    var v2 uint32 = 233
    v3 := make(map[uint32]string)
    var v11 uint64 = 1
    if (2147483647 == (v1 & -2147483648)) {
        var v12 float32 = 0.0
        for fi13 := 0; fi13 < 3; fi13++ {
            v11 = v11
            _ = fi13
        }
        _ = v12
    }
    if (!(((3.14159 / 1.5) * (1.5 / (-2.25 - 1.5))) == -2.25)) {
        var v14 uint8 = 1
        for fi15 := 0; fi15 < 1; fi15++ {
            v16 := GenS2{F0: 32767, F1: (v11 * (v11 % ((v11) | 1)))}
            var v17 float64 = (3.14159 + (1.5 + 1e10))
            v18 := v3[v2]
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
        _ = v14
    }
    var v19 float32 = 1e10
    var v20 int16 = 0
    v1 += 502
    v2 = v2
    v1 += v1
    v21 := v3[v2]
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
