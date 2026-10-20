package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    v1 := []uint32{0, 878, 4294967295, 2}
    var v2 uint64 = 4294967295
    var v3 int16 = -32768
    var v4 int64 = -904
    var v5 uint8 = 255
    v6 := []uint32{}
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := GenS1{F0: (v5 << ((13) & 63))}
            _ = fi8
            _ = v9
        }
        var v10 uint64 = v2
        v5 += v5
        _ = fi7
        _ = v10
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        if len(v6) > 0 {
            v6[0] = 328
        }
        for fi12 := 0; fi12 < 6; fi12++ {
            if (int(v5) < (fi12 * fi12)) {
                var v13 string = (("" + "\n\t") + ("a" + "foo bar"))
                _ = v13
            }
            fmt.Println(fi11, fi12)
            _ = fi12
        }
        _ = fi11
    }
    v5 = (v5 * v5)
    v3 -= -32768
    fmt.Println(v3, v4)
    v14 := GenS1{F0: 58}
    v5 = (v5 ^ (v5 + v5))
    var v15 int16 = 32767
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println(v15, v5, v15)
    if (0 > 42) {
        v2 += uint64(v5)
        v15 -= 32767
    } else {
        v15 = 812
        for fi16 := 0; fi16 < 9; fi16++ {
            var v17 int32 = 0
            _ = fi16
            _ = v17
        }
        v6 = append(v6, 2147483648)
    }
    var v18 float64 = 1e10
    v15 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
