package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0() int {
    var v6 float32 = (3.14159 + (0.1 * 1.5))
    v7 := []int{341, 2147483647, 9223372036854775807}
    if false {
        if (862 == 0) {
            v8 := GenS1{F0: -886}
            if ("\n\t" != "foo bar") {
                v6 -= v6
                fmt.Println(v6)
                if (!true) {
                    v6 -= (v6 + (1e10 * (((((float32(float64(v6)) - (v6 - v6)) * v6) / 1e10) + v6) + (3.14159 * v6))))
                    var v9 bool = (32767 < 0)
                    _ = v9
                } else {
                    var v10 string = "a"
                    fmt.Println(v6, v10)
                    for fi11 := 0; fi11 < 8; fi11++ {
                        v10 = v10
                        v10 = "foo bar"
                        _ = fi11
                    }
                    _ = v10
                }
            } else {
                var v12 bool = (2 <= 255)
                v6 = -2.25
                _ = v12
            }
            _ = len(v7)
            _ = v8
        } else {
            v6 += v6
        }
    } else {
        v6 = v6
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        var v13 uint64 = 42
        _ = v13
    }
    _ = v6
    _ = v7
    return -420
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[string]bool)
    v3 := make(map[int]int)
    var v4 int16 = 32767
    v5 := make(map[uint32]string)
    fmt.Println(v1, v4)
    var v14 uint64 = 0
    v14 = 470
    v14 = (^(v14))
    var v15 int8 = 42
    var v16 float64 = ((float64(1e10) / v1) + (1.5 + v1))
    for fi17 := 0; fi17 < 17; fi17++ {
        v1 = v1
        _ = fi17
    }
    v18 := []int{genFn0(), genFn0(), 32768, -177}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
