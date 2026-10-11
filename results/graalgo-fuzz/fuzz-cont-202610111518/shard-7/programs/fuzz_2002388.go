package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int8; F1 bool; F2 string }
type GenS3 struct { F0 int16; F1 int8 }
func genFn0() uint8 {
    v7 := []int{4294967295, -1}
    var v8 float64 = (1.5 * ((float64(((3.14159 + 1e10) / ((-2.25 / 1.5) / 1.0))) / 3.14159) * 3.14159))
    v9 := GenS1{F0: 607, F1: 327}
    v8 = v8
    if true {
        if (42 == 2) {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            var v10 uint16 = 1
            _ = v10
        } else {
            fmt.Println(v8)
            fmt.Println(v8)
        }
    } else {
        fmt.Println(v8)
        v11 := GenS1{F0: 0, F1: 1}
        v8 = v8
        _ = v11
    }
    _ = v7
    _ = v8
    _ = v9
    return 255
}

func main() {
    v1 := []string{("foo bar" + "a"), ("hello" + ("\n\t" + "a"))}
    var v2 uint = 537
    var v3 int32 = 449
    var v4 uint16 = 1
    v5 := make(map[int]bool)
    var v6 uint16 = (^(v4))
    v12 := GenS2{F0: 1, F1: false, F2: ("a" + "hello")}
    v13 := GenS3{F0: 1, F1: 127}
    var v14 int32 = (-(v3))
    v4 += (v4 * v4)
    for fi15 := 0; fi15 < 5; fi15++ {
        v14 -= v3
        delete(v5, fi15)
        _ = fi15
    }
    fmt.Println(v2, v4)
    v2 = (v2 << (((v2 / (((^(v2))) | 1))) & 63))
    v14 = v3
    v3 -= 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
