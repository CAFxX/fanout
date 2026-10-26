package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
func genFn0(p0 int16, p1 int32, p2 int16) int {
    v4 := GenS1{F0: 42, F1: 42}
    var v5 int8 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 292
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = 1.5
    var v3 uint8 = 128
    fmt.Println(v3, v2)
    v6 := []int{1}
    var v7 uint8 = v3
    var v8 float32 = (0.1 / ((-2.25 / 1.5) / -2.25))
    v8 = v8
    var v9 int = 2147483647
    v2 -= 1e10
    if (-9223372036854775808 < -415) {
        delete(v1, ((("" + "foo bar") + (((("a" + ("" + "foo bar")) + "a") + "hello") + "hello")) + ((("a" + "hello") + ("foo bar" + "\n\t")) + ("foo bar" + (("hello" + "hello") + (("a" + "hello") + "a"))))))
        if false {
            var v10 float32 = (((v8 + v8) / v8) / 1.0)
            v10 -= v10
            _ = v10
        }
        if (v2 > (((v2 + 1.5) * float64(v8)) / float64((1e10 / v8)))) {
            v11 := GenS1{F0: 0, F1: 1000}
            for fi12 := 0; fi12 < 6; fi12++ {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = fi12
            }
            v8 = v8
            _ = v11
        }
    }
    v2 = 0.0
    if len(v6) > 0 {
        v6[0] = genFn0(0, 2147483647, 0)
    }
    v6 = append(v6, -536)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
