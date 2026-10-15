package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
func genFn0(p0 uint64) uint16 {
    v7 := []int{70, -1, 32768}
    fmt.Println(p0)
    for fi8 := 0; fi8 < 20; fi8++ {
        for fi9 := 0; fi9 < 2; fi9++ {
            p0 += (p0 - p0)
            _ = fi9
        }
        p0 = p0
        _ = fi8
    }
    _ = p0
    _ = v7
    return 242
}
func genFn1() int32 {
    var v10 uint8 = 1
    v10 = 10
    v10 = 0
    var v11 bool = (32767 == 853)
    _ = v10
    _ = v11
    return 494
}

func main() {
    var v1 uint16 = 65535
    var v2 int16 = 18
    var v3 string = ("a" + "foo bar")
    v4 := make(map[uint32]string)
    var v5 float32 = 1.5
    v6 := make(map[uint32]int)
    var v12 uint32 = 667
    var v13 int32 = 42
    v13 = genFn1()
    v14 := GenS1{F0: (v13 / (((^(v13))) | 1)), F1: 802}
    var v15 bool = true
    var v16 int64 = 9223372036854775807
    var v17 float32 = (3.14159 - v5)
    v18 := GenS1{F0: (v13 % ((-370) | 1)), F1: genFn0(95)}
    v5 = float32(((0.0 + (3.14159 * (0.0 / 1.5))) - 3.14159))
    var v19 uint16 = (v1 + 65535)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
