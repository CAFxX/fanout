package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int32 }
func genFn0(p0 int32, p1 uint64) uint64 {
    var v7 string = ("" + "foo bar")
    v8 := GenS1{F0: (^(p1)), F1: 0}
    v9 := GenS1{F0: p1, F1: 178}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return (p1 | 4294967295)
}

func main() {
    var v1 string = "\n\t"
    var v2 float32 = 0.0
    var v3 uint64 = 42
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]uint32)
    v6 := make(map[string]string)
    var v10 string = v1
    var v11 uint8 = 40
    fmt.Println(v2)
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 int8 = 127
        fmt.Println(fi12)
        for fi14 := 0; fi14 < 5; fi14++ {
            if true {
                v13 -= v13
            } else {
                v15 := GenS1{F0: (v3 * v3), F1: 255}
                v16 := GenS2{F0: -747}
                if true {
                    if (v3 == v3) {
                        var v17 string = (("\n\t" + (v10 + "")) + ("foo bar" + ((((("" + v1) + v10) + (("foo bar" + "foo bar") + (((v10 + v1) + v1) + ((("hello" + v1) + v1) + v10)))) + v1) + (v1 + "a"))))
                        var v18 float64 = -2.25
                        _ = v17
                        _ = v18
                    }
                    if true {
                        _ = len(v6)
                    } else {
                        var v19 uint64 = v3
                        _ = v19
                    }
                } else {
                    var v20 uint8 = v11
                    _ = v20
                }
                var v21 bool = false
                _ = v15
                _ = v16
                _ = v21
            }
            _ = fi14
        }
        _ = fi12
        _ = v13
    }
    v2 = v2
    var v22 uint32 = 4294967295
    v1 = v10
    v23 := GenS1{F0: 2, F1: 400}
    var v24 uint = 802
    v22 = (v22 / ((v22) | 1))
    v24 = (v24 + (v24 / ((v24) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
