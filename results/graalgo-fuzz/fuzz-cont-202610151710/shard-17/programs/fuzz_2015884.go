package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 uint64, p1 int8) int64 {
    v5 := GenS1{F0: 42}
    var v6 int = 257
    var v7 uint8 = 255
    if ((p0 + p0) < (p0 << ((42) & 63))) {
        v7 += uint8(p0)
        v8 := GenS1{F0: 444}
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 uint32 = 1
    var v2 bool = true
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]int)
    var v9 int16 = 1000
    fmt.Println(v9, v9, v2)
    var v10 uint8 = 42
    var v11 int16 = 32767
    v1 = 4294967295
    for fi12 := 0; fi12 < 12; fi12++ {
        var v13 int16 = (v11 ^ (v9 >> ((2) & 63)))
        v14 := []int{(fi12 % ((fi12) | 1))}
        _ = fi12
        _ = v13
        _ = v14
    }
    if ("\n\t" < "foo bar") {
        v11 += (v9 ^ v11)
        if false {
            v1 += 1
            v1 = 0
        } else {
            v15 := []int{32767, 9223372036854775807, 0}
            _ = v15
        }
    }
    v1 = (v1 - (v1 / ((2147483647) | 1)))
    for fi16 := 0; fi16 < 10; fi16++ {
        v17 := GenS1{F0: v1}
        if v2 {
            for fi18 := 0; fi18 < 7; fi18++ {
                v3[v1] = true
                var v19 float64 = 1.5
                v19 += float64(3.14159)
                _ = fi18
                _ = v19
            }
            v10 = v10
        } else {
            v11 -= (v9 % ((v11) | 1))
            v4[(v1 << ((uint(v1)) & 63))] = (fi16 >> ((683) & 63))
        }
        _ = fi16
        _ = v17
    }
    v20 := GenS1{F0: (v1 ^ (v1 % (((v1 << ((4294967295) & 63))) | 1)))}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v9
}
