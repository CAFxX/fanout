package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 int64 }
func genFn0(p0 int, p1 uint8, p2 uint64) bool {
    fmt.Println(p2, p0)
    fmt.Println(p2, p1, p2)
    var v7 uint8 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return true
}

func main() {
    v1 := []uint32{549, 2147483648, 0, 2}
    var v2 int16 = 1000
    var v3 uint8 = 128
    var v4 uint32 = 693
    var v5 uint = 659
    v6 := []string{"foo bar", ("hello" + "foo bar"), "foo bar", ("\n\t" + ("hello" + "hello"))}
    var v8 float32 = 3.14159
    if (629 <= 1) {
        v9 := GenS1{F0: false}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v9
    }
    v8 -= 0.1
    v10 := GenS1{F0: (601 > 127)}
    v5 = (v5 / ((v5) | 1))
    for fi11 := 0; fi11 < 18; fi11++ {
        for fi12 := 0; fi12 < 1; fi12++ {
            if len(v1) > 0 {
                v1[0] = v4
            }
            fmt.Println(fi12, v5, v4)
            var v13 bool = true
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    var v14 int8 = 42
    v15 := GenS1{F0: true}
    if len(v6) > 0 {
        v6[0] = "hello"
    }
    v16 := GenS1{F0: ((!true) && true)}
    v17 := GenS2{F0: v4, F1: 42}
    v14 -= 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
