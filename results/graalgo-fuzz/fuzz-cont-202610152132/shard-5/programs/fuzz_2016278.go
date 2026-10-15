package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 bool }
func genFn0() int64 {
    var v6 float64 = 3.14159
    fmt.Println(v6)
    _ = v6
    return 42
}

func main() {
    var v1 int8 = -1
    var v2 bool = true
    var v3 int64 = 0
    v4 := []uint32{4294967295, 706}
    v5 := make(map[int]int)
    for fi7 := 0; fi7 < 10; fi7++ {
        if v2 {
            var v8 int32 = 0
            _ = v8
        }
        v9 := GenS1{F0: 56, F1: 65535}
        v1 = v1
        _ = fi7
        _ = v9
    }
    v10 := GenS1{F0: 65535, F1: 32767}
    v3 -= v3
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = len(v5)
    var v11 bool = v2
    fmt.Println(v11)
    var v12 bool = (-2147483648 == 617)
    v13 := v5[42]
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 uint8 = 115
        fmt.Println(fi14, v3)
        _ = fi14
        _ = v15
    }
    if true {
        for fi16 := 0; fi16 < 17; fi16++ {
            v17 := GenS1{F0: 2147483647, F1: 2}
            fmt.Println(fi16, v11, v2)
            v11 = false
            _ = fi16
            _ = v17
        }
    }
    v18 := GenS2{F0: 2147483647, F1: 4294967295}
    if (42 >= 363) {
        switch "" {
        case "a":
            v19 := GenS1{F0: 631, F1: 255}
            _ = v19
        }
    } else {
        _ = len(v4)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    var v20 float32 = 0.1
    v11 = (4294967295 >= 100)
    fmt.Println(v13, v11, v13)
    v21 := GenS3{F0: (42 < 0)}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
