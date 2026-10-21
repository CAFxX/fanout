package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 string; F2 int64 }
type GenS3 struct { F0 uint8; F1 int8 }
func genFn0(p0 bool, p1 bool, p2 uint32) uint64 {
    v4 := []int{671, 127}
    p0 = ((53 > 77) && false)
    v5 := GenS1{F0: 1}
    _ = len(v4)
    fmt.Println(p0, p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 uint) uint64 {
    v6 := GenS2{F0: -1, F1: "\n\t", F2: -417}
    v7 := GenS3{F0: 142, F1: 100}
    v8 := []int{42, -1, 128}
    v9 := GenS2{F0: 42, F1: "", F2: -1}
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := make(map[string]string)
    var v2 int = -814
    var v3 uint64 = 4294967295
    for fi10 := 0; fi10 < 11; fi10++ {
        v3 -= (v3 | (v3 % (((v3 & 18446744073709551615)) | 1)))
        fmt.Println(v3, v2)
        v11 := []int{v2, 128, (fi10 | (v2 << ((774) & 63)))}
        _ = fi10
        _ = v11
    }
    var v12 uint32 = 4294967295
    var v13 uint = 890
    var v14 uint8 = 0
    if true {
        var v15 float64 = -2.25
        var v16 int8 = -128
        v13 = v13
        _ = v15
        _ = v16
    }
    switch "" {
    case "foo bar":
        for fi17 := 0; fi17 < 8; fi17++ {
            v18 := v1[("foo bar" + "foo bar")]
            _ = fi17
            _ = v18
        }
    default:
        var v19 float32 = (3.14159 * 1.5)
        _ = v19
    }
    v1["\n\t"] = ("a" + "")
    for fi20 := 0; fi20 < 3; fi20++ {
        var v21 int8 = 100
        switch (("foo bar" + ("a" + ("a" + "hello"))) + "hello") {
        case "foo bar":
            v12 += v12
        case "hello":
            v2 -= v2
        }
        _ = fi20
        _ = v21
    }
    v22 := GenS1{F0: -96}
    v13 = (v13 | 0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v3
}
