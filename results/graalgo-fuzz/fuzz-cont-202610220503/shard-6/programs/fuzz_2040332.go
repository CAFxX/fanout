package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 string }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 uint; F1 uint8; F2 int8 }
func genFn0(p0 uint8, p1 int, p2 bool) int16 {
    v6 := []int{}
    p1 += p1
    v7 := GenS1{F0: 1, F1: 101, F2: ""}
    p0 -= 0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -32768
}
func genFn1(p0 int8, p1 uint8) int8 {
    var v8 uint8 = p1
    v9 := GenS2{F0: 0, F1: 2147483647}
    v10 := []int{1, 434}
    var v11 uint32 = 2147483648
    switch "\n\t" {
    case "\n\t":
        for fi12 := 0; fi12 < 3; fi12++ {
            var v13 int32 = 0
            v14 := GenS2{F0: 0, F1: v13}
            _ = fi12
            _ = v13
            _ = v14
        }
    default:
        switch -9223372036854775808 {
        case 32768:
            var v15 float64 = 1.5
            _ = v15
        default:
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
        }
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 127
}

func main() {
    var v1 uint64 = 904
    var v2 uint = 685
    var v3 uint32 = 65535
    v4 := make(map[uint32]int)
    v5 := []string{"\n\t"}
    for fi16 := 0; fi16 < 4; fi16++ {
        v17 := GenS2{F0: 475, F1: -1}
        v3 = (^(v3))
        var v18 int = fi16
        _ = fi16
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 19; fi19++ {
        fmt.Println(v3, v1, v2)
        _ = fi19
    }
    var v20 int64 = 42
    for fi21 := 0; fi21 < 3; fi21++ {
        switch ("" + ("" + "")) {
        case "hello":
            if (4294967295 < v2) {
                var v22 uint = (^(v2))
                _ = v22
            }
        case "\n\t":
            var v23 int8 = genFn1(127, 245)
            _ = v23
        case "":
            v1 = (v1 % (((v1 * (v1 / (((v1 - 1)) | 1)))) | 1))
        default:
            v3 -= 2
        }
        var v24 int32 = -468
        fmt.Println(v24)
        _ = fi21
        _ = v24
    }
    v5 = append(v5, "a")
    v25 := GenS1{F0: 54, F1: uint(v3), F2: "a"}
    switch -2147483648 {
    case 9223372036854775807:
        for fi26 := 0; fi26 < 14; fi26++ {
            v27 := GenS3{F0: 1, F1: 128, F2: -1}
            v20 -= v20
            var v28 bool = true
            _ = fi26
            _ = v27
            _ = v28
        }
    }
    v2 += v2
    var v29 uint16 = 2
    for fi30 := 0; fi30 < 1; fi30++ {
        v20 -= (-(v20))
        var v31 uint8 = 1
        _ = fi30
        _ = v31
    }
    v1 = 18446744073709551615
    v4[v3] = 4294967295
    v5 = append(v5, "\n\t")
    switch "foo bar" {
    case "hello":
        var v32 uint32 = v3
        _ = v32
    case "\n\t":
        v20 = 0
    case "foo bar":
        var v33 int16 = -32768
        _ = v33
    default:
        fmt.Println(v29)
    }
    var v34 int8 = genFn1(-1, 133)
    for fi35 := 0; fi35 < 16; fi35++ {
        _ = len(v4)
        var v36 uint32 = v3
        _ = fi35
        _ = v36
    }
    v37 := v4[v3]
    fmt.Println(v37, v34, v37)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v29
    _ = v3
    _ = v34
    _ = v37
    _ = v4
    _ = v5
}
