package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int; F1 int32 }
func genFn0() int8 {
    var v6 int8 = 127
    v7 := GenS1{F0: 2, F1: 767}
    v6 = (v6 % ((-128) | 1))
    var v8 uint64 = 18446744073709551615
    _ = v6
    _ = v7
    _ = v8
    return -121
}
func genFn1(p0 uint64, p1 int8, p2 uint32) uint16 {
    fmt.Println(p0)
    v9 := []int{32768}
    switch 127 {
    case -714:
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 int8 = p1
            _ = fi10
            _ = v11
        }
    case 1:
        if (!(false || (-1 < -258))) {
            var v12 uint16 = 32767
            _ = v12
        } else {
            p1 = 1
            var v13 uint32 = p2
            _ = v13
        }
    }
    fmt.Println(p1, p1, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return uint16(p2)
}

func main() {
    v1 := []string{(("hello" + ("a" + (("" + "hello") + ""))) + "a"), ("" + ("foo bar" + ("hello" + (("a" + (("\n\t" + "foo bar") + "a")) + "a")))), "hello", "foo bar"}
    v2 := []uint32{4294967295}
    v3 := []uint32{42}
    var v4 float64 = 0.0
    v5 := make(map[int]string)
    v14 := GenS2{F0: 127}
    for fi15 := 0; fi15 < 18; fi15++ {
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 float64 = float64(float32(0.1))
            v18 := GenS2{F0: 104}
            _ = fi16
            _ = v17
            _ = v18
        }
        if false {
            v4 = v4
        }
        v19 := GenS3{F0: 1000}
        _ = fi15
        _ = v19
    }
    var v20 string = "a"
    if false {
        var v21 bool = (!true)
        v22 := GenS4{F0: 127, F1: 42}
        _ = v21
        _ = v22
    } else {
        if len(v1) > 0 {
            v1[0] = ""
        }
        var v23 uint32 = 2147483647
        v4 -= (((float64(1.5) + 0.0) * v4) / (v4 - v4))
        _ = v23
    }
    v20 = v20
    for fi24 := 0; fi24 < 15; fi24++ {
        v25 := GenS3{F0: -1}
        _ = fi24
        _ = v25
    }
    var v26 bool = (!(!true))
    v20 = v20
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
