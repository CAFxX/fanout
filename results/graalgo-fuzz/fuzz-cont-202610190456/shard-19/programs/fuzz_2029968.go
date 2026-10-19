package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
func genFn0(p0 int32, p1 uint16, p2 int) bool {
    p2 = (p2 - (p2 + (^(p2))))
    p2 += p2
    fmt.Println(p1)
    v7 := []int{p2, 431, (^(p2)), p2}
    var v8 uint = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return true
}
func genFn1(p0 int16) int {
    var v9 uint8 = 17
    var v10 int8 = 100
    var v11 uint16 = 42
    var v12 int32 = 1
    v11 += v11
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return 32767
}

func main() {
    v1 := []string{("\n\t" + ""), ("hello" + "foo bar"), ""}
    var v2 string = ("foo bar" + (("\n\t" + "hello") + "foo bar"))
    var v3 float64 = (-2.25 + (1.5 + 0.0))
    var v4 int8 = 127
    var v5 uint32 = 415
    var v6 uint64 = 42
    var v13 float32 = 3.14159
    v2 = v2
    v4 -= 42
    v14 := GenS1{F0: 2, F1: 42}
    if false {
        v2 = v2
        for fi15 := 0; fi15 < 10; fi15++ {
            var v16 uint8 = 0
            _ = fi15
            _ = v16
        }
    } else {
        v5 += (v5 >> ((185) & 63))
        if len(v1) > 0 {
            v1[0] = v2
        }
    }
    _ = len(v1)
    v4 = v4
    v17 := GenS1{F0: 1, F1: (v4 << ((1) & 63))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
