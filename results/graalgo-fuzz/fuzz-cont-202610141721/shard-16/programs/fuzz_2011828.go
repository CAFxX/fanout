package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 string }

func main() {
    var v1 string = ("hello" + (("\n\t" + "") + "hello"))
    v2 := make(map[int]uint32)
    var v3 int16 = 1
    var v4 uint16 = 42
    v5 := make(map[int]bool)
    var v6 float32 = -2.25
    v3 -= v3
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 int16 = int16(fi7)
        _ = fi7
        _ = v8
    }
    v3 -= v3
    var v9 float64 = float64(float32(0.0))
    delete(v5, -9223372036854775808)
    v6 -= 0.0
    if false {
        v4 -= v4
    } else {
        for fi10 := 0; fi10 < 5; fi10++ {
            for fi11 := 0; fi11 < 6; fi11++ {
                v2[fi10] = 98
                v5[fi10] = ((9223372036854775807 <= 1) && false)
                v9 -= float64(v6)
                _ = fi11
            }
            _ = fi10
        }
    }
    fmt.Println(v9, v1, v1)
    fmt.Println(v4, v6)
    v12 := []int{32768, 42, 127}
    for fi13 := 0; fi13 < 7; fi13++ {
        var v14 int16 = (v3 ^ (v3 ^ (v3 ^ (v3 ^ v3))))
        if (1 == 1) {
            v4 += 255
            for fi15 := 0; fi15 < 3; fi15++ {
                v16 := GenS1{F0: 1000, F1: 4294967295, F2: ""}
                _ = fi15
                _ = v16
            }
            var v17 string = v1
            _ = v17
        }
        for fi18 := 0; fi18 < 16; fi18++ {
            fmt.Println(v4, v14)
            fmt.Println(v3, v1, v14)
            _ = fi18
        }
        _ = fi13
        _ = v14
    }
    if (49 <= 126) {
        var v19 uint = 18446744073709551615
        if (("\n\t" + "hello") > v1) {
            v20 := v2[452]
            v5[int(v4)] = false
            _ = v20
        }
        fmt.Println(v1, v6, v9)
        _ = v19
    } else {
        fmt.Println(v4, v9)
        var v21 int16 = 0
        var v22 int = 957
        _ = v21
        _ = v22
    }
    var v23 bool = ((!true) && (!(true || false)))
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
