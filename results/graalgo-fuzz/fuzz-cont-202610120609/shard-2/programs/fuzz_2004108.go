package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
func genFn0(p0 uint8, p1 bool, p2 int32) uint16 {
    v7 := []int{}
    var v8 bool = ((true || (42 <= 0)) && (p1 && true))
    fmt.Println(p2, p2)
    v8 = false
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return uint16(p0)
}

func main() {
    var v1 int16 = 1
    var v2 int32 = 741
    v3 := make(map[int]string)
    var v4 int = 718
    v5 := make(map[uint32]string)
    v6 := []string{"\n\t", "foo bar", "", "\n\t"}
    v4 -= v4
    var v9 int64 = -1
    if ((0.1 / 1.0) != 0.1) {
        if (!true) {
            var v10 float32 = -2.25
            _ = v10
        }
        var v11 string = ""
        _ = v11
    }
    var v12 bool = false
    var v13 int8 = 1
    fmt.Println(v1, v9, v2)
    fmt.Println(v4)
    if false {
        v14 := GenS1{F0: 1, F1: 42}
        v9 = (v9 & (v9 * v9))
        _ = v14
    } else {
        if (!((^(v1)) != 42)) {
            var v15 int32 = v2
            for fi16 := 0; fi16 < 10; fi16++ {
                fmt.Println(v15, fi16, fi16)
                _ = fi16
            }
            for fi17 := 0; fi17 < 6; fi17++ {
                v13 = (v13 ^ (v13 - (v13 + (v13 / ((v13) | 1)))))
                switch (fi17 / ((128) | 1)) {
                case 127:
                    v5[0] = "\n\t"
                default:
                    if (2147483648 < 2) {
                        _ = len(v5)
                    }
                }
                v18 := GenS1{F0: 1, F1: 2147483648}
                _ = fi17
                _ = v18
            }
            _ = v15
        }
    }
    v12 = v12
    v4 += (v4 % (((v4 << ((1) & 63))) | 1))
    v19 := v3[-276]
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
