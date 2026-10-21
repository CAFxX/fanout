package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 int; F2 uint32 }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint16) string {
    v6 := GenS1{F0: (42 > -1)}
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 uint64 = 273
        var v9 int32 = 849
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 bool = false
    if v10 {
        fmt.Println(p0)
    }
    if v10 {
        v11 := []int{-2147483648, -494, 1, 4294967295}
        v12 := GenS2{F0: 18446744073709551615, F1: -2147483648, F2: 42}
        _ = v11
        _ = v12
    }
    _ = p0
    _ = v10
    _ = v6
    return "a"
}

func main() {
    var v1 uint16 = 32767
    var v2 int8 = 36
    v3 := make(map[int]string)
    v4 := make(map[int]string)
    v5 := []string{("a" + "\n\t"), "a"}
    for fi13 := 0; fi13 < 1; fi13++ {
        v1 = (v1 >> ((2) & 63))
        switch fi13 {
        case 127:
            fmt.Println(v2, v1, fi13)
        }
        _ = fi13
    }
    if (-659 < 42) {
        v2 = v2
    } else {
        if true {
            v14 := GenS2{F0: 1, F1: 128, F2: 2}
            v1 = v1
            _ = v14
        }
        v15 := GenS1{F0: (1000 < 1000)}
        v2 += 42
        _ = v15
    }
    fmt.Println(v1)
    var v16 uint16 = v1
    var v17 int = 9223372036854775807
    v18 := GenS1{F0: (759 >= 2)}
    v19 := GenS3{F0: -2147483648}
    var v20 float32 = 0.1
    switch ("a" + "hello") {
    case "a":
        delete(v4, -992)
    case "":
        var v21 uint8 = 2
        _ = v21
    }
    if ((-(v2)) < (v2 + (v2 % (((^(v2))) | 1)))) {
        var v22 int32 = -1
        _ = v22
    }
    var v23 uint16 = v1
    v1 -= 32767
    v24 := GenS2{F0: 1, F1: v17, F2: 583}
    var v25 int64 = -1
    if (128 != 128) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    } else {
        v17 = -898
        var v26 string = "hello"
        v5 = append(v5, (v26 + v26))
        _ = v26
    }
    v17 = v17
    v16 -= v23
    var v27 uint64 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
