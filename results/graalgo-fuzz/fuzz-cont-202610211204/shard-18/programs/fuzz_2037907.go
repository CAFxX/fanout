package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() uint32 {
    var v4 int16 = 1
    var v5 float64 = 3.14159
    v6 := []int{32767, 530, -1, 128}
    _ = v4
    _ = v5
    _ = v6
    return 891
}
func genFn1() string {
    if true {
        var v7 uint8 = 30
        v7 += 18
        _ = v7
    }
    v8 := []int{238, -2147483648, 1, 2147483647}
    _ = v8
    return "a"
}

func main() {
    var v1 string = "hello"
    var v2 float64 = 3.14159
    v3 := make(map[uint32]bool)
    var v9 int8 = 0
    for fi10 := 0; fi10 < 10; fi10++ {
        v11 := GenS1{F0: -1}
        _ = fi10
        _ = v11
    }
    var v12 float64 = 3.14159
    v13 := []int{}
    v2 = v2
    if (false && (0 >= 9223372036854775807)) {
        fmt.Println(v9)
    }
    var v14 int64 = -1
    for fi15 := 0; fi15 < 10; fi15++ {
        for fi16 := 0; fi16 < 8; fi16++ {
            v13 = append(v13, -1)
            _ = fi16
        }
        for fi17 := 0; fi17 < 17; fi17++ {
            v18 := v3[2]
            _ = fi17
            _ = v18
        }
        _ = fi15
    }
    v14 = v14
    v1 = "\n\t"
    v2 += (3.14159 + (v2 + (float64(1.5) + (v2 * ((1e10 * float64(((1e10 - 1.5) + 0.1))) / (1e10 - (float64(0.1) * (0.0 / 0.1))))))))
    v3[2] = ((true && true) || false)
    if true {
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        if true {
            fmt.Println(v14, v12)
            v1 = "a"
        } else {
            fmt.Println(v1, v14)
            delete(v3, 4294967295)
            v2 = (0.0 / v12)
        }
        var v19 string = v1
        _ = v19
    }
    var v20 uint32 = 834
    if (!false) {
        var v21 float32 = 3.14159
        _ = v21
    }
    var v22 int8 = int8(v14)
    var v23 uint = 93
    v13 = append(v13, 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v9
}
