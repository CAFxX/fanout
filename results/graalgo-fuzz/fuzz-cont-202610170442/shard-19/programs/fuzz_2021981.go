package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 uint64; F1 uint }
func genFn0() uint64 {
    var v5 int16 = 42
    v5 = 1000
    v6 := GenS1{F0: 26}
    v7 := GenS2{F0: 2, F1: 242}
    if (1 > 65535) {
        for fi8 := 0; fi8 < 7; fi8++ {
            if false {
                var v9 bool = (1 >= 42)
                v10 := GenS3{F0: 18446744073709551615, F1: 255}
                _ = v10
                _ = v9
            }
            if false {
                v5 = (^(v5))
            }
            _ = fi8
        }
        if (650 <= 0) {
            if (true && (-2147483648 <= 1)) {
                v11 := GenS1{F0: 255}
                _ = v11
            }
        } else {
            var v12 uint8 = 42
            _ = v12
        }
    }
    _ = v5
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := []string{(("a" + ("\n\t" + "hello")) + ""), ("" + "a"), "hello"}
    v2 := make(map[uint32]string)
    v3 := []uint32{2, 2147483647, 343}
    var v4 int8 = -26
    fmt.Println(v4)
    for fi13 := 0; fi13 < 9; fi13++ {
        v4 = (v4 << ((255) & 63))
        if (0 >= 2147483647) {
            if false {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
        }
        v14 := GenS2{F0: 42, F1: 816}
        _ = fi13
        _ = v14
    }
    var v15 int16 = -32768
    v16 := v2[65535]
    fmt.Println(v15)
    v17 := GenS3{F0: 0, F1: 255}
    var v18 int16 = v15
    var v19 float64 = ((1e10 / float64(float32(0.1))) - (1.5 + 1.5))
    switch 32768 {
    case 128:
        v16 = v16
    }
    v4 -= 1
    var v20 int16 = v18
    if false {
        if len(v1) > 0 {
            v1[0] = "a"
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
