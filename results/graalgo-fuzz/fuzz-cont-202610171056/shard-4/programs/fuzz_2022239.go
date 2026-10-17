package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 uint }
func genFn0(p0 int64) uint16 {
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 int64 = (p0 + p0)
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: -2147483648}
    var v10 bool = (!false)
    if v10 {
        var v11 int32 = 52
        _ = v11
    } else {
        if ((("" + "\n\t") <= (("foo bar" + "\n\t") + ("hello" + ("\n\t" + "\n\t")))) || true) {
            for fi12 := 0; fi12 < 15; fi12++ {
                p0 = p0
                _ = fi12
            }
        }
    }
    _ = p0
    _ = v10
    _ = v9
    return 1
}
func genFn1(p0 uint, p1 int8) uint8 {
    p0 += p0
    p1 += 42
    v13 := GenS2{F0: true}
    _ = p0
    _ = p1
    _ = v13
    return 110
}

func main() {
    v1 := []int{2147483647, 127}
    v2 := []int{-9223372036854775808, 1, 128, 4294967295}
    v3 := make(map[int]string)
    v4 := make(map[uint32]string)
    v5 := make(map[int]uint32)
    v6 := []string{}
    _ = len(v4)
    var v14 int16 = 42
    v15 := GenS3{F0: 428, F1: 255}
    _ = len(v6)
    for fi16 := 0; fi16 < 14; fi16++ {
        _ = len(v6)
        _ = fi16
    }
    fmt.Println(v14)
    switch ("hello" + ("foo bar" + "foo bar")) {
    case "foo bar":
        if true {
            var v17 int32 = 50
            v17 = (v17 | v17)
            _ = v17
        }
    default:
        var v18 string = ("\n\t" + "\n\t")
        _ = v18
    }
    if (!false) {
        for fi19 := 0; fi19 < 16; fi19++ {
            var v20 int8 = 68
            fmt.Println(v14, v14, v20)
            _ = fi19
            _ = v20
        }
        v14 = v14
    }
    v14 = 1
    var v21 int8 = -63
    var v22 int32 = 905
    v23 := GenS1{F0: -2147483648}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
