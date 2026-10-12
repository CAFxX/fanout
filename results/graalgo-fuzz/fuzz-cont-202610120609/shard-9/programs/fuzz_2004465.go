package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
func genFn0(p0 uint) int8 {
    p0 -= (p0 - p0)
    var v7 float64 = 0.0
    v7 += float64((0.1 + 1e10))
    fmt.Println(v7, p0)
    _ = p0
    _ = v7
    return -15
}

func main() {
    var v1 int64 = 151
    v2 := make(map[int]uint32)
    v3 := []uint32{}
    var v4 float32 = 1e10
    var v5 int32 = int32(v1)
    var v6 int8 = 66
    var v8 float64 = 1e10
    var v9 uint = 82
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 uint = 0
        _ = fi10
        _ = v11
    }
    v8 -= float64(1.5)
    v12 := GenS1{F0: 433, F1: 128}
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 uint8 = 128
        _ = fi13
        _ = v14
    }
    switch ((("" + "a") + ("foo bar" + ("foo bar" + "foo bar"))) + "hello") {
    case "":
        v6 = v6
    default:
        v6 += (-(v6))
    }
    v1 = (^(v1))
    v5 = v5
    delete(v2, -1)
    fmt.Println(v9, v4, v6)
    for fi15 := 0; fi15 < 10; fi15++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi15
    }
    var v16 uint = v9
    v1 += (v1 ^ 402)
    v1 -= 9223372036854775807
    var v17 int64 = v1
    fmt.Println(v1)
    if ((false && ((32767 != 0) || true)) || true) {
        for fi18 := 0; fi18 < 20; fi18++ {
            for fi19 := 0; fi19 < 16; fi19++ {
                var v20 uint64 = 0
                fmt.Println(v17, v8, v4)
                v8 = v8
                _ = fi19
                _ = v20
            }
            var v21 bool = true
            for fi22 := 0; fi22 < 10; fi22++ {
                v3 = append(v3, 0)
                if true {
                    v16 = v9
                    var v23 int16 = -32768
                    _ = v23
                } else {
                    v17 += v17
                    v6 = v6
                    var v24 float64 = (float64(v4) * 1e10)
                    _ = v24
                }
                var v25 uint8 = 0
                _ = fi22
                _ = v25
            }
            _ = fi18
            _ = v21
        }
        v1 -= v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
