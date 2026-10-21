package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 bool }

func main() {
    var v1 uint32 = 2
    var v2 float32 = (((1.5 * 0.0) + (float32(float64((3.14159 * -2.25))) + 0.0)) + 1e10)
    v3 := make(map[uint32]int)
    var v4 bool = (!((("hello" + "hello") + ("\n\t" + ("\n\t" + ("\n\t" + "hello")))) < "hello"))
    var v5 uint = 0
    if true {
        var v6 uint8 = 255
        var v7 int32 = 0
        _ = v6
        _ = v7
    } else {
        v1 = v1
        v8 := GenS1{F0: 42, F1: 255, F2: true}
        _ = v8
    }
    v9 := []int{}
    for fi10 := 0; fi10 < 13; fi10++ {
        v9 = append(v9, (fi10 | 9223372036854775807))
        var v11 bool = ((968 != 32767) && (95 < -28))
        for fi12 := 0; fi12 < 5; fi12++ {
            fmt.Println(v4, fi12, v1)
            var v13 uint16 = 1
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    v14 := v3[4294967295]
    v2 -= (((0.0 + v2) + ((((0.1 * v2) * v2) - 1e10) - v2)) - v2)
    switch "" {
    case "":
        fmt.Println(v14, v2)
    case "a":
        v5 = 81
    case "foo bar":
        for fi15 := 0; fi15 < 1; fi15++ {
            v2 = v2
            fmt.Println(v5, v2, v2)
            if len(v9) > 0 {
                v9[0] = (v14 * (^(fi15)))
            }
            _ = fi15
        }
    default:
        fmt.Println(v5, v5, v14)
    }
    v1 -= (v1 % (((^(v1))) | 1))
    fmt.Println(v2, v1)
    v3[343] = v14
    v5 = v5
    for fi16 := 0; fi16 < 17; fi16++ {
        v1 = v1
        var v17 uint32 = (v1 % ((v1) | 1))
        var v18 float64 = (0.1 / (((((0.0 + (float64((float32(float64(float32(1e10))) - 0.1)) * 0.0)) + 0.0) / -2.25) - 3.14159) - (1.5 - -2.25)))
        _ = fi16
        _ = v17
        _ = v18
    }
    switch 618 {
    case 9223372036854775807:
        for fi19 := 0; fi19 < 20; fi19++ {
            fmt.Println(fi19)
            v1 = v1
            var v20 uint16 = 0
            _ = fi19
            _ = v20
        }
    case -365:
        if false {
            for fi21 := 0; fi21 < 10; fi21++ {
                v22 := v3[361]
                if (v4 && ((v5 / ((949) | 1)) < 743)) {
                    var v23 uint64 = 0
                    v24 := GenS1{F0: 1, F1: 29, F2: true}
                    _ = v23
                    _ = v24
                } else {
                    if true {
                        fmt.Println(v14, v22, v22)
                    }
                    v4 = v4
                }
                _ = fi21
                _ = v22
            }
            v2 -= v2
        }
    default:
        delete(v3, 139)
    }
    var v25 uint32 = (v1 + (v1 | 2147483648))
    var v26 uint32 = v25
    switch (^(v14)) {
    case 0:
        v27 := v3[v26]
        _ = v27
    case -271:
        v14 -= v14
    }
    if false {
        v4 = true
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
