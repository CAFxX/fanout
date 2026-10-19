package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint16; F1 uint64 }
func genFn0(p0 uint16, p1 uint) uint {
    for fi4 := 0; fi4 < 3; fi4++ {
        v5 := GenS1{F0: (-1 < 1), F1: 64}
        for fi6 := 0; fi6 < 15; fi6++ {
            p0 = 1
            var v7 string = (("foo bar" + "hello") + ("a" + "foo bar"))
            _ = fi6
            _ = v7
        }
        var v8 uint32 = 65535
        _ = fi4
        _ = v5
        _ = v8
    }
    if false {
        var v9 uint8 = 127
        v10 := []int{-9223372036854775808, 129, -2147483648}
        _ = v10
        _ = v9
    } else {
        var v11 string = ("" + "hello")
        _ = v11
    }
    v12 := GenS2{F0: 27, F1: 42, F2: 249}
    v13 := []int{127, 1}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return p1
}

func main() {
    var v1 string = "foo bar"
    var v2 string = "hello"
    v3 := make(map[int]uint32)
    v1 = "hello"
    for fi14 := 0; fi14 < 18; fi14++ {
        v15 := GenS3{F0: 32767, F1: 42}
        _ = fi14
        _ = v15
    }
    var v16 float32 = 0.1
    var v17 float32 = v16
    for fi18 := 0; fi18 < 10; fi18++ {
        v2 = "\n\t"
        var v19 uint8 = 191
        _ = fi18
        _ = v19
    }
    fmt.Println(v1, v2)
    var v20 string = v2
    v1 = v1
    var v21 uint = 18446744073709551615
    var v22 bool = ((true || true) || true)
    v23 := GenS3{F0: 15, F1: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
