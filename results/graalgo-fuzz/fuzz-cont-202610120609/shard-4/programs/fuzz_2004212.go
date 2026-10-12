package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 int; F2 uint }
func genFn0(p0 int, p1 int32, p2 string) string {
    fmt.Println(p1)
    v6 := GenS1{F0: 296, F1: "foo bar", F2: 254}
    v7 := GenS1{F0: 926, F1: p2, F2: 36}
    var v8 int16 = -32768
    var v9 int8 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return "a"
}
func genFn1(p0 uint8, p1 int8) uint8 {
    for fi10 := 0; fi10 < 19; fi10++ {
        if false {
            v11 := []int{(fi10 - (fi10 * (fi10 / ((fi10) | 1)))), (fi10 / ((fi10) | 1))}
            fmt.Println(fi10, p1, p1)
            _ = v11
        } else {
            p1 -= p1
        }
        if (true || (!true)) {
            p1 -= (p1 + (p1 / (((p1 / ((1) | 1))) | 1)))
            fmt.Println(p1, fi10)
            var v12 int16 = 304
            _ = v12
        } else {
            var v13 bool = ((!(false || true)) && (float64(3.14159) > 0.0))
            _ = v13
        }
        _ = fi10
    }
    var v14 int = 32767
    _ = p0
    _ = p1
    _ = v14
    return p0
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[uint32]int)
    v3 := make(map[string]string)
    var v4 float32 = 3.14159
    var v5 string = "\n\t"
    v15 := GenS1{F0: 1, F1: v5, F2: 0}
    for fi16 := 0; fi16 < 12; fi16++ {
        v1 = 919
        fmt.Println(v1, v1)
        _ = fi16
    }
    if true {
        if true {
            for fi17 := 0; fi17 < 13; fi17++ {
                v18 := GenS2{F0: (^(v1))}
                _ = fi17
                _ = v18
            }
        } else {
            v3[(v5 + ("" + ((("hello" + "") + "\n\t") + v5)))] = "foo bar"
        }
    }
    v5 = ("a" + v5)
    v19 := GenS1{F0: 922, F1: "foo bar", F2: 1}
    var v20 int8 = 127
    v21 := []int{}
    v22 := GenS3{F0: -1, F1: -1, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
