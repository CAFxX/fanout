package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 string }
type GenS2 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint64; F1 uint; F2 int64 }
func genFn0(p0 int8) int64 {
    if (2 >= 4294967295) {
        v7 := []int{42, 145, 9223372036854775807}
        _ = v7
    } else {
        var v8 float32 = 1.5
        _ = v8
    }
    p0 = (p0 + -128)
    p0 -= 106
    p0 = p0
    v9 := GenS1{F0: 1, F1: -2147483648, F2: "a"}
    _ = p0
    _ = v9
    return -217
}

func main() {
    var v1 int64 = 1
    v2 := make(map[int]uint32)
    v3 := make(map[string]bool)
    v4 := []string{(("a" + "\n\t") + "foo bar"), "hello", (("a" + "") + ("hello" + "a"))}
    var v5 int = 208
    v6 := []uint32{936, 604, 44}
    fmt.Println(v1, v1)
    v5 += (v5 << ((903) & 63))
    v10 := GenS1{F0: 2, F1: 1, F2: ""}
    for fi11 := 0; fi11 < 8; fi11++ {
        var v12 int32 = 0
        _ = fi11
        _ = v12
    }
    fmt.Println(v1, v5)
    var v13 bool = false
    v5 -= v5
    v14 := GenS2{F0: 65535, F1: -1, F2: 2}
    switch v5 {
    case 2147483647:
        for fi15 := 0; fi15 < 1; fi15++ {
            var v16 float32 = (float32(float64(float32(-2.25))) + 1e10)
            _ = fi15
            _ = v16
        }
    case -1:
        var v17 uint16 = 42
        _ = v17
    case 128:
        var v18 float64 = 1e10
        _ = v18
    }
    switch -512 {
    case 210:
        v19 := GenS3{F0: 1, F1: 4294967295, F2: -698}
        _ = v19
    case -765:
        if ((0 != 0) || ((!((99 >= 87) && (!(611 >= 20)))) && false)) {
            delete(v3, "\n\t")
        }
    case -957:
        var v20 int8 = 0
        _ = v20
    default:
        delete(v2, v5)
    }
    v21 := GenS3{F0: 246, F1: 1, F2: v1}
    var v22 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
