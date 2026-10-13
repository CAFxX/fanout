package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 int16 }
func genFn0() uint64 {
    var v4 uint32 = 2147483648
    fmt.Println(v4)
    if false {
        fmt.Println(v4)
    }
    fmt.Println(v4)
    _ = v4
    return 18446744073709551615
}
func genFn1(p0 uint32) bool {
    v5 := []int{42, 32768, 127}
    fmt.Println(p0)
    v6 := GenS1{F0: 1, F1: "\n\t"}
    _ = p0
    _ = v5
    _ = v6
    return (1000 <= 1)
}

func main() {
    var v1 int = 955
    v2 := make(map[string]string)
    var v3 uint = 2
    if true {
        for fi7 := 0; fi7 < 1; fi7++ {
            v8 := GenS2{F0: 625, F1: (1.5 > 1.5)}
            _ = fi7
            _ = v8
        }
    } else {
        v1 += v1
    }
    v1 += v1
    fmt.Println(v1, v1)
    v1 += -2147483648
    v2["foo bar"] = "\n\t"
    var v9 uint8 = 1
    fmt.Println(v3)
    switch (v1 % ((4294967295) | 1)) {
    case 0:
        v9 += v9
    case 42:
        delete(v2, ("" + ("\n\t" + "\n\t")))
    default:
        if true {
            v10 := v2["hello"]
            _ = v10
        } else {
            var v11 int64 = 1
            _ = v11
        }
    }
    switch 1 {
    case 32768:
        v12 := v2["hello"]
        _ = v12
    }
    if false {
        v13 := []int{}
        _ = v13
    } else {
        delete(v2, ("foo bar" + "foo bar"))
        if false {
            var v14 float64 = 0.1
            v15 := []int{1, -2147483648, v1, v1}
            v16 := GenS3{F0: 920}
            _ = v14
            _ = v15
            _ = v16
        } else {
            v3 = 42
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v9
}
