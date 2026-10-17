package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 int64, p1 int16) int16 {
    var v5 uint16 = 255
    var v6 uint8 = 255
    v5 -= 1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (p1 >> ((326) & 63))
}
func genFn1(p0 uint32, p1 bool) int32 {
    var v7 uint64 = 204
    v8 := GenS1{F0: (1e10 != -2.25)}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -1
}

func main() {
    v1 := []uint32{}
    v2 := []int{-2147483648, 32768, 568}
    v3 := make(map[uint32]uint32)
    var v4 float32 = float32(float64(0.1))
    v4 = v4
    _ = len(v3)
    for fi9 := 0; fi9 < 14; fi9++ {
        if (fi9 >= fi9) {
            v4 -= (-2.25 - float32(float64(0.1)))
            v4 -= v4
        }
        _ = fi9
    }
    if true {
        v4 += (v4 / 1.0)
    }
    if false {
        delete(v3, 2147483647)
        for fi10 := 0; fi10 < 10; fi10++ {
            v2 = append(v2, fi10)
            fmt.Println(v4, v4)
            var v11 int64 = 9223372036854775807
            _ = fi10
            _ = v11
        }
        var v12 float32 = (1.5 - ((v4 / float32(-2.25)) * float32((float64(float32(((3.14159 + (((0.0 / ((((1e10 / 1e10) + 1e10) + 1.5) - float64((3.14159 * v4)))) + float64(v4)) / ((3.14159 / 1e10) - 1e10))) + 1.5))) - 0.0))))
        _ = v12
    } else {
        var v13 uint64 = 4294967295
        _ = v13
    }
    delete(v3, 65535)
    v4 -= v4
    var v14 float32 = v4
    if false {
        switch 0 {
        case 1:
            if (4294967295 > 167) {
                v14 -= 3.14159
            }
        case -1:
            fmt.Println(v4, v14)
        default:
            fmt.Println(v4, v14)
        }
        var v15 int8 = 127
        _ = v15
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
