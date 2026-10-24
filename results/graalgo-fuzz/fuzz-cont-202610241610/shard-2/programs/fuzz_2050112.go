package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint }
type GenS2 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int32) bool {
    var v4 int64 = 0
    for fi5 := 0; fi5 < 11; fi5++ {
        p0 = p0
        p0 = (p0 % (((p0 ^ 1)) | 1))
        var v6 float64 = 0.1
        _ = fi5
        _ = v6
    }
    if (!true) {
        v7 := GenS1{F0: "", F1: true, F2: 69}
        fmt.Println(v4)
        _ = v7
    } else {
        v4 += -479
    }
    _ = p0
    _ = v4
    return (v4 >= v4)
}
func genFn1(p0 bool, p1 uint32, p2 int) uint {
    fmt.Println(p0, p1, p2)
    p0 = (255 > 0)
    _ = p0
    _ = p1
    _ = p2
    return 584
}

func main() {
    var v1 int32 = -1
    v2 := []string{}
    v3 := make(map[string]bool)
    v8 := GenS1{F0: ("foo bar" + "foo bar"), F1: true, F2: 65535}
    v9 := GenS2{F0: -128, F1: 487, F2: 1}
    var v10 int = -1
    var v11 float32 = 0.0
    v12 := GenS3{F0: 32767, F1: 0}
    v13 := GenS4{F0: 204}
    var v14 string = "hello"
    v1 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
