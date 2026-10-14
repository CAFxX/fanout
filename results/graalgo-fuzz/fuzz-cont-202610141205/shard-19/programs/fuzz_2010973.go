package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint8; F1 int }
func genFn0(p0 string, p1 uint64, p2 uint) bool {
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 bool = (false && true)
        _ = fi6
        _ = v7
    }
    if (-722 != 9223372036854775807) {
        var v8 uint8 = 234
        _ = v8
    }
    p1 = (p1 + 4294967295)
    p0 = (("hello" + p0) + p0)
    for fi9 := 0; fi9 < 20; fi9++ {
        if true {
            p0 = p0
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    return (!((true || true) && false))
}

func main() {
    var v1 string = "foo bar"
    var v2 int32 = -921
    v3 := make(map[uint32]string)
    var v4 uint8 = 1
    var v5 int64 = -1
    if (0.1 <= 0.0) {
        if (3.14159 >= float32(((-2.25 / -2.25) / 3.14159))) {
            var v10 bool = (!true)
            for fi11 := 0; fi11 < 6; fi11++ {
                fmt.Println(v10, v5)
                _ = fi11
            }
            v12 := GenS1{F0: v4, F1: 1}
            _ = v10
            _ = v12
        }
    } else {
        if true {
            v13 := []int{128, 127}
            _ = v13
        } else {
            var v14 int = 430
            _ = v14
        }
        v15 := GenS2{F0: v4, F1: 128}
        v4 = (v4 % ((v4) | 1))
        _ = v15
    }
    var v16 int32 = (v2 / (((v2 - v2)) | 1))
    var v17 uint16 = 0
    v16 = -1
    var v18 uint64 = 18446744073709551615
    fmt.Println(v4, v4, v5)
    var v19 uint64 = v18
    for fi20 := 0; fi20 < 2; fi20++ {
        v21 := []int{(-(fi20)), 0}
        _ = fi20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
