package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 uint32; F2 int8 }

func main() {
    var v1 float64 = (1.5 * ((3.14159 + -2.25) * 0.1))
    var v2 uint16 = 1
    v3 := make(map[string]bool)
    v4 := make(map[string]bool)
    v5 := make(map[string]uint32)
    var v6 uint8 = 42
    fmt.Println(v1, v1)
    v7 := []int{50, 4294967295, 785}
    for fi8 := 0; fi8 < 6; fi8++ {
        v1 = (v1 / ((v1 * v1) - 1e10))
        _ = fi8
    }
    v6 -= v6
    v9 := GenS1{F0: 0}
    var v10 uint = 2
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v1 = 0.0
    v1 -= v1
    for fi11 := 0; fi11 < 3; fi11++ {
        v12 := GenS2{F0: v6}
        _ = fi11
        _ = v12
    }
    var v13 float32 = float32(v1)
    var v14 int8 = 100
    v4[((("" + "") + "") + "foo bar")] = false
    var v15 int32 = -2147483648
    if (881 != 2147483647) {
        if true {
            var v16 int64 = 1
            _ = v16
        }
        for fi17 := 0; fi17 < 10; fi17++ {
            delete(v5, "hello")
            delete(v5, "\n\t")
            _ = fi17
        }
    }
    if false {
        if len(v7) > 0 {
            v7[0] = 127
        }
        v13 = float32(v1)
    } else {
        for fi18 := 0; fi18 < 12; fi18++ {
            v19 := GenS3{F0: (v2 % (((v2 - 42)) | 1)), F1: 1, F2: v14}
            v20 := GenS1{F0: v14}
            var v21 uint16 = 1
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
        for fi22 := 0; fi22 < 7; fi22++ {
            var v23 uint = (v10 / ((42) | 1))
            _ = fi22
            _ = v23
        }
    }
    var v24 uint32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
