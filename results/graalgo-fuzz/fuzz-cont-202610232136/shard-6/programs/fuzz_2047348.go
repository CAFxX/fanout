package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 bool }
type GenS2 struct { F0 uint32; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 uint32; F1 int16 }
func genFn0() int16 {
    v4 := []int{996, 32768, -505}
    v5 := GenS1{F0: 0, F1: 322, F2: (-928 == -822)}
    var v6 uint16 = 255
    v7 := GenS2{F0: 4294967295, F1: 873, F2: int64(v6)}
    fmt.Println(v6)
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 654
    v3 := []int{127, 127}
    v2 = (^(v2))
    if (-1 > 127) {
        v2 = (v2 ^ 176)
    } else {
        v8 := v1[uint32(v2)]
        _ = v8
    }
    fmt.Println(v2)
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := GenS3{F0: 115}
        _ = fi9
        _ = v10
    }
    v11 := GenS2{F0: 2147483647, F1: 42, F2: 1}
    v12 := GenS2{F0: 2, F1: 4294967295, F2: 42}
    v2 += 0
    if false {
        if true {
            if len(v3) > 0 {
                v3[0] = -255
            }
            v13 := GenS4{F0: -95, F1: 1}
            v14 := GenS5{F0: 236, F1: -610}
            _ = v13
            _ = v14
        } else {
            v2 += v2
            var v15 uint = 1
            _ = v15
        }
        var v16 string = "\n\t"
        _ = v16
    } else {
        var v17 float32 = 1.5
        v17 += v17
        _ = v17
    }
    switch (("" + "hello") + "\n\t") {
    case "hello":
        for fi18 := 0; fi18 < 1; fi18++ {
            var v19 float64 = (1.5 + 0.0)
            _ = fi18
            _ = v19
        }
    case "a":
        if (false && false) {
            v2 += v2
            var v20 int64 = 524
            v21 := v1[42]
            _ = v20
            _ = v21
        }
    }
    var v22 uint64 = 715
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v23 uint16 = uint16(v22)
    v2 += v2
    var v24 uint32 = 2147483648
    if (0 > -9223372036854775808) {
        fmt.Println(v24, v24)
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v3 = append(v3, 0)
    }
    switch "" {
    case "foo bar":
        if true {
            fmt.Println(v24, v24, v2)
        }
    }
    v2 -= 113
    var v25 string = "\n\t"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
