package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
func genFn0(p0 uint32, p1 uint32, p2 uint64) int32 {
    p1 -= 712
    switch "hello" {
    case "\n\t":
        if (!(false && (!false))) {
            var v4 uint16 = 65535
            _ = v4
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return -2147483648
}

func main() {
    var v1 uint8 = 0
    var v2 bool = ((!true) && false)
    v3 := make(map[int]uint32)
    v5 := GenS1{F0: 32767, F1: 255}
    v6 := GenS1{F0: 142, F1: 1}
    v1 = 42
    if v2 {
        var v7 int = 128
        v3[v7] = 2
        _ = v7
    } else {
        v1 = 128
        v2 = (568 == 0)
        if (-128 <= 0) {
            v1 = (v1 / ((42) | 1))
            for fi8 := 0; fi8 < 1; fi8++ {
                if (423 >= 2147483647) {
                    var v9 bool = true
                    var v10 int8 = 42
                    delete(v3, fi8)
                    _ = v10
                    _ = v9
                } else {
                    var v11 float32 = -2.25
                    v12 := []int{}
                    v12 = append(v12, -154)
                    _ = v11
                    _ = v12
                }
                var v13 string = "hello"
                _ = fi8
                _ = v13
            }
        }
    }
    v1 += (v1 >> ((4294967295) & 63))
    fmt.Println(v1)
    fmt.Println(v2)
    var v14 float32 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
