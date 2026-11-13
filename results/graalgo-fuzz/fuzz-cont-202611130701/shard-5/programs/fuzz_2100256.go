package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 int64; F1 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int8; F1 uint16 }
func genFn0(p0 int32) uint64 {
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 uint16 = 255
        var v8 int64 = -116
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 int16 = -1
    v10 := GenS1{F0: 1, F1: ("a" != ("foo bar" + "hello"))}
    v11 := GenS2{F0: -9223372036854775808, F1: 42}
    v12 := GenS2{F0: 277, F1: 585}
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return 42
}

func main() {
    var v1 float32 = 3.14159
    var v2 int32 = 2147483647
    var v3 uint32 = 324
    var v4 float64 = -2.25
    v5 := make(map[string]uint32)
    v13 := v5["a"]
    var v14 uint32 = (v3 >> ((42) & 63))
    for fi15 := 0; fi15 < 20; fi15++ {
        var v16 int = -898
        _ = fi15
        _ = v16
    }
    v1 = (float32(float64(v1)) * v1)
    v17 := GenS3{F0: 9223372036854775807, F1: 9223372036854775807}
    v18 := []int{612, -282, 390}
    var v19 int = 1
    v20 := GenS4{F0: (true || false)}
    var v21 uint8 = 2
    for fi22 := 0; fi22 < 11; fi22++ {
        switch (^(v19)) {
        case 4294967295:
            for fi23 := 0; fi23 < 9; fi23++ {
                v24 := GenS1{F0: 547, F1: false}
                _ = fi23
                _ = v24
            }
        case 642:
            var v25 uint64 = 42
            _ = v25
        }
        var v26 float64 = 0.1
        v27 := GenS1{F0: 2147483647, F1: true}
        _ = fi22
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 6; fi28++ {
        var v29 int32 = (v2 / ((1) | 1))
        _ = fi28
        _ = v29
    }
    for fi30 := 0; fi30 < 5; fi30++ {
        var v31 uint8 = v21
        fmt.Println(v31, v3, v31)
        if true {
            if false {
                fmt.Println(v4, v1, v19)
                if ((42 >= 4294967295) && false) {
                    v32 := GenS5{F0: 127}
                    _ = v32
                }
                var v33 string = ""
                _ = v33
            } else {
                v3 -= v3
                var v34 int16 = 42
                _ = v34
            }
            v35 := GenS6{F0: int8(fi30), F1: 65535}
            _ = v35
        }
        _ = fi30
        _ = v31
    }
    var v36 string = ("\n\t" + "hello")
    v37 := GenS6{F0: 42, F1: 576}
    var v38 string = v36
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
