package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int64 }
func genFn0(p0 uint16) int64 {
    v4 := []int{42, 0}
    p0 += p0
    v5 := GenS1{F0: (p0 / (((p0 << ((4294967295) & 63))) | 1)), F1: -1, F2: 360}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v6 uint8 = 233
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return -904
}

func main() {
    var v1 bool = false
    var v2 uint8 = 255
    v3 := []int{}
    v2 += 255
    var v7 string = "foo bar"
    var v8 int64 = 1
    var v9 uint32 = 2
    _ = len(v3)
    var v10 int16 = 0
    v1 = v1
    for fi11 := 0; fi11 < 20; fi11++ {
        if (!v1) {
            var v12 uint = 255
            v7 = "foo bar"
            for fi13 := 0; fi13 < 5; fi13++ {
                switch fi13 {
                case 1:
                    if (!(42 <= (v12 << (((v12 ^ 18446744073709551615)) & 63)))) {
                        var v14 int8 = 42
                        _ = v14
                    } else {
                        v1 = (-269 != -892)
                    }
                default:
                    var v15 int16 = (-(v10))
                    _ = v15
                }
                switch (fi13 + (^(fi11))) {
                case 127:
                    var v16 int64 = genFn0(665)
                    _ = v16
                case -2147483648:
                    var v17 uint16 = 2
                    _ = v17
                case 128:
                    v9 += (^(v9))
                default:
                    var v18 uint32 = 65535
                    _ = v18
                }
                if v1 {
                    if (genFn0(262) == 0) {
                        v19 := GenS1{F0: 1, F1: 100, F2: (v8 | (-(v8)))}
                        v10 = v10
                        _ = v19
                    } else {
                        var v20 int64 = 1
                        var v21 int8 = 0
                        v10 = int16(v9)
                        _ = v20
                        _ = v21
                    }
                }
                v22 := GenS1{F0: 255, F1: -1, F2: v8}
                v2 = (^(v2))
                _ = fi13
                _ = v22
            }
            _ = v12
        }
        _ = fi11
    }
    v23 := GenS1{F0: 951, F1: -1, F2: v8}
    if len(v3) > 0 {
        v3[0] = 32767
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
