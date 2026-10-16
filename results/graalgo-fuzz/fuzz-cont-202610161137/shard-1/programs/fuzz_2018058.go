package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]string)
    var v3 float32 = 0.1
    v4 := []int{}
    if false {
        if len(v4) > 0 {
            v4[0] = 127
        }
    } else {
        var v5 float64 = ((-2.25 / float64(v3)) - float64((1e10 / v3)))
        var v6 uint8 = 255
        var v7 int64 = 42
        _ = v5
        _ = v6
        _ = v7
    }
    var v8 uint8 = 42
    var v9 uint32 = 42
    if len(v4) > 0 {
        v4[0] = 253
    }
    fmt.Println(v9, v3)
    fmt.Println(v8, v9, v3)
    var v10 uint16 = 956
    var v11 uint = 2
    if (3.14159 == 3.14159) {
        fmt.Println(v9, v10, v11)
        v11 -= (v11 * v11)
        if (true || (false && true)) {
            fmt.Println(v9, v8)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if ((-395 < 716) && true) {
                _ = len(v4)
            }
        }
    }
    switch 553 {
    case 1:
        v12 := GenS1{F0: "a"}
        _ = v12
    case -9223372036854775808:
        for fi13 := 0; fi13 < 6; fi13++ {
            var v14 string = ("" + "hello")
            _ = fi13
            _ = v14
        }
    }
    fmt.Println(v11, v3, v10)
    delete(v1, "a")
    v11 -= v11
    v15 := GenS1{F0: "\n\t"}
    if true {
        v16 := GenS1{F0: "hello"}
        _ = v16
    } else {
        for fi17 := 0; fi17 < 20; fi17++ {
            v8 -= v8
            _ = fi17
        }
    }
    v18 := GenS2{F0: 327}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
