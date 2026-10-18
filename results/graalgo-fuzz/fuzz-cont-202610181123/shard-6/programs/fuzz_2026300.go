package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 uint64) uint8 {
    fmt.Println(p0)
    var v6 uint8 = 0
    _ = p0
    _ = v6
    return v6
}
func genFn1() int64 {
    v7 := []int{9223372036854775807}
    var v8 uint32 = 794
    var v9 int64 = 0
    if (42 != 870) {
        if len(v7) > 0 {
            v7[0] = 32767
        }
        if len(v7) > 0 {
            v7[0] = 1
        }
    } else {
        var v10 int8 = 127
        _ = v10
    }
    _ = v7
    _ = v8
    _ = v9
    return (v9 << ((254) & 63))
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint = 65535
    var v3 int8 = -117
    v4 := []uint32{65535}
    v5 := []uint32{0, 957}
    if (-1 < 9223372036854775807) {
        var v11 int16 = 888
        _ = v11
    } else {
        for fi12 := 0; fi12 < 18; fi12++ {
            v2 -= 1
            v2 = v2
            _ = fi12
        }
        v13 := GenS1{F0: v3}
        _ = v13
    }
    var v14 float64 = (3.14159 / 0.1)
    var v15 float64 = (v14 - float64(0.1))
    v14 -= ((float64(-2.25) * v14) - v15)
    for fi16 := 0; fi16 < 11; fi16++ {
        v2 += 841
        _ = fi16
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        if false {
            v15 = 0.1
        }
        var v18 float32 = (((3.14159 - -2.25) + float32((3.14159 / v15))) * (1e10 / (((1.5 * 0.1) + 0.0) - 0.0)))
        _ = fi17
        _ = v18
    }
    if (true || (!true)) {
        var v19 uint = 1
        _ = v19
    } else {
        v20 := GenS1{F0: v3}
        var v21 int16 = 42
        _ = v20
        _ = v21
    }
    v14 = 0.0
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v22 int16 = 0
    var v23 int = 830
    v1[(v23 >> ((v2) & 63))] = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
