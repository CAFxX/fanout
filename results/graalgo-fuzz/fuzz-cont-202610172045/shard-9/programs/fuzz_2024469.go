package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8; F1 uint8 }
type GenS5 struct { F0 uint; F1 uint32 }
type GenS6 struct { F0 int8 }
func genFn0(p0 string, p1 int64) uint16 {
    var v7 int8 = 42
    if false {
        for fi8 := 0; fi8 < 9; fi8++ {
            var v9 uint16 = 1
            _ = fi8
            _ = v9
        }
    } else {
        fmt.Println(p1, p1, v7)
    }
    v10 := []int{9223372036854775807, 1}
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    return 32767
}

func main() {
    var v1 float32 = 3.14159
    v2 := []int{}
    v3 := make(map[int]int)
    var v4 int16 = -32768
    var v5 string = "a"
    v6 := make(map[uint32]uint32)
    for fi11 := 0; fi11 < 1; fi11++ {
        fmt.Println(fi11, v4)
        v12 := GenS1{F0: 65535, F1: 65535, F2: (fi11 << ((2) & 63))}
        for fi13 := 0; fi13 < 1; fi13++ {
            v14 := GenS2{F0: 531}
            if (((("" + "a") + (("foo bar" + (((("\n\t" + "foo bar") + "") + "\n\t") + "\n\t")) + "hello")) + "a") <= v5) {
                if true {
                    if (2147483648 >= 0) {
                        var v15 uint = 18446744073709551615
                        v16 := GenS3{F0: 153}
                        _ = v15
                        _ = v16
                    }
                } else {
                    v17 := v6[521]
                    v4 += v4
                    _ = v17
                }
                var v18 string = "hello"
                v18 = v18
                _ = v18
            } else {
                v19 := GenS4{F0: 255, F1: 255}
                v1 = v1
                _ = v19
            }
            v20 := GenS5{F0: 42, F1: 1}
            _ = fi13
            _ = v14
            _ = v20
        }
        _ = fi11
        _ = v12
    }
    var v21 int16 = (v4 * (v4 >> ((1) & 63)))
    fmt.Println(v5, v1)
    v22 := GenS6{F0: 127}
    v3[-650] = 42
    v23 := v3[-1]
    v24 := GenS5{F0: 42, F1: 42}
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
