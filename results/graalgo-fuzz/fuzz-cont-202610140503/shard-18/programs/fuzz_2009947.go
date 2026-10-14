package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string; F1 uint64 }
func genFn0(p0 uint8, p1 uint32, p2 uint) int {
    fmt.Println(p0)
    if (((!(548 <= -2147483648)) && (uint16(p2) > 2)) && true) {
        v6 := GenS1{F0: -48, F1: -32768, F2: p0}
        _ = v6
    } else {
        if false {
            var v7 int64 = 42
            v8 := GenS2{F0: p2}
            _ = v7
            _ = v8
        }
    }
    v9 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return -932
}

func main() {
    var v1 int8 = -1
    v2 := []uint32{2147483648, 224}
    v3 := make(map[int]string)
    var v4 bool = (true && false)
    v5 := make(map[int]uint32)
    if (!true) {
        v1 = v1
        v4 = v4
    } else {
        for fi10 := 0; fi10 < 16; fi10++ {
            v1 -= 29
            v1 += v1
            _ = fi10
        }
    }
    var v11 float32 = 1e10
    v5[110] = 4294967295
    var v12 uint32 = 2
    var v13 uint8 = 0
    fmt.Println(v11, v1)
    v14 := GenS3{F0: 42}
    v3[944] = (("hello" + "") + "foo bar")
    if len(v2) > 0 {
        v2[0] = 203
    }
    if v4 {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi15 := 0; fi15 < 2; fi15++ {
            v16 := GenS4{F0: "", F1: 4294967295}
            _ = fi15
            _ = v16
        }
        v12 -= v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
