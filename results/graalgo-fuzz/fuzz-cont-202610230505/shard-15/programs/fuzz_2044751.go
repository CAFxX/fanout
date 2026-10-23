package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }

func main() {
    var v1 int8 = -36
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]bool)
    var v4 uint32 = 65535
    var v5 int = 32767
    fmt.Println(v4)
    if true {
        var v6 int16 = -1
        _ = v6
    } else {
        v7 := v2[v4]
        v4 -= v4
        _ = v7
    }
    var v8 int64 = -9223372036854775808
    var v9 string = "\n\t"
    fmt.Println(v5)
    v5 -= (^(v5))
    v9 = ""
    fmt.Println(v8, v9)
    v9 = v9
    v3[(^(v4))] = false
    var v10 uint8 = 127
    var v11 int = v5
    var v12 float64 = (1.5 + -2.25)
    for fi13 := 0; fi13 < 14; fi13++ {
        fmt.Println(v12, v9, fi13)
        for fi14 := 0; fi14 < 15; fi14++ {
            v15 := GenS1{F0: (^(v10))}
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    _ = len(v2)
    var v16 uint8 = v10
    if ((v11 ^ v11) == 575) {
        for fi17 := 0; fi17 < 9; fi17++ {
            v16 += 42
            v10 = uint8(v4)
            var v18 string = v9
            _ = fi17
            _ = v18
        }
        v16 += (v10 + 27)
        fmt.Println(v4, v16)
    } else {
        v19 := GenS2{F0: 40}
        var v20 int8 = -1
        _ = v19
        _ = v20
    }
    v11 += v11
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
