package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 bool; F1 int }
func genFn0(p0 int) int {
    v6 := GenS1{F0: 42, F1: 127}
    v7 := GenS1{F0: 0, F1: -75}
    _ = p0
    _ = v6
    _ = v7
    return p0
}

func main() {
    var v1 string = ("\n\t" + "")
    v2 := []uint32{}
    var v3 string = "\n\t"
    var v4 string = "hello"
    var v5 uint32 = 980
    var v8 uint = 18446744073709551615
    v2 = append(v2, v5)
    var v9 int32 = -2147483648
    if true {
        if false {
            fmt.Println(v4, v5)
            var v10 uint8 = 1
            _ = v10
        } else {
            if false {
                v4 = "foo bar"
            } else {
                if false {
                    v5 = v5
                    for fi11 := 0; fi11 < 10; fi11++ {
                        var v12 uint16 = 482
                        v3 = "a"
                        v2 = append(v2, v5)
                        _ = fi11
                        _ = v12
                    }
                } else {
                    if true {
                        var v13 float64 = -2.25
                        var v14 string = "foo bar"
                        _ = v13
                        _ = v14
                    } else {
                        v15 := GenS2{F0: false, F1: -9223372036854775808}
                        _ = v15
                    }
                }
                var v16 int32 = -472
                var v17 string = (("foo bar" + "\n\t") + v4)
                _ = v16
                _ = v17
            }
            v9 -= (v9 ^ v9)
        }
        v2 = append(v2, 4294967295)
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        var v19 float32 = (0.1 / 1.0)
        fmt.Println(fi18, v8)
        _ = fi18
        _ = v19
    }
    fmt.Println(v5)
    var v20 int16 = 0
    var v21 uint8 = 1
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
