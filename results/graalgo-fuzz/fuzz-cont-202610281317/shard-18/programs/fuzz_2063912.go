package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 string }
func genFn0() int8 {
    if ((18446744073709551615 == 73) && true) {
        v7 := GenS1{F0: 255, F1: false}
        _ = v7
    } else {
        if true {
            var v8 uint8 = 0
            _ = v8
        } else {
            var v9 string = ("" + "foo bar")
            var v10 int = 0
            v11 := []int{}
            _ = v10
            _ = v11
            _ = v9
        }
    }
    var v12 int = 748
    v13 := GenS2{F0: true, F1: 0}
    _ = v12
    _ = v13
    return 42
}

func main() {
    var v1 int8 = 0
    v2 := make(map[int]uint32)
    var v3 uint64 = 18446744073709551615
    v4 := []uint32{138, 64}
    v5 := make(map[int]string)
    var v6 string = "foo bar"
    v2[0] = 2147483648
    fmt.Println(v1, v1)
    if ((((681 <= 18446744073709551615) || true) && (2 < 255)) && true) {
        v14 := GenS3{F0: (v6 + "hello")}
        var v15 uint64 = v3
        v3 = v3
        _ = v14
        _ = v15
    }
    var v16 int8 = 127
    switch 2147483647 {
    case -479:
        fmt.Println(v3, v1, v6)
    case 127:
        var v17 uint = 65535
        _ = v17
    }
    v18 := GenS2{F0: ((-(v16)) <= v1), F1: 42}
    v19 := GenS1{F0: 4294967295, F1: true}
    v3 = v3
    v20 := v5[32768]
    v21 := GenS3{F0: v20}
    fmt.Println(v20, v3, v1)
    for fi22 := 0; fi22 < 6; fi22++ {
        _ = len(v4)
        _ = fi22
    }
    _ = len(v2)
    v16 -= (-(v16))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
