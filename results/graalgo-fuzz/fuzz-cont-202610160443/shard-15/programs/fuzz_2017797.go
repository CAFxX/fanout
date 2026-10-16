package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 uint; F2 int32 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 uint32; F1 string }

func main() {
    var v1 int16 = 93
    var v2 uint32 = 42
    v3 := make(map[uint32]bool)
    var v4 int64 = -1
    var v5 uint64 = 53
    var v6 uint64 = (v5 ^ 18446744073709551615)
    v7 := GenS1{F0: -440}
    for fi8 := 0; fi8 < 3; fi8++ {
        for fi9 := 0; fi9 < 9; fi9++ {
            fmt.Println(v4)
            if (true && true) {
                var v10 uint16 = 32767
                delete(v3, (^(v2)))
                _ = v10
            }
            _ = fi9
        }
        fmt.Println(v5)
        v11 := []int{(^(fi8))}
        _ = fi8
        _ = v11
    }
    fmt.Println(v2)
    switch "" {
    case "foo bar":
        v2 += v2
    case "hello":
        var v12 string = "\n\t"
        _ = v12
    case "":
        fmt.Println(v5)
    }
    if false {
        var v13 int = 128
        var v14 uint16 = 42
        _ = v13
        _ = v14
    } else {
        if (128 >= 78) {
            var v15 uint8 = 42
            _ = v15
        }
        v16 := GenS2{F0: (v1 << ((2) & 63)), F1: 2, F2: -174}
        _ = v16
    }
    var v17 bool = ((0.0 * (float64(float32(0.0)) / 1.5)) != float64(1e10))
    var v18 uint8 = 158
    var v19 float32 = (3.14159 / (3.14159 - ((0.1 / 0.1) * ((0.0 + ((float32((3.14159 - 0.0)) - 0.1) + 1.5)) + 1e10))))
    var v20 uint16 = 0
    v17 = v17
    v21 := GenS3{F0: 0, F1: (^(v20))}
    var v22 uint = 255
    v23 := GenS4{F0: (v2 << (((v22 << ((v22) & 63))) & 63)), F1: ("\n\t" + ((("hello" + "\n\t") + (("a" + ("\n\t" + "hello")) + "foo bar")) + ("" + ("" + "a"))))}
    v24 := GenS1{F0: v1}
    var v25 uint32 = v2
    var v26 uint = v22
    var v27 uint32 = 1
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
