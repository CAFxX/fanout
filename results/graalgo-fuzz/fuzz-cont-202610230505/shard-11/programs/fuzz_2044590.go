package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 uint8; F2 int32 }
func genFn0() uint32 {
    var v6 uint = 65535
    if (true || ("\n\t" != "hello")) {
        if true {
            v7 := GenS1{F0: 0, F1: -530}
            v6 += v6
            var v8 uint8 = 128
            _ = v7
            _ = v8
        }
    }
    v9 := []int{720, -901}
    var v10 uint32 = 0
    fmt.Println(v6, v6)
    _ = v10
    _ = v6
    _ = v9
    return 4294967295
}

func main() {
    var v1 uint = 65535
    var v2 float64 = 3.14159
    v3 := []int{0, -942, 270, -2147483648}
    v4 := []string{"foo bar", "foo bar"}
    var v5 string = "foo bar"
    v11 := GenS1{F0: 1, F1: -9223372036854775808}
    if (!true) {
        fmt.Println(v2)
        v2 -= float64(0.1)
        var v12 int8 = -97
        _ = v12
    } else {
        v2 += 0.1
        for fi13 := 0; fi13 < 3; fi13++ {
            if (32767 <= -127) {
                v4 = append(v4, v5)
            }
            _ = len(v3)
            fmt.Println(v2, v1)
            _ = fi13
        }
    }
    v14 := GenS2{F0: genFn0()}
    if false {
        var v15 int16 = 1
        fmt.Println(v5)
        _ = v15
    } else {
        v2 += float64(((((float32((v2 / float64(3.14159))) / 1.0) / 1e10) - 3.14159) * 1e10))
    }
    v5 = "\n\t"
    v4 = append(v4, (v5 + "\n\t"))
    if (v1 != v1) {
        var v16 uint8 = 255
        if false {
            var v17 float64 = (v2 + v2)
            _ = len(v3)
            _ = v17
        }
        v18 := GenS3{F0: true, F1: v16, F2: 0}
        _ = v16
        _ = v18
    }
    var v19 int64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
