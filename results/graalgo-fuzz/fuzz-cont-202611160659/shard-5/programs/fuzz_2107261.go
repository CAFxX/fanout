package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }

func main() {
    var v1 uint32 = 649
    var v2 uint8 = 2
    v3 := make(map[string]uint32)
    for fi4 := 0; fi4 < 16; fi4++ {
        for fi5 := 0; fi5 < 6; fi5++ {
            v2 -= v2
            var v6 float32 = 3.14159
            _ = fi5
            _ = v6
        }
        fmt.Println(fi4, v1)
        _ = fi4
    }
    v7 := []int{}
    if ((true && false) || (-124 >= 0)) {
        if len(v7) > 0 {
            v7[0] = -2147483648
        }
        var v8 uint16 = 65535
        _ = v8
    }
    v1 = 4294967295
    var v9 int16 = 42
    var v10 uint16 = uint16(v2)
    var v11 bool = true
    v10 = 255
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 bool = ("\n\t" <= "\n\t")
        if v13 {
            var v14 string = "\n\t"
            if v13 {
                var v15 bool = (((^(fi12)) >= (-(fi12))) || (true || v13))
                _ = v15
            } else {
                v13 = true
                if true {
                    v10 -= v10
                    var v16 int = (-(fi12))
                    v17 := GenS1{F0: -97, F1: 2}
                    _ = v16
                    _ = v17
                }
                for fi18 := 0; fi18 < 5; fi18++ {
                    v19 := GenS1{F0: -1, F1: 0}
                    var v20 float32 = ((-2.25 * 1e10) * (-2.25 * 1e10))
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            }
            var v21 int16 = 0
            _ = v14
            _ = v21
        } else {
            fmt.Println(fi12, v11)
        }
        v2 = uint8(v1)
        _ = fi12
        _ = v13
    }
    var v22 float32 = float32((0.1 / 0.1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v3
    _ = v7
    _ = v9
}
