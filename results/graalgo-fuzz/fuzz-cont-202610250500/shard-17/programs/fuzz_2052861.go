package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int64; F1 int16 }
type GenS4 struct { F0 int32; F1 int64 }
func genFn0() int {
    var v5 int8 = -1
    if ("" > ("a" + (("foo bar" + ("a" + (("a" + "a") + ("" + "a")))) + "a"))) {
        if (!(!true)) {
            v6 := GenS1{F0: 677, F1: ((0 > 65535) && false)}
            _ = v6
        }
        fmt.Println(v5)
    }
    _ = v5
    return -1
}

func main() {
    var v1 uint64 = 42
    var v2 uint8 = 208
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]uint32)
    v7 := []int{genFn0(), 618, 32768, 9223372036854775807}
    if ((0.1 - (3.14159 / 1.5)) == (1e10 / 1e10)) {
        v1 -= 0
        if (-2.25 != -2.25) {
            var v8 uint = 32
            _ = v8
        } else {
            if true {
                v9 := GenS1{F0: 32767, F1: (-32768 >= 0)}
                _ = v9
            }
            delete(v4, 83)
            v2 = (^(v2))
        }
        var v10 int8 = 42
        _ = v10
    } else {
        var v11 uint64 = v1
        fmt.Println(v2, v2, v1)
        v12 := GenS2{F0: 4294967295, F1: 65535}
        _ = v11
        _ = v12
    }
    fmt.Println(v2, v2)
    var v13 float32 = float32((float64(1.5) * ((((float64(-2.25) * (float64(1.5) + 3.14159)) * float64((3.14159 * ((-2.25 * (-2.25 + 0.0)) * 3.14159)))) / -2.25) / 1e10)))
    var v14 uint16 = 42
    v1 += 0
    if (!(1 < 42)) {
        v2 = (^(v2))
        fmt.Println(v2, v13)
    } else {
        var v15 int = 9223372036854775807
        _ = v15
    }
    var v16 float32 = v13
    v17 := GenS3{F0: -1, F1: -336}
    _ = len(v3)
    fmt.Println(v14, v13, v16)
    var v18 float64 = 1.5
    v19 := GenS4{F0: 537, F1: 0}
    var v20 uint = 255
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v7
}
