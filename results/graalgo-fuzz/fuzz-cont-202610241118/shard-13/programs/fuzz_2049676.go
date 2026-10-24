package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint32 }
func genFn0(p0 string) uint {
    v5 := GenS1{F0: 0, F1: 127, F2: 4294967295}
    p0 = "hello"
    p0 = ""
    if false {
        p0 = ((p0 + p0) + "foo bar")
        v6 := GenS1{F0: 4294967295, F1: 127, F2: 65535}
        _ = v6
    } else {
        v7 := []int{-1, 92, 32768}
        p0 = p0
        p0 = p0
        _ = v7
    }
    p0 = p0
    _ = p0
    _ = v5
    return 2
}
func genFn1(p0 uint64) uint16 {
    p0 += p0
    p0 -= p0
    fmt.Println(p0)
    _ = p0
    return 42
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 864
    v4 := []string{("hello" + "a"), ("\n\t" + ("hello" + "a"))}
    v8 := GenS1{F0: 1, F1: 1, F2: 217}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v9 := v2[262]
    delete(v2, v9)
    var v10 uint16 = 32767
    v10 = 255
    v3 += (v3 << ((626) & 63))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v11 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
