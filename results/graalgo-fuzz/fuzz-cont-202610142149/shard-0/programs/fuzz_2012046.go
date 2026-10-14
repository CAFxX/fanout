package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 int8; F2 uint64 }
type GenS4 struct { F0 uint16 }

func main() {
    v1 := []int{32767, 0, -971, -1}
    v2 := []uint32{2147483648}
    var v3 bool = (255 != 128)
    var v4 uint64 = 607
    var v5 uint64 = (v4 << ((404) & 63))
    v6 := make(map[string]uint32)
    var v7 string = "foo bar"
    for fi8 := 0; fi8 < 1; fi8++ {
        var v9 int32 = -29
        _ = fi8
        _ = v9
    }
    var v10 string = "\n\t"
    v11 := GenS1{F0: 0}
    var v12 uint8 = 42
    for fi13 := 0; fi13 < 8; fi13++ {
        v3 = (!("hello" <= v10))
        v14 := GenS2{F0: 255}
        for fi15 := 0; fi15 < 11; fi15++ {
            v5 -= v4
            v16 := GenS1{F0: int32(v5)}
            if ((fi13 << ((780) & 63)) != fi13) {
                v6[v7] = 4294967295
                var v17 float32 = -2.25
                _ = v17
            }
            _ = fi15
            _ = v16
        }
        _ = fi13
        _ = v14
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        for fi19 := 0; fi19 < 10; fi19++ {
            var v20 uint32 = 65535
            var v21 float64 = 3.14159
            v22 := GenS3{F0: 1, F1: 100, F2: 763}
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
        v12 += (v12 ^ (v12 >> ((65535) & 63)))
        v10 = "foo bar"
        _ = fi18
    }
    v3 = v3
    for fi23 := 0; fi23 < 18; fi23++ {
        v24 := GenS4{F0: 321}
        for fi25 := 0; fi25 < 13; fi25++ {
            v6[v7] = 2
            var v26 uint16 = 42
            _ = fi25
            _ = v26
        }
        _ = fi23
        _ = v24
    }
    var v27 int64 = 933
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
