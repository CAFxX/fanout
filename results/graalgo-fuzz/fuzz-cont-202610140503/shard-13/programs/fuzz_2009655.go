package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 int }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 int; F1 uint32; F2 bool }
func genFn0() int8 {
    v7 := GenS1{F0: "a", F1: (true && false), F2: 4294967295}
    if true {
        var v8 float32 = 0.0
        if false {
            v8 += float32((1.5 / float64(float32(-2.25))))
        }
        _ = v8
    } else {
        var v9 float32 = float32(1e10)
        _ = v9
    }
    var v10 int32 = 2147483647
    _ = v10
    _ = v7
    return 56
}

func main() {
    var v1 int32 = 1
    var v2 uint32 = 302
    v3 := []uint32{v2, 2}
    v4 := make(map[uint32]string)
    var v5 int8 = 127
    var v6 uint32 = 42
    v11 := GenS1{F0: (("\n\t" + "\n\t") + (("" + ("a" + "hello")) + (("" + ("foo bar" + ("" + "\n\t"))) + "hello"))), F1: false, F2: 2147483647}
    for fi12 := 0; fi12 < 6; fi12++ {
        fmt.Println(v2, v1)
        for fi13 := 0; fi13 < 9; fi13++ {
            v6 += 65535
            _ = fi13
        }
        _ = fi12
    }
    fmt.Println(v2, v1)
    if false {
        if false {
            var v14 uint64 = 655
            _ = v14
        }
        v2 = 1
    }
    v15 := GenS1{F0: "foo bar", F1: false, F2: 1}
    var v16 float64 = (1.5 - ((-2.25 + 1e10) - ((-2.25 - 0.1) / 3.14159)))
    var v17 float32 = 1e10
    for fi18 := 0; fi18 < 2; fi18++ {
        switch (fi18 | fi18) {
        case 127:
            v3 = append(v3, 4294967295)
        default:
            if (!(-826 <= 42)) {
                _ = len(v4)
                for fi19 := 0; fi19 < 9; fi19++ {
                    v20 := GenS2{F0: v1, F1: 159}
                    _ = fi19
                    _ = v20
                }
                fmt.Println(fi18)
            }
        }
        _ = fi18
    }
    v21 := GenS3{F0: 4294967295, F1: (v6 >> ((0) & 63)), F2: true}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
