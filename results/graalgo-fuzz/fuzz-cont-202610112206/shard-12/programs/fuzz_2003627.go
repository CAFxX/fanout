package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 string }
type GenS2 struct { F0 int64; F1 int64; F2 uint8 }
type GenS3 struct { F0 bool; F1 int64; F2 string }

func main() {
    v1 := []int{-9223372036854775808, 32767, -245, -906}
    v2 := make(map[string]bool)
    var v3 uint32 = 211
    var v4 uint = 699
    var v5 uint8 = 255
    switch 32768 {
    case 127:
        if ((0.1 + (0.1 + (((0.0 * float32(1.5)) * (-2.25 + 1.5)) / 3.14159))) > 0.1) {
            for fi6 := 0; fi6 < 20; fi6++ {
                v7 := GenS1{F0: -531, F1: v3, F2: "foo bar"}
                v4 += v4
                var v8 uint = v4
                _ = fi6
                _ = v7
                _ = v8
            }
            var v9 int64 = 325
            if (true || false) {
                var v10 float64 = 0.1
                _ = v10
            } else {
                if (true || (uint16(v3) != 322)) {
                    v11 := GenS2{F0: -271, F1: v9, F2: 0}
                    fmt.Println(v4, v9, v3)
                    var v12 uint32 = (v3 | v3)
                    _ = v11
                    _ = v12
                }
            }
            _ = v9
        }
    case 4294967295:
        var v13 int8 = 0
        _ = v13
    }
    var v14 float32 = 0.0
    if false {
        v15 := GenS3{F0: false, F1: -63, F2: (("" + "foo bar") + "")}
        _ = v15
    }
    v16 := GenS3{F0: ((v3 * v3) > (^(v3))), F1: 956, F2: "a"}
    if true {
        for fi17 := 0; fi17 < 14; fi17++ {
            var v18 uint8 = (v5 - 63)
            _ = fi17
            _ = v18
        }
    } else {
        fmt.Println(v5, v14, v4)
    }
    var v19 int = 4294967295
    delete(v2, "hello")
    var v20 uint32 = v3
    fmt.Println(v3, v3)
    v21 := GenS3{F0: ((v20 > v20) && true), F1: 0, F2: "hello"}
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
