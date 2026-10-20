package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    v1 := []uint32{42, 0, 489, 0}
    v2 := make(map[int]int)
    var v3 bool = true
    var v4 int = -9223372036854775808
    var v5 int32 = -331
    v4 += (v4 - v4)
    var v6 int16 = -1
    v3 = (!v3)
    v3 = true
    var v7 int8 = 1
    for fi8 := 0; fi8 < 4; fi8++ {
        v9 := GenS1{F0: -229}
        v3 = v3
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi8
        _ = v9
    }
    if (!false) {
        if (v4 != (v4 & (v4 % ((v4) | 1)))) {
            if (false && (v3 && v3)) {
                _ = len(v2)
                var v10 uint32 = 2147483648
                _ = v10
            } else {
                var v11 uint32 = 764
                fmt.Println(v4, v4)
                switch (v4 >> ((1) & 63)) {
                case 127:
                    for fi12 := 0; fi12 < 14; fi12++ {
                        for fi13 := 0; fi13 < 16; fi13++ {
                            fmt.Println(v4, v5)
                            v7 += (v7 / (((v7 / (((^(v7))) | 1))) | 1))
                            var v14 uint8 = 128
                            _ = fi13
                            _ = v14
                        }
                        _ = fi12
                    }
                }
                _ = v11
            }
            var v15 int8 = v7
            _ = v15
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
