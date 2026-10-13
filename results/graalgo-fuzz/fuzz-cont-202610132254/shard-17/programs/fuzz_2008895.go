package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 uint16 }
func genFn0(p0 int16, p1 int16, p2 string) int32 {
    var v4 int32 = 2147483647
    fmt.Println(p1, p0, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (v4 * v4)
}
func genFn1(p0 int64) uint64 {
    fmt.Println(p0)
    p0 -= (p0 | p0)
    if (false || (9223372036854775807 <= p0)) {
        for fi5 := 0; fi5 < 2; fi5++ {
            for fi6 := 0; fi6 < 20; fi6++ {
                var v7 uint = 839
                _ = fi6
                _ = v7
            }
            if true {
                fmt.Println(fi5, fi5)
            }
            _ = fi5
        }
    } else {
        if (!true) {
            switch "" {
            case "":
                var v8 uint = 18446744073709551615
                _ = v8
            case "foo bar":
                p0 -= (^(p0))
            }
            var v9 uint16 = 255
            _ = v9
        }
        v10 := []int{4294967295, 4294967295}
        fmt.Println(p0)
        _ = v10
    }
    p0 = (p0 + p0)
    fmt.Println(p0)
    _ = p0
    return 4294967295
}

func main() {
    var v1 float64 = ((((-2.25 + 3.14159) + 0.1) * 3.14159) - (float64(((((0.1 / 3.14159) + 1e10) - -2.25) + (3.14159 - (0.0 + float32(1.5))))) / -2.25))
    var v2 float64 = float64(0.1)
    v3 := make(map[uint32]uint32)
    var v11 string = "hello"
    var v12 string = ("foo bar" + "")
    v1 += (v2 + (0.0 + 1.5))
    if false {
        if (!((2 < 42) || true)) {
            var v13 uint32 = 37
            _ = v13
        }
    }
    v14 := GenS1{F0: genFn1(-241)}
    v12 = ((v12 + ("a" + "\n\t")) + v11)
    switch -2147483648 {
    case 0:
        fmt.Println(v2, v11, v1)
    case 127:
        v15 := GenS1{F0: 0}
        _ = v15
    case 128:
        var v16 int16 = -1
        _ = v16
    default:
        v17 := []int{345}
        _ = v17
    }
    if (2147483648 == 746) {
        var v18 string = (v12 + ("hello" + v11))
        v19 := []int{826, 990, -9223372036854775808}
        v12 = (v11 + v12)
        _ = v18
        _ = v19
    }
    if (!(!(42 <= 18446744073709551615))) {
        var v20 uint64 = 299
        v21 := []int{4294967295}
        v22 := GenS1{F0: (^(v20))}
        _ = v20
        _ = v21
        _ = v22
    }
    v23 := []int{-43, -1}
    var v24 int32 = 42
    v25 := GenS1{F0: 183}
    v26 := GenS2{F0: (v24 - v24), F1: 255}
    _ = len(v23)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
