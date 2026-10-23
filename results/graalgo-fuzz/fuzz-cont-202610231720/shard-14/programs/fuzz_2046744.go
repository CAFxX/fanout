package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string }
func genFn0(p0 uint32, p1 int64, p2 uint) int16 {
    v7 := []int{989}
    for fi8 := 0; fi8 < 17; fi8++ {
        v9 := GenS1{F0: 127}
        if (!true) {
            p0 = p0
            p2 = 42
            var v10 uint8 = 255
            _ = v10
        }
        _ = fi8
        _ = v9
    }
    fmt.Println(p1, p1)
    p2 = 42
    var v11 int8 = -62
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v7
    return 42
}

func main() {
    var v1 uint8 = 38
    v2 := []uint32{}
    v3 := make(map[int]int)
    v4 := make(map[uint32]int)
    v5 := make(map[int]string)
    var v6 uint = 736
    var v12 uint16 = 973
    v13 := GenS2{F0: ""}
    fmt.Println(v12)
    fmt.Println(v6, v6, v12)
    if true {
        var v14 float64 = 0.0
        v15 := GenS1{F0: (v1 & (v1 & uint8(v12)))}
        v12 += v12
        _ = v14
        _ = v15
    } else {
        v1 -= v1
    }
    delete(v3, -423)
    var v16 float32 = 1e10
    if (true || (1 <= 4294967295)) {
        var v17 bool = (32767 < -32768)
        for fi18 := 0; fi18 < 18; fi18++ {
            v1 = (v1 << ((v6) & 63))
            v5[-9223372036854775808] = ("foo bar" + ("foo bar" + ("" + "\n\t")))
            v1 = v1
            _ = fi18
        }
        fmt.Println(v16, v12)
        _ = v17
    }
    fmt.Println(v16, v16)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
