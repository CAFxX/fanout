package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint32 }

func main() {
    var v1 string = "a"
    v2 := make(map[uint32]int)
    var v3 float64 = ((float64((0.1 - -2.25)) - 1.5) + -2.25)
    var v4 int32 = 0
    v5 := GenS1{F0: 255, F1: 32767, F2: 2}
    fmt.Println(v3, v3)
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 uint32 = 1
        v3 = (-2.25 - (1e10 - (v3 - (1.5 + 0.1))))
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 uint8 = 255
            var v10 uint64 = 0
            if ((18446744073709551615 != 995) && (v1 >= ((v1 + v1) + ((v1 + ((v1 + "foo bar") + v1)) + "foo bar")))) {
                var v11 bool = (v1 >= "foo bar")
                var v12 int8 = 127
                _ = v11
                _ = v12
            } else {
                var v13 uint32 = (^(v7))
                _ = v13
            }
            v1 = v1
            var v14 int8 = -36
            _ = fi8
            _ = v10
            _ = v14
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    var v15 uint64 = 0
    v16 := GenS1{F0: 218, F1: 0, F2: 4294967295}
    v4 += v4
    var v17 uint64 = 18446744073709551615
    var v18 int16 = 1000
    for fi19 := 0; fi19 < 17; fi19++ {
        if (!((float32((v3 * v3)) / 1e10) < 0.0)) {
            v2[981] = (^(fi19))
        } else {
            var v20 uint64 = v15
            fmt.Println(v4, v20, fi19)
            _ = v20
        }
        var v21 uint16 = 32767
        _ = fi19
        _ = v21
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        v3 = 0.1
        fmt.Println(v18, v1)
        _ = fi22
    }
    for fi23 := 0; fi23 < 1; fi23++ {
        v1 = v1
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
