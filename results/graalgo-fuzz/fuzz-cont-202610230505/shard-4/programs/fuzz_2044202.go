package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 string = "foo bar"
    v2 := make(map[string]string)
    v3 := []int{0, -564, 32768, 2147483647}
    var v4 int32 = -1
    v5 := make(map[string]int)
    for fi6 := 0; fi6 < 16; fi6++ {
        fmt.Println(fi6, fi6, v1)
        var v7 uint32 = 0
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 13; fi8++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v9 string = v1
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: 0}
    for fi11 := 0; fi11 < 5; fi11++ {
        v4 -= -1
        v1 = ("" + "")
        v12 := GenS1{F0: 1000}
        _ = fi11
        _ = v12
    }
    if ((!false) || (-387 >= 4294967295)) {
        var v13 int64 = 9223372036854775807
        if false {
            var v14 uint16 = 592
            _ = v14
        }
        for fi15 := 0; fi15 < 9; fi15++ {
            var v16 uint8 = 47
            v2[(((((v1 + v1) + (v1 + v1)) + v1) + (v1 + ((v1 + "foo bar") + ""))) + ((("foo bar" + "") + (v1 + v1)) + (v1 + (v1 + v1))))] = "hello"
            _ = fi15
            _ = v16
        }
        _ = v13
    } else {
        var v17 int8 = -1
        _ = v17
    }
    v18 := GenS1{F0: 32767}
    v4 -= v4
    if false {
        var v19 int8 = 100
        _ = v19
    }
    var v20 float32 = float32(0.1)
    v4 -= 2147483647
    switch 473 {
    case 42:
        var v21 int16 = -951
        _ = v21
    case 4294967295:
        _ = len(v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
