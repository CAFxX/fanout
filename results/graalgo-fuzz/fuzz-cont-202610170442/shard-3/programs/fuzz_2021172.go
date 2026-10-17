package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 string; F2 int16 }
type GenS3 struct { F0 string }
func genFn0(p0 int16, p1 int16) uint {
    if (!true) {
        var v7 int = 952
        v8 := GenS1{F0: (p0 & (p0 - -825))}
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: p0}
    fmt.Println(p0, p0)
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 bool = false
        var v12 int32 = 42
        v13 := []int{32768}
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    v14 := GenS2{F0: 944, F1: "hello", F2: (p1 % (((^(p0))) | 1))}
    _ = p0
    _ = p1
    _ = v14
    _ = v9
    return 807
}
func genFn1(p0 uint, p1 uint64) uint {
    for fi15 := 0; fi15 < 4; fi15++ {
        v16 := GenS2{F0: 255, F1: (("" + "a") + "\n\t"), F2: -1}
        _ = fi15
        _ = v16
    }
    v17 := GenS2{F0: 0, F1: "a", F2: -1}
    var v18 uint = p0
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    return (p0 ^ v18)
}

func main() {
    v1 := make(map[string]uint32)
    var v2 string = ("hello" + "")
    v3 := []uint32{2, 1, 159, 960}
    var v4 float64 = (1.5 + -2.25)
    v5 := []string{v2, (("hello" + v2) + "hello"), (v2 + v2), v2}
    v6 := make(map[int]uint32)
    if (!(!false)) {
        _ = len(v6)
        fmt.Println(v4, v2)
    } else {
        for fi19 := 0; fi19 < 10; fi19++ {
            var v20 uint8 = 42
            var v21 int32 = 0
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    v4 = v4
    v22 := GenS3{F0: "a"}
    v23 := GenS3{F0: ""}
    var v24 string = "foo bar"
    var v25 int16 = 133
    if true {
        v2 = (v24 + (("a" + "a") + (v2 + "a")))
    }
    v25 += v25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
