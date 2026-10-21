package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 string }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 int }

func main() {
    v1 := make(map[int]uint32)
    var v2 string = "a"
    var v3 uint16 = 65535
    var v4 int64 = 9223372036854775807
    var v5 float32 = 3.14159
    var v6 uint16 = v3
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 int16 = -878
        v6 -= v3
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 16; fi9++ {
        v10 := GenS1{F0: v2, F1: fi9, F2: ("hello" + v2)}
        v6 = (^(v3))
        _ = fi9
        _ = v10
    }
    if false {
        if true {
            var v11 uint32 = 2147483647
            _ = v11
        }
    }
    v4 = (v4 % (((v4 / (((v4 + v4)) | 1))) | 1))
    v6 += 2
    v12 := GenS2{F0: v2, F1: (v4 * -1)}
    v5 += (v5 * (-2.25 / v5))
    v6 -= 2
    v13 := GenS1{F0: v2, F1: -657, F2: (v2 + (("\n\t" + v2) + ""))}
    for fi14 := 0; fi14 < 2; fi14++ {
        var v15 float32 = v5
        _ = fi14
        _ = v15
    }
    var v16 float32 = v5
    v17 := GenS3{F0: 9223372036854775807}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
