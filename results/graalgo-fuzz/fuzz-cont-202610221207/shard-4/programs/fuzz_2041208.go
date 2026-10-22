package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 bool }
type GenS2 struct { F0 uint16 }
func genFn0() int {
    for fi4 := 0; fi4 < 3; fi4++ {
        v5 := []int{}
        if ((1 == 52) && true) {
            v5 = append(v5, 82)
            var v6 uint = 2
            var v7 string = "hello"
            _ = v6
            _ = v7
        } else {
            for fi8 := 0; fi8 < 5; fi8++ {
                v9 := GenS1{F0: 32767, F1: "", F2: false}
                _ = fi8
                _ = v9
            }
            var v10 int64 = 792
            _ = v10
        }
        _ = fi4
        _ = v5
    }
    if (1 <= -480) {
        var v11 int8 = 42
        v11 -= v11
        v12 := GenS1{F0: 65535, F1: "a", F2: (false || true)}
        _ = v11
        _ = v12
    } else {
        var v13 int8 = -1
        switch ("foo bar" + "foo bar") {
        case "foo bar":
            var v14 uint = 65535
            _ = v14
        case "a":
            var v15 string = ""
            _ = v15
        case "hello":
            v16 := GenS2{F0: 32767}
            _ = v16
        default:
            for fi17 := 0; fi17 < 3; fi17++ {
                if false {
                    v18 := GenS2{F0: 255}
                    _ = v18
                } else {
                    fmt.Println(v13)
                }
                var v19 uint32 = 42
                fmt.Println(v13, v13)
                _ = fi17
                _ = v19
            }
        }
        _ = v13
    }
    v20 := []int{2147483647, -294, 2147483647}
    var v21 uint16 = 32767
    _ = v20
    _ = v21
    return 0
}

func main() {
    var v1 int32 = 0
    var v2 uint64 = 812
    v3 := make(map[int]uint32)
    var v22 int8 = 1
    v23 := []int{genFn0(), 42, 1}
    v2 = 18446744073709551615
    var v24 int = 2147483647
    v25 := GenS2{F0: 1}
    delete(v3, v24)
    var v26 uint = 65535
    fmt.Println(v26, v1)
    var v27 uint16 = 32767
    fmt.Println(v2, v2)
    switch (v24 + v24) {
    case 945:
        v3[v24] = 707
    case -612:
        v28 := GenS1{F0: v27, F1: "", F2: false}
        _ = v28
    default:
        fmt.Println(v22)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
