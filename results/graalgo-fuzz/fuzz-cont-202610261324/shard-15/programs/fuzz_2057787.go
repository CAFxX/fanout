package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int8 }
func genFn0() uint8 {
    switch (("\n\t" + "a") + (("foo bar" + "hello") + ("" + ""))) {
    case "foo bar":
        v5 := []int{}
        _ = v5
    case "a":
        var v6 int64 = 1
        _ = v6
    case "":
        var v7 uint = 1
        _ = v7
    }
    var v8 uint = 171
    _ = v8
    return 255
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]uint32)
    var v3 uint8 = 42
    var v4 uint32 = 2147483647
    v9 := []int{}
    v4 = (v4 - (v4 + v4))
    v10 := GenS1{F0: "a", F1: 2, F2: 12}
    v9 = append(v9, 0)
    var v11 float32 = (-2.25 / 1.5)
    _ = len(v9)
    if ((v4 - 731) >= (v4 - (v4 >> ((0) & 63)))) {
        if true {
            v11 = v11
            v11 += 0.0
            var v12 int8 = 96
            _ = v12
        } else {
            v3 += 2
            var v13 float64 = -2.25
            var v14 string = ("\n\t" + "hello")
            _ = v13
            _ = v14
        }
    }
    switch "" {
    case "foo bar":
        v4 += v4
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        var v16 uint16 = 575
        var v17 int32 = -1
        _ = fi15
        _ = v16
        _ = v17
    }
    fmt.Println(v4, v3)
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    fmt.Println(v3, v11)
    if len(v9) > 0 {
        v9[0] = -2147483648
    }
    for fi18 := 0; fi18 < 19; fi18++ {
        delete(v2, "hello")
        var v19 int8 = 118
        v2["\n\t"] = (^(v4))
        _ = fi18
        _ = v19
    }
    fmt.Println(v3)
    var v20 float64 = -2.25
    v20 -= (float64(((v11 / (v11 * v11)) + v11)) * float64(float32((v20 * v20))))
    v20 -= ((v20 / 1.0) - (v20 - (v20 * ((v20 * (v20 - 0.1)) - v20))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v9
}
