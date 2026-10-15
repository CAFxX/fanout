package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int8; F1 uint16; F2 string }
func genFn0(p0 uint16, p1 string) int {
    for fi6 := 0; fi6 < 15; fi6++ {
        for fi7 := 0; fi7 < 1; fi7++ {
            v8 := GenS1{F0: 0, F1: (p0 >> ((4294967295) & 63)), F2: 1}
            v9 := GenS1{F0: 0, F1: 1, F2: 824}
            v10 := []int{}
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    v11 := GenS1{F0: 42, F1: 65535, F2: 4294967295}
    fmt.Println(p1, p1)
    fmt.Println(p0, p0)
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v11
    return 2147483647
}

func main() {
    var v1 uint64 = 149
    v2 := make(map[int]string)
    v3 := []int{-2147483648}
    v4 := []uint32{2147483648}
    var v5 int64 = -1
    for fi12 := 0; fi12 < 5; fi12++ {
        v1 = 42
        v1 -= 18446744073709551615
        v2[(fi12 % ((-932) | 1))] = (("" + "a") + "foo bar")
        _ = fi12
    }
    if true {
        v13 := GenS2{F0: false, F1: ("" + "foo bar")}
        var v14 float32 = float32(0.1)
        _ = v13
        _ = v14
    } else {
        var v15 uint32 = 2
        for fi16 := 0; fi16 < 11; fi16++ {
            _ = len(v3)
            v15 = v15
            var v17 uint32 = (^(v15))
            _ = fi16
            _ = v17
        }
        v18 := GenS3{F0: -1, F1: 491, F2: ("foo bar" + "foo bar")}
        _ = v15
        _ = v18
    }
    if true {
        if false {
            fmt.Println(v1, v5)
            v1 -= 1
        } else {
            v1 += v1
            var v19 float32 = (1e10 / 1e10)
            _ = v19
        }
        v20 := GenS3{F0: 127, F1: 32767, F2: ""}
        _ = v20
    } else {
        var v21 uint8 = 0
        _ = v21
    }
    var v22 string = ("foo bar" + "hello")
    v1 = 0
    var v23 uint32 = 0
    if (810 != 2147483647) {
        fmt.Println(v5, v23, v23)
    }
    v2[genFn0(65535, v22)] = v22
    v22 = (v22 + v22)
    var v24 int64 = v5
    v24 -= v24
    v25 := GenS2{F0: false, F1: "a"}
    v5 += v24
    if true {
        if (65535 == 943) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
