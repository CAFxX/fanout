package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint32 }
type GenS2 struct { F0 int8; F1 uint16; F2 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8; F1 string }
func genFn0(p0 uint) uint {
    var v5 float64 = 1.5
    fmt.Println(p0, v5)
    v6 := GenS1{F0: 2147483647, F1: p0, F2: 1}
    p0 = (p0 - p0)
    var v7 string = "foo bar"
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 65535
}
func genFn1(p0 string, p1 string, p2 int8) string {
    v8 := GenS2{F0: p2, F1: 227, F2: 42}
    v9 := GenS2{F0: p2, F1: 414, F2: 4294967295}
    switch "hello" {
    case "a":
        v10 := []int{4294967295, 32768, 32768, 881}
        _ = v10
    default:
        p0 = (p0 + "")
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 string = p0
        _ = fi11
        _ = v12
    }
    p0 = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return "hello"
}

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{2147483647, 4294967295, 148}
    var v3 int32 = -2147483648
    v4 := []uint32{}
    delete(v1, -971)
    fmt.Println(v3)
    v13 := GenS3{F0: 4294967295}
    var v14 int32 = -2147483648
    var v15 float64 = (0.1 - (0.0 + 3.14159))
    fmt.Println(v3, v15, v3)
    v16 := GenS2{F0: 100, F1: 1, F2: 1}
    v17 := GenS4{F0: 184, F1: "hello"}
    for fi18 := 0; fi18 < 14; fi18++ {
        v14 = v3
        v19 := GenS1{F0: (v14 % (((v3 >> ((4294967295) & 63))) | 1)), F1: 4294967295, F2: 465}
        _ = fi18
        _ = v19
    }
    var v20 int = -2147483648
    var v21 bool = false
    var v22 bool = (v21 || (v14 >= v14))
    var v23 uint64 = 603
    for fi24 := 0; fi24 < 9; fi24++ {
        for fi25 := 0; fi25 < 7; fi25++ {
            fmt.Println(fi24)
            v2 = append(v2, 178)
            _ = fi25
        }
        _ = fi24
    }
    v14 = v3
    _ = len(v4)
    v1[v20] = true
    v26 := GenS2{F0: 81, F1: 360, F2: (v23 >> ((genFn0(0)) & 63))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v4
}
