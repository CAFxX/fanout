package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 int16; F1 uint; F2 int64 }
type GenS4 struct { F0 uint32; F1 bool; F2 uint8 }
func genFn0() int8 {
    var v7 int16 = 32767
    fmt.Println(v7)
    var v8 uint8 = 0
    _ = v7
    _ = v8
    return 22
}
func genFn1() int64 {
    v9 := GenS1{F0: -1, F1: 4294967295, F2: 255}
    if (-2147483648 != 287) {
        var v10 uint16 = 65535
        var v11 int16 = 696
        _ = v10
        _ = v11
    }
    var v12 int8 = 42
    v13 := []int{-1, 432, 127}
    _ = v12
    _ = v13
    _ = v9
    return 285
}

func main() {
    var v1 int16 = 32767
    v2 := []string{}
    v3 := []string{""}
    v4 := make(map[int]string)
    v5 := make(map[int]uint32)
    v6 := []string{("" + ("foo bar" + "")), ("" + "\n\t")}
    _ = len(v2)
    v1 = v1
    v14 := GenS2{F0: genFn1(), F1: -32768}
    v15 := GenS3{F0: v1, F1: 673, F2: -1}
    if (true || true) {
        v1 = 1
        v16 := GenS4{F0: 42, F1: true, F2: 255}
        v1 = v1
        _ = v16
    }
    v1 += v1
    for fi17 := 0; fi17 < 11; fi17++ {
        v18 := GenS4{F0: 697, F1: true, F2: 56}
        var v19 int8 = 42
        _ = fi17
        _ = v18
        _ = v19
    }
    var v20 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
