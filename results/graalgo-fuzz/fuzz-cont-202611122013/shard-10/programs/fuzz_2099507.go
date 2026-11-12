package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 bool = (1 > 130)
    var v2 int = 356
    var v3 uint16 = 2
    v4 := []string{}
    v3 -= (^(v3))
    var v5 uint32 = 748
    _ = len(v4)
    for fi6 := 0; fi6 < 2; fi6++ {
        if (!(561 == 4294967295)) {
            v5 = (v5 - (v5 & (v5 + 65535)))
            if ((v3 / ((v3) | 1)) >= v3) {
                for fi7 := 0; fi7 < 12; fi7++ {
                    v8 := GenS1{F0: 4294967295, F1: -1}
                    v5 = 2147483648
                    _ = len(v4)
                    _ = fi7
                    _ = v8
                }
                var v9 float32 = ((3.14159 + 0.0) * (0.0 - 3.14159))
                if len(v4) > 0 {
                    v4[0] = ("" + "\n\t")
                }
                _ = v9
            }
            v5 += (v5 * v5)
        }
        v3 = v3
        var v10 float64 = 1e10
        _ = fi6
        _ = v10
    }
    _ = len(v4)
    v5 = (^(v5))
    var v11 int64 = 1
    var v12 uint16 = (v3 << ((215) & 63))
    v3 -= 255
    for fi13 := 0; fi13 < 4; fi13++ {
        fmt.Println(fi13, v5)
        v14 := GenS2{F0: 0, F1: ""}
        _ = fi13
        _ = v14
    }
    v3 -= 2
    switch (^(v2)) {
    case 127:
        if v1 {
            v15 := GenS3{F0: v12}
            _ = v15
        } else {
            v12 -= 42
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
