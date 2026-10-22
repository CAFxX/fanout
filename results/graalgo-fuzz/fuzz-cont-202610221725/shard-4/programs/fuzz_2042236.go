package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8 }
func genFn0(p0 uint8) int32 {
    p0 = 127
    p0 -= (p0 & p0)
    _ = p0
    return -1
}
func genFn1(p0 int16, p1 uint) int16 {
    p1 -= p1
    v4 := []int{}
    var v5 int8 = 127
    _ = len(v4)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (p0 % (((p0 << ((p1) & 63))) | 1))
}

func main() {
    var v1 float32 = 0.0
    var v2 int32 = 0
    var v3 float64 = float64(v1)
    var v6 uint = 1
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := []int{924}
        var v9 uint = v6
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(v6, v1, v3)
    v10 := GenS1{F0: -1}
    v11 := GenS2{F0: genFn1(842, 0)}
    v1 = 0.1
    for fi12 := 0; fi12 < 3; fi12++ {
        var v13 uint16 = 42
        fmt.Println(v6, v6)
        _ = fi12
        _ = v13
    }
    v14 := []int{32768, 87, 602, -867}
    if len(v14) > 0 {
        v14[0] = -1
    }
    v3 = v3
    for fi15 := 0; fi15 < 8; fi15++ {
        if ((("" + "foo bar") + ("" + "a")) != "foo bar") {
            v16 := GenS3{F0: 1}
            var v17 uint32 = 245
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            _ = v16
            _ = v17
        }
        _ = fi15
    }
    fmt.Println(v6, v3, v3)
    v1 = (v1 * 1e10)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v6
}
